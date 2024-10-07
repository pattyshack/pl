// Auto-generated from source: grammar.lr

package lr

import (
	fmt "fmt"
	lexutil "github.com/pattyshack/gt/lexutil"
	ast "github.com/pattyshack/pl/ast"
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
	RepeatToken          = SymbolId(271)
	ForToken             = SymbolId(272)
	DoToken              = SymbolId(273)
	InToken              = SymbolId(274)
	SelectToken          = SymbolId(275)
	ReturnToken          = SymbolId(276)
	BreakToken           = SymbolId(277)
	ContinueToken        = SymbolId(278)
	FallthroughToken     = SymbolId(279)
	ImportToken          = SymbolId(280)
	UnsafeToken          = SymbolId(281)
	TypeToken            = SymbolId(282)
	ImplementsToken      = SymbolId(283)
	StructToken          = SymbolId(284)
	EnumToken            = SymbolId(285)
	TraitToken           = SymbolId(286)
	FuncToken            = SymbolId(287)
	AsyncToken           = SymbolId(288)
	DeferToken           = SymbolId(289)
	VarToken             = SymbolId(290)
	LetToken             = SymbolId(291)
	AsToken              = SymbolId(292)
	NotToken             = SymbolId(293)
	AndToken             = SymbolId(294)
	OrToken              = SymbolId(295)
	AtToken              = SymbolId(296)
	LbraceToken          = SymbolId(297)
	RbraceToken          = SymbolId(298)
	LparenToken          = SymbolId(299)
	RparenToken          = SymbolId(300)
	LbracketToken        = SymbolId(301)
	RbracketToken        = SymbolId(302)
	DotToken             = SymbolId(303)
	CommaToken           = SymbolId(304)
	QuestionToken        = SymbolId(305)
	SemicolonToken       = SymbolId(306)
	ColonToken           = SymbolId(307)
	ExclaimToken         = SymbolId(308)
	DollarLbracketToken  = SymbolId(309)
	EllipsisToken        = SymbolId(310)
	TildeToken           = SymbolId(311)
	TildeTildeToken      = SymbolId(312)
	AssignToken          = SymbolId(313)
	ArrowToken           = SymbolId(314)
	AddAssignToken       = SymbolId(315)
	SubAssignToken       = SymbolId(316)
	MulAssignToken       = SymbolId(317)
	DivAssignToken       = SymbolId(318)
	ModAssignToken       = SymbolId(319)
	AddOneAssignToken    = SymbolId(320)
	SubOneAssignToken    = SymbolId(321)
	BitAndAssignToken    = SymbolId(322)
	BitOrAssignToken     = SymbolId(323)
	BitXorAssignToken    = SymbolId(324)
	BitLshiftAssignToken = SymbolId(325)
	BitRshiftAssignToken = SymbolId(326)
	AddToken             = SymbolId(327)
	SubToken             = SymbolId(328)
	MulToken             = SymbolId(329)
	DivToken             = SymbolId(330)
	ModToken             = SymbolId(331)
	BitAndToken          = SymbolId(332)
	BitXorToken          = SymbolId(333)
	BitOrToken           = SymbolId(334)
	BitLshiftToken       = SymbolId(335)
	BitRshiftToken       = SymbolId(336)
	EqualToken           = SymbolId(337)
	NotEqualToken        = SymbolId(338)
	LessToken            = SymbolId(339)
	LessOrEqualToken     = SymbolId(340)
	GreaterToken         = SymbolId(341)
	GreaterOrEqualToken  = SymbolId(342)
	ParseErrorToken      = SymbolId(343)
)

type Location = lexutil.Location

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

type FloatingCommentReducer interface {
	// 95:31: floating_comment -> ...
	ToFloatingComment(CommentGroups_ CommentGroupsTok) (ast.Statement, error)
}

type BranchStmtReducer interface {
	// 107:2: branch_stmt -> case_branch: ...
	CaseBranchToBranchStmt(Case_ *TokenValue, CasePatterns_ *ast.ExpressionList, Colon_ *TokenValue, OptionalStatement_ ast.Statement) (ast.Statement, error)

	// 108:2: branch_stmt -> default_branch: ...
	DefaultBranchToBranchStmt(Default_ *TokenValue, Colon_ *TokenValue, OptionalStatement_ ast.Statement) (ast.Statement, error)
}

type UnsafeStmtReducer interface {
	// 117:2: unsafe_stmt -> ...
	ToUnsafeStmt(Unsafe_ *TokenValue, Less_ *TokenValue, Identifier_ *TokenValue, Greater_ *TokenValue, StringLiteral_ *TokenValue) (*ast.UnsafeStmt, error)
}

type JumpStmtReducer interface {
	// 124:2: jump_stmt -> unlabeled_no_value: ...
	UnlabeledNoValueToJumpStmt(JumpOp_ *TokenValue) (ast.Statement, error)

	// 125:2: jump_stmt -> unlabeled_valued: ...
	UnlabeledValuedToJumpStmt(JumpOp_ *TokenValue, ReturnableExpr_ ast.Expression) (ast.Statement, error)

	// 126:2: jump_stmt -> labeled_no_value: ...
	LabeledNoValueToJumpStmt(JumpOp_ *TokenValue, At_ *TokenValue, Identifier_ *TokenValue) (ast.Statement, error)

	// 127:2: jump_stmt -> labeled_valued: ...
	LabeledValuedToJumpStmt(JumpOp_ *TokenValue, At_ *TokenValue, Identifier_ *TokenValue, ReturnableExpr_ ast.Expression) (ast.Statement, error)

	// 130:2: jump_stmt -> FALLTHROUGH: ...
	FallthroughToJumpStmt(Fallthrough_ *TokenValue) (ast.Statement, error)
}

type AssignStmtReducer interface {
	// 145:26: assign_stmt -> ...
	ToAssignStmt(ReturnableExpr_ ast.Expression, Assign_ *TokenValue, ReturnableExpr_2 ast.Expression) (ast.Statement, error)
}

type ImportStmtReducer interface {
	// 152:2: import_stmt -> single: ...
	SingleToImportStmt(Import_ *TokenValue, ImportClause_ *ast.ImportClause) (ast.Statement, error)

	// 153:2: import_stmt -> multiple: ...
	MultipleToImportStmt(Import_ *TokenValue, Lparen_ *TokenValue, ImportClauses_ *ast.ImportClauseList, Rparen_ *TokenValue) (ast.Statement, error)
}

type ProperImportClausesReducer interface {
	// 156:2: proper_import_clauses -> add_implicit: ...
	AddImplicitToProperImportClauses(ProperImportClauses_ *ast.ImportClauseList, Newlines_ TokenCount, ImportClause_ *ast.ImportClause) (*ast.ImportClauseList, error)

	// 157:2: proper_import_clauses -> add_explicit: ...
	AddExplicitToProperImportClauses(ProperImportClauses_ *ast.ImportClauseList, Comma_ *TokenValue, ImportClause_ *ast.ImportClause) (*ast.ImportClauseList, error)

	// 158:2: proper_import_clauses -> import_clause: ...
	ImportClauseToProperImportClauses(ImportClause_ *ast.ImportClause) (*ast.ImportClauseList, error)
}

type ImportClausesReducer interface {

	// 162:2: import_clauses -> implicit: ...
	ImplicitToImportClauses(ProperImportClauses_ *ast.ImportClauseList, Newlines_ TokenCount) (*ast.ImportClauseList, error)

	// 163:2: import_clauses -> explicit: ...
	ExplicitToImportClauses(ProperImportClauses_ *ast.ImportClauseList, Comma_ *TokenValue) (*ast.ImportClauseList, error)
}

type ImportClauseReducer interface {
	// 166:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ *TokenValue) (*ast.ImportClause, error)

	// 167:2: import_clause -> alias: ...
	AliasToImportClause(Identifier_ *TokenValue, StringLiteral_ *TokenValue) (*ast.ImportClause, error)

	// 168:2: import_clause -> unusable_import: ...
	UnusableImportToImportClause(Underscore_ *TokenValue, StringLiteral_ *TokenValue) (*ast.ImportClause, error)

	// 169:2: import_clause -> import_to_local: ...
	ImportToLocalToImportClause(Dot_ *TokenValue, StringLiteral_ *TokenValue) (*ast.ImportClause, error)
}

type AddrDeclPatternReducer interface {
	// 179:2: addr_decl_pattern -> new_inferred: ...
	NewInferredToAddrDeclPattern(NewVarType_ *TokenValue, NewAddressable_ ast.Expression) (ast.Expression, error)

	// 180:2: addr_decl_pattern -> new_typed: ...
	NewTypedToAddrDeclPattern(NewVarType_ *TokenValue, NewAddressable_ ast.Expression, TypeExpr_ ast.TypeExpression) (ast.Expression, error)
}

type AssignToAddrPatternReducer interface {
	// 182:38: assign_to_addr_pattern -> ...
	ToAssignToAddrPattern(Greater_ *TokenValue, Expr_ ast.Expression) (ast.Expression, error)
}

type CasePatternsReducer interface {

	// 195:2: case_patterns -> ...
	ToCasePatterns(AssignSelectablePattern_ ast.Expression) (*ast.ExpressionList, error)
}

type AssignSelectablePatternReducer interface {
	// 201:2: assign_selectable_pattern -> ...
	ToAssignSelectablePattern(SwitchableCasePatterns_ *ast.ExpressionList, Assign_ *TokenValue, Expr_ ast.Expression) (ast.Expression, error)
}

type SwitchableCasePatternsReducer interface {
	// 204:2: switchable_case_patterns -> switchable_case_pattern: ...
	SwitchableCasePatternToSwitchableCasePatterns(SwitchableCasePattern_ ast.Expression) (*ast.ExpressionList, error)

	// 205:2: switchable_case_patterns -> add: ...
	AddToSwitchableCasePatterns(SwitchableCasePatterns_ *ast.ExpressionList, Comma_ *TokenValue, SwitchableCasePattern_ ast.Expression) (*ast.ExpressionList, error)
}

type EnumPatternReducer interface {
	// 237:2: enum_pattern -> named: ...
	NamedToEnumPattern(Dot_ *TokenValue, Identifier_ *TokenValue, ImplicitStructExpr_ ast.Expression) (ast.Expression, error)

	// 238:2: enum_pattern -> unnamed_struct: ...
	UnnamedStructToEnumPattern(Dot_ *TokenValue, ImplicitStructExpr_ ast.Expression) (ast.Expression, error)
}

type ParseErrorExprReducer interface {
	// 256:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ *ParseErrorSymbol) (ast.Expression, error)
}

type LiteralExprReducer interface {
	// 259:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ *TokenValue) (ast.Expression, error)

	// 260:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ *TokenValue) (ast.Expression, error)

	// 261:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ *TokenValue) (ast.Expression, error)

	// 262:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ *TokenValue) (ast.Expression, error)

	// 263:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ *TokenValue) (ast.Expression, error)

	// 264:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ *TokenValue) (ast.Expression, error)
}

type NamedExprReducer interface {
	// 267:2: named_expr -> IDENTIFIER: ...
	IdentifierToNamedExpr(Identifier_ *TokenValue) (ast.Expression, error)

	// 268:2: named_expr -> UNDERSCORE: ...
	UnderscoreToNamedExpr(Underscore_ *TokenValue) (ast.Expression, error)
}

type InitializeExprReducer interface {
	// 270:31: initialize_expr -> ...
	ToInitializeExpr(InitializableTypeExpr_ ast.TypeExpression, Lparen_ *TokenValue, Arguments_ *ast.ArgumentList, Rparen_ *TokenValue) (ast.Expression, error)
}

type ImplicitStructExprReducer interface {
	// 272:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ *TokenValue, Arguments_ *ast.ArgumentList, Rparen_ *TokenValue) (ast.Expression, error)
}

type AccessExprReducer interface {
	// 285:27: access_expr -> ...
	ToAccessExpr(AccessibleExpr_ ast.Expression, Dot_ *TokenValue, Identifier_ *TokenValue) (ast.Expression, error)
}

type IndexExprReducer interface {
	// 289:26: index_expr -> ...
	ToIndexExpr(AccessibleExpr_ ast.Expression, Lbracket_ *TokenValue, Index_ ast.Expression, Rbracket_ *TokenValue) (ast.Expression, error)
}

type AsExprReducer interface {
	// 296:23: as_expr -> ...
	ToAsExpr(AccessibleExpr_ ast.Expression, Dot_ *TokenValue, As_ *TokenValue, Lparen_ *TokenValue, TypeExpr_ ast.TypeExpression, Rparen_ *TokenValue) (ast.Expression, error)
}

type CallExprReducer interface {
	// 299:2: call_expr -> ...
	ToCallExpr(AccessibleExpr_ ast.Expression, GenericArguments_ *ast.TypeExpressionList, Lparen_ *TokenValue, Arguments_ *ast.ArgumentList, Rparen_ *TokenValue) (ast.Expression, error)
}

type ProperArgumentsReducer interface {
	// 302:2: proper_arguments -> add: ...
	AddToProperArguments(ProperArguments_ *ast.ArgumentList, Comma_ *TokenValue, Argument_ *ast.Argument) (*ast.ArgumentList, error)

	// 303:2: proper_arguments -> argument: ...
	ArgumentToProperArguments(Argument_ *ast.Argument) (*ast.ArgumentList, error)
}

type ArgumentsReducer interface {

	// 307:2: arguments -> improper: ...
	ImproperToArguments(ProperArguments_ *ast.ArgumentList, Comma_ *TokenValue) (*ast.ArgumentList, error)

	// 308:2: arguments -> nil: ...
	NilToArguments() (*ast.ArgumentList, error)
}

type ArgumentReducer interface {
	// 311:2: argument -> positional: ...
	PositionalToArgument(Expr_ ast.Expression) (*ast.Argument, error)

	// 312:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ *ast.ImplicitStructExpr) (*ast.Argument, error)

	// 313:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ *TokenValue, Assign_ *TokenValue, Expr_ ast.Expression) (*ast.Argument, error)

	// 317:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expr_ ast.Expression, Ellipsis_ *TokenValue) (*ast.Argument, error)

	// 320:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ *TokenValue) (*ast.Argument, error)
}

type ColonExprReducer interface {
	// 324:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ *TokenValue) (*ast.ImplicitStructExpr, error)

	// 325:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expr_ ast.Expression, Colon_ *TokenValue) (*ast.ImplicitStructExpr, error)

	// 326:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ *TokenValue, Expr_ ast.Expression) (*ast.ImplicitStructExpr, error)

	// 327:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expr_ ast.Expression, Colon_ *TokenValue, Expr_2 ast.Expression) (*ast.ImplicitStructExpr, error)

	// 328:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ *ast.ImplicitStructExpr, Colon_ *TokenValue) (*ast.ImplicitStructExpr, error)

	// 329:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ *ast.ImplicitStructExpr, Colon_ *TokenValue, Expr_ ast.Expression) (*ast.ImplicitStructExpr, error)
}

type PostfixUnaryExprReducer interface {
	// 339:34: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessibleExpr_ ast.Expression, PostfixUnaryOp_ *TokenValue) (ast.Expression, error)
}

type PrefixUnaryExprReducer interface {
	// 358:33: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ *TokenValue, PrefixableExpr_ ast.Expression) (ast.Expression, error)
}

type BinaryMulExprReducer interface {
	// 399:31: binary_mul_expr -> ...
	ToBinaryMulExpr(MulExpr_ ast.Expression, MulOp_ *TokenValue, PrefixableExpr_ ast.Expression) (ast.Expression, error)
}

type BinaryAddExprReducer interface {
	// 417:31: binary_add_expr -> ...
	ToBinaryAddExpr(AddExpr_ ast.Expression, AddOp_ *TokenValue, MulExpr_ ast.Expression) (ast.Expression, error)
}

type BinaryCmpExprReducer interface {
	// 433:31: binary_cmp_expr -> ...
	ToBinaryCmpExpr(CmpExpr_ ast.Expression, CmpOp_ *TokenValue, AddExpr_ ast.Expression) (ast.Expression, error)
}

type BinaryAndExprReducer interface {
	// 451:31: binary_and_expr -> ...
	ToBinaryAndExpr(AndExpr_ ast.Expression, And_ *TokenValue, CmpExpr_ ast.Expression) (ast.Expression, error)
}

type BinaryOrExprReducer interface {
	// 461:30: binary_or_expr -> ...
	ToBinaryOrExpr(OrExpr_ ast.Expression, Or_ *TokenValue, AndExpr_ ast.Expression) (ast.Expression, error)
}

type SendExprReducer interface {
	// 472:25: send_expr -> ...
	ToSendExpr(SendRecvExpr_ ast.Expression, Arrow_ *TokenValue, OrExpr_ ast.Expression) (ast.Expression, error)
}

type RecvExprReducer interface {
	// 474:25: recv_expr -> ...
	ToRecvExpr(Arrow_ *TokenValue, OrExpr_ ast.Expression) (ast.Expression, error)
}

type BinaryAssignOpExprReducer interface {
	// 487:2: binary_assign_op_expr -> ...
	ToBinaryAssignOpExpr(SendRecvExpr_ ast.Expression, BinaryAssignOp_ *TokenValue, SendRecvExpr_2 ast.Expression) (ast.Expression, error)
}

type ControlFlowExprReducer interface {

	// 514:2: control_flow_expr -> labelled: ...
	LabelledToControlFlowExpr(Identifier_ *TokenValue, At_ *TokenValue, UnlabelledControlFlowExpr_ ast.ControlFlowExpr) (ast.ControlFlowExpr, error)
}

type StatementsReducer interface {
	// 539:30: statements -> ...
	ToStatements(Lbrace_ *TokenValue, StatementList_ *ast.StatementList, Rbrace_ *TokenValue) (*ast.StatementsExpr, error)
}

type ProperStatementListReducer interface {
	// 542:2: proper_statement_list -> add_implicit: ...
	AddImplicitToProperStatementList(ProperStatementList_ *ast.StatementList, Newlines_ TokenCount, Statement_ ast.Statement) (*ast.StatementList, error)

	// 543:2: proper_statement_list -> add_explicit: ...
	AddExplicitToProperStatementList(ProperStatementList_ *ast.StatementList, Semicolon_ *TokenValue, Statement_ ast.Statement) (*ast.StatementList, error)

	// 544:2: proper_statement_list -> statement: ...
	StatementToProperStatementList(Statement_ ast.Statement) (*ast.StatementList, error)
}

type StatementListReducer interface {

	// 548:2: statement_list -> improper_implicit: ...
	ImproperImplicitToStatementList(ProperStatementList_ *ast.StatementList, Newlines_ TokenCount) (*ast.StatementList, error)

	// 549:2: statement_list -> improper_explicit: ...
	ImproperExplicitToStatementList(ProperStatementList_ *ast.StatementList, Semicolon_ *TokenValue) (*ast.StatementList, error)

	// 550:2: statement_list -> nil: ...
	NilToStatementList() (*ast.StatementList, error)
}

type IfElseExprReducer interface {

	// 554:2: if_else_expr -> else: ...
	ElseToIfElseExpr(IfElifExpr_ *ast.IfExpr, Else_ *TokenValue, Statements_ *ast.StatementsExpr) (*ast.IfExpr, error)
}

type IfElifExprReducer interface {

	// 558:2: if_elif_expr -> elif: ...
	ElifToIfElifExpr(IfElifExpr_ *ast.IfExpr, Else_ *TokenValue, If_ *TokenValue, Condition_ ast.Expression, Statements_ *ast.StatementsExpr) (*ast.IfExpr, error)
}

type IfOnlyExprReducer interface {
	// 561:2: if_only_expr -> ...
	ToIfOnlyExpr(If_ *TokenValue, Condition_ ast.Expression, Statements_ *ast.StatementsExpr) (*ast.IfExpr, error)
}

type CasePatternExprReducer interface {
	// 567:33: case_pattern_expr -> ...
	ToCasePatternExpr(Case_ *TokenValue, SwitchableCasePatterns_ *ast.ExpressionList, Assign_ *TokenValue, Expr_ ast.Expression) (ast.Expression, error)
}

type SwitchExprBodyReducer interface {
	// 586:37: switch_expr_body -> ...
	ToSwitchExprBody(Switch_ *TokenValue, Expr_ ast.Expression, Statements_ *ast.StatementsExpr) (ast.ControlFlowExpr, error)
}

type SelectExprBodyReducer interface {
	// 588:37: select_expr_body -> ...
	ToSelectExprBody(Select_ *TokenValue, Statements_ *ast.StatementsExpr) (ast.ControlFlowExpr, error)
}

type LoopExprBodyReducer interface {
	// 591:2: loop_expr_body -> infinite: ...
	InfiniteToLoopExprBody(RepeatLoopBody_ *ast.StatementsExpr) (ast.ControlFlowExpr, error)

	// 592:2: loop_expr_body -> do_while: ...
	DoWhileToLoopExprBody(RepeatLoopBody_ *ast.StatementsExpr, For_ *TokenValue, Expr_ ast.Expression) (ast.ControlFlowExpr, error)

	// 593:2: loop_expr_body -> while: ...
	WhileToLoopExprBody(For_ *TokenValue, Expr_ ast.Expression, ForLoopBody_ *ast.StatementsExpr) (ast.ControlFlowExpr, error)

	// 594:2: loop_expr_body -> iterator: ...
	IteratorToLoopExprBody(For_ *TokenValue, ReturnableExpr_ ast.Expression, In_ *TokenValue, Expr_ ast.Expression, ForLoopBody_ *ast.StatementsExpr) (ast.ControlFlowExpr, error)

	// 595:2: loop_expr_body -> for: ...
	ForToLoopExprBody(For_ *TokenValue, OptionalStatement_ ast.Statement, Semicolon_ *TokenValue, OptionalExpr_ ast.Expression, Semicolon_2 *TokenValue, OptionalStatement_2 ast.Statement, ForLoopBody_ *ast.StatementsExpr) (ast.ControlFlowExpr, error)
}

type OptionalStatementReducer interface {

	// 599:2: optional_statement -> nil: ...
	NilToOptionalStatement() (ast.Statement, error)
}

type OptionalExprReducer interface {

	// 603:2: optional_expr -> nil: ...
	NilToOptionalExpr() (ast.Expression, error)
}

type RepeatLoopBodyReducer interface {
	// 605:36: repeat_loop_body -> ...
	ToRepeatLoopBody(Repeat_ *TokenValue, Statements_ *ast.StatementsExpr) (*ast.StatementsExpr, error)
}

type ForLoopBodyReducer interface {
	// 607:33: for_loop_body -> ...
	ToForLoopBody(Do_ *TokenValue, Statements_ *ast.StatementsExpr) (*ast.StatementsExpr, error)
}

type ImproperExprStructReducer interface {
	// 619:2: improper_expr_struct -> pair: ...
	PairToImproperExprStruct(Expr_ ast.Expression, Comma_ *TokenValue, Expr_2 ast.Expression) (*ast.ImplicitStructExpr, error)

	// 620:2: improper_expr_struct -> add: ...
	AddToImproperExprStruct(ImproperExprStruct_ *ast.ImplicitStructExpr, Comma_ *TokenValue, Expr_ ast.Expression) (*ast.ImplicitStructExpr, error)
}

type SliceTypeExprReducer interface {
	// 635:35: slice_type_expr -> ...
	ToSliceTypeExpr(Lbracket_ *TokenValue, TypeExpr_ ast.TypeExpression, Rbracket_ *TokenValue) (ast.TypeExpression, error)
}

type ArrayTypeExprReducer interface {
	// 638:2: array_type_expr -> ...
	ToArrayTypeExpr(Lbracket_ *TokenValue, TypeExpr_ ast.TypeExpression, Comma_ *TokenValue, IntegerLiteral_ *TokenValue, Rbracket_ *TokenValue) (ast.TypeExpression, error)
}

type MapTypeExprReducer interface {
	// 641:33: map_type_expr -> ...
	ToMapTypeExpr(Lbracket_ *TokenValue, TypeExpr_ ast.TypeExpression, Colon_ *TokenValue, TypeExpr_2 ast.TypeExpression, Rbracket_ *TokenValue) (ast.TypeExpression, error)
}

type NamedTypeExprReducer interface {
	// 654:2: named_type_expr -> local: ...
	LocalToNamedTypeExpr(Identifier_ *TokenValue, GenericArguments_ *ast.TypeExpressionList) (ast.TypeExpression, error)

	// 655:2: named_type_expr -> external: ...
	ExternalToNamedTypeExpr(Identifier_ *TokenValue, Dot_ *TokenValue, Identifier_2 *TokenValue, GenericArguments_ *ast.TypeExpressionList) (ast.TypeExpression, error)
}

type InferredTypeExprReducer interface {
	// 663:2: inferred_type_expr -> DOT: ...
	DotToInferredTypeExpr(Dot_ *TokenValue) (ast.TypeExpression, error)

	// 664:2: inferred_type_expr -> UNDERSCORE: ...
	UnderscoreToInferredTypeExpr(Underscore_ *TokenValue) (ast.TypeExpression, error)
}

type PrefixUnaryTypeExprReducer interface {
	// 674:2: prefix_unary_type_expr -> ...
	ToPrefixUnaryTypeExpr(PrefixUnaryTypeOp_ *TokenValue, ReturnableTypeExpr_ ast.TypeExpression) (ast.TypeExpression, error)
}

type BinaryTypeExprReducer interface {
	// 690:2: binary_type_expr -> ...
	ToBinaryTypeExpr(TypeExpr_ ast.TypeExpression, BinaryTypeOp_ *TokenValue, ReturnableTypeExpr_ ast.TypeExpression) (ast.TypeExpression, error)
}

type TypeDefReducer interface {
	// 698:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *TokenValue, Identifier_ *TokenValue, GenericParameters_ *ast.GenericParameterList, TypeExpr_ ast.TypeExpression) (ast.Statement, error)

	// 699:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *TokenValue, Identifier_ *TokenValue, GenericParameters_ *ast.GenericParameterList, TypeExpr_ ast.TypeExpression, Implements_ *TokenValue, TypeExpr_2 ast.TypeExpression) (ast.Statement, error)

	// 700:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *TokenValue, Identifier_ *TokenValue, Assign_ *TokenValue, TypeExpr_ ast.TypeExpression) (ast.Statement, error)
}

type GenericParameterReducer interface {
	// 708:2: generic_parameter -> unconstrained: ...
	UnconstrainedToGenericParameter(Identifier_ *TokenValue) (*ast.GenericParameter, error)

	// 709:2: generic_parameter -> constrained: ...
	ConstrainedToGenericParameter(Identifier_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.GenericParameter, error)
}

type GenericParametersReducer interface {
	// 712:2: generic_parameters -> generic: ...
	GenericToGenericParameters(DollarLbracket_ *TokenValue, GenericParameterList_ *ast.GenericParameterList, Rbracket_ *TokenValue) (*ast.GenericParameterList, error)

	// 713:2: generic_parameters -> nil: ...
	NilToGenericParameters() (*ast.GenericParameterList, error)
}

type ProperGenericParameterListReducer interface {
	// 716:2: proper_generic_parameter_list -> add: ...
	AddToProperGenericParameterList(ProperGenericParameterList_ *ast.GenericParameterList, Comma_ *TokenValue, GenericParameter_ *ast.GenericParameter) (*ast.GenericParameterList, error)

	// 717:2: proper_generic_parameter_list -> generic_parameter: ...
	GenericParameterToProperGenericParameterList(GenericParameter_ *ast.GenericParameter) (*ast.GenericParameterList, error)
}

type GenericParameterListReducer interface {

	// 721:2: generic_parameter_list -> improper: ...
	ImproperToGenericParameterList(ProperGenericParameterList_ *ast.GenericParameterList, Comma_ *TokenValue) (*ast.GenericParameterList, error)

	// 722:2: generic_parameter_list -> nil: ...
	NilToGenericParameterList() (*ast.GenericParameterList, error)
}

type GenericArgumentsReducer interface {
	// 725:2: generic_arguments -> binding: ...
	BindingToGenericArguments(DollarLbracket_ *TokenValue, GenericArgumentList_ *ast.TypeExpressionList, Rbracket_ *TokenValue) (*ast.TypeExpressionList, error)

	// 726:2: generic_arguments -> nil: ...
	NilToGenericArguments() (*ast.TypeExpressionList, error)
}

type ProperGenericArgumentListReducer interface {
	// 729:2: proper_generic_argument_list -> add: ...
	AddToProperGenericArgumentList(ProperGenericArgumentList_ *ast.TypeExpressionList, Comma_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.TypeExpressionList, error)

	// 730:2: proper_generic_argument_list -> type_expr: ...
	TypeExprToProperGenericArgumentList(TypeExpr_ ast.TypeExpression) (*ast.TypeExpressionList, error)
}

type GenericArgumentListReducer interface {

	// 734:2: generic_argument_list -> improper: ...
	ImproperToGenericArgumentList(ProperGenericArgumentList_ *ast.TypeExpressionList, Comma_ *TokenValue) (*ast.TypeExpressionList, error)

	// 735:2: generic_argument_list -> nil: ...
	NilToGenericArgumentList() (*ast.TypeExpressionList, error)
}

type TypePropertyReducer interface {
	// 749:2: type_property -> unnamed_field: ...
	UnnamedFieldToTypeProperty(TypeExpr_ ast.TypeExpression) (ast.TypeProperty, error)

	// 750:2: type_property -> named_field: ...
	NamedFieldToTypeProperty(Identifier_ *TokenValue, TypeExpr_ ast.TypeExpression) (ast.TypeProperty, error)

	// 751:2: type_property -> padding_field: ...
	PaddingFieldToTypeProperty(Underscore_ *TokenValue, TypeExpr_ ast.TypeExpression) (ast.TypeProperty, error)

	// 752:2: type_property -> default_named_enum_field: ...
	DefaultNamedEnumFieldToTypeProperty(Default_ *TokenValue, Identifier_ *TokenValue, TypeExpr_ ast.TypeExpression) (ast.TypeProperty, error)

	// 753:2: type_property -> default_unnamed_enum_field: ...
	DefaultUnnamedEnumFieldToTypeProperty(Default_ *TokenValue, TypeExpr_ ast.TypeExpression) (ast.TypeProperty, error)
}

type ProperImplicitTypePropertiesReducer interface {
	// 756:2: proper_implicit_type_properties -> add: ...
	AddToProperImplicitTypeProperties(ProperImplicitTypeProperties_ *ast.TypePropertyList, Comma_ *TokenValue, TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)

	// 757:2: proper_implicit_type_properties -> type_property: ...
	TypePropertyToProperImplicitTypeProperties(TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)
}

type ImplicitTypePropertiesReducer interface {

	// 761:2: implicit_type_properties -> improper: ...
	ImproperToImplicitTypeProperties(ProperImplicitTypeProperties_ *ast.TypePropertyList, Comma_ *TokenValue) (*ast.TypePropertyList, error)

	// 762:2: implicit_type_properties -> nil: ...
	NilToImplicitTypeProperties() (*ast.TypePropertyList, error)
}

type ImplicitStructTypeExprReducer interface {
	// 765:2: implicit_struct_type_expr -> ...
	ToImplicitStructTypeExpr(Lparen_ *TokenValue, ImplicitTypeProperties_ *ast.TypePropertyList, Rparen_ *TokenValue) (ast.TypeExpression, error)
}

type ProperExplicitTypePropertiesReducer interface {
	// 768:2: proper_explicit_type_properties -> add_implicit: ...
	AddImplicitToProperExplicitTypeProperties(ProperExplicitTypeProperties_ *ast.TypePropertyList, Newlines_ TokenCount, TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)

	// 769:2: proper_explicit_type_properties -> add_explicit: ...
	AddExplicitToProperExplicitTypeProperties(ProperExplicitTypeProperties_ *ast.TypePropertyList, Comma_ *TokenValue, TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)

	// 770:2: proper_explicit_type_properties -> type_property: ...
	TypePropertyToProperExplicitTypeProperties(TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)
}

type ExplicitTypePropertiesReducer interface {

	// 774:2: explicit_type_properties -> improper_implicit: ...
	ImproperImplicitToExplicitTypeProperties(ProperExplicitTypeProperties_ *ast.TypePropertyList, Newlines_ TokenCount) (*ast.TypePropertyList, error)

	// 775:2: explicit_type_properties -> improper_explicit: ...
	ImproperExplicitToExplicitTypeProperties(ProperExplicitTypeProperties_ *ast.TypePropertyList, Comma_ *TokenValue) (*ast.TypePropertyList, error)

	// 776:2: explicit_type_properties -> nil: ...
	NilToExplicitTypeProperties() (*ast.TypePropertyList, error)
}

type ExplicitStructTypeExprReducer interface {
	// 779:2: explicit_struct_type_expr -> ...
	ToExplicitStructTypeExpr(Struct_ *TokenValue, Lparen_ *TokenValue, ExplicitTypeProperties_ *ast.TypePropertyList, Rparen_ *TokenValue) (ast.TypeExpression, error)
}

type TraitTypeExprReducer interface {
	// 782:2: trait_type_expr -> ...
	ToTraitTypeExpr(Trait_ *TokenValue, Lparen_ *TokenValue, ExplicitTypeProperties_ *ast.TypePropertyList, Rparen_ *TokenValue) (ast.TypeExpression, error)
}

type ProperImplicitEnumTypePropertiesReducer interface {
	// 793:2: proper_implicit_enum_type_properties -> pair: ...
	PairToProperImplicitEnumTypeProperties(TypeProperty_ ast.TypeProperty, Or_ *TokenValue, TypeProperty_2 ast.TypeProperty) (*ast.TypePropertyList, error)

	// 794:2: proper_implicit_enum_type_properties -> add: ...
	AddToProperImplicitEnumTypeProperties(ProperImplicitEnumTypeProperties_ *ast.TypePropertyList, Or_ *TokenValue, TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)
}

type ImplicitEnumTypePropertiesReducer interface {

	// 799:2: implicit_enum_type_properties -> improper: ...
	ImproperToImplicitEnumTypeProperties(ProperImplicitEnumTypeProperties_ *ast.TypePropertyList, Newlines_ TokenCount) (*ast.TypePropertyList, error)
}

type ImplicitEnumTypeExprReducer interface {
	// 802:2: implicit_enum_type_expr -> ...
	ToImplicitEnumTypeExpr(Lparen_ *TokenValue, ImplicitEnumTypeProperties_ *ast.TypePropertyList, Rparen_ *TokenValue) (ast.TypeExpression, error)
}

type ProperExplicitEnumTypePropertiesReducer interface {
	// 805:2: proper_explicit_enum_type_properties -> explicit_pair: ...
	ExplicitPairToProperExplicitEnumTypeProperties(TypeProperty_ ast.TypeProperty, Or_ *TokenValue, TypeProperty_2 ast.TypeProperty) (*ast.TypePropertyList, error)

	// 806:2: proper_explicit_enum_type_properties -> implicit_pair: ...
	ImplicitPairToProperExplicitEnumTypeProperties(TypeProperty_ ast.TypeProperty, Newlines_ TokenCount, TypeProperty_2 ast.TypeProperty) (*ast.TypePropertyList, error)

	// 807:2: proper_explicit_enum_type_properties -> explicit_add: ...
	ExplicitAddToProperExplicitEnumTypeProperties(ProperExplicitEnumTypeProperties_ *ast.TypePropertyList, Or_ *TokenValue, TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)

	// 808:2: proper_explicit_enum_type_properties -> implicit_add: ...
	ImplicitAddToProperExplicitEnumTypeProperties(ProperExplicitEnumTypeProperties_ *ast.TypePropertyList, Newlines_ TokenCount, TypeProperty_ ast.TypeProperty) (*ast.TypePropertyList, error)
}

type ExplicitEnumTypePropertiesReducer interface {

	// 813:2: explicit_enum_type_properties -> improper: ...
	ImproperToExplicitEnumTypeProperties(ProperExplicitEnumTypeProperties_ *ast.TypePropertyList, Newlines_ TokenCount) (*ast.TypePropertyList, error)
}

type ExplicitEnumTypeExprReducer interface {
	// 816:2: explicit_enum_type_expr -> ...
	ToExplicitEnumTypeExpr(Enum_ *TokenValue, Lparen_ *TokenValue, ExplicitEnumTypeProperties_ *ast.TypePropertyList, Rparen_ *TokenValue) (ast.TypeExpression, error)
}

type ReturnTypeReducer interface {

	// 825:2: return_type -> nil: ...
	NilToReturnType() (ast.TypeExpression, error)
}

type ParameterReducer interface {
	// 830:2: parameter -> named_arg: ...
	NamedArgToParameter(Identifier_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 831:2: parameter -> named_receiver: ...
	NamedReceiverToParameter(Identifier_ *TokenValue, Greater_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 832:2: parameter -> named_vararg: ...
	NamedVarargToParameter(Identifier_ *TokenValue, Ellipsis_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 833:2: parameter -> ignore_arg: ...
	IgnoreArgToParameter(Underscore_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 834:2: parameter -> ignore_receiver: ...
	IgnoreReceiverToParameter(Underscore_ *TokenValue, Greater_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 835:2: parameter -> ignore_vararg: ...
	IgnoreVarargToParameter(Underscore_ *TokenValue, Ellipsis_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 836:2: parameter -> unnamed_arg: ...
	UnnamedArgToParameter(TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 837:2: parameter -> unnamed_receiver: ...
	UnnamedReceiverToParameter(Greater_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)

	// 838:2: parameter -> unnamed_vararg: ...
	UnnamedVarargToParameter(Ellipsis_ *TokenValue, TypeExpr_ ast.TypeExpression) (*ast.Parameter, error)
}

type ProperParameterListReducer interface {
	// 841:2: proper_parameter_list -> add: ...
	AddToProperParameterList(ProperParameterList_ *ast.ParameterList, Comma_ *TokenValue, Parameter_ *ast.Parameter) (*ast.ParameterList, error)

	// 842:2: proper_parameter_list -> parameter: ...
	ParameterToProperParameterList(Parameter_ *ast.Parameter) (*ast.ParameterList, error)
}

type ParameterListReducer interface {

	// 846:2: parameter_list -> improper: ...
	ImproperToParameterList(ProperParameterList_ *ast.ParameterList, Comma_ *TokenValue) (*ast.ParameterList, error)

	// 847:2: parameter_list -> nil: ...
	NilToParameterList() (*ast.ParameterList, error)
}

type ParametersReducer interface {
	// 849:26: parameters -> ...
	ToParameters(Lparen_ *TokenValue, ParameterList_ *ast.ParameterList, Rparen_ *TokenValue) (*ast.ParameterList, error)
}

type FuncSignatureReducer interface {
	// 856:2: func_signature -> anonymous: ...
	AnonymousToFuncSignature(Func_ *TokenValue, Parameters_ *ast.ParameterList, ReturnType_ ast.TypeExpression) (*ast.FuncSignature, error)

	// 857:2: func_signature -> named: ...
	NamedToFuncSignature(Func_ *TokenValue, Identifier_ *TokenValue, GenericParameters_ *ast.GenericParameterList, Parameters_ *ast.ParameterList, ReturnType_ ast.TypeExpression) (*ast.FuncSignature, error)
}

type FuncDefReducer interface {
	// 859:24: func_def -> ...
	ToFuncDef(FuncSignature_ *ast.FuncSignature, Statements_ *ast.StatementsExpr) (ast.Expression, error)
}

type Reducer interface {
	FloatingCommentReducer
	BranchStmtReducer
	UnsafeStmtReducer
	JumpStmtReducer
	AssignStmtReducer
	ImportStmtReducer
	ProperImportClausesReducer
	ImportClausesReducer
	ImportClauseReducer
	AddrDeclPatternReducer
	AssignToAddrPatternReducer
	CasePatternsReducer
	AssignSelectablePatternReducer
	SwitchableCasePatternsReducer
	EnumPatternReducer
	ParseErrorExprReducer
	LiteralExprReducer
	NamedExprReducer
	InitializeExprReducer
	ImplicitStructExprReducer
	AccessExprReducer
	IndexExprReducer
	AsExprReducer
	CallExprReducer
	ProperArgumentsReducer
	ArgumentsReducer
	ArgumentReducer
	ColonExprReducer
	PostfixUnaryExprReducer
	PrefixUnaryExprReducer
	BinaryMulExprReducer
	BinaryAddExprReducer
	BinaryCmpExprReducer
	BinaryAndExprReducer
	BinaryOrExprReducer
	SendExprReducer
	RecvExprReducer
	BinaryAssignOpExprReducer
	ControlFlowExprReducer
	StatementsReducer
	ProperStatementListReducer
	StatementListReducer
	IfElseExprReducer
	IfElifExprReducer
	IfOnlyExprReducer
	CasePatternExprReducer
	SwitchExprBodyReducer
	SelectExprBodyReducer
	LoopExprBodyReducer
	OptionalStatementReducer
	OptionalExprReducer
	RepeatLoopBodyReducer
	ForLoopBodyReducer
	ImproperExprStructReducer
	SliceTypeExprReducer
	ArrayTypeExprReducer
	MapTypeExprReducer
	NamedTypeExprReducer
	InferredTypeExprReducer
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
	ParameterReducer
	ProperParameterListReducer
	ParameterListReducer
	ParametersReducer
	FuncSignatureReducer
	FuncDefReducer
}

type ParseErrorHandler interface {
	Error(nextToken Token, parseStack _Stack) error
}

type DefaultParseErrorHandler struct{}

func (DefaultParseErrorHandler) Error(nextToken Token, stack _Stack) error {
	return lexutil.NewLocationError(
		nextToken.Loc(),
		"syntax error: unexpected symbol %s. expecting %v",
		nextToken.Id(),
		ExpectedTerminals(stack[len(stack)-1].StateId))
}

func ExpectedTerminals(id _StateId) []SymbolId {
	switch id {
	case _State2:
		return []SymbolId{CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, CaseToken, DefaultToken, RepeatToken, ForToken, SelectToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, ImportToken, UnsafeToken, TypeToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State3:
		return []SymbolId{LbraceToken}
	case _State4:
		return []SymbolId{_EndMarker}
	case _State5:
		return []SymbolId{_EndMarker}
	case _State6:
		return []SymbolId{_EndMarker}
	case _State7:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State8:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, DotToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State9:
		return []SymbolId{ColonToken}
	case _State11:
		return []SymbolId{IdentifierToken, LparenToken}
	case _State12:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State14:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, CaseToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State15:
		return []SymbolId{StringLiteralToken, IdentifierToken, UnderscoreToken, LparenToken, DotToken}
	case _State17:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State19:
		return []SymbolId{LbraceToken}
	case _State20:
		return []SymbolId{LbraceToken}
	case _State21:
		return []SymbolId{LparenToken}
	case _State22:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State23:
		return []SymbolId{IdentifierToken}
	case _State24:
		return []SymbolId{LessToken}
	case _State30:
		return []SymbolId{LbraceToken}
	case _State33:
		return []SymbolId{LparenToken}
	case _State36:
		return []SymbolId{IdentifierToken, UnderscoreToken, LparenToken}
	case _State38:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State44:
		return []SymbolId{IdentifierToken, LparenToken}
	case _State45:
		return []SymbolId{ColonToken}
	case _State49:
		return []SymbolId{SemicolonToken}
	case _State54:
		return []SymbolId{IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, LbraceToken}
	case _State55:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, DotToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State56:
		return []SymbolId{LbraceToken}
	case _State57:
		return []SymbolId{StringLiteralToken}
	case _State58:
		return []SymbolId{StringLiteralToken}
	case _State59:
		return []SymbolId{StringLiteralToken, IdentifierToken, UnderscoreToken, DotToken}
	case _State60:
		return []SymbolId{StringLiteralToken}
	case _State61:
		return []SymbolId{RbraceToken}
	case _State62:
		return []SymbolId{LparenToken}
	case _State65:
		return []SymbolId{LparenToken}
	case _State66:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State67:
		return []SymbolId{RbracketToken, CommaToken, ColonToken, AddToken, SubToken, MulToken}
	case _State70:
		return []SymbolId{RparenToken}
	case _State75:
		return []SymbolId{LbraceToken}
	case _State77:
		return []SymbolId{IdentifierToken}
	case _State79:
		return []SymbolId{IdentifierToken, AsToken}
	case _State80:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ColonToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State81:
		return []SymbolId{LparenToken}
	case _State82:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State83:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State84:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State85:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State86:
		return []SymbolId{IfToken, LbraceToken}
	case _State87:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State89:
		return []SymbolId{IdentifierToken}
	case _State90:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State92:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State95:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State96:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State97:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State98:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, StructToken, FuncToken, AsyncToken, DeferToken, NotToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, ParseErrorToken}
	case _State99:
		return []SymbolId{LparenToken}
	case _State101:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State102:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, DotToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State103:
		return []SymbolId{LbraceToken}
	case _State105:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State107:
		return []SymbolId{LparenToken}
	case _State108:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State109:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State112:
		return []SymbolId{RparenToken}
	case _State115:
		return []SymbolId{CommaToken, AssignToken}
	case _State116:
		return []SymbolId{RparenToken}
	case _State118:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State119:
		return []SymbolId{IdentifierToken}
	case _State120:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State123:
		return []SymbolId{RparenToken}
	case _State124:
		return []SymbolId{RparenToken}
	case _State130:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State131:
		return []SymbolId{IntegerLiteralToken}
	case _State132:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State133:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State137:
		return []SymbolId{RparenToken}
	case _State139:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State140:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State141:
		return []SymbolId{GreaterToken}
	case _State142:
		return []SymbolId{RbracketToken}
	case _State145:
		return []SymbolId{LparenToken}
	case _State148:
		return []SymbolId{RbracketToken}
	case _State153:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, CaseToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State154:
		return []SymbolId{RparenToken}
	case _State160:
		return []SymbolId{SemicolonToken}
	case _State161:
		return []SymbolId{DoToken}
	case _State163:
		return []SymbolId{RbracketToken}
	case _State169:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State170:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State172:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State173:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State176:
		return []SymbolId{IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, SwitchToken, RepeatToken, ForToken, SelectToken, StructToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, LbraceToken, LparenToken, LbracketToken, ArrowToken, AddToken, SubToken, MulToken, BitAndToken, BitXorToken, GreaterToken, ParseErrorToken}
	case _State179:
		return []SymbolId{RparenToken}
	case _State181:
		return []SymbolId{NewlinesToken, OrToken}
	case _State187:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State189:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State190:
		return []SymbolId{RparenToken}
	case _State191:
		return []SymbolId{RbracketToken, AddToken, SubToken, MulToken}
	case _State192:
		return []SymbolId{RbracketToken}
	case _State197:
		return []SymbolId{StringLiteralToken}
	case _State199:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State200:
		return []SymbolId{RparenToken}
	case _State201:
		return []SymbolId{LbraceToken}
	case _State210:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State211:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State212:
		return []SymbolId{IdentifierToken, UnderscoreToken, DefaultToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State214:
		return []SymbolId{IdentifierToken, UnderscoreToken, StructToken, EnumToken, TraitToken, FuncToken, LparenToken, LbracketToken, DotToken, QuestionToken, ExclaimToken, TildeToken, TildeTildeToken, BitAndToken}
	case _State216:
		return []SymbolId{RparenToken, AddToken, SubToken, MulToken}
	case _State217:
		return []SymbolId{DoToken}
	}

	return nil
}

func ParseSource(lexer Lexer, reducer Reducer) (*ast.StatementList, error) {

	return ParseSourceWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseSourceWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
) (
	*ast.StatementList,
	error,
) {
	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal *ast.StatementList
		return errRetVal, err
	}
	return item.StatementList, nil
}

func ParseStatement(lexer Lexer, reducer Reducer) (ast.Statement, error) {

	return ParseStatementWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseStatementWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
) (
	ast.Statement,
	error,
) {
	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal ast.Statement
		return errRetVal, err
	}
	return item.Statement, nil
}

func ParseStatements(lexer Lexer, reducer Reducer) (*ast.StatementsExpr, error) {

	return ParseStatementsWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseStatementsWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
) (
	*ast.StatementsExpr,
	error,
) {
	item, err := _Parse(lexer, reducer, errHandler, _State3)
	if err != nil {
		var errRetVal *ast.StatementsExpr
		return errRetVal, err
	}
	return item.StatementsExpr, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _Parse(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
	startState _StateId,
) (
	*_StackItem,
	error,
) {
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
	case RepeatToken:
		return "REPEAT"
	case ForToken:
		return "FOR"
	case DoToken:
		return "DO"
	case InToken:
		return "IN"
	case SelectToken:
		return "SELECT"
	case ReturnToken:
		return "RETURN"
	case BreakToken:
		return "BREAK"
	case ContinueToken:
		return "CONTINUE"
	case FallthroughToken:
		return "FALLTHROUGH"
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
	case AsToken:
		return "AS"
	case NotToken:
		return "NOT"
	case AndToken:
		return "AND"
	case OrToken:
		return "OR"
	case AtToken:
		return "AT"
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
	case TildeToken:
		return "TILDE"
	case TildeTildeToken:
		return "TILDE_TILDE"
	case AssignToken:
		return "ASSIGN"
	case ArrowToken:
		return "ARROW"
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
	case StatementType:
		return "statement"
	case FloatingCommentType:
		return "floating_comment"
	case BranchStmtType:
		return "branch_stmt"
	case UnsafeStmtType:
		return "unsafe_stmt"
	case JumpStmtType:
		return "jump_stmt"
	case JumpOpType:
		return "jump_op"
	case AssignStmtType:
		return "assign_stmt"
	case ImportStmtType:
		return "import_stmt"
	case ProperImportClausesType:
		return "proper_import_clauses"
	case ImportClausesType:
		return "import_clauses"
	case ImportClauseType:
		return "import_clause"
	case AddrDeclPatternType:
		return "addr_decl_pattern"
	case AssignToAddrPatternType:
		return "assign_to_addr_pattern"
	case NewVarTypeType:
		return "new_var_type"
	case NewAddressableType:
		return "new_addressable"
	case CasePatternsType:
		return "case_patterns"
	case AssignSelectablePatternType:
		return "assign_selectable_pattern"
	case SwitchableCasePatternsType:
		return "switchable_case_patterns"
	case SwitchableCasePatternType:
		return "switchable_case_pattern"
	case EnumPatternType:
		return "enum_pattern"
	case AtomExprType:
		return "atom_expr"
	case ParseErrorExprType:
		return "parse_error_expr"
	case LiteralExprType:
		return "literal_expr"
	case NamedExprType:
		return "named_expr"
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
	case IndexType:
		return "index"
	case AsExprType:
		return "as_expr"
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
	case PostfixableExprType:
		return "postfixable_expr"
	case PostfixUnaryExprType:
		return "postfix_unary_expr"
	case PostfixUnaryOpType:
		return "postfix_unary_op"
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
	case SendRecvExprType:
		return "send_recv_expr"
	case SendExprType:
		return "send_expr"
	case RecvExprType:
		return "recv_expr"
	case AssignOpExprType:
		return "assign_op_expr"
	case BinaryAssignOpExprType:
		return "binary_assign_op_expr"
	case BinaryAssignOpType:
		return "binary_assign_op"
	case UnlabelledControlFlowExprType:
		return "unlabelled_control_flow_expr"
	case ControlFlowExprType:
		return "control_flow_expr"
	case ExprType:
		return "expr"
	case StatementsType:
		return "statements"
	case ProperStatementListType:
		return "proper_statement_list"
	case StatementListType:
		return "statement_list"
	case IfElseExprType:
		return "if_else_expr"
	case IfElifExprType:
		return "if_elif_expr"
	case IfOnlyExprType:
		return "if_only_expr"
	case ConditionType:
		return "condition"
	case CasePatternExprType:
		return "case_pattern_expr"
	case SwitchExprBodyType:
		return "switch_expr_body"
	case SelectExprBodyType:
		return "select_expr_body"
	case LoopExprBodyType:
		return "loop_expr_body"
	case OptionalStatementType:
		return "optional_statement"
	case OptionalExprType:
		return "optional_expr"
	case RepeatLoopBodyType:
		return "repeat_loop_body"
	case ForLoopBodyType:
		return "for_loop_body"
	case ReturnableExprType:
		return "returnable_expr"
	case ImproperExprStructType:
		return "improper_expr_struct"
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
	case ParameterType:
		return "parameter"
	case ProperParameterListType:
		return "proper_parameter_list"
	case ParameterListType:
		return "parameter_list"
	case ParametersType:
		return "parameters"
	case FuncSignatureType:
		return "func_signature"
	case FuncDefType:
		return "func_def"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	SourceType                           = SymbolId(344)
	StatementType                        = SymbolId(345)
	FloatingCommentType                  = SymbolId(346)
	BranchStmtType                       = SymbolId(347)
	UnsafeStmtType                       = SymbolId(348)
	JumpStmtType                         = SymbolId(349)
	JumpOpType                           = SymbolId(350)
	AssignStmtType                       = SymbolId(351)
	ImportStmtType                       = SymbolId(352)
	ProperImportClausesType              = SymbolId(353)
	ImportClausesType                    = SymbolId(354)
	ImportClauseType                     = SymbolId(355)
	AddrDeclPatternType                  = SymbolId(356)
	AssignToAddrPatternType              = SymbolId(357)
	NewVarTypeType                       = SymbolId(358)
	NewAddressableType                   = SymbolId(359)
	CasePatternsType                     = SymbolId(360)
	AssignSelectablePatternType          = SymbolId(361)
	SwitchableCasePatternsType           = SymbolId(362)
	SwitchableCasePatternType            = SymbolId(363)
	EnumPatternType                      = SymbolId(364)
	AtomExprType                         = SymbolId(365)
	ParseErrorExprType                   = SymbolId(366)
	LiteralExprType                      = SymbolId(367)
	NamedExprType                        = SymbolId(368)
	InitializeExprType                   = SymbolId(369)
	ImplicitStructExprType               = SymbolId(370)
	AccessibleExprType                   = SymbolId(371)
	AccessExprType                       = SymbolId(372)
	IndexExprType                        = SymbolId(373)
	IndexType                            = SymbolId(374)
	AsExprType                           = SymbolId(375)
	CallExprType                         = SymbolId(376)
	ProperArgumentsType                  = SymbolId(377)
	ArgumentsType                        = SymbolId(378)
	ArgumentType                         = SymbolId(379)
	ColonExprType                        = SymbolId(380)
	PostfixableExprType                  = SymbolId(381)
	PostfixUnaryExprType                 = SymbolId(382)
	PostfixUnaryOpType                   = SymbolId(383)
	PrefixableExprType                   = SymbolId(384)
	PrefixUnaryExprType                  = SymbolId(385)
	PrefixUnaryOpType                    = SymbolId(386)
	MulExprType                          = SymbolId(387)
	BinaryMulExprType                    = SymbolId(388)
	MulOpType                            = SymbolId(389)
	AddExprType                          = SymbolId(390)
	BinaryAddExprType                    = SymbolId(391)
	AddOpType                            = SymbolId(392)
	CmpExprType                          = SymbolId(393)
	BinaryCmpExprType                    = SymbolId(394)
	CmpOpType                            = SymbolId(395)
	AndExprType                          = SymbolId(396)
	BinaryAndExprType                    = SymbolId(397)
	OrExprType                           = SymbolId(398)
	BinaryOrExprType                     = SymbolId(399)
	SendRecvExprType                     = SymbolId(400)
	SendExprType                         = SymbolId(401)
	RecvExprType                         = SymbolId(402)
	AssignOpExprType                     = SymbolId(403)
	BinaryAssignOpExprType               = SymbolId(404)
	BinaryAssignOpType                   = SymbolId(405)
	UnlabelledControlFlowExprType        = SymbolId(406)
	ControlFlowExprType                  = SymbolId(407)
	ExprType                             = SymbolId(408)
	StatementsType                       = SymbolId(409)
	ProperStatementListType              = SymbolId(410)
	StatementListType                    = SymbolId(411)
	IfElseExprType                       = SymbolId(412)
	IfElifExprType                       = SymbolId(413)
	IfOnlyExprType                       = SymbolId(414)
	ConditionType                        = SymbolId(415)
	CasePatternExprType                  = SymbolId(416)
	SwitchExprBodyType                   = SymbolId(417)
	SelectExprBodyType                   = SymbolId(418)
	LoopExprBodyType                     = SymbolId(419)
	OptionalStatementType                = SymbolId(420)
	OptionalExprType                     = SymbolId(421)
	RepeatLoopBodyType                   = SymbolId(422)
	ForLoopBodyType                      = SymbolId(423)
	ReturnableExprType                   = SymbolId(424)
	ImproperExprStructType               = SymbolId(425)
	InitializableTypeExprType            = SymbolId(426)
	SliceTypeExprType                    = SymbolId(427)
	ArrayTypeExprType                    = SymbolId(428)
	MapTypeExprType                      = SymbolId(429)
	AtomTypeExprType                     = SymbolId(430)
	NamedTypeExprType                    = SymbolId(431)
	InferredTypeExprType                 = SymbolId(432)
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
	TypePropertyType                     = SymbolId(447)
	ProperImplicitTypePropertiesType     = SymbolId(448)
	ImplicitTypePropertiesType           = SymbolId(449)
	ImplicitStructTypeExprType           = SymbolId(450)
	ProperExplicitTypePropertiesType     = SymbolId(451)
	ExplicitTypePropertiesType           = SymbolId(452)
	ExplicitStructTypeExprType           = SymbolId(453)
	TraitTypeExprType                    = SymbolId(454)
	ProperImplicitEnumTypePropertiesType = SymbolId(455)
	ImplicitEnumTypePropertiesType       = SymbolId(456)
	ImplicitEnumTypeExprType             = SymbolId(457)
	ProperExplicitEnumTypePropertiesType = SymbolId(458)
	ExplicitEnumTypePropertiesType       = SymbolId(459)
	ExplicitEnumTypeExprType             = SymbolId(460)
	ReturnTypeType                       = SymbolId(461)
	ParameterType                        = SymbolId(462)
	ProperParameterListType              = SymbolId(463)
	ParameterListType                    = SymbolId(464)
	ParametersType                       = SymbolId(465)
	FuncSignatureType                    = SymbolId(466)
	FuncDefType                          = SymbolId(467)
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
	_ReduceStatementListToSource                                        = _ReduceType(1)
	_ReduceUnsafeStmtToStatement                                        = _ReduceType(2)
	_ReduceImportStmtToStatement                                        = _ReduceType(3)
	_ReduceFloatingCommentToStatement                                   = _ReduceType(4)
	_ReduceTypeDefToStatement                                           = _ReduceType(5)
	_ReduceBranchStmtToStatement                                        = _ReduceType(6)
	_ReduceReturnableExprToStatement                                    = _ReduceType(7)
	_ReduceJumpStmtToStatement                                          = _ReduceType(8)
	_ReduceAssignStmtToStatement                                        = _ReduceType(9)
	_ReduceToFloatingComment                                            = _ReduceType(10)
	_ReduceCaseBranchToBranchStmt                                       = _ReduceType(11)
	_ReduceDefaultBranchToBranchStmt                                    = _ReduceType(12)
	_ReduceToUnsafeStmt                                                 = _ReduceType(13)
	_ReduceUnlabeledNoValueToJumpStmt                                   = _ReduceType(14)
	_ReduceUnlabeledValuedToJumpStmt                                    = _ReduceType(15)
	_ReduceLabeledNoValueToJumpStmt                                     = _ReduceType(16)
	_ReduceLabeledValuedToJumpStmt                                      = _ReduceType(17)
	_ReduceFallthroughToJumpStmt                                        = _ReduceType(18)
	_ReduceReturnToJumpOp                                               = _ReduceType(19)
	_ReduceBreakToJumpOp                                                = _ReduceType(20)
	_ReduceContinueToJumpOp                                             = _ReduceType(21)
	_ReduceToAssignStmt                                                 = _ReduceType(22)
	_ReduceSingleToImportStmt                                           = _ReduceType(23)
	_ReduceMultipleToImportStmt                                         = _ReduceType(24)
	_ReduceAddImplicitToProperImportClauses                             = _ReduceType(25)
	_ReduceAddExplicitToProperImportClauses                             = _ReduceType(26)
	_ReduceImportClauseToProperImportClauses                            = _ReduceType(27)
	_ReduceProperImportClausesToImportClauses                           = _ReduceType(28)
	_ReduceImplicitToImportClauses                                      = _ReduceType(29)
	_ReduceExplicitToImportClauses                                      = _ReduceType(30)
	_ReduceStringLiteralToImportClause                                  = _ReduceType(31)
	_ReduceAliasToImportClause                                          = _ReduceType(32)
	_ReduceUnusableImportToImportClause                                 = _ReduceType(33)
	_ReduceImportToLocalToImportClause                                  = _ReduceType(34)
	_ReduceNewInferredToAddrDeclPattern                                 = _ReduceType(35)
	_ReduceNewTypedToAddrDeclPattern                                    = _ReduceType(36)
	_ReduceToAssignToAddrPattern                                        = _ReduceType(37)
	_ReduceVarToNewVarType                                              = _ReduceType(38)
	_ReduceLetToNewVarType                                              = _ReduceType(39)
	_ReduceNamedExprToNewAddressable                                    = _ReduceType(40)
	_ReduceImplicitStructExprToNewAddressable                           = _ReduceType(41)
	_ReduceSwitchableCasePatternsToCasePatterns                         = _ReduceType(42)
	_ReduceToCasePatterns                                               = _ReduceType(43)
	_ReduceToAssignSelectablePattern                                    = _ReduceType(44)
	_ReduceSwitchableCasePatternToSwitchableCasePatterns                = _ReduceType(45)
	_ReduceAddToSwitchableCasePatterns                                  = _ReduceType(46)
	_ReduceExprToSwitchableCasePattern                                  = _ReduceType(47)
	_ReduceEnumPatternToSwitchableCasePattern                           = _ReduceType(48)
	_ReduceNamedToEnumPattern                                           = _ReduceType(49)
	_ReduceUnnamedStructToEnumPattern                                   = _ReduceType(50)
	_ReduceParseErrorExprToAtomExpr                                     = _ReduceType(51)
	_ReduceLiteralExprToAtomExpr                                        = _ReduceType(52)
	_ReduceNamedExprToAtomExpr                                          = _ReduceType(53)
	_ReduceFuncDefToAtomExpr                                            = _ReduceType(54)
	_ReduceInitializeExprToAtomExpr                                     = _ReduceType(55)
	_ReduceImplicitStructExprToAtomExpr                                 = _ReduceType(56)
	_ReduceToParseErrorExpr                                             = _ReduceType(57)
	_ReduceTrueToLiteralExpr                                            = _ReduceType(58)
	_ReduceFalseToLiteralExpr                                           = _ReduceType(59)
	_ReduceIntegerLiteralToLiteralExpr                                  = _ReduceType(60)
	_ReduceFloatLiteralToLiteralExpr                                    = _ReduceType(61)
	_ReduceRuneLiteralToLiteralExpr                                     = _ReduceType(62)
	_ReduceStringLiteralToLiteralExpr                                   = _ReduceType(63)
	_ReduceIdentifierToNamedExpr                                        = _ReduceType(64)
	_ReduceUnderscoreToNamedExpr                                        = _ReduceType(65)
	_ReduceToInitializeExpr                                             = _ReduceType(66)
	_ReduceToImplicitStructExpr                                         = _ReduceType(67)
	_ReduceAtomExprToAccessibleExpr                                     = _ReduceType(68)
	_ReduceAccessExprToAccessibleExpr                                   = _ReduceType(69)
	_ReduceCallExprToAccessibleExpr                                     = _ReduceType(70)
	_ReduceIndexExprToAccessibleExpr                                    = _ReduceType(71)
	_ReduceAsExprToAccessibleExpr                                       = _ReduceType(72)
	_ReduceToAccessExpr                                                 = _ReduceType(73)
	_ReduceToIndexExpr                                                  = _ReduceType(74)
	_ReduceExprToIndex                                                  = _ReduceType(75)
	_ReduceColonExprToIndex                                             = _ReduceType(76)
	_ReduceToAsExpr                                                     = _ReduceType(77)
	_ReduceToCallExpr                                                   = _ReduceType(78)
	_ReduceAddToProperArguments                                         = _ReduceType(79)
	_ReduceArgumentToProperArguments                                    = _ReduceType(80)
	_ReduceProperArgumentsToArguments                                   = _ReduceType(81)
	_ReduceImproperToArguments                                          = _ReduceType(82)
	_ReduceNilToArguments                                               = _ReduceType(83)
	_ReducePositionalToArgument                                         = _ReduceType(84)
	_ReduceColonExprToArgument                                          = _ReduceType(85)
	_ReduceNamedAssignmentToArgument                                    = _ReduceType(86)
	_ReduceVarargAssignmentToArgument                                   = _ReduceType(87)
	_ReduceSkipPatternToArgument                                        = _ReduceType(88)
	_ReduceUnitUnitPairToColonExpr                                      = _ReduceType(89)
	_ReduceExprUnitPairToColonExpr                                      = _ReduceType(90)
	_ReduceUnitExprPairToColonExpr                                      = _ReduceType(91)
	_ReduceExprExprPairToColonExpr                                      = _ReduceType(92)
	_ReduceColonExprUnitTupleToColonExpr                                = _ReduceType(93)
	_ReduceColonExprExprTupleToColonExpr                                = _ReduceType(94)
	_ReduceAccessibleExprToPostfixableExpr                              = _ReduceType(95)
	_ReducePostfixUnaryExprToPostfixableExpr                            = _ReduceType(96)
	_ReduceToPostfixUnaryExpr                                           = _ReduceType(97)
	_ReduceQuestionToPostfixUnaryOp                                     = _ReduceType(98)
	_ReduceExclaimToPostfixUnaryOp                                      = _ReduceType(99)
	_ReduceAddOneAssignToPostfixUnaryOp                                 = _ReduceType(100)
	_ReduceSubOneAssignToPostfixUnaryOp                                 = _ReduceType(101)
	_ReducePostfixableExprToPrefixableExpr                              = _ReduceType(102)
	_ReducePrefixUnaryExprToPrefixableExpr                              = _ReduceType(103)
	_ReduceToPrefixUnaryExpr                                            = _ReduceType(104)
	_ReduceNotToPrefixUnaryOp                                           = _ReduceType(105)
	_ReduceBitXorToPrefixUnaryOp                                        = _ReduceType(106)
	_ReduceAddToPrefixUnaryOp                                           = _ReduceType(107)
	_ReduceSubToPrefixUnaryOp                                           = _ReduceType(108)
	_ReduceMulToPrefixUnaryOp                                           = _ReduceType(109)
	_ReduceBitAndToPrefixUnaryOp                                        = _ReduceType(110)
	_ReduceAsyncToPrefixUnaryOp                                         = _ReduceType(111)
	_ReduceDeferToPrefixUnaryOp                                         = _ReduceType(112)
	_ReducePrefixableExprToMulExpr                                      = _ReduceType(113)
	_ReduceBinaryMulExprToMulExpr                                       = _ReduceType(114)
	_ReduceToBinaryMulExpr                                              = _ReduceType(115)
	_ReduceMulToMulOp                                                   = _ReduceType(116)
	_ReduceDivToMulOp                                                   = _ReduceType(117)
	_ReduceModToMulOp                                                   = _ReduceType(118)
	_ReduceBitAndToMulOp                                                = _ReduceType(119)
	_ReduceBitLshiftToMulOp                                             = _ReduceType(120)
	_ReduceBitRshiftToMulOp                                             = _ReduceType(121)
	_ReduceMulExprToAddExpr                                             = _ReduceType(122)
	_ReduceBinaryAddExprToAddExpr                                       = _ReduceType(123)
	_ReduceToBinaryAddExpr                                              = _ReduceType(124)
	_ReduceAddToAddOp                                                   = _ReduceType(125)
	_ReduceSubToAddOp                                                   = _ReduceType(126)
	_ReduceBitOrToAddOp                                                 = _ReduceType(127)
	_ReduceBitXorToAddOp                                                = _ReduceType(128)
	_ReduceAddExprToCmpExpr                                             = _ReduceType(129)
	_ReduceBinaryCmpExprToCmpExpr                                       = _ReduceType(130)
	_ReduceToBinaryCmpExpr                                              = _ReduceType(131)
	_ReduceEqualToCmpOp                                                 = _ReduceType(132)
	_ReduceNotEqualToCmpOp                                              = _ReduceType(133)
	_ReduceLessToCmpOp                                                  = _ReduceType(134)
	_ReduceLessOrEqualToCmpOp                                           = _ReduceType(135)
	_ReduceGreaterToCmpOp                                               = _ReduceType(136)
	_ReduceGreaterOrEqualToCmpOp                                        = _ReduceType(137)
	_ReduceCmpExprToAndExpr                                             = _ReduceType(138)
	_ReduceBinaryAndExprToAndExpr                                       = _ReduceType(139)
	_ReduceToBinaryAndExpr                                              = _ReduceType(140)
	_ReduceAndExprToOrExpr                                              = _ReduceType(141)
	_ReduceBinaryOrExprToOrExpr                                         = _ReduceType(142)
	_ReduceToBinaryOrExpr                                               = _ReduceType(143)
	_ReduceOrExprToSendRecvExpr                                         = _ReduceType(144)
	_ReduceSendExprToSendRecvExpr                                       = _ReduceType(145)
	_ReduceRecvExprToSendRecvExpr                                       = _ReduceType(146)
	_ReduceToSendExpr                                                   = _ReduceType(147)
	_ReduceToRecvExpr                                                   = _ReduceType(148)
	_ReduceSendRecvExprToAssignOpExpr                                   = _ReduceType(149)
	_ReduceBinaryAssignOpExprToAssignOpExpr                             = _ReduceType(150)
	_ReduceToBinaryAssignOpExpr                                         = _ReduceType(151)
	_ReduceAddAssignToBinaryAssignOp                                    = _ReduceType(152)
	_ReduceSubAssignToBinaryAssignOp                                    = _ReduceType(153)
	_ReduceMulAssignToBinaryAssignOp                                    = _ReduceType(154)
	_ReduceDivAssignToBinaryAssignOp                                    = _ReduceType(155)
	_ReduceModAssignToBinaryAssignOp                                    = _ReduceType(156)
	_ReduceBitAndAssignToBinaryAssignOp                                 = _ReduceType(157)
	_ReduceBitOrAssignToBinaryAssignOp                                  = _ReduceType(158)
	_ReduceBitXorAssignToBinaryAssignOp                                 = _ReduceType(159)
	_ReduceBitLshiftAssignToBinaryAssignOp                              = _ReduceType(160)
	_ReduceBitRshiftAssignToBinaryAssignOp                              = _ReduceType(161)
	_ReduceStatementsToUnlabelledControlFlowExpr                        = _ReduceType(162)
	_ReduceIfElseExprToUnlabelledControlFlowExpr                        = _ReduceType(163)
	_ReduceSwitchExprBodyToUnlabelledControlFlowExpr                    = _ReduceType(164)
	_ReduceSelectExprBodyToUnlabelledControlFlowExpr                    = _ReduceType(165)
	_ReduceLoopExprBodyToUnlabelledControlFlowExpr                      = _ReduceType(166)
	_ReduceUnlabelledControlFlowExprToControlFlowExpr                   = _ReduceType(167)
	_ReduceLabelledToControlFlowExpr                                    = _ReduceType(168)
	_ReduceAssignOpExprToExpr                                           = _ReduceType(169)
	_ReduceControlFlowExprToExpr                                        = _ReduceType(170)
	_ReduceAddrDeclPatternToExpr                                        = _ReduceType(171)
	_ReduceAssignToAddrPatternToExpr                                    = _ReduceType(172)
	_ReduceToStatements                                                 = _ReduceType(173)
	_ReduceAddImplicitToProperStatementList                             = _ReduceType(174)
	_ReduceAddExplicitToProperStatementList                             = _ReduceType(175)
	_ReduceStatementToProperStatementList                               = _ReduceType(176)
	_ReduceProperStatementListToStatementList                           = _ReduceType(177)
	_ReduceImproperImplicitToStatementList                              = _ReduceType(178)
	_ReduceImproperExplicitToStatementList                              = _ReduceType(179)
	_ReduceNilToStatementList                                           = _ReduceType(180)
	_ReduceIfElifExprToIfElseExpr                                       = _ReduceType(181)
	_ReduceElseToIfElseExpr                                             = _ReduceType(182)
	_ReduceIfOnlyExprToIfElifExpr                                       = _ReduceType(183)
	_ReduceElifToIfElifExpr                                             = _ReduceType(184)
	_ReduceToIfOnlyExpr                                                 = _ReduceType(185)
	_ReduceExprToCondition                                              = _ReduceType(186)
	_ReduceCasePatternExprToCondition                                   = _ReduceType(187)
	_ReduceToCasePatternExpr                                            = _ReduceType(188)
	_ReduceToSwitchExprBody                                             = _ReduceType(189)
	_ReduceToSelectExprBody                                             = _ReduceType(190)
	_ReduceInfiniteToLoopExprBody                                       = _ReduceType(191)
	_ReduceDoWhileToLoopExprBody                                        = _ReduceType(192)
	_ReduceWhileToLoopExprBody                                          = _ReduceType(193)
	_ReduceIteratorToLoopExprBody                                       = _ReduceType(194)
	_ReduceForToLoopExprBody                                            = _ReduceType(195)
	_ReduceStatementToOptionalStatement                                 = _ReduceType(196)
	_ReduceNilToOptionalStatement                                       = _ReduceType(197)
	_ReduceExprToOptionalExpr                                           = _ReduceType(198)
	_ReduceNilToOptionalExpr                                            = _ReduceType(199)
	_ReduceToRepeatLoopBody                                             = _ReduceType(200)
	_ReduceToForLoopBody                                                = _ReduceType(201)
	_ReduceExprToReturnableExpr                                         = _ReduceType(202)
	_ReduceImproperExprStructToReturnableExpr                           = _ReduceType(203)
	_ReducePairToImproperExprStruct                                     = _ReduceType(204)
	_ReduceAddToImproperExprStruct                                      = _ReduceType(205)
	_ReduceExplicitStructTypeExprToInitializableTypeExpr                = _ReduceType(206)
	_ReduceSliceTypeExprToInitializableTypeExpr                         = _ReduceType(207)
	_ReduceArrayTypeExprToInitializableTypeExpr                         = _ReduceType(208)
	_ReduceMapTypeExprToInitializableTypeExpr                           = _ReduceType(209)
	_ReduceToSliceTypeExpr                                              = _ReduceType(210)
	_ReduceToArrayTypeExpr                                              = _ReduceType(211)
	_ReduceToMapTypeExpr                                                = _ReduceType(212)
	_ReduceInitializableTypeExprToAtomTypeExpr                          = _ReduceType(213)
	_ReduceNamedTypeExprToAtomTypeExpr                                  = _ReduceType(214)
	_ReduceInferredTypeExprToAtomTypeExpr                               = _ReduceType(215)
	_ReduceImplicitStructTypeExprToAtomTypeExpr                         = _ReduceType(216)
	_ReduceExplicitEnumTypeExprToAtomTypeExpr                           = _ReduceType(217)
	_ReduceImplicitEnumTypeExprToAtomTypeExpr                           = _ReduceType(218)
	_ReduceTraitTypeExprToAtomTypeExpr                                  = _ReduceType(219)
	_ReduceFuncSignatureToAtomTypeExpr                                  = _ReduceType(220)
	_ReduceLocalToNamedTypeExpr                                         = _ReduceType(221)
	_ReduceExternalToNamedTypeExpr                                      = _ReduceType(222)
	_ReduceDotToInferredTypeExpr                                        = _ReduceType(223)
	_ReduceUnderscoreToInferredTypeExpr                                 = _ReduceType(224)
	_ReduceAtomTypeExprToReturnableTypeExpr                             = _ReduceType(225)
	_ReducePrefixUnaryTypeExprToReturnableTypeExpr                      = _ReduceType(226)
	_ReduceToPrefixUnaryTypeExpr                                        = _ReduceType(227)
	_ReduceQuestionToPrefixUnaryTypeOp                                  = _ReduceType(228)
	_ReduceExclaimToPrefixUnaryTypeOp                                   = _ReduceType(229)
	_ReduceBitAndToPrefixUnaryTypeOp                                    = _ReduceType(230)
	_ReduceTildeToPrefixUnaryTypeOp                                     = _ReduceType(231)
	_ReduceTildeTildeToPrefixUnaryTypeOp                                = _ReduceType(232)
	_ReduceReturnableTypeExprToTypeExpr                                 = _ReduceType(233)
	_ReduceBinaryTypeExprToTypeExpr                                     = _ReduceType(234)
	_ReduceToBinaryTypeExpr                                             = _ReduceType(235)
	_ReduceMulToBinaryTypeOp                                            = _ReduceType(236)
	_ReduceAddToBinaryTypeOp                                            = _ReduceType(237)
	_ReduceSubToBinaryTypeOp                                            = _ReduceType(238)
	_ReduceDefinitionToTypeDef                                          = _ReduceType(239)
	_ReduceConstrainedDefToTypeDef                                      = _ReduceType(240)
	_ReduceAliasToTypeDef                                               = _ReduceType(241)
	_ReduceUnconstrainedToGenericParameter                              = _ReduceType(242)
	_ReduceConstrainedToGenericParameter                                = _ReduceType(243)
	_ReduceGenericToGenericParameters                                   = _ReduceType(244)
	_ReduceNilToGenericParameters                                       = _ReduceType(245)
	_ReduceAddToProperGenericParameterList                              = _ReduceType(246)
	_ReduceGenericParameterToProperGenericParameterList                 = _ReduceType(247)
	_ReduceProperGenericParameterListToGenericParameterList             = _ReduceType(248)
	_ReduceImproperToGenericParameterList                               = _ReduceType(249)
	_ReduceNilToGenericParameterList                                    = _ReduceType(250)
	_ReduceBindingToGenericArguments                                    = _ReduceType(251)
	_ReduceNilToGenericArguments                                        = _ReduceType(252)
	_ReduceAddToProperGenericArgumentList                               = _ReduceType(253)
	_ReduceTypeExprToProperGenericArgumentList                          = _ReduceType(254)
	_ReduceProperGenericArgumentListToGenericArgumentList               = _ReduceType(255)
	_ReduceImproperToGenericArgumentList                                = _ReduceType(256)
	_ReduceNilToGenericArgumentList                                     = _ReduceType(257)
	_ReduceUnnamedFieldToTypeProperty                                   = _ReduceType(258)
	_ReduceNamedFieldToTypeProperty                                     = _ReduceType(259)
	_ReducePaddingFieldToTypeProperty                                   = _ReduceType(260)
	_ReduceDefaultNamedEnumFieldToTypeProperty                          = _ReduceType(261)
	_ReduceDefaultUnnamedEnumFieldToTypeProperty                        = _ReduceType(262)
	_ReduceAddToProperImplicitTypeProperties                            = _ReduceType(263)
	_ReduceTypePropertyToProperImplicitTypeProperties                   = _ReduceType(264)
	_ReduceProperImplicitTypePropertiesToImplicitTypeProperties         = _ReduceType(265)
	_ReduceImproperToImplicitTypeProperties                             = _ReduceType(266)
	_ReduceNilToImplicitTypeProperties                                  = _ReduceType(267)
	_ReduceToImplicitStructTypeExpr                                     = _ReduceType(268)
	_ReduceAddImplicitToProperExplicitTypeProperties                    = _ReduceType(269)
	_ReduceAddExplicitToProperExplicitTypeProperties                    = _ReduceType(270)
	_ReduceTypePropertyToProperExplicitTypeProperties                   = _ReduceType(271)
	_ReduceProperExplicitTypePropertiesToExplicitTypeProperties         = _ReduceType(272)
	_ReduceImproperImplicitToExplicitTypeProperties                     = _ReduceType(273)
	_ReduceImproperExplicitToExplicitTypeProperties                     = _ReduceType(274)
	_ReduceNilToExplicitTypeProperties                                  = _ReduceType(275)
	_ReduceToExplicitStructTypeExpr                                     = _ReduceType(276)
	_ReduceToTraitTypeExpr                                              = _ReduceType(277)
	_ReducePairToProperImplicitEnumTypeProperties                       = _ReduceType(278)
	_ReduceAddToProperImplicitEnumTypeProperties                        = _ReduceType(279)
	_ReduceProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties = _ReduceType(280)
	_ReduceImproperToImplicitEnumTypeProperties                         = _ReduceType(281)
	_ReduceToImplicitEnumTypeExpr                                       = _ReduceType(282)
	_ReduceExplicitPairToProperExplicitEnumTypeProperties               = _ReduceType(283)
	_ReduceImplicitPairToProperExplicitEnumTypeProperties               = _ReduceType(284)
	_ReduceExplicitAddToProperExplicitEnumTypeProperties                = _ReduceType(285)
	_ReduceImplicitAddToProperExplicitEnumTypeProperties                = _ReduceType(286)
	_ReduceProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties = _ReduceType(287)
	_ReduceImproperToExplicitEnumTypeProperties                         = _ReduceType(288)
	_ReduceToExplicitEnumTypeExpr                                       = _ReduceType(289)
	_ReduceReturnableTypeExprToReturnType                               = _ReduceType(290)
	_ReduceNilToReturnType                                              = _ReduceType(291)
	_ReduceNamedArgToParameter                                          = _ReduceType(292)
	_ReduceNamedReceiverToParameter                                     = _ReduceType(293)
	_ReduceNamedVarargToParameter                                       = _ReduceType(294)
	_ReduceIgnoreArgToParameter                                         = _ReduceType(295)
	_ReduceIgnoreReceiverToParameter                                    = _ReduceType(296)
	_ReduceIgnoreVarargToParameter                                      = _ReduceType(297)
	_ReduceUnnamedArgToParameter                                        = _ReduceType(298)
	_ReduceUnnamedReceiverToParameter                                   = _ReduceType(299)
	_ReduceUnnamedVarargToParameter                                     = _ReduceType(300)
	_ReduceAddToProperParameterList                                     = _ReduceType(301)
	_ReduceParameterToProperParameterList                               = _ReduceType(302)
	_ReduceProperParameterListToParameterList                           = _ReduceType(303)
	_ReduceImproperToParameterList                                      = _ReduceType(304)
	_ReduceNilToParameterList                                           = _ReduceType(305)
	_ReduceToParameters                                                 = _ReduceType(306)
	_ReduceAnonymousToFuncSignature                                     = _ReduceType(307)
	_ReduceNamedToFuncSignature                                         = _ReduceType(308)
	_ReduceToFuncDef                                                    = _ReduceType(309)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceStatementListToSource:
		return "StatementListToSource"
	case _ReduceUnsafeStmtToStatement:
		return "UnsafeStmtToStatement"
	case _ReduceImportStmtToStatement:
		return "ImportStmtToStatement"
	case _ReduceFloatingCommentToStatement:
		return "FloatingCommentToStatement"
	case _ReduceTypeDefToStatement:
		return "TypeDefToStatement"
	case _ReduceBranchStmtToStatement:
		return "BranchStmtToStatement"
	case _ReduceReturnableExprToStatement:
		return "ReturnableExprToStatement"
	case _ReduceJumpStmtToStatement:
		return "JumpStmtToStatement"
	case _ReduceAssignStmtToStatement:
		return "AssignStmtToStatement"
	case _ReduceToFloatingComment:
		return "ToFloatingComment"
	case _ReduceCaseBranchToBranchStmt:
		return "CaseBranchToBranchStmt"
	case _ReduceDefaultBranchToBranchStmt:
		return "DefaultBranchToBranchStmt"
	case _ReduceToUnsafeStmt:
		return "ToUnsafeStmt"
	case _ReduceUnlabeledNoValueToJumpStmt:
		return "UnlabeledNoValueToJumpStmt"
	case _ReduceUnlabeledValuedToJumpStmt:
		return "UnlabeledValuedToJumpStmt"
	case _ReduceLabeledNoValueToJumpStmt:
		return "LabeledNoValueToJumpStmt"
	case _ReduceLabeledValuedToJumpStmt:
		return "LabeledValuedToJumpStmt"
	case _ReduceFallthroughToJumpStmt:
		return "FallthroughToJumpStmt"
	case _ReduceReturnToJumpOp:
		return "ReturnToJumpOp"
	case _ReduceBreakToJumpOp:
		return "BreakToJumpOp"
	case _ReduceContinueToJumpOp:
		return "ContinueToJumpOp"
	case _ReduceToAssignStmt:
		return "ToAssignStmt"
	case _ReduceSingleToImportStmt:
		return "SingleToImportStmt"
	case _ReduceMultipleToImportStmt:
		return "MultipleToImportStmt"
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
	case _ReduceStringLiteralToImportClause:
		return "StringLiteralToImportClause"
	case _ReduceAliasToImportClause:
		return "AliasToImportClause"
	case _ReduceUnusableImportToImportClause:
		return "UnusableImportToImportClause"
	case _ReduceImportToLocalToImportClause:
		return "ImportToLocalToImportClause"
	case _ReduceNewInferredToAddrDeclPattern:
		return "NewInferredToAddrDeclPattern"
	case _ReduceNewTypedToAddrDeclPattern:
		return "NewTypedToAddrDeclPattern"
	case _ReduceToAssignToAddrPattern:
		return "ToAssignToAddrPattern"
	case _ReduceVarToNewVarType:
		return "VarToNewVarType"
	case _ReduceLetToNewVarType:
		return "LetToNewVarType"
	case _ReduceNamedExprToNewAddressable:
		return "NamedExprToNewAddressable"
	case _ReduceImplicitStructExprToNewAddressable:
		return "ImplicitStructExprToNewAddressable"
	case _ReduceSwitchableCasePatternsToCasePatterns:
		return "SwitchableCasePatternsToCasePatterns"
	case _ReduceToCasePatterns:
		return "ToCasePatterns"
	case _ReduceToAssignSelectablePattern:
		return "ToAssignSelectablePattern"
	case _ReduceSwitchableCasePatternToSwitchableCasePatterns:
		return "SwitchableCasePatternToSwitchableCasePatterns"
	case _ReduceAddToSwitchableCasePatterns:
		return "AddToSwitchableCasePatterns"
	case _ReduceExprToSwitchableCasePattern:
		return "ExprToSwitchableCasePattern"
	case _ReduceEnumPatternToSwitchableCasePattern:
		return "EnumPatternToSwitchableCasePattern"
	case _ReduceNamedToEnumPattern:
		return "NamedToEnumPattern"
	case _ReduceUnnamedStructToEnumPattern:
		return "UnnamedStructToEnumPattern"
	case _ReduceParseErrorExprToAtomExpr:
		return "ParseErrorExprToAtomExpr"
	case _ReduceLiteralExprToAtomExpr:
		return "LiteralExprToAtomExpr"
	case _ReduceNamedExprToAtomExpr:
		return "NamedExprToAtomExpr"
	case _ReduceFuncDefToAtomExpr:
		return "FuncDefToAtomExpr"
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
	case _ReduceAsExprToAccessibleExpr:
		return "AsExprToAccessibleExpr"
	case _ReduceToAccessExpr:
		return "ToAccessExpr"
	case _ReduceToIndexExpr:
		return "ToIndexExpr"
	case _ReduceExprToIndex:
		return "ExprToIndex"
	case _ReduceColonExprToIndex:
		return "ColonExprToIndex"
	case _ReduceToAsExpr:
		return "ToAsExpr"
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
	case _ReduceAccessibleExprToPostfixableExpr:
		return "AccessibleExprToPostfixableExpr"
	case _ReducePostfixUnaryExprToPostfixableExpr:
		return "PostfixUnaryExprToPostfixableExpr"
	case _ReduceToPostfixUnaryExpr:
		return "ToPostfixUnaryExpr"
	case _ReduceQuestionToPostfixUnaryOp:
		return "QuestionToPostfixUnaryOp"
	case _ReduceExclaimToPostfixUnaryOp:
		return "ExclaimToPostfixUnaryOp"
	case _ReduceAddOneAssignToPostfixUnaryOp:
		return "AddOneAssignToPostfixUnaryOp"
	case _ReduceSubOneAssignToPostfixUnaryOp:
		return "SubOneAssignToPostfixUnaryOp"
	case _ReducePostfixableExprToPrefixableExpr:
		return "PostfixableExprToPrefixableExpr"
	case _ReducePrefixUnaryExprToPrefixableExpr:
		return "PrefixUnaryExprToPrefixableExpr"
	case _ReduceToPrefixUnaryExpr:
		return "ToPrefixUnaryExpr"
	case _ReduceNotToPrefixUnaryOp:
		return "NotToPrefixUnaryOp"
	case _ReduceBitXorToPrefixUnaryOp:
		return "BitXorToPrefixUnaryOp"
	case _ReduceAddToPrefixUnaryOp:
		return "AddToPrefixUnaryOp"
	case _ReduceSubToPrefixUnaryOp:
		return "SubToPrefixUnaryOp"
	case _ReduceMulToPrefixUnaryOp:
		return "MulToPrefixUnaryOp"
	case _ReduceBitAndToPrefixUnaryOp:
		return "BitAndToPrefixUnaryOp"
	case _ReduceAsyncToPrefixUnaryOp:
		return "AsyncToPrefixUnaryOp"
	case _ReduceDeferToPrefixUnaryOp:
		return "DeferToPrefixUnaryOp"
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
	case _ReduceOrExprToSendRecvExpr:
		return "OrExprToSendRecvExpr"
	case _ReduceSendExprToSendRecvExpr:
		return "SendExprToSendRecvExpr"
	case _ReduceRecvExprToSendRecvExpr:
		return "RecvExprToSendRecvExpr"
	case _ReduceToSendExpr:
		return "ToSendExpr"
	case _ReduceToRecvExpr:
		return "ToRecvExpr"
	case _ReduceSendRecvExprToAssignOpExpr:
		return "SendRecvExprToAssignOpExpr"
	case _ReduceBinaryAssignOpExprToAssignOpExpr:
		return "BinaryAssignOpExprToAssignOpExpr"
	case _ReduceToBinaryAssignOpExpr:
		return "ToBinaryAssignOpExpr"
	case _ReduceAddAssignToBinaryAssignOp:
		return "AddAssignToBinaryAssignOp"
	case _ReduceSubAssignToBinaryAssignOp:
		return "SubAssignToBinaryAssignOp"
	case _ReduceMulAssignToBinaryAssignOp:
		return "MulAssignToBinaryAssignOp"
	case _ReduceDivAssignToBinaryAssignOp:
		return "DivAssignToBinaryAssignOp"
	case _ReduceModAssignToBinaryAssignOp:
		return "ModAssignToBinaryAssignOp"
	case _ReduceBitAndAssignToBinaryAssignOp:
		return "BitAndAssignToBinaryAssignOp"
	case _ReduceBitOrAssignToBinaryAssignOp:
		return "BitOrAssignToBinaryAssignOp"
	case _ReduceBitXorAssignToBinaryAssignOp:
		return "BitXorAssignToBinaryAssignOp"
	case _ReduceBitLshiftAssignToBinaryAssignOp:
		return "BitLshiftAssignToBinaryAssignOp"
	case _ReduceBitRshiftAssignToBinaryAssignOp:
		return "BitRshiftAssignToBinaryAssignOp"
	case _ReduceStatementsToUnlabelledControlFlowExpr:
		return "StatementsToUnlabelledControlFlowExpr"
	case _ReduceIfElseExprToUnlabelledControlFlowExpr:
		return "IfElseExprToUnlabelledControlFlowExpr"
	case _ReduceSwitchExprBodyToUnlabelledControlFlowExpr:
		return "SwitchExprBodyToUnlabelledControlFlowExpr"
	case _ReduceSelectExprBodyToUnlabelledControlFlowExpr:
		return "SelectExprBodyToUnlabelledControlFlowExpr"
	case _ReduceLoopExprBodyToUnlabelledControlFlowExpr:
		return "LoopExprBodyToUnlabelledControlFlowExpr"
	case _ReduceUnlabelledControlFlowExprToControlFlowExpr:
		return "UnlabelledControlFlowExprToControlFlowExpr"
	case _ReduceLabelledToControlFlowExpr:
		return "LabelledToControlFlowExpr"
	case _ReduceAssignOpExprToExpr:
		return "AssignOpExprToExpr"
	case _ReduceControlFlowExprToExpr:
		return "ControlFlowExprToExpr"
	case _ReduceAddrDeclPatternToExpr:
		return "AddrDeclPatternToExpr"
	case _ReduceAssignToAddrPatternToExpr:
		return "AssignToAddrPatternToExpr"
	case _ReduceToStatements:
		return "ToStatements"
	case _ReduceAddImplicitToProperStatementList:
		return "AddImplicitToProperStatementList"
	case _ReduceAddExplicitToProperStatementList:
		return "AddExplicitToProperStatementList"
	case _ReduceStatementToProperStatementList:
		return "StatementToProperStatementList"
	case _ReduceProperStatementListToStatementList:
		return "ProperStatementListToStatementList"
	case _ReduceImproperImplicitToStatementList:
		return "ImproperImplicitToStatementList"
	case _ReduceImproperExplicitToStatementList:
		return "ImproperExplicitToStatementList"
	case _ReduceNilToStatementList:
		return "NilToStatementList"
	case _ReduceIfElifExprToIfElseExpr:
		return "IfElifExprToIfElseExpr"
	case _ReduceElseToIfElseExpr:
		return "ElseToIfElseExpr"
	case _ReduceIfOnlyExprToIfElifExpr:
		return "IfOnlyExprToIfElifExpr"
	case _ReduceElifToIfElifExpr:
		return "ElifToIfElifExpr"
	case _ReduceToIfOnlyExpr:
		return "ToIfOnlyExpr"
	case _ReduceExprToCondition:
		return "ExprToCondition"
	case _ReduceCasePatternExprToCondition:
		return "CasePatternExprToCondition"
	case _ReduceToCasePatternExpr:
		return "ToCasePatternExpr"
	case _ReduceToSwitchExprBody:
		return "ToSwitchExprBody"
	case _ReduceToSelectExprBody:
		return "ToSelectExprBody"
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
	case _ReduceStatementToOptionalStatement:
		return "StatementToOptionalStatement"
	case _ReduceNilToOptionalStatement:
		return "NilToOptionalStatement"
	case _ReduceExprToOptionalExpr:
		return "ExprToOptionalExpr"
	case _ReduceNilToOptionalExpr:
		return "NilToOptionalExpr"
	case _ReduceToRepeatLoopBody:
		return "ToRepeatLoopBody"
	case _ReduceToForLoopBody:
		return "ToForLoopBody"
	case _ReduceExprToReturnableExpr:
		return "ExprToReturnableExpr"
	case _ReduceImproperExprStructToReturnableExpr:
		return "ImproperExprStructToReturnableExpr"
	case _ReducePairToImproperExprStruct:
		return "PairToImproperExprStruct"
	case _ReduceAddToImproperExprStruct:
		return "AddToImproperExprStruct"
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
	case _ReduceFuncSignatureToAtomTypeExpr:
		return "FuncSignatureToAtomTypeExpr"
	case _ReduceLocalToNamedTypeExpr:
		return "LocalToNamedTypeExpr"
	case _ReduceExternalToNamedTypeExpr:
		return "ExternalToNamedTypeExpr"
	case _ReduceDotToInferredTypeExpr:
		return "DotToInferredTypeExpr"
	case _ReduceUnderscoreToInferredTypeExpr:
		return "UnderscoreToInferredTypeExpr"
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
	case _ReduceTildeToPrefixUnaryTypeOp:
		return "TildeToPrefixUnaryTypeOp"
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
	case _ReduceUnnamedFieldToTypeProperty:
		return "UnnamedFieldToTypeProperty"
	case _ReduceNamedFieldToTypeProperty:
		return "NamedFieldToTypeProperty"
	case _ReducePaddingFieldToTypeProperty:
		return "PaddingFieldToTypeProperty"
	case _ReduceDefaultNamedEnumFieldToTypeProperty:
		return "DefaultNamedEnumFieldToTypeProperty"
	case _ReduceDefaultUnnamedEnumFieldToTypeProperty:
		return "DefaultUnnamedEnumFieldToTypeProperty"
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
	case _ReduceNamedArgToParameter:
		return "NamedArgToParameter"
	case _ReduceNamedReceiverToParameter:
		return "NamedReceiverToParameter"
	case _ReduceNamedVarargToParameter:
		return "NamedVarargToParameter"
	case _ReduceIgnoreArgToParameter:
		return "IgnoreArgToParameter"
	case _ReduceIgnoreReceiverToParameter:
		return "IgnoreReceiverToParameter"
	case _ReduceIgnoreVarargToParameter:
		return "IgnoreVarargToParameter"
	case _ReduceUnnamedArgToParameter:
		return "UnnamedArgToParameter"
	case _ReduceUnnamedReceiverToParameter:
		return "UnnamedReceiverToParameter"
	case _ReduceUnnamedVarargToParameter:
		return "UnnamedVarargToParameter"
	case _ReduceAddToProperParameterList:
		return "AddToProperParameterList"
	case _ReduceParameterToProperParameterList:
		return "ParameterToProperParameterList"
	case _ReduceProperParameterListToParameterList:
		return "ProperParameterListToParameterList"
	case _ReduceImproperToParameterList:
		return "ImproperToParameterList"
	case _ReduceNilToParameterList:
		return "NilToParameterList"
	case _ReduceToParameters:
		return "ToParameters"
	case _ReduceAnonymousToFuncSignature:
		return "AnonymousToFuncSignature"
	case _ReduceNamedToFuncSignature:
		return "NamedToFuncSignature"
	case _ReduceToFuncDef:
		return "ToFuncDef"
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
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Argument             *ast.Argument
	ArgumentList         *ast.ArgumentList
	CommentGroups        CommentGroupsTok
	ControlFlowExpr      ast.ControlFlowExpr
	Count                TokenCount
	Expression           ast.Expression
	ExpressionList       *ast.ExpressionList
	FuncSignature        *ast.FuncSignature
	GenericParameter     *ast.GenericParameter
	GenericParameterList *ast.GenericParameterList
	IfExpr               *ast.IfExpr
	ImplicitStructExpr   *ast.ImplicitStructExpr
	ImportClause         *ast.ImportClause
	ImportClauseList     *ast.ImportClauseList
	Parameter            *ast.Parameter
	Parameters           *ast.ParameterList
	ParseError           *ParseErrorSymbol
	Statement            ast.Statement
	StatementList        *ast.StatementList
	StatementsExpr       *ast.StatementsExpr
	TypeExpression       ast.TypeExpression
	TypeExpressionList   *ast.TypeExpressionList
	TypeProperties       *ast.TypePropertyList
	TypeProperty         ast.TypeProperty
	UnsafeStmt           *ast.UnsafeStmt
	Value                *TokenValue
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case CommentGroupsToken:
		val, ok := token.(CommentGroupsTok)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting CommentGroupsTok",
				token.Id())
		}
		symbol.CommentGroups = val
	case NewlinesToken:
		val, ok := token.(TokenCount)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting TokenCount",
				token.Id())
		}
		symbol.Count = val
	case _EndMarker:
		val, ok := token.(GenericSymbol)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting GenericSymbol",
				token.Id())
		}
		symbol.Generic_ = val
	case ParseErrorToken:
		val, ok := token.(*ParseErrorSymbol)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *ParseErrorSymbol",
				token.Id())
		}
		symbol.ParseError = val
	case IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, RepeatToken, ForToken, DoToken, InToken, SelectToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, AsToken, NotToken, AndToken, OrToken, AtToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeToken, TildeTildeToken, AssignToken, ArrowToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken:
		val, ok := token.(*TokenValue)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *TokenValue",
				token.Id())
		}
		symbol.Value = val
	default:
		return nil, lexutil.NewLocationError(
			token.Loc(),
			"unexpected token type: %s",
			token.Id())
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
	case CommentGroupsToken:
		loc, ok := interface{}(s.CommentGroups).(locator)
		if ok {
			return loc.Loc()
		}
	case UnlabelledControlFlowExprType, ControlFlowExprType, SwitchExprBodyType, SelectExprBodyType, LoopExprBodyType:
		loc, ok := interface{}(s.ControlFlowExpr).(locator)
		if ok {
			return loc.Loc()
		}
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.Loc()
		}
	case AddrDeclPatternType, AssignToAddrPatternType, NewAddressableType, AssignSelectablePatternType, SwitchableCasePatternType, EnumPatternType, AtomExprType, ParseErrorExprType, LiteralExprType, NamedExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, IndexType, AsExprType, CallExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, SendRecvExprType, SendExprType, RecvExprType, AssignOpExprType, BinaryAssignOpExprType, ExprType, ConditionType, CasePatternExprType, OptionalExprType, ReturnableExprType, FuncDefType:
		loc, ok := interface{}(s.Expression).(locator)
		if ok {
			return loc.Loc()
		}
	case CasePatternsType, SwitchableCasePatternsType:
		loc, ok := interface{}(s.ExpressionList).(locator)
		if ok {
			return loc.Loc()
		}
	case FuncSignatureType:
		loc, ok := interface{}(s.FuncSignature).(locator)
		if ok {
			return loc.Loc()
		}
	case GenericParameterType:
		loc, ok := interface{}(s.GenericParameter).(locator)
		if ok {
			return loc.Loc()
		}
	case GenericParametersType, ProperGenericParameterListType, GenericParameterListType:
		loc, ok := interface{}(s.GenericParameterList).(locator)
		if ok {
			return loc.Loc()
		}
	case IfElseExprType, IfElifExprType, IfOnlyExprType:
		loc, ok := interface{}(s.IfExpr).(locator)
		if ok {
			return loc.Loc()
		}
	case ColonExprType, ImproperExprStructType:
		loc, ok := interface{}(s.ImplicitStructExpr).(locator)
		if ok {
			return loc.Loc()
		}
	case ImportClauseType:
		loc, ok := interface{}(s.ImportClause).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperImportClausesType, ImportClausesType:
		loc, ok := interface{}(s.ImportClauseList).(locator)
		if ok {
			return loc.Loc()
		}
	case ParameterType:
		loc, ok := interface{}(s.Parameter).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperParameterListType, ParameterListType, ParametersType:
		loc, ok := interface{}(s.Parameters).(locator)
		if ok {
			return loc.Loc()
		}
	case ParseErrorToken:
		loc, ok := interface{}(s.ParseError).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementType, FloatingCommentType, BranchStmtType, JumpStmtType, AssignStmtType, ImportStmtType, OptionalStatementType, TypeDefType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case SourceType, ProperStatementListType, StatementListType:
		loc, ok := interface{}(s.StatementList).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementsType, RepeatLoopBodyType, ForLoopBodyType:
		loc, ok := interface{}(s.StatementsExpr).(locator)
		if ok {
			return loc.Loc()
		}
	case InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ReturnableTypeExprType, PrefixUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case GenericArgumentsType, ProperGenericArgumentListType, GenericArgumentListType:
		loc, ok := interface{}(s.TypeExpressionList).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperImplicitTypePropertiesType, ImplicitTypePropertiesType, ProperExplicitTypePropertiesType, ExplicitTypePropertiesType, ProperImplicitEnumTypePropertiesType, ImplicitEnumTypePropertiesType, ProperExplicitEnumTypePropertiesType, ExplicitEnumTypePropertiesType:
		loc, ok := interface{}(s.TypeProperties).(locator)
		if ok {
			return loc.Loc()
		}
	case TypePropertyType:
		loc, ok := interface{}(s.TypeProperty).(locator)
		if ok {
			return loc.Loc()
		}
	case UnsafeStmtType:
		loc, ok := interface{}(s.UnsafeStmt).(locator)
		if ok {
			return loc.Loc()
		}
	case IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, RepeatToken, ForToken, DoToken, InToken, SelectToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, AsToken, NotToken, AndToken, OrToken, AtToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeToken, TildeTildeToken, AssignToken, ArrowToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, JumpOpType, NewVarTypeType, PostfixUnaryOpType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, BinaryAssignOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
	case CommentGroupsToken:
		loc, ok := interface{}(s.CommentGroups).(locator)
		if ok {
			return loc.End()
		}
	case UnlabelledControlFlowExprType, ControlFlowExprType, SwitchExprBodyType, SelectExprBodyType, LoopExprBodyType:
		loc, ok := interface{}(s.ControlFlowExpr).(locator)
		if ok {
			return loc.End()
		}
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.End()
		}
	case AddrDeclPatternType, AssignToAddrPatternType, NewAddressableType, AssignSelectablePatternType, SwitchableCasePatternType, EnumPatternType, AtomExprType, ParseErrorExprType, LiteralExprType, NamedExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, IndexType, AsExprType, CallExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, SendRecvExprType, SendExprType, RecvExprType, AssignOpExprType, BinaryAssignOpExprType, ExprType, ConditionType, CasePatternExprType, OptionalExprType, ReturnableExprType, FuncDefType:
		loc, ok := interface{}(s.Expression).(locator)
		if ok {
			return loc.End()
		}
	case CasePatternsType, SwitchableCasePatternsType:
		loc, ok := interface{}(s.ExpressionList).(locator)
		if ok {
			return loc.End()
		}
	case FuncSignatureType:
		loc, ok := interface{}(s.FuncSignature).(locator)
		if ok {
			return loc.End()
		}
	case GenericParameterType:
		loc, ok := interface{}(s.GenericParameter).(locator)
		if ok {
			return loc.End()
		}
	case GenericParametersType, ProperGenericParameterListType, GenericParameterListType:
		loc, ok := interface{}(s.GenericParameterList).(locator)
		if ok {
			return loc.End()
		}
	case IfElseExprType, IfElifExprType, IfOnlyExprType:
		loc, ok := interface{}(s.IfExpr).(locator)
		if ok {
			return loc.End()
		}
	case ColonExprType, ImproperExprStructType:
		loc, ok := interface{}(s.ImplicitStructExpr).(locator)
		if ok {
			return loc.End()
		}
	case ImportClauseType:
		loc, ok := interface{}(s.ImportClause).(locator)
		if ok {
			return loc.End()
		}
	case ProperImportClausesType, ImportClausesType:
		loc, ok := interface{}(s.ImportClauseList).(locator)
		if ok {
			return loc.End()
		}
	case ParameterType:
		loc, ok := interface{}(s.Parameter).(locator)
		if ok {
			return loc.End()
		}
	case ProperParameterListType, ParameterListType, ParametersType:
		loc, ok := interface{}(s.Parameters).(locator)
		if ok {
			return loc.End()
		}
	case ParseErrorToken:
		loc, ok := interface{}(s.ParseError).(locator)
		if ok {
			return loc.End()
		}
	case StatementType, FloatingCommentType, BranchStmtType, JumpStmtType, AssignStmtType, ImportStmtType, OptionalStatementType, TypeDefType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.End()
		}
	case SourceType, ProperStatementListType, StatementListType:
		loc, ok := interface{}(s.StatementList).(locator)
		if ok {
			return loc.End()
		}
	case StatementsType, RepeatLoopBodyType, ForLoopBodyType:
		loc, ok := interface{}(s.StatementsExpr).(locator)
		if ok {
			return loc.End()
		}
	case InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ReturnableTypeExprType, PrefixUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case GenericArgumentsType, ProperGenericArgumentListType, GenericArgumentListType:
		loc, ok := interface{}(s.TypeExpressionList).(locator)
		if ok {
			return loc.End()
		}
	case ProperImplicitTypePropertiesType, ImplicitTypePropertiesType, ProperExplicitTypePropertiesType, ExplicitTypePropertiesType, ProperImplicitEnumTypePropertiesType, ImplicitEnumTypePropertiesType, ProperExplicitEnumTypePropertiesType, ExplicitEnumTypePropertiesType:
		loc, ok := interface{}(s.TypeProperties).(locator)
		if ok {
			return loc.End()
		}
	case TypePropertyType:
		loc, ok := interface{}(s.TypeProperty).(locator)
		if ok {
			return loc.End()
		}
	case UnsafeStmtType:
		loc, ok := interface{}(s.UnsafeStmt).(locator)
		if ok {
			return loc.End()
		}
	case IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, RepeatToken, ForToken, DoToken, InToken, SelectToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, AsToken, NotToken, AndToken, OrToken, AtToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeToken, TildeTildeToken, AssignToken, ArrowToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, JumpOpType, NewVarTypeType, PostfixUnaryOpType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, BinaryAssignOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
				return nil, lexutil.NewLocationError(
					stack.lexer.CurrentLocation(),
					"unexpected lex error: %w",
					err)
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
	stack _Stack,
) (
	_Stack,
	*Symbol,
	error,
) {
	var err error
	symbol := &Symbol{}
	switch act.ReduceType {
	case _ReduceStatementListToSource:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SourceType
		//line grammar.lr:65:4
		symbol.StatementList = args[0].StatementList
		err = nil
	case _ReduceUnsafeStmtToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:72:4
		symbol.Statement = args[0].UnsafeStmt
		err = nil
	case _ReduceImportStmtToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:76:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceFloatingCommentToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:77:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceTypeDefToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:78:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceBranchStmtToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:81:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceReturnableExprToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:85:4
		symbol.Statement = args[0].Expression
		err = nil
	case _ReduceJumpStmtToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:86:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceAssignStmtToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:90:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceToFloatingComment:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FloatingCommentType
		symbol.Statement, err = reducer.ToFloatingComment(args[0].CommentGroups)
	case _ReduceCaseBranchToBranchStmt:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = BranchStmtType
		symbol.Statement, err = reducer.CaseBranchToBranchStmt(args[0].Value, args[1].ExpressionList, args[2].Value, args[3].Statement)
	case _ReduceDefaultBranchToBranchStmt:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BranchStmtType
		symbol.Statement, err = reducer.DefaultBranchToBranchStmt(args[0].Value, args[1].Value, args[2].Statement)
	case _ReduceToUnsafeStmt:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = UnsafeStmtType
		symbol.UnsafeStmt, err = reducer.ToUnsafeStmt(args[0].Value, args[1].Value, args[2].Value, args[3].Value, args[4].Value)
	case _ReduceUnlabeledNoValueToJumpStmt:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpStmtType
		symbol.Statement, err = reducer.UnlabeledNoValueToJumpStmt(args[0].Value)
	case _ReduceUnlabeledValuedToJumpStmt:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = JumpStmtType
		symbol.Statement, err = reducer.UnlabeledValuedToJumpStmt(args[0].Value, args[1].Expression)
	case _ReduceLabeledNoValueToJumpStmt:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = JumpStmtType
		symbol.Statement, err = reducer.LabeledNoValueToJumpStmt(args[0].Value, args[1].Value, args[2].Value)
	case _ReduceLabeledValuedToJumpStmt:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = JumpStmtType
		symbol.Statement, err = reducer.LabeledValuedToJumpStmt(args[0].Value, args[1].Value, args[2].Value, args[3].Expression)
	case _ReduceFallthroughToJumpStmt:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpStmtType
		symbol.Statement, err = reducer.FallthroughToJumpStmt(args[0].Value)
	case _ReduceReturnToJumpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpOpType
		//line grammar.lr:133:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBreakToJumpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpOpType
		//line grammar.lr:134:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceContinueToJumpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpOpType
		//line grammar.lr:135:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceToAssignStmt:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AssignStmtType
		symbol.Statement, err = reducer.ToAssignStmt(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceSingleToImportStmt:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportStmtType
		symbol.Statement, err = reducer.SingleToImportStmt(args[0].Value, args[1].ImportClause)
	case _ReduceMultipleToImportStmt:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ImportStmtType
		symbol.Statement, err = reducer.MultipleToImportStmt(args[0].Value, args[1].Value, args[2].ImportClauseList, args[3].Value)
	case _ReduceAddImplicitToProperImportClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImportClausesType
		symbol.ImportClauseList, err = reducer.AddImplicitToProperImportClauses(args[0].ImportClauseList, args[1].Count, args[2].ImportClause)
	case _ReduceAddExplicitToProperImportClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImportClausesType
		symbol.ImportClauseList, err = reducer.AddExplicitToProperImportClauses(args[0].ImportClauseList, args[1].Value, args[2].ImportClause)
	case _ReduceImportClauseToProperImportClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperImportClausesType
		symbol.ImportClauseList, err = reducer.ImportClauseToProperImportClauses(args[0].ImportClause)
	case _ReduceProperImportClausesToImportClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImportClausesType
		//line grammar.lr:161:4
		symbol.ImportClauseList = args[0].ImportClauseList
		err = nil
	case _ReduceImplicitToImportClauses:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClausesType
		symbol.ImportClauseList, err = reducer.ImplicitToImportClauses(args[0].ImportClauseList, args[1].Count)
	case _ReduceExplicitToImportClauses:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClausesType
		symbol.ImportClauseList, err = reducer.ExplicitToImportClauses(args[0].ImportClauseList, args[1].Value)
	case _ReduceStringLiteralToImportClause:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImportClauseType
		symbol.ImportClause, err = reducer.StringLiteralToImportClause(args[0].Value)
	case _ReduceAliasToImportClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseType
		symbol.ImportClause, err = reducer.AliasToImportClause(args[0].Value, args[1].Value)
	case _ReduceUnusableImportToImportClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseType
		symbol.ImportClause, err = reducer.UnusableImportToImportClause(args[0].Value, args[1].Value)
	case _ReduceImportToLocalToImportClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseType
		symbol.ImportClause, err = reducer.ImportToLocalToImportClause(args[0].Value, args[1].Value)
	case _ReduceNewInferredToAddrDeclPattern:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AddrDeclPatternType
		symbol.Expression, err = reducer.NewInferredToAddrDeclPattern(args[0].Value, args[1].Expression)
	case _ReduceNewTypedToAddrDeclPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AddrDeclPatternType
		symbol.Expression, err = reducer.NewTypedToAddrDeclPattern(args[0].Value, args[1].Expression, args[2].TypeExpression)
	case _ReduceToAssignToAddrPattern:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AssignToAddrPatternType
		symbol.Expression, err = reducer.ToAssignToAddrPattern(args[0].Value, args[1].Expression)
	case _ReduceVarToNewVarType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NewVarTypeType
		//line grammar.lr:185:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLetToNewVarType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NewVarTypeType
		//line grammar.lr:186:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNamedExprToNewAddressable:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NewAddressableType
		//line grammar.lr:189:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToNewAddressable:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NewAddressableType
		//line grammar.lr:190:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceSwitchableCasePatternsToCasePatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CasePatternsType
		//line grammar.lr:194:4
		symbol.ExpressionList = args[0].ExpressionList
		err = nil
	case _ReduceToCasePatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CasePatternsType
		symbol.ExpressionList, err = reducer.ToCasePatterns(args[0].Expression)
	case _ReduceToAssignSelectablePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AssignSelectablePatternType
		symbol.Expression, err = reducer.ToAssignSelectablePattern(args[0].ExpressionList, args[1].Value, args[2].Expression)
	case _ReduceSwitchableCasePatternToSwitchableCasePatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SwitchableCasePatternsType
		symbol.ExpressionList, err = reducer.SwitchableCasePatternToSwitchableCasePatterns(args[0].Expression)
	case _ReduceAddToSwitchableCasePatterns:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SwitchableCasePatternsType
		symbol.ExpressionList, err = reducer.AddToSwitchableCasePatterns(args[0].ExpressionList, args[1].Value, args[2].Expression)
	case _ReduceExprToSwitchableCasePattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SwitchableCasePatternType
		//line grammar.lr:228:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceEnumPatternToSwitchableCasePattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SwitchableCasePatternType
		//line grammar.lr:229:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceNamedToEnumPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = EnumPatternType
		symbol.Expression, err = reducer.NamedToEnumPattern(args[0].Value, args[1].Value, args[2].Expression)
	case _ReduceUnnamedStructToEnumPattern:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = EnumPatternType
		symbol.Expression, err = reducer.UnnamedStructToEnumPattern(args[0].Value, args[1].Expression)
	case _ReduceParseErrorExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:249:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:250:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceNamedExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:251:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceFuncDefToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:252:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:253:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:254:4
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
		//line grammar.lr:279:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:280:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCallExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:281:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:282:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAsExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:283:4
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
		symbol.Expression, err = reducer.ToIndexExpr(args[0].Expression, args[1].Value, args[2].Expression, args[3].Value)
	case _ReduceExprToIndex:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = IndexType
		//line grammar.lr:292:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceColonExprToIndex:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = IndexType
		//line grammar.lr:293:4
		symbol.Expression = args[0].ImplicitStructExpr
		err = nil
	case _ReduceToAsExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = AsExprType
		symbol.Expression, err = reducer.ToAsExpr(args[0].Expression, args[1].Value, args[2].Value, args[3].Value, args[4].TypeExpression, args[5].Value)
	case _ReduceToCallExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CallExprType
		symbol.Expression, err = reducer.ToCallExpr(args[0].Expression, args[1].TypeExpressionList, args[2].Value, args[3].ArgumentList, args[4].Value)
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
		//line grammar.lr:306:4
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
		symbol.Argument, err = reducer.ColonExprToArgument(args[0].ImplicitStructExpr)
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
		symbol.ImplicitStructExpr, err = reducer.UnitUnitPairToColonExpr(args[0].Value)
	case _ReduceExprUnitPairToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.ImplicitStructExpr, err = reducer.ExprUnitPairToColonExpr(args[0].Expression, args[1].Value)
	case _ReduceUnitExprPairToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.ImplicitStructExpr, err = reducer.UnitExprPairToColonExpr(args[0].Value, args[1].Expression)
	case _ReduceExprExprPairToColonExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExprType
		symbol.ImplicitStructExpr, err = reducer.ExprExprPairToColonExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceColonExprUnitTupleToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.ImplicitStructExpr, err = reducer.ColonExprUnitTupleToColonExpr(args[0].ImplicitStructExpr, args[1].Value)
	case _ReduceColonExprExprTupleToColonExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExprType
		symbol.ImplicitStructExpr, err = reducer.ColonExprExprTupleToColonExpr(args[0].ImplicitStructExpr, args[1].Value, args[2].Expression)
	case _ReduceAccessibleExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:336:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePostfixUnaryExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:337:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToPostfixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Expression, err = reducer.ToPostfixUnaryExpr(args[0].Expression, args[1].Value)
	case _ReduceQuestionToPostfixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryOpType
		//line grammar.lr:342:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPostfixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryOpType
		//line grammar.lr:343:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddOneAssignToPostfixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryOpType
		//line grammar.lr:347:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubOneAssignToPostfixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryOpType
		//line grammar.lr:348:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePostfixableExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:355:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePrefixUnaryExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:356:4
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
		//line grammar.lr:361:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:362:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:363:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:364:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:367:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:370:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAsyncToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:388:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDeferToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:389:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixableExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:396:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryMulExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:397:4
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
		//line grammar.lr:402:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:403:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:404:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:405:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:406:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:407:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:414:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAddExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:415:4
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
		//line grammar.lr:420:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:421:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:422:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:423:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:430:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryCmpExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:431:4
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
		//line grammar.lr:436:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:437:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:438:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:439:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:440:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:441:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:448:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAndExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:449:4
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
		//line grammar.lr:458:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryOrExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:459:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryOrExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryOrExprType
		symbol.Expression, err = reducer.ToBinaryOrExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceOrExprToSendRecvExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SendRecvExprType
		//line grammar.lr:468:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceSendExprToSendRecvExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SendRecvExprType
		//line grammar.lr:469:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceRecvExprToSendRecvExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SendRecvExprType
		//line grammar.lr:470:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToSendExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SendExprType
		symbol.Expression, err = reducer.ToSendExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceToRecvExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = RecvExprType
		symbol.Expression, err = reducer.ToRecvExpr(args[0].Value, args[1].Expression)
	case _ReduceSendRecvExprToAssignOpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AssignOpExprType
		//line grammar.lr:483:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAssignOpExprToAssignOpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AssignOpExprType
		//line grammar.lr:484:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryAssignOpExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryAssignOpExprType
		symbol.Expression, err = reducer.ToBinaryAssignOpExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAddAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:490:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:491:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:492:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:493:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:494:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:495:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:496:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:497:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:498:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftAssignToBinaryAssignOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryAssignOpType
		//line grammar.lr:499:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceStatementsToUnlabelledControlFlowExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnlabelledControlFlowExprType
		//line grammar.lr:506:4
		symbol.ControlFlowExpr = args[0].StatementsExpr
		err = nil
	case _ReduceIfElseExprToUnlabelledControlFlowExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnlabelledControlFlowExprType
		//line grammar.lr:507:4
		symbol.ControlFlowExpr = args[0].IfExpr
		err = nil
	case _ReduceSwitchExprBodyToUnlabelledControlFlowExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnlabelledControlFlowExprType
		//line grammar.lr:508:4
		symbol.ControlFlowExpr = args[0].ControlFlowExpr
		err = nil
	case _ReduceSelectExprBodyToUnlabelledControlFlowExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnlabelledControlFlowExprType
		//line grammar.lr:509:4
		symbol.ControlFlowExpr = args[0].ControlFlowExpr
		err = nil
	case _ReduceLoopExprBodyToUnlabelledControlFlowExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnlabelledControlFlowExprType
		//line grammar.lr:510:4
		symbol.ControlFlowExpr = args[0].ControlFlowExpr
		err = nil
	case _ReduceUnlabelledControlFlowExprToControlFlowExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ControlFlowExprType
		//line grammar.lr:513:4
		symbol.ControlFlowExpr = args[0].ControlFlowExpr
		err = nil
	case _ReduceLabelledToControlFlowExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ControlFlowExprType
		symbol.ControlFlowExpr, err = reducer.LabelledToControlFlowExpr(args[0].Value, args[1].Value, args[2].ControlFlowExpr)
	case _ReduceAssignOpExprToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:517:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceControlFlowExprToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:518:4
		symbol.Expression = args[0].ControlFlowExpr
		err = nil
	case _ReduceAddrDeclPatternToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:521:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAssignToAddrPatternToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:522:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToStatements:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementsType
		symbol.StatementsExpr, err = reducer.ToStatements(args[0].Value, args[1].StatementList, args[2].Value)
	case _ReduceAddImplicitToProperStatementList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperStatementListType
		symbol.StatementList, err = reducer.AddImplicitToProperStatementList(args[0].StatementList, args[1].Count, args[2].Statement)
	case _ReduceAddExplicitToProperStatementList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperStatementListType
		symbol.StatementList, err = reducer.AddExplicitToProperStatementList(args[0].StatementList, args[1].Value, args[2].Statement)
	case _ReduceStatementToProperStatementList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperStatementListType
		symbol.StatementList, err = reducer.StatementToProperStatementList(args[0].Statement)
	case _ReduceProperStatementListToStatementList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementListType
		//line grammar.lr:547:4
		symbol.StatementList = args[0].StatementList
		err = nil
	case _ReduceImproperImplicitToStatementList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementListType
		symbol.StatementList, err = reducer.ImproperImplicitToStatementList(args[0].StatementList, args[1].Count)
	case _ReduceImproperExplicitToStatementList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementListType
		symbol.StatementList, err = reducer.ImproperExplicitToStatementList(args[0].StatementList, args[1].Value)
	case _ReduceNilToStatementList:
		symbol.SymbolId_ = StatementListType
		symbol.StatementList, err = reducer.NilToStatementList()
	case _ReduceIfElifExprToIfElseExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = IfElseExprType
		//line grammar.lr:553:4
		symbol.IfExpr = args[0].IfExpr
		err = nil
	case _ReduceElseToIfElseExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfElseExprType
		symbol.IfExpr, err = reducer.ElseToIfElseExpr(args[0].IfExpr, args[1].Value, args[2].StatementsExpr)
	case _ReduceIfOnlyExprToIfElifExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = IfElifExprType
		//line grammar.lr:557:4
		symbol.IfExpr = args[0].IfExpr
		err = nil
	case _ReduceElifToIfElifExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfElifExprType
		symbol.IfExpr, err = reducer.ElifToIfElifExpr(args[0].IfExpr, args[1].Value, args[2].Value, args[3].Expression, args[4].StatementsExpr)
	case _ReduceToIfOnlyExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfOnlyExprType
		symbol.IfExpr, err = reducer.ToIfOnlyExpr(args[0].Value, args[1].Expression, args[2].StatementsExpr)
	case _ReduceExprToCondition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ConditionType
		//line grammar.lr:564:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCasePatternExprToCondition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ConditionType
		//line grammar.lr:565:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToCasePatternExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CasePatternExprType
		symbol.Expression, err = reducer.ToCasePatternExpr(args[0].Value, args[1].ExpressionList, args[2].Value, args[3].Expression)
	case _ReduceToSwitchExprBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SwitchExprBodyType
		symbol.ControlFlowExpr, err = reducer.ToSwitchExprBody(args[0].Value, args[1].Expression, args[2].StatementsExpr)
	case _ReduceToSelectExprBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SelectExprBodyType
		symbol.ControlFlowExpr, err = reducer.ToSelectExprBody(args[0].Value, args[1].StatementsExpr)
	case _ReduceInfiniteToLoopExprBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.ControlFlowExpr, err = reducer.InfiniteToLoopExprBody(args[0].StatementsExpr)
	case _ReduceDoWhileToLoopExprBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.ControlFlowExpr, err = reducer.DoWhileToLoopExprBody(args[0].StatementsExpr, args[1].Value, args[2].Expression)
	case _ReduceWhileToLoopExprBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.ControlFlowExpr, err = reducer.WhileToLoopExprBody(args[0].Value, args[1].Expression, args[2].StatementsExpr)
	case _ReduceIteratorToLoopExprBody:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.ControlFlowExpr, err = reducer.IteratorToLoopExprBody(args[0].Value, args[1].Expression, args[2].Value, args[3].Expression, args[4].StatementsExpr)
	case _ReduceForToLoopExprBody:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.ControlFlowExpr, err = reducer.ForToLoopExprBody(args[0].Value, args[1].Statement, args[2].Value, args[3].Expression, args[4].Value, args[5].Statement, args[6].StatementsExpr)
	case _ReduceStatementToOptionalStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalStatementType
		//line grammar.lr:598:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceNilToOptionalStatement:
		symbol.SymbolId_ = OptionalStatementType
		symbol.Statement, err = reducer.NilToOptionalStatement()
	case _ReduceExprToOptionalExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExprType
		//line grammar.lr:602:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceNilToOptionalExpr:
		symbol.SymbolId_ = OptionalExprType
		symbol.Expression, err = reducer.NilToOptionalExpr()
	case _ReduceToRepeatLoopBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = RepeatLoopBodyType
		symbol.StatementsExpr, err = reducer.ToRepeatLoopBody(args[0].Value, args[1].StatementsExpr)
	case _ReduceToForLoopBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ForLoopBodyType
		symbol.StatementsExpr, err = reducer.ToForLoopBody(args[0].Value, args[1].StatementsExpr)
	case _ReduceExprToReturnableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableExprType
		//line grammar.lr:614:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImproperExprStructToReturnableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableExprType
		//line grammar.lr:615:4
		symbol.Expression = args[0].ImplicitStructExpr
		err = nil
	case _ReducePairToImproperExprStruct:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImproperExprStructType
		symbol.ImplicitStructExpr, err = reducer.PairToImproperExprStruct(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAddToImproperExprStruct:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImproperExprStructType
		symbol.ImplicitStructExpr, err = reducer.AddToImproperExprStruct(args[0].ImplicitStructExpr, args[1].Value, args[2].Expression)
	case _ReduceExplicitStructTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:629:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:630:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceArrayTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:631:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceMapTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:632:4
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
		//line grammar.lr:644:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNamedTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:645:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceInferredTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:646:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitStructTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:647:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:648:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:649:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:650:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncSignatureToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:651:4
		symbol.TypeExpression = args[0].FuncSignature
		err = nil
	case _ReduceLocalToNamedTypeExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = NamedTypeExprType
		symbol.TypeExpression, err = reducer.LocalToNamedTypeExpr(args[0].Value, args[1].TypeExpressionList)
	case _ReduceExternalToNamedTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedTypeExprType
		symbol.TypeExpression, err = reducer.ExternalToNamedTypeExpr(args[0].Value, args[1].Value, args[2].Value, args[3].TypeExpressionList)
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
	case _ReduceAtomTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:670:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixUnaryTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:671:4
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
		//line grammar.lr:677:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:678:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:679:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:680:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:681:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:686:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceBinaryTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:687:4
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
		//line grammar.lr:693:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:694:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:695:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDefinitionToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Statement, err = reducer.DefinitionToTypeDef(args[0].Value, args[1].Value, args[2].GenericParameterList, args[3].TypeExpression)
	case _ReduceConstrainedDefToTypeDef:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = TypeDefType
		symbol.Statement, err = reducer.ConstrainedDefToTypeDef(args[0].Value, args[1].Value, args[2].GenericParameterList, args[3].TypeExpression, args[4].Value, args[5].TypeExpression)
	case _ReduceAliasToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Statement, err = reducer.AliasToTypeDef(args[0].Value, args[1].Value, args[2].Value, args[3].TypeExpression)
	case _ReduceUnconstrainedToGenericParameter:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterType
		symbol.GenericParameter, err = reducer.UnconstrainedToGenericParameter(args[0].Value)
	case _ReduceConstrainedToGenericParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterType
		symbol.GenericParameter, err = reducer.ConstrainedToGenericParameter(args[0].Value, args[1].TypeExpression)
	case _ReduceGenericToGenericParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParametersType
		symbol.GenericParameterList, err = reducer.GenericToGenericParameters(args[0].Value, args[1].GenericParameterList, args[2].Value)
	case _ReduceNilToGenericParameters:
		symbol.SymbolId_ = GenericParametersType
		symbol.GenericParameterList, err = reducer.NilToGenericParameters()
	case _ReduceAddToProperGenericParameterList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericParameterListType
		symbol.GenericParameterList, err = reducer.AddToProperGenericParameterList(args[0].GenericParameterList, args[1].Value, args[2].GenericParameter)
	case _ReduceGenericParameterToProperGenericParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericParameterListType
		symbol.GenericParameterList, err = reducer.GenericParameterToProperGenericParameterList(args[0].GenericParameter)
	case _ReduceProperGenericParameterListToGenericParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterListType
		//line grammar.lr:720:4
		symbol.GenericParameterList = args[0].GenericParameterList
		err = nil
	case _ReduceImproperToGenericParameterList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterListType
		symbol.GenericParameterList, err = reducer.ImproperToGenericParameterList(args[0].GenericParameterList, args[1].Value)
	case _ReduceNilToGenericParameterList:
		symbol.SymbolId_ = GenericParameterListType
		symbol.GenericParameterList, err = reducer.NilToGenericParameterList()
	case _ReduceBindingToGenericArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.TypeExpressionList, err = reducer.BindingToGenericArguments(args[0].Value, args[1].TypeExpressionList, args[2].Value)
	case _ReduceNilToGenericArguments:
		symbol.SymbolId_ = GenericArgumentsType
		symbol.TypeExpressionList, err = reducer.NilToGenericArguments()
	case _ReduceAddToProperGenericArgumentList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericArgumentListType
		symbol.TypeExpressionList, err = reducer.AddToProperGenericArgumentList(args[0].TypeExpressionList, args[1].Value, args[2].TypeExpression)
	case _ReduceTypeExprToProperGenericArgumentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericArgumentListType
		symbol.TypeExpressionList, err = reducer.TypeExprToProperGenericArgumentList(args[0].TypeExpression)
	case _ReduceProperGenericArgumentListToGenericArgumentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericArgumentListType
		//line grammar.lr:733:4
		symbol.TypeExpressionList = args[0].TypeExpressionList
		err = nil
	case _ReduceImproperToGenericArgumentList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericArgumentListType
		symbol.TypeExpressionList, err = reducer.ImproperToGenericArgumentList(args[0].TypeExpressionList, args[1].Value)
	case _ReduceNilToGenericArgumentList:
		symbol.SymbolId_ = GenericArgumentListType
		symbol.TypeExpressionList, err = reducer.NilToGenericArgumentList()
	case _ReduceUnnamedFieldToTypeProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.UnnamedFieldToTypeProperty(args[0].TypeExpression)
	case _ReduceNamedFieldToTypeProperty:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.NamedFieldToTypeProperty(args[0].Value, args[1].TypeExpression)
	case _ReducePaddingFieldToTypeProperty:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.PaddingFieldToTypeProperty(args[0].Value, args[1].TypeExpression)
	case _ReduceDefaultNamedEnumFieldToTypeProperty:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.DefaultNamedEnumFieldToTypeProperty(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceDefaultUnnamedEnumFieldToTypeProperty:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TypePropertyType
		symbol.TypeProperty, err = reducer.DefaultUnnamedEnumFieldToTypeProperty(args[0].Value, args[1].TypeExpression)
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
		//line grammar.lr:760:4
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
		//line grammar.lr:773:4
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
		//line grammar.lr:797:4
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
		//line grammar.lr:811:4
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
		//line grammar.lr:824:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNilToReturnType:
		symbol.SymbolId_ = ReturnTypeType
		symbol.TypeExpression, err = reducer.NilToReturnType()
	case _ReduceNamedArgToParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.NamedArgToParameter(args[0].Value, args[1].TypeExpression)
	case _ReduceNamedReceiverToParameter:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.NamedReceiverToParameter(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceNamedVarargToParameter:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.NamedVarargToParameter(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceIgnoreArgToParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.IgnoreArgToParameter(args[0].Value, args[1].TypeExpression)
	case _ReduceIgnoreReceiverToParameter:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.IgnoreReceiverToParameter(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceIgnoreVarargToParameter:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.IgnoreVarargToParameter(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceUnnamedArgToParameter:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.UnnamedArgToParameter(args[0].TypeExpression)
	case _ReduceUnnamedReceiverToParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.UnnamedReceiverToParameter(args[0].Value, args[1].TypeExpression)
	case _ReduceUnnamedVarargToParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterType
		symbol.Parameter, err = reducer.UnnamedVarargToParameter(args[0].Value, args[1].TypeExpression)
	case _ReduceAddToProperParameterList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterListType
		symbol.Parameters, err = reducer.AddToProperParameterList(args[0].Parameters, args[1].Value, args[2].Parameter)
	case _ReduceParameterToProperParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperParameterListType
		symbol.Parameters, err = reducer.ParameterToProperParameterList(args[0].Parameter)
	case _ReduceProperParameterListToParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterListType
		//line grammar.lr:845:4
		symbol.Parameters = args[0].Parameters
		err = nil
	case _ReduceImproperToParameterList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterListType
		symbol.Parameters, err = reducer.ImproperToParameterList(args[0].Parameters, args[1].Value)
	case _ReduceNilToParameterList:
		symbol.SymbolId_ = ParameterListType
		symbol.Parameters, err = reducer.NilToParameterList()
	case _ReduceToParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParametersType
		symbol.Parameters, err = reducer.ToParameters(args[0].Value, args[1].Parameters, args[2].Value)
	case _ReduceAnonymousToFuncSignature:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FuncSignatureType
		symbol.FuncSignature, err = reducer.AnonymousToFuncSignature(args[0].Value, args[1].Parameters, args[2].TypeExpression)
	case _ReduceNamedToFuncSignature:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FuncSignatureType
		symbol.FuncSignature, err = reducer.NamedToFuncSignature(args[0].Value, args[1].Value, args[2].GenericParameterList, args[3].Parameters, args[4].TypeExpression)
	case _ReduceToFuncDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FuncDefType
		symbol.Expression, err = reducer.ToFuncDef(args[0].FuncSignature, args[1].StatementsExpr)
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("unexpected %s reduce error: %w", act.ReduceType, err)
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
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case SourceType:
			return _Action{_ShiftAction, _State4, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case ProperStatementListType:
			return _Action{_ShiftAction, _State39, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToProperStatementList}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case StatementListType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementListToSource}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToStatementList}, true
		}
	case _State2:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case StatementType:
			return _Action{_ShiftAction, _State5, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State3:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAction, _State6, 0}, true
		}
	case _State4:
		switch symbolId {
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
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
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State43, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State8:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case CasePatternsType:
			return _Action{_ShiftAction, _State45, 0}, true
		case SwitchableCasePatternsType:
			return _Action{_ShiftAction, _State46, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AssignSelectablePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToCasePatterns}, true
		case SwitchableCasePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchableCasePatternToSwitchableCasePatterns}, true
		case EnumPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumPatternToSwitchableCasePattern}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToSwitchableCasePattern}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State9:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State47, 0}, true
		}
	case _State10:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State48, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case OptionalStatementType:
			return _Action{_ShiftAction, _State49, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State50, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToOptionalStatement}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalStatement}, true
		}
	case _State11:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State51, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State52, 0}, true
		case ParametersType:
			return _Action{_ShiftAction, _State53, 0}, true
		}
	case _State12:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignToAddrPattern}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State13:
		switch symbolId {
		case AtToken:
			return _Action{_ShiftAction, _State54, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		}
	case _State14:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State55, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case ConditionType:
			return _Action{_ShiftAction, _State56, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToCondition}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case CasePatternExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCasePatternExprToCondition}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State15:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State58, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State60, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State59, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State57, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSingleToImportStmt}, true
		}
	case _State16:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case ProperStatementListType:
			return _Action{_ShiftAction, _State39, 0}, true
		case StatementListType:
			return _Action{_ShiftAction, _State61, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToProperStatementList}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToStatementList}, true
		}
	case _State17:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State67, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State18:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State69, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State68, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State73, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State70, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State71, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State72, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArgumentToProperArguments}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToArguments}, true
		}
	case _State19:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToRepeatLoopBody}, true
		}
	case _State20:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSelectExprBody}, true
		}
	case _State21:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State74, 0}, true
		}
	case _State22:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State75, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State23:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State76, 0}, true
		}
	case _State24:
		switch symbolId {
		case LessToken:
			return _Action{_ShiftAction, _State77, 0}, true
		}
	case _State25:
		switch symbolId {
		case LbracketToken:
			return _Action{_ShiftAction, _State80, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State79, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State78, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAction, _State81, 0}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPostfixUnaryOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPostfixUnaryOp}, true
		case AddOneAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddOneAssignToPostfixUnaryOp}, true
		case SubOneAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubOneAssignToPostfixUnaryOp}, true
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
			return _Action{_ShiftAction, _State82, 0}, true
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
			return _Action{_ShiftAction, _State83, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}, true
		}
	case _State28:
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
	case _State29:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State85, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExprToReturnableExpr}, true
		}
	case _State30:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFuncDef}, true
		}
	case _State31:
		switch symbolId {
		case ElseToken:
			return _Action{_ShiftAction, _State86, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIfElifExprToIfElseExpr}, true
		}
	case _State32:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State87, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperExprStructToReturnableExpr}, true
		}
	case _State33:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State88, 0}, true
		}
	case _State34:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case AtToken:
			return _Action{_ShiftAction, _State89, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case ReturnableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabeledValuedToJumpStmt}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStmt}, true
		}
	case _State35:
		switch symbolId {
		case MulOpType:
			return _Action{_ShiftAction, _State90, 0}, true
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
	case _State36:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case NewAddressableType:
			return _Action{_ShiftAction, _State91, 0}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToNamedExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToNewAddressable}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToNewAddressable}, true
		}
	case _State37:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State92, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceOrExprToSendRecvExpr}, true
		}
	case _State38:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State39:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State93, 0}, true
		case SemicolonToken:
			return _Action{_ShiftAction, _State94, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperStatementListToStatementList}, true
		}
	case _State40:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State95, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceInfiniteToLoopExprBody}, true
		}
	case _State41:
		switch symbolId {
		case AssignToken:
			return _Action{_ShiftAction, _State96, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceReturnableExprToStatement}, true
		}
	case _State42:
		switch symbolId {
		case ArrowToken:
			return _Action{_ShiftAction, _State97, 0}, true
		case BinaryAssignOpType:
			return _Action{_ShiftAction, _State98, 0}, true
		case AddAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddAssignToBinaryAssignOp}, true
		case SubAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubAssignToBinaryAssignOp}, true
		case MulAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulAssignToBinaryAssignOp}, true
		case DivAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDivAssignToBinaryAssignOp}, true
		case ModAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceModAssignToBinaryAssignOp}, true
		case BitAndAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndAssignToBinaryAssignOp}, true
		case BitOrAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitOrAssignToBinaryAssignOp}, true
		case BitXorAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorAssignToBinaryAssignOp}, true
		case BitLshiftAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitLshiftAssignToBinaryAssignOp}, true
		case BitRshiftAssignToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitRshiftAssignToBinaryAssignOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceSendRecvExprToAssignOpExpr}, true
		}
	case _State43:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State92, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToRecvExpr}, true
		}
	case _State44:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State99, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnnamedStructToEnumPattern}, true
		}
	case _State45:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State100, 0}, true
		}
	case _State46:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State102, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State101, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceSwitchableCasePatternsToCasePatterns}, true
		}
	case _State47:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToOptionalStatement}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case OptionalStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDefaultBranchToBranchStmt}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalStatement}, true
		}
	case _State48:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State103, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State85, 0}, true
		case ForLoopBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWhileToLoopExprBody}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExprToReturnableExpr}, true
		}
	case _State49:
		switch symbolId {
		case SemicolonToken:
			return _Action{_ShiftAction, _State104, 0}, true
		}
	case _State50:
		switch symbolId {
		case InToken:
			return _Action{_ShiftAction, _State105, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State96, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceReturnableExprToStatement}, true
		}
	case _State51:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State106, 0}, true
		case GenericParametersType:
			return _Action{_ShiftAction, _State107, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameters}, true
		}
	case _State52:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State110, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State111, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State108, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State109, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State114, 0}, true
		case ProperParameterListType:
			return _Action{_ShiftAction, _State113, 0}, true
		case ParameterListType:
			return _Action{_ShiftAction, _State112, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ParameterType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParameterToProperParameterList}, true
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToParameterList}, true
		}
	case _State53:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceAnonymousToFuncSignature}, true
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State54:
		switch symbolId {
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabelledToControlFlowExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		}
	case _State55:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case SwitchableCasePatternsType:
			return _Action{_ShiftAction, _State115, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case SwitchableCasePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchableCasePatternToSwitchableCasePatterns}, true
		case EnumPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumPatternToSwitchableCasePattern}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToSwitchableCasePattern}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State56:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIfOnlyExpr}, true
		}
	case _State57:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportToLocalToImportClause}, true
		}
	case _State58:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAliasToImportClause}, true
		}
	case _State59:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State58, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State60, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State57, 0}, true
		case ProperImportClausesType:
			return _Action{_ShiftAction, _State117, 0}, true
		case ImportClausesType:
			return _Action{_ShiftAction, _State116, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportClauseToProperImportClauses}, true
		}
	case _State60:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnusableImportToImportClause}, true
		}
	case _State61:
		switch symbolId {
		case RbraceToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToStatements}, true
		}
	case _State62:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State118, 0}, true
		}
	case _State63:
		switch symbolId {
		case DotToken:
			return _Action{_ShiftAction, _State119, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State78, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLocalToNamedTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State64:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case TypePropertyType:
			return _Action{_ShiftAction, _State128, 0}, true
		case ProperImplicitTypePropertiesType:
			return _Action{_ShiftAction, _State126, 0}, true
		case ImplicitTypePropertiesType:
			return _Action{_ShiftAction, _State124, 0}, true
		case ProperImplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State125, 0}, true
		case ImplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State123, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToImplicitTypeProperties}, true
		}
	case _State65:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State129, 0}, true
		}
	case _State66:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State67:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State131, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State130, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSliceTypeExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State68:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnitExprPairToColonExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}, true
		}
	case _State69:
		switch symbolId {
		case AtToken:
			return _Action{_ShiftAction, _State54, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State133, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIdentifierToNamedExpr}, true
		}
	case _State70:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitStructExpr}, true
		}
	case _State71:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State134, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceColonExprToArgument}, true
		}
	case _State72:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State135, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarargAssignmentToArgument}, true

		default:
			return _Action{_ReduceAction, 0, _ReducePositionalToArgument}, true
		}
	case _State73:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State136, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}, true
		}
	case _State74:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
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
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToExplicitTypeProperties}, true
		}
	case _State75:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToSwitchExprBody}, true
		}
	case _State76:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State106, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State139, 0}, true
		case GenericParametersType:
			return _Action{_ShiftAction, _State140, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameters}, true
		}
	case _State77:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State141, 0}, true
		}
	case _State78:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State144, 0}, true
		case ProperGenericArgumentListType:
			return _Action{_ShiftAction, _State143, 0}, true
		case GenericArgumentListType:
			return _Action{_ShiftAction, _State142, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArgumentList}, true
		}
	case _State79:
		switch symbolId {
		case AsToken:
			return _Action{_ShiftAction, _State145, 0}, true
		case IdentifierToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAccessExpr}, true
		}
	case _State80:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State68, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case IndexType:
			return _Action{_ShiftAction, _State148, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State146, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State147, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State81:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State149, 0}, true
		}
	case _State82:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State150, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State83:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State151, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State84:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State152, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State85:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePairToImproperExprStruct}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State86:
		switch symbolId {
		case IfToken:
			return _Action{_ShiftAction, _State153, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceElseToIfElseExpr}, true
		}
	case _State87:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToImproperExprStruct}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State88:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State69, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State68, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State73, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State154, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State71, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State72, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArgumentToProperArguments}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToArguments}, true
		}
	case _State89:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State155, 0}, true
		}
	case _State90:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State91:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State156, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNewInferredToAddrDeclPattern}, true
		}
	case _State92:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State157, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State93:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperStatementList}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperImplicitToStatementList}, true
		}
	case _State94:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperStatementList}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperExplicitToStatementList}, true
		}
	case _State95:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDoWhileToLoopExprBody}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State96:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case ReturnableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignStmt}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State97:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State158, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State98:
		switch symbolId {
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State159, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
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
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State99:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedToEnumPattern}, true
		}
	case _State100:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToOptionalStatement}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case OptionalStatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCaseBranchToBranchStmt}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalStatement}, true
		}
	case _State101:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAssignSelectablePattern}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State102:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State44, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case SwitchableCasePatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToSwitchableCasePatterns}, true
		case EnumPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEnumPatternToSwitchableCasePattern}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToSwitchableCasePattern}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State103:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToForLoopBody}, true
		}
	case _State104:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case OptionalExprType:
			return _Action{_ShiftAction, _State160, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToOptionalExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalExpr}, true
		}
	case _State105:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State161, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State106:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case ProperGenericParameterListType:
			return _Action{_ShiftAction, _State164, 0}, true
		case GenericParameterListType:
			return _Action{_ShiftAction, _State163, 0}, true
		case GenericParameterType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGenericParameterToProperGenericParameterList}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericParameterList}, true
		}
	case _State107:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State52, 0}, true
		case ParametersType:
			return _Action{_ShiftAction, _State165, 0}, true
		}
	case _State108:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State166, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State109:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
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
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State110:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State78, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State169, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State170, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State171, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State111:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State172, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State173, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State174, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		}
	case _State112:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParameters}, true
		}
	case _State113:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State175, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperParameterListToParameterList}, true
		}
	case _State114:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedArgToParameter}, true
		}
	case _State115:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State102, 0}, true
		case AssignToken:
			return _Action{_ShiftAction, _State176, 0}, true
		}
	case _State116:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMultipleToImportStmt}, true
		}
	case _State117:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State178, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State177, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}, true
		}
	case _State118:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case TypePropertyType:
			return _Action{_ShiftAction, _State181, 0}, true
		case ProperExplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State180, 0}, true
		case ExplicitEnumTypePropertiesType:
			return _Action{_ShiftAction, _State179, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State119:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State182, 0}, true
		}
	case _State120:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State183, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State184, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State121:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State78, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State185, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State122:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
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
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		}
	case _State123:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitEnumTypeExpr}, true
		}
	case _State124:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToImplicitStructTypeExpr}, true
		}
	case _State125:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State187, 0}, true
		case NewlinesToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImproperToImplicitEnumTypeProperties}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImplicitEnumTypePropertiesToImplicitEnumTypeProperties}, true
		}
	case _State126:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State188, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperImplicitTypePropertiesToImplicitTypeProperties}, true
		}
	case _State127:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedFieldToTypeProperty}, true
		}
	case _State128:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State189, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypePropertyToProperImplicitTypeProperties}, true
		}
	case _State129:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case ProperExplicitTypePropertiesType:
			return _Action{_ShiftAction, _State138, 0}, true
		case ExplicitTypePropertiesType:
			return _Action{_ShiftAction, _State190, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToExplicitTypeProperties}, true
		}
	case _State130:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State191, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State131:
		switch symbolId {
		case IntegerLiteralToken:
			return _Action{_ShiftAction, _State192, 0}, true
		}
	case _State132:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State133:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedAssignmentToArgument}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State134:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceColonExprExprTupleToColonExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}, true
		}
	case _State135:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprExprPairToColonExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}, true
		}
	case _State136:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State69, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State68, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State71, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State72, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperArguments}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToArguments}, true
		}
	case _State137:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToExplicitStructTypeExpr}, true
		}
	case _State138:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State194, 0}, true
		case CommaToken:
			return _Action{_ShiftAction, _State193, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperExplicitTypePropertiesToExplicitTypeProperties}, true
		}
	case _State139:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State195, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State140:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State196, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State141:
		switch symbolId {
		case GreaterToken:
			return _Action{_ShiftAction, _State197, 0}, true
		}
	case _State142:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBindingToGenericArguments}, true
		}
	case _State143:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State198, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperGenericArgumentListToGenericArgumentList}, true
		}
	case _State144:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceTypeExprToProperGenericArgumentList}, true
		}
	case _State145:
		switch symbolId {
		case LparenToken:
			return _Action{_ShiftAction, _State199, 0}, true
		}
	case _State146:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State134, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceColonExprToIndex}, true
		}
	case _State147:
		switch symbolId {
		case ColonToken:
			return _Action{_ShiftAction, _State135, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExprToIndex}, true
		}
	case _State148:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIndexExpr}, true
		}
	case _State149:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State69, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ColonToken:
			return _Action{_ShiftAction, _State68, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case ProperArgumentsType:
			return _Action{_ShiftAction, _State73, 0}, true
		case ArgumentsType:
			return _Action{_ShiftAction, _State200, 0}, true
		case ColonExprType:
			return _Action{_ShiftAction, _State71, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State72, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case EllipsisToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSkipPatternToArgument}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
		case ArgumentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArgumentToProperArguments}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToArguments}, true
		}
	case _State150:
		switch symbolId {
		case MulOpType:
			return _Action{_ShiftAction, _State90, 0}, true
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
	case _State151:
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
	case _State152:
		switch symbolId {
		case AddOpType:
			return _Action{_ShiftAction, _State82, 0}, true
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
	case _State153:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State55, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case ConditionType:
			return _Action{_ShiftAction, _State201, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExprToCondition}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case CasePatternExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCasePatternExprToCondition}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State154:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToInitializeExpr}, true
		}
	case _State155:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case ReturnableExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLabeledValuedToJumpStmt}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStmt}, true
		}
	case _State156:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNewTypedToAddrDeclPattern}, true
		}
	case _State157:
		switch symbolId {
		case AndToken:
			return _Action{_ShiftAction, _State83, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}, true
		}
	case _State158:
		switch symbolId {
		case OrToken:
			return _Action{_ShiftAction, _State92, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToSendExpr}, true
		}
	case _State159:
		switch symbolId {
		case ArrowToken:
			return _Action{_ShiftAction, _State97, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToBinaryAssignOpExpr}, true
		}
	case _State160:
		switch symbolId {
		case SemicolonToken:
			return _Action{_ShiftAction, _State202, 0}, true
		}
	case _State161:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State103, 0}, true
		case ForLoopBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIteratorToLoopExprBody}, true
		}
	case _State162:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State203, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameter}, true
		}
	case _State163:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceGenericToGenericParameters}, true
		}
	case _State164:
		switch symbolId {
		case CommaToken:
			return _Action{_ShiftAction, _State204, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperGenericParameterListToGenericParameterList}, true
		}
	case _State165:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedToFuncSignature}, true
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToReturnType}, true
		}
	case _State166:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameter}, true
		}
	case _State167:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceUnnamedReceiverToParameter}, true
		}
	case _State168:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State182, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		}
	case _State169:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State205, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State170:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State206, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State171:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedArgToParameter}, true
		}
	case _State172:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State207, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State173:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State208, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State174:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreArgToParameter}, true
		}
	case _State175:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State110, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State111, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case EllipsisToken:
			return _Action{_ShiftAction, _State108, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State109, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State114, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ParameterType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperParameterList}, true
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToParameterList}, true
		}
	case _State176:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
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
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case ExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToCasePatternExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true
		}
	case _State177:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State58, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State60, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State57, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddExplicitToProperImportClauses}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}, true
		}
	case _State178:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State58, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State60, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State57, 0}, true
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStringLiteralToImportClause}, true
		case ImportClauseType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddImplicitToProperImportClauses}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}, true
		}
	case _State179:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToExplicitEnumTypeExpr}, true
		}
	case _State180:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State209, 0}, true
		case OrToken:
			return _Action{_ShiftAction, _State210, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceProperExplicitEnumTypePropertiesToExplicitEnumTypeProperties}, true
		}
	case _State181:
		switch symbolId {
		case NewlinesToken:
			return _Action{_ShiftAction, _State211, 0}, true
		case OrToken:
			return _Action{_ShiftAction, _State212, 0}, true
		}
	case _State182:
		switch symbolId {
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State78, 0}, true
		case GenericArgumentsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExternalToNamedTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State183:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case DotToken:
			return _Action{_ShiftAction, _State168, 0}, true
		case DollarLbracketToken:
			return _Action{_ShiftAction, _State78, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State213, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToGenericArguments}, true
		}
	case _State184:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDefaultUnnamedEnumFieldToTypeProperty}, true
		}
	case _State185:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedFieldToTypeProperty}, true
		}
	case _State186:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReducePaddingFieldToTypeProperty}, true
		}
	case _State187:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State188:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToImplicitTypeProperties}, true
		}
	case _State189:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State190:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToTraitTypeExpr}, true
		}
	case _State191:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToMapTypeExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State192:
		switch symbolId {
		case RbracketToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToArrayTypeExpr}, true
		}
	case _State193:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitTypeProperties}, true
		}
	case _State194:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitTypeProperties}, true
		}
	case _State195:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAliasToTypeDef}, true
		}
	case _State196:
		switch symbolId {
		case ImplementsToken:
			return _Action{_ShiftAction, _State214, 0}, true
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}, true
		}
	case _State197:
		switch symbolId {
		case StringLiteralToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToUnsafeStmt}, true
		}
	case _State198:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State215, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToGenericArgumentList}, true
		}
	case _State199:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
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
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State200:
		switch symbolId {
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToCallExpr}, true
		}
	case _State201:
		switch symbolId {
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceElifToIfElifExpr}, true
		}
	case _State202:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case CaseToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case RepeatToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case ForToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case SelectToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case ImportToken:
			return _Action{_ShiftAction, _State15, 0}, true
		case UnsafeToken:
			return _Action{_ShiftAction, _State24, 0}, true
		case TypeToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LbraceToken:
			return _Action{_ShiftAction, _State16, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State18, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case ArrowToken:
			return _Action{_ShiftAction, _State7, 0}, true
		case GreaterToken:
			return _Action{_ShiftAction, _State12, 0}, true
		case JumpOpType:
			return _Action{_ShiftAction, _State34, 0}, true
		case NewVarTypeType:
			return _Action{_ShiftAction, _State36, 0}, true
		case AccessibleExprType:
			return _Action{_ShiftAction, _State25, 0}, true
		case PrefixUnaryOpType:
			return _Action{_ShiftAction, _State38, 0}, true
		case MulExprType:
			return _Action{_ShiftAction, _State35, 0}, true
		case AddExprType:
			return _Action{_ShiftAction, _State26, 0}, true
		case CmpExprType:
			return _Action{_ShiftAction, _State28, 0}, true
		case AndExprType:
			return _Action{_ShiftAction, _State27, 0}, true
		case OrExprType:
			return _Action{_ShiftAction, _State37, 0}, true
		case SendRecvExprType:
			return _Action{_ShiftAction, _State42, 0}, true
		case ExprType:
			return _Action{_ShiftAction, _State29, 0}, true
		case IfElifExprType:
			return _Action{_ShiftAction, _State31, 0}, true
		case OptionalStatementType:
			return _Action{_ShiftAction, _State217, 0}, true
		case RepeatLoopBodyType:
			return _Action{_ShiftAction, _State40, 0}, true
		case ReturnableExprType:
			return _Action{_ShiftAction, _State41, 0}, true
		case ImproperExprStructType:
			return _Action{_ShiftAction, _State32, 0}, true
		case InitializableTypeExprType:
			return _Action{_ShiftAction, _State33, 0}, true
		case FuncSignatureType:
			return _Action{_ShiftAction, _State30, 0}, true
		case CommentGroupsToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFloatingComment}, true
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
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToJumpOp}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToJumpOp}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToJumpOp}, true
		case FallthroughToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFallthroughToJumpStmt}, true
		case AsyncToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsyncToPrefixUnaryOp}, true
		case DeferToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDeferToPrefixUnaryOp}, true
		case VarToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceVarToNewVarType}, true
		case LetToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLetToNewVarType}, true
		case NotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNotToPrefixUnaryOp}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToPrefixUnaryOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToPrefixUnaryOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToPrefixUnaryOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}, true
		case BitXorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitXorToPrefixUnaryOp}, true
		case ParseErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToParseErrorExpr}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementToOptionalStatement}, true
		case FloatingCommentType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFloatingCommentToStatement}, true
		case BranchStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBranchStmtToStatement}, true
		case UnsafeStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnsafeStmtToStatement}, true
		case JumpStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceJumpStmtToStatement}, true
		case AssignStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignStmtToStatement}, true
		case ImportStmtType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportStmtToStatement}, true
		case AddrDeclPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddrDeclPatternToExpr}, true
		case AssignToAddrPatternType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignToAddrPatternToExpr}, true
		case AtomExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomExprToAccessibleExpr}, true
		case ParseErrorExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceParseErrorExprToAtomExpr}, true
		case LiteralExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLiteralExprToAtomExpr}, true
		case NamedExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceNamedExprToAtomExpr}, true
		case InitializeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceInitializeExprToAtomExpr}, true
		case ImplicitStructExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}, true
		case AccessExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAccessExprToAccessibleExpr}, true
		case IndexExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIndexExprToAccessibleExpr}, true
		case AsExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAsExprToAccessibleExpr}, true
		case CallExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCallExprToAccessibleExpr}, true
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
		case SendExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSendExprToSendRecvExpr}, true
		case RecvExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceRecvExprToSendRecvExpr}, true
		case AssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAssignOpExprToExpr}, true
		case BinaryAssignOpExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryAssignOpExprToAssignOpExpr}, true
		case UnlabelledControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnlabelledControlFlowExprToControlFlowExpr}, true
		case ControlFlowExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceControlFlowExprToExpr}, true
		case StatementsType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceStatementsToUnlabelledControlFlowExpr}, true
		case IfElseExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfElseExprToUnlabelledControlFlowExpr}, true
		case IfOnlyExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfOnlyExprToIfElifExpr}, true
		case SwitchExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchExprBodyToUnlabelledControlFlowExpr}, true
		case SelectExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSelectExprBodyToUnlabelledControlFlowExpr}, true
		case LoopExprBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceLoopExprBodyToUnlabelledControlFlowExpr}, true
		case SliceTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}, true
		case ArrayTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}, true
		case MapTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}, true
		case TypeDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTypeDefToStatement}, true
		case ExplicitStructTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}, true
		case FuncDefType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncDefToAtomExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalStatement}, true
		}
	case _State203:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameter}, true
		}
	case _State204:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State162, 0}, true
		case GenericParameterType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToProperGenericParameterList}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToGenericParameterList}, true
		}
	case _State205:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedVarargToParameter}, true
		}
	case _State206:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNamedReceiverToParameter}, true
		}
	case _State207:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreVarargToParameter}, true
		}
	case _State208:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceIgnoreReceiverToParameter}, true
		}
	case _State209:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceImproperToExplicitEnumTypeProperties}, true
		}
	case _State210:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State211:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State212:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State121, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAction, _State122, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State120, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State127, 0}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case ReturnableTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}, true
		case PrefixUnaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}, true
		case BinaryTypeExprType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State213:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDefaultNamedEnumFieldToTypeProperty}, true
		}
	case _State214:
		switch symbolId {
		case IdentifierToken:
			return _Action{_ShiftAction, _State63, 0}, true
		case StructToken:
			return _Action{_ShiftAction, _State21, 0}, true
		case EnumToken:
			return _Action{_ShiftAction, _State62, 0}, true
		case TraitToken:
			return _Action{_ShiftAction, _State65, 0}, true
		case FuncToken:
			return _Action{_ShiftAction, _State11, 0}, true
		case LparenToken:
			return _Action{_ShiftAction, _State64, 0}, true
		case LbracketToken:
			return _Action{_ShiftAction, _State17, 0}, true
		case PrefixUnaryTypeOpType:
			return _Action{_ShiftAction, _State66, 0}, true
		case TypeExprType:
			return _Action{_ShiftAction, _State218, 0}, true
		case UnderscoreToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}, true
		case DotToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceDotToInferredTypeExpr}, true
		case QuestionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}, true
		case ExclaimToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}, true
		case TildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeToPrefixUnaryTypeOp}, true
		case TildeTildeToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}, true
		case BitAndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}, true
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
		case FuncSignatureType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceFuncSignatureToAtomTypeExpr}, true
		}
	case _State215:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAddToProperGenericArgumentList}, true
		}
	case _State216:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case RparenToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToAsExpr}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true
		}
	case _State217:
		switch symbolId {
		case DoToken:
			return _Action{_ShiftAction, _State103, 0}, true
		case ForLoopBodyType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToLoopExprBody}, true
		}
	case _State218:
		switch symbolId {
		case BinaryTypeOpType:
			return _Action{_ShiftAction, _State132, 0}, true
		case AddToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBinaryTypeOp}, true
		case SubToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubToBinaryTypeOp}, true
		case MulToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceMulToBinaryTypeOp}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}, true
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
      * -> [statement_list]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [proper_statement_list]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      statement_list -> [source]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      source -> State 4
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      proper_statement_list -> State 39
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 2:
    Kernel Items:
      #accept: ^.statement
    Reduce:
      (nil)
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      statement -> State 5
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 3:
    Kernel Items:
      #accept: ^.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LBRACE -> State 16
      statements -> State 6

  State 4:
    Kernel Items:
      #accept: ^ source., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 5:
    Kernel Items:
      #accept: ^ statement., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 6:
    Kernel Items:
      #accept: ^ statements., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 7:
    Kernel Items:
      recv_expr: ARROW.or_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
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
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 43
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 8:
    Kernel Items:
      branch_stmt: CASE.case_patterns COLON optional_statement
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      assign_selectable_pattern -> [case_patterns]
      switchable_case_pattern -> [switchable_case_patterns]
      enum_pattern -> [switchable_case_pattern]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [switchable_case_pattern]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      DOT -> State 44
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      case_patterns -> State 45
      switchable_case_patterns -> State 46
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 9:
    Kernel Items:
      branch_stmt: DEFAULT.COLON optional_statement
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 47

  State 10:
    Kernel Items:
      loop_expr_body: FOR.expr for_loop_body
      loop_expr_body: FOR.returnable_expr IN expr for_loop_body
      loop_expr_body: FOR.optional_statement SEMICOLON optional_expr SEMICOLON optional_statement for_loop_body
    Reduce:
      * -> [optional_statement]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [optional_statement]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 48
      if_elif_expr -> State 31
      optional_statement -> State 49
      repeat_loop_body -> State 40
      returnable_expr -> State 50
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 11:
    Kernel Items:
      func_signature: FUNC.parameters return_type
      func_signature: FUNC.IDENTIFIER generic_parameters parameters return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 51
      LPAREN -> State 52
      parameters -> State 53

  State 12:
    Kernel Items:
      assign_to_addr_pattern: GREATER.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [assign_to_addr_pattern]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 13:
    Kernel Items:
      named_expr: IDENTIFIER., *
      control_flow_expr: IDENTIFIER.AT unlabelled_control_flow_expr
    Reduce:
      * -> [named_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AT -> State 54

  State 14:
    Kernel Items:
      if_only_expr: IF.condition statements
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [condition]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      case_pattern_expr -> [condition]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 55
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      condition -> State 56
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 15:
    Kernel Items:
      import_stmt: IMPORT.import_clause
      import_stmt: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [import_stmt]
    Goto:
      IDENTIFIER -> State 58
      UNDERSCORE -> State 60
      LPAREN -> State 59
      DOT -> State 57

  State 16:
    Kernel Items:
      statements: LBRACE.statement_list RBRACE
    Reduce:
      * -> [statement_list]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [proper_statement_list]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      proper_statement_list -> State 39
      statement_list -> State 61
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 17:
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
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 67

  State 18:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ELLIPSIS -> [argument]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 69
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      COLON -> State 68
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      proper_arguments -> State 73
      arguments -> State 70
      colon_expr -> State 71
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 72
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 19:
    Kernel Items:
      repeat_loop_body: REPEAT.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [repeat_loop_body]
    Goto:
      LBRACE -> State 16

  State 20:
    Kernel Items:
      select_expr_body: SELECT.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [select_expr_body]
    Goto:
      LBRACE -> State 16

  State 21:
    Kernel Items:
      explicit_struct_type_expr: STRUCT.LPAREN explicit_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 74

  State 22:
    Kernel Items:
      switch_expr_body: SWITCH.expr statements
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 75
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 23:
    Kernel Items:
      type_def: TYPE.IDENTIFIER generic_parameters type_expr
      type_def: TYPE.IDENTIFIER generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE.IDENTIFIER ASSIGN type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76

  State 24:
    Kernel Items:
      unsafe_stmt: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LESS -> State 77

  State 25:
    Kernel Items:
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET index RBRACKET
      as_expr: accessible_expr.DOT AS LPAREN type_expr RPAREN
      call_expr: accessible_expr.generic_arguments LPAREN arguments RPAREN
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.postfix_unary_op
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_arguments]
    ShiftAndReduce:
      QUESTION -> [postfix_unary_op]
      EXCLAIM -> [postfix_unary_op]
      ADD_ONE_ASSIGN -> [postfix_unary_op]
      SUB_ONE_ASSIGN -> [postfix_unary_op]
      postfix_unary_op -> [postfix_unary_expr]
    Goto:
      LBRACKET -> State 80
      DOT -> State 79
      DOLLAR_LBRACKET -> State 78
      generic_arguments -> State 81

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
      add_op -> State 82

  State 27:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AND -> State 83

  State 28:
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

  State 29:
    Kernel Items:
      returnable_expr: expr., *
      improper_expr_struct: expr.COMMA expr
    Reduce:
      * -> [returnable_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 85

  State 30:
    Kernel Items:
      func_def: func_signature.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [func_def]
    Goto:
      LBRACE -> State 16

  State 31:
    Kernel Items:
      if_else_expr: if_elif_expr., *
      if_else_expr: if_elif_expr.ELSE statements
      if_elif_expr: if_elif_expr.ELSE IF condition statements
    Reduce:
      * -> [if_else_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      ELSE -> State 86

  State 32:
    Kernel Items:
      returnable_expr: improper_expr_struct., *
      improper_expr_struct: improper_expr_struct.COMMA expr
    Reduce:
      * -> [returnable_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 87

  State 33:
    Kernel Items:
      initialize_expr: initializable_type_expr.LPAREN arguments RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 88

  State 34:
    Kernel Items:
      jump_stmt: jump_op., *
      jump_stmt: jump_op.returnable_expr
      jump_stmt: jump_op.AT IDENTIFIER
      jump_stmt: jump_op.AT IDENTIFIER returnable_expr
    Reduce:
      * -> [jump_stmt]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      returnable_expr -> [jump_stmt]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      AT -> State 89
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 35:
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
      mul_op -> State 90

  State 36:
    Kernel Items:
      addr_decl_pattern: new_var_type.new_addressable
      addr_decl_pattern: new_var_type.new_addressable type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      named_expr -> [new_addressable]
      implicit_struct_expr -> [new_addressable]
    Goto:
      LPAREN -> State 18
      new_addressable -> State 91

  State 37:
    Kernel Items:
      binary_or_expr: or_expr.OR and_expr
      send_recv_expr: or_expr., *
    Reduce:
      * -> [send_recv_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 92

  State 38:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefixable_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [prefix_unary_expr]
      prefix_unary_expr -> [prefixable_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 39:
    Kernel Items:
      proper_statement_list: proper_statement_list.NEWLINES statement
      proper_statement_list: proper_statement_list.SEMICOLON statement
      statement_list: proper_statement_list., *
      statement_list: proper_statement_list.NEWLINES
      statement_list: proper_statement_list.SEMICOLON
    Reduce:
      * -> [statement_list]
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 93
      SEMICOLON -> State 94

  State 40:
    Kernel Items:
      loop_expr_body: repeat_loop_body., *
      loop_expr_body: repeat_loop_body.FOR expr
    Reduce:
      * -> [loop_expr_body]
    ShiftAndReduce:
      (nil)
    Goto:
      FOR -> State 95

  State 41:
    Kernel Items:
      statement: returnable_expr., *
      assign_stmt: returnable_expr.ASSIGN returnable_expr
    Reduce:
      * -> [statement]
    ShiftAndReduce:
      (nil)
    Goto:
      ASSIGN -> State 96

  State 42:
    Kernel Items:
      send_expr: send_recv_expr.ARROW or_expr
      assign_op_expr: send_recv_expr., *
      binary_assign_op_expr: send_recv_expr.binary_assign_op send_recv_expr
    Reduce:
      * -> [assign_op_expr]
    ShiftAndReduce:
      ADD_ASSIGN -> [binary_assign_op]
      SUB_ASSIGN -> [binary_assign_op]
      MUL_ASSIGN -> [binary_assign_op]
      DIV_ASSIGN -> [binary_assign_op]
      MOD_ASSIGN -> [binary_assign_op]
      BIT_AND_ASSIGN -> [binary_assign_op]
      BIT_OR_ASSIGN -> [binary_assign_op]
      BIT_XOR_ASSIGN -> [binary_assign_op]
      BIT_LSHIFT_ASSIGN -> [binary_assign_op]
      BIT_RSHIFT_ASSIGN -> [binary_assign_op]
    Goto:
      ARROW -> State 97
      binary_assign_op -> State 98

  State 43:
    Kernel Items:
      binary_or_expr: or_expr.OR and_expr
      recv_expr: ARROW or_expr., *
    Reduce:
      * -> [recv_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 92

  State 44:
    Kernel Items:
      enum_pattern: DOT.IDENTIFIER implicit_struct_expr
      enum_pattern: DOT.implicit_struct_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      implicit_struct_expr -> [enum_pattern]
    Goto:
      IDENTIFIER -> State 99
      LPAREN -> State 18

  State 45:
    Kernel Items:
      branch_stmt: CASE case_patterns.COLON optional_statement
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 100

  State 46:
    Kernel Items:
      case_patterns: switchable_case_patterns., *
      assign_selectable_pattern: switchable_case_patterns.ASSIGN expr
      switchable_case_patterns: switchable_case_patterns.COMMA switchable_case_pattern
    Reduce:
      * -> [case_patterns]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 102
      ASSIGN -> State 101

  State 47:
    Kernel Items:
      branch_stmt: DEFAULT COLON.optional_statement
    Reduce:
      * -> [optional_statement]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [optional_statement]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      optional_statement -> [branch_stmt]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 48:
    Kernel Items:
      loop_expr_body: FOR expr.for_loop_body
      returnable_expr: expr., *
      improper_expr_struct: expr.COMMA expr
    Reduce:
      * -> [returnable_expr]
    ShiftAndReduce:
      for_loop_body -> [loop_expr_body]
    Goto:
      DO -> State 103
      COMMA -> State 85

  State 49:
    Kernel Items:
      loop_expr_body: FOR optional_statement.SEMICOLON optional_expr SEMICOLON optional_statement for_loop_body
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      SEMICOLON -> State 104

  State 50:
    Kernel Items:
      statement: returnable_expr., *
      assign_stmt: returnable_expr.ASSIGN returnable_expr
      loop_expr_body: FOR returnable_expr.IN expr for_loop_body
    Reduce:
      * -> [statement]
    ShiftAndReduce:
      (nil)
    Goto:
      IN -> State 105
      ASSIGN -> State 96

  State 51:
    Kernel Items:
      func_signature: FUNC IDENTIFIER.generic_parameters parameters return_type
    Reduce:
      * -> [generic_parameters]
    ShiftAndReduce:
      (nil)
    Goto:
      DOLLAR_LBRACKET -> State 106
      generic_parameters -> State 107

  State 52:
    Kernel Items:
      parameters: LPAREN.parameter_list RPAREN
    Reduce:
      * -> [parameter_list]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      parameter -> [proper_parameter_list]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 110
      UNDERSCORE -> State 111
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      ELLIPSIS -> State 108
      GREATER -> State 109
      prefix_unary_type_op -> State 66
      type_expr -> State 114
      proper_parameter_list -> State 113
      parameter_list -> State 112

  State 53:
    Kernel Items:
      func_signature: FUNC parameters.return_type
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      return_type -> [func_signature]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66

  State 54:
    Kernel Items:
      control_flow_expr: IDENTIFIER AT.unlabelled_control_flow_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      unlabelled_control_flow_expr -> [control_flow_expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
    Goto:
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      LBRACE -> State 16
      if_elif_expr -> State 31
      repeat_loop_body -> State 40

  State 55:
    Kernel Items:
      case_pattern_expr: CASE.switchable_case_patterns ASSIGN expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      switchable_case_pattern -> [switchable_case_patterns]
      enum_pattern -> [switchable_case_pattern]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [switchable_case_pattern]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      DOT -> State 44
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      switchable_case_patterns -> State 115
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 56:
    Kernel Items:
      if_only_expr: IF condition.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [if_only_expr]
    Goto:
      LBRACE -> State 16

  State 57:
    Kernel Items:
      import_clause: DOT.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      import_clause: IDENTIFIER.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      import_stmt: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 58
      UNDERSCORE -> State 60
      DOT -> State 57
      proper_import_clauses -> State 117
      import_clauses -> State 116

  State 60:
    Kernel Items:
      import_clause: UNDERSCORE.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      statements: LBRACE statement_list.RBRACE
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACE -> [statements]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      explicit_enum_type_expr: ENUM.LPAREN explicit_enum_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 118

  State 63:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      generic_arguments -> [named_type_expr]
    Goto:
      DOT -> State 119
      DOLLAR_LBRACKET -> State 78

  State 64:
    Kernel Items:
      implicit_struct_type_expr: LPAREN.implicit_type_properties RPAREN
      implicit_enum_type_expr: LPAREN.implicit_enum_type_properties RPAREN
    Reduce:
      * -> [implicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127
      type_property -> State 128
      proper_implicit_type_properties -> State 126
      implicit_type_properties -> State 124
      proper_implicit_enum_type_properties -> State 125
      implicit_enum_type_properties -> State 123

  State 65:
    Kernel Items:
      trait_type_expr: TRAIT.LPAREN explicit_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 129

  State 66:
    Kernel Items:
      prefix_unary_type_expr: prefix_unary_type_op.returnable_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [prefix_unary_type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66

  State 67:
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
      COMMA -> State 131
      COLON -> State 130
      binary_type_op -> State 132

  State 68:
    Kernel Items:
      colon_expr: COLON., *
      colon_expr: COLON.expr
    Reduce:
      * -> [colon_expr]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [colon_expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 69:
    Kernel Items:
      named_expr: IDENTIFIER., *
      argument: IDENTIFIER.ASSIGN expr
      control_flow_expr: IDENTIFIER.AT unlabelled_control_flow_expr
    Reduce:
      * -> [named_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AT -> State 54
      ASSIGN -> State 133

  State 70:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_struct_expr]
    Goto:
      (nil)

  State 71:
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

  State 72:
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

  State 73:
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

  State 74:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN.explicit_type_properties RPAREN
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127
      proper_explicit_type_properties -> State 138
      explicit_type_properties -> State 137

  State 75:
    Kernel Items:
      switch_expr_body: SWITCH expr.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [switch_expr_body]
    Goto:
      LBRACE -> State 16

  State 76:
    Kernel Items:
      type_def: TYPE IDENTIFIER.generic_parameters type_expr
      type_def: TYPE IDENTIFIER.generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [generic_parameters]
    ShiftAndReduce:
      (nil)
    Goto:
      DOLLAR_LBRACKET -> State 106
      ASSIGN -> State 139
      generic_parameters -> State 140

  State 77:
    Kernel Items:
      unsafe_stmt: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141

  State 78:
    Kernel Items:
      generic_arguments: DOLLAR_LBRACKET.generic_argument_list RBRACKET
    Reduce:
      * -> [generic_argument_list]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 144
      proper_generic_argument_list -> State 143
      generic_argument_list -> State 142

  State 79:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
      as_expr: accessible_expr DOT.AS LPAREN type_expr RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [access_expr]
    Goto:
      AS -> State 145

  State 80:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.index RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      COLON -> State 68
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      index -> State 148
      colon_expr -> State 146
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 147
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 81:
    Kernel Items:
      call_expr: accessible_expr generic_arguments.LPAREN arguments RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 149

  State 82:
    Kernel Items:
      binary_add_expr: add_expr add_op.mul_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 150
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 83:
    Kernel Items:
      binary_and_expr: and_expr AND.cmp_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
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
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 151
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 84:
    Kernel Items:
      binary_cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
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
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 152
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 85:
    Kernel Items:
      improper_expr_struct: expr COMMA.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [improper_expr_struct]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 86:
    Kernel Items:
      if_else_expr: if_elif_expr ELSE.statements
      if_elif_expr: if_elif_expr ELSE.IF condition statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [if_else_expr]
    Goto:
      IF -> State 153
      LBRACE -> State 16

  State 87:
    Kernel Items:
      improper_expr_struct: improper_expr_struct COMMA.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [improper_expr_struct]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 88:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN.arguments RPAREN
    Reduce:
      * -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ELLIPSIS -> [argument]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 69
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      COLON -> State 68
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      proper_arguments -> State 73
      arguments -> State 154
      colon_expr -> State 71
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 72
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 89:
    Kernel Items:
      jump_stmt: jump_op AT.IDENTIFIER
      jump_stmt: jump_op AT.IDENTIFIER returnable_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155

  State 90:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op.prefixable_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [binary_mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 91:
    Kernel Items:
      addr_decl_pattern: new_var_type new_addressable., *
      addr_decl_pattern: new_var_type new_addressable.type_expr
    Reduce:
      * -> [addr_decl_pattern]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 156

  State 92:
    Kernel Items:
      binary_or_expr: or_expr OR.and_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
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
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 157
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 93:
    Kernel Items:
      proper_statement_list: proper_statement_list NEWLINES.statement
      statement_list: proper_statement_list NEWLINES., *
    Reduce:
      * -> [statement_list]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [proper_statement_list]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 94:
    Kernel Items:
      proper_statement_list: proper_statement_list SEMICOLON.statement
      statement_list: proper_statement_list SEMICOLON., *
    Reduce:
      * -> [statement_list]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [proper_statement_list]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 95:
    Kernel Items:
      loop_expr_body: repeat_loop_body FOR.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [loop_expr_body]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 96:
    Kernel Items:
      assign_stmt: returnable_expr ASSIGN.returnable_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      returnable_expr -> [assign_stmt]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 97:
    Kernel Items:
      send_expr: send_recv_expr ARROW.or_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
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
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 158
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 98:
    Kernel Items:
      binary_assign_op_expr: send_recv_expr binary_assign_op.send_recv_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      IDENTIFIER -> [named_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      STRUCT -> State 21
      FUNC -> State 11
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 159
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 99:
    Kernel Items:
      enum_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      implicit_struct_expr -> [enum_pattern]
    Goto:
      LPAREN -> State 18

  State 100:
    Kernel Items:
      branch_stmt: CASE case_patterns COLON.optional_statement
    Reduce:
      * -> [optional_statement]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [optional_statement]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      optional_statement -> [branch_stmt]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 101:
    Kernel Items:
      assign_selectable_pattern: switchable_case_patterns ASSIGN.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [assign_selectable_pattern]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 102:
    Kernel Items:
      switchable_case_patterns: switchable_case_patterns COMMA.switchable_case_pattern
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      switchable_case_pattern -> [switchable_case_patterns]
      enum_pattern -> [switchable_case_pattern]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [switchable_case_pattern]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      DOT -> State 44
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 103:
    Kernel Items:
      for_loop_body: DO.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [for_loop_body]
    Goto:
      LBRACE -> State 16

  State 104:
    Kernel Items:
      loop_expr_body: FOR optional_statement SEMICOLON.optional_expr SEMICOLON optional_statement for_loop_body
    Reduce:
      * -> [optional_expr]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [optional_expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      optional_expr -> State 160
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 105:
    Kernel Items:
      loop_expr_body: FOR returnable_expr IN.expr for_loop_body
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 161
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 106:
    Kernel Items:
      generic_parameters: DOLLAR_LBRACKET.generic_parameter_list RBRACKET
    Reduce:
      * -> [generic_parameter_list]
    ShiftAndReduce:
      generic_parameter -> [proper_generic_parameter_list]
    Goto:
      IDENTIFIER -> State 162
      proper_generic_parameter_list -> State 164
      generic_parameter_list -> State 163

  State 107:
    Kernel Items:
      func_signature: FUNC IDENTIFIER generic_parameters.parameters return_type
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 52
      parameters -> State 165

  State 108:
    Kernel Items:
      parameter: ELLIPSIS.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 166

  State 109:
    Kernel Items:
      parameter: GREATER.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 167

  State 110:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
      parameter: IDENTIFIER.type_expr
      parameter: IDENTIFIER.GREATER type_expr
      parameter: IDENTIFIER.ELLIPSIS type_expr
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      generic_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      DOT -> State 168
      DOLLAR_LBRACKET -> State 78
      ELLIPSIS -> State 169
      GREATER -> State 170
      prefix_unary_type_op -> State 66
      type_expr -> State 171

  State 111:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      parameter: UNDERSCORE.type_expr
      parameter: UNDERSCORE.GREATER type_expr
      parameter: UNDERSCORE.ELLIPSIS type_expr
    Reduce:
      * -> [inferred_type_expr]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      ELLIPSIS -> State 172
      GREATER -> State 173
      prefix_unary_type_op -> State 66
      type_expr -> State 174

  State 112:
    Kernel Items:
      parameters: LPAREN parameter_list.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      proper_parameter_list: proper_parameter_list.COMMA parameter
      parameter_list: proper_parameter_list., *
      parameter_list: proper_parameter_list.COMMA
    Reduce:
      * -> [parameter_list]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 175

  State 114:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 115:
    Kernel Items:
      switchable_case_patterns: switchable_case_patterns.COMMA switchable_case_pattern
      case_pattern_expr: CASE switchable_case_patterns.ASSIGN expr
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 102
      ASSIGN -> State 176

  State 116:
    Kernel Items:
      import_stmt: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [import_stmt]
    Goto:
      (nil)

  State 117:
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
      NEWLINES -> State 178
      COMMA -> State 177

  State 118:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN.explicit_enum_type_properties RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127
      type_property -> State 181
      proper_explicit_enum_type_properties -> State 180
      explicit_enum_type_properties -> State 179

  State 119:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_arguments
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 182

  State 120:
    Kernel Items:
      type_property: DEFAULT.IDENTIFIER type_expr
      type_property: DEFAULT.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 183
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 184

  State 121:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
      type_property: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      generic_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      DOT -> State 168
      DOLLAR_LBRACKET -> State 78
      prefix_unary_type_op -> State 66
      type_expr -> State 185

  State 122:
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
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 186

  State 123:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      proper_implicit_enum_type_properties: proper_implicit_enum_type_properties.OR type_property
      implicit_enum_type_properties: proper_implicit_enum_type_properties., *
      implicit_enum_type_properties: proper_implicit_enum_type_properties.NEWLINES
    Reduce:
      * -> [implicit_enum_type_properties]
    ShiftAndReduce:
      NEWLINES -> [implicit_enum_type_properties]
    Goto:
      OR -> State 187

  State 126:
    Kernel Items:
      proper_implicit_type_properties: proper_implicit_type_properties.COMMA type_property
      implicit_type_properties: proper_implicit_type_properties., *
      implicit_type_properties: proper_implicit_type_properties.COMMA
    Reduce:
      * -> [implicit_type_properties]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 188

  State 127:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_property: type_expr., *
    Reduce:
      * -> [type_property]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 128:
    Kernel Items:
      proper_implicit_type_properties: type_property., *
      proper_implicit_enum_type_properties: type_property.OR type_property
    Reduce:
      * -> [proper_implicit_type_properties]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 189

  State 129:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN.explicit_type_properties RPAREN
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127
      proper_explicit_type_properties -> State 138
      explicit_type_properties -> State 190

  State 130:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON.type_expr RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 191

  State 131:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 192

  State 132:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op.returnable_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [binary_type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66

  State 133:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [argument]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 134:
    Kernel Items:
      colon_expr: colon_expr COLON., *
      colon_expr: colon_expr COLON.expr
    Reduce:
      * -> [colon_expr]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [colon_expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 135:
    Kernel Items:
      colon_expr: expr COLON., *
      colon_expr: expr COLON.expr
    Reduce:
      * -> [colon_expr]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [colon_expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 136:
    Kernel Items:
      proper_arguments: proper_arguments COMMA.argument
      arguments: proper_arguments COMMA., *
    Reduce:
      * -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ELLIPSIS -> [argument]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 69
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      COLON -> State 68
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      colon_expr -> State 71
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 72
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

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
      NEWLINES -> State 194
      COMMA -> State 193

  State 139:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 195

  State 140:
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
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 196

  State 141:
    Kernel Items:
      unsafe_stmt: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      GREATER -> State 197

  State 142:
    Kernel Items:
      generic_arguments: DOLLAR_LBRACKET generic_argument_list.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [generic_arguments]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      proper_generic_argument_list: proper_generic_argument_list.COMMA type_expr
      generic_argument_list: proper_generic_argument_list., *
      generic_argument_list: proper_generic_argument_list.COMMA
    Reduce:
      * -> [generic_argument_list]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 198

  State 144:
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
      binary_type_op -> State 132

  State 145:
    Kernel Items:
      as_expr: accessible_expr DOT AS.LPAREN type_expr RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      LPAREN -> State 199

  State 146:
    Kernel Items:
      index: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [index]
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 134

  State 147:
    Kernel Items:
      index: expr., *
      colon_expr: expr.COLON
      colon_expr: expr.COLON expr
    Reduce:
      * -> [index]
    ShiftAndReduce:
      (nil)
    Goto:
      COLON -> State 135

  State 148:
    Kernel Items:
      index_expr: accessible_expr LBRACKET index.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [index_expr]
    Goto:
      (nil)

  State 149:
    Kernel Items:
      call_expr: accessible_expr generic_arguments LPAREN.arguments RPAREN
    Reduce:
      * -> [arguments]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ELLIPSIS -> [argument]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      argument -> [proper_arguments]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 69
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      COLON -> State 68
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      proper_arguments -> State 73
      arguments -> State 200
      colon_expr -> State 71
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 72
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 150:
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
      mul_op -> State 90

  State 151:
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

  State 152:
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
      add_op -> State 82

  State 153:
    Kernel Items:
      if_elif_expr: if_elif_expr ELSE IF.condition statements
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [condition]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      case_pattern_expr -> [condition]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 55
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      condition -> State 201
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 154:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [initialize_expr]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      jump_stmt: jump_op AT IDENTIFIER., *
      jump_stmt: jump_op AT IDENTIFIER.returnable_expr
    Reduce:
      * -> [jump_stmt]
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      returnable_expr -> [jump_stmt]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 156:
    Kernel Items:
      addr_decl_pattern: new_var_type new_addressable type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [addr_decl_pattern]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 157:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      AND -> State 83

  State 158:
    Kernel Items:
      binary_or_expr: or_expr.OR and_expr
      send_expr: send_recv_expr ARROW or_expr., *
    Reduce:
      * -> [send_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      OR -> State 92

  State 159:
    Kernel Items:
      send_expr: send_recv_expr.ARROW or_expr
      binary_assign_op_expr: send_recv_expr binary_assign_op send_recv_expr., *
    Reduce:
      * -> [binary_assign_op_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      ARROW -> State 97

  State 160:
    Kernel Items:
      loop_expr_body: FOR optional_statement SEMICOLON optional_expr.SEMICOLON optional_statement for_loop_body
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      SEMICOLON -> State 202

  State 161:
    Kernel Items:
      loop_expr_body: FOR returnable_expr IN expr.for_loop_body
    Reduce:
      (nil)
    ShiftAndReduce:
      for_loop_body -> [loop_expr_body]
    Goto:
      DO -> State 103

  State 162:
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
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 203

  State 163:
    Kernel Items:
      generic_parameters: DOLLAR_LBRACKET generic_parameter_list.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [generic_parameters]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      proper_generic_parameter_list: proper_generic_parameter_list.COMMA generic_parameter
      generic_parameter_list: proper_generic_parameter_list., *
      generic_parameter_list: proper_generic_parameter_list.COMMA
    Reduce:
      * -> [generic_parameter_list]
    ShiftAndReduce:
      (nil)
    Goto:
      COMMA -> State 204

  State 165:
    Kernel Items:
      func_signature: FUNC IDENTIFIER generic_parameters parameters.return_type
    Reduce:
      * -> [return_type]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [return_type]
      prefix_unary_type_expr -> [returnable_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      return_type -> [func_signature]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66

  State 166:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 167:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: GREATER type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 168:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 182

  State 169:
    Kernel Items:
      parameter: IDENTIFIER ELLIPSIS.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 205

  State 170:
    Kernel Items:
      parameter: IDENTIFIER GREATER.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 206

  State 171:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: IDENTIFIER type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 172:
    Kernel Items:
      parameter: UNDERSCORE ELLIPSIS.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 207

  State 173:
    Kernel Items:
      parameter: UNDERSCORE GREATER.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 208

  State 174:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: UNDERSCORE type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 175:
    Kernel Items:
      proper_parameter_list: proper_parameter_list COMMA.parameter
      parameter_list: proper_parameter_list COMMA., *
    Reduce:
      * -> [parameter_list]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      parameter -> [proper_parameter_list]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 110
      UNDERSCORE -> State 111
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      ELLIPSIS -> State 108
      GREATER -> State 109
      prefix_unary_type_op -> State 66
      type_expr -> State 114

  State 176:
    Kernel Items:
      case_pattern_expr: CASE switchable_case_patterns ASSIGN.expr
    Reduce:
      (nil)
    ShiftAndReduce:
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      expr -> [case_pattern_expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      if_elif_expr -> State 31
      repeat_loop_body -> State 40
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 177:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., *
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 58
      UNDERSCORE -> State 60
      DOT -> State 57

  State 178:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., *
    Reduce:
      * -> [import_clauses]
    ShiftAndReduce:
      STRING_LITERAL -> [import_clause]
      import_clause -> [proper_import_clauses]
    Goto:
      IDENTIFIER -> State 58
      UNDERSCORE -> State 60
      DOT -> State 57

  State 179:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 180:
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
      NEWLINES -> State 209
      OR -> State 210

  State 181:
    Kernel Items:
      proper_explicit_enum_type_properties: type_property.OR type_property
      proper_explicit_enum_type_properties: type_property.NEWLINES type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      NEWLINES -> State 211
      OR -> State 212

  State 182:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_arguments
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      generic_arguments -> [named_type_expr]
    Goto:
      DOLLAR_LBRACKET -> State 78

  State 183:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_arguments
      type_property: DEFAULT IDENTIFIER.type_expr
    Reduce:
      * -> [generic_arguments]
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      generic_arguments -> [named_type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      DOT -> State 168
      DOLLAR_LBRACKET -> State 78
      prefix_unary_type_op -> State 66
      type_expr -> State 213

  State 184:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_property: DEFAULT type_expr., *
    Reduce:
      * -> [type_property]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 185:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_property: IDENTIFIER type_expr., *
    Reduce:
      * -> [type_property]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 186:
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
      binary_type_op -> State 132

  State 187:
    Kernel Items:
      proper_implicit_enum_type_properties: proper_implicit_enum_type_properties OR.type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_implicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 188:
    Kernel Items:
      proper_implicit_type_properties: proper_implicit_type_properties COMMA.type_property
      implicit_type_properties: proper_implicit_type_properties COMMA., *
    Reduce:
      * -> [implicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_implicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 189:
    Kernel Items:
      proper_implicit_enum_type_properties: type_property OR.type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_implicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 190:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_type_properties.RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [trait_type_expr]
    Goto:
      (nil)

  State 191:
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
      binary_type_op -> State 132

  State 192:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    ShiftAndReduce:
      RBRACKET -> [array_type_expr]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      proper_explicit_type_properties: proper_explicit_type_properties COMMA.type_property
      explicit_type_properties: proper_explicit_type_properties COMMA., *
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 194:
    Kernel Items:
      proper_explicit_type_properties: proper_explicit_type_properties NEWLINES.type_property
      explicit_type_properties: proper_explicit_type_properties NEWLINES., *
    Reduce:
      * -> [explicit_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 195:
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
      binary_type_op -> State 132

  State 196:
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
      IMPLEMENTS -> State 214
      binary_type_op -> State 132

  State 197:
    Kernel Items:
      unsafe_stmt: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    ShiftAndReduce:
      STRING_LITERAL -> [unsafe_stmt]
    Goto:
      (nil)

  State 198:
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
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 215

  State 199:
    Kernel Items:
      as_expr: accessible_expr DOT AS LPAREN.type_expr RPAREN
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 216

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
      if_elif_expr: if_elif_expr ELSE IF condition.statements
    Reduce:
      (nil)
    ShiftAndReduce:
      statements -> [if_elif_expr]
    Goto:
      LBRACE -> State 16

  State 202:
    Kernel Items:
      loop_expr_body: FOR optional_statement SEMICOLON optional_expr SEMICOLON.optional_statement for_loop_body
    Reduce:
      * -> [optional_statement]
    ShiftAndReduce:
      COMMENT_GROUPS -> [floating_comment]
      INTEGER_LITERAL -> [literal_expr]
      FLOAT_LITERAL -> [literal_expr]
      RUNE_LITERAL -> [literal_expr]
      STRING_LITERAL -> [literal_expr]
      UNDERSCORE -> [named_expr]
      TRUE -> [literal_expr]
      FALSE -> [literal_expr]
      RETURN -> [jump_op]
      BREAK -> [jump_op]
      CONTINUE -> [jump_op]
      FALLTHROUGH -> [jump_stmt]
      ASYNC -> [prefix_unary_op]
      DEFER -> [prefix_unary_op]
      VAR -> [new_var_type]
      LET -> [new_var_type]
      NOT -> [prefix_unary_op]
      ADD -> [prefix_unary_op]
      SUB -> [prefix_unary_op]
      MUL -> [prefix_unary_op]
      BIT_AND -> [prefix_unary_op]
      BIT_XOR -> [prefix_unary_op]
      PARSE_ERROR -> [parse_error_expr]
      statement -> [optional_statement]
      floating_comment -> [statement]
      branch_stmt -> [statement]
      unsafe_stmt -> [statement]
      jump_stmt -> [statement]
      assign_stmt -> [statement]
      import_stmt -> [statement]
      addr_decl_pattern -> [expr]
      assign_to_addr_pattern -> [expr]
      atom_expr -> [accessible_expr]
      parse_error_expr -> [atom_expr]
      literal_expr -> [atom_expr]
      named_expr -> [atom_expr]
      initialize_expr -> [atom_expr]
      implicit_struct_expr -> [atom_expr]
      access_expr -> [accessible_expr]
      index_expr -> [accessible_expr]
      as_expr -> [accessible_expr]
      call_expr -> [accessible_expr]
      postfixable_expr -> [prefixable_expr]
      postfix_unary_expr -> [postfixable_expr]
      prefixable_expr -> [mul_expr]
      prefix_unary_expr -> [prefixable_expr]
      binary_mul_expr -> [mul_expr]
      binary_add_expr -> [add_expr]
      binary_cmp_expr -> [cmp_expr]
      binary_and_expr -> [and_expr]
      binary_or_expr -> [or_expr]
      send_expr -> [send_recv_expr]
      recv_expr -> [send_recv_expr]
      assign_op_expr -> [expr]
      binary_assign_op_expr -> [assign_op_expr]
      unlabelled_control_flow_expr -> [control_flow_expr]
      control_flow_expr -> [expr]
      statements -> [unlabelled_control_flow_expr]
      if_else_expr -> [unlabelled_control_flow_expr]
      if_only_expr -> [if_elif_expr]
      switch_expr_body -> [unlabelled_control_flow_expr]
      select_expr_body -> [unlabelled_control_flow_expr]
      loop_expr_body -> [unlabelled_control_flow_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      type_def -> [statement]
      explicit_struct_type_expr -> [initializable_type_expr]
      func_def -> [atom_expr]
    Goto:
      IDENTIFIER -> State 13
      IF -> State 14
      SWITCH -> State 22
      CASE -> State 8
      DEFAULT -> State 9
      REPEAT -> State 19
      FOR -> State 10
      SELECT -> State 20
      IMPORT -> State 15
      UNSAFE -> State 24
      TYPE -> State 23
      STRUCT -> State 21
      FUNC -> State 11
      LBRACE -> State 16
      LPAREN -> State 18
      LBRACKET -> State 17
      ARROW -> State 7
      GREATER -> State 12
      jump_op -> State 34
      new_var_type -> State 36
      accessible_expr -> State 25
      prefix_unary_op -> State 38
      mul_expr -> State 35
      add_expr -> State 26
      cmp_expr -> State 28
      and_expr -> State 27
      or_expr -> State 37
      send_recv_expr -> State 42
      expr -> State 29
      if_elif_expr -> State 31
      optional_statement -> State 217
      repeat_loop_body -> State 40
      returnable_expr -> State 41
      improper_expr_struct -> State 32
      initializable_type_expr -> State 33
      func_signature -> State 30

  State 203:
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
      binary_type_op -> State 132

  State 204:
    Kernel Items:
      proper_generic_parameter_list: proper_generic_parameter_list COMMA.generic_parameter
      generic_parameter_list: proper_generic_parameter_list COMMA., *
    Reduce:
      * -> [generic_parameter_list]
    ShiftAndReduce:
      generic_parameter -> [proper_generic_parameter_list]
    Goto:
      IDENTIFIER -> State 162

  State 205:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 206:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: IDENTIFIER GREATER type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 207:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: UNDERSCORE ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 208:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter: UNDERSCORE GREATER type_expr., *
    Reduce:
      * -> [parameter]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 209:
    Kernel Items:
      proper_explicit_enum_type_properties: proper_explicit_enum_type_properties NEWLINES.type_property
      explicit_enum_type_properties: proper_explicit_enum_type_properties NEWLINES., *
    Reduce:
      * -> [explicit_enum_type_properties]
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 210:
    Kernel Items:
      proper_explicit_enum_type_properties: proper_explicit_enum_type_properties OR.type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 211:
    Kernel Items:
      proper_explicit_enum_type_properties: type_property NEWLINES.type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 212:
    Kernel Items:
      proper_explicit_enum_type_properties: type_property OR.type_property
    Reduce:
      (nil)
    ShiftAndReduce:
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      type_property -> [proper_explicit_enum_type_properties]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 121
      UNDERSCORE -> State 122
      DEFAULT -> State 120
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 127

  State 213:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_property: DEFAULT IDENTIFIER type_expr., *
    Reduce:
      * -> [type_property]
    ShiftAndReduce:
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 214:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameters type_expr IMPLEMENTS.type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      UNDERSCORE -> [inferred_type_expr]
      DOT -> [inferred_type_expr]
      QUESTION -> [prefix_unary_type_op]
      EXCLAIM -> [prefix_unary_type_op]
      TILDE -> [prefix_unary_type_op]
      TILDE_TILDE -> [prefix_unary_type_op]
      BIT_AND -> [prefix_unary_type_op]
      initializable_type_expr -> [atom_type_expr]
      slice_type_expr -> [initializable_type_expr]
      array_type_expr -> [initializable_type_expr]
      map_type_expr -> [initializable_type_expr]
      atom_type_expr -> [returnable_type_expr]
      named_type_expr -> [atom_type_expr]
      inferred_type_expr -> [atom_type_expr]
      returnable_type_expr -> [type_expr]
      prefix_unary_type_expr -> [returnable_type_expr]
      binary_type_expr -> [type_expr]
      implicit_struct_type_expr -> [atom_type_expr]
      explicit_struct_type_expr -> [initializable_type_expr]
      trait_type_expr -> [atom_type_expr]
      implicit_enum_type_expr -> [atom_type_expr]
      explicit_enum_type_expr -> [atom_type_expr]
      func_signature -> [atom_type_expr]
    Goto:
      IDENTIFIER -> State 63
      STRUCT -> State 21
      ENUM -> State 62
      TRAIT -> State 65
      FUNC -> State 11
      LPAREN -> State 64
      LBRACKET -> State 17
      prefix_unary_type_op -> State 66
      type_expr -> State 218

  State 215:
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
      binary_type_op -> State 132

  State 216:
    Kernel Items:
      as_expr: accessible_expr DOT AS LPAREN type_expr.RPAREN
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    ShiftAndReduce:
      RPAREN -> [as_expr]
      ADD -> [binary_type_op]
      SUB -> [binary_type_op]
      MUL -> [binary_type_op]
    Goto:
      binary_type_op -> State 132

  State 217:
    Kernel Items:
      loop_expr_body: FOR optional_statement SEMICOLON optional_expr SEMICOLON optional_statement.for_loop_body
    Reduce:
      (nil)
    ShiftAndReduce:
      for_loop_body -> [loop_expr_body]
    Goto:
      DO -> State 103

  State 218:
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
      binary_type_op -> State 132

Number of states: 218
Number of shift actions: 1721
Number of reduce actions: 111
Number of shift-and-reduce actions: 3585
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 8469
Number of unoptimized shift actions: 98088
Number of unoptimized reduce actions: 84792
*/
