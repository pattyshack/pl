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
	GlobalVarAssignmentToDefinition(VarDeclPattern_ GenericSymbol, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)

	// 72:2: definition -> statement_block: ...
	StatementBlockToDefinition(StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 76:2: definition -> COMMENT_GROUPS: ...
	CommentGroupsToDefinition(CommentGroups_ TokenValue) (SourceDefinition, error)
}

type StatementBlockReducer interface {
	// 96:19: statement_block -> ...
	ToStatementBlock(Lbrace_ TokenValue, Statements_ *StatementList, Rbrace_ TokenValue) (GenericSymbol, error)
}

type ProperStatementsReducer interface {
	// 99:2: proper_statements -> add_implicit: ...
	AddImplicitToProperStatements(ProperStatements_ *StatementList, Newlines_ TokenCount, Statement_ Statement) (*StatementList, error)

	// 100:2: proper_statements -> add_explicit: ...
	AddExplicitToProperStatements(ProperStatements_ *StatementList, Semicolon_ TokenValue, Statement_ Statement) (*StatementList, error)

	// 101:2: proper_statements -> statement: ...
	StatementToProperStatements(Statement_ Statement) (*StatementList, error)
}

type StatementsReducer interface {

	// 105:2: statements -> improper_implicit: ...
	ImproperImplicitToStatements(ProperStatements_ *StatementList, Newlines_ TokenCount) (*StatementList, error)

	// 106:2: statements -> improper_explicit: ...
	ImproperExplicitToStatements(ProperStatements_ *StatementList, Semicolon_ TokenValue) (*StatementList, error)

	// 107:2: statements -> nil: ...
	NilToStatements() (*StatementList, error)
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

type ExprOrImproperStructStatementReducer interface {
	// 156:48: expr_or_improper_struct_statement -> ...
	ToExprOrImproperStructStatement(Exprs_ GenericSymbol) (Statement, error)
}

type ExprsReducer interface {
	// 159:2: exprs -> expr: ...
	ExprToExprs(Expr_ Expression) (GenericSymbol, error)

	// 160:2: exprs -> add: ...
	AddToExprs(Exprs_ GenericSymbol, Comma_ TokenValue, Expr_ Expression) (GenericSymbol, error)
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
	UnlabeledValuedToJumpStatement(JumpType_ TokenValue, Exprs_ GenericSymbol) (Statement, error)

	// 203:2: jump_statement -> labeled_no_value: ...
	LabeledNoValueToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue) (Statement, error)

	// 204:2: jump_statement -> labeled_valued: ...
	LabeledValuedToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue, Exprs_ GenericSymbol) (Statement, error)
}

type FallthroughStatementReducer interface {
	// 211:36: fallthrough_statement -> ...
	ToFallthroughStatement(Fallthrough_ TokenValue) (Statement, error)
}

type AssignStatementReducer interface {
	// 217:31: assign_statement -> ...
	ToAssignStatement(AssignPattern_ GenericSymbol, Assign_ TokenValue, Expr_ Expression) (Statement, error)
}

type UnaryOpAssignStatementReducer interface {
	// 219:40: unary_op_assign_statement -> ...
	ToUnaryOpAssignStatement(AccessibleExpr_ Expression, UnaryOpAssign_ TokenValue) (Statement, error)
}

type BinaryOpAssignStatementReducer interface {
	// 225:41: binary_op_assign_statement -> ...
	ToBinaryOpAssignStatement(AccessibleExpr_ Expression, BinaryOpAssign_ TokenValue, Expr_ Expression) (Statement, error)
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
	ToVarDeclPattern(VarOrLet_ TokenValue, VarPattern_ GenericSymbol, OptionalTypeExpr_ TypeExpression) (GenericSymbol, error)
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

type OptionalTypeExprReducer interface {

	// 300:2: optional_type_expr -> nil: ...
	NilToOptionalTypeExpr() (TypeExpression, error)
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

type ExprReducer interface {
	// 325:2: expr -> if_expr: ...
	IfExprToExpr(OptionalLabelDecl_ TokenValue, IfExpr_ Expression) (Expression, error)

	// 326:2: expr -> switch_expr: ...
	SwitchExprToExpr(OptionalLabelDecl_ TokenValue, SwitchExpr_ Expression) (Expression, error)

	// 327:2: expr -> loop_expr: ...
	LoopExprToExpr(OptionalLabelDecl_ TokenValue, LoopExpr_ Expression) (Expression, error)
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
	ToCallExpr(AccessibleExpr_ Expression, GenericTypeArguments_ *TypeArgumentList, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type GenericTypeArgumentsReducer interface {
	// 421:2: generic_type_arguments -> binding: ...
	BindingToGenericTypeArguments(DollarLbracket_ TokenValue, TypeArguments_ *TypeArgumentList, Rbracket_ TokenValue) (*TypeArgumentList, error)

	// 422:2: generic_type_arguments -> nil: ...
	NilToGenericTypeArguments() (*TypeArgumentList, error)
}

type ProperTypeArgumentsReducer interface {
	// 425:2: proper_type_arguments -> add: ...
	AddToProperTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue, TypeExpr_ TypeExpression) (*TypeArgumentList, error)

	// 426:2: proper_type_arguments -> type_expr: ...
	TypeExprToProperTypeArguments(TypeExpr_ TypeExpression) (*TypeArgumentList, error)
}

type TypeArgumentsReducer interface {

	// 430:2: type_arguments -> improper: ...
	ImproperToTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue) (*TypeArgumentList, error)

	// 431:2: type_arguments -> nil: ...
	NilToTypeArguments() (*TypeArgumentList, error)
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
	PositionalToArgument(Expr_ Expression) (*Argument, error)

	// 444:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ *ColonExpr) (*Argument, error)

	// 445:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (*Argument, error)

	// 449:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expr_ Expression, Ellipsis_ TokenValue) (*Argument, error)

	// 452:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ TokenValue) (*Argument, error)
}

type ColonExprReducer interface {
	// 456:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ TokenValue) (*ColonExpr, error)

	// 457:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expr_ Expression, Colon_ TokenValue) (*ColonExpr, error)

	// 458:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)

	// 459:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expr_ Expression, Colon_ TokenValue, Expr_2 Expression) (*ColonExpr, error)

	// 460:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue) (*ColonExpr, error)

	// 461:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)
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
	// 625:38: inferred_type_expr -> ...
	ToInferredTypeExpr(Dot_ TokenValue) (TypeExpression, error)
}

type ParseErrorTypeExprReducer interface {
	// 627:41: parse_error_type_expr -> ...
	ToParseErrorTypeExpr(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixedUnaryTypeExprReducer interface {
	// 637:2: prefixed_unary_type_expr -> ...
	ToPrefixedUnaryTypeExpr(PrefixUnaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type BinaryTypeExprReducer interface {
	// 653:2: binary_type_expr -> ...
	ToBinaryTypeExpr(TypeExpr_ TypeExpression, BinaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 661:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, TypeExpr_ TypeExpression) (SourceDefinition, error)

	// 662:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, TypeExpr_ TypeExpression, Implements_ TokenValue, TypeExpr_2 TypeExpression) (SourceDefinition, error)

	// 663:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, TypeExpr_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 671:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 672:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type ProperGenericParameterDefsReducer interface {
	// 675:2: proper_generic_parameter_defs -> add: ...
	AddToProperGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 676:2: proper_generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToProperGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {

	// 680:2: generic_parameter_defs -> improper: ...
	ImproperToGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 681:2: generic_parameter_defs -> nil: ...
	NilToGenericParameterDefs() (GenericSymbol, error)
}

type OptionalGenericParametersReducer interface {
	// 684:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ TokenValue, GenericParameterDefs_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 685:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)
}

type FieldDefReducer interface {
	// 692:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 693:2: field_def -> implicit: ...
	ImplicitToFieldDef(TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 694:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type ProperImplicitFieldDefsReducer interface {
	// 697:2: proper_implicit_field_defs -> add: ...
	AddToProperImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 698:2: proper_implicit_field_defs -> field_def: ...
	FieldDefToProperImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {

	// 702:2: implicit_field_defs -> improper: ...
	ImproperToImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 703:2: implicit_field_defs -> nil: ...
	NilToImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructTypeExprReducer interface {
	// 706:2: implicit_struct_type_expr -> ...
	ToImplicitStructTypeExpr(Lparen_ TokenValue, ImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitFieldDefsReducer interface {
	// 709:2: proper_explicit_field_defs -> add_implicit: ...
	AddImplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 710:2: proper_explicit_field_defs -> add_explicit: ...
	AddExplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 711:2: proper_explicit_field_defs -> field_def: ...
	FieldDefToProperExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitFieldDefsReducer interface {

	// 715:2: explicit_field_defs -> improper_implicit: ...
	ImproperImplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 716:2: explicit_field_defs -> improper_explicit: ...
	ImproperExplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 717:2: explicit_field_defs -> nil: ...
	NilToExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructTypeExprReducer interface {
	// 720:2: explicit_struct_type_expr -> ...
	ToExplicitStructTypeExpr(Struct_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type EnumValueDefReducer interface {
	// 728:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 729:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ GenericSymbol, Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {
	// 741:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 742:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumTypeExprReducer interface {
	// 744:43: implicit_enum_type_expr -> ...
	ToImplicitEnumTypeExpr(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitEnumValueDefsReducer interface {
	// 747:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 748:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 749:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 750:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumTypeExprReducer interface {
	// 753:2: explicit_enum_type_expr -> ...
	ToExplicitEnumTypeExpr(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitPropertyReducer interface {
	// 760:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 761:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ GenericSymbol) (GenericSymbol, error)
}

type ProperTraitPropertiesReducer interface {
	// 764:2: proper_trait_properties -> implicit: ...
	ImplicitToProperTraitProperties(ProperTraitProperties_ GenericSymbol, Newlines_ TokenCount, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 765:2: proper_trait_properties -> explicit: ...
	ExplicitToProperTraitProperties(ProperTraitProperties_ GenericSymbol, Comma_ TokenValue, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 766:2: proper_trait_properties -> trait_property: ...
	TraitPropertyToProperTraitProperties(TraitProperty_ GenericSymbol) (GenericSymbol, error)
}

type TraitPropertiesReducer interface {

	// 770:2: trait_properties -> improper_implicit: ...
	ImproperImplicitToTraitProperties(ProperTraitProperties_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 771:2: trait_properties -> improper_explicit: ...
	ImproperExplicitToTraitProperties(ProperTraitProperties_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 772:2: trait_properties -> nil: ...
	NilToTraitProperties() (GenericSymbol, error)
}

type TraitTypeExprReducer interface {
	// 774:35: trait_type_expr -> ...
	ToTraitTypeExpr(Trait_ TokenValue, Lparen_ TokenValue, TraitProperties_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 782:2: return_type -> returnable_type_expr: ...
	ReturnableTypeExprToReturnType(ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)

	// 783:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ParameterDeclReducer interface {
	// 786:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 787:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 788:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 789:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDeclsReducer interface {
	// 792:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue, ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 793:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDeclsReducer interface {

	// 797:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 798:2: parameter_decls -> nil: ...
	NilToParameterDecls() (GenericSymbol, error)
}

type FuncTypeExprReducer interface {
	// 801:2: func_type_expr -> ...
	ToFuncTypeExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 812:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefReducer interface {
	// 818:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 819:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (GenericSymbol, error)

	// 820:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 821:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDefsReducer interface {
	// 824:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue, ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 825:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDefsReducer interface {

	// 829:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 830:2: parameter_defs -> nil: ...
	NilToParameterDefs() (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 833:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 834:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ GenericSymbol, Rparen_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_2 TokenValue, ParameterDefs_ GenericSymbol, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 835:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 839:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 851:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 852:2: package_def -> with_spec: ...
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
	ProperImplicitFieldDefsReducer
	ImplicitFieldDefsReducer
	ImplicitStructTypeExprReducer
	ProperExplicitFieldDefsReducer
	ExplicitFieldDefsReducer
	ExplicitStructTypeExprReducer
	EnumValueDefReducer
	ImplicitEnumValueDefsReducer
	ImplicitEnumTypeExprReducer
	ExplicitEnumValueDefsReducer
	ExplicitEnumTypeExprReducer
	TraitPropertyReducer
	ProperTraitPropertiesReducer
	TraitPropertiesReducer
	TraitTypeExprReducer
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
	case EnumValueDefType:
		return "enum_value_def"
	case ImplicitEnumValueDefsType:
		return "implicit_enum_value_defs"
	case ImplicitEnumTypeExprType:
		return "implicit_enum_type_expr"
	case ExplicitEnumValueDefsType:
		return "explicit_enum_value_defs"
	case ExplicitEnumTypeExprType:
		return "explicit_enum_type_expr"
	case TraitPropertyType:
		return "trait_property"
	case ProperTraitPropertiesType:
		return "proper_trait_properties"
	case TraitPropertiesType:
		return "trait_properties"
	case TraitTypeExprType:
		return "trait_type_expr"
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
	IdentifierExprType                = SymbolId(400)
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
	ProperImplicitFieldDefsType       = SymbolId(445)
	ImplicitFieldDefsType             = SymbolId(446)
	ImplicitStructTypeExprType        = SymbolId(447)
	ProperExplicitFieldDefsType       = SymbolId(448)
	ExplicitFieldDefsType             = SymbolId(449)
	ExplicitStructTypeExprType        = SymbolId(450)
	EnumValueDefType                  = SymbolId(451)
	ImplicitEnumValueDefsType         = SymbolId(452)
	ImplicitEnumTypeExprType          = SymbolId(453)
	ExplicitEnumValueDefsType         = SymbolId(454)
	ExplicitEnumTypeExprType          = SymbolId(455)
	TraitPropertyType                 = SymbolId(456)
	ProperTraitPropertiesType         = SymbolId(457)
	TraitPropertiesType               = SymbolId(458)
	TraitTypeExprType                 = SymbolId(459)
	ReturnTypeType                    = SymbolId(460)
	ParameterDeclType                 = SymbolId(461)
	ProperParameterDeclsType          = SymbolId(462)
	ParameterDeclsType                = SymbolId(463)
	FuncTypeExprType                  = SymbolId(464)
	MethodSignatureType               = SymbolId(465)
	ParameterDefType                  = SymbolId(466)
	ProperParameterDefsType           = SymbolId(467)
	ParameterDefsType                 = SymbolId(468)
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
	_ReduceToSource                                         = _ReduceType(1)
	_ReduceAddToProperDefinitions                           = _ReduceType(2)
	_ReduceDefinitionToProperDefinitions                    = _ReduceType(3)
	_ReduceProperDefinitionsToDefinitions                   = _ReduceType(4)
	_ReduceImproperToDefinitions                            = _ReduceType(5)
	_ReduceNilToDefinitions                                 = _ReduceType(6)
	_ReducePackageDefToDefinition                           = _ReduceType(7)
	_ReduceTypeDefToDefinition                              = _ReduceType(8)
	_ReduceNamedFuncDefToDefinition                         = _ReduceType(9)
	_ReduceGlobalVarDefToDefinition                         = _ReduceType(10)
	_ReduceGlobalVarAssignmentToDefinition                  = _ReduceType(11)
	_ReduceStatementBlockToDefinition                       = _ReduceType(12)
	_ReduceCommentGroupsToDefinition                        = _ReduceType(13)
	_ReduceToStatementBlock                                 = _ReduceType(14)
	_ReduceAddImplicitToProperStatements                    = _ReduceType(15)
	_ReduceAddExplicitToProperStatements                    = _ReduceType(16)
	_ReduceStatementToProperStatements                      = _ReduceType(17)
	_ReduceProperStatementsToStatements                     = _ReduceType(18)
	_ReduceImproperImplicitToStatements                     = _ReduceType(19)
	_ReduceImproperExplicitToStatements                     = _ReduceType(20)
	_ReduceNilToStatements                                  = _ReduceType(21)
	_ReduceUnsafeStatementToSimpleStatement                 = _ReduceType(22)
	_ReduceExprOrImproperStructStatementToSimpleStatement   = _ReduceType(23)
	_ReduceCallbackOpStatementToSimpleStatement             = _ReduceType(24)
	_ReduceJumpStatementToSimpleStatement                   = _ReduceType(25)
	_ReduceAssignStatementToSimpleStatement                 = _ReduceType(26)
	_ReduceUnaryOpAssignStatementToSimpleStatement          = _ReduceType(27)
	_ReduceBinaryOpAssignStatementToSimpleStatement         = _ReduceType(28)
	_ReduceFallthroughStatementToSimpleStatement            = _ReduceType(29)
	_ReduceSimpleStatementToStatement                       = _ReduceType(30)
	_ReduceImportStatementToStatement                       = _ReduceType(31)
	_ReduceCaseBranchStatementToStatement                   = _ReduceType(32)
	_ReduceDefaultBranchStatementToStatement                = _ReduceType(33)
	_ReduceSimpleStatementToOptionalSimpleStatement         = _ReduceType(34)
	_ReduceNilToOptionalSimpleStatement                     = _ReduceType(35)
	_ReduceToExprOrImproperStructStatement                  = _ReduceType(36)
	_ReduceExprToExprs                                      = _ReduceType(37)
	_ReduceAddToExprs                                       = _ReduceType(38)
	_ReduceAsyncToCallbackOp                                = _ReduceType(39)
	_ReduceDeferToCallbackOp                                = _ReduceType(40)
	_ReduceToCallbackOpStatement                            = _ReduceType(41)
	_ReduceToUnsafeStatement                                = _ReduceType(42)
	_ReduceUnlabeledNoValueToJumpStatement                  = _ReduceType(43)
	_ReduceUnlabeledValuedToJumpStatement                   = _ReduceType(44)
	_ReduceLabeledNoValueToJumpStatement                    = _ReduceType(45)
	_ReduceLabeledValuedToJumpStatement                     = _ReduceType(46)
	_ReduceReturnToJumpType                                 = _ReduceType(47)
	_ReduceBreakToJumpType                                  = _ReduceType(48)
	_ReduceContinueToJumpType                               = _ReduceType(49)
	_ReduceToFallthroughStatement                           = _ReduceType(50)
	_ReduceToAssignStatement                                = _ReduceType(51)
	_ReduceToUnaryOpAssignStatement                         = _ReduceType(52)
	_ReduceAddOneAssignToUnaryOpAssign                      = _ReduceType(53)
	_ReduceSubOneAssignToUnaryOpAssign                      = _ReduceType(54)
	_ReduceToBinaryOpAssignStatement                        = _ReduceType(55)
	_ReduceAddAssignToBinaryOpAssign                        = _ReduceType(56)
	_ReduceSubAssignToBinaryOpAssign                        = _ReduceType(57)
	_ReduceMulAssignToBinaryOpAssign                        = _ReduceType(58)
	_ReduceDivAssignToBinaryOpAssign                        = _ReduceType(59)
	_ReduceModAssignToBinaryOpAssign                        = _ReduceType(60)
	_ReduceBitNegAssignToBinaryOpAssign                     = _ReduceType(61)
	_ReduceBitAndAssignToBinaryOpAssign                     = _ReduceType(62)
	_ReduceBitOrAssignToBinaryOpAssign                      = _ReduceType(63)
	_ReduceBitXorAssignToBinaryOpAssign                     = _ReduceType(64)
	_ReduceBitLshiftAssignToBinaryOpAssign                  = _ReduceType(65)
	_ReduceBitRshiftAssignToBinaryOpAssign                  = _ReduceType(66)
	_ReduceSingleToImportStatement                          = _ReduceType(67)
	_ReduceMultipleToImportStatement                        = _ReduceType(68)
	_ReduceStringLiteralToImportClause                      = _ReduceType(69)
	_ReduceAliasToImportClause                              = _ReduceType(70)
	_ReduceAddImplicitToProperImportClauses                 = _ReduceType(71)
	_ReduceAddExplicitToProperImportClauses                 = _ReduceType(72)
	_ReduceImportClauseToProperImportClauses                = _ReduceType(73)
	_ReduceProperImportClausesToImportClauses               = _ReduceType(74)
	_ReduceImplicitToImportClauses                          = _ReduceType(75)
	_ReduceExplicitToImportClauses                          = _ReduceType(76)
	_ReduceCasePatternToCasePatterns                        = _ReduceType(77)
	_ReduceMultipleToCasePatterns                           = _ReduceType(78)
	_ReduceToVarDeclPattern                                 = _ReduceType(79)
	_ReduceVarToVarOrLet                                    = _ReduceType(80)
	_ReduceLetToVarOrLet                                    = _ReduceType(81)
	_ReduceIdentifierToVarPattern                           = _ReduceType(82)
	_ReduceTuplePatternToVarPattern                         = _ReduceType(83)
	_ReduceToTuplePattern                                   = _ReduceType(84)
	_ReduceFieldVarPatternToFieldVarPatterns                = _ReduceType(85)
	_ReduceAddToFieldVarPatterns                            = _ReduceType(86)
	_ReducePositionalToFieldVarPattern                      = _ReduceType(87)
	_ReduceNamedToFieldVarPattern                           = _ReduceType(88)
	_ReduceEllipsisToFieldVarPattern                        = _ReduceType(89)
	_ReduceTypeExprToOptionalTypeExpr                       = _ReduceType(90)
	_ReduceNilToOptionalTypeExpr                            = _ReduceType(91)
	_ReduceToAssignPattern                                  = _ReduceType(92)
	_ReduceAssignPatternToCasePattern                       = _ReduceType(93)
	_ReduceEnumMatchPatternToCasePattern                    = _ReduceType(94)
	_ReduceEnumVarDeclPatternToCasePattern                  = _ReduceType(95)
	_ReduceIfExprToExpr                                     = _ReduceType(96)
	_ReduceSwitchExprToExpr                                 = _ReduceType(97)
	_ReduceLoopExprToExpr                                   = _ReduceType(98)
	_ReduceSequenceExprToExpr                               = _ReduceType(99)
	_ReduceLabelDeclToOptionalLabelDecl                     = _ReduceType(100)
	_ReduceUnlabelledToOptionalLabelDecl                    = _ReduceType(101)
	_ReduceOrExprToSequenceExpr                             = _ReduceType(102)
	_ReduceVarDeclPatternToSequenceExpr                     = _ReduceType(103)
	_ReduceAssignVarPatternToSequenceExpr                   = _ReduceType(104)
	_ReduceNoElseToIfExpr                                   = _ReduceType(105)
	_ReduceIfElseToIfExpr                                   = _ReduceType(106)
	_ReduceMultiIfElseToIfExpr                              = _ReduceType(107)
	_ReduceSequenceExprToCondition                          = _ReduceType(108)
	_ReduceCaseToCondition                                  = _ReduceType(109)
	_ReduceToSwitchExpr                                     = _ReduceType(110)
	_ReduceInfiniteToLoopExpr                               = _ReduceType(111)
	_ReduceDoWhileToLoopExpr                                = _ReduceType(112)
	_ReduceWhileToLoopExpr                                  = _ReduceType(113)
	_ReduceIteratorToLoopExpr                               = _ReduceType(114)
	_ReduceForToLoopExpr                                    = _ReduceType(115)
	_ReduceSequenceExprToOptionalSequenceExpr               = _ReduceType(116)
	_ReduceNilToOptionalSequenceExpr                        = _ReduceType(117)
	_ReduceSequenceExprToForAssignment                      = _ReduceType(118)
	_ReduceAssignToForAssignment                            = _ReduceType(119)
	_ReduceToCallExpr                                       = _ReduceType(120)
	_ReduceBindingToGenericTypeArguments                    = _ReduceType(121)
	_ReduceNilToGenericTypeArguments                        = _ReduceType(122)
	_ReduceAddToProperTypeArguments                         = _ReduceType(123)
	_ReduceTypeExprToProperTypeArguments                    = _ReduceType(124)
	_ReduceProperTypeArgumentsToTypeArguments               = _ReduceType(125)
	_ReduceImproperToTypeArguments                          = _ReduceType(126)
	_ReduceNilToTypeArguments                               = _ReduceType(127)
	_ReduceAddToProperArguments                             = _ReduceType(128)
	_ReduceArgumentToProperArguments                        = _ReduceType(129)
	_ReduceProperArgumentsToArguments                       = _ReduceType(130)
	_ReduceImproperToArguments                              = _ReduceType(131)
	_ReduceNilToArguments                                   = _ReduceType(132)
	_ReducePositionalToArgument                             = _ReduceType(133)
	_ReduceColonExprToArgument                              = _ReduceType(134)
	_ReduceNamedAssignmentToArgument                        = _ReduceType(135)
	_ReduceVarargAssignmentToArgument                       = _ReduceType(136)
	_ReduceSkipPatternToArgument                            = _ReduceType(137)
	_ReduceUnitUnitPairToColonExpr                          = _ReduceType(138)
	_ReduceExprUnitPairToColonExpr                          = _ReduceType(139)
	_ReduceUnitExprPairToColonExpr                          = _ReduceType(140)
	_ReduceExprExprPairToColonExpr                          = _ReduceType(141)
	_ReduceColonExprUnitTupleToColonExpr                    = _ReduceType(142)
	_ReduceColonExprExprTupleToColonExpr                    = _ReduceType(143)
	_ReduceParseErrorExprToAtomExpr                         = _ReduceType(144)
	_ReduceLiteralExprToAtomExpr                            = _ReduceType(145)
	_ReduceIdentifierExprToAtomExpr                         = _ReduceType(146)
	_ReduceBlockExprToAtomExpr                              = _ReduceType(147)
	_ReduceAnonymousFuncExprToAtomExpr                      = _ReduceType(148)
	_ReduceInitializeExprToAtomExpr                         = _ReduceType(149)
	_ReduceImplicitStructExprToAtomExpr                     = _ReduceType(150)
	_ReduceToParseErrorExpr                                 = _ReduceType(151)
	_ReduceTrueToLiteralExpr                                = _ReduceType(152)
	_ReduceFalseToLiteralExpr                               = _ReduceType(153)
	_ReduceIntegerLiteralToLiteralExpr                      = _ReduceType(154)
	_ReduceFloatLiteralToLiteralExpr                        = _ReduceType(155)
	_ReduceRuneLiteralToLiteralExpr                         = _ReduceType(156)
	_ReduceStringLiteralToLiteralExpr                       = _ReduceType(157)
	_ReduceToIdentifierExpr                                 = _ReduceType(158)
	_ReduceToBlockExpr                                      = _ReduceType(159)
	_ReduceToInitializeExpr                                 = _ReduceType(160)
	_ReduceToImplicitStructExpr                             = _ReduceType(161)
	_ReduceAtomExprToAccessibleExpr                         = _ReduceType(162)
	_ReduceAccessExprToAccessibleExpr                       = _ReduceType(163)
	_ReduceCallExprToAccessibleExpr                         = _ReduceType(164)
	_ReduceIndexExprToAccessibleExpr                        = _ReduceType(165)
	_ReduceToAccessExpr                                     = _ReduceType(166)
	_ReduceToIndexExpr                                      = _ReduceType(167)
	_ReduceAccessibleExprToPostfixableExpr                  = _ReduceType(168)
	_ReducePostfixUnaryExprToPostfixableExpr                = _ReduceType(169)
	_ReduceToPostfixUnaryExpr                               = _ReduceType(170)
	_ReducePostfixableExprToPrefixableExpr                  = _ReduceType(171)
	_ReducePrefixUnaryExprToPrefixableExpr                  = _ReduceType(172)
	_ReduceToPrefixUnaryExpr                                = _ReduceType(173)
	_ReduceNotToPrefixUnaryOp                               = _ReduceType(174)
	_ReduceBitNegToPrefixUnaryOp                            = _ReduceType(175)
	_ReduceSubToPrefixUnaryOp                               = _ReduceType(176)
	_ReduceMulToPrefixUnaryOp                               = _ReduceType(177)
	_ReduceBitAndToPrefixUnaryOp                            = _ReduceType(178)
	_ReducePrefixableExprToMulExpr                          = _ReduceType(179)
	_ReduceBinaryMulExprToMulExpr                           = _ReduceType(180)
	_ReduceToBinaryMulExpr                                  = _ReduceType(181)
	_ReduceMulToMulOp                                       = _ReduceType(182)
	_ReduceDivToMulOp                                       = _ReduceType(183)
	_ReduceModToMulOp                                       = _ReduceType(184)
	_ReduceBitAndToMulOp                                    = _ReduceType(185)
	_ReduceBitLshiftToMulOp                                 = _ReduceType(186)
	_ReduceBitRshiftToMulOp                                 = _ReduceType(187)
	_ReduceMulExprToAddExpr                                 = _ReduceType(188)
	_ReduceBinaryAddExprToAddExpr                           = _ReduceType(189)
	_ReduceToBinaryAddExpr                                  = _ReduceType(190)
	_ReduceAddToAddOp                                       = _ReduceType(191)
	_ReduceSubToAddOp                                       = _ReduceType(192)
	_ReduceBitOrToAddOp                                     = _ReduceType(193)
	_ReduceBitXorToAddOp                                    = _ReduceType(194)
	_ReduceAddExprToCmpExpr                                 = _ReduceType(195)
	_ReduceBinaryCmpExprToCmpExpr                           = _ReduceType(196)
	_ReduceToBinaryCmpExpr                                  = _ReduceType(197)
	_ReduceEqualToCmpOp                                     = _ReduceType(198)
	_ReduceNotEqualToCmpOp                                  = _ReduceType(199)
	_ReduceLessToCmpOp                                      = _ReduceType(200)
	_ReduceLessOrEqualToCmpOp                               = _ReduceType(201)
	_ReduceGreaterToCmpOp                                   = _ReduceType(202)
	_ReduceGreaterOrEqualToCmpOp                            = _ReduceType(203)
	_ReduceCmpExprToAndExpr                                 = _ReduceType(204)
	_ReduceBinaryAndExprToAndExpr                           = _ReduceType(205)
	_ReduceToBinaryAndExpr                                  = _ReduceType(206)
	_ReduceAndExprToOrExpr                                  = _ReduceType(207)
	_ReduceBinaryOrExprToOrExpr                             = _ReduceType(208)
	_ReduceToBinaryOrExpr                                   = _ReduceType(209)
	_ReduceExplicitStructTypeExprToInitializableTypeExpr    = _ReduceType(210)
	_ReduceSliceTypeExprToInitializableTypeExpr             = _ReduceType(211)
	_ReduceArrayTypeExprToInitializableTypeExpr             = _ReduceType(212)
	_ReduceMapTypeExprToInitializableTypeExpr               = _ReduceType(213)
	_ReduceToSliceTypeExpr                                  = _ReduceType(214)
	_ReduceToArrayTypeExpr                                  = _ReduceType(215)
	_ReduceToMapTypeExpr                                    = _ReduceType(216)
	_ReduceInitializableTypeExprToAtomTypeExpr              = _ReduceType(217)
	_ReduceNamedTypeExprToAtomTypeExpr                      = _ReduceType(218)
	_ReduceInferredTypeExprToAtomTypeExpr                   = _ReduceType(219)
	_ReduceImplicitStructTypeExprToAtomTypeExpr             = _ReduceType(220)
	_ReduceExplicitEnumTypeExprToAtomTypeExpr               = _ReduceType(221)
	_ReduceImplicitEnumTypeExprToAtomTypeExpr               = _ReduceType(222)
	_ReduceTraitTypeExprToAtomTypeExpr                      = _ReduceType(223)
	_ReduceFuncTypeExprToAtomTypeExpr                       = _ReduceType(224)
	_ReduceParseErrorTypeExprToAtomTypeExpr                 = _ReduceType(225)
	_ReduceLocalToNamedTypeExpr                             = _ReduceType(226)
	_ReduceExternalToNamedTypeExpr                          = _ReduceType(227)
	_ReduceToInferredTypeExpr                               = _ReduceType(228)
	_ReduceToParseErrorTypeExpr                             = _ReduceType(229)
	_ReduceAtomTypeExprToReturnableTypeExpr                 = _ReduceType(230)
	_ReducePrefixedUnaryTypeExprToReturnableTypeExpr        = _ReduceType(231)
	_ReduceToPrefixedUnaryTypeExpr                          = _ReduceType(232)
	_ReduceQuestionToPrefixUnaryTypeOp                      = _ReduceType(233)
	_ReduceExclaimToPrefixUnaryTypeOp                       = _ReduceType(234)
	_ReduceBitAndToPrefixUnaryTypeOp                        = _ReduceType(235)
	_ReduceBitNegToPrefixUnaryTypeOp                        = _ReduceType(236)
	_ReduceTildeTildeToPrefixUnaryTypeOp                    = _ReduceType(237)
	_ReduceReturnableTypeExprToTypeExpr                     = _ReduceType(238)
	_ReduceBinaryTypeExprToTypeExpr                         = _ReduceType(239)
	_ReduceToBinaryTypeExpr                                 = _ReduceType(240)
	_ReduceMulToBinaryTypeOp                                = _ReduceType(241)
	_ReduceAddToBinaryTypeOp                                = _ReduceType(242)
	_ReduceSubToBinaryTypeOp                                = _ReduceType(243)
	_ReduceDefinitionToTypeDef                              = _ReduceType(244)
	_ReduceConstrainedDefToTypeDef                          = _ReduceType(245)
	_ReduceAliasToTypeDef                                   = _ReduceType(246)
	_ReduceUnconstrainedToGenericParameterDef               = _ReduceType(247)
	_ReduceConstrainedToGenericParameterDef                 = _ReduceType(248)
	_ReduceAddToProperGenericParameterDefs                  = _ReduceType(249)
	_ReduceGenericParameterDefToProperGenericParameterDefs  = _ReduceType(250)
	_ReduceProperGenericParameterDefsToGenericParameterDefs = _ReduceType(251)
	_ReduceImproperToGenericParameterDefs                   = _ReduceType(252)
	_ReduceNilToGenericParameterDefs                        = _ReduceType(253)
	_ReduceGenericToOptionalGenericParameters               = _ReduceType(254)
	_ReduceNilToOptionalGenericParameters                   = _ReduceType(255)
	_ReduceExplicitToFieldDef                               = _ReduceType(256)
	_ReduceImplicitToFieldDef                               = _ReduceType(257)
	_ReduceUnsafeStatementToFieldDef                        = _ReduceType(258)
	_ReduceAddToProperImplicitFieldDefs                     = _ReduceType(259)
	_ReduceFieldDefToProperImplicitFieldDefs                = _ReduceType(260)
	_ReduceProperImplicitFieldDefsToImplicitFieldDefs       = _ReduceType(261)
	_ReduceImproperToImplicitFieldDefs                      = _ReduceType(262)
	_ReduceNilToImplicitFieldDefs                           = _ReduceType(263)
	_ReduceToImplicitStructTypeExpr                         = _ReduceType(264)
	_ReduceAddImplicitToProperExplicitFieldDefs             = _ReduceType(265)
	_ReduceAddExplicitToProperExplicitFieldDefs             = _ReduceType(266)
	_ReduceFieldDefToProperExplicitFieldDefs                = _ReduceType(267)
	_ReduceProperExplicitFieldDefsToExplicitFieldDefs       = _ReduceType(268)
	_ReduceImproperImplicitToExplicitFieldDefs              = _ReduceType(269)
	_ReduceImproperExplicitToExplicitFieldDefs              = _ReduceType(270)
	_ReduceNilToExplicitFieldDefs                           = _ReduceType(271)
	_ReduceToExplicitStructTypeExpr                         = _ReduceType(272)
	_ReduceFieldDefToEnumValueDef                           = _ReduceType(273)
	_ReduceDefaultToEnumValueDef                            = _ReduceType(274)
	_ReducePairToImplicitEnumValueDefs                      = _ReduceType(275)
	_ReduceAddToImplicitEnumValueDefs                       = _ReduceType(276)
	_ReduceToImplicitEnumTypeExpr                           = _ReduceType(277)
	_ReduceExplicitPairToExplicitEnumValueDefs              = _ReduceType(278)
	_ReduceImplicitPairToExplicitEnumValueDefs              = _ReduceType(279)
	_ReduceExplicitAddToExplicitEnumValueDefs               = _ReduceType(280)
	_ReduceImplicitAddToExplicitEnumValueDefs               = _ReduceType(281)
	_ReduceToExplicitEnumTypeExpr                           = _ReduceType(282)
	_ReduceFieldDefToTraitProperty                          = _ReduceType(283)
	_ReduceMethodSignatureToTraitProperty                   = _ReduceType(284)
	_ReduceImplicitToProperTraitProperties                  = _ReduceType(285)
	_ReduceExplicitToProperTraitProperties                  = _ReduceType(286)
	_ReduceTraitPropertyToProperTraitProperties             = _ReduceType(287)
	_ReduceProperTraitPropertiesToTraitProperties           = _ReduceType(288)
	_ReduceImproperImplicitToTraitProperties                = _ReduceType(289)
	_ReduceImproperExplicitToTraitProperties                = _ReduceType(290)
	_ReduceNilToTraitProperties                             = _ReduceType(291)
	_ReduceToTraitTypeExpr                                  = _ReduceType(292)
	_ReduceReturnableTypeExprToReturnType                   = _ReduceType(293)
	_ReduceNilToReturnType                                  = _ReduceType(294)
	_ReduceArgToParameterDecl                               = _ReduceType(295)
	_ReduceVarargToParameterDecl                            = _ReduceType(296)
	_ReduceUnamedToParameterDecl                            = _ReduceType(297)
	_ReduceUnnamedVarargToParameterDecl                     = _ReduceType(298)
	_ReduceAddToProperParameterDecls                        = _ReduceType(299)
	_ReduceParameterDeclToProperParameterDecls              = _ReduceType(300)
	_ReduceProperParameterDeclsToParameterDecls             = _ReduceType(301)
	_ReduceImproperToParameterDecls                         = _ReduceType(302)
	_ReduceNilToParameterDecls                              = _ReduceType(303)
	_ReduceToFuncTypeExpr                                   = _ReduceType(304)
	_ReduceToMethodSignature                                = _ReduceType(305)
	_ReduceInferredRefArgToParameterDef                     = _ReduceType(306)
	_ReduceInferredRefVarargToParameterDef                  = _ReduceType(307)
	_ReduceArgToParameterDef                                = _ReduceType(308)
	_ReduceVarargToParameterDef                             = _ReduceType(309)
	_ReduceAddToProperParameterDefs                         = _ReduceType(310)
	_ReduceParameterDefToProperParameterDefs                = _ReduceType(311)
	_ReduceProperParameterDefsToParameterDefs               = _ReduceType(312)
	_ReduceImproperToParameterDefs                          = _ReduceType(313)
	_ReduceNilToParameterDefs                               = _ReduceType(314)
	_ReduceFuncDefToNamedFuncDef                            = _ReduceType(315)
	_ReduceMethodDefToNamedFuncDef                          = _ReduceType(316)
	_ReduceAliasToNamedFuncDef                              = _ReduceType(317)
	_ReduceToAnonymousFuncExpr                              = _ReduceType(318)
	_ReduceNoSpecToPackageDef                               = _ReduceType(319)
	_ReduceWithSpecToPackageDef                             = _ReduceType(320)
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
	case _ReduceToInferredTypeExpr:
		return "ToInferredTypeExpr"
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
	case _ReduceFieldDefToEnumValueDef:
		return "FieldDefToEnumValueDef"
	case _ReduceDefaultToEnumValueDef:
		return "DefaultToEnumValueDef"
	case _ReducePairToImplicitEnumValueDefs:
		return "PairToImplicitEnumValueDefs"
	case _ReduceAddToImplicitEnumValueDefs:
		return "AddToImplicitEnumValueDefs"
	case _ReduceToImplicitEnumTypeExpr:
		return "ToImplicitEnumTypeExpr"
	case _ReduceExplicitPairToExplicitEnumValueDefs:
		return "ExplicitPairToExplicitEnumValueDefs"
	case _ReduceImplicitPairToExplicitEnumValueDefs:
		return "ImplicitPairToExplicitEnumValueDefs"
	case _ReduceExplicitAddToExplicitEnumValueDefs:
		return "ExplicitAddToExplicitEnumValueDefs"
	case _ReduceImplicitAddToExplicitEnumValueDefs:
		return "ImplicitAddToExplicitEnumValueDefs"
	case _ReduceToExplicitEnumTypeExpr:
		return "ToExplicitEnumTypeExpr"
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
	case _ReduceToTraitTypeExpr:
		return "ToTraitTypeExpr"
	case _ReduceReturnableTypeExprToReturnType:
		return "ReturnableTypeExprToReturnType"
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
	case ExprType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, AtomExprType, ParseErrorExprType, LiteralExprType, IdentifierExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
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
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixedUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, TraitTypeExprType, ReturnTypeType, FuncTypeExprType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
	case ExprType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, AtomExprType, ParseErrorExprType, LiteralExprType, IdentifierExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
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
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixedUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, TraitTypeExprType, ReturnTypeType, FuncTypeExprType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
	case _ReduceExprOrImproperStructStatementToSimpleStatement:
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
	case _ReduceTypeExprToOptionalTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalTypeExprType
		//line grammar.lr:299:4
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
		//line grammar.lr:429:4
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
	case _ReduceToInferredTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InferredTypeExprType
		symbol.TypeExpression, err = reducer.ToInferredTypeExpr(args[0].Value)
	case _ReduceToParseErrorTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParseErrorTypeExprType
		symbol.TypeExpression, err = reducer.ToParseErrorTypeExpr(args[0].ParseError)
	case _ReduceAtomTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:633:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixedUnaryTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:634:4
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
		//line grammar.lr:640:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:641:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:642:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:643:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:644:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:649:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceBinaryTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:650:4
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
		//line grammar.lr:656:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:657:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:658:4
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
		//line grammar.lr:679:4
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
		//line grammar.lr:701:4
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
		//line grammar.lr:714:4
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
	case _ReduceToImplicitEnumTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumTypeExprType
		symbol.TypeExpression, err = reducer.ToImplicitEnumTypeExpr(args[0].Value, args[1].Generic_, args[2].Value)
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
	case _ReduceToExplicitEnumTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitEnumTypeExprType
		symbol.TypeExpression, err = reducer.ToExplicitEnumTypeExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
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
		//line grammar.lr:769:4
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
	case _ReduceToTraitTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TraitTypeExprType
		symbol.TypeExpression, err = reducer.ToTraitTypeExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceReturnableTypeExprToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.TypeExpression, err = reducer.ReturnableTypeExprToReturnType(args[0].TypeExpression)
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
		//line grammar.lr:796:4
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
		//line grammar.lr:828:4
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
	_GotoState1Action                                             = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                                             = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                                             = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                                             = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                                             = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                                             = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                                             = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                                             = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                                             = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                                            = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                                            = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                                            = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                                            = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                                            = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                                            = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                                            = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                                            = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                                            = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                                            = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                                            = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                                            = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                                            = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                                            = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                                            = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                                            = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                                            = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                                            = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                                            = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                                            = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                                            = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                                            = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                                            = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                                            = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                                            = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                                            = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                                            = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                                            = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                                            = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                                            = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                                            = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                                            = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                                            = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                                            = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                                            = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                                            = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                                            = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                                            = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                                            = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                                            = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                                            = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                                            = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                                            = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                                            = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                                            = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                                            = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                                            = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                                            = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                                            = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                                            = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                                            = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                                            = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                                            = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                                            = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                                            = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                                            = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                                            = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                                            = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                                            = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                                            = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                                            = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                                            = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                                            = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                                            = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                                            = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                                            = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                                            = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                                            = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                                            = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                                            = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                                            = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                                            = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                                            = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                                            = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                                            = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                                            = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                                            = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                                            = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                                            = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                                            = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                                            = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                                            = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                                            = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                                            = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                                            = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                                            = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                                            = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                                            = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                                            = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                                            = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                                           = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                                           = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                                           = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                                           = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                                           = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                                           = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                                           = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                                           = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                                           = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                                           = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                                           = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                                           = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                                           = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                                           = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                                           = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                                           = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                                           = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                                           = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                                           = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                                           = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                                           = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                                           = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                                           = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                                           = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                                           = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                                           = &_Action{_ShiftAction, _State125, 0}
	_GotoState126Action                                           = &_Action{_ShiftAction, _State126, 0}
	_GotoState127Action                                           = &_Action{_ShiftAction, _State127, 0}
	_GotoState128Action                                           = &_Action{_ShiftAction, _State128, 0}
	_GotoState129Action                                           = &_Action{_ShiftAction, _State129, 0}
	_GotoState130Action                                           = &_Action{_ShiftAction, _State130, 0}
	_GotoState131Action                                           = &_Action{_ShiftAction, _State131, 0}
	_GotoState132Action                                           = &_Action{_ShiftAction, _State132, 0}
	_GotoState133Action                                           = &_Action{_ShiftAction, _State133, 0}
	_GotoState134Action                                           = &_Action{_ShiftAction, _State134, 0}
	_GotoState135Action                                           = &_Action{_ShiftAction, _State135, 0}
	_GotoState136Action                                           = &_Action{_ShiftAction, _State136, 0}
	_GotoState137Action                                           = &_Action{_ShiftAction, _State137, 0}
	_GotoState138Action                                           = &_Action{_ShiftAction, _State138, 0}
	_GotoState139Action                                           = &_Action{_ShiftAction, _State139, 0}
	_GotoState140Action                                           = &_Action{_ShiftAction, _State140, 0}
	_GotoState141Action                                           = &_Action{_ShiftAction, _State141, 0}
	_GotoState142Action                                           = &_Action{_ShiftAction, _State142, 0}
	_GotoState143Action                                           = &_Action{_ShiftAction, _State143, 0}
	_GotoState144Action                                           = &_Action{_ShiftAction, _State144, 0}
	_GotoState145Action                                           = &_Action{_ShiftAction, _State145, 0}
	_GotoState146Action                                           = &_Action{_ShiftAction, _State146, 0}
	_GotoState147Action                                           = &_Action{_ShiftAction, _State147, 0}
	_GotoState148Action                                           = &_Action{_ShiftAction, _State148, 0}
	_GotoState149Action                                           = &_Action{_ShiftAction, _State149, 0}
	_GotoState150Action                                           = &_Action{_ShiftAction, _State150, 0}
	_GotoState151Action                                           = &_Action{_ShiftAction, _State151, 0}
	_GotoState152Action                                           = &_Action{_ShiftAction, _State152, 0}
	_GotoState153Action                                           = &_Action{_ShiftAction, _State153, 0}
	_GotoState154Action                                           = &_Action{_ShiftAction, _State154, 0}
	_GotoState155Action                                           = &_Action{_ShiftAction, _State155, 0}
	_GotoState156Action                                           = &_Action{_ShiftAction, _State156, 0}
	_GotoState157Action                                           = &_Action{_ShiftAction, _State157, 0}
	_GotoState158Action                                           = &_Action{_ShiftAction, _State158, 0}
	_GotoState159Action                                           = &_Action{_ShiftAction, _State159, 0}
	_GotoState160Action                                           = &_Action{_ShiftAction, _State160, 0}
	_GotoState161Action                                           = &_Action{_ShiftAction, _State161, 0}
	_GotoState162Action                                           = &_Action{_ShiftAction, _State162, 0}
	_GotoState163Action                                           = &_Action{_ShiftAction, _State163, 0}
	_GotoState164Action                                           = &_Action{_ShiftAction, _State164, 0}
	_GotoState165Action                                           = &_Action{_ShiftAction, _State165, 0}
	_GotoState166Action                                           = &_Action{_ShiftAction, _State166, 0}
	_GotoState167Action                                           = &_Action{_ShiftAction, _State167, 0}
	_GotoState168Action                                           = &_Action{_ShiftAction, _State168, 0}
	_GotoState169Action                                           = &_Action{_ShiftAction, _State169, 0}
	_GotoState170Action                                           = &_Action{_ShiftAction, _State170, 0}
	_GotoState171Action                                           = &_Action{_ShiftAction, _State171, 0}
	_GotoState172Action                                           = &_Action{_ShiftAction, _State172, 0}
	_GotoState173Action                                           = &_Action{_ShiftAction, _State173, 0}
	_GotoState174Action                                           = &_Action{_ShiftAction, _State174, 0}
	_GotoState175Action                                           = &_Action{_ShiftAction, _State175, 0}
	_GotoState176Action                                           = &_Action{_ShiftAction, _State176, 0}
	_GotoState177Action                                           = &_Action{_ShiftAction, _State177, 0}
	_GotoState178Action                                           = &_Action{_ShiftAction, _State178, 0}
	_GotoState179Action                                           = &_Action{_ShiftAction, _State179, 0}
	_GotoState180Action                                           = &_Action{_ShiftAction, _State180, 0}
	_GotoState181Action                                           = &_Action{_ShiftAction, _State181, 0}
	_GotoState182Action                                           = &_Action{_ShiftAction, _State182, 0}
	_GotoState183Action                                           = &_Action{_ShiftAction, _State183, 0}
	_GotoState184Action                                           = &_Action{_ShiftAction, _State184, 0}
	_GotoState185Action                                           = &_Action{_ShiftAction, _State185, 0}
	_GotoState186Action                                           = &_Action{_ShiftAction, _State186, 0}
	_GotoState187Action                                           = &_Action{_ShiftAction, _State187, 0}
	_GotoState188Action                                           = &_Action{_ShiftAction, _State188, 0}
	_GotoState189Action                                           = &_Action{_ShiftAction, _State189, 0}
	_GotoState190Action                                           = &_Action{_ShiftAction, _State190, 0}
	_GotoState191Action                                           = &_Action{_ShiftAction, _State191, 0}
	_GotoState192Action                                           = &_Action{_ShiftAction, _State192, 0}
	_GotoState193Action                                           = &_Action{_ShiftAction, _State193, 0}
	_GotoState194Action                                           = &_Action{_ShiftAction, _State194, 0}
	_GotoState195Action                                           = &_Action{_ShiftAction, _State195, 0}
	_GotoState196Action                                           = &_Action{_ShiftAction, _State196, 0}
	_GotoState197Action                                           = &_Action{_ShiftAction, _State197, 0}
	_GotoState198Action                                           = &_Action{_ShiftAction, _State198, 0}
	_GotoState199Action                                           = &_Action{_ShiftAction, _State199, 0}
	_GotoState200Action                                           = &_Action{_ShiftAction, _State200, 0}
	_GotoState201Action                                           = &_Action{_ShiftAction, _State201, 0}
	_GotoState202Action                                           = &_Action{_ShiftAction, _State202, 0}
	_GotoState203Action                                           = &_Action{_ShiftAction, _State203, 0}
	_GotoState204Action                                           = &_Action{_ShiftAction, _State204, 0}
	_GotoState205Action                                           = &_Action{_ShiftAction, _State205, 0}
	_GotoState206Action                                           = &_Action{_ShiftAction, _State206, 0}
	_GotoState207Action                                           = &_Action{_ShiftAction, _State207, 0}
	_GotoState208Action                                           = &_Action{_ShiftAction, _State208, 0}
	_GotoState209Action                                           = &_Action{_ShiftAction, _State209, 0}
	_GotoState210Action                                           = &_Action{_ShiftAction, _State210, 0}
	_GotoState211Action                                           = &_Action{_ShiftAction, _State211, 0}
	_GotoState212Action                                           = &_Action{_ShiftAction, _State212, 0}
	_GotoState213Action                                           = &_Action{_ShiftAction, _State213, 0}
	_GotoState214Action                                           = &_Action{_ShiftAction, _State214, 0}
	_GotoState215Action                                           = &_Action{_ShiftAction, _State215, 0}
	_GotoState216Action                                           = &_Action{_ShiftAction, _State216, 0}
	_GotoState217Action                                           = &_Action{_ShiftAction, _State217, 0}
	_GotoState218Action                                           = &_Action{_ShiftAction, _State218, 0}
	_GotoState219Action                                           = &_Action{_ShiftAction, _State219, 0}
	_GotoState220Action                                           = &_Action{_ShiftAction, _State220, 0}
	_GotoState221Action                                           = &_Action{_ShiftAction, _State221, 0}
	_GotoState222Action                                           = &_Action{_ShiftAction, _State222, 0}
	_GotoState223Action                                           = &_Action{_ShiftAction, _State223, 0}
	_GotoState224Action                                           = &_Action{_ShiftAction, _State224, 0}
	_GotoState225Action                                           = &_Action{_ShiftAction, _State225, 0}
	_GotoState226Action                                           = &_Action{_ShiftAction, _State226, 0}
	_GotoState227Action                                           = &_Action{_ShiftAction, _State227, 0}
	_GotoState228Action                                           = &_Action{_ShiftAction, _State228, 0}
	_GotoState229Action                                           = &_Action{_ShiftAction, _State229, 0}
	_GotoState230Action                                           = &_Action{_ShiftAction, _State230, 0}
	_GotoState231Action                                           = &_Action{_ShiftAction, _State231, 0}
	_GotoState232Action                                           = &_Action{_ShiftAction, _State232, 0}
	_GotoState233Action                                           = &_Action{_ShiftAction, _State233, 0}
	_GotoState234Action                                           = &_Action{_ShiftAction, _State234, 0}
	_GotoState235Action                                           = &_Action{_ShiftAction, _State235, 0}
	_GotoState236Action                                           = &_Action{_ShiftAction, _State236, 0}
	_GotoState237Action                                           = &_Action{_ShiftAction, _State237, 0}
	_GotoState238Action                                           = &_Action{_ShiftAction, _State238, 0}
	_GotoState239Action                                           = &_Action{_ShiftAction, _State239, 0}
	_GotoState240Action                                           = &_Action{_ShiftAction, _State240, 0}
	_GotoState241Action                                           = &_Action{_ShiftAction, _State241, 0}
	_GotoState242Action                                           = &_Action{_ShiftAction, _State242, 0}
	_GotoState243Action                                           = &_Action{_ShiftAction, _State243, 0}
	_GotoState244Action                                           = &_Action{_ShiftAction, _State244, 0}
	_GotoState245Action                                           = &_Action{_ShiftAction, _State245, 0}
	_GotoState246Action                                           = &_Action{_ShiftAction, _State246, 0}
	_GotoState247Action                                           = &_Action{_ShiftAction, _State247, 0}
	_GotoState248Action                                           = &_Action{_ShiftAction, _State248, 0}
	_GotoState249Action                                           = &_Action{_ShiftAction, _State249, 0}
	_GotoState250Action                                           = &_Action{_ShiftAction, _State250, 0}
	_GotoState251Action                                           = &_Action{_ShiftAction, _State251, 0}
	_GotoState252Action                                           = &_Action{_ShiftAction, _State252, 0}
	_GotoState253Action                                           = &_Action{_ShiftAction, _State253, 0}
	_GotoState254Action                                           = &_Action{_ShiftAction, _State254, 0}
	_GotoState255Action                                           = &_Action{_ShiftAction, _State255, 0}
	_GotoState256Action                                           = &_Action{_ShiftAction, _State256, 0}
	_GotoState257Action                                           = &_Action{_ShiftAction, _State257, 0}
	_GotoState258Action                                           = &_Action{_ShiftAction, _State258, 0}
	_GotoState259Action                                           = &_Action{_ShiftAction, _State259, 0}
	_GotoState260Action                                           = &_Action{_ShiftAction, _State260, 0}
	_GotoState261Action                                           = &_Action{_ShiftAction, _State261, 0}
	_GotoState262Action                                           = &_Action{_ShiftAction, _State262, 0}
	_GotoState263Action                                           = &_Action{_ShiftAction, _State263, 0}
	_GotoState264Action                                           = &_Action{_ShiftAction, _State264, 0}
	_GotoState265Action                                           = &_Action{_ShiftAction, _State265, 0}
	_GotoState266Action                                           = &_Action{_ShiftAction, _State266, 0}
	_GotoState267Action                                           = &_Action{_ShiftAction, _State267, 0}
	_GotoState268Action                                           = &_Action{_ShiftAction, _State268, 0}
	_GotoState269Action                                           = &_Action{_ShiftAction, _State269, 0}
	_GotoState270Action                                           = &_Action{_ShiftAction, _State270, 0}
	_GotoState271Action                                           = &_Action{_ShiftAction, _State271, 0}
	_GotoState272Action                                           = &_Action{_ShiftAction, _State272, 0}
	_GotoState273Action                                           = &_Action{_ShiftAction, _State273, 0}
	_GotoState274Action                                           = &_Action{_ShiftAction, _State274, 0}
	_GotoState275Action                                           = &_Action{_ShiftAction, _State275, 0}
	_GotoState276Action                                           = &_Action{_ShiftAction, _State276, 0}
	_GotoState277Action                                           = &_Action{_ShiftAction, _State277, 0}
	_GotoState278Action                                           = &_Action{_ShiftAction, _State278, 0}
	_GotoState279Action                                           = &_Action{_ShiftAction, _State279, 0}
	_GotoState280Action                                           = &_Action{_ShiftAction, _State280, 0}
	_GotoState281Action                                           = &_Action{_ShiftAction, _State281, 0}
	_GotoState282Action                                           = &_Action{_ShiftAction, _State282, 0}
	_GotoState283Action                                           = &_Action{_ShiftAction, _State283, 0}
	_GotoState284Action                                           = &_Action{_ShiftAction, _State284, 0}
	_GotoState285Action                                           = &_Action{_ShiftAction, _State285, 0}
	_GotoState286Action                                           = &_Action{_ShiftAction, _State286, 0}
	_GotoState287Action                                           = &_Action{_ShiftAction, _State287, 0}
	_GotoState288Action                                           = &_Action{_ShiftAction, _State288, 0}
	_GotoState289Action                                           = &_Action{_ShiftAction, _State289, 0}
	_GotoState290Action                                           = &_Action{_ShiftAction, _State290, 0}
	_GotoState291Action                                           = &_Action{_ShiftAction, _State291, 0}
	_GotoState292Action                                           = &_Action{_ShiftAction, _State292, 0}
	_GotoState293Action                                           = &_Action{_ShiftAction, _State293, 0}
	_GotoState294Action                                           = &_Action{_ShiftAction, _State294, 0}
	_GotoState295Action                                           = &_Action{_ShiftAction, _State295, 0}
	_GotoState296Action                                           = &_Action{_ShiftAction, _State296, 0}
	_GotoState297Action                                           = &_Action{_ShiftAction, _State297, 0}
	_GotoState298Action                                           = &_Action{_ShiftAction, _State298, 0}
	_GotoState299Action                                           = &_Action{_ShiftAction, _State299, 0}
	_GotoState300Action                                           = &_Action{_ShiftAction, _State300, 0}
	_GotoState301Action                                           = &_Action{_ShiftAction, _State301, 0}
	_GotoState302Action                                           = &_Action{_ShiftAction, _State302, 0}
	_GotoState303Action                                           = &_Action{_ShiftAction, _State303, 0}
	_GotoState304Action                                           = &_Action{_ShiftAction, _State304, 0}
	_GotoState305Action                                           = &_Action{_ShiftAction, _State305, 0}
	_GotoState306Action                                           = &_Action{_ShiftAction, _State306, 0}
	_GotoState307Action                                           = &_Action{_ShiftAction, _State307, 0}
	_GotoState308Action                                           = &_Action{_ShiftAction, _State308, 0}
	_GotoState309Action                                           = &_Action{_ShiftAction, _State309, 0}
	_GotoState310Action                                           = &_Action{_ShiftAction, _State310, 0}
	_GotoState311Action                                           = &_Action{_ShiftAction, _State311, 0}
	_GotoState312Action                                           = &_Action{_ShiftAction, _State312, 0}
	_GotoState313Action                                           = &_Action{_ShiftAction, _State313, 0}
	_GotoState314Action                                           = &_Action{_ShiftAction, _State314, 0}
	_GotoState315Action                                           = &_Action{_ShiftAction, _State315, 0}
	_GotoState316Action                                           = &_Action{_ShiftAction, _State316, 0}
	_GotoState317Action                                           = &_Action{_ShiftAction, _State317, 0}
	_GotoState318Action                                           = &_Action{_ShiftAction, _State318, 0}
	_GotoState319Action                                           = &_Action{_ShiftAction, _State319, 0}
	_GotoState320Action                                           = &_Action{_ShiftAction, _State320, 0}
	_GotoState321Action                                           = &_Action{_ShiftAction, _State321, 0}
	_GotoState322Action                                           = &_Action{_ShiftAction, _State322, 0}
	_GotoState323Action                                           = &_Action{_ShiftAction, _State323, 0}
	_GotoState324Action                                           = &_Action{_ShiftAction, _State324, 0}
	_GotoState325Action                                           = &_Action{_ShiftAction, _State325, 0}
	_GotoState326Action                                           = &_Action{_ShiftAction, _State326, 0}
	_GotoState327Action                                           = &_Action{_ShiftAction, _State327, 0}
	_GotoState328Action                                           = &_Action{_ShiftAction, _State328, 0}
	_GotoState329Action                                           = &_Action{_ShiftAction, _State329, 0}
	_GotoState330Action                                           = &_Action{_ShiftAction, _State330, 0}
	_GotoState331Action                                           = &_Action{_ShiftAction, _State331, 0}
	_GotoState332Action                                           = &_Action{_ShiftAction, _State332, 0}
	_GotoState333Action                                           = &_Action{_ShiftAction, _State333, 0}
	_GotoState334Action                                           = &_Action{_ShiftAction, _State334, 0}
	_GotoState335Action                                           = &_Action{_ShiftAction, _State335, 0}
	_GotoState336Action                                           = &_Action{_ShiftAction, _State336, 0}
	_GotoState337Action                                           = &_Action{_ShiftAction, _State337, 0}
	_GotoState338Action                                           = &_Action{_ShiftAction, _State338, 0}
	_GotoState339Action                                           = &_Action{_ShiftAction, _State339, 0}
	_GotoState340Action                                           = &_Action{_ShiftAction, _State340, 0}
	_GotoState341Action                                           = &_Action{_ShiftAction, _State341, 0}
	_GotoState342Action                                           = &_Action{_ShiftAction, _State342, 0}
	_GotoState343Action                                           = &_Action{_ShiftAction, _State343, 0}
	_GotoState344Action                                           = &_Action{_ShiftAction, _State344, 0}
	_GotoState345Action                                           = &_Action{_ShiftAction, _State345, 0}
	_GotoState346Action                                           = &_Action{_ShiftAction, _State346, 0}
	_GotoState347Action                                           = &_Action{_ShiftAction, _State347, 0}
	_GotoState348Action                                           = &_Action{_ShiftAction, _State348, 0}
	_GotoState349Action                                           = &_Action{_ShiftAction, _State349, 0}
	_GotoState350Action                                           = &_Action{_ShiftAction, _State350, 0}
	_GotoState351Action                                           = &_Action{_ShiftAction, _State351, 0}
	_GotoState352Action                                           = &_Action{_ShiftAction, _State352, 0}
	_GotoState353Action                                           = &_Action{_ShiftAction, _State353, 0}
	_GotoState354Action                                           = &_Action{_ShiftAction, _State354, 0}
	_GotoState355Action                                           = &_Action{_ShiftAction, _State355, 0}
	_GotoState356Action                                           = &_Action{_ShiftAction, _State356, 0}
	_GotoState357Action                                           = &_Action{_ShiftAction, _State357, 0}
	_GotoState358Action                                           = &_Action{_ShiftAction, _State358, 0}
	_GotoState359Action                                           = &_Action{_ShiftAction, _State359, 0}
	_GotoState360Action                                           = &_Action{_ShiftAction, _State360, 0}
	_GotoState361Action                                           = &_Action{_ShiftAction, _State361, 0}
	_GotoState362Action                                           = &_Action{_ShiftAction, _State362, 0}
	_GotoState363Action                                           = &_Action{_ShiftAction, _State363, 0}
	_GotoState364Action                                           = &_Action{_ShiftAction, _State364, 0}
	_GotoState365Action                                           = &_Action{_ShiftAction, _State365, 0}
	_GotoState366Action                                           = &_Action{_ShiftAction, _State366, 0}
	_GotoState367Action                                           = &_Action{_ShiftAction, _State367, 0}
	_GotoState368Action                                           = &_Action{_ShiftAction, _State368, 0}
	_GotoState369Action                                           = &_Action{_ShiftAction, _State369, 0}
	_GotoState370Action                                           = &_Action{_ShiftAction, _State370, 0}
	_GotoState371Action                                           = &_Action{_ShiftAction, _State371, 0}
	_GotoState372Action                                           = &_Action{_ShiftAction, _State372, 0}
	_GotoState373Action                                           = &_Action{_ShiftAction, _State373, 0}
	_GotoState374Action                                           = &_Action{_ShiftAction, _State374, 0}
	_GotoState375Action                                           = &_Action{_ShiftAction, _State375, 0}
	_GotoState376Action                                           = &_Action{_ShiftAction, _State376, 0}
	_GotoState377Action                                           = &_Action{_ShiftAction, _State377, 0}
	_GotoState378Action                                           = &_Action{_ShiftAction, _State378, 0}
	_GotoState379Action                                           = &_Action{_ShiftAction, _State379, 0}
	_GotoState380Action                                           = &_Action{_ShiftAction, _State380, 0}
	_GotoState381Action                                           = &_Action{_ShiftAction, _State381, 0}
	_GotoState382Action                                           = &_Action{_ShiftAction, _State382, 0}
	_GotoState383Action                                           = &_Action{_ShiftAction, _State383, 0}
	_GotoState384Action                                           = &_Action{_ShiftAction, _State384, 0}
	_GotoState385Action                                           = &_Action{_ShiftAction, _State385, 0}
	_GotoState386Action                                           = &_Action{_ShiftAction, _State386, 0}
	_GotoState387Action                                           = &_Action{_ShiftAction, _State387, 0}
	_GotoState388Action                                           = &_Action{_ShiftAction, _State388, 0}
	_GotoState389Action                                           = &_Action{_ShiftAction, _State389, 0}
	_GotoState390Action                                           = &_Action{_ShiftAction, _State390, 0}
	_GotoState391Action                                           = &_Action{_ShiftAction, _State391, 0}
	_GotoState392Action                                           = &_Action{_ShiftAction, _State392, 0}
	_GotoState393Action                                           = &_Action{_ShiftAction, _State393, 0}
	_GotoState394Action                                           = &_Action{_ShiftAction, _State394, 0}
	_GotoState395Action                                           = &_Action{_ShiftAction, _State395, 0}
	_GotoState396Action                                           = &_Action{_ShiftAction, _State396, 0}
	_GotoState397Action                                           = &_Action{_ShiftAction, _State397, 0}
	_GotoState398Action                                           = &_Action{_ShiftAction, _State398, 0}
	_GotoState399Action                                           = &_Action{_ShiftAction, _State399, 0}
	_GotoState400Action                                           = &_Action{_ShiftAction, _State400, 0}
	_GotoState401Action                                           = &_Action{_ShiftAction, _State401, 0}
	_GotoState402Action                                           = &_Action{_ShiftAction, _State402, 0}
	_GotoState403Action                                           = &_Action{_ShiftAction, _State403, 0}
	_GotoState404Action                                           = &_Action{_ShiftAction, _State404, 0}
	_GotoState405Action                                           = &_Action{_ShiftAction, _State405, 0}
	_GotoState406Action                                           = &_Action{_ShiftAction, _State406, 0}
	_GotoState407Action                                           = &_Action{_ShiftAction, _State407, 0}
	_GotoState408Action                                           = &_Action{_ShiftAction, _State408, 0}
	_GotoState409Action                                           = &_Action{_ShiftAction, _State409, 0}
	_GotoState410Action                                           = &_Action{_ShiftAction, _State410, 0}
	_GotoState411Action                                           = &_Action{_ShiftAction, _State411, 0}
	_GotoState412Action                                           = &_Action{_ShiftAction, _State412, 0}
	_GotoState413Action                                           = &_Action{_ShiftAction, _State413, 0}
	_GotoState414Action                                           = &_Action{_ShiftAction, _State414, 0}
	_GotoState415Action                                           = &_Action{_ShiftAction, _State415, 0}
	_GotoState416Action                                           = &_Action{_ShiftAction, _State416, 0}
	_GotoState417Action                                           = &_Action{_ShiftAction, _State417, 0}
	_GotoState418Action                                           = &_Action{_ShiftAction, _State418, 0}
	_GotoState419Action                                           = &_Action{_ShiftAction, _State419, 0}
	_GotoState420Action                                           = &_Action{_ShiftAction, _State420, 0}
	_GotoState421Action                                           = &_Action{_ShiftAction, _State421, 0}
	_GotoState422Action                                           = &_Action{_ShiftAction, _State422, 0}
	_GotoState423Action                                           = &_Action{_ShiftAction, _State423, 0}
	_GotoState424Action                                           = &_Action{_ShiftAction, _State424, 0}
	_GotoState425Action                                           = &_Action{_ShiftAction, _State425, 0}
	_GotoState426Action                                           = &_Action{_ShiftAction, _State426, 0}
	_GotoState427Action                                           = &_Action{_ShiftAction, _State427, 0}
	_GotoState428Action                                           = &_Action{_ShiftAction, _State428, 0}
	_GotoState429Action                                           = &_Action{_ShiftAction, _State429, 0}
	_GotoState430Action                                           = &_Action{_ShiftAction, _State430, 0}
	_GotoState431Action                                           = &_Action{_ShiftAction, _State431, 0}
	_GotoState432Action                                           = &_Action{_ShiftAction, _State432, 0}
	_GotoState433Action                                           = &_Action{_ShiftAction, _State433, 0}
	_GotoState434Action                                           = &_Action{_ShiftAction, _State434, 0}
	_GotoState435Action                                           = &_Action{_ShiftAction, _State435, 0}
	_GotoState436Action                                           = &_Action{_ShiftAction, _State436, 0}
	_GotoState437Action                                           = &_Action{_ShiftAction, _State437, 0}
	_GotoState438Action                                           = &_Action{_ShiftAction, _State438, 0}
	_GotoState439Action                                           = &_Action{_ShiftAction, _State439, 0}
	_GotoState440Action                                           = &_Action{_ShiftAction, _State440, 0}
	_GotoState441Action                                           = &_Action{_ShiftAction, _State441, 0}
	_GotoState442Action                                           = &_Action{_ShiftAction, _State442, 0}
	_GotoState443Action                                           = &_Action{_ShiftAction, _State443, 0}
	_GotoState444Action                                           = &_Action{_ShiftAction, _State444, 0}
	_GotoState445Action                                           = &_Action{_ShiftAction, _State445, 0}
	_GotoState446Action                                           = &_Action{_ShiftAction, _State446, 0}
	_GotoState447Action                                           = &_Action{_ShiftAction, _State447, 0}
	_GotoState448Action                                           = &_Action{_ShiftAction, _State448, 0}
	_GotoState449Action                                           = &_Action{_ShiftAction, _State449, 0}
	_GotoState450Action                                           = &_Action{_ShiftAction, _State450, 0}
	_GotoState451Action                                           = &_Action{_ShiftAction, _State451, 0}
	_GotoState452Action                                           = &_Action{_ShiftAction, _State452, 0}
	_GotoState453Action                                           = &_Action{_ShiftAction, _State453, 0}
	_GotoState454Action                                           = &_Action{_ShiftAction, _State454, 0}
	_GotoState455Action                                           = &_Action{_ShiftAction, _State455, 0}
	_GotoState456Action                                           = &_Action{_ShiftAction, _State456, 0}
	_GotoState457Action                                           = &_Action{_ShiftAction, _State457, 0}
	_GotoState458Action                                           = &_Action{_ShiftAction, _State458, 0}
	_GotoState459Action                                           = &_Action{_ShiftAction, _State459, 0}
	_GotoState460Action                                           = &_Action{_ShiftAction, _State460, 0}
	_GotoState461Action                                           = &_Action{_ShiftAction, _State461, 0}
	_GotoState462Action                                           = &_Action{_ShiftAction, _State462, 0}
	_GotoState463Action                                           = &_Action{_ShiftAction, _State463, 0}
	_GotoState464Action                                           = &_Action{_ShiftAction, _State464, 0}
	_GotoState465Action                                           = &_Action{_ShiftAction, _State465, 0}
	_GotoState466Action                                           = &_Action{_ShiftAction, _State466, 0}
	_GotoState467Action                                           = &_Action{_ShiftAction, _State467, 0}
	_GotoState468Action                                           = &_Action{_ShiftAction, _State468, 0}
	_GotoState469Action                                           = &_Action{_ShiftAction, _State469, 0}
	_GotoState470Action                                           = &_Action{_ShiftAction, _State470, 0}
	_GotoState471Action                                           = &_Action{_ShiftAction, _State471, 0}
	_ReduceToSourceAction                                         = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceAddToProperDefinitionsAction                           = &_Action{_ReduceAction, 0, _ReduceAddToProperDefinitions}
	_ReduceDefinitionToProperDefinitionsAction                    = &_Action{_ReduceAction, 0, _ReduceDefinitionToProperDefinitions}
	_ReduceProperDefinitionsToDefinitionsAction                   = &_Action{_ReduceAction, 0, _ReduceProperDefinitionsToDefinitions}
	_ReduceImproperToDefinitionsAction                            = &_Action{_ReduceAction, 0, _ReduceImproperToDefinitions}
	_ReduceNilToDefinitionsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToDefinitions}
	_ReducePackageDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                              = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                         = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                         = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                  = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                       = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
	_ReduceCommentGroupsToDefinitionAction                        = &_Action{_ReduceAction, 0, _ReduceCommentGroupsToDefinition}
	_ReduceToStatementBlockAction                                 = &_Action{_ReduceAction, 0, _ReduceToStatementBlock}
	_ReduceAddImplicitToProperStatementsAction                    = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperStatements}
	_ReduceAddExplicitToProperStatementsAction                    = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperStatements}
	_ReduceStatementToProperStatementsAction                      = &_Action{_ReduceAction, 0, _ReduceStatementToProperStatements}
	_ReduceProperStatementsToStatementsAction                     = &_Action{_ReduceAction, 0, _ReduceProperStatementsToStatements}
	_ReduceImproperImplicitToStatementsAction                     = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToStatements}
	_ReduceImproperExplicitToStatementsAction                     = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToStatements}
	_ReduceNilToStatementsAction                                  = &_Action{_ReduceAction, 0, _ReduceNilToStatements}
	_ReduceUnsafeStatementToSimpleStatementAction                 = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}
	_ReduceExprOrImproperStructStatementToSimpleStatementAction   = &_Action{_ReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}
	_ReduceCallbackOpStatementToSimpleStatementAction             = &_Action{_ReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}
	_ReduceJumpStatementToSimpleStatementAction                   = &_Action{_ReduceAction, 0, _ReduceJumpStatementToSimpleStatement}
	_ReduceAssignStatementToSimpleStatementAction                 = &_Action{_ReduceAction, 0, _ReduceAssignStatementToSimpleStatement}
	_ReduceUnaryOpAssignStatementToSimpleStatementAction          = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}
	_ReduceBinaryOpAssignStatementToSimpleStatementAction         = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}
	_ReduceFallthroughStatementToSimpleStatementAction            = &_Action{_ReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}
	_ReduceSimpleStatementToStatementAction                       = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToStatement}
	_ReduceImportStatementToStatementAction                       = &_Action{_ReduceAction, 0, _ReduceImportStatementToStatement}
	_ReduceCaseBranchStatementToStatementAction                   = &_Action{_ReduceAction, 0, _ReduceCaseBranchStatementToStatement}
	_ReduceDefaultBranchStatementToStatementAction                = &_Action{_ReduceAction, 0, _ReduceDefaultBranchStatementToStatement}
	_ReduceSimpleStatementToOptionalSimpleStatementAction         = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToOptionalSimpleStatement}
	_ReduceNilToOptionalSimpleStatementAction                     = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}
	_ReduceToExprOrImproperStructStatementAction                  = &_Action{_ReduceAction, 0, _ReduceToExprOrImproperStructStatement}
	_ReduceExprToExprsAction                                      = &_Action{_ReduceAction, 0, _ReduceExprToExprs}
	_ReduceAddToExprsAction                                       = &_Action{_ReduceAction, 0, _ReduceAddToExprs}
	_ReduceAsyncToCallbackOpAction                                = &_Action{_ReduceAction, 0, _ReduceAsyncToCallbackOp}
	_ReduceDeferToCallbackOpAction                                = &_Action{_ReduceAction, 0, _ReduceDeferToCallbackOp}
	_ReduceToCallbackOpStatementAction                            = &_Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}
	_ReduceToUnsafeStatementAction                                = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceUnlabeledNoValueToJumpStatementAction                  = &_Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}
	_ReduceUnlabeledValuedToJumpStatementAction                   = &_Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}
	_ReduceLabeledNoValueToJumpStatementAction                    = &_Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}
	_ReduceLabeledValuedToJumpStatementAction                     = &_Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}
	_ReduceReturnToJumpTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                  = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                               = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceToFallthroughStatementAction                           = &_Action{_ReduceAction, 0, _ReduceToFallthroughStatement}
	_ReduceToAssignStatementAction                                = &_Action{_ReduceAction, 0, _ReduceToAssignStatement}
	_ReduceToUnaryOpAssignStatementAction                         = &_Action{_ReduceAction, 0, _ReduceToUnaryOpAssignStatement}
	_ReduceAddOneAssignToUnaryOpAssignAction                      = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}
	_ReduceSubOneAssignToUnaryOpAssignAction                      = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}
	_ReduceToBinaryOpAssignStatementAction                        = &_Action{_ReduceAction, 0, _ReduceToBinaryOpAssignStatement}
	_ReduceAddAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction                     = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction                     = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction                      = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction                     = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction                  = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction                  = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceSingleToImportStatementAction                          = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                        = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                      = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                              = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceAddImplicitToProperImportClausesAction                 = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperImportClauses}
	_ReduceAddExplicitToProperImportClausesAction                 = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperImportClauses}
	_ReduceImportClauseToProperImportClausesAction                = &_Action{_ReduceAction, 0, _ReduceImportClauseToProperImportClauses}
	_ReduceProperImportClausesToImportClausesAction               = &_Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}
	_ReduceImplicitToImportClausesAction                          = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}
	_ReduceExplicitToImportClausesAction                          = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}
	_ReduceCasePatternToCasePatternsAction                        = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                           = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceToVarDeclPatternAction                                 = &_Action{_ReduceAction, 0, _ReduceToVarDeclPattern}
	_ReduceVarToVarOrLetAction                                    = &_Action{_ReduceAction, 0, _ReduceVarToVarOrLet}
	_ReduceLetToVarOrLetAction                                    = &_Action{_ReduceAction, 0, _ReduceLetToVarOrLet}
	_ReduceIdentifierToVarPatternAction                           = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceTuplePatternToVarPatternAction                         = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                   = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                            = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                      = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                           = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceEllipsisToFieldVarPatternAction                        = &_Action{_ReduceAction, 0, _ReduceEllipsisToFieldVarPattern}
	_ReduceTypeExprToOptionalTypeExprAction                       = &_Action{_ReduceAction, 0, _ReduceTypeExprToOptionalTypeExpr}
	_ReduceNilToOptionalTypeExprAction                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalTypeExpr}
	_ReduceToAssignPatternAction                                  = &_Action{_ReduceAction, 0, _ReduceToAssignPattern}
	_ReduceAssignPatternToCasePatternAction                       = &_Action{_ReduceAction, 0, _ReduceAssignPatternToCasePattern}
	_ReduceEnumMatchPatternToCasePatternAction                    = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}
	_ReduceEnumVarDeclPatternToCasePatternAction                  = &_Action{_ReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}
	_ReduceIfExprToExprAction                                     = &_Action{_ReduceAction, 0, _ReduceIfExprToExpr}
	_ReduceSwitchExprToExprAction                                 = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpr}
	_ReduceLoopExprToExprAction                                   = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpr}
	_ReduceSequenceExprToExprAction                               = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpr}
	_ReduceLabelDeclToOptionalLabelDeclAction                     = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                    = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceOrExprToSequenceExprAction                             = &_Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}
	_ReduceVarDeclPatternToSequenceExprAction                     = &_Action{_ReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}
	_ReduceAssignVarPatternToSequenceExprAction                   = &_Action{_ReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}
	_ReduceNoElseToIfExprAction                                   = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                   = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                              = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceSequenceExprToConditionAction                          = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                  = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                               = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                  = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                               = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                    = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction               = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceSequenceExprToForAssignmentAction                      = &_Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}
	_ReduceAssignToForAssignmentAction                            = &_Action{_ReduceAction, 0, _ReduceAssignToForAssignment}
	_ReduceToCallExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToCallExpr}
	_ReduceBindingToGenericTypeArgumentsAction                    = &_Action{_ReduceAction, 0, _ReduceBindingToGenericTypeArguments}
	_ReduceNilToGenericTypeArgumentsAction                        = &_Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}
	_ReduceAddToProperTypeArgumentsAction                         = &_Action{_ReduceAction, 0, _ReduceAddToProperTypeArguments}
	_ReduceTypeExprToProperTypeArgumentsAction                    = &_Action{_ReduceAction, 0, _ReduceTypeExprToProperTypeArguments}
	_ReduceProperTypeArgumentsToTypeArgumentsAction               = &_Action{_ReduceAction, 0, _ReduceProperTypeArgumentsToTypeArguments}
	_ReduceImproperToTypeArgumentsAction                          = &_Action{_ReduceAction, 0, _ReduceImproperToTypeArguments}
	_ReduceNilToTypeArgumentsAction                               = &_Action{_ReduceAction, 0, _ReduceNilToTypeArguments}
	_ReduceAddToProperArgumentsAction                             = &_Action{_ReduceAction, 0, _ReduceAddToProperArguments}
	_ReduceArgumentToProperArgumentsAction                        = &_Action{_ReduceAction, 0, _ReduceArgumentToProperArguments}
	_ReduceProperArgumentsToArgumentsAction                       = &_Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}
	_ReduceImproperToArgumentsAction                              = &_Action{_ReduceAction, 0, _ReduceImproperToArguments}
	_ReduceNilToArgumentsAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToArguments}
	_ReducePositionalToArgumentAction                             = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceColonExprToArgumentAction                              = &_Action{_ReduceAction, 0, _ReduceColonExprToArgument}
	_ReduceNamedAssignmentToArgumentAction                        = &_Action{_ReduceAction, 0, _ReduceNamedAssignmentToArgument}
	_ReduceVarargAssignmentToArgumentAction                       = &_Action{_ReduceAction, 0, _ReduceVarargAssignmentToArgument}
	_ReduceSkipPatternToArgumentAction                            = &_Action{_ReduceAction, 0, _ReduceSkipPatternToArgument}
	_ReduceUnitUnitPairToColonExprAction                          = &_Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}
	_ReduceExprUnitPairToColonExprAction                          = &_Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}
	_ReduceUnitExprPairToColonExprAction                          = &_Action{_ReduceAction, 0, _ReduceUnitExprPairToColonExpr}
	_ReduceExprExprPairToColonExprAction                          = &_Action{_ReduceAction, 0, _ReduceExprExprPairToColonExpr}
	_ReduceColonExprUnitTupleToColonExprAction                    = &_Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}
	_ReduceColonExprExprTupleToColonExprAction                    = &_Action{_ReduceAction, 0, _ReduceColonExprExprTupleToColonExpr}
	_ReduceParseErrorExprToAtomExprAction                         = &_Action{_ReduceAction, 0, _ReduceParseErrorExprToAtomExpr}
	_ReduceLiteralExprToAtomExprAction                            = &_Action{_ReduceAction, 0, _ReduceLiteralExprToAtomExpr}
	_ReduceIdentifierExprToAtomExprAction                         = &_Action{_ReduceAction, 0, _ReduceIdentifierExprToAtomExpr}
	_ReduceBlockExprToAtomExprAction                              = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                      = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                         = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                     = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceToParseErrorExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToParseErrorExpr}
	_ReduceTrueToLiteralExprAction                                = &_Action{_ReduceAction, 0, _ReduceTrueToLiteralExpr}
	_ReduceFalseToLiteralExprAction                               = &_Action{_ReduceAction, 0, _ReduceFalseToLiteralExpr}
	_ReduceIntegerLiteralToLiteralExprAction                      = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}
	_ReduceFloatLiteralToLiteralExprAction                        = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}
	_ReduceRuneLiteralToLiteralExprAction                         = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}
	_ReduceStringLiteralToLiteralExprAction                       = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteralExpr}
	_ReduceToIdentifierExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToIdentifierExpr}
	_ReduceToBlockExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToInitializeExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToInitializeExpr}
	_ReduceToImplicitStructExprAction                             = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessibleExprAction                         = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessibleExpr}
	_ReduceAccessExprToAccessibleExprAction                       = &_Action{_ReduceAction, 0, _ReduceAccessExprToAccessibleExpr}
	_ReduceCallExprToAccessibleExprAction                         = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessibleExpr}
	_ReduceIndexExprToAccessibleExprAction                        = &_Action{_ReduceAction, 0, _ReduceIndexExprToAccessibleExpr}
	_ReduceToAccessExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToAccessExpr}
	_ReduceToIndexExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToIndexExpr}
	_ReduceAccessibleExprToPostfixableExprAction                  = &_Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}
	_ReducePostfixUnaryExprToPostfixableExprAction                = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}
	_ReduceToPostfixUnaryExprAction                               = &_Action{_ReduceAction, 0, _ReduceToPostfixUnaryExpr}
	_ReducePostfixableExprToPrefixableExprAction                  = &_Action{_ReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}
	_ReducePrefixUnaryExprToPrefixableExprAction                  = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}
	_ReduceToPrefixUnaryExprAction                                = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                               = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                            = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                               = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                               = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                            = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePrefixableExprToMulExprAction                          = &_Action{_ReduceAction, 0, _ReducePrefixableExprToMulExpr}
	_ReduceBinaryMulExprToMulExprAction                           = &_Action{_ReduceAction, 0, _ReduceBinaryMulExprToMulExpr}
	_ReduceToBinaryMulExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToBinaryMulExpr}
	_ReduceMulToMulOpAction                                       = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                       = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                       = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                    = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                 = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                 = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceMulExprToAddExprAction                                 = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceBinaryAddExprToAddExprAction                           = &_Action{_ReduceAction, 0, _ReduceBinaryAddExprToAddExpr}
	_ReduceToBinaryAddExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToBinaryAddExpr}
	_ReduceAddToAddOpAction                                       = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                       = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                     = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                    = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceAddExprToCmpExprAction                                 = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceBinaryCmpExprToCmpExprAction                           = &_Action{_ReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}
	_ReduceToBinaryCmpExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToBinaryCmpExpr}
	_ReduceEqualToCmpOpAction                                     = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                  = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                      = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                               = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                   = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                            = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceCmpExprToAndExprAction                                 = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceBinaryAndExprToAndExprAction                           = &_Action{_ReduceAction, 0, _ReduceBinaryAndExprToAndExpr}
	_ReduceToBinaryAndExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToBinaryAndExpr}
	_ReduceAndExprToOrExprAction                                  = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceBinaryOrExprToOrExprAction                             = &_Action{_ReduceAction, 0, _ReduceBinaryOrExprToOrExpr}
	_ReduceToBinaryOrExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}
	_ReduceExplicitStructTypeExprToInitializableTypeExprAction    = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}
	_ReduceSliceTypeExprToInitializableTypeExprAction             = &_Action{_ReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}
	_ReduceArrayTypeExprToInitializableTypeExprAction             = &_Action{_ReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}
	_ReduceMapTypeExprToInitializableTypeExprAction               = &_Action{_ReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}
	_ReduceToSliceTypeExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToSliceTypeExpr}
	_ReduceToArrayTypeExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToArrayTypeExpr}
	_ReduceToMapTypeExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToMapTypeExpr}
	_ReduceInitializableTypeExprToAtomTypeExprAction              = &_Action{_ReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}
	_ReduceNamedTypeExprToAtomTypeExprAction                      = &_Action{_ReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}
	_ReduceInferredTypeExprToAtomTypeExprAction                   = &_Action{_ReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}
	_ReduceImplicitStructTypeExprToAtomTypeExprAction             = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}
	_ReduceExplicitEnumTypeExprToAtomTypeExprAction               = &_Action{_ReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}
	_ReduceImplicitEnumTypeExprToAtomTypeExprAction               = &_Action{_ReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}
	_ReduceTraitTypeExprToAtomTypeExprAction                      = &_Action{_ReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}
	_ReduceFuncTypeExprToAtomTypeExprAction                       = &_Action{_ReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}
	_ReduceParseErrorTypeExprToAtomTypeExprAction                 = &_Action{_ReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}
	_ReduceLocalToNamedTypeExprAction                             = &_Action{_ReduceAction, 0, _ReduceLocalToNamedTypeExpr}
	_ReduceExternalToNamedTypeExprAction                          = &_Action{_ReduceAction, 0, _ReduceExternalToNamedTypeExpr}
	_ReduceToInferredTypeExprAction                               = &_Action{_ReduceAction, 0, _ReduceToInferredTypeExpr}
	_ReduceToParseErrorTypeExprAction                             = &_Action{_ReduceAction, 0, _ReduceToParseErrorTypeExpr}
	_ReduceAtomTypeExprToReturnableTypeExprAction                 = &_Action{_ReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}
	_ReducePrefixedUnaryTypeExprToReturnableTypeExprAction        = &_Action{_ReduceAction, 0, _ReducePrefixedUnaryTypeExprToReturnableTypeExpr}
	_ReduceToPrefixedUnaryTypeExprAction                          = &_Action{_ReduceAction, 0, _ReduceToPrefixedUnaryTypeExpr}
	_ReduceQuestionToPrefixUnaryTypeOpAction                      = &_Action{_ReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}
	_ReduceExclaimToPrefixUnaryTypeOpAction                       = &_Action{_ReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}
	_ReduceBitAndToPrefixUnaryTypeOpAction                        = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}
	_ReduceBitNegToPrefixUnaryTypeOpAction                        = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}
	_ReduceTildeTildeToPrefixUnaryTypeOpAction                    = &_Action{_ReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}
	_ReduceReturnableTypeExprToTypeExprAction                     = &_Action{_ReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}
	_ReduceBinaryTypeExprToTypeExprAction                         = &_Action{_ReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}
	_ReduceToBinaryTypeExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToBinaryTypeExpr}
	_ReduceMulToBinaryTypeOpAction                                = &_Action{_ReduceAction, 0, _ReduceMulToBinaryTypeOp}
	_ReduceAddToBinaryTypeOpAction                                = &_Action{_ReduceAction, 0, _ReduceAddToBinaryTypeOp}
	_ReduceSubToBinaryTypeOpAction                                = &_Action{_ReduceAction, 0, _ReduceSubToBinaryTypeOp}
	_ReduceDefinitionToTypeDefAction                              = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                          = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                   = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction               = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                 = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceAddToProperGenericParameterDefsAction                  = &_Action{_ReduceAction, 0, _ReduceAddToProperGenericParameterDefs}
	_ReduceGenericParameterDefToProperGenericParameterDefsAction  = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToProperGenericParameterDefs}
	_ReduceProperGenericParameterDefsToGenericParameterDefsAction = &_Action{_ReduceAction, 0, _ReduceProperGenericParameterDefsToGenericParameterDefs}
	_ReduceImproperToGenericParameterDefsAction                   = &_Action{_ReduceAction, 0, _ReduceImproperToGenericParameterDefs}
	_ReduceNilToGenericParameterDefsAction                        = &_Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction               = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                   = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceExplicitToFieldDefAction                               = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDef}
	_ReduceImplicitToFieldDefAction                               = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDef}
	_ReduceUnsafeStatementToFieldDefAction                        = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToFieldDef}
	_ReduceAddToProperImplicitFieldDefsAction                     = &_Action{_ReduceAction, 0, _ReduceAddToProperImplicitFieldDefs}
	_ReduceFieldDefToProperImplicitFieldDefsAction                = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperImplicitFieldDefs}
	_ReduceProperImplicitFieldDefsToImplicitFieldDefsAction       = &_Action{_ReduceAction, 0, _ReduceProperImplicitFieldDefsToImplicitFieldDefs}
	_ReduceImproperToImplicitFieldDefsAction                      = &_Action{_ReduceAction, 0, _ReduceImproperToImplicitFieldDefs}
	_ReduceNilToImplicitFieldDefsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToImplicitFieldDefs}
	_ReduceToImplicitStructTypeExprAction                         = &_Action{_ReduceAction, 0, _ReduceToImplicitStructTypeExpr}
	_ReduceAddImplicitToProperExplicitFieldDefsAction             = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperExplicitFieldDefs}
	_ReduceAddExplicitToProperExplicitFieldDefsAction             = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperExplicitFieldDefs}
	_ReduceFieldDefToProperExplicitFieldDefsAction                = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperExplicitFieldDefs}
	_ReduceProperExplicitFieldDefsToExplicitFieldDefsAction       = &_Action{_ReduceAction, 0, _ReduceProperExplicitFieldDefsToExplicitFieldDefs}
	_ReduceImproperImplicitToExplicitFieldDefsAction              = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitFieldDefs}
	_ReduceImproperExplicitToExplicitFieldDefsAction              = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitFieldDefs}
	_ReduceNilToExplicitFieldDefsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToExplicitFieldDefs}
	_ReduceToExplicitStructTypeExprAction                         = &_Action{_ReduceAction, 0, _ReduceToExplicitStructTypeExpr}
	_ReduceFieldDefToEnumValueDefAction                           = &_Action{_ReduceAction, 0, _ReduceFieldDefToEnumValueDef}
	_ReduceDefaultToEnumValueDefAction                            = &_Action{_ReduceAction, 0, _ReduceDefaultToEnumValueDef}
	_ReducePairToImplicitEnumValueDefsAction                      = &_Action{_ReduceAction, 0, _ReducePairToImplicitEnumValueDefs}
	_ReduceAddToImplicitEnumValueDefsAction                       = &_Action{_ReduceAction, 0, _ReduceAddToImplicitEnumValueDefs}
	_ReduceToImplicitEnumTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumTypeExpr}
	_ReduceExplicitPairToExplicitEnumValueDefsAction              = &_Action{_ReduceAction, 0, _ReduceExplicitPairToExplicitEnumValueDefs}
	_ReduceImplicitPairToExplicitEnumValueDefsAction              = &_Action{_ReduceAction, 0, _ReduceImplicitPairToExplicitEnumValueDefs}
	_ReduceExplicitAddToExplicitEnumValueDefsAction               = &_Action{_ReduceAction, 0, _ReduceExplicitAddToExplicitEnumValueDefs}
	_ReduceImplicitAddToExplicitEnumValueDefsAction               = &_Action{_ReduceAction, 0, _ReduceImplicitAddToExplicitEnumValueDefs}
	_ReduceToExplicitEnumTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumTypeExpr}
	_ReduceFieldDefToTraitPropertyAction                          = &_Action{_ReduceAction, 0, _ReduceFieldDefToTraitProperty}
	_ReduceMethodSignatureToTraitPropertyAction                   = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToTraitProperty}
	_ReduceImplicitToProperTraitPropertiesAction                  = &_Action{_ReduceAction, 0, _ReduceImplicitToProperTraitProperties}
	_ReduceExplicitToProperTraitPropertiesAction                  = &_Action{_ReduceAction, 0, _ReduceExplicitToProperTraitProperties}
	_ReduceTraitPropertyToProperTraitPropertiesAction             = &_Action{_ReduceAction, 0, _ReduceTraitPropertyToProperTraitProperties}
	_ReduceProperTraitPropertiesToTraitPropertiesAction           = &_Action{_ReduceAction, 0, _ReduceProperTraitPropertiesToTraitProperties}
	_ReduceImproperImplicitToTraitPropertiesAction                = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToTraitProperties}
	_ReduceImproperExplicitToTraitPropertiesAction                = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToTraitProperties}
	_ReduceNilToTraitPropertiesAction                             = &_Action{_ReduceAction, 0, _ReduceNilToTraitProperties}
	_ReduceToTraitTypeExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToTraitTypeExpr}
	_ReduceReturnableTypeExprToReturnTypeAction                   = &_Action{_ReduceAction, 0, _ReduceReturnableTypeExprToReturnType}
	_ReduceNilToReturnTypeAction                                  = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceArgToParameterDeclAction                               = &_Action{_ReduceAction, 0, _ReduceArgToParameterDecl}
	_ReduceVarargToParameterDeclAction                            = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDecl}
	_ReduceUnamedToParameterDeclAction                            = &_Action{_ReduceAction, 0, _ReduceUnamedToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                     = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceAddToProperParameterDeclsAction                        = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDecls}
	_ReduceParameterDeclToProperParameterDeclsAction              = &_Action{_ReduceAction, 0, _ReduceParameterDeclToProperParameterDecls}
	_ReduceProperParameterDeclsToParameterDeclsAction             = &_Action{_ReduceAction, 0, _ReduceProperParameterDeclsToParameterDecls}
	_ReduceImproperToParameterDeclsAction                         = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDecls}
	_ReduceNilToParameterDeclsAction                              = &_Action{_ReduceAction, 0, _ReduceNilToParameterDecls}
	_ReduceToFuncTypeExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToFuncTypeExpr}
	_ReduceToMethodSignatureAction                                = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
	_ReduceInferredRefArgToParameterDefAction                     = &_Action{_ReduceAction, 0, _ReduceInferredRefArgToParameterDef}
	_ReduceInferredRefVarargToParameterDefAction                  = &_Action{_ReduceAction, 0, _ReduceInferredRefVarargToParameterDef}
	_ReduceArgToParameterDefAction                                = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                             = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceAddToProperParameterDefsAction                         = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDefs}
	_ReduceParameterDefToProperParameterDefsAction                = &_Action{_ReduceAction, 0, _ReduceParameterDefToProperParameterDefs}
	_ReduceProperParameterDefsToParameterDefsAction               = &_Action{_ReduceAction, 0, _ReduceProperParameterDefsToParameterDefs}
	_ReduceImproperToParameterDefsAction                          = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDefs}
	_ReduceNilToParameterDefsAction                               = &_Action{_ReduceAction, 0, _ReduceNilToParameterDefs}
	_ReduceFuncDefToNamedFuncDefAction                            = &_Action{_ReduceAction, 0, _ReduceFuncDefToNamedFuncDef}
	_ReduceMethodDefToNamedFuncDefAction                          = &_Action{_ReduceAction, 0, _ReduceMethodDefToNamedFuncDef}
	_ReduceAliasToNamedFuncDefAction                              = &_Action{_ReduceAction, 0, _ReduceAliasToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                              = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                               = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                             = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
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
	{_State2, TrueToken}:                           _GotoState52Action,
	{_State2, FalseToken}:                          _GotoState34Action,
	{_State2, CaseToken}:                           _GotoState29Action,
	{_State2, DefaultToken}:                        _GotoState31Action,
	{_State2, ReturnToken}:                         _GotoState47Action,
	{_State2, BreakToken}:                          _GotoState28Action,
	{_State2, ContinueToken}:                       _GotoState30Action,
	{_State2, FallthroughToken}:                    _GotoState33Action,
	{_State2, ImportToken}:                         _GotoState39Action,
	{_State2, UnsafeToken}:                         _GotoState53Action,
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
	{_State2, SimpleStatementType}:                 _GotoState99Action,
	{_State2, StatementType}:                       _GotoState6Action,
	{_State2, ExprOrImproperStructStatementType}:   _GotoState76Action,
	{_State2, ExprsType}:                           _GotoState77Action,
	{_State2, CallbackOpType}:                      _GotoState71Action,
	{_State2, CallbackOpStatementType}:             _GotoState72Action,
	{_State2, UnsafeStatementType}:                 _GotoState102Action,
	{_State2, JumpStatementType}:                   _GotoState85Action,
	{_State2, JumpTypeType}:                        _GotoState86Action,
	{_State2, FallthroughStatementType}:            _GotoState78Action,
	{_State2, AssignStatementType}:                 _GotoState61Action,
	{_State2, UnaryOpAssignStatementType}:          _GotoState101Action,
	{_State2, BinaryOpAssignStatementType}:         _GotoState67Action,
	{_State2, ImportStatementType}:                 _GotoState81Action,
	{_State2, VarDeclPatternType}:                  _GotoState103Action,
	{_State2, VarOrLetType}:                        _GotoState24Action,
	{_State2, AssignPatternType}:                   _GotoState60Action,
	{_State2, ExprType}:                            _GotoState75Action,
	{_State2, OptionalLabelDeclType}:               _GotoState90Action,
	{_State2, SequenceExprType}:                    _GotoState98Action,
	{_State2, CallExprType}:                        _GotoState70Action,
	{_State2, AtomExprType}:                        _GotoState62Action,
	{_State2, ParseErrorExprType}:                  _GotoState92Action,
	{_State2, LiteralExprType}:                     _GotoState87Action,
	{_State2, IdentifierExprType}:                  _GotoState79Action,
	{_State2, BlockExprType}:                       _GotoState69Action,
	{_State2, InitializeExprType}:                  _GotoState84Action,
	{_State2, ImplicitStructExprType}:              _GotoState80Action,
	{_State2, AccessibleExprType}:                  _GotoState55Action,
	{_State2, AccessExprType}:                      _GotoState54Action,
	{_State2, IndexExprType}:                       _GotoState82Action,
	{_State2, PostfixableExprType}:                 _GotoState94Action,
	{_State2, PostfixUnaryExprType}:                _GotoState93Action,
	{_State2, PrefixableExprType}:                  _GotoState97Action,
	{_State2, PrefixUnaryExprType}:                 _GotoState95Action,
	{_State2, PrefixUnaryOpType}:                   _GotoState96Action,
	{_State2, MulExprType}:                         _GotoState89Action,
	{_State2, BinaryMulExprType}:                   _GotoState66Action,
	{_State2, AddExprType}:                         _GotoState56Action,
	{_State2, BinaryAddExprType}:                   _GotoState63Action,
	{_State2, CmpExprType}:                         _GotoState73Action,
	{_State2, BinaryCmpExprType}:                   _GotoState65Action,
	{_State2, AndExprType}:                         _GotoState57Action,
	{_State2, BinaryAndExprType}:                   _GotoState64Action,
	{_State2, OrExprType}:                          _GotoState91Action,
	{_State2, BinaryOrExprType}:                    _GotoState68Action,
	{_State2, InitializableTypeExprType}:           _GotoState83Action,
	{_State2, SliceTypeExprType}:                   _GotoState100Action,
	{_State2, ArrayTypeExprType}:                   _GotoState59Action,
	{_State2, MapTypeExprType}:                     _GotoState88Action,
	{_State2, ExplicitStructTypeExprType}:          _GotoState74Action,
	{_State2, AnonymousFuncExprType}:               _GotoState58Action,
	{_State3, IntegerLiteralToken}:                 _GotoState40Action,
	{_State3, FloatLiteralToken}:                   _GotoState35Action,
	{_State3, RuneLiteralToken}:                    _GotoState48Action,
	{_State3, StringLiteralToken}:                  _GotoState49Action,
	{_State3, IdentifierToken}:                     _GotoState38Action,
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
	{_State3, VarDeclPatternType}:                  _GotoState103Action,
	{_State3, VarOrLetType}:                        _GotoState24Action,
	{_State3, ExprType}:                            _GotoState7Action,
	{_State3, OptionalLabelDeclType}:               _GotoState90Action,
	{_State3, SequenceExprType}:                    _GotoState105Action,
	{_State3, CallExprType}:                        _GotoState70Action,
	{_State3, AtomExprType}:                        _GotoState62Action,
	{_State3, ParseErrorExprType}:                  _GotoState92Action,
	{_State3, LiteralExprType}:                     _GotoState87Action,
	{_State3, IdentifierExprType}:                  _GotoState79Action,
	{_State3, BlockExprType}:                       _GotoState69Action,
	{_State3, InitializeExprType}:                  _GotoState84Action,
	{_State3, ImplicitStructExprType}:              _GotoState80Action,
	{_State3, AccessibleExprType}:                  _GotoState104Action,
	{_State3, AccessExprType}:                      _GotoState54Action,
	{_State3, IndexExprType}:                       _GotoState82Action,
	{_State3, PostfixableExprType}:                 _GotoState94Action,
	{_State3, PostfixUnaryExprType}:                _GotoState93Action,
	{_State3, PrefixableExprType}:                  _GotoState97Action,
	{_State3, PrefixUnaryExprType}:                 _GotoState95Action,
	{_State3, PrefixUnaryOpType}:                   _GotoState96Action,
	{_State3, MulExprType}:                         _GotoState89Action,
	{_State3, BinaryMulExprType}:                   _GotoState66Action,
	{_State3, AddExprType}:                         _GotoState56Action,
	{_State3, BinaryAddExprType}:                   _GotoState63Action,
	{_State3, CmpExprType}:                         _GotoState73Action,
	{_State3, BinaryCmpExprType}:                   _GotoState65Action,
	{_State3, AndExprType}:                         _GotoState57Action,
	{_State3, BinaryAndExprType}:                   _GotoState64Action,
	{_State3, OrExprType}:                          _GotoState91Action,
	{_State3, BinaryOrExprType}:                    _GotoState68Action,
	{_State3, InitializableTypeExprType}:           _GotoState83Action,
	{_State3, SliceTypeExprType}:                   _GotoState100Action,
	{_State3, ArrayTypeExprType}:                   _GotoState59Action,
	{_State3, MapTypeExprType}:                     _GotoState88Action,
	{_State3, ExplicitStructTypeExprType}:          _GotoState74Action,
	{_State3, AnonymousFuncExprType}:               _GotoState58Action,
	{_State4, IdentifierToken}:                     _GotoState112Action,
	{_State4, StructToken}:                         _GotoState50Action,
	{_State4, EnumToken}:                           _GotoState109Action,
	{_State4, TraitToken}:                          _GotoState117Action,
	{_State4, FuncToken}:                           _GotoState111Action,
	{_State4, LparenToken}:                         _GotoState113Action,
	{_State4, LbracketToken}:                       _GotoState42Action,
	{_State4, DotToken}:                            _GotoState108Action,
	{_State4, QuestionToken}:                       _GotoState115Action,
	{_State4, ExclaimToken}:                        _GotoState110Action,
	{_State4, TildeTildeToken}:                     _GotoState116Action,
	{_State4, BitNegToken}:                         _GotoState107Action,
	{_State4, BitAndToken}:                         _GotoState106Action,
	{_State4, ParseErrorToken}:                     _GotoState114Action,
	{_State4, InitializableTypeExprType}:           _GotoState125Action,
	{_State4, SliceTypeExprType}:                   _GotoState100Action,
	{_State4, ArrayTypeExprType}:                   _GotoState59Action,
	{_State4, MapTypeExprType}:                     _GotoState88Action,
	{_State4, AtomTypeExprType}:                    _GotoState118Action,
	{_State4, NamedTypeExprType}:                   _GotoState126Action,
	{_State4, InferredTypeExprType}:                _GotoState124Action,
	{_State4, ParseErrorTypeExprType}:              _GotoState127Action,
	{_State4, ReturnableTypeExprType}:              _GotoState130Action,
	{_State4, PrefixedUnaryTypeExprType}:           _GotoState129Action,
	{_State4, PrefixUnaryTypeOpType}:               _GotoState128Action,
	{_State4, TypeExprType}:                        _GotoState8Action,
	{_State4, BinaryTypeExprType}:                  _GotoState119Action,
	{_State4, ImplicitStructTypeExprType}:          _GotoState123Action,
	{_State4, ExplicitStructTypeExprType}:          _GotoState74Action,
	{_State4, ImplicitEnumTypeExprType}:            _GotoState122Action,
	{_State4, ExplicitEnumTypeExprType}:            _GotoState120Action,
	{_State4, TraitTypeExprType}:                   _GotoState131Action,
	{_State4, FuncTypeExprType}:                    _GotoState121Action,
	{_State8, AddToken}:                            _GotoState240Action,
	{_State8, SubToken}:                            _GotoState242Action,
	{_State8, MulToken}:                            _GotoState241Action,
	{_State8, BinaryTypeOpType}:                    _GotoState243Action,
	{_State10, IdentifierToken}:                    _GotoState132Action,
	{_State10, LparenToken}:                        _GotoState133Action,
	{_State11, IntegerLiteralToken}:                _GotoState40Action,
	{_State11, FloatLiteralToken}:                  _GotoState35Action,
	{_State11, RuneLiteralToken}:                   _GotoState48Action,
	{_State11, StringLiteralToken}:                 _GotoState49Action,
	{_State11, IdentifierToken}:                    _GotoState38Action,
	{_State11, TrueToken}:                          _GotoState52Action,
	{_State11, FalseToken}:                         _GotoState34Action,
	{_State11, CaseToken}:                          _GotoState29Action,
	{_State11, DefaultToken}:                       _GotoState31Action,
	{_State11, ReturnToken}:                        _GotoState47Action,
	{_State11, BreakToken}:                         _GotoState28Action,
	{_State11, ContinueToken}:                      _GotoState30Action,
	{_State11, FallthroughToken}:                   _GotoState33Action,
	{_State11, ImportToken}:                        _GotoState39Action,
	{_State11, UnsafeToken}:                        _GotoState53Action,
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
	{_State11, ProperStatementsType}:               _GotoState134Action,
	{_State11, StatementsType}:                     _GotoState136Action,
	{_State11, SimpleStatementType}:                _GotoState99Action,
	{_State11, StatementType}:                      _GotoState135Action,
	{_State11, ExprOrImproperStructStatementType}:  _GotoState76Action,
	{_State11, ExprsType}:                          _GotoState77Action,
	{_State11, CallbackOpType}:                     _GotoState71Action,
	{_State11, CallbackOpStatementType}:            _GotoState72Action,
	{_State11, UnsafeStatementType}:                _GotoState102Action,
	{_State11, JumpStatementType}:                  _GotoState85Action,
	{_State11, JumpTypeType}:                       _GotoState86Action,
	{_State11, FallthroughStatementType}:           _GotoState78Action,
	{_State11, AssignStatementType}:                _GotoState61Action,
	{_State11, UnaryOpAssignStatementType}:         _GotoState101Action,
	{_State11, BinaryOpAssignStatementType}:        _GotoState67Action,
	{_State11, ImportStatementType}:                _GotoState81Action,
	{_State11, VarDeclPatternType}:                 _GotoState103Action,
	{_State11, VarOrLetType}:                       _GotoState24Action,
	{_State11, AssignPatternType}:                  _GotoState60Action,
	{_State11, ExprType}:                           _GotoState75Action,
	{_State11, OptionalLabelDeclType}:              _GotoState90Action,
	{_State11, SequenceExprType}:                   _GotoState98Action,
	{_State11, CallExprType}:                       _GotoState70Action,
	{_State11, AtomExprType}:                       _GotoState62Action,
	{_State11, ParseErrorExprType}:                 _GotoState92Action,
	{_State11, LiteralExprType}:                    _GotoState87Action,
	{_State11, IdentifierExprType}:                 _GotoState79Action,
	{_State11, BlockExprType}:                      _GotoState69Action,
	{_State11, InitializeExprType}:                 _GotoState84Action,
	{_State11, ImplicitStructExprType}:             _GotoState80Action,
	{_State11, AccessibleExprType}:                 _GotoState55Action,
	{_State11, AccessExprType}:                     _GotoState54Action,
	{_State11, IndexExprType}:                      _GotoState82Action,
	{_State11, PostfixableExprType}:                _GotoState94Action,
	{_State11, PostfixUnaryExprType}:               _GotoState93Action,
	{_State11, PrefixableExprType}:                 _GotoState97Action,
	{_State11, PrefixUnaryExprType}:                _GotoState95Action,
	{_State11, PrefixUnaryOpType}:                  _GotoState96Action,
	{_State11, MulExprType}:                        _GotoState89Action,
	{_State11, BinaryMulExprType}:                  _GotoState66Action,
	{_State11, AddExprType}:                        _GotoState56Action,
	{_State11, BinaryAddExprType}:                  _GotoState63Action,
	{_State11, CmpExprType}:                        _GotoState73Action,
	{_State11, BinaryCmpExprType}:                  _GotoState65Action,
	{_State11, AndExprType}:                        _GotoState57Action,
	{_State11, BinaryAndExprType}:                  _GotoState64Action,
	{_State11, OrExprType}:                         _GotoState91Action,
	{_State11, BinaryOrExprType}:                   _GotoState68Action,
	{_State11, InitializableTypeExprType}:          _GotoState83Action,
	{_State11, SliceTypeExprType}:                  _GotoState100Action,
	{_State11, ArrayTypeExprType}:                  _GotoState59Action,
	{_State11, MapTypeExprType}:                    _GotoState88Action,
	{_State11, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State11, AnonymousFuncExprType}:              _GotoState58Action,
	{_State13, LbraceToken}:                        _GotoState11Action,
	{_State13, StatementBlockType}:                 _GotoState137Action,
	{_State14, IdentifierToken}:                    _GotoState138Action,
	{_State20, NewlinesToken}:                      _GotoState139Action,
	{_State23, AssignToken}:                        _GotoState140Action,
	{_State24, IdentifierToken}:                    _GotoState141Action,
	{_State24, LparenToken}:                        _GotoState142Action,
	{_State24, VarPatternType}:                     _GotoState144Action,
	{_State24, TuplePatternType}:                   _GotoState143Action,
	{_State29, IntegerLiteralToken}:                _GotoState40Action,
	{_State29, FloatLiteralToken}:                  _GotoState35Action,
	{_State29, RuneLiteralToken}:                   _GotoState48Action,
	{_State29, StringLiteralToken}:                 _GotoState49Action,
	{_State29, IdentifierToken}:                    _GotoState38Action,
	{_State29, TrueToken}:                          _GotoState52Action,
	{_State29, FalseToken}:                         _GotoState34Action,
	{_State29, StructToken}:                        _GotoState50Action,
	{_State29, FuncToken}:                          _GotoState36Action,
	{_State29, VarToken}:                           _GotoState146Action,
	{_State29, LetToken}:                           _GotoState12Action,
	{_State29, NotToken}:                           _GotoState45Action,
	{_State29, LabelDeclToken}:                     _GotoState41Action,
	{_State29, LparenToken}:                        _GotoState43Action,
	{_State29, LbracketToken}:                      _GotoState42Action,
	{_State29, DotToken}:                           _GotoState145Action,
	{_State29, SubToken}:                           _GotoState51Action,
	{_State29, MulToken}:                           _GotoState44Action,
	{_State29, BitNegToken}:                        _GotoState27Action,
	{_State29, BitAndToken}:                        _GotoState26Action,
	{_State29, GreaterToken}:                       _GotoState37Action,
	{_State29, ParseErrorToken}:                    _GotoState46Action,
	{_State29, CasePatternsType}:                   _GotoState149Action,
	{_State29, VarDeclPatternType}:                 _GotoState103Action,
	{_State29, VarOrLetType}:                       _GotoState24Action,
	{_State29, AssignPatternType}:                  _GotoState147Action,
	{_State29, CasePatternType}:                    _GotoState148Action,
	{_State29, OptionalLabelDeclType}:              _GotoState150Action,
	{_State29, SequenceExprType}:                   _GotoState151Action,
	{_State29, CallExprType}:                       _GotoState70Action,
	{_State29, AtomExprType}:                       _GotoState62Action,
	{_State29, ParseErrorExprType}:                 _GotoState92Action,
	{_State29, LiteralExprType}:                    _GotoState87Action,
	{_State29, IdentifierExprType}:                 _GotoState79Action,
	{_State29, BlockExprType}:                      _GotoState69Action,
	{_State29, InitializeExprType}:                 _GotoState84Action,
	{_State29, ImplicitStructExprType}:             _GotoState80Action,
	{_State29, AccessibleExprType}:                 _GotoState104Action,
	{_State29, AccessExprType}:                     _GotoState54Action,
	{_State29, IndexExprType}:                      _GotoState82Action,
	{_State29, PostfixableExprType}:                _GotoState94Action,
	{_State29, PostfixUnaryExprType}:               _GotoState93Action,
	{_State29, PrefixableExprType}:                 _GotoState97Action,
	{_State29, PrefixUnaryExprType}:                _GotoState95Action,
	{_State29, PrefixUnaryOpType}:                  _GotoState96Action,
	{_State29, MulExprType}:                        _GotoState89Action,
	{_State29, BinaryMulExprType}:                  _GotoState66Action,
	{_State29, AddExprType}:                        _GotoState56Action,
	{_State29, BinaryAddExprType}:                  _GotoState63Action,
	{_State29, CmpExprType}:                        _GotoState73Action,
	{_State29, BinaryCmpExprType}:                  _GotoState65Action,
	{_State29, AndExprType}:                        _GotoState57Action,
	{_State29, BinaryAndExprType}:                  _GotoState64Action,
	{_State29, OrExprType}:                         _GotoState91Action,
	{_State29, BinaryOrExprType}:                   _GotoState68Action,
	{_State29, InitializableTypeExprType}:          _GotoState83Action,
	{_State29, SliceTypeExprType}:                  _GotoState100Action,
	{_State29, ArrayTypeExprType}:                  _GotoState59Action,
	{_State29, MapTypeExprType}:                    _GotoState88Action,
	{_State29, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State29, AnonymousFuncExprType}:              _GotoState58Action,
	{_State31, ColonToken}:                         _GotoState152Action,
	{_State36, LparenToken}:                        _GotoState153Action,
	{_State37, IntegerLiteralToken}:                _GotoState40Action,
	{_State37, FloatLiteralToken}:                  _GotoState35Action,
	{_State37, RuneLiteralToken}:                   _GotoState48Action,
	{_State37, StringLiteralToken}:                 _GotoState49Action,
	{_State37, IdentifierToken}:                    _GotoState38Action,
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
	{_State37, VarDeclPatternType}:                 _GotoState103Action,
	{_State37, VarOrLetType}:                       _GotoState24Action,
	{_State37, OptionalLabelDeclType}:              _GotoState150Action,
	{_State37, SequenceExprType}:                   _GotoState154Action,
	{_State37, CallExprType}:                       _GotoState70Action,
	{_State37, AtomExprType}:                       _GotoState62Action,
	{_State37, ParseErrorExprType}:                 _GotoState92Action,
	{_State37, LiteralExprType}:                    _GotoState87Action,
	{_State37, IdentifierExprType}:                 _GotoState79Action,
	{_State37, BlockExprType}:                      _GotoState69Action,
	{_State37, InitializeExprType}:                 _GotoState84Action,
	{_State37, ImplicitStructExprType}:             _GotoState80Action,
	{_State37, AccessibleExprType}:                 _GotoState104Action,
	{_State37, AccessExprType}:                     _GotoState54Action,
	{_State37, IndexExprType}:                      _GotoState82Action,
	{_State37, PostfixableExprType}:                _GotoState94Action,
	{_State37, PostfixUnaryExprType}:               _GotoState93Action,
	{_State37, PrefixableExprType}:                 _GotoState97Action,
	{_State37, PrefixUnaryExprType}:                _GotoState95Action,
	{_State37, PrefixUnaryOpType}:                  _GotoState96Action,
	{_State37, MulExprType}:                        _GotoState89Action,
	{_State37, BinaryMulExprType}:                  _GotoState66Action,
	{_State37, AddExprType}:                        _GotoState56Action,
	{_State37, BinaryAddExprType}:                  _GotoState63Action,
	{_State37, CmpExprType}:                        _GotoState73Action,
	{_State37, BinaryCmpExprType}:                  _GotoState65Action,
	{_State37, AndExprType}:                        _GotoState57Action,
	{_State37, BinaryAndExprType}:                  _GotoState64Action,
	{_State37, OrExprType}:                         _GotoState91Action,
	{_State37, BinaryOrExprType}:                   _GotoState68Action,
	{_State37, InitializableTypeExprType}:          _GotoState83Action,
	{_State37, SliceTypeExprType}:                  _GotoState100Action,
	{_State37, ArrayTypeExprType}:                  _GotoState59Action,
	{_State37, MapTypeExprType}:                    _GotoState88Action,
	{_State37, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State37, AnonymousFuncExprType}:              _GotoState58Action,
	{_State39, StringLiteralToken}:                 _GotoState156Action,
	{_State39, LparenToken}:                        _GotoState155Action,
	{_State39, ImportClauseType}:                   _GotoState157Action,
	{_State42, IdentifierToken}:                    _GotoState112Action,
	{_State42, StructToken}:                        _GotoState50Action,
	{_State42, EnumToken}:                          _GotoState109Action,
	{_State42, TraitToken}:                         _GotoState117Action,
	{_State42, FuncToken}:                          _GotoState111Action,
	{_State42, LparenToken}:                        _GotoState113Action,
	{_State42, LbracketToken}:                      _GotoState42Action,
	{_State42, DotToken}:                           _GotoState108Action,
	{_State42, QuestionToken}:                      _GotoState115Action,
	{_State42, ExclaimToken}:                       _GotoState110Action,
	{_State42, TildeTildeToken}:                    _GotoState116Action,
	{_State42, BitNegToken}:                        _GotoState107Action,
	{_State42, BitAndToken}:                        _GotoState106Action,
	{_State42, ParseErrorToken}:                    _GotoState114Action,
	{_State42, InitializableTypeExprType}:          _GotoState125Action,
	{_State42, SliceTypeExprType}:                  _GotoState100Action,
	{_State42, ArrayTypeExprType}:                  _GotoState59Action,
	{_State42, MapTypeExprType}:                    _GotoState88Action,
	{_State42, AtomTypeExprType}:                   _GotoState118Action,
	{_State42, NamedTypeExprType}:                  _GotoState126Action,
	{_State42, InferredTypeExprType}:               _GotoState124Action,
	{_State42, ParseErrorTypeExprType}:             _GotoState127Action,
	{_State42, ReturnableTypeExprType}:             _GotoState130Action,
	{_State42, PrefixedUnaryTypeExprType}:          _GotoState129Action,
	{_State42, PrefixUnaryTypeOpType}:              _GotoState128Action,
	{_State42, TypeExprType}:                       _GotoState158Action,
	{_State42, BinaryTypeExprType}:                 _GotoState119Action,
	{_State42, ImplicitStructTypeExprType}:         _GotoState123Action,
	{_State42, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State42, ImplicitEnumTypeExprType}:           _GotoState122Action,
	{_State42, ExplicitEnumTypeExprType}:           _GotoState120Action,
	{_State42, TraitTypeExprType}:                  _GotoState131Action,
	{_State42, FuncTypeExprType}:                   _GotoState121Action,
	{_State43, IntegerLiteralToken}:                _GotoState40Action,
	{_State43, FloatLiteralToken}:                  _GotoState35Action,
	{_State43, RuneLiteralToken}:                   _GotoState48Action,
	{_State43, StringLiteralToken}:                 _GotoState49Action,
	{_State43, IdentifierToken}:                    _GotoState161Action,
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
	{_State43, ColonToken}:                         _GotoState159Action,
	{_State43, EllipsisToken}:                      _GotoState160Action,
	{_State43, SubToken}:                           _GotoState51Action,
	{_State43, MulToken}:                           _GotoState44Action,
	{_State43, BitNegToken}:                        _GotoState27Action,
	{_State43, BitAndToken}:                        _GotoState26Action,
	{_State43, GreaterToken}:                       _GotoState37Action,
	{_State43, ParseErrorToken}:                    _GotoState46Action,
	{_State43, VarDeclPatternType}:                 _GotoState103Action,
	{_State43, VarOrLetType}:                       _GotoState24Action,
	{_State43, ExprType}:                           _GotoState165Action,
	{_State43, OptionalLabelDeclType}:              _GotoState90Action,
	{_State43, SequenceExprType}:                   _GotoState105Action,
	{_State43, CallExprType}:                       _GotoState70Action,
	{_State43, ProperArgumentsType}:                _GotoState166Action,
	{_State43, ArgumentsType}:                      _GotoState163Action,
	{_State43, ArgumentType}:                       _GotoState162Action,
	{_State43, ColonExprType}:                      _GotoState164Action,
	{_State43, AtomExprType}:                       _GotoState62Action,
	{_State43, ParseErrorExprType}:                 _GotoState92Action,
	{_State43, LiteralExprType}:                    _GotoState87Action,
	{_State43, IdentifierExprType}:                 _GotoState79Action,
	{_State43, BlockExprType}:                      _GotoState69Action,
	{_State43, InitializeExprType}:                 _GotoState84Action,
	{_State43, ImplicitStructExprType}:             _GotoState80Action,
	{_State43, AccessibleExprType}:                 _GotoState104Action,
	{_State43, AccessExprType}:                     _GotoState54Action,
	{_State43, IndexExprType}:                      _GotoState82Action,
	{_State43, PostfixableExprType}:                _GotoState94Action,
	{_State43, PostfixUnaryExprType}:               _GotoState93Action,
	{_State43, PrefixableExprType}:                 _GotoState97Action,
	{_State43, PrefixUnaryExprType}:                _GotoState95Action,
	{_State43, PrefixUnaryOpType}:                  _GotoState96Action,
	{_State43, MulExprType}:                        _GotoState89Action,
	{_State43, BinaryMulExprType}:                  _GotoState66Action,
	{_State43, AddExprType}:                        _GotoState56Action,
	{_State43, BinaryAddExprType}:                  _GotoState63Action,
	{_State43, CmpExprType}:                        _GotoState73Action,
	{_State43, BinaryCmpExprType}:                  _GotoState65Action,
	{_State43, AndExprType}:                        _GotoState57Action,
	{_State43, BinaryAndExprType}:                  _GotoState64Action,
	{_State43, OrExprType}:                         _GotoState91Action,
	{_State43, BinaryOrExprType}:                   _GotoState68Action,
	{_State43, InitializableTypeExprType}:          _GotoState83Action,
	{_State43, SliceTypeExprType}:                  _GotoState100Action,
	{_State43, ArrayTypeExprType}:                  _GotoState59Action,
	{_State43, MapTypeExprType}:                    _GotoState88Action,
	{_State43, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State43, AnonymousFuncExprType}:              _GotoState58Action,
	{_State50, LparenToken}:                        _GotoState167Action,
	{_State53, LessToken}:                          _GotoState168Action,
	{_State55, LbracketToken}:                      _GotoState180Action,
	{_State55, DotToken}:                           _GotoState179Action,
	{_State55, QuestionToken}:                      _GotoState183Action,
	{_State55, DollarLbracketToken}:                _GotoState178Action,
	{_State55, AddAssignToken}:                     _GotoState169Action,
	{_State55, SubAssignToken}:                     _GotoState184Action,
	{_State55, MulAssignToken}:                     _GotoState182Action,
	{_State55, DivAssignToken}:                     _GotoState177Action,
	{_State55, ModAssignToken}:                     _GotoState181Action,
	{_State55, AddOneAssignToken}:                  _GotoState170Action,
	{_State55, SubOneAssignToken}:                  _GotoState185Action,
	{_State55, BitNegAssignToken}:                  _GotoState173Action,
	{_State55, BitAndAssignToken}:                  _GotoState171Action,
	{_State55, BitOrAssignToken}:                   _GotoState174Action,
	{_State55, BitXorAssignToken}:                  _GotoState176Action,
	{_State55, BitLshiftAssignToken}:               _GotoState172Action,
	{_State55, BitRshiftAssignToken}:               _GotoState175Action,
	{_State55, UnaryOpAssignType}:                  _GotoState188Action,
	{_State55, BinaryOpAssignType}:                 _GotoState186Action,
	{_State55, GenericTypeArgumentsType}:           _GotoState187Action,
	{_State56, AddToken}:                           _GotoState189Action,
	{_State56, SubToken}:                           _GotoState192Action,
	{_State56, BitXorToken}:                        _GotoState191Action,
	{_State56, BitOrToken}:                         _GotoState190Action,
	{_State56, AddOpType}:                          _GotoState193Action,
	{_State57, AndToken}:                           _GotoState194Action,
	{_State60, AssignToken}:                        _GotoState195Action,
	{_State71, IntegerLiteralToken}:                _GotoState40Action,
	{_State71, FloatLiteralToken}:                  _GotoState35Action,
	{_State71, RuneLiteralToken}:                   _GotoState48Action,
	{_State71, StringLiteralToken}:                 _GotoState49Action,
	{_State71, IdentifierToken}:                    _GotoState38Action,
	{_State71, TrueToken}:                          _GotoState52Action,
	{_State71, FalseToken}:                         _GotoState34Action,
	{_State71, StructToken}:                        _GotoState50Action,
	{_State71, FuncToken}:                          _GotoState36Action,
	{_State71, LabelDeclToken}:                     _GotoState41Action,
	{_State71, LparenToken}:                        _GotoState43Action,
	{_State71, LbracketToken}:                      _GotoState42Action,
	{_State71, ParseErrorToken}:                    _GotoState46Action,
	{_State71, OptionalLabelDeclType}:              _GotoState150Action,
	{_State71, CallExprType}:                       _GotoState197Action,
	{_State71, AtomExprType}:                       _GotoState62Action,
	{_State71, ParseErrorExprType}:                 _GotoState92Action,
	{_State71, LiteralExprType}:                    _GotoState87Action,
	{_State71, IdentifierExprType}:                 _GotoState79Action,
	{_State71, BlockExprType}:                      _GotoState69Action,
	{_State71, InitializeExprType}:                 _GotoState84Action,
	{_State71, ImplicitStructExprType}:             _GotoState80Action,
	{_State71, AccessibleExprType}:                 _GotoState196Action,
	{_State71, AccessExprType}:                     _GotoState54Action,
	{_State71, IndexExprType}:                      _GotoState82Action,
	{_State71, InitializableTypeExprType}:          _GotoState83Action,
	{_State71, SliceTypeExprType}:                  _GotoState100Action,
	{_State71, ArrayTypeExprType}:                  _GotoState59Action,
	{_State71, MapTypeExprType}:                    _GotoState88Action,
	{_State71, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State71, AnonymousFuncExprType}:              _GotoState58Action,
	{_State73, EqualToken}:                         _GotoState198Action,
	{_State73, NotEqualToken}:                      _GotoState203Action,
	{_State73, LessToken}:                          _GotoState201Action,
	{_State73, LessOrEqualToken}:                   _GotoState202Action,
	{_State73, GreaterToken}:                       _GotoState199Action,
	{_State73, GreaterOrEqualToken}:                _GotoState200Action,
	{_State73, CmpOpType}:                          _GotoState204Action,
	{_State77, CommaToken}:                         _GotoState205Action,
	{_State83, LparenToken}:                        _GotoState206Action,
	{_State86, IntegerLiteralToken}:                _GotoState40Action,
	{_State86, FloatLiteralToken}:                  _GotoState35Action,
	{_State86, RuneLiteralToken}:                   _GotoState48Action,
	{_State86, StringLiteralToken}:                 _GotoState49Action,
	{_State86, IdentifierToken}:                    _GotoState38Action,
	{_State86, TrueToken}:                          _GotoState52Action,
	{_State86, FalseToken}:                         _GotoState34Action,
	{_State86, StructToken}:                        _GotoState50Action,
	{_State86, FuncToken}:                          _GotoState36Action,
	{_State86, VarToken}:                           _GotoState15Action,
	{_State86, LetToken}:                           _GotoState12Action,
	{_State86, NotToken}:                           _GotoState45Action,
	{_State86, LabelDeclToken}:                     _GotoState41Action,
	{_State86, JumpLabelToken}:                     _GotoState207Action,
	{_State86, LparenToken}:                        _GotoState43Action,
	{_State86, LbracketToken}:                      _GotoState42Action,
	{_State86, SubToken}:                           _GotoState51Action,
	{_State86, MulToken}:                           _GotoState44Action,
	{_State86, BitNegToken}:                        _GotoState27Action,
	{_State86, BitAndToken}:                        _GotoState26Action,
	{_State86, GreaterToken}:                       _GotoState37Action,
	{_State86, ParseErrorToken}:                    _GotoState46Action,
	{_State86, ExprsType}:                          _GotoState208Action,
	{_State86, VarDeclPatternType}:                 _GotoState103Action,
	{_State86, VarOrLetType}:                       _GotoState24Action,
	{_State86, ExprType}:                           _GotoState75Action,
	{_State86, OptionalLabelDeclType}:              _GotoState90Action,
	{_State86, SequenceExprType}:                   _GotoState105Action,
	{_State86, CallExprType}:                       _GotoState70Action,
	{_State86, AtomExprType}:                       _GotoState62Action,
	{_State86, ParseErrorExprType}:                 _GotoState92Action,
	{_State86, LiteralExprType}:                    _GotoState87Action,
	{_State86, IdentifierExprType}:                 _GotoState79Action,
	{_State86, BlockExprType}:                      _GotoState69Action,
	{_State86, InitializeExprType}:                 _GotoState84Action,
	{_State86, ImplicitStructExprType}:             _GotoState80Action,
	{_State86, AccessibleExprType}:                 _GotoState104Action,
	{_State86, AccessExprType}:                     _GotoState54Action,
	{_State86, IndexExprType}:                      _GotoState82Action,
	{_State86, PostfixableExprType}:                _GotoState94Action,
	{_State86, PostfixUnaryExprType}:               _GotoState93Action,
	{_State86, PrefixableExprType}:                 _GotoState97Action,
	{_State86, PrefixUnaryExprType}:                _GotoState95Action,
	{_State86, PrefixUnaryOpType}:                  _GotoState96Action,
	{_State86, MulExprType}:                        _GotoState89Action,
	{_State86, BinaryMulExprType}:                  _GotoState66Action,
	{_State86, AddExprType}:                        _GotoState56Action,
	{_State86, BinaryAddExprType}:                  _GotoState63Action,
	{_State86, CmpExprType}:                        _GotoState73Action,
	{_State86, BinaryCmpExprType}:                  _GotoState65Action,
	{_State86, AndExprType}:                        _GotoState57Action,
	{_State86, BinaryAndExprType}:                  _GotoState64Action,
	{_State86, OrExprType}:                         _GotoState91Action,
	{_State86, BinaryOrExprType}:                   _GotoState68Action,
	{_State86, InitializableTypeExprType}:          _GotoState83Action,
	{_State86, SliceTypeExprType}:                  _GotoState100Action,
	{_State86, ArrayTypeExprType}:                  _GotoState59Action,
	{_State86, MapTypeExprType}:                    _GotoState88Action,
	{_State86, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State86, AnonymousFuncExprType}:              _GotoState58Action,
	{_State89, MulToken}:                           _GotoState214Action,
	{_State89, DivToken}:                           _GotoState212Action,
	{_State89, ModToken}:                           _GotoState213Action,
	{_State89, BitAndToken}:                        _GotoState209Action,
	{_State89, BitLshiftToken}:                     _GotoState210Action,
	{_State89, BitRshiftToken}:                     _GotoState211Action,
	{_State89, MulOpType}:                          _GotoState215Action,
	{_State90, IfToken}:                            _GotoState218Action,
	{_State90, SwitchToken}:                        _GotoState219Action,
	{_State90, ForToken}:                           _GotoState217Action,
	{_State90, DoToken}:                            _GotoState216Action,
	{_State90, LbraceToken}:                        _GotoState11Action,
	{_State90, StatementBlockType}:                 _GotoState222Action,
	{_State90, IfExprType}:                         _GotoState220Action,
	{_State90, SwitchExprType}:                     _GotoState223Action,
	{_State90, LoopExprType}:                       _GotoState221Action,
	{_State91, OrToken}:                            _GotoState224Action,
	{_State96, IntegerLiteralToken}:                _GotoState40Action,
	{_State96, FloatLiteralToken}:                  _GotoState35Action,
	{_State96, RuneLiteralToken}:                   _GotoState48Action,
	{_State96, StringLiteralToken}:                 _GotoState49Action,
	{_State96, IdentifierToken}:                    _GotoState38Action,
	{_State96, TrueToken}:                          _GotoState52Action,
	{_State96, FalseToken}:                         _GotoState34Action,
	{_State96, StructToken}:                        _GotoState50Action,
	{_State96, FuncToken}:                          _GotoState36Action,
	{_State96, NotToken}:                           _GotoState45Action,
	{_State96, LabelDeclToken}:                     _GotoState41Action,
	{_State96, LparenToken}:                        _GotoState43Action,
	{_State96, LbracketToken}:                      _GotoState42Action,
	{_State96, SubToken}:                           _GotoState51Action,
	{_State96, MulToken}:                           _GotoState44Action,
	{_State96, BitNegToken}:                        _GotoState27Action,
	{_State96, BitAndToken}:                        _GotoState26Action,
	{_State96, ParseErrorToken}:                    _GotoState46Action,
	{_State96, OptionalLabelDeclType}:              _GotoState150Action,
	{_State96, CallExprType}:                       _GotoState70Action,
	{_State96, AtomExprType}:                       _GotoState62Action,
	{_State96, ParseErrorExprType}:                 _GotoState92Action,
	{_State96, LiteralExprType}:                    _GotoState87Action,
	{_State96, IdentifierExprType}:                 _GotoState79Action,
	{_State96, BlockExprType}:                      _GotoState69Action,
	{_State96, InitializeExprType}:                 _GotoState84Action,
	{_State96, ImplicitStructExprType}:             _GotoState80Action,
	{_State96, AccessibleExprType}:                 _GotoState104Action,
	{_State96, AccessExprType}:                     _GotoState54Action,
	{_State96, IndexExprType}:                      _GotoState82Action,
	{_State96, PostfixableExprType}:                _GotoState94Action,
	{_State96, PostfixUnaryExprType}:               _GotoState93Action,
	{_State96, PrefixableExprType}:                 _GotoState225Action,
	{_State96, PrefixUnaryExprType}:                _GotoState95Action,
	{_State96, PrefixUnaryOpType}:                  _GotoState96Action,
	{_State96, InitializableTypeExprType}:          _GotoState83Action,
	{_State96, SliceTypeExprType}:                  _GotoState100Action,
	{_State96, ArrayTypeExprType}:                  _GotoState59Action,
	{_State96, MapTypeExprType}:                    _GotoState88Action,
	{_State96, ExplicitStructTypeExprType}:         _GotoState74Action,
	{_State96, AnonymousFuncExprType}:              _GotoState58Action,
	{_State104, LbracketToken}:                     _GotoState180Action,
	{_State104, DotToken}:                          _GotoState179Action,
	{_State104, QuestionToken}:                     _GotoState183Action,
	{_State104, DollarLbracketToken}:               _GotoState178Action,
	{_State104, GenericTypeArgumentsType}:          _GotoState187Action,
	{_State109, LparenToken}:                       _GotoState226Action,
	{_State111, LparenToken}:                       _GotoState227Action,
	{_State112, DotToken}:                          _GotoState228Action,
	{_State112, DollarLbracketToken}:               _GotoState178Action,
	{_State112, GenericTypeArgumentsType}:          _GotoState229Action,
	{_State113, IdentifierToken}:                   _GotoState230Action,
	{_State113, UnsafeToken}:                       _GotoState53Action,
	{_State113, StructToken}:                       _GotoState50Action,
	{_State113, EnumToken}:                         _GotoState109Action,
	{_State113, TraitToken}:                        _GotoState117Action,
	{_State113, FuncToken}:                         _GotoState111Action,
	{_State113, LparenToken}:                       _GotoState113Action,
	{_State113, LbracketToken}:                     _GotoState42Action,
	{_State113, DotToken}:                          _GotoState108Action,
	{_State113, QuestionToken}:                     _GotoState115Action,
	{_State113, ExclaimToken}:                      _GotoState110Action,
	{_State113, TildeTildeToken}:                   _GotoState116Action,
	{_State113, BitNegToken}:                       _GotoState107Action,
	{_State113, BitAndToken}:                       _GotoState106Action,
	{_State113, ParseErrorToken}:                   _GotoState114Action,
	{_State113, UnsafeStatementType}:               _GotoState237Action,
	{_State113, InitializableTypeExprType}:         _GotoState125Action,
	{_State113, SliceTypeExprType}:                 _GotoState100Action,
	{_State113, ArrayTypeExprType}:                 _GotoState59Action,
	{_State113, MapTypeExprType}:                   _GotoState88Action,
	{_State113, AtomTypeExprType}:                  _GotoState118Action,
	{_State113, NamedTypeExprType}:                 _GotoState126Action,
	{_State113, InferredTypeExprType}:              _GotoState124Action,
	{_State113, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State113, ReturnableTypeExprType}:            _GotoState130Action,
	{_State113, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State113, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State113, TypeExprType}:                      _GotoState236Action,
	{_State113, BinaryTypeExprType}:                _GotoState119Action,
	{_State113, FieldDefType}:                      _GotoState232Action,
	{_State113, ProperImplicitFieldDefsType}:       _GotoState235Action,
	{_State113, ImplicitFieldDefsType}:             _GotoState234Action,
	{_State113, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State113, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State113, EnumValueDefType}:                  _GotoState231Action,
	{_State113, ImplicitEnumValueDefsType}:         _GotoState233Action,
	{_State113, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State113, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State113, TraitTypeExprType}:                 _GotoState131Action,
	{_State113, FuncTypeExprType}:                  _GotoState121Action,
	{_State117, LparenToken}:                       _GotoState238Action,
	{_State128, IdentifierToken}:                   _GotoState112Action,
	{_State128, StructToken}:                       _GotoState50Action,
	{_State128, EnumToken}:                         _GotoState109Action,
	{_State128, TraitToken}:                        _GotoState117Action,
	{_State128, FuncToken}:                         _GotoState111Action,
	{_State128, LparenToken}:                       _GotoState113Action,
	{_State128, LbracketToken}:                     _GotoState42Action,
	{_State128, DotToken}:                          _GotoState108Action,
	{_State128, QuestionToken}:                     _GotoState115Action,
	{_State128, ExclaimToken}:                      _GotoState110Action,
	{_State128, TildeTildeToken}:                   _GotoState116Action,
	{_State128, BitNegToken}:                       _GotoState107Action,
	{_State128, BitAndToken}:                       _GotoState106Action,
	{_State128, ParseErrorToken}:                   _GotoState114Action,
	{_State128, InitializableTypeExprType}:         _GotoState125Action,
	{_State128, SliceTypeExprType}:                 _GotoState100Action,
	{_State128, ArrayTypeExprType}:                 _GotoState59Action,
	{_State128, MapTypeExprType}:                   _GotoState88Action,
	{_State128, AtomTypeExprType}:                  _GotoState118Action,
	{_State128, NamedTypeExprType}:                 _GotoState126Action,
	{_State128, InferredTypeExprType}:              _GotoState124Action,
	{_State128, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State128, ReturnableTypeExprType}:            _GotoState239Action,
	{_State128, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State128, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State128, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State128, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State128, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State128, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State128, TraitTypeExprType}:                 _GotoState131Action,
	{_State128, FuncTypeExprType}:                  _GotoState121Action,
	{_State132, DollarLbracketToken}:               _GotoState245Action,
	{_State132, AssignToken}:                       _GotoState244Action,
	{_State132, OptionalGenericParametersType}:     _GotoState246Action,
	{_State133, IdentifierToken}:                   _GotoState247Action,
	{_State133, ParameterDefType}:                  _GotoState248Action,
	{_State134, NewlinesToken}:                     _GotoState249Action,
	{_State134, SemicolonToken}:                    _GotoState250Action,
	{_State136, RbraceToken}:                       _GotoState251Action,
	{_State138, DollarLbracketToken}:               _GotoState245Action,
	{_State138, AssignToken}:                       _GotoState252Action,
	{_State138, OptionalGenericParametersType}:     _GotoState253Action,
	{_State139, CommentGroupsToken}:                _GotoState9Action,
	{_State139, PackageToken}:                      _GotoState13Action,
	{_State139, TypeToken}:                         _GotoState14Action,
	{_State139, FuncToken}:                         _GotoState10Action,
	{_State139, VarToken}:                          _GotoState15Action,
	{_State139, LetToken}:                          _GotoState12Action,
	{_State139, LbraceToken}:                       _GotoState11Action,
	{_State139, DefinitionType}:                    _GotoState254Action,
	{_State139, StatementBlockType}:                _GotoState21Action,
	{_State139, VarDeclPatternType}:                _GotoState23Action,
	{_State139, VarOrLetType}:                      _GotoState24Action,
	{_State139, TypeDefType}:                       _GotoState22Action,
	{_State139, NamedFuncDefType}:                  _GotoState18Action,
	{_State139, PackageDefType}:                    _GotoState19Action,
	{_State140, IntegerLiteralToken}:               _GotoState40Action,
	{_State140, FloatLiteralToken}:                 _GotoState35Action,
	{_State140, RuneLiteralToken}:                  _GotoState48Action,
	{_State140, StringLiteralToken}:                _GotoState49Action,
	{_State140, IdentifierToken}:                   _GotoState38Action,
	{_State140, TrueToken}:                         _GotoState52Action,
	{_State140, FalseToken}:                        _GotoState34Action,
	{_State140, StructToken}:                       _GotoState50Action,
	{_State140, FuncToken}:                         _GotoState36Action,
	{_State140, VarToken}:                          _GotoState15Action,
	{_State140, LetToken}:                          _GotoState12Action,
	{_State140, NotToken}:                          _GotoState45Action,
	{_State140, LabelDeclToken}:                    _GotoState41Action,
	{_State140, LparenToken}:                       _GotoState43Action,
	{_State140, LbracketToken}:                     _GotoState42Action,
	{_State140, SubToken}:                          _GotoState51Action,
	{_State140, MulToken}:                          _GotoState44Action,
	{_State140, BitNegToken}:                       _GotoState27Action,
	{_State140, BitAndToken}:                       _GotoState26Action,
	{_State140, GreaterToken}:                      _GotoState37Action,
	{_State140, ParseErrorToken}:                   _GotoState46Action,
	{_State140, VarDeclPatternType}:                _GotoState103Action,
	{_State140, VarOrLetType}:                      _GotoState24Action,
	{_State140, ExprType}:                          _GotoState255Action,
	{_State140, OptionalLabelDeclType}:             _GotoState90Action,
	{_State140, SequenceExprType}:                  _GotoState105Action,
	{_State140, CallExprType}:                      _GotoState70Action,
	{_State140, AtomExprType}:                      _GotoState62Action,
	{_State140, ParseErrorExprType}:                _GotoState92Action,
	{_State140, LiteralExprType}:                   _GotoState87Action,
	{_State140, IdentifierExprType}:                _GotoState79Action,
	{_State140, BlockExprType}:                     _GotoState69Action,
	{_State140, InitializeExprType}:                _GotoState84Action,
	{_State140, ImplicitStructExprType}:            _GotoState80Action,
	{_State140, AccessibleExprType}:                _GotoState104Action,
	{_State140, AccessExprType}:                    _GotoState54Action,
	{_State140, IndexExprType}:                     _GotoState82Action,
	{_State140, PostfixableExprType}:               _GotoState94Action,
	{_State140, PostfixUnaryExprType}:              _GotoState93Action,
	{_State140, PrefixableExprType}:                _GotoState97Action,
	{_State140, PrefixUnaryExprType}:               _GotoState95Action,
	{_State140, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State140, MulExprType}:                       _GotoState89Action,
	{_State140, BinaryMulExprType}:                 _GotoState66Action,
	{_State140, AddExprType}:                       _GotoState56Action,
	{_State140, BinaryAddExprType}:                 _GotoState63Action,
	{_State140, CmpExprType}:                       _GotoState73Action,
	{_State140, BinaryCmpExprType}:                 _GotoState65Action,
	{_State140, AndExprType}:                       _GotoState57Action,
	{_State140, BinaryAndExprType}:                 _GotoState64Action,
	{_State140, OrExprType}:                        _GotoState91Action,
	{_State140, BinaryOrExprType}:                  _GotoState68Action,
	{_State140, InitializableTypeExprType}:         _GotoState83Action,
	{_State140, SliceTypeExprType}:                 _GotoState100Action,
	{_State140, ArrayTypeExprType}:                 _GotoState59Action,
	{_State140, MapTypeExprType}:                   _GotoState88Action,
	{_State140, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State140, AnonymousFuncExprType}:             _GotoState58Action,
	{_State142, IdentifierToken}:                   _GotoState257Action,
	{_State142, LparenToken}:                       _GotoState142Action,
	{_State142, EllipsisToken}:                     _GotoState256Action,
	{_State142, VarPatternType}:                    _GotoState260Action,
	{_State142, TuplePatternType}:                  _GotoState143Action,
	{_State142, FieldVarPatternsType}:              _GotoState259Action,
	{_State142, FieldVarPatternType}:               _GotoState258Action,
	{_State144, IdentifierToken}:                   _GotoState112Action,
	{_State144, StructToken}:                       _GotoState50Action,
	{_State144, EnumToken}:                         _GotoState109Action,
	{_State144, TraitToken}:                        _GotoState117Action,
	{_State144, FuncToken}:                         _GotoState111Action,
	{_State144, LparenToken}:                       _GotoState113Action,
	{_State144, LbracketToken}:                     _GotoState42Action,
	{_State144, DotToken}:                          _GotoState108Action,
	{_State144, QuestionToken}:                     _GotoState115Action,
	{_State144, ExclaimToken}:                      _GotoState110Action,
	{_State144, TildeTildeToken}:                   _GotoState116Action,
	{_State144, BitNegToken}:                       _GotoState107Action,
	{_State144, BitAndToken}:                       _GotoState106Action,
	{_State144, ParseErrorToken}:                   _GotoState114Action,
	{_State144, OptionalTypeExprType}:              _GotoState261Action,
	{_State144, InitializableTypeExprType}:         _GotoState125Action,
	{_State144, SliceTypeExprType}:                 _GotoState100Action,
	{_State144, ArrayTypeExprType}:                 _GotoState59Action,
	{_State144, MapTypeExprType}:                   _GotoState88Action,
	{_State144, AtomTypeExprType}:                  _GotoState118Action,
	{_State144, NamedTypeExprType}:                 _GotoState126Action,
	{_State144, InferredTypeExprType}:              _GotoState124Action,
	{_State144, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State144, ReturnableTypeExprType}:            _GotoState130Action,
	{_State144, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State144, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State144, TypeExprType}:                      _GotoState262Action,
	{_State144, BinaryTypeExprType}:                _GotoState119Action,
	{_State144, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State144, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State144, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State144, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State144, TraitTypeExprType}:                 _GotoState131Action,
	{_State144, FuncTypeExprType}:                  _GotoState121Action,
	{_State145, IdentifierToken}:                   _GotoState263Action,
	{_State146, DotToken}:                          _GotoState264Action,
	{_State149, CommaToken}:                        _GotoState266Action,
	{_State149, ColonToken}:                        _GotoState265Action,
	{_State150, LbraceToken}:                       _GotoState11Action,
	{_State150, StatementBlockType}:                _GotoState222Action,
	{_State152, IntegerLiteralToken}:               _GotoState40Action,
	{_State152, FloatLiteralToken}:                 _GotoState35Action,
	{_State152, RuneLiteralToken}:                  _GotoState48Action,
	{_State152, StringLiteralToken}:                _GotoState49Action,
	{_State152, IdentifierToken}:                   _GotoState38Action,
	{_State152, TrueToken}:                         _GotoState52Action,
	{_State152, FalseToken}:                        _GotoState34Action,
	{_State152, ReturnToken}:                       _GotoState47Action,
	{_State152, BreakToken}:                        _GotoState28Action,
	{_State152, ContinueToken}:                     _GotoState30Action,
	{_State152, FallthroughToken}:                  _GotoState33Action,
	{_State152, UnsafeToken}:                       _GotoState53Action,
	{_State152, StructToken}:                       _GotoState50Action,
	{_State152, FuncToken}:                         _GotoState36Action,
	{_State152, AsyncToken}:                        _GotoState25Action,
	{_State152, DeferToken}:                        _GotoState32Action,
	{_State152, VarToken}:                          _GotoState15Action,
	{_State152, LetToken}:                          _GotoState12Action,
	{_State152, NotToken}:                          _GotoState45Action,
	{_State152, LabelDeclToken}:                    _GotoState41Action,
	{_State152, LparenToken}:                       _GotoState43Action,
	{_State152, LbracketToken}:                     _GotoState42Action,
	{_State152, SubToken}:                          _GotoState51Action,
	{_State152, MulToken}:                          _GotoState44Action,
	{_State152, BitNegToken}:                       _GotoState27Action,
	{_State152, BitAndToken}:                       _GotoState26Action,
	{_State152, GreaterToken}:                      _GotoState37Action,
	{_State152, ParseErrorToken}:                   _GotoState46Action,
	{_State152, SimpleStatementType}:               _GotoState268Action,
	{_State152, OptionalSimpleStatementType}:       _GotoState267Action,
	{_State152, ExprOrImproperStructStatementType}: _GotoState76Action,
	{_State152, ExprsType}:                         _GotoState77Action,
	{_State152, CallbackOpType}:                    _GotoState71Action,
	{_State152, CallbackOpStatementType}:           _GotoState72Action,
	{_State152, UnsafeStatementType}:               _GotoState102Action,
	{_State152, JumpStatementType}:                 _GotoState85Action,
	{_State152, JumpTypeType}:                      _GotoState86Action,
	{_State152, FallthroughStatementType}:          _GotoState78Action,
	{_State152, AssignStatementType}:               _GotoState61Action,
	{_State152, UnaryOpAssignStatementType}:        _GotoState101Action,
	{_State152, BinaryOpAssignStatementType}:       _GotoState67Action,
	{_State152, VarDeclPatternType}:                _GotoState103Action,
	{_State152, VarOrLetType}:                      _GotoState24Action,
	{_State152, AssignPatternType}:                 _GotoState60Action,
	{_State152, ExprType}:                          _GotoState75Action,
	{_State152, OptionalLabelDeclType}:             _GotoState90Action,
	{_State152, SequenceExprType}:                  _GotoState98Action,
	{_State152, CallExprType}:                      _GotoState70Action,
	{_State152, AtomExprType}:                      _GotoState62Action,
	{_State152, ParseErrorExprType}:                _GotoState92Action,
	{_State152, LiteralExprType}:                   _GotoState87Action,
	{_State152, IdentifierExprType}:                _GotoState79Action,
	{_State152, BlockExprType}:                     _GotoState69Action,
	{_State152, InitializeExprType}:                _GotoState84Action,
	{_State152, ImplicitStructExprType}:            _GotoState80Action,
	{_State152, AccessibleExprType}:                _GotoState55Action,
	{_State152, AccessExprType}:                    _GotoState54Action,
	{_State152, IndexExprType}:                     _GotoState82Action,
	{_State152, PostfixableExprType}:               _GotoState94Action,
	{_State152, PostfixUnaryExprType}:              _GotoState93Action,
	{_State152, PrefixableExprType}:                _GotoState97Action,
	{_State152, PrefixUnaryExprType}:               _GotoState95Action,
	{_State152, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State152, MulExprType}:                       _GotoState89Action,
	{_State152, BinaryMulExprType}:                 _GotoState66Action,
	{_State152, AddExprType}:                       _GotoState56Action,
	{_State152, BinaryAddExprType}:                 _GotoState63Action,
	{_State152, CmpExprType}:                       _GotoState73Action,
	{_State152, BinaryCmpExprType}:                 _GotoState65Action,
	{_State152, AndExprType}:                       _GotoState57Action,
	{_State152, BinaryAndExprType}:                 _GotoState64Action,
	{_State152, OrExprType}:                        _GotoState91Action,
	{_State152, BinaryOrExprType}:                  _GotoState68Action,
	{_State152, InitializableTypeExprType}:         _GotoState83Action,
	{_State152, SliceTypeExprType}:                 _GotoState100Action,
	{_State152, ArrayTypeExprType}:                 _GotoState59Action,
	{_State152, MapTypeExprType}:                   _GotoState88Action,
	{_State152, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State152, AnonymousFuncExprType}:             _GotoState58Action,
	{_State153, IdentifierToken}:                   _GotoState247Action,
	{_State153, ParameterDefType}:                  _GotoState269Action,
	{_State153, ProperParameterDefsType}:           _GotoState271Action,
	{_State153, ParameterDefsType}:                 _GotoState270Action,
	{_State155, StringLiteralToken}:                _GotoState156Action,
	{_State155, ImportClauseType}:                  _GotoState272Action,
	{_State155, ProperImportClausesType}:           _GotoState274Action,
	{_State155, ImportClausesType}:                 _GotoState273Action,
	{_State156, AsToken}:                           _GotoState275Action,
	{_State158, RbracketToken}:                     _GotoState278Action,
	{_State158, CommaToken}:                        _GotoState277Action,
	{_State158, ColonToken}:                        _GotoState276Action,
	{_State158, AddToken}:                          _GotoState240Action,
	{_State158, SubToken}:                          _GotoState242Action,
	{_State158, MulToken}:                          _GotoState241Action,
	{_State158, BinaryTypeOpType}:                  _GotoState243Action,
	{_State159, IntegerLiteralToken}:               _GotoState40Action,
	{_State159, FloatLiteralToken}:                 _GotoState35Action,
	{_State159, RuneLiteralToken}:                  _GotoState48Action,
	{_State159, StringLiteralToken}:                _GotoState49Action,
	{_State159, IdentifierToken}:                   _GotoState38Action,
	{_State159, TrueToken}:                         _GotoState52Action,
	{_State159, FalseToken}:                        _GotoState34Action,
	{_State159, StructToken}:                       _GotoState50Action,
	{_State159, FuncToken}:                         _GotoState36Action,
	{_State159, VarToken}:                          _GotoState15Action,
	{_State159, LetToken}:                          _GotoState12Action,
	{_State159, NotToken}:                          _GotoState45Action,
	{_State159, LabelDeclToken}:                    _GotoState41Action,
	{_State159, LparenToken}:                       _GotoState43Action,
	{_State159, LbracketToken}:                     _GotoState42Action,
	{_State159, SubToken}:                          _GotoState51Action,
	{_State159, MulToken}:                          _GotoState44Action,
	{_State159, BitNegToken}:                       _GotoState27Action,
	{_State159, BitAndToken}:                       _GotoState26Action,
	{_State159, GreaterToken}:                      _GotoState37Action,
	{_State159, ParseErrorToken}:                   _GotoState46Action,
	{_State159, VarDeclPatternType}:                _GotoState103Action,
	{_State159, VarOrLetType}:                      _GotoState24Action,
	{_State159, ExprType}:                          _GotoState279Action,
	{_State159, OptionalLabelDeclType}:             _GotoState90Action,
	{_State159, SequenceExprType}:                  _GotoState105Action,
	{_State159, CallExprType}:                      _GotoState70Action,
	{_State159, AtomExprType}:                      _GotoState62Action,
	{_State159, ParseErrorExprType}:                _GotoState92Action,
	{_State159, LiteralExprType}:                   _GotoState87Action,
	{_State159, IdentifierExprType}:                _GotoState79Action,
	{_State159, BlockExprType}:                     _GotoState69Action,
	{_State159, InitializeExprType}:                _GotoState84Action,
	{_State159, ImplicitStructExprType}:            _GotoState80Action,
	{_State159, AccessibleExprType}:                _GotoState104Action,
	{_State159, AccessExprType}:                    _GotoState54Action,
	{_State159, IndexExprType}:                     _GotoState82Action,
	{_State159, PostfixableExprType}:               _GotoState94Action,
	{_State159, PostfixUnaryExprType}:              _GotoState93Action,
	{_State159, PrefixableExprType}:                _GotoState97Action,
	{_State159, PrefixUnaryExprType}:               _GotoState95Action,
	{_State159, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State159, MulExprType}:                       _GotoState89Action,
	{_State159, BinaryMulExprType}:                 _GotoState66Action,
	{_State159, AddExprType}:                       _GotoState56Action,
	{_State159, BinaryAddExprType}:                 _GotoState63Action,
	{_State159, CmpExprType}:                       _GotoState73Action,
	{_State159, BinaryCmpExprType}:                 _GotoState65Action,
	{_State159, AndExprType}:                       _GotoState57Action,
	{_State159, BinaryAndExprType}:                 _GotoState64Action,
	{_State159, OrExprType}:                        _GotoState91Action,
	{_State159, BinaryOrExprType}:                  _GotoState68Action,
	{_State159, InitializableTypeExprType}:         _GotoState83Action,
	{_State159, SliceTypeExprType}:                 _GotoState100Action,
	{_State159, ArrayTypeExprType}:                 _GotoState59Action,
	{_State159, MapTypeExprType}:                   _GotoState88Action,
	{_State159, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State159, AnonymousFuncExprType}:             _GotoState58Action,
	{_State161, AssignToken}:                       _GotoState280Action,
	{_State163, RparenToken}:                       _GotoState281Action,
	{_State164, ColonToken}:                        _GotoState282Action,
	{_State165, ColonToken}:                        _GotoState283Action,
	{_State165, EllipsisToken}:                     _GotoState284Action,
	{_State166, CommaToken}:                        _GotoState285Action,
	{_State167, IdentifierToken}:                   _GotoState230Action,
	{_State167, UnsafeToken}:                       _GotoState53Action,
	{_State167, StructToken}:                       _GotoState50Action,
	{_State167, EnumToken}:                         _GotoState109Action,
	{_State167, TraitToken}:                        _GotoState117Action,
	{_State167, FuncToken}:                         _GotoState111Action,
	{_State167, LparenToken}:                       _GotoState113Action,
	{_State167, LbracketToken}:                     _GotoState42Action,
	{_State167, DotToken}:                          _GotoState108Action,
	{_State167, QuestionToken}:                     _GotoState115Action,
	{_State167, ExclaimToken}:                      _GotoState110Action,
	{_State167, TildeTildeToken}:                   _GotoState116Action,
	{_State167, BitNegToken}:                       _GotoState107Action,
	{_State167, BitAndToken}:                       _GotoState106Action,
	{_State167, ParseErrorToken}:                   _GotoState114Action,
	{_State167, UnsafeStatementType}:               _GotoState237Action,
	{_State167, InitializableTypeExprType}:         _GotoState125Action,
	{_State167, SliceTypeExprType}:                 _GotoState100Action,
	{_State167, ArrayTypeExprType}:                 _GotoState59Action,
	{_State167, MapTypeExprType}:                   _GotoState88Action,
	{_State167, AtomTypeExprType}:                  _GotoState118Action,
	{_State167, NamedTypeExprType}:                 _GotoState126Action,
	{_State167, InferredTypeExprType}:              _GotoState124Action,
	{_State167, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State167, ReturnableTypeExprType}:            _GotoState130Action,
	{_State167, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State167, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State167, TypeExprType}:                      _GotoState236Action,
	{_State167, BinaryTypeExprType}:                _GotoState119Action,
	{_State167, FieldDefType}:                      _GotoState287Action,
	{_State167, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State167, ProperExplicitFieldDefsType}:       _GotoState288Action,
	{_State167, ExplicitFieldDefsType}:             _GotoState286Action,
	{_State167, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State167, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State167, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State167, TraitTypeExprType}:                 _GotoState131Action,
	{_State167, FuncTypeExprType}:                  _GotoState121Action,
	{_State168, IdentifierToken}:                   _GotoState289Action,
	{_State178, IdentifierToken}:                   _GotoState112Action,
	{_State178, StructToken}:                       _GotoState50Action,
	{_State178, EnumToken}:                         _GotoState109Action,
	{_State178, TraitToken}:                        _GotoState117Action,
	{_State178, FuncToken}:                         _GotoState111Action,
	{_State178, LparenToken}:                       _GotoState113Action,
	{_State178, LbracketToken}:                     _GotoState42Action,
	{_State178, DotToken}:                          _GotoState108Action,
	{_State178, QuestionToken}:                     _GotoState115Action,
	{_State178, ExclaimToken}:                      _GotoState110Action,
	{_State178, TildeTildeToken}:                   _GotoState116Action,
	{_State178, BitNegToken}:                       _GotoState107Action,
	{_State178, BitAndToken}:                       _GotoState106Action,
	{_State178, ParseErrorToken}:                   _GotoState114Action,
	{_State178, ProperTypeArgumentsType}:           _GotoState290Action,
	{_State178, TypeArgumentsType}:                 _GotoState291Action,
	{_State178, InitializableTypeExprType}:         _GotoState125Action,
	{_State178, SliceTypeExprType}:                 _GotoState100Action,
	{_State178, ArrayTypeExprType}:                 _GotoState59Action,
	{_State178, MapTypeExprType}:                   _GotoState88Action,
	{_State178, AtomTypeExprType}:                  _GotoState118Action,
	{_State178, NamedTypeExprType}:                 _GotoState126Action,
	{_State178, InferredTypeExprType}:              _GotoState124Action,
	{_State178, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State178, ReturnableTypeExprType}:            _GotoState130Action,
	{_State178, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State178, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State178, TypeExprType}:                      _GotoState292Action,
	{_State178, BinaryTypeExprType}:                _GotoState119Action,
	{_State178, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State178, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State178, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State178, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State178, TraitTypeExprType}:                 _GotoState131Action,
	{_State178, FuncTypeExprType}:                  _GotoState121Action,
	{_State179, IdentifierToken}:                   _GotoState293Action,
	{_State180, IntegerLiteralToken}:               _GotoState40Action,
	{_State180, FloatLiteralToken}:                 _GotoState35Action,
	{_State180, RuneLiteralToken}:                  _GotoState48Action,
	{_State180, StringLiteralToken}:                _GotoState49Action,
	{_State180, IdentifierToken}:                   _GotoState161Action,
	{_State180, TrueToken}:                         _GotoState52Action,
	{_State180, FalseToken}:                        _GotoState34Action,
	{_State180, StructToken}:                       _GotoState50Action,
	{_State180, FuncToken}:                         _GotoState36Action,
	{_State180, VarToken}:                          _GotoState15Action,
	{_State180, LetToken}:                          _GotoState12Action,
	{_State180, NotToken}:                          _GotoState45Action,
	{_State180, LabelDeclToken}:                    _GotoState41Action,
	{_State180, LparenToken}:                       _GotoState43Action,
	{_State180, LbracketToken}:                     _GotoState42Action,
	{_State180, ColonToken}:                        _GotoState159Action,
	{_State180, EllipsisToken}:                     _GotoState160Action,
	{_State180, SubToken}:                          _GotoState51Action,
	{_State180, MulToken}:                          _GotoState44Action,
	{_State180, BitNegToken}:                       _GotoState27Action,
	{_State180, BitAndToken}:                       _GotoState26Action,
	{_State180, GreaterToken}:                      _GotoState37Action,
	{_State180, ParseErrorToken}:                   _GotoState46Action,
	{_State180, VarDeclPatternType}:                _GotoState103Action,
	{_State180, VarOrLetType}:                      _GotoState24Action,
	{_State180, ExprType}:                          _GotoState165Action,
	{_State180, OptionalLabelDeclType}:             _GotoState90Action,
	{_State180, SequenceExprType}:                  _GotoState105Action,
	{_State180, CallExprType}:                      _GotoState70Action,
	{_State180, ArgumentType}:                      _GotoState294Action,
	{_State180, ColonExprType}:                     _GotoState164Action,
	{_State180, AtomExprType}:                      _GotoState62Action,
	{_State180, ParseErrorExprType}:                _GotoState92Action,
	{_State180, LiteralExprType}:                   _GotoState87Action,
	{_State180, IdentifierExprType}:                _GotoState79Action,
	{_State180, BlockExprType}:                     _GotoState69Action,
	{_State180, InitializeExprType}:                _GotoState84Action,
	{_State180, ImplicitStructExprType}:            _GotoState80Action,
	{_State180, AccessibleExprType}:                _GotoState104Action,
	{_State180, AccessExprType}:                    _GotoState54Action,
	{_State180, IndexExprType}:                     _GotoState82Action,
	{_State180, PostfixableExprType}:               _GotoState94Action,
	{_State180, PostfixUnaryExprType}:              _GotoState93Action,
	{_State180, PrefixableExprType}:                _GotoState97Action,
	{_State180, PrefixUnaryExprType}:               _GotoState95Action,
	{_State180, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State180, MulExprType}:                       _GotoState89Action,
	{_State180, BinaryMulExprType}:                 _GotoState66Action,
	{_State180, AddExprType}:                       _GotoState56Action,
	{_State180, BinaryAddExprType}:                 _GotoState63Action,
	{_State180, CmpExprType}:                       _GotoState73Action,
	{_State180, BinaryCmpExprType}:                 _GotoState65Action,
	{_State180, AndExprType}:                       _GotoState57Action,
	{_State180, BinaryAndExprType}:                 _GotoState64Action,
	{_State180, OrExprType}:                        _GotoState91Action,
	{_State180, BinaryOrExprType}:                  _GotoState68Action,
	{_State180, InitializableTypeExprType}:         _GotoState83Action,
	{_State180, SliceTypeExprType}:                 _GotoState100Action,
	{_State180, ArrayTypeExprType}:                 _GotoState59Action,
	{_State180, MapTypeExprType}:                   _GotoState88Action,
	{_State180, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State180, AnonymousFuncExprType}:             _GotoState58Action,
	{_State186, IntegerLiteralToken}:               _GotoState40Action,
	{_State186, FloatLiteralToken}:                 _GotoState35Action,
	{_State186, RuneLiteralToken}:                  _GotoState48Action,
	{_State186, StringLiteralToken}:                _GotoState49Action,
	{_State186, IdentifierToken}:                   _GotoState38Action,
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
	{_State186, SubToken}:                          _GotoState51Action,
	{_State186, MulToken}:                          _GotoState44Action,
	{_State186, BitNegToken}:                       _GotoState27Action,
	{_State186, BitAndToken}:                       _GotoState26Action,
	{_State186, GreaterToken}:                      _GotoState37Action,
	{_State186, ParseErrorToken}:                   _GotoState46Action,
	{_State186, VarDeclPatternType}:                _GotoState103Action,
	{_State186, VarOrLetType}:                      _GotoState24Action,
	{_State186, ExprType}:                          _GotoState295Action,
	{_State186, OptionalLabelDeclType}:             _GotoState90Action,
	{_State186, SequenceExprType}:                  _GotoState105Action,
	{_State186, CallExprType}:                      _GotoState70Action,
	{_State186, AtomExprType}:                      _GotoState62Action,
	{_State186, ParseErrorExprType}:                _GotoState92Action,
	{_State186, LiteralExprType}:                   _GotoState87Action,
	{_State186, IdentifierExprType}:                _GotoState79Action,
	{_State186, BlockExprType}:                     _GotoState69Action,
	{_State186, InitializeExprType}:                _GotoState84Action,
	{_State186, ImplicitStructExprType}:            _GotoState80Action,
	{_State186, AccessibleExprType}:                _GotoState104Action,
	{_State186, AccessExprType}:                    _GotoState54Action,
	{_State186, IndexExprType}:                     _GotoState82Action,
	{_State186, PostfixableExprType}:               _GotoState94Action,
	{_State186, PostfixUnaryExprType}:              _GotoState93Action,
	{_State186, PrefixableExprType}:                _GotoState97Action,
	{_State186, PrefixUnaryExprType}:               _GotoState95Action,
	{_State186, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State186, MulExprType}:                       _GotoState89Action,
	{_State186, BinaryMulExprType}:                 _GotoState66Action,
	{_State186, AddExprType}:                       _GotoState56Action,
	{_State186, BinaryAddExprType}:                 _GotoState63Action,
	{_State186, CmpExprType}:                       _GotoState73Action,
	{_State186, BinaryCmpExprType}:                 _GotoState65Action,
	{_State186, AndExprType}:                       _GotoState57Action,
	{_State186, BinaryAndExprType}:                 _GotoState64Action,
	{_State186, OrExprType}:                        _GotoState91Action,
	{_State186, BinaryOrExprType}:                  _GotoState68Action,
	{_State186, InitializableTypeExprType}:         _GotoState83Action,
	{_State186, SliceTypeExprType}:                 _GotoState100Action,
	{_State186, ArrayTypeExprType}:                 _GotoState59Action,
	{_State186, MapTypeExprType}:                   _GotoState88Action,
	{_State186, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State186, AnonymousFuncExprType}:             _GotoState58Action,
	{_State187, LparenToken}:                       _GotoState296Action,
	{_State193, IntegerLiteralToken}:               _GotoState40Action,
	{_State193, FloatLiteralToken}:                 _GotoState35Action,
	{_State193, RuneLiteralToken}:                  _GotoState48Action,
	{_State193, StringLiteralToken}:                _GotoState49Action,
	{_State193, IdentifierToken}:                   _GotoState38Action,
	{_State193, TrueToken}:                         _GotoState52Action,
	{_State193, FalseToken}:                        _GotoState34Action,
	{_State193, StructToken}:                       _GotoState50Action,
	{_State193, FuncToken}:                         _GotoState36Action,
	{_State193, NotToken}:                          _GotoState45Action,
	{_State193, LabelDeclToken}:                    _GotoState41Action,
	{_State193, LparenToken}:                       _GotoState43Action,
	{_State193, LbracketToken}:                     _GotoState42Action,
	{_State193, SubToken}:                          _GotoState51Action,
	{_State193, MulToken}:                          _GotoState44Action,
	{_State193, BitNegToken}:                       _GotoState27Action,
	{_State193, BitAndToken}:                       _GotoState26Action,
	{_State193, ParseErrorToken}:                   _GotoState46Action,
	{_State193, OptionalLabelDeclType}:             _GotoState150Action,
	{_State193, CallExprType}:                      _GotoState70Action,
	{_State193, AtomExprType}:                      _GotoState62Action,
	{_State193, ParseErrorExprType}:                _GotoState92Action,
	{_State193, LiteralExprType}:                   _GotoState87Action,
	{_State193, IdentifierExprType}:                _GotoState79Action,
	{_State193, BlockExprType}:                     _GotoState69Action,
	{_State193, InitializeExprType}:                _GotoState84Action,
	{_State193, ImplicitStructExprType}:            _GotoState80Action,
	{_State193, AccessibleExprType}:                _GotoState104Action,
	{_State193, AccessExprType}:                    _GotoState54Action,
	{_State193, IndexExprType}:                     _GotoState82Action,
	{_State193, PostfixableExprType}:               _GotoState94Action,
	{_State193, PostfixUnaryExprType}:              _GotoState93Action,
	{_State193, PrefixableExprType}:                _GotoState97Action,
	{_State193, PrefixUnaryExprType}:               _GotoState95Action,
	{_State193, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State193, MulExprType}:                       _GotoState297Action,
	{_State193, BinaryMulExprType}:                 _GotoState66Action,
	{_State193, InitializableTypeExprType}:         _GotoState83Action,
	{_State193, SliceTypeExprType}:                 _GotoState100Action,
	{_State193, ArrayTypeExprType}:                 _GotoState59Action,
	{_State193, MapTypeExprType}:                   _GotoState88Action,
	{_State193, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State193, AnonymousFuncExprType}:             _GotoState58Action,
	{_State194, IntegerLiteralToken}:               _GotoState40Action,
	{_State194, FloatLiteralToken}:                 _GotoState35Action,
	{_State194, RuneLiteralToken}:                  _GotoState48Action,
	{_State194, StringLiteralToken}:                _GotoState49Action,
	{_State194, IdentifierToken}:                   _GotoState38Action,
	{_State194, TrueToken}:                         _GotoState52Action,
	{_State194, FalseToken}:                        _GotoState34Action,
	{_State194, StructToken}:                       _GotoState50Action,
	{_State194, FuncToken}:                         _GotoState36Action,
	{_State194, NotToken}:                          _GotoState45Action,
	{_State194, LabelDeclToken}:                    _GotoState41Action,
	{_State194, LparenToken}:                       _GotoState43Action,
	{_State194, LbracketToken}:                     _GotoState42Action,
	{_State194, SubToken}:                          _GotoState51Action,
	{_State194, MulToken}:                          _GotoState44Action,
	{_State194, BitNegToken}:                       _GotoState27Action,
	{_State194, BitAndToken}:                       _GotoState26Action,
	{_State194, ParseErrorToken}:                   _GotoState46Action,
	{_State194, OptionalLabelDeclType}:             _GotoState150Action,
	{_State194, CallExprType}:                      _GotoState70Action,
	{_State194, AtomExprType}:                      _GotoState62Action,
	{_State194, ParseErrorExprType}:                _GotoState92Action,
	{_State194, LiteralExprType}:                   _GotoState87Action,
	{_State194, IdentifierExprType}:                _GotoState79Action,
	{_State194, BlockExprType}:                     _GotoState69Action,
	{_State194, InitializeExprType}:                _GotoState84Action,
	{_State194, ImplicitStructExprType}:            _GotoState80Action,
	{_State194, AccessibleExprType}:                _GotoState104Action,
	{_State194, AccessExprType}:                    _GotoState54Action,
	{_State194, IndexExprType}:                     _GotoState82Action,
	{_State194, PostfixableExprType}:               _GotoState94Action,
	{_State194, PostfixUnaryExprType}:              _GotoState93Action,
	{_State194, PrefixableExprType}:                _GotoState97Action,
	{_State194, PrefixUnaryExprType}:               _GotoState95Action,
	{_State194, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State194, MulExprType}:                       _GotoState89Action,
	{_State194, BinaryMulExprType}:                 _GotoState66Action,
	{_State194, AddExprType}:                       _GotoState56Action,
	{_State194, BinaryAddExprType}:                 _GotoState63Action,
	{_State194, CmpExprType}:                       _GotoState298Action,
	{_State194, BinaryCmpExprType}:                 _GotoState65Action,
	{_State194, InitializableTypeExprType}:         _GotoState83Action,
	{_State194, SliceTypeExprType}:                 _GotoState100Action,
	{_State194, ArrayTypeExprType}:                 _GotoState59Action,
	{_State194, MapTypeExprType}:                   _GotoState88Action,
	{_State194, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State194, AnonymousFuncExprType}:             _GotoState58Action,
	{_State195, IntegerLiteralToken}:               _GotoState40Action,
	{_State195, FloatLiteralToken}:                 _GotoState35Action,
	{_State195, RuneLiteralToken}:                  _GotoState48Action,
	{_State195, StringLiteralToken}:                _GotoState49Action,
	{_State195, IdentifierToken}:                   _GotoState38Action,
	{_State195, TrueToken}:                         _GotoState52Action,
	{_State195, FalseToken}:                        _GotoState34Action,
	{_State195, StructToken}:                       _GotoState50Action,
	{_State195, FuncToken}:                         _GotoState36Action,
	{_State195, VarToken}:                          _GotoState15Action,
	{_State195, LetToken}:                          _GotoState12Action,
	{_State195, NotToken}:                          _GotoState45Action,
	{_State195, LabelDeclToken}:                    _GotoState41Action,
	{_State195, LparenToken}:                       _GotoState43Action,
	{_State195, LbracketToken}:                     _GotoState42Action,
	{_State195, SubToken}:                          _GotoState51Action,
	{_State195, MulToken}:                          _GotoState44Action,
	{_State195, BitNegToken}:                       _GotoState27Action,
	{_State195, BitAndToken}:                       _GotoState26Action,
	{_State195, GreaterToken}:                      _GotoState37Action,
	{_State195, ParseErrorToken}:                   _GotoState46Action,
	{_State195, VarDeclPatternType}:                _GotoState103Action,
	{_State195, VarOrLetType}:                      _GotoState24Action,
	{_State195, ExprType}:                          _GotoState299Action,
	{_State195, OptionalLabelDeclType}:             _GotoState90Action,
	{_State195, SequenceExprType}:                  _GotoState105Action,
	{_State195, CallExprType}:                      _GotoState70Action,
	{_State195, AtomExprType}:                      _GotoState62Action,
	{_State195, ParseErrorExprType}:                _GotoState92Action,
	{_State195, LiteralExprType}:                   _GotoState87Action,
	{_State195, IdentifierExprType}:                _GotoState79Action,
	{_State195, BlockExprType}:                     _GotoState69Action,
	{_State195, InitializeExprType}:                _GotoState84Action,
	{_State195, ImplicitStructExprType}:            _GotoState80Action,
	{_State195, AccessibleExprType}:                _GotoState104Action,
	{_State195, AccessExprType}:                    _GotoState54Action,
	{_State195, IndexExprType}:                     _GotoState82Action,
	{_State195, PostfixableExprType}:               _GotoState94Action,
	{_State195, PostfixUnaryExprType}:              _GotoState93Action,
	{_State195, PrefixableExprType}:                _GotoState97Action,
	{_State195, PrefixUnaryExprType}:               _GotoState95Action,
	{_State195, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State195, MulExprType}:                       _GotoState89Action,
	{_State195, BinaryMulExprType}:                 _GotoState66Action,
	{_State195, AddExprType}:                       _GotoState56Action,
	{_State195, BinaryAddExprType}:                 _GotoState63Action,
	{_State195, CmpExprType}:                       _GotoState73Action,
	{_State195, BinaryCmpExprType}:                 _GotoState65Action,
	{_State195, AndExprType}:                       _GotoState57Action,
	{_State195, BinaryAndExprType}:                 _GotoState64Action,
	{_State195, OrExprType}:                        _GotoState91Action,
	{_State195, BinaryOrExprType}:                  _GotoState68Action,
	{_State195, InitializableTypeExprType}:         _GotoState83Action,
	{_State195, SliceTypeExprType}:                 _GotoState100Action,
	{_State195, ArrayTypeExprType}:                 _GotoState59Action,
	{_State195, MapTypeExprType}:                   _GotoState88Action,
	{_State195, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State195, AnonymousFuncExprType}:             _GotoState58Action,
	{_State196, LbracketToken}:                     _GotoState180Action,
	{_State196, DotToken}:                          _GotoState179Action,
	{_State196, DollarLbracketToken}:               _GotoState178Action,
	{_State196, GenericTypeArgumentsType}:          _GotoState187Action,
	{_State204, IntegerLiteralToken}:               _GotoState40Action,
	{_State204, FloatLiteralToken}:                 _GotoState35Action,
	{_State204, RuneLiteralToken}:                  _GotoState48Action,
	{_State204, StringLiteralToken}:                _GotoState49Action,
	{_State204, IdentifierToken}:                   _GotoState38Action,
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
	{_State204, OptionalLabelDeclType}:             _GotoState150Action,
	{_State204, CallExprType}:                      _GotoState70Action,
	{_State204, AtomExprType}:                      _GotoState62Action,
	{_State204, ParseErrorExprType}:                _GotoState92Action,
	{_State204, LiteralExprType}:                   _GotoState87Action,
	{_State204, IdentifierExprType}:                _GotoState79Action,
	{_State204, BlockExprType}:                     _GotoState69Action,
	{_State204, InitializeExprType}:                _GotoState84Action,
	{_State204, ImplicitStructExprType}:            _GotoState80Action,
	{_State204, AccessibleExprType}:                _GotoState104Action,
	{_State204, AccessExprType}:                    _GotoState54Action,
	{_State204, IndexExprType}:                     _GotoState82Action,
	{_State204, PostfixableExprType}:               _GotoState94Action,
	{_State204, PostfixUnaryExprType}:              _GotoState93Action,
	{_State204, PrefixableExprType}:                _GotoState97Action,
	{_State204, PrefixUnaryExprType}:               _GotoState95Action,
	{_State204, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State204, MulExprType}:                       _GotoState89Action,
	{_State204, BinaryMulExprType}:                 _GotoState66Action,
	{_State204, AddExprType}:                       _GotoState300Action,
	{_State204, BinaryAddExprType}:                 _GotoState63Action,
	{_State204, InitializableTypeExprType}:         _GotoState83Action,
	{_State204, SliceTypeExprType}:                 _GotoState100Action,
	{_State204, ArrayTypeExprType}:                 _GotoState59Action,
	{_State204, MapTypeExprType}:                   _GotoState88Action,
	{_State204, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State204, AnonymousFuncExprType}:             _GotoState58Action,
	{_State205, IntegerLiteralToken}:               _GotoState40Action,
	{_State205, FloatLiteralToken}:                 _GotoState35Action,
	{_State205, RuneLiteralToken}:                  _GotoState48Action,
	{_State205, StringLiteralToken}:                _GotoState49Action,
	{_State205, IdentifierToken}:                   _GotoState38Action,
	{_State205, TrueToken}:                         _GotoState52Action,
	{_State205, FalseToken}:                        _GotoState34Action,
	{_State205, StructToken}:                       _GotoState50Action,
	{_State205, FuncToken}:                         _GotoState36Action,
	{_State205, VarToken}:                          _GotoState15Action,
	{_State205, LetToken}:                          _GotoState12Action,
	{_State205, NotToken}:                          _GotoState45Action,
	{_State205, LabelDeclToken}:                    _GotoState41Action,
	{_State205, LparenToken}:                       _GotoState43Action,
	{_State205, LbracketToken}:                     _GotoState42Action,
	{_State205, SubToken}:                          _GotoState51Action,
	{_State205, MulToken}:                          _GotoState44Action,
	{_State205, BitNegToken}:                       _GotoState27Action,
	{_State205, BitAndToken}:                       _GotoState26Action,
	{_State205, GreaterToken}:                      _GotoState37Action,
	{_State205, ParseErrorToken}:                   _GotoState46Action,
	{_State205, VarDeclPatternType}:                _GotoState103Action,
	{_State205, VarOrLetType}:                      _GotoState24Action,
	{_State205, ExprType}:                          _GotoState301Action,
	{_State205, OptionalLabelDeclType}:             _GotoState90Action,
	{_State205, SequenceExprType}:                  _GotoState105Action,
	{_State205, CallExprType}:                      _GotoState70Action,
	{_State205, AtomExprType}:                      _GotoState62Action,
	{_State205, ParseErrorExprType}:                _GotoState92Action,
	{_State205, LiteralExprType}:                   _GotoState87Action,
	{_State205, IdentifierExprType}:                _GotoState79Action,
	{_State205, BlockExprType}:                     _GotoState69Action,
	{_State205, InitializeExprType}:                _GotoState84Action,
	{_State205, ImplicitStructExprType}:            _GotoState80Action,
	{_State205, AccessibleExprType}:                _GotoState104Action,
	{_State205, AccessExprType}:                    _GotoState54Action,
	{_State205, IndexExprType}:                     _GotoState82Action,
	{_State205, PostfixableExprType}:               _GotoState94Action,
	{_State205, PostfixUnaryExprType}:              _GotoState93Action,
	{_State205, PrefixableExprType}:                _GotoState97Action,
	{_State205, PrefixUnaryExprType}:               _GotoState95Action,
	{_State205, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State205, MulExprType}:                       _GotoState89Action,
	{_State205, BinaryMulExprType}:                 _GotoState66Action,
	{_State205, AddExprType}:                       _GotoState56Action,
	{_State205, BinaryAddExprType}:                 _GotoState63Action,
	{_State205, CmpExprType}:                       _GotoState73Action,
	{_State205, BinaryCmpExprType}:                 _GotoState65Action,
	{_State205, AndExprType}:                       _GotoState57Action,
	{_State205, BinaryAndExprType}:                 _GotoState64Action,
	{_State205, OrExprType}:                        _GotoState91Action,
	{_State205, BinaryOrExprType}:                  _GotoState68Action,
	{_State205, InitializableTypeExprType}:         _GotoState83Action,
	{_State205, SliceTypeExprType}:                 _GotoState100Action,
	{_State205, ArrayTypeExprType}:                 _GotoState59Action,
	{_State205, MapTypeExprType}:                   _GotoState88Action,
	{_State205, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State205, AnonymousFuncExprType}:             _GotoState58Action,
	{_State206, IntegerLiteralToken}:               _GotoState40Action,
	{_State206, FloatLiteralToken}:                 _GotoState35Action,
	{_State206, RuneLiteralToken}:                  _GotoState48Action,
	{_State206, StringLiteralToken}:                _GotoState49Action,
	{_State206, IdentifierToken}:                   _GotoState161Action,
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
	{_State206, ColonToken}:                        _GotoState159Action,
	{_State206, EllipsisToken}:                     _GotoState160Action,
	{_State206, SubToken}:                          _GotoState51Action,
	{_State206, MulToken}:                          _GotoState44Action,
	{_State206, BitNegToken}:                       _GotoState27Action,
	{_State206, BitAndToken}:                       _GotoState26Action,
	{_State206, GreaterToken}:                      _GotoState37Action,
	{_State206, ParseErrorToken}:                   _GotoState46Action,
	{_State206, VarDeclPatternType}:                _GotoState103Action,
	{_State206, VarOrLetType}:                      _GotoState24Action,
	{_State206, ExprType}:                          _GotoState165Action,
	{_State206, OptionalLabelDeclType}:             _GotoState90Action,
	{_State206, SequenceExprType}:                  _GotoState105Action,
	{_State206, CallExprType}:                      _GotoState70Action,
	{_State206, ProperArgumentsType}:               _GotoState166Action,
	{_State206, ArgumentsType}:                     _GotoState302Action,
	{_State206, ArgumentType}:                      _GotoState162Action,
	{_State206, ColonExprType}:                     _GotoState164Action,
	{_State206, AtomExprType}:                      _GotoState62Action,
	{_State206, ParseErrorExprType}:                _GotoState92Action,
	{_State206, LiteralExprType}:                   _GotoState87Action,
	{_State206, IdentifierExprType}:                _GotoState79Action,
	{_State206, BlockExprType}:                     _GotoState69Action,
	{_State206, InitializeExprType}:                _GotoState84Action,
	{_State206, ImplicitStructExprType}:            _GotoState80Action,
	{_State206, AccessibleExprType}:                _GotoState104Action,
	{_State206, AccessExprType}:                    _GotoState54Action,
	{_State206, IndexExprType}:                     _GotoState82Action,
	{_State206, PostfixableExprType}:               _GotoState94Action,
	{_State206, PostfixUnaryExprType}:              _GotoState93Action,
	{_State206, PrefixableExprType}:                _GotoState97Action,
	{_State206, PrefixUnaryExprType}:               _GotoState95Action,
	{_State206, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State206, MulExprType}:                       _GotoState89Action,
	{_State206, BinaryMulExprType}:                 _GotoState66Action,
	{_State206, AddExprType}:                       _GotoState56Action,
	{_State206, BinaryAddExprType}:                 _GotoState63Action,
	{_State206, CmpExprType}:                       _GotoState73Action,
	{_State206, BinaryCmpExprType}:                 _GotoState65Action,
	{_State206, AndExprType}:                       _GotoState57Action,
	{_State206, BinaryAndExprType}:                 _GotoState64Action,
	{_State206, OrExprType}:                        _GotoState91Action,
	{_State206, BinaryOrExprType}:                  _GotoState68Action,
	{_State206, InitializableTypeExprType}:         _GotoState83Action,
	{_State206, SliceTypeExprType}:                 _GotoState100Action,
	{_State206, ArrayTypeExprType}:                 _GotoState59Action,
	{_State206, MapTypeExprType}:                   _GotoState88Action,
	{_State206, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State206, AnonymousFuncExprType}:             _GotoState58Action,
	{_State207, IntegerLiteralToken}:               _GotoState40Action,
	{_State207, FloatLiteralToken}:                 _GotoState35Action,
	{_State207, RuneLiteralToken}:                  _GotoState48Action,
	{_State207, StringLiteralToken}:                _GotoState49Action,
	{_State207, IdentifierToken}:                   _GotoState38Action,
	{_State207, TrueToken}:                         _GotoState52Action,
	{_State207, FalseToken}:                        _GotoState34Action,
	{_State207, StructToken}:                       _GotoState50Action,
	{_State207, FuncToken}:                         _GotoState36Action,
	{_State207, VarToken}:                          _GotoState15Action,
	{_State207, LetToken}:                          _GotoState12Action,
	{_State207, NotToken}:                          _GotoState45Action,
	{_State207, LabelDeclToken}:                    _GotoState41Action,
	{_State207, LparenToken}:                       _GotoState43Action,
	{_State207, LbracketToken}:                     _GotoState42Action,
	{_State207, SubToken}:                          _GotoState51Action,
	{_State207, MulToken}:                          _GotoState44Action,
	{_State207, BitNegToken}:                       _GotoState27Action,
	{_State207, BitAndToken}:                       _GotoState26Action,
	{_State207, GreaterToken}:                      _GotoState37Action,
	{_State207, ParseErrorToken}:                   _GotoState46Action,
	{_State207, ExprsType}:                         _GotoState303Action,
	{_State207, VarDeclPatternType}:                _GotoState103Action,
	{_State207, VarOrLetType}:                      _GotoState24Action,
	{_State207, ExprType}:                          _GotoState75Action,
	{_State207, OptionalLabelDeclType}:             _GotoState90Action,
	{_State207, SequenceExprType}:                  _GotoState105Action,
	{_State207, CallExprType}:                      _GotoState70Action,
	{_State207, AtomExprType}:                      _GotoState62Action,
	{_State207, ParseErrorExprType}:                _GotoState92Action,
	{_State207, LiteralExprType}:                   _GotoState87Action,
	{_State207, IdentifierExprType}:                _GotoState79Action,
	{_State207, BlockExprType}:                     _GotoState69Action,
	{_State207, InitializeExprType}:                _GotoState84Action,
	{_State207, ImplicitStructExprType}:            _GotoState80Action,
	{_State207, AccessibleExprType}:                _GotoState104Action,
	{_State207, AccessExprType}:                    _GotoState54Action,
	{_State207, IndexExprType}:                     _GotoState82Action,
	{_State207, PostfixableExprType}:               _GotoState94Action,
	{_State207, PostfixUnaryExprType}:              _GotoState93Action,
	{_State207, PrefixableExprType}:                _GotoState97Action,
	{_State207, PrefixUnaryExprType}:               _GotoState95Action,
	{_State207, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State207, MulExprType}:                       _GotoState89Action,
	{_State207, BinaryMulExprType}:                 _GotoState66Action,
	{_State207, AddExprType}:                       _GotoState56Action,
	{_State207, BinaryAddExprType}:                 _GotoState63Action,
	{_State207, CmpExprType}:                       _GotoState73Action,
	{_State207, BinaryCmpExprType}:                 _GotoState65Action,
	{_State207, AndExprType}:                       _GotoState57Action,
	{_State207, BinaryAndExprType}:                 _GotoState64Action,
	{_State207, OrExprType}:                        _GotoState91Action,
	{_State207, BinaryOrExprType}:                  _GotoState68Action,
	{_State207, InitializableTypeExprType}:         _GotoState83Action,
	{_State207, SliceTypeExprType}:                 _GotoState100Action,
	{_State207, ArrayTypeExprType}:                 _GotoState59Action,
	{_State207, MapTypeExprType}:                   _GotoState88Action,
	{_State207, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State207, AnonymousFuncExprType}:             _GotoState58Action,
	{_State208, CommaToken}:                        _GotoState205Action,
	{_State215, IntegerLiteralToken}:               _GotoState40Action,
	{_State215, FloatLiteralToken}:                 _GotoState35Action,
	{_State215, RuneLiteralToken}:                  _GotoState48Action,
	{_State215, StringLiteralToken}:                _GotoState49Action,
	{_State215, IdentifierToken}:                   _GotoState38Action,
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
	{_State215, OptionalLabelDeclType}:             _GotoState150Action,
	{_State215, CallExprType}:                      _GotoState70Action,
	{_State215, AtomExprType}:                      _GotoState62Action,
	{_State215, ParseErrorExprType}:                _GotoState92Action,
	{_State215, LiteralExprType}:                   _GotoState87Action,
	{_State215, IdentifierExprType}:                _GotoState79Action,
	{_State215, BlockExprType}:                     _GotoState69Action,
	{_State215, InitializeExprType}:                _GotoState84Action,
	{_State215, ImplicitStructExprType}:            _GotoState80Action,
	{_State215, AccessibleExprType}:                _GotoState104Action,
	{_State215, AccessExprType}:                    _GotoState54Action,
	{_State215, IndexExprType}:                     _GotoState82Action,
	{_State215, PostfixableExprType}:               _GotoState94Action,
	{_State215, PostfixUnaryExprType}:              _GotoState93Action,
	{_State215, PrefixableExprType}:                _GotoState304Action,
	{_State215, PrefixUnaryExprType}:               _GotoState95Action,
	{_State215, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State215, InitializableTypeExprType}:         _GotoState83Action,
	{_State215, SliceTypeExprType}:                 _GotoState100Action,
	{_State215, ArrayTypeExprType}:                 _GotoState59Action,
	{_State215, MapTypeExprType}:                   _GotoState88Action,
	{_State215, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State215, AnonymousFuncExprType}:             _GotoState58Action,
	{_State216, LbraceToken}:                       _GotoState11Action,
	{_State216, StatementBlockType}:                _GotoState305Action,
	{_State217, IntegerLiteralToken}:               _GotoState40Action,
	{_State217, FloatLiteralToken}:                 _GotoState35Action,
	{_State217, RuneLiteralToken}:                  _GotoState48Action,
	{_State217, StringLiteralToken}:                _GotoState49Action,
	{_State217, IdentifierToken}:                   _GotoState38Action,
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
	{_State217, SubToken}:                          _GotoState51Action,
	{_State217, MulToken}:                          _GotoState44Action,
	{_State217, BitNegToken}:                       _GotoState27Action,
	{_State217, BitAndToken}:                       _GotoState26Action,
	{_State217, GreaterToken}:                      _GotoState37Action,
	{_State217, ParseErrorToken}:                   _GotoState46Action,
	{_State217, VarDeclPatternType}:                _GotoState103Action,
	{_State217, VarOrLetType}:                      _GotoState24Action,
	{_State217, AssignPatternType}:                 _GotoState306Action,
	{_State217, OptionalLabelDeclType}:             _GotoState150Action,
	{_State217, SequenceExprType}:                  _GotoState308Action,
	{_State217, ForAssignmentType}:                 _GotoState307Action,
	{_State217, CallExprType}:                      _GotoState70Action,
	{_State217, AtomExprType}:                      _GotoState62Action,
	{_State217, ParseErrorExprType}:                _GotoState92Action,
	{_State217, LiteralExprType}:                   _GotoState87Action,
	{_State217, IdentifierExprType}:                _GotoState79Action,
	{_State217, BlockExprType}:                     _GotoState69Action,
	{_State217, InitializeExprType}:                _GotoState84Action,
	{_State217, ImplicitStructExprType}:            _GotoState80Action,
	{_State217, AccessibleExprType}:                _GotoState104Action,
	{_State217, AccessExprType}:                    _GotoState54Action,
	{_State217, IndexExprType}:                     _GotoState82Action,
	{_State217, PostfixableExprType}:               _GotoState94Action,
	{_State217, PostfixUnaryExprType}:              _GotoState93Action,
	{_State217, PrefixableExprType}:                _GotoState97Action,
	{_State217, PrefixUnaryExprType}:               _GotoState95Action,
	{_State217, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State217, MulExprType}:                       _GotoState89Action,
	{_State217, BinaryMulExprType}:                 _GotoState66Action,
	{_State217, AddExprType}:                       _GotoState56Action,
	{_State217, BinaryAddExprType}:                 _GotoState63Action,
	{_State217, CmpExprType}:                       _GotoState73Action,
	{_State217, BinaryCmpExprType}:                 _GotoState65Action,
	{_State217, AndExprType}:                       _GotoState57Action,
	{_State217, BinaryAndExprType}:                 _GotoState64Action,
	{_State217, OrExprType}:                        _GotoState91Action,
	{_State217, BinaryOrExprType}:                  _GotoState68Action,
	{_State217, InitializableTypeExprType}:         _GotoState83Action,
	{_State217, SliceTypeExprType}:                 _GotoState100Action,
	{_State217, ArrayTypeExprType}:                 _GotoState59Action,
	{_State217, MapTypeExprType}:                   _GotoState88Action,
	{_State217, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State217, AnonymousFuncExprType}:             _GotoState58Action,
	{_State218, IntegerLiteralToken}:               _GotoState40Action,
	{_State218, FloatLiteralToken}:                 _GotoState35Action,
	{_State218, RuneLiteralToken}:                  _GotoState48Action,
	{_State218, StringLiteralToken}:                _GotoState49Action,
	{_State218, IdentifierToken}:                   _GotoState38Action,
	{_State218, TrueToken}:                         _GotoState52Action,
	{_State218, FalseToken}:                        _GotoState34Action,
	{_State218, CaseToken}:                         _GotoState309Action,
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
	{_State218, VarDeclPatternType}:                _GotoState103Action,
	{_State218, VarOrLetType}:                      _GotoState24Action,
	{_State218, OptionalLabelDeclType}:             _GotoState150Action,
	{_State218, SequenceExprType}:                  _GotoState311Action,
	{_State218, ConditionType}:                     _GotoState310Action,
	{_State218, CallExprType}:                      _GotoState70Action,
	{_State218, AtomExprType}:                      _GotoState62Action,
	{_State218, ParseErrorExprType}:                _GotoState92Action,
	{_State218, LiteralExprType}:                   _GotoState87Action,
	{_State218, IdentifierExprType}:                _GotoState79Action,
	{_State218, BlockExprType}:                     _GotoState69Action,
	{_State218, InitializeExprType}:                _GotoState84Action,
	{_State218, ImplicitStructExprType}:            _GotoState80Action,
	{_State218, AccessibleExprType}:                _GotoState104Action,
	{_State218, AccessExprType}:                    _GotoState54Action,
	{_State218, IndexExprType}:                     _GotoState82Action,
	{_State218, PostfixableExprType}:               _GotoState94Action,
	{_State218, PostfixUnaryExprType}:              _GotoState93Action,
	{_State218, PrefixableExprType}:                _GotoState97Action,
	{_State218, PrefixUnaryExprType}:               _GotoState95Action,
	{_State218, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State218, MulExprType}:                       _GotoState89Action,
	{_State218, BinaryMulExprType}:                 _GotoState66Action,
	{_State218, AddExprType}:                       _GotoState56Action,
	{_State218, BinaryAddExprType}:                 _GotoState63Action,
	{_State218, CmpExprType}:                       _GotoState73Action,
	{_State218, BinaryCmpExprType}:                 _GotoState65Action,
	{_State218, AndExprType}:                       _GotoState57Action,
	{_State218, BinaryAndExprType}:                 _GotoState64Action,
	{_State218, OrExprType}:                        _GotoState91Action,
	{_State218, BinaryOrExprType}:                  _GotoState68Action,
	{_State218, InitializableTypeExprType}:         _GotoState83Action,
	{_State218, SliceTypeExprType}:                 _GotoState100Action,
	{_State218, ArrayTypeExprType}:                 _GotoState59Action,
	{_State218, MapTypeExprType}:                   _GotoState88Action,
	{_State218, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State218, AnonymousFuncExprType}:             _GotoState58Action,
	{_State219, IntegerLiteralToken}:               _GotoState40Action,
	{_State219, FloatLiteralToken}:                 _GotoState35Action,
	{_State219, RuneLiteralToken}:                  _GotoState48Action,
	{_State219, StringLiteralToken}:                _GotoState49Action,
	{_State219, IdentifierToken}:                   _GotoState38Action,
	{_State219, TrueToken}:                         _GotoState52Action,
	{_State219, FalseToken}:                        _GotoState34Action,
	{_State219, StructToken}:                       _GotoState50Action,
	{_State219, FuncToken}:                         _GotoState36Action,
	{_State219, VarToken}:                          _GotoState15Action,
	{_State219, LetToken}:                          _GotoState12Action,
	{_State219, NotToken}:                          _GotoState45Action,
	{_State219, LabelDeclToken}:                    _GotoState41Action,
	{_State219, LparenToken}:                       _GotoState43Action,
	{_State219, LbracketToken}:                     _GotoState42Action,
	{_State219, SubToken}:                          _GotoState51Action,
	{_State219, MulToken}:                          _GotoState44Action,
	{_State219, BitNegToken}:                       _GotoState27Action,
	{_State219, BitAndToken}:                       _GotoState26Action,
	{_State219, GreaterToken}:                      _GotoState37Action,
	{_State219, ParseErrorToken}:                   _GotoState46Action,
	{_State219, VarDeclPatternType}:                _GotoState103Action,
	{_State219, VarOrLetType}:                      _GotoState24Action,
	{_State219, OptionalLabelDeclType}:             _GotoState150Action,
	{_State219, SequenceExprType}:                  _GotoState312Action,
	{_State219, CallExprType}:                      _GotoState70Action,
	{_State219, AtomExprType}:                      _GotoState62Action,
	{_State219, ParseErrorExprType}:                _GotoState92Action,
	{_State219, LiteralExprType}:                   _GotoState87Action,
	{_State219, IdentifierExprType}:                _GotoState79Action,
	{_State219, BlockExprType}:                     _GotoState69Action,
	{_State219, InitializeExprType}:                _GotoState84Action,
	{_State219, ImplicitStructExprType}:            _GotoState80Action,
	{_State219, AccessibleExprType}:                _GotoState104Action,
	{_State219, AccessExprType}:                    _GotoState54Action,
	{_State219, IndexExprType}:                     _GotoState82Action,
	{_State219, PostfixableExprType}:               _GotoState94Action,
	{_State219, PostfixUnaryExprType}:              _GotoState93Action,
	{_State219, PrefixableExprType}:                _GotoState97Action,
	{_State219, PrefixUnaryExprType}:               _GotoState95Action,
	{_State219, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State219, MulExprType}:                       _GotoState89Action,
	{_State219, BinaryMulExprType}:                 _GotoState66Action,
	{_State219, AddExprType}:                       _GotoState56Action,
	{_State219, BinaryAddExprType}:                 _GotoState63Action,
	{_State219, CmpExprType}:                       _GotoState73Action,
	{_State219, BinaryCmpExprType}:                 _GotoState65Action,
	{_State219, AndExprType}:                       _GotoState57Action,
	{_State219, BinaryAndExprType}:                 _GotoState64Action,
	{_State219, OrExprType}:                        _GotoState91Action,
	{_State219, BinaryOrExprType}:                  _GotoState68Action,
	{_State219, InitializableTypeExprType}:         _GotoState83Action,
	{_State219, SliceTypeExprType}:                 _GotoState100Action,
	{_State219, ArrayTypeExprType}:                 _GotoState59Action,
	{_State219, MapTypeExprType}:                   _GotoState88Action,
	{_State219, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State219, AnonymousFuncExprType}:             _GotoState58Action,
	{_State224, IntegerLiteralToken}:               _GotoState40Action,
	{_State224, FloatLiteralToken}:                 _GotoState35Action,
	{_State224, RuneLiteralToken}:                  _GotoState48Action,
	{_State224, StringLiteralToken}:                _GotoState49Action,
	{_State224, IdentifierToken}:                   _GotoState38Action,
	{_State224, TrueToken}:                         _GotoState52Action,
	{_State224, FalseToken}:                        _GotoState34Action,
	{_State224, StructToken}:                       _GotoState50Action,
	{_State224, FuncToken}:                         _GotoState36Action,
	{_State224, NotToken}:                          _GotoState45Action,
	{_State224, LabelDeclToken}:                    _GotoState41Action,
	{_State224, LparenToken}:                       _GotoState43Action,
	{_State224, LbracketToken}:                     _GotoState42Action,
	{_State224, SubToken}:                          _GotoState51Action,
	{_State224, MulToken}:                          _GotoState44Action,
	{_State224, BitNegToken}:                       _GotoState27Action,
	{_State224, BitAndToken}:                       _GotoState26Action,
	{_State224, ParseErrorToken}:                   _GotoState46Action,
	{_State224, OptionalLabelDeclType}:             _GotoState150Action,
	{_State224, CallExprType}:                      _GotoState70Action,
	{_State224, AtomExprType}:                      _GotoState62Action,
	{_State224, ParseErrorExprType}:                _GotoState92Action,
	{_State224, LiteralExprType}:                   _GotoState87Action,
	{_State224, IdentifierExprType}:                _GotoState79Action,
	{_State224, BlockExprType}:                     _GotoState69Action,
	{_State224, InitializeExprType}:                _GotoState84Action,
	{_State224, ImplicitStructExprType}:            _GotoState80Action,
	{_State224, AccessibleExprType}:                _GotoState104Action,
	{_State224, AccessExprType}:                    _GotoState54Action,
	{_State224, IndexExprType}:                     _GotoState82Action,
	{_State224, PostfixableExprType}:               _GotoState94Action,
	{_State224, PostfixUnaryExprType}:              _GotoState93Action,
	{_State224, PrefixableExprType}:                _GotoState97Action,
	{_State224, PrefixUnaryExprType}:               _GotoState95Action,
	{_State224, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State224, MulExprType}:                       _GotoState89Action,
	{_State224, BinaryMulExprType}:                 _GotoState66Action,
	{_State224, AddExprType}:                       _GotoState56Action,
	{_State224, BinaryAddExprType}:                 _GotoState63Action,
	{_State224, CmpExprType}:                       _GotoState73Action,
	{_State224, BinaryCmpExprType}:                 _GotoState65Action,
	{_State224, AndExprType}:                       _GotoState313Action,
	{_State224, BinaryAndExprType}:                 _GotoState64Action,
	{_State224, InitializableTypeExprType}:         _GotoState83Action,
	{_State224, SliceTypeExprType}:                 _GotoState100Action,
	{_State224, ArrayTypeExprType}:                 _GotoState59Action,
	{_State224, MapTypeExprType}:                   _GotoState88Action,
	{_State224, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State224, AnonymousFuncExprType}:             _GotoState58Action,
	{_State226, IdentifierToken}:                   _GotoState230Action,
	{_State226, UnsafeToken}:                       _GotoState53Action,
	{_State226, StructToken}:                       _GotoState50Action,
	{_State226, EnumToken}:                         _GotoState109Action,
	{_State226, TraitToken}:                        _GotoState117Action,
	{_State226, FuncToken}:                         _GotoState111Action,
	{_State226, LparenToken}:                       _GotoState113Action,
	{_State226, LbracketToken}:                     _GotoState42Action,
	{_State226, DotToken}:                          _GotoState108Action,
	{_State226, QuestionToken}:                     _GotoState115Action,
	{_State226, ExclaimToken}:                      _GotoState110Action,
	{_State226, TildeTildeToken}:                   _GotoState116Action,
	{_State226, BitNegToken}:                       _GotoState107Action,
	{_State226, BitAndToken}:                       _GotoState106Action,
	{_State226, ParseErrorToken}:                   _GotoState114Action,
	{_State226, UnsafeStatementType}:               _GotoState237Action,
	{_State226, InitializableTypeExprType}:         _GotoState125Action,
	{_State226, SliceTypeExprType}:                 _GotoState100Action,
	{_State226, ArrayTypeExprType}:                 _GotoState59Action,
	{_State226, MapTypeExprType}:                   _GotoState88Action,
	{_State226, AtomTypeExprType}:                  _GotoState118Action,
	{_State226, NamedTypeExprType}:                 _GotoState126Action,
	{_State226, InferredTypeExprType}:              _GotoState124Action,
	{_State226, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State226, ReturnableTypeExprType}:            _GotoState130Action,
	{_State226, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State226, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State226, TypeExprType}:                      _GotoState236Action,
	{_State226, BinaryTypeExprType}:                _GotoState119Action,
	{_State226, FieldDefType}:                      _GotoState316Action,
	{_State226, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State226, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State226, EnumValueDefType}:                  _GotoState314Action,
	{_State226, ImplicitEnumValueDefsType}:         _GotoState317Action,
	{_State226, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State226, ExplicitEnumValueDefsType}:         _GotoState315Action,
	{_State226, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State226, TraitTypeExprType}:                 _GotoState131Action,
	{_State226, FuncTypeExprType}:                  _GotoState121Action,
	{_State227, IdentifierToken}:                   _GotoState319Action,
	{_State227, StructToken}:                       _GotoState50Action,
	{_State227, EnumToken}:                         _GotoState109Action,
	{_State227, TraitToken}:                        _GotoState117Action,
	{_State227, FuncToken}:                         _GotoState111Action,
	{_State227, LparenToken}:                       _GotoState113Action,
	{_State227, LbracketToken}:                     _GotoState42Action,
	{_State227, DotToken}:                          _GotoState108Action,
	{_State227, QuestionToken}:                     _GotoState115Action,
	{_State227, ExclaimToken}:                      _GotoState110Action,
	{_State227, EllipsisToken}:                     _GotoState318Action,
	{_State227, TildeTildeToken}:                   _GotoState116Action,
	{_State227, BitNegToken}:                       _GotoState107Action,
	{_State227, BitAndToken}:                       _GotoState106Action,
	{_State227, ParseErrorToken}:                   _GotoState114Action,
	{_State227, InitializableTypeExprType}:         _GotoState125Action,
	{_State227, SliceTypeExprType}:                 _GotoState100Action,
	{_State227, ArrayTypeExprType}:                 _GotoState59Action,
	{_State227, MapTypeExprType}:                   _GotoState88Action,
	{_State227, AtomTypeExprType}:                  _GotoState118Action,
	{_State227, NamedTypeExprType}:                 _GotoState126Action,
	{_State227, InferredTypeExprType}:              _GotoState124Action,
	{_State227, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State227, ReturnableTypeExprType}:            _GotoState130Action,
	{_State227, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State227, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State227, TypeExprType}:                      _GotoState323Action,
	{_State227, BinaryTypeExprType}:                _GotoState119Action,
	{_State227, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State227, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State227, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State227, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State227, TraitTypeExprType}:                 _GotoState131Action,
	{_State227, ParameterDeclType}:                 _GotoState320Action,
	{_State227, ProperParameterDeclsType}:          _GotoState322Action,
	{_State227, ParameterDeclsType}:                _GotoState321Action,
	{_State227, FuncTypeExprType}:                  _GotoState121Action,
	{_State228, IdentifierToken}:                   _GotoState324Action,
	{_State230, IdentifierToken}:                   _GotoState112Action,
	{_State230, StructToken}:                       _GotoState50Action,
	{_State230, EnumToken}:                         _GotoState109Action,
	{_State230, TraitToken}:                        _GotoState117Action,
	{_State230, FuncToken}:                         _GotoState111Action,
	{_State230, LparenToken}:                       _GotoState113Action,
	{_State230, LbracketToken}:                     _GotoState42Action,
	{_State230, DotToken}:                          _GotoState325Action,
	{_State230, QuestionToken}:                     _GotoState115Action,
	{_State230, ExclaimToken}:                      _GotoState110Action,
	{_State230, DollarLbracketToken}:               _GotoState178Action,
	{_State230, TildeTildeToken}:                   _GotoState116Action,
	{_State230, BitNegToken}:                       _GotoState107Action,
	{_State230, BitAndToken}:                       _GotoState106Action,
	{_State230, ParseErrorToken}:                   _GotoState114Action,
	{_State230, GenericTypeArgumentsType}:          _GotoState229Action,
	{_State230, InitializableTypeExprType}:         _GotoState125Action,
	{_State230, SliceTypeExprType}:                 _GotoState100Action,
	{_State230, ArrayTypeExprType}:                 _GotoState59Action,
	{_State230, MapTypeExprType}:                   _GotoState88Action,
	{_State230, AtomTypeExprType}:                  _GotoState118Action,
	{_State230, NamedTypeExprType}:                 _GotoState126Action,
	{_State230, InferredTypeExprType}:              _GotoState124Action,
	{_State230, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State230, ReturnableTypeExprType}:            _GotoState130Action,
	{_State230, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State230, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State230, TypeExprType}:                      _GotoState326Action,
	{_State230, BinaryTypeExprType}:                _GotoState119Action,
	{_State230, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State230, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State230, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State230, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State230, TraitTypeExprType}:                 _GotoState131Action,
	{_State230, FuncTypeExprType}:                  _GotoState121Action,
	{_State231, OrToken}:                           _GotoState327Action,
	{_State232, AssignToken}:                       _GotoState328Action,
	{_State233, OrToken}:                           _GotoState329Action,
	{_State233, RparenToken}:                       _GotoState330Action,
	{_State234, RparenToken}:                       _GotoState331Action,
	{_State235, CommaToken}:                        _GotoState332Action,
	{_State236, AddToken}:                          _GotoState240Action,
	{_State236, SubToken}:                          _GotoState242Action,
	{_State236, MulToken}:                          _GotoState241Action,
	{_State236, BinaryTypeOpType}:                  _GotoState243Action,
	{_State238, IdentifierToken}:                   _GotoState230Action,
	{_State238, UnsafeToken}:                       _GotoState53Action,
	{_State238, StructToken}:                       _GotoState50Action,
	{_State238, EnumToken}:                         _GotoState109Action,
	{_State238, TraitToken}:                        _GotoState117Action,
	{_State238, FuncToken}:                         _GotoState333Action,
	{_State238, LparenToken}:                       _GotoState113Action,
	{_State238, LbracketToken}:                     _GotoState42Action,
	{_State238, DotToken}:                          _GotoState108Action,
	{_State238, QuestionToken}:                     _GotoState115Action,
	{_State238, ExclaimToken}:                      _GotoState110Action,
	{_State238, TildeTildeToken}:                   _GotoState116Action,
	{_State238, BitNegToken}:                       _GotoState107Action,
	{_State238, BitAndToken}:                       _GotoState106Action,
	{_State238, ParseErrorToken}:                   _GotoState114Action,
	{_State238, UnsafeStatementType}:               _GotoState237Action,
	{_State238, InitializableTypeExprType}:         _GotoState125Action,
	{_State238, SliceTypeExprType}:                 _GotoState100Action,
	{_State238, ArrayTypeExprType}:                 _GotoState59Action,
	{_State238, MapTypeExprType}:                   _GotoState88Action,
	{_State238, AtomTypeExprType}:                  _GotoState118Action,
	{_State238, NamedTypeExprType}:                 _GotoState126Action,
	{_State238, InferredTypeExprType}:              _GotoState124Action,
	{_State238, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State238, ReturnableTypeExprType}:            _GotoState130Action,
	{_State238, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State238, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State238, TypeExprType}:                      _GotoState236Action,
	{_State238, BinaryTypeExprType}:                _GotoState119Action,
	{_State238, FieldDefType}:                      _GotoState334Action,
	{_State238, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State238, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State238, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State238, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State238, TraitPropertyType}:                 _GotoState338Action,
	{_State238, ProperTraitPropertiesType}:         _GotoState336Action,
	{_State238, TraitPropertiesType}:               _GotoState337Action,
	{_State238, TraitTypeExprType}:                 _GotoState131Action,
	{_State238, FuncTypeExprType}:                  _GotoState121Action,
	{_State238, MethodSignatureType}:               _GotoState335Action,
	{_State243, IdentifierToken}:                   _GotoState112Action,
	{_State243, StructToken}:                       _GotoState50Action,
	{_State243, EnumToken}:                         _GotoState109Action,
	{_State243, TraitToken}:                        _GotoState117Action,
	{_State243, FuncToken}:                         _GotoState111Action,
	{_State243, LparenToken}:                       _GotoState113Action,
	{_State243, LbracketToken}:                     _GotoState42Action,
	{_State243, DotToken}:                          _GotoState108Action,
	{_State243, QuestionToken}:                     _GotoState115Action,
	{_State243, ExclaimToken}:                      _GotoState110Action,
	{_State243, TildeTildeToken}:                   _GotoState116Action,
	{_State243, BitNegToken}:                       _GotoState107Action,
	{_State243, BitAndToken}:                       _GotoState106Action,
	{_State243, ParseErrorToken}:                   _GotoState114Action,
	{_State243, InitializableTypeExprType}:         _GotoState125Action,
	{_State243, SliceTypeExprType}:                 _GotoState100Action,
	{_State243, ArrayTypeExprType}:                 _GotoState59Action,
	{_State243, MapTypeExprType}:                   _GotoState88Action,
	{_State243, AtomTypeExprType}:                  _GotoState118Action,
	{_State243, NamedTypeExprType}:                 _GotoState126Action,
	{_State243, InferredTypeExprType}:              _GotoState124Action,
	{_State243, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State243, ReturnableTypeExprType}:            _GotoState339Action,
	{_State243, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State243, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State243, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State243, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State243, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State243, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State243, TraitTypeExprType}:                 _GotoState131Action,
	{_State243, FuncTypeExprType}:                  _GotoState121Action,
	{_State244, IntegerLiteralToken}:               _GotoState40Action,
	{_State244, FloatLiteralToken}:                 _GotoState35Action,
	{_State244, RuneLiteralToken}:                  _GotoState48Action,
	{_State244, StringLiteralToken}:                _GotoState49Action,
	{_State244, IdentifierToken}:                   _GotoState38Action,
	{_State244, TrueToken}:                         _GotoState52Action,
	{_State244, FalseToken}:                        _GotoState34Action,
	{_State244, StructToken}:                       _GotoState50Action,
	{_State244, FuncToken}:                         _GotoState36Action,
	{_State244, VarToken}:                          _GotoState15Action,
	{_State244, LetToken}:                          _GotoState12Action,
	{_State244, NotToken}:                          _GotoState45Action,
	{_State244, LabelDeclToken}:                    _GotoState41Action,
	{_State244, LparenToken}:                       _GotoState43Action,
	{_State244, LbracketToken}:                     _GotoState42Action,
	{_State244, SubToken}:                          _GotoState51Action,
	{_State244, MulToken}:                          _GotoState44Action,
	{_State244, BitNegToken}:                       _GotoState27Action,
	{_State244, BitAndToken}:                       _GotoState26Action,
	{_State244, GreaterToken}:                      _GotoState37Action,
	{_State244, ParseErrorToken}:                   _GotoState46Action,
	{_State244, VarDeclPatternType}:                _GotoState103Action,
	{_State244, VarOrLetType}:                      _GotoState24Action,
	{_State244, ExprType}:                          _GotoState340Action,
	{_State244, OptionalLabelDeclType}:             _GotoState90Action,
	{_State244, SequenceExprType}:                  _GotoState105Action,
	{_State244, CallExprType}:                      _GotoState70Action,
	{_State244, AtomExprType}:                      _GotoState62Action,
	{_State244, ParseErrorExprType}:                _GotoState92Action,
	{_State244, LiteralExprType}:                   _GotoState87Action,
	{_State244, IdentifierExprType}:                _GotoState79Action,
	{_State244, BlockExprType}:                     _GotoState69Action,
	{_State244, InitializeExprType}:                _GotoState84Action,
	{_State244, ImplicitStructExprType}:            _GotoState80Action,
	{_State244, AccessibleExprType}:                _GotoState104Action,
	{_State244, AccessExprType}:                    _GotoState54Action,
	{_State244, IndexExprType}:                     _GotoState82Action,
	{_State244, PostfixableExprType}:               _GotoState94Action,
	{_State244, PostfixUnaryExprType}:              _GotoState93Action,
	{_State244, PrefixableExprType}:                _GotoState97Action,
	{_State244, PrefixUnaryExprType}:               _GotoState95Action,
	{_State244, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State244, MulExprType}:                       _GotoState89Action,
	{_State244, BinaryMulExprType}:                 _GotoState66Action,
	{_State244, AddExprType}:                       _GotoState56Action,
	{_State244, BinaryAddExprType}:                 _GotoState63Action,
	{_State244, CmpExprType}:                       _GotoState73Action,
	{_State244, BinaryCmpExprType}:                 _GotoState65Action,
	{_State244, AndExprType}:                       _GotoState57Action,
	{_State244, BinaryAndExprType}:                 _GotoState64Action,
	{_State244, OrExprType}:                        _GotoState91Action,
	{_State244, BinaryOrExprType}:                  _GotoState68Action,
	{_State244, InitializableTypeExprType}:         _GotoState83Action,
	{_State244, SliceTypeExprType}:                 _GotoState100Action,
	{_State244, ArrayTypeExprType}:                 _GotoState59Action,
	{_State244, MapTypeExprType}:                   _GotoState88Action,
	{_State244, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State244, AnonymousFuncExprType}:             _GotoState58Action,
	{_State245, IdentifierToken}:                   _GotoState341Action,
	{_State245, GenericParameterDefType}:           _GotoState342Action,
	{_State245, ProperGenericParameterDefsType}:    _GotoState344Action,
	{_State245, GenericParameterDefsType}:          _GotoState343Action,
	{_State246, LparenToken}:                       _GotoState345Action,
	{_State247, IdentifierToken}:                   _GotoState112Action,
	{_State247, StructToken}:                       _GotoState50Action,
	{_State247, EnumToken}:                         _GotoState109Action,
	{_State247, TraitToken}:                        _GotoState117Action,
	{_State247, FuncToken}:                         _GotoState111Action,
	{_State247, LparenToken}:                       _GotoState113Action,
	{_State247, LbracketToken}:                     _GotoState42Action,
	{_State247, DotToken}:                          _GotoState108Action,
	{_State247, QuestionToken}:                     _GotoState115Action,
	{_State247, ExclaimToken}:                      _GotoState110Action,
	{_State247, EllipsisToken}:                     _GotoState346Action,
	{_State247, TildeTildeToken}:                   _GotoState116Action,
	{_State247, BitNegToken}:                       _GotoState107Action,
	{_State247, BitAndToken}:                       _GotoState106Action,
	{_State247, ParseErrorToken}:                   _GotoState114Action,
	{_State247, InitializableTypeExprType}:         _GotoState125Action,
	{_State247, SliceTypeExprType}:                 _GotoState100Action,
	{_State247, ArrayTypeExprType}:                 _GotoState59Action,
	{_State247, MapTypeExprType}:                   _GotoState88Action,
	{_State247, AtomTypeExprType}:                  _GotoState118Action,
	{_State247, NamedTypeExprType}:                 _GotoState126Action,
	{_State247, InferredTypeExprType}:              _GotoState124Action,
	{_State247, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State247, ReturnableTypeExprType}:            _GotoState130Action,
	{_State247, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State247, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State247, TypeExprType}:                      _GotoState347Action,
	{_State247, BinaryTypeExprType}:                _GotoState119Action,
	{_State247, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State247, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State247, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State247, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State247, TraitTypeExprType}:                 _GotoState131Action,
	{_State247, FuncTypeExprType}:                  _GotoState121Action,
	{_State248, RparenToken}:                       _GotoState348Action,
	{_State249, IntegerLiteralToken}:               _GotoState40Action,
	{_State249, FloatLiteralToken}:                 _GotoState35Action,
	{_State249, RuneLiteralToken}:                  _GotoState48Action,
	{_State249, StringLiteralToken}:                _GotoState49Action,
	{_State249, IdentifierToken}:                   _GotoState38Action,
	{_State249, TrueToken}:                         _GotoState52Action,
	{_State249, FalseToken}:                        _GotoState34Action,
	{_State249, CaseToken}:                         _GotoState29Action,
	{_State249, DefaultToken}:                      _GotoState31Action,
	{_State249, ReturnToken}:                       _GotoState47Action,
	{_State249, BreakToken}:                        _GotoState28Action,
	{_State249, ContinueToken}:                     _GotoState30Action,
	{_State249, FallthroughToken}:                  _GotoState33Action,
	{_State249, ImportToken}:                       _GotoState39Action,
	{_State249, UnsafeToken}:                       _GotoState53Action,
	{_State249, StructToken}:                       _GotoState50Action,
	{_State249, FuncToken}:                         _GotoState36Action,
	{_State249, AsyncToken}:                        _GotoState25Action,
	{_State249, DeferToken}:                        _GotoState32Action,
	{_State249, VarToken}:                          _GotoState15Action,
	{_State249, LetToken}:                          _GotoState12Action,
	{_State249, NotToken}:                          _GotoState45Action,
	{_State249, LabelDeclToken}:                    _GotoState41Action,
	{_State249, LparenToken}:                       _GotoState43Action,
	{_State249, LbracketToken}:                     _GotoState42Action,
	{_State249, SubToken}:                          _GotoState51Action,
	{_State249, MulToken}:                          _GotoState44Action,
	{_State249, BitNegToken}:                       _GotoState27Action,
	{_State249, BitAndToken}:                       _GotoState26Action,
	{_State249, GreaterToken}:                      _GotoState37Action,
	{_State249, ParseErrorToken}:                   _GotoState46Action,
	{_State249, SimpleStatementType}:               _GotoState99Action,
	{_State249, StatementType}:                     _GotoState349Action,
	{_State249, ExprOrImproperStructStatementType}: _GotoState76Action,
	{_State249, ExprsType}:                         _GotoState77Action,
	{_State249, CallbackOpType}:                    _GotoState71Action,
	{_State249, CallbackOpStatementType}:           _GotoState72Action,
	{_State249, UnsafeStatementType}:               _GotoState102Action,
	{_State249, JumpStatementType}:                 _GotoState85Action,
	{_State249, JumpTypeType}:                      _GotoState86Action,
	{_State249, FallthroughStatementType}:          _GotoState78Action,
	{_State249, AssignStatementType}:               _GotoState61Action,
	{_State249, UnaryOpAssignStatementType}:        _GotoState101Action,
	{_State249, BinaryOpAssignStatementType}:       _GotoState67Action,
	{_State249, ImportStatementType}:               _GotoState81Action,
	{_State249, VarDeclPatternType}:                _GotoState103Action,
	{_State249, VarOrLetType}:                      _GotoState24Action,
	{_State249, AssignPatternType}:                 _GotoState60Action,
	{_State249, ExprType}:                          _GotoState75Action,
	{_State249, OptionalLabelDeclType}:             _GotoState90Action,
	{_State249, SequenceExprType}:                  _GotoState98Action,
	{_State249, CallExprType}:                      _GotoState70Action,
	{_State249, AtomExprType}:                      _GotoState62Action,
	{_State249, ParseErrorExprType}:                _GotoState92Action,
	{_State249, LiteralExprType}:                   _GotoState87Action,
	{_State249, IdentifierExprType}:                _GotoState79Action,
	{_State249, BlockExprType}:                     _GotoState69Action,
	{_State249, InitializeExprType}:                _GotoState84Action,
	{_State249, ImplicitStructExprType}:            _GotoState80Action,
	{_State249, AccessibleExprType}:                _GotoState55Action,
	{_State249, AccessExprType}:                    _GotoState54Action,
	{_State249, IndexExprType}:                     _GotoState82Action,
	{_State249, PostfixableExprType}:               _GotoState94Action,
	{_State249, PostfixUnaryExprType}:              _GotoState93Action,
	{_State249, PrefixableExprType}:                _GotoState97Action,
	{_State249, PrefixUnaryExprType}:               _GotoState95Action,
	{_State249, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State249, MulExprType}:                       _GotoState89Action,
	{_State249, BinaryMulExprType}:                 _GotoState66Action,
	{_State249, AddExprType}:                       _GotoState56Action,
	{_State249, BinaryAddExprType}:                 _GotoState63Action,
	{_State249, CmpExprType}:                       _GotoState73Action,
	{_State249, BinaryCmpExprType}:                 _GotoState65Action,
	{_State249, AndExprType}:                       _GotoState57Action,
	{_State249, BinaryAndExprType}:                 _GotoState64Action,
	{_State249, OrExprType}:                        _GotoState91Action,
	{_State249, BinaryOrExprType}:                  _GotoState68Action,
	{_State249, InitializableTypeExprType}:         _GotoState83Action,
	{_State249, SliceTypeExprType}:                 _GotoState100Action,
	{_State249, ArrayTypeExprType}:                 _GotoState59Action,
	{_State249, MapTypeExprType}:                   _GotoState88Action,
	{_State249, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State249, AnonymousFuncExprType}:             _GotoState58Action,
	{_State250, IntegerLiteralToken}:               _GotoState40Action,
	{_State250, FloatLiteralToken}:                 _GotoState35Action,
	{_State250, RuneLiteralToken}:                  _GotoState48Action,
	{_State250, StringLiteralToken}:                _GotoState49Action,
	{_State250, IdentifierToken}:                   _GotoState38Action,
	{_State250, TrueToken}:                         _GotoState52Action,
	{_State250, FalseToken}:                        _GotoState34Action,
	{_State250, CaseToken}:                         _GotoState29Action,
	{_State250, DefaultToken}:                      _GotoState31Action,
	{_State250, ReturnToken}:                       _GotoState47Action,
	{_State250, BreakToken}:                        _GotoState28Action,
	{_State250, ContinueToken}:                     _GotoState30Action,
	{_State250, FallthroughToken}:                  _GotoState33Action,
	{_State250, ImportToken}:                       _GotoState39Action,
	{_State250, UnsafeToken}:                       _GotoState53Action,
	{_State250, StructToken}:                       _GotoState50Action,
	{_State250, FuncToken}:                         _GotoState36Action,
	{_State250, AsyncToken}:                        _GotoState25Action,
	{_State250, DeferToken}:                        _GotoState32Action,
	{_State250, VarToken}:                          _GotoState15Action,
	{_State250, LetToken}:                          _GotoState12Action,
	{_State250, NotToken}:                          _GotoState45Action,
	{_State250, LabelDeclToken}:                    _GotoState41Action,
	{_State250, LparenToken}:                       _GotoState43Action,
	{_State250, LbracketToken}:                     _GotoState42Action,
	{_State250, SubToken}:                          _GotoState51Action,
	{_State250, MulToken}:                          _GotoState44Action,
	{_State250, BitNegToken}:                       _GotoState27Action,
	{_State250, BitAndToken}:                       _GotoState26Action,
	{_State250, GreaterToken}:                      _GotoState37Action,
	{_State250, ParseErrorToken}:                   _GotoState46Action,
	{_State250, SimpleStatementType}:               _GotoState99Action,
	{_State250, StatementType}:                     _GotoState350Action,
	{_State250, ExprOrImproperStructStatementType}: _GotoState76Action,
	{_State250, ExprsType}:                         _GotoState77Action,
	{_State250, CallbackOpType}:                    _GotoState71Action,
	{_State250, CallbackOpStatementType}:           _GotoState72Action,
	{_State250, UnsafeStatementType}:               _GotoState102Action,
	{_State250, JumpStatementType}:                 _GotoState85Action,
	{_State250, JumpTypeType}:                      _GotoState86Action,
	{_State250, FallthroughStatementType}:          _GotoState78Action,
	{_State250, AssignStatementType}:               _GotoState61Action,
	{_State250, UnaryOpAssignStatementType}:        _GotoState101Action,
	{_State250, BinaryOpAssignStatementType}:       _GotoState67Action,
	{_State250, ImportStatementType}:               _GotoState81Action,
	{_State250, VarDeclPatternType}:                _GotoState103Action,
	{_State250, VarOrLetType}:                      _GotoState24Action,
	{_State250, AssignPatternType}:                 _GotoState60Action,
	{_State250, ExprType}:                          _GotoState75Action,
	{_State250, OptionalLabelDeclType}:             _GotoState90Action,
	{_State250, SequenceExprType}:                  _GotoState98Action,
	{_State250, CallExprType}:                      _GotoState70Action,
	{_State250, AtomExprType}:                      _GotoState62Action,
	{_State250, ParseErrorExprType}:                _GotoState92Action,
	{_State250, LiteralExprType}:                   _GotoState87Action,
	{_State250, IdentifierExprType}:                _GotoState79Action,
	{_State250, BlockExprType}:                     _GotoState69Action,
	{_State250, InitializeExprType}:                _GotoState84Action,
	{_State250, ImplicitStructExprType}:            _GotoState80Action,
	{_State250, AccessibleExprType}:                _GotoState55Action,
	{_State250, AccessExprType}:                    _GotoState54Action,
	{_State250, IndexExprType}:                     _GotoState82Action,
	{_State250, PostfixableExprType}:               _GotoState94Action,
	{_State250, PostfixUnaryExprType}:              _GotoState93Action,
	{_State250, PrefixableExprType}:                _GotoState97Action,
	{_State250, PrefixUnaryExprType}:               _GotoState95Action,
	{_State250, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State250, MulExprType}:                       _GotoState89Action,
	{_State250, BinaryMulExprType}:                 _GotoState66Action,
	{_State250, AddExprType}:                       _GotoState56Action,
	{_State250, BinaryAddExprType}:                 _GotoState63Action,
	{_State250, CmpExprType}:                       _GotoState73Action,
	{_State250, BinaryCmpExprType}:                 _GotoState65Action,
	{_State250, AndExprType}:                       _GotoState57Action,
	{_State250, BinaryAndExprType}:                 _GotoState64Action,
	{_State250, OrExprType}:                        _GotoState91Action,
	{_State250, BinaryOrExprType}:                  _GotoState68Action,
	{_State250, InitializableTypeExprType}:         _GotoState83Action,
	{_State250, SliceTypeExprType}:                 _GotoState100Action,
	{_State250, ArrayTypeExprType}:                 _GotoState59Action,
	{_State250, MapTypeExprType}:                   _GotoState88Action,
	{_State250, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State250, AnonymousFuncExprType}:             _GotoState58Action,
	{_State252, IdentifierToken}:                   _GotoState112Action,
	{_State252, StructToken}:                       _GotoState50Action,
	{_State252, EnumToken}:                         _GotoState109Action,
	{_State252, TraitToken}:                        _GotoState117Action,
	{_State252, FuncToken}:                         _GotoState111Action,
	{_State252, LparenToken}:                       _GotoState113Action,
	{_State252, LbracketToken}:                     _GotoState42Action,
	{_State252, DotToken}:                          _GotoState108Action,
	{_State252, QuestionToken}:                     _GotoState115Action,
	{_State252, ExclaimToken}:                      _GotoState110Action,
	{_State252, TildeTildeToken}:                   _GotoState116Action,
	{_State252, BitNegToken}:                       _GotoState107Action,
	{_State252, BitAndToken}:                       _GotoState106Action,
	{_State252, ParseErrorToken}:                   _GotoState114Action,
	{_State252, InitializableTypeExprType}:         _GotoState125Action,
	{_State252, SliceTypeExprType}:                 _GotoState100Action,
	{_State252, ArrayTypeExprType}:                 _GotoState59Action,
	{_State252, MapTypeExprType}:                   _GotoState88Action,
	{_State252, AtomTypeExprType}:                  _GotoState118Action,
	{_State252, NamedTypeExprType}:                 _GotoState126Action,
	{_State252, InferredTypeExprType}:              _GotoState124Action,
	{_State252, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State252, ReturnableTypeExprType}:            _GotoState130Action,
	{_State252, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State252, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State252, TypeExprType}:                      _GotoState351Action,
	{_State252, BinaryTypeExprType}:                _GotoState119Action,
	{_State252, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State252, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State252, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State252, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State252, TraitTypeExprType}:                 _GotoState131Action,
	{_State252, FuncTypeExprType}:                  _GotoState121Action,
	{_State253, IdentifierToken}:                   _GotoState112Action,
	{_State253, StructToken}:                       _GotoState50Action,
	{_State253, EnumToken}:                         _GotoState109Action,
	{_State253, TraitToken}:                        _GotoState117Action,
	{_State253, FuncToken}:                         _GotoState111Action,
	{_State253, LparenToken}:                       _GotoState113Action,
	{_State253, LbracketToken}:                     _GotoState42Action,
	{_State253, DotToken}:                          _GotoState108Action,
	{_State253, QuestionToken}:                     _GotoState115Action,
	{_State253, ExclaimToken}:                      _GotoState110Action,
	{_State253, TildeTildeToken}:                   _GotoState116Action,
	{_State253, BitNegToken}:                       _GotoState107Action,
	{_State253, BitAndToken}:                       _GotoState106Action,
	{_State253, ParseErrorToken}:                   _GotoState114Action,
	{_State253, InitializableTypeExprType}:         _GotoState125Action,
	{_State253, SliceTypeExprType}:                 _GotoState100Action,
	{_State253, ArrayTypeExprType}:                 _GotoState59Action,
	{_State253, MapTypeExprType}:                   _GotoState88Action,
	{_State253, AtomTypeExprType}:                  _GotoState118Action,
	{_State253, NamedTypeExprType}:                 _GotoState126Action,
	{_State253, InferredTypeExprType}:              _GotoState124Action,
	{_State253, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State253, ReturnableTypeExprType}:            _GotoState130Action,
	{_State253, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State253, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State253, TypeExprType}:                      _GotoState352Action,
	{_State253, BinaryTypeExprType}:                _GotoState119Action,
	{_State253, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State253, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State253, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State253, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State253, TraitTypeExprType}:                 _GotoState131Action,
	{_State253, FuncTypeExprType}:                  _GotoState121Action,
	{_State257, AssignToken}:                       _GotoState353Action,
	{_State259, RparenToken}:                       _GotoState355Action,
	{_State259, CommaToken}:                        _GotoState354Action,
	{_State262, AddToken}:                          _GotoState240Action,
	{_State262, SubToken}:                          _GotoState242Action,
	{_State262, MulToken}:                          _GotoState241Action,
	{_State262, BinaryTypeOpType}:                  _GotoState243Action,
	{_State263, LparenToken}:                       _GotoState43Action,
	{_State263, ImplicitStructExprType}:            _GotoState356Action,
	{_State264, IdentifierToken}:                   _GotoState357Action,
	{_State265, IntegerLiteralToken}:               _GotoState40Action,
	{_State265, FloatLiteralToken}:                 _GotoState35Action,
	{_State265, RuneLiteralToken}:                  _GotoState48Action,
	{_State265, StringLiteralToken}:                _GotoState49Action,
	{_State265, IdentifierToken}:                   _GotoState38Action,
	{_State265, TrueToken}:                         _GotoState52Action,
	{_State265, FalseToken}:                        _GotoState34Action,
	{_State265, ReturnToken}:                       _GotoState47Action,
	{_State265, BreakToken}:                        _GotoState28Action,
	{_State265, ContinueToken}:                     _GotoState30Action,
	{_State265, FallthroughToken}:                  _GotoState33Action,
	{_State265, UnsafeToken}:                       _GotoState53Action,
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
	{_State265, SimpleStatementType}:               _GotoState268Action,
	{_State265, OptionalSimpleStatementType}:       _GotoState358Action,
	{_State265, ExprOrImproperStructStatementType}: _GotoState76Action,
	{_State265, ExprsType}:                         _GotoState77Action,
	{_State265, CallbackOpType}:                    _GotoState71Action,
	{_State265, CallbackOpStatementType}:           _GotoState72Action,
	{_State265, UnsafeStatementType}:               _GotoState102Action,
	{_State265, JumpStatementType}:                 _GotoState85Action,
	{_State265, JumpTypeType}:                      _GotoState86Action,
	{_State265, FallthroughStatementType}:          _GotoState78Action,
	{_State265, AssignStatementType}:               _GotoState61Action,
	{_State265, UnaryOpAssignStatementType}:        _GotoState101Action,
	{_State265, BinaryOpAssignStatementType}:       _GotoState67Action,
	{_State265, VarDeclPatternType}:                _GotoState103Action,
	{_State265, VarOrLetType}:                      _GotoState24Action,
	{_State265, AssignPatternType}:                 _GotoState60Action,
	{_State265, ExprType}:                          _GotoState75Action,
	{_State265, OptionalLabelDeclType}:             _GotoState90Action,
	{_State265, SequenceExprType}:                  _GotoState98Action,
	{_State265, CallExprType}:                      _GotoState70Action,
	{_State265, AtomExprType}:                      _GotoState62Action,
	{_State265, ParseErrorExprType}:                _GotoState92Action,
	{_State265, LiteralExprType}:                   _GotoState87Action,
	{_State265, IdentifierExprType}:                _GotoState79Action,
	{_State265, BlockExprType}:                     _GotoState69Action,
	{_State265, InitializeExprType}:                _GotoState84Action,
	{_State265, ImplicitStructExprType}:            _GotoState80Action,
	{_State265, AccessibleExprType}:                _GotoState55Action,
	{_State265, AccessExprType}:                    _GotoState54Action,
	{_State265, IndexExprType}:                     _GotoState82Action,
	{_State265, PostfixableExprType}:               _GotoState94Action,
	{_State265, PostfixUnaryExprType}:              _GotoState93Action,
	{_State265, PrefixableExprType}:                _GotoState97Action,
	{_State265, PrefixUnaryExprType}:               _GotoState95Action,
	{_State265, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State265, MulExprType}:                       _GotoState89Action,
	{_State265, BinaryMulExprType}:                 _GotoState66Action,
	{_State265, AddExprType}:                       _GotoState56Action,
	{_State265, BinaryAddExprType}:                 _GotoState63Action,
	{_State265, CmpExprType}:                       _GotoState73Action,
	{_State265, BinaryCmpExprType}:                 _GotoState65Action,
	{_State265, AndExprType}:                       _GotoState57Action,
	{_State265, BinaryAndExprType}:                 _GotoState64Action,
	{_State265, OrExprType}:                        _GotoState91Action,
	{_State265, BinaryOrExprType}:                  _GotoState68Action,
	{_State265, InitializableTypeExprType}:         _GotoState83Action,
	{_State265, SliceTypeExprType}:                 _GotoState100Action,
	{_State265, ArrayTypeExprType}:                 _GotoState59Action,
	{_State265, MapTypeExprType}:                   _GotoState88Action,
	{_State265, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State265, AnonymousFuncExprType}:             _GotoState58Action,
	{_State266, IntegerLiteralToken}:               _GotoState40Action,
	{_State266, FloatLiteralToken}:                 _GotoState35Action,
	{_State266, RuneLiteralToken}:                  _GotoState48Action,
	{_State266, StringLiteralToken}:                _GotoState49Action,
	{_State266, IdentifierToken}:                   _GotoState38Action,
	{_State266, TrueToken}:                         _GotoState52Action,
	{_State266, FalseToken}:                        _GotoState34Action,
	{_State266, StructToken}:                       _GotoState50Action,
	{_State266, FuncToken}:                         _GotoState36Action,
	{_State266, VarToken}:                          _GotoState146Action,
	{_State266, LetToken}:                          _GotoState12Action,
	{_State266, NotToken}:                          _GotoState45Action,
	{_State266, LabelDeclToken}:                    _GotoState41Action,
	{_State266, LparenToken}:                       _GotoState43Action,
	{_State266, LbracketToken}:                     _GotoState42Action,
	{_State266, DotToken}:                          _GotoState145Action,
	{_State266, SubToken}:                          _GotoState51Action,
	{_State266, MulToken}:                          _GotoState44Action,
	{_State266, BitNegToken}:                       _GotoState27Action,
	{_State266, BitAndToken}:                       _GotoState26Action,
	{_State266, GreaterToken}:                      _GotoState37Action,
	{_State266, ParseErrorToken}:                   _GotoState46Action,
	{_State266, VarDeclPatternType}:                _GotoState103Action,
	{_State266, VarOrLetType}:                      _GotoState24Action,
	{_State266, AssignPatternType}:                 _GotoState147Action,
	{_State266, CasePatternType}:                   _GotoState359Action,
	{_State266, OptionalLabelDeclType}:             _GotoState150Action,
	{_State266, SequenceExprType}:                  _GotoState151Action,
	{_State266, CallExprType}:                      _GotoState70Action,
	{_State266, AtomExprType}:                      _GotoState62Action,
	{_State266, ParseErrorExprType}:                _GotoState92Action,
	{_State266, LiteralExprType}:                   _GotoState87Action,
	{_State266, IdentifierExprType}:                _GotoState79Action,
	{_State266, BlockExprType}:                     _GotoState69Action,
	{_State266, InitializeExprType}:                _GotoState84Action,
	{_State266, ImplicitStructExprType}:            _GotoState80Action,
	{_State266, AccessibleExprType}:                _GotoState104Action,
	{_State266, AccessExprType}:                    _GotoState54Action,
	{_State266, IndexExprType}:                     _GotoState82Action,
	{_State266, PostfixableExprType}:               _GotoState94Action,
	{_State266, PostfixUnaryExprType}:              _GotoState93Action,
	{_State266, PrefixableExprType}:                _GotoState97Action,
	{_State266, PrefixUnaryExprType}:               _GotoState95Action,
	{_State266, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State266, MulExprType}:                       _GotoState89Action,
	{_State266, BinaryMulExprType}:                 _GotoState66Action,
	{_State266, AddExprType}:                       _GotoState56Action,
	{_State266, BinaryAddExprType}:                 _GotoState63Action,
	{_State266, CmpExprType}:                       _GotoState73Action,
	{_State266, BinaryCmpExprType}:                 _GotoState65Action,
	{_State266, AndExprType}:                       _GotoState57Action,
	{_State266, BinaryAndExprType}:                 _GotoState64Action,
	{_State266, OrExprType}:                        _GotoState91Action,
	{_State266, BinaryOrExprType}:                  _GotoState68Action,
	{_State266, InitializableTypeExprType}:         _GotoState83Action,
	{_State266, SliceTypeExprType}:                 _GotoState100Action,
	{_State266, ArrayTypeExprType}:                 _GotoState59Action,
	{_State266, MapTypeExprType}:                   _GotoState88Action,
	{_State266, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State266, AnonymousFuncExprType}:             _GotoState58Action,
	{_State270, RparenToken}:                       _GotoState360Action,
	{_State271, CommaToken}:                        _GotoState361Action,
	{_State273, RparenToken}:                       _GotoState362Action,
	{_State274, NewlinesToken}:                     _GotoState364Action,
	{_State274, CommaToken}:                        _GotoState363Action,
	{_State275, IdentifierToken}:                   _GotoState365Action,
	{_State276, IdentifierToken}:                   _GotoState112Action,
	{_State276, StructToken}:                       _GotoState50Action,
	{_State276, EnumToken}:                         _GotoState109Action,
	{_State276, TraitToken}:                        _GotoState117Action,
	{_State276, FuncToken}:                         _GotoState111Action,
	{_State276, LparenToken}:                       _GotoState113Action,
	{_State276, LbracketToken}:                     _GotoState42Action,
	{_State276, DotToken}:                          _GotoState108Action,
	{_State276, QuestionToken}:                     _GotoState115Action,
	{_State276, ExclaimToken}:                      _GotoState110Action,
	{_State276, TildeTildeToken}:                   _GotoState116Action,
	{_State276, BitNegToken}:                       _GotoState107Action,
	{_State276, BitAndToken}:                       _GotoState106Action,
	{_State276, ParseErrorToken}:                   _GotoState114Action,
	{_State276, InitializableTypeExprType}:         _GotoState125Action,
	{_State276, SliceTypeExprType}:                 _GotoState100Action,
	{_State276, ArrayTypeExprType}:                 _GotoState59Action,
	{_State276, MapTypeExprType}:                   _GotoState88Action,
	{_State276, AtomTypeExprType}:                  _GotoState118Action,
	{_State276, NamedTypeExprType}:                 _GotoState126Action,
	{_State276, InferredTypeExprType}:              _GotoState124Action,
	{_State276, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State276, ReturnableTypeExprType}:            _GotoState130Action,
	{_State276, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State276, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State276, TypeExprType}:                      _GotoState366Action,
	{_State276, BinaryTypeExprType}:                _GotoState119Action,
	{_State276, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State276, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State276, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State276, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State276, TraitTypeExprType}:                 _GotoState131Action,
	{_State276, FuncTypeExprType}:                  _GotoState121Action,
	{_State277, IntegerLiteralToken}:               _GotoState367Action,
	{_State280, IntegerLiteralToken}:               _GotoState40Action,
	{_State280, FloatLiteralToken}:                 _GotoState35Action,
	{_State280, RuneLiteralToken}:                  _GotoState48Action,
	{_State280, StringLiteralToken}:                _GotoState49Action,
	{_State280, IdentifierToken}:                   _GotoState38Action,
	{_State280, TrueToken}:                         _GotoState52Action,
	{_State280, FalseToken}:                        _GotoState34Action,
	{_State280, StructToken}:                       _GotoState50Action,
	{_State280, FuncToken}:                         _GotoState36Action,
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
	{_State280, VarDeclPatternType}:                _GotoState103Action,
	{_State280, VarOrLetType}:                      _GotoState24Action,
	{_State280, ExprType}:                          _GotoState368Action,
	{_State280, OptionalLabelDeclType}:             _GotoState90Action,
	{_State280, SequenceExprType}:                  _GotoState105Action,
	{_State280, CallExprType}:                      _GotoState70Action,
	{_State280, AtomExprType}:                      _GotoState62Action,
	{_State280, ParseErrorExprType}:                _GotoState92Action,
	{_State280, LiteralExprType}:                   _GotoState87Action,
	{_State280, IdentifierExprType}:                _GotoState79Action,
	{_State280, BlockExprType}:                     _GotoState69Action,
	{_State280, InitializeExprType}:                _GotoState84Action,
	{_State280, ImplicitStructExprType}:            _GotoState80Action,
	{_State280, AccessibleExprType}:                _GotoState104Action,
	{_State280, AccessExprType}:                    _GotoState54Action,
	{_State280, IndexExprType}:                     _GotoState82Action,
	{_State280, PostfixableExprType}:               _GotoState94Action,
	{_State280, PostfixUnaryExprType}:              _GotoState93Action,
	{_State280, PrefixableExprType}:                _GotoState97Action,
	{_State280, PrefixUnaryExprType}:               _GotoState95Action,
	{_State280, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State280, MulExprType}:                       _GotoState89Action,
	{_State280, BinaryMulExprType}:                 _GotoState66Action,
	{_State280, AddExprType}:                       _GotoState56Action,
	{_State280, BinaryAddExprType}:                 _GotoState63Action,
	{_State280, CmpExprType}:                       _GotoState73Action,
	{_State280, BinaryCmpExprType}:                 _GotoState65Action,
	{_State280, AndExprType}:                       _GotoState57Action,
	{_State280, BinaryAndExprType}:                 _GotoState64Action,
	{_State280, OrExprType}:                        _GotoState91Action,
	{_State280, BinaryOrExprType}:                  _GotoState68Action,
	{_State280, InitializableTypeExprType}:         _GotoState83Action,
	{_State280, SliceTypeExprType}:                 _GotoState100Action,
	{_State280, ArrayTypeExprType}:                 _GotoState59Action,
	{_State280, MapTypeExprType}:                   _GotoState88Action,
	{_State280, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State280, AnonymousFuncExprType}:             _GotoState58Action,
	{_State282, IntegerLiteralToken}:               _GotoState40Action,
	{_State282, FloatLiteralToken}:                 _GotoState35Action,
	{_State282, RuneLiteralToken}:                  _GotoState48Action,
	{_State282, StringLiteralToken}:                _GotoState49Action,
	{_State282, IdentifierToken}:                   _GotoState38Action,
	{_State282, TrueToken}:                         _GotoState52Action,
	{_State282, FalseToken}:                        _GotoState34Action,
	{_State282, StructToken}:                       _GotoState50Action,
	{_State282, FuncToken}:                         _GotoState36Action,
	{_State282, VarToken}:                          _GotoState15Action,
	{_State282, LetToken}:                          _GotoState12Action,
	{_State282, NotToken}:                          _GotoState45Action,
	{_State282, LabelDeclToken}:                    _GotoState41Action,
	{_State282, LparenToken}:                       _GotoState43Action,
	{_State282, LbracketToken}:                     _GotoState42Action,
	{_State282, SubToken}:                          _GotoState51Action,
	{_State282, MulToken}:                          _GotoState44Action,
	{_State282, BitNegToken}:                       _GotoState27Action,
	{_State282, BitAndToken}:                       _GotoState26Action,
	{_State282, GreaterToken}:                      _GotoState37Action,
	{_State282, ParseErrorToken}:                   _GotoState46Action,
	{_State282, VarDeclPatternType}:                _GotoState103Action,
	{_State282, VarOrLetType}:                      _GotoState24Action,
	{_State282, ExprType}:                          _GotoState369Action,
	{_State282, OptionalLabelDeclType}:             _GotoState90Action,
	{_State282, SequenceExprType}:                  _GotoState105Action,
	{_State282, CallExprType}:                      _GotoState70Action,
	{_State282, AtomExprType}:                      _GotoState62Action,
	{_State282, ParseErrorExprType}:                _GotoState92Action,
	{_State282, LiteralExprType}:                   _GotoState87Action,
	{_State282, IdentifierExprType}:                _GotoState79Action,
	{_State282, BlockExprType}:                     _GotoState69Action,
	{_State282, InitializeExprType}:                _GotoState84Action,
	{_State282, ImplicitStructExprType}:            _GotoState80Action,
	{_State282, AccessibleExprType}:                _GotoState104Action,
	{_State282, AccessExprType}:                    _GotoState54Action,
	{_State282, IndexExprType}:                     _GotoState82Action,
	{_State282, PostfixableExprType}:               _GotoState94Action,
	{_State282, PostfixUnaryExprType}:              _GotoState93Action,
	{_State282, PrefixableExprType}:                _GotoState97Action,
	{_State282, PrefixUnaryExprType}:               _GotoState95Action,
	{_State282, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State282, MulExprType}:                       _GotoState89Action,
	{_State282, BinaryMulExprType}:                 _GotoState66Action,
	{_State282, AddExprType}:                       _GotoState56Action,
	{_State282, BinaryAddExprType}:                 _GotoState63Action,
	{_State282, CmpExprType}:                       _GotoState73Action,
	{_State282, BinaryCmpExprType}:                 _GotoState65Action,
	{_State282, AndExprType}:                       _GotoState57Action,
	{_State282, BinaryAndExprType}:                 _GotoState64Action,
	{_State282, OrExprType}:                        _GotoState91Action,
	{_State282, BinaryOrExprType}:                  _GotoState68Action,
	{_State282, InitializableTypeExprType}:         _GotoState83Action,
	{_State282, SliceTypeExprType}:                 _GotoState100Action,
	{_State282, ArrayTypeExprType}:                 _GotoState59Action,
	{_State282, MapTypeExprType}:                   _GotoState88Action,
	{_State282, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State282, AnonymousFuncExprType}:             _GotoState58Action,
	{_State283, IntegerLiteralToken}:               _GotoState40Action,
	{_State283, FloatLiteralToken}:                 _GotoState35Action,
	{_State283, RuneLiteralToken}:                  _GotoState48Action,
	{_State283, StringLiteralToken}:                _GotoState49Action,
	{_State283, IdentifierToken}:                   _GotoState38Action,
	{_State283, TrueToken}:                         _GotoState52Action,
	{_State283, FalseToken}:                        _GotoState34Action,
	{_State283, StructToken}:                       _GotoState50Action,
	{_State283, FuncToken}:                         _GotoState36Action,
	{_State283, VarToken}:                          _GotoState15Action,
	{_State283, LetToken}:                          _GotoState12Action,
	{_State283, NotToken}:                          _GotoState45Action,
	{_State283, LabelDeclToken}:                    _GotoState41Action,
	{_State283, LparenToken}:                       _GotoState43Action,
	{_State283, LbracketToken}:                     _GotoState42Action,
	{_State283, SubToken}:                          _GotoState51Action,
	{_State283, MulToken}:                          _GotoState44Action,
	{_State283, BitNegToken}:                       _GotoState27Action,
	{_State283, BitAndToken}:                       _GotoState26Action,
	{_State283, GreaterToken}:                      _GotoState37Action,
	{_State283, ParseErrorToken}:                   _GotoState46Action,
	{_State283, VarDeclPatternType}:                _GotoState103Action,
	{_State283, VarOrLetType}:                      _GotoState24Action,
	{_State283, ExprType}:                          _GotoState370Action,
	{_State283, OptionalLabelDeclType}:             _GotoState90Action,
	{_State283, SequenceExprType}:                  _GotoState105Action,
	{_State283, CallExprType}:                      _GotoState70Action,
	{_State283, AtomExprType}:                      _GotoState62Action,
	{_State283, ParseErrorExprType}:                _GotoState92Action,
	{_State283, LiteralExprType}:                   _GotoState87Action,
	{_State283, IdentifierExprType}:                _GotoState79Action,
	{_State283, BlockExprType}:                     _GotoState69Action,
	{_State283, InitializeExprType}:                _GotoState84Action,
	{_State283, ImplicitStructExprType}:            _GotoState80Action,
	{_State283, AccessibleExprType}:                _GotoState104Action,
	{_State283, AccessExprType}:                    _GotoState54Action,
	{_State283, IndexExprType}:                     _GotoState82Action,
	{_State283, PostfixableExprType}:               _GotoState94Action,
	{_State283, PostfixUnaryExprType}:              _GotoState93Action,
	{_State283, PrefixableExprType}:                _GotoState97Action,
	{_State283, PrefixUnaryExprType}:               _GotoState95Action,
	{_State283, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State283, MulExprType}:                       _GotoState89Action,
	{_State283, BinaryMulExprType}:                 _GotoState66Action,
	{_State283, AddExprType}:                       _GotoState56Action,
	{_State283, BinaryAddExprType}:                 _GotoState63Action,
	{_State283, CmpExprType}:                       _GotoState73Action,
	{_State283, BinaryCmpExprType}:                 _GotoState65Action,
	{_State283, AndExprType}:                       _GotoState57Action,
	{_State283, BinaryAndExprType}:                 _GotoState64Action,
	{_State283, OrExprType}:                        _GotoState91Action,
	{_State283, BinaryOrExprType}:                  _GotoState68Action,
	{_State283, InitializableTypeExprType}:         _GotoState83Action,
	{_State283, SliceTypeExprType}:                 _GotoState100Action,
	{_State283, ArrayTypeExprType}:                 _GotoState59Action,
	{_State283, MapTypeExprType}:                   _GotoState88Action,
	{_State283, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State283, AnonymousFuncExprType}:             _GotoState58Action,
	{_State285, IntegerLiteralToken}:               _GotoState40Action,
	{_State285, FloatLiteralToken}:                 _GotoState35Action,
	{_State285, RuneLiteralToken}:                  _GotoState48Action,
	{_State285, StringLiteralToken}:                _GotoState49Action,
	{_State285, IdentifierToken}:                   _GotoState161Action,
	{_State285, TrueToken}:                         _GotoState52Action,
	{_State285, FalseToken}:                        _GotoState34Action,
	{_State285, StructToken}:                       _GotoState50Action,
	{_State285, FuncToken}:                         _GotoState36Action,
	{_State285, VarToken}:                          _GotoState15Action,
	{_State285, LetToken}:                          _GotoState12Action,
	{_State285, NotToken}:                          _GotoState45Action,
	{_State285, LabelDeclToken}:                    _GotoState41Action,
	{_State285, LparenToken}:                       _GotoState43Action,
	{_State285, LbracketToken}:                     _GotoState42Action,
	{_State285, ColonToken}:                        _GotoState159Action,
	{_State285, EllipsisToken}:                     _GotoState160Action,
	{_State285, SubToken}:                          _GotoState51Action,
	{_State285, MulToken}:                          _GotoState44Action,
	{_State285, BitNegToken}:                       _GotoState27Action,
	{_State285, BitAndToken}:                       _GotoState26Action,
	{_State285, GreaterToken}:                      _GotoState37Action,
	{_State285, ParseErrorToken}:                   _GotoState46Action,
	{_State285, VarDeclPatternType}:                _GotoState103Action,
	{_State285, VarOrLetType}:                      _GotoState24Action,
	{_State285, ExprType}:                          _GotoState165Action,
	{_State285, OptionalLabelDeclType}:             _GotoState90Action,
	{_State285, SequenceExprType}:                  _GotoState105Action,
	{_State285, CallExprType}:                      _GotoState70Action,
	{_State285, ArgumentType}:                      _GotoState371Action,
	{_State285, ColonExprType}:                     _GotoState164Action,
	{_State285, AtomExprType}:                      _GotoState62Action,
	{_State285, ParseErrorExprType}:                _GotoState92Action,
	{_State285, LiteralExprType}:                   _GotoState87Action,
	{_State285, IdentifierExprType}:                _GotoState79Action,
	{_State285, BlockExprType}:                     _GotoState69Action,
	{_State285, InitializeExprType}:                _GotoState84Action,
	{_State285, ImplicitStructExprType}:            _GotoState80Action,
	{_State285, AccessibleExprType}:                _GotoState104Action,
	{_State285, AccessExprType}:                    _GotoState54Action,
	{_State285, IndexExprType}:                     _GotoState82Action,
	{_State285, PostfixableExprType}:               _GotoState94Action,
	{_State285, PostfixUnaryExprType}:              _GotoState93Action,
	{_State285, PrefixableExprType}:                _GotoState97Action,
	{_State285, PrefixUnaryExprType}:               _GotoState95Action,
	{_State285, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State285, MulExprType}:                       _GotoState89Action,
	{_State285, BinaryMulExprType}:                 _GotoState66Action,
	{_State285, AddExprType}:                       _GotoState56Action,
	{_State285, BinaryAddExprType}:                 _GotoState63Action,
	{_State285, CmpExprType}:                       _GotoState73Action,
	{_State285, BinaryCmpExprType}:                 _GotoState65Action,
	{_State285, AndExprType}:                       _GotoState57Action,
	{_State285, BinaryAndExprType}:                 _GotoState64Action,
	{_State285, OrExprType}:                        _GotoState91Action,
	{_State285, BinaryOrExprType}:                  _GotoState68Action,
	{_State285, InitializableTypeExprType}:         _GotoState83Action,
	{_State285, SliceTypeExprType}:                 _GotoState100Action,
	{_State285, ArrayTypeExprType}:                 _GotoState59Action,
	{_State285, MapTypeExprType}:                   _GotoState88Action,
	{_State285, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State285, AnonymousFuncExprType}:             _GotoState58Action,
	{_State286, RparenToken}:                       _GotoState372Action,
	{_State288, NewlinesToken}:                     _GotoState374Action,
	{_State288, CommaToken}:                        _GotoState373Action,
	{_State289, GreaterToken}:                      _GotoState375Action,
	{_State290, CommaToken}:                        _GotoState376Action,
	{_State291, RbracketToken}:                     _GotoState377Action,
	{_State292, AddToken}:                          _GotoState240Action,
	{_State292, SubToken}:                          _GotoState242Action,
	{_State292, MulToken}:                          _GotoState241Action,
	{_State292, BinaryTypeOpType}:                  _GotoState243Action,
	{_State294, RbracketToken}:                     _GotoState378Action,
	{_State296, IntegerLiteralToken}:               _GotoState40Action,
	{_State296, FloatLiteralToken}:                 _GotoState35Action,
	{_State296, RuneLiteralToken}:                  _GotoState48Action,
	{_State296, StringLiteralToken}:                _GotoState49Action,
	{_State296, IdentifierToken}:                   _GotoState161Action,
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
	{_State296, ColonToken}:                        _GotoState159Action,
	{_State296, EllipsisToken}:                     _GotoState160Action,
	{_State296, SubToken}:                          _GotoState51Action,
	{_State296, MulToken}:                          _GotoState44Action,
	{_State296, BitNegToken}:                       _GotoState27Action,
	{_State296, BitAndToken}:                       _GotoState26Action,
	{_State296, GreaterToken}:                      _GotoState37Action,
	{_State296, ParseErrorToken}:                   _GotoState46Action,
	{_State296, VarDeclPatternType}:                _GotoState103Action,
	{_State296, VarOrLetType}:                      _GotoState24Action,
	{_State296, ExprType}:                          _GotoState165Action,
	{_State296, OptionalLabelDeclType}:             _GotoState90Action,
	{_State296, SequenceExprType}:                  _GotoState105Action,
	{_State296, CallExprType}:                      _GotoState70Action,
	{_State296, ProperArgumentsType}:               _GotoState166Action,
	{_State296, ArgumentsType}:                     _GotoState379Action,
	{_State296, ArgumentType}:                      _GotoState162Action,
	{_State296, ColonExprType}:                     _GotoState164Action,
	{_State296, AtomExprType}:                      _GotoState62Action,
	{_State296, ParseErrorExprType}:                _GotoState92Action,
	{_State296, LiteralExprType}:                   _GotoState87Action,
	{_State296, IdentifierExprType}:                _GotoState79Action,
	{_State296, BlockExprType}:                     _GotoState69Action,
	{_State296, InitializeExprType}:                _GotoState84Action,
	{_State296, ImplicitStructExprType}:            _GotoState80Action,
	{_State296, AccessibleExprType}:                _GotoState104Action,
	{_State296, AccessExprType}:                    _GotoState54Action,
	{_State296, IndexExprType}:                     _GotoState82Action,
	{_State296, PostfixableExprType}:               _GotoState94Action,
	{_State296, PostfixUnaryExprType}:              _GotoState93Action,
	{_State296, PrefixableExprType}:                _GotoState97Action,
	{_State296, PrefixUnaryExprType}:               _GotoState95Action,
	{_State296, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State296, MulExprType}:                       _GotoState89Action,
	{_State296, BinaryMulExprType}:                 _GotoState66Action,
	{_State296, AddExprType}:                       _GotoState56Action,
	{_State296, BinaryAddExprType}:                 _GotoState63Action,
	{_State296, CmpExprType}:                       _GotoState73Action,
	{_State296, BinaryCmpExprType}:                 _GotoState65Action,
	{_State296, AndExprType}:                       _GotoState57Action,
	{_State296, BinaryAndExprType}:                 _GotoState64Action,
	{_State296, OrExprType}:                        _GotoState91Action,
	{_State296, BinaryOrExprType}:                  _GotoState68Action,
	{_State296, InitializableTypeExprType}:         _GotoState83Action,
	{_State296, SliceTypeExprType}:                 _GotoState100Action,
	{_State296, ArrayTypeExprType}:                 _GotoState59Action,
	{_State296, MapTypeExprType}:                   _GotoState88Action,
	{_State296, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State296, AnonymousFuncExprType}:             _GotoState58Action,
	{_State297, MulToken}:                          _GotoState214Action,
	{_State297, DivToken}:                          _GotoState212Action,
	{_State297, ModToken}:                          _GotoState213Action,
	{_State297, BitAndToken}:                       _GotoState209Action,
	{_State297, BitLshiftToken}:                    _GotoState210Action,
	{_State297, BitRshiftToken}:                    _GotoState211Action,
	{_State297, MulOpType}:                         _GotoState215Action,
	{_State298, EqualToken}:                        _GotoState198Action,
	{_State298, NotEqualToken}:                     _GotoState203Action,
	{_State298, LessToken}:                         _GotoState201Action,
	{_State298, LessOrEqualToken}:                  _GotoState202Action,
	{_State298, GreaterToken}:                      _GotoState199Action,
	{_State298, GreaterOrEqualToken}:               _GotoState200Action,
	{_State298, CmpOpType}:                         _GotoState204Action,
	{_State300, AddToken}:                          _GotoState189Action,
	{_State300, SubToken}:                          _GotoState192Action,
	{_State300, BitXorToken}:                       _GotoState191Action,
	{_State300, BitOrToken}:                        _GotoState190Action,
	{_State300, AddOpType}:                         _GotoState193Action,
	{_State302, RparenToken}:                       _GotoState380Action,
	{_State303, CommaToken}:                        _GotoState205Action,
	{_State305, ForToken}:                          _GotoState381Action,
	{_State306, InToken}:                           _GotoState383Action,
	{_State306, AssignToken}:                       _GotoState382Action,
	{_State307, SemicolonToken}:                    _GotoState384Action,
	{_State308, DoToken}:                           _GotoState385Action,
	{_State309, IntegerLiteralToken}:               _GotoState40Action,
	{_State309, FloatLiteralToken}:                 _GotoState35Action,
	{_State309, RuneLiteralToken}:                  _GotoState48Action,
	{_State309, StringLiteralToken}:                _GotoState49Action,
	{_State309, IdentifierToken}:                   _GotoState38Action,
	{_State309, TrueToken}:                         _GotoState52Action,
	{_State309, FalseToken}:                        _GotoState34Action,
	{_State309, StructToken}:                       _GotoState50Action,
	{_State309, FuncToken}:                         _GotoState36Action,
	{_State309, VarToken}:                          _GotoState146Action,
	{_State309, LetToken}:                          _GotoState12Action,
	{_State309, NotToken}:                          _GotoState45Action,
	{_State309, LabelDeclToken}:                    _GotoState41Action,
	{_State309, LparenToken}:                       _GotoState43Action,
	{_State309, LbracketToken}:                     _GotoState42Action,
	{_State309, DotToken}:                          _GotoState145Action,
	{_State309, SubToken}:                          _GotoState51Action,
	{_State309, MulToken}:                          _GotoState44Action,
	{_State309, BitNegToken}:                       _GotoState27Action,
	{_State309, BitAndToken}:                       _GotoState26Action,
	{_State309, GreaterToken}:                      _GotoState37Action,
	{_State309, ParseErrorToken}:                   _GotoState46Action,
	{_State309, CasePatternsType}:                  _GotoState386Action,
	{_State309, VarDeclPatternType}:                _GotoState103Action,
	{_State309, VarOrLetType}:                      _GotoState24Action,
	{_State309, AssignPatternType}:                 _GotoState147Action,
	{_State309, CasePatternType}:                   _GotoState148Action,
	{_State309, OptionalLabelDeclType}:             _GotoState150Action,
	{_State309, SequenceExprType}:                  _GotoState151Action,
	{_State309, CallExprType}:                      _GotoState70Action,
	{_State309, AtomExprType}:                      _GotoState62Action,
	{_State309, ParseErrorExprType}:                _GotoState92Action,
	{_State309, LiteralExprType}:                   _GotoState87Action,
	{_State309, IdentifierExprType}:                _GotoState79Action,
	{_State309, BlockExprType}:                     _GotoState69Action,
	{_State309, InitializeExprType}:                _GotoState84Action,
	{_State309, ImplicitStructExprType}:            _GotoState80Action,
	{_State309, AccessibleExprType}:                _GotoState104Action,
	{_State309, AccessExprType}:                    _GotoState54Action,
	{_State309, IndexExprType}:                     _GotoState82Action,
	{_State309, PostfixableExprType}:               _GotoState94Action,
	{_State309, PostfixUnaryExprType}:              _GotoState93Action,
	{_State309, PrefixableExprType}:                _GotoState97Action,
	{_State309, PrefixUnaryExprType}:               _GotoState95Action,
	{_State309, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State309, MulExprType}:                       _GotoState89Action,
	{_State309, BinaryMulExprType}:                 _GotoState66Action,
	{_State309, AddExprType}:                       _GotoState56Action,
	{_State309, BinaryAddExprType}:                 _GotoState63Action,
	{_State309, CmpExprType}:                       _GotoState73Action,
	{_State309, BinaryCmpExprType}:                 _GotoState65Action,
	{_State309, AndExprType}:                       _GotoState57Action,
	{_State309, BinaryAndExprType}:                 _GotoState64Action,
	{_State309, OrExprType}:                        _GotoState91Action,
	{_State309, BinaryOrExprType}:                  _GotoState68Action,
	{_State309, InitializableTypeExprType}:         _GotoState83Action,
	{_State309, SliceTypeExprType}:                 _GotoState100Action,
	{_State309, ArrayTypeExprType}:                 _GotoState59Action,
	{_State309, MapTypeExprType}:                   _GotoState88Action,
	{_State309, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State309, AnonymousFuncExprType}:             _GotoState58Action,
	{_State310, LbraceToken}:                       _GotoState11Action,
	{_State310, StatementBlockType}:                _GotoState387Action,
	{_State312, LbraceToken}:                       _GotoState11Action,
	{_State312, StatementBlockType}:                _GotoState388Action,
	{_State313, AndToken}:                          _GotoState194Action,
	{_State314, NewlinesToken}:                     _GotoState389Action,
	{_State314, OrToken}:                           _GotoState390Action,
	{_State315, RparenToken}:                       _GotoState391Action,
	{_State316, AssignToken}:                       _GotoState328Action,
	{_State317, NewlinesToken}:                     _GotoState392Action,
	{_State317, OrToken}:                           _GotoState393Action,
	{_State318, IdentifierToken}:                   _GotoState112Action,
	{_State318, StructToken}:                       _GotoState50Action,
	{_State318, EnumToken}:                         _GotoState109Action,
	{_State318, TraitToken}:                        _GotoState117Action,
	{_State318, FuncToken}:                         _GotoState111Action,
	{_State318, LparenToken}:                       _GotoState113Action,
	{_State318, LbracketToken}:                     _GotoState42Action,
	{_State318, DotToken}:                          _GotoState108Action,
	{_State318, QuestionToken}:                     _GotoState115Action,
	{_State318, ExclaimToken}:                      _GotoState110Action,
	{_State318, TildeTildeToken}:                   _GotoState116Action,
	{_State318, BitNegToken}:                       _GotoState107Action,
	{_State318, BitAndToken}:                       _GotoState106Action,
	{_State318, ParseErrorToken}:                   _GotoState114Action,
	{_State318, InitializableTypeExprType}:         _GotoState125Action,
	{_State318, SliceTypeExprType}:                 _GotoState100Action,
	{_State318, ArrayTypeExprType}:                 _GotoState59Action,
	{_State318, MapTypeExprType}:                   _GotoState88Action,
	{_State318, AtomTypeExprType}:                  _GotoState118Action,
	{_State318, NamedTypeExprType}:                 _GotoState126Action,
	{_State318, InferredTypeExprType}:              _GotoState124Action,
	{_State318, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State318, ReturnableTypeExprType}:            _GotoState130Action,
	{_State318, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State318, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State318, TypeExprType}:                      _GotoState394Action,
	{_State318, BinaryTypeExprType}:                _GotoState119Action,
	{_State318, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State318, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State318, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State318, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State318, TraitTypeExprType}:                 _GotoState131Action,
	{_State318, FuncTypeExprType}:                  _GotoState121Action,
	{_State319, IdentifierToken}:                   _GotoState112Action,
	{_State319, StructToken}:                       _GotoState50Action,
	{_State319, EnumToken}:                         _GotoState109Action,
	{_State319, TraitToken}:                        _GotoState117Action,
	{_State319, FuncToken}:                         _GotoState111Action,
	{_State319, LparenToken}:                       _GotoState113Action,
	{_State319, LbracketToken}:                     _GotoState42Action,
	{_State319, DotToken}:                          _GotoState325Action,
	{_State319, QuestionToken}:                     _GotoState115Action,
	{_State319, ExclaimToken}:                      _GotoState110Action,
	{_State319, DollarLbracketToken}:               _GotoState178Action,
	{_State319, EllipsisToken}:                     _GotoState395Action,
	{_State319, TildeTildeToken}:                   _GotoState116Action,
	{_State319, BitNegToken}:                       _GotoState107Action,
	{_State319, BitAndToken}:                       _GotoState106Action,
	{_State319, ParseErrorToken}:                   _GotoState114Action,
	{_State319, GenericTypeArgumentsType}:          _GotoState229Action,
	{_State319, InitializableTypeExprType}:         _GotoState125Action,
	{_State319, SliceTypeExprType}:                 _GotoState100Action,
	{_State319, ArrayTypeExprType}:                 _GotoState59Action,
	{_State319, MapTypeExprType}:                   _GotoState88Action,
	{_State319, AtomTypeExprType}:                  _GotoState118Action,
	{_State319, NamedTypeExprType}:                 _GotoState126Action,
	{_State319, InferredTypeExprType}:              _GotoState124Action,
	{_State319, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State319, ReturnableTypeExprType}:            _GotoState130Action,
	{_State319, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State319, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State319, TypeExprType}:                      _GotoState396Action,
	{_State319, BinaryTypeExprType}:                _GotoState119Action,
	{_State319, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State319, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State319, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State319, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State319, TraitTypeExprType}:                 _GotoState131Action,
	{_State319, FuncTypeExprType}:                  _GotoState121Action,
	{_State321, RparenToken}:                       _GotoState397Action,
	{_State322, CommaToken}:                        _GotoState398Action,
	{_State323, AddToken}:                          _GotoState240Action,
	{_State323, SubToken}:                          _GotoState242Action,
	{_State323, MulToken}:                          _GotoState241Action,
	{_State323, BinaryTypeOpType}:                  _GotoState243Action,
	{_State324, DollarLbracketToken}:               _GotoState178Action,
	{_State324, GenericTypeArgumentsType}:          _GotoState399Action,
	{_State325, IdentifierToken}:                   _GotoState324Action,
	{_State326, AddToken}:                          _GotoState240Action,
	{_State326, SubToken}:                          _GotoState242Action,
	{_State326, MulToken}:                          _GotoState241Action,
	{_State326, BinaryTypeOpType}:                  _GotoState243Action,
	{_State327, IdentifierToken}:                   _GotoState230Action,
	{_State327, UnsafeToken}:                       _GotoState53Action,
	{_State327, StructToken}:                       _GotoState50Action,
	{_State327, EnumToken}:                         _GotoState109Action,
	{_State327, TraitToken}:                        _GotoState117Action,
	{_State327, FuncToken}:                         _GotoState111Action,
	{_State327, LparenToken}:                       _GotoState113Action,
	{_State327, LbracketToken}:                     _GotoState42Action,
	{_State327, DotToken}:                          _GotoState108Action,
	{_State327, QuestionToken}:                     _GotoState115Action,
	{_State327, ExclaimToken}:                      _GotoState110Action,
	{_State327, TildeTildeToken}:                   _GotoState116Action,
	{_State327, BitNegToken}:                       _GotoState107Action,
	{_State327, BitAndToken}:                       _GotoState106Action,
	{_State327, ParseErrorToken}:                   _GotoState114Action,
	{_State327, UnsafeStatementType}:               _GotoState237Action,
	{_State327, InitializableTypeExprType}:         _GotoState125Action,
	{_State327, SliceTypeExprType}:                 _GotoState100Action,
	{_State327, ArrayTypeExprType}:                 _GotoState59Action,
	{_State327, MapTypeExprType}:                   _GotoState88Action,
	{_State327, AtomTypeExprType}:                  _GotoState118Action,
	{_State327, NamedTypeExprType}:                 _GotoState126Action,
	{_State327, InferredTypeExprType}:              _GotoState124Action,
	{_State327, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State327, ReturnableTypeExprType}:            _GotoState130Action,
	{_State327, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State327, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State327, TypeExprType}:                      _GotoState236Action,
	{_State327, BinaryTypeExprType}:                _GotoState119Action,
	{_State327, FieldDefType}:                      _GotoState316Action,
	{_State327, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State327, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State327, EnumValueDefType}:                  _GotoState400Action,
	{_State327, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State327, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State327, TraitTypeExprType}:                 _GotoState131Action,
	{_State327, FuncTypeExprType}:                  _GotoState121Action,
	{_State328, DefaultToken}:                      _GotoState401Action,
	{_State329, IdentifierToken}:                   _GotoState230Action,
	{_State329, UnsafeToken}:                       _GotoState53Action,
	{_State329, StructToken}:                       _GotoState50Action,
	{_State329, EnumToken}:                         _GotoState109Action,
	{_State329, TraitToken}:                        _GotoState117Action,
	{_State329, FuncToken}:                         _GotoState111Action,
	{_State329, LparenToken}:                       _GotoState113Action,
	{_State329, LbracketToken}:                     _GotoState42Action,
	{_State329, DotToken}:                          _GotoState108Action,
	{_State329, QuestionToken}:                     _GotoState115Action,
	{_State329, ExclaimToken}:                      _GotoState110Action,
	{_State329, TildeTildeToken}:                   _GotoState116Action,
	{_State329, BitNegToken}:                       _GotoState107Action,
	{_State329, BitAndToken}:                       _GotoState106Action,
	{_State329, ParseErrorToken}:                   _GotoState114Action,
	{_State329, UnsafeStatementType}:               _GotoState237Action,
	{_State329, InitializableTypeExprType}:         _GotoState125Action,
	{_State329, SliceTypeExprType}:                 _GotoState100Action,
	{_State329, ArrayTypeExprType}:                 _GotoState59Action,
	{_State329, MapTypeExprType}:                   _GotoState88Action,
	{_State329, AtomTypeExprType}:                  _GotoState118Action,
	{_State329, NamedTypeExprType}:                 _GotoState126Action,
	{_State329, InferredTypeExprType}:              _GotoState124Action,
	{_State329, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State329, ReturnableTypeExprType}:            _GotoState130Action,
	{_State329, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State329, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State329, TypeExprType}:                      _GotoState236Action,
	{_State329, BinaryTypeExprType}:                _GotoState119Action,
	{_State329, FieldDefType}:                      _GotoState316Action,
	{_State329, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State329, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State329, EnumValueDefType}:                  _GotoState402Action,
	{_State329, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State329, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State329, TraitTypeExprType}:                 _GotoState131Action,
	{_State329, FuncTypeExprType}:                  _GotoState121Action,
	{_State332, IdentifierToken}:                   _GotoState230Action,
	{_State332, UnsafeToken}:                       _GotoState53Action,
	{_State332, StructToken}:                       _GotoState50Action,
	{_State332, EnumToken}:                         _GotoState109Action,
	{_State332, TraitToken}:                        _GotoState117Action,
	{_State332, FuncToken}:                         _GotoState111Action,
	{_State332, LparenToken}:                       _GotoState113Action,
	{_State332, LbracketToken}:                     _GotoState42Action,
	{_State332, DotToken}:                          _GotoState108Action,
	{_State332, QuestionToken}:                     _GotoState115Action,
	{_State332, ExclaimToken}:                      _GotoState110Action,
	{_State332, TildeTildeToken}:                   _GotoState116Action,
	{_State332, BitNegToken}:                       _GotoState107Action,
	{_State332, BitAndToken}:                       _GotoState106Action,
	{_State332, ParseErrorToken}:                   _GotoState114Action,
	{_State332, UnsafeStatementType}:               _GotoState237Action,
	{_State332, InitializableTypeExprType}:         _GotoState125Action,
	{_State332, SliceTypeExprType}:                 _GotoState100Action,
	{_State332, ArrayTypeExprType}:                 _GotoState59Action,
	{_State332, MapTypeExprType}:                   _GotoState88Action,
	{_State332, AtomTypeExprType}:                  _GotoState118Action,
	{_State332, NamedTypeExprType}:                 _GotoState126Action,
	{_State332, InferredTypeExprType}:              _GotoState124Action,
	{_State332, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State332, ReturnableTypeExprType}:            _GotoState130Action,
	{_State332, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State332, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State332, TypeExprType}:                      _GotoState236Action,
	{_State332, BinaryTypeExprType}:                _GotoState119Action,
	{_State332, FieldDefType}:                      _GotoState403Action,
	{_State332, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State332, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State332, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State332, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State332, TraitTypeExprType}:                 _GotoState131Action,
	{_State332, FuncTypeExprType}:                  _GotoState121Action,
	{_State333, IdentifierToken}:                   _GotoState404Action,
	{_State333, LparenToken}:                       _GotoState227Action,
	{_State336, NewlinesToken}:                     _GotoState406Action,
	{_State336, CommaToken}:                        _GotoState405Action,
	{_State337, RparenToken}:                       _GotoState407Action,
	{_State341, IdentifierToken}:                   _GotoState112Action,
	{_State341, StructToken}:                       _GotoState50Action,
	{_State341, EnumToken}:                         _GotoState109Action,
	{_State341, TraitToken}:                        _GotoState117Action,
	{_State341, FuncToken}:                         _GotoState111Action,
	{_State341, LparenToken}:                       _GotoState113Action,
	{_State341, LbracketToken}:                     _GotoState42Action,
	{_State341, DotToken}:                          _GotoState108Action,
	{_State341, QuestionToken}:                     _GotoState115Action,
	{_State341, ExclaimToken}:                      _GotoState110Action,
	{_State341, TildeTildeToken}:                   _GotoState116Action,
	{_State341, BitNegToken}:                       _GotoState107Action,
	{_State341, BitAndToken}:                       _GotoState106Action,
	{_State341, ParseErrorToken}:                   _GotoState114Action,
	{_State341, InitializableTypeExprType}:         _GotoState125Action,
	{_State341, SliceTypeExprType}:                 _GotoState100Action,
	{_State341, ArrayTypeExprType}:                 _GotoState59Action,
	{_State341, MapTypeExprType}:                   _GotoState88Action,
	{_State341, AtomTypeExprType}:                  _GotoState118Action,
	{_State341, NamedTypeExprType}:                 _GotoState126Action,
	{_State341, InferredTypeExprType}:              _GotoState124Action,
	{_State341, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State341, ReturnableTypeExprType}:            _GotoState130Action,
	{_State341, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State341, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State341, TypeExprType}:                      _GotoState408Action,
	{_State341, BinaryTypeExprType}:                _GotoState119Action,
	{_State341, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State341, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State341, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State341, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State341, TraitTypeExprType}:                 _GotoState131Action,
	{_State341, FuncTypeExprType}:                  _GotoState121Action,
	{_State343, RbracketToken}:                     _GotoState409Action,
	{_State344, CommaToken}:                        _GotoState410Action,
	{_State345, IdentifierToken}:                   _GotoState247Action,
	{_State345, ParameterDefType}:                  _GotoState269Action,
	{_State345, ProperParameterDefsType}:           _GotoState271Action,
	{_State345, ParameterDefsType}:                 _GotoState411Action,
	{_State346, IdentifierToken}:                   _GotoState112Action,
	{_State346, StructToken}:                       _GotoState50Action,
	{_State346, EnumToken}:                         _GotoState109Action,
	{_State346, TraitToken}:                        _GotoState117Action,
	{_State346, FuncToken}:                         _GotoState111Action,
	{_State346, LparenToken}:                       _GotoState113Action,
	{_State346, LbracketToken}:                     _GotoState42Action,
	{_State346, DotToken}:                          _GotoState108Action,
	{_State346, QuestionToken}:                     _GotoState115Action,
	{_State346, ExclaimToken}:                      _GotoState110Action,
	{_State346, TildeTildeToken}:                   _GotoState116Action,
	{_State346, BitNegToken}:                       _GotoState107Action,
	{_State346, BitAndToken}:                       _GotoState106Action,
	{_State346, ParseErrorToken}:                   _GotoState114Action,
	{_State346, InitializableTypeExprType}:         _GotoState125Action,
	{_State346, SliceTypeExprType}:                 _GotoState100Action,
	{_State346, ArrayTypeExprType}:                 _GotoState59Action,
	{_State346, MapTypeExprType}:                   _GotoState88Action,
	{_State346, AtomTypeExprType}:                  _GotoState118Action,
	{_State346, NamedTypeExprType}:                 _GotoState126Action,
	{_State346, InferredTypeExprType}:              _GotoState124Action,
	{_State346, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State346, ReturnableTypeExprType}:            _GotoState130Action,
	{_State346, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State346, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State346, TypeExprType}:                      _GotoState412Action,
	{_State346, BinaryTypeExprType}:                _GotoState119Action,
	{_State346, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State346, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State346, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State346, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State346, TraitTypeExprType}:                 _GotoState131Action,
	{_State346, FuncTypeExprType}:                  _GotoState121Action,
	{_State347, AddToken}:                          _GotoState240Action,
	{_State347, SubToken}:                          _GotoState242Action,
	{_State347, MulToken}:                          _GotoState241Action,
	{_State347, BinaryTypeOpType}:                  _GotoState243Action,
	{_State348, IdentifierToken}:                   _GotoState413Action,
	{_State351, AddToken}:                          _GotoState240Action,
	{_State351, SubToken}:                          _GotoState242Action,
	{_State351, MulToken}:                          _GotoState241Action,
	{_State351, BinaryTypeOpType}:                  _GotoState243Action,
	{_State352, ImplementsToken}:                   _GotoState414Action,
	{_State352, AddToken}:                          _GotoState240Action,
	{_State352, SubToken}:                          _GotoState242Action,
	{_State352, MulToken}:                          _GotoState241Action,
	{_State352, BinaryTypeOpType}:                  _GotoState243Action,
	{_State353, IdentifierToken}:                   _GotoState141Action,
	{_State353, LparenToken}:                       _GotoState142Action,
	{_State353, VarPatternType}:                    _GotoState415Action,
	{_State353, TuplePatternType}:                  _GotoState143Action,
	{_State354, IdentifierToken}:                   _GotoState257Action,
	{_State354, LparenToken}:                       _GotoState142Action,
	{_State354, EllipsisToken}:                     _GotoState256Action,
	{_State354, VarPatternType}:                    _GotoState260Action,
	{_State354, TuplePatternType}:                  _GotoState143Action,
	{_State354, FieldVarPatternType}:               _GotoState416Action,
	{_State357, LparenToken}:                       _GotoState142Action,
	{_State357, TuplePatternType}:                  _GotoState417Action,
	{_State360, IdentifierToken}:                   _GotoState112Action,
	{_State360, StructToken}:                       _GotoState50Action,
	{_State360, EnumToken}:                         _GotoState109Action,
	{_State360, TraitToken}:                        _GotoState117Action,
	{_State360, FuncToken}:                         _GotoState111Action,
	{_State360, LparenToken}:                       _GotoState113Action,
	{_State360, LbracketToken}:                     _GotoState42Action,
	{_State360, DotToken}:                          _GotoState108Action,
	{_State360, QuestionToken}:                     _GotoState115Action,
	{_State360, ExclaimToken}:                      _GotoState110Action,
	{_State360, TildeTildeToken}:                   _GotoState116Action,
	{_State360, BitNegToken}:                       _GotoState107Action,
	{_State360, BitAndToken}:                       _GotoState106Action,
	{_State360, ParseErrorToken}:                   _GotoState114Action,
	{_State360, InitializableTypeExprType}:         _GotoState125Action,
	{_State360, SliceTypeExprType}:                 _GotoState100Action,
	{_State360, ArrayTypeExprType}:                 _GotoState59Action,
	{_State360, MapTypeExprType}:                   _GotoState88Action,
	{_State360, AtomTypeExprType}:                  _GotoState118Action,
	{_State360, NamedTypeExprType}:                 _GotoState126Action,
	{_State360, InferredTypeExprType}:              _GotoState124Action,
	{_State360, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State360, ReturnableTypeExprType}:            _GotoState419Action,
	{_State360, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State360, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State360, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State360, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State360, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State360, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State360, TraitTypeExprType}:                 _GotoState131Action,
	{_State360, ReturnTypeType}:                    _GotoState418Action,
	{_State360, FuncTypeExprType}:                  _GotoState121Action,
	{_State361, IdentifierToken}:                   _GotoState247Action,
	{_State361, ParameterDefType}:                  _GotoState420Action,
	{_State363, StringLiteralToken}:                _GotoState156Action,
	{_State363, ImportClauseType}:                  _GotoState421Action,
	{_State364, StringLiteralToken}:                _GotoState156Action,
	{_State364, ImportClauseType}:                  _GotoState422Action,
	{_State366, RbracketToken}:                     _GotoState423Action,
	{_State366, AddToken}:                          _GotoState240Action,
	{_State366, SubToken}:                          _GotoState242Action,
	{_State366, MulToken}:                          _GotoState241Action,
	{_State366, BinaryTypeOpType}:                  _GotoState243Action,
	{_State367, RbracketToken}:                     _GotoState424Action,
	{_State373, IdentifierToken}:                   _GotoState230Action,
	{_State373, UnsafeToken}:                       _GotoState53Action,
	{_State373, StructToken}:                       _GotoState50Action,
	{_State373, EnumToken}:                         _GotoState109Action,
	{_State373, TraitToken}:                        _GotoState117Action,
	{_State373, FuncToken}:                         _GotoState111Action,
	{_State373, LparenToken}:                       _GotoState113Action,
	{_State373, LbracketToken}:                     _GotoState42Action,
	{_State373, DotToken}:                          _GotoState108Action,
	{_State373, QuestionToken}:                     _GotoState115Action,
	{_State373, ExclaimToken}:                      _GotoState110Action,
	{_State373, TildeTildeToken}:                   _GotoState116Action,
	{_State373, BitNegToken}:                       _GotoState107Action,
	{_State373, BitAndToken}:                       _GotoState106Action,
	{_State373, ParseErrorToken}:                   _GotoState114Action,
	{_State373, UnsafeStatementType}:               _GotoState237Action,
	{_State373, InitializableTypeExprType}:         _GotoState125Action,
	{_State373, SliceTypeExprType}:                 _GotoState100Action,
	{_State373, ArrayTypeExprType}:                 _GotoState59Action,
	{_State373, MapTypeExprType}:                   _GotoState88Action,
	{_State373, AtomTypeExprType}:                  _GotoState118Action,
	{_State373, NamedTypeExprType}:                 _GotoState126Action,
	{_State373, InferredTypeExprType}:              _GotoState124Action,
	{_State373, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State373, ReturnableTypeExprType}:            _GotoState130Action,
	{_State373, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State373, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State373, TypeExprType}:                      _GotoState236Action,
	{_State373, BinaryTypeExprType}:                _GotoState119Action,
	{_State373, FieldDefType}:                      _GotoState425Action,
	{_State373, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State373, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State373, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State373, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State373, TraitTypeExprType}:                 _GotoState131Action,
	{_State373, FuncTypeExprType}:                  _GotoState121Action,
	{_State374, IdentifierToken}:                   _GotoState230Action,
	{_State374, UnsafeToken}:                       _GotoState53Action,
	{_State374, StructToken}:                       _GotoState50Action,
	{_State374, EnumToken}:                         _GotoState109Action,
	{_State374, TraitToken}:                        _GotoState117Action,
	{_State374, FuncToken}:                         _GotoState111Action,
	{_State374, LparenToken}:                       _GotoState113Action,
	{_State374, LbracketToken}:                     _GotoState42Action,
	{_State374, DotToken}:                          _GotoState108Action,
	{_State374, QuestionToken}:                     _GotoState115Action,
	{_State374, ExclaimToken}:                      _GotoState110Action,
	{_State374, TildeTildeToken}:                   _GotoState116Action,
	{_State374, BitNegToken}:                       _GotoState107Action,
	{_State374, BitAndToken}:                       _GotoState106Action,
	{_State374, ParseErrorToken}:                   _GotoState114Action,
	{_State374, UnsafeStatementType}:               _GotoState237Action,
	{_State374, InitializableTypeExprType}:         _GotoState125Action,
	{_State374, SliceTypeExprType}:                 _GotoState100Action,
	{_State374, ArrayTypeExprType}:                 _GotoState59Action,
	{_State374, MapTypeExprType}:                   _GotoState88Action,
	{_State374, AtomTypeExprType}:                  _GotoState118Action,
	{_State374, NamedTypeExprType}:                 _GotoState126Action,
	{_State374, InferredTypeExprType}:              _GotoState124Action,
	{_State374, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State374, ReturnableTypeExprType}:            _GotoState130Action,
	{_State374, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State374, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State374, TypeExprType}:                      _GotoState236Action,
	{_State374, BinaryTypeExprType}:                _GotoState119Action,
	{_State374, FieldDefType}:                      _GotoState426Action,
	{_State374, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State374, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State374, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State374, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State374, TraitTypeExprType}:                 _GotoState131Action,
	{_State374, FuncTypeExprType}:                  _GotoState121Action,
	{_State375, StringLiteralToken}:                _GotoState427Action,
	{_State376, IdentifierToken}:                   _GotoState112Action,
	{_State376, StructToken}:                       _GotoState50Action,
	{_State376, EnumToken}:                         _GotoState109Action,
	{_State376, TraitToken}:                        _GotoState117Action,
	{_State376, FuncToken}:                         _GotoState111Action,
	{_State376, LparenToken}:                       _GotoState113Action,
	{_State376, LbracketToken}:                     _GotoState42Action,
	{_State376, DotToken}:                          _GotoState108Action,
	{_State376, QuestionToken}:                     _GotoState115Action,
	{_State376, ExclaimToken}:                      _GotoState110Action,
	{_State376, TildeTildeToken}:                   _GotoState116Action,
	{_State376, BitNegToken}:                       _GotoState107Action,
	{_State376, BitAndToken}:                       _GotoState106Action,
	{_State376, ParseErrorToken}:                   _GotoState114Action,
	{_State376, InitializableTypeExprType}:         _GotoState125Action,
	{_State376, SliceTypeExprType}:                 _GotoState100Action,
	{_State376, ArrayTypeExprType}:                 _GotoState59Action,
	{_State376, MapTypeExprType}:                   _GotoState88Action,
	{_State376, AtomTypeExprType}:                  _GotoState118Action,
	{_State376, NamedTypeExprType}:                 _GotoState126Action,
	{_State376, InferredTypeExprType}:              _GotoState124Action,
	{_State376, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State376, ReturnableTypeExprType}:            _GotoState130Action,
	{_State376, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State376, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State376, TypeExprType}:                      _GotoState428Action,
	{_State376, BinaryTypeExprType}:                _GotoState119Action,
	{_State376, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State376, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State376, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State376, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State376, TraitTypeExprType}:                 _GotoState131Action,
	{_State376, FuncTypeExprType}:                  _GotoState121Action,
	{_State379, RparenToken}:                       _GotoState429Action,
	{_State381, IntegerLiteralToken}:               _GotoState40Action,
	{_State381, FloatLiteralToken}:                 _GotoState35Action,
	{_State381, RuneLiteralToken}:                  _GotoState48Action,
	{_State381, StringLiteralToken}:                _GotoState49Action,
	{_State381, IdentifierToken}:                   _GotoState38Action,
	{_State381, TrueToken}:                         _GotoState52Action,
	{_State381, FalseToken}:                        _GotoState34Action,
	{_State381, StructToken}:                       _GotoState50Action,
	{_State381, FuncToken}:                         _GotoState36Action,
	{_State381, VarToken}:                          _GotoState15Action,
	{_State381, LetToken}:                          _GotoState12Action,
	{_State381, NotToken}:                          _GotoState45Action,
	{_State381, LabelDeclToken}:                    _GotoState41Action,
	{_State381, LparenToken}:                       _GotoState43Action,
	{_State381, LbracketToken}:                     _GotoState42Action,
	{_State381, SubToken}:                          _GotoState51Action,
	{_State381, MulToken}:                          _GotoState44Action,
	{_State381, BitNegToken}:                       _GotoState27Action,
	{_State381, BitAndToken}:                       _GotoState26Action,
	{_State381, GreaterToken}:                      _GotoState37Action,
	{_State381, ParseErrorToken}:                   _GotoState46Action,
	{_State381, VarDeclPatternType}:                _GotoState103Action,
	{_State381, VarOrLetType}:                      _GotoState24Action,
	{_State381, OptionalLabelDeclType}:             _GotoState150Action,
	{_State381, SequenceExprType}:                  _GotoState430Action,
	{_State381, CallExprType}:                      _GotoState70Action,
	{_State381, AtomExprType}:                      _GotoState62Action,
	{_State381, ParseErrorExprType}:                _GotoState92Action,
	{_State381, LiteralExprType}:                   _GotoState87Action,
	{_State381, IdentifierExprType}:                _GotoState79Action,
	{_State381, BlockExprType}:                     _GotoState69Action,
	{_State381, InitializeExprType}:                _GotoState84Action,
	{_State381, ImplicitStructExprType}:            _GotoState80Action,
	{_State381, AccessibleExprType}:                _GotoState104Action,
	{_State381, AccessExprType}:                    _GotoState54Action,
	{_State381, IndexExprType}:                     _GotoState82Action,
	{_State381, PostfixableExprType}:               _GotoState94Action,
	{_State381, PostfixUnaryExprType}:              _GotoState93Action,
	{_State381, PrefixableExprType}:                _GotoState97Action,
	{_State381, PrefixUnaryExprType}:               _GotoState95Action,
	{_State381, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State381, MulExprType}:                       _GotoState89Action,
	{_State381, BinaryMulExprType}:                 _GotoState66Action,
	{_State381, AddExprType}:                       _GotoState56Action,
	{_State381, BinaryAddExprType}:                 _GotoState63Action,
	{_State381, CmpExprType}:                       _GotoState73Action,
	{_State381, BinaryCmpExprType}:                 _GotoState65Action,
	{_State381, AndExprType}:                       _GotoState57Action,
	{_State381, BinaryAndExprType}:                 _GotoState64Action,
	{_State381, OrExprType}:                        _GotoState91Action,
	{_State381, BinaryOrExprType}:                  _GotoState68Action,
	{_State381, InitializableTypeExprType}:         _GotoState83Action,
	{_State381, SliceTypeExprType}:                 _GotoState100Action,
	{_State381, ArrayTypeExprType}:                 _GotoState59Action,
	{_State381, MapTypeExprType}:                   _GotoState88Action,
	{_State381, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State381, AnonymousFuncExprType}:             _GotoState58Action,
	{_State382, IntegerLiteralToken}:               _GotoState40Action,
	{_State382, FloatLiteralToken}:                 _GotoState35Action,
	{_State382, RuneLiteralToken}:                  _GotoState48Action,
	{_State382, StringLiteralToken}:                _GotoState49Action,
	{_State382, IdentifierToken}:                   _GotoState38Action,
	{_State382, TrueToken}:                         _GotoState52Action,
	{_State382, FalseToken}:                        _GotoState34Action,
	{_State382, StructToken}:                       _GotoState50Action,
	{_State382, FuncToken}:                         _GotoState36Action,
	{_State382, VarToken}:                          _GotoState15Action,
	{_State382, LetToken}:                          _GotoState12Action,
	{_State382, NotToken}:                          _GotoState45Action,
	{_State382, LabelDeclToken}:                    _GotoState41Action,
	{_State382, LparenToken}:                       _GotoState43Action,
	{_State382, LbracketToken}:                     _GotoState42Action,
	{_State382, SubToken}:                          _GotoState51Action,
	{_State382, MulToken}:                          _GotoState44Action,
	{_State382, BitNegToken}:                       _GotoState27Action,
	{_State382, BitAndToken}:                       _GotoState26Action,
	{_State382, GreaterToken}:                      _GotoState37Action,
	{_State382, ParseErrorToken}:                   _GotoState46Action,
	{_State382, VarDeclPatternType}:                _GotoState103Action,
	{_State382, VarOrLetType}:                      _GotoState24Action,
	{_State382, OptionalLabelDeclType}:             _GotoState150Action,
	{_State382, SequenceExprType}:                  _GotoState431Action,
	{_State382, CallExprType}:                      _GotoState70Action,
	{_State382, AtomExprType}:                      _GotoState62Action,
	{_State382, ParseErrorExprType}:                _GotoState92Action,
	{_State382, LiteralExprType}:                   _GotoState87Action,
	{_State382, IdentifierExprType}:                _GotoState79Action,
	{_State382, BlockExprType}:                     _GotoState69Action,
	{_State382, InitializeExprType}:                _GotoState84Action,
	{_State382, ImplicitStructExprType}:            _GotoState80Action,
	{_State382, AccessibleExprType}:                _GotoState104Action,
	{_State382, AccessExprType}:                    _GotoState54Action,
	{_State382, IndexExprType}:                     _GotoState82Action,
	{_State382, PostfixableExprType}:               _GotoState94Action,
	{_State382, PostfixUnaryExprType}:              _GotoState93Action,
	{_State382, PrefixableExprType}:                _GotoState97Action,
	{_State382, PrefixUnaryExprType}:               _GotoState95Action,
	{_State382, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State382, MulExprType}:                       _GotoState89Action,
	{_State382, BinaryMulExprType}:                 _GotoState66Action,
	{_State382, AddExprType}:                       _GotoState56Action,
	{_State382, BinaryAddExprType}:                 _GotoState63Action,
	{_State382, CmpExprType}:                       _GotoState73Action,
	{_State382, BinaryCmpExprType}:                 _GotoState65Action,
	{_State382, AndExprType}:                       _GotoState57Action,
	{_State382, BinaryAndExprType}:                 _GotoState64Action,
	{_State382, OrExprType}:                        _GotoState91Action,
	{_State382, BinaryOrExprType}:                  _GotoState68Action,
	{_State382, InitializableTypeExprType}:         _GotoState83Action,
	{_State382, SliceTypeExprType}:                 _GotoState100Action,
	{_State382, ArrayTypeExprType}:                 _GotoState59Action,
	{_State382, MapTypeExprType}:                   _GotoState88Action,
	{_State382, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State382, AnonymousFuncExprType}:             _GotoState58Action,
	{_State383, IntegerLiteralToken}:               _GotoState40Action,
	{_State383, FloatLiteralToken}:                 _GotoState35Action,
	{_State383, RuneLiteralToken}:                  _GotoState48Action,
	{_State383, StringLiteralToken}:                _GotoState49Action,
	{_State383, IdentifierToken}:                   _GotoState38Action,
	{_State383, TrueToken}:                         _GotoState52Action,
	{_State383, FalseToken}:                        _GotoState34Action,
	{_State383, StructToken}:                       _GotoState50Action,
	{_State383, FuncToken}:                         _GotoState36Action,
	{_State383, VarToken}:                          _GotoState15Action,
	{_State383, LetToken}:                          _GotoState12Action,
	{_State383, NotToken}:                          _GotoState45Action,
	{_State383, LabelDeclToken}:                    _GotoState41Action,
	{_State383, LparenToken}:                       _GotoState43Action,
	{_State383, LbracketToken}:                     _GotoState42Action,
	{_State383, SubToken}:                          _GotoState51Action,
	{_State383, MulToken}:                          _GotoState44Action,
	{_State383, BitNegToken}:                       _GotoState27Action,
	{_State383, BitAndToken}:                       _GotoState26Action,
	{_State383, GreaterToken}:                      _GotoState37Action,
	{_State383, ParseErrorToken}:                   _GotoState46Action,
	{_State383, VarDeclPatternType}:                _GotoState103Action,
	{_State383, VarOrLetType}:                      _GotoState24Action,
	{_State383, OptionalLabelDeclType}:             _GotoState150Action,
	{_State383, SequenceExprType}:                  _GotoState432Action,
	{_State383, CallExprType}:                      _GotoState70Action,
	{_State383, AtomExprType}:                      _GotoState62Action,
	{_State383, ParseErrorExprType}:                _GotoState92Action,
	{_State383, LiteralExprType}:                   _GotoState87Action,
	{_State383, IdentifierExprType}:                _GotoState79Action,
	{_State383, BlockExprType}:                     _GotoState69Action,
	{_State383, InitializeExprType}:                _GotoState84Action,
	{_State383, ImplicitStructExprType}:            _GotoState80Action,
	{_State383, AccessibleExprType}:                _GotoState104Action,
	{_State383, AccessExprType}:                    _GotoState54Action,
	{_State383, IndexExprType}:                     _GotoState82Action,
	{_State383, PostfixableExprType}:               _GotoState94Action,
	{_State383, PostfixUnaryExprType}:              _GotoState93Action,
	{_State383, PrefixableExprType}:                _GotoState97Action,
	{_State383, PrefixUnaryExprType}:               _GotoState95Action,
	{_State383, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State383, MulExprType}:                       _GotoState89Action,
	{_State383, BinaryMulExprType}:                 _GotoState66Action,
	{_State383, AddExprType}:                       _GotoState56Action,
	{_State383, BinaryAddExprType}:                 _GotoState63Action,
	{_State383, CmpExprType}:                       _GotoState73Action,
	{_State383, BinaryCmpExprType}:                 _GotoState65Action,
	{_State383, AndExprType}:                       _GotoState57Action,
	{_State383, BinaryAndExprType}:                 _GotoState64Action,
	{_State383, OrExprType}:                        _GotoState91Action,
	{_State383, BinaryOrExprType}:                  _GotoState68Action,
	{_State383, InitializableTypeExprType}:         _GotoState83Action,
	{_State383, SliceTypeExprType}:                 _GotoState100Action,
	{_State383, ArrayTypeExprType}:                 _GotoState59Action,
	{_State383, MapTypeExprType}:                   _GotoState88Action,
	{_State383, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State383, AnonymousFuncExprType}:             _GotoState58Action,
	{_State384, IntegerLiteralToken}:               _GotoState40Action,
	{_State384, FloatLiteralToken}:                 _GotoState35Action,
	{_State384, RuneLiteralToken}:                  _GotoState48Action,
	{_State384, StringLiteralToken}:                _GotoState49Action,
	{_State384, IdentifierToken}:                   _GotoState38Action,
	{_State384, TrueToken}:                         _GotoState52Action,
	{_State384, FalseToken}:                        _GotoState34Action,
	{_State384, StructToken}:                       _GotoState50Action,
	{_State384, FuncToken}:                         _GotoState36Action,
	{_State384, VarToken}:                          _GotoState15Action,
	{_State384, LetToken}:                          _GotoState12Action,
	{_State384, NotToken}:                          _GotoState45Action,
	{_State384, LabelDeclToken}:                    _GotoState41Action,
	{_State384, LparenToken}:                       _GotoState43Action,
	{_State384, LbracketToken}:                     _GotoState42Action,
	{_State384, SubToken}:                          _GotoState51Action,
	{_State384, MulToken}:                          _GotoState44Action,
	{_State384, BitNegToken}:                       _GotoState27Action,
	{_State384, BitAndToken}:                       _GotoState26Action,
	{_State384, GreaterToken}:                      _GotoState37Action,
	{_State384, ParseErrorToken}:                   _GotoState46Action,
	{_State384, VarDeclPatternType}:                _GotoState103Action,
	{_State384, VarOrLetType}:                      _GotoState24Action,
	{_State384, OptionalLabelDeclType}:             _GotoState150Action,
	{_State384, SequenceExprType}:                  _GotoState434Action,
	{_State384, OptionalSequenceExprType}:          _GotoState433Action,
	{_State384, CallExprType}:                      _GotoState70Action,
	{_State384, AtomExprType}:                      _GotoState62Action,
	{_State384, ParseErrorExprType}:                _GotoState92Action,
	{_State384, LiteralExprType}:                   _GotoState87Action,
	{_State384, IdentifierExprType}:                _GotoState79Action,
	{_State384, BlockExprType}:                     _GotoState69Action,
	{_State384, InitializeExprType}:                _GotoState84Action,
	{_State384, ImplicitStructExprType}:            _GotoState80Action,
	{_State384, AccessibleExprType}:                _GotoState104Action,
	{_State384, AccessExprType}:                    _GotoState54Action,
	{_State384, IndexExprType}:                     _GotoState82Action,
	{_State384, PostfixableExprType}:               _GotoState94Action,
	{_State384, PostfixUnaryExprType}:              _GotoState93Action,
	{_State384, PrefixableExprType}:                _GotoState97Action,
	{_State384, PrefixUnaryExprType}:               _GotoState95Action,
	{_State384, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State384, MulExprType}:                       _GotoState89Action,
	{_State384, BinaryMulExprType}:                 _GotoState66Action,
	{_State384, AddExprType}:                       _GotoState56Action,
	{_State384, BinaryAddExprType}:                 _GotoState63Action,
	{_State384, CmpExprType}:                       _GotoState73Action,
	{_State384, BinaryCmpExprType}:                 _GotoState65Action,
	{_State384, AndExprType}:                       _GotoState57Action,
	{_State384, BinaryAndExprType}:                 _GotoState64Action,
	{_State384, OrExprType}:                        _GotoState91Action,
	{_State384, BinaryOrExprType}:                  _GotoState68Action,
	{_State384, InitializableTypeExprType}:         _GotoState83Action,
	{_State384, SliceTypeExprType}:                 _GotoState100Action,
	{_State384, ArrayTypeExprType}:                 _GotoState59Action,
	{_State384, MapTypeExprType}:                   _GotoState88Action,
	{_State384, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State384, AnonymousFuncExprType}:             _GotoState58Action,
	{_State385, LbraceToken}:                       _GotoState11Action,
	{_State385, StatementBlockType}:                _GotoState435Action,
	{_State386, CommaToken}:                        _GotoState266Action,
	{_State386, AssignToken}:                       _GotoState436Action,
	{_State387, ElseToken}:                         _GotoState437Action,
	{_State389, IdentifierToken}:                   _GotoState230Action,
	{_State389, UnsafeToken}:                       _GotoState53Action,
	{_State389, StructToken}:                       _GotoState50Action,
	{_State389, EnumToken}:                         _GotoState109Action,
	{_State389, TraitToken}:                        _GotoState117Action,
	{_State389, FuncToken}:                         _GotoState111Action,
	{_State389, LparenToken}:                       _GotoState113Action,
	{_State389, LbracketToken}:                     _GotoState42Action,
	{_State389, DotToken}:                          _GotoState108Action,
	{_State389, QuestionToken}:                     _GotoState115Action,
	{_State389, ExclaimToken}:                      _GotoState110Action,
	{_State389, TildeTildeToken}:                   _GotoState116Action,
	{_State389, BitNegToken}:                       _GotoState107Action,
	{_State389, BitAndToken}:                       _GotoState106Action,
	{_State389, ParseErrorToken}:                   _GotoState114Action,
	{_State389, UnsafeStatementType}:               _GotoState237Action,
	{_State389, InitializableTypeExprType}:         _GotoState125Action,
	{_State389, SliceTypeExprType}:                 _GotoState100Action,
	{_State389, ArrayTypeExprType}:                 _GotoState59Action,
	{_State389, MapTypeExprType}:                   _GotoState88Action,
	{_State389, AtomTypeExprType}:                  _GotoState118Action,
	{_State389, NamedTypeExprType}:                 _GotoState126Action,
	{_State389, InferredTypeExprType}:              _GotoState124Action,
	{_State389, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State389, ReturnableTypeExprType}:            _GotoState130Action,
	{_State389, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State389, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State389, TypeExprType}:                      _GotoState236Action,
	{_State389, BinaryTypeExprType}:                _GotoState119Action,
	{_State389, FieldDefType}:                      _GotoState316Action,
	{_State389, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State389, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State389, EnumValueDefType}:                  _GotoState438Action,
	{_State389, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State389, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State389, TraitTypeExprType}:                 _GotoState131Action,
	{_State389, FuncTypeExprType}:                  _GotoState121Action,
	{_State390, IdentifierToken}:                   _GotoState230Action,
	{_State390, UnsafeToken}:                       _GotoState53Action,
	{_State390, StructToken}:                       _GotoState50Action,
	{_State390, EnumToken}:                         _GotoState109Action,
	{_State390, TraitToken}:                        _GotoState117Action,
	{_State390, FuncToken}:                         _GotoState111Action,
	{_State390, LparenToken}:                       _GotoState113Action,
	{_State390, LbracketToken}:                     _GotoState42Action,
	{_State390, DotToken}:                          _GotoState108Action,
	{_State390, QuestionToken}:                     _GotoState115Action,
	{_State390, ExclaimToken}:                      _GotoState110Action,
	{_State390, TildeTildeToken}:                   _GotoState116Action,
	{_State390, BitNegToken}:                       _GotoState107Action,
	{_State390, BitAndToken}:                       _GotoState106Action,
	{_State390, ParseErrorToken}:                   _GotoState114Action,
	{_State390, UnsafeStatementType}:               _GotoState237Action,
	{_State390, InitializableTypeExprType}:         _GotoState125Action,
	{_State390, SliceTypeExprType}:                 _GotoState100Action,
	{_State390, ArrayTypeExprType}:                 _GotoState59Action,
	{_State390, MapTypeExprType}:                   _GotoState88Action,
	{_State390, AtomTypeExprType}:                  _GotoState118Action,
	{_State390, NamedTypeExprType}:                 _GotoState126Action,
	{_State390, InferredTypeExprType}:              _GotoState124Action,
	{_State390, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State390, ReturnableTypeExprType}:            _GotoState130Action,
	{_State390, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State390, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State390, TypeExprType}:                      _GotoState236Action,
	{_State390, BinaryTypeExprType}:                _GotoState119Action,
	{_State390, FieldDefType}:                      _GotoState316Action,
	{_State390, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State390, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State390, EnumValueDefType}:                  _GotoState439Action,
	{_State390, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State390, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State390, TraitTypeExprType}:                 _GotoState131Action,
	{_State390, FuncTypeExprType}:                  _GotoState121Action,
	{_State392, IdentifierToken}:                   _GotoState230Action,
	{_State392, UnsafeToken}:                       _GotoState53Action,
	{_State392, StructToken}:                       _GotoState50Action,
	{_State392, EnumToken}:                         _GotoState109Action,
	{_State392, TraitToken}:                        _GotoState117Action,
	{_State392, FuncToken}:                         _GotoState111Action,
	{_State392, LparenToken}:                       _GotoState113Action,
	{_State392, LbracketToken}:                     _GotoState42Action,
	{_State392, DotToken}:                          _GotoState108Action,
	{_State392, QuestionToken}:                     _GotoState115Action,
	{_State392, ExclaimToken}:                      _GotoState110Action,
	{_State392, TildeTildeToken}:                   _GotoState116Action,
	{_State392, BitNegToken}:                       _GotoState107Action,
	{_State392, BitAndToken}:                       _GotoState106Action,
	{_State392, ParseErrorToken}:                   _GotoState114Action,
	{_State392, UnsafeStatementType}:               _GotoState237Action,
	{_State392, InitializableTypeExprType}:         _GotoState125Action,
	{_State392, SliceTypeExprType}:                 _GotoState100Action,
	{_State392, ArrayTypeExprType}:                 _GotoState59Action,
	{_State392, MapTypeExprType}:                   _GotoState88Action,
	{_State392, AtomTypeExprType}:                  _GotoState118Action,
	{_State392, NamedTypeExprType}:                 _GotoState126Action,
	{_State392, InferredTypeExprType}:              _GotoState124Action,
	{_State392, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State392, ReturnableTypeExprType}:            _GotoState130Action,
	{_State392, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State392, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State392, TypeExprType}:                      _GotoState236Action,
	{_State392, BinaryTypeExprType}:                _GotoState119Action,
	{_State392, FieldDefType}:                      _GotoState316Action,
	{_State392, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State392, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State392, EnumValueDefType}:                  _GotoState440Action,
	{_State392, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State392, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State392, TraitTypeExprType}:                 _GotoState131Action,
	{_State392, FuncTypeExprType}:                  _GotoState121Action,
	{_State393, IdentifierToken}:                   _GotoState230Action,
	{_State393, UnsafeToken}:                       _GotoState53Action,
	{_State393, StructToken}:                       _GotoState50Action,
	{_State393, EnumToken}:                         _GotoState109Action,
	{_State393, TraitToken}:                        _GotoState117Action,
	{_State393, FuncToken}:                         _GotoState111Action,
	{_State393, LparenToken}:                       _GotoState113Action,
	{_State393, LbracketToken}:                     _GotoState42Action,
	{_State393, DotToken}:                          _GotoState108Action,
	{_State393, QuestionToken}:                     _GotoState115Action,
	{_State393, ExclaimToken}:                      _GotoState110Action,
	{_State393, TildeTildeToken}:                   _GotoState116Action,
	{_State393, BitNegToken}:                       _GotoState107Action,
	{_State393, BitAndToken}:                       _GotoState106Action,
	{_State393, ParseErrorToken}:                   _GotoState114Action,
	{_State393, UnsafeStatementType}:               _GotoState237Action,
	{_State393, InitializableTypeExprType}:         _GotoState125Action,
	{_State393, SliceTypeExprType}:                 _GotoState100Action,
	{_State393, ArrayTypeExprType}:                 _GotoState59Action,
	{_State393, MapTypeExprType}:                   _GotoState88Action,
	{_State393, AtomTypeExprType}:                  _GotoState118Action,
	{_State393, NamedTypeExprType}:                 _GotoState126Action,
	{_State393, InferredTypeExprType}:              _GotoState124Action,
	{_State393, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State393, ReturnableTypeExprType}:            _GotoState130Action,
	{_State393, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State393, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State393, TypeExprType}:                      _GotoState236Action,
	{_State393, BinaryTypeExprType}:                _GotoState119Action,
	{_State393, FieldDefType}:                      _GotoState316Action,
	{_State393, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State393, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State393, EnumValueDefType}:                  _GotoState441Action,
	{_State393, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State393, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State393, TraitTypeExprType}:                 _GotoState131Action,
	{_State393, FuncTypeExprType}:                  _GotoState121Action,
	{_State394, AddToken}:                          _GotoState240Action,
	{_State394, SubToken}:                          _GotoState242Action,
	{_State394, MulToken}:                          _GotoState241Action,
	{_State394, BinaryTypeOpType}:                  _GotoState243Action,
	{_State395, IdentifierToken}:                   _GotoState112Action,
	{_State395, StructToken}:                       _GotoState50Action,
	{_State395, EnumToken}:                         _GotoState109Action,
	{_State395, TraitToken}:                        _GotoState117Action,
	{_State395, FuncToken}:                         _GotoState111Action,
	{_State395, LparenToken}:                       _GotoState113Action,
	{_State395, LbracketToken}:                     _GotoState42Action,
	{_State395, DotToken}:                          _GotoState108Action,
	{_State395, QuestionToken}:                     _GotoState115Action,
	{_State395, ExclaimToken}:                      _GotoState110Action,
	{_State395, TildeTildeToken}:                   _GotoState116Action,
	{_State395, BitNegToken}:                       _GotoState107Action,
	{_State395, BitAndToken}:                       _GotoState106Action,
	{_State395, ParseErrorToken}:                   _GotoState114Action,
	{_State395, InitializableTypeExprType}:         _GotoState125Action,
	{_State395, SliceTypeExprType}:                 _GotoState100Action,
	{_State395, ArrayTypeExprType}:                 _GotoState59Action,
	{_State395, MapTypeExprType}:                   _GotoState88Action,
	{_State395, AtomTypeExprType}:                  _GotoState118Action,
	{_State395, NamedTypeExprType}:                 _GotoState126Action,
	{_State395, InferredTypeExprType}:              _GotoState124Action,
	{_State395, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State395, ReturnableTypeExprType}:            _GotoState130Action,
	{_State395, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State395, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State395, TypeExprType}:                      _GotoState442Action,
	{_State395, BinaryTypeExprType}:                _GotoState119Action,
	{_State395, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State395, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State395, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State395, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State395, TraitTypeExprType}:                 _GotoState131Action,
	{_State395, FuncTypeExprType}:                  _GotoState121Action,
	{_State396, AddToken}:                          _GotoState240Action,
	{_State396, SubToken}:                          _GotoState242Action,
	{_State396, MulToken}:                          _GotoState241Action,
	{_State396, BinaryTypeOpType}:                  _GotoState243Action,
	{_State397, IdentifierToken}:                   _GotoState112Action,
	{_State397, StructToken}:                       _GotoState50Action,
	{_State397, EnumToken}:                         _GotoState109Action,
	{_State397, TraitToken}:                        _GotoState117Action,
	{_State397, FuncToken}:                         _GotoState111Action,
	{_State397, LparenToken}:                       _GotoState113Action,
	{_State397, LbracketToken}:                     _GotoState42Action,
	{_State397, DotToken}:                          _GotoState108Action,
	{_State397, QuestionToken}:                     _GotoState115Action,
	{_State397, ExclaimToken}:                      _GotoState110Action,
	{_State397, TildeTildeToken}:                   _GotoState116Action,
	{_State397, BitNegToken}:                       _GotoState107Action,
	{_State397, BitAndToken}:                       _GotoState106Action,
	{_State397, ParseErrorToken}:                   _GotoState114Action,
	{_State397, InitializableTypeExprType}:         _GotoState125Action,
	{_State397, SliceTypeExprType}:                 _GotoState100Action,
	{_State397, ArrayTypeExprType}:                 _GotoState59Action,
	{_State397, MapTypeExprType}:                   _GotoState88Action,
	{_State397, AtomTypeExprType}:                  _GotoState118Action,
	{_State397, NamedTypeExprType}:                 _GotoState126Action,
	{_State397, InferredTypeExprType}:              _GotoState124Action,
	{_State397, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State397, ReturnableTypeExprType}:            _GotoState419Action,
	{_State397, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State397, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State397, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State397, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State397, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State397, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State397, TraitTypeExprType}:                 _GotoState131Action,
	{_State397, ReturnTypeType}:                    _GotoState443Action,
	{_State397, FuncTypeExprType}:                  _GotoState121Action,
	{_State398, IdentifierToken}:                   _GotoState319Action,
	{_State398, StructToken}:                       _GotoState50Action,
	{_State398, EnumToken}:                         _GotoState109Action,
	{_State398, TraitToken}:                        _GotoState117Action,
	{_State398, FuncToken}:                         _GotoState111Action,
	{_State398, LparenToken}:                       _GotoState113Action,
	{_State398, LbracketToken}:                     _GotoState42Action,
	{_State398, DotToken}:                          _GotoState108Action,
	{_State398, QuestionToken}:                     _GotoState115Action,
	{_State398, ExclaimToken}:                      _GotoState110Action,
	{_State398, EllipsisToken}:                     _GotoState318Action,
	{_State398, TildeTildeToken}:                   _GotoState116Action,
	{_State398, BitNegToken}:                       _GotoState107Action,
	{_State398, BitAndToken}:                       _GotoState106Action,
	{_State398, ParseErrorToken}:                   _GotoState114Action,
	{_State398, InitializableTypeExprType}:         _GotoState125Action,
	{_State398, SliceTypeExprType}:                 _GotoState100Action,
	{_State398, ArrayTypeExprType}:                 _GotoState59Action,
	{_State398, MapTypeExprType}:                   _GotoState88Action,
	{_State398, AtomTypeExprType}:                  _GotoState118Action,
	{_State398, NamedTypeExprType}:                 _GotoState126Action,
	{_State398, InferredTypeExprType}:              _GotoState124Action,
	{_State398, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State398, ReturnableTypeExprType}:            _GotoState130Action,
	{_State398, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State398, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State398, TypeExprType}:                      _GotoState323Action,
	{_State398, BinaryTypeExprType}:                _GotoState119Action,
	{_State398, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State398, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State398, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State398, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State398, TraitTypeExprType}:                 _GotoState131Action,
	{_State398, ParameterDeclType}:                 _GotoState444Action,
	{_State398, FuncTypeExprType}:                  _GotoState121Action,
	{_State404, LparenToken}:                       _GotoState445Action,
	{_State405, IdentifierToken}:                   _GotoState230Action,
	{_State405, UnsafeToken}:                       _GotoState53Action,
	{_State405, StructToken}:                       _GotoState50Action,
	{_State405, EnumToken}:                         _GotoState109Action,
	{_State405, TraitToken}:                        _GotoState117Action,
	{_State405, FuncToken}:                         _GotoState333Action,
	{_State405, LparenToken}:                       _GotoState113Action,
	{_State405, LbracketToken}:                     _GotoState42Action,
	{_State405, DotToken}:                          _GotoState108Action,
	{_State405, QuestionToken}:                     _GotoState115Action,
	{_State405, ExclaimToken}:                      _GotoState110Action,
	{_State405, TildeTildeToken}:                   _GotoState116Action,
	{_State405, BitNegToken}:                       _GotoState107Action,
	{_State405, BitAndToken}:                       _GotoState106Action,
	{_State405, ParseErrorToken}:                   _GotoState114Action,
	{_State405, UnsafeStatementType}:               _GotoState237Action,
	{_State405, InitializableTypeExprType}:         _GotoState125Action,
	{_State405, SliceTypeExprType}:                 _GotoState100Action,
	{_State405, ArrayTypeExprType}:                 _GotoState59Action,
	{_State405, MapTypeExprType}:                   _GotoState88Action,
	{_State405, AtomTypeExprType}:                  _GotoState118Action,
	{_State405, NamedTypeExprType}:                 _GotoState126Action,
	{_State405, InferredTypeExprType}:              _GotoState124Action,
	{_State405, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State405, ReturnableTypeExprType}:            _GotoState130Action,
	{_State405, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State405, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State405, TypeExprType}:                      _GotoState236Action,
	{_State405, BinaryTypeExprType}:                _GotoState119Action,
	{_State405, FieldDefType}:                      _GotoState334Action,
	{_State405, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State405, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State405, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State405, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State405, TraitPropertyType}:                 _GotoState446Action,
	{_State405, TraitTypeExprType}:                 _GotoState131Action,
	{_State405, FuncTypeExprType}:                  _GotoState121Action,
	{_State405, MethodSignatureType}:               _GotoState335Action,
	{_State406, IdentifierToken}:                   _GotoState230Action,
	{_State406, UnsafeToken}:                       _GotoState53Action,
	{_State406, StructToken}:                       _GotoState50Action,
	{_State406, EnumToken}:                         _GotoState109Action,
	{_State406, TraitToken}:                        _GotoState117Action,
	{_State406, FuncToken}:                         _GotoState333Action,
	{_State406, LparenToken}:                       _GotoState113Action,
	{_State406, LbracketToken}:                     _GotoState42Action,
	{_State406, DotToken}:                          _GotoState108Action,
	{_State406, QuestionToken}:                     _GotoState115Action,
	{_State406, ExclaimToken}:                      _GotoState110Action,
	{_State406, TildeTildeToken}:                   _GotoState116Action,
	{_State406, BitNegToken}:                       _GotoState107Action,
	{_State406, BitAndToken}:                       _GotoState106Action,
	{_State406, ParseErrorToken}:                   _GotoState114Action,
	{_State406, UnsafeStatementType}:               _GotoState237Action,
	{_State406, InitializableTypeExprType}:         _GotoState125Action,
	{_State406, SliceTypeExprType}:                 _GotoState100Action,
	{_State406, ArrayTypeExprType}:                 _GotoState59Action,
	{_State406, MapTypeExprType}:                   _GotoState88Action,
	{_State406, AtomTypeExprType}:                  _GotoState118Action,
	{_State406, NamedTypeExprType}:                 _GotoState126Action,
	{_State406, InferredTypeExprType}:              _GotoState124Action,
	{_State406, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State406, ReturnableTypeExprType}:            _GotoState130Action,
	{_State406, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State406, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State406, TypeExprType}:                      _GotoState236Action,
	{_State406, BinaryTypeExprType}:                _GotoState119Action,
	{_State406, FieldDefType}:                      _GotoState334Action,
	{_State406, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State406, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State406, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State406, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State406, TraitPropertyType}:                 _GotoState447Action,
	{_State406, TraitTypeExprType}:                 _GotoState131Action,
	{_State406, FuncTypeExprType}:                  _GotoState121Action,
	{_State406, MethodSignatureType}:               _GotoState335Action,
	{_State408, AddToken}:                          _GotoState240Action,
	{_State408, SubToken}:                          _GotoState242Action,
	{_State408, MulToken}:                          _GotoState241Action,
	{_State408, BinaryTypeOpType}:                  _GotoState243Action,
	{_State410, IdentifierToken}:                   _GotoState341Action,
	{_State410, GenericParameterDefType}:           _GotoState448Action,
	{_State411, RparenToken}:                       _GotoState449Action,
	{_State412, AddToken}:                          _GotoState240Action,
	{_State412, SubToken}:                          _GotoState242Action,
	{_State412, MulToken}:                          _GotoState241Action,
	{_State412, BinaryTypeOpType}:                  _GotoState243Action,
	{_State413, DollarLbracketToken}:               _GotoState245Action,
	{_State413, OptionalGenericParametersType}:     _GotoState450Action,
	{_State414, IdentifierToken}:                   _GotoState112Action,
	{_State414, StructToken}:                       _GotoState50Action,
	{_State414, EnumToken}:                         _GotoState109Action,
	{_State414, TraitToken}:                        _GotoState117Action,
	{_State414, FuncToken}:                         _GotoState111Action,
	{_State414, LparenToken}:                       _GotoState113Action,
	{_State414, LbracketToken}:                     _GotoState42Action,
	{_State414, DotToken}:                          _GotoState108Action,
	{_State414, QuestionToken}:                     _GotoState115Action,
	{_State414, ExclaimToken}:                      _GotoState110Action,
	{_State414, TildeTildeToken}:                   _GotoState116Action,
	{_State414, BitNegToken}:                       _GotoState107Action,
	{_State414, BitAndToken}:                       _GotoState106Action,
	{_State414, ParseErrorToken}:                   _GotoState114Action,
	{_State414, InitializableTypeExprType}:         _GotoState125Action,
	{_State414, SliceTypeExprType}:                 _GotoState100Action,
	{_State414, ArrayTypeExprType}:                 _GotoState59Action,
	{_State414, MapTypeExprType}:                   _GotoState88Action,
	{_State414, AtomTypeExprType}:                  _GotoState118Action,
	{_State414, NamedTypeExprType}:                 _GotoState126Action,
	{_State414, InferredTypeExprType}:              _GotoState124Action,
	{_State414, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State414, ReturnableTypeExprType}:            _GotoState130Action,
	{_State414, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State414, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State414, TypeExprType}:                      _GotoState451Action,
	{_State414, BinaryTypeExprType}:                _GotoState119Action,
	{_State414, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State414, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State414, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State414, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State414, TraitTypeExprType}:                 _GotoState131Action,
	{_State414, FuncTypeExprType}:                  _GotoState121Action,
	{_State418, LbraceToken}:                       _GotoState11Action,
	{_State418, StatementBlockType}:                _GotoState452Action,
	{_State428, AddToken}:                          _GotoState240Action,
	{_State428, SubToken}:                          _GotoState242Action,
	{_State428, MulToken}:                          _GotoState241Action,
	{_State428, BinaryTypeOpType}:                  _GotoState243Action,
	{_State432, DoToken}:                           _GotoState453Action,
	{_State433, SemicolonToken}:                    _GotoState454Action,
	{_State436, IntegerLiteralToken}:               _GotoState40Action,
	{_State436, FloatLiteralToken}:                 _GotoState35Action,
	{_State436, RuneLiteralToken}:                  _GotoState48Action,
	{_State436, StringLiteralToken}:                _GotoState49Action,
	{_State436, IdentifierToken}:                   _GotoState38Action,
	{_State436, TrueToken}:                         _GotoState52Action,
	{_State436, FalseToken}:                        _GotoState34Action,
	{_State436, StructToken}:                       _GotoState50Action,
	{_State436, FuncToken}:                         _GotoState36Action,
	{_State436, VarToken}:                          _GotoState15Action,
	{_State436, LetToken}:                          _GotoState12Action,
	{_State436, NotToken}:                          _GotoState45Action,
	{_State436, LabelDeclToken}:                    _GotoState41Action,
	{_State436, LparenToken}:                       _GotoState43Action,
	{_State436, LbracketToken}:                     _GotoState42Action,
	{_State436, SubToken}:                          _GotoState51Action,
	{_State436, MulToken}:                          _GotoState44Action,
	{_State436, BitNegToken}:                       _GotoState27Action,
	{_State436, BitAndToken}:                       _GotoState26Action,
	{_State436, GreaterToken}:                      _GotoState37Action,
	{_State436, ParseErrorToken}:                   _GotoState46Action,
	{_State436, VarDeclPatternType}:                _GotoState103Action,
	{_State436, VarOrLetType}:                      _GotoState24Action,
	{_State436, OptionalLabelDeclType}:             _GotoState150Action,
	{_State436, SequenceExprType}:                  _GotoState455Action,
	{_State436, CallExprType}:                      _GotoState70Action,
	{_State436, AtomExprType}:                      _GotoState62Action,
	{_State436, ParseErrorExprType}:                _GotoState92Action,
	{_State436, LiteralExprType}:                   _GotoState87Action,
	{_State436, IdentifierExprType}:                _GotoState79Action,
	{_State436, BlockExprType}:                     _GotoState69Action,
	{_State436, InitializeExprType}:                _GotoState84Action,
	{_State436, ImplicitStructExprType}:            _GotoState80Action,
	{_State436, AccessibleExprType}:                _GotoState104Action,
	{_State436, AccessExprType}:                    _GotoState54Action,
	{_State436, IndexExprType}:                     _GotoState82Action,
	{_State436, PostfixableExprType}:               _GotoState94Action,
	{_State436, PostfixUnaryExprType}:              _GotoState93Action,
	{_State436, PrefixableExprType}:                _GotoState97Action,
	{_State436, PrefixUnaryExprType}:               _GotoState95Action,
	{_State436, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State436, MulExprType}:                       _GotoState89Action,
	{_State436, BinaryMulExprType}:                 _GotoState66Action,
	{_State436, AddExprType}:                       _GotoState56Action,
	{_State436, BinaryAddExprType}:                 _GotoState63Action,
	{_State436, CmpExprType}:                       _GotoState73Action,
	{_State436, BinaryCmpExprType}:                 _GotoState65Action,
	{_State436, AndExprType}:                       _GotoState57Action,
	{_State436, BinaryAndExprType}:                 _GotoState64Action,
	{_State436, OrExprType}:                        _GotoState91Action,
	{_State436, BinaryOrExprType}:                  _GotoState68Action,
	{_State436, InitializableTypeExprType}:         _GotoState83Action,
	{_State436, SliceTypeExprType}:                 _GotoState100Action,
	{_State436, ArrayTypeExprType}:                 _GotoState59Action,
	{_State436, MapTypeExprType}:                   _GotoState88Action,
	{_State436, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State436, AnonymousFuncExprType}:             _GotoState58Action,
	{_State437, IfToken}:                           _GotoState218Action,
	{_State437, LbraceToken}:                       _GotoState11Action,
	{_State437, StatementBlockType}:                _GotoState457Action,
	{_State437, IfExprType}:                        _GotoState456Action,
	{_State442, AddToken}:                          _GotoState240Action,
	{_State442, SubToken}:                          _GotoState242Action,
	{_State442, MulToken}:                          _GotoState241Action,
	{_State442, BinaryTypeOpType}:                  _GotoState243Action,
	{_State445, IdentifierToken}:                   _GotoState319Action,
	{_State445, StructToken}:                       _GotoState50Action,
	{_State445, EnumToken}:                         _GotoState109Action,
	{_State445, TraitToken}:                        _GotoState117Action,
	{_State445, FuncToken}:                         _GotoState111Action,
	{_State445, LparenToken}:                       _GotoState113Action,
	{_State445, LbracketToken}:                     _GotoState42Action,
	{_State445, DotToken}:                          _GotoState108Action,
	{_State445, QuestionToken}:                     _GotoState115Action,
	{_State445, ExclaimToken}:                      _GotoState110Action,
	{_State445, EllipsisToken}:                     _GotoState318Action,
	{_State445, TildeTildeToken}:                   _GotoState116Action,
	{_State445, BitNegToken}:                       _GotoState107Action,
	{_State445, BitAndToken}:                       _GotoState106Action,
	{_State445, ParseErrorToken}:                   _GotoState114Action,
	{_State445, InitializableTypeExprType}:         _GotoState125Action,
	{_State445, SliceTypeExprType}:                 _GotoState100Action,
	{_State445, ArrayTypeExprType}:                 _GotoState59Action,
	{_State445, MapTypeExprType}:                   _GotoState88Action,
	{_State445, AtomTypeExprType}:                  _GotoState118Action,
	{_State445, NamedTypeExprType}:                 _GotoState126Action,
	{_State445, InferredTypeExprType}:              _GotoState124Action,
	{_State445, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State445, ReturnableTypeExprType}:            _GotoState130Action,
	{_State445, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State445, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State445, TypeExprType}:                      _GotoState323Action,
	{_State445, BinaryTypeExprType}:                _GotoState119Action,
	{_State445, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State445, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State445, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State445, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State445, TraitTypeExprType}:                 _GotoState131Action,
	{_State445, ParameterDeclType}:                 _GotoState320Action,
	{_State445, ProperParameterDeclsType}:          _GotoState322Action,
	{_State445, ParameterDeclsType}:                _GotoState458Action,
	{_State445, FuncTypeExprType}:                  _GotoState121Action,
	{_State449, IdentifierToken}:                   _GotoState112Action,
	{_State449, StructToken}:                       _GotoState50Action,
	{_State449, EnumToken}:                         _GotoState109Action,
	{_State449, TraitToken}:                        _GotoState117Action,
	{_State449, FuncToken}:                         _GotoState111Action,
	{_State449, LparenToken}:                       _GotoState113Action,
	{_State449, LbracketToken}:                     _GotoState42Action,
	{_State449, DotToken}:                          _GotoState108Action,
	{_State449, QuestionToken}:                     _GotoState115Action,
	{_State449, ExclaimToken}:                      _GotoState110Action,
	{_State449, TildeTildeToken}:                   _GotoState116Action,
	{_State449, BitNegToken}:                       _GotoState107Action,
	{_State449, BitAndToken}:                       _GotoState106Action,
	{_State449, ParseErrorToken}:                   _GotoState114Action,
	{_State449, InitializableTypeExprType}:         _GotoState125Action,
	{_State449, SliceTypeExprType}:                 _GotoState100Action,
	{_State449, ArrayTypeExprType}:                 _GotoState59Action,
	{_State449, MapTypeExprType}:                   _GotoState88Action,
	{_State449, AtomTypeExprType}:                  _GotoState118Action,
	{_State449, NamedTypeExprType}:                 _GotoState126Action,
	{_State449, InferredTypeExprType}:              _GotoState124Action,
	{_State449, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State449, ReturnableTypeExprType}:            _GotoState419Action,
	{_State449, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State449, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State449, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State449, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State449, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State449, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State449, TraitTypeExprType}:                 _GotoState131Action,
	{_State449, ReturnTypeType}:                    _GotoState459Action,
	{_State449, FuncTypeExprType}:                  _GotoState121Action,
	{_State450, LparenToken}:                       _GotoState460Action,
	{_State451, AddToken}:                          _GotoState240Action,
	{_State451, SubToken}:                          _GotoState242Action,
	{_State451, MulToken}:                          _GotoState241Action,
	{_State451, BinaryTypeOpType}:                  _GotoState243Action,
	{_State453, LbraceToken}:                       _GotoState11Action,
	{_State453, StatementBlockType}:                _GotoState461Action,
	{_State454, IntegerLiteralToken}:               _GotoState40Action,
	{_State454, FloatLiteralToken}:                 _GotoState35Action,
	{_State454, RuneLiteralToken}:                  _GotoState48Action,
	{_State454, StringLiteralToken}:                _GotoState49Action,
	{_State454, IdentifierToken}:                   _GotoState38Action,
	{_State454, TrueToken}:                         _GotoState52Action,
	{_State454, FalseToken}:                        _GotoState34Action,
	{_State454, StructToken}:                       _GotoState50Action,
	{_State454, FuncToken}:                         _GotoState36Action,
	{_State454, VarToken}:                          _GotoState15Action,
	{_State454, LetToken}:                          _GotoState12Action,
	{_State454, NotToken}:                          _GotoState45Action,
	{_State454, LabelDeclToken}:                    _GotoState41Action,
	{_State454, LparenToken}:                       _GotoState43Action,
	{_State454, LbracketToken}:                     _GotoState42Action,
	{_State454, SubToken}:                          _GotoState51Action,
	{_State454, MulToken}:                          _GotoState44Action,
	{_State454, BitNegToken}:                       _GotoState27Action,
	{_State454, BitAndToken}:                       _GotoState26Action,
	{_State454, GreaterToken}:                      _GotoState37Action,
	{_State454, ParseErrorToken}:                   _GotoState46Action,
	{_State454, VarDeclPatternType}:                _GotoState103Action,
	{_State454, VarOrLetType}:                      _GotoState24Action,
	{_State454, OptionalLabelDeclType}:             _GotoState150Action,
	{_State454, SequenceExprType}:                  _GotoState434Action,
	{_State454, OptionalSequenceExprType}:          _GotoState462Action,
	{_State454, CallExprType}:                      _GotoState70Action,
	{_State454, AtomExprType}:                      _GotoState62Action,
	{_State454, ParseErrorExprType}:                _GotoState92Action,
	{_State454, LiteralExprType}:                   _GotoState87Action,
	{_State454, IdentifierExprType}:                _GotoState79Action,
	{_State454, BlockExprType}:                     _GotoState69Action,
	{_State454, InitializeExprType}:                _GotoState84Action,
	{_State454, ImplicitStructExprType}:            _GotoState80Action,
	{_State454, AccessibleExprType}:                _GotoState104Action,
	{_State454, AccessExprType}:                    _GotoState54Action,
	{_State454, IndexExprType}:                     _GotoState82Action,
	{_State454, PostfixableExprType}:               _GotoState94Action,
	{_State454, PostfixUnaryExprType}:              _GotoState93Action,
	{_State454, PrefixableExprType}:                _GotoState97Action,
	{_State454, PrefixUnaryExprType}:               _GotoState95Action,
	{_State454, PrefixUnaryOpType}:                 _GotoState96Action,
	{_State454, MulExprType}:                       _GotoState89Action,
	{_State454, BinaryMulExprType}:                 _GotoState66Action,
	{_State454, AddExprType}:                       _GotoState56Action,
	{_State454, BinaryAddExprType}:                 _GotoState63Action,
	{_State454, CmpExprType}:                       _GotoState73Action,
	{_State454, BinaryCmpExprType}:                 _GotoState65Action,
	{_State454, AndExprType}:                       _GotoState57Action,
	{_State454, BinaryAndExprType}:                 _GotoState64Action,
	{_State454, OrExprType}:                        _GotoState91Action,
	{_State454, BinaryOrExprType}:                  _GotoState68Action,
	{_State454, InitializableTypeExprType}:         _GotoState83Action,
	{_State454, SliceTypeExprType}:                 _GotoState100Action,
	{_State454, ArrayTypeExprType}:                 _GotoState59Action,
	{_State454, MapTypeExprType}:                   _GotoState88Action,
	{_State454, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State454, AnonymousFuncExprType}:             _GotoState58Action,
	{_State458, RparenToken}:                       _GotoState463Action,
	{_State459, LbraceToken}:                       _GotoState11Action,
	{_State459, StatementBlockType}:                _GotoState464Action,
	{_State460, IdentifierToken}:                   _GotoState247Action,
	{_State460, ParameterDefType}:                  _GotoState269Action,
	{_State460, ProperParameterDefsType}:           _GotoState271Action,
	{_State460, ParameterDefsType}:                 _GotoState465Action,
	{_State462, DoToken}:                           _GotoState466Action,
	{_State463, IdentifierToken}:                   _GotoState112Action,
	{_State463, StructToken}:                       _GotoState50Action,
	{_State463, EnumToken}:                         _GotoState109Action,
	{_State463, TraitToken}:                        _GotoState117Action,
	{_State463, FuncToken}:                         _GotoState111Action,
	{_State463, LparenToken}:                       _GotoState113Action,
	{_State463, LbracketToken}:                     _GotoState42Action,
	{_State463, DotToken}:                          _GotoState108Action,
	{_State463, QuestionToken}:                     _GotoState115Action,
	{_State463, ExclaimToken}:                      _GotoState110Action,
	{_State463, TildeTildeToken}:                   _GotoState116Action,
	{_State463, BitNegToken}:                       _GotoState107Action,
	{_State463, BitAndToken}:                       _GotoState106Action,
	{_State463, ParseErrorToken}:                   _GotoState114Action,
	{_State463, InitializableTypeExprType}:         _GotoState125Action,
	{_State463, SliceTypeExprType}:                 _GotoState100Action,
	{_State463, ArrayTypeExprType}:                 _GotoState59Action,
	{_State463, MapTypeExprType}:                   _GotoState88Action,
	{_State463, AtomTypeExprType}:                  _GotoState118Action,
	{_State463, NamedTypeExprType}:                 _GotoState126Action,
	{_State463, InferredTypeExprType}:              _GotoState124Action,
	{_State463, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State463, ReturnableTypeExprType}:            _GotoState419Action,
	{_State463, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State463, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State463, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State463, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State463, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State463, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State463, TraitTypeExprType}:                 _GotoState131Action,
	{_State463, ReturnTypeType}:                    _GotoState467Action,
	{_State463, FuncTypeExprType}:                  _GotoState121Action,
	{_State465, RparenToken}:                       _GotoState468Action,
	{_State466, LbraceToken}:                       _GotoState11Action,
	{_State466, StatementBlockType}:                _GotoState469Action,
	{_State468, IdentifierToken}:                   _GotoState112Action,
	{_State468, StructToken}:                       _GotoState50Action,
	{_State468, EnumToken}:                         _GotoState109Action,
	{_State468, TraitToken}:                        _GotoState117Action,
	{_State468, FuncToken}:                         _GotoState111Action,
	{_State468, LparenToken}:                       _GotoState113Action,
	{_State468, LbracketToken}:                     _GotoState42Action,
	{_State468, DotToken}:                          _GotoState108Action,
	{_State468, QuestionToken}:                     _GotoState115Action,
	{_State468, ExclaimToken}:                      _GotoState110Action,
	{_State468, TildeTildeToken}:                   _GotoState116Action,
	{_State468, BitNegToken}:                       _GotoState107Action,
	{_State468, BitAndToken}:                       _GotoState106Action,
	{_State468, ParseErrorToken}:                   _GotoState114Action,
	{_State468, InitializableTypeExprType}:         _GotoState125Action,
	{_State468, SliceTypeExprType}:                 _GotoState100Action,
	{_State468, ArrayTypeExprType}:                 _GotoState59Action,
	{_State468, MapTypeExprType}:                   _GotoState88Action,
	{_State468, AtomTypeExprType}:                  _GotoState118Action,
	{_State468, NamedTypeExprType}:                 _GotoState126Action,
	{_State468, InferredTypeExprType}:              _GotoState124Action,
	{_State468, ParseErrorTypeExprType}:            _GotoState127Action,
	{_State468, ReturnableTypeExprType}:            _GotoState419Action,
	{_State468, PrefixedUnaryTypeExprType}:         _GotoState129Action,
	{_State468, PrefixUnaryTypeOpType}:             _GotoState128Action,
	{_State468, ImplicitStructTypeExprType}:        _GotoState123Action,
	{_State468, ExplicitStructTypeExprType}:        _GotoState74Action,
	{_State468, ImplicitEnumTypeExprType}:          _GotoState122Action,
	{_State468, ExplicitEnumTypeExprType}:          _GotoState120Action,
	{_State468, TraitTypeExprType}:                 _GotoState131Action,
	{_State468, ReturnTypeType}:                    _GotoState470Action,
	{_State468, FuncTypeExprType}:                  _GotoState121Action,
	{_State470, LbraceToken}:                       _GotoState11Action,
	{_State470, StatementBlockType}:                _GotoState471Action,
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
	{_State38, _WildcardMarker}:                    _ReduceToIdentifierExprAction,
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
	{_State54, _WildcardMarker}:                    _ReduceAccessExprToAccessibleExprAction,
	{_State55, _WildcardMarker}:                    _ReduceAccessibleExprToPostfixableExprAction,
	{_State55, LparenToken}:                        _ReduceNilToGenericTypeArgumentsAction,
	{_State56, _WildcardMarker}:                    _ReduceAddExprToCmpExprAction,
	{_State57, _WildcardMarker}:                    _ReduceAndExprToOrExprAction,
	{_State58, _WildcardMarker}:                    _ReduceAnonymousFuncExprToAtomExprAction,
	{_State59, _WildcardMarker}:                    _ReduceArrayTypeExprToInitializableTypeExprAction,
	{_State61, _EndMarker}:                         _ReduceAssignStatementToSimpleStatementAction,
	{_State62, _WildcardMarker}:                    _ReduceAtomExprToAccessibleExprAction,
	{_State63, _WildcardMarker}:                    _ReduceBinaryAddExprToAddExprAction,
	{_State64, _WildcardMarker}:                    _ReduceBinaryAndExprToAndExprAction,
	{_State65, _WildcardMarker}:                    _ReduceBinaryCmpExprToCmpExprAction,
	{_State66, _WildcardMarker}:                    _ReduceBinaryMulExprToMulExprAction,
	{_State67, _EndMarker}:                         _ReduceBinaryOpAssignStatementToSimpleStatementAction,
	{_State68, _WildcardMarker}:                    _ReduceBinaryOrExprToOrExprAction,
	{_State69, _WildcardMarker}:                    _ReduceBlockExprToAtomExprAction,
	{_State70, _WildcardMarker}:                    _ReduceCallExprToAccessibleExprAction,
	{_State71, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State72, _EndMarker}:                         _ReduceCallbackOpStatementToSimpleStatementAction,
	{_State73, _WildcardMarker}:                    _ReduceCmpExprToAndExprAction,
	{_State74, _WildcardMarker}:                    _ReduceExplicitStructTypeExprToInitializableTypeExprAction,
	{_State75, _WildcardMarker}:                    _ReduceExprToExprsAction,
	{_State76, _EndMarker}:                         _ReduceExprOrImproperStructStatementToSimpleStatementAction,
	{_State77, _WildcardMarker}:                    _ReduceToExprOrImproperStructStatementAction,
	{_State78, _EndMarker}:                         _ReduceFallthroughStatementToSimpleStatementAction,
	{_State79, _WildcardMarker}:                    _ReduceIdentifierExprToAtomExprAction,
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
	{_State91, _WildcardMarker}:                    _ReduceOrExprToSequenceExprAction,
	{_State92, _WildcardMarker}:                    _ReduceParseErrorExprToAtomExprAction,
	{_State93, _WildcardMarker}:                    _ReducePostfixUnaryExprToPostfixableExprAction,
	{_State94, _WildcardMarker}:                    _ReducePostfixableExprToPrefixableExprAction,
	{_State95, _WildcardMarker}:                    _ReducePrefixUnaryExprToPrefixableExprAction,
	{_State96, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State97, _WildcardMarker}:                    _ReducePrefixableExprToMulExprAction,
	{_State98, AssignToken}:                        _ReduceToAssignPatternAction,
	{_State98, _WildcardMarker}:                    _ReduceSequenceExprToExprAction,
	{_State99, _EndMarker}:                         _ReduceSimpleStatementToStatementAction,
	{_State100, _WildcardMarker}:                   _ReduceSliceTypeExprToInitializableTypeExprAction,
	{_State101, _EndMarker}:                        _ReduceUnaryOpAssignStatementToSimpleStatementAction,
	{_State102, _EndMarker}:                        _ReduceUnsafeStatementToSimpleStatementAction,
	{_State103, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State104, _WildcardMarker}:                   _ReduceAccessibleExprToPostfixableExprAction,
	{_State104, LparenToken}:                       _ReduceNilToGenericTypeArgumentsAction,
	{_State105, _EndMarker}:                        _ReduceSequenceExprToExprAction,
	{_State106, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryTypeOpAction,
	{_State107, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryTypeOpAction,
	{_State108, _WildcardMarker}:                   _ReduceToInferredTypeExprAction,
	{_State110, _WildcardMarker}:                   _ReduceExclaimToPrefixUnaryTypeOpAction,
	{_State112, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State113, RparenToken}:                       _ReduceNilToImplicitFieldDefsAction,
	{_State114, _WildcardMarker}:                   _ReduceToParseErrorTypeExprAction,
	{_State115, _WildcardMarker}:                   _ReduceQuestionToPrefixUnaryTypeOpAction,
	{_State116, _WildcardMarker}:                   _ReduceTildeTildeToPrefixUnaryTypeOpAction,
	{_State118, _WildcardMarker}:                   _ReduceAtomTypeExprToReturnableTypeExprAction,
	{_State119, _WildcardMarker}:                   _ReduceBinaryTypeExprToTypeExprAction,
	{_State120, _WildcardMarker}:                   _ReduceExplicitEnumTypeExprToAtomTypeExprAction,
	{_State121, _WildcardMarker}:                   _ReduceFuncTypeExprToAtomTypeExprAction,
	{_State122, _WildcardMarker}:                   _ReduceImplicitEnumTypeExprToAtomTypeExprAction,
	{_State123, _WildcardMarker}:                   _ReduceImplicitStructTypeExprToAtomTypeExprAction,
	{_State124, _WildcardMarker}:                   _ReduceInferredTypeExprToAtomTypeExprAction,
	{_State125, _WildcardMarker}:                   _ReduceInitializableTypeExprToAtomTypeExprAction,
	{_State126, _WildcardMarker}:                   _ReduceNamedTypeExprToAtomTypeExprAction,
	{_State127, _WildcardMarker}:                   _ReduceParseErrorTypeExprToAtomTypeExprAction,
	{_State129, _WildcardMarker}:                   _ReducePrefixedUnaryTypeExprToReturnableTypeExprAction,
	{_State130, _WildcardMarker}:                   _ReduceReturnableTypeExprToTypeExprAction,
	{_State131, _WildcardMarker}:                   _ReduceTraitTypeExprToAtomTypeExprAction,
	{_State132, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State134, RbraceToken}:                       _ReduceProperStatementsToStatementsAction,
	{_State135, _WildcardMarker}:                   _ReduceStatementToProperStatementsAction,
	{_State137, _WildcardMarker}:                   _ReduceWithSpecToPackageDefAction,
	{_State138, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State139, _EndMarker}:                        _ReduceImproperToDefinitionsAction,
	{_State140, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State143, _WildcardMarker}:                   _ReduceTuplePatternToVarPatternAction,
	{_State144, _WildcardMarker}:                   _ReduceNilToOptionalTypeExprAction,
	{_State146, _WildcardMarker}:                   _ReduceVarToVarOrLetAction,
	{_State147, _WildcardMarker}:                   _ReduceAssignPatternToCasePatternAction,
	{_State148, _WildcardMarker}:                   _ReduceCasePatternToCasePatternsAction,
	{_State151, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State152, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State152, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State152, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State152, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State152, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State153, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State154, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State156, _WildcardMarker}:                   _ReduceStringLiteralToImportClauseAction,
	{_State157, _EndMarker}:                        _ReduceSingleToImportStatementAction,
	{_State159, ColonToken}:                        _ReduceUnitUnitPairToColonExprAction,
	{_State159, CommaToken}:                        _ReduceUnitUnitPairToColonExprAction,
	{_State159, RbracketToken}:                     _ReduceUnitUnitPairToColonExprAction,
	{_State159, RparenToken}:                       _ReduceUnitUnitPairToColonExprAction,
	{_State159, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State160, _WildcardMarker}:                   _ReduceSkipPatternToArgumentAction,
	{_State161, _WildcardMarker}:                   _ReduceToIdentifierExprAction,
	{_State162, _WildcardMarker}:                   _ReduceArgumentToProperArgumentsAction,
	{_State164, _WildcardMarker}:                   _ReduceColonExprToArgumentAction,
	{_State165, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State166, RparenToken}:                       _ReduceProperArgumentsToArgumentsAction,
	{_State167, RparenToken}:                       _ReduceNilToExplicitFieldDefsAction,
	{_State169, _WildcardMarker}:                   _ReduceAddAssignToBinaryOpAssignAction,
	{_State170, _EndMarker}:                        _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State171, _WildcardMarker}:                   _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State172, _WildcardMarker}:                   _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State173, _WildcardMarker}:                   _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State174, _WildcardMarker}:                   _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State175, _WildcardMarker}:                   _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State176, _WildcardMarker}:                   _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State177, _WildcardMarker}:                   _ReduceDivAssignToBinaryOpAssignAction,
	{_State178, RbracketToken}:                     _ReduceNilToTypeArgumentsAction,
	{_State180, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State181, _WildcardMarker}:                   _ReduceModAssignToBinaryOpAssignAction,
	{_State182, _WildcardMarker}:                   _ReduceMulAssignToBinaryOpAssignAction,
	{_State183, _WildcardMarker}:                   _ReduceToPostfixUnaryExprAction,
	{_State184, _WildcardMarker}:                   _ReduceSubAssignToBinaryOpAssignAction,
	{_State185, _EndMarker}:                        _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State186, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State188, _EndMarker}:                        _ReduceToUnaryOpAssignStatementAction,
	{_State189, _WildcardMarker}:                   _ReduceAddToAddOpAction,
	{_State190, _WildcardMarker}:                   _ReduceBitOrToAddOpAction,
	{_State191, _WildcardMarker}:                   _ReduceBitXorToAddOpAction,
	{_State192, _WildcardMarker}:                   _ReduceSubToAddOpAction,
	{_State193, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State194, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State195, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State196, LparenToken}:                       _ReduceNilToGenericTypeArgumentsAction,
	{_State197, _EndMarker}:                        _ReduceToCallbackOpStatementAction,
	{_State197, NewlinesToken}:                     _ReduceToCallbackOpStatementAction,
	{_State197, RbraceToken}:                       _ReduceToCallbackOpStatementAction,
	{_State197, SemicolonToken}:                    _ReduceToCallbackOpStatementAction,
	{_State197, _WildcardMarker}:                   _ReduceCallExprToAccessibleExprAction,
	{_State198, _WildcardMarker}:                   _ReduceEqualToCmpOpAction,
	{_State199, _WildcardMarker}:                   _ReduceGreaterToCmpOpAction,
	{_State200, _WildcardMarker}:                   _ReduceGreaterOrEqualToCmpOpAction,
	{_State201, _WildcardMarker}:                   _ReduceLessToCmpOpAction,
	{_State202, _WildcardMarker}:                   _ReduceLessOrEqualToCmpOpAction,
	{_State203, _WildcardMarker}:                   _ReduceNotEqualToCmpOpAction,
	{_State204, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State205, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State206, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State206, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State207, _EndMarker}:                        _ReduceLabeledNoValueToJumpStatementAction,
	{_State207, NewlinesToken}:                     _ReduceLabeledNoValueToJumpStatementAction,
	{_State207, RbraceToken}:                       _ReduceLabeledNoValueToJumpStatementAction,
	{_State207, SemicolonToken}:                    _ReduceLabeledNoValueToJumpStatementAction,
	{_State207, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State208, _EndMarker}:                        _ReduceUnlabeledValuedToJumpStatementAction,
	{_State209, _WildcardMarker}:                   _ReduceBitAndToMulOpAction,
	{_State210, _WildcardMarker}:                   _ReduceBitLshiftToMulOpAction,
	{_State211, _WildcardMarker}:                   _ReduceBitRshiftToMulOpAction,
	{_State212, _WildcardMarker}:                   _ReduceDivToMulOpAction,
	{_State213, _WildcardMarker}:                   _ReduceModToMulOpAction,
	{_State214, _WildcardMarker}:                   _ReduceMulToMulOpAction,
	{_State215, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State217, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State218, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State219, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State220, _EndMarker}:                        _ReduceIfExprToExprAction,
	{_State221, _EndMarker}:                        _ReduceLoopExprToExprAction,
	{_State222, _WildcardMarker}:                   _ReduceToBlockExprAction,
	{_State223, _EndMarker}:                        _ReduceSwitchExprToExprAction,
	{_State224, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State225, _WildcardMarker}:                   _ReduceToPrefixUnaryExprAction,
	{_State227, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State229, _WildcardMarker}:                   _ReduceLocalToNamedTypeExprAction,
	{_State230, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State232, _WildcardMarker}:                   _ReduceFieldDefToProperImplicitFieldDefsAction,
	{_State232, OrToken}:                           _ReduceFieldDefToEnumValueDefAction,
	{_State235, RparenToken}:                       _ReduceProperImplicitFieldDefsToImplicitFieldDefsAction,
	{_State236, _WildcardMarker}:                   _ReduceImplicitToFieldDefAction,
	{_State237, _WildcardMarker}:                   _ReduceUnsafeStatementToFieldDefAction,
	{_State238, RparenToken}:                       _ReduceNilToTraitPropertiesAction,
	{_State239, _WildcardMarker}:                   _ReduceToPrefixedUnaryTypeExprAction,
	{_State240, _WildcardMarker}:                   _ReduceAddToBinaryTypeOpAction,
	{_State241, _WildcardMarker}:                   _ReduceMulToBinaryTypeOpAction,
	{_State242, _WildcardMarker}:                   _ReduceSubToBinaryTypeOpAction,
	{_State244, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State245, RbracketToken}:                     _ReduceNilToGenericParameterDefsAction,
	{_State247, _WildcardMarker}:                   _ReduceInferredRefArgToParameterDefAction,
	{_State249, RbraceToken}:                       _ReduceImproperImplicitToStatementsAction,
	{_State249, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State250, RbraceToken}:                       _ReduceImproperExplicitToStatementsAction,
	{_State250, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State251, _EndMarker}:                        _ReduceToStatementBlockAction,
	{_State254, _WildcardMarker}:                   _ReduceAddToProperDefinitionsAction,
	{_State255, _WildcardMarker}:                   _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State256, _WildcardMarker}:                   _ReduceEllipsisToFieldVarPatternAction,
	{_State257, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State258, _WildcardMarker}:                   _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State260, _WildcardMarker}:                   _ReducePositionalToFieldVarPatternAction,
	{_State261, _EndMarker}:                        _ReduceToVarDeclPatternAction,
	{_State262, _WildcardMarker}:                   _ReduceTypeExprToOptionalTypeExprAction,
	{_State265, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State265, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State265, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State265, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State265, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State266, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State267, _EndMarker}:                        _ReduceDefaultBranchStatementToStatementAction,
	{_State268, _EndMarker}:                        _ReduceSimpleStatementToOptionalSimpleStatementAction,
	{_State269, _WildcardMarker}:                   _ReduceParameterDefToProperParameterDefsAction,
	{_State271, RparenToken}:                       _ReduceProperParameterDefsToParameterDefsAction,
	{_State272, _WildcardMarker}:                   _ReduceImportClauseToProperImportClausesAction,
	{_State274, RparenToken}:                       _ReduceProperImportClausesToImportClausesAction,
	{_State278, _WildcardMarker}:                   _ReduceToSliceTypeExprAction,
	{_State279, _WildcardMarker}:                   _ReduceUnitExprPairToColonExprAction,
	{_State280, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State281, _WildcardMarker}:                   _ReduceToImplicitStructExprAction,
	{_State282, ColonToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State282, CommaToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State282, RbracketToken}:                     _ReduceColonExprUnitTupleToColonExprAction,
	{_State282, RparenToken}:                       _ReduceColonExprUnitTupleToColonExprAction,
	{_State282, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State283, ColonToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State283, CommaToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State283, RbracketToken}:                     _ReduceExprUnitPairToColonExprAction,
	{_State283, RparenToken}:                       _ReduceExprUnitPairToColonExprAction,
	{_State283, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State284, _WildcardMarker}:                   _ReduceVarargAssignmentToArgumentAction,
	{_State285, RparenToken}:                       _ReduceImproperToArgumentsAction,
	{_State285, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State287, _WildcardMarker}:                   _ReduceFieldDefToProperExplicitFieldDefsAction,
	{_State288, RparenToken}:                       _ReduceProperExplicitFieldDefsToExplicitFieldDefsAction,
	{_State290, RbracketToken}:                     _ReduceProperTypeArgumentsToTypeArgumentsAction,
	{_State292, _WildcardMarker}:                   _ReduceTypeExprToProperTypeArgumentsAction,
	{_State293, _WildcardMarker}:                   _ReduceToAccessExprAction,
	{_State295, _EndMarker}:                        _ReduceToBinaryOpAssignStatementAction,
	{_State296, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State296, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State297, _WildcardMarker}:                   _ReduceToBinaryAddExprAction,
	{_State298, _WildcardMarker}:                   _ReduceToBinaryAndExprAction,
	{_State299, _EndMarker}:                        _ReduceToAssignStatementAction,
	{_State300, _WildcardMarker}:                   _ReduceToBinaryCmpExprAction,
	{_State301, _WildcardMarker}:                   _ReduceAddToExprsAction,
	{_State303, _EndMarker}:                        _ReduceLabeledValuedToJumpStatementAction,
	{_State304, _WildcardMarker}:                   _ReduceToBinaryMulExprAction,
	{_State305, _WildcardMarker}:                   _ReduceInfiniteToLoopExprAction,
	{_State308, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State308, SemicolonToken}:                    _ReduceSequenceExprToForAssignmentAction,
	{_State309, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State311, LbraceToken}:                       _ReduceSequenceExprToConditionAction,
	{_State313, _WildcardMarker}:                   _ReduceToBinaryOrExprAction,
	{_State316, _WildcardMarker}:                   _ReduceFieldDefToEnumValueDefAction,
	{_State319, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State320, _WildcardMarker}:                   _ReduceParameterDeclToProperParameterDeclsAction,
	{_State322, RparenToken}:                       _ReduceProperParameterDeclsToParameterDeclsAction,
	{_State323, _WildcardMarker}:                   _ReduceUnamedToParameterDeclAction,
	{_State324, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State325, _WildcardMarker}:                   _ReduceToInferredTypeExprAction,
	{_State326, _WildcardMarker}:                   _ReduceExplicitToFieldDefAction,
	{_State330, _WildcardMarker}:                   _ReduceToImplicitEnumTypeExprAction,
	{_State331, _WildcardMarker}:                   _ReduceToImplicitStructTypeExprAction,
	{_State332, RparenToken}:                       _ReduceImproperToImplicitFieldDefsAction,
	{_State334, _WildcardMarker}:                   _ReduceFieldDefToTraitPropertyAction,
	{_State335, _WildcardMarker}:                   _ReduceMethodSignatureToTraitPropertyAction,
	{_State336, RparenToken}:                       _ReduceProperTraitPropertiesToTraitPropertiesAction,
	{_State338, _WildcardMarker}:                   _ReduceTraitPropertyToProperTraitPropertiesAction,
	{_State339, _WildcardMarker}:                   _ReduceToBinaryTypeExprAction,
	{_State340, _WildcardMarker}:                   _ReduceAliasToNamedFuncDefAction,
	{_State341, _WildcardMarker}:                   _ReduceUnconstrainedToGenericParameterDefAction,
	{_State342, _WildcardMarker}:                   _ReduceGenericParameterDefToProperGenericParameterDefsAction,
	{_State344, RbracketToken}:                     _ReduceProperGenericParameterDefsToGenericParameterDefsAction,
	{_State345, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State346, _WildcardMarker}:                   _ReduceInferredRefVarargToParameterDefAction,
	{_State347, _WildcardMarker}:                   _ReduceArgToParameterDefAction,
	{_State349, _WildcardMarker}:                   _ReduceAddImplicitToProperStatementsAction,
	{_State350, _WildcardMarker}:                   _ReduceAddExplicitToProperStatementsAction,
	{_State351, _WildcardMarker}:                   _ReduceAliasToTypeDefAction,
	{_State352, _WildcardMarker}:                   _ReduceDefinitionToTypeDefAction,
	{_State355, _WildcardMarker}:                   _ReduceToTuplePatternAction,
	{_State356, _WildcardMarker}:                   _ReduceEnumMatchPatternToCasePatternAction,
	{_State358, _EndMarker}:                        _ReduceCaseBranchStatementToStatementAction,
	{_State359, _WildcardMarker}:                   _ReduceMultipleToCasePatternsAction,
	{_State360, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State361, RparenToken}:                       _ReduceImproperToParameterDefsAction,
	{_State362, _EndMarker}:                        _ReduceMultipleToImportStatementAction,
	{_State363, RparenToken}:                       _ReduceExplicitToImportClausesAction,
	{_State364, RparenToken}:                       _ReduceImplicitToImportClausesAction,
	{_State365, _EndMarker}:                        _ReduceAliasToImportClauseAction,
	{_State368, _WildcardMarker}:                   _ReduceNamedAssignmentToArgumentAction,
	{_State369, _WildcardMarker}:                   _ReduceColonExprExprTupleToColonExprAction,
	{_State370, _WildcardMarker}:                   _ReduceExprExprPairToColonExprAction,
	{_State371, _WildcardMarker}:                   _ReduceAddToProperArgumentsAction,
	{_State372, _WildcardMarker}:                   _ReduceToExplicitStructTypeExprAction,
	{_State373, RparenToken}:                       _ReduceImproperExplicitToExplicitFieldDefsAction,
	{_State374, RparenToken}:                       _ReduceImproperImplicitToExplicitFieldDefsAction,
	{_State376, RbracketToken}:                     _ReduceImproperToTypeArgumentsAction,
	{_State377, _WildcardMarker}:                   _ReduceBindingToGenericTypeArgumentsAction,
	{_State378, _WildcardMarker}:                   _ReduceToIndexExprAction,
	{_State380, _WildcardMarker}:                   _ReduceToInitializeExprAction,
	{_State381, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State382, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State383, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State384, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State384, SemicolonToken}:                    _ReduceNilToOptionalSequenceExprAction,
	{_State387, _WildcardMarker}:                   _ReduceNoElseToIfExprAction,
	{_State388, _EndMarker}:                        _ReduceToSwitchExprAction,
	{_State391, _WildcardMarker}:                   _ReduceToExplicitEnumTypeExprAction,
	{_State394, _WildcardMarker}:                   _ReduceUnnamedVarargToParameterDeclAction,
	{_State396, _WildcardMarker}:                   _ReduceArgToParameterDeclAction,
	{_State397, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State398, RparenToken}:                       _ReduceImproperToParameterDeclsAction,
	{_State399, _WildcardMarker}:                   _ReduceExternalToNamedTypeExprAction,
	{_State400, _WildcardMarker}:                   _ReducePairToImplicitEnumValueDefsAction,
	{_State401, _WildcardMarker}:                   _ReduceDefaultToEnumValueDefAction,
	{_State402, _WildcardMarker}:                   _ReduceAddToImplicitEnumValueDefsAction,
	{_State403, _WildcardMarker}:                   _ReduceAddToProperImplicitFieldDefsAction,
	{_State405, RparenToken}:                       _ReduceImproperExplicitToTraitPropertiesAction,
	{_State406, RparenToken}:                       _ReduceImproperImplicitToTraitPropertiesAction,
	{_State407, _WildcardMarker}:                   _ReduceToTraitTypeExprAction,
	{_State408, _WildcardMarker}:                   _ReduceConstrainedToGenericParameterDefAction,
	{_State409, _WildcardMarker}:                   _ReduceGenericToOptionalGenericParametersAction,
	{_State410, RbracketToken}:                     _ReduceImproperToGenericParameterDefsAction,
	{_State412, _WildcardMarker}:                   _ReduceVarargToParameterDefAction,
	{_State413, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State415, _WildcardMarker}:                   _ReduceNamedToFieldVarPatternAction,
	{_State416, _WildcardMarker}:                   _ReduceAddToFieldVarPatternsAction,
	{_State417, _WildcardMarker}:                   _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State419, _WildcardMarker}:                   _ReduceReturnableTypeExprToReturnTypeAction,
	{_State420, _WildcardMarker}:                   _ReduceAddToProperParameterDefsAction,
	{_State421, _WildcardMarker}:                   _ReduceAddExplicitToProperImportClausesAction,
	{_State422, _WildcardMarker}:                   _ReduceAddImplicitToProperImportClausesAction,
	{_State423, _WildcardMarker}:                   _ReduceToMapTypeExprAction,
	{_State424, _WildcardMarker}:                   _ReduceToArrayTypeExprAction,
	{_State425, _WildcardMarker}:                   _ReduceAddExplicitToProperExplicitFieldDefsAction,
	{_State426, _WildcardMarker}:                   _ReduceAddImplicitToProperExplicitFieldDefsAction,
	{_State427, _EndMarker}:                        _ReduceToUnsafeStatementAction,
	{_State428, _WildcardMarker}:                   _ReduceAddToProperTypeArgumentsAction,
	{_State429, _WildcardMarker}:                   _ReduceToCallExprAction,
	{_State430, _EndMarker}:                        _ReduceDoWhileToLoopExprAction,
	{_State431, SemicolonToken}:                    _ReduceAssignToForAssignmentAction,
	{_State434, DoToken}:                           _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State435, _EndMarker}:                        _ReduceWhileToLoopExprAction,
	{_State436, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State438, RparenToken}:                       _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State439, _WildcardMarker}:                   _ReducePairToImplicitEnumValueDefsAction,
	{_State439, RparenToken}:                       _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State440, RparenToken}:                       _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State441, _WildcardMarker}:                   _ReduceAddToImplicitEnumValueDefsAction,
	{_State441, RparenToken}:                       _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State442, _WildcardMarker}:                   _ReduceVarargToParameterDeclAction,
	{_State443, _WildcardMarker}:                   _ReduceToFuncTypeExprAction,
	{_State444, _WildcardMarker}:                   _ReduceAddToProperParameterDeclsAction,
	{_State445, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State446, _WildcardMarker}:                   _ReduceExplicitToProperTraitPropertiesAction,
	{_State447, _WildcardMarker}:                   _ReduceImplicitToProperTraitPropertiesAction,
	{_State448, _WildcardMarker}:                   _ReduceAddToProperGenericParameterDefsAction,
	{_State449, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State451, _WildcardMarker}:                   _ReduceConstrainedDefToTypeDefAction,
	{_State452, _WildcardMarker}:                   _ReduceToAnonymousFuncExprAction,
	{_State454, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State454, DoToken}:                           _ReduceNilToOptionalSequenceExprAction,
	{_State455, LbraceToken}:                       _ReduceCaseToConditionAction,
	{_State456, _EndMarker}:                        _ReduceMultiIfElseToIfExprAction,
	{_State457, _EndMarker}:                        _ReduceIfElseToIfExprAction,
	{_State460, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State461, _EndMarker}:                        _ReduceIteratorToLoopExprAction,
	{_State463, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State464, _WildcardMarker}:                   _ReduceFuncDefToNamedFuncDefAction,
	{_State467, _WildcardMarker}:                   _ReduceToMethodSignatureAction,
	{_State468, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State469, _EndMarker}:                        _ReduceForToLoopExprAction,
	{_State471, _WildcardMarker}:                   _ReduceMethodDefToNamedFuncDefAction,
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
      expr_or_improper_struct_statement -> State 76
      exprs -> State 77
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
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

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
      expr -> State 7
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 4:
    Kernel Items:
      #accept: ^.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 8
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

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
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

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
      IDENTIFIER -> State 132
      LPAREN -> State 133

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
      proper_statements -> State 134
      statements -> State 136
      simple_statement -> State 99
      statement -> State 135
      expr_or_improper_struct_statement -> State 76
      exprs -> State 77
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
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      statement_block -> State 137

  State 14:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters type_expr
      type_def: TYPE.IDENTIFIER optional_generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE.IDENTIFIER ASSIGN type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 138

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
      NEWLINES -> State 139

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
      ASSIGN -> State 140

  State 24:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      LPAREN -> State 142
      var_pattern -> State 144
      tuple_pattern -> State 143

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
      VAR -> State 146
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 145
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 149
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 147
      case_pattern -> State 148
      optional_label_decl -> State 150
      sequence_expr -> State 151
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      COLON -> State 152

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
      LPAREN -> State 153

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
      optional_label_decl -> State 150
      sequence_expr -> State 154
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      STRING_LITERAL -> State 156
      LPAREN -> State 155
      import_clause -> State 157

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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 158
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

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
      IDENTIFIER -> State 161
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
      COLON -> State 159
      ELLIPSIS -> State 160
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 165
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      proper_arguments -> State 166
      arguments -> State 163
      argument -> State 162
      colon_expr -> State 164
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      explicit_struct_type_expr: STRUCT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 167

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
      LESS -> State 168

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
      LBRACKET -> State 180
      DOT -> State 179
      QUESTION -> State 183
      DOLLAR_LBRACKET -> State 178
      ADD_ASSIGN -> State 169
      SUB_ASSIGN -> State 184
      MUL_ASSIGN -> State 182
      DIV_ASSIGN -> State 177
      MOD_ASSIGN -> State 181
      ADD_ONE_ASSIGN -> State 170
      SUB_ONE_ASSIGN -> State 185
      BIT_NEG_ASSIGN -> State 173
      BIT_AND_ASSIGN -> State 171
      BIT_OR_ASSIGN -> State 174
      BIT_XOR_ASSIGN -> State 176
      BIT_LSHIFT_ASSIGN -> State 172
      BIT_RSHIFT_ASSIGN -> State 175
      unary_op_assign -> State 188
      binary_op_assign -> State 186
      generic_type_arguments -> State 187

  State 56:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 189
      SUB -> State 192
      BIT_XOR -> State 191
      BIT_OR -> State 190
      add_op -> State 193

  State 57:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 194

  State 58:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      initializable_type_expr: array_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expr
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 195

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
      optional_label_decl -> State 150
      call_expr -> State 197
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 196
      access_expr -> State 54
      index_expr -> State 82
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      EQUAL -> State 198
      NOT_EQUAL -> State 203
      LESS -> State 201
      LESS_OR_EQUAL -> State 202
      GREATER -> State 199
      GREATER_OR_EQUAL -> State 200
      cmp_op -> State 204

  State 74:
    Kernel Items:
      initializable_type_expr: explicit_struct_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      exprs: expr., *
    Reduce:
      * -> [exprs]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      simple_statement: expr_or_improper_struct_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      expr_or_improper_struct_statement: exprs., *
      exprs: exprs.COMMA expr
    Reduce:
      * -> [expr_or_improper_struct_statement]
    Goto:
      COMMA -> State 205

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
      initialize_expr: initializable_type_expr.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 206

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
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      JUMP_LABEL -> State 207
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      exprs -> State 208
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      MUL -> State 214
      DIV -> State 212
      MOD -> State 213
      BIT_AND -> State 209
      BIT_LSHIFT -> State 210
      BIT_RSHIFT -> State 211
      mul_op -> State 215

  State 90:
    Kernel Items:
      expr: optional_label_decl.if_expr
      expr: optional_label_decl.switch_expr
      expr: optional_label_decl.loop_expr
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 218
      SWITCH -> State 219
      FOR -> State 217
      DO -> State 216
      LBRACE -> State 11
      statement_block -> State 222
      if_expr -> State 220
      switch_expr -> State 223
      loop_expr -> State 221

  State 91:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 224

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
      optional_label_decl -> State 150
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
      prefixable_expr -> State 225
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
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
      expr: sequence_expr., *
    Reduce:
      * -> [expr]
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
      initializable_type_expr: slice_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
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
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.QUESTION
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 180
      DOT -> State 179
      QUESTION -> State 183
      DOLLAR_LBRACKET -> State 178
      generic_type_arguments -> State 187

  State 105:
    Kernel Items:
      expr: sequence_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      prefix_unary_type_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      prefix_unary_type_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      explicit_enum_type_expr: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 226

  State 110:
    Kernel Items:
      prefix_unary_type_op: EXCLAIM., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 227

  State 112:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOT -> State 228
      DOLLAR_LBRACKET -> State 178
      generic_type_arguments -> State 229

  State 113:
    Kernel Items:
      implicit_struct_type_expr: LPAREN.implicit_field_defs RPAREN
      implicit_enum_type_expr: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 232
      proper_implicit_field_defs -> State 235
      implicit_field_defs -> State 234
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 231
      implicit_enum_value_defs -> State 233
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 114:
    Kernel Items:
      parse_error_type_expr: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_type_expr]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      prefix_unary_type_op: QUESTION., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      prefix_unary_type_op: TILDE_TILDE., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      trait_type_expr: TRAIT.LPAREN trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 238

  State 118:
    Kernel Items:
      returnable_type_expr: atom_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      type_expr: binary_type_expr., *
    Reduce:
      * -> [type_expr]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      atom_type_expr: explicit_enum_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      atom_type_expr: func_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      atom_type_expr: implicit_enum_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      atom_type_expr: implicit_struct_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      atom_type_expr: inferred_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      atom_type_expr: initializable_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      atom_type_expr: named_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      atom_type_expr: parse_error_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      prefixed_unary_type_expr: prefix_unary_type_op.returnable_type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 239
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 129:
    Kernel Items:
      returnable_type_expr: prefixed_unary_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      type_expr: returnable_type_expr., *
    Reduce:
      * -> [type_expr]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      atom_type_expr: trait_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expr
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 245
      ASSIGN -> State 244
      optional_generic_parameters -> State 246

  State 133:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 247
      parameter_def -> State 248

  State 134:
    Kernel Items:
      proper_statements: proper_statements.NEWLINES statement
      proper_statements: proper_statements.SEMICOLON statement
      statements: proper_statements., RBRACE
      statements: proper_statements.NEWLINES
      statements: proper_statements.SEMICOLON
    Reduce:
      RBRACE -> [statements]
    Goto:
      NEWLINES -> State 249
      SEMICOLON -> State 250

  State 135:
    Kernel Items:
      proper_statements: statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 251

  State 137:
    Kernel Items:
      package_def: PACKAGE statement_block., *
    Reduce:
      * -> [package_def]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters type_expr
      type_def: TYPE IDENTIFIER.optional_generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 245
      ASSIGN -> State 252
      optional_generic_parameters -> State 253

  State 139:
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
      definition -> State 254
      statement_block -> State 21
      var_decl_pattern -> State 23
      var_or_let -> State 24
      type_def -> State 22
      named_func_def -> State 18
      package_def -> State 19

  State 140:
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
      expr -> State 255
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 141:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 142:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 257
      LPAREN -> State 142
      ELLIPSIS -> State 256
      var_pattern -> State 260
      tuple_pattern -> State 143
      field_var_patterns -> State 259
      field_var_pattern -> State 258

  State 143:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 144:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_type_expr
    Reduce:
      * -> [optional_type_expr]
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
      optional_type_expr -> State 261
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 262
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 145:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 263

  State 146:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 264

  State 147:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 148:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 149:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 266
      COLON -> State 265

  State 150:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 222

  State 151:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 152:
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
      simple_statement -> State 268
      optional_simple_statement -> State 267
      expr_or_improper_struct_statement -> State 76
      exprs -> State 77
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
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 153:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 247
      parameter_def -> State 269
      proper_parameter_defs -> State 271
      parameter_defs -> State 270

  State 154:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 156
      import_clause -> State 272
      proper_import_clauses -> State 274
      import_clauses -> State 273

  State 156:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 275

  State 157:
    Kernel Items:
      import_statement: IMPORT import_clause., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr.RBRACKET
      array_type_expr: LBRACKET type_expr.COMMA INTEGER_LITERAL RBRACKET
      map_type_expr: LBRACKET type_expr.COLON type_expr RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 278
      COMMA -> State 277
      COLON -> State 276
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 159:
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
      expr -> State 279
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 160:
    Kernel Items:
      argument: ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expr
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      ASSIGN -> State 280

  State 162:
    Kernel Items:
      proper_arguments: argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 281

  State 164:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 282

  State 165:
    Kernel Items:
      argument: expr., *
      argument: expr.ELLIPSIS
      colon_expr: expr.COLON
      colon_expr: expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 283
      ELLIPSIS -> State 284

  State 166:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., RPAREN
      arguments: proper_arguments.COMMA
    Reduce:
      RPAREN -> [arguments]
    Goto:
      COMMA -> State 285

  State 167:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 287
      implicit_struct_type_expr -> State 123
      proper_explicit_field_defs -> State 288
      explicit_field_defs -> State 286
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 168:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 289

  State 169:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 171:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET.type_arguments RBRACKET
    Reduce:
      RBRACKET -> [type_arguments]
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
      proper_type_arguments -> State 290
      type_arguments -> State 291
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 292
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 179:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 293

  State 180:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 161
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
      COLON -> State 159
      ELLIPSIS -> State 160
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 165
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      argument -> State 294
      colon_expr -> State 164
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 181:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      postfix_unary_expr: accessible_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 186:
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
      expr -> State 295
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 187:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 296

  State 188:
    Kernel Items:
      unary_op_assign_statement: accessible_expr unary_op_assign., $
    Reduce:
      $ -> [unary_op_assign_statement]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 193:
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
      optional_label_decl -> State 150
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
      mul_expr -> State 297
      binary_mul_expr -> State 66
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 194:
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
      optional_label_decl -> State 150
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
      cmp_expr -> State 298
      binary_cmp_expr -> State 65
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 195:
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
      expr -> State 299
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 196:
    Kernel Items:
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 180
      DOT -> State 179
      DOLLAR_LBRACKET -> State 178
      generic_type_arguments -> State 187

  State 197:
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

  State 198:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 204:
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
      optional_label_decl -> State 150
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
      add_expr -> State 300
      binary_add_expr -> State 63
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 205:
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
      expr -> State 301
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 206:
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
      IDENTIFIER -> State 161
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
      COLON -> State 159
      ELLIPSIS -> State 160
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 165
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      proper_arguments -> State 166
      arguments -> State 302
      argument -> State 162
      colon_expr -> State 164
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 207:
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
      exprs -> State 303
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 208:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 205

  State 209:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 215:
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
      optional_label_decl -> State 150
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
      prefixable_expr -> State 304
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 216:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 305

  State 217:
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
      assign_pattern -> State 306
      optional_label_decl -> State 150
      sequence_expr -> State 308
      for_assignment -> State 307
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 218:
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
      CASE -> State 309
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
      optional_label_decl -> State 150
      sequence_expr -> State 311
      condition -> State 310
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 219:
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
      optional_label_decl -> State 150
      sequence_expr -> State 312
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 220:
    Kernel Items:
      expr: optional_label_decl if_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      expr: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      block_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      expr: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 224:
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
      optional_label_decl -> State 150
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
      and_expr -> State 313
      binary_and_expr -> State 64
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 225:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefixable_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 314
      implicit_enum_value_defs -> State 317
      implicit_enum_type_expr -> State 122
      explicit_enum_value_defs -> State 315
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 227:
    Kernel Items:
      func_type_expr: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 319
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 318
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 323
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      parameter_decl -> State 320
      proper_parameter_decls -> State 322
      parameter_decls -> State 321
      func_type_expr -> State 121

  State 228:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 324

  State 229:
    Kernel Items:
      named_type_expr: IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      field_def: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 325
      QUESTION -> State 115
      EXCLAIM -> State 110
      DOLLAR_LBRACKET -> State 178
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      generic_type_arguments -> State 229
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 326
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 231:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 327

  State 232:
    Kernel Items:
      proper_implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [proper_implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 328

  State 233:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      OR -> State 329
      RPAREN -> State 330

  State 234:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 331

  State 235:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., RPAREN
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      COMMA -> State 332

  State 236:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: type_expr., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 237:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN.trait_properties RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 230
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 333
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 334
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_property -> State 338
      proper_trait_properties -> State 336
      trait_properties -> State 337
      trait_type_expr -> State 131
      func_type_expr -> State 121
      method_signature -> State 335

  State 239:
    Kernel Items:
      prefixed_unary_type_expr: prefix_unary_type_op returnable_type_expr., *
    Reduce:
      * -> [prefixed_unary_type_expr]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      binary_type_op: ADD., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      binary_type_op: MUL., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      binary_type_op: SUB., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op.returnable_type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 339
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 244:
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
      expr -> State 340
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 245:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 341
      generic_parameter_def -> State 342
      proper_generic_parameter_defs -> State 344
      generic_parameter_defs -> State 343

  State 246:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 345

  State 247:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.ELLIPSIS
      parameter_def: IDENTIFIER.type_expr
      parameter_def: IDENTIFIER.ELLIPSIS type_expr
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
      ELLIPSIS -> State 346
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 347
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 248:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 348

  State 249:
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
      statement -> State 349
      expr_or_improper_struct_statement -> State 76
      exprs -> State 77
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
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 250:
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
      statement -> State 350
      expr_or_improper_struct_statement -> State 76
      exprs -> State 77
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
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 251:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 351
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 253:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.type_expr
      type_def: TYPE IDENTIFIER optional_generic_parameters.type_expr IMPLEMENTS type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 352
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 254:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES definition., *
    Reduce:
      * -> [proper_definitions]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expr., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      field_var_pattern: ELLIPSIS., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 353

  State 258:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 259:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 355
      COMMA -> State 354

  State 260:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_type_expr., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      optional_type_expr: type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [optional_type_expr]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 263:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 43
      implicit_struct_expr -> State 356

  State 264:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 357

  State 265:
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
      simple_statement -> State 268
      optional_simple_statement -> State 358
      expr_or_improper_struct_statement -> State 76
      exprs -> State 77
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
      expr -> State 75
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 266:
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
      VAR -> State 146
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 145
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 147
      case_pattern -> State 359
      optional_label_decl -> State 150
      sequence_expr -> State 151
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 267:
    Kernel Items:
      statement: DEFAULT COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      optional_simple_statement: simple_statement., $
    Reduce:
      $ -> [optional_simple_statement]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      proper_parameter_defs: parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 360

  State 271:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., RPAREN
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      COMMA -> State 361

  State 272:
    Kernel Items:
      proper_import_clauses: import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 362

  State 274:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., RPAREN
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      NEWLINES -> State 364
      COMMA -> State 363

  State 275:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 365

  State 276:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON.type_expr RBRACKET
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 366
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 277:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 367

  State 278:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr RBRACKET., *
    Reduce:
      * -> [slice_type_expr]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      colon_expr: COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 280:
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
      expr -> State 368
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 281:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 282:
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
      expr -> State 369
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 283:
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
      expr -> State 370
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 284:
    Kernel Items:
      argument: expr ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 285:
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
      IDENTIFIER -> State 161
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
      COLON -> State 159
      ELLIPSIS -> State 160
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 165
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      argument -> State 371
      colon_expr -> State 164
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 286:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 372

  State 287:
    Kernel Items:
      proper_explicit_field_defs: field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., RPAREN
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      NEWLINES -> State 374
      COMMA -> State 373

  State 289:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 375

  State 290:
    Kernel Items:
      proper_type_arguments: proper_type_arguments.COMMA type_expr
      type_arguments: proper_type_arguments., RBRACKET
      type_arguments: proper_type_arguments.COMMA
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      COMMA -> State 376

  State 291:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 377

  State 292:
    Kernel Items:
      proper_type_arguments: type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 293:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 378

  State 295:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expr., $
    Reduce:
      $ -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 296:
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
      IDENTIFIER -> State 161
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
      COLON -> State 159
      ELLIPSIS -> State 160
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expr -> State 165
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      proper_arguments -> State 166
      arguments -> State 379
      argument -> State 162
      colon_expr -> State 164
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 297:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      binary_add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [binary_add_expr]
    Goto:
      MUL -> State 214
      DIV -> State 212
      MOD -> State 213
      BIT_AND -> State 209
      BIT_LSHIFT -> State 210
      BIT_RSHIFT -> State 211
      mul_op -> State 215

  State 298:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      binary_and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [binary_and_expr]
    Goto:
      EQUAL -> State 198
      NOT_EQUAL -> State 203
      LESS -> State 201
      LESS_OR_EQUAL -> State 202
      GREATER -> State 199
      GREATER_OR_EQUAL -> State 200
      cmp_op -> State 204

  State 299:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expr., $
    Reduce:
      $ -> [assign_statement]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      binary_cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [binary_cmp_expr]
    Goto:
      ADD -> State 189
      SUB -> State 192
      BIT_XOR -> State 191
      BIT_OR -> State 190
      add_op -> State 193

  State 301:
    Kernel Items:
      exprs: exprs COMMA expr., *
    Reduce:
      * -> [exprs]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 380

  State 303:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type JUMP_LABEL exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 205

  State 304:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 381

  State 306:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 383
      ASSIGN -> State 382

  State 307:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 384

  State 308:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 385

  State 309:
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
      VAR -> State 146
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 145
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 386
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 147
      case_pattern -> State 148
      optional_label_decl -> State 150
      sequence_expr -> State 151
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 310:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 387

  State 311:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 388

  State 313:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 194

  State 314:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 389
      OR -> State 390

  State 315:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 391

  State 316:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 328

  State 317:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 392
      OR -> State 393

  State 318:
    Kernel Items:
      parameter_decl: ELLIPSIS.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 394
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 319:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      parameter_decl: IDENTIFIER.type_expr
      parameter_decl: IDENTIFIER.ELLIPSIS type_expr
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 325
      QUESTION -> State 115
      EXCLAIM -> State 110
      DOLLAR_LBRACKET -> State 178
      ELLIPSIS -> State 395
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      generic_type_arguments -> State 229
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 396
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 320:
    Kernel Items:
      proper_parameter_decls: parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 397

  State 322:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., RPAREN
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      COMMA -> State 398

  State 323:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 324:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOLLAR_LBRACKET -> State 178
      generic_type_arguments -> State 399

  State 325:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 324

  State 326:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 327:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 400
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 328:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 401

  State 329:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 402
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 330:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA.field_def
      implicit_field_defs: proper_implicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 403
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 333:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 404
      LPAREN -> State 227

  State 334:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 336:
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

  State 337:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 407

  State 338:
    Kernel Items:
      proper_trait_properties: trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op returnable_type_expr., *
    Reduce:
      * -> [binary_type_expr]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 408
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 342:
    Kernel Items:
      proper_generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 409

  State 344:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs.COMMA generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs., RBRACKET
      generic_parameter_defs: proper_generic_parameter_defs.COMMA
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      COMMA -> State 410

  State 345:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 247
      parameter_def -> State 269
      proper_parameter_defs -> State 271
      parameter_defs -> State 411

  State 346:
    Kernel Items:
      parameter_def: IDENTIFIER ELLIPSIS., *
      parameter_def: IDENTIFIER ELLIPSIS.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 412
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 347:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 348:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 413

  State 349:
    Kernel Items:
      proper_statements: proper_statements NEWLINES statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER ASSIGN type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 352:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr., *
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr.IMPLEMENTS type_expr
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 414
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 353:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      LPAREN -> State 142
      var_pattern -> State 415
      tuple_pattern -> State 143

  State 354:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 257
      LPAREN -> State 142
      ELLIPSIS -> State 256
      var_pattern -> State 260
      tuple_pattern -> State 143
      field_var_pattern -> State 416

  State 355:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 142
      tuple_pattern -> State 417

  State 358:
    Kernel Items:
      statement: CASE case_patterns COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 360:
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 419
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      return_type -> State 418
      func_type_expr -> State 121

  State 361:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 247
      parameter_def -> State 420

  State 362:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 156
      import_clause -> State 421

  State 364:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 156
      import_clause -> State 422

  State 365:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 366:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr.RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 423
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 367:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 424

  State 368:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      colon_expr: colon_expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      colon_expr: expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      proper_arguments: proper_arguments COMMA argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_type_expr]
    Goto:
      (nil)

  State 373:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA.field_def
      explicit_field_defs: proper_explicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 425
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 374:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES.field_def
      explicit_field_defs: proper_explicit_field_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 426
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 375:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 427

  State 376:
    Kernel Items:
      proper_type_arguments: proper_type_arguments COMMA.type_expr
      type_arguments: proper_type_arguments COMMA., RBRACKET
    Reduce:
      RBRACKET -> [type_arguments]
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 428
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 377:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments RBRACKET., *
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 429

  State 380:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 381:
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
      optional_label_decl -> State 150
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 382:
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
      optional_label_decl -> State 150
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 383:
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
      optional_label_decl -> State 150
      sequence_expr -> State 432
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 384:
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
      optional_label_decl -> State 150
      sequence_expr -> State 434
      optional_sequence_expr -> State 433
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 385:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 435

  State 386:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 266
      ASSIGN -> State 436

  State 387:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 437

  State 388:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 389:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 438
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 390:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 439
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 391:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 440
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 393:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 230
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
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 316
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      enum_value_def -> State 441
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 394:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 395:
    Kernel Items:
      parameter_decl: IDENTIFIER ELLIPSIS.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 442
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 396:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: IDENTIFIER type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 397:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN.return_type
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 419
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      return_type -> State 443
      func_type_expr -> State 121

  State 398:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 319
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 318
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 323
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      parameter_decl -> State 444
      func_type_expr -> State 121

  State 399:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
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
      IDENTIFIER -> State 230
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 333
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 334
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_property -> State 446
      trait_type_expr -> State 131
      func_type_expr -> State 121
      method_signature -> State 335

  State 406:
    Kernel Items:
      proper_trait_properties: proper_trait_properties NEWLINES.trait_property
      trait_properties: proper_trait_properties NEWLINES., RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 230
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 333
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 237
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 236
      binary_type_expr -> State 119
      field_def -> State 334
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_property -> State 447
      trait_type_expr -> State 131
      func_type_expr -> State 121
      method_signature -> State 335

  State 407:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN trait_properties RPAREN., *
    Reduce:
      * -> [trait_type_expr]
    Goto:
      (nil)

  State 408:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      generic_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 409:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 410:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA.generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 341
      generic_parameter_def -> State 448

  State 411:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 449

  State 412:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_def: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 413:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 245
      optional_generic_parameters -> State 450

  State 414:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr IMPLEMENTS.type_expr
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 451
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      func_type_expr -> State 121

  State 415:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 452

  State 419:
    Kernel Items:
      return_type: returnable_type_expr., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 422:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr RBRACKET., *
    Reduce:
      * -> [map_type_expr]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [array_type_expr]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., $
    Reduce:
      $ -> [unsafe_statement]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      proper_type_arguments: proper_type_arguments COMMA type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 429:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 453

  State 433:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 454

  State 434:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 436:
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
      optional_label_decl -> State 150
      sequence_expr -> State 455
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 437:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 218
      LBRACE -> State 11
      statement_block -> State 457
      if_expr -> State 456

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
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 443:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type_expr]
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
      IDENTIFIER -> State 319
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 318
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 130
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      type_expr -> State 323
      binary_type_expr -> State 119
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      parameter_decl -> State 320
      proper_parameter_decls -> State 322
      parameter_decls -> State 458
      func_type_expr -> State 121

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
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 449:
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 419
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      return_type -> State 459
      func_type_expr -> State 121

  State 450:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 460

  State 451:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr IMPLEMENTS type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 240
      SUB -> State 242
      MUL -> State 241
      binary_type_op -> State 243

  State 452:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 453:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 461

  State 454:
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
      optional_label_decl -> State 150
      sequence_expr -> State 434
      optional_sequence_expr -> State 462
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
      initializable_type_expr -> State 83
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 74
      anonymous_func_expr -> State 58

  State 455:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 457:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 458:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 463

  State 459:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 464

  State 460:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 247
      parameter_def -> State 269
      proper_parameter_defs -> State 271
      parameter_defs -> State 465

  State 461:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 462:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 466

  State 463:
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 419
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      return_type -> State 467
      func_type_expr -> State 121

  State 464:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 465:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 468

  State 466:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 469

  State 467:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 468:
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
      initializable_type_expr -> State 125
      slice_type_expr -> State 100
      array_type_expr -> State 59
      map_type_expr -> State 88
      atom_type_expr -> State 118
      named_type_expr -> State 126
      inferred_type_expr -> State 124
      parse_error_type_expr -> State 127
      returnable_type_expr -> State 419
      prefixed_unary_type_expr -> State 129
      prefix_unary_type_op -> State 128
      implicit_struct_type_expr -> State 123
      explicit_struct_type_expr -> State 74
      implicit_enum_type_expr -> State 122
      explicit_enum_type_expr -> State 120
      trait_type_expr -> State 131
      return_type -> State 470
      func_type_expr -> State 121

  State 469:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 470:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 471

  State 471:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

Number of states: 471
Number of shift actions: 4376
Number of reduce actions: 421
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 4759
Number of unoptimized shift actions: 42038
Number of unoptimized reduce actions: 31995
*/
