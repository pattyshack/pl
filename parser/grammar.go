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
	IntegerLiteralToken  = SymbolId(257)
	FloatLiteralToken    = SymbolId(258)
	RuneLiteralToken     = SymbolId(259)
	StringLiteralToken   = SymbolId(260)
	IdentifierToken      = SymbolId(261)
	TrueToken            = SymbolId(262)
	FalseToken           = SymbolId(263)
	IfToken              = SymbolId(264)
	ElseToken            = SymbolId(265)
	SwitchToken          = SymbolId(266)
	CaseToken            = SymbolId(267)
	DefaultToken         = SymbolId(268)
	ForToken             = SymbolId(269)
	DoToken              = SymbolId(270)
	InToken              = SymbolId(271)
	ReturnToken          = SymbolId(272)
	BreakToken           = SymbolId(273)
	ContinueToken        = SymbolId(274)
	FallthroughToken     = SymbolId(275)
	PackageToken         = SymbolId(276)
	ImportToken          = SymbolId(277)
	AsToken              = SymbolId(278)
	UnsafeToken          = SymbolId(279)
	TypeToken            = SymbolId(280)
	ImplementsToken      = SymbolId(281)
	StructToken          = SymbolId(282)
	EnumToken            = SymbolId(283)
	TraitToken           = SymbolId(284)
	FuncToken            = SymbolId(285)
	AsyncToken           = SymbolId(286)
	DeferToken           = SymbolId(287)
	VarToken             = SymbolId(288)
	LetToken             = SymbolId(289)
	NotToken             = SymbolId(290)
	AndToken             = SymbolId(291)
	OrToken              = SymbolId(292)
	LabelDeclToken       = SymbolId(293)
	JumpLabelToken       = SymbolId(294)
	LbraceToken          = SymbolId(295)
	RbraceToken          = SymbolId(296)
	LparenToken          = SymbolId(297)
	RparenToken          = SymbolId(298)
	LbracketToken        = SymbolId(299)
	RbracketToken        = SymbolId(300)
	DotToken             = SymbolId(301)
	CommaToken           = SymbolId(302)
	QuestionToken        = SymbolId(303)
	SemicolonToken       = SymbolId(304)
	ColonToken           = SymbolId(305)
	ExclaimToken         = SymbolId(306)
	DollarLbracketToken  = SymbolId(307)
	DotDotDotToken       = SymbolId(308)
	TildeTildeToken      = SymbolId(309)
	AssignToken          = SymbolId(310)
	AddAssignToken       = SymbolId(311)
	SubAssignToken       = SymbolId(312)
	MulAssignToken       = SymbolId(313)
	DivAssignToken       = SymbolId(314)
	ModAssignToken       = SymbolId(315)
	AddOneAssignToken    = SymbolId(316)
	SubOneAssignToken    = SymbolId(317)
	BitNegAssignToken    = SymbolId(318)
	BitAndAssignToken    = SymbolId(319)
	BitOrAssignToken     = SymbolId(320)
	BitXorAssignToken    = SymbolId(321)
	BitLshiftAssignToken = SymbolId(322)
	BitRshiftAssignToken = SymbolId(323)
	AddToken             = SymbolId(324)
	SubToken             = SymbolId(325)
	MulToken             = SymbolId(326)
	DivToken             = SymbolId(327)
	ModToken             = SymbolId(328)
	BitNegToken          = SymbolId(329)
	BitAndToken          = SymbolId(330)
	BitXorToken          = SymbolId(331)
	BitOrToken           = SymbolId(332)
	BitLshiftToken       = SymbolId(333)
	BitRshiftToken       = SymbolId(334)
	EqualToken           = SymbolId(335)
	NotEqualToken        = SymbolId(336)
	LessToken            = SymbolId(337)
	LessOrEqualToken     = SymbolId(338)
	GreaterToken         = SymbolId(339)
	GreaterOrEqualToken  = SymbolId(340)
	ParseErrorToken      = SymbolId(341)
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

type Reducer interface {
	// 50:10: source -> ...
	ToSource(OptionalDefinitions_ GenericSymbol) (GenericSymbol, error)

	// 53:2: optional_definitions -> NEWLINES: ...
	NewlinesToOptionalDefinitions(Newlines_ CountSymbol) (GenericSymbol, error)

	// 54:2: optional_definitions -> definitions: ...
	DefinitionsToOptionalDefinitions(OptionalNewlines_ GenericSymbol, Definitions_ GenericSymbol, OptionalNewlines_2 GenericSymbol) (GenericSymbol, error)

	// 55:2: optional_definitions -> nil: ...
	NilToOptionalDefinitions() (GenericSymbol, error)

	// 58:2: optional_newlines -> NEWLINES: ...
	NewlinesToOptionalNewlines(Newlines_ CountSymbol) (GenericSymbol, error)

	// 59:2: optional_newlines -> nil: ...
	NilToOptionalNewlines() (GenericSymbol, error)

	// 62:2: definitions -> nil: ...
	NilToDefinitions(Definition_ GenericSymbol) (GenericSymbol, error)

	// 63:2: definitions -> add: ...
	AddToDefinitions(Definitions_ GenericSymbol, Newlines_ CountSymbol, Definition_ GenericSymbol) (GenericSymbol, error)

	// 67:2: definition -> package_def: ...
	PackageDefToDefinition(PackageDef_ GenericSymbol) (GenericSymbol, error)

	// 68:2: definition -> type_def: ...
	TypeDefToDefinition(TypeDef_ GenericSymbol) (GenericSymbol, error)

	// 69:2: definition -> named_func_def: ...
	NamedFuncDefToDefinition(NamedFuncDef_ GenericSymbol) (GenericSymbol, error)

	// 70:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ GenericSymbol) (GenericSymbol, error)

	// 71:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ GenericSymbol, Assign_ GenericSymbol, Expression_ GenericSymbol) (GenericSymbol, error)

	// 74:2: definition -> statement_block: ...
	StatementBlockToDefinition(StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 94:19: statement_block -> ...
	ToStatementBlock(Lbrace_ GenericSymbol, Statements_ GenericSymbol, Rbrace_ GenericSymbol) (GenericSymbol, error)

	// 97:2: statements -> empty_list: ...
	EmptyListToStatements() (GenericSymbol, error)

	// 98:2: statements -> add: ...
	AddToStatements(Statements_ GenericSymbol, Statement_ GenericSymbol) (GenericSymbol, error)

	// 101:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ GenericSymbol, Newlines_ CountSymbol) (GenericSymbol, error)

	// 102:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ GenericSymbol, Semicolon_ GenericSymbol) (GenericSymbol, error)

	// 105:2: simple_statement_body -> unsafe_statement: ...
	UnsafeStatementToSimpleStatementBody(UnsafeStatement_ GenericSymbol) (GenericSymbol, error)

	// 107:2: simple_statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToSimpleStatementBody(Expressions_ GenericSymbol) (GenericSymbol, error)

	// 115:2: simple_statement_body -> async: ...
	AsyncToSimpleStatementBody(Async_ GenericSymbol, CallExpr_ GenericSymbol) (GenericSymbol, error)

	// 133:2: simple_statement_body -> defer: ...
	DeferToSimpleStatementBody(Defer_ GenericSymbol, CallExpr_ GenericSymbol) (GenericSymbol, error)

	// 137:2: simple_statement_body -> jump_statement: ...
	JumpStatementToSimpleStatementBody(JumpType_ GenericSymbol, OptionalJumpLabel_ GenericSymbol, OptionalExpressions_ GenericSymbol) (GenericSymbol, error)

	// 142:2: simple_statement_body -> assign_statement: ...
	AssignStatementToSimpleStatementBody(AssignPattern_ GenericSymbol, Assign_ GenericSymbol, Expression_ GenericSymbol) (GenericSymbol, error)

	// 148:2: simple_statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToSimpleStatementBody(AccessExpr_ GenericSymbol, UnaryOpAssign_ GenericSymbol) (GenericSymbol, error)

	// 149:2: simple_statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToSimpleStatementBody(AccessExpr_ GenericSymbol, BinaryOpAssign_ GenericSymbol, Expression_ GenericSymbol) (GenericSymbol, error)

	// 153:2: simple_statement_body -> FALLTHROUGH: ...
	FallthroughToSimpleStatementBody(Fallthrough_ GenericSymbol) (GenericSymbol, error)

	// 156:2: statement_body -> simple_statement_body: ...
	SimpleStatementBodyToStatementBody(SimpleStatementBody_ GenericSymbol) (GenericSymbol, error)

	// 159:2: statement_body -> import_statement: ...
	ImportStatementToStatementBody(ImportStatement_ GenericSymbol) (GenericSymbol, error)

	// 164:2: statement_body -> case_branch_statement: ...
	CaseBranchStatementToStatementBody(Case_ GenericSymbol, CasePatterns_ GenericSymbol, Colon_ GenericSymbol, OptionalSimpleStatementBody_ GenericSymbol) (GenericSymbol, error)

	// 165:2: statement_body -> default_branch_statement: ...
	DefaultBranchStatementToStatementBody(Default_ GenericSymbol, Colon_ GenericSymbol, OptionalSimpleStatementBody_ GenericSymbol) (GenericSymbol, error)

	// 168:2: optional_simple_statement_body -> simple_statement_body: ...
	SimpleStatementBodyToOptionalSimpleStatementBody(SimpleStatementBody_ GenericSymbol) (GenericSymbol, error)

	// 169:2: optional_simple_statement_body -> nil: ...
	NilToOptionalSimpleStatementBody() (GenericSymbol, error)

	// 176:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ GenericSymbol) (GenericSymbol, error)

	// 177:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ GenericSymbol) (GenericSymbol, error)

	// 180:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ GenericSymbol) (GenericSymbol, error)

	// 181:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ GenericSymbol) (GenericSymbol, error)

	// 182:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ GenericSymbol) (GenericSymbol, error)

	// 183:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ GenericSymbol) (GenericSymbol, error)

	// 184:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ GenericSymbol) (GenericSymbol, error)

	// 185:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ GenericSymbol) (GenericSymbol, error)

	// 186:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ GenericSymbol) (GenericSymbol, error)

	// 187:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ GenericSymbol) (GenericSymbol, error)

	// 188:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ GenericSymbol) (GenericSymbol, error)

	// 189:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ GenericSymbol) (GenericSymbol, error)

	// 190:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ GenericSymbol) (GenericSymbol, error)

	// 198:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ GenericSymbol, Less_ GenericSymbol, Identifier_ ValueSymbol, Greater_ GenericSymbol, StringLiteral_ GenericSymbol) (GenericSymbol, error)

	// 205:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ GenericSymbol) (GenericSymbol, error)

	// 206:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ GenericSymbol) (GenericSymbol, error)

	// 207:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ GenericSymbol) (GenericSymbol, error)

	// 210:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ ValueSymbol) (GenericSymbol, error)

	// 211:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (GenericSymbol, error)

	// 214:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ GenericSymbol) (GenericSymbol, error)

	// 215:2: expressions -> add: ...
	AddToExpressions(Expressions_ GenericSymbol, Comma_ GenericSymbol, Expression_ GenericSymbol) (GenericSymbol, error)

	// 218:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ GenericSymbol) (GenericSymbol, error)

	// 219:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (GenericSymbol, error)

	// 226:2: import_statement -> single: ...
	SingleToImportStatement(Import_ GenericSymbol, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 227:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ GenericSymbol, Lparen_ GenericSymbol, ImportClauses_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 230:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ GenericSymbol) (GenericSymbol, error)

	// 231:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ GenericSymbol, As_ GenericSymbol, Identifier_ ValueSymbol) (GenericSymbol, error)

	// 234:2: import_clause_terminal -> implicit: ...
	ImplicitToImportClauseTerminal(ImportClause_ GenericSymbol, Newlines_ CountSymbol) (GenericSymbol, error)

	// 235:2: import_clause_terminal -> explicit: ...
	ExplicitToImportClauseTerminal(ImportClause_ GenericSymbol, Comma_ GenericSymbol) (GenericSymbol, error)

	// 238:2: import_clauses -> first: ...
	FirstToImportClauses(ImportClauseTerminal_ GenericSymbol) (GenericSymbol, error)

	// 239:2: import_clauses -> add: ...
	AddToImportClauses(ImportClauses_ GenericSymbol, ImportClauseTerminal_ GenericSymbol) (GenericSymbol, error)

	// 246:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 247:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ GenericSymbol, Comma_ GenericSymbol, CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 256:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ GenericSymbol, VarPattern_ GenericSymbol, OptionalValueType_ GenericSymbol) (GenericSymbol, error)

	// 258:14: var_or_let -> VAR: ...
	VarToVarOrLet(Var_ GenericSymbol) (GenericSymbol, error)

	// 258:20: var_or_let -> LET: ...
	LetToVarOrLet(Let_ GenericSymbol) (GenericSymbol, error)

	// 261:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ ValueSymbol) (GenericSymbol, error)

	// 262:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ GenericSymbol) (GenericSymbol, error)

	// 264:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ GenericSymbol, FieldVarPatterns_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 267:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 268:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ GenericSymbol, Comma_ GenericSymbol, FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 271:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 272:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ ValueSymbol, Assign_ GenericSymbol, VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 273:2: field_var_pattern -> DOT_DOT_DOT: ...
	DotDotDotToFieldVarPattern(DotDotDot_ GenericSymbol) (GenericSymbol, error)

	// 276:2: optional_value_type -> value_type: ...
	ValueTypeToOptionalValueType(ValueType_ GenericSymbol) (GenericSymbol, error)

	// 277:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (GenericSymbol, error)

	// 283:18: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 286:2: case_pattern -> assign_pattern: ...
	AssignPatternToCasePattern(AssignPattern_ GenericSymbol) (GenericSymbol, error)

	// 293:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ GenericSymbol, Identifier_ ValueSymbol, ImplicitStructExpr_ GenericSymbol) (GenericSymbol, error)

	// 294:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ GenericSymbol, Dot_ GenericSymbol, Identifier_ ValueSymbol, TuplePattern_ GenericSymbol) (GenericSymbol, error)

	// 301:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ GenericSymbol, IfExpr_ GenericSymbol) (GenericSymbol, error)

	// 302:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ GenericSymbol, SwitchExpr_ GenericSymbol) (GenericSymbol, error)

	// 303:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ GenericSymbol, LoopExpr_ GenericSymbol) (GenericSymbol, error)

	// 304:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 307:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ ValueSymbol) (GenericSymbol, error)

	// 308:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (GenericSymbol, error)

	// 315:2: sequence_expr -> or_expr: ...
	OrExprToSequenceExpr(OrExpr_ GenericSymbol) (GenericSymbol, error)

	// 318:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ GenericSymbol) (GenericSymbol, error)

	// 322:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ GenericSymbol, SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 331:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ GenericSymbol, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 332:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ GenericSymbol, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ GenericSymbol, StatementBlock_2 GenericSymbol) (GenericSymbol, error)

	// 333:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ GenericSymbol, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ GenericSymbol, IfExpr_ GenericSymbol) (GenericSymbol, error)

	// 336:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 337:2: condition -> case: ...
	CaseToCondition(Case_ GenericSymbol, CasePatterns_ GenericSymbol, Assign_ GenericSymbol, SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 361:15: switch_expr -> ...
	ToSwitchExpr(Switch_ GenericSymbol, SequenceExpr_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 375:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 376:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ GenericSymbol, StatementBlock_ GenericSymbol, For_ GenericSymbol, SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 377:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ GenericSymbol, SequenceExpr_ GenericSymbol, Do_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 378:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ GenericSymbol, AssignPattern_ GenericSymbol, In_ GenericSymbol, SequenceExpr_ GenericSymbol, Do_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 379:2: loop_expr -> for: ...
	ForToLoopExpr(For_ GenericSymbol, ForAssignment_ GenericSymbol, Semicolon_ GenericSymbol, OptionalSequenceExpr_ GenericSymbol, Semicolon_2 GenericSymbol, OptionalSequenceExpr_2 GenericSymbol, Do_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 382:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 383:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (GenericSymbol, error)

	// 386:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 387:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ GenericSymbol, SequenceExpr_ GenericSymbol) (GenericSymbol, error)

	// 393:13: call_expr -> ...
	ToCallExpr(AccessExpr_ GenericSymbol, OptionalGenericBinding_ GenericSymbol, Lparen_ GenericSymbol, OptionalArguments_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 396:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ GenericSymbol, OptionalGenericArguments_ GenericSymbol, Rbracket_ GenericSymbol) (GenericSymbol, error)

	// 397:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (GenericSymbol, error)

	// 400:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ GenericSymbol) (GenericSymbol, error)

	// 401:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (GenericSymbol, error)

	// 405:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ GenericSymbol) (GenericSymbol, error)

	// 406:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ GenericSymbol, Comma_ GenericSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 409:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ GenericSymbol) (GenericSymbol, error)

	// 410:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (GenericSymbol, error)

	// 413:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ GenericSymbol) (GenericSymbol, error)

	// 414:2: arguments -> add: ...
	AddToArguments(Arguments_ GenericSymbol, Comma_ GenericSymbol, Argument_ GenericSymbol) (GenericSymbol, error)

	// 417:2: argument -> positional: ...
	PositionalToArgument(Expression_ GenericSymbol) (GenericSymbol, error)

	// 418:2: argument -> named: ...
	NamedToArgument(Identifier_ ValueSymbol, Assign_ GenericSymbol, Expression_ GenericSymbol) (GenericSymbol, error)

	// 419:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ GenericSymbol) (GenericSymbol, error)

	// 422:2: argument -> DOT_DOT_DOT: ...
	DotDotDotToArgument(DotDotDot_ GenericSymbol) (GenericSymbol, error)

	// 425:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ GenericSymbol, Colon_ GenericSymbol, OptionalExpression_2 GenericSymbol) (GenericSymbol, error)

	// 426:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ GenericSymbol, Colon_ GenericSymbol, OptionalExpression_ GenericSymbol) (GenericSymbol, error)

	// 429:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ GenericSymbol) (GenericSymbol, error)

	// 430:2: optional_expression -> nil: ...
	NilToOptionalExpression() (GenericSymbol, error)

	// 441:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ GenericSymbol) (GenericSymbol, error)

	// 442:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ ValueSymbol) (GenericSymbol, error)

	// 443:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(OptionalLabelDecl_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 444:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ GenericSymbol) (GenericSymbol, error)

	// 445:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializableType_ GenericSymbol, Lparen_ GenericSymbol, Arguments_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 446:2: atom_expr -> implicit_struct_expr: ...
	ImplicitStructExprToAtomExpr(ImplicitStructExpr_ GenericSymbol) (GenericSymbol, error)

	// 447:2: atom_expr -> PARSE_ERROR: ...
	ParseErrorToAtomExpr(ParseError_ ParseErrorSymbol) (GenericSymbol, error)

	// 450:2: literal -> TRUE: ...
	TrueToLiteral(True_ GenericSymbol) (GenericSymbol, error)

	// 451:2: literal -> FALSE: ...
	FalseToLiteral(False_ GenericSymbol) (GenericSymbol, error)

	// 452:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ GenericSymbol) (GenericSymbol, error)

	// 453:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ GenericSymbol) (GenericSymbol, error)

	// 454:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ GenericSymbol) (GenericSymbol, error)

	// 455:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ GenericSymbol) (GenericSymbol, error)

	// 457:24: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ GenericSymbol, Arguments_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 460:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ GenericSymbol) (GenericSymbol, error)

	// 461:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ GenericSymbol, Dot_ GenericSymbol, Identifier_ ValueSymbol) (GenericSymbol, error)

	// 462:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ GenericSymbol) (GenericSymbol, error)

	// 463:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ GenericSymbol, Lbracket_ GenericSymbol, Argument_ GenericSymbol, Rbracket_ GenericSymbol) (GenericSymbol, error)

	// 466:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ GenericSymbol) (GenericSymbol, error)

	// 467:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ GenericSymbol, Question_ GenericSymbol) (GenericSymbol, error)

	// 470:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ GenericSymbol) (GenericSymbol, error)

	// 471:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ GenericSymbol) (GenericSymbol, error)

	// 472:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ GenericSymbol) (GenericSymbol, error)

	// 475:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ GenericSymbol) (GenericSymbol, error)

	// 478:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ GenericSymbol) (GenericSymbol, error)

	// 481:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ GenericSymbol) (GenericSymbol, error)

	// 482:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ GenericSymbol, PrefixUnaryExpr_ GenericSymbol) (GenericSymbol, error)

	// 485:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ GenericSymbol) (GenericSymbol, error)

	// 486:2: mul_op -> DIV: ...
	DivToMulOp(Div_ GenericSymbol) (GenericSymbol, error)

	// 487:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ GenericSymbol) (GenericSymbol, error)

	// 488:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ GenericSymbol) (GenericSymbol, error)

	// 489:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ GenericSymbol) (GenericSymbol, error)

	// 490:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ GenericSymbol) (GenericSymbol, error)

	// 493:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ GenericSymbol) (GenericSymbol, error)

	// 494:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ GenericSymbol, MulOp_ GenericSymbol, PrefixUnaryExpr_ GenericSymbol) (GenericSymbol, error)

	// 497:2: add_op -> ADD: ...
	AddToAddOp(Add_ GenericSymbol) (GenericSymbol, error)

	// 498:2: add_op -> SUB: ...
	SubToAddOp(Sub_ GenericSymbol) (GenericSymbol, error)

	// 499:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ GenericSymbol) (GenericSymbol, error)

	// 500:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ GenericSymbol) (GenericSymbol, error)

	// 503:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ GenericSymbol) (GenericSymbol, error)

	// 504:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ GenericSymbol, AddOp_ GenericSymbol, MulExpr_ GenericSymbol) (GenericSymbol, error)

	// 507:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ GenericSymbol) (GenericSymbol, error)

	// 508:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ GenericSymbol) (GenericSymbol, error)

	// 509:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ GenericSymbol) (GenericSymbol, error)

	// 510:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ GenericSymbol) (GenericSymbol, error)

	// 511:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ GenericSymbol) (GenericSymbol, error)

	// 512:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ GenericSymbol) (GenericSymbol, error)

	// 515:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ GenericSymbol) (GenericSymbol, error)

	// 516:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ GenericSymbol, CmpOp_ GenericSymbol, AddExpr_ GenericSymbol) (GenericSymbol, error)

	// 519:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ GenericSymbol) (GenericSymbol, error)

	// 520:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ GenericSymbol, And_ GenericSymbol, CmpExpr_ GenericSymbol) (GenericSymbol, error)

	// 523:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ GenericSymbol) (GenericSymbol, error)

	// 524:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ GenericSymbol, Or_ GenericSymbol, AndExpr_ GenericSymbol) (GenericSymbol, error)

	// 533:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ GenericSymbol) (GenericSymbol, error)

	// 534:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ GenericSymbol, ValueType_ GenericSymbol, Rbracket_ GenericSymbol) (GenericSymbol, error)

	// 535:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ GenericSymbol, ValueType_ GenericSymbol, Comma_ GenericSymbol, IntegerLiteral_ GenericSymbol, Rbracket_ GenericSymbol) (GenericSymbol, error)

	// 536:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ GenericSymbol, ValueType_ GenericSymbol, Colon_ GenericSymbol, ValueType_2 GenericSymbol, Rbracket_ GenericSymbol) (GenericSymbol, error)

	// 539:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ GenericSymbol) (GenericSymbol, error)

	// 540:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ ValueSymbol, OptionalGenericBinding_ GenericSymbol) (GenericSymbol, error)

	// 541:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ ValueSymbol, Dot_ GenericSymbol, Identifier_2 ValueSymbol, OptionalGenericBinding_ GenericSymbol) (GenericSymbol, error)

	// 542:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ GenericSymbol, OptionalGenericBinding_ GenericSymbol) (GenericSymbol, error)

	// 543:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ GenericSymbol) (GenericSymbol, error)

	// 544:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ GenericSymbol) (GenericSymbol, error)

	// 545:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ GenericSymbol) (GenericSymbol, error)

	// 546:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ GenericSymbol) (GenericSymbol, error)

	// 547:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ GenericSymbol) (GenericSymbol, error)

	// 548:2: atom_type -> PARSE_ERROR: ...
	ParseErrorToAtomType(ParseError_ ParseErrorSymbol) (GenericSymbol, error)

	// 554:2: returnable_type -> atom_type: ...
	AtomTypeToReturnableType(AtomType_ GenericSymbol) (GenericSymbol, error)

	// 555:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 556:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 557:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 558:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 559:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 564:2: value_type -> returnable_type: ...
	ReturnableTypeToValueType(ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 565:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ GenericSymbol, Mul_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 566:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ GenericSymbol, Add_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 567:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ GenericSymbol, Sub_ GenericSymbol, ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 570:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ GenericSymbol, Identifier_ ValueSymbol, OptionalGenericParameters_ GenericSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 571:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ GenericSymbol, Identifier_ ValueSymbol, OptionalGenericParameters_ GenericSymbol, ValueType_ GenericSymbol, Implements_ GenericSymbol, ValueType_2 GenericSymbol) (GenericSymbol, error)

	// 572:2: type_def -> alias: ...
	AliasToTypeDef(Type_ GenericSymbol, Identifier_ ValueSymbol, Assign_ GenericSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 580:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ ValueSymbol) (GenericSymbol, error)

	// 581:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ ValueSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 584:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 585:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ GenericSymbol, Comma_ GenericSymbol, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 588:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ GenericSymbol) (GenericSymbol, error)

	// 589:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (GenericSymbol, error)

	// 592:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ GenericSymbol, OptionalGenericParameterDefs_ GenericSymbol, Rbracket_ GenericSymbol) (GenericSymbol, error)

	// 593:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)

	// 600:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ ValueSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 601:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ GenericSymbol) (GenericSymbol, error)

	// 602:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ GenericSymbol) (GenericSymbol, error)

	// 605:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 606:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ GenericSymbol, Comma_ GenericSymbol, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 609:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ GenericSymbol) (GenericSymbol, error)

	// 610:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (GenericSymbol, error)

	// 612:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ GenericSymbol, OptionalImplicitFieldDefs_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 615:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 616:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol, Newlines_ CountSymbol, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 617:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol, Comma_ GenericSymbol, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 620:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol) (GenericSymbol, error)

	// 621:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (GenericSymbol, error)

	// 623:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ GenericSymbol, Lparen_ GenericSymbol, OptionalExplicitFieldDefs_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 631:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 632:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ GenericSymbol, Assign_ GenericSymbol, Default_ GenericSymbol) (GenericSymbol, error)

	// 644:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ GenericSymbol, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 645:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ GenericSymbol, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 647:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ GenericSymbol, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 650:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ GenericSymbol, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 651:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Newlines_ CountSymbol, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 652:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ GenericSymbol, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 653:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Newlines_ CountSymbol, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 655:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ GenericSymbol, Lparen_ GenericSymbol, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 662:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 663:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ GenericSymbol) (GenericSymbol, error)

	// 666:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 667:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ GenericSymbol, Newlines_ CountSymbol, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 668:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ GenericSymbol, Comma_ GenericSymbol, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 671:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ GenericSymbol) (GenericSymbol, error)

	// 672:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (GenericSymbol, error)

	// 674:13: trait_def -> ...
	ToTraitDef(Trait_ GenericSymbol, Lparen_ GenericSymbol, OptionalTraitProperties_ GenericSymbol, Rparen_ GenericSymbol) (GenericSymbol, error)

	// 682:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ GenericSymbol) (GenericSymbol, error)

	// 683:2: return_type -> nil: ...
	NilToReturnType() (GenericSymbol, error)

	// 686:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ ValueSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 687:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ ValueSymbol, DotDotDot_ GenericSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 688:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ GenericSymbol) (GenericSymbol, error)

	// 689:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(DotDotDot_ GenericSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 692:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 693:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ GenericSymbol, Comma_ GenericSymbol, ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 696:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ GenericSymbol) (GenericSymbol, error)

	// 697:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (GenericSymbol, error)

	// 699:13: func_type -> ...
	ToFuncType(Func_ GenericSymbol, Lparen_ GenericSymbol, OptionalParameterDecls_ GenericSymbol, Rparen_ GenericSymbol, ReturnType_ GenericSymbol) (GenericSymbol, error)

	// 710:20: method_signature -> ...
	ToMethodSignature(Func_ GenericSymbol, Identifier_ ValueSymbol, Lparen_ GenericSymbol, OptionalParameterDecls_ GenericSymbol, Rparen_ GenericSymbol, ReturnType_ GenericSymbol) (GenericSymbol, error)

	// 716:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ ValueSymbol) (GenericSymbol, error)

	// 717:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ ValueSymbol, DotDotDot_ GenericSymbol) (GenericSymbol, error)

	// 718:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ ValueSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 719:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ ValueSymbol, DotDotDot_ GenericSymbol, ValueType_ GenericSymbol) (GenericSymbol, error)

	// 722:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 723:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ GenericSymbol, Comma_ GenericSymbol, ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 726:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ GenericSymbol) (GenericSymbol, error)

	// 727:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (GenericSymbol, error)

	// 730:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ GenericSymbol, Identifier_ ValueSymbol, OptionalGenericParameters_ GenericSymbol, Lparen_ GenericSymbol, OptionalParameterDefs_ GenericSymbol, Rparen_ GenericSymbol, ReturnType_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 731:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ GenericSymbol, Lparen_ GenericSymbol, ParameterDef_ GenericSymbol, Rparen_ GenericSymbol, Identifier_ ValueSymbol, OptionalGenericParameters_ GenericSymbol, Lparen_2 GenericSymbol, OptionalParameterDefs_ GenericSymbol, Rparen_2 GenericSymbol, ReturnType_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 732:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ GenericSymbol, Identifier_ ValueSymbol, Assign_ GenericSymbol, Expression_ GenericSymbol) (GenericSymbol, error)

	// 736:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ GenericSymbol, Lparen_ GenericSymbol, OptionalParameterDefs_ GenericSymbol, Rparen_ GenericSymbol, ReturnType_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)

	// 748:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ GenericSymbol) (GenericSymbol, error)

	// 749:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ GenericSymbol, StatementBlock_ GenericSymbol) (GenericSymbol, error)
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

func ParseSource(lexer Lexer, reducer Reducer) (GenericSymbol, error) {

	return ParseSourceWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseSourceWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParsePackageDef(lexer Lexer, reducer Reducer) (GenericSymbol, error) {

	return ParsePackageDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParsePackageDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseTypeDef(lexer Lexer, reducer Reducer) (GenericSymbol, error) {

	return ParseTypeDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTypeDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State3)
	if err != nil {
		var errRetVal GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseNamedFuncDef(lexer Lexer, reducer Reducer) (GenericSymbol, error) {

	return ParseNamedFuncDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseNamedFuncDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State4)
	if err != nil {
		var errRetVal GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseExpression(lexer Lexer, reducer Reducer) (GenericSymbol, error) {

	return ParseExpressionWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseExpressionWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State5)
	if err != nil {
		var errRetVal GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
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
	case LiteralType:
		return "literal"
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

	SourceType                       = SymbolId(342)
	OptionalDefinitionsType          = SymbolId(343)
	OptionalNewlinesType             = SymbolId(344)
	DefinitionsType                  = SymbolId(345)
	DefinitionType                   = SymbolId(346)
	StatementBlockType               = SymbolId(347)
	StatementsType                   = SymbolId(348)
	StatementType                    = SymbolId(349)
	SimpleStatementBodyType          = SymbolId(350)
	StatementBodyType                = SymbolId(351)
	OptionalSimpleStatementBodyType  = SymbolId(352)
	UnaryOpAssignType                = SymbolId(353)
	BinaryOpAssignType               = SymbolId(354)
	UnsafeStatementType              = SymbolId(355)
	JumpTypeType                     = SymbolId(356)
	OptionalJumpLabelType            = SymbolId(357)
	ExpressionsType                  = SymbolId(358)
	OptionalExpressionsType          = SymbolId(359)
	ImportStatementType              = SymbolId(360)
	ImportClauseType                 = SymbolId(361)
	ImportClauseTerminalType         = SymbolId(362)
	ImportClausesType                = SymbolId(363)
	CasePatternsType                 = SymbolId(364)
	VarDeclPatternType               = SymbolId(365)
	VarOrLetType                     = SymbolId(366)
	VarPatternType                   = SymbolId(367)
	TuplePatternType                 = SymbolId(368)
	FieldVarPatternsType             = SymbolId(369)
	FieldVarPatternType              = SymbolId(370)
	OptionalValueTypeType            = SymbolId(371)
	AssignPatternType                = SymbolId(372)
	CasePatternType                  = SymbolId(373)
	ExpressionType                   = SymbolId(374)
	OptionalLabelDeclType            = SymbolId(375)
	SequenceExprType                 = SymbolId(376)
	IfExprType                       = SymbolId(377)
	ConditionType                    = SymbolId(378)
	SwitchExprType                   = SymbolId(379)
	LoopExprType                     = SymbolId(380)
	OptionalSequenceExprType         = SymbolId(381)
	ForAssignmentType                = SymbolId(382)
	CallExprType                     = SymbolId(383)
	OptionalGenericBindingType       = SymbolId(384)
	OptionalGenericArgumentsType     = SymbolId(385)
	GenericArgumentsType             = SymbolId(386)
	OptionalArgumentsType            = SymbolId(387)
	ArgumentsType                    = SymbolId(388)
	ArgumentType                     = SymbolId(389)
	ColonExpressionsType             = SymbolId(390)
	OptionalExpressionType           = SymbolId(391)
	AtomExprType                     = SymbolId(392)
	LiteralType                      = SymbolId(393)
	ImplicitStructExprType           = SymbolId(394)
	AccessExprType                   = SymbolId(395)
	PostfixUnaryExprType             = SymbolId(396)
	PrefixUnaryOpType                = SymbolId(397)
	PrefixUnaryExprType              = SymbolId(398)
	MulOpType                        = SymbolId(399)
	MulExprType                      = SymbolId(400)
	AddOpType                        = SymbolId(401)
	AddExprType                      = SymbolId(402)
	CmpOpType                        = SymbolId(403)
	CmpExprType                      = SymbolId(404)
	AndExprType                      = SymbolId(405)
	OrExprType                       = SymbolId(406)
	InitializableTypeType            = SymbolId(407)
	AtomTypeType                     = SymbolId(408)
	ReturnableTypeType               = SymbolId(409)
	ValueTypeType                    = SymbolId(410)
	TypeDefType                      = SymbolId(411)
	GenericParameterDefType          = SymbolId(412)
	GenericParameterDefsType         = SymbolId(413)
	OptionalGenericParameterDefsType = SymbolId(414)
	OptionalGenericParametersType    = SymbolId(415)
	FieldDefType                     = SymbolId(416)
	ImplicitFieldDefsType            = SymbolId(417)
	OptionalImplicitFieldDefsType    = SymbolId(418)
	ImplicitStructDefType            = SymbolId(419)
	ExplicitFieldDefsType            = SymbolId(420)
	OptionalExplicitFieldDefsType    = SymbolId(421)
	ExplicitStructDefType            = SymbolId(422)
	EnumValueDefType                 = SymbolId(423)
	ImplicitEnumValueDefsType        = SymbolId(424)
	ImplicitEnumDefType              = SymbolId(425)
	ExplicitEnumValueDefsType        = SymbolId(426)
	ExplicitEnumDefType              = SymbolId(427)
	TraitPropertyType                = SymbolId(428)
	TraitPropertiesType              = SymbolId(429)
	OptionalTraitPropertiesType      = SymbolId(430)
	TraitDefType                     = SymbolId(431)
	ReturnTypeType                   = SymbolId(432)
	ParameterDeclType                = SymbolId(433)
	ParameterDeclsType               = SymbolId(434)
	OptionalParameterDeclsType       = SymbolId(435)
	FuncTypeType                     = SymbolId(436)
	MethodSignatureType              = SymbolId(437)
	ParameterDefType                 = SymbolId(438)
	ParameterDefsType                = SymbolId(439)
	OptionalParameterDefsType        = SymbolId(440)
	NamedFuncDefType                 = SymbolId(441)
	AnonymousFuncExprType            = SymbolId(442)
	PackageDefType                   = SymbolId(443)
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
	_ReduceNilToDefinitions                                   = _ReduceType(7)
	_ReduceAddToDefinitions                                   = _ReduceType(8)
	_ReducePackageDefToDefinition                             = _ReduceType(9)
	_ReduceTypeDefToDefinition                                = _ReduceType(10)
	_ReduceNamedFuncDefToDefinition                           = _ReduceType(11)
	_ReduceGlobalVarDefToDefinition                           = _ReduceType(12)
	_ReduceGlobalVarAssignmentToDefinition                    = _ReduceType(13)
	_ReduceStatementBlockToDefinition                         = _ReduceType(14)
	_ReduceToStatementBlock                                   = _ReduceType(15)
	_ReduceEmptyListToStatements                              = _ReduceType(16)
	_ReduceAddToStatements                                    = _ReduceType(17)
	_ReduceImplicitToStatement                                = _ReduceType(18)
	_ReduceExplicitToStatement                                = _ReduceType(19)
	_ReduceUnsafeStatementToSimpleStatementBody               = _ReduceType(20)
	_ReduceExpressionOrImplicitStructToSimpleStatementBody    = _ReduceType(21)
	_ReduceAsyncToSimpleStatementBody                         = _ReduceType(22)
	_ReduceDeferToSimpleStatementBody                         = _ReduceType(23)
	_ReduceJumpStatementToSimpleStatementBody                 = _ReduceType(24)
	_ReduceAssignStatementToSimpleStatementBody               = _ReduceType(25)
	_ReduceUnaryOpAssignStatementToSimpleStatementBody        = _ReduceType(26)
	_ReduceBinaryOpAssignStatementToSimpleStatementBody       = _ReduceType(27)
	_ReduceFallthroughToSimpleStatementBody                   = _ReduceType(28)
	_ReduceSimpleStatementBodyToStatementBody                 = _ReduceType(29)
	_ReduceImportStatementToStatementBody                     = _ReduceType(30)
	_ReduceCaseBranchStatementToStatementBody                 = _ReduceType(31)
	_ReduceDefaultBranchStatementToStatementBody              = _ReduceType(32)
	_ReduceSimpleStatementBodyToOptionalSimpleStatementBody   = _ReduceType(33)
	_ReduceNilToOptionalSimpleStatementBody                   = _ReduceType(34)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(35)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(36)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(37)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(38)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(39)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(40)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(41)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(42)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(43)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(44)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(45)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(46)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(47)
	_ReduceToUnsafeStatement                                  = _ReduceType(48)
	_ReduceReturnToJumpType                                   = _ReduceType(49)
	_ReduceBreakToJumpType                                    = _ReduceType(50)
	_ReduceContinueToJumpType                                 = _ReduceType(51)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(52)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(53)
	_ReduceExpressionToExpressions                            = _ReduceType(54)
	_ReduceAddToExpressions                                   = _ReduceType(55)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(56)
	_ReduceNilToOptionalExpressions                           = _ReduceType(57)
	_ReduceSingleToImportStatement                            = _ReduceType(58)
	_ReduceMultipleToImportStatement                          = _ReduceType(59)
	_ReduceStringLiteralToImportClause                        = _ReduceType(60)
	_ReduceAliasToImportClause                                = _ReduceType(61)
	_ReduceImplicitToImportClauseTerminal                     = _ReduceType(62)
	_ReduceExplicitToImportClauseTerminal                     = _ReduceType(63)
	_ReduceFirstToImportClauses                               = _ReduceType(64)
	_ReduceAddToImportClauses                                 = _ReduceType(65)
	_ReduceCasePatternToCasePatterns                          = _ReduceType(66)
	_ReduceMultipleToCasePatterns                             = _ReduceType(67)
	_ReduceToVarDeclPattern                                   = _ReduceType(68)
	_ReduceVarToVarOrLet                                      = _ReduceType(69)
	_ReduceLetToVarOrLet                                      = _ReduceType(70)
	_ReduceIdentifierToVarPattern                             = _ReduceType(71)
	_ReduceTuplePatternToVarPattern                           = _ReduceType(72)
	_ReduceToTuplePattern                                     = _ReduceType(73)
	_ReduceFieldVarPatternToFieldVarPatterns                  = _ReduceType(74)
	_ReduceAddToFieldVarPatterns                              = _ReduceType(75)
	_ReducePositionalToFieldVarPattern                        = _ReduceType(76)
	_ReduceNamedToFieldVarPattern                             = _ReduceType(77)
	_ReduceDotDotDotToFieldVarPattern                         = _ReduceType(78)
	_ReduceValueTypeToOptionalValueType                       = _ReduceType(79)
	_ReduceNilToOptionalValueType                             = _ReduceType(80)
	_ReduceToAssignPattern                                    = _ReduceType(81)
	_ReduceAssignPatternToCasePattern                         = _ReduceType(82)
	_ReduceEnumMatchPatternToCasePattern                      = _ReduceType(83)
	_ReduceEnumVarDeclPatternToCasePattern                    = _ReduceType(84)
	_ReduceIfExprToExpression                                 = _ReduceType(85)
	_ReduceSwitchExprToExpression                             = _ReduceType(86)
	_ReduceLoopExprToExpression                               = _ReduceType(87)
	_ReduceSequenceExprToExpression                           = _ReduceType(88)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(89)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(90)
	_ReduceOrExprToSequenceExpr                               = _ReduceType(91)
	_ReduceVarDeclPatternToSequenceExpr                       = _ReduceType(92)
	_ReduceAssignVarPatternToSequenceExpr                     = _ReduceType(93)
	_ReduceNoElseToIfExpr                                     = _ReduceType(94)
	_ReduceIfElseToIfExpr                                     = _ReduceType(95)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(96)
	_ReduceSequenceExprToCondition                            = _ReduceType(97)
	_ReduceCaseToCondition                                    = _ReduceType(98)
	_ReduceToSwitchExpr                                       = _ReduceType(99)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(100)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(101)
	_ReduceWhileToLoopExpr                                    = _ReduceType(102)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(103)
	_ReduceForToLoopExpr                                      = _ReduceType(104)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(105)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(106)
	_ReduceSequenceExprToForAssignment                        = _ReduceType(107)
	_ReduceAssignToForAssignment                              = _ReduceType(108)
	_ReduceToCallExpr                                         = _ReduceType(109)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(110)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(111)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(112)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(113)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(114)
	_ReduceAddToGenericArguments                              = _ReduceType(115)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(116)
	_ReduceNilToOptionalArguments                             = _ReduceType(117)
	_ReduceArgumentToArguments                                = _ReduceType(118)
	_ReduceAddToArguments                                     = _ReduceType(119)
	_ReducePositionalToArgument                               = _ReduceType(120)
	_ReduceNamedToArgument                                    = _ReduceType(121)
	_ReduceColonExpressionsToArgument                         = _ReduceType(122)
	_ReduceDotDotDotToArgument                                = _ReduceType(123)
	_ReducePairToColonExpressions                             = _ReduceType(124)
	_ReduceAddToColonExpressions                              = _ReduceType(125)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(126)
	_ReduceNilToOptionalExpression                            = _ReduceType(127)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(128)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(129)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(130)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(131)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(132)
	_ReduceImplicitStructExprToAtomExpr                       = _ReduceType(133)
	_ReduceParseErrorToAtomExpr                               = _ReduceType(134)
	_ReduceTrueToLiteral                                      = _ReduceType(135)
	_ReduceFalseToLiteral                                     = _ReduceType(136)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(137)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(138)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(139)
	_ReduceStringLiteralToLiteral                             = _ReduceType(140)
	_ReduceToImplicitStructExpr                               = _ReduceType(141)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(142)
	_ReduceAccessToAccessExpr                                 = _ReduceType(143)
	_ReduceCallExprToAccessExpr                               = _ReduceType(144)
	_ReduceIndexToAccessExpr                                  = _ReduceType(145)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(146)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(147)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(148)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(149)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(150)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(151)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(152)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(153)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(154)
	_ReduceMulToMulOp                                         = _ReduceType(155)
	_ReduceDivToMulOp                                         = _ReduceType(156)
	_ReduceModToMulOp                                         = _ReduceType(157)
	_ReduceBitAndToMulOp                                      = _ReduceType(158)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(159)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(160)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(161)
	_ReduceOpToMulExpr                                        = _ReduceType(162)
	_ReduceAddToAddOp                                         = _ReduceType(163)
	_ReduceSubToAddOp                                         = _ReduceType(164)
	_ReduceBitOrToAddOp                                       = _ReduceType(165)
	_ReduceBitXorToAddOp                                      = _ReduceType(166)
	_ReduceMulExprToAddExpr                                   = _ReduceType(167)
	_ReduceOpToAddExpr                                        = _ReduceType(168)
	_ReduceEqualToCmpOp                                       = _ReduceType(169)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(170)
	_ReduceLessToCmpOp                                        = _ReduceType(171)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(172)
	_ReduceGreaterToCmpOp                                     = _ReduceType(173)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(174)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(175)
	_ReduceOpToCmpExpr                                        = _ReduceType(176)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(177)
	_ReduceOpToAndExpr                                        = _ReduceType(178)
	_ReduceAndExprToOrExpr                                    = _ReduceType(179)
	_ReduceOpToOrExpr                                         = _ReduceType(180)
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(181)
	_ReduceSliceToInitializableType                           = _ReduceType(182)
	_ReduceArrayToInitializableType                           = _ReduceType(183)
	_ReduceMapToInitializableType                             = _ReduceType(184)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(185)
	_ReduceNamedToAtomType                                    = _ReduceType(186)
	_ReduceExternNamedToAtomType                              = _ReduceType(187)
	_ReduceInferredToAtomType                                 = _ReduceType(188)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(189)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(190)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(191)
	_ReduceTraitDefToAtomType                                 = _ReduceType(192)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(193)
	_ReduceParseErrorToAtomType                               = _ReduceType(194)
	_ReduceAtomTypeToReturnableType                           = _ReduceType(195)
	_ReduceOptionalToReturnableType                           = _ReduceType(196)
	_ReduceResultToReturnableType                             = _ReduceType(197)
	_ReduceReferenceToReturnableType                          = _ReduceType(198)
	_ReducePublicMethodsTraitToReturnableType                 = _ReduceType(199)
	_ReducePublicTraitToReturnableType                        = _ReduceType(200)
	_ReduceReturnableTypeToValueType                          = _ReduceType(201)
	_ReduceTraitIntersectToValueType                          = _ReduceType(202)
	_ReduceTraitUnionToValueType                              = _ReduceType(203)
	_ReduceTraitDifferenceToValueType                         = _ReduceType(204)
	_ReduceDefinitionToTypeDef                                = _ReduceType(205)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(206)
	_ReduceAliasToTypeDef                                     = _ReduceType(207)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(208)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(209)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(210)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(211)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(212)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(213)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(214)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(215)
	_ReduceExplicitToFieldDef                                 = _ReduceType(216)
	_ReduceImplicitToFieldDef                                 = _ReduceType(217)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(218)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(219)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(220)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(221)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(222)
	_ReduceToImplicitStructDef                                = _ReduceType(223)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(224)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(225)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(226)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(227)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(228)
	_ReduceToExplicitStructDef                                = _ReduceType(229)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(230)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(231)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(232)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(233)
	_ReduceToImplicitEnumDef                                  = _ReduceType(234)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(235)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(236)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(237)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(238)
	_ReduceToExplicitEnumDef                                  = _ReduceType(239)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(240)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(241)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(242)
	_ReduceImplicitToTraitProperties                          = _ReduceType(243)
	_ReduceExplicitToTraitProperties                          = _ReduceType(244)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(245)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(246)
	_ReduceToTraitDef                                         = _ReduceType(247)
	_ReduceReturnableTypeToReturnType                         = _ReduceType(248)
	_ReduceNilToReturnType                                    = _ReduceType(249)
	_ReduceArgToParameterDecl                                 = _ReduceType(250)
	_ReduceVarargToParameterDecl                              = _ReduceType(251)
	_ReduceUnamedToParameterDecl                              = _ReduceType(252)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(253)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(254)
	_ReduceAddToParameterDecls                                = _ReduceType(255)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(256)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(257)
	_ReduceToFuncType                                         = _ReduceType(258)
	_ReduceToMethodSignature                                  = _ReduceType(259)
	_ReduceInferredRefArgToParameterDef                       = _ReduceType(260)
	_ReduceInferredRefVarargToParameterDef                    = _ReduceType(261)
	_ReduceArgToParameterDef                                  = _ReduceType(262)
	_ReduceVarargToParameterDef                               = _ReduceType(263)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(264)
	_ReduceAddToParameterDefs                                 = _ReduceType(265)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(266)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(267)
	_ReduceFuncDefToNamedFuncDef                              = _ReduceType(268)
	_ReduceMethodDefToNamedFuncDef                            = _ReduceType(269)
	_ReduceAliasToNamedFuncDef                                = _ReduceType(270)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(271)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(272)
	_ReduceWithSpecToPackageDef                               = _ReduceType(273)
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
	case _ReduceNilToDefinitions:
		return "NilToDefinitions"
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
	case _ReduceLiteralToAtomExpr:
		return "LiteralToAtomExpr"
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
	case _ReduceTrueToLiteral:
		return "TrueToLiteral"
	case _ReduceFalseToLiteral:
		return "FalseToLiteral"
	case _ReduceIntegerLiteralToLiteral:
		return "IntegerLiteralToLiteral"
	case _ReduceFloatLiteralToLiteral:
		return "FloatLiteralToLiteral"
	case _ReduceRuneLiteralToLiteral:
		return "RuneLiteralToLiteral"
	case _ReduceStringLiteralToLiteral:
		return "StringLiteralToLiteral"
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
	case _ReduceQuestionToPostfixUnaryExpr:
		return "QuestionToPostfixUnaryExpr"
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
	case _ReducePrefixOpToPrefixUnaryExpr:
		return "PrefixOpToPrefixUnaryExpr"
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
	case _ReduceOpToMulExpr:
		return "OpToMulExpr"
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
	case _ReduceOpToAddExpr:
		return "OpToAddExpr"
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
	case _ReduceOpToCmpExpr:
		return "OpToCmpExpr"
	case _ReduceCmpExprToAndExpr:
		return "CmpExprToAndExpr"
	case _ReduceOpToAndExpr:
		return "OpToAndExpr"
	case _ReduceAndExprToOrExpr:
		return "AndExprToOrExpr"
	case _ReduceOpToOrExpr:
		return "OpToOrExpr"
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
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	CountSymbol CountSymbol
	ParseError  ParseErrorSymbol
	ValueSymbol ValueSymbol
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case NewlinesToken:
		val, ok := token.(CountSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting CountSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.CountSymbol = val
	case _EndMarker, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken:
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
	case IdentifierToken, LabelDeclToken, JumpLabelToken:
		val, ok := token.(ValueSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting ValueSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.ValueSymbol = val
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
		loc, ok := interface{}(s.CountSymbol).(locator)
		if ok {
			return loc.Loc()
		}
	case ParseErrorToken:
		loc, ok := interface{}(s.ParseError).(locator)
		if ok {
			return loc.Loc()
		}
	case IdentifierToken, LabelDeclToken, JumpLabelToken:
		loc, ok := interface{}(s.ValueSymbol).(locator)
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
		loc, ok := interface{}(s.CountSymbol).(locator)
		if ok {
			return loc.End()
		}
	case ParseErrorToken:
		loc, ok := interface{}(s.ParseError).(locator)
		if ok {
			return loc.End()
		}
	case IdentifierToken, LabelDeclToken, JumpLabelToken:
		loc, ok := interface{}(s.ValueSymbol).(locator)
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
		symbol.Generic_, err = reducer.ToSource(args[0].Generic_)
	case _ReduceNewlinesToOptionalDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.NewlinesToOptionalDefinitions(args[0].CountSymbol)
	case _ReduceDefinitionsToOptionalDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.DefinitionsToOptionalDefinitions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalDefinitions:
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.NilToOptionalDefinitions()
	case _ReduceNewlinesToOptionalNewlines:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalNewlinesType
		symbol.Generic_, err = reducer.NewlinesToOptionalNewlines(args[0].CountSymbol)
	case _ReduceNilToOptionalNewlines:
		symbol.SymbolId_ = OptionalNewlinesType
		symbol.Generic_, err = reducer.NilToOptionalNewlines()
	case _ReduceNilToDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionsType
		symbol.Generic_, err = reducer.NilToDefinitions(args[0].Generic_)
	case _ReduceAddToDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefinitionsType
		symbol.Generic_, err = reducer.AddToDefinitions(args[0].Generic_, args[1].CountSymbol, args[2].Generic_)
	case _ReducePackageDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.PackageDefToDefinition(args[0].Generic_)
	case _ReduceTypeDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.TypeDefToDefinition(args[0].Generic_)
	case _ReduceNamedFuncDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.NamedFuncDefToDefinition(args[0].Generic_)
	case _ReduceGlobalVarDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.GlobalVarDefToDefinition(args[0].Generic_)
	case _ReduceGlobalVarAssignmentToDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.GlobalVarAssignmentToDefinition(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceStatementBlockToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.StatementBlockToDefinition(args[0].Generic_)
	case _ReduceToStatementBlock:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBlockType
		symbol.Generic_, err = reducer.ToStatementBlock(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceEmptyListToStatements:
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.EmptyListToStatements()
	case _ReduceAddToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.AddToStatements(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Generic_, err = reducer.ImplicitToStatement(args[0].Generic_, args[1].CountSymbol)
	case _ReduceExplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Generic_, err = reducer.ExplicitToStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceUnsafeStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.UnsafeStatementToSimpleStatementBody(args[0].Generic_)
	case _ReduceExpressionOrImplicitStructToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.ExpressionOrImplicitStructToSimpleStatementBody(args[0].Generic_)
	case _ReduceAsyncToSimpleStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.AsyncToSimpleStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceDeferToSimpleStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.DeferToSimpleStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceJumpStatementToSimpleStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.JumpStatementToSimpleStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.AssignStatementToSimpleStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceUnaryOpAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.UnaryOpAssignStatementToSimpleStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceBinaryOpAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.BinaryOpAssignStatementToSimpleStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceFallthroughToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Generic_, err = reducer.FallthroughToSimpleStatementBody(args[0].Generic_)
	case _ReduceSimpleStatementBodyToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.SimpleStatementBodyToStatementBody(args[0].Generic_)
	case _ReduceImportStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.ImportStatementToStatementBody(args[0].Generic_)
	case _ReduceCaseBranchStatementToStatementBody:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.CaseBranchStatementToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceDefaultBranchStatementToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.DefaultBranchStatementToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceSimpleStatementBodyToOptionalSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalSimpleStatementBodyType
		symbol.Generic_, err = reducer.SimpleStatementBodyToOptionalSimpleStatementBody(args[0].Generic_)
	case _ReduceNilToOptionalSimpleStatementBody:
		symbol.SymbolId_ = OptionalSimpleStatementBodyType
		symbol.Generic_, err = reducer.NilToOptionalSimpleStatementBody()
	case _ReduceAddOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		symbol.Generic_, err = reducer.AddOneAssignToUnaryOpAssign(args[0].Generic_)
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		symbol.Generic_, err = reducer.SubOneAssignToUnaryOpAssign(args[0].Generic_)
	case _ReduceAddAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.AddAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.SubAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.MulAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.DivAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.ModAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitNegAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitAndAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitOrAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitXorAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitLshiftAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitRshiftAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceToUnsafeStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = UnsafeStatementType
		symbol.Generic_, err = reducer.ToUnsafeStatement(args[0].Generic_, args[1].Generic_, args[2].ValueSymbol, args[3].Generic_, args[4].Generic_)
	case _ReduceReturnToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		symbol.Generic_, err = reducer.ReturnToJumpType(args[0].Generic_)
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		symbol.Generic_, err = reducer.BreakToJumpType(args[0].Generic_)
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		symbol.Generic_, err = reducer.ContinueToJumpType(args[0].Generic_)
	case _ReduceJumpLabelToOptionalJumpLabel:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.JumpLabelToOptionalJumpLabel(args[0].ValueSymbol)
	case _ReduceUnlabelledToOptionalJumpLabel:
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.UnlabelledToOptionalJumpLabel()
	case _ReduceExpressionToExpressions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionsType
		symbol.Generic_, err = reducer.ExpressionToExpressions(args[0].Generic_)
	case _ReduceAddToExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExpressionsType
		symbol.Generic_, err = reducer.AddToExpressions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.SingleToImportStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceMultipleToImportStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ImportStatementType
		symbol.Generic_, err = reducer.MultipleToImportStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceStringLiteralToImportClause:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.StringLiteralToImportClause(args[0].Generic_)
	case _ReduceAliasToImportClause:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.AliasToImportClause(args[0].Generic_, args[1].Generic_, args[2].ValueSymbol)
	case _ReduceImplicitToImportClauseTerminal:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseTerminalType
		symbol.Generic_, err = reducer.ImplicitToImportClauseTerminal(args[0].Generic_, args[1].CountSymbol)
	case _ReduceExplicitToImportClauseTerminal:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseTerminalType
		symbol.Generic_, err = reducer.ExplicitToImportClauseTerminal(args[0].Generic_, args[1].Generic_)
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
		symbol.Generic_, err = reducer.MultipleToCasePatterns(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToVarDeclPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = VarDeclPatternType
		symbol.Generic_, err = reducer.ToVarDeclPattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceVarToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		symbol.Generic_, err = reducer.VarToVarOrLet(args[0].Generic_)
	case _ReduceLetToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		symbol.Generic_, err = reducer.LetToVarOrLet(args[0].Generic_)
	case _ReduceIdentifierToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.IdentifierToVarPattern(args[0].ValueSymbol)
	case _ReduceTuplePatternToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.TuplePatternToVarPattern(args[0].Generic_)
	case _ReduceToTuplePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TuplePatternType
		symbol.Generic_, err = reducer.ToTuplePattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceFieldVarPatternToFieldVarPatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternsType
		symbol.Generic_, err = reducer.FieldVarPatternToFieldVarPatterns(args[0].Generic_)
	case _ReduceAddToFieldVarPatterns:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldVarPatternsType
		symbol.Generic_, err = reducer.AddToFieldVarPatterns(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePositionalToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.PositionalToFieldVarPattern(args[0].Generic_)
	case _ReduceNamedToFieldVarPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.NamedToFieldVarPattern(args[0].ValueSymbol, args[1].Generic_, args[2].Generic_)
	case _ReduceDotDotDotToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.DotDotDotToFieldVarPattern(args[0].Generic_)
	case _ReduceValueTypeToOptionalValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalValueTypeType
		symbol.Generic_, err = reducer.ValueTypeToOptionalValueType(args[0].Generic_)
	case _ReduceNilToOptionalValueType:
		symbol.SymbolId_ = OptionalValueTypeType
		symbol.Generic_, err = reducer.NilToOptionalValueType()
	case _ReduceToAssignPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AssignPatternType
		symbol.Generic_, err = reducer.ToAssignPattern(args[0].Generic_)
	case _ReduceAssignPatternToCasePattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.AssignPatternToCasePattern(args[0].Generic_)
	case _ReduceEnumMatchPatternToCasePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumMatchPatternToCasePattern(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_)
	case _ReduceEnumVarDeclPatternToCasePattern:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumVarDeclPatternToCasePattern(args[0].Generic_, args[1].Generic_, args[2].ValueSymbol, args[3].Generic_)
	case _ReduceIfExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.IfExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceSwitchExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.SwitchExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceLoopExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LoopExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceSequenceExprToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.SequenceExprToExpression(args[0].Generic_)
	case _ReduceLabelDeclToOptionalLabelDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.LabelDeclToOptionalLabelDecl(args[0].ValueSymbol)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.UnlabelledToOptionalLabelDecl()
	case _ReduceOrExprToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		symbol.Generic_, err = reducer.OrExprToSequenceExpr(args[0].Generic_)
	case _ReduceVarDeclPatternToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		symbol.Generic_, err = reducer.VarDeclPatternToSequenceExpr(args[0].Generic_)
	case _ReduceAssignVarPatternToSequenceExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SequenceExprType
		symbol.Generic_, err = reducer.AssignVarPatternToSequenceExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceNoElseToIfExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfExprType
		symbol.Generic_, err = reducer.NoElseToIfExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceIfElseToIfExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprType
		symbol.Generic_, err = reducer.IfElseToIfExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceMultiIfElseToIfExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprType
		symbol.Generic_, err = reducer.MultiIfElseToIfExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceSequenceExprToCondition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ConditionType
		symbol.Generic_, err = reducer.SequenceExprToCondition(args[0].Generic_)
	case _ReduceCaseToCondition:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ConditionType
		symbol.Generic_, err = reducer.CaseToCondition(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceToSwitchExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SwitchExprType
		symbol.Generic_, err = reducer.ToSwitchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceInfiniteToLoopExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.InfiniteToLoopExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceDoWhileToLoopExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.DoWhileToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceWhileToLoopExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.WhileToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceIteratorToLoopExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.IteratorToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceForToLoopExpr:
		args := stack[len(stack)-8:]
		stack = stack[:len(stack)-8]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.ForToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_, args[7].Generic_)
	case _ReduceSequenceExprToOptionalSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalSequenceExprType
		symbol.Generic_, err = reducer.SequenceExprToOptionalSequenceExpr(args[0].Generic_)
	case _ReduceNilToOptionalSequenceExpr:
		symbol.SymbolId_ = OptionalSequenceExprType
		symbol.Generic_, err = reducer.NilToOptionalSequenceExpr()
	case _ReduceSequenceExprToForAssignment:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ForAssignmentType
		symbol.Generic_, err = reducer.SequenceExprToForAssignment(args[0].Generic_)
	case _ReduceAssignToForAssignment:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ForAssignmentType
		symbol.Generic_, err = reducer.AssignToForAssignment(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToCallExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CallExprType
		symbol.Generic_, err = reducer.ToCallExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceBindingToOptionalGenericBinding:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.BindingToOptionalGenericBinding(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ValueTypeToGenericArguments(args[0].Generic_)
	case _ReduceAddToGenericArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.AddToGenericArguments(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.AddToArguments(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePositionalToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.PositionalToArgument(args[0].Generic_)
	case _ReduceNamedToArgument:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.NamedToArgument(args[0].ValueSymbol, args[1].Generic_, args[2].Generic_)
	case _ReduceColonExpressionsToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.ColonExpressionsToArgument(args[0].Generic_)
	case _ReduceDotDotDotToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.DotDotDotToArgument(args[0].Generic_)
	case _ReducePairToColonExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExpressionsType
		symbol.Generic_, err = reducer.PairToColonExpressions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToColonExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExpressionsType
		symbol.Generic_, err = reducer.AddToColonExpressions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExpressionToOptionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.ExpressionToOptionalExpression(args[0].Generic_)
	case _ReduceNilToOptionalExpression:
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.NilToOptionalExpression()
	case _ReduceLiteralToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.LiteralToAtomExpr(args[0].Generic_)
	case _ReduceIdentifierToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.IdentifierToAtomExpr(args[0].ValueSymbol)
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.BlockExprToAtomExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.AnonymousFuncExprToAtomExpr(args[0].Generic_)
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.InitializeExprToAtomExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.ImplicitStructExprToAtomExpr(args[0].Generic_)
	case _ReduceParseErrorToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.ParseErrorToAtomExpr(args[0].ParseError)
	case _ReduceTrueToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.TrueToLiteral(args[0].Generic_)
	case _ReduceFalseToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.FalseToLiteral(args[0].Generic_)
	case _ReduceIntegerLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.IntegerLiteralToLiteral(args[0].Generic_)
	case _ReduceFloatLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.FloatLiteralToLiteral(args[0].Generic_)
	case _ReduceRuneLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.RuneLiteralToLiteral(args[0].Generic_)
	case _ReduceStringLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.StringLiteralToLiteral(args[0].Generic_)
	case _ReduceToImplicitStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructExprType
		symbol.Generic_, err = reducer.ToImplicitStructExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAtomExprToAccessExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.AtomExprToAccessExpr(args[0].Generic_)
	case _ReduceAccessToAccessExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.AccessToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].ValueSymbol)
	case _ReduceCallExprToAccessExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.CallExprToAccessExpr(args[0].Generic_)
	case _ReduceIndexToAccessExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.IndexToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceAccessExprToPostfixUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Generic_, err = reducer.AccessExprToPostfixUnaryExpr(args[0].Generic_)
	case _ReduceQuestionToPostfixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Generic_, err = reducer.QuestionToPostfixUnaryExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceNotToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.NotToPrefixUnaryOp(args[0].Generic_)
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.BitNegToPrefixUnaryOp(args[0].Generic_)
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.SubToPrefixUnaryOp(args[0].Generic_)
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.MulToPrefixUnaryOp(args[0].Generic_)
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.BitAndToPrefixUnaryOp(args[0].Generic_)
	case _ReducePostfixUnaryExprToPrefixUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryExprType
		symbol.Generic_, err = reducer.PostfixUnaryExprToPrefixUnaryExpr(args[0].Generic_)
	case _ReducePrefixOpToPrefixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixUnaryExprType
		symbol.Generic_, err = reducer.PrefixOpToPrefixUnaryExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceMulToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.MulToMulOp(args[0].Generic_)
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.DivToMulOp(args[0].Generic_)
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.ModToMulOp(args[0].Generic_)
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.BitAndToMulOp(args[0].Generic_)
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.BitLshiftToMulOp(args[0].Generic_)
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.BitRshiftToMulOp(args[0].Generic_)
	case _ReducePrefixUnaryExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		symbol.Generic_, err = reducer.PrefixUnaryExprToMulExpr(args[0].Generic_)
	case _ReduceOpToMulExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = MulExprType
		symbol.Generic_, err = reducer.OpToMulExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.AddToAddOp(args[0].Generic_)
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.SubToAddOp(args[0].Generic_)
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.BitOrToAddOp(args[0].Generic_)
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.BitXorToAddOp(args[0].Generic_)
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		symbol.Generic_, err = reducer.MulExprToAddExpr(args[0].Generic_)
	case _ReduceOpToAddExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AddExprType
		symbol.Generic_, err = reducer.OpToAddExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.EqualToCmpOp(args[0].Generic_)
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.NotEqualToCmpOp(args[0].Generic_)
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.LessToCmpOp(args[0].Generic_)
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.LessOrEqualToCmpOp(args[0].Generic_)
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.GreaterToCmpOp(args[0].Generic_)
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.GreaterOrEqualToCmpOp(args[0].Generic_)
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		symbol.Generic_, err = reducer.AddExprToCmpExpr(args[0].Generic_)
	case _ReduceOpToCmpExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CmpExprType
		symbol.Generic_, err = reducer.OpToCmpExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		symbol.Generic_, err = reducer.CmpExprToAndExpr(args[0].Generic_)
	case _ReduceOpToAndExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AndExprType
		symbol.Generic_, err = reducer.OpToAndExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAndExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		symbol.Generic_, err = reducer.AndExprToOrExpr(args[0].Generic_)
	case _ReduceOpToOrExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OrExprType
		symbol.Generic_, err = reducer.OpToOrExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitStructDefToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.ExplicitStructDefToInitializableType(args[0].Generic_)
	case _ReduceSliceToInitializableType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.SliceToInitializableType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceArrayToInitializableType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.ArrayToInitializableType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceMapToInitializableType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.MapToInitializableType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceInitializableTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.InitializableTypeToAtomType(args[0].Generic_)
	case _ReduceNamedToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.NamedToAtomType(args[0].ValueSymbol, args[1].Generic_)
	case _ReduceExternNamedToAtomType:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExternNamedToAtomType(args[0].ValueSymbol, args[1].Generic_, args[2].ValueSymbol, args[3].Generic_)
	case _ReduceInferredToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.InferredToAtomType(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitStructDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitStructDefToAtomType(args[0].Generic_)
	case _ReduceExplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExplicitEnumDefToAtomType(args[0].Generic_)
	case _ReduceImplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitEnumDefToAtomType(args[0].Generic_)
	case _ReduceTraitDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.TraitDefToAtomType(args[0].Generic_)
	case _ReduceFuncTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.FuncTypeToAtomType(args[0].Generic_)
	case _ReduceParseErrorToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ParseErrorToAtomType(args[0].ParseError)
	case _ReduceAtomTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.AtomTypeToReturnableType(args[0].Generic_)
	case _ReduceOptionalToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.OptionalToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReduceResultToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.ResultToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReduceReferenceToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.ReferenceToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicMethodsTraitToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.PublicMethodsTraitToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicTraitToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.PublicTraitToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReduceReturnableTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ReturnableTypeToValueType(args[0].Generic_)
	case _ReduceTraitIntersectToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitIntersectToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitUnionToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitUnionToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitDifferenceToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitDifferenceToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceDefinitionToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.DefinitionToTypeDef(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_, args[3].Generic_)
	case _ReduceConstrainedDefToTypeDef:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.ConstrainedDefToTypeDef(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceAliasToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.AliasToTypeDef(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_, args[3].Generic_)
	case _ReduceUnconstrainedToGenericParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.UnconstrainedToGenericParameterDef(args[0].ValueSymbol)
	case _ReduceConstrainedToGenericParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.ConstrainedToGenericParameterDef(args[0].ValueSymbol, args[1].Generic_)
	case _ReduceGenericParameterDefToGenericParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.GenericParameterDefToGenericParameterDefs(args[0].Generic_)
	case _ReduceAddToGenericParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.AddToGenericParameterDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.GenericToOptionalGenericParameters(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalGenericParameters:
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameters()
	case _ReduceExplicitToFieldDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.ExplicitToFieldDef(args[0].ValueSymbol, args[1].Generic_)
	case _ReduceImplicitToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.ImplicitToFieldDef(args[0].Generic_)
	case _ReduceUnsafeStatementToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.UnsafeStatementToFieldDef(args[0].Generic_)
	case _ReduceFieldDefToImplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToImplicitFieldDefs(args[0].Generic_)
	case _ReduceAddToImplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.AddToImplicitFieldDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ToImplicitStructDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceFieldDefToExplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToExplicitFieldDefs(args[0].Generic_)
	case _ReduceImplicitToExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ImplicitToExplicitFieldDefs(args[0].Generic_, args[1].CountSymbol, args[2].Generic_)
	case _ReduceExplicitToExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ExplicitToExplicitFieldDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ToExplicitStructDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceFieldDefToEnumValueDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = EnumValueDefType
		symbol.Generic_, err = reducer.FieldDefToEnumValueDef(args[0].Generic_)
	case _ReduceDefaultToEnumValueDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = EnumValueDefType
		symbol.Generic_, err = reducer.DefaultToEnumValueDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePairToImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.PairToImplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.AddToImplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToImplicitEnumDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumDefType
		symbol.Generic_, err = reducer.ToImplicitEnumDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitPairToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitPairToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].CountSymbol, args[2].Generic_)
	case _ReduceExplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].CountSymbol, args[2].Generic_)
	case _ReduceToExplicitEnumDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitEnumDefType
		symbol.Generic_, err = reducer.ToExplicitEnumDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
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
		symbol.Generic_, err = reducer.ImplicitToTraitProperties(args[0].Generic_, args[1].CountSymbol, args[2].Generic_)
	case _ReduceExplicitToTraitProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.ExplicitToTraitProperties(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ToTraitDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceReturnableTypeToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.ReturnableTypeToReturnType(args[0].Generic_)
	case _ReduceNilToReturnType:
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.NilToReturnType()
	case _ReduceArgToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.ArgToParameterDecl(args[0].ValueSymbol, args[1].Generic_)
	case _ReduceVarargToParameterDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.VarargToParameterDecl(args[0].ValueSymbol, args[1].Generic_, args[2].Generic_)
	case _ReduceUnamedToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnamedToParameterDecl(args[0].Generic_)
	case _ReduceUnnamedVarargToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnnamedVarargToParameterDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceParameterDeclToParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.ParameterDeclToParameterDecls(args[0].Generic_)
	case _ReduceAddToParameterDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.AddToParameterDecls(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ToFuncType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceToMethodSignature:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = MethodSignatureType
		symbol.Generic_, err = reducer.ToMethodSignature(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceInferredRefArgToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.InferredRefArgToParameterDef(args[0].ValueSymbol)
	case _ReduceInferredRefVarargToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.InferredRefVarargToParameterDef(args[0].ValueSymbol, args[1].Generic_)
	case _ReduceArgToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.ArgToParameterDef(args[0].ValueSymbol, args[1].Generic_)
	case _ReduceVarargToParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.VarargToParameterDef(args[0].ValueSymbol, args[1].Generic_, args[2].Generic_)
	case _ReduceParameterDefToParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.ParameterDefToParameterDefs(args[0].Generic_)
	case _ReduceAddToParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.AddToParameterDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.FuncDefToNamedFuncDef(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_, args[7].Generic_)
	case _ReduceMethodDefToNamedFuncDef:
		args := stack[len(stack)-11:]
		stack = stack[:len(stack)-11]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.Generic_, err = reducer.MethodDefToNamedFuncDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].ValueSymbol, args[5].Generic_, args[6].Generic_, args[7].Generic_, args[8].Generic_, args[9].Generic_, args[10].Generic_)
	case _ReduceAliasToNamedFuncDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.Generic_, err = reducer.AliasToNamedFuncDef(args[0].Generic_, args[1].ValueSymbol, args[2].Generic_, args[3].Generic_)
	case _ReduceToAnonymousFuncExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = AnonymousFuncExprType
		symbol.Generic_, err = reducer.ToAnonymousFuncExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceNoSpecToPackageDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PackageDefType
		symbol.Generic_, err = reducer.NoSpecToPackageDef(args[0].Generic_)
	case _ReduceWithSpecToPackageDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageDefType
		symbol.Generic_, err = reducer.WithSpecToPackageDef(args[0].Generic_, args[1].Generic_)
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
	_ReduceToSourceAction                                           = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceNewlinesToOptionalDefinitionsAction                      = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalDefinitions}
	_ReduceDefinitionsToOptionalDefinitionsAction                   = &_Action{_ReduceAction, 0, _ReduceDefinitionsToOptionalDefinitions}
	_ReduceNilToOptionalDefinitionsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefinitions}
	_ReduceNewlinesToOptionalNewlinesAction                         = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalNewlines}
	_ReduceNilToOptionalNewlinesAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalNewlines}
	_ReduceNilToDefinitionsAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToDefinitions}
	_ReduceAddToDefinitionsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToDefinitions}
	_ReducePackageDefToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                    = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                         = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
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
	_ReduceLiteralToAtomExprAction                                  = &_Action{_ReduceAction, 0, _ReduceLiteralToAtomExpr}
	_ReduceIdentifierToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                        = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                       = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceParseErrorToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceParseErrorToAtomExpr}
	_ReduceTrueToLiteralAction                                      = &_Action{_ReduceAction, 0, _ReduceTrueToLiteral}
	_ReduceFalseToLiteralAction                                     = &_Action{_ReduceAction, 0, _ReduceFalseToLiteral}
	_ReduceIntegerLiteralToLiteralAction                            = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction                              = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction                               = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction                             = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceToImplicitStructExprAction                               = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessExprAction                               = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessExpr}
	_ReduceAccessToAccessExprAction                                 = &_Action{_ReduceAction, 0, _ReduceAccessToAccessExpr}
	_ReduceCallExprToAccessExprAction                               = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessExpr}
	_ReduceIndexToAccessExprAction                                  = &_Action{_ReduceAction, 0, _ReduceIndexToAccessExpr}
	_ReduceAccessExprToPostfixUnaryExprAction                       = &_Action{_ReduceAction, 0, _ReduceAccessExprToPostfixUnaryExpr}
	_ReduceQuestionToPostfixUnaryExprAction                         = &_Action{_ReduceAction, 0, _ReduceQuestionToPostfixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePostfixUnaryExprToPrefixUnaryExprAction                  = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPrefixUnaryExpr}
	_ReducePrefixOpToPrefixUnaryExprAction                          = &_Action{_ReduceAction, 0, _ReducePrefixOpToPrefixUnaryExpr}
	_ReduceMulToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReducePrefixUnaryExprToMulExprAction                           = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToMulExpr}
	_ReduceOpToMulExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToMulExpr}
	_ReduceAddToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                       = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceMulExprToAddExprAction                                   = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceOpToAddExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToAddExpr}
	_ReduceEqualToCmpOpAction                                       = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                    = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                        = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                 = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                     = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                              = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceAddExprToCmpExprAction                                   = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceOpToCmpExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToCmpExpr}
	_ReduceCmpExprToAndExprAction                                   = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceOpToAndExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToAndExpr}
	_ReduceAndExprToOrExprAction                                    = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceOpToOrExprAction                                         = &_Action{_ReduceAction, 0, _ReduceOpToOrExpr}
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
	{_State5, LiteralType}:                        _GotoState48Action,
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
	{_State13, PackageToken}:                      _GotoState14Action,
	{_State13, TypeToken}:                         _GotoState15Action,
	{_State13, FuncToken}:                         _GotoState16Action,
	{_State13, VarToken}:                          _GotoState37Action,
	{_State13, LetToken}:                          _GotoState27Action,
	{_State13, LbraceToken}:                       _GotoState58Action,
	{_State13, DefinitionsType}:                   _GotoState60Action,
	{_State13, DefinitionType}:                    _GotoState59Action,
	{_State13, StatementBlockType}:                _GotoState63Action,
	{_State13, VarDeclPatternType}:                _GotoState65Action,
	{_State13, VarOrLetType}:                      _GotoState57Action,
	{_State13, TypeDefType}:                       _GotoState64Action,
	{_State13, NamedFuncDefType}:                  _GotoState61Action,
	{_State13, PackageDefType}:                    _GotoState62Action,
	{_State14, LbraceToken}:                       _GotoState58Action,
	{_State14, StatementBlockType}:                _GotoState66Action,
	{_State15, IdentifierToken}:                   _GotoState67Action,
	{_State16, IdentifierToken}:                   _GotoState68Action,
	{_State16, LparenToken}:                       _GotoState69Action,
	{_State21, LparenToken}:                       _GotoState70Action,
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
	{_State22, OptionalLabelDeclType}:             _GotoState71Action,
	{_State22, SequenceExprType}:                  _GotoState72Action,
	{_State22, CallExprType}:                      _GotoState43Action,
	{_State22, AtomExprType}:                      _GotoState42Action,
	{_State22, LiteralType}:                       _GotoState48Action,
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
	{_State26, IdentifierToken}:                   _GotoState79Action,
	{_State26, StructToken}:                       _GotoState34Action,
	{_State26, EnumToken}:                         _GotoState76Action,
	{_State26, TraitToken}:                        _GotoState84Action,
	{_State26, FuncToken}:                         _GotoState78Action,
	{_State26, LparenToken}:                       _GotoState80Action,
	{_State26, LbracketToken}:                     _GotoState26Action,
	{_State26, DotToken}:                          _GotoState75Action,
	{_State26, QuestionToken}:                     _GotoState82Action,
	{_State26, ExclaimToken}:                      _GotoState77Action,
	{_State26, TildeTildeToken}:                   _GotoState83Action,
	{_State26, BitNegToken}:                       _GotoState74Action,
	{_State26, BitAndToken}:                       _GotoState73Action,
	{_State26, ParseErrorToken}:                   _GotoState81Action,
	{_State26, InitializableTypeType}:             _GotoState90Action,
	{_State26, AtomTypeType}:                      _GotoState85Action,
	{_State26, ReturnableTypeType}:                _GotoState91Action,
	{_State26, ValueTypeType}:                     _GotoState93Action,
	{_State26, ImplicitStructDefType}:             _GotoState89Action,
	{_State26, ExplicitStructDefType}:             _GotoState45Action,
	{_State26, ImplicitEnumDefType}:               _GotoState88Action,
	{_State26, ExplicitEnumDefType}:               _GotoState86Action,
	{_State26, TraitDefType}:                      _GotoState92Action,
	{_State26, FuncTypeType}:                      _GotoState87Action,
	{_State28, IntegerLiteralToken}:               _GotoState24Action,
	{_State28, FloatLiteralToken}:                 _GotoState20Action,
	{_State28, RuneLiteralToken}:                  _GotoState32Action,
	{_State28, StringLiteralToken}:                _GotoState33Action,
	{_State28, IdentifierToken}:                   _GotoState95Action,
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
	{_State28, DotDotDotToken}:                    _GotoState94Action,
	{_State28, SubToken}:                          _GotoState35Action,
	{_State28, MulToken}:                          _GotoState29Action,
	{_State28, BitNegToken}:                       _GotoState18Action,
	{_State28, BitAndToken}:                       _GotoState17Action,
	{_State28, GreaterToken}:                      _GotoState22Action,
	{_State28, ParseErrorToken}:                   _GotoState31Action,
	{_State28, VarDeclPatternType}:                _GotoState56Action,
	{_State28, VarOrLetType}:                      _GotoState57Action,
	{_State28, ExpressionType}:                    _GotoState99Action,
	{_State28, OptionalLabelDeclType}:             _GotoState50Action,
	{_State28, SequenceExprType}:                  _GotoState55Action,
	{_State28, CallExprType}:                      _GotoState43Action,
	{_State28, ArgumentsType}:                     _GotoState97Action,
	{_State28, ArgumentType}:                      _GotoState96Action,
	{_State28, ColonExpressionsType}:              _GotoState98Action,
	{_State28, OptionalExpressionType}:            _GotoState100Action,
	{_State28, AtomExprType}:                      _GotoState42Action,
	{_State28, LiteralType}:                       _GotoState48Action,
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
	{_State34, LparenToken}:                       _GotoState101Action,
	{_State38, LbracketToken}:                     _GotoState104Action,
	{_State38, DotToken}:                          _GotoState103Action,
	{_State38, QuestionToken}:                     _GotoState105Action,
	{_State38, DollarLbracketToken}:               _GotoState102Action,
	{_State38, OptionalGenericBindingType}:        _GotoState106Action,
	{_State39, AddToken}:                          _GotoState107Action,
	{_State39, SubToken}:                          _GotoState110Action,
	{_State39, BitXorToken}:                       _GotoState109Action,
	{_State39, BitOrToken}:                        _GotoState108Action,
	{_State39, AddOpType}:                         _GotoState111Action,
	{_State40, AndToken}:                          _GotoState112Action,
	{_State44, EqualToken}:                        _GotoState113Action,
	{_State44, NotEqualToken}:                     _GotoState118Action,
	{_State44, LessToken}:                         _GotoState116Action,
	{_State44, LessOrEqualToken}:                  _GotoState117Action,
	{_State44, GreaterToken}:                      _GotoState114Action,
	{_State44, GreaterOrEqualToken}:               _GotoState115Action,
	{_State44, CmpOpType}:                         _GotoState119Action,
	{_State47, LparenToken}:                       _GotoState120Action,
	{_State49, MulToken}:                          _GotoState126Action,
	{_State49, DivToken}:                          _GotoState124Action,
	{_State49, ModToken}:                          _GotoState125Action,
	{_State49, BitAndToken}:                       _GotoState121Action,
	{_State49, BitLshiftToken}:                    _GotoState122Action,
	{_State49, BitRshiftToken}:                    _GotoState123Action,
	{_State49, MulOpType}:                         _GotoState127Action,
	{_State50, IfToken}:                           _GotoState130Action,
	{_State50, SwitchToken}:                       _GotoState131Action,
	{_State50, ForToken}:                          _GotoState129Action,
	{_State50, DoToken}:                           _GotoState128Action,
	{_State50, LbraceToken}:                       _GotoState58Action,
	{_State50, StatementBlockType}:                _GotoState134Action,
	{_State50, IfExprType}:                        _GotoState132Action,
	{_State50, SwitchExprType}:                    _GotoState135Action,
	{_State50, LoopExprType}:                      _GotoState133Action,
	{_State51, OrToken}:                           _GotoState136Action,
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
	{_State54, OptionalLabelDeclType}:             _GotoState71Action,
	{_State54, CallExprType}:                      _GotoState43Action,
	{_State54, AtomExprType}:                      _GotoState42Action,
	{_State54, LiteralType}:                       _GotoState48Action,
	{_State54, ImplicitStructExprType}:            _GotoState46Action,
	{_State54, AccessExprType}:                    _GotoState38Action,
	{_State54, PostfixUnaryExprType}:              _GotoState52Action,
	{_State54, PrefixUnaryOpType}:                 _GotoState54Action,
	{_State54, PrefixUnaryExprType}:               _GotoState137Action,
	{_State54, InitializableTypeType}:             _GotoState47Action,
	{_State54, ExplicitStructDefType}:             _GotoState45Action,
	{_State54, AnonymousFuncExprType}:             _GotoState41Action,
	{_State57, IdentifierToken}:                   _GotoState138Action,
	{_State57, LparenToken}:                       _GotoState139Action,
	{_State57, VarPatternType}:                    _GotoState141Action,
	{_State57, TuplePatternType}:                  _GotoState140Action,
	{_State58, StatementsType}:                    _GotoState142Action,
	{_State60, NewlinesToken}:                     _GotoState143Action,
	{_State60, OptionalNewlinesType}:              _GotoState144Action,
	{_State65, AssignToken}:                       _GotoState145Action,
	{_State67, DollarLbracketToken}:               _GotoState147Action,
	{_State67, AssignToken}:                       _GotoState146Action,
	{_State67, OptionalGenericParametersType}:     _GotoState148Action,
	{_State68, DollarLbracketToken}:               _GotoState147Action,
	{_State68, AssignToken}:                       _GotoState149Action,
	{_State68, OptionalGenericParametersType}:     _GotoState150Action,
	{_State69, IdentifierToken}:                   _GotoState151Action,
	{_State69, ParameterDefType}:                  _GotoState152Action,
	{_State70, IdentifierToken}:                   _GotoState151Action,
	{_State70, ParameterDefType}:                  _GotoState154Action,
	{_State70, ParameterDefsType}:                 _GotoState155Action,
	{_State70, OptionalParameterDefsType}:         _GotoState153Action,
	{_State71, LbraceToken}:                       _GotoState58Action,
	{_State71, StatementBlockType}:                _GotoState134Action,
	{_State73, IdentifierToken}:                   _GotoState79Action,
	{_State73, StructToken}:                       _GotoState34Action,
	{_State73, EnumToken}:                         _GotoState76Action,
	{_State73, TraitToken}:                        _GotoState84Action,
	{_State73, FuncToken}:                         _GotoState78Action,
	{_State73, LparenToken}:                       _GotoState80Action,
	{_State73, LbracketToken}:                     _GotoState26Action,
	{_State73, DotToken}:                          _GotoState75Action,
	{_State73, QuestionToken}:                     _GotoState82Action,
	{_State73, ExclaimToken}:                      _GotoState77Action,
	{_State73, TildeTildeToken}:                   _GotoState83Action,
	{_State73, BitNegToken}:                       _GotoState74Action,
	{_State73, BitAndToken}:                       _GotoState73Action,
	{_State73, ParseErrorToken}:                   _GotoState81Action,
	{_State73, InitializableTypeType}:             _GotoState90Action,
	{_State73, AtomTypeType}:                      _GotoState85Action,
	{_State73, ReturnableTypeType}:                _GotoState156Action,
	{_State73, ImplicitStructDefType}:             _GotoState89Action,
	{_State73, ExplicitStructDefType}:             _GotoState45Action,
	{_State73, ImplicitEnumDefType}:               _GotoState88Action,
	{_State73, ExplicitEnumDefType}:               _GotoState86Action,
	{_State73, TraitDefType}:                      _GotoState92Action,
	{_State73, FuncTypeType}:                      _GotoState87Action,
	{_State74, IdentifierToken}:                   _GotoState79Action,
	{_State74, StructToken}:                       _GotoState34Action,
	{_State74, EnumToken}:                         _GotoState76Action,
	{_State74, TraitToken}:                        _GotoState84Action,
	{_State74, FuncToken}:                         _GotoState78Action,
	{_State74, LparenToken}:                       _GotoState80Action,
	{_State74, LbracketToken}:                     _GotoState26Action,
	{_State74, DotToken}:                          _GotoState75Action,
	{_State74, QuestionToken}:                     _GotoState82Action,
	{_State74, ExclaimToken}:                      _GotoState77Action,
	{_State74, TildeTildeToken}:                   _GotoState83Action,
	{_State74, BitNegToken}:                       _GotoState74Action,
	{_State74, BitAndToken}:                       _GotoState73Action,
	{_State74, ParseErrorToken}:                   _GotoState81Action,
	{_State74, InitializableTypeType}:             _GotoState90Action,
	{_State74, AtomTypeType}:                      _GotoState85Action,
	{_State74, ReturnableTypeType}:                _GotoState157Action,
	{_State74, ImplicitStructDefType}:             _GotoState89Action,
	{_State74, ExplicitStructDefType}:             _GotoState45Action,
	{_State74, ImplicitEnumDefType}:               _GotoState88Action,
	{_State74, ExplicitEnumDefType}:               _GotoState86Action,
	{_State74, TraitDefType}:                      _GotoState92Action,
	{_State74, FuncTypeType}:                      _GotoState87Action,
	{_State75, DollarLbracketToken}:               _GotoState102Action,
	{_State75, OptionalGenericBindingType}:        _GotoState158Action,
	{_State76, LparenToken}:                       _GotoState159Action,
	{_State77, IdentifierToken}:                   _GotoState79Action,
	{_State77, StructToken}:                       _GotoState34Action,
	{_State77, EnumToken}:                         _GotoState76Action,
	{_State77, TraitToken}:                        _GotoState84Action,
	{_State77, FuncToken}:                         _GotoState78Action,
	{_State77, LparenToken}:                       _GotoState80Action,
	{_State77, LbracketToken}:                     _GotoState26Action,
	{_State77, DotToken}:                          _GotoState75Action,
	{_State77, QuestionToken}:                     _GotoState82Action,
	{_State77, ExclaimToken}:                      _GotoState77Action,
	{_State77, TildeTildeToken}:                   _GotoState83Action,
	{_State77, BitNegToken}:                       _GotoState74Action,
	{_State77, BitAndToken}:                       _GotoState73Action,
	{_State77, ParseErrorToken}:                   _GotoState81Action,
	{_State77, InitializableTypeType}:             _GotoState90Action,
	{_State77, AtomTypeType}:                      _GotoState85Action,
	{_State77, ReturnableTypeType}:                _GotoState160Action,
	{_State77, ImplicitStructDefType}:             _GotoState89Action,
	{_State77, ExplicitStructDefType}:             _GotoState45Action,
	{_State77, ImplicitEnumDefType}:               _GotoState88Action,
	{_State77, ExplicitEnumDefType}:               _GotoState86Action,
	{_State77, TraitDefType}:                      _GotoState92Action,
	{_State77, FuncTypeType}:                      _GotoState87Action,
	{_State78, LparenToken}:                       _GotoState161Action,
	{_State79, DotToken}:                          _GotoState162Action,
	{_State79, DollarLbracketToken}:               _GotoState102Action,
	{_State79, OptionalGenericBindingType}:        _GotoState163Action,
	{_State80, IdentifierToken}:                   _GotoState164Action,
	{_State80, UnsafeToken}:                       _GotoState165Action,
	{_State80, StructToken}:                       _GotoState34Action,
	{_State80, EnumToken}:                         _GotoState76Action,
	{_State80, TraitToken}:                        _GotoState84Action,
	{_State80, FuncToken}:                         _GotoState78Action,
	{_State80, LparenToken}:                       _GotoState80Action,
	{_State80, LbracketToken}:                     _GotoState26Action,
	{_State80, DotToken}:                          _GotoState75Action,
	{_State80, QuestionToken}:                     _GotoState82Action,
	{_State80, ExclaimToken}:                      _GotoState77Action,
	{_State80, TildeTildeToken}:                   _GotoState83Action,
	{_State80, BitNegToken}:                       _GotoState74Action,
	{_State80, BitAndToken}:                       _GotoState73Action,
	{_State80, ParseErrorToken}:                   _GotoState81Action,
	{_State80, UnsafeStatementType}:               _GotoState171Action,
	{_State80, InitializableTypeType}:             _GotoState90Action,
	{_State80, AtomTypeType}:                      _GotoState85Action,
	{_State80, ReturnableTypeType}:                _GotoState91Action,
	{_State80, ValueTypeType}:                     _GotoState172Action,
	{_State80, FieldDefType}:                      _GotoState167Action,
	{_State80, ImplicitFieldDefsType}:             _GotoState169Action,
	{_State80, OptionalImplicitFieldDefsType}:     _GotoState170Action,
	{_State80, ImplicitStructDefType}:             _GotoState89Action,
	{_State80, ExplicitStructDefType}:             _GotoState45Action,
	{_State80, EnumValueDefType}:                  _GotoState166Action,
	{_State80, ImplicitEnumValueDefsType}:         _GotoState168Action,
	{_State80, ImplicitEnumDefType}:               _GotoState88Action,
	{_State80, ExplicitEnumDefType}:               _GotoState86Action,
	{_State80, TraitDefType}:                      _GotoState92Action,
	{_State80, FuncTypeType}:                      _GotoState87Action,
	{_State82, IdentifierToken}:                   _GotoState79Action,
	{_State82, StructToken}:                       _GotoState34Action,
	{_State82, EnumToken}:                         _GotoState76Action,
	{_State82, TraitToken}:                        _GotoState84Action,
	{_State82, FuncToken}:                         _GotoState78Action,
	{_State82, LparenToken}:                       _GotoState80Action,
	{_State82, LbracketToken}:                     _GotoState26Action,
	{_State82, DotToken}:                          _GotoState75Action,
	{_State82, QuestionToken}:                     _GotoState82Action,
	{_State82, ExclaimToken}:                      _GotoState77Action,
	{_State82, TildeTildeToken}:                   _GotoState83Action,
	{_State82, BitNegToken}:                       _GotoState74Action,
	{_State82, BitAndToken}:                       _GotoState73Action,
	{_State82, ParseErrorToken}:                   _GotoState81Action,
	{_State82, InitializableTypeType}:             _GotoState90Action,
	{_State82, AtomTypeType}:                      _GotoState85Action,
	{_State82, ReturnableTypeType}:                _GotoState173Action,
	{_State82, ImplicitStructDefType}:             _GotoState89Action,
	{_State82, ExplicitStructDefType}:             _GotoState45Action,
	{_State82, ImplicitEnumDefType}:               _GotoState88Action,
	{_State82, ExplicitEnumDefType}:               _GotoState86Action,
	{_State82, TraitDefType}:                      _GotoState92Action,
	{_State82, FuncTypeType}:                      _GotoState87Action,
	{_State83, IdentifierToken}:                   _GotoState79Action,
	{_State83, StructToken}:                       _GotoState34Action,
	{_State83, EnumToken}:                         _GotoState76Action,
	{_State83, TraitToken}:                        _GotoState84Action,
	{_State83, FuncToken}:                         _GotoState78Action,
	{_State83, LparenToken}:                       _GotoState80Action,
	{_State83, LbracketToken}:                     _GotoState26Action,
	{_State83, DotToken}:                          _GotoState75Action,
	{_State83, QuestionToken}:                     _GotoState82Action,
	{_State83, ExclaimToken}:                      _GotoState77Action,
	{_State83, TildeTildeToken}:                   _GotoState83Action,
	{_State83, BitNegToken}:                       _GotoState74Action,
	{_State83, BitAndToken}:                       _GotoState73Action,
	{_State83, ParseErrorToken}:                   _GotoState81Action,
	{_State83, InitializableTypeType}:             _GotoState90Action,
	{_State83, AtomTypeType}:                      _GotoState85Action,
	{_State83, ReturnableTypeType}:                _GotoState174Action,
	{_State83, ImplicitStructDefType}:             _GotoState89Action,
	{_State83, ExplicitStructDefType}:             _GotoState45Action,
	{_State83, ImplicitEnumDefType}:               _GotoState88Action,
	{_State83, ExplicitEnumDefType}:               _GotoState86Action,
	{_State83, TraitDefType}:                      _GotoState92Action,
	{_State83, FuncTypeType}:                      _GotoState87Action,
	{_State84, LparenToken}:                       _GotoState175Action,
	{_State93, RbracketToken}:                     _GotoState180Action,
	{_State93, CommaToken}:                        _GotoState178Action,
	{_State93, ColonToken}:                        _GotoState177Action,
	{_State93, AddToken}:                          _GotoState176Action,
	{_State93, SubToken}:                          _GotoState181Action,
	{_State93, MulToken}:                          _GotoState179Action,
	{_State95, AssignToken}:                       _GotoState182Action,
	{_State97, RparenToken}:                       _GotoState184Action,
	{_State97, CommaToken}:                        _GotoState183Action,
	{_State98, ColonToken}:                        _GotoState185Action,
	{_State100, ColonToken}:                       _GotoState186Action,
	{_State101, IdentifierToken}:                  _GotoState164Action,
	{_State101, UnsafeToken}:                      _GotoState165Action,
	{_State101, StructToken}:                      _GotoState34Action,
	{_State101, EnumToken}:                        _GotoState76Action,
	{_State101, TraitToken}:                       _GotoState84Action,
	{_State101, FuncToken}:                        _GotoState78Action,
	{_State101, LparenToken}:                      _GotoState80Action,
	{_State101, LbracketToken}:                    _GotoState26Action,
	{_State101, DotToken}:                         _GotoState75Action,
	{_State101, QuestionToken}:                    _GotoState82Action,
	{_State101, ExclaimToken}:                     _GotoState77Action,
	{_State101, TildeTildeToken}:                  _GotoState83Action,
	{_State101, BitNegToken}:                      _GotoState74Action,
	{_State101, BitAndToken}:                      _GotoState73Action,
	{_State101, ParseErrorToken}:                  _GotoState81Action,
	{_State101, UnsafeStatementType}:              _GotoState171Action,
	{_State101, InitializableTypeType}:            _GotoState90Action,
	{_State101, AtomTypeType}:                     _GotoState85Action,
	{_State101, ReturnableTypeType}:               _GotoState91Action,
	{_State101, ValueTypeType}:                    _GotoState172Action,
	{_State101, FieldDefType}:                     _GotoState188Action,
	{_State101, ImplicitStructDefType}:            _GotoState89Action,
	{_State101, ExplicitFieldDefsType}:            _GotoState187Action,
	{_State101, OptionalExplicitFieldDefsType}:    _GotoState189Action,
	{_State101, ExplicitStructDefType}:            _GotoState45Action,
	{_State101, ImplicitEnumDefType}:              _GotoState88Action,
	{_State101, ExplicitEnumDefType}:              _GotoState86Action,
	{_State101, TraitDefType}:                     _GotoState92Action,
	{_State101, FuncTypeType}:                     _GotoState87Action,
	{_State102, IdentifierToken}:                  _GotoState79Action,
	{_State102, StructToken}:                      _GotoState34Action,
	{_State102, EnumToken}:                        _GotoState76Action,
	{_State102, TraitToken}:                       _GotoState84Action,
	{_State102, FuncToken}:                        _GotoState78Action,
	{_State102, LparenToken}:                      _GotoState80Action,
	{_State102, LbracketToken}:                    _GotoState26Action,
	{_State102, DotToken}:                         _GotoState75Action,
	{_State102, QuestionToken}:                    _GotoState82Action,
	{_State102, ExclaimToken}:                     _GotoState77Action,
	{_State102, TildeTildeToken}:                  _GotoState83Action,
	{_State102, BitNegToken}:                      _GotoState74Action,
	{_State102, BitAndToken}:                      _GotoState73Action,
	{_State102, ParseErrorToken}:                  _GotoState81Action,
	{_State102, OptionalGenericArgumentsType}:     _GotoState191Action,
	{_State102, GenericArgumentsType}:             _GotoState190Action,
	{_State102, InitializableTypeType}:            _GotoState90Action,
	{_State102, AtomTypeType}:                     _GotoState85Action,
	{_State102, ReturnableTypeType}:               _GotoState91Action,
	{_State102, ValueTypeType}:                    _GotoState192Action,
	{_State102, ImplicitStructDefType}:            _GotoState89Action,
	{_State102, ExplicitStructDefType}:            _GotoState45Action,
	{_State102, ImplicitEnumDefType}:              _GotoState88Action,
	{_State102, ExplicitEnumDefType}:              _GotoState86Action,
	{_State102, TraitDefType}:                     _GotoState92Action,
	{_State102, FuncTypeType}:                     _GotoState87Action,
	{_State103, IdentifierToken}:                  _GotoState193Action,
	{_State104, IntegerLiteralToken}:              _GotoState24Action,
	{_State104, FloatLiteralToken}:                _GotoState20Action,
	{_State104, RuneLiteralToken}:                 _GotoState32Action,
	{_State104, StringLiteralToken}:               _GotoState33Action,
	{_State104, IdentifierToken}:                  _GotoState95Action,
	{_State104, TrueToken}:                        _GotoState36Action,
	{_State104, FalseToken}:                       _GotoState19Action,
	{_State104, StructToken}:                      _GotoState34Action,
	{_State104, FuncToken}:                        _GotoState21Action,
	{_State104, VarToken}:                         _GotoState37Action,
	{_State104, LetToken}:                         _GotoState27Action,
	{_State104, NotToken}:                         _GotoState30Action,
	{_State104, LabelDeclToken}:                   _GotoState25Action,
	{_State104, LparenToken}:                      _GotoState28Action,
	{_State104, LbracketToken}:                    _GotoState26Action,
	{_State104, DotDotDotToken}:                   _GotoState94Action,
	{_State104, SubToken}:                         _GotoState35Action,
	{_State104, MulToken}:                         _GotoState29Action,
	{_State104, BitNegToken}:                      _GotoState18Action,
	{_State104, BitAndToken}:                      _GotoState17Action,
	{_State104, GreaterToken}:                     _GotoState22Action,
	{_State104, ParseErrorToken}:                  _GotoState31Action,
	{_State104, VarDeclPatternType}:               _GotoState56Action,
	{_State104, VarOrLetType}:                     _GotoState57Action,
	{_State104, ExpressionType}:                   _GotoState99Action,
	{_State104, OptionalLabelDeclType}:            _GotoState50Action,
	{_State104, SequenceExprType}:                 _GotoState55Action,
	{_State104, CallExprType}:                     _GotoState43Action,
	{_State104, ArgumentType}:                     _GotoState194Action,
	{_State104, ColonExpressionsType}:             _GotoState98Action,
	{_State104, OptionalExpressionType}:           _GotoState100Action,
	{_State104, AtomExprType}:                     _GotoState42Action,
	{_State104, LiteralType}:                      _GotoState48Action,
	{_State104, ImplicitStructExprType}:           _GotoState46Action,
	{_State104, AccessExprType}:                   _GotoState38Action,
	{_State104, PostfixUnaryExprType}:             _GotoState52Action,
	{_State104, PrefixUnaryOpType}:                _GotoState54Action,
	{_State104, PrefixUnaryExprType}:              _GotoState53Action,
	{_State104, MulExprType}:                      _GotoState49Action,
	{_State104, AddExprType}:                      _GotoState39Action,
	{_State104, CmpExprType}:                      _GotoState44Action,
	{_State104, AndExprType}:                      _GotoState40Action,
	{_State104, OrExprType}:                       _GotoState51Action,
	{_State104, InitializableTypeType}:            _GotoState47Action,
	{_State104, ExplicitStructDefType}:            _GotoState45Action,
	{_State104, AnonymousFuncExprType}:            _GotoState41Action,
	{_State106, LparenToken}:                      _GotoState195Action,
	{_State111, IntegerLiteralToken}:              _GotoState24Action,
	{_State111, FloatLiteralToken}:                _GotoState20Action,
	{_State111, RuneLiteralToken}:                 _GotoState32Action,
	{_State111, StringLiteralToken}:               _GotoState33Action,
	{_State111, IdentifierToken}:                  _GotoState23Action,
	{_State111, TrueToken}:                        _GotoState36Action,
	{_State111, FalseToken}:                       _GotoState19Action,
	{_State111, StructToken}:                      _GotoState34Action,
	{_State111, FuncToken}:                        _GotoState21Action,
	{_State111, NotToken}:                         _GotoState30Action,
	{_State111, LabelDeclToken}:                   _GotoState25Action,
	{_State111, LparenToken}:                      _GotoState28Action,
	{_State111, LbracketToken}:                    _GotoState26Action,
	{_State111, SubToken}:                         _GotoState35Action,
	{_State111, MulToken}:                         _GotoState29Action,
	{_State111, BitNegToken}:                      _GotoState18Action,
	{_State111, BitAndToken}:                      _GotoState17Action,
	{_State111, ParseErrorToken}:                  _GotoState31Action,
	{_State111, OptionalLabelDeclType}:            _GotoState71Action,
	{_State111, CallExprType}:                     _GotoState43Action,
	{_State111, AtomExprType}:                     _GotoState42Action,
	{_State111, LiteralType}:                      _GotoState48Action,
	{_State111, ImplicitStructExprType}:           _GotoState46Action,
	{_State111, AccessExprType}:                   _GotoState38Action,
	{_State111, PostfixUnaryExprType}:             _GotoState52Action,
	{_State111, PrefixUnaryOpType}:                _GotoState54Action,
	{_State111, PrefixUnaryExprType}:              _GotoState53Action,
	{_State111, MulExprType}:                      _GotoState196Action,
	{_State111, InitializableTypeType}:            _GotoState47Action,
	{_State111, ExplicitStructDefType}:            _GotoState45Action,
	{_State111, AnonymousFuncExprType}:            _GotoState41Action,
	{_State112, IntegerLiteralToken}:              _GotoState24Action,
	{_State112, FloatLiteralToken}:                _GotoState20Action,
	{_State112, RuneLiteralToken}:                 _GotoState32Action,
	{_State112, StringLiteralToken}:               _GotoState33Action,
	{_State112, IdentifierToken}:                  _GotoState23Action,
	{_State112, TrueToken}:                        _GotoState36Action,
	{_State112, FalseToken}:                       _GotoState19Action,
	{_State112, StructToken}:                      _GotoState34Action,
	{_State112, FuncToken}:                        _GotoState21Action,
	{_State112, NotToken}:                         _GotoState30Action,
	{_State112, LabelDeclToken}:                   _GotoState25Action,
	{_State112, LparenToken}:                      _GotoState28Action,
	{_State112, LbracketToken}:                    _GotoState26Action,
	{_State112, SubToken}:                         _GotoState35Action,
	{_State112, MulToken}:                         _GotoState29Action,
	{_State112, BitNegToken}:                      _GotoState18Action,
	{_State112, BitAndToken}:                      _GotoState17Action,
	{_State112, ParseErrorToken}:                  _GotoState31Action,
	{_State112, OptionalLabelDeclType}:            _GotoState71Action,
	{_State112, CallExprType}:                     _GotoState43Action,
	{_State112, AtomExprType}:                     _GotoState42Action,
	{_State112, LiteralType}:                      _GotoState48Action,
	{_State112, ImplicitStructExprType}:           _GotoState46Action,
	{_State112, AccessExprType}:                   _GotoState38Action,
	{_State112, PostfixUnaryExprType}:             _GotoState52Action,
	{_State112, PrefixUnaryOpType}:                _GotoState54Action,
	{_State112, PrefixUnaryExprType}:              _GotoState53Action,
	{_State112, MulExprType}:                      _GotoState49Action,
	{_State112, AddExprType}:                      _GotoState39Action,
	{_State112, CmpExprType}:                      _GotoState197Action,
	{_State112, InitializableTypeType}:            _GotoState47Action,
	{_State112, ExplicitStructDefType}:            _GotoState45Action,
	{_State112, AnonymousFuncExprType}:            _GotoState41Action,
	{_State119, IntegerLiteralToken}:              _GotoState24Action,
	{_State119, FloatLiteralToken}:                _GotoState20Action,
	{_State119, RuneLiteralToken}:                 _GotoState32Action,
	{_State119, StringLiteralToken}:               _GotoState33Action,
	{_State119, IdentifierToken}:                  _GotoState23Action,
	{_State119, TrueToken}:                        _GotoState36Action,
	{_State119, FalseToken}:                       _GotoState19Action,
	{_State119, StructToken}:                      _GotoState34Action,
	{_State119, FuncToken}:                        _GotoState21Action,
	{_State119, NotToken}:                         _GotoState30Action,
	{_State119, LabelDeclToken}:                   _GotoState25Action,
	{_State119, LparenToken}:                      _GotoState28Action,
	{_State119, LbracketToken}:                    _GotoState26Action,
	{_State119, SubToken}:                         _GotoState35Action,
	{_State119, MulToken}:                         _GotoState29Action,
	{_State119, BitNegToken}:                      _GotoState18Action,
	{_State119, BitAndToken}:                      _GotoState17Action,
	{_State119, ParseErrorToken}:                  _GotoState31Action,
	{_State119, OptionalLabelDeclType}:            _GotoState71Action,
	{_State119, CallExprType}:                     _GotoState43Action,
	{_State119, AtomExprType}:                     _GotoState42Action,
	{_State119, LiteralType}:                      _GotoState48Action,
	{_State119, ImplicitStructExprType}:           _GotoState46Action,
	{_State119, AccessExprType}:                   _GotoState38Action,
	{_State119, PostfixUnaryExprType}:             _GotoState52Action,
	{_State119, PrefixUnaryOpType}:                _GotoState54Action,
	{_State119, PrefixUnaryExprType}:              _GotoState53Action,
	{_State119, MulExprType}:                      _GotoState49Action,
	{_State119, AddExprType}:                      _GotoState198Action,
	{_State119, InitializableTypeType}:            _GotoState47Action,
	{_State119, ExplicitStructDefType}:            _GotoState45Action,
	{_State119, AnonymousFuncExprType}:            _GotoState41Action,
	{_State120, IntegerLiteralToken}:              _GotoState24Action,
	{_State120, FloatLiteralToken}:                _GotoState20Action,
	{_State120, RuneLiteralToken}:                 _GotoState32Action,
	{_State120, StringLiteralToken}:               _GotoState33Action,
	{_State120, IdentifierToken}:                  _GotoState95Action,
	{_State120, TrueToken}:                        _GotoState36Action,
	{_State120, FalseToken}:                       _GotoState19Action,
	{_State120, StructToken}:                      _GotoState34Action,
	{_State120, FuncToken}:                        _GotoState21Action,
	{_State120, VarToken}:                         _GotoState37Action,
	{_State120, LetToken}:                         _GotoState27Action,
	{_State120, NotToken}:                         _GotoState30Action,
	{_State120, LabelDeclToken}:                   _GotoState25Action,
	{_State120, LparenToken}:                      _GotoState28Action,
	{_State120, LbracketToken}:                    _GotoState26Action,
	{_State120, DotDotDotToken}:                   _GotoState94Action,
	{_State120, SubToken}:                         _GotoState35Action,
	{_State120, MulToken}:                         _GotoState29Action,
	{_State120, BitNegToken}:                      _GotoState18Action,
	{_State120, BitAndToken}:                      _GotoState17Action,
	{_State120, GreaterToken}:                     _GotoState22Action,
	{_State120, ParseErrorToken}:                  _GotoState31Action,
	{_State120, VarDeclPatternType}:               _GotoState56Action,
	{_State120, VarOrLetType}:                     _GotoState57Action,
	{_State120, ExpressionType}:                   _GotoState99Action,
	{_State120, OptionalLabelDeclType}:            _GotoState50Action,
	{_State120, SequenceExprType}:                 _GotoState55Action,
	{_State120, CallExprType}:                     _GotoState43Action,
	{_State120, ArgumentsType}:                    _GotoState199Action,
	{_State120, ArgumentType}:                     _GotoState96Action,
	{_State120, ColonExpressionsType}:             _GotoState98Action,
	{_State120, OptionalExpressionType}:           _GotoState100Action,
	{_State120, AtomExprType}:                     _GotoState42Action,
	{_State120, LiteralType}:                      _GotoState48Action,
	{_State120, ImplicitStructExprType}:           _GotoState46Action,
	{_State120, AccessExprType}:                   _GotoState38Action,
	{_State120, PostfixUnaryExprType}:             _GotoState52Action,
	{_State120, PrefixUnaryOpType}:                _GotoState54Action,
	{_State120, PrefixUnaryExprType}:              _GotoState53Action,
	{_State120, MulExprType}:                      _GotoState49Action,
	{_State120, AddExprType}:                      _GotoState39Action,
	{_State120, CmpExprType}:                      _GotoState44Action,
	{_State120, AndExprType}:                      _GotoState40Action,
	{_State120, OrExprType}:                       _GotoState51Action,
	{_State120, InitializableTypeType}:            _GotoState47Action,
	{_State120, ExplicitStructDefType}:            _GotoState45Action,
	{_State120, AnonymousFuncExprType}:            _GotoState41Action,
	{_State127, IntegerLiteralToken}:              _GotoState24Action,
	{_State127, FloatLiteralToken}:                _GotoState20Action,
	{_State127, RuneLiteralToken}:                 _GotoState32Action,
	{_State127, StringLiteralToken}:               _GotoState33Action,
	{_State127, IdentifierToken}:                  _GotoState23Action,
	{_State127, TrueToken}:                        _GotoState36Action,
	{_State127, FalseToken}:                       _GotoState19Action,
	{_State127, StructToken}:                      _GotoState34Action,
	{_State127, FuncToken}:                        _GotoState21Action,
	{_State127, NotToken}:                         _GotoState30Action,
	{_State127, LabelDeclToken}:                   _GotoState25Action,
	{_State127, LparenToken}:                      _GotoState28Action,
	{_State127, LbracketToken}:                    _GotoState26Action,
	{_State127, SubToken}:                         _GotoState35Action,
	{_State127, MulToken}:                         _GotoState29Action,
	{_State127, BitNegToken}:                      _GotoState18Action,
	{_State127, BitAndToken}:                      _GotoState17Action,
	{_State127, ParseErrorToken}:                  _GotoState31Action,
	{_State127, OptionalLabelDeclType}:            _GotoState71Action,
	{_State127, CallExprType}:                     _GotoState43Action,
	{_State127, AtomExprType}:                     _GotoState42Action,
	{_State127, LiteralType}:                      _GotoState48Action,
	{_State127, ImplicitStructExprType}:           _GotoState46Action,
	{_State127, AccessExprType}:                   _GotoState38Action,
	{_State127, PostfixUnaryExprType}:             _GotoState52Action,
	{_State127, PrefixUnaryOpType}:                _GotoState54Action,
	{_State127, PrefixUnaryExprType}:              _GotoState200Action,
	{_State127, InitializableTypeType}:            _GotoState47Action,
	{_State127, ExplicitStructDefType}:            _GotoState45Action,
	{_State127, AnonymousFuncExprType}:            _GotoState41Action,
	{_State128, LbraceToken}:                      _GotoState58Action,
	{_State128, StatementBlockType}:               _GotoState201Action,
	{_State129, IntegerLiteralToken}:              _GotoState24Action,
	{_State129, FloatLiteralToken}:                _GotoState20Action,
	{_State129, RuneLiteralToken}:                 _GotoState32Action,
	{_State129, StringLiteralToken}:               _GotoState33Action,
	{_State129, IdentifierToken}:                  _GotoState23Action,
	{_State129, TrueToken}:                        _GotoState36Action,
	{_State129, FalseToken}:                       _GotoState19Action,
	{_State129, StructToken}:                      _GotoState34Action,
	{_State129, FuncToken}:                        _GotoState21Action,
	{_State129, VarToken}:                         _GotoState37Action,
	{_State129, LetToken}:                         _GotoState27Action,
	{_State129, NotToken}:                         _GotoState30Action,
	{_State129, LabelDeclToken}:                   _GotoState25Action,
	{_State129, LparenToken}:                      _GotoState28Action,
	{_State129, LbracketToken}:                    _GotoState26Action,
	{_State129, SubToken}:                         _GotoState35Action,
	{_State129, MulToken}:                         _GotoState29Action,
	{_State129, BitNegToken}:                      _GotoState18Action,
	{_State129, BitAndToken}:                      _GotoState17Action,
	{_State129, GreaterToken}:                     _GotoState22Action,
	{_State129, ParseErrorToken}:                  _GotoState31Action,
	{_State129, VarDeclPatternType}:               _GotoState56Action,
	{_State129, VarOrLetType}:                     _GotoState57Action,
	{_State129, AssignPatternType}:                _GotoState202Action,
	{_State129, OptionalLabelDeclType}:            _GotoState71Action,
	{_State129, SequenceExprType}:                 _GotoState204Action,
	{_State129, ForAssignmentType}:                _GotoState203Action,
	{_State129, CallExprType}:                     _GotoState43Action,
	{_State129, AtomExprType}:                     _GotoState42Action,
	{_State129, LiteralType}:                      _GotoState48Action,
	{_State129, ImplicitStructExprType}:           _GotoState46Action,
	{_State129, AccessExprType}:                   _GotoState38Action,
	{_State129, PostfixUnaryExprType}:             _GotoState52Action,
	{_State129, PrefixUnaryOpType}:                _GotoState54Action,
	{_State129, PrefixUnaryExprType}:              _GotoState53Action,
	{_State129, MulExprType}:                      _GotoState49Action,
	{_State129, AddExprType}:                      _GotoState39Action,
	{_State129, CmpExprType}:                      _GotoState44Action,
	{_State129, AndExprType}:                      _GotoState40Action,
	{_State129, OrExprType}:                       _GotoState51Action,
	{_State129, InitializableTypeType}:            _GotoState47Action,
	{_State129, ExplicitStructDefType}:            _GotoState45Action,
	{_State129, AnonymousFuncExprType}:            _GotoState41Action,
	{_State130, IntegerLiteralToken}:              _GotoState24Action,
	{_State130, FloatLiteralToken}:                _GotoState20Action,
	{_State130, RuneLiteralToken}:                 _GotoState32Action,
	{_State130, StringLiteralToken}:               _GotoState33Action,
	{_State130, IdentifierToken}:                  _GotoState23Action,
	{_State130, TrueToken}:                        _GotoState36Action,
	{_State130, FalseToken}:                       _GotoState19Action,
	{_State130, CaseToken}:                        _GotoState205Action,
	{_State130, StructToken}:                      _GotoState34Action,
	{_State130, FuncToken}:                        _GotoState21Action,
	{_State130, VarToken}:                         _GotoState37Action,
	{_State130, LetToken}:                         _GotoState27Action,
	{_State130, NotToken}:                         _GotoState30Action,
	{_State130, LabelDeclToken}:                   _GotoState25Action,
	{_State130, LparenToken}:                      _GotoState28Action,
	{_State130, LbracketToken}:                    _GotoState26Action,
	{_State130, SubToken}:                         _GotoState35Action,
	{_State130, MulToken}:                         _GotoState29Action,
	{_State130, BitNegToken}:                      _GotoState18Action,
	{_State130, BitAndToken}:                      _GotoState17Action,
	{_State130, GreaterToken}:                     _GotoState22Action,
	{_State130, ParseErrorToken}:                  _GotoState31Action,
	{_State130, VarDeclPatternType}:               _GotoState56Action,
	{_State130, VarOrLetType}:                     _GotoState57Action,
	{_State130, OptionalLabelDeclType}:            _GotoState71Action,
	{_State130, SequenceExprType}:                 _GotoState207Action,
	{_State130, ConditionType}:                    _GotoState206Action,
	{_State130, CallExprType}:                     _GotoState43Action,
	{_State130, AtomExprType}:                     _GotoState42Action,
	{_State130, LiteralType}:                      _GotoState48Action,
	{_State130, ImplicitStructExprType}:           _GotoState46Action,
	{_State130, AccessExprType}:                   _GotoState38Action,
	{_State130, PostfixUnaryExprType}:             _GotoState52Action,
	{_State130, PrefixUnaryOpType}:                _GotoState54Action,
	{_State130, PrefixUnaryExprType}:              _GotoState53Action,
	{_State130, MulExprType}:                      _GotoState49Action,
	{_State130, AddExprType}:                      _GotoState39Action,
	{_State130, CmpExprType}:                      _GotoState44Action,
	{_State130, AndExprType}:                      _GotoState40Action,
	{_State130, OrExprType}:                       _GotoState51Action,
	{_State130, InitializableTypeType}:            _GotoState47Action,
	{_State130, ExplicitStructDefType}:            _GotoState45Action,
	{_State130, AnonymousFuncExprType}:            _GotoState41Action,
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
	{_State131, OptionalLabelDeclType}:            _GotoState71Action,
	{_State131, SequenceExprType}:                 _GotoState208Action,
	{_State131, CallExprType}:                     _GotoState43Action,
	{_State131, AtomExprType}:                     _GotoState42Action,
	{_State131, LiteralType}:                      _GotoState48Action,
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
	{_State136, IntegerLiteralToken}:              _GotoState24Action,
	{_State136, FloatLiteralToken}:                _GotoState20Action,
	{_State136, RuneLiteralToken}:                 _GotoState32Action,
	{_State136, StringLiteralToken}:               _GotoState33Action,
	{_State136, IdentifierToken}:                  _GotoState23Action,
	{_State136, TrueToken}:                        _GotoState36Action,
	{_State136, FalseToken}:                       _GotoState19Action,
	{_State136, StructToken}:                      _GotoState34Action,
	{_State136, FuncToken}:                        _GotoState21Action,
	{_State136, NotToken}:                         _GotoState30Action,
	{_State136, LabelDeclToken}:                   _GotoState25Action,
	{_State136, LparenToken}:                      _GotoState28Action,
	{_State136, LbracketToken}:                    _GotoState26Action,
	{_State136, SubToken}:                         _GotoState35Action,
	{_State136, MulToken}:                         _GotoState29Action,
	{_State136, BitNegToken}:                      _GotoState18Action,
	{_State136, BitAndToken}:                      _GotoState17Action,
	{_State136, ParseErrorToken}:                  _GotoState31Action,
	{_State136, OptionalLabelDeclType}:            _GotoState71Action,
	{_State136, CallExprType}:                     _GotoState43Action,
	{_State136, AtomExprType}:                     _GotoState42Action,
	{_State136, LiteralType}:                      _GotoState48Action,
	{_State136, ImplicitStructExprType}:           _GotoState46Action,
	{_State136, AccessExprType}:                   _GotoState38Action,
	{_State136, PostfixUnaryExprType}:             _GotoState52Action,
	{_State136, PrefixUnaryOpType}:                _GotoState54Action,
	{_State136, PrefixUnaryExprType}:              _GotoState53Action,
	{_State136, MulExprType}:                      _GotoState49Action,
	{_State136, AddExprType}:                      _GotoState39Action,
	{_State136, CmpExprType}:                      _GotoState44Action,
	{_State136, AndExprType}:                      _GotoState209Action,
	{_State136, InitializableTypeType}:            _GotoState47Action,
	{_State136, ExplicitStructDefType}:            _GotoState45Action,
	{_State136, AnonymousFuncExprType}:            _GotoState41Action,
	{_State139, IdentifierToken}:                  _GotoState211Action,
	{_State139, LparenToken}:                      _GotoState139Action,
	{_State139, DotDotDotToken}:                   _GotoState210Action,
	{_State139, VarPatternType}:                   _GotoState214Action,
	{_State139, TuplePatternType}:                 _GotoState140Action,
	{_State139, FieldVarPatternsType}:             _GotoState213Action,
	{_State139, FieldVarPatternType}:              _GotoState212Action,
	{_State141, IdentifierToken}:                  _GotoState79Action,
	{_State141, StructToken}:                      _GotoState34Action,
	{_State141, EnumToken}:                        _GotoState76Action,
	{_State141, TraitToken}:                       _GotoState84Action,
	{_State141, FuncToken}:                        _GotoState78Action,
	{_State141, LparenToken}:                      _GotoState80Action,
	{_State141, LbracketToken}:                    _GotoState26Action,
	{_State141, DotToken}:                         _GotoState75Action,
	{_State141, QuestionToken}:                    _GotoState82Action,
	{_State141, ExclaimToken}:                     _GotoState77Action,
	{_State141, TildeTildeToken}:                  _GotoState83Action,
	{_State141, BitNegToken}:                      _GotoState74Action,
	{_State141, BitAndToken}:                      _GotoState73Action,
	{_State141, ParseErrorToken}:                  _GotoState81Action,
	{_State141, OptionalValueTypeType}:            _GotoState215Action,
	{_State141, InitializableTypeType}:            _GotoState90Action,
	{_State141, AtomTypeType}:                     _GotoState85Action,
	{_State141, ReturnableTypeType}:               _GotoState91Action,
	{_State141, ValueTypeType}:                    _GotoState216Action,
	{_State141, ImplicitStructDefType}:            _GotoState89Action,
	{_State141, ExplicitStructDefType}:            _GotoState45Action,
	{_State141, ImplicitEnumDefType}:              _GotoState88Action,
	{_State141, ExplicitEnumDefType}:              _GotoState86Action,
	{_State141, TraitDefType}:                     _GotoState92Action,
	{_State141, FuncTypeType}:                     _GotoState87Action,
	{_State142, IntegerLiteralToken}:              _GotoState24Action,
	{_State142, FloatLiteralToken}:                _GotoState20Action,
	{_State142, RuneLiteralToken}:                 _GotoState32Action,
	{_State142, StringLiteralToken}:               _GotoState33Action,
	{_State142, IdentifierToken}:                  _GotoState23Action,
	{_State142, TrueToken}:                        _GotoState36Action,
	{_State142, FalseToken}:                       _GotoState19Action,
	{_State142, CaseToken}:                        _GotoState219Action,
	{_State142, DefaultToken}:                     _GotoState221Action,
	{_State142, ReturnToken}:                      _GotoState226Action,
	{_State142, BreakToken}:                       _GotoState218Action,
	{_State142, ContinueToken}:                    _GotoState220Action,
	{_State142, FallthroughToken}:                 _GotoState223Action,
	{_State142, ImportToken}:                      _GotoState224Action,
	{_State142, UnsafeToken}:                      _GotoState165Action,
	{_State142, StructToken}:                      _GotoState34Action,
	{_State142, FuncToken}:                        _GotoState21Action,
	{_State142, AsyncToken}:                       _GotoState217Action,
	{_State142, DeferToken}:                       _GotoState222Action,
	{_State142, VarToken}:                         _GotoState37Action,
	{_State142, LetToken}:                         _GotoState27Action,
	{_State142, NotToken}:                         _GotoState30Action,
	{_State142, LabelDeclToken}:                   _GotoState25Action,
	{_State142, RbraceToken}:                      _GotoState225Action,
	{_State142, LparenToken}:                      _GotoState28Action,
	{_State142, LbracketToken}:                    _GotoState26Action,
	{_State142, SubToken}:                         _GotoState35Action,
	{_State142, MulToken}:                         _GotoState29Action,
	{_State142, BitNegToken}:                      _GotoState18Action,
	{_State142, BitAndToken}:                      _GotoState17Action,
	{_State142, GreaterToken}:                     _GotoState22Action,
	{_State142, ParseErrorToken}:                  _GotoState31Action,
	{_State142, StatementType}:                    _GotoState235Action,
	{_State142, SimpleStatementBodyType}:          _GotoState234Action,
	{_State142, StatementBodyType}:                _GotoState236Action,
	{_State142, UnsafeStatementType}:              _GotoState237Action,
	{_State142, JumpTypeType}:                     _GotoState232Action,
	{_State142, ExpressionsType}:                  _GotoState230Action,
	{_State142, ImportStatementType}:              _GotoState231Action,
	{_State142, VarDeclPatternType}:               _GotoState56Action,
	{_State142, VarOrLetType}:                     _GotoState57Action,
	{_State142, AssignPatternType}:                _GotoState228Action,
	{_State142, ExpressionType}:                   _GotoState229Action,
	{_State142, OptionalLabelDeclType}:            _GotoState50Action,
	{_State142, SequenceExprType}:                 _GotoState233Action,
	{_State142, CallExprType}:                     _GotoState43Action,
	{_State142, AtomExprType}:                     _GotoState42Action,
	{_State142, LiteralType}:                      _GotoState48Action,
	{_State142, ImplicitStructExprType}:           _GotoState46Action,
	{_State142, AccessExprType}:                   _GotoState227Action,
	{_State142, PostfixUnaryExprType}:             _GotoState52Action,
	{_State142, PrefixUnaryOpType}:                _GotoState54Action,
	{_State142, PrefixUnaryExprType}:              _GotoState53Action,
	{_State142, MulExprType}:                      _GotoState49Action,
	{_State142, AddExprType}:                      _GotoState39Action,
	{_State142, CmpExprType}:                      _GotoState44Action,
	{_State142, AndExprType}:                      _GotoState40Action,
	{_State142, OrExprType}:                       _GotoState51Action,
	{_State142, InitializableTypeType}:            _GotoState47Action,
	{_State142, ExplicitStructDefType}:            _GotoState45Action,
	{_State142, AnonymousFuncExprType}:            _GotoState41Action,
	{_State143, PackageToken}:                     _GotoState14Action,
	{_State143, TypeToken}:                        _GotoState15Action,
	{_State143, FuncToken}:                        _GotoState16Action,
	{_State143, VarToken}:                         _GotoState37Action,
	{_State143, LetToken}:                         _GotoState27Action,
	{_State143, LbraceToken}:                      _GotoState58Action,
	{_State143, DefinitionType}:                   _GotoState238Action,
	{_State143, StatementBlockType}:               _GotoState63Action,
	{_State143, VarDeclPatternType}:               _GotoState65Action,
	{_State143, VarOrLetType}:                     _GotoState57Action,
	{_State143, TypeDefType}:                      _GotoState64Action,
	{_State143, NamedFuncDefType}:                 _GotoState61Action,
	{_State143, PackageDefType}:                   _GotoState62Action,
	{_State145, IntegerLiteralToken}:              _GotoState24Action,
	{_State145, FloatLiteralToken}:                _GotoState20Action,
	{_State145, RuneLiteralToken}:                 _GotoState32Action,
	{_State145, StringLiteralToken}:               _GotoState33Action,
	{_State145, IdentifierToken}:                  _GotoState23Action,
	{_State145, TrueToken}:                        _GotoState36Action,
	{_State145, FalseToken}:                       _GotoState19Action,
	{_State145, StructToken}:                      _GotoState34Action,
	{_State145, FuncToken}:                        _GotoState21Action,
	{_State145, VarToken}:                         _GotoState37Action,
	{_State145, LetToken}:                         _GotoState27Action,
	{_State145, NotToken}:                         _GotoState30Action,
	{_State145, LabelDeclToken}:                   _GotoState25Action,
	{_State145, LparenToken}:                      _GotoState28Action,
	{_State145, LbracketToken}:                    _GotoState26Action,
	{_State145, SubToken}:                         _GotoState35Action,
	{_State145, MulToken}:                         _GotoState29Action,
	{_State145, BitNegToken}:                      _GotoState18Action,
	{_State145, BitAndToken}:                      _GotoState17Action,
	{_State145, GreaterToken}:                     _GotoState22Action,
	{_State145, ParseErrorToken}:                  _GotoState31Action,
	{_State145, VarDeclPatternType}:               _GotoState56Action,
	{_State145, VarOrLetType}:                     _GotoState57Action,
	{_State145, ExpressionType}:                   _GotoState239Action,
	{_State145, OptionalLabelDeclType}:            _GotoState50Action,
	{_State145, SequenceExprType}:                 _GotoState55Action,
	{_State145, CallExprType}:                     _GotoState43Action,
	{_State145, AtomExprType}:                     _GotoState42Action,
	{_State145, LiteralType}:                      _GotoState48Action,
	{_State145, ImplicitStructExprType}:           _GotoState46Action,
	{_State145, AccessExprType}:                   _GotoState38Action,
	{_State145, PostfixUnaryExprType}:             _GotoState52Action,
	{_State145, PrefixUnaryOpType}:                _GotoState54Action,
	{_State145, PrefixUnaryExprType}:              _GotoState53Action,
	{_State145, MulExprType}:                      _GotoState49Action,
	{_State145, AddExprType}:                      _GotoState39Action,
	{_State145, CmpExprType}:                      _GotoState44Action,
	{_State145, AndExprType}:                      _GotoState40Action,
	{_State145, OrExprType}:                       _GotoState51Action,
	{_State145, InitializableTypeType}:            _GotoState47Action,
	{_State145, ExplicitStructDefType}:            _GotoState45Action,
	{_State145, AnonymousFuncExprType}:            _GotoState41Action,
	{_State146, IdentifierToken}:                  _GotoState79Action,
	{_State146, StructToken}:                      _GotoState34Action,
	{_State146, EnumToken}:                        _GotoState76Action,
	{_State146, TraitToken}:                       _GotoState84Action,
	{_State146, FuncToken}:                        _GotoState78Action,
	{_State146, LparenToken}:                      _GotoState80Action,
	{_State146, LbracketToken}:                    _GotoState26Action,
	{_State146, DotToken}:                         _GotoState75Action,
	{_State146, QuestionToken}:                    _GotoState82Action,
	{_State146, ExclaimToken}:                     _GotoState77Action,
	{_State146, TildeTildeToken}:                  _GotoState83Action,
	{_State146, BitNegToken}:                      _GotoState74Action,
	{_State146, BitAndToken}:                      _GotoState73Action,
	{_State146, ParseErrorToken}:                  _GotoState81Action,
	{_State146, InitializableTypeType}:            _GotoState90Action,
	{_State146, AtomTypeType}:                     _GotoState85Action,
	{_State146, ReturnableTypeType}:               _GotoState91Action,
	{_State146, ValueTypeType}:                    _GotoState240Action,
	{_State146, ImplicitStructDefType}:            _GotoState89Action,
	{_State146, ExplicitStructDefType}:            _GotoState45Action,
	{_State146, ImplicitEnumDefType}:              _GotoState88Action,
	{_State146, ExplicitEnumDefType}:              _GotoState86Action,
	{_State146, TraitDefType}:                     _GotoState92Action,
	{_State146, FuncTypeType}:                     _GotoState87Action,
	{_State147, IdentifierToken}:                  _GotoState241Action,
	{_State147, GenericParameterDefType}:          _GotoState242Action,
	{_State147, GenericParameterDefsType}:         _GotoState243Action,
	{_State147, OptionalGenericParameterDefsType}: _GotoState244Action,
	{_State148, IdentifierToken}:                  _GotoState79Action,
	{_State148, StructToken}:                      _GotoState34Action,
	{_State148, EnumToken}:                        _GotoState76Action,
	{_State148, TraitToken}:                       _GotoState84Action,
	{_State148, FuncToken}:                        _GotoState78Action,
	{_State148, LparenToken}:                      _GotoState80Action,
	{_State148, LbracketToken}:                    _GotoState26Action,
	{_State148, DotToken}:                         _GotoState75Action,
	{_State148, QuestionToken}:                    _GotoState82Action,
	{_State148, ExclaimToken}:                     _GotoState77Action,
	{_State148, TildeTildeToken}:                  _GotoState83Action,
	{_State148, BitNegToken}:                      _GotoState74Action,
	{_State148, BitAndToken}:                      _GotoState73Action,
	{_State148, ParseErrorToken}:                  _GotoState81Action,
	{_State148, InitializableTypeType}:            _GotoState90Action,
	{_State148, AtomTypeType}:                     _GotoState85Action,
	{_State148, ReturnableTypeType}:               _GotoState91Action,
	{_State148, ValueTypeType}:                    _GotoState245Action,
	{_State148, ImplicitStructDefType}:            _GotoState89Action,
	{_State148, ExplicitStructDefType}:            _GotoState45Action,
	{_State148, ImplicitEnumDefType}:              _GotoState88Action,
	{_State148, ExplicitEnumDefType}:              _GotoState86Action,
	{_State148, TraitDefType}:                     _GotoState92Action,
	{_State148, FuncTypeType}:                     _GotoState87Action,
	{_State149, IntegerLiteralToken}:              _GotoState24Action,
	{_State149, FloatLiteralToken}:                _GotoState20Action,
	{_State149, RuneLiteralToken}:                 _GotoState32Action,
	{_State149, StringLiteralToken}:               _GotoState33Action,
	{_State149, IdentifierToken}:                  _GotoState23Action,
	{_State149, TrueToken}:                        _GotoState36Action,
	{_State149, FalseToken}:                       _GotoState19Action,
	{_State149, StructToken}:                      _GotoState34Action,
	{_State149, FuncToken}:                        _GotoState21Action,
	{_State149, VarToken}:                         _GotoState37Action,
	{_State149, LetToken}:                         _GotoState27Action,
	{_State149, NotToken}:                         _GotoState30Action,
	{_State149, LabelDeclToken}:                   _GotoState25Action,
	{_State149, LparenToken}:                      _GotoState28Action,
	{_State149, LbracketToken}:                    _GotoState26Action,
	{_State149, SubToken}:                         _GotoState35Action,
	{_State149, MulToken}:                         _GotoState29Action,
	{_State149, BitNegToken}:                      _GotoState18Action,
	{_State149, BitAndToken}:                      _GotoState17Action,
	{_State149, GreaterToken}:                     _GotoState22Action,
	{_State149, ParseErrorToken}:                  _GotoState31Action,
	{_State149, VarDeclPatternType}:               _GotoState56Action,
	{_State149, VarOrLetType}:                     _GotoState57Action,
	{_State149, ExpressionType}:                   _GotoState246Action,
	{_State149, OptionalLabelDeclType}:            _GotoState50Action,
	{_State149, SequenceExprType}:                 _GotoState55Action,
	{_State149, CallExprType}:                     _GotoState43Action,
	{_State149, AtomExprType}:                     _GotoState42Action,
	{_State149, LiteralType}:                      _GotoState48Action,
	{_State149, ImplicitStructExprType}:           _GotoState46Action,
	{_State149, AccessExprType}:                   _GotoState38Action,
	{_State149, PostfixUnaryExprType}:             _GotoState52Action,
	{_State149, PrefixUnaryOpType}:                _GotoState54Action,
	{_State149, PrefixUnaryExprType}:              _GotoState53Action,
	{_State149, MulExprType}:                      _GotoState49Action,
	{_State149, AddExprType}:                      _GotoState39Action,
	{_State149, CmpExprType}:                      _GotoState44Action,
	{_State149, AndExprType}:                      _GotoState40Action,
	{_State149, OrExprType}:                       _GotoState51Action,
	{_State149, InitializableTypeType}:            _GotoState47Action,
	{_State149, ExplicitStructDefType}:            _GotoState45Action,
	{_State149, AnonymousFuncExprType}:            _GotoState41Action,
	{_State150, LparenToken}:                      _GotoState247Action,
	{_State151, IdentifierToken}:                  _GotoState79Action,
	{_State151, StructToken}:                      _GotoState34Action,
	{_State151, EnumToken}:                        _GotoState76Action,
	{_State151, TraitToken}:                       _GotoState84Action,
	{_State151, FuncToken}:                        _GotoState78Action,
	{_State151, LparenToken}:                      _GotoState80Action,
	{_State151, LbracketToken}:                    _GotoState26Action,
	{_State151, DotToken}:                         _GotoState75Action,
	{_State151, QuestionToken}:                    _GotoState82Action,
	{_State151, ExclaimToken}:                     _GotoState77Action,
	{_State151, DotDotDotToken}:                   _GotoState248Action,
	{_State151, TildeTildeToken}:                  _GotoState83Action,
	{_State151, BitNegToken}:                      _GotoState74Action,
	{_State151, BitAndToken}:                      _GotoState73Action,
	{_State151, ParseErrorToken}:                  _GotoState81Action,
	{_State151, InitializableTypeType}:            _GotoState90Action,
	{_State151, AtomTypeType}:                     _GotoState85Action,
	{_State151, ReturnableTypeType}:               _GotoState91Action,
	{_State151, ValueTypeType}:                    _GotoState249Action,
	{_State151, ImplicitStructDefType}:            _GotoState89Action,
	{_State151, ExplicitStructDefType}:            _GotoState45Action,
	{_State151, ImplicitEnumDefType}:              _GotoState88Action,
	{_State151, ExplicitEnumDefType}:              _GotoState86Action,
	{_State151, TraitDefType}:                     _GotoState92Action,
	{_State151, FuncTypeType}:                     _GotoState87Action,
	{_State152, RparenToken}:                      _GotoState250Action,
	{_State153, RparenToken}:                      _GotoState251Action,
	{_State155, CommaToken}:                       _GotoState252Action,
	{_State159, IdentifierToken}:                  _GotoState164Action,
	{_State159, UnsafeToken}:                      _GotoState165Action,
	{_State159, StructToken}:                      _GotoState34Action,
	{_State159, EnumToken}:                        _GotoState76Action,
	{_State159, TraitToken}:                       _GotoState84Action,
	{_State159, FuncToken}:                        _GotoState78Action,
	{_State159, LparenToken}:                      _GotoState80Action,
	{_State159, LbracketToken}:                    _GotoState26Action,
	{_State159, DotToken}:                         _GotoState75Action,
	{_State159, QuestionToken}:                    _GotoState82Action,
	{_State159, ExclaimToken}:                     _GotoState77Action,
	{_State159, TildeTildeToken}:                  _GotoState83Action,
	{_State159, BitNegToken}:                      _GotoState74Action,
	{_State159, BitAndToken}:                      _GotoState73Action,
	{_State159, ParseErrorToken}:                  _GotoState81Action,
	{_State159, UnsafeStatementType}:              _GotoState171Action,
	{_State159, InitializableTypeType}:            _GotoState90Action,
	{_State159, AtomTypeType}:                     _GotoState85Action,
	{_State159, ReturnableTypeType}:               _GotoState91Action,
	{_State159, ValueTypeType}:                    _GotoState172Action,
	{_State159, FieldDefType}:                     _GotoState255Action,
	{_State159, ImplicitStructDefType}:            _GotoState89Action,
	{_State159, ExplicitStructDefType}:            _GotoState45Action,
	{_State159, EnumValueDefType}:                 _GotoState253Action,
	{_State159, ImplicitEnumValueDefsType}:        _GotoState256Action,
	{_State159, ImplicitEnumDefType}:              _GotoState88Action,
	{_State159, ExplicitEnumValueDefsType}:        _GotoState254Action,
	{_State159, ExplicitEnumDefType}:              _GotoState86Action,
	{_State159, TraitDefType}:                     _GotoState92Action,
	{_State159, FuncTypeType}:                     _GotoState87Action,
	{_State161, IdentifierToken}:                  _GotoState258Action,
	{_State161, StructToken}:                      _GotoState34Action,
	{_State161, EnumToken}:                        _GotoState76Action,
	{_State161, TraitToken}:                       _GotoState84Action,
	{_State161, FuncToken}:                        _GotoState78Action,
	{_State161, LparenToken}:                      _GotoState80Action,
	{_State161, LbracketToken}:                    _GotoState26Action,
	{_State161, DotToken}:                         _GotoState75Action,
	{_State161, QuestionToken}:                    _GotoState82Action,
	{_State161, ExclaimToken}:                     _GotoState77Action,
	{_State161, DotDotDotToken}:                   _GotoState257Action,
	{_State161, TildeTildeToken}:                  _GotoState83Action,
	{_State161, BitNegToken}:                      _GotoState74Action,
	{_State161, BitAndToken}:                      _GotoState73Action,
	{_State161, ParseErrorToken}:                  _GotoState81Action,
	{_State161, InitializableTypeType}:            _GotoState90Action,
	{_State161, AtomTypeType}:                     _GotoState85Action,
	{_State161, ReturnableTypeType}:               _GotoState91Action,
	{_State161, ValueTypeType}:                    _GotoState262Action,
	{_State161, ImplicitStructDefType}:            _GotoState89Action,
	{_State161, ExplicitStructDefType}:            _GotoState45Action,
	{_State161, ImplicitEnumDefType}:              _GotoState88Action,
	{_State161, ExplicitEnumDefType}:              _GotoState86Action,
	{_State161, TraitDefType}:                     _GotoState92Action,
	{_State161, ParameterDeclType}:                _GotoState260Action,
	{_State161, ParameterDeclsType}:               _GotoState261Action,
	{_State161, OptionalParameterDeclsType}:       _GotoState259Action,
	{_State161, FuncTypeType}:                     _GotoState87Action,
	{_State162, IdentifierToken}:                  _GotoState263Action,
	{_State164, IdentifierToken}:                  _GotoState79Action,
	{_State164, StructToken}:                      _GotoState34Action,
	{_State164, EnumToken}:                        _GotoState76Action,
	{_State164, TraitToken}:                       _GotoState84Action,
	{_State164, FuncToken}:                        _GotoState78Action,
	{_State164, LparenToken}:                      _GotoState80Action,
	{_State164, LbracketToken}:                    _GotoState26Action,
	{_State164, DotToken}:                         _GotoState264Action,
	{_State164, QuestionToken}:                    _GotoState82Action,
	{_State164, ExclaimToken}:                     _GotoState77Action,
	{_State164, DollarLbracketToken}:              _GotoState102Action,
	{_State164, TildeTildeToken}:                  _GotoState83Action,
	{_State164, BitNegToken}:                      _GotoState74Action,
	{_State164, BitAndToken}:                      _GotoState73Action,
	{_State164, ParseErrorToken}:                  _GotoState81Action,
	{_State164, OptionalGenericBindingType}:       _GotoState163Action,
	{_State164, InitializableTypeType}:            _GotoState90Action,
	{_State164, AtomTypeType}:                     _GotoState85Action,
	{_State164, ReturnableTypeType}:               _GotoState91Action,
	{_State164, ValueTypeType}:                    _GotoState265Action,
	{_State164, ImplicitStructDefType}:            _GotoState89Action,
	{_State164, ExplicitStructDefType}:            _GotoState45Action,
	{_State164, ImplicitEnumDefType}:              _GotoState88Action,
	{_State164, ExplicitEnumDefType}:              _GotoState86Action,
	{_State164, TraitDefType}:                     _GotoState92Action,
	{_State164, FuncTypeType}:                     _GotoState87Action,
	{_State165, LessToken}:                        _GotoState266Action,
	{_State166, OrToken}:                          _GotoState267Action,
	{_State167, AssignToken}:                      _GotoState268Action,
	{_State168, OrToken}:                          _GotoState269Action,
	{_State168, RparenToken}:                      _GotoState270Action,
	{_State169, CommaToken}:                       _GotoState271Action,
	{_State170, RparenToken}:                      _GotoState272Action,
	{_State172, AddToken}:                         _GotoState176Action,
	{_State172, SubToken}:                         _GotoState181Action,
	{_State172, MulToken}:                         _GotoState179Action,
	{_State175, IdentifierToken}:                  _GotoState164Action,
	{_State175, UnsafeToken}:                      _GotoState165Action,
	{_State175, StructToken}:                      _GotoState34Action,
	{_State175, EnumToken}:                        _GotoState76Action,
	{_State175, TraitToken}:                       _GotoState84Action,
	{_State175, FuncToken}:                        _GotoState273Action,
	{_State175, LparenToken}:                      _GotoState80Action,
	{_State175, LbracketToken}:                    _GotoState26Action,
	{_State175, DotToken}:                         _GotoState75Action,
	{_State175, QuestionToken}:                    _GotoState82Action,
	{_State175, ExclaimToken}:                     _GotoState77Action,
	{_State175, TildeTildeToken}:                  _GotoState83Action,
	{_State175, BitNegToken}:                      _GotoState74Action,
	{_State175, BitAndToken}:                      _GotoState73Action,
	{_State175, ParseErrorToken}:                  _GotoState81Action,
	{_State175, UnsafeStatementType}:              _GotoState171Action,
	{_State175, InitializableTypeType}:            _GotoState90Action,
	{_State175, AtomTypeType}:                     _GotoState85Action,
	{_State175, ReturnableTypeType}:               _GotoState91Action,
	{_State175, ValueTypeType}:                    _GotoState172Action,
	{_State175, FieldDefType}:                     _GotoState274Action,
	{_State175, ImplicitStructDefType}:            _GotoState89Action,
	{_State175, ExplicitStructDefType}:            _GotoState45Action,
	{_State175, ImplicitEnumDefType}:              _GotoState88Action,
	{_State175, ExplicitEnumDefType}:              _GotoState86Action,
	{_State175, TraitPropertyType}:                _GotoState278Action,
	{_State175, TraitPropertiesType}:              _GotoState277Action,
	{_State175, OptionalTraitPropertiesType}:      _GotoState276Action,
	{_State175, TraitDefType}:                     _GotoState92Action,
	{_State175, FuncTypeType}:                     _GotoState87Action,
	{_State175, MethodSignatureType}:              _GotoState275Action,
	{_State176, IdentifierToken}:                  _GotoState79Action,
	{_State176, StructToken}:                      _GotoState34Action,
	{_State176, EnumToken}:                        _GotoState76Action,
	{_State176, TraitToken}:                       _GotoState84Action,
	{_State176, FuncToken}:                        _GotoState78Action,
	{_State176, LparenToken}:                      _GotoState80Action,
	{_State176, LbracketToken}:                    _GotoState26Action,
	{_State176, DotToken}:                         _GotoState75Action,
	{_State176, QuestionToken}:                    _GotoState82Action,
	{_State176, ExclaimToken}:                     _GotoState77Action,
	{_State176, TildeTildeToken}:                  _GotoState83Action,
	{_State176, BitNegToken}:                      _GotoState74Action,
	{_State176, BitAndToken}:                      _GotoState73Action,
	{_State176, ParseErrorToken}:                  _GotoState81Action,
	{_State176, InitializableTypeType}:            _GotoState90Action,
	{_State176, AtomTypeType}:                     _GotoState85Action,
	{_State176, ReturnableTypeType}:               _GotoState279Action,
	{_State176, ImplicitStructDefType}:            _GotoState89Action,
	{_State176, ExplicitStructDefType}:            _GotoState45Action,
	{_State176, ImplicitEnumDefType}:              _GotoState88Action,
	{_State176, ExplicitEnumDefType}:              _GotoState86Action,
	{_State176, TraitDefType}:                     _GotoState92Action,
	{_State176, FuncTypeType}:                     _GotoState87Action,
	{_State177, IdentifierToken}:                  _GotoState79Action,
	{_State177, StructToken}:                      _GotoState34Action,
	{_State177, EnumToken}:                        _GotoState76Action,
	{_State177, TraitToken}:                       _GotoState84Action,
	{_State177, FuncToken}:                        _GotoState78Action,
	{_State177, LparenToken}:                      _GotoState80Action,
	{_State177, LbracketToken}:                    _GotoState26Action,
	{_State177, DotToken}:                         _GotoState75Action,
	{_State177, QuestionToken}:                    _GotoState82Action,
	{_State177, ExclaimToken}:                     _GotoState77Action,
	{_State177, TildeTildeToken}:                  _GotoState83Action,
	{_State177, BitNegToken}:                      _GotoState74Action,
	{_State177, BitAndToken}:                      _GotoState73Action,
	{_State177, ParseErrorToken}:                  _GotoState81Action,
	{_State177, InitializableTypeType}:            _GotoState90Action,
	{_State177, AtomTypeType}:                     _GotoState85Action,
	{_State177, ReturnableTypeType}:               _GotoState91Action,
	{_State177, ValueTypeType}:                    _GotoState280Action,
	{_State177, ImplicitStructDefType}:            _GotoState89Action,
	{_State177, ExplicitStructDefType}:            _GotoState45Action,
	{_State177, ImplicitEnumDefType}:              _GotoState88Action,
	{_State177, ExplicitEnumDefType}:              _GotoState86Action,
	{_State177, TraitDefType}:                     _GotoState92Action,
	{_State177, FuncTypeType}:                     _GotoState87Action,
	{_State178, IntegerLiteralToken}:              _GotoState281Action,
	{_State179, IdentifierToken}:                  _GotoState79Action,
	{_State179, StructToken}:                      _GotoState34Action,
	{_State179, EnumToken}:                        _GotoState76Action,
	{_State179, TraitToken}:                       _GotoState84Action,
	{_State179, FuncToken}:                        _GotoState78Action,
	{_State179, LparenToken}:                      _GotoState80Action,
	{_State179, LbracketToken}:                    _GotoState26Action,
	{_State179, DotToken}:                         _GotoState75Action,
	{_State179, QuestionToken}:                    _GotoState82Action,
	{_State179, ExclaimToken}:                     _GotoState77Action,
	{_State179, TildeTildeToken}:                  _GotoState83Action,
	{_State179, BitNegToken}:                      _GotoState74Action,
	{_State179, BitAndToken}:                      _GotoState73Action,
	{_State179, ParseErrorToken}:                  _GotoState81Action,
	{_State179, InitializableTypeType}:            _GotoState90Action,
	{_State179, AtomTypeType}:                     _GotoState85Action,
	{_State179, ReturnableTypeType}:               _GotoState282Action,
	{_State179, ImplicitStructDefType}:            _GotoState89Action,
	{_State179, ExplicitStructDefType}:            _GotoState45Action,
	{_State179, ImplicitEnumDefType}:              _GotoState88Action,
	{_State179, ExplicitEnumDefType}:              _GotoState86Action,
	{_State179, TraitDefType}:                     _GotoState92Action,
	{_State179, FuncTypeType}:                     _GotoState87Action,
	{_State181, IdentifierToken}:                  _GotoState79Action,
	{_State181, StructToken}:                      _GotoState34Action,
	{_State181, EnumToken}:                        _GotoState76Action,
	{_State181, TraitToken}:                       _GotoState84Action,
	{_State181, FuncToken}:                        _GotoState78Action,
	{_State181, LparenToken}:                      _GotoState80Action,
	{_State181, LbracketToken}:                    _GotoState26Action,
	{_State181, DotToken}:                         _GotoState75Action,
	{_State181, QuestionToken}:                    _GotoState82Action,
	{_State181, ExclaimToken}:                     _GotoState77Action,
	{_State181, TildeTildeToken}:                  _GotoState83Action,
	{_State181, BitNegToken}:                      _GotoState74Action,
	{_State181, BitAndToken}:                      _GotoState73Action,
	{_State181, ParseErrorToken}:                  _GotoState81Action,
	{_State181, InitializableTypeType}:            _GotoState90Action,
	{_State181, AtomTypeType}:                     _GotoState85Action,
	{_State181, ReturnableTypeType}:               _GotoState283Action,
	{_State181, ImplicitStructDefType}:            _GotoState89Action,
	{_State181, ExplicitStructDefType}:            _GotoState45Action,
	{_State181, ImplicitEnumDefType}:              _GotoState88Action,
	{_State181, ExplicitEnumDefType}:              _GotoState86Action,
	{_State181, TraitDefType}:                     _GotoState92Action,
	{_State181, FuncTypeType}:                     _GotoState87Action,
	{_State182, IntegerLiteralToken}:              _GotoState24Action,
	{_State182, FloatLiteralToken}:                _GotoState20Action,
	{_State182, RuneLiteralToken}:                 _GotoState32Action,
	{_State182, StringLiteralToken}:               _GotoState33Action,
	{_State182, IdentifierToken}:                  _GotoState23Action,
	{_State182, TrueToken}:                        _GotoState36Action,
	{_State182, FalseToken}:                       _GotoState19Action,
	{_State182, StructToken}:                      _GotoState34Action,
	{_State182, FuncToken}:                        _GotoState21Action,
	{_State182, VarToken}:                         _GotoState37Action,
	{_State182, LetToken}:                         _GotoState27Action,
	{_State182, NotToken}:                         _GotoState30Action,
	{_State182, LabelDeclToken}:                   _GotoState25Action,
	{_State182, LparenToken}:                      _GotoState28Action,
	{_State182, LbracketToken}:                    _GotoState26Action,
	{_State182, SubToken}:                         _GotoState35Action,
	{_State182, MulToken}:                         _GotoState29Action,
	{_State182, BitNegToken}:                      _GotoState18Action,
	{_State182, BitAndToken}:                      _GotoState17Action,
	{_State182, GreaterToken}:                     _GotoState22Action,
	{_State182, ParseErrorToken}:                  _GotoState31Action,
	{_State182, VarDeclPatternType}:               _GotoState56Action,
	{_State182, VarOrLetType}:                     _GotoState57Action,
	{_State182, ExpressionType}:                   _GotoState284Action,
	{_State182, OptionalLabelDeclType}:            _GotoState50Action,
	{_State182, SequenceExprType}:                 _GotoState55Action,
	{_State182, CallExprType}:                     _GotoState43Action,
	{_State182, AtomExprType}:                     _GotoState42Action,
	{_State182, LiteralType}:                      _GotoState48Action,
	{_State182, ImplicitStructExprType}:           _GotoState46Action,
	{_State182, AccessExprType}:                   _GotoState38Action,
	{_State182, PostfixUnaryExprType}:             _GotoState52Action,
	{_State182, PrefixUnaryOpType}:                _GotoState54Action,
	{_State182, PrefixUnaryExprType}:              _GotoState53Action,
	{_State182, MulExprType}:                      _GotoState49Action,
	{_State182, AddExprType}:                      _GotoState39Action,
	{_State182, CmpExprType}:                      _GotoState44Action,
	{_State182, AndExprType}:                      _GotoState40Action,
	{_State182, OrExprType}:                       _GotoState51Action,
	{_State182, InitializableTypeType}:            _GotoState47Action,
	{_State182, ExplicitStructDefType}:            _GotoState45Action,
	{_State182, AnonymousFuncExprType}:            _GotoState41Action,
	{_State183, IntegerLiteralToken}:              _GotoState24Action,
	{_State183, FloatLiteralToken}:                _GotoState20Action,
	{_State183, RuneLiteralToken}:                 _GotoState32Action,
	{_State183, StringLiteralToken}:               _GotoState33Action,
	{_State183, IdentifierToken}:                  _GotoState95Action,
	{_State183, TrueToken}:                        _GotoState36Action,
	{_State183, FalseToken}:                       _GotoState19Action,
	{_State183, StructToken}:                      _GotoState34Action,
	{_State183, FuncToken}:                        _GotoState21Action,
	{_State183, VarToken}:                         _GotoState37Action,
	{_State183, LetToken}:                         _GotoState27Action,
	{_State183, NotToken}:                         _GotoState30Action,
	{_State183, LabelDeclToken}:                   _GotoState25Action,
	{_State183, LparenToken}:                      _GotoState28Action,
	{_State183, LbracketToken}:                    _GotoState26Action,
	{_State183, DotDotDotToken}:                   _GotoState94Action,
	{_State183, SubToken}:                         _GotoState35Action,
	{_State183, MulToken}:                         _GotoState29Action,
	{_State183, BitNegToken}:                      _GotoState18Action,
	{_State183, BitAndToken}:                      _GotoState17Action,
	{_State183, GreaterToken}:                     _GotoState22Action,
	{_State183, ParseErrorToken}:                  _GotoState31Action,
	{_State183, VarDeclPatternType}:               _GotoState56Action,
	{_State183, VarOrLetType}:                     _GotoState57Action,
	{_State183, ExpressionType}:                   _GotoState99Action,
	{_State183, OptionalLabelDeclType}:            _GotoState50Action,
	{_State183, SequenceExprType}:                 _GotoState55Action,
	{_State183, CallExprType}:                     _GotoState43Action,
	{_State183, ArgumentType}:                     _GotoState285Action,
	{_State183, ColonExpressionsType}:             _GotoState98Action,
	{_State183, OptionalExpressionType}:           _GotoState100Action,
	{_State183, AtomExprType}:                     _GotoState42Action,
	{_State183, LiteralType}:                      _GotoState48Action,
	{_State183, ImplicitStructExprType}:           _GotoState46Action,
	{_State183, AccessExprType}:                   _GotoState38Action,
	{_State183, PostfixUnaryExprType}:             _GotoState52Action,
	{_State183, PrefixUnaryOpType}:                _GotoState54Action,
	{_State183, PrefixUnaryExprType}:              _GotoState53Action,
	{_State183, MulExprType}:                      _GotoState49Action,
	{_State183, AddExprType}:                      _GotoState39Action,
	{_State183, CmpExprType}:                      _GotoState44Action,
	{_State183, AndExprType}:                      _GotoState40Action,
	{_State183, OrExprType}:                       _GotoState51Action,
	{_State183, InitializableTypeType}:            _GotoState47Action,
	{_State183, ExplicitStructDefType}:            _GotoState45Action,
	{_State183, AnonymousFuncExprType}:            _GotoState41Action,
	{_State185, IntegerLiteralToken}:              _GotoState24Action,
	{_State185, FloatLiteralToken}:                _GotoState20Action,
	{_State185, RuneLiteralToken}:                 _GotoState32Action,
	{_State185, StringLiteralToken}:               _GotoState33Action,
	{_State185, IdentifierToken}:                  _GotoState23Action,
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
	{_State185, SubToken}:                         _GotoState35Action,
	{_State185, MulToken}:                         _GotoState29Action,
	{_State185, BitNegToken}:                      _GotoState18Action,
	{_State185, BitAndToken}:                      _GotoState17Action,
	{_State185, GreaterToken}:                     _GotoState22Action,
	{_State185, ParseErrorToken}:                  _GotoState31Action,
	{_State185, VarDeclPatternType}:               _GotoState56Action,
	{_State185, VarOrLetType}:                     _GotoState57Action,
	{_State185, ExpressionType}:                   _GotoState286Action,
	{_State185, OptionalLabelDeclType}:            _GotoState50Action,
	{_State185, SequenceExprType}:                 _GotoState55Action,
	{_State185, CallExprType}:                     _GotoState43Action,
	{_State185, OptionalExpressionType}:           _GotoState287Action,
	{_State185, AtomExprType}:                     _GotoState42Action,
	{_State185, LiteralType}:                      _GotoState48Action,
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
	{_State186, ExpressionType}:                   _GotoState286Action,
	{_State186, OptionalLabelDeclType}:            _GotoState50Action,
	{_State186, SequenceExprType}:                 _GotoState55Action,
	{_State186, CallExprType}:                     _GotoState43Action,
	{_State186, OptionalExpressionType}:           _GotoState288Action,
	{_State186, AtomExprType}:                     _GotoState42Action,
	{_State186, LiteralType}:                      _GotoState48Action,
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
	{_State187, NewlinesToken}:                    _GotoState290Action,
	{_State187, CommaToken}:                       _GotoState289Action,
	{_State189, RparenToken}:                      _GotoState291Action,
	{_State190, CommaToken}:                       _GotoState292Action,
	{_State191, RbracketToken}:                    _GotoState293Action,
	{_State192, AddToken}:                         _GotoState176Action,
	{_State192, SubToken}:                         _GotoState181Action,
	{_State192, MulToken}:                         _GotoState179Action,
	{_State194, RbracketToken}:                    _GotoState294Action,
	{_State195, IntegerLiteralToken}:              _GotoState24Action,
	{_State195, FloatLiteralToken}:                _GotoState20Action,
	{_State195, RuneLiteralToken}:                 _GotoState32Action,
	{_State195, StringLiteralToken}:               _GotoState33Action,
	{_State195, IdentifierToken}:                  _GotoState95Action,
	{_State195, TrueToken}:                        _GotoState36Action,
	{_State195, FalseToken}:                       _GotoState19Action,
	{_State195, StructToken}:                      _GotoState34Action,
	{_State195, FuncToken}:                        _GotoState21Action,
	{_State195, VarToken}:                         _GotoState37Action,
	{_State195, LetToken}:                         _GotoState27Action,
	{_State195, NotToken}:                         _GotoState30Action,
	{_State195, LabelDeclToken}:                   _GotoState25Action,
	{_State195, LparenToken}:                      _GotoState28Action,
	{_State195, LbracketToken}:                    _GotoState26Action,
	{_State195, DotDotDotToken}:                   _GotoState94Action,
	{_State195, SubToken}:                         _GotoState35Action,
	{_State195, MulToken}:                         _GotoState29Action,
	{_State195, BitNegToken}:                      _GotoState18Action,
	{_State195, BitAndToken}:                      _GotoState17Action,
	{_State195, GreaterToken}:                     _GotoState22Action,
	{_State195, ParseErrorToken}:                  _GotoState31Action,
	{_State195, VarDeclPatternType}:               _GotoState56Action,
	{_State195, VarOrLetType}:                     _GotoState57Action,
	{_State195, ExpressionType}:                   _GotoState99Action,
	{_State195, OptionalLabelDeclType}:            _GotoState50Action,
	{_State195, SequenceExprType}:                 _GotoState55Action,
	{_State195, CallExprType}:                     _GotoState43Action,
	{_State195, OptionalArgumentsType}:            _GotoState296Action,
	{_State195, ArgumentsType}:                    _GotoState295Action,
	{_State195, ArgumentType}:                     _GotoState96Action,
	{_State195, ColonExpressionsType}:             _GotoState98Action,
	{_State195, OptionalExpressionType}:           _GotoState100Action,
	{_State195, AtomExprType}:                     _GotoState42Action,
	{_State195, LiteralType}:                      _GotoState48Action,
	{_State195, ImplicitStructExprType}:           _GotoState46Action,
	{_State195, AccessExprType}:                   _GotoState38Action,
	{_State195, PostfixUnaryExprType}:             _GotoState52Action,
	{_State195, PrefixUnaryOpType}:                _GotoState54Action,
	{_State195, PrefixUnaryExprType}:              _GotoState53Action,
	{_State195, MulExprType}:                      _GotoState49Action,
	{_State195, AddExprType}:                      _GotoState39Action,
	{_State195, CmpExprType}:                      _GotoState44Action,
	{_State195, AndExprType}:                      _GotoState40Action,
	{_State195, OrExprType}:                       _GotoState51Action,
	{_State195, InitializableTypeType}:            _GotoState47Action,
	{_State195, ExplicitStructDefType}:            _GotoState45Action,
	{_State195, AnonymousFuncExprType}:            _GotoState41Action,
	{_State196, MulToken}:                         _GotoState126Action,
	{_State196, DivToken}:                         _GotoState124Action,
	{_State196, ModToken}:                         _GotoState125Action,
	{_State196, BitAndToken}:                      _GotoState121Action,
	{_State196, BitLshiftToken}:                   _GotoState122Action,
	{_State196, BitRshiftToken}:                   _GotoState123Action,
	{_State196, MulOpType}:                        _GotoState127Action,
	{_State197, EqualToken}:                       _GotoState113Action,
	{_State197, NotEqualToken}:                    _GotoState118Action,
	{_State197, LessToken}:                        _GotoState116Action,
	{_State197, LessOrEqualToken}:                 _GotoState117Action,
	{_State197, GreaterToken}:                     _GotoState114Action,
	{_State197, GreaterOrEqualToken}:              _GotoState115Action,
	{_State197, CmpOpType}:                        _GotoState119Action,
	{_State198, AddToken}:                         _GotoState107Action,
	{_State198, SubToken}:                         _GotoState110Action,
	{_State198, BitXorToken}:                      _GotoState109Action,
	{_State198, BitOrToken}:                       _GotoState108Action,
	{_State198, AddOpType}:                        _GotoState111Action,
	{_State199, RparenToken}:                      _GotoState297Action,
	{_State199, CommaToken}:                       _GotoState183Action,
	{_State201, ForToken}:                         _GotoState298Action,
	{_State202, InToken}:                          _GotoState300Action,
	{_State202, AssignToken}:                      _GotoState299Action,
	{_State203, SemicolonToken}:                   _GotoState301Action,
	{_State204, DoToken}:                          _GotoState302Action,
	{_State205, IntegerLiteralToken}:              _GotoState24Action,
	{_State205, FloatLiteralToken}:                _GotoState20Action,
	{_State205, RuneLiteralToken}:                 _GotoState32Action,
	{_State205, StringLiteralToken}:               _GotoState33Action,
	{_State205, IdentifierToken}:                  _GotoState23Action,
	{_State205, TrueToken}:                        _GotoState36Action,
	{_State205, FalseToken}:                       _GotoState19Action,
	{_State205, StructToken}:                      _GotoState34Action,
	{_State205, FuncToken}:                        _GotoState21Action,
	{_State205, VarToken}:                         _GotoState304Action,
	{_State205, LetToken}:                         _GotoState27Action,
	{_State205, NotToken}:                         _GotoState30Action,
	{_State205, LabelDeclToken}:                   _GotoState25Action,
	{_State205, LparenToken}:                      _GotoState28Action,
	{_State205, LbracketToken}:                    _GotoState26Action,
	{_State205, DotToken}:                         _GotoState303Action,
	{_State205, SubToken}:                         _GotoState35Action,
	{_State205, MulToken}:                         _GotoState29Action,
	{_State205, BitNegToken}:                      _GotoState18Action,
	{_State205, BitAndToken}:                      _GotoState17Action,
	{_State205, GreaterToken}:                     _GotoState22Action,
	{_State205, ParseErrorToken}:                  _GotoState31Action,
	{_State205, CasePatternsType}:                 _GotoState307Action,
	{_State205, VarDeclPatternType}:               _GotoState56Action,
	{_State205, VarOrLetType}:                     _GotoState57Action,
	{_State205, AssignPatternType}:                _GotoState305Action,
	{_State205, CasePatternType}:                  _GotoState306Action,
	{_State205, OptionalLabelDeclType}:            _GotoState71Action,
	{_State205, SequenceExprType}:                 _GotoState308Action,
	{_State205, CallExprType}:                     _GotoState43Action,
	{_State205, AtomExprType}:                     _GotoState42Action,
	{_State205, LiteralType}:                      _GotoState48Action,
	{_State205, ImplicitStructExprType}:           _GotoState46Action,
	{_State205, AccessExprType}:                   _GotoState38Action,
	{_State205, PostfixUnaryExprType}:             _GotoState52Action,
	{_State205, PrefixUnaryOpType}:                _GotoState54Action,
	{_State205, PrefixUnaryExprType}:              _GotoState53Action,
	{_State205, MulExprType}:                      _GotoState49Action,
	{_State205, AddExprType}:                      _GotoState39Action,
	{_State205, CmpExprType}:                      _GotoState44Action,
	{_State205, AndExprType}:                      _GotoState40Action,
	{_State205, OrExprType}:                       _GotoState51Action,
	{_State205, InitializableTypeType}:            _GotoState47Action,
	{_State205, ExplicitStructDefType}:            _GotoState45Action,
	{_State205, AnonymousFuncExprType}:            _GotoState41Action,
	{_State206, LbraceToken}:                      _GotoState58Action,
	{_State206, StatementBlockType}:               _GotoState309Action,
	{_State208, LbraceToken}:                      _GotoState58Action,
	{_State208, StatementBlockType}:               _GotoState310Action,
	{_State209, AndToken}:                         _GotoState112Action,
	{_State211, AssignToken}:                      _GotoState311Action,
	{_State213, RparenToken}:                      _GotoState313Action,
	{_State213, CommaToken}:                       _GotoState312Action,
	{_State216, AddToken}:                         _GotoState176Action,
	{_State216, SubToken}:                         _GotoState181Action,
	{_State216, MulToken}:                         _GotoState179Action,
	{_State217, IntegerLiteralToken}:              _GotoState24Action,
	{_State217, FloatLiteralToken}:                _GotoState20Action,
	{_State217, RuneLiteralToken}:                 _GotoState32Action,
	{_State217, StringLiteralToken}:               _GotoState33Action,
	{_State217, IdentifierToken}:                  _GotoState23Action,
	{_State217, TrueToken}:                        _GotoState36Action,
	{_State217, FalseToken}:                       _GotoState19Action,
	{_State217, StructToken}:                      _GotoState34Action,
	{_State217, FuncToken}:                        _GotoState21Action,
	{_State217, LabelDeclToken}:                   _GotoState25Action,
	{_State217, LparenToken}:                      _GotoState28Action,
	{_State217, LbracketToken}:                    _GotoState26Action,
	{_State217, ParseErrorToken}:                  _GotoState31Action,
	{_State217, OptionalLabelDeclType}:            _GotoState71Action,
	{_State217, CallExprType}:                     _GotoState315Action,
	{_State217, AtomExprType}:                     _GotoState42Action,
	{_State217, LiteralType}:                      _GotoState48Action,
	{_State217, ImplicitStructExprType}:           _GotoState46Action,
	{_State217, AccessExprType}:                   _GotoState314Action,
	{_State217, InitializableTypeType}:            _GotoState47Action,
	{_State217, ExplicitStructDefType}:            _GotoState45Action,
	{_State217, AnonymousFuncExprType}:            _GotoState41Action,
	{_State219, IntegerLiteralToken}:              _GotoState24Action,
	{_State219, FloatLiteralToken}:                _GotoState20Action,
	{_State219, RuneLiteralToken}:                 _GotoState32Action,
	{_State219, StringLiteralToken}:               _GotoState33Action,
	{_State219, IdentifierToken}:                  _GotoState23Action,
	{_State219, TrueToken}:                        _GotoState36Action,
	{_State219, FalseToken}:                       _GotoState19Action,
	{_State219, StructToken}:                      _GotoState34Action,
	{_State219, FuncToken}:                        _GotoState21Action,
	{_State219, VarToken}:                         _GotoState304Action,
	{_State219, LetToken}:                         _GotoState27Action,
	{_State219, NotToken}:                         _GotoState30Action,
	{_State219, LabelDeclToken}:                   _GotoState25Action,
	{_State219, LparenToken}:                      _GotoState28Action,
	{_State219, LbracketToken}:                    _GotoState26Action,
	{_State219, DotToken}:                         _GotoState303Action,
	{_State219, SubToken}:                         _GotoState35Action,
	{_State219, MulToken}:                         _GotoState29Action,
	{_State219, BitNegToken}:                      _GotoState18Action,
	{_State219, BitAndToken}:                      _GotoState17Action,
	{_State219, GreaterToken}:                     _GotoState22Action,
	{_State219, ParseErrorToken}:                  _GotoState31Action,
	{_State219, CasePatternsType}:                 _GotoState316Action,
	{_State219, VarDeclPatternType}:               _GotoState56Action,
	{_State219, VarOrLetType}:                     _GotoState57Action,
	{_State219, AssignPatternType}:                _GotoState305Action,
	{_State219, CasePatternType}:                  _GotoState306Action,
	{_State219, OptionalLabelDeclType}:            _GotoState71Action,
	{_State219, SequenceExprType}:                 _GotoState308Action,
	{_State219, CallExprType}:                     _GotoState43Action,
	{_State219, AtomExprType}:                     _GotoState42Action,
	{_State219, LiteralType}:                      _GotoState48Action,
	{_State219, ImplicitStructExprType}:           _GotoState46Action,
	{_State219, AccessExprType}:                   _GotoState38Action,
	{_State219, PostfixUnaryExprType}:             _GotoState52Action,
	{_State219, PrefixUnaryOpType}:                _GotoState54Action,
	{_State219, PrefixUnaryExprType}:              _GotoState53Action,
	{_State219, MulExprType}:                      _GotoState49Action,
	{_State219, AddExprType}:                      _GotoState39Action,
	{_State219, CmpExprType}:                      _GotoState44Action,
	{_State219, AndExprType}:                      _GotoState40Action,
	{_State219, OrExprType}:                       _GotoState51Action,
	{_State219, InitializableTypeType}:            _GotoState47Action,
	{_State219, ExplicitStructDefType}:            _GotoState45Action,
	{_State219, AnonymousFuncExprType}:            _GotoState41Action,
	{_State221, ColonToken}:                       _GotoState317Action,
	{_State222, IntegerLiteralToken}:              _GotoState24Action,
	{_State222, FloatLiteralToken}:                _GotoState20Action,
	{_State222, RuneLiteralToken}:                 _GotoState32Action,
	{_State222, StringLiteralToken}:               _GotoState33Action,
	{_State222, IdentifierToken}:                  _GotoState23Action,
	{_State222, TrueToken}:                        _GotoState36Action,
	{_State222, FalseToken}:                       _GotoState19Action,
	{_State222, StructToken}:                      _GotoState34Action,
	{_State222, FuncToken}:                        _GotoState21Action,
	{_State222, LabelDeclToken}:                   _GotoState25Action,
	{_State222, LparenToken}:                      _GotoState28Action,
	{_State222, LbracketToken}:                    _GotoState26Action,
	{_State222, ParseErrorToken}:                  _GotoState31Action,
	{_State222, OptionalLabelDeclType}:            _GotoState71Action,
	{_State222, CallExprType}:                     _GotoState318Action,
	{_State222, AtomExprType}:                     _GotoState42Action,
	{_State222, LiteralType}:                      _GotoState48Action,
	{_State222, ImplicitStructExprType}:           _GotoState46Action,
	{_State222, AccessExprType}:                   _GotoState314Action,
	{_State222, InitializableTypeType}:            _GotoState47Action,
	{_State222, ExplicitStructDefType}:            _GotoState45Action,
	{_State222, AnonymousFuncExprType}:            _GotoState41Action,
	{_State224, StringLiteralToken}:               _GotoState320Action,
	{_State224, LparenToken}:                      _GotoState319Action,
	{_State224, ImportClauseType}:                 _GotoState321Action,
	{_State227, LbracketToken}:                    _GotoState104Action,
	{_State227, DotToken}:                         _GotoState103Action,
	{_State227, QuestionToken}:                    _GotoState105Action,
	{_State227, DollarLbracketToken}:              _GotoState102Action,
	{_State227, AddAssignToken}:                   _GotoState322Action,
	{_State227, SubAssignToken}:                   _GotoState333Action,
	{_State227, MulAssignToken}:                   _GotoState332Action,
	{_State227, DivAssignToken}:                   _GotoState330Action,
	{_State227, ModAssignToken}:                   _GotoState331Action,
	{_State227, AddOneAssignToken}:                _GotoState323Action,
	{_State227, SubOneAssignToken}:                _GotoState334Action,
	{_State227, BitNegAssignToken}:                _GotoState326Action,
	{_State227, BitAndAssignToken}:                _GotoState324Action,
	{_State227, BitOrAssignToken}:                 _GotoState327Action,
	{_State227, BitXorAssignToken}:                _GotoState329Action,
	{_State227, BitLshiftAssignToken}:             _GotoState325Action,
	{_State227, BitRshiftAssignToken}:             _GotoState328Action,
	{_State227, UnaryOpAssignType}:                _GotoState336Action,
	{_State227, BinaryOpAssignType}:               _GotoState335Action,
	{_State227, OptionalGenericBindingType}:       _GotoState106Action,
	{_State228, AssignToken}:                      _GotoState337Action,
	{_State230, CommaToken}:                       _GotoState338Action,
	{_State232, JumpLabelToken}:                   _GotoState339Action,
	{_State232, OptionalJumpLabelType}:            _GotoState340Action,
	{_State236, NewlinesToken}:                    _GotoState341Action,
	{_State236, SemicolonToken}:                   _GotoState342Action,
	{_State240, AddToken}:                         _GotoState176Action,
	{_State240, SubToken}:                         _GotoState181Action,
	{_State240, MulToken}:                         _GotoState179Action,
	{_State241, IdentifierToken}:                  _GotoState79Action,
	{_State241, StructToken}:                      _GotoState34Action,
	{_State241, EnumToken}:                        _GotoState76Action,
	{_State241, TraitToken}:                       _GotoState84Action,
	{_State241, FuncToken}:                        _GotoState78Action,
	{_State241, LparenToken}:                      _GotoState80Action,
	{_State241, LbracketToken}:                    _GotoState26Action,
	{_State241, DotToken}:                         _GotoState75Action,
	{_State241, QuestionToken}:                    _GotoState82Action,
	{_State241, ExclaimToken}:                     _GotoState77Action,
	{_State241, TildeTildeToken}:                  _GotoState83Action,
	{_State241, BitNegToken}:                      _GotoState74Action,
	{_State241, BitAndToken}:                      _GotoState73Action,
	{_State241, ParseErrorToken}:                  _GotoState81Action,
	{_State241, InitializableTypeType}:            _GotoState90Action,
	{_State241, AtomTypeType}:                     _GotoState85Action,
	{_State241, ReturnableTypeType}:               _GotoState91Action,
	{_State241, ValueTypeType}:                    _GotoState343Action,
	{_State241, ImplicitStructDefType}:            _GotoState89Action,
	{_State241, ExplicitStructDefType}:            _GotoState45Action,
	{_State241, ImplicitEnumDefType}:              _GotoState88Action,
	{_State241, ExplicitEnumDefType}:              _GotoState86Action,
	{_State241, TraitDefType}:                     _GotoState92Action,
	{_State241, FuncTypeType}:                     _GotoState87Action,
	{_State243, CommaToken}:                       _GotoState344Action,
	{_State244, RbracketToken}:                    _GotoState345Action,
	{_State245, ImplementsToken}:                  _GotoState346Action,
	{_State245, AddToken}:                         _GotoState176Action,
	{_State245, SubToken}:                         _GotoState181Action,
	{_State245, MulToken}:                         _GotoState179Action,
	{_State247, IdentifierToken}:                  _GotoState151Action,
	{_State247, ParameterDefType}:                 _GotoState154Action,
	{_State247, ParameterDefsType}:                _GotoState155Action,
	{_State247, OptionalParameterDefsType}:        _GotoState347Action,
	{_State248, IdentifierToken}:                  _GotoState79Action,
	{_State248, StructToken}:                      _GotoState34Action,
	{_State248, EnumToken}:                        _GotoState76Action,
	{_State248, TraitToken}:                       _GotoState84Action,
	{_State248, FuncToken}:                        _GotoState78Action,
	{_State248, LparenToken}:                      _GotoState80Action,
	{_State248, LbracketToken}:                    _GotoState26Action,
	{_State248, DotToken}:                         _GotoState75Action,
	{_State248, QuestionToken}:                    _GotoState82Action,
	{_State248, ExclaimToken}:                     _GotoState77Action,
	{_State248, TildeTildeToken}:                  _GotoState83Action,
	{_State248, BitNegToken}:                      _GotoState74Action,
	{_State248, BitAndToken}:                      _GotoState73Action,
	{_State248, ParseErrorToken}:                  _GotoState81Action,
	{_State248, InitializableTypeType}:            _GotoState90Action,
	{_State248, AtomTypeType}:                     _GotoState85Action,
	{_State248, ReturnableTypeType}:               _GotoState91Action,
	{_State248, ValueTypeType}:                    _GotoState348Action,
	{_State248, ImplicitStructDefType}:            _GotoState89Action,
	{_State248, ExplicitStructDefType}:            _GotoState45Action,
	{_State248, ImplicitEnumDefType}:              _GotoState88Action,
	{_State248, ExplicitEnumDefType}:              _GotoState86Action,
	{_State248, TraitDefType}:                     _GotoState92Action,
	{_State248, FuncTypeType}:                     _GotoState87Action,
	{_State249, AddToken}:                         _GotoState176Action,
	{_State249, SubToken}:                         _GotoState181Action,
	{_State249, MulToken}:                         _GotoState179Action,
	{_State250, IdentifierToken}:                  _GotoState349Action,
	{_State251, IdentifierToken}:                  _GotoState79Action,
	{_State251, StructToken}:                      _GotoState34Action,
	{_State251, EnumToken}:                        _GotoState76Action,
	{_State251, TraitToken}:                       _GotoState84Action,
	{_State251, FuncToken}:                        _GotoState78Action,
	{_State251, LparenToken}:                      _GotoState80Action,
	{_State251, LbracketToken}:                    _GotoState26Action,
	{_State251, DotToken}:                         _GotoState75Action,
	{_State251, QuestionToken}:                    _GotoState82Action,
	{_State251, ExclaimToken}:                     _GotoState77Action,
	{_State251, TildeTildeToken}:                  _GotoState83Action,
	{_State251, BitNegToken}:                      _GotoState74Action,
	{_State251, BitAndToken}:                      _GotoState73Action,
	{_State251, ParseErrorToken}:                  _GotoState81Action,
	{_State251, InitializableTypeType}:            _GotoState90Action,
	{_State251, AtomTypeType}:                     _GotoState85Action,
	{_State251, ReturnableTypeType}:               _GotoState351Action,
	{_State251, ImplicitStructDefType}:            _GotoState89Action,
	{_State251, ExplicitStructDefType}:            _GotoState45Action,
	{_State251, ImplicitEnumDefType}:              _GotoState88Action,
	{_State251, ExplicitEnumDefType}:              _GotoState86Action,
	{_State251, TraitDefType}:                     _GotoState92Action,
	{_State251, ReturnTypeType}:                   _GotoState350Action,
	{_State251, FuncTypeType}:                     _GotoState87Action,
	{_State252, IdentifierToken}:                  _GotoState151Action,
	{_State252, ParameterDefType}:                 _GotoState352Action,
	{_State253, NewlinesToken}:                    _GotoState353Action,
	{_State253, OrToken}:                          _GotoState354Action,
	{_State254, RparenToken}:                      _GotoState355Action,
	{_State255, AssignToken}:                      _GotoState268Action,
	{_State256, NewlinesToken}:                    _GotoState356Action,
	{_State256, OrToken}:                          _GotoState357Action,
	{_State257, IdentifierToken}:                  _GotoState79Action,
	{_State257, StructToken}:                      _GotoState34Action,
	{_State257, EnumToken}:                        _GotoState76Action,
	{_State257, TraitToken}:                       _GotoState84Action,
	{_State257, FuncToken}:                        _GotoState78Action,
	{_State257, LparenToken}:                      _GotoState80Action,
	{_State257, LbracketToken}:                    _GotoState26Action,
	{_State257, DotToken}:                         _GotoState75Action,
	{_State257, QuestionToken}:                    _GotoState82Action,
	{_State257, ExclaimToken}:                     _GotoState77Action,
	{_State257, TildeTildeToken}:                  _GotoState83Action,
	{_State257, BitNegToken}:                      _GotoState74Action,
	{_State257, BitAndToken}:                      _GotoState73Action,
	{_State257, ParseErrorToken}:                  _GotoState81Action,
	{_State257, InitializableTypeType}:            _GotoState90Action,
	{_State257, AtomTypeType}:                     _GotoState85Action,
	{_State257, ReturnableTypeType}:               _GotoState91Action,
	{_State257, ValueTypeType}:                    _GotoState358Action,
	{_State257, ImplicitStructDefType}:            _GotoState89Action,
	{_State257, ExplicitStructDefType}:            _GotoState45Action,
	{_State257, ImplicitEnumDefType}:              _GotoState88Action,
	{_State257, ExplicitEnumDefType}:              _GotoState86Action,
	{_State257, TraitDefType}:                     _GotoState92Action,
	{_State257, FuncTypeType}:                     _GotoState87Action,
	{_State258, IdentifierToken}:                  _GotoState79Action,
	{_State258, StructToken}:                      _GotoState34Action,
	{_State258, EnumToken}:                        _GotoState76Action,
	{_State258, TraitToken}:                       _GotoState84Action,
	{_State258, FuncToken}:                        _GotoState78Action,
	{_State258, LparenToken}:                      _GotoState80Action,
	{_State258, LbracketToken}:                    _GotoState26Action,
	{_State258, DotToken}:                         _GotoState264Action,
	{_State258, QuestionToken}:                    _GotoState82Action,
	{_State258, ExclaimToken}:                     _GotoState77Action,
	{_State258, DollarLbracketToken}:              _GotoState102Action,
	{_State258, DotDotDotToken}:                   _GotoState359Action,
	{_State258, TildeTildeToken}:                  _GotoState83Action,
	{_State258, BitNegToken}:                      _GotoState74Action,
	{_State258, BitAndToken}:                      _GotoState73Action,
	{_State258, ParseErrorToken}:                  _GotoState81Action,
	{_State258, OptionalGenericBindingType}:       _GotoState163Action,
	{_State258, InitializableTypeType}:            _GotoState90Action,
	{_State258, AtomTypeType}:                     _GotoState85Action,
	{_State258, ReturnableTypeType}:               _GotoState91Action,
	{_State258, ValueTypeType}:                    _GotoState360Action,
	{_State258, ImplicitStructDefType}:            _GotoState89Action,
	{_State258, ExplicitStructDefType}:            _GotoState45Action,
	{_State258, ImplicitEnumDefType}:              _GotoState88Action,
	{_State258, ExplicitEnumDefType}:              _GotoState86Action,
	{_State258, TraitDefType}:                     _GotoState92Action,
	{_State258, FuncTypeType}:                     _GotoState87Action,
	{_State259, RparenToken}:                      _GotoState361Action,
	{_State261, CommaToken}:                       _GotoState362Action,
	{_State262, AddToken}:                         _GotoState176Action,
	{_State262, SubToken}:                         _GotoState181Action,
	{_State262, MulToken}:                         _GotoState179Action,
	{_State263, DollarLbracketToken}:              _GotoState102Action,
	{_State263, OptionalGenericBindingType}:       _GotoState363Action,
	{_State264, IdentifierToken}:                  _GotoState263Action,
	{_State264, DollarLbracketToken}:              _GotoState102Action,
	{_State264, OptionalGenericBindingType}:       _GotoState158Action,
	{_State265, AddToken}:                         _GotoState176Action,
	{_State265, SubToken}:                         _GotoState181Action,
	{_State265, MulToken}:                         _GotoState179Action,
	{_State266, IdentifierToken}:                  _GotoState364Action,
	{_State267, IdentifierToken}:                  _GotoState164Action,
	{_State267, UnsafeToken}:                      _GotoState165Action,
	{_State267, StructToken}:                      _GotoState34Action,
	{_State267, EnumToken}:                        _GotoState76Action,
	{_State267, TraitToken}:                       _GotoState84Action,
	{_State267, FuncToken}:                        _GotoState78Action,
	{_State267, LparenToken}:                      _GotoState80Action,
	{_State267, LbracketToken}:                    _GotoState26Action,
	{_State267, DotToken}:                         _GotoState75Action,
	{_State267, QuestionToken}:                    _GotoState82Action,
	{_State267, ExclaimToken}:                     _GotoState77Action,
	{_State267, TildeTildeToken}:                  _GotoState83Action,
	{_State267, BitNegToken}:                      _GotoState74Action,
	{_State267, BitAndToken}:                      _GotoState73Action,
	{_State267, ParseErrorToken}:                  _GotoState81Action,
	{_State267, UnsafeStatementType}:              _GotoState171Action,
	{_State267, InitializableTypeType}:            _GotoState90Action,
	{_State267, AtomTypeType}:                     _GotoState85Action,
	{_State267, ReturnableTypeType}:               _GotoState91Action,
	{_State267, ValueTypeType}:                    _GotoState172Action,
	{_State267, FieldDefType}:                     _GotoState255Action,
	{_State267, ImplicitStructDefType}:            _GotoState89Action,
	{_State267, ExplicitStructDefType}:            _GotoState45Action,
	{_State267, EnumValueDefType}:                 _GotoState365Action,
	{_State267, ImplicitEnumDefType}:              _GotoState88Action,
	{_State267, ExplicitEnumDefType}:              _GotoState86Action,
	{_State267, TraitDefType}:                     _GotoState92Action,
	{_State267, FuncTypeType}:                     _GotoState87Action,
	{_State268, DefaultToken}:                     _GotoState366Action,
	{_State269, IdentifierToken}:                  _GotoState164Action,
	{_State269, UnsafeToken}:                      _GotoState165Action,
	{_State269, StructToken}:                      _GotoState34Action,
	{_State269, EnumToken}:                        _GotoState76Action,
	{_State269, TraitToken}:                       _GotoState84Action,
	{_State269, FuncToken}:                        _GotoState78Action,
	{_State269, LparenToken}:                      _GotoState80Action,
	{_State269, LbracketToken}:                    _GotoState26Action,
	{_State269, DotToken}:                         _GotoState75Action,
	{_State269, QuestionToken}:                    _GotoState82Action,
	{_State269, ExclaimToken}:                     _GotoState77Action,
	{_State269, TildeTildeToken}:                  _GotoState83Action,
	{_State269, BitNegToken}:                      _GotoState74Action,
	{_State269, BitAndToken}:                      _GotoState73Action,
	{_State269, ParseErrorToken}:                  _GotoState81Action,
	{_State269, UnsafeStatementType}:              _GotoState171Action,
	{_State269, InitializableTypeType}:            _GotoState90Action,
	{_State269, AtomTypeType}:                     _GotoState85Action,
	{_State269, ReturnableTypeType}:               _GotoState91Action,
	{_State269, ValueTypeType}:                    _GotoState172Action,
	{_State269, FieldDefType}:                     _GotoState255Action,
	{_State269, ImplicitStructDefType}:            _GotoState89Action,
	{_State269, ExplicitStructDefType}:            _GotoState45Action,
	{_State269, EnumValueDefType}:                 _GotoState367Action,
	{_State269, ImplicitEnumDefType}:              _GotoState88Action,
	{_State269, ExplicitEnumDefType}:              _GotoState86Action,
	{_State269, TraitDefType}:                     _GotoState92Action,
	{_State269, FuncTypeType}:                     _GotoState87Action,
	{_State271, IdentifierToken}:                  _GotoState164Action,
	{_State271, UnsafeToken}:                      _GotoState165Action,
	{_State271, StructToken}:                      _GotoState34Action,
	{_State271, EnumToken}:                        _GotoState76Action,
	{_State271, TraitToken}:                       _GotoState84Action,
	{_State271, FuncToken}:                        _GotoState78Action,
	{_State271, LparenToken}:                      _GotoState80Action,
	{_State271, LbracketToken}:                    _GotoState26Action,
	{_State271, DotToken}:                         _GotoState75Action,
	{_State271, QuestionToken}:                    _GotoState82Action,
	{_State271, ExclaimToken}:                     _GotoState77Action,
	{_State271, TildeTildeToken}:                  _GotoState83Action,
	{_State271, BitNegToken}:                      _GotoState74Action,
	{_State271, BitAndToken}:                      _GotoState73Action,
	{_State271, ParseErrorToken}:                  _GotoState81Action,
	{_State271, UnsafeStatementType}:              _GotoState171Action,
	{_State271, InitializableTypeType}:            _GotoState90Action,
	{_State271, AtomTypeType}:                     _GotoState85Action,
	{_State271, ReturnableTypeType}:               _GotoState91Action,
	{_State271, ValueTypeType}:                    _GotoState172Action,
	{_State271, FieldDefType}:                     _GotoState368Action,
	{_State271, ImplicitStructDefType}:            _GotoState89Action,
	{_State271, ExplicitStructDefType}:            _GotoState45Action,
	{_State271, ImplicitEnumDefType}:              _GotoState88Action,
	{_State271, ExplicitEnumDefType}:              _GotoState86Action,
	{_State271, TraitDefType}:                     _GotoState92Action,
	{_State271, FuncTypeType}:                     _GotoState87Action,
	{_State273, IdentifierToken}:                  _GotoState369Action,
	{_State273, LparenToken}:                      _GotoState161Action,
	{_State276, RparenToken}:                      _GotoState370Action,
	{_State277, NewlinesToken}:                    _GotoState372Action,
	{_State277, CommaToken}:                       _GotoState371Action,
	{_State280, RbracketToken}:                    _GotoState373Action,
	{_State280, AddToken}:                         _GotoState176Action,
	{_State280, SubToken}:                         _GotoState181Action,
	{_State280, MulToken}:                         _GotoState179Action,
	{_State281, RbracketToken}:                    _GotoState374Action,
	{_State289, IdentifierToken}:                  _GotoState164Action,
	{_State289, UnsafeToken}:                      _GotoState165Action,
	{_State289, StructToken}:                      _GotoState34Action,
	{_State289, EnumToken}:                        _GotoState76Action,
	{_State289, TraitToken}:                       _GotoState84Action,
	{_State289, FuncToken}:                        _GotoState78Action,
	{_State289, LparenToken}:                      _GotoState80Action,
	{_State289, LbracketToken}:                    _GotoState26Action,
	{_State289, DotToken}:                         _GotoState75Action,
	{_State289, QuestionToken}:                    _GotoState82Action,
	{_State289, ExclaimToken}:                     _GotoState77Action,
	{_State289, TildeTildeToken}:                  _GotoState83Action,
	{_State289, BitNegToken}:                      _GotoState74Action,
	{_State289, BitAndToken}:                      _GotoState73Action,
	{_State289, ParseErrorToken}:                  _GotoState81Action,
	{_State289, UnsafeStatementType}:              _GotoState171Action,
	{_State289, InitializableTypeType}:            _GotoState90Action,
	{_State289, AtomTypeType}:                     _GotoState85Action,
	{_State289, ReturnableTypeType}:               _GotoState91Action,
	{_State289, ValueTypeType}:                    _GotoState172Action,
	{_State289, FieldDefType}:                     _GotoState375Action,
	{_State289, ImplicitStructDefType}:            _GotoState89Action,
	{_State289, ExplicitStructDefType}:            _GotoState45Action,
	{_State289, ImplicitEnumDefType}:              _GotoState88Action,
	{_State289, ExplicitEnumDefType}:              _GotoState86Action,
	{_State289, TraitDefType}:                     _GotoState92Action,
	{_State289, FuncTypeType}:                     _GotoState87Action,
	{_State290, IdentifierToken}:                  _GotoState164Action,
	{_State290, UnsafeToken}:                      _GotoState165Action,
	{_State290, StructToken}:                      _GotoState34Action,
	{_State290, EnumToken}:                        _GotoState76Action,
	{_State290, TraitToken}:                       _GotoState84Action,
	{_State290, FuncToken}:                        _GotoState78Action,
	{_State290, LparenToken}:                      _GotoState80Action,
	{_State290, LbracketToken}:                    _GotoState26Action,
	{_State290, DotToken}:                         _GotoState75Action,
	{_State290, QuestionToken}:                    _GotoState82Action,
	{_State290, ExclaimToken}:                     _GotoState77Action,
	{_State290, TildeTildeToken}:                  _GotoState83Action,
	{_State290, BitNegToken}:                      _GotoState74Action,
	{_State290, BitAndToken}:                      _GotoState73Action,
	{_State290, ParseErrorToken}:                  _GotoState81Action,
	{_State290, UnsafeStatementType}:              _GotoState171Action,
	{_State290, InitializableTypeType}:            _GotoState90Action,
	{_State290, AtomTypeType}:                     _GotoState85Action,
	{_State290, ReturnableTypeType}:               _GotoState91Action,
	{_State290, ValueTypeType}:                    _GotoState172Action,
	{_State290, FieldDefType}:                     _GotoState376Action,
	{_State290, ImplicitStructDefType}:            _GotoState89Action,
	{_State290, ExplicitStructDefType}:            _GotoState45Action,
	{_State290, ImplicitEnumDefType}:              _GotoState88Action,
	{_State290, ExplicitEnumDefType}:              _GotoState86Action,
	{_State290, TraitDefType}:                     _GotoState92Action,
	{_State290, FuncTypeType}:                     _GotoState87Action,
	{_State292, IdentifierToken}:                  _GotoState79Action,
	{_State292, StructToken}:                      _GotoState34Action,
	{_State292, EnumToken}:                        _GotoState76Action,
	{_State292, TraitToken}:                       _GotoState84Action,
	{_State292, FuncToken}:                        _GotoState78Action,
	{_State292, LparenToken}:                      _GotoState80Action,
	{_State292, LbracketToken}:                    _GotoState26Action,
	{_State292, DotToken}:                         _GotoState75Action,
	{_State292, QuestionToken}:                    _GotoState82Action,
	{_State292, ExclaimToken}:                     _GotoState77Action,
	{_State292, TildeTildeToken}:                  _GotoState83Action,
	{_State292, BitNegToken}:                      _GotoState74Action,
	{_State292, BitAndToken}:                      _GotoState73Action,
	{_State292, ParseErrorToken}:                  _GotoState81Action,
	{_State292, InitializableTypeType}:            _GotoState90Action,
	{_State292, AtomTypeType}:                     _GotoState85Action,
	{_State292, ReturnableTypeType}:               _GotoState91Action,
	{_State292, ValueTypeType}:                    _GotoState377Action,
	{_State292, ImplicitStructDefType}:            _GotoState89Action,
	{_State292, ExplicitStructDefType}:            _GotoState45Action,
	{_State292, ImplicitEnumDefType}:              _GotoState88Action,
	{_State292, ExplicitEnumDefType}:              _GotoState86Action,
	{_State292, TraitDefType}:                     _GotoState92Action,
	{_State292, FuncTypeType}:                     _GotoState87Action,
	{_State295, CommaToken}:                       _GotoState183Action,
	{_State296, RparenToken}:                      _GotoState378Action,
	{_State298, IntegerLiteralToken}:              _GotoState24Action,
	{_State298, FloatLiteralToken}:                _GotoState20Action,
	{_State298, RuneLiteralToken}:                 _GotoState32Action,
	{_State298, StringLiteralToken}:               _GotoState33Action,
	{_State298, IdentifierToken}:                  _GotoState23Action,
	{_State298, TrueToken}:                        _GotoState36Action,
	{_State298, FalseToken}:                       _GotoState19Action,
	{_State298, StructToken}:                      _GotoState34Action,
	{_State298, FuncToken}:                        _GotoState21Action,
	{_State298, VarToken}:                         _GotoState37Action,
	{_State298, LetToken}:                         _GotoState27Action,
	{_State298, NotToken}:                         _GotoState30Action,
	{_State298, LabelDeclToken}:                   _GotoState25Action,
	{_State298, LparenToken}:                      _GotoState28Action,
	{_State298, LbracketToken}:                    _GotoState26Action,
	{_State298, SubToken}:                         _GotoState35Action,
	{_State298, MulToken}:                         _GotoState29Action,
	{_State298, BitNegToken}:                      _GotoState18Action,
	{_State298, BitAndToken}:                      _GotoState17Action,
	{_State298, GreaterToken}:                     _GotoState22Action,
	{_State298, ParseErrorToken}:                  _GotoState31Action,
	{_State298, VarDeclPatternType}:               _GotoState56Action,
	{_State298, VarOrLetType}:                     _GotoState57Action,
	{_State298, OptionalLabelDeclType}:            _GotoState71Action,
	{_State298, SequenceExprType}:                 _GotoState379Action,
	{_State298, CallExprType}:                     _GotoState43Action,
	{_State298, AtomExprType}:                     _GotoState42Action,
	{_State298, LiteralType}:                      _GotoState48Action,
	{_State298, ImplicitStructExprType}:           _GotoState46Action,
	{_State298, AccessExprType}:                   _GotoState38Action,
	{_State298, PostfixUnaryExprType}:             _GotoState52Action,
	{_State298, PrefixUnaryOpType}:                _GotoState54Action,
	{_State298, PrefixUnaryExprType}:              _GotoState53Action,
	{_State298, MulExprType}:                      _GotoState49Action,
	{_State298, AddExprType}:                      _GotoState39Action,
	{_State298, CmpExprType}:                      _GotoState44Action,
	{_State298, AndExprType}:                      _GotoState40Action,
	{_State298, OrExprType}:                       _GotoState51Action,
	{_State298, InitializableTypeType}:            _GotoState47Action,
	{_State298, ExplicitStructDefType}:            _GotoState45Action,
	{_State298, AnonymousFuncExprType}:            _GotoState41Action,
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
	{_State299, OptionalLabelDeclType}:            _GotoState71Action,
	{_State299, SequenceExprType}:                 _GotoState380Action,
	{_State299, CallExprType}:                     _GotoState43Action,
	{_State299, AtomExprType}:                     _GotoState42Action,
	{_State299, LiteralType}:                      _GotoState48Action,
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
	{_State300, OptionalLabelDeclType}:            _GotoState71Action,
	{_State300, SequenceExprType}:                 _GotoState381Action,
	{_State300, CallExprType}:                     _GotoState43Action,
	{_State300, AtomExprType}:                     _GotoState42Action,
	{_State300, LiteralType}:                      _GotoState48Action,
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
	{_State301, OptionalLabelDeclType}:            _GotoState71Action,
	{_State301, SequenceExprType}:                 _GotoState383Action,
	{_State301, OptionalSequenceExprType}:         _GotoState382Action,
	{_State301, CallExprType}:                     _GotoState43Action,
	{_State301, AtomExprType}:                     _GotoState42Action,
	{_State301, LiteralType}:                      _GotoState48Action,
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
	{_State302, LbraceToken}:                      _GotoState58Action,
	{_State302, StatementBlockType}:               _GotoState384Action,
	{_State303, IdentifierToken}:                  _GotoState385Action,
	{_State304, DotToken}:                         _GotoState386Action,
	{_State307, CommaToken}:                       _GotoState388Action,
	{_State307, AssignToken}:                      _GotoState387Action,
	{_State309, ElseToken}:                        _GotoState389Action,
	{_State311, IdentifierToken}:                  _GotoState138Action,
	{_State311, LparenToken}:                      _GotoState139Action,
	{_State311, VarPatternType}:                   _GotoState390Action,
	{_State311, TuplePatternType}:                 _GotoState140Action,
	{_State312, IdentifierToken}:                  _GotoState211Action,
	{_State312, LparenToken}:                      _GotoState139Action,
	{_State312, DotDotDotToken}:                   _GotoState210Action,
	{_State312, VarPatternType}:                   _GotoState214Action,
	{_State312, TuplePatternType}:                 _GotoState140Action,
	{_State312, FieldVarPatternType}:              _GotoState391Action,
	{_State314, LbracketToken}:                    _GotoState104Action,
	{_State314, DotToken}:                         _GotoState103Action,
	{_State314, DollarLbracketToken}:              _GotoState102Action,
	{_State314, OptionalGenericBindingType}:       _GotoState106Action,
	{_State316, CommaToken}:                       _GotoState388Action,
	{_State316, ColonToken}:                       _GotoState392Action,
	{_State317, IntegerLiteralToken}:              _GotoState24Action,
	{_State317, FloatLiteralToken}:                _GotoState20Action,
	{_State317, RuneLiteralToken}:                 _GotoState32Action,
	{_State317, StringLiteralToken}:               _GotoState33Action,
	{_State317, IdentifierToken}:                  _GotoState23Action,
	{_State317, TrueToken}:                        _GotoState36Action,
	{_State317, FalseToken}:                       _GotoState19Action,
	{_State317, ReturnToken}:                      _GotoState226Action,
	{_State317, BreakToken}:                       _GotoState218Action,
	{_State317, ContinueToken}:                    _GotoState220Action,
	{_State317, FallthroughToken}:                 _GotoState223Action,
	{_State317, UnsafeToken}:                      _GotoState165Action,
	{_State317, StructToken}:                      _GotoState34Action,
	{_State317, FuncToken}:                        _GotoState21Action,
	{_State317, AsyncToken}:                       _GotoState217Action,
	{_State317, DeferToken}:                       _GotoState222Action,
	{_State317, VarToken}:                         _GotoState37Action,
	{_State317, LetToken}:                         _GotoState27Action,
	{_State317, NotToken}:                         _GotoState30Action,
	{_State317, LabelDeclToken}:                   _GotoState25Action,
	{_State317, LparenToken}:                      _GotoState28Action,
	{_State317, LbracketToken}:                    _GotoState26Action,
	{_State317, SubToken}:                         _GotoState35Action,
	{_State317, MulToken}:                         _GotoState29Action,
	{_State317, BitNegToken}:                      _GotoState18Action,
	{_State317, BitAndToken}:                      _GotoState17Action,
	{_State317, GreaterToken}:                     _GotoState22Action,
	{_State317, ParseErrorToken}:                  _GotoState31Action,
	{_State317, SimpleStatementBodyType}:          _GotoState394Action,
	{_State317, OptionalSimpleStatementBodyType}:  _GotoState393Action,
	{_State317, UnsafeStatementType}:              _GotoState237Action,
	{_State317, JumpTypeType}:                     _GotoState232Action,
	{_State317, ExpressionsType}:                  _GotoState230Action,
	{_State317, VarDeclPatternType}:               _GotoState56Action,
	{_State317, VarOrLetType}:                     _GotoState57Action,
	{_State317, AssignPatternType}:                _GotoState228Action,
	{_State317, ExpressionType}:                   _GotoState229Action,
	{_State317, OptionalLabelDeclType}:            _GotoState50Action,
	{_State317, SequenceExprType}:                 _GotoState233Action,
	{_State317, CallExprType}:                     _GotoState43Action,
	{_State317, AtomExprType}:                     _GotoState42Action,
	{_State317, LiteralType}:                      _GotoState48Action,
	{_State317, ImplicitStructExprType}:           _GotoState46Action,
	{_State317, AccessExprType}:                   _GotoState227Action,
	{_State317, PostfixUnaryExprType}:             _GotoState52Action,
	{_State317, PrefixUnaryOpType}:                _GotoState54Action,
	{_State317, PrefixUnaryExprType}:              _GotoState53Action,
	{_State317, MulExprType}:                      _GotoState49Action,
	{_State317, AddExprType}:                      _GotoState39Action,
	{_State317, CmpExprType}:                      _GotoState44Action,
	{_State317, AndExprType}:                      _GotoState40Action,
	{_State317, OrExprType}:                       _GotoState51Action,
	{_State317, InitializableTypeType}:            _GotoState47Action,
	{_State317, ExplicitStructDefType}:            _GotoState45Action,
	{_State317, AnonymousFuncExprType}:            _GotoState41Action,
	{_State319, StringLiteralToken}:               _GotoState320Action,
	{_State319, ImportClauseType}:                 _GotoState395Action,
	{_State319, ImportClauseTerminalType}:         _GotoState396Action,
	{_State319, ImportClausesType}:                _GotoState397Action,
	{_State320, AsToken}:                          _GotoState398Action,
	{_State335, IntegerLiteralToken}:              _GotoState24Action,
	{_State335, FloatLiteralToken}:                _GotoState20Action,
	{_State335, RuneLiteralToken}:                 _GotoState32Action,
	{_State335, StringLiteralToken}:               _GotoState33Action,
	{_State335, IdentifierToken}:                  _GotoState23Action,
	{_State335, TrueToken}:                        _GotoState36Action,
	{_State335, FalseToken}:                       _GotoState19Action,
	{_State335, StructToken}:                      _GotoState34Action,
	{_State335, FuncToken}:                        _GotoState21Action,
	{_State335, VarToken}:                         _GotoState37Action,
	{_State335, LetToken}:                         _GotoState27Action,
	{_State335, NotToken}:                         _GotoState30Action,
	{_State335, LabelDeclToken}:                   _GotoState25Action,
	{_State335, LparenToken}:                      _GotoState28Action,
	{_State335, LbracketToken}:                    _GotoState26Action,
	{_State335, SubToken}:                         _GotoState35Action,
	{_State335, MulToken}:                         _GotoState29Action,
	{_State335, BitNegToken}:                      _GotoState18Action,
	{_State335, BitAndToken}:                      _GotoState17Action,
	{_State335, GreaterToken}:                     _GotoState22Action,
	{_State335, ParseErrorToken}:                  _GotoState31Action,
	{_State335, VarDeclPatternType}:               _GotoState56Action,
	{_State335, VarOrLetType}:                     _GotoState57Action,
	{_State335, ExpressionType}:                   _GotoState399Action,
	{_State335, OptionalLabelDeclType}:            _GotoState50Action,
	{_State335, SequenceExprType}:                 _GotoState55Action,
	{_State335, CallExprType}:                     _GotoState43Action,
	{_State335, AtomExprType}:                     _GotoState42Action,
	{_State335, LiteralType}:                      _GotoState48Action,
	{_State335, ImplicitStructExprType}:           _GotoState46Action,
	{_State335, AccessExprType}:                   _GotoState38Action,
	{_State335, PostfixUnaryExprType}:             _GotoState52Action,
	{_State335, PrefixUnaryOpType}:                _GotoState54Action,
	{_State335, PrefixUnaryExprType}:              _GotoState53Action,
	{_State335, MulExprType}:                      _GotoState49Action,
	{_State335, AddExprType}:                      _GotoState39Action,
	{_State335, CmpExprType}:                      _GotoState44Action,
	{_State335, AndExprType}:                      _GotoState40Action,
	{_State335, OrExprType}:                       _GotoState51Action,
	{_State335, InitializableTypeType}:            _GotoState47Action,
	{_State335, ExplicitStructDefType}:            _GotoState45Action,
	{_State335, AnonymousFuncExprType}:            _GotoState41Action,
	{_State337, IntegerLiteralToken}:              _GotoState24Action,
	{_State337, FloatLiteralToken}:                _GotoState20Action,
	{_State337, RuneLiteralToken}:                 _GotoState32Action,
	{_State337, StringLiteralToken}:               _GotoState33Action,
	{_State337, IdentifierToken}:                  _GotoState23Action,
	{_State337, TrueToken}:                        _GotoState36Action,
	{_State337, FalseToken}:                       _GotoState19Action,
	{_State337, StructToken}:                      _GotoState34Action,
	{_State337, FuncToken}:                        _GotoState21Action,
	{_State337, VarToken}:                         _GotoState37Action,
	{_State337, LetToken}:                         _GotoState27Action,
	{_State337, NotToken}:                         _GotoState30Action,
	{_State337, LabelDeclToken}:                   _GotoState25Action,
	{_State337, LparenToken}:                      _GotoState28Action,
	{_State337, LbracketToken}:                    _GotoState26Action,
	{_State337, SubToken}:                         _GotoState35Action,
	{_State337, MulToken}:                         _GotoState29Action,
	{_State337, BitNegToken}:                      _GotoState18Action,
	{_State337, BitAndToken}:                      _GotoState17Action,
	{_State337, GreaterToken}:                     _GotoState22Action,
	{_State337, ParseErrorToken}:                  _GotoState31Action,
	{_State337, VarDeclPatternType}:               _GotoState56Action,
	{_State337, VarOrLetType}:                     _GotoState57Action,
	{_State337, ExpressionType}:                   _GotoState400Action,
	{_State337, OptionalLabelDeclType}:            _GotoState50Action,
	{_State337, SequenceExprType}:                 _GotoState55Action,
	{_State337, CallExprType}:                     _GotoState43Action,
	{_State337, AtomExprType}:                     _GotoState42Action,
	{_State337, LiteralType}:                      _GotoState48Action,
	{_State337, ImplicitStructExprType}:           _GotoState46Action,
	{_State337, AccessExprType}:                   _GotoState38Action,
	{_State337, PostfixUnaryExprType}:             _GotoState52Action,
	{_State337, PrefixUnaryOpType}:                _GotoState54Action,
	{_State337, PrefixUnaryExprType}:              _GotoState53Action,
	{_State337, MulExprType}:                      _GotoState49Action,
	{_State337, AddExprType}:                      _GotoState39Action,
	{_State337, CmpExprType}:                      _GotoState44Action,
	{_State337, AndExprType}:                      _GotoState40Action,
	{_State337, OrExprType}:                       _GotoState51Action,
	{_State337, InitializableTypeType}:            _GotoState47Action,
	{_State337, ExplicitStructDefType}:            _GotoState45Action,
	{_State337, AnonymousFuncExprType}:            _GotoState41Action,
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
	{_State338, LiteralType}:                      _GotoState48Action,
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
	{_State340, IntegerLiteralToken}:              _GotoState24Action,
	{_State340, FloatLiteralToken}:                _GotoState20Action,
	{_State340, RuneLiteralToken}:                 _GotoState32Action,
	{_State340, StringLiteralToken}:               _GotoState33Action,
	{_State340, IdentifierToken}:                  _GotoState23Action,
	{_State340, TrueToken}:                        _GotoState36Action,
	{_State340, FalseToken}:                       _GotoState19Action,
	{_State340, StructToken}:                      _GotoState34Action,
	{_State340, FuncToken}:                        _GotoState21Action,
	{_State340, VarToken}:                         _GotoState37Action,
	{_State340, LetToken}:                         _GotoState27Action,
	{_State340, NotToken}:                         _GotoState30Action,
	{_State340, LabelDeclToken}:                   _GotoState25Action,
	{_State340, LparenToken}:                      _GotoState28Action,
	{_State340, LbracketToken}:                    _GotoState26Action,
	{_State340, SubToken}:                         _GotoState35Action,
	{_State340, MulToken}:                         _GotoState29Action,
	{_State340, BitNegToken}:                      _GotoState18Action,
	{_State340, BitAndToken}:                      _GotoState17Action,
	{_State340, GreaterToken}:                     _GotoState22Action,
	{_State340, ParseErrorToken}:                  _GotoState31Action,
	{_State340, ExpressionsType}:                  _GotoState402Action,
	{_State340, OptionalExpressionsType}:          _GotoState403Action,
	{_State340, VarDeclPatternType}:               _GotoState56Action,
	{_State340, VarOrLetType}:                     _GotoState57Action,
	{_State340, ExpressionType}:                   _GotoState229Action,
	{_State340, OptionalLabelDeclType}:            _GotoState50Action,
	{_State340, SequenceExprType}:                 _GotoState55Action,
	{_State340, CallExprType}:                     _GotoState43Action,
	{_State340, AtomExprType}:                     _GotoState42Action,
	{_State340, LiteralType}:                      _GotoState48Action,
	{_State340, ImplicitStructExprType}:           _GotoState46Action,
	{_State340, AccessExprType}:                   _GotoState38Action,
	{_State340, PostfixUnaryExprType}:             _GotoState52Action,
	{_State340, PrefixUnaryOpType}:                _GotoState54Action,
	{_State340, PrefixUnaryExprType}:              _GotoState53Action,
	{_State340, MulExprType}:                      _GotoState49Action,
	{_State340, AddExprType}:                      _GotoState39Action,
	{_State340, CmpExprType}:                      _GotoState44Action,
	{_State340, AndExprType}:                      _GotoState40Action,
	{_State340, OrExprType}:                       _GotoState51Action,
	{_State340, InitializableTypeType}:            _GotoState47Action,
	{_State340, ExplicitStructDefType}:            _GotoState45Action,
	{_State340, AnonymousFuncExprType}:            _GotoState41Action,
	{_State343, AddToken}:                         _GotoState176Action,
	{_State343, SubToken}:                         _GotoState181Action,
	{_State343, MulToken}:                         _GotoState179Action,
	{_State344, IdentifierToken}:                  _GotoState241Action,
	{_State344, GenericParameterDefType}:          _GotoState404Action,
	{_State346, IdentifierToken}:                  _GotoState79Action,
	{_State346, StructToken}:                      _GotoState34Action,
	{_State346, EnumToken}:                        _GotoState76Action,
	{_State346, TraitToken}:                       _GotoState84Action,
	{_State346, FuncToken}:                        _GotoState78Action,
	{_State346, LparenToken}:                      _GotoState80Action,
	{_State346, LbracketToken}:                    _GotoState26Action,
	{_State346, DotToken}:                         _GotoState75Action,
	{_State346, QuestionToken}:                    _GotoState82Action,
	{_State346, ExclaimToken}:                     _GotoState77Action,
	{_State346, TildeTildeToken}:                  _GotoState83Action,
	{_State346, BitNegToken}:                      _GotoState74Action,
	{_State346, BitAndToken}:                      _GotoState73Action,
	{_State346, ParseErrorToken}:                  _GotoState81Action,
	{_State346, InitializableTypeType}:            _GotoState90Action,
	{_State346, AtomTypeType}:                     _GotoState85Action,
	{_State346, ReturnableTypeType}:               _GotoState91Action,
	{_State346, ValueTypeType}:                    _GotoState405Action,
	{_State346, ImplicitStructDefType}:            _GotoState89Action,
	{_State346, ExplicitStructDefType}:            _GotoState45Action,
	{_State346, ImplicitEnumDefType}:              _GotoState88Action,
	{_State346, ExplicitEnumDefType}:              _GotoState86Action,
	{_State346, TraitDefType}:                     _GotoState92Action,
	{_State346, FuncTypeType}:                     _GotoState87Action,
	{_State347, RparenToken}:                      _GotoState406Action,
	{_State348, AddToken}:                         _GotoState176Action,
	{_State348, SubToken}:                         _GotoState181Action,
	{_State348, MulToken}:                         _GotoState179Action,
	{_State349, DollarLbracketToken}:              _GotoState147Action,
	{_State349, OptionalGenericParametersType}:    _GotoState407Action,
	{_State350, LbraceToken}:                      _GotoState58Action,
	{_State350, StatementBlockType}:               _GotoState408Action,
	{_State353, IdentifierToken}:                  _GotoState164Action,
	{_State353, UnsafeToken}:                      _GotoState165Action,
	{_State353, StructToken}:                      _GotoState34Action,
	{_State353, EnumToken}:                        _GotoState76Action,
	{_State353, TraitToken}:                       _GotoState84Action,
	{_State353, FuncToken}:                        _GotoState78Action,
	{_State353, LparenToken}:                      _GotoState80Action,
	{_State353, LbracketToken}:                    _GotoState26Action,
	{_State353, DotToken}:                         _GotoState75Action,
	{_State353, QuestionToken}:                    _GotoState82Action,
	{_State353, ExclaimToken}:                     _GotoState77Action,
	{_State353, TildeTildeToken}:                  _GotoState83Action,
	{_State353, BitNegToken}:                      _GotoState74Action,
	{_State353, BitAndToken}:                      _GotoState73Action,
	{_State353, ParseErrorToken}:                  _GotoState81Action,
	{_State353, UnsafeStatementType}:              _GotoState171Action,
	{_State353, InitializableTypeType}:            _GotoState90Action,
	{_State353, AtomTypeType}:                     _GotoState85Action,
	{_State353, ReturnableTypeType}:               _GotoState91Action,
	{_State353, ValueTypeType}:                    _GotoState172Action,
	{_State353, FieldDefType}:                     _GotoState255Action,
	{_State353, ImplicitStructDefType}:            _GotoState89Action,
	{_State353, ExplicitStructDefType}:            _GotoState45Action,
	{_State353, EnumValueDefType}:                 _GotoState409Action,
	{_State353, ImplicitEnumDefType}:              _GotoState88Action,
	{_State353, ExplicitEnumDefType}:              _GotoState86Action,
	{_State353, TraitDefType}:                     _GotoState92Action,
	{_State353, FuncTypeType}:                     _GotoState87Action,
	{_State354, IdentifierToken}:                  _GotoState164Action,
	{_State354, UnsafeToken}:                      _GotoState165Action,
	{_State354, StructToken}:                      _GotoState34Action,
	{_State354, EnumToken}:                        _GotoState76Action,
	{_State354, TraitToken}:                       _GotoState84Action,
	{_State354, FuncToken}:                        _GotoState78Action,
	{_State354, LparenToken}:                      _GotoState80Action,
	{_State354, LbracketToken}:                    _GotoState26Action,
	{_State354, DotToken}:                         _GotoState75Action,
	{_State354, QuestionToken}:                    _GotoState82Action,
	{_State354, ExclaimToken}:                     _GotoState77Action,
	{_State354, TildeTildeToken}:                  _GotoState83Action,
	{_State354, BitNegToken}:                      _GotoState74Action,
	{_State354, BitAndToken}:                      _GotoState73Action,
	{_State354, ParseErrorToken}:                  _GotoState81Action,
	{_State354, UnsafeStatementType}:              _GotoState171Action,
	{_State354, InitializableTypeType}:            _GotoState90Action,
	{_State354, AtomTypeType}:                     _GotoState85Action,
	{_State354, ReturnableTypeType}:               _GotoState91Action,
	{_State354, ValueTypeType}:                    _GotoState172Action,
	{_State354, FieldDefType}:                     _GotoState255Action,
	{_State354, ImplicitStructDefType}:            _GotoState89Action,
	{_State354, ExplicitStructDefType}:            _GotoState45Action,
	{_State354, EnumValueDefType}:                 _GotoState410Action,
	{_State354, ImplicitEnumDefType}:              _GotoState88Action,
	{_State354, ExplicitEnumDefType}:              _GotoState86Action,
	{_State354, TraitDefType}:                     _GotoState92Action,
	{_State354, FuncTypeType}:                     _GotoState87Action,
	{_State356, IdentifierToken}:                  _GotoState164Action,
	{_State356, UnsafeToken}:                      _GotoState165Action,
	{_State356, StructToken}:                      _GotoState34Action,
	{_State356, EnumToken}:                        _GotoState76Action,
	{_State356, TraitToken}:                       _GotoState84Action,
	{_State356, FuncToken}:                        _GotoState78Action,
	{_State356, LparenToken}:                      _GotoState80Action,
	{_State356, LbracketToken}:                    _GotoState26Action,
	{_State356, DotToken}:                         _GotoState75Action,
	{_State356, QuestionToken}:                    _GotoState82Action,
	{_State356, ExclaimToken}:                     _GotoState77Action,
	{_State356, TildeTildeToken}:                  _GotoState83Action,
	{_State356, BitNegToken}:                      _GotoState74Action,
	{_State356, BitAndToken}:                      _GotoState73Action,
	{_State356, ParseErrorToken}:                  _GotoState81Action,
	{_State356, UnsafeStatementType}:              _GotoState171Action,
	{_State356, InitializableTypeType}:            _GotoState90Action,
	{_State356, AtomTypeType}:                     _GotoState85Action,
	{_State356, ReturnableTypeType}:               _GotoState91Action,
	{_State356, ValueTypeType}:                    _GotoState172Action,
	{_State356, FieldDefType}:                     _GotoState255Action,
	{_State356, ImplicitStructDefType}:            _GotoState89Action,
	{_State356, ExplicitStructDefType}:            _GotoState45Action,
	{_State356, EnumValueDefType}:                 _GotoState411Action,
	{_State356, ImplicitEnumDefType}:              _GotoState88Action,
	{_State356, ExplicitEnumDefType}:              _GotoState86Action,
	{_State356, TraitDefType}:                     _GotoState92Action,
	{_State356, FuncTypeType}:                     _GotoState87Action,
	{_State357, IdentifierToken}:                  _GotoState164Action,
	{_State357, UnsafeToken}:                      _GotoState165Action,
	{_State357, StructToken}:                      _GotoState34Action,
	{_State357, EnumToken}:                        _GotoState76Action,
	{_State357, TraitToken}:                       _GotoState84Action,
	{_State357, FuncToken}:                        _GotoState78Action,
	{_State357, LparenToken}:                      _GotoState80Action,
	{_State357, LbracketToken}:                    _GotoState26Action,
	{_State357, DotToken}:                         _GotoState75Action,
	{_State357, QuestionToken}:                    _GotoState82Action,
	{_State357, ExclaimToken}:                     _GotoState77Action,
	{_State357, TildeTildeToken}:                  _GotoState83Action,
	{_State357, BitNegToken}:                      _GotoState74Action,
	{_State357, BitAndToken}:                      _GotoState73Action,
	{_State357, ParseErrorToken}:                  _GotoState81Action,
	{_State357, UnsafeStatementType}:              _GotoState171Action,
	{_State357, InitializableTypeType}:            _GotoState90Action,
	{_State357, AtomTypeType}:                     _GotoState85Action,
	{_State357, ReturnableTypeType}:               _GotoState91Action,
	{_State357, ValueTypeType}:                    _GotoState172Action,
	{_State357, FieldDefType}:                     _GotoState255Action,
	{_State357, ImplicitStructDefType}:            _GotoState89Action,
	{_State357, ExplicitStructDefType}:            _GotoState45Action,
	{_State357, EnumValueDefType}:                 _GotoState412Action,
	{_State357, ImplicitEnumDefType}:              _GotoState88Action,
	{_State357, ExplicitEnumDefType}:              _GotoState86Action,
	{_State357, TraitDefType}:                     _GotoState92Action,
	{_State357, FuncTypeType}:                     _GotoState87Action,
	{_State358, AddToken}:                         _GotoState176Action,
	{_State358, SubToken}:                         _GotoState181Action,
	{_State358, MulToken}:                         _GotoState179Action,
	{_State359, IdentifierToken}:                  _GotoState79Action,
	{_State359, StructToken}:                      _GotoState34Action,
	{_State359, EnumToken}:                        _GotoState76Action,
	{_State359, TraitToken}:                       _GotoState84Action,
	{_State359, FuncToken}:                        _GotoState78Action,
	{_State359, LparenToken}:                      _GotoState80Action,
	{_State359, LbracketToken}:                    _GotoState26Action,
	{_State359, DotToken}:                         _GotoState75Action,
	{_State359, QuestionToken}:                    _GotoState82Action,
	{_State359, ExclaimToken}:                     _GotoState77Action,
	{_State359, TildeTildeToken}:                  _GotoState83Action,
	{_State359, BitNegToken}:                      _GotoState74Action,
	{_State359, BitAndToken}:                      _GotoState73Action,
	{_State359, ParseErrorToken}:                  _GotoState81Action,
	{_State359, InitializableTypeType}:            _GotoState90Action,
	{_State359, AtomTypeType}:                     _GotoState85Action,
	{_State359, ReturnableTypeType}:               _GotoState91Action,
	{_State359, ValueTypeType}:                    _GotoState413Action,
	{_State359, ImplicitStructDefType}:            _GotoState89Action,
	{_State359, ExplicitStructDefType}:            _GotoState45Action,
	{_State359, ImplicitEnumDefType}:              _GotoState88Action,
	{_State359, ExplicitEnumDefType}:              _GotoState86Action,
	{_State359, TraitDefType}:                     _GotoState92Action,
	{_State359, FuncTypeType}:                     _GotoState87Action,
	{_State360, AddToken}:                         _GotoState176Action,
	{_State360, SubToken}:                         _GotoState181Action,
	{_State360, MulToken}:                         _GotoState179Action,
	{_State361, IdentifierToken}:                  _GotoState79Action,
	{_State361, StructToken}:                      _GotoState34Action,
	{_State361, EnumToken}:                        _GotoState76Action,
	{_State361, TraitToken}:                       _GotoState84Action,
	{_State361, FuncToken}:                        _GotoState78Action,
	{_State361, LparenToken}:                      _GotoState80Action,
	{_State361, LbracketToken}:                    _GotoState26Action,
	{_State361, DotToken}:                         _GotoState75Action,
	{_State361, QuestionToken}:                    _GotoState82Action,
	{_State361, ExclaimToken}:                     _GotoState77Action,
	{_State361, TildeTildeToken}:                  _GotoState83Action,
	{_State361, BitNegToken}:                      _GotoState74Action,
	{_State361, BitAndToken}:                      _GotoState73Action,
	{_State361, ParseErrorToken}:                  _GotoState81Action,
	{_State361, InitializableTypeType}:            _GotoState90Action,
	{_State361, AtomTypeType}:                     _GotoState85Action,
	{_State361, ReturnableTypeType}:               _GotoState351Action,
	{_State361, ImplicitStructDefType}:            _GotoState89Action,
	{_State361, ExplicitStructDefType}:            _GotoState45Action,
	{_State361, ImplicitEnumDefType}:              _GotoState88Action,
	{_State361, ExplicitEnumDefType}:              _GotoState86Action,
	{_State361, TraitDefType}:                     _GotoState92Action,
	{_State361, ReturnTypeType}:                   _GotoState414Action,
	{_State361, FuncTypeType}:                     _GotoState87Action,
	{_State362, IdentifierToken}:                  _GotoState258Action,
	{_State362, StructToken}:                      _GotoState34Action,
	{_State362, EnumToken}:                        _GotoState76Action,
	{_State362, TraitToken}:                       _GotoState84Action,
	{_State362, FuncToken}:                        _GotoState78Action,
	{_State362, LparenToken}:                      _GotoState80Action,
	{_State362, LbracketToken}:                    _GotoState26Action,
	{_State362, DotToken}:                         _GotoState75Action,
	{_State362, QuestionToken}:                    _GotoState82Action,
	{_State362, ExclaimToken}:                     _GotoState77Action,
	{_State362, DotDotDotToken}:                   _GotoState257Action,
	{_State362, TildeTildeToken}:                  _GotoState83Action,
	{_State362, BitNegToken}:                      _GotoState74Action,
	{_State362, BitAndToken}:                      _GotoState73Action,
	{_State362, ParseErrorToken}:                  _GotoState81Action,
	{_State362, InitializableTypeType}:            _GotoState90Action,
	{_State362, AtomTypeType}:                     _GotoState85Action,
	{_State362, ReturnableTypeType}:               _GotoState91Action,
	{_State362, ValueTypeType}:                    _GotoState262Action,
	{_State362, ImplicitStructDefType}:            _GotoState89Action,
	{_State362, ExplicitStructDefType}:            _GotoState45Action,
	{_State362, ImplicitEnumDefType}:              _GotoState88Action,
	{_State362, ExplicitEnumDefType}:              _GotoState86Action,
	{_State362, TraitDefType}:                     _GotoState92Action,
	{_State362, ParameterDeclType}:                _GotoState415Action,
	{_State362, FuncTypeType}:                     _GotoState87Action,
	{_State364, GreaterToken}:                     _GotoState416Action,
	{_State369, LparenToken}:                      _GotoState417Action,
	{_State371, IdentifierToken}:                  _GotoState164Action,
	{_State371, UnsafeToken}:                      _GotoState165Action,
	{_State371, StructToken}:                      _GotoState34Action,
	{_State371, EnumToken}:                        _GotoState76Action,
	{_State371, TraitToken}:                       _GotoState84Action,
	{_State371, FuncToken}:                        _GotoState273Action,
	{_State371, LparenToken}:                      _GotoState80Action,
	{_State371, LbracketToken}:                    _GotoState26Action,
	{_State371, DotToken}:                         _GotoState75Action,
	{_State371, QuestionToken}:                    _GotoState82Action,
	{_State371, ExclaimToken}:                     _GotoState77Action,
	{_State371, TildeTildeToken}:                  _GotoState83Action,
	{_State371, BitNegToken}:                      _GotoState74Action,
	{_State371, BitAndToken}:                      _GotoState73Action,
	{_State371, ParseErrorToken}:                  _GotoState81Action,
	{_State371, UnsafeStatementType}:              _GotoState171Action,
	{_State371, InitializableTypeType}:            _GotoState90Action,
	{_State371, AtomTypeType}:                     _GotoState85Action,
	{_State371, ReturnableTypeType}:               _GotoState91Action,
	{_State371, ValueTypeType}:                    _GotoState172Action,
	{_State371, FieldDefType}:                     _GotoState274Action,
	{_State371, ImplicitStructDefType}:            _GotoState89Action,
	{_State371, ExplicitStructDefType}:            _GotoState45Action,
	{_State371, ImplicitEnumDefType}:              _GotoState88Action,
	{_State371, ExplicitEnumDefType}:              _GotoState86Action,
	{_State371, TraitPropertyType}:                _GotoState418Action,
	{_State371, TraitDefType}:                     _GotoState92Action,
	{_State371, FuncTypeType}:                     _GotoState87Action,
	{_State371, MethodSignatureType}:              _GotoState275Action,
	{_State372, IdentifierToken}:                  _GotoState164Action,
	{_State372, UnsafeToken}:                      _GotoState165Action,
	{_State372, StructToken}:                      _GotoState34Action,
	{_State372, EnumToken}:                        _GotoState76Action,
	{_State372, TraitToken}:                       _GotoState84Action,
	{_State372, FuncToken}:                        _GotoState273Action,
	{_State372, LparenToken}:                      _GotoState80Action,
	{_State372, LbracketToken}:                    _GotoState26Action,
	{_State372, DotToken}:                         _GotoState75Action,
	{_State372, QuestionToken}:                    _GotoState82Action,
	{_State372, ExclaimToken}:                     _GotoState77Action,
	{_State372, TildeTildeToken}:                  _GotoState83Action,
	{_State372, BitNegToken}:                      _GotoState74Action,
	{_State372, BitAndToken}:                      _GotoState73Action,
	{_State372, ParseErrorToken}:                  _GotoState81Action,
	{_State372, UnsafeStatementType}:              _GotoState171Action,
	{_State372, InitializableTypeType}:            _GotoState90Action,
	{_State372, AtomTypeType}:                     _GotoState85Action,
	{_State372, ReturnableTypeType}:               _GotoState91Action,
	{_State372, ValueTypeType}:                    _GotoState172Action,
	{_State372, FieldDefType}:                     _GotoState274Action,
	{_State372, ImplicitStructDefType}:            _GotoState89Action,
	{_State372, ExplicitStructDefType}:            _GotoState45Action,
	{_State372, ImplicitEnumDefType}:              _GotoState88Action,
	{_State372, ExplicitEnumDefType}:              _GotoState86Action,
	{_State372, TraitPropertyType}:                _GotoState419Action,
	{_State372, TraitDefType}:                     _GotoState92Action,
	{_State372, FuncTypeType}:                     _GotoState87Action,
	{_State372, MethodSignatureType}:              _GotoState275Action,
	{_State377, AddToken}:                         _GotoState176Action,
	{_State377, SubToken}:                         _GotoState181Action,
	{_State377, MulToken}:                         _GotoState179Action,
	{_State381, DoToken}:                          _GotoState420Action,
	{_State382, SemicolonToken}:                   _GotoState421Action,
	{_State385, LparenToken}:                      _GotoState28Action,
	{_State385, ImplicitStructExprType}:           _GotoState422Action,
	{_State386, IdentifierToken}:                  _GotoState423Action,
	{_State387, IntegerLiteralToken}:              _GotoState24Action,
	{_State387, FloatLiteralToken}:                _GotoState20Action,
	{_State387, RuneLiteralToken}:                 _GotoState32Action,
	{_State387, StringLiteralToken}:               _GotoState33Action,
	{_State387, IdentifierToken}:                  _GotoState23Action,
	{_State387, TrueToken}:                        _GotoState36Action,
	{_State387, FalseToken}:                       _GotoState19Action,
	{_State387, StructToken}:                      _GotoState34Action,
	{_State387, FuncToken}:                        _GotoState21Action,
	{_State387, VarToken}:                         _GotoState37Action,
	{_State387, LetToken}:                         _GotoState27Action,
	{_State387, NotToken}:                         _GotoState30Action,
	{_State387, LabelDeclToken}:                   _GotoState25Action,
	{_State387, LparenToken}:                      _GotoState28Action,
	{_State387, LbracketToken}:                    _GotoState26Action,
	{_State387, SubToken}:                         _GotoState35Action,
	{_State387, MulToken}:                         _GotoState29Action,
	{_State387, BitNegToken}:                      _GotoState18Action,
	{_State387, BitAndToken}:                      _GotoState17Action,
	{_State387, GreaterToken}:                     _GotoState22Action,
	{_State387, ParseErrorToken}:                  _GotoState31Action,
	{_State387, VarDeclPatternType}:               _GotoState56Action,
	{_State387, VarOrLetType}:                     _GotoState57Action,
	{_State387, OptionalLabelDeclType}:            _GotoState71Action,
	{_State387, SequenceExprType}:                 _GotoState424Action,
	{_State387, CallExprType}:                     _GotoState43Action,
	{_State387, AtomExprType}:                     _GotoState42Action,
	{_State387, LiteralType}:                      _GotoState48Action,
	{_State387, ImplicitStructExprType}:           _GotoState46Action,
	{_State387, AccessExprType}:                   _GotoState38Action,
	{_State387, PostfixUnaryExprType}:             _GotoState52Action,
	{_State387, PrefixUnaryOpType}:                _GotoState54Action,
	{_State387, PrefixUnaryExprType}:              _GotoState53Action,
	{_State387, MulExprType}:                      _GotoState49Action,
	{_State387, AddExprType}:                      _GotoState39Action,
	{_State387, CmpExprType}:                      _GotoState44Action,
	{_State387, AndExprType}:                      _GotoState40Action,
	{_State387, OrExprType}:                       _GotoState51Action,
	{_State387, InitializableTypeType}:            _GotoState47Action,
	{_State387, ExplicitStructDefType}:            _GotoState45Action,
	{_State387, AnonymousFuncExprType}:            _GotoState41Action,
	{_State388, IntegerLiteralToken}:              _GotoState24Action,
	{_State388, FloatLiteralToken}:                _GotoState20Action,
	{_State388, RuneLiteralToken}:                 _GotoState32Action,
	{_State388, StringLiteralToken}:               _GotoState33Action,
	{_State388, IdentifierToken}:                  _GotoState23Action,
	{_State388, TrueToken}:                        _GotoState36Action,
	{_State388, FalseToken}:                       _GotoState19Action,
	{_State388, StructToken}:                      _GotoState34Action,
	{_State388, FuncToken}:                        _GotoState21Action,
	{_State388, VarToken}:                         _GotoState304Action,
	{_State388, LetToken}:                         _GotoState27Action,
	{_State388, NotToken}:                         _GotoState30Action,
	{_State388, LabelDeclToken}:                   _GotoState25Action,
	{_State388, LparenToken}:                      _GotoState28Action,
	{_State388, LbracketToken}:                    _GotoState26Action,
	{_State388, DotToken}:                         _GotoState303Action,
	{_State388, SubToken}:                         _GotoState35Action,
	{_State388, MulToken}:                         _GotoState29Action,
	{_State388, BitNegToken}:                      _GotoState18Action,
	{_State388, BitAndToken}:                      _GotoState17Action,
	{_State388, GreaterToken}:                     _GotoState22Action,
	{_State388, ParseErrorToken}:                  _GotoState31Action,
	{_State388, VarDeclPatternType}:               _GotoState56Action,
	{_State388, VarOrLetType}:                     _GotoState57Action,
	{_State388, AssignPatternType}:                _GotoState305Action,
	{_State388, CasePatternType}:                  _GotoState425Action,
	{_State388, OptionalLabelDeclType}:            _GotoState71Action,
	{_State388, SequenceExprType}:                 _GotoState308Action,
	{_State388, CallExprType}:                     _GotoState43Action,
	{_State388, AtomExprType}:                     _GotoState42Action,
	{_State388, LiteralType}:                      _GotoState48Action,
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
	{_State389, IfToken}:                          _GotoState130Action,
	{_State389, LbraceToken}:                      _GotoState58Action,
	{_State389, StatementBlockType}:               _GotoState427Action,
	{_State389, IfExprType}:                       _GotoState426Action,
	{_State392, IntegerLiteralToken}:              _GotoState24Action,
	{_State392, FloatLiteralToken}:                _GotoState20Action,
	{_State392, RuneLiteralToken}:                 _GotoState32Action,
	{_State392, StringLiteralToken}:               _GotoState33Action,
	{_State392, IdentifierToken}:                  _GotoState23Action,
	{_State392, TrueToken}:                        _GotoState36Action,
	{_State392, FalseToken}:                       _GotoState19Action,
	{_State392, ReturnToken}:                      _GotoState226Action,
	{_State392, BreakToken}:                       _GotoState218Action,
	{_State392, ContinueToken}:                    _GotoState220Action,
	{_State392, FallthroughToken}:                 _GotoState223Action,
	{_State392, UnsafeToken}:                      _GotoState165Action,
	{_State392, StructToken}:                      _GotoState34Action,
	{_State392, FuncToken}:                        _GotoState21Action,
	{_State392, AsyncToken}:                       _GotoState217Action,
	{_State392, DeferToken}:                       _GotoState222Action,
	{_State392, VarToken}:                         _GotoState37Action,
	{_State392, LetToken}:                         _GotoState27Action,
	{_State392, NotToken}:                         _GotoState30Action,
	{_State392, LabelDeclToken}:                   _GotoState25Action,
	{_State392, LparenToken}:                      _GotoState28Action,
	{_State392, LbracketToken}:                    _GotoState26Action,
	{_State392, SubToken}:                         _GotoState35Action,
	{_State392, MulToken}:                         _GotoState29Action,
	{_State392, BitNegToken}:                      _GotoState18Action,
	{_State392, BitAndToken}:                      _GotoState17Action,
	{_State392, GreaterToken}:                     _GotoState22Action,
	{_State392, ParseErrorToken}:                  _GotoState31Action,
	{_State392, SimpleStatementBodyType}:          _GotoState394Action,
	{_State392, OptionalSimpleStatementBodyType}:  _GotoState428Action,
	{_State392, UnsafeStatementType}:              _GotoState237Action,
	{_State392, JumpTypeType}:                     _GotoState232Action,
	{_State392, ExpressionsType}:                  _GotoState230Action,
	{_State392, VarDeclPatternType}:               _GotoState56Action,
	{_State392, VarOrLetType}:                     _GotoState57Action,
	{_State392, AssignPatternType}:                _GotoState228Action,
	{_State392, ExpressionType}:                   _GotoState229Action,
	{_State392, OptionalLabelDeclType}:            _GotoState50Action,
	{_State392, SequenceExprType}:                 _GotoState233Action,
	{_State392, CallExprType}:                     _GotoState43Action,
	{_State392, AtomExprType}:                     _GotoState42Action,
	{_State392, LiteralType}:                      _GotoState48Action,
	{_State392, ImplicitStructExprType}:           _GotoState46Action,
	{_State392, AccessExprType}:                   _GotoState227Action,
	{_State392, PostfixUnaryExprType}:             _GotoState52Action,
	{_State392, PrefixUnaryOpType}:                _GotoState54Action,
	{_State392, PrefixUnaryExprType}:              _GotoState53Action,
	{_State392, MulExprType}:                      _GotoState49Action,
	{_State392, AddExprType}:                      _GotoState39Action,
	{_State392, CmpExprType}:                      _GotoState44Action,
	{_State392, AndExprType}:                      _GotoState40Action,
	{_State392, OrExprType}:                       _GotoState51Action,
	{_State392, InitializableTypeType}:            _GotoState47Action,
	{_State392, ExplicitStructDefType}:            _GotoState45Action,
	{_State392, AnonymousFuncExprType}:            _GotoState41Action,
	{_State395, NewlinesToken}:                    _GotoState430Action,
	{_State395, CommaToken}:                       _GotoState429Action,
	{_State397, StringLiteralToken}:               _GotoState320Action,
	{_State397, RparenToken}:                      _GotoState431Action,
	{_State397, ImportClauseType}:                 _GotoState395Action,
	{_State397, ImportClauseTerminalType}:         _GotoState432Action,
	{_State398, IdentifierToken}:                  _GotoState433Action,
	{_State402, CommaToken}:                       _GotoState338Action,
	{_State405, AddToken}:                         _GotoState176Action,
	{_State405, SubToken}:                         _GotoState181Action,
	{_State405, MulToken}:                         _GotoState179Action,
	{_State406, IdentifierToken}:                  _GotoState79Action,
	{_State406, StructToken}:                      _GotoState34Action,
	{_State406, EnumToken}:                        _GotoState76Action,
	{_State406, TraitToken}:                       _GotoState84Action,
	{_State406, FuncToken}:                        _GotoState78Action,
	{_State406, LparenToken}:                      _GotoState80Action,
	{_State406, LbracketToken}:                    _GotoState26Action,
	{_State406, DotToken}:                         _GotoState75Action,
	{_State406, QuestionToken}:                    _GotoState82Action,
	{_State406, ExclaimToken}:                     _GotoState77Action,
	{_State406, TildeTildeToken}:                  _GotoState83Action,
	{_State406, BitNegToken}:                      _GotoState74Action,
	{_State406, BitAndToken}:                      _GotoState73Action,
	{_State406, ParseErrorToken}:                  _GotoState81Action,
	{_State406, InitializableTypeType}:            _GotoState90Action,
	{_State406, AtomTypeType}:                     _GotoState85Action,
	{_State406, ReturnableTypeType}:               _GotoState351Action,
	{_State406, ImplicitStructDefType}:            _GotoState89Action,
	{_State406, ExplicitStructDefType}:            _GotoState45Action,
	{_State406, ImplicitEnumDefType}:              _GotoState88Action,
	{_State406, ExplicitEnumDefType}:              _GotoState86Action,
	{_State406, TraitDefType}:                     _GotoState92Action,
	{_State406, ReturnTypeType}:                   _GotoState434Action,
	{_State406, FuncTypeType}:                     _GotoState87Action,
	{_State407, LparenToken}:                      _GotoState435Action,
	{_State413, AddToken}:                         _GotoState176Action,
	{_State413, SubToken}:                         _GotoState181Action,
	{_State413, MulToken}:                         _GotoState179Action,
	{_State416, StringLiteralToken}:               _GotoState436Action,
	{_State417, IdentifierToken}:                  _GotoState258Action,
	{_State417, StructToken}:                      _GotoState34Action,
	{_State417, EnumToken}:                        _GotoState76Action,
	{_State417, TraitToken}:                       _GotoState84Action,
	{_State417, FuncToken}:                        _GotoState78Action,
	{_State417, LparenToken}:                      _GotoState80Action,
	{_State417, LbracketToken}:                    _GotoState26Action,
	{_State417, DotToken}:                         _GotoState75Action,
	{_State417, QuestionToken}:                    _GotoState82Action,
	{_State417, ExclaimToken}:                     _GotoState77Action,
	{_State417, DotDotDotToken}:                   _GotoState257Action,
	{_State417, TildeTildeToken}:                  _GotoState83Action,
	{_State417, BitNegToken}:                      _GotoState74Action,
	{_State417, BitAndToken}:                      _GotoState73Action,
	{_State417, ParseErrorToken}:                  _GotoState81Action,
	{_State417, InitializableTypeType}:            _GotoState90Action,
	{_State417, AtomTypeType}:                     _GotoState85Action,
	{_State417, ReturnableTypeType}:               _GotoState91Action,
	{_State417, ValueTypeType}:                    _GotoState262Action,
	{_State417, ImplicitStructDefType}:            _GotoState89Action,
	{_State417, ExplicitStructDefType}:            _GotoState45Action,
	{_State417, ImplicitEnumDefType}:              _GotoState88Action,
	{_State417, ExplicitEnumDefType}:              _GotoState86Action,
	{_State417, TraitDefType}:                     _GotoState92Action,
	{_State417, ParameterDeclType}:                _GotoState260Action,
	{_State417, ParameterDeclsType}:               _GotoState261Action,
	{_State417, OptionalParameterDeclsType}:       _GotoState437Action,
	{_State417, FuncTypeType}:                     _GotoState87Action,
	{_State420, LbraceToken}:                      _GotoState58Action,
	{_State420, StatementBlockType}:               _GotoState438Action,
	{_State421, IntegerLiteralToken}:              _GotoState24Action,
	{_State421, FloatLiteralToken}:                _GotoState20Action,
	{_State421, RuneLiteralToken}:                 _GotoState32Action,
	{_State421, StringLiteralToken}:               _GotoState33Action,
	{_State421, IdentifierToken}:                  _GotoState23Action,
	{_State421, TrueToken}:                        _GotoState36Action,
	{_State421, FalseToken}:                       _GotoState19Action,
	{_State421, StructToken}:                      _GotoState34Action,
	{_State421, FuncToken}:                        _GotoState21Action,
	{_State421, VarToken}:                         _GotoState37Action,
	{_State421, LetToken}:                         _GotoState27Action,
	{_State421, NotToken}:                         _GotoState30Action,
	{_State421, LabelDeclToken}:                   _GotoState25Action,
	{_State421, LparenToken}:                      _GotoState28Action,
	{_State421, LbracketToken}:                    _GotoState26Action,
	{_State421, SubToken}:                         _GotoState35Action,
	{_State421, MulToken}:                         _GotoState29Action,
	{_State421, BitNegToken}:                      _GotoState18Action,
	{_State421, BitAndToken}:                      _GotoState17Action,
	{_State421, GreaterToken}:                     _GotoState22Action,
	{_State421, ParseErrorToken}:                  _GotoState31Action,
	{_State421, VarDeclPatternType}:               _GotoState56Action,
	{_State421, VarOrLetType}:                     _GotoState57Action,
	{_State421, OptionalLabelDeclType}:            _GotoState71Action,
	{_State421, SequenceExprType}:                 _GotoState383Action,
	{_State421, OptionalSequenceExprType}:         _GotoState439Action,
	{_State421, CallExprType}:                     _GotoState43Action,
	{_State421, AtomExprType}:                     _GotoState42Action,
	{_State421, LiteralType}:                      _GotoState48Action,
	{_State421, ImplicitStructExprType}:           _GotoState46Action,
	{_State421, AccessExprType}:                   _GotoState38Action,
	{_State421, PostfixUnaryExprType}:             _GotoState52Action,
	{_State421, PrefixUnaryOpType}:                _GotoState54Action,
	{_State421, PrefixUnaryExprType}:              _GotoState53Action,
	{_State421, MulExprType}:                      _GotoState49Action,
	{_State421, AddExprType}:                      _GotoState39Action,
	{_State421, CmpExprType}:                      _GotoState44Action,
	{_State421, AndExprType}:                      _GotoState40Action,
	{_State421, OrExprType}:                       _GotoState51Action,
	{_State421, InitializableTypeType}:            _GotoState47Action,
	{_State421, ExplicitStructDefType}:            _GotoState45Action,
	{_State421, AnonymousFuncExprType}:            _GotoState41Action,
	{_State423, LparenToken}:                      _GotoState139Action,
	{_State423, TuplePatternType}:                 _GotoState440Action,
	{_State434, LbraceToken}:                      _GotoState58Action,
	{_State434, StatementBlockType}:               _GotoState441Action,
	{_State435, IdentifierToken}:                  _GotoState151Action,
	{_State435, ParameterDefType}:                 _GotoState154Action,
	{_State435, ParameterDefsType}:                _GotoState155Action,
	{_State435, OptionalParameterDefsType}:        _GotoState442Action,
	{_State437, RparenToken}:                      _GotoState443Action,
	{_State439, DoToken}:                          _GotoState444Action,
	{_State442, RparenToken}:                      _GotoState445Action,
	{_State443, IdentifierToken}:                  _GotoState79Action,
	{_State443, StructToken}:                      _GotoState34Action,
	{_State443, EnumToken}:                        _GotoState76Action,
	{_State443, TraitToken}:                       _GotoState84Action,
	{_State443, FuncToken}:                        _GotoState78Action,
	{_State443, LparenToken}:                      _GotoState80Action,
	{_State443, LbracketToken}:                    _GotoState26Action,
	{_State443, DotToken}:                         _GotoState75Action,
	{_State443, QuestionToken}:                    _GotoState82Action,
	{_State443, ExclaimToken}:                     _GotoState77Action,
	{_State443, TildeTildeToken}:                  _GotoState83Action,
	{_State443, BitNegToken}:                      _GotoState74Action,
	{_State443, BitAndToken}:                      _GotoState73Action,
	{_State443, ParseErrorToken}:                  _GotoState81Action,
	{_State443, InitializableTypeType}:            _GotoState90Action,
	{_State443, AtomTypeType}:                     _GotoState85Action,
	{_State443, ReturnableTypeType}:               _GotoState351Action,
	{_State443, ImplicitStructDefType}:            _GotoState89Action,
	{_State443, ExplicitStructDefType}:            _GotoState45Action,
	{_State443, ImplicitEnumDefType}:              _GotoState88Action,
	{_State443, ExplicitEnumDefType}:              _GotoState86Action,
	{_State443, TraitDefType}:                     _GotoState92Action,
	{_State443, ReturnTypeType}:                   _GotoState446Action,
	{_State443, FuncTypeType}:                     _GotoState87Action,
	{_State444, LbraceToken}:                      _GotoState58Action,
	{_State444, StatementBlockType}:               _GotoState447Action,
	{_State445, IdentifierToken}:                  _GotoState79Action,
	{_State445, StructToken}:                      _GotoState34Action,
	{_State445, EnumToken}:                        _GotoState76Action,
	{_State445, TraitToken}:                       _GotoState84Action,
	{_State445, FuncToken}:                        _GotoState78Action,
	{_State445, LparenToken}:                      _GotoState80Action,
	{_State445, LbracketToken}:                    _GotoState26Action,
	{_State445, DotToken}:                         _GotoState75Action,
	{_State445, QuestionToken}:                    _GotoState82Action,
	{_State445, ExclaimToken}:                     _GotoState77Action,
	{_State445, TildeTildeToken}:                  _GotoState83Action,
	{_State445, BitNegToken}:                      _GotoState74Action,
	{_State445, BitAndToken}:                      _GotoState73Action,
	{_State445, ParseErrorToken}:                  _GotoState81Action,
	{_State445, InitializableTypeType}:            _GotoState90Action,
	{_State445, AtomTypeType}:                     _GotoState85Action,
	{_State445, ReturnableTypeType}:               _GotoState351Action,
	{_State445, ImplicitStructDefType}:            _GotoState89Action,
	{_State445, ExplicitStructDefType}:            _GotoState45Action,
	{_State445, ImplicitEnumDefType}:              _GotoState88Action,
	{_State445, ExplicitEnumDefType}:              _GotoState86Action,
	{_State445, TraitDefType}:                     _GotoState92Action,
	{_State445, ReturnTypeType}:                   _GotoState448Action,
	{_State445, FuncTypeType}:                     _GotoState87Action,
	{_State448, LbraceToken}:                      _GotoState58Action,
	{_State448, StatementBlockType}:               _GotoState449Action,
	{_State1, _EndMarker}:                         _ReduceNilToOptionalDefinitionsAction,
	{_State1, _WildcardMarker}:                    _ReduceNilToOptionalNewlinesAction,
	{_State5, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State11, _EndMarker}:                        _ReduceNewlinesToOptionalDefinitionsAction,
	{_State11, _WildcardMarker}:                   _ReduceNewlinesToOptionalNewlinesAction,
	{_State12, _EndMarker}:                        _ReduceToSourceAction,
	{_State14, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State17, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryOpAction,
	{_State18, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryOpAction,
	{_State19, _WildcardMarker}:                   _ReduceFalseToLiteralAction,
	{_State20, _WildcardMarker}:                   _ReduceFloatLiteralToLiteralAction,
	{_State22, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State23, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State24, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralAction,
	{_State25, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State27, _WildcardMarker}:                   _ReduceLetToVarOrLetAction,
	{_State28, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State28, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State29, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State30, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State31, _WildcardMarker}:                   _ReduceParseErrorToAtomExprAction,
	{_State32, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State33, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State35, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State36, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
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
	{_State48, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State49, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State51, _WildcardMarker}:                   _ReduceOrExprToSequenceExprAction,
	{_State52, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State53, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State54, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State55, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State56, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State58, _WildcardMarker}:                   _ReduceEmptyListToStatementsAction,
	{_State59, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State60, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State61, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State62, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State63, _WildcardMarker}:                   _ReduceStatementBlockToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceGlobalVarDefToDefinitionAction,
	{_State66, _EndMarker}:                        _ReduceWithSpecToPackageDefAction,
	{_State67, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State68, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State70, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State72, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State75, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State79, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State80, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State81, _WildcardMarker}:                   _ReduceParseErrorToAtomTypeAction,
	{_State85, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State86, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State87, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State88, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State89, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State90, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State91, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State92, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
	{_State94, _WildcardMarker}:                   _ReduceDotDotDotToArgumentAction,
	{_State95, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State96, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State98, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State99, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State99, ColonToken}:                        _ReduceExpressionToOptionalExpressionAction,
	{_State101, RparenToken}:                      _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State102, RbracketToken}:                    _ReduceNilToOptionalGenericArgumentsAction,
	{_State104, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State104, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State105, _WildcardMarker}:                  _ReduceQuestionToPostfixUnaryExprAction,
	{_State107, _WildcardMarker}:                  _ReduceAddToAddOpAction,
	{_State108, _WildcardMarker}:                  _ReduceBitOrToAddOpAction,
	{_State109, _WildcardMarker}:                  _ReduceBitXorToAddOpAction,
	{_State110, _WildcardMarker}:                  _ReduceSubToAddOpAction,
	{_State111, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State112, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State113, _WildcardMarker}:                  _ReduceEqualToCmpOpAction,
	{_State114, _WildcardMarker}:                  _ReduceGreaterToCmpOpAction,
	{_State115, _WildcardMarker}:                  _ReduceGreaterOrEqualToCmpOpAction,
	{_State116, _WildcardMarker}:                  _ReduceLessToCmpOpAction,
	{_State117, _WildcardMarker}:                  _ReduceLessOrEqualToCmpOpAction,
	{_State118, _WildcardMarker}:                  _ReduceNotEqualToCmpOpAction,
	{_State119, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State120, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State120, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State121, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State122, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State123, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State124, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State125, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State126, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State127, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State129, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State130, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State131, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State132, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State133, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State134, _WildcardMarker}:                  _ReduceBlockExprToAtomExprAction,
	{_State135, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State136, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State137, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State138, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State140, _WildcardMarker}:                  _ReduceTuplePatternToVarPatternAction,
	{_State141, _WildcardMarker}:                  _ReduceNilToOptionalValueTypeAction,
	{_State142, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State143, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State144, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State145, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State147, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State149, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State151, _WildcardMarker}:                  _ReduceInferredRefArgToParameterDefAction,
	{_State154, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State155, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State156, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State157, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State158, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State160, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State161, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State163, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State164, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State167, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State167, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State169, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State171, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State172, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State173, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State174, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State175, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State180, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
	{_State182, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State183, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State183, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State184, _WildcardMarker}:                  _ReduceToImplicitStructExprAction,
	{_State185, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State185, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State185, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State185, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State185, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State186, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State186, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State186, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State186, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State186, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State187, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State188, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State190, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State192, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State193, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State195, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State195, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State195, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State196, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State197, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State198, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State200, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State201, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State204, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State204, SemicolonToken}:                   _ReduceSequenceExprToForAssignmentAction,
	{_State205, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State207, LbraceToken}:                      _ReduceSequenceExprToConditionAction,
	{_State209, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State210, _WildcardMarker}:                  _ReduceDotDotDotToFieldVarPatternAction,
	{_State211, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State212, _WildcardMarker}:                  _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State214, _WildcardMarker}:                  _ReducePositionalToFieldVarPatternAction,
	{_State215, _EndMarker}:                       _ReduceToVarDeclPatternAction,
	{_State216, _WildcardMarker}:                  _ReduceValueTypeToOptionalValueTypeAction,
	{_State217, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State218, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State219, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State220, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State222, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, _WildcardMarker}:                  _ReduceFallthroughToSimpleStatementBodyAction,
	{_State225, _EndMarker}:                       _ReduceToStatementBlockAction,
	{_State226, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State227, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State227, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State229, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State230, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToSimpleStatementBodyAction,
	{_State231, _WildcardMarker}:                  _ReduceImportStatementToStatementBodyAction,
	{_State232, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State233, AssignToken}:                      _ReduceToAssignPatternAction,
	{_State233, _WildcardMarker}:                  _ReduceSequenceExprToExpressionAction,
	{_State234, _WildcardMarker}:                  _ReduceSimpleStatementBodyToStatementBodyAction,
	{_State235, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State237, _WildcardMarker}:                  _ReduceUnsafeStatementToSimpleStatementBodyAction,
	{_State238, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State239, _WildcardMarker}:                  _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State240, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State241, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State242, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State243, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State245, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State246, _EndMarker}:                       _ReduceAliasToNamedFuncDefAction,
	{_State247, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State248, _WildcardMarker}:                  _ReduceInferredRefVarargToParameterDefAction,
	{_State249, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State251, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State255, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State258, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State260, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State261, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State262, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State263, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State264, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State265, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State270, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State272, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State274, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State275, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State277, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State278, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State279, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State282, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State283, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State284, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State285, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State286, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State287, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State288, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State291, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State293, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State294, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State295, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State297, _WildcardMarker}:                  _ReduceInitializeExprToAtomExprAction,
	{_State298, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State299, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State300, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State301, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State301, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State304, _WildcardMarker}:                  _ReduceVarToVarOrLetAction,
	{_State305, _WildcardMarker}:                  _ReduceAssignPatternToCasePatternAction,
	{_State306, _WildcardMarker}:                  _ReduceCasePatternToCasePatternsAction,
	{_State308, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State309, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State310, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State313, _WildcardMarker}:                  _ReduceToTuplePatternAction,
	{_State314, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State315, NewlinesToken}:                    _ReduceAsyncToSimpleStatementBodyAction,
	{_State315, SemicolonToken}:                   _ReduceAsyncToSimpleStatementBodyAction,
	{_State315, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State317, NewlinesToken}:                    _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State317, SemicolonToken}:                   _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State317, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State318, NewlinesToken}:                    _ReduceDeferToSimpleStatementBodyAction,
	{_State318, SemicolonToken}:                   _ReduceDeferToSimpleStatementBodyAction,
	{_State318, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State320, _WildcardMarker}:                  _ReduceStringLiteralToImportClauseAction,
	{_State321, _WildcardMarker}:                  _ReduceSingleToImportStatementAction,
	{_State322, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State323, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State324, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State325, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State326, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State327, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State328, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State329, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State330, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State331, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State332, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State333, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State334, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State335, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State336, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToSimpleStatementBodyAction,
	{_State337, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State338, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State339, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State340, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State340, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State340, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State341, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State342, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State343, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State345, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State348, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State349, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State351, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State352, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State355, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State358, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State360, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State361, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State363, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State365, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State366, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State367, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State368, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State370, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State373, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State374, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State375, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State376, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State377, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State378, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State379, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State380, SemicolonToken}:                   _ReduceAssignToForAssignmentAction,
	{_State383, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State384, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State387, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State388, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State390, _WildcardMarker}:                  _ReduceNamedToFieldVarPatternAction,
	{_State391, _WildcardMarker}:                  _ReduceAddToFieldVarPatternsAction,
	{_State392, NewlinesToken}:                    _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State392, SemicolonToken}:                   _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State392, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State393, _WildcardMarker}:                  _ReduceDefaultBranchStatementToStatementBodyAction,
	{_State394, _WildcardMarker}:                  _ReduceSimpleStatementBodyToOptionalSimpleStatementBodyAction,
	{_State396, _WildcardMarker}:                  _ReduceFirstToImportClausesAction,
	{_State399, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToSimpleStatementBodyAction,
	{_State400, _WildcardMarker}:                  _ReduceAssignStatementToSimpleStatementBodyAction,
	{_State401, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State402, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State403, _WildcardMarker}:                  _ReduceJumpStatementToSimpleStatementBodyAction,
	{_State404, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State405, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State406, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State408, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State409, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State410, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State410, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State411, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State412, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State412, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State413, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State414, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State415, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State417, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State418, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State419, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State421, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State421, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State422, _WildcardMarker}:                  _ReduceEnumMatchPatternToCasePatternAction,
	{_State424, LbraceToken}:                      _ReduceCaseToConditionAction,
	{_State425, _WildcardMarker}:                  _ReduceMultipleToCasePatternsAction,
	{_State426, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State427, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State428, _WildcardMarker}:                  _ReduceCaseBranchStatementToStatementBodyAction,
	{_State429, _WildcardMarker}:                  _ReduceExplicitToImportClauseTerminalAction,
	{_State430, _WildcardMarker}:                  _ReduceImplicitToImportClauseTerminalAction,
	{_State431, _WildcardMarker}:                  _ReduceMultipleToImportStatementAction,
	{_State432, _WildcardMarker}:                  _ReduceAddToImportClausesAction,
	{_State433, _WildcardMarker}:                  _ReduceAliasToImportClauseAction,
	{_State435, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State436, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State438, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State440, _WildcardMarker}:                  _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State441, _EndMarker}:                       _ReduceFuncDefToNamedFuncDefAction,
	{_State443, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State445, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State446, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State447, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State449, _EndMarker}:                       _ReduceMethodDefToNamedFuncDefAction,
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
      literal -> State 48
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
      PACKAGE -> State 14
      TYPE -> State 15
      FUNC -> State 16
      VAR -> State 37
      LET -> State 27
      LBRACE -> State 58
      definitions -> State 60
      definition -> State 59
      statement_block -> State 63
      var_decl_pattern -> State 65
      var_or_let -> State 57
      type_def -> State 64
      named_func_def -> State 61
      package_def -> State 62

  State 14:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    Goto:
      LBRACE -> State 58
      statement_block -> State 66

  State 15:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 67

  State 16:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 68
      LPAREN -> State 69

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
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 70

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
      optional_label_decl -> State 71
      sequence_expr -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
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
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 93
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 27:
    Kernel Items:
      var_or_let: LET., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 95
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
      DOT_DOT_DOT -> State 94
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 99
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      arguments -> State 97
      argument -> State 96
      colon_expressions -> State 98
      optional_expression -> State 100
      atom_expr -> State 42
      literal -> State 48
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
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 101

  State 35:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
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
      LBRACKET -> State 104
      DOT -> State 103
      QUESTION -> State 105
      DOLLAR_LBRACKET -> State 102
      optional_generic_binding -> State 106

  State 39:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 107
      SUB -> State 110
      BIT_XOR -> State 109
      BIT_OR -> State 108
      add_op -> State 111

  State 40:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 112

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
      EQUAL -> State 113
      NOT_EQUAL -> State 118
      LESS -> State 116
      LESS_OR_EQUAL -> State 117
      GREATER -> State 114
      GREATER_OR_EQUAL -> State 115
      cmp_op -> State 119

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
      atom_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 120

  State 48:
    Kernel Items:
      atom_expr: literal., *
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
      MUL -> State 126
      DIV -> State 124
      MOD -> State 125
      BIT_AND -> State 121
      BIT_LSHIFT -> State 122
      BIT_RSHIFT -> State 123
      mul_op -> State 127

  State 50:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      atom_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 130
      SWITCH -> State 131
      FOR -> State 129
      DO -> State 128
      LBRACE -> State 58
      statement_block -> State 134
      if_expr -> State 132
      switch_expr -> State 135
      loop_expr -> State 133

  State 51:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 136

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
      optional_label_decl -> State 71
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 137
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
      IDENTIFIER -> State 138
      LPAREN -> State 139
      var_pattern -> State 141
      tuple_pattern -> State 140

  State 58:
    Kernel Items:
      statement_block: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 142

  State 59:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 143
      optional_newlines -> State 144

  State 61:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      definition: statement_block., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 145

  State 66:
    Kernel Items:
      package_def: PACKAGE statement_block., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 147
      ASSIGN -> State 146
      optional_generic_parameters -> State 148

  State 68:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 147
      ASSIGN -> State 149
      optional_generic_parameters -> State 150

  State 69:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 151
      parameter_def -> State 152

  State 70:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 151
      parameter_def -> State 154
      parameter_defs -> State 155
      optional_parameter_defs -> State 153

  State 71:
    Kernel Items:
      atom_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 134

  State 72:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 156
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 74:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 157
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 75:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 102
      optional_generic_binding -> State 158

  State 76:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 159

  State 77:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 160
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 78:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 161

  State 79:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 162
      DOLLAR_LBRACKET -> State 102
      optional_generic_binding -> State 163

  State 80:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 167
      implicit_field_defs -> State 169
      optional_implicit_field_defs -> State 170
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 166
      implicit_enum_value_defs -> State 168
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 81:
    Kernel Items:
      atom_type: PARSE_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 173
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 83:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 174
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 84:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 175

  State 85:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 93:
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
      RBRACKET -> State 180
      COMMA -> State 178
      COLON -> State 177
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 94:
    Kernel Items:
      argument: DOT_DOT_DOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 182

  State 96:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      arguments: arguments.COMMA argument
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 184
      COMMA -> State 183

  State 98:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 185

  State 99:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 186

  State 101:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 188
      implicit_struct_def -> State 89
      explicit_field_defs -> State 187
      optional_explicit_field_defs -> State 189
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 102:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      optional_generic_arguments -> State 191
      generic_arguments -> State 190
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 192
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 103:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 193

  State 104:
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
      IDENTIFIER -> State 95
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
      DOT_DOT_DOT -> State 94
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 99
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      argument -> State 194
      colon_expressions -> State 98
      optional_expression -> State 100
      atom_expr -> State 42
      literal -> State 48
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

  State 105:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 195

  State 107:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 111:
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
      optional_label_decl -> State 71
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 196
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 112:
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
      optional_label_decl -> State 71
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 197
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 113:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 119:
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
      optional_label_decl -> State 71
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 198
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 120:
    Kernel Items:
      atom_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 95
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
      DOT_DOT_DOT -> State 94
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 99
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      arguments -> State 199
      argument -> State 96
      colon_expressions -> State 98
      optional_expression -> State 100
      atom_expr -> State 42
      literal -> State 48
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

  State 121:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 127:
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
      optional_label_decl -> State 71
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 200
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 128:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 201

  State 129:
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
      assign_pattern -> State 202
      optional_label_decl -> State 71
      sequence_expr -> State 204
      for_assignment -> State 203
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 130:
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
      CASE -> State 205
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
      optional_label_decl -> State 71
      sequence_expr -> State 207
      condition -> State 206
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 131:
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
      optional_label_decl -> State 71
      sequence_expr -> State 208
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      atom_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 136:
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
      optional_label_decl -> State 71
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 209
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 137:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 211
      LPAREN -> State 139
      DOT_DOT_DOT -> State 210
      var_pattern -> State 214
      tuple_pattern -> State 140
      field_var_patterns -> State 213
      field_var_pattern -> State 212

  State 140:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      optional_value_type -> State 215
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 216
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 142:
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
      CASE -> State 219
      DEFAULT -> State 221
      RETURN -> State 226
      BREAK -> State 218
      CONTINUE -> State 220
      FALLTHROUGH -> State 223
      IMPORT -> State 224
      UNSAFE -> State 165
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 217
      DEFER -> State 222
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      RBRACE -> State 225
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      statement -> State 235
      simple_statement_body -> State 234
      statement_body -> State 236
      unsafe_statement -> State 237
      jump_type -> State 232
      expressions -> State 230
      import_statement -> State 231
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 228
      expression -> State 229
      optional_label_decl -> State 50
      sequence_expr -> State 233
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 227
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

  State 143:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      PACKAGE -> State 14
      TYPE -> State 15
      FUNC -> State 16
      VAR -> State 37
      LET -> State 27
      LBRACE -> State 58
      definition -> State 238
      statement_block -> State 63
      var_decl_pattern -> State 65
      var_or_let -> State 57
      type_def -> State 64
      named_func_def -> State 61
      package_def -> State 62

  State 144:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 145:
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
      expression -> State 239
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 146:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 240
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 147:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 241
      generic_parameter_def -> State 242
      generic_parameter_defs -> State 243
      optional_generic_parameter_defs -> State 244

  State 148:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 245
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 149:
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
      expression -> State 246
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 150:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 247

  State 151:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.DOT_DOT_DOT
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      DOT_DOT_DOT -> State 248
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 249
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 152:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 250

  State 153:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 251

  State 154:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 252

  State 156:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 253
      implicit_enum_value_defs -> State 256
      implicit_enum_def -> State 88
      explicit_enum_value_defs -> State 254
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 160:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 258
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      DOT_DOT_DOT -> State 257
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 262
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      parameter_decl -> State 260
      parameter_decls -> State 261
      optional_parameter_decls -> State 259
      func_type -> State 87

  State 162:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 263

  State 163:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 264
      QUESTION -> State 82
      EXCLAIM -> State 77
      DOLLAR_LBRACKET -> State 102
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      optional_generic_binding -> State 163
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 265
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 165:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 266

  State 166:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 267

  State 167:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 268

  State 168:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      OR -> State 269
      RPAREN -> State 270

  State 169:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 271

  State 170:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 272

  State 171:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 173:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 273
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 274
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_property -> State 278
      trait_properties -> State 277
      optional_trait_properties -> State 276
      trait_def -> State 92
      func_type -> State 87
      method_signature -> State 275

  State 176:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 279
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 177:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 280
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 178:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 281

  State 179:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 282
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 180:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 283
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 182:
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
      expression -> State 284
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 183:
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
      IDENTIFIER -> State 95
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
      DOT_DOT_DOT -> State 94
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 99
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      argument -> State 285
      colon_expressions -> State 98
      optional_expression -> State 100
      atom_expr -> State 42
      literal -> State 48
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

  State 184:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 185:
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
      expression -> State 286
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_expression -> State 287
      atom_expr -> State 42
      literal -> State 48
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
      expression -> State 286
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_expression -> State 288
      atom_expr -> State 42
      literal -> State 48
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
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 290
      COMMA -> State 289

  State 188:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 291

  State 190:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 292

  State 191:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 293

  State 192:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 193:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 294

  State 195:
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
      IDENTIFIER -> State 95
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
      DOT_DOT_DOT -> State 94
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 99
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_arguments -> State 296
      arguments -> State 295
      argument -> State 96
      colon_expressions -> State 98
      optional_expression -> State 100
      atom_expr -> State 42
      literal -> State 48
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

  State 196:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 126
      DIV -> State 124
      MOD -> State 125
      BIT_AND -> State 121
      BIT_LSHIFT -> State 122
      BIT_RSHIFT -> State 123
      mul_op -> State 127

  State 197:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 113
      NOT_EQUAL -> State 118
      LESS -> State 116
      LESS_OR_EQUAL -> State 117
      GREATER -> State 114
      GREATER_OR_EQUAL -> State 115
      cmp_op -> State 119

  State 198:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 107
      SUB -> State 110
      BIT_XOR -> State 109
      BIT_OR -> State 108
      add_op -> State 111

  State 199:
    Kernel Items:
      arguments: arguments.COMMA argument
      atom_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 297
      COMMA -> State 183

  State 200:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 298

  State 202:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 300
      ASSIGN -> State 299

  State 203:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 301

  State 204:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 302

  State 205:
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
      VAR -> State 304
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 303
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      case_patterns -> State 307
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 305
      case_pattern -> State 306
      optional_label_decl -> State 71
      sequence_expr -> State 308
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 206:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 309

  State 207:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 310

  State 209:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 112

  State 210:
    Kernel Items:
      field_var_pattern: DOT_DOT_DOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 311

  State 212:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 313
      COMMA -> State 312

  State 214:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      optional_value_type: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 217:
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
      optional_label_decl -> State 71
      call_expr -> State 315
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 314
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 218:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 219:
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
      VAR -> State 304
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 303
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      case_patterns -> State 316
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 305
      case_pattern -> State 306
      optional_label_decl -> State 71
      sequence_expr -> State 308
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 220:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      statement_body: DEFAULT.COLON optional_simple_statement_body
    Reduce:
      (nil)
    Goto:
      COLON -> State 317

  State 222:
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
      optional_label_decl -> State 71
      call_expr -> State 318
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 314
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 223:
    Kernel Items:
      simple_statement_body: FALLTHROUGH., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 320
      LPAREN -> State 319
      import_clause -> State 321

  State 225:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 227:
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
      LBRACKET -> State 104
      DOT -> State 103
      QUESTION -> State 105
      DOLLAR_LBRACKET -> State 102
      ADD_ASSIGN -> State 322
      SUB_ASSIGN -> State 333
      MUL_ASSIGN -> State 332
      DIV_ASSIGN -> State 330
      MOD_ASSIGN -> State 331
      ADD_ONE_ASSIGN -> State 323
      SUB_ONE_ASSIGN -> State 334
      BIT_NEG_ASSIGN -> State 326
      BIT_AND_ASSIGN -> State 324
      BIT_OR_ASSIGN -> State 327
      BIT_XOR_ASSIGN -> State 329
      BIT_LSHIFT_ASSIGN -> State 325
      BIT_RSHIFT_ASSIGN -> State 328
      unary_op_assign -> State 336
      binary_op_assign -> State 335
      optional_generic_binding -> State 106

  State 228:
    Kernel Items:
      simple_statement_body: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 337

  State 229:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      simple_statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [simple_statement_body]
    Goto:
      COMMA -> State 338

  State 231:
    Kernel Items:
      statement_body: import_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      simple_statement_body: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 339
      optional_jump_label -> State 340

  State 233:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      statement_body: simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 341
      SEMICOLON -> State 342

  State 237:
    Kernel Items:
      simple_statement_body: unsafe_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 241:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 343
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 242:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 344

  State 244:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 345

  State 245:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 346
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 246:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 151
      parameter_def -> State 154
      parameter_defs -> State 155
      optional_parameter_defs -> State 347

  State 248:
    Kernel Items:
      parameter_def: IDENTIFIER DOT_DOT_DOT., *
      parameter_def: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 348
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 249:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 250:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 349

  State 251:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 351
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      return_type -> State 350
      func_type -> State 87

  State 252:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 151
      parameter_def -> State 352

  State 253:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 353
      OR -> State 354

  State 254:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 355

  State 255:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 268

  State 256:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 356
      OR -> State 357

  State 257:
    Kernel Items:
      parameter_decl: DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 358
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 258:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 264
      QUESTION -> State 82
      EXCLAIM -> State 77
      DOLLAR_LBRACKET -> State 102
      DOT_DOT_DOT -> State 359
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      optional_generic_binding -> State 163
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 360
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 259:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 361

  State 260:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 362

  State 262:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 263:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 102
      optional_generic_binding -> State 363

  State 264:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 263
      DOLLAR_LBRACKET -> State 102
      optional_generic_binding -> State 158

  State 265:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 266:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 364

  State 267:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 365
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 268:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 366

  State 269:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 367
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 270:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 368
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 272:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 369
      LPAREN -> State 161

  State 274:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 370

  State 277:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 372
      COMMA -> State 371

  State 278:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 373
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 281:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 374

  State 282:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 375
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 290:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 376
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 291:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 377
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 293:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 183

  State 296:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 378

  State 297:
    Kernel Items:
      atom_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 298:
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
      optional_label_decl -> State 71
      sequence_expr -> State 379
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 299:
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
      optional_label_decl -> State 71
      sequence_expr -> State 380
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      optional_label_decl -> State 71
      sequence_expr -> State 381
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      optional_label_decl -> State 71
      sequence_expr -> State 383
      optional_sequence_expr -> State 382
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 384

  State 303:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 385

  State 304:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 386

  State 305:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 388
      ASSIGN -> State 387

  State 308:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 389

  State 310:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 138
      LPAREN -> State 139
      var_pattern -> State 390
      tuple_pattern -> State 140

  State 312:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 211
      LPAREN -> State 139
      DOT_DOT_DOT -> State 210
      var_pattern -> State 214
      tuple_pattern -> State 140
      field_var_pattern -> State 391

  State 313:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 104
      DOT -> State 103
      DOLLAR_LBRACKET -> State 102
      optional_generic_binding -> State 106

  State 315:
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

  State 316:
    Kernel Items:
      statement_body: CASE case_patterns.COLON optional_simple_statement_body
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 388
      COLON -> State 392

  State 317:
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
      RETURN -> State 226
      BREAK -> State 218
      CONTINUE -> State 220
      FALLTHROUGH -> State 223
      UNSAFE -> State 165
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 217
      DEFER -> State 222
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
      simple_statement_body -> State 394
      optional_simple_statement_body -> State 393
      unsafe_statement -> State 237
      jump_type -> State 232
      expressions -> State 230
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 228
      expression -> State 229
      optional_label_decl -> State 50
      sequence_expr -> State 233
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 227
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

  State 318:
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

  State 319:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 320
      import_clause -> State 395
      import_clause_terminal -> State 396
      import_clauses -> State 397

  State 320:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 398

  State 321:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 335:
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
      expression -> State 399
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 336:
    Kernel Items:
      simple_statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 337:
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
      expression -> State 400
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 338:
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
      expression -> State 401
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 340:
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
      expressions -> State 402
      optional_expressions -> State 403
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 229
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 341:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 344:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      generic_parameter_def -> State 404

  State 345:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 405
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 347:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 406

  State 348:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 349:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 147
      optional_generic_parameters -> State 407

  State 350:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 408

  State 351:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 352:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 409
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 354:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 410
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 355:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 411
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 357:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 255
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      enum_value_def -> State 412
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 358:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 359:
    Kernel Items:
      parameter_decl: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 413
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      func_type -> State 87

  State 360:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 361:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 351
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      return_type -> State 414
      func_type -> State 87

  State 362:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 258
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      DOT_DOT_DOT -> State 257
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 262
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      parameter_decl -> State 415
      func_type -> State 87

  State 363:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 364:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 416

  State 365:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 366:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 367:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 368:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 417

  State 370:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 273
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 274
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_property -> State 418
      trait_def -> State 92
      func_type -> State 87
      method_signature -> State 275

  State 372:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 164
      UNSAFE -> State 165
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 273
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      unsafe_statement -> State 171
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 172
      field_def -> State 274
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_property -> State 419
      trait_def -> State 92
      func_type -> State 87
      method_signature -> State 275

  State 373:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 377:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 378:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 420

  State 382:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 421

  State 383:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 385:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 28
      implicit_struct_expr -> State 422

  State 386:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 423

  State 387:
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
      optional_label_decl -> State 71
      sequence_expr -> State 424
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 388:
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
      VAR -> State 304
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 303
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 305
      case_pattern -> State 425
      optional_label_decl -> State 71
      sequence_expr -> State 308
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 130
      LBRACE -> State 58
      statement_block -> State 427
      if_expr -> State 426

  State 390:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 392:
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
      RETURN -> State 226
      BREAK -> State 218
      CONTINUE -> State 220
      FALLTHROUGH -> State 223
      UNSAFE -> State 165
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 217
      DEFER -> State 222
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
      simple_statement_body -> State 394
      optional_simple_statement_body -> State 428
      unsafe_statement -> State 237
      jump_type -> State 232
      expressions -> State 230
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 228
      expression -> State 229
      optional_label_decl -> State 50
      sequence_expr -> State 233
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 227
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

  State 393:
    Kernel Items:
      statement_body: DEFAULT COLON optional_simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 394:
    Kernel Items:
      optional_simple_statement_body: simple_statement_body., *
    Reduce:
      * -> [optional_simple_statement_body]
    Goto:
      (nil)

  State 395:
    Kernel Items:
      import_clause_terminal: import_clause.NEWLINES
      import_clause_terminal: import_clause.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 430
      COMMA -> State 429

  State 396:
    Kernel Items:
      import_clauses: import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 397:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
      import_clauses: import_clauses.import_clause_terminal
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 320
      RPAREN -> State 431
      import_clause -> State 395
      import_clause_terminal -> State 432

  State 398:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 433

  State 399:
    Kernel Items:
      simple_statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      simple_statement_body: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 338

  State 403:
    Kernel Items:
      simple_statement_body: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 404:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 406:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 351
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      return_type -> State 434
      func_type -> State 87

  State 407:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 435

  State 408:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 409:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 410:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 411:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 413:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 176
      SUB -> State 181
      MUL -> State 179

  State 414:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 436

  State 417:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 258
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      DOT_DOT_DOT -> State 257
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 91
      value_type -> State 262
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      parameter_decl -> State 260
      parameter_decls -> State 261
      optional_parameter_decls -> State 437
      func_type -> State 87

  State 418:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 438

  State 421:
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
      optional_label_decl -> State 71
      sequence_expr -> State 383
      optional_sequence_expr -> State 439
      call_expr -> State 43
      atom_expr -> State 42
      literal -> State 48
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

  State 422:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 139
      tuple_pattern -> State 440

  State 424:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      statement_body: CASE case_patterns COLON optional_simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      import_clause_terminal: import_clause COMMA., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      import_clause_terminal: import_clause NEWLINES., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      import_clauses: import_clauses import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 433:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 441

  State 435:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 151
      parameter_def -> State 154
      parameter_defs -> State 155
      optional_parameter_defs -> State 442

  State 436:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 437:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 443

  State 438:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 444

  State 440:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 445

  State 443:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 351
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      return_type -> State 446
      func_type -> State 87

  State 444:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 447

  State 445:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 79
      STRUCT -> State 34
      ENUM -> State 76
      TRAIT -> State 84
      FUNC -> State 78
      LPAREN -> State 80
      LBRACKET -> State 26
      DOT -> State 75
      QUESTION -> State 82
      EXCLAIM -> State 77
      TILDE_TILDE -> State 83
      BIT_NEG -> State 74
      BIT_AND -> State 73
      PARSE_ERROR -> State 81
      initializable_type -> State 90
      atom_type -> State 85
      returnable_type -> State 351
      implicit_struct_def -> State 89
      explicit_struct_def -> State 45
      implicit_enum_def -> State 88
      explicit_enum_def -> State 86
      trait_def -> State 92
      return_type -> State 448
      func_type -> State 87

  State 446:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 447:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 58
      statement_block -> State 449

  State 449:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 449
Number of shift actions: 3170
Number of reduce actions: 367
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 3935
Number of unoptimized shift actions: 29898
Number of unoptimized reduce actions: 22419
*/
