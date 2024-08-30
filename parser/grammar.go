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
	PackageToken         = SymbolId(275)
	ImportToken          = SymbolId(276)
	AsToken              = SymbolId(277)
	UnsafeToken          = SymbolId(278)
	TypeToken            = SymbolId(279)
	ImplementsToken      = SymbolId(280)
	StructToken          = SymbolId(281)
	EnumToken            = SymbolId(282)
	TraitToken           = SymbolId(283)
	FuncToken            = SymbolId(284)
	AsyncToken           = SymbolId(285)
	DeferToken           = SymbolId(286)
	VarToken             = SymbolId(287)
	LetToken             = SymbolId(288)
	LabelDeclToken       = SymbolId(289)
	JumpLabelToken       = SymbolId(290)
	LbraceToken          = SymbolId(291)
	RbraceToken          = SymbolId(292)
	LparenToken          = SymbolId(293)
	RparenToken          = SymbolId(294)
	LbracketToken        = SymbolId(295)
	RbracketToken        = SymbolId(296)
	DotToken             = SymbolId(297)
	CommaToken           = SymbolId(298)
	QuestionToken        = SymbolId(299)
	SemicolonToken       = SymbolId(300)
	ColonToken           = SymbolId(301)
	ExclaimToken         = SymbolId(302)
	DollarLbracketToken  = SymbolId(303)
	DotDotDotToken       = SymbolId(304)
	TildeTildeToken      = SymbolId(305)
	AssignToken          = SymbolId(306)
	AddAssignToken       = SymbolId(307)
	SubAssignToken       = SymbolId(308)
	MulAssignToken       = SymbolId(309)
	DivAssignToken       = SymbolId(310)
	ModAssignToken       = SymbolId(311)
	AddOneAssignToken    = SymbolId(312)
	SubOneAssignToken    = SymbolId(313)
	BitNegAssignToken    = SymbolId(314)
	BitAndAssignToken    = SymbolId(315)
	BitOrAssignToken     = SymbolId(316)
	BitXorAssignToken    = SymbolId(317)
	BitLshiftAssignToken = SymbolId(318)
	BitRshiftAssignToken = SymbolId(319)
	NotToken             = SymbolId(320)
	AndToken             = SymbolId(321)
	OrToken              = SymbolId(322)
	AddToken             = SymbolId(323)
	SubToken             = SymbolId(324)
	MulToken             = SymbolId(325)
	DivToken             = SymbolId(326)
	ModToken             = SymbolId(327)
	BitNegToken          = SymbolId(328)
	BitAndToken          = SymbolId(329)
	BitXorToken          = SymbolId(330)
	BitOrToken           = SymbolId(331)
	BitLshiftToken       = SymbolId(332)
	BitRshiftToken       = SymbolId(333)
	EqualToken           = SymbolId(334)
	NotEqualToken        = SymbolId(335)
	LessToken            = SymbolId(336)
	LessOrEqualToken     = SymbolId(337)
	GreaterToken         = SymbolId(338)
	GreaterOrEqualToken  = SymbolId(339)
	ParseErrorToken      = SymbolId(340)
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
}

type GenericSymbol struct {
	SymbolId
	Location
}

func (t *GenericSymbol) Id() SymbolId { return t.SymbolId }

func (t *GenericSymbol) Loc() Location { return t.Location }

type Lexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *GenericSymbol
	Next() (Token, error)

	CurrentLocation() Location
}

type Reducer interface {
	// 49:10: source -> ...
	ToSource(OptionalDefinitions_ *GenericSymbol) (*GenericSymbol, error)

	// 52:2: optional_definitions -> NEWLINES: ...
	NewlinesToOptionalDefinitions(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 53:2: optional_definitions -> definitions: ...
	DefinitionsToOptionalDefinitions(OptionalNewlines_ *GenericSymbol, Definitions_ *GenericSymbol, OptionalNewlines_2 *GenericSymbol) (*GenericSymbol, error)

	// 54:2: optional_definitions -> nil: ...
	NilToOptionalDefinitions() (*GenericSymbol, error)

	// 57:2: optional_newlines -> NEWLINES: ...
	NewlinesToOptionalNewlines(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 58:2: optional_newlines -> nil: ...
	NilToOptionalNewlines() (*GenericSymbol, error)

	// 61:2: definitions -> nil: ...
	NilToDefinitions(Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 62:2: definitions -> add: ...
	AddToDefinitions(Definitions_ *GenericSymbol, Newlines_ *GenericSymbol, Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 66:2: definition -> package_def: ...
	PackageDefToDefinition(PackageDef_ *GenericSymbol) (*GenericSymbol, error)

	// 67:2: definition -> type_def: ...
	TypeDefToDefinition(TypeDef_ *GenericSymbol) (*GenericSymbol, error)

	// 68:2: definition -> named_func_def: ...
	NamedFuncDefToDefinition(NamedFuncDef_ *GenericSymbol) (*GenericSymbol, error)

	// 69:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 70:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 73:2: definition -> block_body: ...
	BlockBodyToDefinition(BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 82:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ *GenericSymbol, VarPattern_ *GenericSymbol, OptionalValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 84:14: var_or_let -> VAR: ...
	VarToVarOrLet(Var_ *GenericSymbol) (*GenericSymbol, error)

	// 84:20: var_or_let -> LET: ...
	LetToVarOrLet(Let_ *GenericSymbol) (*GenericSymbol, error)

	// 87:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 88:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 90:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ *GenericSymbol, FieldVarPatterns_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 93:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 94:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ *GenericSymbol, Comma_ *GenericSymbol, FieldVarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 97:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 98:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, VarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 99:2: field_var_pattern -> DOT_DOT_DOT: ...
	DotDotDotToFieldVarPattern(DotDotDot_ *GenericSymbol) (*GenericSymbol, error)

	// 102:2: optional_value_type -> value_type: ...
	ValueTypeToOptionalValueType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 103:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (*GenericSymbol, error)

	// 109:18: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: case_pattern -> assign_pattern: ...
	AssignPatternToCasePattern(AssignPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ *GenericSymbol, Identifier_ *GenericSymbol, ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol, TuplePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 134:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 143:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, Condition_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 144:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, Condition_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 145:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, Condition_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 148:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 149:2: condition -> case: ...
	CaseToCondition(Case_ *GenericSymbol, CasePatterns_ *GenericSymbol, Assign_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 173:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, CaseBranches_ *GenericSymbol, OptionalDefaultBranch_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 176:2: case_branches -> case_branch: ...
	CaseBranchToCaseBranches(CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: case_branches -> add: ...
	AddToCaseBranches(CaseBranches_ *GenericSymbol, CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 179:15: case_branch -> ...
	ToCaseBranch(Case_ *GenericSymbol, CasePatterns_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 183:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ *GenericSymbol, Comma_ *GenericSymbol, CasePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: optional_default_branch -> default_branch: ...
	DefaultBranchToOptionalDefaultBranch(DefaultBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: optional_default_branch -> nil: ...
	NilToOptionalDefaultBranch() (*GenericSymbol, error)

	// 189:18: default_branch -> ...
	ToDefaultBranch(Default_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol, For_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ *GenericSymbol, AssignPattern_ *GenericSymbol, In_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: loop_expr -> for: ...
	ForToLoopExpr(For_ *GenericSymbol, ForAssignment_ *GenericSymbol, Semicolon_ *GenericSymbol, OptionalSequenceExpr_ *GenericSymbol, Semicolon_2 *GenericSymbol, OptionalSequenceExpr_2 *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (*GenericSymbol, error)

	// 214:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 215:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ *GenericSymbol, Assign_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: sequence_expr -> or_expr: ...
	OrExprToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 231:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 233:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 236:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 237:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 240:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 241:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 244:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 252:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 270:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 273:2: statement_body -> assign_statement: ...
	AssignStatementToStatementBody(AssignPattern_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 277:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 278:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 285:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 291:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 295:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 296:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 299:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 307:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 314:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 317:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 318:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 322:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 326:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 327:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 337:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 341:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 344:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 349:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 350:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 354:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 357:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 358:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 363:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 366:2: argument -> DOT_DOT_DOT: ...
	DotDotDotToArgument(DotDotDot_ *GenericSymbol) (*GenericSymbol, error)

	// 369:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 370:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 374:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 384:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 386:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 387:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializableType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: atom_expr -> implicit_struct_expr: ...
	ImplicitStructExprToAtomExpr(ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: atom_expr -> PARSE_ERROR: ...
	ParseErrorToAtomExpr(ParseError_ *GenericSymbol) (*GenericSymbol, error)

	// 393:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 397:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 398:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 400:24: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 410:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 414:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 415:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 418:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 421:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 425:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 428:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 429:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 430:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 431:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 432:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 436:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 437:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 440:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 441:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 442:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 443:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 447:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 451:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 452:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 453:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 455:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 462:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 463:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 466:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 467:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 476:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 477:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 478:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Comma_ *GenericSymbol, IntegerLiteral_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 479:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Colon_ *GenericSymbol, ValueType_2 *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 482:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ *GenericSymbol) (*GenericSymbol, error)

	// 483:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_2 *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 486:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 487:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 490:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 491:2: atom_type -> PARSE_ERROR: ...
	ParseErrorToAtomType(ParseError_ *GenericSymbol) (*GenericSymbol, error)

	// 497:2: returnable_type -> atom_type: ...
	AtomTypeToReturnableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 498:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 499:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 500:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 501:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 502:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: value_type -> returnable_type: ...
	ReturnableTypeToValueType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 508:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ *GenericSymbol, Mul_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 509:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ *GenericSymbol, Add_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ *GenericSymbol, Sub_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 513:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 514:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 515:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Assign_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 523:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 524:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 527:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 528:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 531:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 532:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 535:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 536:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 543:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 544:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 545:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 548:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 549:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 552:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 553:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 555:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 558:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 559:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 560:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 563:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 564:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 566:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 574:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 575:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 587:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 588:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 590:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 593:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 594:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 595:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 596:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 598:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 605:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 606:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 609:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 610:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 611:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 614:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 615:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 617:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 625:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 626:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 629:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 630:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, DotDotDot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 631:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 632:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(DotDotDot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 635:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 636:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 639:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 640:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 642:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 653:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 659:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 660:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ *GenericSymbol, DotDotDot_ *GenericSymbol) (*GenericSymbol, error)

	// 661:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 662:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, DotDotDot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 665:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 666:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 669:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 670:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 673:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 674:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ *GenericSymbol, Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_2 *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_2 *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 675:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 679:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 691:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol) (*GenericSymbol, error)

	// 692:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 695:2: package_statement_body -> unsafe_statement: ...
	UnsafeStatementToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 696:2: package_statement_body -> import_statement: ...
	ImportStatementToPackageStatementBody(ImportStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 699:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 700:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 703:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 704:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 707:2: import_statement -> single: ...
	SingleToImportStatement(Import_ *GenericSymbol, ImportClause_ *GenericSymbol) (*GenericSymbol, error)

	// 708:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ *GenericSymbol, Lparen_ *GenericSymbol, ImportClauses_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 711:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 712:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ *GenericSymbol, As_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 715:2: import_clause_terminal -> implicit: ...
	ImplicitToImportClauseTerminal(ImportClause_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 716:2: import_clause_terminal -> explicit: ...
	ExplicitToImportClauseTerminal(ImportClause_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 719:2: import_clauses -> first: ...
	FirstToImportClauses(ImportClauseTerminal_ *GenericSymbol) (*GenericSymbol, error)

	// 720:2: import_clauses -> add: ...
	AddToImportClauses(ImportClauses_ *GenericSymbol, ImportClauseTerminal_ *GenericSymbol) (*GenericSymbol, error)
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

func ParseSource(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseSourceWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseSourceWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParsePackageDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParsePackageDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParsePackageDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseTypeDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseTypeDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTypeDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State3)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseNamedFuncDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseNamedFuncDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseNamedFuncDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State4)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseExpression(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseExpressionWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseExpressionWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State5)
	if err != nil {
		var errRetVal *GenericSymbol
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
	case NotToken:
		return "NOT"
	case AndToken:
		return "AND"
	case OrToken:
		return "OR"
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
	case IfExprType:
		return "if_expr"
	case ConditionType:
		return "condition"
	case SwitchExprType:
		return "switch_expr"
	case CaseBranchesType:
		return "case_branches"
	case CaseBranchType:
		return "case_branch"
	case CasePatternsType:
		return "case_patterns"
	case OptionalDefaultBranchType:
		return "optional_default_branch"
	case DefaultBranchType:
		return "default_branch"
	case LoopExprType:
		return "loop_expr"
	case OptionalSequenceExprType:
		return "optional_sequence_expr"
	case ForAssignmentType:
		return "for_assignment"
	case SequenceExprType:
		return "sequence_expr"
	case BlockExprType:
		return "block_expr"
	case BlockBodyType:
		return "block_body"
	case StatementsType:
		return "statements"
	case StatementType:
		return "statement"
	case StatementBodyType:
		return "statement_body"
	case UnaryOpAssignType:
		return "unary_op_assign"
	case BinaryOpAssignType:
		return "binary_op_assign"
	case UnsafeStatementType:
		return "unsafe_statement"
	case JumpStatementType:
		return "jump_statement"
	case JumpTypeType:
		return "jump_type"
	case OptionalJumpLabelType:
		return "optional_jump_label"
	case ExpressionsType:
		return "expressions"
	case OptionalExpressionsType:
		return "optional_expressions"
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
	case PackageStatementBodyType:
		return "package_statement_body"
	case PackageStatementType:
		return "package_statement"
	case PackageStatementsType:
		return "package_statements"
	case ImportStatementType:
		return "import_statement"
	case ImportClauseType:
		return "import_clause"
	case ImportClauseTerminalType:
		return "import_clause_terminal"
	case ImportClausesType:
		return "import_clauses"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	SourceType                       = SymbolId(341)
	OptionalDefinitionsType          = SymbolId(342)
	OptionalNewlinesType             = SymbolId(343)
	DefinitionsType                  = SymbolId(344)
	DefinitionType                   = SymbolId(345)
	VarDeclPatternType               = SymbolId(346)
	VarOrLetType                     = SymbolId(347)
	VarPatternType                   = SymbolId(348)
	TuplePatternType                 = SymbolId(349)
	FieldVarPatternsType             = SymbolId(350)
	FieldVarPatternType              = SymbolId(351)
	OptionalValueTypeType            = SymbolId(352)
	AssignPatternType                = SymbolId(353)
	CasePatternType                  = SymbolId(354)
	ExpressionType                   = SymbolId(355)
	OptionalLabelDeclType            = SymbolId(356)
	IfExprType                       = SymbolId(357)
	ConditionType                    = SymbolId(358)
	SwitchExprType                   = SymbolId(359)
	CaseBranchesType                 = SymbolId(360)
	CaseBranchType                   = SymbolId(361)
	CasePatternsType                 = SymbolId(362)
	OptionalDefaultBranchType        = SymbolId(363)
	DefaultBranchType                = SymbolId(364)
	LoopExprType                     = SymbolId(365)
	OptionalSequenceExprType         = SymbolId(366)
	ForAssignmentType                = SymbolId(367)
	SequenceExprType                 = SymbolId(368)
	BlockExprType                    = SymbolId(369)
	BlockBodyType                    = SymbolId(370)
	StatementsType                   = SymbolId(371)
	StatementType                    = SymbolId(372)
	StatementBodyType                = SymbolId(373)
	UnaryOpAssignType                = SymbolId(374)
	BinaryOpAssignType               = SymbolId(375)
	UnsafeStatementType              = SymbolId(376)
	JumpStatementType                = SymbolId(377)
	JumpTypeType                     = SymbolId(378)
	OptionalJumpLabelType            = SymbolId(379)
	ExpressionsType                  = SymbolId(380)
	OptionalExpressionsType          = SymbolId(381)
	CallExprType                     = SymbolId(382)
	OptionalGenericBindingType       = SymbolId(383)
	OptionalGenericArgumentsType     = SymbolId(384)
	GenericArgumentsType             = SymbolId(385)
	OptionalArgumentsType            = SymbolId(386)
	ArgumentsType                    = SymbolId(387)
	ArgumentType                     = SymbolId(388)
	ColonExpressionsType             = SymbolId(389)
	OptionalExpressionType           = SymbolId(390)
	AtomExprType                     = SymbolId(391)
	LiteralType                      = SymbolId(392)
	ImplicitStructExprType           = SymbolId(393)
	AccessExprType                   = SymbolId(394)
	PostfixUnaryExprType             = SymbolId(395)
	PrefixUnaryOpType                = SymbolId(396)
	PrefixUnaryExprType              = SymbolId(397)
	MulOpType                        = SymbolId(398)
	MulExprType                      = SymbolId(399)
	AddOpType                        = SymbolId(400)
	AddExprType                      = SymbolId(401)
	CmpOpType                        = SymbolId(402)
	CmpExprType                      = SymbolId(403)
	AndExprType                      = SymbolId(404)
	OrExprType                       = SymbolId(405)
	InitializableTypeType            = SymbolId(406)
	AtomTypeType                     = SymbolId(407)
	ReturnableTypeType               = SymbolId(408)
	ValueTypeType                    = SymbolId(409)
	TypeDefType                      = SymbolId(410)
	GenericParameterDefType          = SymbolId(411)
	GenericParameterDefsType         = SymbolId(412)
	OptionalGenericParameterDefsType = SymbolId(413)
	OptionalGenericParametersType    = SymbolId(414)
	FieldDefType                     = SymbolId(415)
	ImplicitFieldDefsType            = SymbolId(416)
	OptionalImplicitFieldDefsType    = SymbolId(417)
	ImplicitStructDefType            = SymbolId(418)
	ExplicitFieldDefsType            = SymbolId(419)
	OptionalExplicitFieldDefsType    = SymbolId(420)
	ExplicitStructDefType            = SymbolId(421)
	EnumValueDefType                 = SymbolId(422)
	ImplicitEnumValueDefsType        = SymbolId(423)
	ImplicitEnumDefType              = SymbolId(424)
	ExplicitEnumValueDefsType        = SymbolId(425)
	ExplicitEnumDefType              = SymbolId(426)
	TraitPropertyType                = SymbolId(427)
	TraitPropertiesType              = SymbolId(428)
	OptionalTraitPropertiesType      = SymbolId(429)
	TraitDefType                     = SymbolId(430)
	ReturnTypeType                   = SymbolId(431)
	ParameterDeclType                = SymbolId(432)
	ParameterDeclsType               = SymbolId(433)
	OptionalParameterDeclsType       = SymbolId(434)
	FuncTypeType                     = SymbolId(435)
	MethodSignatureType              = SymbolId(436)
	ParameterDefType                 = SymbolId(437)
	ParameterDefsType                = SymbolId(438)
	OptionalParameterDefsType        = SymbolId(439)
	NamedFuncDefType                 = SymbolId(440)
	AnonymousFuncExprType            = SymbolId(441)
	PackageDefType                   = SymbolId(442)
	PackageStatementBodyType         = SymbolId(443)
	PackageStatementType             = SymbolId(444)
	PackageStatementsType            = SymbolId(445)
	ImportStatementType              = SymbolId(446)
	ImportClauseType                 = SymbolId(447)
	ImportClauseTerminalType         = SymbolId(448)
	ImportClausesType                = SymbolId(449)
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
	_ReduceBlockBodyToDefinition                              = _ReduceType(14)
	_ReduceToVarDeclPattern                                   = _ReduceType(15)
	_ReduceVarToVarOrLet                                      = _ReduceType(16)
	_ReduceLetToVarOrLet                                      = _ReduceType(17)
	_ReduceIdentifierToVarPattern                             = _ReduceType(18)
	_ReduceTuplePatternToVarPattern                           = _ReduceType(19)
	_ReduceToTuplePattern                                     = _ReduceType(20)
	_ReduceFieldVarPatternToFieldVarPatterns                  = _ReduceType(21)
	_ReduceAddToFieldVarPatterns                              = _ReduceType(22)
	_ReducePositionalToFieldVarPattern                        = _ReduceType(23)
	_ReduceNamedToFieldVarPattern                             = _ReduceType(24)
	_ReduceDotDotDotToFieldVarPattern                         = _ReduceType(25)
	_ReduceValueTypeToOptionalValueType                       = _ReduceType(26)
	_ReduceNilToOptionalValueType                             = _ReduceType(27)
	_ReduceToAssignPattern                                    = _ReduceType(28)
	_ReduceAssignPatternToCasePattern                         = _ReduceType(29)
	_ReduceEnumMatchPatternToCasePattern                      = _ReduceType(30)
	_ReduceEnumVarDeclPatternToCasePattern                    = _ReduceType(31)
	_ReduceIfExprToExpression                                 = _ReduceType(32)
	_ReduceSwitchExprToExpression                             = _ReduceType(33)
	_ReduceLoopExprToExpression                               = _ReduceType(34)
	_ReduceSequenceExprToExpression                           = _ReduceType(35)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(36)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(37)
	_ReduceNoElseToIfExpr                                     = _ReduceType(38)
	_ReduceIfElseToIfExpr                                     = _ReduceType(39)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(40)
	_ReduceSequenceExprToCondition                            = _ReduceType(41)
	_ReduceCaseToCondition                                    = _ReduceType(42)
	_ReduceToSwitchExpr                                       = _ReduceType(43)
	_ReduceCaseBranchToCaseBranches                           = _ReduceType(44)
	_ReduceAddToCaseBranches                                  = _ReduceType(45)
	_ReduceToCaseBranch                                       = _ReduceType(46)
	_ReduceCasePatternToCasePatterns                          = _ReduceType(47)
	_ReduceMultipleToCasePatterns                             = _ReduceType(48)
	_ReduceDefaultBranchToOptionalDefaultBranch               = _ReduceType(49)
	_ReduceNilToOptionalDefaultBranch                         = _ReduceType(50)
	_ReduceToDefaultBranch                                    = _ReduceType(51)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(52)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(53)
	_ReduceWhileToLoopExpr                                    = _ReduceType(54)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(55)
	_ReduceForToLoopExpr                                      = _ReduceType(56)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(57)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(58)
	_ReduceSequenceExprToForAssignment                        = _ReduceType(59)
	_ReduceAssignToForAssignment                              = _ReduceType(60)
	_ReduceOrExprToSequenceExpr                               = _ReduceType(61)
	_ReduceVarDeclPatternToSequenceExpr                       = _ReduceType(62)
	_ReduceAssignVarPatternToSequenceExpr                     = _ReduceType(63)
	_ReduceToBlockExpr                                        = _ReduceType(64)
	_ReduceToBlockBody                                        = _ReduceType(65)
	_ReduceEmptyListToStatements                              = _ReduceType(66)
	_ReduceAddToStatements                                    = _ReduceType(67)
	_ReduceImplicitToStatement                                = _ReduceType(68)
	_ReduceExplicitToStatement                                = _ReduceType(69)
	_ReduceUnsafeStatementToStatementBody                     = _ReduceType(70)
	_ReduceExpressionOrImplicitStructToStatementBody          = _ReduceType(71)
	_ReduceAsyncToStatementBody                               = _ReduceType(72)
	_ReduceDeferToStatementBody                               = _ReduceType(73)
	_ReduceJumpStatementToStatementBody                       = _ReduceType(74)
	_ReduceAssignStatementToStatementBody                     = _ReduceType(75)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(76)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(77)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(78)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(79)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(80)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(81)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(82)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(83)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(84)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(85)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(86)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(87)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(88)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(89)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(90)
	_ReduceToUnsafeStatement                                  = _ReduceType(91)
	_ReduceToJumpStatement                                    = _ReduceType(92)
	_ReduceReturnToJumpType                                   = _ReduceType(93)
	_ReduceBreakToJumpType                                    = _ReduceType(94)
	_ReduceContinueToJumpType                                 = _ReduceType(95)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(96)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(97)
	_ReduceExpressionToExpressions                            = _ReduceType(98)
	_ReduceAddToExpressions                                   = _ReduceType(99)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(100)
	_ReduceNilToOptionalExpressions                           = _ReduceType(101)
	_ReduceToCallExpr                                         = _ReduceType(102)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(103)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(104)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(105)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(106)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(107)
	_ReduceAddToGenericArguments                              = _ReduceType(108)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(109)
	_ReduceNilToOptionalArguments                             = _ReduceType(110)
	_ReduceArgumentToArguments                                = _ReduceType(111)
	_ReduceAddToArguments                                     = _ReduceType(112)
	_ReducePositionalToArgument                               = _ReduceType(113)
	_ReduceNamedToArgument                                    = _ReduceType(114)
	_ReduceColonExpressionsToArgument                         = _ReduceType(115)
	_ReduceDotDotDotToArgument                                = _ReduceType(116)
	_ReducePairToColonExpressions                             = _ReduceType(117)
	_ReduceAddToColonExpressions                              = _ReduceType(118)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(119)
	_ReduceNilToOptionalExpression                            = _ReduceType(120)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(121)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(122)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(123)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(124)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(125)
	_ReduceImplicitStructExprToAtomExpr                       = _ReduceType(126)
	_ReduceParseErrorToAtomExpr                               = _ReduceType(127)
	_ReduceTrueToLiteral                                      = _ReduceType(128)
	_ReduceFalseToLiteral                                     = _ReduceType(129)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(130)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(131)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(132)
	_ReduceStringLiteralToLiteral                             = _ReduceType(133)
	_ReduceToImplicitStructExpr                               = _ReduceType(134)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(135)
	_ReduceAccessToAccessExpr                                 = _ReduceType(136)
	_ReduceCallExprToAccessExpr                               = _ReduceType(137)
	_ReduceIndexToAccessExpr                                  = _ReduceType(138)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(139)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(140)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(141)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(142)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(143)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(144)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(145)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(146)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(147)
	_ReduceMulToMulOp                                         = _ReduceType(148)
	_ReduceDivToMulOp                                         = _ReduceType(149)
	_ReduceModToMulOp                                         = _ReduceType(150)
	_ReduceBitAndToMulOp                                      = _ReduceType(151)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(152)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(153)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(154)
	_ReduceOpToMulExpr                                        = _ReduceType(155)
	_ReduceAddToAddOp                                         = _ReduceType(156)
	_ReduceSubToAddOp                                         = _ReduceType(157)
	_ReduceBitOrToAddOp                                       = _ReduceType(158)
	_ReduceBitXorToAddOp                                      = _ReduceType(159)
	_ReduceMulExprToAddExpr                                   = _ReduceType(160)
	_ReduceOpToAddExpr                                        = _ReduceType(161)
	_ReduceEqualToCmpOp                                       = _ReduceType(162)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(163)
	_ReduceLessToCmpOp                                        = _ReduceType(164)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(165)
	_ReduceGreaterToCmpOp                                     = _ReduceType(166)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(167)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(168)
	_ReduceOpToCmpExpr                                        = _ReduceType(169)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(170)
	_ReduceOpToAndExpr                                        = _ReduceType(171)
	_ReduceAndExprToOrExpr                                    = _ReduceType(172)
	_ReduceOpToOrExpr                                         = _ReduceType(173)
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(174)
	_ReduceSliceToInitializableType                           = _ReduceType(175)
	_ReduceArrayToInitializableType                           = _ReduceType(176)
	_ReduceMapToInitializableType                             = _ReduceType(177)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(178)
	_ReduceNamedToAtomType                                    = _ReduceType(179)
	_ReduceExternNamedToAtomType                              = _ReduceType(180)
	_ReduceInferredToAtomType                                 = _ReduceType(181)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(182)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(183)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(184)
	_ReduceTraitDefToAtomType                                 = _ReduceType(185)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(186)
	_ReduceParseErrorToAtomType                               = _ReduceType(187)
	_ReduceAtomTypeToReturnableType                           = _ReduceType(188)
	_ReduceOptionalToReturnableType                           = _ReduceType(189)
	_ReduceResultToReturnableType                             = _ReduceType(190)
	_ReduceReferenceToReturnableType                          = _ReduceType(191)
	_ReducePublicMethodsTraitToReturnableType                 = _ReduceType(192)
	_ReducePublicTraitToReturnableType                        = _ReduceType(193)
	_ReduceReturnableTypeToValueType                          = _ReduceType(194)
	_ReduceTraitIntersectToValueType                          = _ReduceType(195)
	_ReduceTraitUnionToValueType                              = _ReduceType(196)
	_ReduceTraitDifferenceToValueType                         = _ReduceType(197)
	_ReduceDefinitionToTypeDef                                = _ReduceType(198)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(199)
	_ReduceAliasToTypeDef                                     = _ReduceType(200)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(201)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(202)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(203)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(204)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(205)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(206)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(207)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(208)
	_ReduceExplicitToFieldDef                                 = _ReduceType(209)
	_ReduceImplicitToFieldDef                                 = _ReduceType(210)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(211)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(212)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(213)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(214)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(215)
	_ReduceToImplicitStructDef                                = _ReduceType(216)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(217)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(218)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(219)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(220)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(221)
	_ReduceToExplicitStructDef                                = _ReduceType(222)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(223)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(224)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(225)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(226)
	_ReduceToImplicitEnumDef                                  = _ReduceType(227)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(228)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(229)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(230)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(231)
	_ReduceToExplicitEnumDef                                  = _ReduceType(232)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(233)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(234)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(235)
	_ReduceImplicitToTraitProperties                          = _ReduceType(236)
	_ReduceExplicitToTraitProperties                          = _ReduceType(237)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(238)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(239)
	_ReduceToTraitDef                                         = _ReduceType(240)
	_ReduceReturnableTypeToReturnType                         = _ReduceType(241)
	_ReduceNilToReturnType                                    = _ReduceType(242)
	_ReduceArgToParameterDecl                                 = _ReduceType(243)
	_ReduceVarargToParameterDecl                              = _ReduceType(244)
	_ReduceUnamedToParameterDecl                              = _ReduceType(245)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(246)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(247)
	_ReduceAddToParameterDecls                                = _ReduceType(248)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(249)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(250)
	_ReduceToFuncType                                         = _ReduceType(251)
	_ReduceToMethodSignature                                  = _ReduceType(252)
	_ReduceInferredRefArgToParameterDef                       = _ReduceType(253)
	_ReduceInferredRefVarargToParameterDef                    = _ReduceType(254)
	_ReduceArgToParameterDef                                  = _ReduceType(255)
	_ReduceVarargToParameterDef                               = _ReduceType(256)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(257)
	_ReduceAddToParameterDefs                                 = _ReduceType(258)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(259)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(260)
	_ReduceFuncDefToNamedFuncDef                              = _ReduceType(261)
	_ReduceMethodDefToNamedFuncDef                            = _ReduceType(262)
	_ReduceAliasToNamedFuncDef                                = _ReduceType(263)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(264)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(265)
	_ReduceWithSpecToPackageDef                               = _ReduceType(266)
	_ReduceUnsafeStatementToPackageStatementBody              = _ReduceType(267)
	_ReduceImportStatementToPackageStatementBody              = _ReduceType(268)
	_ReduceImplicitToPackageStatement                         = _ReduceType(269)
	_ReduceExplicitToPackageStatement                         = _ReduceType(270)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(271)
	_ReduceAddToPackageStatements                             = _ReduceType(272)
	_ReduceSingleToImportStatement                            = _ReduceType(273)
	_ReduceMultipleToImportStatement                          = _ReduceType(274)
	_ReduceStringLiteralToImportClause                        = _ReduceType(275)
	_ReduceAliasToImportClause                                = _ReduceType(276)
	_ReduceImplicitToImportClauseTerminal                     = _ReduceType(277)
	_ReduceExplicitToImportClauseTerminal                     = _ReduceType(278)
	_ReduceFirstToImportClauses                               = _ReduceType(279)
	_ReduceAddToImportClauses                                 = _ReduceType(280)
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
	case _ReduceBlockBodyToDefinition:
		return "BlockBodyToDefinition"
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
	case _ReduceCaseBranchToCaseBranches:
		return "CaseBranchToCaseBranches"
	case _ReduceAddToCaseBranches:
		return "AddToCaseBranches"
	case _ReduceToCaseBranch:
		return "ToCaseBranch"
	case _ReduceCasePatternToCasePatterns:
		return "CasePatternToCasePatterns"
	case _ReduceMultipleToCasePatterns:
		return "MultipleToCasePatterns"
	case _ReduceDefaultBranchToOptionalDefaultBranch:
		return "DefaultBranchToOptionalDefaultBranch"
	case _ReduceNilToOptionalDefaultBranch:
		return "NilToOptionalDefaultBranch"
	case _ReduceToDefaultBranch:
		return "ToDefaultBranch"
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
	case _ReduceOrExprToSequenceExpr:
		return "OrExprToSequenceExpr"
	case _ReduceVarDeclPatternToSequenceExpr:
		return "VarDeclPatternToSequenceExpr"
	case _ReduceAssignVarPatternToSequenceExpr:
		return "AssignVarPatternToSequenceExpr"
	case _ReduceToBlockExpr:
		return "ToBlockExpr"
	case _ReduceToBlockBody:
		return "ToBlockBody"
	case _ReduceEmptyListToStatements:
		return "EmptyListToStatements"
	case _ReduceAddToStatements:
		return "AddToStatements"
	case _ReduceImplicitToStatement:
		return "ImplicitToStatement"
	case _ReduceExplicitToStatement:
		return "ExplicitToStatement"
	case _ReduceUnsafeStatementToStatementBody:
		return "UnsafeStatementToStatementBody"
	case _ReduceExpressionOrImplicitStructToStatementBody:
		return "ExpressionOrImplicitStructToStatementBody"
	case _ReduceAsyncToStatementBody:
		return "AsyncToStatementBody"
	case _ReduceDeferToStatementBody:
		return "DeferToStatementBody"
	case _ReduceJumpStatementToStatementBody:
		return "JumpStatementToStatementBody"
	case _ReduceAssignStatementToStatementBody:
		return "AssignStatementToStatementBody"
	case _ReduceUnaryOpAssignStatementToStatementBody:
		return "UnaryOpAssignStatementToStatementBody"
	case _ReduceBinaryOpAssignStatementToStatementBody:
		return "BinaryOpAssignStatementToStatementBody"
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
	case _ReduceToJumpStatement:
		return "ToJumpStatement"
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
	case _ReduceUnsafeStatementToPackageStatementBody:
		return "UnsafeStatementToPackageStatementBody"
	case _ReduceImportStatementToPackageStatementBody:
		return "ImportStatementToPackageStatementBody"
	case _ReduceImplicitToPackageStatement:
		return "ImplicitToPackageStatement"
	case _ReduceExplicitToPackageStatement:
		return "ExplicitToPackageStatement"
	case _ReduceEmptyListToPackageStatements:
		return "EmptyListToPackageStatements"
	case _ReduceAddToPackageStatements:
		return "AddToPackageStatements"
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
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ *GenericSymbol
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _EndMarker, NewlinesToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, ParseErrorToken:
		val, ok := token.(*GenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *GenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
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
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return Location{}
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
			token = &GenericSymbol{_EndMarker, stack.lexer.CurrentLocation()}
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
		symbol.Generic_, err = reducer.NewlinesToOptionalDefinitions(args[0].Generic_)
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
		symbol.Generic_, err = reducer.NewlinesToOptionalNewlines(args[0].Generic_)
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
		symbol.Generic_, err = reducer.AddToDefinitions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceBlockBodyToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.BlockBodyToDefinition(args[0].Generic_)
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
		symbol.Generic_, err = reducer.IdentifierToVarPattern(args[0].Generic_)
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
		symbol.Generic_, err = reducer.NamedToFieldVarPattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.EnumMatchPatternToCasePattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceEnumVarDeclPatternToCasePattern:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumVarDeclPatternToCasePattern(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
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
		symbol.Generic_, err = reducer.LabelDeclToOptionalLabelDecl(args[0].Generic_)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.UnlabelledToOptionalLabelDecl()
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
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = SwitchExprType
		symbol.Generic_, err = reducer.ToSwitchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceCaseBranchToCaseBranches:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CaseBranchesType
		symbol.Generic_, err = reducer.CaseBranchToCaseBranches(args[0].Generic_)
	case _ReduceAddToCaseBranches:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CaseBranchesType
		symbol.Generic_, err = reducer.AddToCaseBranches(args[0].Generic_, args[1].Generic_)
	case _ReduceToCaseBranch:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CaseBranchType
		symbol.Generic_, err = reducer.ToCaseBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
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
	case _ReduceDefaultBranchToOptionalDefaultBranch:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalDefaultBranchType
		symbol.Generic_, err = reducer.DefaultBranchToOptionalDefaultBranch(args[0].Generic_)
	case _ReduceNilToOptionalDefaultBranch:
		symbol.SymbolId_ = OptionalDefaultBranchType
		symbol.Generic_, err = reducer.NilToOptionalDefaultBranch()
	case _ReduceToDefaultBranch:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefaultBranchType
		symbol.Generic_, err = reducer.ToDefaultBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceToBlockExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BlockExprType
		symbol.Generic_, err = reducer.ToBlockExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceToBlockBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BlockBodyType
		symbol.Generic_, err = reducer.ToBlockBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ImplicitToStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Generic_, err = reducer.ExplicitToStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceUnsafeStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.UnsafeStatementToStatementBody(args[0].Generic_)
	case _ReduceExpressionOrImplicitStructToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.ExpressionOrImplicitStructToStatementBody(args[0].Generic_)
	case _ReduceAsyncToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.AsyncToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceDeferToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.DeferToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceJumpStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.JumpStatementToStatementBody(args[0].Generic_)
	case _ReduceAssignStatementToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.AssignStatementToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceUnaryOpAssignStatementToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.UnaryOpAssignStatementToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceBinaryOpAssignStatementToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.BinaryOpAssignStatementToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ToUnsafeStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = JumpStatementType
		symbol.Generic_, err = reducer.ToJumpStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.JumpLabelToOptionalJumpLabel(args[0].Generic_)
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
		symbol.Generic_, err = reducer.NamedToArgument(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.IdentifierToAtomExpr(args[0].Generic_)
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.BlockExprToAtomExpr(args[0].Generic_)
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
		symbol.Generic_, err = reducer.ParseErrorToAtomExpr(args[0].Generic_)
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
		symbol.Generic_, err = reducer.AccessToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.NamedToAtomType(args[0].Generic_, args[1].Generic_)
	case _ReduceExternNamedToAtomType:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExternNamedToAtomType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
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
		symbol.Generic_, err = reducer.ParseErrorToAtomType(args[0].Generic_)
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
		symbol.Generic_, err = reducer.DefinitionToTypeDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceConstrainedDefToTypeDef:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.ConstrainedDefToTypeDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceAliasToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.AliasToTypeDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceUnconstrainedToGenericParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.UnconstrainedToGenericParameterDef(args[0].Generic_)
	case _ReduceConstrainedToGenericParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.ConstrainedToGenericParameterDef(args[0].Generic_, args[1].Generic_)
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
		symbol.Generic_, err = reducer.ExplicitToFieldDef(args[0].Generic_, args[1].Generic_)
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
		symbol.Generic_, err = reducer.ImplicitToExplicitFieldDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ImplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ImplicitToTraitProperties(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ArgToParameterDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceVarargToParameterDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.VarargToParameterDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.ToMethodSignature(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceInferredRefArgToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.InferredRefArgToParameterDef(args[0].Generic_)
	case _ReduceInferredRefVarargToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.InferredRefVarargToParameterDef(args[0].Generic_, args[1].Generic_)
	case _ReduceArgToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.ArgToParameterDef(args[0].Generic_, args[1].Generic_)
	case _ReduceVarargToParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.VarargToParameterDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		symbol.Generic_, err = reducer.FuncDefToNamedFuncDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_, args[7].Generic_)
	case _ReduceMethodDefToNamedFuncDef:
		args := stack[len(stack)-11:]
		stack = stack[:len(stack)-11]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.Generic_, err = reducer.MethodDefToNamedFuncDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_, args[7].Generic_, args[8].Generic_, args[9].Generic_, args[10].Generic_)
	case _ReduceAliasToNamedFuncDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.Generic_, err = reducer.AliasToNamedFuncDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
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
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = PackageDefType
		symbol.Generic_, err = reducer.WithSpecToPackageDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceUnsafeStatementToPackageStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PackageStatementBodyType
		symbol.Generic_, err = reducer.UnsafeStatementToPackageStatementBody(args[0].Generic_)
	case _ReduceImportStatementToPackageStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PackageStatementBodyType
		symbol.Generic_, err = reducer.ImportStatementToPackageStatementBody(args[0].Generic_)
	case _ReduceImplicitToPackageStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageStatementType
		symbol.Generic_, err = reducer.ImplicitToPackageStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitToPackageStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageStatementType
		symbol.Generic_, err = reducer.ExplicitToPackageStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceEmptyListToPackageStatements:
		symbol.SymbolId_ = PackageStatementsType
		symbol.Generic_, err = reducer.EmptyListToPackageStatements()
	case _ReduceAddToPackageStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageStatementsType
		symbol.Generic_, err = reducer.AddToPackageStatements(args[0].Generic_, args[1].Generic_)
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
		symbol.Generic_, err = reducer.AliasToImportClause(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitToImportClauseTerminal:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseTerminalType
		symbol.Generic_, err = reducer.ImplicitToImportClauseTerminal(args[0].Generic_, args[1].Generic_)
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
	_ReduceBlockBodyToDefinitionAction                              = &_Action{_ReduceAction, 0, _ReduceBlockBodyToDefinition}
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
	_ReduceNoElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceSequenceExprToConditionAction                            = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                    = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceCaseBranchToCaseBranchesAction                           = &_Action{_ReduceAction, 0, _ReduceCaseBranchToCaseBranches}
	_ReduceAddToCaseBranchesAction                                  = &_Action{_ReduceAction, 0, _ReduceAddToCaseBranches}
	_ReduceToCaseBranchAction                                       = &_Action{_ReduceAction, 0, _ReduceToCaseBranch}
	_ReduceCasePatternToCasePatternsAction                          = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                             = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceDefaultBranchToOptionalDefaultBranchAction               = &_Action{_ReduceAction, 0, _ReduceDefaultBranchToOptionalDefaultBranch}
	_ReduceNilToOptionalDefaultBranchAction                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefaultBranch}
	_ReduceToDefaultBranchAction                                    = &_Action{_ReduceAction, 0, _ReduceToDefaultBranch}
	_ReduceInfiniteToLoopExprAction                                 = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                  = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                    = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                                 = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                      = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction                 = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                          = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceSequenceExprToForAssignmentAction                        = &_Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}
	_ReduceAssignToForAssignmentAction                              = &_Action{_ReduceAction, 0, _ReduceAssignToForAssignment}
	_ReduceOrExprToSequenceExprAction                               = &_Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}
	_ReduceVarDeclPatternToSequenceExprAction                       = &_Action{_ReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}
	_ReduceAssignVarPatternToSequenceExprAction                     = &_Action{_ReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}
	_ReduceToBlockExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToBlockBodyAction                                        = &_Action{_ReduceAction, 0, _ReduceToBlockBody}
	_ReduceEmptyListToStatementsAction                              = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                                    = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceImplicitToStatementAction                                = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction                                = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceUnsafeStatementToStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToStatementBody}
	_ReduceExpressionOrImplicitStructToStatementBodyAction          = &_Action{_ReduceAction, 0, _ReduceExpressionOrImplicitStructToStatementBody}
	_ReduceAsyncToStatementBodyAction                               = &_Action{_ReduceAction, 0, _ReduceAsyncToStatementBody}
	_ReduceDeferToStatementBodyAction                               = &_Action{_ReduceAction, 0, _ReduceDeferToStatementBody}
	_ReduceJumpStatementToStatementBodyAction                       = &_Action{_ReduceAction, 0, _ReduceJumpStatementToStatementBody}
	_ReduceAssignStatementToStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceAssignStatementToStatementBody}
	_ReduceUnaryOpAssignStatementToStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToStatementBody}
	_ReduceBinaryOpAssignStatementToStatementBodyAction             = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToStatementBody}
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
	_ReduceToJumpStatementAction                                    = &_Action{_ReduceAction, 0, _ReduceToJumpStatement}
	_ReduceReturnToJumpTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceJumpLabelToOptionalJumpLabelAction                       = &_Action{_ReduceAction, 0, _ReduceJumpLabelToOptionalJumpLabel}
	_ReduceUnlabelledToOptionalJumpLabelAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalJumpLabel}
	_ReduceExpressionToExpressionsAction                            = &_Action{_ReduceAction, 0, _ReduceExpressionToExpressions}
	_ReduceAddToExpressionsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToExpressions}
	_ReduceExpressionsToOptionalExpressionsAction                   = &_Action{_ReduceAction, 0, _ReduceExpressionsToOptionalExpressions}
	_ReduceNilToOptionalExpressionsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpressions}
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
	_ReduceUnsafeStatementToPackageStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToPackageStatementBody}
	_ReduceImportStatementToPackageStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceImportStatementToPackageStatementBody}
	_ReduceImplicitToPackageStatementAction                         = &_Action{_ReduceAction, 0, _ReduceImplicitToPackageStatement}
	_ReduceExplicitToPackageStatementAction                         = &_Action{_ReduceAction, 0, _ReduceExplicitToPackageStatement}
	_ReduceEmptyListToPackageStatementsAction                       = &_Action{_ReduceAction, 0, _ReduceEmptyListToPackageStatements}
	_ReduceAddToPackageStatementsAction                             = &_Action{_ReduceAction, 0, _ReduceAddToPackageStatements}
	_ReduceSingleToImportStatementAction                            = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                          = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                        = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                                = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceImplicitToImportClauseTerminalAction                     = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauseTerminal}
	_ReduceExplicitToImportClauseTerminalAction                     = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauseTerminal}
	_ReduceFirstToImportClausesAction                               = &_Action{_ReduceAction, 0, _ReduceFirstToImportClauses}
	_ReduceAddToImportClausesAction                                 = &_Action{_ReduceAction, 0, _ReduceAddToImportClauses}
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
	{_State5, LabelDeclToken}:                     _GotoState25Action,
	{_State5, LparenToken}:                        _GotoState28Action,
	{_State5, LbracketToken}:                      _GotoState26Action,
	{_State5, NotToken}:                           _GotoState30Action,
	{_State5, SubToken}:                           _GotoState35Action,
	{_State5, MulToken}:                           _GotoState29Action,
	{_State5, BitNegToken}:                        _GotoState18Action,
	{_State5, BitAndToken}:                        _GotoState17Action,
	{_State5, GreaterToken}:                       _GotoState22Action,
	{_State5, ParseErrorToken}:                    _GotoState31Action,
	{_State5, VarDeclPatternType}:                 _GotoState57Action,
	{_State5, VarOrLetType}:                       _GotoState58Action,
	{_State5, ExpressionType}:                     _GotoState10Action,
	{_State5, OptionalLabelDeclType}:              _GotoState51Action,
	{_State5, SequenceExprType}:                   _GotoState56Action,
	{_State5, BlockExprType}:                      _GotoState43Action,
	{_State5, CallExprType}:                       _GotoState44Action,
	{_State5, AtomExprType}:                       _GotoState42Action,
	{_State5, LiteralType}:                        _GotoState49Action,
	{_State5, ImplicitStructExprType}:             _GotoState47Action,
	{_State5, AccessExprType}:                     _GotoState38Action,
	{_State5, PostfixUnaryExprType}:               _GotoState53Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState55Action,
	{_State5, PrefixUnaryExprType}:                _GotoState54Action,
	{_State5, MulExprType}:                        _GotoState50Action,
	{_State5, AddExprType}:                        _GotoState39Action,
	{_State5, CmpExprType}:                        _GotoState45Action,
	{_State5, AndExprType}:                        _GotoState40Action,
	{_State5, OrExprType}:                         _GotoState52Action,
	{_State5, InitializableTypeType}:              _GotoState48Action,
	{_State5, ExplicitStructDefType}:              _GotoState46Action,
	{_State5, AnonymousFuncExprType}:              _GotoState41Action,
	{_State13, PackageToken}:                      _GotoState14Action,
	{_State13, TypeToken}:                         _GotoState15Action,
	{_State13, FuncToken}:                         _GotoState16Action,
	{_State13, VarToken}:                          _GotoState37Action,
	{_State13, LetToken}:                          _GotoState27Action,
	{_State13, LbraceToken}:                       _GotoState59Action,
	{_State13, DefinitionsType}:                   _GotoState62Action,
	{_State13, DefinitionType}:                    _GotoState61Action,
	{_State13, VarDeclPatternType}:                _GotoState66Action,
	{_State13, VarOrLetType}:                      _GotoState58Action,
	{_State13, BlockBodyType}:                     _GotoState60Action,
	{_State13, TypeDefType}:                       _GotoState65Action,
	{_State13, NamedFuncDefType}:                  _GotoState63Action,
	{_State13, PackageDefType}:                    _GotoState64Action,
	{_State14, LparenToken}:                       _GotoState67Action,
	{_State15, IdentifierToken}:                   _GotoState68Action,
	{_State16, IdentifierToken}:                   _GotoState69Action,
	{_State16, LparenToken}:                       _GotoState70Action,
	{_State21, LparenToken}:                       _GotoState71Action,
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
	{_State22, LabelDeclToken}:                    _GotoState25Action,
	{_State22, LparenToken}:                       _GotoState28Action,
	{_State22, LbracketToken}:                     _GotoState26Action,
	{_State22, NotToken}:                          _GotoState30Action,
	{_State22, SubToken}:                          _GotoState35Action,
	{_State22, MulToken}:                          _GotoState29Action,
	{_State22, BitNegToken}:                       _GotoState18Action,
	{_State22, BitAndToken}:                       _GotoState17Action,
	{_State22, GreaterToken}:                      _GotoState22Action,
	{_State22, ParseErrorToken}:                   _GotoState31Action,
	{_State22, VarDeclPatternType}:                _GotoState57Action,
	{_State22, VarOrLetType}:                      _GotoState58Action,
	{_State22, OptionalLabelDeclType}:             _GotoState72Action,
	{_State22, SequenceExprType}:                  _GotoState73Action,
	{_State22, BlockExprType}:                     _GotoState43Action,
	{_State22, CallExprType}:                      _GotoState44Action,
	{_State22, AtomExprType}:                      _GotoState42Action,
	{_State22, LiteralType}:                       _GotoState49Action,
	{_State22, ImplicitStructExprType}:            _GotoState47Action,
	{_State22, AccessExprType}:                    _GotoState38Action,
	{_State22, PostfixUnaryExprType}:              _GotoState53Action,
	{_State22, PrefixUnaryOpType}:                 _GotoState55Action,
	{_State22, PrefixUnaryExprType}:               _GotoState54Action,
	{_State22, MulExprType}:                       _GotoState50Action,
	{_State22, AddExprType}:                       _GotoState39Action,
	{_State22, CmpExprType}:                       _GotoState45Action,
	{_State22, AndExprType}:                       _GotoState40Action,
	{_State22, OrExprType}:                        _GotoState52Action,
	{_State22, InitializableTypeType}:             _GotoState48Action,
	{_State22, ExplicitStructDefType}:             _GotoState46Action,
	{_State22, AnonymousFuncExprType}:             _GotoState41Action,
	{_State26, IdentifierToken}:                   _GotoState80Action,
	{_State26, StructToken}:                       _GotoState34Action,
	{_State26, EnumToken}:                         _GotoState77Action,
	{_State26, TraitToken}:                        _GotoState85Action,
	{_State26, FuncToken}:                         _GotoState79Action,
	{_State26, LparenToken}:                       _GotoState81Action,
	{_State26, LbracketToken}:                     _GotoState26Action,
	{_State26, DotToken}:                          _GotoState76Action,
	{_State26, QuestionToken}:                     _GotoState83Action,
	{_State26, ExclaimToken}:                      _GotoState78Action,
	{_State26, TildeTildeToken}:                   _GotoState84Action,
	{_State26, BitNegToken}:                       _GotoState75Action,
	{_State26, BitAndToken}:                       _GotoState74Action,
	{_State26, ParseErrorToken}:                   _GotoState82Action,
	{_State26, InitializableTypeType}:             _GotoState91Action,
	{_State26, AtomTypeType}:                      _GotoState86Action,
	{_State26, ReturnableTypeType}:                _GotoState92Action,
	{_State26, ValueTypeType}:                     _GotoState94Action,
	{_State26, ImplicitStructDefType}:             _GotoState90Action,
	{_State26, ExplicitStructDefType}:             _GotoState46Action,
	{_State26, ImplicitEnumDefType}:               _GotoState89Action,
	{_State26, ExplicitEnumDefType}:               _GotoState87Action,
	{_State26, TraitDefType}:                      _GotoState93Action,
	{_State26, FuncTypeType}:                      _GotoState88Action,
	{_State28, IntegerLiteralToken}:               _GotoState24Action,
	{_State28, FloatLiteralToken}:                 _GotoState20Action,
	{_State28, RuneLiteralToken}:                  _GotoState32Action,
	{_State28, StringLiteralToken}:                _GotoState33Action,
	{_State28, IdentifierToken}:                   _GotoState96Action,
	{_State28, TrueToken}:                         _GotoState36Action,
	{_State28, FalseToken}:                        _GotoState19Action,
	{_State28, StructToken}:                       _GotoState34Action,
	{_State28, FuncToken}:                         _GotoState21Action,
	{_State28, VarToken}:                          _GotoState37Action,
	{_State28, LetToken}:                          _GotoState27Action,
	{_State28, LabelDeclToken}:                    _GotoState25Action,
	{_State28, LparenToken}:                       _GotoState28Action,
	{_State28, LbracketToken}:                     _GotoState26Action,
	{_State28, DotDotDotToken}:                    _GotoState95Action,
	{_State28, NotToken}:                          _GotoState30Action,
	{_State28, SubToken}:                          _GotoState35Action,
	{_State28, MulToken}:                          _GotoState29Action,
	{_State28, BitNegToken}:                       _GotoState18Action,
	{_State28, BitAndToken}:                       _GotoState17Action,
	{_State28, GreaterToken}:                      _GotoState22Action,
	{_State28, ParseErrorToken}:                   _GotoState31Action,
	{_State28, VarDeclPatternType}:                _GotoState57Action,
	{_State28, VarOrLetType}:                      _GotoState58Action,
	{_State28, ExpressionType}:                    _GotoState100Action,
	{_State28, OptionalLabelDeclType}:             _GotoState51Action,
	{_State28, SequenceExprType}:                  _GotoState56Action,
	{_State28, BlockExprType}:                     _GotoState43Action,
	{_State28, CallExprType}:                      _GotoState44Action,
	{_State28, ArgumentsType}:                     _GotoState98Action,
	{_State28, ArgumentType}:                      _GotoState97Action,
	{_State28, ColonExpressionsType}:              _GotoState99Action,
	{_State28, OptionalExpressionType}:            _GotoState101Action,
	{_State28, AtomExprType}:                      _GotoState42Action,
	{_State28, LiteralType}:                       _GotoState49Action,
	{_State28, ImplicitStructExprType}:            _GotoState47Action,
	{_State28, AccessExprType}:                    _GotoState38Action,
	{_State28, PostfixUnaryExprType}:              _GotoState53Action,
	{_State28, PrefixUnaryOpType}:                 _GotoState55Action,
	{_State28, PrefixUnaryExprType}:               _GotoState54Action,
	{_State28, MulExprType}:                       _GotoState50Action,
	{_State28, AddExprType}:                       _GotoState39Action,
	{_State28, CmpExprType}:                       _GotoState45Action,
	{_State28, AndExprType}:                       _GotoState40Action,
	{_State28, OrExprType}:                        _GotoState52Action,
	{_State28, InitializableTypeType}:             _GotoState48Action,
	{_State28, ExplicitStructDefType}:             _GotoState46Action,
	{_State28, AnonymousFuncExprType}:             _GotoState41Action,
	{_State34, LparenToken}:                       _GotoState102Action,
	{_State38, LbracketToken}:                     _GotoState105Action,
	{_State38, DotToken}:                          _GotoState104Action,
	{_State38, QuestionToken}:                     _GotoState106Action,
	{_State38, DollarLbracketToken}:               _GotoState103Action,
	{_State38, OptionalGenericBindingType}:        _GotoState107Action,
	{_State39, AddToken}:                          _GotoState108Action,
	{_State39, SubToken}:                          _GotoState111Action,
	{_State39, BitXorToken}:                       _GotoState110Action,
	{_State39, BitOrToken}:                        _GotoState109Action,
	{_State39, AddOpType}:                         _GotoState112Action,
	{_State40, AndToken}:                          _GotoState113Action,
	{_State45, EqualToken}:                        _GotoState114Action,
	{_State45, NotEqualToken}:                     _GotoState119Action,
	{_State45, LessToken}:                         _GotoState117Action,
	{_State45, LessOrEqualToken}:                  _GotoState118Action,
	{_State45, GreaterToken}:                      _GotoState115Action,
	{_State45, GreaterOrEqualToken}:               _GotoState116Action,
	{_State45, CmpOpType}:                         _GotoState120Action,
	{_State48, LparenToken}:                       _GotoState121Action,
	{_State50, MulToken}:                          _GotoState127Action,
	{_State50, DivToken}:                          _GotoState125Action,
	{_State50, ModToken}:                          _GotoState126Action,
	{_State50, BitAndToken}:                       _GotoState122Action,
	{_State50, BitLshiftToken}:                    _GotoState123Action,
	{_State50, BitRshiftToken}:                    _GotoState124Action,
	{_State50, MulOpType}:                         _GotoState128Action,
	{_State51, IfToken}:                           _GotoState131Action,
	{_State51, SwitchToken}:                       _GotoState132Action,
	{_State51, ForToken}:                          _GotoState130Action,
	{_State51, DoToken}:                           _GotoState129Action,
	{_State51, LbraceToken}:                       _GotoState59Action,
	{_State51, IfExprType}:                        _GotoState134Action,
	{_State51, SwitchExprType}:                    _GotoState136Action,
	{_State51, LoopExprType}:                      _GotoState135Action,
	{_State51, BlockBodyType}:                     _GotoState133Action,
	{_State52, OrToken}:                           _GotoState137Action,
	{_State55, IntegerLiteralToken}:               _GotoState24Action,
	{_State55, FloatLiteralToken}:                 _GotoState20Action,
	{_State55, RuneLiteralToken}:                  _GotoState32Action,
	{_State55, StringLiteralToken}:                _GotoState33Action,
	{_State55, IdentifierToken}:                   _GotoState23Action,
	{_State55, TrueToken}:                         _GotoState36Action,
	{_State55, FalseToken}:                        _GotoState19Action,
	{_State55, StructToken}:                       _GotoState34Action,
	{_State55, FuncToken}:                         _GotoState21Action,
	{_State55, LabelDeclToken}:                    _GotoState25Action,
	{_State55, LparenToken}:                       _GotoState28Action,
	{_State55, LbracketToken}:                     _GotoState26Action,
	{_State55, NotToken}:                          _GotoState30Action,
	{_State55, SubToken}:                          _GotoState35Action,
	{_State55, MulToken}:                          _GotoState29Action,
	{_State55, BitNegToken}:                       _GotoState18Action,
	{_State55, BitAndToken}:                       _GotoState17Action,
	{_State55, ParseErrorToken}:                   _GotoState31Action,
	{_State55, OptionalLabelDeclType}:             _GotoState72Action,
	{_State55, BlockExprType}:                     _GotoState43Action,
	{_State55, CallExprType}:                      _GotoState44Action,
	{_State55, AtomExprType}:                      _GotoState42Action,
	{_State55, LiteralType}:                       _GotoState49Action,
	{_State55, ImplicitStructExprType}:            _GotoState47Action,
	{_State55, AccessExprType}:                    _GotoState38Action,
	{_State55, PostfixUnaryExprType}:              _GotoState53Action,
	{_State55, PrefixUnaryOpType}:                 _GotoState55Action,
	{_State55, PrefixUnaryExprType}:               _GotoState138Action,
	{_State55, InitializableTypeType}:             _GotoState48Action,
	{_State55, ExplicitStructDefType}:             _GotoState46Action,
	{_State55, AnonymousFuncExprType}:             _GotoState41Action,
	{_State58, IdentifierToken}:                   _GotoState139Action,
	{_State58, LparenToken}:                       _GotoState140Action,
	{_State58, VarPatternType}:                    _GotoState142Action,
	{_State58, TuplePatternType}:                  _GotoState141Action,
	{_State59, StatementsType}:                    _GotoState143Action,
	{_State62, NewlinesToken}:                     _GotoState144Action,
	{_State62, OptionalNewlinesType}:              _GotoState145Action,
	{_State66, AssignToken}:                       _GotoState146Action,
	{_State67, PackageStatementsType}:             _GotoState147Action,
	{_State68, DollarLbracketToken}:               _GotoState149Action,
	{_State68, AssignToken}:                       _GotoState148Action,
	{_State68, OptionalGenericParametersType}:     _GotoState150Action,
	{_State69, DollarLbracketToken}:               _GotoState149Action,
	{_State69, AssignToken}:                       _GotoState151Action,
	{_State69, OptionalGenericParametersType}:     _GotoState152Action,
	{_State70, IdentifierToken}:                   _GotoState153Action,
	{_State70, ParameterDefType}:                  _GotoState154Action,
	{_State71, IdentifierToken}:                   _GotoState153Action,
	{_State71, ParameterDefType}:                  _GotoState156Action,
	{_State71, ParameterDefsType}:                 _GotoState157Action,
	{_State71, OptionalParameterDefsType}:         _GotoState155Action,
	{_State72, LbraceToken}:                       _GotoState59Action,
	{_State72, BlockBodyType}:                     _GotoState133Action,
	{_State74, IdentifierToken}:                   _GotoState80Action,
	{_State74, StructToken}:                       _GotoState34Action,
	{_State74, EnumToken}:                         _GotoState77Action,
	{_State74, TraitToken}:                        _GotoState85Action,
	{_State74, FuncToken}:                         _GotoState79Action,
	{_State74, LparenToken}:                       _GotoState81Action,
	{_State74, LbracketToken}:                     _GotoState26Action,
	{_State74, DotToken}:                          _GotoState76Action,
	{_State74, QuestionToken}:                     _GotoState83Action,
	{_State74, ExclaimToken}:                      _GotoState78Action,
	{_State74, TildeTildeToken}:                   _GotoState84Action,
	{_State74, BitNegToken}:                       _GotoState75Action,
	{_State74, BitAndToken}:                       _GotoState74Action,
	{_State74, ParseErrorToken}:                   _GotoState82Action,
	{_State74, InitializableTypeType}:             _GotoState91Action,
	{_State74, AtomTypeType}:                      _GotoState86Action,
	{_State74, ReturnableTypeType}:                _GotoState158Action,
	{_State74, ImplicitStructDefType}:             _GotoState90Action,
	{_State74, ExplicitStructDefType}:             _GotoState46Action,
	{_State74, ImplicitEnumDefType}:               _GotoState89Action,
	{_State74, ExplicitEnumDefType}:               _GotoState87Action,
	{_State74, TraitDefType}:                      _GotoState93Action,
	{_State74, FuncTypeType}:                      _GotoState88Action,
	{_State75, IdentifierToken}:                   _GotoState80Action,
	{_State75, StructToken}:                       _GotoState34Action,
	{_State75, EnumToken}:                         _GotoState77Action,
	{_State75, TraitToken}:                        _GotoState85Action,
	{_State75, FuncToken}:                         _GotoState79Action,
	{_State75, LparenToken}:                       _GotoState81Action,
	{_State75, LbracketToken}:                     _GotoState26Action,
	{_State75, DotToken}:                          _GotoState76Action,
	{_State75, QuestionToken}:                     _GotoState83Action,
	{_State75, ExclaimToken}:                      _GotoState78Action,
	{_State75, TildeTildeToken}:                   _GotoState84Action,
	{_State75, BitNegToken}:                       _GotoState75Action,
	{_State75, BitAndToken}:                       _GotoState74Action,
	{_State75, ParseErrorToken}:                   _GotoState82Action,
	{_State75, InitializableTypeType}:             _GotoState91Action,
	{_State75, AtomTypeType}:                      _GotoState86Action,
	{_State75, ReturnableTypeType}:                _GotoState159Action,
	{_State75, ImplicitStructDefType}:             _GotoState90Action,
	{_State75, ExplicitStructDefType}:             _GotoState46Action,
	{_State75, ImplicitEnumDefType}:               _GotoState89Action,
	{_State75, ExplicitEnumDefType}:               _GotoState87Action,
	{_State75, TraitDefType}:                      _GotoState93Action,
	{_State75, FuncTypeType}:                      _GotoState88Action,
	{_State76, DollarLbracketToken}:               _GotoState103Action,
	{_State76, OptionalGenericBindingType}:        _GotoState160Action,
	{_State77, LparenToken}:                       _GotoState161Action,
	{_State78, IdentifierToken}:                   _GotoState80Action,
	{_State78, StructToken}:                       _GotoState34Action,
	{_State78, EnumToken}:                         _GotoState77Action,
	{_State78, TraitToken}:                        _GotoState85Action,
	{_State78, FuncToken}:                         _GotoState79Action,
	{_State78, LparenToken}:                       _GotoState81Action,
	{_State78, LbracketToken}:                     _GotoState26Action,
	{_State78, DotToken}:                          _GotoState76Action,
	{_State78, QuestionToken}:                     _GotoState83Action,
	{_State78, ExclaimToken}:                      _GotoState78Action,
	{_State78, TildeTildeToken}:                   _GotoState84Action,
	{_State78, BitNegToken}:                       _GotoState75Action,
	{_State78, BitAndToken}:                       _GotoState74Action,
	{_State78, ParseErrorToken}:                   _GotoState82Action,
	{_State78, InitializableTypeType}:             _GotoState91Action,
	{_State78, AtomTypeType}:                      _GotoState86Action,
	{_State78, ReturnableTypeType}:                _GotoState162Action,
	{_State78, ImplicitStructDefType}:             _GotoState90Action,
	{_State78, ExplicitStructDefType}:             _GotoState46Action,
	{_State78, ImplicitEnumDefType}:               _GotoState89Action,
	{_State78, ExplicitEnumDefType}:               _GotoState87Action,
	{_State78, TraitDefType}:                      _GotoState93Action,
	{_State78, FuncTypeType}:                      _GotoState88Action,
	{_State79, LparenToken}:                       _GotoState163Action,
	{_State80, DotToken}:                          _GotoState164Action,
	{_State80, DollarLbracketToken}:               _GotoState103Action,
	{_State80, OptionalGenericBindingType}:        _GotoState165Action,
	{_State81, IdentifierToken}:                   _GotoState166Action,
	{_State81, UnsafeToken}:                       _GotoState167Action,
	{_State81, StructToken}:                       _GotoState34Action,
	{_State81, EnumToken}:                         _GotoState77Action,
	{_State81, TraitToken}:                        _GotoState85Action,
	{_State81, FuncToken}:                         _GotoState79Action,
	{_State81, LparenToken}:                       _GotoState81Action,
	{_State81, LbracketToken}:                     _GotoState26Action,
	{_State81, DotToken}:                          _GotoState76Action,
	{_State81, QuestionToken}:                     _GotoState83Action,
	{_State81, ExclaimToken}:                      _GotoState78Action,
	{_State81, TildeTildeToken}:                   _GotoState84Action,
	{_State81, BitNegToken}:                       _GotoState75Action,
	{_State81, BitAndToken}:                       _GotoState74Action,
	{_State81, ParseErrorToken}:                   _GotoState82Action,
	{_State81, UnsafeStatementType}:               _GotoState173Action,
	{_State81, InitializableTypeType}:             _GotoState91Action,
	{_State81, AtomTypeType}:                      _GotoState86Action,
	{_State81, ReturnableTypeType}:                _GotoState92Action,
	{_State81, ValueTypeType}:                     _GotoState174Action,
	{_State81, FieldDefType}:                      _GotoState169Action,
	{_State81, ImplicitFieldDefsType}:             _GotoState171Action,
	{_State81, OptionalImplicitFieldDefsType}:     _GotoState172Action,
	{_State81, ImplicitStructDefType}:             _GotoState90Action,
	{_State81, ExplicitStructDefType}:             _GotoState46Action,
	{_State81, EnumValueDefType}:                  _GotoState168Action,
	{_State81, ImplicitEnumValueDefsType}:         _GotoState170Action,
	{_State81, ImplicitEnumDefType}:               _GotoState89Action,
	{_State81, ExplicitEnumDefType}:               _GotoState87Action,
	{_State81, TraitDefType}:                      _GotoState93Action,
	{_State81, FuncTypeType}:                      _GotoState88Action,
	{_State83, IdentifierToken}:                   _GotoState80Action,
	{_State83, StructToken}:                       _GotoState34Action,
	{_State83, EnumToken}:                         _GotoState77Action,
	{_State83, TraitToken}:                        _GotoState85Action,
	{_State83, FuncToken}:                         _GotoState79Action,
	{_State83, LparenToken}:                       _GotoState81Action,
	{_State83, LbracketToken}:                     _GotoState26Action,
	{_State83, DotToken}:                          _GotoState76Action,
	{_State83, QuestionToken}:                     _GotoState83Action,
	{_State83, ExclaimToken}:                      _GotoState78Action,
	{_State83, TildeTildeToken}:                   _GotoState84Action,
	{_State83, BitNegToken}:                       _GotoState75Action,
	{_State83, BitAndToken}:                       _GotoState74Action,
	{_State83, ParseErrorToken}:                   _GotoState82Action,
	{_State83, InitializableTypeType}:             _GotoState91Action,
	{_State83, AtomTypeType}:                      _GotoState86Action,
	{_State83, ReturnableTypeType}:                _GotoState175Action,
	{_State83, ImplicitStructDefType}:             _GotoState90Action,
	{_State83, ExplicitStructDefType}:             _GotoState46Action,
	{_State83, ImplicitEnumDefType}:               _GotoState89Action,
	{_State83, ExplicitEnumDefType}:               _GotoState87Action,
	{_State83, TraitDefType}:                      _GotoState93Action,
	{_State83, FuncTypeType}:                      _GotoState88Action,
	{_State84, IdentifierToken}:                   _GotoState80Action,
	{_State84, StructToken}:                       _GotoState34Action,
	{_State84, EnumToken}:                         _GotoState77Action,
	{_State84, TraitToken}:                        _GotoState85Action,
	{_State84, FuncToken}:                         _GotoState79Action,
	{_State84, LparenToken}:                       _GotoState81Action,
	{_State84, LbracketToken}:                     _GotoState26Action,
	{_State84, DotToken}:                          _GotoState76Action,
	{_State84, QuestionToken}:                     _GotoState83Action,
	{_State84, ExclaimToken}:                      _GotoState78Action,
	{_State84, TildeTildeToken}:                   _GotoState84Action,
	{_State84, BitNegToken}:                       _GotoState75Action,
	{_State84, BitAndToken}:                       _GotoState74Action,
	{_State84, ParseErrorToken}:                   _GotoState82Action,
	{_State84, InitializableTypeType}:             _GotoState91Action,
	{_State84, AtomTypeType}:                      _GotoState86Action,
	{_State84, ReturnableTypeType}:                _GotoState176Action,
	{_State84, ImplicitStructDefType}:             _GotoState90Action,
	{_State84, ExplicitStructDefType}:             _GotoState46Action,
	{_State84, ImplicitEnumDefType}:               _GotoState89Action,
	{_State84, ExplicitEnumDefType}:               _GotoState87Action,
	{_State84, TraitDefType}:                      _GotoState93Action,
	{_State84, FuncTypeType}:                      _GotoState88Action,
	{_State85, LparenToken}:                       _GotoState177Action,
	{_State94, RbracketToken}:                     _GotoState182Action,
	{_State94, CommaToken}:                        _GotoState180Action,
	{_State94, ColonToken}:                        _GotoState179Action,
	{_State94, AddToken}:                          _GotoState178Action,
	{_State94, SubToken}:                          _GotoState183Action,
	{_State94, MulToken}:                          _GotoState181Action,
	{_State96, AssignToken}:                       _GotoState184Action,
	{_State98, RparenToken}:                       _GotoState186Action,
	{_State98, CommaToken}:                        _GotoState185Action,
	{_State99, ColonToken}:                        _GotoState187Action,
	{_State101, ColonToken}:                       _GotoState188Action,
	{_State102, IdentifierToken}:                  _GotoState166Action,
	{_State102, UnsafeToken}:                      _GotoState167Action,
	{_State102, StructToken}:                      _GotoState34Action,
	{_State102, EnumToken}:                        _GotoState77Action,
	{_State102, TraitToken}:                       _GotoState85Action,
	{_State102, FuncToken}:                        _GotoState79Action,
	{_State102, LparenToken}:                      _GotoState81Action,
	{_State102, LbracketToken}:                    _GotoState26Action,
	{_State102, DotToken}:                         _GotoState76Action,
	{_State102, QuestionToken}:                    _GotoState83Action,
	{_State102, ExclaimToken}:                     _GotoState78Action,
	{_State102, TildeTildeToken}:                  _GotoState84Action,
	{_State102, BitNegToken}:                      _GotoState75Action,
	{_State102, BitAndToken}:                      _GotoState74Action,
	{_State102, ParseErrorToken}:                  _GotoState82Action,
	{_State102, UnsafeStatementType}:              _GotoState173Action,
	{_State102, InitializableTypeType}:            _GotoState91Action,
	{_State102, AtomTypeType}:                     _GotoState86Action,
	{_State102, ReturnableTypeType}:               _GotoState92Action,
	{_State102, ValueTypeType}:                    _GotoState174Action,
	{_State102, FieldDefType}:                     _GotoState190Action,
	{_State102, ImplicitStructDefType}:            _GotoState90Action,
	{_State102, ExplicitFieldDefsType}:            _GotoState189Action,
	{_State102, OptionalExplicitFieldDefsType}:    _GotoState191Action,
	{_State102, ExplicitStructDefType}:            _GotoState46Action,
	{_State102, ImplicitEnumDefType}:              _GotoState89Action,
	{_State102, ExplicitEnumDefType}:              _GotoState87Action,
	{_State102, TraitDefType}:                     _GotoState93Action,
	{_State102, FuncTypeType}:                     _GotoState88Action,
	{_State103, IdentifierToken}:                  _GotoState80Action,
	{_State103, StructToken}:                      _GotoState34Action,
	{_State103, EnumToken}:                        _GotoState77Action,
	{_State103, TraitToken}:                       _GotoState85Action,
	{_State103, FuncToken}:                        _GotoState79Action,
	{_State103, LparenToken}:                      _GotoState81Action,
	{_State103, LbracketToken}:                    _GotoState26Action,
	{_State103, DotToken}:                         _GotoState76Action,
	{_State103, QuestionToken}:                    _GotoState83Action,
	{_State103, ExclaimToken}:                     _GotoState78Action,
	{_State103, TildeTildeToken}:                  _GotoState84Action,
	{_State103, BitNegToken}:                      _GotoState75Action,
	{_State103, BitAndToken}:                      _GotoState74Action,
	{_State103, ParseErrorToken}:                  _GotoState82Action,
	{_State103, OptionalGenericArgumentsType}:     _GotoState193Action,
	{_State103, GenericArgumentsType}:             _GotoState192Action,
	{_State103, InitializableTypeType}:            _GotoState91Action,
	{_State103, AtomTypeType}:                     _GotoState86Action,
	{_State103, ReturnableTypeType}:               _GotoState92Action,
	{_State103, ValueTypeType}:                    _GotoState194Action,
	{_State103, ImplicitStructDefType}:            _GotoState90Action,
	{_State103, ExplicitStructDefType}:            _GotoState46Action,
	{_State103, ImplicitEnumDefType}:              _GotoState89Action,
	{_State103, ExplicitEnumDefType}:              _GotoState87Action,
	{_State103, TraitDefType}:                     _GotoState93Action,
	{_State103, FuncTypeType}:                     _GotoState88Action,
	{_State104, IdentifierToken}:                  _GotoState195Action,
	{_State105, IntegerLiteralToken}:              _GotoState24Action,
	{_State105, FloatLiteralToken}:                _GotoState20Action,
	{_State105, RuneLiteralToken}:                 _GotoState32Action,
	{_State105, StringLiteralToken}:               _GotoState33Action,
	{_State105, IdentifierToken}:                  _GotoState96Action,
	{_State105, TrueToken}:                        _GotoState36Action,
	{_State105, FalseToken}:                       _GotoState19Action,
	{_State105, StructToken}:                      _GotoState34Action,
	{_State105, FuncToken}:                        _GotoState21Action,
	{_State105, VarToken}:                         _GotoState37Action,
	{_State105, LetToken}:                         _GotoState27Action,
	{_State105, LabelDeclToken}:                   _GotoState25Action,
	{_State105, LparenToken}:                      _GotoState28Action,
	{_State105, LbracketToken}:                    _GotoState26Action,
	{_State105, DotDotDotToken}:                   _GotoState95Action,
	{_State105, NotToken}:                         _GotoState30Action,
	{_State105, SubToken}:                         _GotoState35Action,
	{_State105, MulToken}:                         _GotoState29Action,
	{_State105, BitNegToken}:                      _GotoState18Action,
	{_State105, BitAndToken}:                      _GotoState17Action,
	{_State105, GreaterToken}:                     _GotoState22Action,
	{_State105, ParseErrorToken}:                  _GotoState31Action,
	{_State105, VarDeclPatternType}:               _GotoState57Action,
	{_State105, VarOrLetType}:                     _GotoState58Action,
	{_State105, ExpressionType}:                   _GotoState100Action,
	{_State105, OptionalLabelDeclType}:            _GotoState51Action,
	{_State105, SequenceExprType}:                 _GotoState56Action,
	{_State105, BlockExprType}:                    _GotoState43Action,
	{_State105, CallExprType}:                     _GotoState44Action,
	{_State105, ArgumentType}:                     _GotoState196Action,
	{_State105, ColonExpressionsType}:             _GotoState99Action,
	{_State105, OptionalExpressionType}:           _GotoState101Action,
	{_State105, AtomExprType}:                     _GotoState42Action,
	{_State105, LiteralType}:                      _GotoState49Action,
	{_State105, ImplicitStructExprType}:           _GotoState47Action,
	{_State105, AccessExprType}:                   _GotoState38Action,
	{_State105, PostfixUnaryExprType}:             _GotoState53Action,
	{_State105, PrefixUnaryOpType}:                _GotoState55Action,
	{_State105, PrefixUnaryExprType}:              _GotoState54Action,
	{_State105, MulExprType}:                      _GotoState50Action,
	{_State105, AddExprType}:                      _GotoState39Action,
	{_State105, CmpExprType}:                      _GotoState45Action,
	{_State105, AndExprType}:                      _GotoState40Action,
	{_State105, OrExprType}:                       _GotoState52Action,
	{_State105, InitializableTypeType}:            _GotoState48Action,
	{_State105, ExplicitStructDefType}:            _GotoState46Action,
	{_State105, AnonymousFuncExprType}:            _GotoState41Action,
	{_State107, LparenToken}:                      _GotoState197Action,
	{_State112, IntegerLiteralToken}:              _GotoState24Action,
	{_State112, FloatLiteralToken}:                _GotoState20Action,
	{_State112, RuneLiteralToken}:                 _GotoState32Action,
	{_State112, StringLiteralToken}:               _GotoState33Action,
	{_State112, IdentifierToken}:                  _GotoState23Action,
	{_State112, TrueToken}:                        _GotoState36Action,
	{_State112, FalseToken}:                       _GotoState19Action,
	{_State112, StructToken}:                      _GotoState34Action,
	{_State112, FuncToken}:                        _GotoState21Action,
	{_State112, LabelDeclToken}:                   _GotoState25Action,
	{_State112, LparenToken}:                      _GotoState28Action,
	{_State112, LbracketToken}:                    _GotoState26Action,
	{_State112, NotToken}:                         _GotoState30Action,
	{_State112, SubToken}:                         _GotoState35Action,
	{_State112, MulToken}:                         _GotoState29Action,
	{_State112, BitNegToken}:                      _GotoState18Action,
	{_State112, BitAndToken}:                      _GotoState17Action,
	{_State112, ParseErrorToken}:                  _GotoState31Action,
	{_State112, OptionalLabelDeclType}:            _GotoState72Action,
	{_State112, BlockExprType}:                    _GotoState43Action,
	{_State112, CallExprType}:                     _GotoState44Action,
	{_State112, AtomExprType}:                     _GotoState42Action,
	{_State112, LiteralType}:                      _GotoState49Action,
	{_State112, ImplicitStructExprType}:           _GotoState47Action,
	{_State112, AccessExprType}:                   _GotoState38Action,
	{_State112, PostfixUnaryExprType}:             _GotoState53Action,
	{_State112, PrefixUnaryOpType}:                _GotoState55Action,
	{_State112, PrefixUnaryExprType}:              _GotoState54Action,
	{_State112, MulExprType}:                      _GotoState198Action,
	{_State112, InitializableTypeType}:            _GotoState48Action,
	{_State112, ExplicitStructDefType}:            _GotoState46Action,
	{_State112, AnonymousFuncExprType}:            _GotoState41Action,
	{_State113, IntegerLiteralToken}:              _GotoState24Action,
	{_State113, FloatLiteralToken}:                _GotoState20Action,
	{_State113, RuneLiteralToken}:                 _GotoState32Action,
	{_State113, StringLiteralToken}:               _GotoState33Action,
	{_State113, IdentifierToken}:                  _GotoState23Action,
	{_State113, TrueToken}:                        _GotoState36Action,
	{_State113, FalseToken}:                       _GotoState19Action,
	{_State113, StructToken}:                      _GotoState34Action,
	{_State113, FuncToken}:                        _GotoState21Action,
	{_State113, LabelDeclToken}:                   _GotoState25Action,
	{_State113, LparenToken}:                      _GotoState28Action,
	{_State113, LbracketToken}:                    _GotoState26Action,
	{_State113, NotToken}:                         _GotoState30Action,
	{_State113, SubToken}:                         _GotoState35Action,
	{_State113, MulToken}:                         _GotoState29Action,
	{_State113, BitNegToken}:                      _GotoState18Action,
	{_State113, BitAndToken}:                      _GotoState17Action,
	{_State113, ParseErrorToken}:                  _GotoState31Action,
	{_State113, OptionalLabelDeclType}:            _GotoState72Action,
	{_State113, BlockExprType}:                    _GotoState43Action,
	{_State113, CallExprType}:                     _GotoState44Action,
	{_State113, AtomExprType}:                     _GotoState42Action,
	{_State113, LiteralType}:                      _GotoState49Action,
	{_State113, ImplicitStructExprType}:           _GotoState47Action,
	{_State113, AccessExprType}:                   _GotoState38Action,
	{_State113, PostfixUnaryExprType}:             _GotoState53Action,
	{_State113, PrefixUnaryOpType}:                _GotoState55Action,
	{_State113, PrefixUnaryExprType}:              _GotoState54Action,
	{_State113, MulExprType}:                      _GotoState50Action,
	{_State113, AddExprType}:                      _GotoState39Action,
	{_State113, CmpExprType}:                      _GotoState199Action,
	{_State113, InitializableTypeType}:            _GotoState48Action,
	{_State113, ExplicitStructDefType}:            _GotoState46Action,
	{_State113, AnonymousFuncExprType}:            _GotoState41Action,
	{_State120, IntegerLiteralToken}:              _GotoState24Action,
	{_State120, FloatLiteralToken}:                _GotoState20Action,
	{_State120, RuneLiteralToken}:                 _GotoState32Action,
	{_State120, StringLiteralToken}:               _GotoState33Action,
	{_State120, IdentifierToken}:                  _GotoState23Action,
	{_State120, TrueToken}:                        _GotoState36Action,
	{_State120, FalseToken}:                       _GotoState19Action,
	{_State120, StructToken}:                      _GotoState34Action,
	{_State120, FuncToken}:                        _GotoState21Action,
	{_State120, LabelDeclToken}:                   _GotoState25Action,
	{_State120, LparenToken}:                      _GotoState28Action,
	{_State120, LbracketToken}:                    _GotoState26Action,
	{_State120, NotToken}:                         _GotoState30Action,
	{_State120, SubToken}:                         _GotoState35Action,
	{_State120, MulToken}:                         _GotoState29Action,
	{_State120, BitNegToken}:                      _GotoState18Action,
	{_State120, BitAndToken}:                      _GotoState17Action,
	{_State120, ParseErrorToken}:                  _GotoState31Action,
	{_State120, OptionalLabelDeclType}:            _GotoState72Action,
	{_State120, BlockExprType}:                    _GotoState43Action,
	{_State120, CallExprType}:                     _GotoState44Action,
	{_State120, AtomExprType}:                     _GotoState42Action,
	{_State120, LiteralType}:                      _GotoState49Action,
	{_State120, ImplicitStructExprType}:           _GotoState47Action,
	{_State120, AccessExprType}:                   _GotoState38Action,
	{_State120, PostfixUnaryExprType}:             _GotoState53Action,
	{_State120, PrefixUnaryOpType}:                _GotoState55Action,
	{_State120, PrefixUnaryExprType}:              _GotoState54Action,
	{_State120, MulExprType}:                      _GotoState50Action,
	{_State120, AddExprType}:                      _GotoState200Action,
	{_State120, InitializableTypeType}:            _GotoState48Action,
	{_State120, ExplicitStructDefType}:            _GotoState46Action,
	{_State120, AnonymousFuncExprType}:            _GotoState41Action,
	{_State121, IntegerLiteralToken}:              _GotoState24Action,
	{_State121, FloatLiteralToken}:                _GotoState20Action,
	{_State121, RuneLiteralToken}:                 _GotoState32Action,
	{_State121, StringLiteralToken}:               _GotoState33Action,
	{_State121, IdentifierToken}:                  _GotoState96Action,
	{_State121, TrueToken}:                        _GotoState36Action,
	{_State121, FalseToken}:                       _GotoState19Action,
	{_State121, StructToken}:                      _GotoState34Action,
	{_State121, FuncToken}:                        _GotoState21Action,
	{_State121, VarToken}:                         _GotoState37Action,
	{_State121, LetToken}:                         _GotoState27Action,
	{_State121, LabelDeclToken}:                   _GotoState25Action,
	{_State121, LparenToken}:                      _GotoState28Action,
	{_State121, LbracketToken}:                    _GotoState26Action,
	{_State121, DotDotDotToken}:                   _GotoState95Action,
	{_State121, NotToken}:                         _GotoState30Action,
	{_State121, SubToken}:                         _GotoState35Action,
	{_State121, MulToken}:                         _GotoState29Action,
	{_State121, BitNegToken}:                      _GotoState18Action,
	{_State121, BitAndToken}:                      _GotoState17Action,
	{_State121, GreaterToken}:                     _GotoState22Action,
	{_State121, ParseErrorToken}:                  _GotoState31Action,
	{_State121, VarDeclPatternType}:               _GotoState57Action,
	{_State121, VarOrLetType}:                     _GotoState58Action,
	{_State121, ExpressionType}:                   _GotoState100Action,
	{_State121, OptionalLabelDeclType}:            _GotoState51Action,
	{_State121, SequenceExprType}:                 _GotoState56Action,
	{_State121, BlockExprType}:                    _GotoState43Action,
	{_State121, CallExprType}:                     _GotoState44Action,
	{_State121, ArgumentsType}:                    _GotoState201Action,
	{_State121, ArgumentType}:                     _GotoState97Action,
	{_State121, ColonExpressionsType}:             _GotoState99Action,
	{_State121, OptionalExpressionType}:           _GotoState101Action,
	{_State121, AtomExprType}:                     _GotoState42Action,
	{_State121, LiteralType}:                      _GotoState49Action,
	{_State121, ImplicitStructExprType}:           _GotoState47Action,
	{_State121, AccessExprType}:                   _GotoState38Action,
	{_State121, PostfixUnaryExprType}:             _GotoState53Action,
	{_State121, PrefixUnaryOpType}:                _GotoState55Action,
	{_State121, PrefixUnaryExprType}:              _GotoState54Action,
	{_State121, MulExprType}:                      _GotoState50Action,
	{_State121, AddExprType}:                      _GotoState39Action,
	{_State121, CmpExprType}:                      _GotoState45Action,
	{_State121, AndExprType}:                      _GotoState40Action,
	{_State121, OrExprType}:                       _GotoState52Action,
	{_State121, InitializableTypeType}:            _GotoState48Action,
	{_State121, ExplicitStructDefType}:            _GotoState46Action,
	{_State121, AnonymousFuncExprType}:            _GotoState41Action,
	{_State128, IntegerLiteralToken}:              _GotoState24Action,
	{_State128, FloatLiteralToken}:                _GotoState20Action,
	{_State128, RuneLiteralToken}:                 _GotoState32Action,
	{_State128, StringLiteralToken}:               _GotoState33Action,
	{_State128, IdentifierToken}:                  _GotoState23Action,
	{_State128, TrueToken}:                        _GotoState36Action,
	{_State128, FalseToken}:                       _GotoState19Action,
	{_State128, StructToken}:                      _GotoState34Action,
	{_State128, FuncToken}:                        _GotoState21Action,
	{_State128, LabelDeclToken}:                   _GotoState25Action,
	{_State128, LparenToken}:                      _GotoState28Action,
	{_State128, LbracketToken}:                    _GotoState26Action,
	{_State128, NotToken}:                         _GotoState30Action,
	{_State128, SubToken}:                         _GotoState35Action,
	{_State128, MulToken}:                         _GotoState29Action,
	{_State128, BitNegToken}:                      _GotoState18Action,
	{_State128, BitAndToken}:                      _GotoState17Action,
	{_State128, ParseErrorToken}:                  _GotoState31Action,
	{_State128, OptionalLabelDeclType}:            _GotoState72Action,
	{_State128, BlockExprType}:                    _GotoState43Action,
	{_State128, CallExprType}:                     _GotoState44Action,
	{_State128, AtomExprType}:                     _GotoState42Action,
	{_State128, LiteralType}:                      _GotoState49Action,
	{_State128, ImplicitStructExprType}:           _GotoState47Action,
	{_State128, AccessExprType}:                   _GotoState38Action,
	{_State128, PostfixUnaryExprType}:             _GotoState53Action,
	{_State128, PrefixUnaryOpType}:                _GotoState55Action,
	{_State128, PrefixUnaryExprType}:              _GotoState202Action,
	{_State128, InitializableTypeType}:            _GotoState48Action,
	{_State128, ExplicitStructDefType}:            _GotoState46Action,
	{_State128, AnonymousFuncExprType}:            _GotoState41Action,
	{_State129, LbraceToken}:                      _GotoState59Action,
	{_State129, BlockBodyType}:                    _GotoState203Action,
	{_State130, IntegerLiteralToken}:              _GotoState24Action,
	{_State130, FloatLiteralToken}:                _GotoState20Action,
	{_State130, RuneLiteralToken}:                 _GotoState32Action,
	{_State130, StringLiteralToken}:               _GotoState33Action,
	{_State130, IdentifierToken}:                  _GotoState23Action,
	{_State130, TrueToken}:                        _GotoState36Action,
	{_State130, FalseToken}:                       _GotoState19Action,
	{_State130, StructToken}:                      _GotoState34Action,
	{_State130, FuncToken}:                        _GotoState21Action,
	{_State130, VarToken}:                         _GotoState37Action,
	{_State130, LetToken}:                         _GotoState27Action,
	{_State130, LabelDeclToken}:                   _GotoState25Action,
	{_State130, LparenToken}:                      _GotoState28Action,
	{_State130, LbracketToken}:                    _GotoState26Action,
	{_State130, NotToken}:                         _GotoState30Action,
	{_State130, SubToken}:                         _GotoState35Action,
	{_State130, MulToken}:                         _GotoState29Action,
	{_State130, BitNegToken}:                      _GotoState18Action,
	{_State130, BitAndToken}:                      _GotoState17Action,
	{_State130, GreaterToken}:                     _GotoState22Action,
	{_State130, ParseErrorToken}:                  _GotoState31Action,
	{_State130, VarDeclPatternType}:               _GotoState57Action,
	{_State130, VarOrLetType}:                     _GotoState58Action,
	{_State130, AssignPatternType}:                _GotoState204Action,
	{_State130, OptionalLabelDeclType}:            _GotoState72Action,
	{_State130, ForAssignmentType}:                _GotoState205Action,
	{_State130, SequenceExprType}:                 _GotoState206Action,
	{_State130, BlockExprType}:                    _GotoState43Action,
	{_State130, CallExprType}:                     _GotoState44Action,
	{_State130, AtomExprType}:                     _GotoState42Action,
	{_State130, LiteralType}:                      _GotoState49Action,
	{_State130, ImplicitStructExprType}:           _GotoState47Action,
	{_State130, AccessExprType}:                   _GotoState38Action,
	{_State130, PostfixUnaryExprType}:             _GotoState53Action,
	{_State130, PrefixUnaryOpType}:                _GotoState55Action,
	{_State130, PrefixUnaryExprType}:              _GotoState54Action,
	{_State130, MulExprType}:                      _GotoState50Action,
	{_State130, AddExprType}:                      _GotoState39Action,
	{_State130, CmpExprType}:                      _GotoState45Action,
	{_State130, AndExprType}:                      _GotoState40Action,
	{_State130, OrExprType}:                       _GotoState52Action,
	{_State130, InitializableTypeType}:            _GotoState48Action,
	{_State130, ExplicitStructDefType}:            _GotoState46Action,
	{_State130, AnonymousFuncExprType}:            _GotoState41Action,
	{_State131, IntegerLiteralToken}:              _GotoState24Action,
	{_State131, FloatLiteralToken}:                _GotoState20Action,
	{_State131, RuneLiteralToken}:                 _GotoState32Action,
	{_State131, StringLiteralToken}:               _GotoState33Action,
	{_State131, IdentifierToken}:                  _GotoState23Action,
	{_State131, TrueToken}:                        _GotoState36Action,
	{_State131, FalseToken}:                       _GotoState19Action,
	{_State131, CaseToken}:                        _GotoState207Action,
	{_State131, StructToken}:                      _GotoState34Action,
	{_State131, FuncToken}:                        _GotoState21Action,
	{_State131, VarToken}:                         _GotoState37Action,
	{_State131, LetToken}:                         _GotoState27Action,
	{_State131, LabelDeclToken}:                   _GotoState25Action,
	{_State131, LparenToken}:                      _GotoState28Action,
	{_State131, LbracketToken}:                    _GotoState26Action,
	{_State131, NotToken}:                         _GotoState30Action,
	{_State131, SubToken}:                         _GotoState35Action,
	{_State131, MulToken}:                         _GotoState29Action,
	{_State131, BitNegToken}:                      _GotoState18Action,
	{_State131, BitAndToken}:                      _GotoState17Action,
	{_State131, GreaterToken}:                     _GotoState22Action,
	{_State131, ParseErrorToken}:                  _GotoState31Action,
	{_State131, VarDeclPatternType}:               _GotoState57Action,
	{_State131, VarOrLetType}:                     _GotoState58Action,
	{_State131, OptionalLabelDeclType}:            _GotoState72Action,
	{_State131, ConditionType}:                    _GotoState208Action,
	{_State131, SequenceExprType}:                 _GotoState209Action,
	{_State131, BlockExprType}:                    _GotoState43Action,
	{_State131, CallExprType}:                     _GotoState44Action,
	{_State131, AtomExprType}:                     _GotoState42Action,
	{_State131, LiteralType}:                      _GotoState49Action,
	{_State131, ImplicitStructExprType}:           _GotoState47Action,
	{_State131, AccessExprType}:                   _GotoState38Action,
	{_State131, PostfixUnaryExprType}:             _GotoState53Action,
	{_State131, PrefixUnaryOpType}:                _GotoState55Action,
	{_State131, PrefixUnaryExprType}:              _GotoState54Action,
	{_State131, MulExprType}:                      _GotoState50Action,
	{_State131, AddExprType}:                      _GotoState39Action,
	{_State131, CmpExprType}:                      _GotoState45Action,
	{_State131, AndExprType}:                      _GotoState40Action,
	{_State131, OrExprType}:                       _GotoState52Action,
	{_State131, InitializableTypeType}:            _GotoState48Action,
	{_State131, ExplicitStructDefType}:            _GotoState46Action,
	{_State131, AnonymousFuncExprType}:            _GotoState41Action,
	{_State132, IntegerLiteralToken}:              _GotoState24Action,
	{_State132, FloatLiteralToken}:                _GotoState20Action,
	{_State132, RuneLiteralToken}:                 _GotoState32Action,
	{_State132, StringLiteralToken}:               _GotoState33Action,
	{_State132, IdentifierToken}:                  _GotoState23Action,
	{_State132, TrueToken}:                        _GotoState36Action,
	{_State132, FalseToken}:                       _GotoState19Action,
	{_State132, StructToken}:                      _GotoState34Action,
	{_State132, FuncToken}:                        _GotoState21Action,
	{_State132, VarToken}:                         _GotoState37Action,
	{_State132, LetToken}:                         _GotoState27Action,
	{_State132, LabelDeclToken}:                   _GotoState25Action,
	{_State132, LparenToken}:                      _GotoState28Action,
	{_State132, LbracketToken}:                    _GotoState26Action,
	{_State132, NotToken}:                         _GotoState30Action,
	{_State132, SubToken}:                         _GotoState35Action,
	{_State132, MulToken}:                         _GotoState29Action,
	{_State132, BitNegToken}:                      _GotoState18Action,
	{_State132, BitAndToken}:                      _GotoState17Action,
	{_State132, GreaterToken}:                     _GotoState22Action,
	{_State132, ParseErrorToken}:                  _GotoState31Action,
	{_State132, VarDeclPatternType}:               _GotoState57Action,
	{_State132, VarOrLetType}:                     _GotoState58Action,
	{_State132, OptionalLabelDeclType}:            _GotoState72Action,
	{_State132, SequenceExprType}:                 _GotoState210Action,
	{_State132, BlockExprType}:                    _GotoState43Action,
	{_State132, CallExprType}:                     _GotoState44Action,
	{_State132, AtomExprType}:                     _GotoState42Action,
	{_State132, LiteralType}:                      _GotoState49Action,
	{_State132, ImplicitStructExprType}:           _GotoState47Action,
	{_State132, AccessExprType}:                   _GotoState38Action,
	{_State132, PostfixUnaryExprType}:             _GotoState53Action,
	{_State132, PrefixUnaryOpType}:                _GotoState55Action,
	{_State132, PrefixUnaryExprType}:              _GotoState54Action,
	{_State132, MulExprType}:                      _GotoState50Action,
	{_State132, AddExprType}:                      _GotoState39Action,
	{_State132, CmpExprType}:                      _GotoState45Action,
	{_State132, AndExprType}:                      _GotoState40Action,
	{_State132, OrExprType}:                       _GotoState52Action,
	{_State132, InitializableTypeType}:            _GotoState48Action,
	{_State132, ExplicitStructDefType}:            _GotoState46Action,
	{_State132, AnonymousFuncExprType}:            _GotoState41Action,
	{_State137, IntegerLiteralToken}:              _GotoState24Action,
	{_State137, FloatLiteralToken}:                _GotoState20Action,
	{_State137, RuneLiteralToken}:                 _GotoState32Action,
	{_State137, StringLiteralToken}:               _GotoState33Action,
	{_State137, IdentifierToken}:                  _GotoState23Action,
	{_State137, TrueToken}:                        _GotoState36Action,
	{_State137, FalseToken}:                       _GotoState19Action,
	{_State137, StructToken}:                      _GotoState34Action,
	{_State137, FuncToken}:                        _GotoState21Action,
	{_State137, LabelDeclToken}:                   _GotoState25Action,
	{_State137, LparenToken}:                      _GotoState28Action,
	{_State137, LbracketToken}:                    _GotoState26Action,
	{_State137, NotToken}:                         _GotoState30Action,
	{_State137, SubToken}:                         _GotoState35Action,
	{_State137, MulToken}:                         _GotoState29Action,
	{_State137, BitNegToken}:                      _GotoState18Action,
	{_State137, BitAndToken}:                      _GotoState17Action,
	{_State137, ParseErrorToken}:                  _GotoState31Action,
	{_State137, OptionalLabelDeclType}:            _GotoState72Action,
	{_State137, BlockExprType}:                    _GotoState43Action,
	{_State137, CallExprType}:                     _GotoState44Action,
	{_State137, AtomExprType}:                     _GotoState42Action,
	{_State137, LiteralType}:                      _GotoState49Action,
	{_State137, ImplicitStructExprType}:           _GotoState47Action,
	{_State137, AccessExprType}:                   _GotoState38Action,
	{_State137, PostfixUnaryExprType}:             _GotoState53Action,
	{_State137, PrefixUnaryOpType}:                _GotoState55Action,
	{_State137, PrefixUnaryExprType}:              _GotoState54Action,
	{_State137, MulExprType}:                      _GotoState50Action,
	{_State137, AddExprType}:                      _GotoState39Action,
	{_State137, CmpExprType}:                      _GotoState45Action,
	{_State137, AndExprType}:                      _GotoState211Action,
	{_State137, InitializableTypeType}:            _GotoState48Action,
	{_State137, ExplicitStructDefType}:            _GotoState46Action,
	{_State137, AnonymousFuncExprType}:            _GotoState41Action,
	{_State140, IdentifierToken}:                  _GotoState213Action,
	{_State140, LparenToken}:                      _GotoState140Action,
	{_State140, DotDotDotToken}:                   _GotoState212Action,
	{_State140, VarPatternType}:                   _GotoState216Action,
	{_State140, TuplePatternType}:                 _GotoState141Action,
	{_State140, FieldVarPatternsType}:             _GotoState215Action,
	{_State140, FieldVarPatternType}:              _GotoState214Action,
	{_State142, IdentifierToken}:                  _GotoState80Action,
	{_State142, StructToken}:                      _GotoState34Action,
	{_State142, EnumToken}:                        _GotoState77Action,
	{_State142, TraitToken}:                       _GotoState85Action,
	{_State142, FuncToken}:                        _GotoState79Action,
	{_State142, LparenToken}:                      _GotoState81Action,
	{_State142, LbracketToken}:                    _GotoState26Action,
	{_State142, DotToken}:                         _GotoState76Action,
	{_State142, QuestionToken}:                    _GotoState83Action,
	{_State142, ExclaimToken}:                     _GotoState78Action,
	{_State142, TildeTildeToken}:                  _GotoState84Action,
	{_State142, BitNegToken}:                      _GotoState75Action,
	{_State142, BitAndToken}:                      _GotoState74Action,
	{_State142, ParseErrorToken}:                  _GotoState82Action,
	{_State142, OptionalValueTypeType}:            _GotoState217Action,
	{_State142, InitializableTypeType}:            _GotoState91Action,
	{_State142, AtomTypeType}:                     _GotoState86Action,
	{_State142, ReturnableTypeType}:               _GotoState92Action,
	{_State142, ValueTypeType}:                    _GotoState218Action,
	{_State142, ImplicitStructDefType}:            _GotoState90Action,
	{_State142, ExplicitStructDefType}:            _GotoState46Action,
	{_State142, ImplicitEnumDefType}:              _GotoState89Action,
	{_State142, ExplicitEnumDefType}:              _GotoState87Action,
	{_State142, TraitDefType}:                     _GotoState93Action,
	{_State142, FuncTypeType}:                     _GotoState88Action,
	{_State143, IntegerLiteralToken}:              _GotoState24Action,
	{_State143, FloatLiteralToken}:                _GotoState20Action,
	{_State143, RuneLiteralToken}:                 _GotoState32Action,
	{_State143, StringLiteralToken}:               _GotoState33Action,
	{_State143, IdentifierToken}:                  _GotoState23Action,
	{_State143, TrueToken}:                        _GotoState36Action,
	{_State143, FalseToken}:                       _GotoState19Action,
	{_State143, ReturnToken}:                      _GotoState224Action,
	{_State143, BreakToken}:                       _GotoState220Action,
	{_State143, ContinueToken}:                    _GotoState221Action,
	{_State143, UnsafeToken}:                      _GotoState167Action,
	{_State143, StructToken}:                      _GotoState34Action,
	{_State143, FuncToken}:                        _GotoState21Action,
	{_State143, AsyncToken}:                       _GotoState219Action,
	{_State143, DeferToken}:                       _GotoState222Action,
	{_State143, VarToken}:                         _GotoState37Action,
	{_State143, LetToken}:                         _GotoState27Action,
	{_State143, LabelDeclToken}:                   _GotoState25Action,
	{_State143, RbraceToken}:                      _GotoState223Action,
	{_State143, LparenToken}:                      _GotoState28Action,
	{_State143, LbracketToken}:                    _GotoState26Action,
	{_State143, NotToken}:                         _GotoState30Action,
	{_State143, SubToken}:                         _GotoState35Action,
	{_State143, MulToken}:                         _GotoState29Action,
	{_State143, BitNegToken}:                      _GotoState18Action,
	{_State143, BitAndToken}:                      _GotoState17Action,
	{_State143, GreaterToken}:                     _GotoState22Action,
	{_State143, ParseErrorToken}:                  _GotoState31Action,
	{_State143, VarDeclPatternType}:               _GotoState57Action,
	{_State143, VarOrLetType}:                     _GotoState58Action,
	{_State143, AssignPatternType}:                _GotoState226Action,
	{_State143, ExpressionType}:                   _GotoState227Action,
	{_State143, OptionalLabelDeclType}:            _GotoState51Action,
	{_State143, SequenceExprType}:                 _GotoState231Action,
	{_State143, BlockExprType}:                    _GotoState43Action,
	{_State143, StatementType}:                    _GotoState232Action,
	{_State143, StatementBodyType}:                _GotoState233Action,
	{_State143, UnsafeStatementType}:              _GotoState234Action,
	{_State143, JumpStatementType}:                _GotoState229Action,
	{_State143, JumpTypeType}:                     _GotoState230Action,
	{_State143, ExpressionsType}:                  _GotoState228Action,
	{_State143, CallExprType}:                     _GotoState44Action,
	{_State143, AtomExprType}:                     _GotoState42Action,
	{_State143, LiteralType}:                      _GotoState49Action,
	{_State143, ImplicitStructExprType}:           _GotoState47Action,
	{_State143, AccessExprType}:                   _GotoState225Action,
	{_State143, PostfixUnaryExprType}:             _GotoState53Action,
	{_State143, PrefixUnaryOpType}:                _GotoState55Action,
	{_State143, PrefixUnaryExprType}:              _GotoState54Action,
	{_State143, MulExprType}:                      _GotoState50Action,
	{_State143, AddExprType}:                      _GotoState39Action,
	{_State143, CmpExprType}:                      _GotoState45Action,
	{_State143, AndExprType}:                      _GotoState40Action,
	{_State143, OrExprType}:                       _GotoState52Action,
	{_State143, InitializableTypeType}:            _GotoState48Action,
	{_State143, ExplicitStructDefType}:            _GotoState46Action,
	{_State143, AnonymousFuncExprType}:            _GotoState41Action,
	{_State144, PackageToken}:                     _GotoState14Action,
	{_State144, TypeToken}:                        _GotoState15Action,
	{_State144, FuncToken}:                        _GotoState16Action,
	{_State144, VarToken}:                         _GotoState37Action,
	{_State144, LetToken}:                         _GotoState27Action,
	{_State144, LbraceToken}:                      _GotoState59Action,
	{_State144, DefinitionType}:                   _GotoState235Action,
	{_State144, VarDeclPatternType}:               _GotoState66Action,
	{_State144, VarOrLetType}:                     _GotoState58Action,
	{_State144, BlockBodyType}:                    _GotoState60Action,
	{_State144, TypeDefType}:                      _GotoState65Action,
	{_State144, NamedFuncDefType}:                 _GotoState63Action,
	{_State144, PackageDefType}:                   _GotoState64Action,
	{_State146, IntegerLiteralToken}:              _GotoState24Action,
	{_State146, FloatLiteralToken}:                _GotoState20Action,
	{_State146, RuneLiteralToken}:                 _GotoState32Action,
	{_State146, StringLiteralToken}:               _GotoState33Action,
	{_State146, IdentifierToken}:                  _GotoState23Action,
	{_State146, TrueToken}:                        _GotoState36Action,
	{_State146, FalseToken}:                       _GotoState19Action,
	{_State146, StructToken}:                      _GotoState34Action,
	{_State146, FuncToken}:                        _GotoState21Action,
	{_State146, VarToken}:                         _GotoState37Action,
	{_State146, LetToken}:                         _GotoState27Action,
	{_State146, LabelDeclToken}:                   _GotoState25Action,
	{_State146, LparenToken}:                      _GotoState28Action,
	{_State146, LbracketToken}:                    _GotoState26Action,
	{_State146, NotToken}:                         _GotoState30Action,
	{_State146, SubToken}:                         _GotoState35Action,
	{_State146, MulToken}:                         _GotoState29Action,
	{_State146, BitNegToken}:                      _GotoState18Action,
	{_State146, BitAndToken}:                      _GotoState17Action,
	{_State146, GreaterToken}:                     _GotoState22Action,
	{_State146, ParseErrorToken}:                  _GotoState31Action,
	{_State146, VarDeclPatternType}:               _GotoState57Action,
	{_State146, VarOrLetType}:                     _GotoState58Action,
	{_State146, ExpressionType}:                   _GotoState236Action,
	{_State146, OptionalLabelDeclType}:            _GotoState51Action,
	{_State146, SequenceExprType}:                 _GotoState56Action,
	{_State146, BlockExprType}:                    _GotoState43Action,
	{_State146, CallExprType}:                     _GotoState44Action,
	{_State146, AtomExprType}:                     _GotoState42Action,
	{_State146, LiteralType}:                      _GotoState49Action,
	{_State146, ImplicitStructExprType}:           _GotoState47Action,
	{_State146, AccessExprType}:                   _GotoState38Action,
	{_State146, PostfixUnaryExprType}:             _GotoState53Action,
	{_State146, PrefixUnaryOpType}:                _GotoState55Action,
	{_State146, PrefixUnaryExprType}:              _GotoState54Action,
	{_State146, MulExprType}:                      _GotoState50Action,
	{_State146, AddExprType}:                      _GotoState39Action,
	{_State146, CmpExprType}:                      _GotoState45Action,
	{_State146, AndExprType}:                      _GotoState40Action,
	{_State146, OrExprType}:                       _GotoState52Action,
	{_State146, InitializableTypeType}:            _GotoState48Action,
	{_State146, ExplicitStructDefType}:            _GotoState46Action,
	{_State146, AnonymousFuncExprType}:            _GotoState41Action,
	{_State147, ImportToken}:                      _GotoState237Action,
	{_State147, UnsafeToken}:                      _GotoState167Action,
	{_State147, RparenToken}:                      _GotoState238Action,
	{_State147, UnsafeStatementType}:              _GotoState242Action,
	{_State147, PackageStatementBodyType}:         _GotoState241Action,
	{_State147, PackageStatementType}:             _GotoState240Action,
	{_State147, ImportStatementType}:              _GotoState239Action,
	{_State148, IdentifierToken}:                  _GotoState80Action,
	{_State148, StructToken}:                      _GotoState34Action,
	{_State148, EnumToken}:                        _GotoState77Action,
	{_State148, TraitToken}:                       _GotoState85Action,
	{_State148, FuncToken}:                        _GotoState79Action,
	{_State148, LparenToken}:                      _GotoState81Action,
	{_State148, LbracketToken}:                    _GotoState26Action,
	{_State148, DotToken}:                         _GotoState76Action,
	{_State148, QuestionToken}:                    _GotoState83Action,
	{_State148, ExclaimToken}:                     _GotoState78Action,
	{_State148, TildeTildeToken}:                  _GotoState84Action,
	{_State148, BitNegToken}:                      _GotoState75Action,
	{_State148, BitAndToken}:                      _GotoState74Action,
	{_State148, ParseErrorToken}:                  _GotoState82Action,
	{_State148, InitializableTypeType}:            _GotoState91Action,
	{_State148, AtomTypeType}:                     _GotoState86Action,
	{_State148, ReturnableTypeType}:               _GotoState92Action,
	{_State148, ValueTypeType}:                    _GotoState243Action,
	{_State148, ImplicitStructDefType}:            _GotoState90Action,
	{_State148, ExplicitStructDefType}:            _GotoState46Action,
	{_State148, ImplicitEnumDefType}:              _GotoState89Action,
	{_State148, ExplicitEnumDefType}:              _GotoState87Action,
	{_State148, TraitDefType}:                     _GotoState93Action,
	{_State148, FuncTypeType}:                     _GotoState88Action,
	{_State149, IdentifierToken}:                  _GotoState244Action,
	{_State149, GenericParameterDefType}:          _GotoState245Action,
	{_State149, GenericParameterDefsType}:         _GotoState246Action,
	{_State149, OptionalGenericParameterDefsType}: _GotoState247Action,
	{_State150, IdentifierToken}:                  _GotoState80Action,
	{_State150, StructToken}:                      _GotoState34Action,
	{_State150, EnumToken}:                        _GotoState77Action,
	{_State150, TraitToken}:                       _GotoState85Action,
	{_State150, FuncToken}:                        _GotoState79Action,
	{_State150, LparenToken}:                      _GotoState81Action,
	{_State150, LbracketToken}:                    _GotoState26Action,
	{_State150, DotToken}:                         _GotoState76Action,
	{_State150, QuestionToken}:                    _GotoState83Action,
	{_State150, ExclaimToken}:                     _GotoState78Action,
	{_State150, TildeTildeToken}:                  _GotoState84Action,
	{_State150, BitNegToken}:                      _GotoState75Action,
	{_State150, BitAndToken}:                      _GotoState74Action,
	{_State150, ParseErrorToken}:                  _GotoState82Action,
	{_State150, InitializableTypeType}:            _GotoState91Action,
	{_State150, AtomTypeType}:                     _GotoState86Action,
	{_State150, ReturnableTypeType}:               _GotoState92Action,
	{_State150, ValueTypeType}:                    _GotoState248Action,
	{_State150, ImplicitStructDefType}:            _GotoState90Action,
	{_State150, ExplicitStructDefType}:            _GotoState46Action,
	{_State150, ImplicitEnumDefType}:              _GotoState89Action,
	{_State150, ExplicitEnumDefType}:              _GotoState87Action,
	{_State150, TraitDefType}:                     _GotoState93Action,
	{_State150, FuncTypeType}:                     _GotoState88Action,
	{_State151, IntegerLiteralToken}:              _GotoState24Action,
	{_State151, FloatLiteralToken}:                _GotoState20Action,
	{_State151, RuneLiteralToken}:                 _GotoState32Action,
	{_State151, StringLiteralToken}:               _GotoState33Action,
	{_State151, IdentifierToken}:                  _GotoState23Action,
	{_State151, TrueToken}:                        _GotoState36Action,
	{_State151, FalseToken}:                       _GotoState19Action,
	{_State151, StructToken}:                      _GotoState34Action,
	{_State151, FuncToken}:                        _GotoState21Action,
	{_State151, VarToken}:                         _GotoState37Action,
	{_State151, LetToken}:                         _GotoState27Action,
	{_State151, LabelDeclToken}:                   _GotoState25Action,
	{_State151, LparenToken}:                      _GotoState28Action,
	{_State151, LbracketToken}:                    _GotoState26Action,
	{_State151, NotToken}:                         _GotoState30Action,
	{_State151, SubToken}:                         _GotoState35Action,
	{_State151, MulToken}:                         _GotoState29Action,
	{_State151, BitNegToken}:                      _GotoState18Action,
	{_State151, BitAndToken}:                      _GotoState17Action,
	{_State151, GreaterToken}:                     _GotoState22Action,
	{_State151, ParseErrorToken}:                  _GotoState31Action,
	{_State151, VarDeclPatternType}:               _GotoState57Action,
	{_State151, VarOrLetType}:                     _GotoState58Action,
	{_State151, ExpressionType}:                   _GotoState249Action,
	{_State151, OptionalLabelDeclType}:            _GotoState51Action,
	{_State151, SequenceExprType}:                 _GotoState56Action,
	{_State151, BlockExprType}:                    _GotoState43Action,
	{_State151, CallExprType}:                     _GotoState44Action,
	{_State151, AtomExprType}:                     _GotoState42Action,
	{_State151, LiteralType}:                      _GotoState49Action,
	{_State151, ImplicitStructExprType}:           _GotoState47Action,
	{_State151, AccessExprType}:                   _GotoState38Action,
	{_State151, PostfixUnaryExprType}:             _GotoState53Action,
	{_State151, PrefixUnaryOpType}:                _GotoState55Action,
	{_State151, PrefixUnaryExprType}:              _GotoState54Action,
	{_State151, MulExprType}:                      _GotoState50Action,
	{_State151, AddExprType}:                      _GotoState39Action,
	{_State151, CmpExprType}:                      _GotoState45Action,
	{_State151, AndExprType}:                      _GotoState40Action,
	{_State151, OrExprType}:                       _GotoState52Action,
	{_State151, InitializableTypeType}:            _GotoState48Action,
	{_State151, ExplicitStructDefType}:            _GotoState46Action,
	{_State151, AnonymousFuncExprType}:            _GotoState41Action,
	{_State152, LparenToken}:                      _GotoState250Action,
	{_State153, IdentifierToken}:                  _GotoState80Action,
	{_State153, StructToken}:                      _GotoState34Action,
	{_State153, EnumToken}:                        _GotoState77Action,
	{_State153, TraitToken}:                       _GotoState85Action,
	{_State153, FuncToken}:                        _GotoState79Action,
	{_State153, LparenToken}:                      _GotoState81Action,
	{_State153, LbracketToken}:                    _GotoState26Action,
	{_State153, DotToken}:                         _GotoState76Action,
	{_State153, QuestionToken}:                    _GotoState83Action,
	{_State153, ExclaimToken}:                     _GotoState78Action,
	{_State153, DotDotDotToken}:                   _GotoState251Action,
	{_State153, TildeTildeToken}:                  _GotoState84Action,
	{_State153, BitNegToken}:                      _GotoState75Action,
	{_State153, BitAndToken}:                      _GotoState74Action,
	{_State153, ParseErrorToken}:                  _GotoState82Action,
	{_State153, InitializableTypeType}:            _GotoState91Action,
	{_State153, AtomTypeType}:                     _GotoState86Action,
	{_State153, ReturnableTypeType}:               _GotoState92Action,
	{_State153, ValueTypeType}:                    _GotoState252Action,
	{_State153, ImplicitStructDefType}:            _GotoState90Action,
	{_State153, ExplicitStructDefType}:            _GotoState46Action,
	{_State153, ImplicitEnumDefType}:              _GotoState89Action,
	{_State153, ExplicitEnumDefType}:              _GotoState87Action,
	{_State153, TraitDefType}:                     _GotoState93Action,
	{_State153, FuncTypeType}:                     _GotoState88Action,
	{_State154, RparenToken}:                      _GotoState253Action,
	{_State155, RparenToken}:                      _GotoState254Action,
	{_State157, CommaToken}:                       _GotoState255Action,
	{_State161, IdentifierToken}:                  _GotoState166Action,
	{_State161, UnsafeToken}:                      _GotoState167Action,
	{_State161, StructToken}:                      _GotoState34Action,
	{_State161, EnumToken}:                        _GotoState77Action,
	{_State161, TraitToken}:                       _GotoState85Action,
	{_State161, FuncToken}:                        _GotoState79Action,
	{_State161, LparenToken}:                      _GotoState81Action,
	{_State161, LbracketToken}:                    _GotoState26Action,
	{_State161, DotToken}:                         _GotoState76Action,
	{_State161, QuestionToken}:                    _GotoState83Action,
	{_State161, ExclaimToken}:                     _GotoState78Action,
	{_State161, TildeTildeToken}:                  _GotoState84Action,
	{_State161, BitNegToken}:                      _GotoState75Action,
	{_State161, BitAndToken}:                      _GotoState74Action,
	{_State161, ParseErrorToken}:                  _GotoState82Action,
	{_State161, UnsafeStatementType}:              _GotoState173Action,
	{_State161, InitializableTypeType}:            _GotoState91Action,
	{_State161, AtomTypeType}:                     _GotoState86Action,
	{_State161, ReturnableTypeType}:               _GotoState92Action,
	{_State161, ValueTypeType}:                    _GotoState174Action,
	{_State161, FieldDefType}:                     _GotoState258Action,
	{_State161, ImplicitStructDefType}:            _GotoState90Action,
	{_State161, ExplicitStructDefType}:            _GotoState46Action,
	{_State161, EnumValueDefType}:                 _GotoState256Action,
	{_State161, ImplicitEnumValueDefsType}:        _GotoState259Action,
	{_State161, ImplicitEnumDefType}:              _GotoState89Action,
	{_State161, ExplicitEnumValueDefsType}:        _GotoState257Action,
	{_State161, ExplicitEnumDefType}:              _GotoState87Action,
	{_State161, TraitDefType}:                     _GotoState93Action,
	{_State161, FuncTypeType}:                     _GotoState88Action,
	{_State163, IdentifierToken}:                  _GotoState261Action,
	{_State163, StructToken}:                      _GotoState34Action,
	{_State163, EnumToken}:                        _GotoState77Action,
	{_State163, TraitToken}:                       _GotoState85Action,
	{_State163, FuncToken}:                        _GotoState79Action,
	{_State163, LparenToken}:                      _GotoState81Action,
	{_State163, LbracketToken}:                    _GotoState26Action,
	{_State163, DotToken}:                         _GotoState76Action,
	{_State163, QuestionToken}:                    _GotoState83Action,
	{_State163, ExclaimToken}:                     _GotoState78Action,
	{_State163, DotDotDotToken}:                   _GotoState260Action,
	{_State163, TildeTildeToken}:                  _GotoState84Action,
	{_State163, BitNegToken}:                      _GotoState75Action,
	{_State163, BitAndToken}:                      _GotoState74Action,
	{_State163, ParseErrorToken}:                  _GotoState82Action,
	{_State163, InitializableTypeType}:            _GotoState91Action,
	{_State163, AtomTypeType}:                     _GotoState86Action,
	{_State163, ReturnableTypeType}:               _GotoState92Action,
	{_State163, ValueTypeType}:                    _GotoState265Action,
	{_State163, ImplicitStructDefType}:            _GotoState90Action,
	{_State163, ExplicitStructDefType}:            _GotoState46Action,
	{_State163, ImplicitEnumDefType}:              _GotoState89Action,
	{_State163, ExplicitEnumDefType}:              _GotoState87Action,
	{_State163, TraitDefType}:                     _GotoState93Action,
	{_State163, ParameterDeclType}:                _GotoState263Action,
	{_State163, ParameterDeclsType}:               _GotoState264Action,
	{_State163, OptionalParameterDeclsType}:       _GotoState262Action,
	{_State163, FuncTypeType}:                     _GotoState88Action,
	{_State164, IdentifierToken}:                  _GotoState266Action,
	{_State166, IdentifierToken}:                  _GotoState80Action,
	{_State166, StructToken}:                      _GotoState34Action,
	{_State166, EnumToken}:                        _GotoState77Action,
	{_State166, TraitToken}:                       _GotoState85Action,
	{_State166, FuncToken}:                        _GotoState79Action,
	{_State166, LparenToken}:                      _GotoState81Action,
	{_State166, LbracketToken}:                    _GotoState26Action,
	{_State166, DotToken}:                         _GotoState267Action,
	{_State166, QuestionToken}:                    _GotoState83Action,
	{_State166, ExclaimToken}:                     _GotoState78Action,
	{_State166, DollarLbracketToken}:              _GotoState103Action,
	{_State166, TildeTildeToken}:                  _GotoState84Action,
	{_State166, BitNegToken}:                      _GotoState75Action,
	{_State166, BitAndToken}:                      _GotoState74Action,
	{_State166, ParseErrorToken}:                  _GotoState82Action,
	{_State166, OptionalGenericBindingType}:       _GotoState165Action,
	{_State166, InitializableTypeType}:            _GotoState91Action,
	{_State166, AtomTypeType}:                     _GotoState86Action,
	{_State166, ReturnableTypeType}:               _GotoState92Action,
	{_State166, ValueTypeType}:                    _GotoState268Action,
	{_State166, ImplicitStructDefType}:            _GotoState90Action,
	{_State166, ExplicitStructDefType}:            _GotoState46Action,
	{_State166, ImplicitEnumDefType}:              _GotoState89Action,
	{_State166, ExplicitEnumDefType}:              _GotoState87Action,
	{_State166, TraitDefType}:                     _GotoState93Action,
	{_State166, FuncTypeType}:                     _GotoState88Action,
	{_State167, LessToken}:                        _GotoState269Action,
	{_State168, OrToken}:                          _GotoState270Action,
	{_State169, AssignToken}:                      _GotoState271Action,
	{_State170, RparenToken}:                      _GotoState273Action,
	{_State170, OrToken}:                          _GotoState272Action,
	{_State171, CommaToken}:                       _GotoState274Action,
	{_State172, RparenToken}:                      _GotoState275Action,
	{_State174, AddToken}:                         _GotoState178Action,
	{_State174, SubToken}:                         _GotoState183Action,
	{_State174, MulToken}:                         _GotoState181Action,
	{_State177, IdentifierToken}:                  _GotoState166Action,
	{_State177, UnsafeToken}:                      _GotoState167Action,
	{_State177, StructToken}:                      _GotoState34Action,
	{_State177, EnumToken}:                        _GotoState77Action,
	{_State177, TraitToken}:                       _GotoState85Action,
	{_State177, FuncToken}:                        _GotoState276Action,
	{_State177, LparenToken}:                      _GotoState81Action,
	{_State177, LbracketToken}:                    _GotoState26Action,
	{_State177, DotToken}:                         _GotoState76Action,
	{_State177, QuestionToken}:                    _GotoState83Action,
	{_State177, ExclaimToken}:                     _GotoState78Action,
	{_State177, TildeTildeToken}:                  _GotoState84Action,
	{_State177, BitNegToken}:                      _GotoState75Action,
	{_State177, BitAndToken}:                      _GotoState74Action,
	{_State177, ParseErrorToken}:                  _GotoState82Action,
	{_State177, UnsafeStatementType}:              _GotoState173Action,
	{_State177, InitializableTypeType}:            _GotoState91Action,
	{_State177, AtomTypeType}:                     _GotoState86Action,
	{_State177, ReturnableTypeType}:               _GotoState92Action,
	{_State177, ValueTypeType}:                    _GotoState174Action,
	{_State177, FieldDefType}:                     _GotoState277Action,
	{_State177, ImplicitStructDefType}:            _GotoState90Action,
	{_State177, ExplicitStructDefType}:            _GotoState46Action,
	{_State177, ImplicitEnumDefType}:              _GotoState89Action,
	{_State177, ExplicitEnumDefType}:              _GotoState87Action,
	{_State177, TraitPropertyType}:                _GotoState281Action,
	{_State177, TraitPropertiesType}:              _GotoState280Action,
	{_State177, OptionalTraitPropertiesType}:      _GotoState279Action,
	{_State177, TraitDefType}:                     _GotoState93Action,
	{_State177, FuncTypeType}:                     _GotoState88Action,
	{_State177, MethodSignatureType}:              _GotoState278Action,
	{_State178, IdentifierToken}:                  _GotoState80Action,
	{_State178, StructToken}:                      _GotoState34Action,
	{_State178, EnumToken}:                        _GotoState77Action,
	{_State178, TraitToken}:                       _GotoState85Action,
	{_State178, FuncToken}:                        _GotoState79Action,
	{_State178, LparenToken}:                      _GotoState81Action,
	{_State178, LbracketToken}:                    _GotoState26Action,
	{_State178, DotToken}:                         _GotoState76Action,
	{_State178, QuestionToken}:                    _GotoState83Action,
	{_State178, ExclaimToken}:                     _GotoState78Action,
	{_State178, TildeTildeToken}:                  _GotoState84Action,
	{_State178, BitNegToken}:                      _GotoState75Action,
	{_State178, BitAndToken}:                      _GotoState74Action,
	{_State178, ParseErrorToken}:                  _GotoState82Action,
	{_State178, InitializableTypeType}:            _GotoState91Action,
	{_State178, AtomTypeType}:                     _GotoState86Action,
	{_State178, ReturnableTypeType}:               _GotoState282Action,
	{_State178, ImplicitStructDefType}:            _GotoState90Action,
	{_State178, ExplicitStructDefType}:            _GotoState46Action,
	{_State178, ImplicitEnumDefType}:              _GotoState89Action,
	{_State178, ExplicitEnumDefType}:              _GotoState87Action,
	{_State178, TraitDefType}:                     _GotoState93Action,
	{_State178, FuncTypeType}:                     _GotoState88Action,
	{_State179, IdentifierToken}:                  _GotoState80Action,
	{_State179, StructToken}:                      _GotoState34Action,
	{_State179, EnumToken}:                        _GotoState77Action,
	{_State179, TraitToken}:                       _GotoState85Action,
	{_State179, FuncToken}:                        _GotoState79Action,
	{_State179, LparenToken}:                      _GotoState81Action,
	{_State179, LbracketToken}:                    _GotoState26Action,
	{_State179, DotToken}:                         _GotoState76Action,
	{_State179, QuestionToken}:                    _GotoState83Action,
	{_State179, ExclaimToken}:                     _GotoState78Action,
	{_State179, TildeTildeToken}:                  _GotoState84Action,
	{_State179, BitNegToken}:                      _GotoState75Action,
	{_State179, BitAndToken}:                      _GotoState74Action,
	{_State179, ParseErrorToken}:                  _GotoState82Action,
	{_State179, InitializableTypeType}:            _GotoState91Action,
	{_State179, AtomTypeType}:                     _GotoState86Action,
	{_State179, ReturnableTypeType}:               _GotoState92Action,
	{_State179, ValueTypeType}:                    _GotoState283Action,
	{_State179, ImplicitStructDefType}:            _GotoState90Action,
	{_State179, ExplicitStructDefType}:            _GotoState46Action,
	{_State179, ImplicitEnumDefType}:              _GotoState89Action,
	{_State179, ExplicitEnumDefType}:              _GotoState87Action,
	{_State179, TraitDefType}:                     _GotoState93Action,
	{_State179, FuncTypeType}:                     _GotoState88Action,
	{_State180, IntegerLiteralToken}:              _GotoState284Action,
	{_State181, IdentifierToken}:                  _GotoState80Action,
	{_State181, StructToken}:                      _GotoState34Action,
	{_State181, EnumToken}:                        _GotoState77Action,
	{_State181, TraitToken}:                       _GotoState85Action,
	{_State181, FuncToken}:                        _GotoState79Action,
	{_State181, LparenToken}:                      _GotoState81Action,
	{_State181, LbracketToken}:                    _GotoState26Action,
	{_State181, DotToken}:                         _GotoState76Action,
	{_State181, QuestionToken}:                    _GotoState83Action,
	{_State181, ExclaimToken}:                     _GotoState78Action,
	{_State181, TildeTildeToken}:                  _GotoState84Action,
	{_State181, BitNegToken}:                      _GotoState75Action,
	{_State181, BitAndToken}:                      _GotoState74Action,
	{_State181, ParseErrorToken}:                  _GotoState82Action,
	{_State181, InitializableTypeType}:            _GotoState91Action,
	{_State181, AtomTypeType}:                     _GotoState86Action,
	{_State181, ReturnableTypeType}:               _GotoState285Action,
	{_State181, ImplicitStructDefType}:            _GotoState90Action,
	{_State181, ExplicitStructDefType}:            _GotoState46Action,
	{_State181, ImplicitEnumDefType}:              _GotoState89Action,
	{_State181, ExplicitEnumDefType}:              _GotoState87Action,
	{_State181, TraitDefType}:                     _GotoState93Action,
	{_State181, FuncTypeType}:                     _GotoState88Action,
	{_State183, IdentifierToken}:                  _GotoState80Action,
	{_State183, StructToken}:                      _GotoState34Action,
	{_State183, EnumToken}:                        _GotoState77Action,
	{_State183, TraitToken}:                       _GotoState85Action,
	{_State183, FuncToken}:                        _GotoState79Action,
	{_State183, LparenToken}:                      _GotoState81Action,
	{_State183, LbracketToken}:                    _GotoState26Action,
	{_State183, DotToken}:                         _GotoState76Action,
	{_State183, QuestionToken}:                    _GotoState83Action,
	{_State183, ExclaimToken}:                     _GotoState78Action,
	{_State183, TildeTildeToken}:                  _GotoState84Action,
	{_State183, BitNegToken}:                      _GotoState75Action,
	{_State183, BitAndToken}:                      _GotoState74Action,
	{_State183, ParseErrorToken}:                  _GotoState82Action,
	{_State183, InitializableTypeType}:            _GotoState91Action,
	{_State183, AtomTypeType}:                     _GotoState86Action,
	{_State183, ReturnableTypeType}:               _GotoState286Action,
	{_State183, ImplicitStructDefType}:            _GotoState90Action,
	{_State183, ExplicitStructDefType}:            _GotoState46Action,
	{_State183, ImplicitEnumDefType}:              _GotoState89Action,
	{_State183, ExplicitEnumDefType}:              _GotoState87Action,
	{_State183, TraitDefType}:                     _GotoState93Action,
	{_State183, FuncTypeType}:                     _GotoState88Action,
	{_State184, IntegerLiteralToken}:              _GotoState24Action,
	{_State184, FloatLiteralToken}:                _GotoState20Action,
	{_State184, RuneLiteralToken}:                 _GotoState32Action,
	{_State184, StringLiteralToken}:               _GotoState33Action,
	{_State184, IdentifierToken}:                  _GotoState23Action,
	{_State184, TrueToken}:                        _GotoState36Action,
	{_State184, FalseToken}:                       _GotoState19Action,
	{_State184, StructToken}:                      _GotoState34Action,
	{_State184, FuncToken}:                        _GotoState21Action,
	{_State184, VarToken}:                         _GotoState37Action,
	{_State184, LetToken}:                         _GotoState27Action,
	{_State184, LabelDeclToken}:                   _GotoState25Action,
	{_State184, LparenToken}:                      _GotoState28Action,
	{_State184, LbracketToken}:                    _GotoState26Action,
	{_State184, NotToken}:                         _GotoState30Action,
	{_State184, SubToken}:                         _GotoState35Action,
	{_State184, MulToken}:                         _GotoState29Action,
	{_State184, BitNegToken}:                      _GotoState18Action,
	{_State184, BitAndToken}:                      _GotoState17Action,
	{_State184, GreaterToken}:                     _GotoState22Action,
	{_State184, ParseErrorToken}:                  _GotoState31Action,
	{_State184, VarDeclPatternType}:               _GotoState57Action,
	{_State184, VarOrLetType}:                     _GotoState58Action,
	{_State184, ExpressionType}:                   _GotoState287Action,
	{_State184, OptionalLabelDeclType}:            _GotoState51Action,
	{_State184, SequenceExprType}:                 _GotoState56Action,
	{_State184, BlockExprType}:                    _GotoState43Action,
	{_State184, CallExprType}:                     _GotoState44Action,
	{_State184, AtomExprType}:                     _GotoState42Action,
	{_State184, LiteralType}:                      _GotoState49Action,
	{_State184, ImplicitStructExprType}:           _GotoState47Action,
	{_State184, AccessExprType}:                   _GotoState38Action,
	{_State184, PostfixUnaryExprType}:             _GotoState53Action,
	{_State184, PrefixUnaryOpType}:                _GotoState55Action,
	{_State184, PrefixUnaryExprType}:              _GotoState54Action,
	{_State184, MulExprType}:                      _GotoState50Action,
	{_State184, AddExprType}:                      _GotoState39Action,
	{_State184, CmpExprType}:                      _GotoState45Action,
	{_State184, AndExprType}:                      _GotoState40Action,
	{_State184, OrExprType}:                       _GotoState52Action,
	{_State184, InitializableTypeType}:            _GotoState48Action,
	{_State184, ExplicitStructDefType}:            _GotoState46Action,
	{_State184, AnonymousFuncExprType}:            _GotoState41Action,
	{_State185, IntegerLiteralToken}:              _GotoState24Action,
	{_State185, FloatLiteralToken}:                _GotoState20Action,
	{_State185, RuneLiteralToken}:                 _GotoState32Action,
	{_State185, StringLiteralToken}:               _GotoState33Action,
	{_State185, IdentifierToken}:                  _GotoState96Action,
	{_State185, TrueToken}:                        _GotoState36Action,
	{_State185, FalseToken}:                       _GotoState19Action,
	{_State185, StructToken}:                      _GotoState34Action,
	{_State185, FuncToken}:                        _GotoState21Action,
	{_State185, VarToken}:                         _GotoState37Action,
	{_State185, LetToken}:                         _GotoState27Action,
	{_State185, LabelDeclToken}:                   _GotoState25Action,
	{_State185, LparenToken}:                      _GotoState28Action,
	{_State185, LbracketToken}:                    _GotoState26Action,
	{_State185, DotDotDotToken}:                   _GotoState95Action,
	{_State185, NotToken}:                         _GotoState30Action,
	{_State185, SubToken}:                         _GotoState35Action,
	{_State185, MulToken}:                         _GotoState29Action,
	{_State185, BitNegToken}:                      _GotoState18Action,
	{_State185, BitAndToken}:                      _GotoState17Action,
	{_State185, GreaterToken}:                     _GotoState22Action,
	{_State185, ParseErrorToken}:                  _GotoState31Action,
	{_State185, VarDeclPatternType}:               _GotoState57Action,
	{_State185, VarOrLetType}:                     _GotoState58Action,
	{_State185, ExpressionType}:                   _GotoState100Action,
	{_State185, OptionalLabelDeclType}:            _GotoState51Action,
	{_State185, SequenceExprType}:                 _GotoState56Action,
	{_State185, BlockExprType}:                    _GotoState43Action,
	{_State185, CallExprType}:                     _GotoState44Action,
	{_State185, ArgumentType}:                     _GotoState288Action,
	{_State185, ColonExpressionsType}:             _GotoState99Action,
	{_State185, OptionalExpressionType}:           _GotoState101Action,
	{_State185, AtomExprType}:                     _GotoState42Action,
	{_State185, LiteralType}:                      _GotoState49Action,
	{_State185, ImplicitStructExprType}:           _GotoState47Action,
	{_State185, AccessExprType}:                   _GotoState38Action,
	{_State185, PostfixUnaryExprType}:             _GotoState53Action,
	{_State185, PrefixUnaryOpType}:                _GotoState55Action,
	{_State185, PrefixUnaryExprType}:              _GotoState54Action,
	{_State185, MulExprType}:                      _GotoState50Action,
	{_State185, AddExprType}:                      _GotoState39Action,
	{_State185, CmpExprType}:                      _GotoState45Action,
	{_State185, AndExprType}:                      _GotoState40Action,
	{_State185, OrExprType}:                       _GotoState52Action,
	{_State185, InitializableTypeType}:            _GotoState48Action,
	{_State185, ExplicitStructDefType}:            _GotoState46Action,
	{_State185, AnonymousFuncExprType}:            _GotoState41Action,
	{_State187, IntegerLiteralToken}:              _GotoState24Action,
	{_State187, FloatLiteralToken}:                _GotoState20Action,
	{_State187, RuneLiteralToken}:                 _GotoState32Action,
	{_State187, StringLiteralToken}:               _GotoState33Action,
	{_State187, IdentifierToken}:                  _GotoState23Action,
	{_State187, TrueToken}:                        _GotoState36Action,
	{_State187, FalseToken}:                       _GotoState19Action,
	{_State187, StructToken}:                      _GotoState34Action,
	{_State187, FuncToken}:                        _GotoState21Action,
	{_State187, VarToken}:                         _GotoState37Action,
	{_State187, LetToken}:                         _GotoState27Action,
	{_State187, LabelDeclToken}:                   _GotoState25Action,
	{_State187, LparenToken}:                      _GotoState28Action,
	{_State187, LbracketToken}:                    _GotoState26Action,
	{_State187, NotToken}:                         _GotoState30Action,
	{_State187, SubToken}:                         _GotoState35Action,
	{_State187, MulToken}:                         _GotoState29Action,
	{_State187, BitNegToken}:                      _GotoState18Action,
	{_State187, BitAndToken}:                      _GotoState17Action,
	{_State187, GreaterToken}:                     _GotoState22Action,
	{_State187, ParseErrorToken}:                  _GotoState31Action,
	{_State187, VarDeclPatternType}:               _GotoState57Action,
	{_State187, VarOrLetType}:                     _GotoState58Action,
	{_State187, ExpressionType}:                   _GotoState289Action,
	{_State187, OptionalLabelDeclType}:            _GotoState51Action,
	{_State187, SequenceExprType}:                 _GotoState56Action,
	{_State187, BlockExprType}:                    _GotoState43Action,
	{_State187, CallExprType}:                     _GotoState44Action,
	{_State187, OptionalExpressionType}:           _GotoState290Action,
	{_State187, AtomExprType}:                     _GotoState42Action,
	{_State187, LiteralType}:                      _GotoState49Action,
	{_State187, ImplicitStructExprType}:           _GotoState47Action,
	{_State187, AccessExprType}:                   _GotoState38Action,
	{_State187, PostfixUnaryExprType}:             _GotoState53Action,
	{_State187, PrefixUnaryOpType}:                _GotoState55Action,
	{_State187, PrefixUnaryExprType}:              _GotoState54Action,
	{_State187, MulExprType}:                      _GotoState50Action,
	{_State187, AddExprType}:                      _GotoState39Action,
	{_State187, CmpExprType}:                      _GotoState45Action,
	{_State187, AndExprType}:                      _GotoState40Action,
	{_State187, OrExprType}:                       _GotoState52Action,
	{_State187, InitializableTypeType}:            _GotoState48Action,
	{_State187, ExplicitStructDefType}:            _GotoState46Action,
	{_State187, AnonymousFuncExprType}:            _GotoState41Action,
	{_State188, IntegerLiteralToken}:              _GotoState24Action,
	{_State188, FloatLiteralToken}:                _GotoState20Action,
	{_State188, RuneLiteralToken}:                 _GotoState32Action,
	{_State188, StringLiteralToken}:               _GotoState33Action,
	{_State188, IdentifierToken}:                  _GotoState23Action,
	{_State188, TrueToken}:                        _GotoState36Action,
	{_State188, FalseToken}:                       _GotoState19Action,
	{_State188, StructToken}:                      _GotoState34Action,
	{_State188, FuncToken}:                        _GotoState21Action,
	{_State188, VarToken}:                         _GotoState37Action,
	{_State188, LetToken}:                         _GotoState27Action,
	{_State188, LabelDeclToken}:                   _GotoState25Action,
	{_State188, LparenToken}:                      _GotoState28Action,
	{_State188, LbracketToken}:                    _GotoState26Action,
	{_State188, NotToken}:                         _GotoState30Action,
	{_State188, SubToken}:                         _GotoState35Action,
	{_State188, MulToken}:                         _GotoState29Action,
	{_State188, BitNegToken}:                      _GotoState18Action,
	{_State188, BitAndToken}:                      _GotoState17Action,
	{_State188, GreaterToken}:                     _GotoState22Action,
	{_State188, ParseErrorToken}:                  _GotoState31Action,
	{_State188, VarDeclPatternType}:               _GotoState57Action,
	{_State188, VarOrLetType}:                     _GotoState58Action,
	{_State188, ExpressionType}:                   _GotoState289Action,
	{_State188, OptionalLabelDeclType}:            _GotoState51Action,
	{_State188, SequenceExprType}:                 _GotoState56Action,
	{_State188, BlockExprType}:                    _GotoState43Action,
	{_State188, CallExprType}:                     _GotoState44Action,
	{_State188, OptionalExpressionType}:           _GotoState291Action,
	{_State188, AtomExprType}:                     _GotoState42Action,
	{_State188, LiteralType}:                      _GotoState49Action,
	{_State188, ImplicitStructExprType}:           _GotoState47Action,
	{_State188, AccessExprType}:                   _GotoState38Action,
	{_State188, PostfixUnaryExprType}:             _GotoState53Action,
	{_State188, PrefixUnaryOpType}:                _GotoState55Action,
	{_State188, PrefixUnaryExprType}:              _GotoState54Action,
	{_State188, MulExprType}:                      _GotoState50Action,
	{_State188, AddExprType}:                      _GotoState39Action,
	{_State188, CmpExprType}:                      _GotoState45Action,
	{_State188, AndExprType}:                      _GotoState40Action,
	{_State188, OrExprType}:                       _GotoState52Action,
	{_State188, InitializableTypeType}:            _GotoState48Action,
	{_State188, ExplicitStructDefType}:            _GotoState46Action,
	{_State188, AnonymousFuncExprType}:            _GotoState41Action,
	{_State189, NewlinesToken}:                    _GotoState293Action,
	{_State189, CommaToken}:                       _GotoState292Action,
	{_State191, RparenToken}:                      _GotoState294Action,
	{_State192, CommaToken}:                       _GotoState295Action,
	{_State193, RbracketToken}:                    _GotoState296Action,
	{_State194, AddToken}:                         _GotoState178Action,
	{_State194, SubToken}:                         _GotoState183Action,
	{_State194, MulToken}:                         _GotoState181Action,
	{_State196, RbracketToken}:                    _GotoState297Action,
	{_State197, IntegerLiteralToken}:              _GotoState24Action,
	{_State197, FloatLiteralToken}:                _GotoState20Action,
	{_State197, RuneLiteralToken}:                 _GotoState32Action,
	{_State197, StringLiteralToken}:               _GotoState33Action,
	{_State197, IdentifierToken}:                  _GotoState96Action,
	{_State197, TrueToken}:                        _GotoState36Action,
	{_State197, FalseToken}:                       _GotoState19Action,
	{_State197, StructToken}:                      _GotoState34Action,
	{_State197, FuncToken}:                        _GotoState21Action,
	{_State197, VarToken}:                         _GotoState37Action,
	{_State197, LetToken}:                         _GotoState27Action,
	{_State197, LabelDeclToken}:                   _GotoState25Action,
	{_State197, LparenToken}:                      _GotoState28Action,
	{_State197, LbracketToken}:                    _GotoState26Action,
	{_State197, DotDotDotToken}:                   _GotoState95Action,
	{_State197, NotToken}:                         _GotoState30Action,
	{_State197, SubToken}:                         _GotoState35Action,
	{_State197, MulToken}:                         _GotoState29Action,
	{_State197, BitNegToken}:                      _GotoState18Action,
	{_State197, BitAndToken}:                      _GotoState17Action,
	{_State197, GreaterToken}:                     _GotoState22Action,
	{_State197, ParseErrorToken}:                  _GotoState31Action,
	{_State197, VarDeclPatternType}:               _GotoState57Action,
	{_State197, VarOrLetType}:                     _GotoState58Action,
	{_State197, ExpressionType}:                   _GotoState100Action,
	{_State197, OptionalLabelDeclType}:            _GotoState51Action,
	{_State197, SequenceExprType}:                 _GotoState56Action,
	{_State197, BlockExprType}:                    _GotoState43Action,
	{_State197, CallExprType}:                     _GotoState44Action,
	{_State197, OptionalArgumentsType}:            _GotoState299Action,
	{_State197, ArgumentsType}:                    _GotoState298Action,
	{_State197, ArgumentType}:                     _GotoState97Action,
	{_State197, ColonExpressionsType}:             _GotoState99Action,
	{_State197, OptionalExpressionType}:           _GotoState101Action,
	{_State197, AtomExprType}:                     _GotoState42Action,
	{_State197, LiteralType}:                      _GotoState49Action,
	{_State197, ImplicitStructExprType}:           _GotoState47Action,
	{_State197, AccessExprType}:                   _GotoState38Action,
	{_State197, PostfixUnaryExprType}:             _GotoState53Action,
	{_State197, PrefixUnaryOpType}:                _GotoState55Action,
	{_State197, PrefixUnaryExprType}:              _GotoState54Action,
	{_State197, MulExprType}:                      _GotoState50Action,
	{_State197, AddExprType}:                      _GotoState39Action,
	{_State197, CmpExprType}:                      _GotoState45Action,
	{_State197, AndExprType}:                      _GotoState40Action,
	{_State197, OrExprType}:                       _GotoState52Action,
	{_State197, InitializableTypeType}:            _GotoState48Action,
	{_State197, ExplicitStructDefType}:            _GotoState46Action,
	{_State197, AnonymousFuncExprType}:            _GotoState41Action,
	{_State198, MulToken}:                         _GotoState127Action,
	{_State198, DivToken}:                         _GotoState125Action,
	{_State198, ModToken}:                         _GotoState126Action,
	{_State198, BitAndToken}:                      _GotoState122Action,
	{_State198, BitLshiftToken}:                   _GotoState123Action,
	{_State198, BitRshiftToken}:                   _GotoState124Action,
	{_State198, MulOpType}:                        _GotoState128Action,
	{_State199, EqualToken}:                       _GotoState114Action,
	{_State199, NotEqualToken}:                    _GotoState119Action,
	{_State199, LessToken}:                        _GotoState117Action,
	{_State199, LessOrEqualToken}:                 _GotoState118Action,
	{_State199, GreaterToken}:                     _GotoState115Action,
	{_State199, GreaterOrEqualToken}:              _GotoState116Action,
	{_State199, CmpOpType}:                        _GotoState120Action,
	{_State200, AddToken}:                         _GotoState108Action,
	{_State200, SubToken}:                         _GotoState111Action,
	{_State200, BitXorToken}:                      _GotoState110Action,
	{_State200, BitOrToken}:                       _GotoState109Action,
	{_State200, AddOpType}:                        _GotoState112Action,
	{_State201, RparenToken}:                      _GotoState300Action,
	{_State201, CommaToken}:                       _GotoState185Action,
	{_State203, ForToken}:                         _GotoState301Action,
	{_State204, InToken}:                          _GotoState303Action,
	{_State204, AssignToken}:                      _GotoState302Action,
	{_State205, SemicolonToken}:                   _GotoState304Action,
	{_State206, DoToken}:                          _GotoState305Action,
	{_State207, IntegerLiteralToken}:              _GotoState24Action,
	{_State207, FloatLiteralToken}:                _GotoState20Action,
	{_State207, RuneLiteralToken}:                 _GotoState32Action,
	{_State207, StringLiteralToken}:               _GotoState33Action,
	{_State207, IdentifierToken}:                  _GotoState23Action,
	{_State207, TrueToken}:                        _GotoState36Action,
	{_State207, FalseToken}:                       _GotoState19Action,
	{_State207, StructToken}:                      _GotoState34Action,
	{_State207, FuncToken}:                        _GotoState21Action,
	{_State207, VarToken}:                         _GotoState307Action,
	{_State207, LetToken}:                         _GotoState27Action,
	{_State207, LabelDeclToken}:                   _GotoState25Action,
	{_State207, LparenToken}:                      _GotoState28Action,
	{_State207, LbracketToken}:                    _GotoState26Action,
	{_State207, DotToken}:                         _GotoState306Action,
	{_State207, NotToken}:                         _GotoState30Action,
	{_State207, SubToken}:                         _GotoState35Action,
	{_State207, MulToken}:                         _GotoState29Action,
	{_State207, BitNegToken}:                      _GotoState18Action,
	{_State207, BitAndToken}:                      _GotoState17Action,
	{_State207, GreaterToken}:                     _GotoState22Action,
	{_State207, ParseErrorToken}:                  _GotoState31Action,
	{_State207, VarDeclPatternType}:               _GotoState57Action,
	{_State207, VarOrLetType}:                     _GotoState58Action,
	{_State207, AssignPatternType}:                _GotoState308Action,
	{_State207, CasePatternType}:                  _GotoState309Action,
	{_State207, OptionalLabelDeclType}:            _GotoState72Action,
	{_State207, CasePatternsType}:                 _GotoState310Action,
	{_State207, SequenceExprType}:                 _GotoState311Action,
	{_State207, BlockExprType}:                    _GotoState43Action,
	{_State207, CallExprType}:                     _GotoState44Action,
	{_State207, AtomExprType}:                     _GotoState42Action,
	{_State207, LiteralType}:                      _GotoState49Action,
	{_State207, ImplicitStructExprType}:           _GotoState47Action,
	{_State207, AccessExprType}:                   _GotoState38Action,
	{_State207, PostfixUnaryExprType}:             _GotoState53Action,
	{_State207, PrefixUnaryOpType}:                _GotoState55Action,
	{_State207, PrefixUnaryExprType}:              _GotoState54Action,
	{_State207, MulExprType}:                      _GotoState50Action,
	{_State207, AddExprType}:                      _GotoState39Action,
	{_State207, CmpExprType}:                      _GotoState45Action,
	{_State207, AndExprType}:                      _GotoState40Action,
	{_State207, OrExprType}:                       _GotoState52Action,
	{_State207, InitializableTypeType}:            _GotoState48Action,
	{_State207, ExplicitStructDefType}:            _GotoState46Action,
	{_State207, AnonymousFuncExprType}:            _GotoState41Action,
	{_State208, LbraceToken}:                      _GotoState59Action,
	{_State208, BlockBodyType}:                    _GotoState312Action,
	{_State210, LbraceToken}:                      _GotoState313Action,
	{_State211, AndToken}:                         _GotoState113Action,
	{_State213, AssignToken}:                      _GotoState314Action,
	{_State215, RparenToken}:                      _GotoState316Action,
	{_State215, CommaToken}:                       _GotoState315Action,
	{_State218, AddToken}:                         _GotoState178Action,
	{_State218, SubToken}:                         _GotoState183Action,
	{_State218, MulToken}:                         _GotoState181Action,
	{_State219, IntegerLiteralToken}:              _GotoState24Action,
	{_State219, FloatLiteralToken}:                _GotoState20Action,
	{_State219, RuneLiteralToken}:                 _GotoState32Action,
	{_State219, StringLiteralToken}:               _GotoState33Action,
	{_State219, IdentifierToken}:                  _GotoState23Action,
	{_State219, TrueToken}:                        _GotoState36Action,
	{_State219, FalseToken}:                       _GotoState19Action,
	{_State219, StructToken}:                      _GotoState34Action,
	{_State219, FuncToken}:                        _GotoState21Action,
	{_State219, LabelDeclToken}:                   _GotoState25Action,
	{_State219, LparenToken}:                      _GotoState28Action,
	{_State219, LbracketToken}:                    _GotoState26Action,
	{_State219, ParseErrorToken}:                  _GotoState31Action,
	{_State219, OptionalLabelDeclType}:            _GotoState72Action,
	{_State219, BlockExprType}:                    _GotoState43Action,
	{_State219, CallExprType}:                     _GotoState318Action,
	{_State219, AtomExprType}:                     _GotoState42Action,
	{_State219, LiteralType}:                      _GotoState49Action,
	{_State219, ImplicitStructExprType}:           _GotoState47Action,
	{_State219, AccessExprType}:                   _GotoState317Action,
	{_State219, InitializableTypeType}:            _GotoState48Action,
	{_State219, ExplicitStructDefType}:            _GotoState46Action,
	{_State219, AnonymousFuncExprType}:            _GotoState41Action,
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
	{_State222, OptionalLabelDeclType}:            _GotoState72Action,
	{_State222, BlockExprType}:                    _GotoState43Action,
	{_State222, CallExprType}:                     _GotoState319Action,
	{_State222, AtomExprType}:                     _GotoState42Action,
	{_State222, LiteralType}:                      _GotoState49Action,
	{_State222, ImplicitStructExprType}:           _GotoState47Action,
	{_State222, AccessExprType}:                   _GotoState317Action,
	{_State222, InitializableTypeType}:            _GotoState48Action,
	{_State222, ExplicitStructDefType}:            _GotoState46Action,
	{_State222, AnonymousFuncExprType}:            _GotoState41Action,
	{_State225, LbracketToken}:                    _GotoState105Action,
	{_State225, DotToken}:                         _GotoState104Action,
	{_State225, QuestionToken}:                    _GotoState106Action,
	{_State225, DollarLbracketToken}:              _GotoState103Action,
	{_State225, AddAssignToken}:                   _GotoState320Action,
	{_State225, SubAssignToken}:                   _GotoState331Action,
	{_State225, MulAssignToken}:                   _GotoState330Action,
	{_State225, DivAssignToken}:                   _GotoState328Action,
	{_State225, ModAssignToken}:                   _GotoState329Action,
	{_State225, AddOneAssignToken}:                _GotoState321Action,
	{_State225, SubOneAssignToken}:                _GotoState332Action,
	{_State225, BitNegAssignToken}:                _GotoState324Action,
	{_State225, BitAndAssignToken}:                _GotoState322Action,
	{_State225, BitOrAssignToken}:                 _GotoState325Action,
	{_State225, BitXorAssignToken}:                _GotoState327Action,
	{_State225, BitLshiftAssignToken}:             _GotoState323Action,
	{_State225, BitRshiftAssignToken}:             _GotoState326Action,
	{_State225, UnaryOpAssignType}:                _GotoState334Action,
	{_State225, BinaryOpAssignType}:               _GotoState333Action,
	{_State225, OptionalGenericBindingType}:       _GotoState107Action,
	{_State226, AssignToken}:                      _GotoState335Action,
	{_State228, CommaToken}:                       _GotoState336Action,
	{_State230, JumpLabelToken}:                   _GotoState337Action,
	{_State230, OptionalJumpLabelType}:            _GotoState338Action,
	{_State233, NewlinesToken}:                    _GotoState339Action,
	{_State233, SemicolonToken}:                   _GotoState340Action,
	{_State237, StringLiteralToken}:               _GotoState342Action,
	{_State237, LparenToken}:                      _GotoState341Action,
	{_State237, ImportClauseType}:                 _GotoState343Action,
	{_State241, NewlinesToken}:                    _GotoState345Action,
	{_State241, CommaToken}:                       _GotoState344Action,
	{_State243, AddToken}:                         _GotoState178Action,
	{_State243, SubToken}:                         _GotoState183Action,
	{_State243, MulToken}:                         _GotoState181Action,
	{_State244, IdentifierToken}:                  _GotoState80Action,
	{_State244, StructToken}:                      _GotoState34Action,
	{_State244, EnumToken}:                        _GotoState77Action,
	{_State244, TraitToken}:                       _GotoState85Action,
	{_State244, FuncToken}:                        _GotoState79Action,
	{_State244, LparenToken}:                      _GotoState81Action,
	{_State244, LbracketToken}:                    _GotoState26Action,
	{_State244, DotToken}:                         _GotoState76Action,
	{_State244, QuestionToken}:                    _GotoState83Action,
	{_State244, ExclaimToken}:                     _GotoState78Action,
	{_State244, TildeTildeToken}:                  _GotoState84Action,
	{_State244, BitNegToken}:                      _GotoState75Action,
	{_State244, BitAndToken}:                      _GotoState74Action,
	{_State244, ParseErrorToken}:                  _GotoState82Action,
	{_State244, InitializableTypeType}:            _GotoState91Action,
	{_State244, AtomTypeType}:                     _GotoState86Action,
	{_State244, ReturnableTypeType}:               _GotoState92Action,
	{_State244, ValueTypeType}:                    _GotoState346Action,
	{_State244, ImplicitStructDefType}:            _GotoState90Action,
	{_State244, ExplicitStructDefType}:            _GotoState46Action,
	{_State244, ImplicitEnumDefType}:              _GotoState89Action,
	{_State244, ExplicitEnumDefType}:              _GotoState87Action,
	{_State244, TraitDefType}:                     _GotoState93Action,
	{_State244, FuncTypeType}:                     _GotoState88Action,
	{_State246, CommaToken}:                       _GotoState347Action,
	{_State247, RbracketToken}:                    _GotoState348Action,
	{_State248, ImplementsToken}:                  _GotoState349Action,
	{_State248, AddToken}:                         _GotoState178Action,
	{_State248, SubToken}:                         _GotoState183Action,
	{_State248, MulToken}:                         _GotoState181Action,
	{_State250, IdentifierToken}:                  _GotoState153Action,
	{_State250, ParameterDefType}:                 _GotoState156Action,
	{_State250, ParameterDefsType}:                _GotoState157Action,
	{_State250, OptionalParameterDefsType}:        _GotoState350Action,
	{_State251, IdentifierToken}:                  _GotoState80Action,
	{_State251, StructToken}:                      _GotoState34Action,
	{_State251, EnumToken}:                        _GotoState77Action,
	{_State251, TraitToken}:                       _GotoState85Action,
	{_State251, FuncToken}:                        _GotoState79Action,
	{_State251, LparenToken}:                      _GotoState81Action,
	{_State251, LbracketToken}:                    _GotoState26Action,
	{_State251, DotToken}:                         _GotoState76Action,
	{_State251, QuestionToken}:                    _GotoState83Action,
	{_State251, ExclaimToken}:                     _GotoState78Action,
	{_State251, TildeTildeToken}:                  _GotoState84Action,
	{_State251, BitNegToken}:                      _GotoState75Action,
	{_State251, BitAndToken}:                      _GotoState74Action,
	{_State251, ParseErrorToken}:                  _GotoState82Action,
	{_State251, InitializableTypeType}:            _GotoState91Action,
	{_State251, AtomTypeType}:                     _GotoState86Action,
	{_State251, ReturnableTypeType}:               _GotoState92Action,
	{_State251, ValueTypeType}:                    _GotoState351Action,
	{_State251, ImplicitStructDefType}:            _GotoState90Action,
	{_State251, ExplicitStructDefType}:            _GotoState46Action,
	{_State251, ImplicitEnumDefType}:              _GotoState89Action,
	{_State251, ExplicitEnumDefType}:              _GotoState87Action,
	{_State251, TraitDefType}:                     _GotoState93Action,
	{_State251, FuncTypeType}:                     _GotoState88Action,
	{_State252, AddToken}:                         _GotoState178Action,
	{_State252, SubToken}:                         _GotoState183Action,
	{_State252, MulToken}:                         _GotoState181Action,
	{_State253, IdentifierToken}:                  _GotoState352Action,
	{_State254, IdentifierToken}:                  _GotoState80Action,
	{_State254, StructToken}:                      _GotoState34Action,
	{_State254, EnumToken}:                        _GotoState77Action,
	{_State254, TraitToken}:                       _GotoState85Action,
	{_State254, FuncToken}:                        _GotoState79Action,
	{_State254, LparenToken}:                      _GotoState81Action,
	{_State254, LbracketToken}:                    _GotoState26Action,
	{_State254, DotToken}:                         _GotoState76Action,
	{_State254, QuestionToken}:                    _GotoState83Action,
	{_State254, ExclaimToken}:                     _GotoState78Action,
	{_State254, TildeTildeToken}:                  _GotoState84Action,
	{_State254, BitNegToken}:                      _GotoState75Action,
	{_State254, BitAndToken}:                      _GotoState74Action,
	{_State254, ParseErrorToken}:                  _GotoState82Action,
	{_State254, InitializableTypeType}:            _GotoState91Action,
	{_State254, AtomTypeType}:                     _GotoState86Action,
	{_State254, ReturnableTypeType}:               _GotoState354Action,
	{_State254, ImplicitStructDefType}:            _GotoState90Action,
	{_State254, ExplicitStructDefType}:            _GotoState46Action,
	{_State254, ImplicitEnumDefType}:              _GotoState89Action,
	{_State254, ExplicitEnumDefType}:              _GotoState87Action,
	{_State254, TraitDefType}:                     _GotoState93Action,
	{_State254, ReturnTypeType}:                   _GotoState353Action,
	{_State254, FuncTypeType}:                     _GotoState88Action,
	{_State255, IdentifierToken}:                  _GotoState153Action,
	{_State255, ParameterDefType}:                 _GotoState355Action,
	{_State256, NewlinesToken}:                    _GotoState356Action,
	{_State256, OrToken}:                          _GotoState357Action,
	{_State257, RparenToken}:                      _GotoState358Action,
	{_State258, AssignToken}:                      _GotoState271Action,
	{_State259, NewlinesToken}:                    _GotoState359Action,
	{_State259, OrToken}:                          _GotoState360Action,
	{_State260, IdentifierToken}:                  _GotoState80Action,
	{_State260, StructToken}:                      _GotoState34Action,
	{_State260, EnumToken}:                        _GotoState77Action,
	{_State260, TraitToken}:                       _GotoState85Action,
	{_State260, FuncToken}:                        _GotoState79Action,
	{_State260, LparenToken}:                      _GotoState81Action,
	{_State260, LbracketToken}:                    _GotoState26Action,
	{_State260, DotToken}:                         _GotoState76Action,
	{_State260, QuestionToken}:                    _GotoState83Action,
	{_State260, ExclaimToken}:                     _GotoState78Action,
	{_State260, TildeTildeToken}:                  _GotoState84Action,
	{_State260, BitNegToken}:                      _GotoState75Action,
	{_State260, BitAndToken}:                      _GotoState74Action,
	{_State260, ParseErrorToken}:                  _GotoState82Action,
	{_State260, InitializableTypeType}:            _GotoState91Action,
	{_State260, AtomTypeType}:                     _GotoState86Action,
	{_State260, ReturnableTypeType}:               _GotoState92Action,
	{_State260, ValueTypeType}:                    _GotoState361Action,
	{_State260, ImplicitStructDefType}:            _GotoState90Action,
	{_State260, ExplicitStructDefType}:            _GotoState46Action,
	{_State260, ImplicitEnumDefType}:              _GotoState89Action,
	{_State260, ExplicitEnumDefType}:              _GotoState87Action,
	{_State260, TraitDefType}:                     _GotoState93Action,
	{_State260, FuncTypeType}:                     _GotoState88Action,
	{_State261, IdentifierToken}:                  _GotoState80Action,
	{_State261, StructToken}:                      _GotoState34Action,
	{_State261, EnumToken}:                        _GotoState77Action,
	{_State261, TraitToken}:                       _GotoState85Action,
	{_State261, FuncToken}:                        _GotoState79Action,
	{_State261, LparenToken}:                      _GotoState81Action,
	{_State261, LbracketToken}:                    _GotoState26Action,
	{_State261, DotToken}:                         _GotoState267Action,
	{_State261, QuestionToken}:                    _GotoState83Action,
	{_State261, ExclaimToken}:                     _GotoState78Action,
	{_State261, DollarLbracketToken}:              _GotoState103Action,
	{_State261, DotDotDotToken}:                   _GotoState362Action,
	{_State261, TildeTildeToken}:                  _GotoState84Action,
	{_State261, BitNegToken}:                      _GotoState75Action,
	{_State261, BitAndToken}:                      _GotoState74Action,
	{_State261, ParseErrorToken}:                  _GotoState82Action,
	{_State261, OptionalGenericBindingType}:       _GotoState165Action,
	{_State261, InitializableTypeType}:            _GotoState91Action,
	{_State261, AtomTypeType}:                     _GotoState86Action,
	{_State261, ReturnableTypeType}:               _GotoState92Action,
	{_State261, ValueTypeType}:                    _GotoState363Action,
	{_State261, ImplicitStructDefType}:            _GotoState90Action,
	{_State261, ExplicitStructDefType}:            _GotoState46Action,
	{_State261, ImplicitEnumDefType}:              _GotoState89Action,
	{_State261, ExplicitEnumDefType}:              _GotoState87Action,
	{_State261, TraitDefType}:                     _GotoState93Action,
	{_State261, FuncTypeType}:                     _GotoState88Action,
	{_State262, RparenToken}:                      _GotoState364Action,
	{_State264, CommaToken}:                       _GotoState365Action,
	{_State265, AddToken}:                         _GotoState178Action,
	{_State265, SubToken}:                         _GotoState183Action,
	{_State265, MulToken}:                         _GotoState181Action,
	{_State266, DollarLbracketToken}:              _GotoState103Action,
	{_State266, OptionalGenericBindingType}:       _GotoState366Action,
	{_State267, IdentifierToken}:                  _GotoState266Action,
	{_State267, DollarLbracketToken}:              _GotoState103Action,
	{_State267, OptionalGenericBindingType}:       _GotoState160Action,
	{_State268, AddToken}:                         _GotoState178Action,
	{_State268, SubToken}:                         _GotoState183Action,
	{_State268, MulToken}:                         _GotoState181Action,
	{_State269, IdentifierToken}:                  _GotoState367Action,
	{_State270, IdentifierToken}:                  _GotoState166Action,
	{_State270, UnsafeToken}:                      _GotoState167Action,
	{_State270, StructToken}:                      _GotoState34Action,
	{_State270, EnumToken}:                        _GotoState77Action,
	{_State270, TraitToken}:                       _GotoState85Action,
	{_State270, FuncToken}:                        _GotoState79Action,
	{_State270, LparenToken}:                      _GotoState81Action,
	{_State270, LbracketToken}:                    _GotoState26Action,
	{_State270, DotToken}:                         _GotoState76Action,
	{_State270, QuestionToken}:                    _GotoState83Action,
	{_State270, ExclaimToken}:                     _GotoState78Action,
	{_State270, TildeTildeToken}:                  _GotoState84Action,
	{_State270, BitNegToken}:                      _GotoState75Action,
	{_State270, BitAndToken}:                      _GotoState74Action,
	{_State270, ParseErrorToken}:                  _GotoState82Action,
	{_State270, UnsafeStatementType}:              _GotoState173Action,
	{_State270, InitializableTypeType}:            _GotoState91Action,
	{_State270, AtomTypeType}:                     _GotoState86Action,
	{_State270, ReturnableTypeType}:               _GotoState92Action,
	{_State270, ValueTypeType}:                    _GotoState174Action,
	{_State270, FieldDefType}:                     _GotoState258Action,
	{_State270, ImplicitStructDefType}:            _GotoState90Action,
	{_State270, ExplicitStructDefType}:            _GotoState46Action,
	{_State270, EnumValueDefType}:                 _GotoState368Action,
	{_State270, ImplicitEnumDefType}:              _GotoState89Action,
	{_State270, ExplicitEnumDefType}:              _GotoState87Action,
	{_State270, TraitDefType}:                     _GotoState93Action,
	{_State270, FuncTypeType}:                     _GotoState88Action,
	{_State271, DefaultToken}:                     _GotoState369Action,
	{_State272, IdentifierToken}:                  _GotoState166Action,
	{_State272, UnsafeToken}:                      _GotoState167Action,
	{_State272, StructToken}:                      _GotoState34Action,
	{_State272, EnumToken}:                        _GotoState77Action,
	{_State272, TraitToken}:                       _GotoState85Action,
	{_State272, FuncToken}:                        _GotoState79Action,
	{_State272, LparenToken}:                      _GotoState81Action,
	{_State272, LbracketToken}:                    _GotoState26Action,
	{_State272, DotToken}:                         _GotoState76Action,
	{_State272, QuestionToken}:                    _GotoState83Action,
	{_State272, ExclaimToken}:                     _GotoState78Action,
	{_State272, TildeTildeToken}:                  _GotoState84Action,
	{_State272, BitNegToken}:                      _GotoState75Action,
	{_State272, BitAndToken}:                      _GotoState74Action,
	{_State272, ParseErrorToken}:                  _GotoState82Action,
	{_State272, UnsafeStatementType}:              _GotoState173Action,
	{_State272, InitializableTypeType}:            _GotoState91Action,
	{_State272, AtomTypeType}:                     _GotoState86Action,
	{_State272, ReturnableTypeType}:               _GotoState92Action,
	{_State272, ValueTypeType}:                    _GotoState174Action,
	{_State272, FieldDefType}:                     _GotoState258Action,
	{_State272, ImplicitStructDefType}:            _GotoState90Action,
	{_State272, ExplicitStructDefType}:            _GotoState46Action,
	{_State272, EnumValueDefType}:                 _GotoState370Action,
	{_State272, ImplicitEnumDefType}:              _GotoState89Action,
	{_State272, ExplicitEnumDefType}:              _GotoState87Action,
	{_State272, TraitDefType}:                     _GotoState93Action,
	{_State272, FuncTypeType}:                     _GotoState88Action,
	{_State274, IdentifierToken}:                  _GotoState166Action,
	{_State274, UnsafeToken}:                      _GotoState167Action,
	{_State274, StructToken}:                      _GotoState34Action,
	{_State274, EnumToken}:                        _GotoState77Action,
	{_State274, TraitToken}:                       _GotoState85Action,
	{_State274, FuncToken}:                        _GotoState79Action,
	{_State274, LparenToken}:                      _GotoState81Action,
	{_State274, LbracketToken}:                    _GotoState26Action,
	{_State274, DotToken}:                         _GotoState76Action,
	{_State274, QuestionToken}:                    _GotoState83Action,
	{_State274, ExclaimToken}:                     _GotoState78Action,
	{_State274, TildeTildeToken}:                  _GotoState84Action,
	{_State274, BitNegToken}:                      _GotoState75Action,
	{_State274, BitAndToken}:                      _GotoState74Action,
	{_State274, ParseErrorToken}:                  _GotoState82Action,
	{_State274, UnsafeStatementType}:              _GotoState173Action,
	{_State274, InitializableTypeType}:            _GotoState91Action,
	{_State274, AtomTypeType}:                     _GotoState86Action,
	{_State274, ReturnableTypeType}:               _GotoState92Action,
	{_State274, ValueTypeType}:                    _GotoState174Action,
	{_State274, FieldDefType}:                     _GotoState371Action,
	{_State274, ImplicitStructDefType}:            _GotoState90Action,
	{_State274, ExplicitStructDefType}:            _GotoState46Action,
	{_State274, ImplicitEnumDefType}:              _GotoState89Action,
	{_State274, ExplicitEnumDefType}:              _GotoState87Action,
	{_State274, TraitDefType}:                     _GotoState93Action,
	{_State274, FuncTypeType}:                     _GotoState88Action,
	{_State276, IdentifierToken}:                  _GotoState372Action,
	{_State276, LparenToken}:                      _GotoState163Action,
	{_State279, RparenToken}:                      _GotoState373Action,
	{_State280, NewlinesToken}:                    _GotoState375Action,
	{_State280, CommaToken}:                       _GotoState374Action,
	{_State283, RbracketToken}:                    _GotoState376Action,
	{_State283, AddToken}:                         _GotoState178Action,
	{_State283, SubToken}:                         _GotoState183Action,
	{_State283, MulToken}:                         _GotoState181Action,
	{_State284, RbracketToken}:                    _GotoState377Action,
	{_State292, IdentifierToken}:                  _GotoState166Action,
	{_State292, UnsafeToken}:                      _GotoState167Action,
	{_State292, StructToken}:                      _GotoState34Action,
	{_State292, EnumToken}:                        _GotoState77Action,
	{_State292, TraitToken}:                       _GotoState85Action,
	{_State292, FuncToken}:                        _GotoState79Action,
	{_State292, LparenToken}:                      _GotoState81Action,
	{_State292, LbracketToken}:                    _GotoState26Action,
	{_State292, DotToken}:                         _GotoState76Action,
	{_State292, QuestionToken}:                    _GotoState83Action,
	{_State292, ExclaimToken}:                     _GotoState78Action,
	{_State292, TildeTildeToken}:                  _GotoState84Action,
	{_State292, BitNegToken}:                      _GotoState75Action,
	{_State292, BitAndToken}:                      _GotoState74Action,
	{_State292, ParseErrorToken}:                  _GotoState82Action,
	{_State292, UnsafeStatementType}:              _GotoState173Action,
	{_State292, InitializableTypeType}:            _GotoState91Action,
	{_State292, AtomTypeType}:                     _GotoState86Action,
	{_State292, ReturnableTypeType}:               _GotoState92Action,
	{_State292, ValueTypeType}:                    _GotoState174Action,
	{_State292, FieldDefType}:                     _GotoState378Action,
	{_State292, ImplicitStructDefType}:            _GotoState90Action,
	{_State292, ExplicitStructDefType}:            _GotoState46Action,
	{_State292, ImplicitEnumDefType}:              _GotoState89Action,
	{_State292, ExplicitEnumDefType}:              _GotoState87Action,
	{_State292, TraitDefType}:                     _GotoState93Action,
	{_State292, FuncTypeType}:                     _GotoState88Action,
	{_State293, IdentifierToken}:                  _GotoState166Action,
	{_State293, UnsafeToken}:                      _GotoState167Action,
	{_State293, StructToken}:                      _GotoState34Action,
	{_State293, EnumToken}:                        _GotoState77Action,
	{_State293, TraitToken}:                       _GotoState85Action,
	{_State293, FuncToken}:                        _GotoState79Action,
	{_State293, LparenToken}:                      _GotoState81Action,
	{_State293, LbracketToken}:                    _GotoState26Action,
	{_State293, DotToken}:                         _GotoState76Action,
	{_State293, QuestionToken}:                    _GotoState83Action,
	{_State293, ExclaimToken}:                     _GotoState78Action,
	{_State293, TildeTildeToken}:                  _GotoState84Action,
	{_State293, BitNegToken}:                      _GotoState75Action,
	{_State293, BitAndToken}:                      _GotoState74Action,
	{_State293, ParseErrorToken}:                  _GotoState82Action,
	{_State293, UnsafeStatementType}:              _GotoState173Action,
	{_State293, InitializableTypeType}:            _GotoState91Action,
	{_State293, AtomTypeType}:                     _GotoState86Action,
	{_State293, ReturnableTypeType}:               _GotoState92Action,
	{_State293, ValueTypeType}:                    _GotoState174Action,
	{_State293, FieldDefType}:                     _GotoState379Action,
	{_State293, ImplicitStructDefType}:            _GotoState90Action,
	{_State293, ExplicitStructDefType}:            _GotoState46Action,
	{_State293, ImplicitEnumDefType}:              _GotoState89Action,
	{_State293, ExplicitEnumDefType}:              _GotoState87Action,
	{_State293, TraitDefType}:                     _GotoState93Action,
	{_State293, FuncTypeType}:                     _GotoState88Action,
	{_State295, IdentifierToken}:                  _GotoState80Action,
	{_State295, StructToken}:                      _GotoState34Action,
	{_State295, EnumToken}:                        _GotoState77Action,
	{_State295, TraitToken}:                       _GotoState85Action,
	{_State295, FuncToken}:                        _GotoState79Action,
	{_State295, LparenToken}:                      _GotoState81Action,
	{_State295, LbracketToken}:                    _GotoState26Action,
	{_State295, DotToken}:                         _GotoState76Action,
	{_State295, QuestionToken}:                    _GotoState83Action,
	{_State295, ExclaimToken}:                     _GotoState78Action,
	{_State295, TildeTildeToken}:                  _GotoState84Action,
	{_State295, BitNegToken}:                      _GotoState75Action,
	{_State295, BitAndToken}:                      _GotoState74Action,
	{_State295, ParseErrorToken}:                  _GotoState82Action,
	{_State295, InitializableTypeType}:            _GotoState91Action,
	{_State295, AtomTypeType}:                     _GotoState86Action,
	{_State295, ReturnableTypeType}:               _GotoState92Action,
	{_State295, ValueTypeType}:                    _GotoState380Action,
	{_State295, ImplicitStructDefType}:            _GotoState90Action,
	{_State295, ExplicitStructDefType}:            _GotoState46Action,
	{_State295, ImplicitEnumDefType}:              _GotoState89Action,
	{_State295, ExplicitEnumDefType}:              _GotoState87Action,
	{_State295, TraitDefType}:                     _GotoState93Action,
	{_State295, FuncTypeType}:                     _GotoState88Action,
	{_State298, CommaToken}:                       _GotoState185Action,
	{_State299, RparenToken}:                      _GotoState381Action,
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
	{_State301, LabelDeclToken}:                   _GotoState25Action,
	{_State301, LparenToken}:                      _GotoState28Action,
	{_State301, LbracketToken}:                    _GotoState26Action,
	{_State301, NotToken}:                         _GotoState30Action,
	{_State301, SubToken}:                         _GotoState35Action,
	{_State301, MulToken}:                         _GotoState29Action,
	{_State301, BitNegToken}:                      _GotoState18Action,
	{_State301, BitAndToken}:                      _GotoState17Action,
	{_State301, GreaterToken}:                     _GotoState22Action,
	{_State301, ParseErrorToken}:                  _GotoState31Action,
	{_State301, VarDeclPatternType}:               _GotoState57Action,
	{_State301, VarOrLetType}:                     _GotoState58Action,
	{_State301, OptionalLabelDeclType}:            _GotoState72Action,
	{_State301, SequenceExprType}:                 _GotoState382Action,
	{_State301, BlockExprType}:                    _GotoState43Action,
	{_State301, CallExprType}:                     _GotoState44Action,
	{_State301, AtomExprType}:                     _GotoState42Action,
	{_State301, LiteralType}:                      _GotoState49Action,
	{_State301, ImplicitStructExprType}:           _GotoState47Action,
	{_State301, AccessExprType}:                   _GotoState38Action,
	{_State301, PostfixUnaryExprType}:             _GotoState53Action,
	{_State301, PrefixUnaryOpType}:                _GotoState55Action,
	{_State301, PrefixUnaryExprType}:              _GotoState54Action,
	{_State301, MulExprType}:                      _GotoState50Action,
	{_State301, AddExprType}:                      _GotoState39Action,
	{_State301, CmpExprType}:                      _GotoState45Action,
	{_State301, AndExprType}:                      _GotoState40Action,
	{_State301, OrExprType}:                       _GotoState52Action,
	{_State301, InitializableTypeType}:            _GotoState48Action,
	{_State301, ExplicitStructDefType}:            _GotoState46Action,
	{_State301, AnonymousFuncExprType}:            _GotoState41Action,
	{_State302, IntegerLiteralToken}:              _GotoState24Action,
	{_State302, FloatLiteralToken}:                _GotoState20Action,
	{_State302, RuneLiteralToken}:                 _GotoState32Action,
	{_State302, StringLiteralToken}:               _GotoState33Action,
	{_State302, IdentifierToken}:                  _GotoState23Action,
	{_State302, TrueToken}:                        _GotoState36Action,
	{_State302, FalseToken}:                       _GotoState19Action,
	{_State302, StructToken}:                      _GotoState34Action,
	{_State302, FuncToken}:                        _GotoState21Action,
	{_State302, VarToken}:                         _GotoState37Action,
	{_State302, LetToken}:                         _GotoState27Action,
	{_State302, LabelDeclToken}:                   _GotoState25Action,
	{_State302, LparenToken}:                      _GotoState28Action,
	{_State302, LbracketToken}:                    _GotoState26Action,
	{_State302, NotToken}:                         _GotoState30Action,
	{_State302, SubToken}:                         _GotoState35Action,
	{_State302, MulToken}:                         _GotoState29Action,
	{_State302, BitNegToken}:                      _GotoState18Action,
	{_State302, BitAndToken}:                      _GotoState17Action,
	{_State302, GreaterToken}:                     _GotoState22Action,
	{_State302, ParseErrorToken}:                  _GotoState31Action,
	{_State302, VarDeclPatternType}:               _GotoState57Action,
	{_State302, VarOrLetType}:                     _GotoState58Action,
	{_State302, OptionalLabelDeclType}:            _GotoState72Action,
	{_State302, SequenceExprType}:                 _GotoState383Action,
	{_State302, BlockExprType}:                    _GotoState43Action,
	{_State302, CallExprType}:                     _GotoState44Action,
	{_State302, AtomExprType}:                     _GotoState42Action,
	{_State302, LiteralType}:                      _GotoState49Action,
	{_State302, ImplicitStructExprType}:           _GotoState47Action,
	{_State302, AccessExprType}:                   _GotoState38Action,
	{_State302, PostfixUnaryExprType}:             _GotoState53Action,
	{_State302, PrefixUnaryOpType}:                _GotoState55Action,
	{_State302, PrefixUnaryExprType}:              _GotoState54Action,
	{_State302, MulExprType}:                      _GotoState50Action,
	{_State302, AddExprType}:                      _GotoState39Action,
	{_State302, CmpExprType}:                      _GotoState45Action,
	{_State302, AndExprType}:                      _GotoState40Action,
	{_State302, OrExprType}:                       _GotoState52Action,
	{_State302, InitializableTypeType}:            _GotoState48Action,
	{_State302, ExplicitStructDefType}:            _GotoState46Action,
	{_State302, AnonymousFuncExprType}:            _GotoState41Action,
	{_State303, IntegerLiteralToken}:              _GotoState24Action,
	{_State303, FloatLiteralToken}:                _GotoState20Action,
	{_State303, RuneLiteralToken}:                 _GotoState32Action,
	{_State303, StringLiteralToken}:               _GotoState33Action,
	{_State303, IdentifierToken}:                  _GotoState23Action,
	{_State303, TrueToken}:                        _GotoState36Action,
	{_State303, FalseToken}:                       _GotoState19Action,
	{_State303, StructToken}:                      _GotoState34Action,
	{_State303, FuncToken}:                        _GotoState21Action,
	{_State303, VarToken}:                         _GotoState37Action,
	{_State303, LetToken}:                         _GotoState27Action,
	{_State303, LabelDeclToken}:                   _GotoState25Action,
	{_State303, LparenToken}:                      _GotoState28Action,
	{_State303, LbracketToken}:                    _GotoState26Action,
	{_State303, NotToken}:                         _GotoState30Action,
	{_State303, SubToken}:                         _GotoState35Action,
	{_State303, MulToken}:                         _GotoState29Action,
	{_State303, BitNegToken}:                      _GotoState18Action,
	{_State303, BitAndToken}:                      _GotoState17Action,
	{_State303, GreaterToken}:                     _GotoState22Action,
	{_State303, ParseErrorToken}:                  _GotoState31Action,
	{_State303, VarDeclPatternType}:               _GotoState57Action,
	{_State303, VarOrLetType}:                     _GotoState58Action,
	{_State303, OptionalLabelDeclType}:            _GotoState72Action,
	{_State303, SequenceExprType}:                 _GotoState384Action,
	{_State303, BlockExprType}:                    _GotoState43Action,
	{_State303, CallExprType}:                     _GotoState44Action,
	{_State303, AtomExprType}:                     _GotoState42Action,
	{_State303, LiteralType}:                      _GotoState49Action,
	{_State303, ImplicitStructExprType}:           _GotoState47Action,
	{_State303, AccessExprType}:                   _GotoState38Action,
	{_State303, PostfixUnaryExprType}:             _GotoState53Action,
	{_State303, PrefixUnaryOpType}:                _GotoState55Action,
	{_State303, PrefixUnaryExprType}:              _GotoState54Action,
	{_State303, MulExprType}:                      _GotoState50Action,
	{_State303, AddExprType}:                      _GotoState39Action,
	{_State303, CmpExprType}:                      _GotoState45Action,
	{_State303, AndExprType}:                      _GotoState40Action,
	{_State303, OrExprType}:                       _GotoState52Action,
	{_State303, InitializableTypeType}:            _GotoState48Action,
	{_State303, ExplicitStructDefType}:            _GotoState46Action,
	{_State303, AnonymousFuncExprType}:            _GotoState41Action,
	{_State304, IntegerLiteralToken}:              _GotoState24Action,
	{_State304, FloatLiteralToken}:                _GotoState20Action,
	{_State304, RuneLiteralToken}:                 _GotoState32Action,
	{_State304, StringLiteralToken}:               _GotoState33Action,
	{_State304, IdentifierToken}:                  _GotoState23Action,
	{_State304, TrueToken}:                        _GotoState36Action,
	{_State304, FalseToken}:                       _GotoState19Action,
	{_State304, StructToken}:                      _GotoState34Action,
	{_State304, FuncToken}:                        _GotoState21Action,
	{_State304, VarToken}:                         _GotoState37Action,
	{_State304, LetToken}:                         _GotoState27Action,
	{_State304, LabelDeclToken}:                   _GotoState25Action,
	{_State304, LparenToken}:                      _GotoState28Action,
	{_State304, LbracketToken}:                    _GotoState26Action,
	{_State304, NotToken}:                         _GotoState30Action,
	{_State304, SubToken}:                         _GotoState35Action,
	{_State304, MulToken}:                         _GotoState29Action,
	{_State304, BitNegToken}:                      _GotoState18Action,
	{_State304, BitAndToken}:                      _GotoState17Action,
	{_State304, GreaterToken}:                     _GotoState22Action,
	{_State304, ParseErrorToken}:                  _GotoState31Action,
	{_State304, VarDeclPatternType}:               _GotoState57Action,
	{_State304, VarOrLetType}:                     _GotoState58Action,
	{_State304, OptionalLabelDeclType}:            _GotoState72Action,
	{_State304, OptionalSequenceExprType}:         _GotoState385Action,
	{_State304, SequenceExprType}:                 _GotoState386Action,
	{_State304, BlockExprType}:                    _GotoState43Action,
	{_State304, CallExprType}:                     _GotoState44Action,
	{_State304, AtomExprType}:                     _GotoState42Action,
	{_State304, LiteralType}:                      _GotoState49Action,
	{_State304, ImplicitStructExprType}:           _GotoState47Action,
	{_State304, AccessExprType}:                   _GotoState38Action,
	{_State304, PostfixUnaryExprType}:             _GotoState53Action,
	{_State304, PrefixUnaryOpType}:                _GotoState55Action,
	{_State304, PrefixUnaryExprType}:              _GotoState54Action,
	{_State304, MulExprType}:                      _GotoState50Action,
	{_State304, AddExprType}:                      _GotoState39Action,
	{_State304, CmpExprType}:                      _GotoState45Action,
	{_State304, AndExprType}:                      _GotoState40Action,
	{_State304, OrExprType}:                       _GotoState52Action,
	{_State304, InitializableTypeType}:            _GotoState48Action,
	{_State304, ExplicitStructDefType}:            _GotoState46Action,
	{_State304, AnonymousFuncExprType}:            _GotoState41Action,
	{_State305, LbraceToken}:                      _GotoState59Action,
	{_State305, BlockBodyType}:                    _GotoState387Action,
	{_State306, IdentifierToken}:                  _GotoState388Action,
	{_State307, DotToken}:                         _GotoState389Action,
	{_State310, CommaToken}:                       _GotoState391Action,
	{_State310, AssignToken}:                      _GotoState390Action,
	{_State312, ElseToken}:                        _GotoState392Action,
	{_State313, CaseToken}:                        _GotoState393Action,
	{_State313, CaseBranchesType}:                 _GotoState395Action,
	{_State313, CaseBranchType}:                   _GotoState394Action,
	{_State314, IdentifierToken}:                  _GotoState139Action,
	{_State314, LparenToken}:                      _GotoState140Action,
	{_State314, VarPatternType}:                   _GotoState396Action,
	{_State314, TuplePatternType}:                 _GotoState141Action,
	{_State315, IdentifierToken}:                  _GotoState213Action,
	{_State315, LparenToken}:                      _GotoState140Action,
	{_State315, DotDotDotToken}:                   _GotoState212Action,
	{_State315, VarPatternType}:                   _GotoState216Action,
	{_State315, TuplePatternType}:                 _GotoState141Action,
	{_State315, FieldVarPatternType}:              _GotoState397Action,
	{_State317, LbracketToken}:                    _GotoState105Action,
	{_State317, DotToken}:                         _GotoState104Action,
	{_State317, DollarLbracketToken}:              _GotoState103Action,
	{_State317, OptionalGenericBindingType}:       _GotoState107Action,
	{_State333, IntegerLiteralToken}:              _GotoState24Action,
	{_State333, FloatLiteralToken}:                _GotoState20Action,
	{_State333, RuneLiteralToken}:                 _GotoState32Action,
	{_State333, StringLiteralToken}:               _GotoState33Action,
	{_State333, IdentifierToken}:                  _GotoState23Action,
	{_State333, TrueToken}:                        _GotoState36Action,
	{_State333, FalseToken}:                       _GotoState19Action,
	{_State333, StructToken}:                      _GotoState34Action,
	{_State333, FuncToken}:                        _GotoState21Action,
	{_State333, VarToken}:                         _GotoState37Action,
	{_State333, LetToken}:                         _GotoState27Action,
	{_State333, LabelDeclToken}:                   _GotoState25Action,
	{_State333, LparenToken}:                      _GotoState28Action,
	{_State333, LbracketToken}:                    _GotoState26Action,
	{_State333, NotToken}:                         _GotoState30Action,
	{_State333, SubToken}:                         _GotoState35Action,
	{_State333, MulToken}:                         _GotoState29Action,
	{_State333, BitNegToken}:                      _GotoState18Action,
	{_State333, BitAndToken}:                      _GotoState17Action,
	{_State333, GreaterToken}:                     _GotoState22Action,
	{_State333, ParseErrorToken}:                  _GotoState31Action,
	{_State333, VarDeclPatternType}:               _GotoState57Action,
	{_State333, VarOrLetType}:                     _GotoState58Action,
	{_State333, ExpressionType}:                   _GotoState398Action,
	{_State333, OptionalLabelDeclType}:            _GotoState51Action,
	{_State333, SequenceExprType}:                 _GotoState56Action,
	{_State333, BlockExprType}:                    _GotoState43Action,
	{_State333, CallExprType}:                     _GotoState44Action,
	{_State333, AtomExprType}:                     _GotoState42Action,
	{_State333, LiteralType}:                      _GotoState49Action,
	{_State333, ImplicitStructExprType}:           _GotoState47Action,
	{_State333, AccessExprType}:                   _GotoState38Action,
	{_State333, PostfixUnaryExprType}:             _GotoState53Action,
	{_State333, PrefixUnaryOpType}:                _GotoState55Action,
	{_State333, PrefixUnaryExprType}:              _GotoState54Action,
	{_State333, MulExprType}:                      _GotoState50Action,
	{_State333, AddExprType}:                      _GotoState39Action,
	{_State333, CmpExprType}:                      _GotoState45Action,
	{_State333, AndExprType}:                      _GotoState40Action,
	{_State333, OrExprType}:                       _GotoState52Action,
	{_State333, InitializableTypeType}:            _GotoState48Action,
	{_State333, ExplicitStructDefType}:            _GotoState46Action,
	{_State333, AnonymousFuncExprType}:            _GotoState41Action,
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
	{_State335, LabelDeclToken}:                   _GotoState25Action,
	{_State335, LparenToken}:                      _GotoState28Action,
	{_State335, LbracketToken}:                    _GotoState26Action,
	{_State335, NotToken}:                         _GotoState30Action,
	{_State335, SubToken}:                         _GotoState35Action,
	{_State335, MulToken}:                         _GotoState29Action,
	{_State335, BitNegToken}:                      _GotoState18Action,
	{_State335, BitAndToken}:                      _GotoState17Action,
	{_State335, GreaterToken}:                     _GotoState22Action,
	{_State335, ParseErrorToken}:                  _GotoState31Action,
	{_State335, VarDeclPatternType}:               _GotoState57Action,
	{_State335, VarOrLetType}:                     _GotoState58Action,
	{_State335, ExpressionType}:                   _GotoState399Action,
	{_State335, OptionalLabelDeclType}:            _GotoState51Action,
	{_State335, SequenceExprType}:                 _GotoState56Action,
	{_State335, BlockExprType}:                    _GotoState43Action,
	{_State335, CallExprType}:                     _GotoState44Action,
	{_State335, AtomExprType}:                     _GotoState42Action,
	{_State335, LiteralType}:                      _GotoState49Action,
	{_State335, ImplicitStructExprType}:           _GotoState47Action,
	{_State335, AccessExprType}:                   _GotoState38Action,
	{_State335, PostfixUnaryExprType}:             _GotoState53Action,
	{_State335, PrefixUnaryOpType}:                _GotoState55Action,
	{_State335, PrefixUnaryExprType}:              _GotoState54Action,
	{_State335, MulExprType}:                      _GotoState50Action,
	{_State335, AddExprType}:                      _GotoState39Action,
	{_State335, CmpExprType}:                      _GotoState45Action,
	{_State335, AndExprType}:                      _GotoState40Action,
	{_State335, OrExprType}:                       _GotoState52Action,
	{_State335, InitializableTypeType}:            _GotoState48Action,
	{_State335, ExplicitStructDefType}:            _GotoState46Action,
	{_State335, AnonymousFuncExprType}:            _GotoState41Action,
	{_State336, IntegerLiteralToken}:              _GotoState24Action,
	{_State336, FloatLiteralToken}:                _GotoState20Action,
	{_State336, RuneLiteralToken}:                 _GotoState32Action,
	{_State336, StringLiteralToken}:               _GotoState33Action,
	{_State336, IdentifierToken}:                  _GotoState23Action,
	{_State336, TrueToken}:                        _GotoState36Action,
	{_State336, FalseToken}:                       _GotoState19Action,
	{_State336, StructToken}:                      _GotoState34Action,
	{_State336, FuncToken}:                        _GotoState21Action,
	{_State336, VarToken}:                         _GotoState37Action,
	{_State336, LetToken}:                         _GotoState27Action,
	{_State336, LabelDeclToken}:                   _GotoState25Action,
	{_State336, LparenToken}:                      _GotoState28Action,
	{_State336, LbracketToken}:                    _GotoState26Action,
	{_State336, NotToken}:                         _GotoState30Action,
	{_State336, SubToken}:                         _GotoState35Action,
	{_State336, MulToken}:                         _GotoState29Action,
	{_State336, BitNegToken}:                      _GotoState18Action,
	{_State336, BitAndToken}:                      _GotoState17Action,
	{_State336, GreaterToken}:                     _GotoState22Action,
	{_State336, ParseErrorToken}:                  _GotoState31Action,
	{_State336, VarDeclPatternType}:               _GotoState57Action,
	{_State336, VarOrLetType}:                     _GotoState58Action,
	{_State336, ExpressionType}:                   _GotoState400Action,
	{_State336, OptionalLabelDeclType}:            _GotoState51Action,
	{_State336, SequenceExprType}:                 _GotoState56Action,
	{_State336, BlockExprType}:                    _GotoState43Action,
	{_State336, CallExprType}:                     _GotoState44Action,
	{_State336, AtomExprType}:                     _GotoState42Action,
	{_State336, LiteralType}:                      _GotoState49Action,
	{_State336, ImplicitStructExprType}:           _GotoState47Action,
	{_State336, AccessExprType}:                   _GotoState38Action,
	{_State336, PostfixUnaryExprType}:             _GotoState53Action,
	{_State336, PrefixUnaryOpType}:                _GotoState55Action,
	{_State336, PrefixUnaryExprType}:              _GotoState54Action,
	{_State336, MulExprType}:                      _GotoState50Action,
	{_State336, AddExprType}:                      _GotoState39Action,
	{_State336, CmpExprType}:                      _GotoState45Action,
	{_State336, AndExprType}:                      _GotoState40Action,
	{_State336, OrExprType}:                       _GotoState52Action,
	{_State336, InitializableTypeType}:            _GotoState48Action,
	{_State336, ExplicitStructDefType}:            _GotoState46Action,
	{_State336, AnonymousFuncExprType}:            _GotoState41Action,
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
	{_State338, LabelDeclToken}:                   _GotoState25Action,
	{_State338, LparenToken}:                      _GotoState28Action,
	{_State338, LbracketToken}:                    _GotoState26Action,
	{_State338, NotToken}:                         _GotoState30Action,
	{_State338, SubToken}:                         _GotoState35Action,
	{_State338, MulToken}:                         _GotoState29Action,
	{_State338, BitNegToken}:                      _GotoState18Action,
	{_State338, BitAndToken}:                      _GotoState17Action,
	{_State338, GreaterToken}:                     _GotoState22Action,
	{_State338, ParseErrorToken}:                  _GotoState31Action,
	{_State338, VarDeclPatternType}:               _GotoState57Action,
	{_State338, VarOrLetType}:                     _GotoState58Action,
	{_State338, ExpressionType}:                   _GotoState227Action,
	{_State338, OptionalLabelDeclType}:            _GotoState51Action,
	{_State338, SequenceExprType}:                 _GotoState56Action,
	{_State338, BlockExprType}:                    _GotoState43Action,
	{_State338, ExpressionsType}:                  _GotoState401Action,
	{_State338, OptionalExpressionsType}:          _GotoState402Action,
	{_State338, CallExprType}:                     _GotoState44Action,
	{_State338, AtomExprType}:                     _GotoState42Action,
	{_State338, LiteralType}:                      _GotoState49Action,
	{_State338, ImplicitStructExprType}:           _GotoState47Action,
	{_State338, AccessExprType}:                   _GotoState38Action,
	{_State338, PostfixUnaryExprType}:             _GotoState53Action,
	{_State338, PrefixUnaryOpType}:                _GotoState55Action,
	{_State338, PrefixUnaryExprType}:              _GotoState54Action,
	{_State338, MulExprType}:                      _GotoState50Action,
	{_State338, AddExprType}:                      _GotoState39Action,
	{_State338, CmpExprType}:                      _GotoState45Action,
	{_State338, AndExprType}:                      _GotoState40Action,
	{_State338, OrExprType}:                       _GotoState52Action,
	{_State338, InitializableTypeType}:            _GotoState48Action,
	{_State338, ExplicitStructDefType}:            _GotoState46Action,
	{_State338, AnonymousFuncExprType}:            _GotoState41Action,
	{_State341, StringLiteralToken}:               _GotoState342Action,
	{_State341, ImportClauseType}:                 _GotoState403Action,
	{_State341, ImportClauseTerminalType}:         _GotoState404Action,
	{_State341, ImportClausesType}:                _GotoState405Action,
	{_State342, AsToken}:                          _GotoState406Action,
	{_State346, AddToken}:                         _GotoState178Action,
	{_State346, SubToken}:                         _GotoState183Action,
	{_State346, MulToken}:                         _GotoState181Action,
	{_State347, IdentifierToken}:                  _GotoState244Action,
	{_State347, GenericParameterDefType}:          _GotoState407Action,
	{_State349, IdentifierToken}:                  _GotoState80Action,
	{_State349, StructToken}:                      _GotoState34Action,
	{_State349, EnumToken}:                        _GotoState77Action,
	{_State349, TraitToken}:                       _GotoState85Action,
	{_State349, FuncToken}:                        _GotoState79Action,
	{_State349, LparenToken}:                      _GotoState81Action,
	{_State349, LbracketToken}:                    _GotoState26Action,
	{_State349, DotToken}:                         _GotoState76Action,
	{_State349, QuestionToken}:                    _GotoState83Action,
	{_State349, ExclaimToken}:                     _GotoState78Action,
	{_State349, TildeTildeToken}:                  _GotoState84Action,
	{_State349, BitNegToken}:                      _GotoState75Action,
	{_State349, BitAndToken}:                      _GotoState74Action,
	{_State349, ParseErrorToken}:                  _GotoState82Action,
	{_State349, InitializableTypeType}:            _GotoState91Action,
	{_State349, AtomTypeType}:                     _GotoState86Action,
	{_State349, ReturnableTypeType}:               _GotoState92Action,
	{_State349, ValueTypeType}:                    _GotoState408Action,
	{_State349, ImplicitStructDefType}:            _GotoState90Action,
	{_State349, ExplicitStructDefType}:            _GotoState46Action,
	{_State349, ImplicitEnumDefType}:              _GotoState89Action,
	{_State349, ExplicitEnumDefType}:              _GotoState87Action,
	{_State349, TraitDefType}:                     _GotoState93Action,
	{_State349, FuncTypeType}:                     _GotoState88Action,
	{_State350, RparenToken}:                      _GotoState409Action,
	{_State351, AddToken}:                         _GotoState178Action,
	{_State351, SubToken}:                         _GotoState183Action,
	{_State351, MulToken}:                         _GotoState181Action,
	{_State352, DollarLbracketToken}:              _GotoState149Action,
	{_State352, OptionalGenericParametersType}:    _GotoState410Action,
	{_State353, LbraceToken}:                      _GotoState59Action,
	{_State353, BlockBodyType}:                    _GotoState411Action,
	{_State356, IdentifierToken}:                  _GotoState166Action,
	{_State356, UnsafeToken}:                      _GotoState167Action,
	{_State356, StructToken}:                      _GotoState34Action,
	{_State356, EnumToken}:                        _GotoState77Action,
	{_State356, TraitToken}:                       _GotoState85Action,
	{_State356, FuncToken}:                        _GotoState79Action,
	{_State356, LparenToken}:                      _GotoState81Action,
	{_State356, LbracketToken}:                    _GotoState26Action,
	{_State356, DotToken}:                         _GotoState76Action,
	{_State356, QuestionToken}:                    _GotoState83Action,
	{_State356, ExclaimToken}:                     _GotoState78Action,
	{_State356, TildeTildeToken}:                  _GotoState84Action,
	{_State356, BitNegToken}:                      _GotoState75Action,
	{_State356, BitAndToken}:                      _GotoState74Action,
	{_State356, ParseErrorToken}:                  _GotoState82Action,
	{_State356, UnsafeStatementType}:              _GotoState173Action,
	{_State356, InitializableTypeType}:            _GotoState91Action,
	{_State356, AtomTypeType}:                     _GotoState86Action,
	{_State356, ReturnableTypeType}:               _GotoState92Action,
	{_State356, ValueTypeType}:                    _GotoState174Action,
	{_State356, FieldDefType}:                     _GotoState258Action,
	{_State356, ImplicitStructDefType}:            _GotoState90Action,
	{_State356, ExplicitStructDefType}:            _GotoState46Action,
	{_State356, EnumValueDefType}:                 _GotoState412Action,
	{_State356, ImplicitEnumDefType}:              _GotoState89Action,
	{_State356, ExplicitEnumDefType}:              _GotoState87Action,
	{_State356, TraitDefType}:                     _GotoState93Action,
	{_State356, FuncTypeType}:                     _GotoState88Action,
	{_State357, IdentifierToken}:                  _GotoState166Action,
	{_State357, UnsafeToken}:                      _GotoState167Action,
	{_State357, StructToken}:                      _GotoState34Action,
	{_State357, EnumToken}:                        _GotoState77Action,
	{_State357, TraitToken}:                       _GotoState85Action,
	{_State357, FuncToken}:                        _GotoState79Action,
	{_State357, LparenToken}:                      _GotoState81Action,
	{_State357, LbracketToken}:                    _GotoState26Action,
	{_State357, DotToken}:                         _GotoState76Action,
	{_State357, QuestionToken}:                    _GotoState83Action,
	{_State357, ExclaimToken}:                     _GotoState78Action,
	{_State357, TildeTildeToken}:                  _GotoState84Action,
	{_State357, BitNegToken}:                      _GotoState75Action,
	{_State357, BitAndToken}:                      _GotoState74Action,
	{_State357, ParseErrorToken}:                  _GotoState82Action,
	{_State357, UnsafeStatementType}:              _GotoState173Action,
	{_State357, InitializableTypeType}:            _GotoState91Action,
	{_State357, AtomTypeType}:                     _GotoState86Action,
	{_State357, ReturnableTypeType}:               _GotoState92Action,
	{_State357, ValueTypeType}:                    _GotoState174Action,
	{_State357, FieldDefType}:                     _GotoState258Action,
	{_State357, ImplicitStructDefType}:            _GotoState90Action,
	{_State357, ExplicitStructDefType}:            _GotoState46Action,
	{_State357, EnumValueDefType}:                 _GotoState413Action,
	{_State357, ImplicitEnumDefType}:              _GotoState89Action,
	{_State357, ExplicitEnumDefType}:              _GotoState87Action,
	{_State357, TraitDefType}:                     _GotoState93Action,
	{_State357, FuncTypeType}:                     _GotoState88Action,
	{_State359, IdentifierToken}:                  _GotoState166Action,
	{_State359, UnsafeToken}:                      _GotoState167Action,
	{_State359, StructToken}:                      _GotoState34Action,
	{_State359, EnumToken}:                        _GotoState77Action,
	{_State359, TraitToken}:                       _GotoState85Action,
	{_State359, FuncToken}:                        _GotoState79Action,
	{_State359, LparenToken}:                      _GotoState81Action,
	{_State359, LbracketToken}:                    _GotoState26Action,
	{_State359, DotToken}:                         _GotoState76Action,
	{_State359, QuestionToken}:                    _GotoState83Action,
	{_State359, ExclaimToken}:                     _GotoState78Action,
	{_State359, TildeTildeToken}:                  _GotoState84Action,
	{_State359, BitNegToken}:                      _GotoState75Action,
	{_State359, BitAndToken}:                      _GotoState74Action,
	{_State359, ParseErrorToken}:                  _GotoState82Action,
	{_State359, UnsafeStatementType}:              _GotoState173Action,
	{_State359, InitializableTypeType}:            _GotoState91Action,
	{_State359, AtomTypeType}:                     _GotoState86Action,
	{_State359, ReturnableTypeType}:               _GotoState92Action,
	{_State359, ValueTypeType}:                    _GotoState174Action,
	{_State359, FieldDefType}:                     _GotoState258Action,
	{_State359, ImplicitStructDefType}:            _GotoState90Action,
	{_State359, ExplicitStructDefType}:            _GotoState46Action,
	{_State359, EnumValueDefType}:                 _GotoState414Action,
	{_State359, ImplicitEnumDefType}:              _GotoState89Action,
	{_State359, ExplicitEnumDefType}:              _GotoState87Action,
	{_State359, TraitDefType}:                     _GotoState93Action,
	{_State359, FuncTypeType}:                     _GotoState88Action,
	{_State360, IdentifierToken}:                  _GotoState166Action,
	{_State360, UnsafeToken}:                      _GotoState167Action,
	{_State360, StructToken}:                      _GotoState34Action,
	{_State360, EnumToken}:                        _GotoState77Action,
	{_State360, TraitToken}:                       _GotoState85Action,
	{_State360, FuncToken}:                        _GotoState79Action,
	{_State360, LparenToken}:                      _GotoState81Action,
	{_State360, LbracketToken}:                    _GotoState26Action,
	{_State360, DotToken}:                         _GotoState76Action,
	{_State360, QuestionToken}:                    _GotoState83Action,
	{_State360, ExclaimToken}:                     _GotoState78Action,
	{_State360, TildeTildeToken}:                  _GotoState84Action,
	{_State360, BitNegToken}:                      _GotoState75Action,
	{_State360, BitAndToken}:                      _GotoState74Action,
	{_State360, ParseErrorToken}:                  _GotoState82Action,
	{_State360, UnsafeStatementType}:              _GotoState173Action,
	{_State360, InitializableTypeType}:            _GotoState91Action,
	{_State360, AtomTypeType}:                     _GotoState86Action,
	{_State360, ReturnableTypeType}:               _GotoState92Action,
	{_State360, ValueTypeType}:                    _GotoState174Action,
	{_State360, FieldDefType}:                     _GotoState258Action,
	{_State360, ImplicitStructDefType}:            _GotoState90Action,
	{_State360, ExplicitStructDefType}:            _GotoState46Action,
	{_State360, EnumValueDefType}:                 _GotoState415Action,
	{_State360, ImplicitEnumDefType}:              _GotoState89Action,
	{_State360, ExplicitEnumDefType}:              _GotoState87Action,
	{_State360, TraitDefType}:                     _GotoState93Action,
	{_State360, FuncTypeType}:                     _GotoState88Action,
	{_State361, AddToken}:                         _GotoState178Action,
	{_State361, SubToken}:                         _GotoState183Action,
	{_State361, MulToken}:                         _GotoState181Action,
	{_State362, IdentifierToken}:                  _GotoState80Action,
	{_State362, StructToken}:                      _GotoState34Action,
	{_State362, EnumToken}:                        _GotoState77Action,
	{_State362, TraitToken}:                       _GotoState85Action,
	{_State362, FuncToken}:                        _GotoState79Action,
	{_State362, LparenToken}:                      _GotoState81Action,
	{_State362, LbracketToken}:                    _GotoState26Action,
	{_State362, DotToken}:                         _GotoState76Action,
	{_State362, QuestionToken}:                    _GotoState83Action,
	{_State362, ExclaimToken}:                     _GotoState78Action,
	{_State362, TildeTildeToken}:                  _GotoState84Action,
	{_State362, BitNegToken}:                      _GotoState75Action,
	{_State362, BitAndToken}:                      _GotoState74Action,
	{_State362, ParseErrorToken}:                  _GotoState82Action,
	{_State362, InitializableTypeType}:            _GotoState91Action,
	{_State362, AtomTypeType}:                     _GotoState86Action,
	{_State362, ReturnableTypeType}:               _GotoState92Action,
	{_State362, ValueTypeType}:                    _GotoState416Action,
	{_State362, ImplicitStructDefType}:            _GotoState90Action,
	{_State362, ExplicitStructDefType}:            _GotoState46Action,
	{_State362, ImplicitEnumDefType}:              _GotoState89Action,
	{_State362, ExplicitEnumDefType}:              _GotoState87Action,
	{_State362, TraitDefType}:                     _GotoState93Action,
	{_State362, FuncTypeType}:                     _GotoState88Action,
	{_State363, AddToken}:                         _GotoState178Action,
	{_State363, SubToken}:                         _GotoState183Action,
	{_State363, MulToken}:                         _GotoState181Action,
	{_State364, IdentifierToken}:                  _GotoState80Action,
	{_State364, StructToken}:                      _GotoState34Action,
	{_State364, EnumToken}:                        _GotoState77Action,
	{_State364, TraitToken}:                       _GotoState85Action,
	{_State364, FuncToken}:                        _GotoState79Action,
	{_State364, LparenToken}:                      _GotoState81Action,
	{_State364, LbracketToken}:                    _GotoState26Action,
	{_State364, DotToken}:                         _GotoState76Action,
	{_State364, QuestionToken}:                    _GotoState83Action,
	{_State364, ExclaimToken}:                     _GotoState78Action,
	{_State364, TildeTildeToken}:                  _GotoState84Action,
	{_State364, BitNegToken}:                      _GotoState75Action,
	{_State364, BitAndToken}:                      _GotoState74Action,
	{_State364, ParseErrorToken}:                  _GotoState82Action,
	{_State364, InitializableTypeType}:            _GotoState91Action,
	{_State364, AtomTypeType}:                     _GotoState86Action,
	{_State364, ReturnableTypeType}:               _GotoState354Action,
	{_State364, ImplicitStructDefType}:            _GotoState90Action,
	{_State364, ExplicitStructDefType}:            _GotoState46Action,
	{_State364, ImplicitEnumDefType}:              _GotoState89Action,
	{_State364, ExplicitEnumDefType}:              _GotoState87Action,
	{_State364, TraitDefType}:                     _GotoState93Action,
	{_State364, ReturnTypeType}:                   _GotoState417Action,
	{_State364, FuncTypeType}:                     _GotoState88Action,
	{_State365, IdentifierToken}:                  _GotoState261Action,
	{_State365, StructToken}:                      _GotoState34Action,
	{_State365, EnumToken}:                        _GotoState77Action,
	{_State365, TraitToken}:                       _GotoState85Action,
	{_State365, FuncToken}:                        _GotoState79Action,
	{_State365, LparenToken}:                      _GotoState81Action,
	{_State365, LbracketToken}:                    _GotoState26Action,
	{_State365, DotToken}:                         _GotoState76Action,
	{_State365, QuestionToken}:                    _GotoState83Action,
	{_State365, ExclaimToken}:                     _GotoState78Action,
	{_State365, DotDotDotToken}:                   _GotoState260Action,
	{_State365, TildeTildeToken}:                  _GotoState84Action,
	{_State365, BitNegToken}:                      _GotoState75Action,
	{_State365, BitAndToken}:                      _GotoState74Action,
	{_State365, ParseErrorToken}:                  _GotoState82Action,
	{_State365, InitializableTypeType}:            _GotoState91Action,
	{_State365, AtomTypeType}:                     _GotoState86Action,
	{_State365, ReturnableTypeType}:               _GotoState92Action,
	{_State365, ValueTypeType}:                    _GotoState265Action,
	{_State365, ImplicitStructDefType}:            _GotoState90Action,
	{_State365, ExplicitStructDefType}:            _GotoState46Action,
	{_State365, ImplicitEnumDefType}:              _GotoState89Action,
	{_State365, ExplicitEnumDefType}:              _GotoState87Action,
	{_State365, TraitDefType}:                     _GotoState93Action,
	{_State365, ParameterDeclType}:                _GotoState418Action,
	{_State365, FuncTypeType}:                     _GotoState88Action,
	{_State367, GreaterToken}:                     _GotoState419Action,
	{_State372, LparenToken}:                      _GotoState420Action,
	{_State374, IdentifierToken}:                  _GotoState166Action,
	{_State374, UnsafeToken}:                      _GotoState167Action,
	{_State374, StructToken}:                      _GotoState34Action,
	{_State374, EnumToken}:                        _GotoState77Action,
	{_State374, TraitToken}:                       _GotoState85Action,
	{_State374, FuncToken}:                        _GotoState276Action,
	{_State374, LparenToken}:                      _GotoState81Action,
	{_State374, LbracketToken}:                    _GotoState26Action,
	{_State374, DotToken}:                         _GotoState76Action,
	{_State374, QuestionToken}:                    _GotoState83Action,
	{_State374, ExclaimToken}:                     _GotoState78Action,
	{_State374, TildeTildeToken}:                  _GotoState84Action,
	{_State374, BitNegToken}:                      _GotoState75Action,
	{_State374, BitAndToken}:                      _GotoState74Action,
	{_State374, ParseErrorToken}:                  _GotoState82Action,
	{_State374, UnsafeStatementType}:              _GotoState173Action,
	{_State374, InitializableTypeType}:            _GotoState91Action,
	{_State374, AtomTypeType}:                     _GotoState86Action,
	{_State374, ReturnableTypeType}:               _GotoState92Action,
	{_State374, ValueTypeType}:                    _GotoState174Action,
	{_State374, FieldDefType}:                     _GotoState277Action,
	{_State374, ImplicitStructDefType}:            _GotoState90Action,
	{_State374, ExplicitStructDefType}:            _GotoState46Action,
	{_State374, ImplicitEnumDefType}:              _GotoState89Action,
	{_State374, ExplicitEnumDefType}:              _GotoState87Action,
	{_State374, TraitPropertyType}:                _GotoState421Action,
	{_State374, TraitDefType}:                     _GotoState93Action,
	{_State374, FuncTypeType}:                     _GotoState88Action,
	{_State374, MethodSignatureType}:              _GotoState278Action,
	{_State375, IdentifierToken}:                  _GotoState166Action,
	{_State375, UnsafeToken}:                      _GotoState167Action,
	{_State375, StructToken}:                      _GotoState34Action,
	{_State375, EnumToken}:                        _GotoState77Action,
	{_State375, TraitToken}:                       _GotoState85Action,
	{_State375, FuncToken}:                        _GotoState276Action,
	{_State375, LparenToken}:                      _GotoState81Action,
	{_State375, LbracketToken}:                    _GotoState26Action,
	{_State375, DotToken}:                         _GotoState76Action,
	{_State375, QuestionToken}:                    _GotoState83Action,
	{_State375, ExclaimToken}:                     _GotoState78Action,
	{_State375, TildeTildeToken}:                  _GotoState84Action,
	{_State375, BitNegToken}:                      _GotoState75Action,
	{_State375, BitAndToken}:                      _GotoState74Action,
	{_State375, ParseErrorToken}:                  _GotoState82Action,
	{_State375, UnsafeStatementType}:              _GotoState173Action,
	{_State375, InitializableTypeType}:            _GotoState91Action,
	{_State375, AtomTypeType}:                     _GotoState86Action,
	{_State375, ReturnableTypeType}:               _GotoState92Action,
	{_State375, ValueTypeType}:                    _GotoState174Action,
	{_State375, FieldDefType}:                     _GotoState277Action,
	{_State375, ImplicitStructDefType}:            _GotoState90Action,
	{_State375, ExplicitStructDefType}:            _GotoState46Action,
	{_State375, ImplicitEnumDefType}:              _GotoState89Action,
	{_State375, ExplicitEnumDefType}:              _GotoState87Action,
	{_State375, TraitPropertyType}:                _GotoState422Action,
	{_State375, TraitDefType}:                     _GotoState93Action,
	{_State375, FuncTypeType}:                     _GotoState88Action,
	{_State375, MethodSignatureType}:              _GotoState278Action,
	{_State380, AddToken}:                         _GotoState178Action,
	{_State380, SubToken}:                         _GotoState183Action,
	{_State380, MulToken}:                         _GotoState181Action,
	{_State384, DoToken}:                          _GotoState423Action,
	{_State385, SemicolonToken}:                   _GotoState424Action,
	{_State388, LparenToken}:                      _GotoState28Action,
	{_State388, ImplicitStructExprType}:           _GotoState425Action,
	{_State389, IdentifierToken}:                  _GotoState426Action,
	{_State390, IntegerLiteralToken}:              _GotoState24Action,
	{_State390, FloatLiteralToken}:                _GotoState20Action,
	{_State390, RuneLiteralToken}:                 _GotoState32Action,
	{_State390, StringLiteralToken}:               _GotoState33Action,
	{_State390, IdentifierToken}:                  _GotoState23Action,
	{_State390, TrueToken}:                        _GotoState36Action,
	{_State390, FalseToken}:                       _GotoState19Action,
	{_State390, StructToken}:                      _GotoState34Action,
	{_State390, FuncToken}:                        _GotoState21Action,
	{_State390, VarToken}:                         _GotoState37Action,
	{_State390, LetToken}:                         _GotoState27Action,
	{_State390, LabelDeclToken}:                   _GotoState25Action,
	{_State390, LparenToken}:                      _GotoState28Action,
	{_State390, LbracketToken}:                    _GotoState26Action,
	{_State390, NotToken}:                         _GotoState30Action,
	{_State390, SubToken}:                         _GotoState35Action,
	{_State390, MulToken}:                         _GotoState29Action,
	{_State390, BitNegToken}:                      _GotoState18Action,
	{_State390, BitAndToken}:                      _GotoState17Action,
	{_State390, GreaterToken}:                     _GotoState22Action,
	{_State390, ParseErrorToken}:                  _GotoState31Action,
	{_State390, VarDeclPatternType}:               _GotoState57Action,
	{_State390, VarOrLetType}:                     _GotoState58Action,
	{_State390, OptionalLabelDeclType}:            _GotoState72Action,
	{_State390, SequenceExprType}:                 _GotoState427Action,
	{_State390, BlockExprType}:                    _GotoState43Action,
	{_State390, CallExprType}:                     _GotoState44Action,
	{_State390, AtomExprType}:                     _GotoState42Action,
	{_State390, LiteralType}:                      _GotoState49Action,
	{_State390, ImplicitStructExprType}:           _GotoState47Action,
	{_State390, AccessExprType}:                   _GotoState38Action,
	{_State390, PostfixUnaryExprType}:             _GotoState53Action,
	{_State390, PrefixUnaryOpType}:                _GotoState55Action,
	{_State390, PrefixUnaryExprType}:              _GotoState54Action,
	{_State390, MulExprType}:                      _GotoState50Action,
	{_State390, AddExprType}:                      _GotoState39Action,
	{_State390, CmpExprType}:                      _GotoState45Action,
	{_State390, AndExprType}:                      _GotoState40Action,
	{_State390, OrExprType}:                       _GotoState52Action,
	{_State390, InitializableTypeType}:            _GotoState48Action,
	{_State390, ExplicitStructDefType}:            _GotoState46Action,
	{_State390, AnonymousFuncExprType}:            _GotoState41Action,
	{_State391, IntegerLiteralToken}:              _GotoState24Action,
	{_State391, FloatLiteralToken}:                _GotoState20Action,
	{_State391, RuneLiteralToken}:                 _GotoState32Action,
	{_State391, StringLiteralToken}:               _GotoState33Action,
	{_State391, IdentifierToken}:                  _GotoState23Action,
	{_State391, TrueToken}:                        _GotoState36Action,
	{_State391, FalseToken}:                       _GotoState19Action,
	{_State391, StructToken}:                      _GotoState34Action,
	{_State391, FuncToken}:                        _GotoState21Action,
	{_State391, VarToken}:                         _GotoState307Action,
	{_State391, LetToken}:                         _GotoState27Action,
	{_State391, LabelDeclToken}:                   _GotoState25Action,
	{_State391, LparenToken}:                      _GotoState28Action,
	{_State391, LbracketToken}:                    _GotoState26Action,
	{_State391, DotToken}:                         _GotoState306Action,
	{_State391, NotToken}:                         _GotoState30Action,
	{_State391, SubToken}:                         _GotoState35Action,
	{_State391, MulToken}:                         _GotoState29Action,
	{_State391, BitNegToken}:                      _GotoState18Action,
	{_State391, BitAndToken}:                      _GotoState17Action,
	{_State391, GreaterToken}:                     _GotoState22Action,
	{_State391, ParseErrorToken}:                  _GotoState31Action,
	{_State391, VarDeclPatternType}:               _GotoState57Action,
	{_State391, VarOrLetType}:                     _GotoState58Action,
	{_State391, AssignPatternType}:                _GotoState308Action,
	{_State391, CasePatternType}:                  _GotoState428Action,
	{_State391, OptionalLabelDeclType}:            _GotoState72Action,
	{_State391, SequenceExprType}:                 _GotoState311Action,
	{_State391, BlockExprType}:                    _GotoState43Action,
	{_State391, CallExprType}:                     _GotoState44Action,
	{_State391, AtomExprType}:                     _GotoState42Action,
	{_State391, LiteralType}:                      _GotoState49Action,
	{_State391, ImplicitStructExprType}:           _GotoState47Action,
	{_State391, AccessExprType}:                   _GotoState38Action,
	{_State391, PostfixUnaryExprType}:             _GotoState53Action,
	{_State391, PrefixUnaryOpType}:                _GotoState55Action,
	{_State391, PrefixUnaryExprType}:              _GotoState54Action,
	{_State391, MulExprType}:                      _GotoState50Action,
	{_State391, AddExprType}:                      _GotoState39Action,
	{_State391, CmpExprType}:                      _GotoState45Action,
	{_State391, AndExprType}:                      _GotoState40Action,
	{_State391, OrExprType}:                       _GotoState52Action,
	{_State391, InitializableTypeType}:            _GotoState48Action,
	{_State391, ExplicitStructDefType}:            _GotoState46Action,
	{_State391, AnonymousFuncExprType}:            _GotoState41Action,
	{_State392, IfToken}:                          _GotoState131Action,
	{_State392, LbraceToken}:                      _GotoState59Action,
	{_State392, IfExprType}:                       _GotoState430Action,
	{_State392, BlockBodyType}:                    _GotoState429Action,
	{_State393, IntegerLiteralToken}:              _GotoState24Action,
	{_State393, FloatLiteralToken}:                _GotoState20Action,
	{_State393, RuneLiteralToken}:                 _GotoState32Action,
	{_State393, StringLiteralToken}:               _GotoState33Action,
	{_State393, IdentifierToken}:                  _GotoState23Action,
	{_State393, TrueToken}:                        _GotoState36Action,
	{_State393, FalseToken}:                       _GotoState19Action,
	{_State393, StructToken}:                      _GotoState34Action,
	{_State393, FuncToken}:                        _GotoState21Action,
	{_State393, VarToken}:                         _GotoState307Action,
	{_State393, LetToken}:                         _GotoState27Action,
	{_State393, LabelDeclToken}:                   _GotoState25Action,
	{_State393, LparenToken}:                      _GotoState28Action,
	{_State393, LbracketToken}:                    _GotoState26Action,
	{_State393, DotToken}:                         _GotoState306Action,
	{_State393, NotToken}:                         _GotoState30Action,
	{_State393, SubToken}:                         _GotoState35Action,
	{_State393, MulToken}:                         _GotoState29Action,
	{_State393, BitNegToken}:                      _GotoState18Action,
	{_State393, BitAndToken}:                      _GotoState17Action,
	{_State393, GreaterToken}:                     _GotoState22Action,
	{_State393, ParseErrorToken}:                  _GotoState31Action,
	{_State393, VarDeclPatternType}:               _GotoState57Action,
	{_State393, VarOrLetType}:                     _GotoState58Action,
	{_State393, AssignPatternType}:                _GotoState308Action,
	{_State393, CasePatternType}:                  _GotoState309Action,
	{_State393, OptionalLabelDeclType}:            _GotoState72Action,
	{_State393, CasePatternsType}:                 _GotoState431Action,
	{_State393, SequenceExprType}:                 _GotoState311Action,
	{_State393, BlockExprType}:                    _GotoState43Action,
	{_State393, CallExprType}:                     _GotoState44Action,
	{_State393, AtomExprType}:                     _GotoState42Action,
	{_State393, LiteralType}:                      _GotoState49Action,
	{_State393, ImplicitStructExprType}:           _GotoState47Action,
	{_State393, AccessExprType}:                   _GotoState38Action,
	{_State393, PostfixUnaryExprType}:             _GotoState53Action,
	{_State393, PrefixUnaryOpType}:                _GotoState55Action,
	{_State393, PrefixUnaryExprType}:              _GotoState54Action,
	{_State393, MulExprType}:                      _GotoState50Action,
	{_State393, AddExprType}:                      _GotoState39Action,
	{_State393, CmpExprType}:                      _GotoState45Action,
	{_State393, AndExprType}:                      _GotoState40Action,
	{_State393, OrExprType}:                       _GotoState52Action,
	{_State393, InitializableTypeType}:            _GotoState48Action,
	{_State393, ExplicitStructDefType}:            _GotoState46Action,
	{_State393, AnonymousFuncExprType}:            _GotoState41Action,
	{_State395, CaseToken}:                        _GotoState393Action,
	{_State395, DefaultToken}:                     _GotoState432Action,
	{_State395, CaseBranchType}:                   _GotoState433Action,
	{_State395, OptionalDefaultBranchType}:        _GotoState435Action,
	{_State395, DefaultBranchType}:                _GotoState434Action,
	{_State401, CommaToken}:                       _GotoState336Action,
	{_State403, NewlinesToken}:                    _GotoState437Action,
	{_State403, CommaToken}:                       _GotoState436Action,
	{_State405, StringLiteralToken}:               _GotoState342Action,
	{_State405, RparenToken}:                      _GotoState438Action,
	{_State405, ImportClauseType}:                 _GotoState403Action,
	{_State405, ImportClauseTerminalType}:         _GotoState439Action,
	{_State406, IdentifierToken}:                  _GotoState440Action,
	{_State408, AddToken}:                         _GotoState178Action,
	{_State408, SubToken}:                         _GotoState183Action,
	{_State408, MulToken}:                         _GotoState181Action,
	{_State409, IdentifierToken}:                  _GotoState80Action,
	{_State409, StructToken}:                      _GotoState34Action,
	{_State409, EnumToken}:                        _GotoState77Action,
	{_State409, TraitToken}:                       _GotoState85Action,
	{_State409, FuncToken}:                        _GotoState79Action,
	{_State409, LparenToken}:                      _GotoState81Action,
	{_State409, LbracketToken}:                    _GotoState26Action,
	{_State409, DotToken}:                         _GotoState76Action,
	{_State409, QuestionToken}:                    _GotoState83Action,
	{_State409, ExclaimToken}:                     _GotoState78Action,
	{_State409, TildeTildeToken}:                  _GotoState84Action,
	{_State409, BitNegToken}:                      _GotoState75Action,
	{_State409, BitAndToken}:                      _GotoState74Action,
	{_State409, ParseErrorToken}:                  _GotoState82Action,
	{_State409, InitializableTypeType}:            _GotoState91Action,
	{_State409, AtomTypeType}:                     _GotoState86Action,
	{_State409, ReturnableTypeType}:               _GotoState354Action,
	{_State409, ImplicitStructDefType}:            _GotoState90Action,
	{_State409, ExplicitStructDefType}:            _GotoState46Action,
	{_State409, ImplicitEnumDefType}:              _GotoState89Action,
	{_State409, ExplicitEnumDefType}:              _GotoState87Action,
	{_State409, TraitDefType}:                     _GotoState93Action,
	{_State409, ReturnTypeType}:                   _GotoState441Action,
	{_State409, FuncTypeType}:                     _GotoState88Action,
	{_State410, LparenToken}:                      _GotoState442Action,
	{_State416, AddToken}:                         _GotoState178Action,
	{_State416, SubToken}:                         _GotoState183Action,
	{_State416, MulToken}:                         _GotoState181Action,
	{_State419, StringLiteralToken}:               _GotoState443Action,
	{_State420, IdentifierToken}:                  _GotoState261Action,
	{_State420, StructToken}:                      _GotoState34Action,
	{_State420, EnumToken}:                        _GotoState77Action,
	{_State420, TraitToken}:                       _GotoState85Action,
	{_State420, FuncToken}:                        _GotoState79Action,
	{_State420, LparenToken}:                      _GotoState81Action,
	{_State420, LbracketToken}:                    _GotoState26Action,
	{_State420, DotToken}:                         _GotoState76Action,
	{_State420, QuestionToken}:                    _GotoState83Action,
	{_State420, ExclaimToken}:                     _GotoState78Action,
	{_State420, DotDotDotToken}:                   _GotoState260Action,
	{_State420, TildeTildeToken}:                  _GotoState84Action,
	{_State420, BitNegToken}:                      _GotoState75Action,
	{_State420, BitAndToken}:                      _GotoState74Action,
	{_State420, ParseErrorToken}:                  _GotoState82Action,
	{_State420, InitializableTypeType}:            _GotoState91Action,
	{_State420, AtomTypeType}:                     _GotoState86Action,
	{_State420, ReturnableTypeType}:               _GotoState92Action,
	{_State420, ValueTypeType}:                    _GotoState265Action,
	{_State420, ImplicitStructDefType}:            _GotoState90Action,
	{_State420, ExplicitStructDefType}:            _GotoState46Action,
	{_State420, ImplicitEnumDefType}:              _GotoState89Action,
	{_State420, ExplicitEnumDefType}:              _GotoState87Action,
	{_State420, TraitDefType}:                     _GotoState93Action,
	{_State420, ParameterDeclType}:                _GotoState263Action,
	{_State420, ParameterDeclsType}:               _GotoState264Action,
	{_State420, OptionalParameterDeclsType}:       _GotoState444Action,
	{_State420, FuncTypeType}:                     _GotoState88Action,
	{_State423, LbraceToken}:                      _GotoState59Action,
	{_State423, BlockBodyType}:                    _GotoState445Action,
	{_State424, IntegerLiteralToken}:              _GotoState24Action,
	{_State424, FloatLiteralToken}:                _GotoState20Action,
	{_State424, RuneLiteralToken}:                 _GotoState32Action,
	{_State424, StringLiteralToken}:               _GotoState33Action,
	{_State424, IdentifierToken}:                  _GotoState23Action,
	{_State424, TrueToken}:                        _GotoState36Action,
	{_State424, FalseToken}:                       _GotoState19Action,
	{_State424, StructToken}:                      _GotoState34Action,
	{_State424, FuncToken}:                        _GotoState21Action,
	{_State424, VarToken}:                         _GotoState37Action,
	{_State424, LetToken}:                         _GotoState27Action,
	{_State424, LabelDeclToken}:                   _GotoState25Action,
	{_State424, LparenToken}:                      _GotoState28Action,
	{_State424, LbracketToken}:                    _GotoState26Action,
	{_State424, NotToken}:                         _GotoState30Action,
	{_State424, SubToken}:                         _GotoState35Action,
	{_State424, MulToken}:                         _GotoState29Action,
	{_State424, BitNegToken}:                      _GotoState18Action,
	{_State424, BitAndToken}:                      _GotoState17Action,
	{_State424, GreaterToken}:                     _GotoState22Action,
	{_State424, ParseErrorToken}:                  _GotoState31Action,
	{_State424, VarDeclPatternType}:               _GotoState57Action,
	{_State424, VarOrLetType}:                     _GotoState58Action,
	{_State424, OptionalLabelDeclType}:            _GotoState72Action,
	{_State424, OptionalSequenceExprType}:         _GotoState446Action,
	{_State424, SequenceExprType}:                 _GotoState386Action,
	{_State424, BlockExprType}:                    _GotoState43Action,
	{_State424, CallExprType}:                     _GotoState44Action,
	{_State424, AtomExprType}:                     _GotoState42Action,
	{_State424, LiteralType}:                      _GotoState49Action,
	{_State424, ImplicitStructExprType}:           _GotoState47Action,
	{_State424, AccessExprType}:                   _GotoState38Action,
	{_State424, PostfixUnaryExprType}:             _GotoState53Action,
	{_State424, PrefixUnaryOpType}:                _GotoState55Action,
	{_State424, PrefixUnaryExprType}:              _GotoState54Action,
	{_State424, MulExprType}:                      _GotoState50Action,
	{_State424, AddExprType}:                      _GotoState39Action,
	{_State424, CmpExprType}:                      _GotoState45Action,
	{_State424, AndExprType}:                      _GotoState40Action,
	{_State424, OrExprType}:                       _GotoState52Action,
	{_State424, InitializableTypeType}:            _GotoState48Action,
	{_State424, ExplicitStructDefType}:            _GotoState46Action,
	{_State424, AnonymousFuncExprType}:            _GotoState41Action,
	{_State426, LparenToken}:                      _GotoState140Action,
	{_State426, TuplePatternType}:                 _GotoState447Action,
	{_State431, CommaToken}:                       _GotoState391Action,
	{_State431, ColonToken}:                       _GotoState448Action,
	{_State432, ColonToken}:                       _GotoState449Action,
	{_State435, RbraceToken}:                      _GotoState450Action,
	{_State441, LbraceToken}:                      _GotoState59Action,
	{_State441, BlockBodyType}:                    _GotoState451Action,
	{_State442, IdentifierToken}:                  _GotoState153Action,
	{_State442, ParameterDefType}:                 _GotoState156Action,
	{_State442, ParameterDefsType}:                _GotoState157Action,
	{_State442, OptionalParameterDefsType}:        _GotoState452Action,
	{_State444, RparenToken}:                      _GotoState453Action,
	{_State446, DoToken}:                          _GotoState454Action,
	{_State448, IntegerLiteralToken}:              _GotoState24Action,
	{_State448, FloatLiteralToken}:                _GotoState20Action,
	{_State448, RuneLiteralToken}:                 _GotoState32Action,
	{_State448, StringLiteralToken}:               _GotoState33Action,
	{_State448, IdentifierToken}:                  _GotoState23Action,
	{_State448, TrueToken}:                        _GotoState36Action,
	{_State448, FalseToken}:                       _GotoState19Action,
	{_State448, StructToken}:                      _GotoState34Action,
	{_State448, FuncToken}:                        _GotoState21Action,
	{_State448, VarToken}:                         _GotoState37Action,
	{_State448, LetToken}:                         _GotoState27Action,
	{_State448, LabelDeclToken}:                   _GotoState25Action,
	{_State448, LparenToken}:                      _GotoState28Action,
	{_State448, LbracketToken}:                    _GotoState26Action,
	{_State448, NotToken}:                         _GotoState30Action,
	{_State448, SubToken}:                         _GotoState35Action,
	{_State448, MulToken}:                         _GotoState29Action,
	{_State448, BitNegToken}:                      _GotoState18Action,
	{_State448, BitAndToken}:                      _GotoState17Action,
	{_State448, GreaterToken}:                     _GotoState22Action,
	{_State448, ParseErrorToken}:                  _GotoState31Action,
	{_State448, VarDeclPatternType}:               _GotoState57Action,
	{_State448, VarOrLetType}:                     _GotoState58Action,
	{_State448, OptionalLabelDeclType}:            _GotoState72Action,
	{_State448, SequenceExprType}:                 _GotoState455Action,
	{_State448, BlockExprType}:                    _GotoState43Action,
	{_State448, CallExprType}:                     _GotoState44Action,
	{_State448, AtomExprType}:                     _GotoState42Action,
	{_State448, LiteralType}:                      _GotoState49Action,
	{_State448, ImplicitStructExprType}:           _GotoState47Action,
	{_State448, AccessExprType}:                   _GotoState38Action,
	{_State448, PostfixUnaryExprType}:             _GotoState53Action,
	{_State448, PrefixUnaryOpType}:                _GotoState55Action,
	{_State448, PrefixUnaryExprType}:              _GotoState54Action,
	{_State448, MulExprType}:                      _GotoState50Action,
	{_State448, AddExprType}:                      _GotoState39Action,
	{_State448, CmpExprType}:                      _GotoState45Action,
	{_State448, AndExprType}:                      _GotoState40Action,
	{_State448, OrExprType}:                       _GotoState52Action,
	{_State448, InitializableTypeType}:            _GotoState48Action,
	{_State448, ExplicitStructDefType}:            _GotoState46Action,
	{_State448, AnonymousFuncExprType}:            _GotoState41Action,
	{_State449, IntegerLiteralToken}:              _GotoState24Action,
	{_State449, FloatLiteralToken}:                _GotoState20Action,
	{_State449, RuneLiteralToken}:                 _GotoState32Action,
	{_State449, StringLiteralToken}:               _GotoState33Action,
	{_State449, IdentifierToken}:                  _GotoState23Action,
	{_State449, TrueToken}:                        _GotoState36Action,
	{_State449, FalseToken}:                       _GotoState19Action,
	{_State449, StructToken}:                      _GotoState34Action,
	{_State449, FuncToken}:                        _GotoState21Action,
	{_State449, VarToken}:                         _GotoState37Action,
	{_State449, LetToken}:                         _GotoState27Action,
	{_State449, LabelDeclToken}:                   _GotoState25Action,
	{_State449, LparenToken}:                      _GotoState28Action,
	{_State449, LbracketToken}:                    _GotoState26Action,
	{_State449, NotToken}:                         _GotoState30Action,
	{_State449, SubToken}:                         _GotoState35Action,
	{_State449, MulToken}:                         _GotoState29Action,
	{_State449, BitNegToken}:                      _GotoState18Action,
	{_State449, BitAndToken}:                      _GotoState17Action,
	{_State449, GreaterToken}:                     _GotoState22Action,
	{_State449, ParseErrorToken}:                  _GotoState31Action,
	{_State449, VarDeclPatternType}:               _GotoState57Action,
	{_State449, VarOrLetType}:                     _GotoState58Action,
	{_State449, OptionalLabelDeclType}:            _GotoState72Action,
	{_State449, SequenceExprType}:                 _GotoState456Action,
	{_State449, BlockExprType}:                    _GotoState43Action,
	{_State449, CallExprType}:                     _GotoState44Action,
	{_State449, AtomExprType}:                     _GotoState42Action,
	{_State449, LiteralType}:                      _GotoState49Action,
	{_State449, ImplicitStructExprType}:           _GotoState47Action,
	{_State449, AccessExprType}:                   _GotoState38Action,
	{_State449, PostfixUnaryExprType}:             _GotoState53Action,
	{_State449, PrefixUnaryOpType}:                _GotoState55Action,
	{_State449, PrefixUnaryExprType}:              _GotoState54Action,
	{_State449, MulExprType}:                      _GotoState50Action,
	{_State449, AddExprType}:                      _GotoState39Action,
	{_State449, CmpExprType}:                      _GotoState45Action,
	{_State449, AndExprType}:                      _GotoState40Action,
	{_State449, OrExprType}:                       _GotoState52Action,
	{_State449, InitializableTypeType}:            _GotoState48Action,
	{_State449, ExplicitStructDefType}:            _GotoState46Action,
	{_State449, AnonymousFuncExprType}:            _GotoState41Action,
	{_State452, RparenToken}:                      _GotoState457Action,
	{_State453, IdentifierToken}:                  _GotoState80Action,
	{_State453, StructToken}:                      _GotoState34Action,
	{_State453, EnumToken}:                        _GotoState77Action,
	{_State453, TraitToken}:                       _GotoState85Action,
	{_State453, FuncToken}:                        _GotoState79Action,
	{_State453, LparenToken}:                      _GotoState81Action,
	{_State453, LbracketToken}:                    _GotoState26Action,
	{_State453, DotToken}:                         _GotoState76Action,
	{_State453, QuestionToken}:                    _GotoState83Action,
	{_State453, ExclaimToken}:                     _GotoState78Action,
	{_State453, TildeTildeToken}:                  _GotoState84Action,
	{_State453, BitNegToken}:                      _GotoState75Action,
	{_State453, BitAndToken}:                      _GotoState74Action,
	{_State453, ParseErrorToken}:                  _GotoState82Action,
	{_State453, InitializableTypeType}:            _GotoState91Action,
	{_State453, AtomTypeType}:                     _GotoState86Action,
	{_State453, ReturnableTypeType}:               _GotoState354Action,
	{_State453, ImplicitStructDefType}:            _GotoState90Action,
	{_State453, ExplicitStructDefType}:            _GotoState46Action,
	{_State453, ImplicitEnumDefType}:              _GotoState89Action,
	{_State453, ExplicitEnumDefType}:              _GotoState87Action,
	{_State453, TraitDefType}:                     _GotoState93Action,
	{_State453, ReturnTypeType}:                   _GotoState458Action,
	{_State453, FuncTypeType}:                     _GotoState88Action,
	{_State454, LbraceToken}:                      _GotoState59Action,
	{_State454, BlockBodyType}:                    _GotoState459Action,
	{_State457, IdentifierToken}:                  _GotoState80Action,
	{_State457, StructToken}:                      _GotoState34Action,
	{_State457, EnumToken}:                        _GotoState77Action,
	{_State457, TraitToken}:                       _GotoState85Action,
	{_State457, FuncToken}:                        _GotoState79Action,
	{_State457, LparenToken}:                      _GotoState81Action,
	{_State457, LbracketToken}:                    _GotoState26Action,
	{_State457, DotToken}:                         _GotoState76Action,
	{_State457, QuestionToken}:                    _GotoState83Action,
	{_State457, ExclaimToken}:                     _GotoState78Action,
	{_State457, TildeTildeToken}:                  _GotoState84Action,
	{_State457, BitNegToken}:                      _GotoState75Action,
	{_State457, BitAndToken}:                      _GotoState74Action,
	{_State457, ParseErrorToken}:                  _GotoState82Action,
	{_State457, InitializableTypeType}:            _GotoState91Action,
	{_State457, AtomTypeType}:                     _GotoState86Action,
	{_State457, ReturnableTypeType}:               _GotoState354Action,
	{_State457, ImplicitStructDefType}:            _GotoState90Action,
	{_State457, ExplicitStructDefType}:            _GotoState46Action,
	{_State457, ImplicitEnumDefType}:              _GotoState89Action,
	{_State457, ExplicitEnumDefType}:              _GotoState87Action,
	{_State457, TraitDefType}:                     _GotoState93Action,
	{_State457, ReturnTypeType}:                   _GotoState460Action,
	{_State457, FuncTypeType}:                     _GotoState88Action,
	{_State460, LbraceToken}:                      _GotoState59Action,
	{_State460, BlockBodyType}:                    _GotoState461Action,
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
	{_State43, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State44, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State45, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State46, _WildcardMarker}:                   _ReduceExplicitStructDefToInitializableTypeAction,
	{_State47, _WildcardMarker}:                   _ReduceImplicitStructExprToAtomExprAction,
	{_State49, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State50, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State52, _WildcardMarker}:                   _ReduceOrExprToSequenceExprAction,
	{_State53, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State54, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State55, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State56, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State57, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State59, _WildcardMarker}:                   _ReduceEmptyListToStatementsAction,
	{_State60, _WildcardMarker}:                   _ReduceBlockBodyToDefinitionAction,
	{_State61, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State62, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State63, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State66, _WildcardMarker}:                   _ReduceGlobalVarDefToDefinitionAction,
	{_State67, _WildcardMarker}:                   _ReduceEmptyListToPackageStatementsAction,
	{_State68, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State69, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State71, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State73, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State76, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State80, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State81, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State82, _WildcardMarker}:                   _ReduceParseErrorToAtomTypeAction,
	{_State86, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State87, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State88, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State89, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State90, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State91, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State92, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State93, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
	{_State95, _WildcardMarker}:                   _ReduceDotDotDotToArgumentAction,
	{_State96, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State97, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State99, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State100, _WildcardMarker}:                  _ReducePositionalToArgumentAction,
	{_State100, ColonToken}:                       _ReduceExpressionToOptionalExpressionAction,
	{_State102, RparenToken}:                      _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State103, RbracketToken}:                    _ReduceNilToOptionalGenericArgumentsAction,
	{_State105, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State105, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State106, _WildcardMarker}:                  _ReduceQuestionToPostfixUnaryExprAction,
	{_State108, _WildcardMarker}:                  _ReduceAddToAddOpAction,
	{_State109, _WildcardMarker}:                  _ReduceBitOrToAddOpAction,
	{_State110, _WildcardMarker}:                  _ReduceBitXorToAddOpAction,
	{_State111, _WildcardMarker}:                  _ReduceSubToAddOpAction,
	{_State112, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State113, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State114, _WildcardMarker}:                  _ReduceEqualToCmpOpAction,
	{_State115, _WildcardMarker}:                  _ReduceGreaterToCmpOpAction,
	{_State116, _WildcardMarker}:                  _ReduceGreaterOrEqualToCmpOpAction,
	{_State117, _WildcardMarker}:                  _ReduceLessToCmpOpAction,
	{_State118, _WildcardMarker}:                  _ReduceLessOrEqualToCmpOpAction,
	{_State119, _WildcardMarker}:                  _ReduceNotEqualToCmpOpAction,
	{_State120, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State121, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State121, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State122, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State123, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State124, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State125, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State126, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State127, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State128, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State130, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State131, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State132, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State133, _WildcardMarker}:                  _ReduceToBlockExprAction,
	{_State134, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State135, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State136, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State137, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State138, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State139, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State141, _WildcardMarker}:                  _ReduceTuplePatternToVarPatternAction,
	{_State142, _WildcardMarker}:                  _ReduceNilToOptionalValueTypeAction,
	{_State143, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State144, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State145, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State146, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State149, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State151, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State153, _WildcardMarker}:                  _ReduceInferredRefArgToParameterDefAction,
	{_State156, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State157, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State158, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State159, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State160, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State162, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State163, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State165, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State166, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State169, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State169, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State171, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State173, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State174, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State175, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State176, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State177, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State182, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
	{_State184, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State185, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State185, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State186, _WildcardMarker}:                  _ReduceToImplicitStructExprAction,
	{_State187, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State187, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State187, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State187, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State187, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State188, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State188, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State188, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State188, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State188, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State189, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State190, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State192, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State194, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State195, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State197, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State197, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State197, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State198, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State199, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State200, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State202, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State203, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State206, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State206, SemicolonToken}:                   _ReduceSequenceExprToForAssignmentAction,
	{_State207, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State209, LbraceToken}:                      _ReduceSequenceExprToConditionAction,
	{_State211, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State212, _WildcardMarker}:                  _ReduceDotDotDotToFieldVarPatternAction,
	{_State213, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State214, _WildcardMarker}:                  _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State216, _WildcardMarker}:                  _ReducePositionalToFieldVarPatternAction,
	{_State217, _EndMarker}:                       _ReduceToVarDeclPatternAction,
	{_State218, _WildcardMarker}:                  _ReduceValueTypeToOptionalValueTypeAction,
	{_State219, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State220, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State221, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State222, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State224, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State225, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State225, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State227, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State228, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State229, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State230, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State231, AssignToken}:                      _ReduceToAssignPatternAction,
	{_State231, _WildcardMarker}:                  _ReduceSequenceExprToExpressionAction,
	{_State232, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State234, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State235, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State236, _WildcardMarker}:                  _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State238, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State239, _WildcardMarker}:                  _ReduceImportStatementToPackageStatementBodyAction,
	{_State240, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State242, _WildcardMarker}:                  _ReduceUnsafeStatementToPackageStatementBodyAction,
	{_State243, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State244, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State245, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State246, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State248, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State249, _EndMarker}:                       _ReduceAliasToNamedFuncDefAction,
	{_State250, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State251, _WildcardMarker}:                  _ReduceInferredRefVarargToParameterDefAction,
	{_State252, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State254, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State258, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State261, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State263, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State264, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State265, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State266, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State267, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State268, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State273, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State275, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State277, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State278, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State280, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State281, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State282, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State285, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State286, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State287, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State288, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State289, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State290, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State291, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State294, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State296, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State297, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State298, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State300, _WildcardMarker}:                  _ReduceInitializeExprToAtomExprAction,
	{_State301, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State302, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State303, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State304, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State304, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State307, _WildcardMarker}:                  _ReduceVarToVarOrLetAction,
	{_State308, _WildcardMarker}:                  _ReduceAssignPatternToCasePatternAction,
	{_State309, _WildcardMarker}:                  _ReduceCasePatternToCasePatternsAction,
	{_State311, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State312, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State316, _WildcardMarker}:                  _ReduceToTuplePatternAction,
	{_State317, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State318, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State318, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State318, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State319, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State319, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State319, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State320, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State321, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State322, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State323, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State324, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State325, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State326, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State327, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State328, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State329, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State330, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State331, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State332, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State333, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State334, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State335, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State336, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State337, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State338, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State338, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State338, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State339, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State340, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State342, _WildcardMarker}:                  _ReduceStringLiteralToImportClauseAction,
	{_State343, _WildcardMarker}:                  _ReduceSingleToImportStatementAction,
	{_State344, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State345, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State346, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State348, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State351, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State352, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State354, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State355, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State358, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State361, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State363, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State364, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State366, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State368, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State369, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State370, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State371, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State373, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State376, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State377, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State378, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State379, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State380, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State381, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State382, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State383, SemicolonToken}:                   _ReduceAssignToForAssignmentAction,
	{_State386, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State387, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State390, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State391, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State393, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State394, _WildcardMarker}:                  _ReduceCaseBranchToCaseBranchesAction,
	{_State395, RbraceToken}:                      _ReduceNilToOptionalDefaultBranchAction,
	{_State396, _WildcardMarker}:                  _ReduceNamedToFieldVarPatternAction,
	{_State397, _WildcardMarker}:                  _ReduceAddToFieldVarPatternsAction,
	{_State398, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State399, _WildcardMarker}:                  _ReduceAssignStatementToStatementBodyAction,
	{_State400, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State401, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State402, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State404, _WildcardMarker}:                  _ReduceFirstToImportClausesAction,
	{_State407, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State408, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State409, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State411, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State412, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State413, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State413, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State414, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State415, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State415, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State416, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State417, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State418, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State420, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State421, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State422, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State424, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State424, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State425, _WildcardMarker}:                  _ReduceEnumMatchPatternToCasePatternAction,
	{_State427, LbraceToken}:                      _ReduceCaseToConditionAction,
	{_State428, _WildcardMarker}:                  _ReduceMultipleToCasePatternsAction,
	{_State429, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State430, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State433, _WildcardMarker}:                  _ReduceAddToCaseBranchesAction,
	{_State434, RbraceToken}:                      _ReduceDefaultBranchToOptionalDefaultBranchAction,
	{_State436, _WildcardMarker}:                  _ReduceExplicitToImportClauseTerminalAction,
	{_State437, _WildcardMarker}:                  _ReduceImplicitToImportClauseTerminalAction,
	{_State438, _WildcardMarker}:                  _ReduceMultipleToImportStatementAction,
	{_State439, _WildcardMarker}:                  _ReduceAddToImportClausesAction,
	{_State440, _WildcardMarker}:                  _ReduceAliasToImportClauseAction,
	{_State442, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State443, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State445, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State447, _WildcardMarker}:                  _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State448, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State449, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State450, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State451, _EndMarker}:                       _ReduceFuncDefToNamedFuncDefAction,
	{_State453, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State455, _WildcardMarker}:                  _ReduceToCaseBranchAction,
	{_State456, RbraceToken}:                      _ReduceToDefaultBranchAction,
	{_State457, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State458, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State459, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State461, _EndMarker}:                       _ReduceMethodDefToNamedFuncDefAction,
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 10
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
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
      LBRACE -> State 59
      definitions -> State 62
      definition -> State 61
      var_decl_pattern -> State 66
      var_or_let -> State 58
      block_body -> State 60
      type_def -> State 65
      named_func_def -> State 63
      package_def -> State 64

  State 14:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 67

  State 15:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 68

  State 16:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 69
      LPAREN -> State 70

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
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 71

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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 73
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
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
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 94
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

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
      IDENTIFIER -> State 96
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 95
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 100
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      arguments -> State 98
      argument -> State 97
      colon_expressions -> State 99
      optional_expression -> State 101
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
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
      LPAREN -> State 102

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
      LBRACKET -> State 105
      DOT -> State 104
      QUESTION -> State 106
      DOLLAR_LBRACKET -> State 103
      optional_generic_binding -> State 107

  State 39:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 108
      SUB -> State 111
      BIT_XOR -> State 110
      BIT_OR -> State 109
      add_op -> State 112

  State 40:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 113

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
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 114
      NOT_EQUAL -> State 119
      LESS -> State 117
      LESS_OR_EQUAL -> State 118
      GREATER -> State 115
      GREATER_OR_EQUAL -> State 116
      cmp_op -> State 120

  State 46:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      atom_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 121

  State 49:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 127
      DIV -> State 125
      MOD -> State 126
      BIT_AND -> State 122
      BIT_LSHIFT -> State 123
      BIT_RSHIFT -> State 124
      mul_op -> State 128

  State 51:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 131
      SWITCH -> State 132
      FOR -> State 130
      DO -> State 129
      LBRACE -> State 59
      if_expr -> State 134
      switch_expr -> State 136
      loop_expr -> State 135
      block_body -> State 133

  State 52:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 137

  State 53:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 55:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 138
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 56:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      LPAREN -> State 140
      var_pattern -> State 142
      tuple_pattern -> State 141

  State 59:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 143

  State 60:
    Kernel Items:
      definition: block_body., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 144
      optional_newlines -> State 145

  State 63:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 146

  State 67:
    Kernel Items:
      package_def: PACKAGE LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 147

  State 68:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 149
      ASSIGN -> State 148
      optional_generic_parameters -> State 150

  State 69:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 149
      ASSIGN -> State 151
      optional_generic_parameters -> State 152

  State 70:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 154

  State 71:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 156
      parameter_defs -> State 157
      optional_parameter_defs -> State 155

  State 72:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 133

  State 73:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 158
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 75:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 159
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 76:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 103
      optional_generic_binding -> State 160

  State 77:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 161

  State 78:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 162
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 79:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 163

  State 80:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 164
      DOLLAR_LBRACKET -> State 103
      optional_generic_binding -> State 165

  State 81:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 169
      implicit_field_defs -> State 171
      optional_implicit_field_defs -> State 172
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 168
      implicit_enum_value_defs -> State 170
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 82:
    Kernel Items:
      atom_type: PARSE_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 175
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 84:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 176
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 85:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 177

  State 86:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 94:
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
      RBRACKET -> State 182
      COMMA -> State 180
      COLON -> State 179
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 95:
    Kernel Items:
      argument: DOT_DOT_DOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 184

  State 97:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      arguments: arguments.COMMA argument
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 186
      COMMA -> State 185

  State 99:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 187

  State 100:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 188

  State 102:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 190
      implicit_struct_def -> State 90
      explicit_field_defs -> State 189
      optional_explicit_field_defs -> State 191
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 103:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_generic_arguments -> State 193
      generic_arguments -> State 192
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 194
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 104:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 195

  State 105:
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
      IDENTIFIER -> State 96
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 95
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 100
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      argument -> State 196
      colon_expressions -> State 99
      optional_expression -> State 101
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 106:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 197

  State 108:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 112:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 198
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 113:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 199
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 114:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 120:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 200
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 121:
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
      IDENTIFIER -> State 96
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 95
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 100
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      arguments -> State 201
      argument -> State 97
      colon_expressions -> State 99
      optional_expression -> State 101
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 122:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 128:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 202
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 129:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 203

  State 130:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.assign_pattern IN sequence_expr DO block_body
      loop_expr: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      assign_pattern -> State 204
      optional_label_decl -> State 72
      for_assignment -> State 205
      sequence_expr -> State 206
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 131:
    Kernel Items:
      if_expr: IF.condition block_body
      if_expr: IF.condition block_body ELSE block_body
      if_expr: IF.condition block_body ELSE if_expr
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
      CASE -> State 207
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      condition -> State 208
      sequence_expr -> State 209
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 132:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE case_branches optional_default_branch RBRACE
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 210
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 133:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 137:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 211
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 138:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

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
      IDENTIFIER -> State 213
      LPAREN -> State 140
      DOT_DOT_DOT -> State 212
      var_pattern -> State 216
      tuple_pattern -> State 141
      field_var_patterns -> State 215
      field_var_pattern -> State 214

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
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_value_type -> State 217
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 218
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 143:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
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
      RETURN -> State 224
      BREAK -> State 220
      CONTINUE -> State 221
      UNSAFE -> State 167
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 219
      DEFER -> State 222
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      RBRACE -> State 223
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      assign_pattern -> State 226
      expression -> State 227
      optional_label_decl -> State 51
      sequence_expr -> State 231
      block_expr -> State 43
      statement -> State 232
      statement_body -> State 233
      unsafe_statement -> State 234
      jump_statement -> State 229
      jump_type -> State 230
      expressions -> State 228
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 225
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 144:
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
      LBRACE -> State 59
      definition -> State 235
      var_decl_pattern -> State 66
      var_or_let -> State 58
      block_body -> State 60
      type_def -> State 65
      named_func_def -> State 63
      package_def -> State 64

  State 145:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 146:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 236
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 147:
    Kernel Items:
      package_def: PACKAGE LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      IMPORT -> State 237
      UNSAFE -> State 167
      RPAREN -> State 238
      unsafe_statement -> State 242
      package_statement_body -> State 241
      package_statement -> State 240
      import_statement -> State 239

  State 148:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 243
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 149:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 244
      generic_parameter_def -> State 245
      generic_parameter_defs -> State 246
      optional_generic_parameter_defs -> State 247

  State 150:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 248
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 151:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 249
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 152:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 250

  State 153:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.DOT_DOT_DOT
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 251
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 252
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 154:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 253

  State 155:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 254

  State 156:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 255

  State 158:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 256
      implicit_enum_value_defs -> State 259
      implicit_enum_def -> State 89
      explicit_enum_value_defs -> State 257
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 162:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 261
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 260
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 265
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      parameter_decl -> State 263
      parameter_decls -> State 264
      optional_parameter_decls -> State 262
      func_type -> State 88

  State 164:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 266

  State 165:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 267
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOLLAR_LBRACKET -> State 103
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_generic_binding -> State 165
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 268
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 167:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 269

  State 168:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 270

  State 169:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 271

  State 170:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 273
      OR -> State 272

  State 171:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 274

  State 172:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 275

  State 173:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 175:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 276
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 277
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_property -> State 281
      trait_properties -> State 280
      optional_trait_properties -> State 279
      trait_def -> State 93
      func_type -> State 88
      method_signature -> State 278

  State 178:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 282
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 179:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 283
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 180:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 284

  State 181:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 285
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 182:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 286
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 184:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 287
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 185:
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
      IDENTIFIER -> State 96
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 95
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 100
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      argument -> State 288
      colon_expressions -> State 99
      optional_expression -> State 101
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 186:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 187:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 289
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      optional_expression -> State 290
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 188:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 289
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      optional_expression -> State 291
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 189:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 293
      COMMA -> State 292

  State 190:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 294

  State 192:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 295

  State 193:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 296

  State 194:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 195:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 297

  State 197:
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
      IDENTIFIER -> State 96
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 95
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 100
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      optional_arguments -> State 299
      arguments -> State 298
      argument -> State 97
      colon_expressions -> State 99
      optional_expression -> State 101
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 198:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 127
      DIV -> State 125
      MOD -> State 126
      BIT_AND -> State 122
      BIT_LSHIFT -> State 123
      BIT_RSHIFT -> State 124
      mul_op -> State 128

  State 199:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 114
      NOT_EQUAL -> State 119
      LESS -> State 117
      LESS_OR_EQUAL -> State 118
      GREATER -> State 115
      GREATER_OR_EQUAL -> State 116
      cmp_op -> State 120

  State 200:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 108
      SUB -> State 111
      BIT_XOR -> State 110
      BIT_OR -> State 109
      add_op -> State 112

  State 201:
    Kernel Items:
      arguments: arguments.COMMA argument
      atom_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 300
      COMMA -> State 185

  State 202:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 301

  State 204:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO block_body
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 303
      ASSIGN -> State 302

  State 205:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 304

  State 206:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO block_body
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 305

  State 207:
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
      VAR -> State 307
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 306
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      assign_pattern -> State 308
      case_pattern -> State 309
      optional_label_decl -> State 72
      case_patterns -> State 310
      sequence_expr -> State 311
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 208:
    Kernel Items:
      if_expr: IF condition.block_body
      if_expr: IF condition.block_body ELSE block_body
      if_expr: IF condition.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 312

  State 209:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 313

  State 211:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 113

  State 212:
    Kernel Items:
      field_var_pattern: DOT_DOT_DOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 314

  State 214:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 316
      COMMA -> State 315

  State 216:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      optional_value_type: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 219:
    Kernel Items:
      statement_body: ASYNC.call_expr
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
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 318
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 317
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 220:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      statement_body: DEFER.call_expr
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
      optional_label_decl -> State 72
      block_expr -> State 43
      call_expr -> State 319
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 317
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 223:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      statement_body: access_expr.unary_op_assign
      statement_body: access_expr.binary_op_assign expression
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
      postfix_unary_expr: access_expr., *
      postfix_unary_expr: access_expr.QUESTION
    Reduce:
      * -> [postfix_unary_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 105
      DOT -> State 104
      QUESTION -> State 106
      DOLLAR_LBRACKET -> State 103
      ADD_ASSIGN -> State 320
      SUB_ASSIGN -> State 331
      MUL_ASSIGN -> State 330
      DIV_ASSIGN -> State 328
      MOD_ASSIGN -> State 329
      ADD_ONE_ASSIGN -> State 321
      SUB_ONE_ASSIGN -> State 332
      BIT_NEG_ASSIGN -> State 324
      BIT_AND_ASSIGN -> State 322
      BIT_OR_ASSIGN -> State 325
      BIT_XOR_ASSIGN -> State 327
      BIT_LSHIFT_ASSIGN -> State 323
      BIT_RSHIFT_ASSIGN -> State 326
      unary_op_assign -> State 334
      binary_op_assign -> State 333
      optional_generic_binding -> State 107

  State 226:
    Kernel Items:
      statement_body: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 335

  State 227:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 336

  State 229:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 337
      optional_jump_label -> State 338

  State 231:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 339
      SEMICOLON -> State 340

  State 234:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 342
      LPAREN -> State 341
      import_clause -> State 343

  State 238:
    Kernel Items:
      package_def: PACKAGE LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      package_statement_body: import_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 345
      COMMA -> State 344

  State 242:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 244:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 346
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 245:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 347

  State 247:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 348

  State 248:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 349
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 249:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 156
      parameter_defs -> State 157
      optional_parameter_defs -> State 350

  State 251:
    Kernel Items:
      parameter_def: IDENTIFIER DOT_DOT_DOT., *
      parameter_def: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 351
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 252:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 253:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 352

  State 254:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 354
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 353
      func_type -> State 88

  State 255:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 355

  State 256:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 356
      OR -> State 357

  State 257:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 358

  State 258:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 271

  State 259:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 359
      OR -> State 360

  State 260:
    Kernel Items:
      parameter_decl: DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 361
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 261:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 267
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOLLAR_LBRACKET -> State 103
      DOT_DOT_DOT -> State 362
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_generic_binding -> State 165
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 363
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 262:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 364

  State 263:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 365

  State 265:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 266:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 103
      optional_generic_binding -> State 366

  State 267:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 266
      DOLLAR_LBRACKET -> State 103
      optional_generic_binding -> State 160

  State 268:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 269:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 367

  State 270:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 368
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 271:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 369

  State 272:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 370
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 273:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 371
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 275:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 372
      LPAREN -> State 163

  State 277:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 373

  State 280:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 375
      COMMA -> State 374

  State 281:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 376
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 284:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 377

  State 285:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 378
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 293:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 379
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 294:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 380
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 296:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 298:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 185

  State 299:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 381

  State 300:
    Kernel Items:
      atom_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 382
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 302:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 383
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 303:
    Kernel Items:
      loop_expr: FOR assign_pattern IN.sequence_expr DO block_body
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 384
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 304:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      optional_sequence_expr -> State 385
      sequence_expr -> State 386
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 305:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 387

  State 306:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 388

  State 307:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 389

  State 308:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      condition: CASE case_patterns.ASSIGN sequence_expr
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 391
      ASSIGN -> State 390

  State 311:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      if_expr: IF condition block_body., *
      if_expr: IF condition block_body.ELSE block_body
      if_expr: IF condition block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 392

  State 313:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 393
      case_branches -> State 395
      case_branch -> State 394

  State 314:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      LPAREN -> State 140
      var_pattern -> State 396
      tuple_pattern -> State 141

  State 315:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 213
      LPAREN -> State 140
      DOT_DOT_DOT -> State 212
      var_pattern -> State 216
      tuple_pattern -> State 141
      field_var_pattern -> State 397

  State 316:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 105
      DOT -> State 104
      DOLLAR_LBRACKET -> State 103
      optional_generic_binding -> State 107

  State 318:
    Kernel Items:
      statement_body: ASYNC call_expr., NEWLINES
      statement_body: ASYNC call_expr., SEMICOLON
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
      NEWLINES -> [statement_body]
      SEMICOLON -> [statement_body]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      statement_body: DEFER call_expr., NEWLINES
      statement_body: DEFER call_expr., SEMICOLON
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
      NEWLINES -> [statement_body]
      SEMICOLON -> [statement_body]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 398
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 334:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      statement_body: assign_pattern ASSIGN.expression
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 399
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 336:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 400
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 337:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      expression -> State 227
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      expressions -> State 401
      optional_expressions -> State 402
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 339:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 342
      import_clause -> State 403
      import_clause_terminal -> State 404
      import_clauses -> State 405

  State 342:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 406

  State 343:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      package_statement: package_statement_body COMMA., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 347:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 244
      generic_parameter_def -> State 407

  State 348:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 408
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 350:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 409

  State 351:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 352:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 149
      optional_generic_parameters -> State 410

  State 353:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 411

  State 354:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 412
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 357:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 413
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 358:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 414
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 360:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 258
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      enum_value_def -> State 415
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 361:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 362:
    Kernel Items:
      parameter_decl: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 416
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 363:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 364:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 354
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 417
      func_type -> State 88

  State 365:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 261
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 260
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 265
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      parameter_decl -> State 418
      func_type -> State 88

  State 366:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 367:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 419

  State 368:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 420

  State 373:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 276
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 277
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_property -> State 421
      trait_def -> State 93
      func_type -> State 88
      method_signature -> State 278

  State 375:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 276
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 277
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_property -> State 422
      trait_def -> State 93
      func_type -> State 88
      method_signature -> State 278

  State 376:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 377:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 381:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 383:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 423

  State 385:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 424

  State 386:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 28
      implicit_struct_expr -> State 425

  State 389:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 426

  State 390:
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 427
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 391:
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
      VAR -> State 307
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 306
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      assign_pattern -> State 308
      case_pattern -> State 428
      optional_label_decl -> State 72
      sequence_expr -> State 311
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 392:
    Kernel Items:
      if_expr: IF condition block_body ELSE.block_body
      if_expr: IF condition block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 131
      LBRACE -> State 59
      if_expr -> State 430
      block_body -> State 429

  State 393:
    Kernel Items:
      case_branch: CASE.case_patterns COLON sequence_expr
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
      VAR -> State 307
      LET -> State 27
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 306
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      assign_pattern -> State 308
      case_pattern -> State 309
      optional_label_decl -> State 72
      case_patterns -> State 431
      sequence_expr -> State 311
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 394:
    Kernel Items:
      case_branches: case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 395:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches.optional_default_branch RBRACE
      case_branches: case_branches.case_branch
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      CASE -> State 393
      DEFAULT -> State 432
      case_branch -> State 433
      optional_default_branch -> State 435
      default_branch -> State 434

  State 396:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 397:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 398:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      statement_body: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 336

  State 402:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      import_clause_terminal: import_clause.NEWLINES
      import_clause_terminal: import_clause.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 437
      COMMA -> State 436

  State 404:
    Kernel Items:
      import_clauses: import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
      import_clauses: import_clauses.import_clause_terminal
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 342
      RPAREN -> State 438
      import_clause -> State 403
      import_clause_terminal -> State 439

  State 406:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 440

  State 407:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 408:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 409:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 354
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 441
      func_type -> State 88

  State 410:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 442

  State 411:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 413:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 417:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 443

  State 420:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 261
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 260
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 265
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      parameter_decl -> State 263
      parameter_decls -> State 264
      optional_parameter_decls -> State 444
      func_type -> State 88

  State 421:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 422:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 445

  State 424:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      optional_sequence_expr -> State 446
      sequence_expr -> State 386
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 425:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 140
      tuple_pattern -> State 447

  State 427:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      if_expr: IF condition block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      if_expr: IF condition block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      case_branch: CASE case_patterns.COLON sequence_expr
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 391
      COLON -> State 448

  State 432:
    Kernel Items:
      default_branch: DEFAULT.COLON sequence_expr
    Reduce:
      (nil)
    Goto:
      COLON -> State 449

  State 433:
    Kernel Items:
      case_branches: case_branches case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      optional_default_branch: default_branch., RBRACE
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 450

  State 436:
    Kernel Items:
      import_clause_terminal: import_clause COMMA., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 437:
    Kernel Items:
      import_clause_terminal: import_clause NEWLINES., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      import_clauses: import_clauses import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 451

  State 442:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 156
      parameter_defs -> State 157
      optional_parameter_defs -> State 452

  State 443:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 453

  State 445:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 446:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 454

  State 447:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      case_branch: CASE case_patterns COLON.sequence_expr
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 455
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 449:
    Kernel Items:
      default_branch: DEFAULT COLON.sequence_expr
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
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      NOT -> State 30
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 57
      var_or_let -> State 58
      optional_label_decl -> State 72
      sequence_expr -> State 456
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 450:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 451:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 452:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 457

  State 453:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 354
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 458
      func_type -> State 88

  State 454:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 459

  State 455:
    Kernel Items:
      case_branch: CASE case_patterns COLON sequence_expr., *
    Reduce:
      * -> [case_branch]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      default_branch: DEFAULT COLON sequence_expr., RBRACE
    Reduce:
      RBRACE -> [default_branch]
    Goto:
      (nil)

  State 457:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 354
      implicit_struct_def -> State 90
      explicit_struct_def -> State 46
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 460
      func_type -> State 88

  State 458:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 459:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 460:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      block_body -> State 461

  State 461:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 461
Number of shift actions: 3193
Number of reduce actions: 371
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 4272
Number of unoptimized shift actions: 32206
Number of unoptimized reduce actions: 24777
*/
