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
	ToCallExpr(AccessibleExpr_ Expression, GenericArguments_ *GenericArgumentList, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
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
	LocalToNamedTypeExpr(Identifier_ TokenValue, GenericArguments_ *GenericArgumentList) (TypeExpression, error)

	// 628:2: named_type_expr -> external: ...
	ExternalToNamedTypeExpr(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, GenericArguments_ *GenericArgumentList) (TypeExpression, error)
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
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, GenericParameters_ GenericSymbol, TypeExpr_ TypeExpression) (SourceDefinition, error)

	// 674:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, GenericParameters_ GenericSymbol, TypeExpr_ TypeExpression, Implements_ TokenValue, TypeExpr_2 TypeExpression) (SourceDefinition, error)

	// 675:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, TypeExpr_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterReducer interface {
	// 683:2: generic_parameter -> unconstrained: ...
	UnconstrainedToGenericParameter(Identifier_ TokenValue) (GenericSymbol, error)

	// 684:2: generic_parameter -> constrained: ...
	ConstrainedToGenericParameter(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type GenericParametersReducer interface {
	// 687:2: generic_parameters -> generic: ...
	GenericToGenericParameters(DollarLbracket_ TokenValue, GenericParameterList_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 688:2: generic_parameters -> nil: ...
	NilToGenericParameters() (GenericSymbol, error)
}

type ProperGenericParameterListReducer interface {
	// 691:2: proper_generic_parameter_list -> add: ...
	AddToProperGenericParameterList(ProperGenericParameterList_ GenericSymbol, Comma_ TokenValue, GenericParameter_ GenericSymbol) (GenericSymbol, error)

	// 692:2: proper_generic_parameter_list -> generic_parameter: ...
	GenericParameterToProperGenericParameterList(GenericParameter_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterListReducer interface {

	// 696:2: generic_parameter_list -> improper: ...
	ImproperToGenericParameterList(ProperGenericParameterList_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 697:2: generic_parameter_list -> nil: ...
	NilToGenericParameterList() (GenericSymbol, error)
}

type GenericArgumentsReducer interface {
	// 700:2: generic_arguments -> binding: ...
	BindingToGenericArguments(DollarLbracket_ TokenValue, GenericArgumentList_ *GenericArgumentList, Rbracket_ TokenValue) (*GenericArgumentList, error)

	// 701:2: generic_arguments -> nil: ...
	NilToGenericArguments() (*GenericArgumentList, error)
}

type ProperGenericArgumentListReducer interface {
	// 704:2: proper_generic_argument_list -> add: ...
	AddToProperGenericArgumentList(ProperGenericArgumentList_ *GenericArgumentList, Comma_ TokenValue, TypeExpr_ TypeExpression) (*GenericArgumentList, error)

	// 705:2: proper_generic_argument_list -> type_expr: ...
	TypeExprToProperGenericArgumentList(TypeExpr_ TypeExpression) (*GenericArgumentList, error)
}

type GenericArgumentListReducer interface {

	// 709:2: generic_argument_list -> improper: ...
	ImproperToGenericArgumentList(ProperGenericArgumentList_ *GenericArgumentList, Comma_ TokenValue) (*GenericArgumentList, error)

	// 710:2: generic_argument_list -> nil: ...
	NilToGenericArgumentList() (*GenericArgumentList, error)
}

type FieldDefReducer interface {
	// 717:2: field_def -> named: ...
	NamedToFieldDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (*FieldDef, error)

	// 718:2: field_def -> unnamed: ...
	UnnamedToFieldDef(TypeExpr_ TypeExpression) (*FieldDef, error)
}

type UnsafeStatementPropertyReducer interface {
	// 720:43: unsafe_statement_property -> ...
	ToUnsafeStatementProperty(UnsafeStatement_ Statement) (TypeProperty, error)
}

type TypePropertyReducer interface {

	// 731:2: type_property -> default_enum_field_def: ...
	DefaultEnumFieldDefToTypeProperty(Default_ TokenValue, FieldDef_ *FieldDef) (TypeProperty, error)

	// 732:2: type_property -> padding_field_def: ...
	PaddingFieldDefToTypeProperty(Underscore_ TokenValue, TypeExpr_ TypeExpression) (TypeProperty, error)
}

type ProperImplicitTypePropertiesReducer interface {
	// 737:2: proper_implicit_type_properties -> add: ...
	AddToProperImplicitTypeProperties(ProperImplicitTypeProperties_ *TypePropertyList, Comma_ TokenValue, TypeProperty_ TypeProperty) (*TypePropertyList, error)

	// 738:2: proper_implicit_type_properties -> type_property: ...
	TypePropertyToProperImplicitTypeProperties(TypeProperty_ TypeProperty) (*TypePropertyList, error)
}

type ImplicitTypePropertiesReducer interface {

	// 742:2: implicit_type_properties -> improper: ...
	ImproperToImplicitTypeProperties(ProperImplicitTypeProperties_ *TypePropertyList, Comma_ TokenValue) (*TypePropertyList, error)

	// 743:2: implicit_type_properties -> nil: ...
	NilToImplicitTypeProperties() (*TypePropertyList, error)
}

type ImplicitStructTypeExprReducer interface {
	// 746:2: implicit_struct_type_expr -> ...
	ToImplicitStructTypeExpr(Lparen_ TokenValue, ImplicitTypeProperties_ *TypePropertyList, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitTypePropertiesReducer interface {
	// 749:2: proper_explicit_type_properties -> add_implicit: ...
	AddImplicitToProperExplicitTypeProperties(ProperExplicitTypeProperties_ *TypePropertyList, Newlines_ TokenCount, TypeProperty_ TypeProperty) (*TypePropertyList, error)

	// 750:2: proper_explicit_type_properties -> add_explicit: ...
	AddExplicitToProperExplicitTypeProperties(ProperExplicitTypeProperties_ *TypePropertyList, Comma_ TokenValue, TypeProperty_ TypeProperty) (*TypePropertyList, error)

	// 751:2: proper_explicit_type_properties -> type_property: ...
	TypePropertyToProperExplicitTypeProperties(TypeProperty_ TypeProperty) (*TypePropertyList, error)
}

type ExplicitTypePropertiesReducer interface {

	// 755:2: explicit_type_properties -> improper_implicit: ...
	ImproperImplicitToExplicitTypeProperties(ProperExplicitTypeProperties_ *TypePropertyList, Newlines_ TokenCount) (*TypePropertyList, error)

	// 756:2: explicit_type_properties -> improper_explicit: ...
	ImproperExplicitToExplicitTypeProperties(ProperExplicitTypeProperties_ *TypePropertyList, Comma_ TokenValue) (*TypePropertyList, error)

	// 757:2: explicit_type_properties -> nil: ...
	NilToExplicitTypeProperties() (*TypePropertyList, error)
}

type ExplicitStructTypeExprReducer interface {
	// 760:2: explicit_struct_type_expr -> ...
	ToExplicitStructTypeExpr(Struct_ TokenValue, Lparen_ TokenValue, ExplicitTypeProperties_ *TypePropertyList, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitTypeExprReducer interface {
	// 763:2: trait_type_expr -> ...
	ToTraitTypeExpr(Trait_ TokenValue, Lparen_ TokenValue, ExplicitTypeProperties_ *TypePropertyList, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperImplicitEnumTypePropertiesReducer interface {
	// 774:2: proper_implicit_enum_type_properties -> pair: ...
	PairToProperImplicitEnumTypeProperties(TypeProperty_ TypeProperty, Or_ TokenValue, TypeProperty_2 TypeProperty) (*TypePropertyList, error)

	// 775:2: proper_implicit_enum_type_properties -> add: ...
	AddToProperImplicitEnumTypeProperties(ProperImplicitEnumTypeProperties_ *TypePropertyList, Or_ TokenValue, TypeProperty_ TypeProperty) (*TypePropertyList, error)
}

type ImplicitEnumTypePropertiesReducer interface {

	// 780:2: implicit_enum_type_properties -> improper: ...
	ImproperToImplicitEnumTypeProperties(ProperImplicitEnumTypeProperties_ *TypePropertyList, Newlines_ TokenCount) (*TypePropertyList, error)
}

type ImplicitEnumTypeExprReducer interface {
	// 783:2: implicit_enum_type_expr -> ...
	ToImplicitEnumTypeExpr(Lparen_ TokenValue, ImplicitEnumTypeProperties_ *TypePropertyList, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitEnumTypePropertiesReducer interface {
	// 786:2: proper_explicit_enum_type_properties -> explicit_pair: ...
	ExplicitPairToProperExplicitEnumTypeProperties(TypeProperty_ TypeProperty, Or_ TokenValue, TypeProperty_2 TypeProperty) (*TypePropertyList, error)

	// 787:2: proper_explicit_enum_type_properties -> implicit_pair: ...
	ImplicitPairToProperExplicitEnumTypeProperties(TypeProperty_ TypeProperty, Newlines_ TokenCount, TypeProperty_2 TypeProperty) (*TypePropertyList, error)

	// 788:2: proper_explicit_enum_type_properties -> explicit_add: ...
	ExplicitAddToProperExplicitEnumTypeProperties(ProperExplicitEnumTypeProperties_ *TypePropertyList, Or_ TokenValue, TypeProperty_ TypeProperty) (*TypePropertyList, error)

	// 789:2: proper_explicit_enum_type_properties -> implicit_add: ...
	ImplicitAddToProperExplicitEnumTypeProperties(ProperExplicitEnumTypeProperties_ *TypePropertyList, Newlines_ TokenCount, TypeProperty_ TypeProperty) (*TypePropertyList, error)
}

type ExplicitEnumTypePropertiesReducer interface {

	// 794:2: explicit_enum_type_properties -> improper: ...
	ImproperToExplicitEnumTypeProperties(ProperExplicitEnumTypeProperties_ *TypePropertyList, Newlines_ TokenCount) (*TypePropertyList, error)
}

type ExplicitEnumTypeExprReducer interface {
	// 797:2: explicit_enum_type_expr -> ...
	ToExplicitEnumTypeExpr(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumTypeProperties_ *TypePropertyList, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 805:2: return_type -> returnable_type_expr: ...
	ReturnableTypeExprToReturnType(ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)

	// 806:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ProperParameterDefReducer interface {
	// 809:2: proper_parameter_def -> named_typed_arg: ...
	NamedTypedArgToProperParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 810:2: proper_parameter_def -> named_typed_vararg: ...
	NamedTypedVarargToProperParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 811:2: proper_parameter_def -> named_inferred_vararg: ...
	NamedInferredVarargToProperParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (*Parameter, error)

	// 812:2: proper_parameter_def -> ignore_typed_arg: ...
	IgnoreTypedArgToProperParameterDef(Underscore_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 813:2: proper_parameter_def -> ignore_inferred_vararg: ...
	IgnoreInferredVarargToProperParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue) (*Parameter, error)

	// 814:2: proper_parameter_def -> ignore_typed_vararg: ...
	IgnoreTypedVarargToProperParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)
}

type ParameterDeclReducer interface {

	// 820:2: parameter_decl -> unnamed_typed_arg: ...
	UnnamedTypedArgToParameterDecl(TypeExpr_ TypeExpression) (*Parameter, error)

	// 821:2: parameter_decl -> unnamed_inferred_vararg: ...
	UnnamedInferredVarargToParameterDecl(Ellipsis_ TokenValue) (*Parameter, error)

	// 822:2: parameter_decl -> unnamed_typed_vararg: ...
	UnnamedTypedVarargToParameterDecl(Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)
}

type ParameterDefReducer interface {

	// 830:2: parameter_def -> named_inferred_arg: ...
	NamedInferredArgToParameterDef(Identifier_ TokenValue) (*Parameter, error)

	// 831:2: parameter_def -> ignore_inferred_arg: ...
	IgnoreInferredArgToParameterDef(Underscore_ TokenValue) (*Parameter, error)
}

type ProperParameterDeclsReducer interface {
	// 834:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ *ParameterList, Comma_ TokenValue, ParameterDecl_ *Parameter) (*ParameterList, error)

	// 835:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ *Parameter) (*ParameterList, error)
}

type ParameterDeclsReducer interface {

	// 839:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ *ParameterList, Comma_ TokenValue) (*ParameterList, error)

	// 840:2: parameter_decls -> nil: ...
	NilToParameterDecls() (*ParameterList, error)
}

type ProperParameterDefsReducer interface {
	// 843:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ *ParameterList, Comma_ TokenValue, ParameterDef_ *Parameter) (*ParameterList, error)

	// 844:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ *Parameter) (*ParameterList, error)
}

type ParameterDefsReducer interface {

	// 848:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ *ParameterList, Comma_ TokenValue) (*ParameterList, error)

	// 849:2: parameter_defs -> nil: ...
	NilToParameterDefs() (*ParameterList, error)
}

type FuncTypeExprReducer interface {
	// 852:2: func_type_expr -> ...
	ToFuncTypeExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 864:2: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeProperty, error)
}

type NamedFuncDefReducer interface {
	// 868:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, GenericParameters_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 869:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ *Parameter, Rparen_ TokenValue, Identifier_ TokenValue, Lparen_2 TokenValue, ParameterDefs_ *ParameterList, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 870:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 873:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 885:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 886:2: package_def -> with_spec: ...
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
	GenericParameterReducer
	GenericParametersReducer
	ProperGenericParameterListReducer
	GenericParameterListReducer
	GenericArgumentsReducer
	ProperGenericArgumentListReducer
	GenericArgumentListReducer
	FieldDefReducer
	UnsafeStatementPropertyReducer
	TypePropertyReducer
	ProperImplicitTypePropertiesReducer
	ImplicitTypePropertiesReducer
	ImplicitStructTypeExprReducer
	ProperExplicitTypePropertiesReducer
	ExplicitTypePropertiesReducer
	ExplicitStructTypeExprReducer
	TraitTypeExprReducer
	ProperImplicitEnumTypePropertiesReducer
	ImplicitEnumTypePropertiesReducer
	ImplicitEnumTypeExprReducer
	ProperExplicitEnumTypePropertiesReducer
	ExplicitEnumTypePropertiesReducer
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
	case _State9:
		return []SymbolId{IdentifierToken, LparenToken}
	case _State12:
		return []SymbolId{IdentifierToken}
	case _State15:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken}
	case _State17:
		return []SymbolId{ColonToken}
	case _State18:
		return []SymbolId{LparenToken}
	case _State20:
		return []SymbolId{StringLiteralToken, IdentifierToken, UnderscoreToken, LparenToken, DotToken}
	case _State21:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State23:
		return []SymbolId{LparenToken}
	case _State24:
		return []SymbolId{LessToken}
	case _State28:
		return []SymbolId{AssignToken}
	case _State32:
		return []SymbolId{LparenToken}
	case _State35:
		return []SymbolId{IfToken, SwitchToken, ForToken, DoToken, LbraceToken}
	case _State40:
		return []SymbolId{LparenToken}
	case _State41:
		return []SymbolId{LparenToken}
	case _State44:
		return []SymbolId{LparenToken}
	case _State45:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State47:
		return []SymbolId{IdentifierToken, UnderscoreToken}
	case _State49:
		return []SymbolId{RbraceToken}
	case _State53:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken, EllipsisToken}
	case _State55:
		return []SymbolId{IdentifierToken}
	case _State57:
		return []SymbolId{CommaToken, ColonToken}
	case _State58:
		return []SymbolId{LbraceToken}
	case _State61:
		return []SymbolId{StringLiteralToken}
	case _State62:
		return []SymbolId{StringLiteralToken}
	case _State63:
		return []SymbolId{StringLiteralToken, IdentifierToken, UnderscoreToken, DotToken}
	case _State64:
		return []SymbolId{StringLiteralToken}
	case _State65:
		return []SymbolId{RbracketToken, CommaToken, ColonToken, AddToken, SubToken, MulToken}
	case _State68:
		return []SymbolId{RparenToken}
	case _State73:
		return []SymbolId{IdentifierToken}
	case _State75:
		return []SymbolId{IdentifierToken}
	case _State78:
		return []SymbolId{LparenToken}
	case _State90:
		return []SymbolId{LbraceToken}
	case _State95:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State97:
		return []SymbolId{IdentifierToken}
	case _State98:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State99:
		return []SymbolId{IdentifierToken, LparenToken}
	case _State102:
		return []SymbolId{RparenToken}
	case _State103:
		return []SymbolId{RparenToken}
	case _State109:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State112:
		return []SymbolId{LparenToken}
	case _State115:
		return []SymbolId{RparenToken}
	case _State118:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State119:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State121:
		return []SymbolId{RparenToken, CommaToken}
	case _State124:
		return []SymbolId{IdentifierToken}
	case _State127:
		return []SymbolId{RparenToken}
	case _State129:
		return []SymbolId{RparenToken}
	case _State131:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State132:
		return []SymbolId{IntegerLiteralToken}
	case _State137:
		return []SymbolId{RparenToken}
	case _State139:
		return []SymbolId{GreaterToken}
	case _State140:
		return []SymbolId{RbracketToken}
	case _State143:
		return []SymbolId{RbracketToken}
	case _State148:
		return []SymbolId{RparenToken}
	case _State151:
		return []SymbolId{InToken, AssignToken}
	case _State152:
		return []SymbolId{SemicolonToken}
	case _State155:
		return []SymbolId{LbraceToken}
	case _State156:
		return []SymbolId{LbraceToken}
	case _State158:
		return []SymbolId{RparenToken}
	case _State160:
		return []SymbolId{NewlinesToken, OrToken}
	case _State164:
		return []SymbolId{RparenToken}
	case _State168:
		return []SymbolId{LparenToken}
	case _State172:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State174:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State175:
		return []SymbolId{RparenToken}
	case _State177:
		return []SymbolId{RbracketToken}
	case _State184:
		return []SymbolId{IdentifierToken}
	case _State187:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken}
	case _State188:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken, EllipsisToken}
	case _State189:
		return []SymbolId{LparenToken}
	case _State194:
		return []SymbolId{RbracketToken, AddToken, SubToken, MulToken}
	case _State195:
		return []SymbolId{RbracketToken}
	case _State198:
		return []SymbolId{StringLiteralToken}
	case _State200:
		return []SymbolId{RparenToken}
	case _State205:
		return []SymbolId{LbraceToken}
	case _State206:
		return []SymbolId{CommaToken, AssignToken}
	case _State209:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State210:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State211:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, UnsafeToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State218:
		return []SymbolId{RparenToken}
	case _State221:
		return []SymbolId{LparenToken}
	case _State222:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeTildeToken, BitNegToken, BitAndToken, ParseErrorToken}
	case _State223:
		return []SymbolId{LbraceToken}
	case _State225:
		return []SymbolId{DoToken}
	case _State226:
		return []SymbolId{SemicolonToken}
	case _State229:
		return []SymbolId{RparenToken}
	case _State233:
		return []SymbolId{LbraceToken}
	case _State235:
		return []SymbolId{IfToken}
	case _State237:
		return []SymbolId{LbraceToken}
	case _State238:
		return []SymbolId{RparenToken}
	case _State239:
		return []SymbolId{DoToken}
	case _State241:
		return []SymbolId{LbraceToken}
	case _State242:
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
	case GenericParameterType:
		return "generic_parameter"
	case GenericParametersType:
		return "generic_parameters"
	case ProperGenericParameterListType:
		return "proper_generic_parameter_list"
	case GenericParameterListType:
		return "generic_parameter_list"
	case GenericArgumentsType:
		return "generic_arguments"
	case ProperGenericArgumentListType:
		return "proper_generic_argument_list"
	case GenericArgumentListType:
		return "generic_argument_list"
	case FieldDefType:
		return "field_def"
	case UnsafeStatementPropertyType:
		return "unsafe_statement_property"
	case TypePropertyType:
		return "type_property"
	case ProperImplicitTypePropertiesType:
		return "proper_implicit_type_properties"
	case ImplicitTypePropertiesType:
		return "implicit_type_properties"
	case ImplicitStructTypeExprType:
		return "implicit_struct_type_expr"
	case ProperExplicitTypePropertiesType:
		return "proper_explicit_type_properties"
	case ExplicitTypePropertiesType:
		return "explicit_type_properties"
	case ExplicitStructTypeExprType:
		return "explicit_struct_type_expr"
	case TraitTypeExprType:
		return "trait_type_expr"
	case ProperImplicitEnumTypePropertiesType:
		return "proper_implicit_enum_type_properties"
	case ImplicitEnumTypePropertiesType:
		return "implicit_enum_type_properties"
	case ImplicitEnumTypeExprType:
		return "implicit_enum_type_expr"
	case ProperExplicitEnumTypePropertiesType:
		return "proper_explicit_enum_type_properties"
	case ExplicitEnumTypePropertiesType:
		return "explicit_enum_type_properties"
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

	SourceType                           = SymbolId(343)
	ProperDefinitionsType                = SymbolId(344)
	DefinitionsType                      = SymbolId(345)
	DefinitionType                       = SymbolId(346)
	StatementBlockType                   = SymbolId(347)
	ProperStatementsType                 = SymbolId(348)
	StatementsType                       = SymbolId(349)
	SimpleStatementType                  = SymbolId(350)
	StatementType                        = SymbolId(351)
	OptionalSimpleStatementType          = SymbolId(352)
	ExprOrImproperStructStatementType    = SymbolId(353)
	ExprsType                            = SymbolId(354)
	CallbackOpType                       = SymbolId(355)
	CallbackOpStatementType              = SymbolId(356)
	UnsafeStatementType                  = SymbolId(357)
	JumpStatementType                    = SymbolId(358)
	JumpTypeType                         = SymbolId(359)
	FallthroughStatementType             = SymbolId(360)
	AssignStatementType                  = SymbolId(361)
	UnaryOpAssignStatementType           = SymbolId(362)
	UnaryOpAssignType                    = SymbolId(363)
	BinaryOpAssignStatementType          = SymbolId(364)
	BinaryOpAssignType                   = SymbolId(365)
	ImportStatementType                  = SymbolId(366)
	ImportClauseType                     = SymbolId(367)
	ProperImportClausesType              = SymbolId(368)
	ImportClausesType                    = SymbolId(369)
	CasePatternsType                     = SymbolId(370)
	VarDeclPatternType                   = SymbolId(371)
	VarOrLetType                         = SymbolId(372)
	VarPatternType                       = SymbolId(373)
	TuplePatternType                     = SymbolId(374)
	FieldVarPatternsType                 = SymbolId(375)
	FieldVarPatternType                  = SymbolId(376)
	OptionalTypeExprType                 = SymbolId(377)
	AssignPatternType                    = SymbolId(378)
	CasePatternType                      = SymbolId(379)
	ExprType                             = SymbolId(380)
	OptionalLabelDeclType                = SymbolId(381)
	SequenceExprType                     = SymbolId(382)
	IfExprType                           = SymbolId(383)
	IfExprBodyType                       = SymbolId(384)
	ConditionType                        = SymbolId(385)
	SwitchExprType                       = SymbolId(386)
	LoopExprType                         = SymbolId(387)
	LoopExprBodyType                     = SymbolId(388)
	OptionalSequenceExprType             = SymbolId(389)
	ForAssignmentType                    = SymbolId(390)
	CallExprType                         = SymbolId(391)
	ProperArgumentsType                  = SymbolId(392)
	ArgumentsType                        = SymbolId(393)
	ArgumentType                         = SymbolId(394)
	ColonExprType                        = SymbolId(395)
	AtomExprType                         = SymbolId(396)
	ParseErrorExprType                   = SymbolId(397)
	LiteralExprType                      = SymbolId(398)
	NamedExprType                        = SymbolId(399)
	BlockExprType                        = SymbolId(400)
	InitializeExprType                   = SymbolId(401)
	ImplicitStructExprType               = SymbolId(402)
	AccessibleExprType                   = SymbolId(403)
	AccessExprType                       = SymbolId(404)
	IndexExprType                        = SymbolId(405)
	PostfixableExprType                  = SymbolId(406)
	PostfixUnaryOpType                   = SymbolId(407)
	PostfixUnaryExprType                 = SymbolId(408)
	PrefixableExprType                   = SymbolId(409)
	PrefixUnaryExprType                  = SymbolId(410)
	PrefixUnaryOpType                    = SymbolId(411)
	MulExprType                          = SymbolId(412)
	BinaryMulExprType                    = SymbolId(413)
	MulOpType                            = SymbolId(414)
	AddExprType                          = SymbolId(415)
	BinaryAddExprType                    = SymbolId(416)
	AddOpType                            = SymbolId(417)
	CmpExprType                          = SymbolId(418)
	BinaryCmpExprType                    = SymbolId(419)
	CmpOpType                            = SymbolId(420)
	AndExprType                          = SymbolId(421)
	BinaryAndExprType                    = SymbolId(422)
	OrExprType                           = SymbolId(423)
	BinaryOrExprType                     = SymbolId(424)
	InitializableTypeExprType            = SymbolId(425)
	SliceTypeExprType                    = SymbolId(426)
	ArrayTypeExprType                    = SymbolId(427)
	MapTypeExprType                      = SymbolId(428)
	AtomTypeExprType                     = SymbolId(429)
	NamedTypeExprType                    = SymbolId(430)
	InferredTypeExprType                 = SymbolId(431)
	ParseErrorTypeExprType               = SymbolId(432)
	ReturnableTypeExprType               = SymbolId(433)
	PrefixUnaryTypeExprType              = SymbolId(434)
	PrefixUnaryTypeOpType                = SymbolId(435)
	TypeExprType                         = SymbolId(436)
	BinaryTypeExprType                   = SymbolId(437)
	BinaryTypeOpType                     = SymbolId(438)
	TypeDefType                          = SymbolId(439)
	GenericParameterType                 = SymbolId(440)
	GenericParametersType                = SymbolId(441)
	ProperGenericParameterListType       = SymbolId(442)
	GenericParameterListType             = SymbolId(443)
	GenericArgumentsType                 = SymbolId(444)
	ProperGenericArgumentListType        = SymbolId(445)
	GenericArgumentListType              = SymbolId(446)
	FieldDefType                         = SymbolId(447)
	UnsafeStatementPropertyType          = SymbolId(448)
	TypePropertyType                     = SymbolId(449)
	ProperImplicitTypePropertiesType     = SymbolId(450)
	ImplicitTypePropertiesType           = SymbolId(451)
	ImplicitStructTypeExprType           = SymbolId(452)
	ProperExplicitTypePropertiesType     = SymbolId(453)
	ExplicitTypePropertiesType           = SymbolId(454)
	ExplicitStructTypeExprType           = SymbolId(455)
	TraitTypeExprType                    = SymbolId(456)
	ProperImplicitEnumTypePropertiesType = SymbolId(457)
	ImplicitEnumTypePropertiesType       = SymbolId(458)
	ImplicitEnumTypeExprType             = SymbolId(459)
	ProperExplicitEnumTypePropertiesType = SymbolId(460)
	ExplicitEnumTypePropertiesType       = SymbolId(461)
	ExplicitEnumTypeExprType             = SymbolId(462)
	ReturnTypeType                       = SymbolId(463)
	ProperParameterDefType               = SymbolId(464)
	ParameterDeclType                    = SymbolId(465)
	ParameterDefType                     = SymbolId(466)
	ProperParameterDeclsType             = SymbolId(467)
	ParameterDeclsType                   = SymbolId(468)
	ProperParameterDefsType              = SymbolId(469)
	ParameterDefsType                    = SymbolId(470)
	FuncTypeExprType                     = SymbolId(471)
	MethodSignatureType                  = SymbolId(472)
	NamedFuncDefType                     = SymbolId(473)
	AnonymousFuncExprType                = SymbolId(474)
	PackageDefType                       = SymbolId(475)
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
	_ReduceToSource                                                     = _ReduceType(1)
	_ReduceAddToProperDefinitions                                       = _ReduceType(2)
	_ReduceDefinitionToProperDefinitions                                = _ReduceType(3)
	_ReduceProperDefinitionsToDefinitions                               = _ReduceType(4)
	_ReduceImproperToDefinitions                                        = _ReduceType(5)
	_ReduceNilToDefinitions                                             = _ReduceType(6)
	_ReducePackageDefToDefinition                                       = _ReduceType(7)
	_ReduceTypeDefToDefinition                                          = _ReduceType(8)
	_ReduceNamedFuncDefToDefinition                                     = _ReduceType(9)
	_ReduceGlobalVarDefToDefinition                                     = _ReduceType(10)
	_ReduceGlobalVarAssignmentToDefinition                              = _ReduceType(11)
	_ReduceStatementBlockToDefinition                                   = _ReduceType(12)
	_ReduceCommentGroupsToDefinition                                    = _ReduceType(13)
	_ReduceToStatementBlock                                             = _ReduceType(14)
	_ReduceAddImplicitToProperStatements                                = _ReduceType(15)
	_ReduceAddExplicitToProperStatements                                = _ReduceType(16)
	_ReduceStatementToProperStatements                                  = _ReduceType(17)
	_ReduceProperStatementsToStatements                                 = _ReduceType(18)
	_ReduceImproperImplicitToStatements                                 = _ReduceType(19)
	_ReduceImproperExplicitToStatements                                 = _ReduceType(20)
	_ReduceNilToStatements                                              = _ReduceType(21)
	_ReduceUnsafeStatementToSimpleStatement                             = _ReduceType(22)
	_ReduceExprOrImproperStructStatementToSimpleStatement               = _ReduceType(23)
	_ReduceCallbackOpStatementToSimpleStatement                         = _ReduceType(24)
	_ReduceJumpStatementToSimpleStatement                               = _ReduceType(25)
	_ReduceAssignStatementToSimpleStatement                             = _ReduceType(26)
	_ReduceUnaryOpAssignStatementToSimpleStatement                      = _ReduceType(27)
	_ReduceBinaryOpAssignStatementToSimpleStatement                     = _ReduceType(28)
	_ReduceFallthroughStatementToSimpleStatement                        = _ReduceType(29)
	_ReduceSimpleStatementToStatement                                   = _ReduceType(30)
	_ReduceImportStatementToStatement                                   = _ReduceType(31)
	_ReduceCaseBranchStatementToStatement                               = _ReduceType(32)
	_ReduceDefaultBranchStatementToStatement                            = _ReduceType(33)
	_ReduceSimpleStatementToOptionalSimpleStatement                     = _ReduceType(34)
	_ReduceNilToOptionalSimpleStatement                                 = _ReduceType(35)
	_ReduceToExprOrImproperStructStatement                              = _ReduceType(36)
	_ReduceExprToExprs                                                  = _ReduceType(37)
	_ReduceAddToExprs                                                   = _ReduceType(38)
	_ReduceAsyncToCallbackOp                                            = _ReduceType(39)
	_ReduceDeferToCallbackOp                                            = _ReduceType(40)
	_ReduceToCallbackOpStatement                                        = _ReduceType(41)
	_ReduceToUnsafeStatement                                            = _ReduceType(42)
	_ReduceUnlabeledNoValueToJumpStatement                              = _ReduceType(43)
	_ReduceUnlabeledValuedToJumpStatement                               = _ReduceType(44)
	_ReduceLabeledNoValueToJumpStatement                                = _ReduceType(45)
	_ReduceLabeledValuedToJumpStatement                                 = _ReduceType(46)
	_ReduceReturnToJumpType                                             = _ReduceType(47)
	_ReduceBreakToJumpType                                              = _ReduceType(48)
	_ReduceContinueToJumpType                                           = _ReduceType(49)
	_ReduceToFallthroughStatement                                       = _ReduceType(50)
	_ReduceToAssignStatement                                            = _ReduceType(51)
	_ReduceToUnaryOpAssignStatement                                     = _ReduceType(52)
	_ReduceAddOneAssignToUnaryOpAssign                                  = _ReduceType(53)
	_ReduceSubOneAssignToUnaryOpAssign                                  = _ReduceType(54)
	_ReduceToBinaryOpAssignStatement                                    = _ReduceType(55)
	_ReduceAddAssignToBinaryOpAssign                                    = _ReduceType(56)
	_ReduceSubAssignToBinaryOpAssign                                    = _ReduceType(57)
	_ReduceMulAssignToBinaryOpAssign                                    = _ReduceType(58)
	_ReduceDivAssignToBinaryOpAssign                                    = _ReduceType(59)
	_ReduceModAssignToBinaryOpAssign                                    = _ReduceType(60)
	_ReduceBitNegAssignToBinaryOpAssign                                 = _ReduceType(61)
	_ReduceBitAndAssignToBinaryOpAssign                                 = _ReduceType(62)
	_ReduceBitOrAssignToBinaryOpAssign                                  = _ReduceType(63)
	_ReduceBitXorAssignToBinaryOpAssign                                 = _ReduceType(64)
	_ReduceBitLshiftAssignToBinaryOpAssign                              = _ReduceType(65)
	_ReduceBitRshiftAssignToBinaryOpAssign                              = _ReduceType(66)
	_ReduceSingleToImportStatement                                      = _ReduceType(67)
	_ReduceMultipleToImportStatement                                    = _ReduceType(68)
	_ReduceStringLiteralToImportClause                                  = _ReduceType(69)
	_ReduceAliasToImportClause                                          = _ReduceType(70)
	_ReduceUnusableImportToImportClause                                 = _ReduceType(71)
	_ReduceImportToLocalToImportClause                                  = _ReduceType(72)
	_ReduceAddImplicitToProperImportClauses                             = _ReduceType(73)
	_ReduceAddExplicitToProperImportClauses                             = _ReduceType(74)
	_ReduceImportClauseToProperImportClauses                            = _ReduceType(75)
	_ReduceProperImportClausesToImportClauses                           = _ReduceType(76)
	_ReduceImplicitToImportClauses                                      = _ReduceType(77)
	_ReduceExplicitToImportClauses                                      = _ReduceType(78)
	_ReduceCasePatternToCasePatterns                                    = _ReduceType(79)
	_ReduceMultipleToCasePatterns                                       = _ReduceType(80)
	_ReduceToVarDeclPattern                                             = _ReduceType(81)
	_ReduceVarToVarOrLet                                                = _ReduceType(82)
	_ReduceLetToVarOrLet                                                = _ReduceType(83)
	_ReduceIdentifierToVarPattern                                       = _ReduceType(84)
	_ReduceUnderscoreToVarPattern                                       = _ReduceType(85)
	_ReduceTuplePatternToVarPattern                                     = _ReduceType(86)
	_ReduceToTuplePattern                                               = _ReduceType(87)
	_ReduceFieldVarPatternToFieldVarPatterns                            = _ReduceType(88)
	_ReduceAddToFieldVarPatterns                                        = _ReduceType(89)
	_ReducePositionalToFieldVarPattern                                  = _ReduceType(90)
	_ReduceNamedToFieldVarPattern                                       = _ReduceType(91)
	_ReduceEllipsisToFieldVarPattern                                    = _ReduceType(92)
	_ReduceTypeExprToOptionalTypeExpr                                   = _ReduceType(93)
	_ReduceNilToOptionalTypeExpr                                        = _ReduceType(94)
	_ReduceToAssignPattern                                              = _ReduceType(95)
	_ReduceAssignPatternToCasePattern                                   = _ReduceType(96)
	_ReduceEnumMatchPatternToCasePattern                                = _ReduceType(97)
	_ReduceEnumNondataMatchPattenToCasePattern                          = _ReduceType(98)
	_ReduceEnumVarDeclPatternToCasePattern                              = _ReduceType(99)
	_ReduceIfExprToExpr                                                 = _ReduceType(100)
	_ReduceSwitchExprToExpr                                             = _ReduceType(101)
	_ReduceLoopExprToExpr                                               = _ReduceType(102)
	_ReduceSequenceExprToExpr                                           = _ReduceType(103)
	_ReduceLabelDeclToOptionalLabelDecl                                 = _ReduceType(104)
	_ReduceUnlabelledToOptionalLabelDecl                                = _ReduceType(105)
	_ReduceOrExprToSequenceExpr                                         = _ReduceType(106)
	_ReduceVarDeclPatternToSequenceExpr                                 = _ReduceType(107)
	_ReduceAssignVarPatternToSequenceExpr                               = _ReduceType(108)
	_ReduceToIfExpr                                                     = _ReduceType(109)
	_ReduceNoElseToIfExprBody                                           = _ReduceType(110)
	_ReduceIfElseToIfExprBody                                           = _ReduceType(111)
	_ReduceMultiIfElseToIfExprBody                                      = _ReduceType(112)
	_ReduceSequenceExprToCondition                                      = _ReduceType(113)
	_ReduceCaseToCondition                                              = _ReduceType(114)
	_ReduceToSwitchExpr                                                 = _ReduceType(115)
	_ReduceToLoopExpr                                                   = _ReduceType(116)
	_ReduceInfiniteToLoopExprBody                                       = _ReduceType(117)
	_ReduceDoWhileToLoopExprBody                                        = _ReduceType(118)
	_ReduceWhileToLoopExprBody                                          = _ReduceType(119)
	_ReduceIteratorToLoopExprBody                                       = _ReduceType(120)
	_ReduceForToLoopExprBody                                            = _ReduceType(121)
	_ReduceSequenceExprToOptionalSequenceExpr                           = _ReduceType(122)
	_ReduceNilToOptionalSequenceExpr                                    = _ReduceType(123)
	_ReduceSequenceExprToForAssignment                                  = _ReduceType(124)
	_ReduceAssignToForAssignment                                        = _ReduceType(125)
	_ReduceToCallExpr                                                   = _ReduceType(126)
	_ReduceAddToProperArguments                                         = _ReduceType(127)
	_ReduceArgumentToProperArguments                                    = _ReduceType(128)
	_ReduceProperArgumentsToArguments                                   = _ReduceType(129)
	_ReduceImproperToArguments                                          = _ReduceType(130)
	_ReduceNilToArguments                                               = _ReduceType(131)
	_ReducePositionalToArgument                                         = _ReduceType(132)
	_ReduceColonExprToArgument                                          = _ReduceType(133)
	_ReduceNamedAssignmentToArgument                                    = _ReduceType(134)
	_ReduceVarargAssignmentToArgument                                   = _ReduceType(135)
	_ReduceSkipPatternToArgument                                        = _ReduceType(136)
	_ReduceUnitUnitPairToColonExpr                                      = _ReduceType(137)
	_ReduceExprUnitPairToColonExpr                                      = _ReduceType(138)
	_ReduceUnitExprPairToColonExpr                                      = _ReduceType(139)
	_ReduceExprExprPairToColonExpr                                      = _ReduceType(140)
	_ReduceColonExprUnitTupleToColonExpr                                = _ReduceType(141)
	_ReduceColonExprExprTupleToColonExpr                                = _ReduceType(142)
	_ReduceParseErrorExprToAtomExpr                                     = _ReduceType(143)
	_ReduceLiteralExprToAtomExpr                                        = _ReduceType(144)
	_ReduceNamedExprToAtomExpr                                          = _ReduceType(145)
	_ReduceBlockExprToAtomExpr                                          = _ReduceType(146)
	_ReduceAnonymousFuncExprToAtomExpr                                  = _ReduceType(147)
	_ReduceInitializeExprToAtomExpr                                     = _ReduceType(148)
	_ReduceImplicitStructExprToAtomExpr                                 = _ReduceType(149)
	_ReduceToParseErrorExpr                                             = _ReduceType(150)
	_ReduceTrueToLiteralExpr                                            = _ReduceType(151)
	_ReduceFalseToLiteralExpr                                           = _ReduceType(152)
	_ReduceIntegerLiteralToLiteralExpr                                  = _ReduceType(153)
	_ReduceFloatLiteralToLiteralExpr                                    = _ReduceType(154)
	_ReduceRuneLiteralToLiteralExpr                                     = _ReduceType(155)
	_ReduceStringLiteralToLiteralExpr                                   = _ReduceType(156)
	_ReduceIdentifierToNamedExpr                                        = _ReduceType(157)
	_ReduceUnderscoreToNamedExpr                                        = _ReduceType(158)
	_ReduceToBlockExpr                                                  = _ReduceType(159)
	_ReduceToInitializeExpr                                             = _ReduceType(160)
	_ReduceToImplicitStructExpr                                         = _ReduceType(161)
	_ReduceAtomExprToAccessibleExpr                                     = _ReduceType(162)
	_ReduceAccessExprToAccessibleExpr                                   = _ReduceType(163)
	_ReduceCallExprToAccessibleExpr                                     = _ReduceType(164)
	_ReduceIndexExprToAccessibleExpr                                    = _ReduceType(165)
	_ReduceToAccessExpr                                                 = _ReduceType(166)
	_ReduceToIndexExpr                                                  = _ReduceType(167)
	_ReduceAccessibleExprToPostfixableExpr                              = _ReduceType(168)
	_ReducePostfixUnaryExprToPostfixableExpr                            = _ReduceType(169)
	_ReduceQuestionToPostfixUnaryOp                                     = _ReduceType(170)
	_ReduceExclaimToPostfixUnaryOp                                      = _ReduceType(171)
	_ReduceToPostfixUnaryExpr                                           = _ReduceType(172)
	_ReducePostfixableExprToPrefixableExpr                              = _ReduceType(173)
	_ReducePrefixUnaryExprToPrefixableExpr                              = _ReduceType(174)
	_ReduceToPrefixUnaryExpr                                            = _ReduceType(175)
	_ReduceNotToPrefixUnaryOp                                           = _ReduceType(176)
	_ReduceBitNegToPrefixUnaryOp                                        = _ReduceType(177)
	_ReduceSubToPrefixUnaryOp                                           = _ReduceType(178)
	_ReduceMulToPrefixUnaryOp                                           = _ReduceType(179)
	_ReduceBitAndToPrefixUnaryOp                                        = _ReduceType(180)
	_ReducePrefixableExprToMulExpr                                      = _ReduceType(181)
	_ReduceBinaryMulExprToMulExpr                                       = _ReduceType(182)
	_ReduceToBinaryMulExpr                                              = _ReduceType(183)
	_ReduceMulToMulOp                                                   = _ReduceType(184)
	_ReduceDivToMulOp                                                   = _ReduceType(185)
	_ReduceModToMulOp                                                   = _ReduceType(186)
	_ReduceBitAndToMulOp                                                = _ReduceType(187)
	_ReduceBitLshiftToMulOp                                             = _ReduceType(188)
	_ReduceBitRshiftToMulOp                                             = _ReduceType(189)
	_ReduceMulExprToAddExpr                                             = _ReduceType(190)
	_ReduceBinaryAddExprToAddExpr                                       = _ReduceType(191)
	_ReduceToBinaryAddExpr                                              = _ReduceType(192)
	_ReduceAddToAddOp                                                   = _ReduceType(193)
	_ReduceSubToAddOp                                                   = _ReduceType(194)
	_ReduceBitOrToAddOp                                                 = _ReduceType(195)
	_ReduceBitXorToAddOp                                                = _ReduceType(196)
	_ReduceAddExprToCmpExpr                                             = _ReduceType(197)
	_ReduceBinaryCmpExprToCmpExpr                                       = _ReduceType(198)
	_ReduceToBinaryCmpExpr                                              = _ReduceType(199)
	_ReduceEqualToCmpOp                                                 = _ReduceType(200)
	_ReduceNotEqualToCmpOp                                              = _ReduceType(201)
	_ReduceLessToCmpOp                                                  = _ReduceType(202)
	_ReduceLessOrEqualToCmpOp                                           = _ReduceType(203)
	_ReduceGreaterToCmpOp                                               = _ReduceType(204)
	_ReduceGreaterOrEqualToCmpOp                                        = _ReduceType(205)
	_ReduceCmpExprToAndExpr                                             = _ReduceType(206)
	_ReduceBinaryAndExprToAndExpr                                       = _ReduceType(207)
	_ReduceToBinaryAndExpr                                              = _ReduceType(208)
	_ReduceAndExprToOrExpr                                              = _ReduceType(209)
	_ReduceBinaryOrExprToOrExpr                                         = _ReduceType(210)
	_ReduceToBinaryOrExpr                                               = _ReduceType(211)
	_ReduceExplicitStructTypeExprToInitializableTypeExpr                = _ReduceType(212)
	_ReduceSliceTypeExprToInitializableTypeExpr                         = _ReduceType(213)
	_ReduceArrayTypeExprToInitializableTypeExpr                         = _ReduceType(214)
	_ReduceMapTypeExprToInitializableTypeExpr                           = _ReduceType(215)
	_ReduceToSliceTypeExpr                                              = _ReduceType(216)
	_ReduceToArrayTypeExpr                                              = _ReduceType(217)
	_ReduceToMapTypeExpr                                                = _ReduceType(218)
	_ReduceInitializableTypeExprToAtomTypeExpr                          = _ReduceType(219)
	_ReduceNamedTypeExprToAtomTypeExpr                                  = _ReduceType(220)
	_ReduceInferredTypeExprToAtomTypeExpr                               = _ReduceType(221)
	_ReduceImplicitStructTypeExprToAtomTypeExpr                         = _ReduceType(222)
	_ReduceExplicitEnumTypeExprToAtomTypeExpr                           = _ReduceType(223)
	_ReduceImplicitEnumTypeExprToAtomTypeExpr                           = _ReduceType(224)
	_ReduceTraitTypeExprToAtomTypeExpr                                  = _ReduceType(225)
	_ReduceFuncTypeExprToAtomTypeExpr                                   = _ReduceType(226)
	_ReduceParseErrorTypeExprToAtomTypeExpr                             = _ReduceType(227)
	_ReduceLocalToNamedTypeExpr                                         = _ReduceType(228)
	_ReduceExternalToNamedTypeExpr                                      = _ReduceType(229)
	_ReduceDotToInferredTypeExpr                                        = _ReduceType(230)
	_ReduceUnderscoreToInferredTypeExpr                                 = _ReduceType(231)
	_ReduceToParseErrorTypeExpr                                         = _ReduceType(232)
	_ReduceAtomTypeExprToReturnableTypeExpr                             = _ReduceType(233)
	_ReducePrefixUnaryTypeExprToReturnableTypeExpr                      = _ReduceType(234)
	_ReduceToPrefixUnaryTypeExpr                                        = _ReduceType(235)
	_ReduceQuestionToPrefixUnaryTypeOp                                  = _ReduceType(236)
	_ReduceExclaimToPrefixUnaryTypeOp                                   = _ReduceType(237)
	_ReduceBitAndToPrefixUnaryTypeOp                                    = _ReduceType(238)
	_ReduceBitNegToPrefixUnaryTypeOp                                    = _ReduceType(239)
	_ReduceTildeTildeToPrefixUnaryTypeOp                                = _ReduceType(240)
	_ReduceReturnableTypeExprToTypeExpr                                 = _ReduceType(241)
	_ReduceBinaryTypeExprToTypeExpr                                     = _ReduceType(242)
	_ReduceToBinaryTypeExpr                                             = _ReduceType(243)
	_ReduceMulToBinaryTypeOp                                            = _ReduceType(244)
	_ReduceAddToBinaryTypeOp                                            = _ReduceType(245)
	_ReduceSubToBinaryTypeOp                                            = _ReduceType(246)
	_ReduceDefinitionToTypeDef                                          = _ReduceType(247)
	_ReduceConstrainedDefToTypeDef                                      = _ReduceType(248)
	_ReduceAliasToTypeDef                                               = _ReduceType(249)
	_ReduceUnconstrainedToGenericParameter                              = _ReduceType(250)
	_ReduceConstrainedToGenericParameter                                = _ReduceType(251)
	_ReduceGenericToGenericParameters                                   = _ReduceType(252)
	_ReduceNilToGenericParameters                                       = _ReduceType(253)
	_ReduceAddToProperGenericParameterList                              = _ReduceType(254)
	_ReduceGenericParameterToProperGenericParameterList                 = _ReduceType(255)
	_ReduceProperGenericParameterListToGenericParameterList             = _ReduceType(256)
	_ReduceImproperToGenericParameterList                               = _ReduceType(257)
	_ReduceNilToGenericParameterList                                    = _ReduceType(258)
	_ReduceBindingToGenericArguments                                    = _ReduceType(259)
	_ReduceNilToGenericArguments                                        = _ReduceType(260)
	_ReduceAddToProperGenericArgumentList                               = _ReduceType(261)
	_ReduceTypeExprToProperGenericArgumentList                          = _ReduceType(262)
	_ReduceProperGenericArgumentListToGenericArgumentList               = _ReduceType(263)
	_ReduceImproperToGenericArgumentList                                = _ReduceType(264)
	_ReduceNilToGenericArgumentList                                     = _ReduceType(265)
	_ReduceNamedToFieldDef                                              = _ReduceType(266)
	_ReduceUnnamedToFieldDef                                            = _ReduceType(267)
	_ReduceToUnsafeStatementProperty                                    = _ReduceType(268)
	_ReduceFieldDefToTypeProperty                                       = _ReduceType(269)
	_ReduceDefaultEnumFieldDefToTypeProperty                            = _ReduceType(270)
	_ReducePaddingFieldDefToTypeProperty                                = _ReduceType(271)
	_ReduceMethodSignatureToTypeProperty                                = _ReduceType(272)
	_ReduceUnsafeStatementPropertyToTypeProperty                        = _ReduceType(273)
	_ReduceAddToProperImplicitTypeProperties                            = _ReduceType(274)
	_ReduceTypePropertyToProperImplicitTypeProperties                   = _ReduceType(275)
	_ReduceProperImplicitTypePropertiesToImplicitTypeProperties         = _ReduceType(276)
	_ReduceImproperToImplicitTypeProperties                             = _ReduceType(277)
	_ReduceNilToImplicitTypeProperties                                  = _ReduceType(278)
	_ReduceToImplicitStructTypeExpr                                     = _ReduceType(279)
	_ReduceAddImplicitToProperExplicitTypeProperties                    = _ReduceType(280)
	_ReduceAddExplicitToProperExplicitTypeProperties                    = _ReduceType(281)
	_ReduceTypePropertyToProperExplicitTypeProperties                   = _ReduceType(282)
	_ReduceProperExplicitTypePropertiesToExplicitTypeProperties         = _ReduceType(283)
	_ReduceImproperImplicitToExplicitTypeProperties                     = _ReduceType(284)
	_ReduceImproperExplicitToExplicitTypeProperties                     = _ReduceType(285)
	_ReduceNilToExplicitTypeProperties                                  = _ReduceType(286)
	_ReduceToExplicitStructTypeExpr                                     = _ReduceType(287)
	_ReduceToTraitTypeExpr                                              = _ReduceType(288)
	_ReducePairToProperImplicitEnumTypeProperties                       = _ReduceType(289)
	_ReduceAddToProperImplicitEnumTypeProperties                        = _ReduceType(290)
	_ReduceProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties = _ReduceType(291)
	_ReduceImproperToImplicitEnumTypeProperties                         = _ReduceType(292)
	_ReduceToImplicitEnumTypeExpr                                       = _ReduceType(293)
	_ReduceExplicitPairToProperExplicitEnumTypeProperties               = _ReduceType(294)
	_ReduceImplicitPairToProperExplicitEnumTypeProperties               = _ReduceType(295)
	_ReduceExplicitAddToProperExplicitEnumTypeProperties                = _ReduceType(296)
	_ReduceImplicitAddToProperExplicitEnumTypeProperties                = _ReduceType(297)
	_ReduceProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties = _ReduceType(298)
	_ReduceImproperToExplicitEnumTypeProperties                         = _ReduceType(299)
	_ReduceToExplicitEnumTypeExpr                                       = _ReduceType(300)
	_ReduceReturnableTypeExprToReturnType                               = _ReduceType(301)
	_ReduceNilToReturnType                                              = _ReduceType(302)
	_ReduceNamedTypedArgToProperParameterDef                            = _ReduceType(303)
	_ReduceNamedTypedVarargToProperParameterDef                         = _ReduceType(304)
	_ReduceNamedInferredVarargToProperParameterDef                      = _ReduceType(305)
	_ReduceIgnoreTypedArgToProperParameterDef                           = _ReduceType(306)
	_ReduceIgnoreInferredVarargToProperParameterDef                     = _ReduceType(307)
	_ReduceIgnoreTypedVarargToProperParameterDef                        = _ReduceType(308)
	_ReduceProperParameterDefToParameterDecl                            = _ReduceType(309)
	_ReduceUnnamedTypedArgToParameterDecl                               = _ReduceType(310)
	_ReduceUnnamedInferredVarargToParameterDecl                         = _ReduceType(311)
	_ReduceUnnamedTypedVarargToParameterDecl                            = _ReduceType(312)
	_ReduceProperParameterDefToParameterDef                             = _ReduceType(313)
	_ReduceNamedInferredArgToParameterDef                               = _ReduceType(314)
	_ReduceIgnoreInferredArgToParameterDef                              = _ReduceType(315)
	_ReduceAddToProperParameterDecls                                    = _ReduceType(316)
	_ReduceParameterDeclToProperParameterDecls                          = _ReduceType(317)
	_ReduceProperParameterDeclsToParameterDecls                         = _ReduceType(318)
	_ReduceImproperToParameterDecls                                     = _ReduceType(319)
	_ReduceNilToParameterDecls                                          = _ReduceType(320)
	_ReduceAddToProperParameterDefs                                     = _ReduceType(321)
	_ReduceParameterDefToProperParameterDefs                            = _ReduceType(322)
	_ReduceProperParameterDefsToParameterDefs                           = _ReduceType(323)
	_ReduceImproperToParameterDefs                                      = _ReduceType(324)
	_ReduceNilToParameterDefs                                           = _ReduceType(325)
	_ReduceToFuncTypeExpr                                               = _ReduceType(326)
	_ReduceToMethodSignature                                            = _ReduceType(327)
	_ReduceFuncDefToNamedFuncDef                                        = _ReduceType(328)
	_ReduceMethodDefToNamedFuncDef                                      = _ReduceType(329)
	_ReduceAliasToNamedFuncDef                                          = _ReduceType(330)
	_ReduceToAnonymousFuncExpr                                          = _ReduceType(331)
	_ReduceNoSpecToPackageDef                                           = _ReduceType(332)
	_ReduceWithSpecToPackageDef                                         = _ReduceType(333)
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
	case _ReduceUnconstrainedToGenericParameter:
		return "UnconstrainedToGenericParameter"
	case _ReduceConstrainedToGenericParameter:
		return "ConstrainedToGenericParameter"
	case _ReduceGenericToGenericParameters:
		return "GenericToGenericParameters"
	case _ReduceNilToGenericParameters:
		return "NilToGenericParameters"
	case _ReduceAddToProperGenericParameterList:
		return "AddToProperGenericParameterList"
	case _ReduceGenericParameterToProperGenericParameterList:
		return "GenericParameterToProperGenericParameterList"
	case _ReduceProperGenericParameterListToGenericParameterList:
		return "ProperGenericParameterListToGenericParameterList"
	case _ReduceImproperToGenericParameterList:
		return "ImproperToGenericParameterList"
	case _ReduceNilToGenericParameterList:
		return "NilToGenericParameterList"
	case _ReduceBindingToGenericArguments:
		return "BindingToGenericArguments"
	case _ReduceNilToGenericArguments:
		return "NilToGenericArguments"
	case _ReduceAddToProperGenericArgumentList:
		return "AddToProperGenericArgumentList"
	case _ReduceTypeExprToProperGenericArgumentList:
		return "TypeExprToProperGenericArgumentList"
	case _ReduceProperGenericArgumentListToGenericArgumentList:
		return "ProperGenericArgumentListToGenericArgumentList"
	case _ReduceImproperToGenericArgumentList:
		return "ImproperToGenericArgumentList"
	case _ReduceNilToGenericArgumentList:
		return "NilToGenericArgumentList"
	case _ReduceNamedToFieldDef:
		return "NamedToFieldDef"
	case _ReduceUnnamedToFieldDef:
		return "UnnamedToFieldDef"
	case _ReduceToUnsafeStatementProperty:
		return "ToUnsafeStatementProperty"
	case _ReduceFieldDefToTypeProperty:
		return "FieldDefToTypeProperty"
	case _ReduceDefaultEnumFieldDefToTypeProperty:
		return "DefaultEnumFieldDefToTypeProperty"
	case _ReducePaddingFieldDefToTypeProperty:
		return "PaddingFieldDefToTypeProperty"
	case _ReduceMethodSignatureToTypeProperty:
		return "MethodSignatureToTypeProperty"
	case _ReduceUnsafeStatementPropertyToTypeProperty:
		return "UnsafeStatementPropertyToTypeProperty"
	case _ReduceAddToProperImplicitTypeProperties:
		return "AddToProperImplicitTypeProperties"
	case _ReduceTypePropertyToProperImplicitTypeProperties:
		return "TypePropertyToProperImplicitTypeProperties"
	case _ReduceProperImplicitTypePropertiesToImplicitTypeProperties:
		return "ProperImplicitTypePropertiesToImplicitTypeProperties"
	case _ReduceImproperToImplicitTypeProperties:
		return "ImproperToImplicitTypeProperties"
	case _ReduceNilToImplicitTypeProperties:
		return "NilToImplicitTypeProperties"
	case _ReduceToImplicitStructTypeExpr:
		return "ToImplicitStructTypeExpr"
	case _ReduceAddImplicitToProperExplicitTypeProperties:
		return "AddImplicitToProperExplicitTypeProperties"
	case _ReduceAddExplicitToProperExplicitTypeProperties:
		return "AddExplicitToProperExplicitTypeProperties"
	case _ReduceTypePropertyToProperExplicitTypeProperties:
		return "TypePropertyToProperExplicitTypeProperties"
	case _ReduceProperExplicitTypePropertiesToExplicitTypeProperties:
		return "ProperExplicitTypePropertiesToExplicitTypeProperties"
	case _ReduceImproperImplicitToExplicitTypeProperties:
		return "ImproperImplicitToExplicitTypeProperties"
	case _ReduceImproperExplicitToExplicitTypeProperties:
		return "ImproperExplicitToExplicitTypeProperties"
	case _ReduceNilToExplicitTypeProperties:
		return "NilToExplicitTypeProperties"
	case _ReduceToExplicitStructTypeExpr:
		return "ToExplicitStructTypeExpr"
	case _ReduceToTraitTypeExpr:
		return "ToTraitTypeExpr"
	case _ReducePairToProperImplicitEnumTypeProperties:
		return "PairToProperImplicitEnumTypeProperties"
	case _ReduceAddToProperImplicitEnumTypeProperties:
		return "AddToProperImplicitEnumTypeProperties"
	case _ReduceProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties:
		return "ProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties"
	case _ReduceImproperToImplicitEnumTypeProperties:
		return "ImproperToImplicitEnumTypeProperties"
	case _ReduceToImplicitEnumTypeExpr:
		return "ToImplicitEnumTypeExpr"
	case _ReduceExplicitPairToProperExplicitEnumTypeProperties:
		return "ExplicitPairToProperExplicitEnumTypeProperties"
	case _ReduceImplicitPairToProperExplicitEnumTypeProperties:
		return "ImplicitPairToProperExplicitEnumTypeProperties"
	case _ReduceExplicitAddToProperExplicitEnumTypeProperties:
		return "ExplicitAddToProperExplicitEnumTypeProperties"
	case _ReduceImplicitAddToProperExplicitEnumTypeProperties:
		return "ImplicitAddToProperExplicitEnumTypeProperties"
	case _ReduceProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties:
		return "ProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties"
	case _ReduceImproperToExplicitEnumTypeProperties:
		return "ImproperToExplicitEnumTypeProperties"
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
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Argument            *Argument
	ArgumentList        *ArgumentList
	ColonExpr           *ColonExpr
	Count               TokenCount
	Expression          Expression
	FieldDef            *FieldDef
	GenericArgumentList *GenericArgumentList
	Parameter           *Parameter
	Parameters          *ParameterList
	ParseError          ParseErrorSymbol
	SourceDefinition    SourceDefinition
	SourceDefinitions   []SourceDefinition
	Statement           Statement
	Statements          *StatementList
	TypeExpression      TypeExpression
	TypeProperties      *TypePropertyList
	TypeProperty        TypeProperty
	Value               TokenValue
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
	case FieldDefType:
		loc, ok := interface{}(s.FieldDef).(locator)
		if ok {
			return loc.Loc()
		}
	case GenericArgumentsType, ProperGenericArgumentListType, GenericArgumentListType:
		loc, ok := interface{}(s.GenericArgumentList).(locator)
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
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType, FuncTypeExprType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperImplicitTypePropertiesType, ImplicitTypePropertiesType, ProperExplicitTypePropertiesType, ExplicitTypePropertiesType, ProperImplicitEnumTypePropertiesType, ImplicitEnumTypePropertiesType, ProperExplicitEnumTypePropertiesType, ExplicitEnumTypePropertiesType:
		loc, ok := interface{}(s.TypeProperties).(locator)
		if ok {
			return loc.Loc()
		}
	case UnsafeStatementPropertyType, TypePropertyType, MethodSignatureType:
		loc, ok := interface{}(s.TypeProperty).(locator)
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
	case FieldDefType:
		loc, ok := interface{}(s.FieldDef).(locator)
		if ok {
			return loc.End()
		}
	case GenericArgumentsType, ProperGenericArgumentListType, GenericArgumentListType:
		loc, ok := interface{}(s.GenericArgumentList).(locator)
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
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType, FuncTypeExprType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case ProperImplicitTypePropertiesType, ImplicitTypePropertiesType, ProperExplicitTypePropertiesType, ExplicitTypePropertiesType, ProperImplicitEnumTypePropertiesType, ImplicitEnumTypePropertiesType, ProperExplicitEnumTypePropertiesType, ExplicitEnumTypePropertiesType:
		loc, ok := interface{}(s.TypeProperties).(locator)
		if ok {
			return loc.End()
		}
	case UnsafeStatementPropertyType, TypePropertyType, MethodSignatureType:
		loc, ok := interface{}(s.TypeProperty).(locator)
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
		symbol.Expression, err = reducer.ToCallExpr(args[0].Expression, args[1].GenericArgumentList, args[2].Value, args[3].ArgumentList, args[4].Value)
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
		symbol.TypeExpression, err = reducer.LocalToNamedTypeExpr(args[0].Value, args[1].GenericArgumentList)
	case _ReduceExternalToNamedTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedTypeExprType
		symbol.TypeExpression, err = reducer.ExternalToNamedTypeExpr(args[0].Value, args[1].Value, args[2].Value, args[3].GenericArgumentList)
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
	case _ReduceUnconstrainedToGenericParameter:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterType
		symbol.Generic_, err = reducer.UnconstrainedToGenericParameter(args[0].Value)
	case _ReduceConstrainedToGenericParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterType
		symbol.Generic_, err = reducer.ConstrainedToGenericParameter(args[0].Value, args[1].TypeExpression)
	case _ReduceGenericToGenericParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParametersType
		symbol.Generic_, err = reducer.GenericToGenericParameters(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceNilToGenericParameters:
		symbol.SymbolId_ = GenericParametersType
		symbol.Generic_, err = reducer.NilToGenericParameters()
	case _ReduceAddToProperGenericParameterList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericParameterListType
		symbol.Generic_, err = reducer.AddToProperGenericParameterList(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceGenericParameterToProperGenericParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericParameterListType
		symbol.Generic_, err = reducer.GenericParameterToProperGenericParameterList(args[0].Generic_)
	case _ReduceProperGenericParameterListToGenericParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterListType
		//line grammar.lr:695:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToGenericParameterList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterListType
		symbol.Generic_, err = reducer.ImproperToGenericParameterList(args[0].Generic_, args[1].Value)
	case _ReduceNilToGenericParameterList:
		symbol.SymbolId_ = GenericParameterListType
		symbol.Generic_, err = reducer.NilToGenericParameterList()
	case _ReduceBindingToGenericArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.GenericArgumentList, err = reducer.BindingToGenericArguments(args[0].Value, args[1].GenericArgumentList, args[2].Value)
	case _ReduceNilToGenericArguments:
		symbol.SymbolId_ = GenericArgumentsType
		symbol.GenericArgumentList, err = reducer.NilToGenericArguments()
	case _ReduceAddToProperGenericArgumentList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericArgumentListType
		symbol.GenericArgumentList, err = reducer.AddToProperGenericArgumentList(args[0].GenericArgumentList, args[1].Value, args[2].TypeExpression)
	case _ReduceTypeExprToProperGenericArgumentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericArgumentListType
		symbol.GenericArgumentList, err = reducer.TypeExprToProperGenericArgumentList(args[0].TypeExpression)
	case _ReduceProperGenericArgumentListToGenericArgumentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericArgumentListType
		//line grammar.lr:708:4
		symbol.GenericArgumentList = args[0].GenericArgumentList
		err = nil
	case _ReduceImproperToGenericArgumentList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericArgumentListType
		symbol.GenericArgumentList, err = reducer.ImproperToGenericArgumentList(args[0].GenericArgumentList, args[1].Value)
	case _ReduceNilToGenericArgumentList:
		symbol.SymbolId_ = GenericArgumentListType
		symbol.GenericArgumentList, err = reducer.NilToGenericArgumentList()
	case _ReduceNamedToFieldDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDefType
		symbol.FieldDef, err = reducer.NamedToFieldDef(args[0].Value, args[1].TypeExpression)
	case _ReduceUnnamedToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.FieldDef, err = reducer.UnnamedToFieldDef(args[0].TypeExpression)
	case _ReduceToUnsafeStatementProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnsafeStatementPropertyType
		symbol.TypeProperty, err = reducer.ToUnsafeStatementProperty(args[0].Statement)
	case _ReduceFieldDefToTypeProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypePropertyType
		//line grammar.lr:730:4
		symbol.TypeProperty = args[0].FieldDef
		err = nil
	case _ReduceDefaultEnumFieldDefToTypeProperty:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.DefaultEnumFieldDefToTypeProperty(args[0].Value, args[1].FieldDef)
	case _ReducePaddingFieldDefToTypeProperty:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.PaddingFieldDefToTypeProperty(args[0].Value, args[1].TypeExpression)
	case _ReduceMethodSignatureToTypeProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypePropertyType
		//line grammar.lr:733:4
		symbol.TypeProperty = args[0].TypeProperty
		err = nil
	case _ReduceUnsafeStatementPropertyToTypeProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypePropertyType
		//line grammar.lr:734:4
		symbol.TypeProperty = args[0].TypeProperty
		err = nil
	case _ReduceAddToProperImplicitTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.AddToProperImplicitTypeProperties(args[0].TypeProperties, args[1].Value, args[2].TypeProperty)
	case _ReduceTypePropertyToProperImplicitTypeProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperImplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.TypePropertyToProperImplicitTypeProperties(args[0].TypeProperty)
	case _ReduceProperImplicitTypePropertiesToImplicitTypeProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitTypePropertiesType
		//line grammar.lr:741:4
		symbol.TypeProperties = args[0].TypeProperties
		err = nil
	case _ReduceImproperToImplicitTypeProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.ImproperToImplicitTypeProperties(args[0].TypeProperties, args[1].Value)
	case _ReduceNilToImplicitTypeProperties:
		symbol.SymbolId_ = ImplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.NilToImplicitTypeProperties()
	case _ReduceToImplicitStructTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructTypeExprType
		symbol.TypeExpression, err = reducer.ToImplicitStructTypeExpr(args[0].Value, args[1].TypeProperties, args[2].Value)
	case _ReduceAddImplicitToProperExplicitTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.AddImplicitToProperExplicitTypeProperties(args[0].TypeProperties, args[1].Count, args[2].TypeProperty)
	case _ReduceAddExplicitToProperExplicitTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.AddExplicitToProperExplicitTypeProperties(args[0].TypeProperties, args[1].Value, args[2].TypeProperty)
	case _ReduceTypePropertyToProperExplicitTypeProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperExplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.TypePropertyToProperExplicitTypeProperties(args[0].TypeProperty)
	case _ReduceProperExplicitTypePropertiesToExplicitTypeProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitTypePropertiesType
		//line grammar.lr:754:4
		symbol.TypeProperties = args[0].TypeProperties
		err = nil
	case _ReduceImproperImplicitToExplicitTypeProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.ImproperImplicitToExplicitTypeProperties(args[0].TypeProperties, args[1].Count)
	case _ReduceImproperExplicitToExplicitTypeProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.ImproperExplicitToExplicitTypeProperties(args[0].TypeProperties, args[1].Value)
	case _ReduceNilToExplicitTypeProperties:
		symbol.SymbolId_ = ExplicitTypePropertiesType
		symbol.TypeProperties, err = reducer.NilToExplicitTypeProperties()
	case _ReduceToExplicitStructTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitStructTypeExprType
		symbol.TypeExpression, err = reducer.ToExplicitStructTypeExpr(args[0].Value, args[1].Value, args[2].TypeProperties, args[3].Value)
	case _ReduceToTraitTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TraitTypeExprType
		symbol.TypeExpression, err = reducer.ToTraitTypeExpr(args[0].Value, args[1].Value, args[2].TypeProperties, args[3].Value)
	case _ReducePairToProperImplicitEnumTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.PairToProperImplicitEnumTypeProperties(args[0].TypeProperty, args[1].Value, args[2].TypeProperty)
	case _ReduceAddToProperImplicitEnumTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.AddToProperImplicitEnumTypeProperties(args[0].TypeProperties, args[1].Value, args[2].TypeProperty)
	case _ReduceProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitEnumTypePropertiesType
		//line grammar.lr:778:4
		symbol.TypeProperties = args[0].TypeProperties
		err = nil
	case _ReduceImproperToImplicitEnumTypeProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.ImproperToImplicitEnumTypeProperties(args[0].TypeProperties, args[1].Count)
	case _ReduceToImplicitEnumTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumTypeExprType
		symbol.TypeExpression, err = reducer.ToImplicitEnumTypeExpr(args[0].Value, args[1].TypeProperties, args[2].Value)
	case _ReduceExplicitPairToProperExplicitEnumTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.ExplicitPairToProperExplicitEnumTypeProperties(args[0].TypeProperty, args[1].Value, args[2].TypeProperty)
	case _ReduceImplicitPairToProperExplicitEnumTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.ImplicitPairToProperExplicitEnumTypeProperties(args[0].TypeProperty, args[1].Count, args[2].TypeProperty)
	case _ReduceExplicitAddToProperExplicitEnumTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.ExplicitAddToProperExplicitEnumTypeProperties(args[0].TypeProperties, args[1].Value, args[2].TypeProperty)
	case _ReduceImplicitAddToProperExplicitEnumTypeProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.ImplicitAddToProperExplicitEnumTypeProperties(args[0].TypeProperties, args[1].Count, args[2].TypeProperty)
	case _ReduceProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitEnumTypePropertiesType
		//line grammar.lr:792:4
		symbol.TypeProperties = args[0].TypeProperties
		err = nil
	case _ReduceImproperToExplicitEnumTypeProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitEnumTypePropertiesType
		symbol.TypeProperties, err = reducer.ImproperToExplicitEnumTypeProperties(args[0].TypeProperties, args[1].Count)
	case _ReduceToExplicitEnumTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitEnumTypeExprType
		symbol.TypeExpression, err = reducer.ToExplicitEnumTypeExpr(args[0].Value, args[1].Value, args[2].TypeProperties, args[3].Value)
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
		//line grammar.lr:819:4
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
		//line grammar.lr:829:4
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
		//line grammar.lr:838:4
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
		//line grammar.lr:847:4
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
		symbol.TypeProperty, err = reducer.ToMethodSignature(args[0].Value, args[1].Value, args[2].Value, args[3].Parameters, args[4].Value, args[5].TypeExpression)
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
			return _Action{_ShiftAction, _State11, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SourceType:
			return _Action{_ShiftAction, _State5, 0}, true
		case ProperDefinitionsType:
			return _Action{_ShiftAction, _State13, 0}, true
		case VarDeclPatternType:
			return _Action{_ShiftAction, _State14, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
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
			return _Action{_ShiftAction, _State16, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case StatementType:
			return _Action{_ShiftAction, _State6, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State31, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State29, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State33, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State28, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State38, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State7, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
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
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
		}
	case _State9:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State46, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State47, 0}, true
		}
	case _State10:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ProperStatementsType:
			return _Action{_ShiftAction, _State48, 0}, true
		case StatementsType:
			return _Action{_ShiftAction, _State49, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State31, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State29, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State33, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State28, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State38, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State11:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWithSpecToPackageDef}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}, true
		}
	case _State12:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State50, 0}, true
		}
	case _State13:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State51, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperDefinitionsToDefinitions}, true
		}
	case _State14:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State52, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}, true
		}
	case _State15:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State53, 0}, true
		case VarPatternType:
			return _Action{_ShiftAction, _State54, 0}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToVarPattern}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToVarPattern}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTuplePatternToVarPattern}, true
		}
	case _State16:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case VarToken:
			return _Action{_ShiftAction, _State56, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State55, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case CasePatternsType:
			return _Action{_ShiftAction, _State57, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State17:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State59, 0}, true
		}
	case _State18:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State60, 0}, true
		}
	case _State19:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State20:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State61, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSingleToImportStatement}, true
		}
	case _State21:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State65, 0}, true
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
	case _State22:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State67, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State66, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State70, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State71, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State68, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State69, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State23:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State72, 0}, true
		}
	case _State24:
		switch symbolId {
		case LessToken:
			return _Action{_ShiftAction, _State73, 0}, true
		}
	case _State25:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State76, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State75, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case BinaryOpAssignType:
			return _Action{_ShiftAction, _State77, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAction, _State78, 0}, true
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
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}, true
		}
	case _State26:
		switch symbolId {
		case AddOpType:
			return _Action{_ShiftAction, _State79, 0}, true
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
	case _State27:
		switch symbolId {
		case AndToken:
			return _Action{_ShiftAction, _State80, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}, true
		}
	case _State28:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State81, 0}, true
		}
	case _State29:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case CallExprType:
			return _Action{_ShiftAction, _State83, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State82, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State30:
		switch symbolId {
		case CmpOpType:
			return _Action{_ShiftAction, _State84, 0}, true
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
	case _State31:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State85, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToExprOrImproperStructStatement}, true
		}
	case _State32:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State86, 0}, true
		}
	case _State33:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case JumpLabelToken:
			return _Action{_ShiftAction, _State87, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State88, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State34:
		switch symbolId {
		case MulOpType:
			return _Action{_ShiftAction, _State89, 0}, true
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
	case _State35:
		switch symbolId {
		case IfToken:
			return _Action{_ShiftAction, _State92, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State93, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State91, 0}, true
		case DoToken:
			return _Action{_ShiftAction, _State90, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBlockExpr}, true
		case IfExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIfExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToLoopExpr}, true
		}
	case _State36:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State94, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}, true
		}
	case _State37:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State38:
		switch symbolId {
		case AssignToken:
			return _Action{_ReduceAction, 0, _ReduceToAssignPattern}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceSequenceExprToExpr}, true
		}
	case _State39:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State76, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State75, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAction, _State78, 0}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPostfixUnaryOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPostfixUnaryOp}, true
		case PostfixUnaryOpType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToPostfixUnaryExpr}, true
		case LparenToken:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}, true
		}
	case _State40:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State95, 0}, true
		}
	case _State41:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State96, 0}, true
		}
	case _State42:
		switch symbolId {
		case DotToken:
			return _Action{_ShiftAction, _State97, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLocalToNamedTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State43:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
		case TypePropertyType:
			return _Action{_ShiftAction, _State107, 0}, true
		case ProperImplicitTypePropertiesType:
			return _Action{_ShiftAction, _State105, 0}, true
		case ImplicitTypePropertiesType:
			return _Action{_ShiftAction, _State103, 0}, true
		case ProperImplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State104, 0}, true
		case ImplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State102, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToImplicitTypeProperties}, true
		}
	case _State44:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State108, 0}, true
		}
	case _State45:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
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
	case _State46:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State111, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State110, 0}, true
		case GenericParametersType:
			return _Action{_ShiftAction, _State112, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameters}, true
		}
	case _State47:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State114, 0}, true
		case ParameterDefType:
			return _Action{_ShiftAction, _State115, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		}
	case _State48:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State116, 0}, true
		case SemicolonToken:
			return _Action{_ShiftAction, _State117, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperStatementsToStatements}, true
		}
	case _State49:
		switch symbolId {
		case RbraceToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToStatementBlock}, true
		}
	case _State50:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State111, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State118, 0}, true
		case GenericParametersType:
			return _Action{_ShiftAction, _State119, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameters}, true
		}
	case _State51:
		switch symbolId {
		case PackageToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case VarDeclPatternType:
			return _Action{_ShiftAction, _State14, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
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
	case _State52:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State53:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State53, 0}, true
		case FieldVarPatternsType:
			return _Action{_ShiftAction, _State121, 0}, true
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
	case _State54:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State122, 0}, true
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
	case _State55:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State123, 0}, true
		}
	case _State56:
		switch symbolId {
		case DotToken:
			return _Action{_ShiftAction, _State124, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceVarToVarOrLet}, true
		}
	case _State57:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State126, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State125, 0}, true
		}
	case _State58:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToBlockExpr}, true
		}
	case _State59:
		switch symbolId {
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State31, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State29, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State33, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State28, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State38, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State60:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State114, 0}, true
		case ProperParameterDefsType:
			return _Action{_ShiftAction, _State128, 0}, true
		case ParameterDefsType:
			return _Action{_ShiftAction, _State127, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDefToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDefs}, true
		}
	case _State61:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportToLocalToImportClause}, true
		}
	case _State62:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAliasToImportClause}, true
		}
	case _State63:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State61, 0}, true
		case ProperImportClausesType:
			return _Action{_ShiftAction, _State130, 0}, true
		case ImportClausesType:
			return _Action{_ShiftAction, _State129, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportClauseToProperImportClauses}, true
		}
	case _State64:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnusableImportToImportClause}, true
		}
	case _State65:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State132, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State131, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSliceTypeExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State66:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State67:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State133, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		}
	case _State68:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitStructExpr}, true
		}
	case _State69:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State134, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceColonExprToArgument}, true
		}
	case _State70:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State135, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarargAssignmentToArgument}, true

		default:
			return _Action{_ReduceAction, 0, _ReducePositionalToArgument}, true
		}
	case _State71:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State136, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}, true
		}
	case _State72:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
		case ProperExplicitTypePropertiesType:
			return _Action{_ShiftAction, _State138, 0}, true
		case ExplicitTypePropertiesType:
			return _Action{_ShiftAction, _State137, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypePropertyToProperExplicitTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToExplicitTypeProperties}, true
		}
	case _State73:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State139, 0}, true
		}
	case _State74:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State142, 0}, true
		case ProperGenericArgumentListType:
			return _Action{_ShiftAction, _State141, 0}, true
		case GenericArgumentListType:
			return _Action{_ShiftAction, _State140, 0}, true
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
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArgumentList}, true
		}
	case _State75:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAccessExpr}, true
		}
	case _State76:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State67, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State66, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State70, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case ArgumentType:
			return _Action{_ShiftAction, _State143, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State69, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State77:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State78:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State144, 0}, true
		}
	case _State79:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State145, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State80:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State146, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State81:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State82:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State76, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State75, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAction, _State78, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State83:
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
	case _State84:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State147, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State85:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State86:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State67, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State66, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State70, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State71, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State148, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State69, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State87:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State149, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State88:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State85, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}, true
		}
	case _State89:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State90:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAction, _State150, 0}, true
		}
	case _State91:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State151, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State153, 0}, true
		case ForAssignmentType:
			return _Action{_ShiftAction, _State152, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State92:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State154, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case ConditionType:
			return _Action{_ShiftAction, _State155, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State93:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State156, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State94:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State157, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State95:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
		case TypePropertyType:
			return _Action{_ShiftAction, _State160, 0}, true
		case ProperExplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State159, 0}, true
		case ExplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State158, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true
		}
	case _State96:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State163, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State166, 0}, true
		case ProperParameterDeclsType:
			return _Action{_ShiftAction, _State165, 0}, true
		case ParameterDeclsType:
			return _Action{_ShiftAction, _State164, 0}, true
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
	case _State97:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State167, 0}, true
		}
	case _State98:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
		case FieldDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDefaultEnumFieldDefToTypeProperty}, true
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
	case _State99:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State96, 0}, true
		}
	case _State100:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State169, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State170, 0}, true
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
		case GenericArgumentsType:
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
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State101:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State171, 0}, true
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
	case _State102:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitEnumTypeExpr}, true
		}
	case _State103:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitStructTypeExpr}, true
		}
	case _State104:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State172, 0}, true
		case NewlinesToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImproperToImplicitEnumTypeProperties}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties}, true
		}
	case _State105:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State173, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImplicitTypePropertiesToImplicitTypeProperties}, true
		}
	case _State106:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedToFieldDef}, true
		}
	case _State107:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State174, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypePropertyToProperImplicitTypeProperties}, true
		}
	case _State108:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
		case ProperExplicitTypePropertiesType:
			return _Action{_ShiftAction, _State138, 0}, true
		case ExplicitTypePropertiesType:
			return _Action{_ShiftAction, _State175, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypePropertyToProperExplicitTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToExplicitTypeProperties}, true
		}
	case _State109:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
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
	case _State110:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State111:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State176, 0}, true
		case ProperGenericParameterListType:
			return _Action{_ShiftAction, _State178, 0}, true
		case GenericParameterListType:
			return _Action{_ShiftAction, _State177, 0}, true
		case GenericParameterType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGenericParameterToProperGenericParameterList}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameterList}, true
		}
	case _State112:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State179, 0}, true
		}
	case _State113:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State180, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State181, 0}, true
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
	case _State114:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State182, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State183, 0}, true
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
	case _State115:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State184, 0}, true
		}
	case _State116:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State31, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State29, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State33, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State28, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State38, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State117:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State31, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State29, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State33, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State28, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State38, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State118:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State185, 0}, true
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
	case _State119:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State186, 0}, true
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
	case _State120:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State187, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}, true
		}
	case _State121:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State188, 0}, true
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToTuplePattern}, true
		}
	case _State122:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypeExprToOptionalTypeExpr}, true
		}
	case _State123:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceEnumNondataMatchPattenToCasePattern}, true
		}
	case _State124:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State189, 0}, true
		}
	case _State125:
		switch symbolId {
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ExprsType:
			return _Action{_ShiftAction, _State31, 0}, true
		case CallbackOpType:
			return _Action{_ShiftAction, _State29, 0}, true
		case JumpTypeType:
			return _Action{_ShiftAction, _State33, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case AssignPatternType:
			return _Action{_ShiftAction, _State28, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State38, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State126:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case VarToken:
			return _Action{_ShiftAction, _State56, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State55, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State127:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State190, 0}, true
		}
	case _State128:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State191, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperParameterDefsToParameterDefs}, true
		}
	case _State129:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMultipleToImportStatement}, true
		}
	case _State130:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State193, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State192, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}, true
		}
	case _State131:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State194, 0}, true
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
	case _State132:
		switch symbolId {
		case IntegerLiteralToken:
			return _Action{_ShiftAction, _State195, 0}, true
		}
	case _State133:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State134:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State135:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State136:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State67, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State66, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State70, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State69, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State137:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToExplicitStructTypeExpr}, true
		}
	case _State138:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State197, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State196, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperExplicitTypePropertiesToExplicitTypeProperties}, true
		}
	case _State139:
		switch symbolId {
		case GreaterToken:
			return _Action{_ShiftAction, _State198, 0}, true
		}
	case _State140:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBindingToGenericArguments}, true
		}
	case _State141:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State199, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperGenericArgumentListToGenericArgumentList}, true
		}
	case _State142:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypeExprToProperGenericArgumentList}, true
		}
	case _State143:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIndexExpr}, true
		}
	case _State144:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State67, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State66, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State70, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State35, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State71, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State200, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State69, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State145:
		switch symbolId {
		case MulOpType:
			return _Action{_ShiftAction, _State89, 0}, true
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
	case _State146:
		switch symbolId {
		case CmpOpType:
			return _Action{_ShiftAction, _State84, 0}, true
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
	case _State147:
		switch symbolId {
		case AddOpType:
			return _Action{_ShiftAction, _State79, 0}, true
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
	case _State148:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToInitializeExpr}, true
		}
	case _State149:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State85, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}, true
		}
	case _State150:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State201, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceInfiniteToLoopExprBody}, true
		}
	case _State151:
		switch symbolId {
		case InToken:
			return _Action{_ShiftAction, _State203, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State202, 0}, true
		}
	case _State152:
		switch symbolId {
		case SemicolonToken:
			return _Action{_ShiftAction, _State204, 0}, true
		}
	case _State153:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State205, 0}, true
		case SemicolonToken:
			return _Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToAssignPattern}, true
		}
	case _State154:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case VarToken:
			return _Action{_ShiftAction, _State56, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State55, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case CasePatternsType:
			return _Action{_ShiftAction, _State206, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State155:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAction, _State207, 0}, true
		}
	case _State156:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSwitchExpr}, true
		}
	case _State157:
		switch symbolId {
		case AndToken:
			return _Action{_ShiftAction, _State80, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}, true
		}
	case _State158:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToExplicitEnumTypeExpr}, true
		}
	case _State159:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State208, 0}, true
		case OrToken:
			return _Action{_ShiftAction, _State209, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties}, true
		}
	case _State160:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State210, 0}, true
		case OrToken:
			return _Action{_ShiftAction, _State211, 0}, true
		}
	case _State161:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State212, 0}, true
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
	case _State162:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State169, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State180, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State181, 0}, true
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
		case GenericArgumentsType:
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
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State163:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State182, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State183, 0}, true
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
	case _State164:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State213, 0}, true
		}
	case _State165:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State214, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperParameterDeclsToParameterDecls}, true
		}
	case _State166:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedTypedArgToParameterDecl}, true
		}
	case _State167:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State74, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExternalToNamedTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State168:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State215, 0}, true
		}
	case _State169:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State167, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		}
	case _State170:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedToFieldDef}, true
		}
	case _State171:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReducePaddingFieldDefToTypeProperty}, true
		}
	case _State172:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperImplicitEnumTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true
		}
	case _State173:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperImplicitTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToImplicitTypeProperties}, true
		}
	case _State174:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePairToProperImplicitEnumTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true
		}
	case _State175:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToTraitTypeExpr}, true
		}
	case _State176:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State216, 0}, true
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
			return _Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameter}, true
		}
	case _State177:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGenericToGenericParameters}, true
		}
	case _State178:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State217, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperGenericParameterListToGenericParameterList}, true
		}
	case _State179:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State114, 0}, true
		case ProperParameterDefsType:
			return _Action{_ShiftAction, _State128, 0}, true
		case ParameterDefsType:
			return _Action{_ShiftAction, _State218, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDefToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDefs}, true
		}
	case _State180:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State219, 0}, true
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
	case _State181:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedTypedArgToProperParameterDef}, true
		}
	case _State182:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State220, 0}, true
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
	case _State183:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreTypedArgToProperParameterDef}, true
		}
	case _State184:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State221, 0}, true
		}
	case _State185:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAliasToTypeDef}, true
		}
	case _State186:
		switch symbolId {
		case ImplementsToken:
			return _Action{_ShiftAction, _State222, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}, true
		}
	case _State187:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State53, 0}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToVarPattern}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToVarPattern}, true
		case VarPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedToFieldVarPattern}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTuplePatternToVarPattern}, true
		}
	case _State188:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State53, 0}, true
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
	case _State189:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State53, 0}, true
		case TuplePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}, true
		}
	case _State190:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case ReturnTypeType:
			return _Action{_ShiftAction, _State223, 0}, true
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
	case _State191:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State114, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToParameterDefs}, true
		}
	case _State192:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State61, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperImportClauses}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}, true
		}
	case _State193:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State61, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperImportClauses}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}, true
		}
	case _State194:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToMapTypeExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State195:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToArrayTypeExpr}, true
		}
	case _State196:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperExplicitTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitTypeProperties}, true
		}
	case _State197:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperExplicitTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitTypeProperties}, true
		}
	case _State198:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatement}, true
		}
	case _State199:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State224, 0}, true
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
			return _Action{_ReduceAction, 0, _ReduceImproperToGenericArgumentList}, true
		}
	case _State200:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToCallExpr}, true
		}
	case _State201:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State202:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State203:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case SequenceExprType:
			return _Action{_ShiftAction, _State225, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State204:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OptionalSequenceExprType:
			return _Action{_ShiftAction, _State226, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State205:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWhileToLoopExprBody}, true
		}
	case _State206:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State126, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State227, 0}, true
		}
	case _State207:
		switch symbolId {
		case ElseToken:
			return _Action{_ShiftAction, _State228, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNoElseToIfExprBody}, true
		}
	case _State208:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitAddToProperExplicitEnumTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToExplicitEnumTypeProperties}, true
		}
	case _State209:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitAddToProperExplicitEnumTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true
		}
	case _State210:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitPairToProperExplicitEnumTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true
		}
	case _State211:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State100, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State101, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State98, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State106, 0}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStatementProperty}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceFieldDefToTypeProperty}, true
		case UnsafeStatementPropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStatementPropertyToTypeProperty}, true
		case TypePropertyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitPairToProperExplicitEnumTypeProperties}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceMethodSignatureToTypeProperty}, true
		}
	case _State212:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedTypedVarargToParameterDecl}, true
		}
	case _State213:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
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
	case _State214:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State163, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State166, 0}, true
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
	case _State215:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State163, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State161, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State166, 0}, true
		case ProperParameterDeclsType:
			return _Action{_ShiftAction, _State165, 0}, true
		case ParameterDeclsType:
			return _Action{_ShiftAction, _State229, 0}, true
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
	case _State216:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameter}, true
		}
	case _State217:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State176, 0}, true
		case GenericParameterType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperGenericParameterList}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToGenericParameterList}, true
		}
	case _State218:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State230, 0}, true
		}
	case _State219:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedTypedVarargToProperParameterDef}, true
		}
	case _State220:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreTypedVarargToProperParameterDef}, true
		}
	case _State221:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State231, 0}, true
		}
	case _State222:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State232, 0}, true
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
	case _State223:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAnonymousFuncExpr}, true
		}
	case _State224:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAddToProperGenericArgumentList}, true
		}
	case _State225:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State233, 0}, true
		}
	case _State226:
		switch symbolId {
		case SemicolonToken:
			return _Action{_ShiftAction, _State234, 0}, true
		}
	case _State227:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State228:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State235, 0}, true
		case LabelDeclToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseToIfExprBody}, true
		case IfExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMultiIfElseToIfExprBody}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}, true
		}
	case _State229:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State236, 0}, true
		}
	case _State230:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case ReturnTypeType:
			return _Action{_ShiftAction, _State237, 0}, true
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
	case _State231:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State113, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State114, 0}, true
		case ProperParameterDefsType:
			return _Action{_ShiftAction, _State128, 0}, true
		case ParameterDefsType:
			return _Action{_ShiftAction, _State238, 0}, true
		case ProperParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceProperParameterDefToParameterDef}, true
		case ParameterDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterDefToProperParameterDefs}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterDefs}, true
		}
	case _State232:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State109, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}, true
		}
	case _State233:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIteratorToLoopExprBody}, true
		}
	case _State234:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case VarOrLetType:
			return _Action{_ShiftAction, _State15, 0}, true
		case OptionalLabelDeclType:
			return _Action{_ShiftAction, _State58, 0}, true
		case OptionalSequenceExprType:
			return _Action{_ShiftAction, _State239, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State39, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State37, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State34, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State30, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State36, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State32, 0}, true
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
	case _State235:
		switch symbolId {
		case IfToken:
			return _Action{_ShiftAction, _State92, 0}, true
		case IfExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIfExpr}, true
		}
	case _State236:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
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
	case _State237:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToNamedFuncDef}, true
		}
	case _State238:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAction, _State240, 0}, true
		}
	case _State239:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State241, 0}, true
		}
	case _State240:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State42, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State40, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State41, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State43, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State45, 0}, true
		case ReturnTypeType:
			return _Action{_ShiftAction, _State242, 0}, true
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
	case _State241:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case StatementBlockType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToLoopExprBody}, true
		}
	case _State242:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State10, 0}, true
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
      PACKAGE -> State 11
      TYPE -> State 12
      FUNC -> State 9
      LBRACE -> State 10
      source -> State 5
      proper_definitions -> State 13
      var_decl_pattern -> State 14
      var_or_let -> State 15

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
      CASE -> State 16
      DEFAULT -> State 17
      IMPORT -> State 20
      UNSAFE -> State 24
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      statement -> State 6
      exprs -> State 31
      callback_op -> State 29
      jump_type -> State 33
      var_or_let -> State 15
      assign_pattern -> State 28
      optional_label_decl -> State 35
      sequence_expr -> State 38
      accessible_expr -> State 25
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      expr -> State 7
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
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
      binary_type_op -> State 109

  State 9:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 46
      LPAREN -> State 47

  State 10:
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
      CASE -> State 16
      DEFAULT -> State 17
      IMPORT -> State 20
      UNSAFE -> State 24
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      proper_statements -> State 48
      statements -> State 49
      exprs -> State 31
      callback_op -> State 29
      jump_type -> State 33
      var_or_let -> State 15
      assign_pattern -> State 28
      optional_label_decl -> State 35
      sequence_expr -> State 38
      accessible_expr -> State 25
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 11:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    ShiftAndReduce:
      statement_block -> [package_def]
    Goto:
      LBRACE -> State 10

  State 12:
    Kernel Items:
      type_def: TYPE.IDENTIFIER generic_parameters type_expr
      type_def: TYPE.IDENTIFIER generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE.IDENTIFIER ASSIGN type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 50

  State 13:
    Kernel Items:
      proper_definitions: proper_definitions.NEWLINES definition
      definitions: proper_definitions., *
      definitions: proper_definitions.NEWLINES
    Reduce:
      * -> [definitions]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 51

  State 14:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expr
    Reduce:
      * -> [definition]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 52

  State 15:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [var_pattern]
      UNDERSCORE -> [var_pattern]
      tuple_pattern -> [var_pattern]
    Goto:
      LPAREN -> State 53
      var_pattern -> State 54

  State 16:
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
      STRUCT -> State 23
      FUNC -> State 18
      VAR -> State 56
      LPAREN -> State 22
      LBRACKET -> State 21
      DOT -> State 55
      GREATER -> State 19
      case_patterns -> State 57
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 17:
    Kernel Items:
      statement: DEFAULT.COLON optional_simple_statement
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 59

  State 18:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 60

  State 19:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 20:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [import_statement]
    Goto:
      IDENTIFIER -> State 62
      UNDERSCORE -> State 64
      LPAREN -> State 63
      DOT -> State 61

  State 21:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 65

  State 22:
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
      IDENTIFIER -> State 67
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      COLON -> State 66
      GREATER -> State 19
      var_or_let -> State 15
      expr -> State 70
      optional_label_decl -> State 35
      proper_arguments -> State 71
      arguments -> State 68
      colon_expr -> State 69
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 23:
    Kernel Items:
      explicit_struct_type_expr: STRUCT.LPAREN explicit_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 72

  State 24:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LESS -> State 73

  State 25:
    Kernel Items:
      unary_op_assign_statement: accessible_expr.unary_op_assign
      binary_op_assign_statement: accessible_expr.binary_op_assign expr
      call_expr: accessible_expr.generic_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.postfix_unary_op
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_arguments]
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
      LBRACKET -> State 76
      DOT -> State 75
      DOLLAR_LBRACKET -> State 74
      binary_op_assign -> State 77
      generic_arguments -> State 78

  State 26:
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
      add_op -> State 79

  State 27:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AND -> State 80

  State 28:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 81

  State 29:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      call_expr -> State 83
      accessible_expr -> State 82
      initializable_type_expr -> State 32

  State 30:
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
      cmp_op -> State 84

  State 31:
    Kernel Items:
      expr_or_improper_struct_statement: exprs., *
      exprs: exprs.COMMA expr
    Reduce:
      * -> [expr_or_improper_struct_statement]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 85

  State 32:
    Kernel Items:
      initialize_expr: initializable_type_expr.LPAREN arguments RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 86

  State 33:
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
      STRUCT -> State 23
      FUNC -> State 18
      JUMP_LABEL -> State 87
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      exprs -> State 88
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 34:
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
      mul_op -> State 89

  State 35:
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
      IF -> State 92
      SWITCH -> State 93
      FOR -> State 91
      DO -> State 90
      LBRACE -> State 10

  State 36:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 94

  State 37:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      initializable_type_expr -> State 32

  State 38:
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

  State 39:
    Kernel Items:
      call_expr: accessible_expr.generic_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.postfix_unary_op
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_arguments]
    ShiftAndReduce:
      QUESTION -> [postfix_unary_op]
      EXCLAIM -> [postfix_unary_op]
      postfix_unary_op -> [postfix_unary_expr]
    Goto:
      LBRACKET -> State 76
      DOT -> State 75
      DOLLAR_LBRACKET -> State 74
      generic_arguments -> State 78

  State 40:
    Kernel Items:
      explicit_enum_type_expr: ENUM.LPAREN explicit_enum_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 95

  State 41:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 96

  State 42:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      generic_arguments -> [named_type_expr]
    Goto:
      DOT -> State 97
      DOLLAR_LBRACKET -> State 74

  State 43:
    Kernel Items:
      implicit_struct_type_expr: LPAREN.implicit_type_properties RPAREN
      implicit_enum_type_expr: LPAREN.implicit_enum_type_properties RPAREN
    Reduce:
      * -> [implicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106
      type_property -> State 107
      proper_implicit_type_properties -> State 105
      implicit_type_properties -> State 103
      proper_implicit_enum_type_properties -> State 104
      implicit_enum_type_properties -> State 102

  State 44:
    Kernel Items:
      trait_type_expr: TRAIT.LPAREN explicit_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 108

  State 45:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45

  State 46:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expr
    Reduce:
      * -> [generic_parameters]
    ShiftAndReduce:
      (nil)
    Goto:
      DOLLAR_LBRACKET -> State 111
      ASSIGN -> State 110
      generic_parameters -> State 112

  State 47:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 114
      parameter_def -> State 115

  State 48:
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
      NEWLINES -> State 116
      SEMICOLON -> State 117

  State 49:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACE -> [statement_block]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      type_def: TYPE IDENTIFIER.generic_parameters type_expr
      type_def: TYPE IDENTIFIER.generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [generic_parameters]
    ShiftAndReduce:
      (nil)
    Goto:
      DOLLAR_LBRACKET -> State 111
      ASSIGN -> State 118
      generic_parameters -> State 119

  State 51:
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
      PACKAGE -> State 11
      TYPE -> State 12
      FUNC -> State 9
      LBRACE -> State 10
      var_decl_pattern -> State 14
      var_or_let -> State 15

  State 52:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 53:
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
      IDENTIFIER -> State 120
      LPAREN -> State 53
      field_var_patterns -> State 121

  State 54:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 122

  State 55:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
      case_pattern: DOT.IDENTIFIER
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 123

  State 56:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    ShiftAndReduce:
      (nil)
    Goto:
      DOT -> State 124

  State 57:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 126
      COLON -> State 125

  State 58:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [block_expr]
    Goto:
      LBRACE -> State 10

  State 59:
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
      UNSAFE -> State 24
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      exprs -> State 31
      callback_op -> State 29
      jump_type -> State 33
      var_or_let -> State 15
      assign_pattern -> State 28
      optional_label_decl -> State 35
      sequence_expr -> State 38
      accessible_expr -> State 25
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 60:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 114
      proper_parameter_defs -> State 128
      parameter_defs -> State 127

  State 61:
    Kernel Items:
      import_clause: DOT.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      import_clause: IDENTIFIER.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 62
      UNDERSCORE -> State 64
      DOT -> State 61
      proper_import_clauses -> State 130
      import_clauses -> State 129

  State 64:
    Kernel Items:
      import_clause: UNDERSCORE.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 65:
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
      COMMA -> State 132
      COLON -> State 131
      binary_type_op -> State 109

  State 66:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 67:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expr
      named_expr: IDENTIFIER., *
    Reduce:
      * -> [named_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 133

  State 68:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_struct_expr]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [argument]
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 134

  State 70:
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
      COLON -> State 135

  State 71:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., *
      arguments: proper_arguments.COMMA
    Reduce:
      * -> [arguments]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 136

  State 72:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN.explicit_type_properties RPAREN
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106
      proper_explicit_type_properties -> State 138
      explicit_type_properties -> State 137

  State 73:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139

  State 74:
    Kernel Items:
      generic_arguments: DOLLAR_LBRACKET.generic_argument_list RBRACKET
    Reduce:
      * -> [generic_argument_list]
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 142
      proper_generic_argument_list -> State 141
      generic_argument_list -> State 140

  State 75:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [access_expr]
    Goto:
      (nil)

  State 76:
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
      IDENTIFIER -> State 67
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      COLON -> State 66
      GREATER -> State 19
      var_or_let -> State 15
      expr -> State 70
      optional_label_decl -> State 35
      argument -> State 143
      colon_expr -> State 69
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 77:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 78:
    Kernel Items:
      call_expr: accessible_expr generic_arguments.LPAREN arguments RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 144

  State 79:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 145
      initializable_type_expr -> State 32

  State 80:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 146
      initializable_type_expr -> State 32

  State 81:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 82:
    Kernel Items:
      call_expr: accessible_expr.generic_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      (nil)
    Goto:
      LBRACKET -> State 76
      DOT -> State 75
      DOLLAR_LBRACKET -> State 74
      generic_arguments -> State 78

  State 83:
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

  State 84:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 147
      initializable_type_expr -> State 32

  State 85:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 86:
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
      IDENTIFIER -> State 67
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      COLON -> State 66
      GREATER -> State 19
      var_or_let -> State 15
      expr -> State 70
      optional_label_decl -> State 35
      proper_arguments -> State 71
      arguments -> State 148
      colon_expr -> State 69
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 87:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      exprs -> State 149
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 88:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type exprs., *
    Reduce:
      * -> [jump_statement]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 85

  State 89:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      initializable_type_expr -> State 32

  State 90:
    Kernel Items:
      loop_expr_body: DO.statement_block
      loop_expr_body: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LBRACE -> State 10
      statement_block -> State 150

  State 91:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      assign_pattern -> State 151
      optional_label_decl -> State 58
      sequence_expr -> State 153
      for_assignment -> State 152
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 92:
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
      CASE -> State 154
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      condition -> State 155
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 93:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      sequence_expr -> State 156
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 94:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 157
      initializable_type_expr -> State 32

  State 95:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN.explicit_enum_type_properties RPAREN
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
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106
      type_property -> State 160
      proper_explicit_enum_type_properties -> State 159
      explicit_enum_type_properties -> State 158

  State 96:
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
      IDENTIFIER -> State 162
      UNDERSCORE -> State 163
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      ELLIPSIS -> State 161
      prefix_unary_type_op -> State 45
      type_expr -> State 166
      proper_parameter_decls -> State 165
      parameter_decls -> State 164

  State 97:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_arguments
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 167

  State 98:
    Kernel Items:
      type_property: DEFAULT.field_def
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
      field_def -> [type_property]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 100
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 99:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      LPAREN -> State 96

  State 100:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
      field_def: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_arguments]
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
      generic_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      DOT -> State 169
      DOLLAR_LBRACKET -> State 74
      prefix_unary_type_op -> State 45
      type_expr -> State 170

  State 101:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      type_property: UNDERSCORE.type_expr
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 171

  State 102:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      proper_implicit_enum_type_properties: proper_implicit_enum_type_properties.OR type_property
      implicit_enum_type_properties: proper_implicit_enum_type_properties., *
      implicit_enum_type_properties: proper_implicit_enum_type_properties.NEWLINES
    Reduce:
      * -> [implicit_enum_type_properties]
    ShiftAndReduce:
      NEWLINES -> [implicit_enum_type_properties]
    Goto:
      OR -> State 172

  State 105:
    Kernel Items:
      proper_implicit_type_properties: proper_implicit_type_properties.COMMA type_property
      implicit_type_properties: proper_implicit_type_properties., *
      implicit_type_properties: proper_implicit_type_properties.COMMA
    Reduce:
      * -> [implicit_type_properties]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 173

  State 106:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: type_expr., *
    Reduce:
      * -> [field_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 107:
    Kernel Items:
      proper_implicit_type_properties: type_property., *
      proper_implicit_enum_type_properties: type_property.OR type_property
    Reduce:
      * -> [proper_implicit_type_properties]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 174

  State 108:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN.explicit_type_properties RPAREN
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106
      proper_explicit_type_properties -> State 138
      explicit_type_properties -> State 175

  State 109:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45

  State 110:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 111:
    Kernel Items:
      generic_parameters: DOLLAR_LBRACKET.generic_parameter_list RBRACKET
    Reduce:
      * -> [generic_parameter_list]
    ShiftAndReduce:
      generic_parameter -> [proper_generic_parameter_list]
    Goto:
      IDENTIFIER -> State 176
      proper_generic_parameter_list -> State 178
      generic_parameter_list -> State 177

  State 112:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 179

  State 113:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      ELLIPSIS -> State 180
      prefix_unary_type_op -> State 45
      type_expr -> State 181

  State 114:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      ELLIPSIS -> State 182
      prefix_unary_type_op -> State 45
      type_expr -> State 183

  State 115:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 184

  State 116:
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
      CASE -> State 16
      DEFAULT -> State 17
      IMPORT -> State 20
      UNSAFE -> State 24
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      exprs -> State 31
      callback_op -> State 29
      jump_type -> State 33
      var_or_let -> State 15
      assign_pattern -> State 28
      optional_label_decl -> State 35
      sequence_expr -> State 38
      accessible_expr -> State 25
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 117:
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
      CASE -> State 16
      DEFAULT -> State 17
      IMPORT -> State 20
      UNSAFE -> State 24
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      exprs -> State 31
      callback_op -> State 29
      jump_type -> State 33
      var_or_let -> State 15
      assign_pattern -> State 28
      optional_label_decl -> State 35
      sequence_expr -> State 38
      accessible_expr -> State 25
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 118:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 185

  State 119:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameters.type_expr
      type_def: TYPE IDENTIFIER generic_parameters.type_expr IMPLEMENTS type_expr
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 186

  State 120:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 187

  State 121:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [tuple_pattern]
    Goto:
      COMMA -> State 188

  State 122:
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
      binary_type_op -> State 109

  State 123:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
      case_pattern: DOT IDENTIFIER., *
    Reduce:
      * -> [case_pattern]
    ShiftAndReduce:
      implicit_struct_expr -> [case_pattern]
    Goto:
      LPAREN -> State 22

  State 124:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 189

  State 125:
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
      UNSAFE -> State 24
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      exprs -> State 31
      callback_op -> State 29
      jump_type -> State 33
      var_or_let -> State 15
      assign_pattern -> State 28
      optional_label_decl -> State 35
      sequence_expr -> State 38
      accessible_expr -> State 25
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 126:
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
      STRUCT -> State 23
      FUNC -> State 18
      VAR -> State 56
      LPAREN -> State 22
      LBRACKET -> State 21
      DOT -> State 55
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 127:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 190

  State 128:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., *
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 191

  State 129:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [import_statement]
    Goto:
      (nil)

  State 130:
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
      NEWLINES -> State 193
      COMMA -> State 192

  State 131:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 194

  State 132:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 195

  State 133:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 134:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 135:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 35
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 136:
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
      IDENTIFIER -> State 67
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      COLON -> State 66
      GREATER -> State 19
      var_or_let -> State 15
      expr -> State 70
      optional_label_decl -> State 35
      colon_expr -> State 69
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 137:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [explicit_struct_type_expr]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      proper_explicit_type_properties: proper_explicit_type_properties.NEWLINES type_property
      proper_explicit_type_properties: proper_explicit_type_properties.COMMA type_property
      explicit_type_properties: proper_explicit_type_properties., *
      explicit_type_properties: proper_explicit_type_properties.NEWLINES
      explicit_type_properties: proper_explicit_type_properties.COMMA
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 197
      COMMA -> State 196

  State 139:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      GREATER -> State 198

  State 140:
    Kernel Items:
      generic_arguments: DOLLAR_LBRACKET generic_argument_list.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [generic_arguments]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      proper_generic_argument_list: proper_generic_argument_list.COMMA type_expr
      generic_argument_list: proper_generic_argument_list., *
      generic_argument_list: proper_generic_argument_list.COMMA
    Reduce:
      * -> [generic_argument_list]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 199

  State 142:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_generic_argument_list: type_expr., *
    Reduce:
      * -> [proper_generic_argument_list]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 143:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [index_expr]
    Goto:
      (nil)

  State 144:
    Kernel Items:
      call_expr: accessible_expr generic_arguments LPAREN.arguments RPAREN
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
      IDENTIFIER -> State 67
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      COLON -> State 66
      GREATER -> State 19
      var_or_let -> State 15
      expr -> State 70
      optional_label_decl -> State 35
      proper_arguments -> State 71
      arguments -> State 200
      colon_expr -> State 69
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 145:
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
      mul_op -> State 89

  State 146:
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
      cmp_op -> State 84

  State 147:
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
      add_op -> State 79

  State 148:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [initialize_expr]
    Goto:
      (nil)

  State 149:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type JUMP_LABEL exprs., *
    Reduce:
      * -> [jump_statement]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 85

  State 150:
    Kernel Items:
      loop_expr_body: DO statement_block., *
      loop_expr_body: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr_body]
    ShiftAndReduce:
      (nil)
    Goto:
      FOR -> State 201

  State 151:
    Kernel Items:
      loop_expr_body: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IN -> State 203
      ASSIGN -> State 202

  State 152:
    Kernel Items:
      loop_expr_body: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      SEMICOLON -> State 204

  State 153:
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
      DO -> State 205

  State 154:
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
      STRUCT -> State 23
      FUNC -> State 18
      VAR -> State 56
      LPAREN -> State 22
      LBRACKET -> State 21
      DOT -> State 55
      GREATER -> State 19
      case_patterns -> State 206
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 155:
    Kernel Items:
      if_expr_body: IF condition.statement_block
      if_expr_body: IF condition.statement_block ELSE statement_block
      if_expr_body: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LBRACE -> State 10
      statement_block -> State 207

  State 156:
    Kernel Items:
      switch_expr: optional_label_decl SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [switch_expr]
    Goto:
      LBRACE -> State 10

  State 157:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AND -> State 80

  State 158:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      proper_explicit_enum_type_properties: proper_explicit_enum_type_properties.OR type_property
      proper_explicit_enum_type_properties: proper_explicit_enum_type_properties.NEWLINES type_property
      explicit_enum_type_properties: proper_explicit_enum_type_properties., *
      explicit_enum_type_properties: proper_explicit_enum_type_properties.NEWLINES
    Reduce:
      * -> [explicit_enum_type_properties]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 208
      OR -> State 209

  State 160:
    Kernel Items:
      proper_explicit_enum_type_properties: type_property.OR type_property
      proper_explicit_enum_type_properties: type_property.NEWLINES type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 210
      OR -> State 211

  State 161:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 212

  State 162:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
      proper_parameter_def: IDENTIFIER.type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS
    Reduce:
      * -> [generic_arguments]
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
      generic_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      DOT -> State 169
      DOLLAR_LBRACKET -> State 74
      ELLIPSIS -> State 180
      prefix_unary_type_op -> State 45
      type_expr -> State 181

  State 163:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      ELLIPSIS -> State 182
      prefix_unary_type_op -> State 45
      type_expr -> State 183

  State 164:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 213

  State 165:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., *
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      * -> [parameter_decls]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 214

  State 166:
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
      binary_type_op -> State 109

  State 167:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_arguments
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      generic_arguments -> [named_type_expr]
    Goto:
      DOLLAR_LBRACKET -> State 74

  State 168:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 215

  State 169:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 167

  State 170:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [field_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 171:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_property: UNDERSCORE type_expr., *
    Reduce:
      * -> [type_property]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 172:
    Kernel Items:
      proper_implicit_enum_type_properties: proper_implicit_enum_type_properties OR.type_property
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
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_implicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 173:
    Kernel Items:
      proper_implicit_type_properties: proper_implicit_type_properties COMMA.type_property
      implicit_type_properties: proper_implicit_type_properties COMMA., *
    Reduce:
      * -> [implicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_implicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 174:
    Kernel Items:
      proper_implicit_enum_type_properties: type_property OR.type_property
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
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_implicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 175:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [trait_type_expr]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      generic_parameter: IDENTIFIER., *
      generic_parameter: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_parameter]
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 216

  State 177:
    Kernel Items:
      generic_parameters: DOLLAR_LBRACKET generic_parameter_list.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [generic_parameters]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      proper_generic_parameter_list: proper_generic_parameter_list.COMMA generic_parameter
      generic_parameter_list: proper_generic_parameter_list., *
      generic_parameter_list: proper_generic_parameter_list.COMMA
    Reduce:
      * -> [generic_parameter_list]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 217

  State 179:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 114
      proper_parameter_defs -> State 128
      parameter_defs -> State 218

  State 180:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 219

  State 181:
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
      binary_type_op -> State 109

  State 182:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 220

  State 183:
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
      binary_type_op -> State 109

  State 184:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 221

  State 185:
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
      binary_type_op -> State 109

  State 186:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameters type_expr., *
      type_def: TYPE IDENTIFIER generic_parameters type_expr.IMPLEMENTS type_expr
    Reduce:
      * -> [type_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      IMPLEMENTS -> State 222
      binary_type_op -> State 109

  State 187:
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
      LPAREN -> State 53

  State 188:
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
      IDENTIFIER -> State 120
      LPAREN -> State 53

  State 189:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      tuple_pattern -> [case_pattern]
    Goto:
      LPAREN -> State 53

  State 190:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      return_type -> State 223

  State 191:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., *
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 114

  State 192:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., *
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 62
      UNDERSCORE -> State 64
      DOT -> State 61

  State 193:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., *
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 62
      UNDERSCORE -> State 64
      DOT -> State 61

  State 194:
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
      binary_type_op -> State 109

  State 195:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [array_type_expr]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      proper_explicit_type_properties: proper_explicit_type_properties COMMA.type_property
      explicit_type_properties: proper_explicit_type_properties COMMA., *
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 197:
    Kernel Items:
      proper_explicit_type_properties: proper_explicit_type_properties NEWLINES.type_property
      explicit_type_properties: proper_explicit_type_properties NEWLINES., *
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 198:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [unsafe_statement]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      proper_generic_argument_list: proper_generic_argument_list COMMA.type_expr
      generic_argument_list: proper_generic_argument_list COMMA., *
    Reduce:
      * -> [generic_argument_list]
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 224

  State 200:
    Kernel Items:
      call_expr: accessible_expr generic_arguments LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [call_expr]
    Goto:
      (nil)

  State 201:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 202:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 203:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      sequence_expr -> State 225
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 204:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      optional_sequence_expr -> State 226
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 205:
    Kernel Items:
      loop_expr_body: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [loop_expr_body]
    Goto:
      LBRACE -> State 10

  State 206:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 126
      ASSIGN -> State 227

  State 207:
    Kernel Items:
      if_expr_body: IF condition statement_block., *
      if_expr_body: IF condition statement_block.ELSE statement_block
      if_expr_body: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr_body]
    ShiftAndReduce:
      (nil)
    Goto:
      ELSE -> State 228

  State 208:
    Kernel Items:
      proper_explicit_enum_type_properties: proper_explicit_enum_type_properties NEWLINES.type_property
      explicit_enum_type_properties: proper_explicit_enum_type_properties NEWLINES., *
    Reduce:
      * -> [explicit_enum_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_NEG -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      PARSE_ERROR -> [parse_error_type_expr]
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 209:
    Kernel Items:
      proper_explicit_enum_type_properties: proper_explicit_enum_type_properties OR.type_property
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
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 210:
    Kernel Items:
      proper_explicit_enum_type_properties: type_property NEWLINES.type_property
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
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 211:
    Kernel Items:
      proper_explicit_enum_type_properties: type_property OR.type_property
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
      unsafe_statement -> [unsafe_statement_property]
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
      field_def -> [type_property]
      unsafe_statement_property -> [type_property]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_type_expr -> [atom_type_expr]
      method_signature -> [type_property]
    Goto:
      IDENTIFIER -> State 100
      UNDERSCORE -> State 101
      DEFAULT -> State 98
      UNSAFE -> State 24
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 99
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 106

  State 212:
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
      binary_type_op -> State 109

  State 213:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45

  State 214:
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
      IDENTIFIER -> State 162
      UNDERSCORE -> State 163
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      ELLIPSIS -> State 161
      prefix_unary_type_op -> State 45
      type_expr -> State 166

  State 215:
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
      IDENTIFIER -> State 162
      UNDERSCORE -> State 163
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      ELLIPSIS -> State 161
      prefix_unary_type_op -> State 45
      type_expr -> State 166
      proper_parameter_decls -> State 165
      parameter_decls -> State 229

  State 216:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      generic_parameter: IDENTIFIER type_expr., *
    Reduce:
      * -> [generic_parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 217:
    Kernel Items:
      proper_generic_parameter_list: proper_generic_parameter_list COMMA.generic_parameter
      generic_parameter_list: proper_generic_parameter_list COMMA., *
    Reduce:
      * -> [generic_parameter_list]
    ShiftAndReduce:
      generic_parameter -> [proper_generic_parameter_list]
    Goto:
      IDENTIFIER -> State 176

  State 218:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 230

  State 219:
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
      binary_type_op -> State 109

  State 220:
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
      binary_type_op -> State 109

  State 221:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 231

  State 222:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameters type_expr IMPLEMENTS.type_expr
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      type_expr -> State 232

  State 223:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [anonymous_func_expr]
    Goto:
      LBRACE -> State 10

  State 224:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_generic_argument_list: proper_generic_argument_list COMMA type_expr., *
    Reduce:
      * -> [proper_generic_argument_list]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 225:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      DO -> State 233

  State 226:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      SEMICOLON -> State 234

  State 227:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 228:
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
      LBRACE -> State 10
      optional_label_decl -> State 235

  State 229:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 236

  State 230:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      return_type -> State 237

  State 231:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      * -> [parameter_defs]
    ShiftAndReduce:
      proper_parameter_def -> [parameter_def]
      parameter_def -> [proper_parameter_defs]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 114
      proper_parameter_defs -> State 128
      parameter_defs -> State 238

  State 232:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameters type_expr IMPLEMENTS type_expr., *
    Reduce:
      * -> [type_def]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 109

  State 233:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [loop_expr_body]
    Goto:
      LBRACE -> State 10

  State 234:
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
      STRUCT -> State 23
      FUNC -> State 18
      LPAREN -> State 22
      LBRACKET -> State 21
      GREATER -> State 19
      var_or_let -> State 15
      optional_label_decl -> State 58
      optional_sequence_expr -> State 239
      accessible_expr -> State 39
      prefix_unary_op -> State 37
      mul_expr -> State 34
      add_expr -> State 26
      cmp_expr -> State 30
      and_expr -> State 27
      or_expr -> State 36
      initializable_type_expr -> State 32

  State 235:
    Kernel Items:
      if_expr: optional_label_decl.if_expr_body
    Reduce:
      (nil)
    ShiftAndReduce:
      if_expr_body -> [if_expr]
    Goto:
      IF -> State 92

  State 236:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45

  State 237:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [named_func_def]
    Goto:
      LBRACE -> State 10

  State 238:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      RPAREN -> State 240

  State 239:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      DO -> State 241

  State 240:
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
      IDENTIFIER -> State 42
      STRUCT -> State 23
      ENUM -> State 40
      TRAIT -> State 44
      FUNC -> State 41
      LPAREN -> State 43
      LBRACKET -> State 21
      prefix_unary_type_op -> State 45
      return_type -> State 242

  State 241:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [loop_expr_body]
    Goto:
      LBRACE -> State 10

  State 242:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    ShiftAndReduce:
      statement_block -> [named_func_def]
    Goto:
      LBRACE -> State 10

Number of states: 242
Number of shift actions: 1367
Number of reduce actions: 193
Number of shift-and-reduce actions: 3300
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 5056
Number of unoptimized shift actions: 45411
Number of unoptimized reduce actions: 34526
*/
