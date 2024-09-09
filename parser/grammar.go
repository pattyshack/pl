// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
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
	switch id {
	case _State4:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State5:
		return []SymbolId{}
	case _State6:
		return []SymbolId{}
	case _State7:
		return []SymbolId{}
	case _State8:
		return []SymbolId{AddToken, SubToken, MulToken}
	case _State10:
		return []SymbolId{IdentifierToken, LparenToken}
	case _State14:
		return []SymbolId{IdentifierToken}
	case _State24:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken}
	case _State31:
		return []SymbolId{ColonToken}
	case _State36:
		return []SymbolId{LparenToken}
	case _State39:
		return []SymbolId{StringLiteralToken, IdentifierToken, UnderscoreToken, LparenToken, DotToken}
	case _State42:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State50:
		return []SymbolId{LparenToken}
	case _State54:
		return []SymbolId{LessToken}
	case _State61:
		return []SymbolId{AssignToken}
	case _State84:
		return []SymbolId{LparenToken}
	case _State93:
		return []SymbolId{IfToken, SwitchToken, ForToken, DoToken, LbraceToken}
	case _State113:
		return []SymbolId{LparenToken}
	case _State115:
		return []SymbolId{LparenToken}
	case _State121:
		return []SymbolId{LparenToken}
	case _State134:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State138:
		return []SymbolId{IdentifierToken, UnderscoreToken}
	case _State141:
		return []SymbolId{RbraceToken}
	case _State147:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken, EllipsisToken}
	case _State151:
		return []SymbolId{IdentifierToken}
	case _State155:
		return []SymbolId{CommaToken, ColonToken}
	case _State156:
		return []SymbolId{LbraceToken}
	case _State161:
		return []SymbolId{StringLiteralToken}
	case _State162:
		return []SymbolId{StringLiteralToken}
	case _State163:
		return []SymbolId{StringLiteralToken, IdentifierToken, UnderscoreToken, DotToken}
	case _State165:
		return []SymbolId{StringLiteralToken}
	case _State167:
		return []SymbolId{RbracketToken, CommaToken, ColonToken, AddToken, SubToken, MulToken}
	case _State172:
		return []SymbolId{RparenToken}
	case _State177:
		return []SymbolId{IdentifierToken}
	case _State188:
		return []SymbolId{IdentifierToken}
	case _State197:
		return []SymbolId{LparenToken}
	case _State227:
		return []SymbolId{LbraceToken}
	case _State236:
		return []SymbolId{IdentifierToken, UnderscoreToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State238:
		return []SymbolId{IdentifierToken}
	case _State240:
		return []SymbolId{IdentifierToken, LparenToken}
	case _State244:
		return []SymbolId{RparenToken}
	case _State245:
		return []SymbolId{RparenToken}
	case _State256:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State259:
		return []SymbolId{LparenToken}
	case _State262:
		return []SymbolId{RparenToken}
	case _State267:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State268:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State274:
		return []SymbolId{RparenToken, CommaToken}
	case _State279:
		return []SymbolId{IdentifierToken}
	case _State285:
		return []SymbolId{RparenToken}
	case _State290:
		return []SymbolId{RparenToken}
	case _State293:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State294:
		return []SymbolId{IntegerLiteralToken}
	case _State303:
		return []SymbolId{RparenToken}
	case _State306:
		return []SymbolId{GreaterToken}
	case _State308:
		return []SymbolId{RbracketToken}
	case _State311:
		return []SymbolId{RbracketToken}
	case _State319:
		return []SymbolId{RparenToken}
	case _State323:
		return []SymbolId{InToken, AssignToken}
	case _State324:
		return []SymbolId{SemicolonToken}
	case _State327:
		return []SymbolId{LbraceToken}
	case _State329:
		return []SymbolId{LbraceToken}
	case _State331:
		return []SymbolId{RparenToken}
	case _State332:
		return []SymbolId{NewlinesToken, OrToken}
	case _State338:
		return []SymbolId{RparenToken}
	case _State343:
		return []SymbolId{LparenToken}
	case _State347:
		return []SymbolId{IdentifierToken, UnderscoreToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State351:
		return []SymbolId{IdentifierToken, UnderscoreToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State353:
		return []SymbolId{DefaultToken}
	case _State355:
		return []SymbolId{RparenToken}
	case _State360:
		return []SymbolId{RbracketToken}
	case _State367:
		return []SymbolId{IdentifierToken}
	case _State372:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken}
	case _State373:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken, EllipsisToken}
	case _State376:
		return []SymbolId{LparenToken}
	case _State384:
		return []SymbolId{RbracketToken, AddToken, SubToken, MulToken}
	case _State385:
		return []SymbolId{RbracketToken}
	case _State393:
		return []SymbolId{StringLiteralToken}
	case _State397:
		return []SymbolId{RparenToken}
	case _State403:
		return []SymbolId{LbraceToken}
	case _State404:
		return []SymbolId{CommaToken, AssignToken}
	case _State408:
		return []SymbolId{IdentifierToken, UnderscoreToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State409:
		return []SymbolId{IdentifierToken, UnderscoreToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State411:
		return []SymbolId{IdentifierToken, UnderscoreToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State426:
		return []SymbolId{RparenToken}
	case _State429:
		return []SymbolId{LparenToken}
	case _State430:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State434:
		return []SymbolId{LbraceToken}
	case _State448:
		return []SymbolId{DoToken}
	case _State449:
		return []SymbolId{SemicolonToken}
	case _State460:
		return []SymbolId{RparenToken}
	case _State466:
		return []SymbolId{LbraceToken}
	case _State470:
		return []SymbolId{IfToken}
	case _State473:
		return []SymbolId{LbraceToken}
	case _State474:
		return []SymbolId{RparenToken}
	case _State476:
		return []SymbolId{DoToken}
	case _State480:
		return []SymbolId{LbraceToken}
	case _State481:
		return []SymbolId{LbraceToken}
	}

	return nil
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
		} else if action.ActionType == _ShiftAndReduceAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}

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
	_ShiftAction          = _ActionType(0)
	_ReduceAction         = _ActionType(1)
	_ShiftAndReduceAction = _ActionType(2)
	_AcceptAction         = _ActionType(3)
)

func (i _ActionType) String() string {
	switch i {
	case _ShiftAction:
		return "shift"
	case _ReduceAction:
		return "reduce"
	case _ShiftAndReduceAction:
		return "shift-and-reduce"
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
	_State10  = _StateId(10)
	_State11  = _StateId(11)
	_State13  = _StateId(13)
	_State14  = _StateId(14)
	_State20  = _StateId(20)
	_State23  = _StateId(23)
	_State24  = _StateId(24)
	_State29  = _StateId(29)
	_State31  = _StateId(31)
	_State36  = _StateId(36)
	_State37  = _StateId(37)
	_State39  = _StateId(39)
	_State42  = _StateId(42)
	_State43  = _StateId(43)
	_State50  = _StateId(50)
	_State54  = _StateId(54)
	_State56  = _StateId(56)
	_State57  = _StateId(57)
	_State58  = _StateId(58)
	_State61  = _StateId(61)
	_State72  = _StateId(72)
	_State74  = _StateId(74)
	_State78  = _StateId(78)
	_State84  = _StateId(84)
	_State87  = _StateId(87)
	_State91  = _StateId(91)
	_State93  = _StateId(93)
	_State94  = _StateId(94)
	_State99  = _StateId(99)
	_State101 = _StateId(101)
	_State108 = _StateId(108)
	_State113 = _StateId(113)
	_State115 = _StateId(115)
	_State116 = _StateId(116)
	_State117 = _StateId(117)
	_State121 = _StateId(121)
	_State134 = _StateId(134)
	_State137 = _StateId(137)
	_State138 = _StateId(138)
	_State139 = _StateId(139)
	_State141 = _StateId(141)
	_State143 = _StateId(143)
	_State144 = _StateId(144)
	_State145 = _StateId(145)
	_State147 = _StateId(147)
	_State150 = _StateId(150)
	_State151 = _StateId(151)
	_State152 = _StateId(152)
	_State155 = _StateId(155)
	_State156 = _StateId(156)
	_State158 = _StateId(158)
	_State159 = _StateId(159)
	_State161 = _StateId(161)
	_State162 = _StateId(162)
	_State163 = _StateId(163)
	_State165 = _StateId(165)
	_State167 = _StateId(167)
	_State168 = _StateId(168)
	_State170 = _StateId(170)
	_State172 = _StateId(172)
	_State173 = _StateId(173)
	_State174 = _StateId(174)
	_State175 = _StateId(175)
	_State176 = _StateId(176)
	_State177 = _StateId(177)
	_State187 = _StateId(187)
	_State188 = _StateId(188)
	_State190 = _StateId(190)
	_State196 = _StateId(196)
	_State197 = _StateId(197)
	_State204 = _StateId(204)
	_State205 = _StateId(205)
	_State206 = _StateId(206)
	_State207 = _StateId(207)
	_State208 = _StateId(208)
	_State215 = _StateId(215)
	_State216 = _StateId(216)
	_State217 = _StateId(217)
	_State218 = _StateId(218)
	_State219 = _StateId(219)
	_State226 = _StateId(226)
	_State227 = _StateId(227)
	_State228 = _StateId(228)
	_State229 = _StateId(229)
	_State230 = _StateId(230)
	_State234 = _StateId(234)
	_State236 = _StateId(236)
	_State237 = _StateId(237)
	_State238 = _StateId(238)
	_State240 = _StateId(240)
	_State241 = _StateId(241)
	_State242 = _StateId(242)
	_State243 = _StateId(243)
	_State244 = _StateId(244)
	_State245 = _StateId(245)
	_State247 = _StateId(247)
	_State248 = _StateId(248)
	_State249 = _StateId(249)
	_State251 = _StateId(251)
	_State256 = _StateId(256)
	_State257 = _StateId(257)
	_State258 = _StateId(258)
	_State259 = _StateId(259)
	_State260 = _StateId(260)
	_State261 = _StateId(261)
	_State262 = _StateId(262)
	_State264 = _StateId(264)
	_State265 = _StateId(265)
	_State267 = _StateId(267)
	_State268 = _StateId(268)
	_State272 = _StateId(272)
	_State274 = _StateId(274)
	_State277 = _StateId(277)
	_State278 = _StateId(278)
	_State279 = _StateId(279)
	_State280 = _StateId(280)
	_State281 = _StateId(281)
	_State285 = _StateId(285)
	_State286 = _StateId(286)
	_State290 = _StateId(290)
	_State291 = _StateId(291)
	_State293 = _StateId(293)
	_State294 = _StateId(294)
	_State297 = _StateId(297)
	_State299 = _StateId(299)
	_State300 = _StateId(300)
	_State302 = _StateId(302)
	_State303 = _StateId(303)
	_State305 = _StateId(305)
	_State306 = _StateId(306)
	_State307 = _StateId(307)
	_State308 = _StateId(308)
	_State309 = _StateId(309)
	_State311 = _StateId(311)
	_State313 = _StateId(313)
	_State314 = _StateId(314)
	_State315 = _StateId(315)
	_State317 = _StateId(317)
	_State319 = _StateId(319)
	_State320 = _StateId(320)
	_State322 = _StateId(322)
	_State323 = _StateId(323)
	_State324 = _StateId(324)
	_State325 = _StateId(325)
	_State326 = _StateId(326)
	_State327 = _StateId(327)
	_State329 = _StateId(329)
	_State330 = _StateId(330)
	_State331 = _StateId(331)
	_State332 = _StateId(332)
	_State333 = _StateId(333)
	_State334 = _StateId(334)
	_State335 = _StateId(335)
	_State336 = _StateId(336)
	_State338 = _StateId(338)
	_State339 = _StateId(339)
	_State341 = _StateId(341)
	_State342 = _StateId(342)
	_State343 = _StateId(343)
	_State344 = _StateId(344)
	_State345 = _StateId(345)
	_State346 = _StateId(346)
	_State347 = _StateId(347)
	_State351 = _StateId(351)
	_State352 = _StateId(352)
	_State353 = _StateId(353)
	_State355 = _StateId(355)
	_State358 = _StateId(358)
	_State360 = _StateId(360)
	_State361 = _StateId(361)
	_State362 = _StateId(362)
	_State363 = _StateId(363)
	_State364 = _StateId(364)
	_State365 = _StateId(365)
	_State366 = _StateId(366)
	_State367 = _StateId(367)
	_State370 = _StateId(370)
	_State371 = _StateId(371)
	_State372 = _StateId(372)
	_State373 = _StateId(373)
	_State376 = _StateId(376)
	_State379 = _StateId(379)
	_State380 = _StateId(380)
	_State382 = _StateId(382)
	_State383 = _StateId(383)
	_State384 = _StateId(384)
	_State385 = _StateId(385)
	_State391 = _StateId(391)
	_State392 = _StateId(392)
	_State393 = _StateId(393)
	_State394 = _StateId(394)
	_State397 = _StateId(397)
	_State399 = _StateId(399)
	_State400 = _StateId(400)
	_State401 = _StateId(401)
	_State402 = _StateId(402)
	_State403 = _StateId(403)
	_State404 = _StateId(404)
	_State405 = _StateId(405)
	_State408 = _StateId(408)
	_State409 = _StateId(409)
	_State410 = _StateId(410)
	_State411 = _StateId(411)
	_State412 = _StateId(412)
	_State413 = _StateId(413)
	_State414 = _StateId(414)
	_State416 = _StateId(416)
	_State423 = _StateId(423)
	_State425 = _StateId(425)
	_State426 = _StateId(426)
	_State427 = _StateId(427)
	_State428 = _StateId(428)
	_State429 = _StateId(429)
	_State430 = _StateId(430)
	_State434 = _StateId(434)
	_State444 = _StateId(444)
	_State448 = _StateId(448)
	_State449 = _StateId(449)
	_State452 = _StateId(452)
	_State453 = _StateId(453)
	_State460 = _StateId(460)
	_State462 = _StateId(462)
	_State463 = _StateId(463)
	_State464 = _StateId(464)
	_State466 = _StateId(466)
	_State467 = _StateId(467)
	_State470 = _StateId(470)
	_State472 = _StateId(472)
	_State473 = _StateId(473)
	_State474 = _StateId(474)
	_State476 = _StateId(476)
	_State479 = _StateId(479)
	_State480 = _StateId(480)
	_State481 = _StateId(481)
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

type _ActionTableType struct{}

func (_ActionTableType) Get(
	stateId _StateId,
	symbolId SymbolId,
) (
	_Action,
	bool,
) {
	switch stateId {
	case _State1:
		switch symbolId {
		case PackageToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case SourceType:
			return _Action{_ShiftAction, _State5, 0}, true
		case ProperDefinitionsType:
			return _Action{_ShiftAction, _State20, 0}, true
		case VarDeclPatternType:
			return _Action{_ShiftAction, _State23, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentGroupsToDefinition}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case DefinitionsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSource}, true
		case DefinitionType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDefinitionToProperDefinitions}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementBlockToDefinition}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToDefinition}, true
		case NamedFuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedFuncDefToDefinition}, true
		case PackageDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePackageDefToDefinition}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToDefinitions}, true
		}
	case _State2:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State29, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State31, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State39, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case StatementType:
			return _Action{_ShiftAction, _State6, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State72, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State87, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State61, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State101, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State56, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpType}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpType}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpType}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFallthroughStatement}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToCallbackOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToCallbackOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case SimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSimpleStatementToStatement}, true
		case ExprOrImproperStructStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}, true
		case CallbackOpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}, true
		case JumpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStatementToSimpleStatement}, true
		case FallthroughStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}, true
		case AssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStatementToSimpleStatement}, true
		case UnaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}, true
		case BinaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}, true
		case ImportStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStatementToStatement}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State3:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State7, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State4:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State8, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State5:
		switch symbolId {
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
		}
	case _State6:
		switch symbolId {
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
		}
	case _State7:
		switch symbolId {
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
		}
	case _State8:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
		}
	case _State10:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State137, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State138, 0}, true
		}
	case _State11:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State29, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State31, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State39, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ProperStatementsType:
			return _Action{_ShiftAction, _State139, 0}, true
		case StatementsType:
			return _Action{_ShiftAction, _State141, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State72, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State87, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State61, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State101, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State56, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpType}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpType}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpType}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFallthroughStatement}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToCallbackOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToCallbackOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case SimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSimpleStatementToStatement}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToProperStatements}, true
		case ExprOrImproperStructStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}, true
		case CallbackOpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}, true
		case JumpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStatementToSimpleStatement}, true
		case FallthroughStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}, true
		case AssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStatementToSimpleStatement}, true
		case UnaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}, true
		case BinaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}, true
		case ImportStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStatementToStatement}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceNilToStatements}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State13:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWithSpecToPackageDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}, true
		}
	case _State14:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State143, 0}, true
		}
	case _State20:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State144, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperDefinitionsToDefinitions}, true
		}
	case _State23:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State145, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}, true
		}
	case _State24:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State147, 0}, true
		case VarPatternType:
			return _Action{_ShiftAction, _State150, 0}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToVarPattern}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToVarPattern}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTuplePatternToVarPattern}, true
		}
	case _State29:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case VarToken:
			return _Action{_ShiftAction, _State152, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State151, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case CasePatternsType:
			return _Action{_ShiftAction, _State155, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case AssignPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignPatternToCasePattern}, true
		case CasePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCasePatternToCasePatterns}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignPattern}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State31:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State158, 0}, true
		}
	case _State36:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State159, 0}, true
		}
	case _State37:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State39:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State165, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State163, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSingleToImportStatement}, true
		}
	case _State42:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State167, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State43:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State170, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State174, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State175, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State172, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State173, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArgumentToProperArguments}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceNilToArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State50:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State176, 0}, true
		}
	case _State54:
		switch symbolId {
		case LessToken:
			return _Action{_ShiftAction, _State177, 0}, true
		}
	case _State56:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State190, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State188, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case BinaryOpAssignType:
			return _Action{_ShiftAction, _State196, 0}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAction, _State197, 0}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPostfixUnaryOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPostfixUnaryOp}, true
		case AddAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}, true
		case SubAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}, true
		case MulAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}, true
		case DivAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}, true
		case ModAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceModAssignToBinaryOpAssign}, true
		case AddOneAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}, true
		case SubOneAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}, true
		case BitNegAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}, true
		case BitAndAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}, true
		case BitOrAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}, true
		case BitXorAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}, true
		case BitLshiftAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}, true
		case BitRshiftAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}, true
		case UnaryOpAssignType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnaryOpAssignStatement}, true
		case PostfixUnaryOpType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToPostfixUnaryExpr}, true
		case LparenToken:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}, true
		}
	case _State57:
		switch symbolId {
		case AddOpType:
			return _Action{_ShiftAction, _State204, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToAddOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToAddOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToAddOp}, true
		case BitOrToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitOrToAddOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}, true
		}
	case _State58:
		switch symbolId {
		case AndToken:
			return _Action{_ShiftAction, _State205, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}, true
		}
	case _State61:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State206, 0}, true
		}
	case _State72:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case CallExprType:
			return _Action{_ShiftAction, _State208, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State207, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State74:
		switch symbolId {
		case CmpOpType:
			return _Action{_ShiftAction, _State215, 0}, true
		case EqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEqualToCmpOp}, true
		case NotEqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotEqualToCmpOp}, true
		case LessToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLessToCmpOp}, true
		case LessOrEqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLessOrEqualToCmpOp}, true
		case GreaterToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGreaterToCmpOp}, true
		case GreaterOrEqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}, true
		}
	case _State78:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State216, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToExprOrImproperStructStatement}, true
		}
	case _State84:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State217, 0}, true
		}
	case _State87:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case JumpLabelToken:
			return _Action{_ShiftAction, _State218, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State219, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case _EndMarker:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}, true
		case NewlinesToken:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State91:
		switch symbolId {
		case MulOpType:
			return _Action{_ShiftAction, _State226, 0}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToMulOp}, true
		case DivToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDivToMulOp}, true
		case ModToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceModToMulOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToMulOp}, true
		case BitLshiftToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitLshiftToMulOp}, true
		case BitRshiftToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitRshiftToMulOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}, true
		}
	case _State93:
		switch symbolId {
		case IfToken:
			return _Action{_ShiftAction, _State229, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State230, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State228, 0}, true
		case DoToken:
			return _Action{_ShiftAction, _State227, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBlockExpr}, true
		case IfExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIfExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToLoopExpr}, true
		}
	case _State94:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State234, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}, true
		}
	case _State99:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToPrefixUnaryExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State101:
		switch symbolId {
		case AssignToken:
			return _Action{_ReduceAction, 0, _ReduceToAssignPattern}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceSequenceExprToExpr}, true
		}
	case _State108:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State190, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State188, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAction, _State197, 0}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPostfixUnaryOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPostfixUnaryOp}, true
		case PostfixUnaryOpType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToPostfixUnaryExpr}, true
		case LparenToken:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}, true
		}
	case _State113:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State236, 0}, true
		}
	case _State115:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State237, 0}, true
		}
	case _State116:
		switch symbolId {
		case DotToken:
			return _Action{_ShiftAction, _State238, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLocalToNamedTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true
		}
	case _State117:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case FieldDefType:
			return _Action{_ShiftAction, _State243, 0}, true
		case ProperImplicitFieldDefsType:
			return _Action{_ShiftAction, _State248, 0}, true
		case ImplicitFieldDefsType:
			return _Action{_ShiftAction, _State245, 0}, true
		case ProperImplicitEnumValueDefsType:
			return _Action{_ShiftAction, _State247, 0}, true
		case ImplicitEnumValueDefsType:
			return _Action{_ShiftAction, _State244, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToImplicitFieldDefs}, true
		}
	case _State121:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State251, 0}, true
		}
	case _State134:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToPrefixUnaryTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State137:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State258, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State257, 0}, true
		case GenericParameterDefsType:
			return _Action{_ShiftAction, _State259, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefs}, true
		}
	case _State138:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State260, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State261, 0}, true
		case ParameterDefType:
			return _Action{_ShiftAction, _State262, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		}
	case _State139:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State264, 0}, true
		case SemicolonToken:
			return _Action{_ShiftAction, _State265, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperStatementsToStatements}, true
		}
	case _State141:
		switch symbolId {
		case RbraceToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToStatementBlock}, true
		}
	case _State143:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State258, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State267, 0}, true
		case GenericParameterDefsType:
			return _Action{_ShiftAction, _State268, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefs}, true
		}
	case _State144:
		switch symbolId {
		case PackageToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case VarDeclPatternType:
			return _Action{_ShiftAction, _State23, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentGroupsToDefinition}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case DefinitionType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperDefinitions}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementBlockToDefinition}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToDefinition}, true
		case NamedFuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedFuncDefToDefinition}, true
		case PackageDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePackageDefToDefinition}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToDefinitions}, true
		}
	case _State145:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State147:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State272, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State147, 0}, true
		case FieldVarPatternsType:
			return _Action{_ShiftAction, _State274, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToVarPattern}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEllipsisToFieldVarPattern}, true
		case VarPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePositionalToFieldVarPattern}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTuplePatternToVarPattern}, true
		case FieldVarPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}, true
		}
	case _State150:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State277, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case OptionalTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToVarDeclPattern}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalTypeExpr}, true
		}
	case _State151:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State278, 0}, true
		}
	case _State152:
		switch symbolId {
		case DotToken:
			return _Action{_ShiftAction, _State279, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceVarToVarOrLet}, true
		}
	case _State155:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State281, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State280, 0}, true
		}
	case _State156:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBlockExpr}, true
		}
	case _State158:
		switch symbolId {
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State72, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State87, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State61, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State101, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State56, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpType}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpType}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpType}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFallthroughStatement}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToCallbackOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToCallbackOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case SimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSimpleStatementToOptionalSimpleStatement}, true
		case OptionalSimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDefaultBranchStatementToStatement}, true
		case ExprOrImproperStructStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}, true
		case CallbackOpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}, true
		case JumpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStatementToSimpleStatement}, true
		case FallthroughStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}, true
		case AssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStatementToSimpleStatement}, true
		case UnaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}, true
		case BinaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case _EndMarker:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true
		case NewlinesToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State159:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State260, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State261, 0}, true
		case ProperParameterDefsType:
			return _Action{_ShiftAction, _State286, 0}, true
		case ParameterDefsType:
			return _Action{_ShiftAction, _State285, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDefToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDefs}, true
		}
	case _State161:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportToLocalToImportClause}, true
		}
	case _State162:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAliasToImportClause}, true
		}
	case _State163:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State165, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case ProperImportClausesType:
			return _Action{_ShiftAction, _State291, 0}, true
		case ImportClausesType:
			return _Action{_ShiftAction, _State290, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportClauseToProperImportClauses}, true
		}
	case _State165:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnusableImportToImportClause}, true
		}
	case _State167:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State294, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State293, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSliceTypeExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State168:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnitExprPairToColonExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}, true
		case RbracketToken:
			return _Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}, true
		case CommaToken:
			return _Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}, true
		case ColonToken:
			return _Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State170:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State297, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		}
	case _State172:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitStructExpr}, true
		}
	case _State173:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State299, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceColonExprToArgument}, true
		}
	case _State174:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State300, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarargAssignmentToArgument}, true

		default:
			return _Action{_ReduceAction, 0, _ReducePositionalToArgument}, true
		}
	case _State175:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State302, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}, true
		}
	case _State176:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case ProperExplicitFieldDefsType:
			return _Action{_ShiftAction, _State305, 0}, true
		case ExplicitFieldDefsType:
			return _Action{_ShiftAction, _State303, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToProperExplicitFieldDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToExplicitFieldDefs}, true
		}
	case _State177:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State306, 0}, true
		}
	case _State187:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State309, 0}, true
		case ProperTypeArgumentsType:
			return _Action{_ShiftAction, _State307, 0}, true
		case TypeArgumentsType:
			return _Action{_ShiftAction, _State308, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToTypeArguments}, true
		}
	case _State188:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAccessExpr}, true
		}
	case _State190:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State170, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State174, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case ArgumentType:
			return _Action{_ShiftAction, _State311, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State173, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State196:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBinaryOpAssignStatement}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State197:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State313, 0}, true
		}
	case _State204:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State314, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State205:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State315, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State206:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignStatement}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State207:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State190, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State188, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAction, _State197, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true
		}
	case _State208:
		switch symbolId {
		case _EndMarker:
			return _Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}, true
		case NewlinesToken:
			return _Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		}
	case _State215:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State317, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State216:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToExprs}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State217:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State170, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State174, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State175, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State319, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State173, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArgumentToProperArguments}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceNilToArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State218:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State320, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case _EndMarker:
			return _Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}, true
		case NewlinesToken:
			return _Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State219:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State216, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}, true
		}
	case _State226:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBinaryMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State227:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAction, _State322, 0}, true
		}
	case _State228:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State323, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State325, 0}, true
		case ForAssignmentType:
			return _Action{_ShiftAction, _State324, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State229:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State326, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case ConditionType:
			return _Action{_ShiftAction, _State327, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToCondition}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State230:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State329, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State234:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State330, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State236:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case FieldDefType:
			return _Action{_ShiftAction, _State332, 0}, true
		case ProperExplicitEnumValueDefsType:
			return _Action{_ShiftAction, _State333, 0}, true
		case ExplicitEnumValueDefsType:
			return _Action{_ShiftAction, _State331, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true
		}
	case _State237:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State335, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State336, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State334, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State341, 0}, true
		case ProperParameterDeclsType:
			return _Action{_ShiftAction, _State339, 0}, true
		case ParameterDeclsType:
			return _Action{_ShiftAction, _State338, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDecl}, true
		case ParameterDeclType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDeclToProperParameterDecls}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDecls}, true
		}
	case _State238:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State342, 0}, true
		}
	case _State240:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State343, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State237, 0}, true
		}
	case _State241:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State344, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State345, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLocalToNamedTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true
		}
	case _State242:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State346, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		}
	case _State243:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State347, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceFieldDefToProperImplicitFieldDefs}, true
		}
	case _State244:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitEnumTypeExpr}, true
		}
	case _State245:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitStructTypeExpr}, true
		}
	case _State247:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State351, 0}, true
		case NewlinesToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImproperToImplicitEnumValueDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs}, true
		}
	case _State248:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State352, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImplicitFieldDefsToImplicitFieldDefs}, true
		}
	case _State249:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State353, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		case OptionalDefaultType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnnamedToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}, true
		}
	case _State251:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case ProperExplicitFieldDefsType:
			return _Action{_ShiftAction, _State305, 0}, true
		case ExplicitFieldDefsType:
			return _Action{_ShiftAction, _State355, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToProperExplicitFieldDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToExplicitFieldDefs}, true
		}
	case _State256:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBinaryTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State257:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAliasToNamedFuncDef}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State258:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State358, 0}, true
		case ProperGenericParameterDefListType:
			return _Action{_ShiftAction, _State361, 0}, true
		case GenericParameterDefListType:
			return _Action{_ShiftAction, _State360, 0}, true
		case GenericParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGenericParameterDefToProperGenericParameterDefList}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefList}, true
		}
	case _State259:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State362, 0}, true
		}
	case _State260:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State363, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State364, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedInferredArgToParameterDef}, true
		}
	case _State261:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State365, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State366, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreInferredArgToParameterDef}, true
		}
	case _State262:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State367, 0}, true
		}
	case _State264:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State29, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State31, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State39, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State72, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State87, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State61, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State101, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State56, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpType}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpType}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpType}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFallthroughStatement}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToCallbackOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToCallbackOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case SimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSimpleStatementToStatement}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperStatements}, true
		case ExprOrImproperStructStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}, true
		case CallbackOpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}, true
		case JumpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStatementToSimpleStatement}, true
		case FallthroughStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}, true
		case AssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStatementToSimpleStatement}, true
		case UnaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}, true
		case BinaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}, true
		case ImportStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStatementToStatement}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceImproperImplicitToStatements}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State265:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State29, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State31, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State39, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State72, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State87, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State61, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State101, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State56, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpType}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpType}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpType}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFallthroughStatement}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToCallbackOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToCallbackOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case SimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSimpleStatementToStatement}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperStatements}, true
		case ExprOrImproperStructStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}, true
		case CallbackOpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}, true
		case JumpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStatementToSimpleStatement}, true
		case FallthroughStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}, true
		case AssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStatementToSimpleStatement}, true
		case UnaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}, true
		case BinaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}, true
		case ImportStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStatementToStatement}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceImproperExplicitToStatements}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State267:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State370, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State268:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State371, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State272:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State372, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}, true
		}
	case _State274:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State373, 0}, true
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToTuplePattern}, true
		}
	case _State277:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypeExprToOptionalTypeExpr}, true
		}
	case _State278:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceEnumNondataMatchPattenToCasePattern}, true
		}
	case _State279:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State376, 0}, true
		}
	case _State280:
		switch symbolId {
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State72, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State87, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State61, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State101, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State56, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpType}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpType}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpType}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFallthroughStatement}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToCallbackOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToCallbackOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case SimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSimpleStatementToOptionalSimpleStatement}, true
		case OptionalSimpleStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCaseBranchStatementToStatement}, true
		case ExprOrImproperStructStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}, true
		case CallbackOpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}, true
		case JumpStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStatementToSimpleStatement}, true
		case FallthroughStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}, true
		case AssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStatementToSimpleStatement}, true
		case UnaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}, true
		case BinaryOpAssignStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToExprs}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case _EndMarker:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true
		case NewlinesToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true
		case RbraceToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State281:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case VarToken:
			return _Action{_ShiftAction, _State152, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State151, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case AssignPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignPatternToCasePattern}, true
		case CasePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMultipleToCasePatterns}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignPattern}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State285:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State379, 0}, true
		}
	case _State286:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State380, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperParameterDefsToParameterDefs}, true
		}
	case _State290:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMultipleToImportStatement}, true
		}
	case _State291:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State383, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State382, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}, true
		}
	case _State293:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State384, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State294:
		switch symbolId {
		case IntegerLiteralToken:
			return _Action{_ShiftAction, _State385, 0}, true
		}
	case _State297:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedAssignmentToArgument}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State299:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceColonExprExprTupleToColonExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}, true
		case RbracketToken:
			return _Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}, true
		case CommaToken:
			return _Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}, true
		case ColonToken:
			return _Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State300:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprExprPairToColonExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}, true
		case RbracketToken:
			return _Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}, true
		case CommaToken:
			return _Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}, true
		case ColonToken:
			return _Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State302:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State170, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State174, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State173, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperArguments}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceImproperToArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State303:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToExplicitStructTypeExpr}, true
		}
	case _State305:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State392, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State391, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperExplicitFieldDefsToExplicitFieldDefs}, true
		}
	case _State306:
		switch symbolId {
		case GreaterToken:
			return _Action{_ShiftAction, _State393, 0}, true
		}
	case _State307:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State394, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperTypeArgumentsToTypeArguments}, true
		}
	case _State308:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBindingToGenericTypeArguments}, true
		}
	case _State309:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypeExprToProperTypeArguments}, true
		}
	case _State311:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIndexExpr}, true
		}
	case _State313:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State170, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State174, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State93, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State175, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State397, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State173, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToExpr}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfExprToExpr}, true
		case SwitchExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprToExpr}, true
		case LoopExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprToExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArgumentToProperArguments}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case RparenToken:
			return _Action{_ReduceAction, 0, _ReduceNilToArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State314:
		switch symbolId {
		case MulOpType:
			return _Action{_ShiftAction, _State226, 0}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToMulOp}, true
		case DivToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDivToMulOp}, true
		case ModToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceModToMulOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToMulOp}, true
		case BitLshiftToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitLshiftToMulOp}, true
		case BitRshiftToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitRshiftToMulOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryAddExpr}, true
		}
	case _State315:
		switch symbolId {
		case CmpOpType:
			return _Action{_ShiftAction, _State215, 0}, true
		case EqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEqualToCmpOp}, true
		case NotEqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotEqualToCmpOp}, true
		case LessToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLessToCmpOp}, true
		case LessOrEqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLessOrEqualToCmpOp}, true
		case GreaterToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGreaterToCmpOp}, true
		case GreaterOrEqualToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryAndExpr}, true
		}
	case _State317:
		switch symbolId {
		case AddOpType:
			return _Action{_ShiftAction, _State204, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToAddOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToAddOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToAddOp}, true
		case BitOrToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitOrToAddOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryCmpExpr}, true
		}
	case _State319:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToInitializeExpr}, true
		}
	case _State320:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State216, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}, true
		}
	case _State322:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State399, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceInfiniteToLoopExprBody}, true
		}
	case _State323:
		switch symbolId {
		case InToken:
			return _Action{_ShiftAction, _State401, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State400, 0}, true
		}
	case _State324:
		switch symbolId {
		case SemicolonToken:
			return _Action{_ShiftAction, _State402, 0}, true
		}
	case _State325:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State403, 0}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToAssignPattern}, true
		}
	case _State326:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case VarToken:
			return _Action{_ShiftAction, _State152, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State151, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case CasePatternsType:
			return _Action{_ShiftAction, _State404, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case AssignPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignPatternToCasePattern}, true
		case CasePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCasePatternToCasePatterns}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignPattern}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State327:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAction, _State405, 0}, true
		}
	case _State329:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSwitchExpr}, true
		}
	case _State330:
		switch symbolId {
		case AndToken:
			return _Action{_ShiftAction, _State205, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}, true
		}
	case _State331:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToExplicitEnumTypeExpr}, true
		}
	case _State332:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State408, 0}, true
		case OrToken:
			return _Action{_ShiftAction, _State409, 0}, true
		}
	case _State333:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State410, 0}, true
		case OrToken:
			return _Action{_ShiftAction, _State411, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs}, true
		}
	case _State334:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State412, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedInferredVarargToParameterDecl}, true
		}
	case _State335:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State344, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State363, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State364, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLocalToNamedTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true
		}
	case _State336:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State365, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State366, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		}
	case _State338:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State413, 0}, true
		}
	case _State339:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State414, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperParameterDeclsToParameterDecls}, true
		}
	case _State341:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedTypedArgToParameterDecl}, true
		}
	case _State342:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case GenericTypeArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExternalToNamedTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}, true
		}
	case _State343:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State416, 0}, true
		}
	case _State344:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State342, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		}
	case _State345:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State353, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		case OptionalDefaultType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}, true
		}
	case _State346:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceStructPaddingToFieldDef}, true
		}
	case _State347:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePairToProperImplicitEnumValueDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true
		}
	case _State351:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperImplicitEnumValueDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true
		}
	case _State352:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperImplicitFieldDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToImplicitFieldDefs}, true
		}
	case _State353:
		switch symbolId {
		case DefaultToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDefaultToOptionalDefault}, true
		}
	case _State355:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToTraitTypeExpr}, true
		}
	case _State358:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State423, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}, true
		}
	case _State360:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGenericToGenericParameterDefs}, true
		}
	case _State361:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State425, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperGenericParameterDefListToGenericParameterDefList}, true
		}
	case _State362:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State260, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State261, 0}, true
		case ProperParameterDefsType:
			return _Action{_ShiftAction, _State286, 0}, true
		case ParameterDefsType:
			return _Action{_ShiftAction, _State426, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDefToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDefs}, true
		}
	case _State363:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State427, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedInferredVarargToProperParameterDef}, true
		}
	case _State364:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedTypedArgToProperParameterDef}, true
		}
	case _State365:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State428, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreInferredVarargToProperParameterDef}, true
		}
	case _State366:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreTypedArgToProperParameterDef}, true
		}
	case _State367:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State429, 0}, true
		}
	case _State370:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAliasToTypeDef}, true
		}
	case _State371:
		switch symbolId {
		case ImplementsToken:
			return _Action{_ShiftAction, _State430, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}, true
		}
	case _State372:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State147, 0}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToVarPattern}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToVarPattern}, true
		case VarPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedToFieldVarPattern}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTuplePatternToVarPattern}, true
		}
	case _State373:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State272, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State147, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToVarPattern}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEllipsisToFieldVarPattern}, true
		case VarPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePositionalToFieldVarPattern}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTuplePatternToVarPattern}, true
		case FieldVarPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToFieldVarPatterns}, true
		}
	case _State376:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State147, 0}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}, true
		}
	case _State379:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case ReturnTypeType:
			return _Action{_ShiftAction, _State434, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToReturnType}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State380:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State260, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State261, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToParameterDefs}, true
		}
	case _State382:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State165, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperImportClauses}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}, true
		}
	case _State383:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State165, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperImportClauses}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}, true
		}
	case _State384:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToMapTypeExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State385:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToArrayTypeExpr}, true
		}
	case _State391:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperExplicitFieldDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitFieldDefs}, true
		}
	case _State392:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperExplicitFieldDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitFieldDefs}, true
		}
	case _State393:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatement}, true
		}
	case _State394:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State444, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToTypeArguments}, true
		}
	case _State397:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToCallExpr}, true
		}
	case _State399:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDoWhileToLoopExprBody}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State400:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToForAssignment}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State401:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State448, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State402:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case OptionalSequenceExprType:
			return _Action{_ShiftAction, _State449, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State403:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWhileToLoopExprBody}, true
		}
	case _State404:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State281, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State452, 0}, true
		}
	case _State405:
		switch symbolId {
		case ElseToken:
			return _Action{_ShiftAction, _State453, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNoElseToIfExprBody}, true
		}
	case _State408:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitPairToProperExplicitEnumValueDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true
		}
	case _State409:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitPairToProperExplicitEnumValueDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true
		}
	case _State410:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitAddToProperExplicitEnumValueDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToExplicitEnumValueDefs}, true
		}
	case _State411:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State241, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State242, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State240, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State249, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case UnsafeStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementToFieldDef}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitAddToProperExplicitEnumValueDefs}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		case MethodSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToFieldDef}, true
		}
	case _State412:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedTypedVarargToParameterDecl}, true
		}
	case _State413:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToReturnType}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case ReturnTypeType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFuncTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State414:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State335, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State336, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State334, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State341, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDecl}, true
		case ParameterDeclType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperParameterDecls}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToParameterDecls}, true
		}
	case _State416:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State335, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State336, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State334, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State341, 0}, true
		case ProperParameterDeclsType:
			return _Action{_ShiftAction, _State339, 0}, true
		case ParameterDeclsType:
			return _Action{_ShiftAction, _State460, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDecl}, true
		case ParameterDeclType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDeclToProperParameterDecls}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDecls}, true
		}
	case _State423:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}, true
		}
	case _State425:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State358, 0}, true
		case GenericParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperGenericParameterDefList}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToGenericParameterDefList}, true
		}
	case _State426:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State462, 0}, true
		}
	case _State427:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedTypedVarargToProperParameterDef}, true
		}
	case _State428:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreTypedVarargToProperParameterDef}, true
		}
	case _State429:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State463, 0}, true
		}
	case _State430:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State464, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true
		}
	case _State434:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAnonymousFuncExpr}, true
		}
	case _State444:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAddToProperTypeArguments}, true
		}
	case _State448:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State466, 0}, true
		}
	case _State449:
		switch symbolId {
		case SemicolonToken:
			return _Action{_ShiftAction, _State467, 0}, true
		}
	case _State452:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCaseToCondition}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State453:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State470, 0}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseToIfExprBody}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMultiIfElseToIfExprBody}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State460:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State472, 0}, true
		}
	case _State462:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case ReturnTypeType:
			return _Action{_ShiftAction, _State473, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToReturnType}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State463:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State260, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State261, 0}, true
		case ProperParameterDefsType:
			return _Action{_ShiftAction, _State286, 0}, true
		case ParameterDefsType:
			return _Action{_ShiftAction, _State474, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDefToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDefs}, true
		}
	case _State464:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State256, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}, true
		}
	case _State466:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIteratorToLoopExprBody}, true
		}
	case _State467:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State36, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State37, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State24, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State156, 0}, true
		case OptionalSequenceExprType:
			return _Action{_ShiftAction, _State476, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State108, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State99, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State91, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State57, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State74, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State94, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State84, 0}, true
		case IntegerLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}, true
		case FloatLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}, true
		case RuneLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToLiteralExpr}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case TrueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTrueToLiteralExpr}, true
		case FalseToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFalseToLiteralExpr}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToVarOrLet}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToVarOrLet}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case VarDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}, true
		case SequenceExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case BlockExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBlockExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case PostfixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}, true
		case PostfixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}, true
		case PrefixableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixableExprToMulExpr}, true
		case PrefixUnaryExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}, true
		case BinaryMulExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryMulExprToMulExpr}, true
		case BinaryAddExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAddExprToAddExpr}, true
		case BinaryCmpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}, true
		case BinaryAndExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAndExprToAndExpr}, true
		case BinaryOrExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryOrExprToOrExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case AnonymousFuncExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}, true
		case DoToken:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State470:
		switch symbolId {
		case IfToken:
			return _Action{_ShiftAction, _State229, 0}, true
		case IfExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIfExpr}, true
		}
	case _State472:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToReturnType}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case ReturnTypeType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToMethodSignature}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State473:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToNamedFuncDef}, true
		}
	case _State474:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State479, 0}, true
		}
	case _State476:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State480, 0}, true
		}
	case _State479:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State50, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State115, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State117, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State134, 0}, true
		case ReturnTypeType:
			return _Action{_ShiftAction, _State481, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitNegToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorTypeExpr}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case AtomTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}, true
		case NamedTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}, true
		case InferredTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}, true
		case ParseErrorTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}, true
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToReturnType}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case ImplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case TraitTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}, true
		case ImplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}, true
		case ExplicitEnumTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}, true
		case FuncTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State480:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToLoopExprBody}, true
		}
	case _State481:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodDefToNamedFuncDef}, true
		}
	}

	return _Action{}, false
}

var _ActionTable = _ActionTableType{}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.source
    Reduce:
      * -> [definitions]
    ShiftAndReduce:
      COMMENT_GROUPS -> [definition]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      definitions -> [source]
      definition -> [proper_definitions]
      statement_block -> [definition]
      type_def -> [definition]
      named_func_def -> [definition]
      package_def -> [definition]
    Goto:
      PACKAGE -> State 13
      TYPE -> State 14
      FUNC -> State 10
      LBRACE -> State 11
      source -> State 5
      proper_definitions -> State 20
      var_decl_pattern -> State 23
      var_or_let -> State 24

  State 2:
    Kernel Items:
      #accept: ^.statement
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_type]
      BREAK -> [jump_type]
      CONTINUE -> [jump_type]
      FALLTHROUGH -> [fallthrough_statement]
      ASYNC -> [callback_op]
      DEFER -> [callback_op]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      simple_statement -> [statement]
      expr_or_improper_struct_statement -> [simple_statement]
      callback_op_statement -> [simple_statement]
      unsafe_statement -> [simple_statement]
      jump_statement -> [simple_statement]
      fallthrough_statement -> [simple_statement]
      assign_statement -> [simple_statement]
      unary_op_assign_statement -> [simple_statement]
      binary_op_assign_statement -> [simple_statement]
      import_statement -> [statement]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      CASE -> State 29
      DEFAULT -> State 31
      IMPORT -> State 39
      UNSAFE -> State 54
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      statement -> State 6
      exprs -> State 78
      callback_op -> State 72
      jump_type -> State 87
      var_or_let -> State 24
      assign_pattern -> State 61
      optional_label_decl -> State 93
      sequence_expr -> State 101
      accessible_expr -> State 56
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 3:
    Kernel Items:
      #accept: ^.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      expr -> State 7
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 4:
    Kernel Items:
      #accept: ^.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 8

  State 5:
    Kernel Items:
      #accept: ^ source., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 6:
    Kernel Items:
      #accept: ^ statement., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 7:
    Kernel Items:
      #accept: ^ expr., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ type_expr., $
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 10:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expr
    Reduce:
      (nil)
    ShiftAndReduce:
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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_type]
      BREAK -> [jump_type]
      CONTINUE -> [jump_type]
      FALLTHROUGH -> [fallthrough_statement]
      ASYNC -> [callback_op]
      DEFER -> [callback_op]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      simple_statement -> [statement]
      statement -> [proper_statements]
      expr_or_improper_struct_statement -> [simple_statement]
      callback_op_statement -> [simple_statement]
      unsafe_statement -> [simple_statement]
      jump_statement -> [simple_statement]
      fallthrough_statement -> [simple_statement]
      assign_statement -> [simple_statement]
      unary_op_assign_statement -> [simple_statement]
      binary_op_assign_statement -> [simple_statement]
      import_statement -> [statement]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      CASE -> State 29
      DEFAULT -> State 31
      IMPORT -> State 39
      UNSAFE -> State 54
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      proper_statements -> State 139
      statements -> State 141
      exprs -> State 78
      callback_op -> State 72
      jump_type -> State 87
      var_or_let -> State 24
      assign_pattern -> State 61
      optional_label_decl -> State 93
      sequence_expr -> State 101
      accessible_expr -> State 56
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 13:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    ShiftAndReduce:
      statement_block -> [package_def]
    Goto:
      LBRACE -> State 11

  State 14:
    Kernel Items:
      type_def: TYPE.IDENTIFIER generic_parameter_defs type_expr
      type_def: TYPE.IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS type_expr
      type_def: TYPE.IDENTIFIER ASSIGN type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143

  State 20:
    Kernel Items:
      proper_definitions: proper_definitions.NEWLINES definition
      definitions: proper_definitions., *
      definitions: proper_definitions.NEWLINES
    Reduce:
      * -> [definitions]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 144

  State 23:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expr
    Reduce:
      * -> [definition]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 145

  State 24:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [var_pattern]
      UNDERSCORE -> [var_pattern]
      tuple_pattern -> [var_pattern]
    Goto:
      LPAREN -> State 147
      var_pattern -> State 150

  State 29:
    Kernel Items:
      statement: CASE.case_patterns COLON optional_simple_statement
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      assign_pattern -> [case_pattern]
      case_pattern -> [case_patterns]
      sequence_expr -> [assign_pattern]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 152
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 151
      GREATER -> State 37
      case_patterns -> State 155
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 31:
    Kernel Items:
      statement: DEFAULT.COLON optional_simple_statement
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 158

  State 36:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 159

  State 37:
    Kernel Items:
      sequence_expr: GREATER.sequence_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [sequence_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 39:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [import_statement]
    Goto:
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      LPAREN -> State 163
      DOT -> State 161

  State 42:
    Kernel Items:
      slice_type_expr: LBRACKET.type_expr RBRACKET
      array_type_expr: LBRACKET.type_expr COMMA INTEGER_LITERAL RBRACKET
      map_type_expr: LBRACKET.type_expr COLON type_expr RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 167

  State 43:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      ELLIPSIS -> [argument]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      IDENTIFIER -> State 170
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 168
      GREATER -> State 37
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      proper_arguments -> State 175
      arguments -> State 172
      colon_expr -> State 173
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 50:
    Kernel Items:
      explicit_struct_type_expr: STRUCT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 176

  State 54:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LESS -> State 177

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
    ShiftAndReduce:
      QUESTION -> [postfix_unary_op]
      EXCLAIM -> [postfix_unary_op]
      ADD_ASSIGN -> [binary_op_assign]
      SUB_ASSIGN -> [binary_op_assign]
      MUL_ASSIGN -> [binary_op_assign]
      DIV_ASSIGN -> [binary_op_assign]
      MOD_ASSIGN -> [binary_op_assign]
      ADD_ONE_ASSIGN -> [unary_op_assign]
      SUB_ONE_ASSIGN -> [unary_op_assign]
      BIT_NEG_ASSIGN -> [binary_op_assign]
      BIT_AND_ASSIGN -> [binary_op_assign]
      BIT_OR_ASSIGN -> [binary_op_assign]
      BIT_XOR_ASSIGN -> [binary_op_assign]
      BIT_LSHIFT_ASSIGN -> [binary_op_assign]
      BIT_RSHIFT_ASSIGN -> [binary_op_assign]
      unary_op_assign -> [unary_op_assign_statement]
      postfix_unary_op -> [postfix_unary_expr]
    Goto:
      LBRACKET -> State 190
      DOT -> State 188
      DOLLAR_LBRACKET -> State 187
      binary_op_assign -> State 196
      generic_type_arguments -> State 197

  State 57:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    ShiftAndReduce:
      ADD -> [add_op]
      SUB -> [add_op]
      BIT_XOR -> [add_op]
      BIT_OR -> [add_op]
    Goto:
      add_op -> State 204

  State 58:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AND -> State 205

  State 61:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 206

  State 72:
    Kernel Items:
      callback_op_statement: callback_op.call_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      LABEL_DECL -> [optional_label_decl]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      call_expr -> State 208
      accessible_expr -> State 207
      initializable_type_expr -> State 84

  State 74:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    ShiftAndReduce:
      EQUAL -> [cmp_op]
      NOT_EQUAL -> [cmp_op]
      LESS -> [cmp_op]
      LESS_OR_EQUAL -> [cmp_op]
      GREATER -> [cmp_op]
      GREATER_OR_EQUAL -> [cmp_op]
    Goto:
      cmp_op -> State 215

  State 78:
    Kernel Items:
      expr_or_improper_struct_statement: exprs., *
      exprs: exprs.COMMA expr
    Reduce:
      * -> [expr_or_improper_struct_statement]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 216

  State 84:
    Kernel Items:
      initialize_expr: initializable_type_expr.LPAREN arguments RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 217

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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      JUMP_LABEL -> State 218
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      exprs -> State 219
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 91:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    ShiftAndReduce:
      MUL -> [mul_op]
      DIV -> [mul_op]
      MOD -> [mul_op]
      BIT_AND -> [mul_op]
      BIT_LSHIFT -> [mul_op]
      BIT_RSHIFT -> [mul_op]
    Goto:
      mul_op -> State 226

  State 93:
    Kernel Items:
      if_expr: optional_label_decl.if_expr_body
      switch_expr: optional_label_decl.SWITCH sequence_expr statement_block
      loop_expr: optional_label_decl.loop_expr_body
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [block_expr]
      if_expr_body -> [if_expr]
      loop_expr_body -> [loop_expr]
    Goto:
      IF -> State 229
      SWITCH -> State 230
      FOR -> State 228
      DO -> State 227
      LBRACE -> State 11

  State 94:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 234

  State 99:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefixable_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [prefix_unary_expr]
      prefix_unary_expr -> [prefixable_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      initializable_type_expr -> State 84

  State 101:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expr: sequence_expr., *
    Reduce:
      * -> [expr]
      ASSIGN -> [assign_pattern]
    ShiftAndReduce:
      (nil)
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
    ShiftAndReduce:
      QUESTION -> [postfix_unary_op]
      EXCLAIM -> [postfix_unary_op]
      postfix_unary_op -> [postfix_unary_expr]
    Goto:
      LBRACKET -> State 190
      DOT -> State 188
      DOLLAR_LBRACKET -> State 187
      generic_type_arguments -> State 197

  State 113:
    Kernel Items:
      explicit_enum_type_expr: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 236

  State 115:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 237

  State 116:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    ShiftAndReduce:
      generic_type_arguments -> [named_type_expr]
    Goto:
      DOT -> State 238
      DOLLAR_LBRACKET -> State 187

  State 117:
    Kernel Items:
      implicit_struct_type_expr: LPAREN.implicit_field_defs RPAREN
      implicit_enum_type_expr: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      * -> [implicit_field_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      field_def -> State 243
      proper_implicit_field_defs -> State 248
      implicit_field_defs -> State 245
      proper_implicit_enum_value_defs -> State 247
      implicit_enum_value_defs -> State 244

  State 121:
    Kernel Items:
      trait_type_expr: TRAIT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 251

  State 134:
    Kernel Items:
      prefix_unary_type_expr: prefix_unary_type_op.returnable_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [prefix_unary_type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134

  State 137:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expr
    Reduce:
      * -> [generic_parameter_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      DOLLAR_LBRACKET -> State 258
      ASSIGN -> State 257
      generic_parameter_defs -> State 259

  State 138:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      parameter_def -> State 262

  State 139:
    Kernel Items:
      proper_statements: proper_statements.NEWLINES statement
      proper_statements: proper_statements.SEMICOLON statement
      statements: proper_statements., *
      statements: proper_statements.NEWLINES
      statements: proper_statements.SEMICOLON
    Reduce:
      * -> [statements]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 264
      SEMICOLON -> State 265

  State 141:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACE -> [statement_block]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      type_def: TYPE IDENTIFIER.generic_parameter_defs type_expr
      type_def: TYPE IDENTIFIER.generic_parameter_defs type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [generic_parameter_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      DOLLAR_LBRACKET -> State 258
      ASSIGN -> State 267
      generic_parameter_defs -> State 268

  State 144:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES.definition
      definitions: proper_definitions NEWLINES., *
    Reduce:
      * -> [definitions]
    ShiftAndReduce:
      COMMENT_GROUPS -> [definition]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      definition -> [proper_definitions]
      statement_block -> [definition]
      type_def -> [definition]
      named_func_def -> [definition]
      package_def -> [definition]
    Goto:
      PACKAGE -> State 13
      TYPE -> State 14
      FUNC -> State 10
      LBRACE -> State 11
      var_decl_pattern -> State 23
      var_or_let -> State 24

  State 145:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [definition]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 147:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [var_pattern]
      ELLIPSIS -> [field_var_pattern]
      var_pattern -> [field_var_pattern]
      tuple_pattern -> [var_pattern]
      field_var_pattern -> [field_var_patterns]
    Goto:
      IDENTIFIER -> State 272
      LPAREN -> State 147
      field_var_patterns -> State 274

  State 150:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_type_expr
    Reduce:
      * -> [optional_type_expr]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      optional_type_expr -> [var_decl_pattern]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 277

  State 151:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
      case_pattern: DOT.IDENTIFIER
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 278

  State 152:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    ShiftAndReduce:
      (nil)
    Goto:
      DOT -> State 279

  State 155:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 281
      COLON -> State 280

  State 156:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [block_expr]
    Goto:
      LBRACE -> State 11

  State 158:
    Kernel Items:
      statement: DEFAULT COLON.optional_simple_statement
    Reduce:
      * -> [optional_label_decl]
      $ -> [optional_simple_statement]
      NEWLINES -> [optional_simple_statement]
      RBRACE -> [optional_simple_statement]
      SEMICOLON -> [optional_simple_statement]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_type]
      BREAK -> [jump_type]
      CONTINUE -> [jump_type]
      FALLTHROUGH -> [fallthrough_statement]
      ASYNC -> [callback_op]
      DEFER -> [callback_op]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      simple_statement -> [optional_simple_statement]
      optional_simple_statement -> [statement]
      expr_or_improper_struct_statement -> [simple_statement]
      callback_op_statement -> [simple_statement]
      unsafe_statement -> [simple_statement]
      jump_statement -> [simple_statement]
      fallthrough_statement -> [simple_statement]
      assign_statement -> [simple_statement]
      unary_op_assign_statement -> [simple_statement]
      binary_op_assign_statement -> [simple_statement]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      UNSAFE -> State 54
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      exprs -> State 78
      callback_op -> State 72
      jump_type -> State 87
      var_or_let -> State 24
      assign_pattern -> State 61
      optional_label_decl -> State 93
      sequence_expr -> State 101
      accessible_expr -> State 56
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 159:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_defs -> State 286
      parameter_defs -> State 285

  State 161:
    Kernel Items:
      import_clause: DOT.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      import_clause: IDENTIFIER.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      DOT -> State 161
      proper_import_clauses -> State 291
      import_clauses -> State 290

  State 165:
    Kernel Items:
      import_clause: UNDERSCORE.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
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
    ShiftAndReduce:
      RBRACKET -> [slice_type_expr]
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      COMMA -> State 294
      COLON -> State 293
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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [colon_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 170:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expr
      named_expr: IDENTIFIER., *
    Reduce:
      * -> [named_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 297

  State 172:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_struct_expr]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [argument]
    ShiftAndReduce:
      (nil)
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
    ShiftAndReduce:
      ELLIPSIS -> [argument]
    Goto:
      COLON -> State 300

  State 175:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., *
      arguments: proper_arguments.COMMA
    Reduce:
      * -> [arguments]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 302

  State 176:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN.explicit_field_defs RPAREN
    Reduce:
      * -> [explicit_field_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_field_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      proper_explicit_field_defs -> State 305
      explicit_field_defs -> State 303

  State 177:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 306

  State 187:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET.type_arguments RBRACKET
    Reduce:
      * -> [type_arguments]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 309
      proper_type_arguments -> State 307
      type_arguments -> State 308

  State 188:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [access_expr]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      ELLIPSIS -> [argument]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      IDENTIFIER -> State 170
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 168
      GREATER -> State 37
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      argument -> State 311
      colon_expr -> State 173
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 196:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [binary_op_assign_statement]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 197:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments.LPAREN arguments RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 313

  State 204:
    Kernel Items:
      binary_add_expr: add_expr add_op.mul_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 314
      initializable_type_expr -> State 84

  State 205:
    Kernel Items:
      binary_and_expr: and_expr AND.cmp_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 315
      initializable_type_expr -> State 84

  State 206:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [assign_statement]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 207:
    Kernel Items:
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      * -> [generic_type_arguments]
    ShiftAndReduce:
      (nil)
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
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 215:
    Kernel Items:
      binary_cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 317
      initializable_type_expr -> State 84

  State 216:
    Kernel Items:
      exprs: exprs COMMA.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 217:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      ELLIPSIS -> [argument]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      IDENTIFIER -> State 170
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 168
      GREATER -> State 37
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      proper_arguments -> State 175
      arguments -> State 319
      colon_expr -> State 173
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      exprs -> State 320
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 219:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type exprs., *
    Reduce:
      * -> [jump_statement]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 216

  State 226:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op.prefixable_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [binary_mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      initializable_type_expr -> State 84

  State 227:
    Kernel Items:
      loop_expr_body: DO.statement_block
      loop_expr_body: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    ShiftAndReduce:
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
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      assign_pattern -> State 323
      optional_label_decl -> State 156
      sequence_expr -> State 325
      for_assignment -> State 324
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 229:
    Kernel Items:
      if_expr_body: IF.condition statement_block
      if_expr_body: IF.condition statement_block ELSE statement_block
      if_expr_body: IF.condition statement_block ELSE if_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [condition]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      CASE -> State 326
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      condition -> State 327
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 230:
    Kernel Items:
      switch_expr: optional_label_decl SWITCH.sequence_expr statement_block
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 329
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 234:
    Kernel Items:
      binary_or_expr: or_expr OR.and_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 330
      initializable_type_expr -> State 84

  State 236:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      field_def -> State 332
      proper_explicit_enum_value_defs -> State 333
      explicit_enum_value_defs -> State 331

  State 237:
    Kernel Items:
      func_type_expr: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      * -> [parameter_decls]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      proper_parameter_def -> [parameter_decl]
      parameter_decl -> [proper_parameter_decls]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 335
      UNDERSCORE -> State 336
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      ELLIPSIS -> State 334
      prefix_unary_type_op -> State 134
      type_expr -> State 341
      proper_parameter_decls -> State 339
      parameter_decls -> State 338

  State 238:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 342

  State 240:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
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
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      generic_type_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 344
      DOLLAR_LBRACKET -> State 187
      prefix_unary_type_op -> State 134
      type_expr -> State 345

  State 242:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      field_def: UNDERSCORE.type_expr
    Reduce:
      * -> [inferred_type_expr]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 346

  State 243:
    Kernel Items:
      proper_implicit_field_defs: field_def., *
      proper_implicit_enum_value_defs: field_def.OR field_def
    Reduce:
      * -> [proper_implicit_field_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 347

  State 244:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs.OR field_def
      implicit_enum_value_defs: proper_implicit_enum_value_defs., *
      implicit_enum_value_defs: proper_implicit_enum_value_defs.NEWLINES
    Reduce:
      * -> [implicit_enum_value_defs]
    ShiftAndReduce:
      NEWLINES -> [implicit_enum_value_defs]
    Goto:
      OR -> State 351

  State 248:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., *
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      * -> [implicit_field_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 352

  State 249:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: type_expr.optional_default
    Reduce:
      * -> [optional_default]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
      optional_default -> [field_def]
    Goto:
      ASSIGN -> State 353
      binary_type_op -> State 256

  State 251:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN.explicit_field_defs RPAREN
    Reduce:
      * -> [explicit_field_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_field_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      proper_explicit_field_defs -> State 305
      explicit_field_defs -> State 355

  State 256:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op.returnable_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [binary_type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134

  State 257:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [named_func_def]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 258:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET.generic_parameter_def_list RBRACKET
    Reduce:
      * -> [generic_parameter_def_list]
    ShiftAndReduce:
      generic_parameter_def -> [proper_generic_parameter_def_list]
    Goto:
      IDENTIFIER -> State 358
      proper_generic_parameter_def_list -> State 361
      generic_parameter_def_list -> State 360

  State 259:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
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
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      ELLIPSIS -> State 363
      prefix_unary_type_op -> State 134
      type_expr -> State 364

  State 261:
    Kernel Items:
      proper_parameter_def: UNDERSCORE.type_expr
      proper_parameter_def: UNDERSCORE.ELLIPSIS
      proper_parameter_def: UNDERSCORE.ELLIPSIS type_expr
      parameter_def: UNDERSCORE., *
    Reduce:
      * -> [parameter_def]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      ELLIPSIS -> State 365
      prefix_unary_type_op -> State 134
      type_expr -> State 366

  State 262:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 367

  State 264:
    Kernel Items:
      proper_statements: proper_statements NEWLINES.statement
      statements: proper_statements NEWLINES., RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_type]
      BREAK -> [jump_type]
      CONTINUE -> [jump_type]
      FALLTHROUGH -> [fallthrough_statement]
      ASYNC -> [callback_op]
      DEFER -> [callback_op]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      simple_statement -> [statement]
      statement -> [proper_statements]
      expr_or_improper_struct_statement -> [simple_statement]
      callback_op_statement -> [simple_statement]
      unsafe_statement -> [simple_statement]
      jump_statement -> [simple_statement]
      fallthrough_statement -> [simple_statement]
      assign_statement -> [simple_statement]
      unary_op_assign_statement -> [simple_statement]
      binary_op_assign_statement -> [simple_statement]
      import_statement -> [statement]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      CASE -> State 29
      DEFAULT -> State 31
      IMPORT -> State 39
      UNSAFE -> State 54
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      exprs -> State 78
      callback_op -> State 72
      jump_type -> State 87
      var_or_let -> State 24
      assign_pattern -> State 61
      optional_label_decl -> State 93
      sequence_expr -> State 101
      accessible_expr -> State 56
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 265:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON.statement
      statements: proper_statements SEMICOLON., RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_type]
      BREAK -> [jump_type]
      CONTINUE -> [jump_type]
      FALLTHROUGH -> [fallthrough_statement]
      ASYNC -> [callback_op]
      DEFER -> [callback_op]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      simple_statement -> [statement]
      statement -> [proper_statements]
      expr_or_improper_struct_statement -> [simple_statement]
      callback_op_statement -> [simple_statement]
      unsafe_statement -> [simple_statement]
      jump_statement -> [simple_statement]
      fallthrough_statement -> [simple_statement]
      assign_statement -> [simple_statement]
      unary_op_assign_statement -> [simple_statement]
      binary_op_assign_statement -> [simple_statement]
      import_statement -> [statement]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      CASE -> State 29
      DEFAULT -> State 31
      IMPORT -> State 39
      UNSAFE -> State 54
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      exprs -> State 78
      callback_op -> State 72
      jump_type -> State 87
      var_or_let -> State 24
      assign_pattern -> State 61
      optional_label_decl -> State 93
      sequence_expr -> State 101
      accessible_expr -> State 56
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 267:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 370

  State 268:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameter_defs.type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs.type_expr IMPLEMENTS type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 371

  State 272:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 372

  State 274:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [tuple_pattern]
    Goto:
      COMMA -> State 373

  State 277:
    Kernel Items:
      optional_type_expr: type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [optional_type_expr]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 278:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
      case_pattern: DOT IDENTIFIER., *
    Reduce:
      * -> [case_pattern]
    ShiftAndReduce:
      implicit_struct_expr -> [case_pattern]
    Goto:
      LPAREN -> State 43

  State 279:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_type]
      BREAK -> [jump_type]
      CONTINUE -> [jump_type]
      FALLTHROUGH -> [fallthrough_statement]
      ASYNC -> [callback_op]
      DEFER -> [callback_op]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      simple_statement -> [optional_simple_statement]
      optional_simple_statement -> [statement]
      expr_or_improper_struct_statement -> [simple_statement]
      callback_op_statement -> [simple_statement]
      unsafe_statement -> [simple_statement]
      jump_statement -> [simple_statement]
      fallthrough_statement -> [simple_statement]
      assign_statement -> [simple_statement]
      unary_op_assign_statement -> [simple_statement]
      binary_op_assign_statement -> [simple_statement]
      var_decl_pattern -> [sequence_expr]
      expr -> [exprs]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      UNSAFE -> State 54
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      exprs -> State 78
      callback_op -> State 72
      jump_type -> State 87
      var_or_let -> State 24
      assign_pattern -> State 61
      optional_label_decl -> State 93
      sequence_expr -> State 101
      accessible_expr -> State 56
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 281:
    Kernel Items:
      case_patterns: case_patterns COMMA.case_pattern
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      assign_pattern -> [case_pattern]
      case_pattern -> [case_patterns]
      sequence_expr -> [assign_pattern]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 152
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 151
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 285:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 379

  State 286:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., *
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 380

  State 290:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [import_statement]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., *
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 383
      COMMA -> State 382

  State 293:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON.type_expr RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 384

  State 294:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 385

  State 297:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [argument]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [colon_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

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
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      expr -> [colon_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 93
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 302:
    Kernel Items:
      proper_arguments: proper_arguments COMMA.argument
      arguments: proper_arguments COMMA., RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      ELLIPSIS -> [argument]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      IDENTIFIER -> State 170
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 168
      GREATER -> State 37
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      colon_expr -> State 173
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 303:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [explicit_struct_type_expr]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., *
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      * -> [explicit_field_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 392
      COMMA -> State 391

  State 306:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      GREATER -> State 393

  State 307:
    Kernel Items:
      proper_type_arguments: proper_type_arguments.COMMA type_expr
      type_arguments: proper_type_arguments., *
      type_arguments: proper_type_arguments.COMMA
    Reduce:
      * -> [type_arguments]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 394

  State 308:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [generic_type_arguments]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_type_arguments: type_expr., *
    Reduce:
      * -> [proper_type_arguments]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 311:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [index_expr]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      ELLIPSIS -> [argument]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [expr]
      if_expr -> [expr]
      switch_expr -> [expr]
      loop_expr -> [expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      IDENTIFIER -> State 170
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 168
      GREATER -> State 37
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      proper_arguments -> State 175
      arguments -> State 397
      colon_expr -> State 173
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 314:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      binary_add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [binary_add_expr]
    ShiftAndReduce:
      MUL -> [mul_op]
      DIV -> [mul_op]
      MOD -> [mul_op]
      BIT_AND -> [mul_op]
      BIT_LSHIFT -> [mul_op]
      BIT_RSHIFT -> [mul_op]
    Goto:
      mul_op -> State 226

  State 315:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      binary_and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [binary_and_expr]
    ShiftAndReduce:
      EQUAL -> [cmp_op]
      NOT_EQUAL -> [cmp_op]
      LESS -> [cmp_op]
      LESS_OR_EQUAL -> [cmp_op]
      GREATER -> [cmp_op]
      GREATER_OR_EQUAL -> [cmp_op]
    Goto:
      cmp_op -> State 215

  State 317:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      binary_cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [binary_cmp_expr]
    ShiftAndReduce:
      ADD -> [add_op]
      SUB -> [add_op]
      BIT_XOR -> [add_op]
      BIT_OR -> [add_op]
    Goto:
      add_op -> State 204

  State 319:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [initialize_expr]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type JUMP_LABEL exprs., *
    Reduce:
      * -> [jump_statement]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 216

  State 322:
    Kernel Items:
      loop_expr_body: DO statement_block., *
      loop_expr_body: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr_body]
    ShiftAndReduce:
      (nil)
    Goto:
      FOR -> State 399

  State 323:
    Kernel Items:
      loop_expr_body: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IN -> State 401
      ASSIGN -> State 400

  State 324:
    Kernel Items:
      loop_expr_body: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
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
    ShiftAndReduce:
      (nil)
    Goto:
      DO -> State 403

  State 326:
    Kernel Items:
      condition: CASE.case_patterns ASSIGN sequence_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      assign_pattern -> [case_pattern]
      case_pattern -> [case_patterns]
      sequence_expr -> [assign_pattern]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 152
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 151
      GREATER -> State 37
      case_patterns -> State 404
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 327:
    Kernel Items:
      if_expr_body: IF condition.statement_block
      if_expr_body: IF condition.statement_block ELSE statement_block
      if_expr_body: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 405

  State 329:
    Kernel Items:
      switch_expr: optional_label_decl SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [switch_expr]
    Goto:
      LBRACE -> State 11

  State 330:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AND -> State 205

  State 331:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def.OR field_def
      proper_explicit_enum_value_defs: field_def.NEWLINES field_def
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 408
      OR -> State 409

  State 333:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.OR field_def
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs., *
      explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES
    Reduce:
      * -> [explicit_enum_value_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 410
      OR -> State 411

  State 334:
    Kernel Items:
      parameter_decl: ELLIPSIS., *
      parameter_decl: ELLIPSIS.type_expr
    Reduce:
      * -> [parameter_decl]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 412

  State 335:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      proper_parameter_def: IDENTIFIER.type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS
    Reduce:
      * -> [generic_type_arguments]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      generic_type_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 344
      DOLLAR_LBRACKET -> State 187
      ELLIPSIS -> State 363
      prefix_unary_type_op -> State 134
      type_expr -> State 364

  State 336:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      proper_parameter_def: UNDERSCORE.type_expr
      proper_parameter_def: UNDERSCORE.ELLIPSIS
      proper_parameter_def: UNDERSCORE.ELLIPSIS type_expr
    Reduce:
      * -> [inferred_type_expr]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      ELLIPSIS -> State 365
      prefix_unary_type_op -> State 134
      type_expr -> State 366

  State 338:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 413

  State 339:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., *
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      * -> [parameter_decls]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 414

  State 341:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: type_expr., *
    Reduce:
      * -> [parameter_decl]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 342:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    ShiftAndReduce:
      generic_type_arguments -> [named_type_expr]
    Goto:
      DOLLAR_LBRACKET -> State 187

  State 343:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 416

  State 344:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 342

  State 345:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: IDENTIFIER type_expr.optional_default
    Reduce:
      * -> [optional_default]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
      optional_default -> [field_def]
    Goto:
      ASSIGN -> State 353
      binary_type_op -> State 256

  State 346:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [field_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 347:
    Kernel Items:
      proper_implicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_implicit_enum_value_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 351:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_implicit_enum_value_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 352:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA.field_def
      implicit_field_defs: proper_implicit_field_defs COMMA., *
    Reduce:
      * -> [implicit_field_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_implicit_field_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 353:
    Kernel Items:
      optional_default: ASSIGN.DEFAULT
    Reduce:
      (nil)
    ShiftAndReduce:
      DEFAULT -> [optional_default]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [trait_type_expr]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_parameter_def]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 423

  State 360:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET generic_parameter_def_list.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      (nil)

  State 361:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list.COMMA generic_parameter_def
      generic_parameter_def_list: proper_generic_parameter_def_list., *
      generic_parameter_def_list: proper_generic_parameter_def_list.COMMA
    Reduce:
      * -> [generic_parameter_def_list]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 425

  State 362:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_defs -> State 286
      parameter_defs -> State 426

  State 363:
    Kernel Items:
      proper_parameter_def: IDENTIFIER ELLIPSIS.type_expr
      proper_parameter_def: IDENTIFIER ELLIPSIS., *
    Reduce:
      * -> [proper_parameter_def]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 427

  State 364:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 365:
    Kernel Items:
      proper_parameter_def: UNDERSCORE ELLIPSIS., *
      proper_parameter_def: UNDERSCORE ELLIPSIS.type_expr
    Reduce:
      * -> [proper_parameter_def]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 428

  State 366:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 367:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 429

  State 370:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER ASSIGN type_expr., *
    Reduce:
      * -> [type_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 371:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr., *
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr.IMPLEMENTS type_expr
    Reduce:
      * -> [type_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      IMPLEMENTS -> State 430
      binary_type_op -> State 256

  State 372:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [var_pattern]
      UNDERSCORE -> [var_pattern]
      var_pattern -> [field_var_pattern]
      tuple_pattern -> [var_pattern]
    Goto:
      LPAREN -> State 147

  State 373:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [var_pattern]
      ELLIPSIS -> [field_var_pattern]
      var_pattern -> [field_var_pattern]
      tuple_pattern -> [var_pattern]
      field_var_pattern -> [field_var_patterns]
    Goto:
      IDENTIFIER -> State 272
      LPAREN -> State 147

  State 376:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      tuple_pattern -> [case_pattern]
    Goto:
      LPAREN -> State 147

  State 379:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      return_type -> State 434

  State 380:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., *
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261

  State 382:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., *
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      DOT -> State 161

  State 383:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., *
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      DOT -> State 161

  State 384:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr.RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [map_type_expr]
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 385:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [array_type_expr]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA.field_def
      explicit_field_defs: proper_explicit_field_defs COMMA., *
    Reduce:
      * -> [explicit_field_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_field_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 392:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES.field_def
      explicit_field_defs: proper_explicit_field_defs NEWLINES., *
    Reduce:
      * -> [explicit_field_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_field_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 393:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [unsafe_statement]
    Goto:
      (nil)

  State 394:
    Kernel Items:
      proper_type_arguments: proper_type_arguments COMMA.type_expr
      type_arguments: proper_type_arguments COMMA., *
    Reduce:
      * -> [type_arguments]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 444

  State 397:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [call_expr]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      loop_expr_body: DO statement_block FOR.sequence_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [loop_expr_body]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 400:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN.sequence_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [for_assignment]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 401:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN.sequence_expr DO statement_block
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 448
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 402:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      * -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [optional_sequence_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      optional_sequence_expr -> State 449
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 403:
    Kernel Items:
      loop_expr_body: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [loop_expr_body]
    Goto:
      LBRACE -> State 11

  State 404:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    ShiftAndReduce:
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
    ShiftAndReduce:
      (nil)
    Goto:
      ELSE -> State 453

  State 408:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def NEWLINES.field_def
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_enum_value_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 409:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_enum_value_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 410:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES.field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES., *
    Reduce:
      * -> [explicit_enum_value_defs]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_enum_value_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 411:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [field_def]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      field_def -> [proper_explicit_enum_value_defs]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [field_def]
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
      prefix_unary_type_op -> State 134
      type_expr -> State 249

  State 412:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 413:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      return_type -> [func_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134

  State 414:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., *
    Reduce:
      * -> [parameter_decls]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      proper_parameter_def -> [parameter_decl]
      parameter_decl -> [proper_parameter_decls]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 335
      UNDERSCORE -> State 336
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      ELLIPSIS -> State 334
      prefix_unary_type_op -> State 134
      type_expr -> State 341

  State 416:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.parameter_decls RPAREN return_type
    Reduce:
      * -> [parameter_decls]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      proper_parameter_def -> [parameter_decl]
      parameter_decl -> [proper_parameter_decls]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 335
      UNDERSCORE -> State 336
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      ELLIPSIS -> State 334
      prefix_unary_type_op -> State 134
      type_expr -> State 341
      proper_parameter_decls -> State 339
      parameter_decls -> State 460

  State 423:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      generic_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [generic_parameter_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 425:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list COMMA.generic_parameter_def
      generic_parameter_def_list: proper_generic_parameter_def_list COMMA., *
    Reduce:
      * -> [generic_parameter_def_list]
    ShiftAndReduce:
      generic_parameter_def -> [proper_generic_parameter_def_list]
    Goto:
      IDENTIFIER -> State 358

  State 426:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 462

  State 427:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 428:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: UNDERSCORE ELLIPSIS type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 429:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 463

  State 430:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      type_expr -> State 464

  State 434:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [anonymous_func_expr]
    Goto:
      LBRACE -> State 11

  State 444:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_type_arguments: proper_type_arguments COMMA type_expr., *
    Reduce:
      * -> [proper_type_arguments]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 448:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      DO -> State 466

  State 449:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      SEMICOLON -> State 467

  State 452:
    Kernel Items:
      condition: CASE case_patterns ASSIGN.sequence_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [condition]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 453:
    Kernel Items:
      if_expr_body: IF condition statement_block ELSE.statement_block
      if_expr_body: IF condition statement_block ELSE.if_expr
    Reduce:
      * -> [optional_label_decl]
    ShiftAndReduce:
      LABEL_DECL -> [optional_label_decl]
      statement_block -> [if_expr_body]
      if_expr -> [if_expr_body]
    Goto:
      LBRACE -> State 11
      optional_label_decl -> State 470

  State 460:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 472

  State 462:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      return_type -> State 473

  State 463:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_defs -> State 286
      parameter_defs -> State 474

  State 464:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS type_expr., *
    Reduce:
      * -> [type_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 256

  State 466:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [loop_expr_body]
    Goto:
      LBRACE -> State 11

  State 467:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO statement_block
    Reduce:
      * -> [optional_label_decl]
      DO -> [optional_sequence_expr]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      VAR -> [var_or_let]
      LET -> [var_or_let]
      NOT -> [prefix_unary_op]
      LABEL_DECL -> [optional_label_decl]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_NEG -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      var_decl_pattern -> [sequence_expr]
      sequence_expr -> [optional_sequence_expr]
      call_expr -> [accessible_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      block_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      anonymous_func_expr -> [atom_expr]
    Goto:
      STRUCT -> State 50
      FUNC -> State 36
      LPAREN -> State 43
      LBRACKET -> State 42
      GREATER -> State 37
      var_or_let -> State 24
      optional_label_decl -> State 156
      optional_sequence_expr -> State 476
      accessible_expr -> State 108
      prefix_unary_op -> State 99
      mul_expr -> State 91
      add_expr -> State 57
      cmp_expr -> State 74
      and_expr -> State 58
      or_expr -> State 94
      initializable_type_expr -> State 84

  State 470:
    Kernel Items:
      if_expr: optional_label_decl.if_expr_body
    Reduce:
      (nil)
    ShiftAndReduce:
      if_expr_body -> [if_expr]
    Goto:
      IF -> State 229

  State 472:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      return_type -> [method_signature]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134

  State 473:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [named_func_def]
    Goto:
      LBRACE -> State 11

  State 474:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 479

  State 476:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      DO -> State 480

  State 479:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      parse_error_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 116
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      prefix_unary_type_op -> State 134
      return_type -> State 481

  State 480:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [loop_expr_body]
    Goto:
      LBRACE -> State 11

  State 481:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [named_func_def]
    Goto:
      LBRACE -> State 11

Number of states: 242
Number of shift actions: 1347
Number of reduce actions: 193
Number of shift-and-reduce actions: 3251
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 5377
Number of unoptimized shift actions: 46989
Number of unoptimized reduce actions: 35825
*/
