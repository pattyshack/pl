// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type SymbolId int

const (
	SpacesToken          = SymbolId(256)
	NewlinesToken        = SymbolId(257)
	CommentToken         = SymbolId(258)
	IntegerLiteralToken  = SymbolId(259)
	FloatLiteralToken    = SymbolId(260)
	RuneLiteralToken     = SymbolId(261)
	StringLiteralToken   = SymbolId(262)
	IdentifierToken      = SymbolId(263)
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
	LabelDeclToken       = SymbolId(291)
	JumpLabelToken       = SymbolId(292)
	LbraceToken          = SymbolId(293)
	RbraceToken          = SymbolId(294)
	LparenToken          = SymbolId(295)
	RparenToken          = SymbolId(296)
	LbracketToken        = SymbolId(297)
	RbracketToken        = SymbolId(298)
	DotToken             = SymbolId(299)
	CommaToken           = SymbolId(300)
	QuestionToken        = SymbolId(301)
	SemicolonToken       = SymbolId(302)
	ColonToken           = SymbolId(303)
	ExclaimToken         = SymbolId(304)
	DollarLbracketToken  = SymbolId(305)
	DotDotDotToken       = SymbolId(306)
	TildeTildeToken      = SymbolId(307)
	AssignToken          = SymbolId(308)
	AddAssignToken       = SymbolId(309)
	SubAssignToken       = SymbolId(310)
	MulAssignToken       = SymbolId(311)
	DivAssignToken       = SymbolId(312)
	ModAssignToken       = SymbolId(313)
	AddOneAssignToken    = SymbolId(314)
	SubOneAssignToken    = SymbolId(315)
	BitNegAssignToken    = SymbolId(316)
	BitAndAssignToken    = SymbolId(317)
	BitOrAssignToken     = SymbolId(318)
	BitXorAssignToken    = SymbolId(319)
	BitLshiftAssignToken = SymbolId(320)
	BitRshiftAssignToken = SymbolId(321)
	NotToken             = SymbolId(322)
	AndToken             = SymbolId(323)
	OrToken              = SymbolId(324)
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
	LexErrorToken        = SymbolId(342)
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
	// 52:10: source -> ...
	ToSource(OptionalDefinitions_ *GenericSymbol) (*GenericSymbol, error)

	// 55:2: optional_definitions -> NEWLINES: ...
	NewlinesToOptionalDefinitions(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 56:2: optional_definitions -> definitions: ...
	DefinitionsToOptionalDefinitions(OptionalNewlines_ *GenericSymbol, Definitions_ *GenericSymbol, OptionalNewlines_2 *GenericSymbol) (*GenericSymbol, error)

	// 57:2: optional_definitions -> nil: ...
	NilToOptionalDefinitions() (*GenericSymbol, error)

	// 60:2: optional_newlines -> NEWLINES: ...
	NewlinesToOptionalNewlines(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 61:2: optional_newlines -> nil: ...
	NilToOptionalNewlines() (*GenericSymbol, error)

	// 64:2: definitions -> nil: ...
	NilToDefinitions(Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 65:2: definitions -> add: ...
	AddToDefinitions(Definitions_ *GenericSymbol, Newlines_ *GenericSymbol, Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 69:2: definition -> package_def: ...
	PackageDefToDefinition(PackageDef_ *GenericSymbol) (*GenericSymbol, error)

	// 70:2: definition -> type_def: ...
	TypeDefToDefinition(TypeDef_ *GenericSymbol) (*GenericSymbol, error)

	// 71:2: definition -> named_func_def: ...
	NamedFuncDefToDefinition(NamedFuncDef_ *GenericSymbol) (*GenericSymbol, error)

	// 72:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 73:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 76:2: definition -> block_body: ...
	BlockBodyToDefinition(BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 85:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ *GenericSymbol, VarPattern_ *GenericSymbol, OptionalValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 87:14: var_or_let -> VAR: ...
	VarToVarOrLet(Var_ *GenericSymbol) (*GenericSymbol, error)

	// 87:20: var_or_let -> LET: ...
	LetToVarOrLet(Let_ *GenericSymbol) (*GenericSymbol, error)

	// 90:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 91:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 93:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ *GenericSymbol, FieldVarPatterns_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 96:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 97:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ *GenericSymbol, Comma_ *GenericSymbol, FieldVarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 100:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, VarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 102:2: field_var_pattern -> DOT_DOT_DOT: ...
	DotDotDotToFieldVarPattern(DotDotDot_ *GenericSymbol) (*GenericSymbol, error)

	// 105:2: optional_value_type -> value_type: ...
	ValueTypeToOptionalValueType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 106:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (*GenericSymbol, error)

	// 112:18: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 115:2: case_pattern -> assign_pattern: ...
	AssignPatternToCasePattern(AssignPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ *GenericSymbol, Identifier_ *GenericSymbol, ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 123:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol, TuplePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 131:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 136:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 137:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 146:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, Condition_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, Condition_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 148:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, Condition_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: condition -> case: ...
	CaseToCondition(Case_ *GenericSymbol, CasePatterns_ *GenericSymbol, Assign_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 176:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, CaseBranches_ *GenericSymbol, OptionalDefaultBranch_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 179:2: case_branches -> case_branch: ...
	CaseBranchToCaseBranches(CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 180:2: case_branches -> add: ...
	AddToCaseBranches(CaseBranches_ *GenericSymbol, CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 182:15: case_branch -> ...
	ToCaseBranch(Case_ *GenericSymbol, CasePatterns_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ *GenericSymbol, Comma_ *GenericSymbol, CasePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: optional_default_branch -> default_branch: ...
	DefaultBranchToOptionalDefaultBranch(DefaultBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 190:2: optional_default_branch -> nil: ...
	NilToOptionalDefaultBranch() (*GenericSymbol, error)

	// 192:18: default_branch -> ...
	ToDefaultBranch(Default_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol, For_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 208:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ *GenericSymbol, AssignPattern_ *GenericSymbol, In_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: loop_expr -> for: ...
	ForToLoopExpr(For_ *GenericSymbol, ForAssignment_ *GenericSymbol, Semicolon_ *GenericSymbol, OptionalSequenceExpr_ *GenericSymbol, Semicolon_2 *GenericSymbol, OptionalSequenceExpr_2 *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 213:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 214:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (*GenericSymbol, error)

	// 217:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ *GenericSymbol, Assign_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: sequence_expr -> or_expr: ...
	OrExprToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 228:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 232:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 234:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 236:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 239:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 240:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 243:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 244:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 247:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 273:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 276:2: statement_body -> assign_statement: ...
	AssignStatementToStatementBody(AssignPattern_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 280:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 281:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 295:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 296:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 299:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 300:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 301:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 302:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 310:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 317:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 321:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 322:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 325:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 326:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 329:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 333:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 334:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 340:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 343:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 347:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 352:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 360:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 364:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 365:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 366:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 369:2: argument -> DOT_DOT_DOT: ...
	DotDotDotToArgument(DotDotDot_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 373:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 376:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 377:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 387:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializableType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 392:2: atom_expr -> implicit_struct_expr: ...
	ImplicitStructExprToAtomExpr(ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 393:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 397:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 398:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 399:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 400:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 401:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 403:24: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 407:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 416:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 417:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 418:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 421:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 427:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 428:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 431:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 432:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 434:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 435:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 436:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 439:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 440:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 443:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 444:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 445:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 449:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 453:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 455:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 456:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 457:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 461:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 462:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 465:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 466:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 469:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 470:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 479:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 480:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 481:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Comma_ *GenericSymbol, IntegerLiteral_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 482:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Colon_ *GenericSymbol, ValueType_2 *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ *GenericSymbol) (*GenericSymbol, error)

	// 486:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 487:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_2 *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 490:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 491:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 492:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 493:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 494:2: atom_type -> LEX_ERROR: ...
	LexErrorToAtomType(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 500:2: returnable_type -> atom_type: ...
	AtomTypeToReturnableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 501:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 502:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 503:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 504:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 505:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: value_type -> returnable_type: ...
	ReturnableTypeToValueType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 511:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ *GenericSymbol, Mul_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 512:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ *GenericSymbol, Add_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 513:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ *GenericSymbol, Sub_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 516:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 517:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 518:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Assign_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 526:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 527:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 530:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 531:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 534:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 535:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 538:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 539:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 546:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 547:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 548:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 551:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 552:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 555:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 556:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 558:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 561:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 562:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 563:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 566:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 567:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 569:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 577:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 578:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 590:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 591:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 593:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 596:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 597:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 598:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 599:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 601:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 608:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 609:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 612:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 613:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 614:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 617:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 618:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 620:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 628:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 629:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 632:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 633:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, DotDotDot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 634:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 635:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(DotDotDot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 638:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 639:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 642:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 643:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 645:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 656:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 662:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 663:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ *GenericSymbol, DotDotDot_ *GenericSymbol) (*GenericSymbol, error)

	// 664:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 665:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, DotDotDot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 668:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 669:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 672:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 673:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 676:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 677:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ *GenericSymbol, Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_2 *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_2 *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 678:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 682:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 694:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol) (*GenericSymbol, error)

	// 695:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 698:2: package_statement_body -> unsafe_statement: ...
	UnsafeStatementToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 699:2: package_statement_body -> import_statement: ...
	ImportStatementToPackageStatementBody(ImportStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 702:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 703:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 706:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 707:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 710:2: import_statement -> single: ...
	SingleToImportStatement(Import_ *GenericSymbol, ImportClause_ *GenericSymbol) (*GenericSymbol, error)

	// 711:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ *GenericSymbol, Lparen_ *GenericSymbol, ImportClauses_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 714:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 715:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ *GenericSymbol, As_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 718:2: import_clause_terminal -> implicit: ...
	ImplicitToImportClauseTerminal(ImportClause_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 719:2: import_clause_terminal -> explicit: ...
	ExplicitToImportClauseTerminal(ImportClause_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 722:2: import_clauses -> first: ...
	FirstToImportClauses(ImportClauseTerminal_ *GenericSymbol) (*GenericSymbol, error)

	// 723:2: import_clauses -> add: ...
	AddToImportClauses(ImportClauses_ *GenericSymbol, ImportClauseTerminal_ *GenericSymbol) (*GenericSymbol, error)

	// 727:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 728:2: lex_internal_tokens -> COMMENT: ...
	CommentToLexInternalTokens(Comment_ *GenericSymbol) (*GenericSymbol, error)
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

func ParseLexInternalTokens(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseLexInternalTokensWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseLexInternalTokensWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State6)
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
	case SpacesToken:
		return "SPACES"
	case NewlinesToken:
		return "NEWLINES"
	case CommentToken:
		return "COMMENT"
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
	case LexErrorToken:
		return "LEX_ERROR"
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
	case LexInternalTokensType:
		return "lex_internal_tokens"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	SourceType                       = SymbolId(343)
	OptionalDefinitionsType          = SymbolId(344)
	OptionalNewlinesType             = SymbolId(345)
	DefinitionsType                  = SymbolId(346)
	DefinitionType                   = SymbolId(347)
	VarDeclPatternType               = SymbolId(348)
	VarOrLetType                     = SymbolId(349)
	VarPatternType                   = SymbolId(350)
	TuplePatternType                 = SymbolId(351)
	FieldVarPatternsType             = SymbolId(352)
	FieldVarPatternType              = SymbolId(353)
	OptionalValueTypeType            = SymbolId(354)
	AssignPatternType                = SymbolId(355)
	CasePatternType                  = SymbolId(356)
	ExpressionType                   = SymbolId(357)
	OptionalLabelDeclType            = SymbolId(358)
	IfExprType                       = SymbolId(359)
	ConditionType                    = SymbolId(360)
	SwitchExprType                   = SymbolId(361)
	CaseBranchesType                 = SymbolId(362)
	CaseBranchType                   = SymbolId(363)
	CasePatternsType                 = SymbolId(364)
	OptionalDefaultBranchType        = SymbolId(365)
	DefaultBranchType                = SymbolId(366)
	LoopExprType                     = SymbolId(367)
	OptionalSequenceExprType         = SymbolId(368)
	ForAssignmentType                = SymbolId(369)
	SequenceExprType                 = SymbolId(370)
	BlockExprType                    = SymbolId(371)
	BlockBodyType                    = SymbolId(372)
	StatementsType                   = SymbolId(373)
	StatementType                    = SymbolId(374)
	StatementBodyType                = SymbolId(375)
	UnaryOpAssignType                = SymbolId(376)
	BinaryOpAssignType               = SymbolId(377)
	UnsafeStatementType              = SymbolId(378)
	JumpStatementType                = SymbolId(379)
	JumpTypeType                     = SymbolId(380)
	OptionalJumpLabelType            = SymbolId(381)
	ExpressionsType                  = SymbolId(382)
	OptionalExpressionsType          = SymbolId(383)
	CallExprType                     = SymbolId(384)
	OptionalGenericBindingType       = SymbolId(385)
	OptionalGenericArgumentsType     = SymbolId(386)
	GenericArgumentsType             = SymbolId(387)
	OptionalArgumentsType            = SymbolId(388)
	ArgumentsType                    = SymbolId(389)
	ArgumentType                     = SymbolId(390)
	ColonExpressionsType             = SymbolId(391)
	OptionalExpressionType           = SymbolId(392)
	AtomExprType                     = SymbolId(393)
	LiteralType                      = SymbolId(394)
	ImplicitStructExprType           = SymbolId(395)
	AccessExprType                   = SymbolId(396)
	PostfixUnaryExprType             = SymbolId(397)
	PrefixUnaryOpType                = SymbolId(398)
	PrefixUnaryExprType              = SymbolId(399)
	MulOpType                        = SymbolId(400)
	MulExprType                      = SymbolId(401)
	AddOpType                        = SymbolId(402)
	AddExprType                      = SymbolId(403)
	CmpOpType                        = SymbolId(404)
	CmpExprType                      = SymbolId(405)
	AndExprType                      = SymbolId(406)
	OrExprType                       = SymbolId(407)
	InitializableTypeType            = SymbolId(408)
	AtomTypeType                     = SymbolId(409)
	ReturnableTypeType               = SymbolId(410)
	ValueTypeType                    = SymbolId(411)
	TypeDefType                      = SymbolId(412)
	GenericParameterDefType          = SymbolId(413)
	GenericParameterDefsType         = SymbolId(414)
	OptionalGenericParameterDefsType = SymbolId(415)
	OptionalGenericParametersType    = SymbolId(416)
	FieldDefType                     = SymbolId(417)
	ImplicitFieldDefsType            = SymbolId(418)
	OptionalImplicitFieldDefsType    = SymbolId(419)
	ImplicitStructDefType            = SymbolId(420)
	ExplicitFieldDefsType            = SymbolId(421)
	OptionalExplicitFieldDefsType    = SymbolId(422)
	ExplicitStructDefType            = SymbolId(423)
	EnumValueDefType                 = SymbolId(424)
	ImplicitEnumValueDefsType        = SymbolId(425)
	ImplicitEnumDefType              = SymbolId(426)
	ExplicitEnumValueDefsType        = SymbolId(427)
	ExplicitEnumDefType              = SymbolId(428)
	TraitPropertyType                = SymbolId(429)
	TraitPropertiesType              = SymbolId(430)
	OptionalTraitPropertiesType      = SymbolId(431)
	TraitDefType                     = SymbolId(432)
	ReturnTypeType                   = SymbolId(433)
	ParameterDeclType                = SymbolId(434)
	ParameterDeclsType               = SymbolId(435)
	OptionalParameterDeclsType       = SymbolId(436)
	FuncTypeType                     = SymbolId(437)
	MethodSignatureType              = SymbolId(438)
	ParameterDefType                 = SymbolId(439)
	ParameterDefsType                = SymbolId(440)
	OptionalParameterDefsType        = SymbolId(441)
	NamedFuncDefType                 = SymbolId(442)
	AnonymousFuncExprType            = SymbolId(443)
	PackageDefType                   = SymbolId(444)
	PackageStatementBodyType         = SymbolId(445)
	PackageStatementType             = SymbolId(446)
	PackageStatementsType            = SymbolId(447)
	ImportStatementType              = SymbolId(448)
	ImportClauseType                 = SymbolId(449)
	ImportClauseTerminalType         = SymbolId(450)
	ImportClausesType                = SymbolId(451)
	LexInternalTokensType            = SymbolId(452)
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
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(127)
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
	_ReduceLexErrorToAtomType                                 = _ReduceType(187)
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
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(281)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(282)
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
	case _ReduceLexErrorToAtomExpr:
		return "LexErrorToAtomExpr"
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
	case _ReduceLexErrorToAtomType:
		return "LexErrorToAtomType"
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
	case _ReduceSpacesToLexInternalTokens:
		return "SpacesToLexInternalTokens"
	case _ReduceCommentToLexInternalTokens:
		return "CommentToLexInternalTokens"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceLexErrorToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.LexErrorToAtomExpr(args[0].Generic_)
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
	case _ReduceLexErrorToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.LexErrorToAtomType(args[0].Generic_)
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
	case _ReduceSpacesToLexInternalTokens:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LexInternalTokensType
		symbol.Generic_, err = reducer.SpacesToLexInternalTokens(args[0].Generic_)
	case _ReduceCommentToLexInternalTokens:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LexInternalTokensType
		symbol.Generic_, err = reducer.CommentToLexInternalTokens(args[0].Generic_)
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
	_ReduceLexErrorToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomExpr}
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
	_ReduceLexErrorToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomType}
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
	_ReduceSpacesToLexInternalTokensAction                          = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction                         = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State7, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State11, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State12, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State1, NewlinesToken}:                      _GotoState13Action,
	{_State1, SourceType}:                         _GotoState7Action,
	{_State1, OptionalDefinitionsType}:            _GotoState14Action,
	{_State1, OptionalNewlinesType}:               _GotoState15Action,
	{_State2, PackageToken}:                       _GotoState16Action,
	{_State2, PackageDefType}:                     _GotoState8Action,
	{_State3, TypeToken}:                          _GotoState17Action,
	{_State3, TypeDefType}:                        _GotoState9Action,
	{_State4, FuncToken}:                          _GotoState18Action,
	{_State4, NamedFuncDefType}:                   _GotoState10Action,
	{_State5, IntegerLiteralToken}:                _GotoState26Action,
	{_State5, FloatLiteralToken}:                  _GotoState22Action,
	{_State5, RuneLiteralToken}:                   _GotoState34Action,
	{_State5, StringLiteralToken}:                 _GotoState35Action,
	{_State5, IdentifierToken}:                    _GotoState25Action,
	{_State5, TrueToken}:                          _GotoState38Action,
	{_State5, FalseToken}:                         _GotoState21Action,
	{_State5, StructToken}:                        _GotoState36Action,
	{_State5, FuncToken}:                          _GotoState23Action,
	{_State5, VarToken}:                           _GotoState39Action,
	{_State5, LetToken}:                           _GotoState29Action,
	{_State5, LabelDeclToken}:                     _GotoState27Action,
	{_State5, LparenToken}:                        _GotoState31Action,
	{_State5, LbracketToken}:                      _GotoState28Action,
	{_State5, NotToken}:                           _GotoState33Action,
	{_State5, SubToken}:                           _GotoState37Action,
	{_State5, MulToken}:                           _GotoState32Action,
	{_State5, BitNegToken}:                        _GotoState20Action,
	{_State5, BitAndToken}:                        _GotoState19Action,
	{_State5, GreaterToken}:                       _GotoState24Action,
	{_State5, LexErrorToken}:                      _GotoState30Action,
	{_State5, VarDeclPatternType}:                 _GotoState59Action,
	{_State5, VarOrLetType}:                       _GotoState60Action,
	{_State5, ExpressionType}:                     _GotoState11Action,
	{_State5, OptionalLabelDeclType}:              _GotoState53Action,
	{_State5, SequenceExprType}:                   _GotoState58Action,
	{_State5, BlockExprType}:                      _GotoState45Action,
	{_State5, CallExprType}:                       _GotoState46Action,
	{_State5, AtomExprType}:                       _GotoState44Action,
	{_State5, LiteralType}:                        _GotoState51Action,
	{_State5, ImplicitStructExprType}:             _GotoState49Action,
	{_State5, AccessExprType}:                     _GotoState40Action,
	{_State5, PostfixUnaryExprType}:               _GotoState55Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState57Action,
	{_State5, PrefixUnaryExprType}:                _GotoState56Action,
	{_State5, MulExprType}:                        _GotoState52Action,
	{_State5, AddExprType}:                        _GotoState41Action,
	{_State5, CmpExprType}:                        _GotoState47Action,
	{_State5, AndExprType}:                        _GotoState42Action,
	{_State5, OrExprType}:                         _GotoState54Action,
	{_State5, InitializableTypeType}:              _GotoState50Action,
	{_State5, ExplicitStructDefType}:              _GotoState48Action,
	{_State5, AnonymousFuncExprType}:              _GotoState43Action,
	{_State6, SpacesToken}:                        _GotoState62Action,
	{_State6, CommentToken}:                       _GotoState61Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State15, PackageToken}:                      _GotoState16Action,
	{_State15, TypeToken}:                         _GotoState17Action,
	{_State15, FuncToken}:                         _GotoState18Action,
	{_State15, VarToken}:                          _GotoState39Action,
	{_State15, LetToken}:                          _GotoState29Action,
	{_State15, LbraceToken}:                       _GotoState63Action,
	{_State15, DefinitionsType}:                   _GotoState66Action,
	{_State15, DefinitionType}:                    _GotoState65Action,
	{_State15, VarDeclPatternType}:                _GotoState70Action,
	{_State15, VarOrLetType}:                      _GotoState60Action,
	{_State15, BlockBodyType}:                     _GotoState64Action,
	{_State15, TypeDefType}:                       _GotoState69Action,
	{_State15, NamedFuncDefType}:                  _GotoState67Action,
	{_State15, PackageDefType}:                    _GotoState68Action,
	{_State16, LparenToken}:                       _GotoState71Action,
	{_State17, IdentifierToken}:                   _GotoState72Action,
	{_State18, IdentifierToken}:                   _GotoState73Action,
	{_State18, LparenToken}:                       _GotoState74Action,
	{_State23, LparenToken}:                       _GotoState75Action,
	{_State24, IntegerLiteralToken}:               _GotoState26Action,
	{_State24, FloatLiteralToken}:                 _GotoState22Action,
	{_State24, RuneLiteralToken}:                  _GotoState34Action,
	{_State24, StringLiteralToken}:                _GotoState35Action,
	{_State24, IdentifierToken}:                   _GotoState25Action,
	{_State24, TrueToken}:                         _GotoState38Action,
	{_State24, FalseToken}:                        _GotoState21Action,
	{_State24, StructToken}:                       _GotoState36Action,
	{_State24, FuncToken}:                         _GotoState23Action,
	{_State24, VarToken}:                          _GotoState39Action,
	{_State24, LetToken}:                          _GotoState29Action,
	{_State24, LabelDeclToken}:                    _GotoState27Action,
	{_State24, LparenToken}:                       _GotoState31Action,
	{_State24, LbracketToken}:                     _GotoState28Action,
	{_State24, NotToken}:                          _GotoState33Action,
	{_State24, SubToken}:                          _GotoState37Action,
	{_State24, MulToken}:                          _GotoState32Action,
	{_State24, BitNegToken}:                       _GotoState20Action,
	{_State24, BitAndToken}:                       _GotoState19Action,
	{_State24, GreaterToken}:                      _GotoState24Action,
	{_State24, LexErrorToken}:                     _GotoState30Action,
	{_State24, VarDeclPatternType}:                _GotoState59Action,
	{_State24, VarOrLetType}:                      _GotoState60Action,
	{_State24, OptionalLabelDeclType}:             _GotoState76Action,
	{_State24, SequenceExprType}:                  _GotoState77Action,
	{_State24, BlockExprType}:                     _GotoState45Action,
	{_State24, CallExprType}:                      _GotoState46Action,
	{_State24, AtomExprType}:                      _GotoState44Action,
	{_State24, LiteralType}:                       _GotoState51Action,
	{_State24, ImplicitStructExprType}:            _GotoState49Action,
	{_State24, AccessExprType}:                    _GotoState40Action,
	{_State24, PostfixUnaryExprType}:              _GotoState55Action,
	{_State24, PrefixUnaryOpType}:                 _GotoState57Action,
	{_State24, PrefixUnaryExprType}:               _GotoState56Action,
	{_State24, MulExprType}:                       _GotoState52Action,
	{_State24, AddExprType}:                       _GotoState41Action,
	{_State24, CmpExprType}:                       _GotoState47Action,
	{_State24, AndExprType}:                       _GotoState42Action,
	{_State24, OrExprType}:                        _GotoState54Action,
	{_State24, InitializableTypeType}:             _GotoState50Action,
	{_State24, ExplicitStructDefType}:             _GotoState48Action,
	{_State24, AnonymousFuncExprType}:             _GotoState43Action,
	{_State28, IdentifierToken}:                   _GotoState84Action,
	{_State28, StructToken}:                       _GotoState36Action,
	{_State28, EnumToken}:                         _GotoState81Action,
	{_State28, TraitToken}:                        _GotoState89Action,
	{_State28, FuncToken}:                         _GotoState83Action,
	{_State28, LparenToken}:                       _GotoState86Action,
	{_State28, LbracketToken}:                     _GotoState28Action,
	{_State28, DotToken}:                          _GotoState80Action,
	{_State28, QuestionToken}:                     _GotoState87Action,
	{_State28, ExclaimToken}:                      _GotoState82Action,
	{_State28, TildeTildeToken}:                   _GotoState88Action,
	{_State28, BitNegToken}:                       _GotoState79Action,
	{_State28, BitAndToken}:                       _GotoState78Action,
	{_State28, LexErrorToken}:                     _GotoState85Action,
	{_State28, InitializableTypeType}:             _GotoState95Action,
	{_State28, AtomTypeType}:                      _GotoState90Action,
	{_State28, ReturnableTypeType}:                _GotoState96Action,
	{_State28, ValueTypeType}:                     _GotoState98Action,
	{_State28, ImplicitStructDefType}:             _GotoState94Action,
	{_State28, ExplicitStructDefType}:             _GotoState48Action,
	{_State28, ImplicitEnumDefType}:               _GotoState93Action,
	{_State28, ExplicitEnumDefType}:               _GotoState91Action,
	{_State28, TraitDefType}:                      _GotoState97Action,
	{_State28, FuncTypeType}:                      _GotoState92Action,
	{_State31, IntegerLiteralToken}:               _GotoState26Action,
	{_State31, FloatLiteralToken}:                 _GotoState22Action,
	{_State31, RuneLiteralToken}:                  _GotoState34Action,
	{_State31, StringLiteralToken}:                _GotoState35Action,
	{_State31, IdentifierToken}:                   _GotoState100Action,
	{_State31, TrueToken}:                         _GotoState38Action,
	{_State31, FalseToken}:                        _GotoState21Action,
	{_State31, StructToken}:                       _GotoState36Action,
	{_State31, FuncToken}:                         _GotoState23Action,
	{_State31, VarToken}:                          _GotoState39Action,
	{_State31, LetToken}:                          _GotoState29Action,
	{_State31, LabelDeclToken}:                    _GotoState27Action,
	{_State31, LparenToken}:                       _GotoState31Action,
	{_State31, LbracketToken}:                     _GotoState28Action,
	{_State31, DotDotDotToken}:                    _GotoState99Action,
	{_State31, NotToken}:                          _GotoState33Action,
	{_State31, SubToken}:                          _GotoState37Action,
	{_State31, MulToken}:                          _GotoState32Action,
	{_State31, BitNegToken}:                       _GotoState20Action,
	{_State31, BitAndToken}:                       _GotoState19Action,
	{_State31, GreaterToken}:                      _GotoState24Action,
	{_State31, LexErrorToken}:                     _GotoState30Action,
	{_State31, VarDeclPatternType}:                _GotoState59Action,
	{_State31, VarOrLetType}:                      _GotoState60Action,
	{_State31, ExpressionType}:                    _GotoState104Action,
	{_State31, OptionalLabelDeclType}:             _GotoState53Action,
	{_State31, SequenceExprType}:                  _GotoState58Action,
	{_State31, BlockExprType}:                     _GotoState45Action,
	{_State31, CallExprType}:                      _GotoState46Action,
	{_State31, ArgumentsType}:                     _GotoState102Action,
	{_State31, ArgumentType}:                      _GotoState101Action,
	{_State31, ColonExpressionsType}:              _GotoState103Action,
	{_State31, OptionalExpressionType}:            _GotoState105Action,
	{_State31, AtomExprType}:                      _GotoState44Action,
	{_State31, LiteralType}:                       _GotoState51Action,
	{_State31, ImplicitStructExprType}:            _GotoState49Action,
	{_State31, AccessExprType}:                    _GotoState40Action,
	{_State31, PostfixUnaryExprType}:              _GotoState55Action,
	{_State31, PrefixUnaryOpType}:                 _GotoState57Action,
	{_State31, PrefixUnaryExprType}:               _GotoState56Action,
	{_State31, MulExprType}:                       _GotoState52Action,
	{_State31, AddExprType}:                       _GotoState41Action,
	{_State31, CmpExprType}:                       _GotoState47Action,
	{_State31, AndExprType}:                       _GotoState42Action,
	{_State31, OrExprType}:                        _GotoState54Action,
	{_State31, InitializableTypeType}:             _GotoState50Action,
	{_State31, ExplicitStructDefType}:             _GotoState48Action,
	{_State31, AnonymousFuncExprType}:             _GotoState43Action,
	{_State36, LparenToken}:                       _GotoState106Action,
	{_State40, LbracketToken}:                     _GotoState109Action,
	{_State40, DotToken}:                          _GotoState108Action,
	{_State40, QuestionToken}:                     _GotoState110Action,
	{_State40, DollarLbracketToken}:               _GotoState107Action,
	{_State40, OptionalGenericBindingType}:        _GotoState111Action,
	{_State41, AddToken}:                          _GotoState112Action,
	{_State41, SubToken}:                          _GotoState115Action,
	{_State41, BitXorToken}:                       _GotoState114Action,
	{_State41, BitOrToken}:                        _GotoState113Action,
	{_State41, AddOpType}:                         _GotoState116Action,
	{_State42, AndToken}:                          _GotoState117Action,
	{_State47, EqualToken}:                        _GotoState118Action,
	{_State47, NotEqualToken}:                     _GotoState123Action,
	{_State47, LessToken}:                         _GotoState121Action,
	{_State47, LessOrEqualToken}:                  _GotoState122Action,
	{_State47, GreaterToken}:                      _GotoState119Action,
	{_State47, GreaterOrEqualToken}:               _GotoState120Action,
	{_State47, CmpOpType}:                         _GotoState124Action,
	{_State50, LparenToken}:                       _GotoState125Action,
	{_State52, MulToken}:                          _GotoState131Action,
	{_State52, DivToken}:                          _GotoState129Action,
	{_State52, ModToken}:                          _GotoState130Action,
	{_State52, BitAndToken}:                       _GotoState126Action,
	{_State52, BitLshiftToken}:                    _GotoState127Action,
	{_State52, BitRshiftToken}:                    _GotoState128Action,
	{_State52, MulOpType}:                         _GotoState132Action,
	{_State53, IfToken}:                           _GotoState135Action,
	{_State53, SwitchToken}:                       _GotoState136Action,
	{_State53, ForToken}:                          _GotoState134Action,
	{_State53, DoToken}:                           _GotoState133Action,
	{_State53, LbraceToken}:                       _GotoState63Action,
	{_State53, IfExprType}:                        _GotoState138Action,
	{_State53, SwitchExprType}:                    _GotoState140Action,
	{_State53, LoopExprType}:                      _GotoState139Action,
	{_State53, BlockBodyType}:                     _GotoState137Action,
	{_State54, OrToken}:                           _GotoState141Action,
	{_State57, IntegerLiteralToken}:               _GotoState26Action,
	{_State57, FloatLiteralToken}:                 _GotoState22Action,
	{_State57, RuneLiteralToken}:                  _GotoState34Action,
	{_State57, StringLiteralToken}:                _GotoState35Action,
	{_State57, IdentifierToken}:                   _GotoState25Action,
	{_State57, TrueToken}:                         _GotoState38Action,
	{_State57, FalseToken}:                        _GotoState21Action,
	{_State57, StructToken}:                       _GotoState36Action,
	{_State57, FuncToken}:                         _GotoState23Action,
	{_State57, LabelDeclToken}:                    _GotoState27Action,
	{_State57, LparenToken}:                       _GotoState31Action,
	{_State57, LbracketToken}:                     _GotoState28Action,
	{_State57, NotToken}:                          _GotoState33Action,
	{_State57, SubToken}:                          _GotoState37Action,
	{_State57, MulToken}:                          _GotoState32Action,
	{_State57, BitNegToken}:                       _GotoState20Action,
	{_State57, BitAndToken}:                       _GotoState19Action,
	{_State57, LexErrorToken}:                     _GotoState30Action,
	{_State57, OptionalLabelDeclType}:             _GotoState76Action,
	{_State57, BlockExprType}:                     _GotoState45Action,
	{_State57, CallExprType}:                      _GotoState46Action,
	{_State57, AtomExprType}:                      _GotoState44Action,
	{_State57, LiteralType}:                       _GotoState51Action,
	{_State57, ImplicitStructExprType}:            _GotoState49Action,
	{_State57, AccessExprType}:                    _GotoState40Action,
	{_State57, PostfixUnaryExprType}:              _GotoState55Action,
	{_State57, PrefixUnaryOpType}:                 _GotoState57Action,
	{_State57, PrefixUnaryExprType}:               _GotoState142Action,
	{_State57, InitializableTypeType}:             _GotoState50Action,
	{_State57, ExplicitStructDefType}:             _GotoState48Action,
	{_State57, AnonymousFuncExprType}:             _GotoState43Action,
	{_State60, IdentifierToken}:                   _GotoState143Action,
	{_State60, LparenToken}:                       _GotoState144Action,
	{_State60, VarPatternType}:                    _GotoState146Action,
	{_State60, TuplePatternType}:                  _GotoState145Action,
	{_State63, StatementsType}:                    _GotoState147Action,
	{_State66, NewlinesToken}:                     _GotoState148Action,
	{_State66, OptionalNewlinesType}:              _GotoState149Action,
	{_State70, AssignToken}:                       _GotoState150Action,
	{_State71, PackageStatementsType}:             _GotoState151Action,
	{_State72, DollarLbracketToken}:               _GotoState153Action,
	{_State72, AssignToken}:                       _GotoState152Action,
	{_State72, OptionalGenericParametersType}:     _GotoState154Action,
	{_State73, DollarLbracketToken}:               _GotoState153Action,
	{_State73, AssignToken}:                       _GotoState155Action,
	{_State73, OptionalGenericParametersType}:     _GotoState156Action,
	{_State74, IdentifierToken}:                   _GotoState157Action,
	{_State74, ParameterDefType}:                  _GotoState158Action,
	{_State75, IdentifierToken}:                   _GotoState157Action,
	{_State75, ParameterDefType}:                  _GotoState160Action,
	{_State75, ParameterDefsType}:                 _GotoState161Action,
	{_State75, OptionalParameterDefsType}:         _GotoState159Action,
	{_State76, LbraceToken}:                       _GotoState63Action,
	{_State76, BlockBodyType}:                     _GotoState137Action,
	{_State78, IdentifierToken}:                   _GotoState84Action,
	{_State78, StructToken}:                       _GotoState36Action,
	{_State78, EnumToken}:                         _GotoState81Action,
	{_State78, TraitToken}:                        _GotoState89Action,
	{_State78, FuncToken}:                         _GotoState83Action,
	{_State78, LparenToken}:                       _GotoState86Action,
	{_State78, LbracketToken}:                     _GotoState28Action,
	{_State78, DotToken}:                          _GotoState80Action,
	{_State78, QuestionToken}:                     _GotoState87Action,
	{_State78, ExclaimToken}:                      _GotoState82Action,
	{_State78, TildeTildeToken}:                   _GotoState88Action,
	{_State78, BitNegToken}:                       _GotoState79Action,
	{_State78, BitAndToken}:                       _GotoState78Action,
	{_State78, LexErrorToken}:                     _GotoState85Action,
	{_State78, InitializableTypeType}:             _GotoState95Action,
	{_State78, AtomTypeType}:                      _GotoState90Action,
	{_State78, ReturnableTypeType}:                _GotoState162Action,
	{_State78, ImplicitStructDefType}:             _GotoState94Action,
	{_State78, ExplicitStructDefType}:             _GotoState48Action,
	{_State78, ImplicitEnumDefType}:               _GotoState93Action,
	{_State78, ExplicitEnumDefType}:               _GotoState91Action,
	{_State78, TraitDefType}:                      _GotoState97Action,
	{_State78, FuncTypeType}:                      _GotoState92Action,
	{_State79, IdentifierToken}:                   _GotoState84Action,
	{_State79, StructToken}:                       _GotoState36Action,
	{_State79, EnumToken}:                         _GotoState81Action,
	{_State79, TraitToken}:                        _GotoState89Action,
	{_State79, FuncToken}:                         _GotoState83Action,
	{_State79, LparenToken}:                       _GotoState86Action,
	{_State79, LbracketToken}:                     _GotoState28Action,
	{_State79, DotToken}:                          _GotoState80Action,
	{_State79, QuestionToken}:                     _GotoState87Action,
	{_State79, ExclaimToken}:                      _GotoState82Action,
	{_State79, TildeTildeToken}:                   _GotoState88Action,
	{_State79, BitNegToken}:                       _GotoState79Action,
	{_State79, BitAndToken}:                       _GotoState78Action,
	{_State79, LexErrorToken}:                     _GotoState85Action,
	{_State79, InitializableTypeType}:             _GotoState95Action,
	{_State79, AtomTypeType}:                      _GotoState90Action,
	{_State79, ReturnableTypeType}:                _GotoState163Action,
	{_State79, ImplicitStructDefType}:             _GotoState94Action,
	{_State79, ExplicitStructDefType}:             _GotoState48Action,
	{_State79, ImplicitEnumDefType}:               _GotoState93Action,
	{_State79, ExplicitEnumDefType}:               _GotoState91Action,
	{_State79, TraitDefType}:                      _GotoState97Action,
	{_State79, FuncTypeType}:                      _GotoState92Action,
	{_State80, DollarLbracketToken}:               _GotoState107Action,
	{_State80, OptionalGenericBindingType}:        _GotoState164Action,
	{_State81, LparenToken}:                       _GotoState165Action,
	{_State82, IdentifierToken}:                   _GotoState84Action,
	{_State82, StructToken}:                       _GotoState36Action,
	{_State82, EnumToken}:                         _GotoState81Action,
	{_State82, TraitToken}:                        _GotoState89Action,
	{_State82, FuncToken}:                         _GotoState83Action,
	{_State82, LparenToken}:                       _GotoState86Action,
	{_State82, LbracketToken}:                     _GotoState28Action,
	{_State82, DotToken}:                          _GotoState80Action,
	{_State82, QuestionToken}:                     _GotoState87Action,
	{_State82, ExclaimToken}:                      _GotoState82Action,
	{_State82, TildeTildeToken}:                   _GotoState88Action,
	{_State82, BitNegToken}:                       _GotoState79Action,
	{_State82, BitAndToken}:                       _GotoState78Action,
	{_State82, LexErrorToken}:                     _GotoState85Action,
	{_State82, InitializableTypeType}:             _GotoState95Action,
	{_State82, AtomTypeType}:                      _GotoState90Action,
	{_State82, ReturnableTypeType}:                _GotoState166Action,
	{_State82, ImplicitStructDefType}:             _GotoState94Action,
	{_State82, ExplicitStructDefType}:             _GotoState48Action,
	{_State82, ImplicitEnumDefType}:               _GotoState93Action,
	{_State82, ExplicitEnumDefType}:               _GotoState91Action,
	{_State82, TraitDefType}:                      _GotoState97Action,
	{_State82, FuncTypeType}:                      _GotoState92Action,
	{_State83, LparenToken}:                       _GotoState167Action,
	{_State84, DotToken}:                          _GotoState168Action,
	{_State84, DollarLbracketToken}:               _GotoState107Action,
	{_State84, OptionalGenericBindingType}:        _GotoState169Action,
	{_State86, IdentifierToken}:                   _GotoState170Action,
	{_State86, UnsafeToken}:                       _GotoState171Action,
	{_State86, StructToken}:                       _GotoState36Action,
	{_State86, EnumToken}:                         _GotoState81Action,
	{_State86, TraitToken}:                        _GotoState89Action,
	{_State86, FuncToken}:                         _GotoState83Action,
	{_State86, LparenToken}:                       _GotoState86Action,
	{_State86, LbracketToken}:                     _GotoState28Action,
	{_State86, DotToken}:                          _GotoState80Action,
	{_State86, QuestionToken}:                     _GotoState87Action,
	{_State86, ExclaimToken}:                      _GotoState82Action,
	{_State86, TildeTildeToken}:                   _GotoState88Action,
	{_State86, BitNegToken}:                       _GotoState79Action,
	{_State86, BitAndToken}:                       _GotoState78Action,
	{_State86, LexErrorToken}:                     _GotoState85Action,
	{_State86, UnsafeStatementType}:               _GotoState177Action,
	{_State86, InitializableTypeType}:             _GotoState95Action,
	{_State86, AtomTypeType}:                      _GotoState90Action,
	{_State86, ReturnableTypeType}:                _GotoState96Action,
	{_State86, ValueTypeType}:                     _GotoState178Action,
	{_State86, FieldDefType}:                      _GotoState173Action,
	{_State86, ImplicitFieldDefsType}:             _GotoState175Action,
	{_State86, OptionalImplicitFieldDefsType}:     _GotoState176Action,
	{_State86, ImplicitStructDefType}:             _GotoState94Action,
	{_State86, ExplicitStructDefType}:             _GotoState48Action,
	{_State86, EnumValueDefType}:                  _GotoState172Action,
	{_State86, ImplicitEnumValueDefsType}:         _GotoState174Action,
	{_State86, ImplicitEnumDefType}:               _GotoState93Action,
	{_State86, ExplicitEnumDefType}:               _GotoState91Action,
	{_State86, TraitDefType}:                      _GotoState97Action,
	{_State86, FuncTypeType}:                      _GotoState92Action,
	{_State87, IdentifierToken}:                   _GotoState84Action,
	{_State87, StructToken}:                       _GotoState36Action,
	{_State87, EnumToken}:                         _GotoState81Action,
	{_State87, TraitToken}:                        _GotoState89Action,
	{_State87, FuncToken}:                         _GotoState83Action,
	{_State87, LparenToken}:                       _GotoState86Action,
	{_State87, LbracketToken}:                     _GotoState28Action,
	{_State87, DotToken}:                          _GotoState80Action,
	{_State87, QuestionToken}:                     _GotoState87Action,
	{_State87, ExclaimToken}:                      _GotoState82Action,
	{_State87, TildeTildeToken}:                   _GotoState88Action,
	{_State87, BitNegToken}:                       _GotoState79Action,
	{_State87, BitAndToken}:                       _GotoState78Action,
	{_State87, LexErrorToken}:                     _GotoState85Action,
	{_State87, InitializableTypeType}:             _GotoState95Action,
	{_State87, AtomTypeType}:                      _GotoState90Action,
	{_State87, ReturnableTypeType}:                _GotoState179Action,
	{_State87, ImplicitStructDefType}:             _GotoState94Action,
	{_State87, ExplicitStructDefType}:             _GotoState48Action,
	{_State87, ImplicitEnumDefType}:               _GotoState93Action,
	{_State87, ExplicitEnumDefType}:               _GotoState91Action,
	{_State87, TraitDefType}:                      _GotoState97Action,
	{_State87, FuncTypeType}:                      _GotoState92Action,
	{_State88, IdentifierToken}:                   _GotoState84Action,
	{_State88, StructToken}:                       _GotoState36Action,
	{_State88, EnumToken}:                         _GotoState81Action,
	{_State88, TraitToken}:                        _GotoState89Action,
	{_State88, FuncToken}:                         _GotoState83Action,
	{_State88, LparenToken}:                       _GotoState86Action,
	{_State88, LbracketToken}:                     _GotoState28Action,
	{_State88, DotToken}:                          _GotoState80Action,
	{_State88, QuestionToken}:                     _GotoState87Action,
	{_State88, ExclaimToken}:                      _GotoState82Action,
	{_State88, TildeTildeToken}:                   _GotoState88Action,
	{_State88, BitNegToken}:                       _GotoState79Action,
	{_State88, BitAndToken}:                       _GotoState78Action,
	{_State88, LexErrorToken}:                     _GotoState85Action,
	{_State88, InitializableTypeType}:             _GotoState95Action,
	{_State88, AtomTypeType}:                      _GotoState90Action,
	{_State88, ReturnableTypeType}:                _GotoState180Action,
	{_State88, ImplicitStructDefType}:             _GotoState94Action,
	{_State88, ExplicitStructDefType}:             _GotoState48Action,
	{_State88, ImplicitEnumDefType}:               _GotoState93Action,
	{_State88, ExplicitEnumDefType}:               _GotoState91Action,
	{_State88, TraitDefType}:                      _GotoState97Action,
	{_State88, FuncTypeType}:                      _GotoState92Action,
	{_State89, LparenToken}:                       _GotoState181Action,
	{_State98, RbracketToken}:                     _GotoState186Action,
	{_State98, CommaToken}:                        _GotoState184Action,
	{_State98, ColonToken}:                        _GotoState183Action,
	{_State98, AddToken}:                          _GotoState182Action,
	{_State98, SubToken}:                          _GotoState187Action,
	{_State98, MulToken}:                          _GotoState185Action,
	{_State100, AssignToken}:                      _GotoState188Action,
	{_State102, RparenToken}:                      _GotoState190Action,
	{_State102, CommaToken}:                       _GotoState189Action,
	{_State103, ColonToken}:                       _GotoState191Action,
	{_State105, ColonToken}:                       _GotoState192Action,
	{_State106, IdentifierToken}:                  _GotoState170Action,
	{_State106, UnsafeToken}:                      _GotoState171Action,
	{_State106, StructToken}:                      _GotoState36Action,
	{_State106, EnumToken}:                        _GotoState81Action,
	{_State106, TraitToken}:                       _GotoState89Action,
	{_State106, FuncToken}:                        _GotoState83Action,
	{_State106, LparenToken}:                      _GotoState86Action,
	{_State106, LbracketToken}:                    _GotoState28Action,
	{_State106, DotToken}:                         _GotoState80Action,
	{_State106, QuestionToken}:                    _GotoState87Action,
	{_State106, ExclaimToken}:                     _GotoState82Action,
	{_State106, TildeTildeToken}:                  _GotoState88Action,
	{_State106, BitNegToken}:                      _GotoState79Action,
	{_State106, BitAndToken}:                      _GotoState78Action,
	{_State106, LexErrorToken}:                    _GotoState85Action,
	{_State106, UnsafeStatementType}:              _GotoState177Action,
	{_State106, InitializableTypeType}:            _GotoState95Action,
	{_State106, AtomTypeType}:                     _GotoState90Action,
	{_State106, ReturnableTypeType}:               _GotoState96Action,
	{_State106, ValueTypeType}:                    _GotoState178Action,
	{_State106, FieldDefType}:                     _GotoState194Action,
	{_State106, ImplicitStructDefType}:            _GotoState94Action,
	{_State106, ExplicitFieldDefsType}:            _GotoState193Action,
	{_State106, OptionalExplicitFieldDefsType}:    _GotoState195Action,
	{_State106, ExplicitStructDefType}:            _GotoState48Action,
	{_State106, ImplicitEnumDefType}:              _GotoState93Action,
	{_State106, ExplicitEnumDefType}:              _GotoState91Action,
	{_State106, TraitDefType}:                     _GotoState97Action,
	{_State106, FuncTypeType}:                     _GotoState92Action,
	{_State107, IdentifierToken}:                  _GotoState84Action,
	{_State107, StructToken}:                      _GotoState36Action,
	{_State107, EnumToken}:                        _GotoState81Action,
	{_State107, TraitToken}:                       _GotoState89Action,
	{_State107, FuncToken}:                        _GotoState83Action,
	{_State107, LparenToken}:                      _GotoState86Action,
	{_State107, LbracketToken}:                    _GotoState28Action,
	{_State107, DotToken}:                         _GotoState80Action,
	{_State107, QuestionToken}:                    _GotoState87Action,
	{_State107, ExclaimToken}:                     _GotoState82Action,
	{_State107, TildeTildeToken}:                  _GotoState88Action,
	{_State107, BitNegToken}:                      _GotoState79Action,
	{_State107, BitAndToken}:                      _GotoState78Action,
	{_State107, LexErrorToken}:                    _GotoState85Action,
	{_State107, OptionalGenericArgumentsType}:     _GotoState197Action,
	{_State107, GenericArgumentsType}:             _GotoState196Action,
	{_State107, InitializableTypeType}:            _GotoState95Action,
	{_State107, AtomTypeType}:                     _GotoState90Action,
	{_State107, ReturnableTypeType}:               _GotoState96Action,
	{_State107, ValueTypeType}:                    _GotoState198Action,
	{_State107, ImplicitStructDefType}:            _GotoState94Action,
	{_State107, ExplicitStructDefType}:            _GotoState48Action,
	{_State107, ImplicitEnumDefType}:              _GotoState93Action,
	{_State107, ExplicitEnumDefType}:              _GotoState91Action,
	{_State107, TraitDefType}:                     _GotoState97Action,
	{_State107, FuncTypeType}:                     _GotoState92Action,
	{_State108, IdentifierToken}:                  _GotoState199Action,
	{_State109, IntegerLiteralToken}:              _GotoState26Action,
	{_State109, FloatLiteralToken}:                _GotoState22Action,
	{_State109, RuneLiteralToken}:                 _GotoState34Action,
	{_State109, StringLiteralToken}:               _GotoState35Action,
	{_State109, IdentifierToken}:                  _GotoState100Action,
	{_State109, TrueToken}:                        _GotoState38Action,
	{_State109, FalseToken}:                       _GotoState21Action,
	{_State109, StructToken}:                      _GotoState36Action,
	{_State109, FuncToken}:                        _GotoState23Action,
	{_State109, VarToken}:                         _GotoState39Action,
	{_State109, LetToken}:                         _GotoState29Action,
	{_State109, LabelDeclToken}:                   _GotoState27Action,
	{_State109, LparenToken}:                      _GotoState31Action,
	{_State109, LbracketToken}:                    _GotoState28Action,
	{_State109, DotDotDotToken}:                   _GotoState99Action,
	{_State109, NotToken}:                         _GotoState33Action,
	{_State109, SubToken}:                         _GotoState37Action,
	{_State109, MulToken}:                         _GotoState32Action,
	{_State109, BitNegToken}:                      _GotoState20Action,
	{_State109, BitAndToken}:                      _GotoState19Action,
	{_State109, GreaterToken}:                     _GotoState24Action,
	{_State109, LexErrorToken}:                    _GotoState30Action,
	{_State109, VarDeclPatternType}:               _GotoState59Action,
	{_State109, VarOrLetType}:                     _GotoState60Action,
	{_State109, ExpressionType}:                   _GotoState104Action,
	{_State109, OptionalLabelDeclType}:            _GotoState53Action,
	{_State109, SequenceExprType}:                 _GotoState58Action,
	{_State109, BlockExprType}:                    _GotoState45Action,
	{_State109, CallExprType}:                     _GotoState46Action,
	{_State109, ArgumentType}:                     _GotoState200Action,
	{_State109, ColonExpressionsType}:             _GotoState103Action,
	{_State109, OptionalExpressionType}:           _GotoState105Action,
	{_State109, AtomExprType}:                     _GotoState44Action,
	{_State109, LiteralType}:                      _GotoState51Action,
	{_State109, ImplicitStructExprType}:           _GotoState49Action,
	{_State109, AccessExprType}:                   _GotoState40Action,
	{_State109, PostfixUnaryExprType}:             _GotoState55Action,
	{_State109, PrefixUnaryOpType}:                _GotoState57Action,
	{_State109, PrefixUnaryExprType}:              _GotoState56Action,
	{_State109, MulExprType}:                      _GotoState52Action,
	{_State109, AddExprType}:                      _GotoState41Action,
	{_State109, CmpExprType}:                      _GotoState47Action,
	{_State109, AndExprType}:                      _GotoState42Action,
	{_State109, OrExprType}:                       _GotoState54Action,
	{_State109, InitializableTypeType}:            _GotoState50Action,
	{_State109, ExplicitStructDefType}:            _GotoState48Action,
	{_State109, AnonymousFuncExprType}:            _GotoState43Action,
	{_State111, LparenToken}:                      _GotoState201Action,
	{_State116, IntegerLiteralToken}:              _GotoState26Action,
	{_State116, FloatLiteralToken}:                _GotoState22Action,
	{_State116, RuneLiteralToken}:                 _GotoState34Action,
	{_State116, StringLiteralToken}:               _GotoState35Action,
	{_State116, IdentifierToken}:                  _GotoState25Action,
	{_State116, TrueToken}:                        _GotoState38Action,
	{_State116, FalseToken}:                       _GotoState21Action,
	{_State116, StructToken}:                      _GotoState36Action,
	{_State116, FuncToken}:                        _GotoState23Action,
	{_State116, LabelDeclToken}:                   _GotoState27Action,
	{_State116, LparenToken}:                      _GotoState31Action,
	{_State116, LbracketToken}:                    _GotoState28Action,
	{_State116, NotToken}:                         _GotoState33Action,
	{_State116, SubToken}:                         _GotoState37Action,
	{_State116, MulToken}:                         _GotoState32Action,
	{_State116, BitNegToken}:                      _GotoState20Action,
	{_State116, BitAndToken}:                      _GotoState19Action,
	{_State116, LexErrorToken}:                    _GotoState30Action,
	{_State116, OptionalLabelDeclType}:            _GotoState76Action,
	{_State116, BlockExprType}:                    _GotoState45Action,
	{_State116, CallExprType}:                     _GotoState46Action,
	{_State116, AtomExprType}:                     _GotoState44Action,
	{_State116, LiteralType}:                      _GotoState51Action,
	{_State116, ImplicitStructExprType}:           _GotoState49Action,
	{_State116, AccessExprType}:                   _GotoState40Action,
	{_State116, PostfixUnaryExprType}:             _GotoState55Action,
	{_State116, PrefixUnaryOpType}:                _GotoState57Action,
	{_State116, PrefixUnaryExprType}:              _GotoState56Action,
	{_State116, MulExprType}:                      _GotoState202Action,
	{_State116, InitializableTypeType}:            _GotoState50Action,
	{_State116, ExplicitStructDefType}:            _GotoState48Action,
	{_State116, AnonymousFuncExprType}:            _GotoState43Action,
	{_State117, IntegerLiteralToken}:              _GotoState26Action,
	{_State117, FloatLiteralToken}:                _GotoState22Action,
	{_State117, RuneLiteralToken}:                 _GotoState34Action,
	{_State117, StringLiteralToken}:               _GotoState35Action,
	{_State117, IdentifierToken}:                  _GotoState25Action,
	{_State117, TrueToken}:                        _GotoState38Action,
	{_State117, FalseToken}:                       _GotoState21Action,
	{_State117, StructToken}:                      _GotoState36Action,
	{_State117, FuncToken}:                        _GotoState23Action,
	{_State117, LabelDeclToken}:                   _GotoState27Action,
	{_State117, LparenToken}:                      _GotoState31Action,
	{_State117, LbracketToken}:                    _GotoState28Action,
	{_State117, NotToken}:                         _GotoState33Action,
	{_State117, SubToken}:                         _GotoState37Action,
	{_State117, MulToken}:                         _GotoState32Action,
	{_State117, BitNegToken}:                      _GotoState20Action,
	{_State117, BitAndToken}:                      _GotoState19Action,
	{_State117, LexErrorToken}:                    _GotoState30Action,
	{_State117, OptionalLabelDeclType}:            _GotoState76Action,
	{_State117, BlockExprType}:                    _GotoState45Action,
	{_State117, CallExprType}:                     _GotoState46Action,
	{_State117, AtomExprType}:                     _GotoState44Action,
	{_State117, LiteralType}:                      _GotoState51Action,
	{_State117, ImplicitStructExprType}:           _GotoState49Action,
	{_State117, AccessExprType}:                   _GotoState40Action,
	{_State117, PostfixUnaryExprType}:             _GotoState55Action,
	{_State117, PrefixUnaryOpType}:                _GotoState57Action,
	{_State117, PrefixUnaryExprType}:              _GotoState56Action,
	{_State117, MulExprType}:                      _GotoState52Action,
	{_State117, AddExprType}:                      _GotoState41Action,
	{_State117, CmpExprType}:                      _GotoState203Action,
	{_State117, InitializableTypeType}:            _GotoState50Action,
	{_State117, ExplicitStructDefType}:            _GotoState48Action,
	{_State117, AnonymousFuncExprType}:            _GotoState43Action,
	{_State124, IntegerLiteralToken}:              _GotoState26Action,
	{_State124, FloatLiteralToken}:                _GotoState22Action,
	{_State124, RuneLiteralToken}:                 _GotoState34Action,
	{_State124, StringLiteralToken}:               _GotoState35Action,
	{_State124, IdentifierToken}:                  _GotoState25Action,
	{_State124, TrueToken}:                        _GotoState38Action,
	{_State124, FalseToken}:                       _GotoState21Action,
	{_State124, StructToken}:                      _GotoState36Action,
	{_State124, FuncToken}:                        _GotoState23Action,
	{_State124, LabelDeclToken}:                   _GotoState27Action,
	{_State124, LparenToken}:                      _GotoState31Action,
	{_State124, LbracketToken}:                    _GotoState28Action,
	{_State124, NotToken}:                         _GotoState33Action,
	{_State124, SubToken}:                         _GotoState37Action,
	{_State124, MulToken}:                         _GotoState32Action,
	{_State124, BitNegToken}:                      _GotoState20Action,
	{_State124, BitAndToken}:                      _GotoState19Action,
	{_State124, LexErrorToken}:                    _GotoState30Action,
	{_State124, OptionalLabelDeclType}:            _GotoState76Action,
	{_State124, BlockExprType}:                    _GotoState45Action,
	{_State124, CallExprType}:                     _GotoState46Action,
	{_State124, AtomExprType}:                     _GotoState44Action,
	{_State124, LiteralType}:                      _GotoState51Action,
	{_State124, ImplicitStructExprType}:           _GotoState49Action,
	{_State124, AccessExprType}:                   _GotoState40Action,
	{_State124, PostfixUnaryExprType}:             _GotoState55Action,
	{_State124, PrefixUnaryOpType}:                _GotoState57Action,
	{_State124, PrefixUnaryExprType}:              _GotoState56Action,
	{_State124, MulExprType}:                      _GotoState52Action,
	{_State124, AddExprType}:                      _GotoState204Action,
	{_State124, InitializableTypeType}:            _GotoState50Action,
	{_State124, ExplicitStructDefType}:            _GotoState48Action,
	{_State124, AnonymousFuncExprType}:            _GotoState43Action,
	{_State125, IntegerLiteralToken}:              _GotoState26Action,
	{_State125, FloatLiteralToken}:                _GotoState22Action,
	{_State125, RuneLiteralToken}:                 _GotoState34Action,
	{_State125, StringLiteralToken}:               _GotoState35Action,
	{_State125, IdentifierToken}:                  _GotoState100Action,
	{_State125, TrueToken}:                        _GotoState38Action,
	{_State125, FalseToken}:                       _GotoState21Action,
	{_State125, StructToken}:                      _GotoState36Action,
	{_State125, FuncToken}:                        _GotoState23Action,
	{_State125, VarToken}:                         _GotoState39Action,
	{_State125, LetToken}:                         _GotoState29Action,
	{_State125, LabelDeclToken}:                   _GotoState27Action,
	{_State125, LparenToken}:                      _GotoState31Action,
	{_State125, LbracketToken}:                    _GotoState28Action,
	{_State125, DotDotDotToken}:                   _GotoState99Action,
	{_State125, NotToken}:                         _GotoState33Action,
	{_State125, SubToken}:                         _GotoState37Action,
	{_State125, MulToken}:                         _GotoState32Action,
	{_State125, BitNegToken}:                      _GotoState20Action,
	{_State125, BitAndToken}:                      _GotoState19Action,
	{_State125, GreaterToken}:                     _GotoState24Action,
	{_State125, LexErrorToken}:                    _GotoState30Action,
	{_State125, VarDeclPatternType}:               _GotoState59Action,
	{_State125, VarOrLetType}:                     _GotoState60Action,
	{_State125, ExpressionType}:                   _GotoState104Action,
	{_State125, OptionalLabelDeclType}:            _GotoState53Action,
	{_State125, SequenceExprType}:                 _GotoState58Action,
	{_State125, BlockExprType}:                    _GotoState45Action,
	{_State125, CallExprType}:                     _GotoState46Action,
	{_State125, ArgumentsType}:                    _GotoState205Action,
	{_State125, ArgumentType}:                     _GotoState101Action,
	{_State125, ColonExpressionsType}:             _GotoState103Action,
	{_State125, OptionalExpressionType}:           _GotoState105Action,
	{_State125, AtomExprType}:                     _GotoState44Action,
	{_State125, LiteralType}:                      _GotoState51Action,
	{_State125, ImplicitStructExprType}:           _GotoState49Action,
	{_State125, AccessExprType}:                   _GotoState40Action,
	{_State125, PostfixUnaryExprType}:             _GotoState55Action,
	{_State125, PrefixUnaryOpType}:                _GotoState57Action,
	{_State125, PrefixUnaryExprType}:              _GotoState56Action,
	{_State125, MulExprType}:                      _GotoState52Action,
	{_State125, AddExprType}:                      _GotoState41Action,
	{_State125, CmpExprType}:                      _GotoState47Action,
	{_State125, AndExprType}:                      _GotoState42Action,
	{_State125, OrExprType}:                       _GotoState54Action,
	{_State125, InitializableTypeType}:            _GotoState50Action,
	{_State125, ExplicitStructDefType}:            _GotoState48Action,
	{_State125, AnonymousFuncExprType}:            _GotoState43Action,
	{_State132, IntegerLiteralToken}:              _GotoState26Action,
	{_State132, FloatLiteralToken}:                _GotoState22Action,
	{_State132, RuneLiteralToken}:                 _GotoState34Action,
	{_State132, StringLiteralToken}:               _GotoState35Action,
	{_State132, IdentifierToken}:                  _GotoState25Action,
	{_State132, TrueToken}:                        _GotoState38Action,
	{_State132, FalseToken}:                       _GotoState21Action,
	{_State132, StructToken}:                      _GotoState36Action,
	{_State132, FuncToken}:                        _GotoState23Action,
	{_State132, LabelDeclToken}:                   _GotoState27Action,
	{_State132, LparenToken}:                      _GotoState31Action,
	{_State132, LbracketToken}:                    _GotoState28Action,
	{_State132, NotToken}:                         _GotoState33Action,
	{_State132, SubToken}:                         _GotoState37Action,
	{_State132, MulToken}:                         _GotoState32Action,
	{_State132, BitNegToken}:                      _GotoState20Action,
	{_State132, BitAndToken}:                      _GotoState19Action,
	{_State132, LexErrorToken}:                    _GotoState30Action,
	{_State132, OptionalLabelDeclType}:            _GotoState76Action,
	{_State132, BlockExprType}:                    _GotoState45Action,
	{_State132, CallExprType}:                     _GotoState46Action,
	{_State132, AtomExprType}:                     _GotoState44Action,
	{_State132, LiteralType}:                      _GotoState51Action,
	{_State132, ImplicitStructExprType}:           _GotoState49Action,
	{_State132, AccessExprType}:                   _GotoState40Action,
	{_State132, PostfixUnaryExprType}:             _GotoState55Action,
	{_State132, PrefixUnaryOpType}:                _GotoState57Action,
	{_State132, PrefixUnaryExprType}:              _GotoState206Action,
	{_State132, InitializableTypeType}:            _GotoState50Action,
	{_State132, ExplicitStructDefType}:            _GotoState48Action,
	{_State132, AnonymousFuncExprType}:            _GotoState43Action,
	{_State133, LbraceToken}:                      _GotoState63Action,
	{_State133, BlockBodyType}:                    _GotoState207Action,
	{_State134, IntegerLiteralToken}:              _GotoState26Action,
	{_State134, FloatLiteralToken}:                _GotoState22Action,
	{_State134, RuneLiteralToken}:                 _GotoState34Action,
	{_State134, StringLiteralToken}:               _GotoState35Action,
	{_State134, IdentifierToken}:                  _GotoState25Action,
	{_State134, TrueToken}:                        _GotoState38Action,
	{_State134, FalseToken}:                       _GotoState21Action,
	{_State134, StructToken}:                      _GotoState36Action,
	{_State134, FuncToken}:                        _GotoState23Action,
	{_State134, VarToken}:                         _GotoState39Action,
	{_State134, LetToken}:                         _GotoState29Action,
	{_State134, LabelDeclToken}:                   _GotoState27Action,
	{_State134, LparenToken}:                      _GotoState31Action,
	{_State134, LbracketToken}:                    _GotoState28Action,
	{_State134, NotToken}:                         _GotoState33Action,
	{_State134, SubToken}:                         _GotoState37Action,
	{_State134, MulToken}:                         _GotoState32Action,
	{_State134, BitNegToken}:                      _GotoState20Action,
	{_State134, BitAndToken}:                      _GotoState19Action,
	{_State134, GreaterToken}:                     _GotoState24Action,
	{_State134, LexErrorToken}:                    _GotoState30Action,
	{_State134, VarDeclPatternType}:               _GotoState59Action,
	{_State134, VarOrLetType}:                     _GotoState60Action,
	{_State134, AssignPatternType}:                _GotoState208Action,
	{_State134, OptionalLabelDeclType}:            _GotoState76Action,
	{_State134, ForAssignmentType}:                _GotoState209Action,
	{_State134, SequenceExprType}:                 _GotoState210Action,
	{_State134, BlockExprType}:                    _GotoState45Action,
	{_State134, CallExprType}:                     _GotoState46Action,
	{_State134, AtomExprType}:                     _GotoState44Action,
	{_State134, LiteralType}:                      _GotoState51Action,
	{_State134, ImplicitStructExprType}:           _GotoState49Action,
	{_State134, AccessExprType}:                   _GotoState40Action,
	{_State134, PostfixUnaryExprType}:             _GotoState55Action,
	{_State134, PrefixUnaryOpType}:                _GotoState57Action,
	{_State134, PrefixUnaryExprType}:              _GotoState56Action,
	{_State134, MulExprType}:                      _GotoState52Action,
	{_State134, AddExprType}:                      _GotoState41Action,
	{_State134, CmpExprType}:                      _GotoState47Action,
	{_State134, AndExprType}:                      _GotoState42Action,
	{_State134, OrExprType}:                       _GotoState54Action,
	{_State134, InitializableTypeType}:            _GotoState50Action,
	{_State134, ExplicitStructDefType}:            _GotoState48Action,
	{_State134, AnonymousFuncExprType}:            _GotoState43Action,
	{_State135, IntegerLiteralToken}:              _GotoState26Action,
	{_State135, FloatLiteralToken}:                _GotoState22Action,
	{_State135, RuneLiteralToken}:                 _GotoState34Action,
	{_State135, StringLiteralToken}:               _GotoState35Action,
	{_State135, IdentifierToken}:                  _GotoState25Action,
	{_State135, TrueToken}:                        _GotoState38Action,
	{_State135, FalseToken}:                       _GotoState21Action,
	{_State135, CaseToken}:                        _GotoState211Action,
	{_State135, StructToken}:                      _GotoState36Action,
	{_State135, FuncToken}:                        _GotoState23Action,
	{_State135, VarToken}:                         _GotoState39Action,
	{_State135, LetToken}:                         _GotoState29Action,
	{_State135, LabelDeclToken}:                   _GotoState27Action,
	{_State135, LparenToken}:                      _GotoState31Action,
	{_State135, LbracketToken}:                    _GotoState28Action,
	{_State135, NotToken}:                         _GotoState33Action,
	{_State135, SubToken}:                         _GotoState37Action,
	{_State135, MulToken}:                         _GotoState32Action,
	{_State135, BitNegToken}:                      _GotoState20Action,
	{_State135, BitAndToken}:                      _GotoState19Action,
	{_State135, GreaterToken}:                     _GotoState24Action,
	{_State135, LexErrorToken}:                    _GotoState30Action,
	{_State135, VarDeclPatternType}:               _GotoState59Action,
	{_State135, VarOrLetType}:                     _GotoState60Action,
	{_State135, OptionalLabelDeclType}:            _GotoState76Action,
	{_State135, ConditionType}:                    _GotoState212Action,
	{_State135, SequenceExprType}:                 _GotoState213Action,
	{_State135, BlockExprType}:                    _GotoState45Action,
	{_State135, CallExprType}:                     _GotoState46Action,
	{_State135, AtomExprType}:                     _GotoState44Action,
	{_State135, LiteralType}:                      _GotoState51Action,
	{_State135, ImplicitStructExprType}:           _GotoState49Action,
	{_State135, AccessExprType}:                   _GotoState40Action,
	{_State135, PostfixUnaryExprType}:             _GotoState55Action,
	{_State135, PrefixUnaryOpType}:                _GotoState57Action,
	{_State135, PrefixUnaryExprType}:              _GotoState56Action,
	{_State135, MulExprType}:                      _GotoState52Action,
	{_State135, AddExprType}:                      _GotoState41Action,
	{_State135, CmpExprType}:                      _GotoState47Action,
	{_State135, AndExprType}:                      _GotoState42Action,
	{_State135, OrExprType}:                       _GotoState54Action,
	{_State135, InitializableTypeType}:            _GotoState50Action,
	{_State135, ExplicitStructDefType}:            _GotoState48Action,
	{_State135, AnonymousFuncExprType}:            _GotoState43Action,
	{_State136, IntegerLiteralToken}:              _GotoState26Action,
	{_State136, FloatLiteralToken}:                _GotoState22Action,
	{_State136, RuneLiteralToken}:                 _GotoState34Action,
	{_State136, StringLiteralToken}:               _GotoState35Action,
	{_State136, IdentifierToken}:                  _GotoState25Action,
	{_State136, TrueToken}:                        _GotoState38Action,
	{_State136, FalseToken}:                       _GotoState21Action,
	{_State136, StructToken}:                      _GotoState36Action,
	{_State136, FuncToken}:                        _GotoState23Action,
	{_State136, VarToken}:                         _GotoState39Action,
	{_State136, LetToken}:                         _GotoState29Action,
	{_State136, LabelDeclToken}:                   _GotoState27Action,
	{_State136, LparenToken}:                      _GotoState31Action,
	{_State136, LbracketToken}:                    _GotoState28Action,
	{_State136, NotToken}:                         _GotoState33Action,
	{_State136, SubToken}:                         _GotoState37Action,
	{_State136, MulToken}:                         _GotoState32Action,
	{_State136, BitNegToken}:                      _GotoState20Action,
	{_State136, BitAndToken}:                      _GotoState19Action,
	{_State136, GreaterToken}:                     _GotoState24Action,
	{_State136, LexErrorToken}:                    _GotoState30Action,
	{_State136, VarDeclPatternType}:               _GotoState59Action,
	{_State136, VarOrLetType}:                     _GotoState60Action,
	{_State136, OptionalLabelDeclType}:            _GotoState76Action,
	{_State136, SequenceExprType}:                 _GotoState214Action,
	{_State136, BlockExprType}:                    _GotoState45Action,
	{_State136, CallExprType}:                     _GotoState46Action,
	{_State136, AtomExprType}:                     _GotoState44Action,
	{_State136, LiteralType}:                      _GotoState51Action,
	{_State136, ImplicitStructExprType}:           _GotoState49Action,
	{_State136, AccessExprType}:                   _GotoState40Action,
	{_State136, PostfixUnaryExprType}:             _GotoState55Action,
	{_State136, PrefixUnaryOpType}:                _GotoState57Action,
	{_State136, PrefixUnaryExprType}:              _GotoState56Action,
	{_State136, MulExprType}:                      _GotoState52Action,
	{_State136, AddExprType}:                      _GotoState41Action,
	{_State136, CmpExprType}:                      _GotoState47Action,
	{_State136, AndExprType}:                      _GotoState42Action,
	{_State136, OrExprType}:                       _GotoState54Action,
	{_State136, InitializableTypeType}:            _GotoState50Action,
	{_State136, ExplicitStructDefType}:            _GotoState48Action,
	{_State136, AnonymousFuncExprType}:            _GotoState43Action,
	{_State141, IntegerLiteralToken}:              _GotoState26Action,
	{_State141, FloatLiteralToken}:                _GotoState22Action,
	{_State141, RuneLiteralToken}:                 _GotoState34Action,
	{_State141, StringLiteralToken}:               _GotoState35Action,
	{_State141, IdentifierToken}:                  _GotoState25Action,
	{_State141, TrueToken}:                        _GotoState38Action,
	{_State141, FalseToken}:                       _GotoState21Action,
	{_State141, StructToken}:                      _GotoState36Action,
	{_State141, FuncToken}:                        _GotoState23Action,
	{_State141, LabelDeclToken}:                   _GotoState27Action,
	{_State141, LparenToken}:                      _GotoState31Action,
	{_State141, LbracketToken}:                    _GotoState28Action,
	{_State141, NotToken}:                         _GotoState33Action,
	{_State141, SubToken}:                         _GotoState37Action,
	{_State141, MulToken}:                         _GotoState32Action,
	{_State141, BitNegToken}:                      _GotoState20Action,
	{_State141, BitAndToken}:                      _GotoState19Action,
	{_State141, LexErrorToken}:                    _GotoState30Action,
	{_State141, OptionalLabelDeclType}:            _GotoState76Action,
	{_State141, BlockExprType}:                    _GotoState45Action,
	{_State141, CallExprType}:                     _GotoState46Action,
	{_State141, AtomExprType}:                     _GotoState44Action,
	{_State141, LiteralType}:                      _GotoState51Action,
	{_State141, ImplicitStructExprType}:           _GotoState49Action,
	{_State141, AccessExprType}:                   _GotoState40Action,
	{_State141, PostfixUnaryExprType}:             _GotoState55Action,
	{_State141, PrefixUnaryOpType}:                _GotoState57Action,
	{_State141, PrefixUnaryExprType}:              _GotoState56Action,
	{_State141, MulExprType}:                      _GotoState52Action,
	{_State141, AddExprType}:                      _GotoState41Action,
	{_State141, CmpExprType}:                      _GotoState47Action,
	{_State141, AndExprType}:                      _GotoState215Action,
	{_State141, InitializableTypeType}:            _GotoState50Action,
	{_State141, ExplicitStructDefType}:            _GotoState48Action,
	{_State141, AnonymousFuncExprType}:            _GotoState43Action,
	{_State144, IdentifierToken}:                  _GotoState217Action,
	{_State144, LparenToken}:                      _GotoState144Action,
	{_State144, DotDotDotToken}:                   _GotoState216Action,
	{_State144, VarPatternType}:                   _GotoState220Action,
	{_State144, TuplePatternType}:                 _GotoState145Action,
	{_State144, FieldVarPatternsType}:             _GotoState219Action,
	{_State144, FieldVarPatternType}:              _GotoState218Action,
	{_State146, IdentifierToken}:                  _GotoState84Action,
	{_State146, StructToken}:                      _GotoState36Action,
	{_State146, EnumToken}:                        _GotoState81Action,
	{_State146, TraitToken}:                       _GotoState89Action,
	{_State146, FuncToken}:                        _GotoState83Action,
	{_State146, LparenToken}:                      _GotoState86Action,
	{_State146, LbracketToken}:                    _GotoState28Action,
	{_State146, DotToken}:                         _GotoState80Action,
	{_State146, QuestionToken}:                    _GotoState87Action,
	{_State146, ExclaimToken}:                     _GotoState82Action,
	{_State146, TildeTildeToken}:                  _GotoState88Action,
	{_State146, BitNegToken}:                      _GotoState79Action,
	{_State146, BitAndToken}:                      _GotoState78Action,
	{_State146, LexErrorToken}:                    _GotoState85Action,
	{_State146, OptionalValueTypeType}:            _GotoState221Action,
	{_State146, InitializableTypeType}:            _GotoState95Action,
	{_State146, AtomTypeType}:                     _GotoState90Action,
	{_State146, ReturnableTypeType}:               _GotoState96Action,
	{_State146, ValueTypeType}:                    _GotoState222Action,
	{_State146, ImplicitStructDefType}:            _GotoState94Action,
	{_State146, ExplicitStructDefType}:            _GotoState48Action,
	{_State146, ImplicitEnumDefType}:              _GotoState93Action,
	{_State146, ExplicitEnumDefType}:              _GotoState91Action,
	{_State146, TraitDefType}:                     _GotoState97Action,
	{_State146, FuncTypeType}:                     _GotoState92Action,
	{_State147, IntegerLiteralToken}:              _GotoState26Action,
	{_State147, FloatLiteralToken}:                _GotoState22Action,
	{_State147, RuneLiteralToken}:                 _GotoState34Action,
	{_State147, StringLiteralToken}:               _GotoState35Action,
	{_State147, IdentifierToken}:                  _GotoState25Action,
	{_State147, TrueToken}:                        _GotoState38Action,
	{_State147, FalseToken}:                       _GotoState21Action,
	{_State147, ReturnToken}:                      _GotoState228Action,
	{_State147, BreakToken}:                       _GotoState224Action,
	{_State147, ContinueToken}:                    _GotoState225Action,
	{_State147, UnsafeToken}:                      _GotoState171Action,
	{_State147, StructToken}:                      _GotoState36Action,
	{_State147, FuncToken}:                        _GotoState23Action,
	{_State147, AsyncToken}:                       _GotoState223Action,
	{_State147, DeferToken}:                       _GotoState226Action,
	{_State147, VarToken}:                         _GotoState39Action,
	{_State147, LetToken}:                         _GotoState29Action,
	{_State147, LabelDeclToken}:                   _GotoState27Action,
	{_State147, RbraceToken}:                      _GotoState227Action,
	{_State147, LparenToken}:                      _GotoState31Action,
	{_State147, LbracketToken}:                    _GotoState28Action,
	{_State147, NotToken}:                         _GotoState33Action,
	{_State147, SubToken}:                         _GotoState37Action,
	{_State147, MulToken}:                         _GotoState32Action,
	{_State147, BitNegToken}:                      _GotoState20Action,
	{_State147, BitAndToken}:                      _GotoState19Action,
	{_State147, GreaterToken}:                     _GotoState24Action,
	{_State147, LexErrorToken}:                    _GotoState30Action,
	{_State147, VarDeclPatternType}:               _GotoState59Action,
	{_State147, VarOrLetType}:                     _GotoState60Action,
	{_State147, AssignPatternType}:                _GotoState230Action,
	{_State147, ExpressionType}:                   _GotoState231Action,
	{_State147, OptionalLabelDeclType}:            _GotoState53Action,
	{_State147, SequenceExprType}:                 _GotoState235Action,
	{_State147, BlockExprType}:                    _GotoState45Action,
	{_State147, StatementType}:                    _GotoState236Action,
	{_State147, StatementBodyType}:                _GotoState237Action,
	{_State147, UnsafeStatementType}:              _GotoState238Action,
	{_State147, JumpStatementType}:                _GotoState233Action,
	{_State147, JumpTypeType}:                     _GotoState234Action,
	{_State147, ExpressionsType}:                  _GotoState232Action,
	{_State147, CallExprType}:                     _GotoState46Action,
	{_State147, AtomExprType}:                     _GotoState44Action,
	{_State147, LiteralType}:                      _GotoState51Action,
	{_State147, ImplicitStructExprType}:           _GotoState49Action,
	{_State147, AccessExprType}:                   _GotoState229Action,
	{_State147, PostfixUnaryExprType}:             _GotoState55Action,
	{_State147, PrefixUnaryOpType}:                _GotoState57Action,
	{_State147, PrefixUnaryExprType}:              _GotoState56Action,
	{_State147, MulExprType}:                      _GotoState52Action,
	{_State147, AddExprType}:                      _GotoState41Action,
	{_State147, CmpExprType}:                      _GotoState47Action,
	{_State147, AndExprType}:                      _GotoState42Action,
	{_State147, OrExprType}:                       _GotoState54Action,
	{_State147, InitializableTypeType}:            _GotoState50Action,
	{_State147, ExplicitStructDefType}:            _GotoState48Action,
	{_State147, AnonymousFuncExprType}:            _GotoState43Action,
	{_State148, PackageToken}:                     _GotoState16Action,
	{_State148, TypeToken}:                        _GotoState17Action,
	{_State148, FuncToken}:                        _GotoState18Action,
	{_State148, VarToken}:                         _GotoState39Action,
	{_State148, LetToken}:                         _GotoState29Action,
	{_State148, LbraceToken}:                      _GotoState63Action,
	{_State148, DefinitionType}:                   _GotoState239Action,
	{_State148, VarDeclPatternType}:               _GotoState70Action,
	{_State148, VarOrLetType}:                     _GotoState60Action,
	{_State148, BlockBodyType}:                    _GotoState64Action,
	{_State148, TypeDefType}:                      _GotoState69Action,
	{_State148, NamedFuncDefType}:                 _GotoState67Action,
	{_State148, PackageDefType}:                   _GotoState68Action,
	{_State150, IntegerLiteralToken}:              _GotoState26Action,
	{_State150, FloatLiteralToken}:                _GotoState22Action,
	{_State150, RuneLiteralToken}:                 _GotoState34Action,
	{_State150, StringLiteralToken}:               _GotoState35Action,
	{_State150, IdentifierToken}:                  _GotoState25Action,
	{_State150, TrueToken}:                        _GotoState38Action,
	{_State150, FalseToken}:                       _GotoState21Action,
	{_State150, StructToken}:                      _GotoState36Action,
	{_State150, FuncToken}:                        _GotoState23Action,
	{_State150, VarToken}:                         _GotoState39Action,
	{_State150, LetToken}:                         _GotoState29Action,
	{_State150, LabelDeclToken}:                   _GotoState27Action,
	{_State150, LparenToken}:                      _GotoState31Action,
	{_State150, LbracketToken}:                    _GotoState28Action,
	{_State150, NotToken}:                         _GotoState33Action,
	{_State150, SubToken}:                         _GotoState37Action,
	{_State150, MulToken}:                         _GotoState32Action,
	{_State150, BitNegToken}:                      _GotoState20Action,
	{_State150, BitAndToken}:                      _GotoState19Action,
	{_State150, GreaterToken}:                     _GotoState24Action,
	{_State150, LexErrorToken}:                    _GotoState30Action,
	{_State150, VarDeclPatternType}:               _GotoState59Action,
	{_State150, VarOrLetType}:                     _GotoState60Action,
	{_State150, ExpressionType}:                   _GotoState240Action,
	{_State150, OptionalLabelDeclType}:            _GotoState53Action,
	{_State150, SequenceExprType}:                 _GotoState58Action,
	{_State150, BlockExprType}:                    _GotoState45Action,
	{_State150, CallExprType}:                     _GotoState46Action,
	{_State150, AtomExprType}:                     _GotoState44Action,
	{_State150, LiteralType}:                      _GotoState51Action,
	{_State150, ImplicitStructExprType}:           _GotoState49Action,
	{_State150, AccessExprType}:                   _GotoState40Action,
	{_State150, PostfixUnaryExprType}:             _GotoState55Action,
	{_State150, PrefixUnaryOpType}:                _GotoState57Action,
	{_State150, PrefixUnaryExprType}:              _GotoState56Action,
	{_State150, MulExprType}:                      _GotoState52Action,
	{_State150, AddExprType}:                      _GotoState41Action,
	{_State150, CmpExprType}:                      _GotoState47Action,
	{_State150, AndExprType}:                      _GotoState42Action,
	{_State150, OrExprType}:                       _GotoState54Action,
	{_State150, InitializableTypeType}:            _GotoState50Action,
	{_State150, ExplicitStructDefType}:            _GotoState48Action,
	{_State150, AnonymousFuncExprType}:            _GotoState43Action,
	{_State151, ImportToken}:                      _GotoState241Action,
	{_State151, UnsafeToken}:                      _GotoState171Action,
	{_State151, RparenToken}:                      _GotoState242Action,
	{_State151, UnsafeStatementType}:              _GotoState246Action,
	{_State151, PackageStatementBodyType}:         _GotoState245Action,
	{_State151, PackageStatementType}:             _GotoState244Action,
	{_State151, ImportStatementType}:              _GotoState243Action,
	{_State152, IdentifierToken}:                  _GotoState84Action,
	{_State152, StructToken}:                      _GotoState36Action,
	{_State152, EnumToken}:                        _GotoState81Action,
	{_State152, TraitToken}:                       _GotoState89Action,
	{_State152, FuncToken}:                        _GotoState83Action,
	{_State152, LparenToken}:                      _GotoState86Action,
	{_State152, LbracketToken}:                    _GotoState28Action,
	{_State152, DotToken}:                         _GotoState80Action,
	{_State152, QuestionToken}:                    _GotoState87Action,
	{_State152, ExclaimToken}:                     _GotoState82Action,
	{_State152, TildeTildeToken}:                  _GotoState88Action,
	{_State152, BitNegToken}:                      _GotoState79Action,
	{_State152, BitAndToken}:                      _GotoState78Action,
	{_State152, LexErrorToken}:                    _GotoState85Action,
	{_State152, InitializableTypeType}:            _GotoState95Action,
	{_State152, AtomTypeType}:                     _GotoState90Action,
	{_State152, ReturnableTypeType}:               _GotoState96Action,
	{_State152, ValueTypeType}:                    _GotoState247Action,
	{_State152, ImplicitStructDefType}:            _GotoState94Action,
	{_State152, ExplicitStructDefType}:            _GotoState48Action,
	{_State152, ImplicitEnumDefType}:              _GotoState93Action,
	{_State152, ExplicitEnumDefType}:              _GotoState91Action,
	{_State152, TraitDefType}:                     _GotoState97Action,
	{_State152, FuncTypeType}:                     _GotoState92Action,
	{_State153, IdentifierToken}:                  _GotoState248Action,
	{_State153, GenericParameterDefType}:          _GotoState249Action,
	{_State153, GenericParameterDefsType}:         _GotoState250Action,
	{_State153, OptionalGenericParameterDefsType}: _GotoState251Action,
	{_State154, IdentifierToken}:                  _GotoState84Action,
	{_State154, StructToken}:                      _GotoState36Action,
	{_State154, EnumToken}:                        _GotoState81Action,
	{_State154, TraitToken}:                       _GotoState89Action,
	{_State154, FuncToken}:                        _GotoState83Action,
	{_State154, LparenToken}:                      _GotoState86Action,
	{_State154, LbracketToken}:                    _GotoState28Action,
	{_State154, DotToken}:                         _GotoState80Action,
	{_State154, QuestionToken}:                    _GotoState87Action,
	{_State154, ExclaimToken}:                     _GotoState82Action,
	{_State154, TildeTildeToken}:                  _GotoState88Action,
	{_State154, BitNegToken}:                      _GotoState79Action,
	{_State154, BitAndToken}:                      _GotoState78Action,
	{_State154, LexErrorToken}:                    _GotoState85Action,
	{_State154, InitializableTypeType}:            _GotoState95Action,
	{_State154, AtomTypeType}:                     _GotoState90Action,
	{_State154, ReturnableTypeType}:               _GotoState96Action,
	{_State154, ValueTypeType}:                    _GotoState252Action,
	{_State154, ImplicitStructDefType}:            _GotoState94Action,
	{_State154, ExplicitStructDefType}:            _GotoState48Action,
	{_State154, ImplicitEnumDefType}:              _GotoState93Action,
	{_State154, ExplicitEnumDefType}:              _GotoState91Action,
	{_State154, TraitDefType}:                     _GotoState97Action,
	{_State154, FuncTypeType}:                     _GotoState92Action,
	{_State155, IntegerLiteralToken}:              _GotoState26Action,
	{_State155, FloatLiteralToken}:                _GotoState22Action,
	{_State155, RuneLiteralToken}:                 _GotoState34Action,
	{_State155, StringLiteralToken}:               _GotoState35Action,
	{_State155, IdentifierToken}:                  _GotoState25Action,
	{_State155, TrueToken}:                        _GotoState38Action,
	{_State155, FalseToken}:                       _GotoState21Action,
	{_State155, StructToken}:                      _GotoState36Action,
	{_State155, FuncToken}:                        _GotoState23Action,
	{_State155, VarToken}:                         _GotoState39Action,
	{_State155, LetToken}:                         _GotoState29Action,
	{_State155, LabelDeclToken}:                   _GotoState27Action,
	{_State155, LparenToken}:                      _GotoState31Action,
	{_State155, LbracketToken}:                    _GotoState28Action,
	{_State155, NotToken}:                         _GotoState33Action,
	{_State155, SubToken}:                         _GotoState37Action,
	{_State155, MulToken}:                         _GotoState32Action,
	{_State155, BitNegToken}:                      _GotoState20Action,
	{_State155, BitAndToken}:                      _GotoState19Action,
	{_State155, GreaterToken}:                     _GotoState24Action,
	{_State155, LexErrorToken}:                    _GotoState30Action,
	{_State155, VarDeclPatternType}:               _GotoState59Action,
	{_State155, VarOrLetType}:                     _GotoState60Action,
	{_State155, ExpressionType}:                   _GotoState253Action,
	{_State155, OptionalLabelDeclType}:            _GotoState53Action,
	{_State155, SequenceExprType}:                 _GotoState58Action,
	{_State155, BlockExprType}:                    _GotoState45Action,
	{_State155, CallExprType}:                     _GotoState46Action,
	{_State155, AtomExprType}:                     _GotoState44Action,
	{_State155, LiteralType}:                      _GotoState51Action,
	{_State155, ImplicitStructExprType}:           _GotoState49Action,
	{_State155, AccessExprType}:                   _GotoState40Action,
	{_State155, PostfixUnaryExprType}:             _GotoState55Action,
	{_State155, PrefixUnaryOpType}:                _GotoState57Action,
	{_State155, PrefixUnaryExprType}:              _GotoState56Action,
	{_State155, MulExprType}:                      _GotoState52Action,
	{_State155, AddExprType}:                      _GotoState41Action,
	{_State155, CmpExprType}:                      _GotoState47Action,
	{_State155, AndExprType}:                      _GotoState42Action,
	{_State155, OrExprType}:                       _GotoState54Action,
	{_State155, InitializableTypeType}:            _GotoState50Action,
	{_State155, ExplicitStructDefType}:            _GotoState48Action,
	{_State155, AnonymousFuncExprType}:            _GotoState43Action,
	{_State156, LparenToken}:                      _GotoState254Action,
	{_State157, IdentifierToken}:                  _GotoState84Action,
	{_State157, StructToken}:                      _GotoState36Action,
	{_State157, EnumToken}:                        _GotoState81Action,
	{_State157, TraitToken}:                       _GotoState89Action,
	{_State157, FuncToken}:                        _GotoState83Action,
	{_State157, LparenToken}:                      _GotoState86Action,
	{_State157, LbracketToken}:                    _GotoState28Action,
	{_State157, DotToken}:                         _GotoState80Action,
	{_State157, QuestionToken}:                    _GotoState87Action,
	{_State157, ExclaimToken}:                     _GotoState82Action,
	{_State157, DotDotDotToken}:                   _GotoState255Action,
	{_State157, TildeTildeToken}:                  _GotoState88Action,
	{_State157, BitNegToken}:                      _GotoState79Action,
	{_State157, BitAndToken}:                      _GotoState78Action,
	{_State157, LexErrorToken}:                    _GotoState85Action,
	{_State157, InitializableTypeType}:            _GotoState95Action,
	{_State157, AtomTypeType}:                     _GotoState90Action,
	{_State157, ReturnableTypeType}:               _GotoState96Action,
	{_State157, ValueTypeType}:                    _GotoState256Action,
	{_State157, ImplicitStructDefType}:            _GotoState94Action,
	{_State157, ExplicitStructDefType}:            _GotoState48Action,
	{_State157, ImplicitEnumDefType}:              _GotoState93Action,
	{_State157, ExplicitEnumDefType}:              _GotoState91Action,
	{_State157, TraitDefType}:                     _GotoState97Action,
	{_State157, FuncTypeType}:                     _GotoState92Action,
	{_State158, RparenToken}:                      _GotoState257Action,
	{_State159, RparenToken}:                      _GotoState258Action,
	{_State161, CommaToken}:                       _GotoState259Action,
	{_State165, IdentifierToken}:                  _GotoState170Action,
	{_State165, UnsafeToken}:                      _GotoState171Action,
	{_State165, StructToken}:                      _GotoState36Action,
	{_State165, EnumToken}:                        _GotoState81Action,
	{_State165, TraitToken}:                       _GotoState89Action,
	{_State165, FuncToken}:                        _GotoState83Action,
	{_State165, LparenToken}:                      _GotoState86Action,
	{_State165, LbracketToken}:                    _GotoState28Action,
	{_State165, DotToken}:                         _GotoState80Action,
	{_State165, QuestionToken}:                    _GotoState87Action,
	{_State165, ExclaimToken}:                     _GotoState82Action,
	{_State165, TildeTildeToken}:                  _GotoState88Action,
	{_State165, BitNegToken}:                      _GotoState79Action,
	{_State165, BitAndToken}:                      _GotoState78Action,
	{_State165, LexErrorToken}:                    _GotoState85Action,
	{_State165, UnsafeStatementType}:              _GotoState177Action,
	{_State165, InitializableTypeType}:            _GotoState95Action,
	{_State165, AtomTypeType}:                     _GotoState90Action,
	{_State165, ReturnableTypeType}:               _GotoState96Action,
	{_State165, ValueTypeType}:                    _GotoState178Action,
	{_State165, FieldDefType}:                     _GotoState262Action,
	{_State165, ImplicitStructDefType}:            _GotoState94Action,
	{_State165, ExplicitStructDefType}:            _GotoState48Action,
	{_State165, EnumValueDefType}:                 _GotoState260Action,
	{_State165, ImplicitEnumValueDefsType}:        _GotoState263Action,
	{_State165, ImplicitEnumDefType}:              _GotoState93Action,
	{_State165, ExplicitEnumValueDefsType}:        _GotoState261Action,
	{_State165, ExplicitEnumDefType}:              _GotoState91Action,
	{_State165, TraitDefType}:                     _GotoState97Action,
	{_State165, FuncTypeType}:                     _GotoState92Action,
	{_State167, IdentifierToken}:                  _GotoState265Action,
	{_State167, StructToken}:                      _GotoState36Action,
	{_State167, EnumToken}:                        _GotoState81Action,
	{_State167, TraitToken}:                       _GotoState89Action,
	{_State167, FuncToken}:                        _GotoState83Action,
	{_State167, LparenToken}:                      _GotoState86Action,
	{_State167, LbracketToken}:                    _GotoState28Action,
	{_State167, DotToken}:                         _GotoState80Action,
	{_State167, QuestionToken}:                    _GotoState87Action,
	{_State167, ExclaimToken}:                     _GotoState82Action,
	{_State167, DotDotDotToken}:                   _GotoState264Action,
	{_State167, TildeTildeToken}:                  _GotoState88Action,
	{_State167, BitNegToken}:                      _GotoState79Action,
	{_State167, BitAndToken}:                      _GotoState78Action,
	{_State167, LexErrorToken}:                    _GotoState85Action,
	{_State167, InitializableTypeType}:            _GotoState95Action,
	{_State167, AtomTypeType}:                     _GotoState90Action,
	{_State167, ReturnableTypeType}:               _GotoState96Action,
	{_State167, ValueTypeType}:                    _GotoState269Action,
	{_State167, ImplicitStructDefType}:            _GotoState94Action,
	{_State167, ExplicitStructDefType}:            _GotoState48Action,
	{_State167, ImplicitEnumDefType}:              _GotoState93Action,
	{_State167, ExplicitEnumDefType}:              _GotoState91Action,
	{_State167, TraitDefType}:                     _GotoState97Action,
	{_State167, ParameterDeclType}:                _GotoState267Action,
	{_State167, ParameterDeclsType}:               _GotoState268Action,
	{_State167, OptionalParameterDeclsType}:       _GotoState266Action,
	{_State167, FuncTypeType}:                     _GotoState92Action,
	{_State168, IdentifierToken}:                  _GotoState270Action,
	{_State170, IdentifierToken}:                  _GotoState84Action,
	{_State170, StructToken}:                      _GotoState36Action,
	{_State170, EnumToken}:                        _GotoState81Action,
	{_State170, TraitToken}:                       _GotoState89Action,
	{_State170, FuncToken}:                        _GotoState83Action,
	{_State170, LparenToken}:                      _GotoState86Action,
	{_State170, LbracketToken}:                    _GotoState28Action,
	{_State170, DotToken}:                         _GotoState271Action,
	{_State170, QuestionToken}:                    _GotoState87Action,
	{_State170, ExclaimToken}:                     _GotoState82Action,
	{_State170, DollarLbracketToken}:              _GotoState107Action,
	{_State170, TildeTildeToken}:                  _GotoState88Action,
	{_State170, BitNegToken}:                      _GotoState79Action,
	{_State170, BitAndToken}:                      _GotoState78Action,
	{_State170, LexErrorToken}:                    _GotoState85Action,
	{_State170, OptionalGenericBindingType}:       _GotoState169Action,
	{_State170, InitializableTypeType}:            _GotoState95Action,
	{_State170, AtomTypeType}:                     _GotoState90Action,
	{_State170, ReturnableTypeType}:               _GotoState96Action,
	{_State170, ValueTypeType}:                    _GotoState272Action,
	{_State170, ImplicitStructDefType}:            _GotoState94Action,
	{_State170, ExplicitStructDefType}:            _GotoState48Action,
	{_State170, ImplicitEnumDefType}:              _GotoState93Action,
	{_State170, ExplicitEnumDefType}:              _GotoState91Action,
	{_State170, TraitDefType}:                     _GotoState97Action,
	{_State170, FuncTypeType}:                     _GotoState92Action,
	{_State171, LessToken}:                        _GotoState273Action,
	{_State172, OrToken}:                          _GotoState274Action,
	{_State173, AssignToken}:                      _GotoState275Action,
	{_State174, RparenToken}:                      _GotoState277Action,
	{_State174, OrToken}:                          _GotoState276Action,
	{_State175, CommaToken}:                       _GotoState278Action,
	{_State176, RparenToken}:                      _GotoState279Action,
	{_State178, AddToken}:                         _GotoState182Action,
	{_State178, SubToken}:                         _GotoState187Action,
	{_State178, MulToken}:                         _GotoState185Action,
	{_State181, IdentifierToken}:                  _GotoState170Action,
	{_State181, UnsafeToken}:                      _GotoState171Action,
	{_State181, StructToken}:                      _GotoState36Action,
	{_State181, EnumToken}:                        _GotoState81Action,
	{_State181, TraitToken}:                       _GotoState89Action,
	{_State181, FuncToken}:                        _GotoState280Action,
	{_State181, LparenToken}:                      _GotoState86Action,
	{_State181, LbracketToken}:                    _GotoState28Action,
	{_State181, DotToken}:                         _GotoState80Action,
	{_State181, QuestionToken}:                    _GotoState87Action,
	{_State181, ExclaimToken}:                     _GotoState82Action,
	{_State181, TildeTildeToken}:                  _GotoState88Action,
	{_State181, BitNegToken}:                      _GotoState79Action,
	{_State181, BitAndToken}:                      _GotoState78Action,
	{_State181, LexErrorToken}:                    _GotoState85Action,
	{_State181, UnsafeStatementType}:              _GotoState177Action,
	{_State181, InitializableTypeType}:            _GotoState95Action,
	{_State181, AtomTypeType}:                     _GotoState90Action,
	{_State181, ReturnableTypeType}:               _GotoState96Action,
	{_State181, ValueTypeType}:                    _GotoState178Action,
	{_State181, FieldDefType}:                     _GotoState281Action,
	{_State181, ImplicitStructDefType}:            _GotoState94Action,
	{_State181, ExplicitStructDefType}:            _GotoState48Action,
	{_State181, ImplicitEnumDefType}:              _GotoState93Action,
	{_State181, ExplicitEnumDefType}:              _GotoState91Action,
	{_State181, TraitPropertyType}:                _GotoState285Action,
	{_State181, TraitPropertiesType}:              _GotoState284Action,
	{_State181, OptionalTraitPropertiesType}:      _GotoState283Action,
	{_State181, TraitDefType}:                     _GotoState97Action,
	{_State181, FuncTypeType}:                     _GotoState92Action,
	{_State181, MethodSignatureType}:              _GotoState282Action,
	{_State182, IdentifierToken}:                  _GotoState84Action,
	{_State182, StructToken}:                      _GotoState36Action,
	{_State182, EnumToken}:                        _GotoState81Action,
	{_State182, TraitToken}:                       _GotoState89Action,
	{_State182, FuncToken}:                        _GotoState83Action,
	{_State182, LparenToken}:                      _GotoState86Action,
	{_State182, LbracketToken}:                    _GotoState28Action,
	{_State182, DotToken}:                         _GotoState80Action,
	{_State182, QuestionToken}:                    _GotoState87Action,
	{_State182, ExclaimToken}:                     _GotoState82Action,
	{_State182, TildeTildeToken}:                  _GotoState88Action,
	{_State182, BitNegToken}:                      _GotoState79Action,
	{_State182, BitAndToken}:                      _GotoState78Action,
	{_State182, LexErrorToken}:                    _GotoState85Action,
	{_State182, InitializableTypeType}:            _GotoState95Action,
	{_State182, AtomTypeType}:                     _GotoState90Action,
	{_State182, ReturnableTypeType}:               _GotoState286Action,
	{_State182, ImplicitStructDefType}:            _GotoState94Action,
	{_State182, ExplicitStructDefType}:            _GotoState48Action,
	{_State182, ImplicitEnumDefType}:              _GotoState93Action,
	{_State182, ExplicitEnumDefType}:              _GotoState91Action,
	{_State182, TraitDefType}:                     _GotoState97Action,
	{_State182, FuncTypeType}:                     _GotoState92Action,
	{_State183, IdentifierToken}:                  _GotoState84Action,
	{_State183, StructToken}:                      _GotoState36Action,
	{_State183, EnumToken}:                        _GotoState81Action,
	{_State183, TraitToken}:                       _GotoState89Action,
	{_State183, FuncToken}:                        _GotoState83Action,
	{_State183, LparenToken}:                      _GotoState86Action,
	{_State183, LbracketToken}:                    _GotoState28Action,
	{_State183, DotToken}:                         _GotoState80Action,
	{_State183, QuestionToken}:                    _GotoState87Action,
	{_State183, ExclaimToken}:                     _GotoState82Action,
	{_State183, TildeTildeToken}:                  _GotoState88Action,
	{_State183, BitNegToken}:                      _GotoState79Action,
	{_State183, BitAndToken}:                      _GotoState78Action,
	{_State183, LexErrorToken}:                    _GotoState85Action,
	{_State183, InitializableTypeType}:            _GotoState95Action,
	{_State183, AtomTypeType}:                     _GotoState90Action,
	{_State183, ReturnableTypeType}:               _GotoState96Action,
	{_State183, ValueTypeType}:                    _GotoState287Action,
	{_State183, ImplicitStructDefType}:            _GotoState94Action,
	{_State183, ExplicitStructDefType}:            _GotoState48Action,
	{_State183, ImplicitEnumDefType}:              _GotoState93Action,
	{_State183, ExplicitEnumDefType}:              _GotoState91Action,
	{_State183, TraitDefType}:                     _GotoState97Action,
	{_State183, FuncTypeType}:                     _GotoState92Action,
	{_State184, IntegerLiteralToken}:              _GotoState288Action,
	{_State185, IdentifierToken}:                  _GotoState84Action,
	{_State185, StructToken}:                      _GotoState36Action,
	{_State185, EnumToken}:                        _GotoState81Action,
	{_State185, TraitToken}:                       _GotoState89Action,
	{_State185, FuncToken}:                        _GotoState83Action,
	{_State185, LparenToken}:                      _GotoState86Action,
	{_State185, LbracketToken}:                    _GotoState28Action,
	{_State185, DotToken}:                         _GotoState80Action,
	{_State185, QuestionToken}:                    _GotoState87Action,
	{_State185, ExclaimToken}:                     _GotoState82Action,
	{_State185, TildeTildeToken}:                  _GotoState88Action,
	{_State185, BitNegToken}:                      _GotoState79Action,
	{_State185, BitAndToken}:                      _GotoState78Action,
	{_State185, LexErrorToken}:                    _GotoState85Action,
	{_State185, InitializableTypeType}:            _GotoState95Action,
	{_State185, AtomTypeType}:                     _GotoState90Action,
	{_State185, ReturnableTypeType}:               _GotoState289Action,
	{_State185, ImplicitStructDefType}:            _GotoState94Action,
	{_State185, ExplicitStructDefType}:            _GotoState48Action,
	{_State185, ImplicitEnumDefType}:              _GotoState93Action,
	{_State185, ExplicitEnumDefType}:              _GotoState91Action,
	{_State185, TraitDefType}:                     _GotoState97Action,
	{_State185, FuncTypeType}:                     _GotoState92Action,
	{_State187, IdentifierToken}:                  _GotoState84Action,
	{_State187, StructToken}:                      _GotoState36Action,
	{_State187, EnumToken}:                        _GotoState81Action,
	{_State187, TraitToken}:                       _GotoState89Action,
	{_State187, FuncToken}:                        _GotoState83Action,
	{_State187, LparenToken}:                      _GotoState86Action,
	{_State187, LbracketToken}:                    _GotoState28Action,
	{_State187, DotToken}:                         _GotoState80Action,
	{_State187, QuestionToken}:                    _GotoState87Action,
	{_State187, ExclaimToken}:                     _GotoState82Action,
	{_State187, TildeTildeToken}:                  _GotoState88Action,
	{_State187, BitNegToken}:                      _GotoState79Action,
	{_State187, BitAndToken}:                      _GotoState78Action,
	{_State187, LexErrorToken}:                    _GotoState85Action,
	{_State187, InitializableTypeType}:            _GotoState95Action,
	{_State187, AtomTypeType}:                     _GotoState90Action,
	{_State187, ReturnableTypeType}:               _GotoState290Action,
	{_State187, ImplicitStructDefType}:            _GotoState94Action,
	{_State187, ExplicitStructDefType}:            _GotoState48Action,
	{_State187, ImplicitEnumDefType}:              _GotoState93Action,
	{_State187, ExplicitEnumDefType}:              _GotoState91Action,
	{_State187, TraitDefType}:                     _GotoState97Action,
	{_State187, FuncTypeType}:                     _GotoState92Action,
	{_State188, IntegerLiteralToken}:              _GotoState26Action,
	{_State188, FloatLiteralToken}:                _GotoState22Action,
	{_State188, RuneLiteralToken}:                 _GotoState34Action,
	{_State188, StringLiteralToken}:               _GotoState35Action,
	{_State188, IdentifierToken}:                  _GotoState25Action,
	{_State188, TrueToken}:                        _GotoState38Action,
	{_State188, FalseToken}:                       _GotoState21Action,
	{_State188, StructToken}:                      _GotoState36Action,
	{_State188, FuncToken}:                        _GotoState23Action,
	{_State188, VarToken}:                         _GotoState39Action,
	{_State188, LetToken}:                         _GotoState29Action,
	{_State188, LabelDeclToken}:                   _GotoState27Action,
	{_State188, LparenToken}:                      _GotoState31Action,
	{_State188, LbracketToken}:                    _GotoState28Action,
	{_State188, NotToken}:                         _GotoState33Action,
	{_State188, SubToken}:                         _GotoState37Action,
	{_State188, MulToken}:                         _GotoState32Action,
	{_State188, BitNegToken}:                      _GotoState20Action,
	{_State188, BitAndToken}:                      _GotoState19Action,
	{_State188, GreaterToken}:                     _GotoState24Action,
	{_State188, LexErrorToken}:                    _GotoState30Action,
	{_State188, VarDeclPatternType}:               _GotoState59Action,
	{_State188, VarOrLetType}:                     _GotoState60Action,
	{_State188, ExpressionType}:                   _GotoState291Action,
	{_State188, OptionalLabelDeclType}:            _GotoState53Action,
	{_State188, SequenceExprType}:                 _GotoState58Action,
	{_State188, BlockExprType}:                    _GotoState45Action,
	{_State188, CallExprType}:                     _GotoState46Action,
	{_State188, AtomExprType}:                     _GotoState44Action,
	{_State188, LiteralType}:                      _GotoState51Action,
	{_State188, ImplicitStructExprType}:           _GotoState49Action,
	{_State188, AccessExprType}:                   _GotoState40Action,
	{_State188, PostfixUnaryExprType}:             _GotoState55Action,
	{_State188, PrefixUnaryOpType}:                _GotoState57Action,
	{_State188, PrefixUnaryExprType}:              _GotoState56Action,
	{_State188, MulExprType}:                      _GotoState52Action,
	{_State188, AddExprType}:                      _GotoState41Action,
	{_State188, CmpExprType}:                      _GotoState47Action,
	{_State188, AndExprType}:                      _GotoState42Action,
	{_State188, OrExprType}:                       _GotoState54Action,
	{_State188, InitializableTypeType}:            _GotoState50Action,
	{_State188, ExplicitStructDefType}:            _GotoState48Action,
	{_State188, AnonymousFuncExprType}:            _GotoState43Action,
	{_State189, IntegerLiteralToken}:              _GotoState26Action,
	{_State189, FloatLiteralToken}:                _GotoState22Action,
	{_State189, RuneLiteralToken}:                 _GotoState34Action,
	{_State189, StringLiteralToken}:               _GotoState35Action,
	{_State189, IdentifierToken}:                  _GotoState100Action,
	{_State189, TrueToken}:                        _GotoState38Action,
	{_State189, FalseToken}:                       _GotoState21Action,
	{_State189, StructToken}:                      _GotoState36Action,
	{_State189, FuncToken}:                        _GotoState23Action,
	{_State189, VarToken}:                         _GotoState39Action,
	{_State189, LetToken}:                         _GotoState29Action,
	{_State189, LabelDeclToken}:                   _GotoState27Action,
	{_State189, LparenToken}:                      _GotoState31Action,
	{_State189, LbracketToken}:                    _GotoState28Action,
	{_State189, DotDotDotToken}:                   _GotoState99Action,
	{_State189, NotToken}:                         _GotoState33Action,
	{_State189, SubToken}:                         _GotoState37Action,
	{_State189, MulToken}:                         _GotoState32Action,
	{_State189, BitNegToken}:                      _GotoState20Action,
	{_State189, BitAndToken}:                      _GotoState19Action,
	{_State189, GreaterToken}:                     _GotoState24Action,
	{_State189, LexErrorToken}:                    _GotoState30Action,
	{_State189, VarDeclPatternType}:               _GotoState59Action,
	{_State189, VarOrLetType}:                     _GotoState60Action,
	{_State189, ExpressionType}:                   _GotoState104Action,
	{_State189, OptionalLabelDeclType}:            _GotoState53Action,
	{_State189, SequenceExprType}:                 _GotoState58Action,
	{_State189, BlockExprType}:                    _GotoState45Action,
	{_State189, CallExprType}:                     _GotoState46Action,
	{_State189, ArgumentType}:                     _GotoState292Action,
	{_State189, ColonExpressionsType}:             _GotoState103Action,
	{_State189, OptionalExpressionType}:           _GotoState105Action,
	{_State189, AtomExprType}:                     _GotoState44Action,
	{_State189, LiteralType}:                      _GotoState51Action,
	{_State189, ImplicitStructExprType}:           _GotoState49Action,
	{_State189, AccessExprType}:                   _GotoState40Action,
	{_State189, PostfixUnaryExprType}:             _GotoState55Action,
	{_State189, PrefixUnaryOpType}:                _GotoState57Action,
	{_State189, PrefixUnaryExprType}:              _GotoState56Action,
	{_State189, MulExprType}:                      _GotoState52Action,
	{_State189, AddExprType}:                      _GotoState41Action,
	{_State189, CmpExprType}:                      _GotoState47Action,
	{_State189, AndExprType}:                      _GotoState42Action,
	{_State189, OrExprType}:                       _GotoState54Action,
	{_State189, InitializableTypeType}:            _GotoState50Action,
	{_State189, ExplicitStructDefType}:            _GotoState48Action,
	{_State189, AnonymousFuncExprType}:            _GotoState43Action,
	{_State191, IntegerLiteralToken}:              _GotoState26Action,
	{_State191, FloatLiteralToken}:                _GotoState22Action,
	{_State191, RuneLiteralToken}:                 _GotoState34Action,
	{_State191, StringLiteralToken}:               _GotoState35Action,
	{_State191, IdentifierToken}:                  _GotoState25Action,
	{_State191, TrueToken}:                        _GotoState38Action,
	{_State191, FalseToken}:                       _GotoState21Action,
	{_State191, StructToken}:                      _GotoState36Action,
	{_State191, FuncToken}:                        _GotoState23Action,
	{_State191, VarToken}:                         _GotoState39Action,
	{_State191, LetToken}:                         _GotoState29Action,
	{_State191, LabelDeclToken}:                   _GotoState27Action,
	{_State191, LparenToken}:                      _GotoState31Action,
	{_State191, LbracketToken}:                    _GotoState28Action,
	{_State191, NotToken}:                         _GotoState33Action,
	{_State191, SubToken}:                         _GotoState37Action,
	{_State191, MulToken}:                         _GotoState32Action,
	{_State191, BitNegToken}:                      _GotoState20Action,
	{_State191, BitAndToken}:                      _GotoState19Action,
	{_State191, GreaterToken}:                     _GotoState24Action,
	{_State191, LexErrorToken}:                    _GotoState30Action,
	{_State191, VarDeclPatternType}:               _GotoState59Action,
	{_State191, VarOrLetType}:                     _GotoState60Action,
	{_State191, ExpressionType}:                   _GotoState293Action,
	{_State191, OptionalLabelDeclType}:            _GotoState53Action,
	{_State191, SequenceExprType}:                 _GotoState58Action,
	{_State191, BlockExprType}:                    _GotoState45Action,
	{_State191, CallExprType}:                     _GotoState46Action,
	{_State191, OptionalExpressionType}:           _GotoState294Action,
	{_State191, AtomExprType}:                     _GotoState44Action,
	{_State191, LiteralType}:                      _GotoState51Action,
	{_State191, ImplicitStructExprType}:           _GotoState49Action,
	{_State191, AccessExprType}:                   _GotoState40Action,
	{_State191, PostfixUnaryExprType}:             _GotoState55Action,
	{_State191, PrefixUnaryOpType}:                _GotoState57Action,
	{_State191, PrefixUnaryExprType}:              _GotoState56Action,
	{_State191, MulExprType}:                      _GotoState52Action,
	{_State191, AddExprType}:                      _GotoState41Action,
	{_State191, CmpExprType}:                      _GotoState47Action,
	{_State191, AndExprType}:                      _GotoState42Action,
	{_State191, OrExprType}:                       _GotoState54Action,
	{_State191, InitializableTypeType}:            _GotoState50Action,
	{_State191, ExplicitStructDefType}:            _GotoState48Action,
	{_State191, AnonymousFuncExprType}:            _GotoState43Action,
	{_State192, IntegerLiteralToken}:              _GotoState26Action,
	{_State192, FloatLiteralToken}:                _GotoState22Action,
	{_State192, RuneLiteralToken}:                 _GotoState34Action,
	{_State192, StringLiteralToken}:               _GotoState35Action,
	{_State192, IdentifierToken}:                  _GotoState25Action,
	{_State192, TrueToken}:                        _GotoState38Action,
	{_State192, FalseToken}:                       _GotoState21Action,
	{_State192, StructToken}:                      _GotoState36Action,
	{_State192, FuncToken}:                        _GotoState23Action,
	{_State192, VarToken}:                         _GotoState39Action,
	{_State192, LetToken}:                         _GotoState29Action,
	{_State192, LabelDeclToken}:                   _GotoState27Action,
	{_State192, LparenToken}:                      _GotoState31Action,
	{_State192, LbracketToken}:                    _GotoState28Action,
	{_State192, NotToken}:                         _GotoState33Action,
	{_State192, SubToken}:                         _GotoState37Action,
	{_State192, MulToken}:                         _GotoState32Action,
	{_State192, BitNegToken}:                      _GotoState20Action,
	{_State192, BitAndToken}:                      _GotoState19Action,
	{_State192, GreaterToken}:                     _GotoState24Action,
	{_State192, LexErrorToken}:                    _GotoState30Action,
	{_State192, VarDeclPatternType}:               _GotoState59Action,
	{_State192, VarOrLetType}:                     _GotoState60Action,
	{_State192, ExpressionType}:                   _GotoState293Action,
	{_State192, OptionalLabelDeclType}:            _GotoState53Action,
	{_State192, SequenceExprType}:                 _GotoState58Action,
	{_State192, BlockExprType}:                    _GotoState45Action,
	{_State192, CallExprType}:                     _GotoState46Action,
	{_State192, OptionalExpressionType}:           _GotoState295Action,
	{_State192, AtomExprType}:                     _GotoState44Action,
	{_State192, LiteralType}:                      _GotoState51Action,
	{_State192, ImplicitStructExprType}:           _GotoState49Action,
	{_State192, AccessExprType}:                   _GotoState40Action,
	{_State192, PostfixUnaryExprType}:             _GotoState55Action,
	{_State192, PrefixUnaryOpType}:                _GotoState57Action,
	{_State192, PrefixUnaryExprType}:              _GotoState56Action,
	{_State192, MulExprType}:                      _GotoState52Action,
	{_State192, AddExprType}:                      _GotoState41Action,
	{_State192, CmpExprType}:                      _GotoState47Action,
	{_State192, AndExprType}:                      _GotoState42Action,
	{_State192, OrExprType}:                       _GotoState54Action,
	{_State192, InitializableTypeType}:            _GotoState50Action,
	{_State192, ExplicitStructDefType}:            _GotoState48Action,
	{_State192, AnonymousFuncExprType}:            _GotoState43Action,
	{_State193, NewlinesToken}:                    _GotoState297Action,
	{_State193, CommaToken}:                       _GotoState296Action,
	{_State195, RparenToken}:                      _GotoState298Action,
	{_State196, CommaToken}:                       _GotoState299Action,
	{_State197, RbracketToken}:                    _GotoState300Action,
	{_State198, AddToken}:                         _GotoState182Action,
	{_State198, SubToken}:                         _GotoState187Action,
	{_State198, MulToken}:                         _GotoState185Action,
	{_State200, RbracketToken}:                    _GotoState301Action,
	{_State201, IntegerLiteralToken}:              _GotoState26Action,
	{_State201, FloatLiteralToken}:                _GotoState22Action,
	{_State201, RuneLiteralToken}:                 _GotoState34Action,
	{_State201, StringLiteralToken}:               _GotoState35Action,
	{_State201, IdentifierToken}:                  _GotoState100Action,
	{_State201, TrueToken}:                        _GotoState38Action,
	{_State201, FalseToken}:                       _GotoState21Action,
	{_State201, StructToken}:                      _GotoState36Action,
	{_State201, FuncToken}:                        _GotoState23Action,
	{_State201, VarToken}:                         _GotoState39Action,
	{_State201, LetToken}:                         _GotoState29Action,
	{_State201, LabelDeclToken}:                   _GotoState27Action,
	{_State201, LparenToken}:                      _GotoState31Action,
	{_State201, LbracketToken}:                    _GotoState28Action,
	{_State201, DotDotDotToken}:                   _GotoState99Action,
	{_State201, NotToken}:                         _GotoState33Action,
	{_State201, SubToken}:                         _GotoState37Action,
	{_State201, MulToken}:                         _GotoState32Action,
	{_State201, BitNegToken}:                      _GotoState20Action,
	{_State201, BitAndToken}:                      _GotoState19Action,
	{_State201, GreaterToken}:                     _GotoState24Action,
	{_State201, LexErrorToken}:                    _GotoState30Action,
	{_State201, VarDeclPatternType}:               _GotoState59Action,
	{_State201, VarOrLetType}:                     _GotoState60Action,
	{_State201, ExpressionType}:                   _GotoState104Action,
	{_State201, OptionalLabelDeclType}:            _GotoState53Action,
	{_State201, SequenceExprType}:                 _GotoState58Action,
	{_State201, BlockExprType}:                    _GotoState45Action,
	{_State201, CallExprType}:                     _GotoState46Action,
	{_State201, OptionalArgumentsType}:            _GotoState303Action,
	{_State201, ArgumentsType}:                    _GotoState302Action,
	{_State201, ArgumentType}:                     _GotoState101Action,
	{_State201, ColonExpressionsType}:             _GotoState103Action,
	{_State201, OptionalExpressionType}:           _GotoState105Action,
	{_State201, AtomExprType}:                     _GotoState44Action,
	{_State201, LiteralType}:                      _GotoState51Action,
	{_State201, ImplicitStructExprType}:           _GotoState49Action,
	{_State201, AccessExprType}:                   _GotoState40Action,
	{_State201, PostfixUnaryExprType}:             _GotoState55Action,
	{_State201, PrefixUnaryOpType}:                _GotoState57Action,
	{_State201, PrefixUnaryExprType}:              _GotoState56Action,
	{_State201, MulExprType}:                      _GotoState52Action,
	{_State201, AddExprType}:                      _GotoState41Action,
	{_State201, CmpExprType}:                      _GotoState47Action,
	{_State201, AndExprType}:                      _GotoState42Action,
	{_State201, OrExprType}:                       _GotoState54Action,
	{_State201, InitializableTypeType}:            _GotoState50Action,
	{_State201, ExplicitStructDefType}:            _GotoState48Action,
	{_State201, AnonymousFuncExprType}:            _GotoState43Action,
	{_State202, MulToken}:                         _GotoState131Action,
	{_State202, DivToken}:                         _GotoState129Action,
	{_State202, ModToken}:                         _GotoState130Action,
	{_State202, BitAndToken}:                      _GotoState126Action,
	{_State202, BitLshiftToken}:                   _GotoState127Action,
	{_State202, BitRshiftToken}:                   _GotoState128Action,
	{_State202, MulOpType}:                        _GotoState132Action,
	{_State203, EqualToken}:                       _GotoState118Action,
	{_State203, NotEqualToken}:                    _GotoState123Action,
	{_State203, LessToken}:                        _GotoState121Action,
	{_State203, LessOrEqualToken}:                 _GotoState122Action,
	{_State203, GreaterToken}:                     _GotoState119Action,
	{_State203, GreaterOrEqualToken}:              _GotoState120Action,
	{_State203, CmpOpType}:                        _GotoState124Action,
	{_State204, AddToken}:                         _GotoState112Action,
	{_State204, SubToken}:                         _GotoState115Action,
	{_State204, BitXorToken}:                      _GotoState114Action,
	{_State204, BitOrToken}:                       _GotoState113Action,
	{_State204, AddOpType}:                        _GotoState116Action,
	{_State205, RparenToken}:                      _GotoState304Action,
	{_State205, CommaToken}:                       _GotoState189Action,
	{_State207, ForToken}:                         _GotoState305Action,
	{_State208, InToken}:                          _GotoState307Action,
	{_State208, AssignToken}:                      _GotoState306Action,
	{_State209, SemicolonToken}:                   _GotoState308Action,
	{_State210, DoToken}:                          _GotoState309Action,
	{_State211, IntegerLiteralToken}:              _GotoState26Action,
	{_State211, FloatLiteralToken}:                _GotoState22Action,
	{_State211, RuneLiteralToken}:                 _GotoState34Action,
	{_State211, StringLiteralToken}:               _GotoState35Action,
	{_State211, IdentifierToken}:                  _GotoState25Action,
	{_State211, TrueToken}:                        _GotoState38Action,
	{_State211, FalseToken}:                       _GotoState21Action,
	{_State211, StructToken}:                      _GotoState36Action,
	{_State211, FuncToken}:                        _GotoState23Action,
	{_State211, VarToken}:                         _GotoState311Action,
	{_State211, LetToken}:                         _GotoState29Action,
	{_State211, LabelDeclToken}:                   _GotoState27Action,
	{_State211, LparenToken}:                      _GotoState31Action,
	{_State211, LbracketToken}:                    _GotoState28Action,
	{_State211, DotToken}:                         _GotoState310Action,
	{_State211, NotToken}:                         _GotoState33Action,
	{_State211, SubToken}:                         _GotoState37Action,
	{_State211, MulToken}:                         _GotoState32Action,
	{_State211, BitNegToken}:                      _GotoState20Action,
	{_State211, BitAndToken}:                      _GotoState19Action,
	{_State211, GreaterToken}:                     _GotoState24Action,
	{_State211, LexErrorToken}:                    _GotoState30Action,
	{_State211, VarDeclPatternType}:               _GotoState59Action,
	{_State211, VarOrLetType}:                     _GotoState60Action,
	{_State211, AssignPatternType}:                _GotoState312Action,
	{_State211, CasePatternType}:                  _GotoState313Action,
	{_State211, OptionalLabelDeclType}:            _GotoState76Action,
	{_State211, CasePatternsType}:                 _GotoState314Action,
	{_State211, SequenceExprType}:                 _GotoState315Action,
	{_State211, BlockExprType}:                    _GotoState45Action,
	{_State211, CallExprType}:                     _GotoState46Action,
	{_State211, AtomExprType}:                     _GotoState44Action,
	{_State211, LiteralType}:                      _GotoState51Action,
	{_State211, ImplicitStructExprType}:           _GotoState49Action,
	{_State211, AccessExprType}:                   _GotoState40Action,
	{_State211, PostfixUnaryExprType}:             _GotoState55Action,
	{_State211, PrefixUnaryOpType}:                _GotoState57Action,
	{_State211, PrefixUnaryExprType}:              _GotoState56Action,
	{_State211, MulExprType}:                      _GotoState52Action,
	{_State211, AddExprType}:                      _GotoState41Action,
	{_State211, CmpExprType}:                      _GotoState47Action,
	{_State211, AndExprType}:                      _GotoState42Action,
	{_State211, OrExprType}:                       _GotoState54Action,
	{_State211, InitializableTypeType}:            _GotoState50Action,
	{_State211, ExplicitStructDefType}:            _GotoState48Action,
	{_State211, AnonymousFuncExprType}:            _GotoState43Action,
	{_State212, LbraceToken}:                      _GotoState63Action,
	{_State212, BlockBodyType}:                    _GotoState316Action,
	{_State214, LbraceToken}:                      _GotoState317Action,
	{_State215, AndToken}:                         _GotoState117Action,
	{_State217, AssignToken}:                      _GotoState318Action,
	{_State219, RparenToken}:                      _GotoState320Action,
	{_State219, CommaToken}:                       _GotoState319Action,
	{_State222, AddToken}:                         _GotoState182Action,
	{_State222, SubToken}:                         _GotoState187Action,
	{_State222, MulToken}:                         _GotoState185Action,
	{_State223, IntegerLiteralToken}:              _GotoState26Action,
	{_State223, FloatLiteralToken}:                _GotoState22Action,
	{_State223, RuneLiteralToken}:                 _GotoState34Action,
	{_State223, StringLiteralToken}:               _GotoState35Action,
	{_State223, IdentifierToken}:                  _GotoState25Action,
	{_State223, TrueToken}:                        _GotoState38Action,
	{_State223, FalseToken}:                       _GotoState21Action,
	{_State223, StructToken}:                      _GotoState36Action,
	{_State223, FuncToken}:                        _GotoState23Action,
	{_State223, LabelDeclToken}:                   _GotoState27Action,
	{_State223, LparenToken}:                      _GotoState31Action,
	{_State223, LbracketToken}:                    _GotoState28Action,
	{_State223, LexErrorToken}:                    _GotoState30Action,
	{_State223, OptionalLabelDeclType}:            _GotoState76Action,
	{_State223, BlockExprType}:                    _GotoState45Action,
	{_State223, CallExprType}:                     _GotoState322Action,
	{_State223, AtomExprType}:                     _GotoState44Action,
	{_State223, LiteralType}:                      _GotoState51Action,
	{_State223, ImplicitStructExprType}:           _GotoState49Action,
	{_State223, AccessExprType}:                   _GotoState321Action,
	{_State223, InitializableTypeType}:            _GotoState50Action,
	{_State223, ExplicitStructDefType}:            _GotoState48Action,
	{_State223, AnonymousFuncExprType}:            _GotoState43Action,
	{_State226, IntegerLiteralToken}:              _GotoState26Action,
	{_State226, FloatLiteralToken}:                _GotoState22Action,
	{_State226, RuneLiteralToken}:                 _GotoState34Action,
	{_State226, StringLiteralToken}:               _GotoState35Action,
	{_State226, IdentifierToken}:                  _GotoState25Action,
	{_State226, TrueToken}:                        _GotoState38Action,
	{_State226, FalseToken}:                       _GotoState21Action,
	{_State226, StructToken}:                      _GotoState36Action,
	{_State226, FuncToken}:                        _GotoState23Action,
	{_State226, LabelDeclToken}:                   _GotoState27Action,
	{_State226, LparenToken}:                      _GotoState31Action,
	{_State226, LbracketToken}:                    _GotoState28Action,
	{_State226, LexErrorToken}:                    _GotoState30Action,
	{_State226, OptionalLabelDeclType}:            _GotoState76Action,
	{_State226, BlockExprType}:                    _GotoState45Action,
	{_State226, CallExprType}:                     _GotoState323Action,
	{_State226, AtomExprType}:                     _GotoState44Action,
	{_State226, LiteralType}:                      _GotoState51Action,
	{_State226, ImplicitStructExprType}:           _GotoState49Action,
	{_State226, AccessExprType}:                   _GotoState321Action,
	{_State226, InitializableTypeType}:            _GotoState50Action,
	{_State226, ExplicitStructDefType}:            _GotoState48Action,
	{_State226, AnonymousFuncExprType}:            _GotoState43Action,
	{_State229, LbracketToken}:                    _GotoState109Action,
	{_State229, DotToken}:                         _GotoState108Action,
	{_State229, QuestionToken}:                    _GotoState110Action,
	{_State229, DollarLbracketToken}:              _GotoState107Action,
	{_State229, AddAssignToken}:                   _GotoState324Action,
	{_State229, SubAssignToken}:                   _GotoState335Action,
	{_State229, MulAssignToken}:                   _GotoState334Action,
	{_State229, DivAssignToken}:                   _GotoState332Action,
	{_State229, ModAssignToken}:                   _GotoState333Action,
	{_State229, AddOneAssignToken}:                _GotoState325Action,
	{_State229, SubOneAssignToken}:                _GotoState336Action,
	{_State229, BitNegAssignToken}:                _GotoState328Action,
	{_State229, BitAndAssignToken}:                _GotoState326Action,
	{_State229, BitOrAssignToken}:                 _GotoState329Action,
	{_State229, BitXorAssignToken}:                _GotoState331Action,
	{_State229, BitLshiftAssignToken}:             _GotoState327Action,
	{_State229, BitRshiftAssignToken}:             _GotoState330Action,
	{_State229, UnaryOpAssignType}:                _GotoState338Action,
	{_State229, BinaryOpAssignType}:               _GotoState337Action,
	{_State229, OptionalGenericBindingType}:       _GotoState111Action,
	{_State230, AssignToken}:                      _GotoState339Action,
	{_State232, CommaToken}:                       _GotoState340Action,
	{_State234, JumpLabelToken}:                   _GotoState341Action,
	{_State234, OptionalJumpLabelType}:            _GotoState342Action,
	{_State237, NewlinesToken}:                    _GotoState343Action,
	{_State237, SemicolonToken}:                   _GotoState344Action,
	{_State241, StringLiteralToken}:               _GotoState346Action,
	{_State241, LparenToken}:                      _GotoState345Action,
	{_State241, ImportClauseType}:                 _GotoState347Action,
	{_State245, NewlinesToken}:                    _GotoState349Action,
	{_State245, CommaToken}:                       _GotoState348Action,
	{_State247, AddToken}:                         _GotoState182Action,
	{_State247, SubToken}:                         _GotoState187Action,
	{_State247, MulToken}:                         _GotoState185Action,
	{_State248, IdentifierToken}:                  _GotoState84Action,
	{_State248, StructToken}:                      _GotoState36Action,
	{_State248, EnumToken}:                        _GotoState81Action,
	{_State248, TraitToken}:                       _GotoState89Action,
	{_State248, FuncToken}:                        _GotoState83Action,
	{_State248, LparenToken}:                      _GotoState86Action,
	{_State248, LbracketToken}:                    _GotoState28Action,
	{_State248, DotToken}:                         _GotoState80Action,
	{_State248, QuestionToken}:                    _GotoState87Action,
	{_State248, ExclaimToken}:                     _GotoState82Action,
	{_State248, TildeTildeToken}:                  _GotoState88Action,
	{_State248, BitNegToken}:                      _GotoState79Action,
	{_State248, BitAndToken}:                      _GotoState78Action,
	{_State248, LexErrorToken}:                    _GotoState85Action,
	{_State248, InitializableTypeType}:            _GotoState95Action,
	{_State248, AtomTypeType}:                     _GotoState90Action,
	{_State248, ReturnableTypeType}:               _GotoState96Action,
	{_State248, ValueTypeType}:                    _GotoState350Action,
	{_State248, ImplicitStructDefType}:            _GotoState94Action,
	{_State248, ExplicitStructDefType}:            _GotoState48Action,
	{_State248, ImplicitEnumDefType}:              _GotoState93Action,
	{_State248, ExplicitEnumDefType}:              _GotoState91Action,
	{_State248, TraitDefType}:                     _GotoState97Action,
	{_State248, FuncTypeType}:                     _GotoState92Action,
	{_State250, CommaToken}:                       _GotoState351Action,
	{_State251, RbracketToken}:                    _GotoState352Action,
	{_State252, ImplementsToken}:                  _GotoState353Action,
	{_State252, AddToken}:                         _GotoState182Action,
	{_State252, SubToken}:                         _GotoState187Action,
	{_State252, MulToken}:                         _GotoState185Action,
	{_State254, IdentifierToken}:                  _GotoState157Action,
	{_State254, ParameterDefType}:                 _GotoState160Action,
	{_State254, ParameterDefsType}:                _GotoState161Action,
	{_State254, OptionalParameterDefsType}:        _GotoState354Action,
	{_State255, IdentifierToken}:                  _GotoState84Action,
	{_State255, StructToken}:                      _GotoState36Action,
	{_State255, EnumToken}:                        _GotoState81Action,
	{_State255, TraitToken}:                       _GotoState89Action,
	{_State255, FuncToken}:                        _GotoState83Action,
	{_State255, LparenToken}:                      _GotoState86Action,
	{_State255, LbracketToken}:                    _GotoState28Action,
	{_State255, DotToken}:                         _GotoState80Action,
	{_State255, QuestionToken}:                    _GotoState87Action,
	{_State255, ExclaimToken}:                     _GotoState82Action,
	{_State255, TildeTildeToken}:                  _GotoState88Action,
	{_State255, BitNegToken}:                      _GotoState79Action,
	{_State255, BitAndToken}:                      _GotoState78Action,
	{_State255, LexErrorToken}:                    _GotoState85Action,
	{_State255, InitializableTypeType}:            _GotoState95Action,
	{_State255, AtomTypeType}:                     _GotoState90Action,
	{_State255, ReturnableTypeType}:               _GotoState96Action,
	{_State255, ValueTypeType}:                    _GotoState355Action,
	{_State255, ImplicitStructDefType}:            _GotoState94Action,
	{_State255, ExplicitStructDefType}:            _GotoState48Action,
	{_State255, ImplicitEnumDefType}:              _GotoState93Action,
	{_State255, ExplicitEnumDefType}:              _GotoState91Action,
	{_State255, TraitDefType}:                     _GotoState97Action,
	{_State255, FuncTypeType}:                     _GotoState92Action,
	{_State256, AddToken}:                         _GotoState182Action,
	{_State256, SubToken}:                         _GotoState187Action,
	{_State256, MulToken}:                         _GotoState185Action,
	{_State257, IdentifierToken}:                  _GotoState356Action,
	{_State258, IdentifierToken}:                  _GotoState84Action,
	{_State258, StructToken}:                      _GotoState36Action,
	{_State258, EnumToken}:                        _GotoState81Action,
	{_State258, TraitToken}:                       _GotoState89Action,
	{_State258, FuncToken}:                        _GotoState83Action,
	{_State258, LparenToken}:                      _GotoState86Action,
	{_State258, LbracketToken}:                    _GotoState28Action,
	{_State258, DotToken}:                         _GotoState80Action,
	{_State258, QuestionToken}:                    _GotoState87Action,
	{_State258, ExclaimToken}:                     _GotoState82Action,
	{_State258, TildeTildeToken}:                  _GotoState88Action,
	{_State258, BitNegToken}:                      _GotoState79Action,
	{_State258, BitAndToken}:                      _GotoState78Action,
	{_State258, LexErrorToken}:                    _GotoState85Action,
	{_State258, InitializableTypeType}:            _GotoState95Action,
	{_State258, AtomTypeType}:                     _GotoState90Action,
	{_State258, ReturnableTypeType}:               _GotoState358Action,
	{_State258, ImplicitStructDefType}:            _GotoState94Action,
	{_State258, ExplicitStructDefType}:            _GotoState48Action,
	{_State258, ImplicitEnumDefType}:              _GotoState93Action,
	{_State258, ExplicitEnumDefType}:              _GotoState91Action,
	{_State258, TraitDefType}:                     _GotoState97Action,
	{_State258, ReturnTypeType}:                   _GotoState357Action,
	{_State258, FuncTypeType}:                     _GotoState92Action,
	{_State259, IdentifierToken}:                  _GotoState157Action,
	{_State259, ParameterDefType}:                 _GotoState359Action,
	{_State260, NewlinesToken}:                    _GotoState360Action,
	{_State260, OrToken}:                          _GotoState361Action,
	{_State261, RparenToken}:                      _GotoState362Action,
	{_State262, AssignToken}:                      _GotoState275Action,
	{_State263, NewlinesToken}:                    _GotoState363Action,
	{_State263, OrToken}:                          _GotoState364Action,
	{_State264, IdentifierToken}:                  _GotoState84Action,
	{_State264, StructToken}:                      _GotoState36Action,
	{_State264, EnumToken}:                        _GotoState81Action,
	{_State264, TraitToken}:                       _GotoState89Action,
	{_State264, FuncToken}:                        _GotoState83Action,
	{_State264, LparenToken}:                      _GotoState86Action,
	{_State264, LbracketToken}:                    _GotoState28Action,
	{_State264, DotToken}:                         _GotoState80Action,
	{_State264, QuestionToken}:                    _GotoState87Action,
	{_State264, ExclaimToken}:                     _GotoState82Action,
	{_State264, TildeTildeToken}:                  _GotoState88Action,
	{_State264, BitNegToken}:                      _GotoState79Action,
	{_State264, BitAndToken}:                      _GotoState78Action,
	{_State264, LexErrorToken}:                    _GotoState85Action,
	{_State264, InitializableTypeType}:            _GotoState95Action,
	{_State264, AtomTypeType}:                     _GotoState90Action,
	{_State264, ReturnableTypeType}:               _GotoState96Action,
	{_State264, ValueTypeType}:                    _GotoState365Action,
	{_State264, ImplicitStructDefType}:            _GotoState94Action,
	{_State264, ExplicitStructDefType}:            _GotoState48Action,
	{_State264, ImplicitEnumDefType}:              _GotoState93Action,
	{_State264, ExplicitEnumDefType}:              _GotoState91Action,
	{_State264, TraitDefType}:                     _GotoState97Action,
	{_State264, FuncTypeType}:                     _GotoState92Action,
	{_State265, IdentifierToken}:                  _GotoState84Action,
	{_State265, StructToken}:                      _GotoState36Action,
	{_State265, EnumToken}:                        _GotoState81Action,
	{_State265, TraitToken}:                       _GotoState89Action,
	{_State265, FuncToken}:                        _GotoState83Action,
	{_State265, LparenToken}:                      _GotoState86Action,
	{_State265, LbracketToken}:                    _GotoState28Action,
	{_State265, DotToken}:                         _GotoState271Action,
	{_State265, QuestionToken}:                    _GotoState87Action,
	{_State265, ExclaimToken}:                     _GotoState82Action,
	{_State265, DollarLbracketToken}:              _GotoState107Action,
	{_State265, DotDotDotToken}:                   _GotoState366Action,
	{_State265, TildeTildeToken}:                  _GotoState88Action,
	{_State265, BitNegToken}:                      _GotoState79Action,
	{_State265, BitAndToken}:                      _GotoState78Action,
	{_State265, LexErrorToken}:                    _GotoState85Action,
	{_State265, OptionalGenericBindingType}:       _GotoState169Action,
	{_State265, InitializableTypeType}:            _GotoState95Action,
	{_State265, AtomTypeType}:                     _GotoState90Action,
	{_State265, ReturnableTypeType}:               _GotoState96Action,
	{_State265, ValueTypeType}:                    _GotoState367Action,
	{_State265, ImplicitStructDefType}:            _GotoState94Action,
	{_State265, ExplicitStructDefType}:            _GotoState48Action,
	{_State265, ImplicitEnumDefType}:              _GotoState93Action,
	{_State265, ExplicitEnumDefType}:              _GotoState91Action,
	{_State265, TraitDefType}:                     _GotoState97Action,
	{_State265, FuncTypeType}:                     _GotoState92Action,
	{_State266, RparenToken}:                      _GotoState368Action,
	{_State268, CommaToken}:                       _GotoState369Action,
	{_State269, AddToken}:                         _GotoState182Action,
	{_State269, SubToken}:                         _GotoState187Action,
	{_State269, MulToken}:                         _GotoState185Action,
	{_State270, DollarLbracketToken}:              _GotoState107Action,
	{_State270, OptionalGenericBindingType}:       _GotoState370Action,
	{_State271, IdentifierToken}:                  _GotoState270Action,
	{_State271, DollarLbracketToken}:              _GotoState107Action,
	{_State271, OptionalGenericBindingType}:       _GotoState164Action,
	{_State272, AddToken}:                         _GotoState182Action,
	{_State272, SubToken}:                         _GotoState187Action,
	{_State272, MulToken}:                         _GotoState185Action,
	{_State273, IdentifierToken}:                  _GotoState371Action,
	{_State274, IdentifierToken}:                  _GotoState170Action,
	{_State274, UnsafeToken}:                      _GotoState171Action,
	{_State274, StructToken}:                      _GotoState36Action,
	{_State274, EnumToken}:                        _GotoState81Action,
	{_State274, TraitToken}:                       _GotoState89Action,
	{_State274, FuncToken}:                        _GotoState83Action,
	{_State274, LparenToken}:                      _GotoState86Action,
	{_State274, LbracketToken}:                    _GotoState28Action,
	{_State274, DotToken}:                         _GotoState80Action,
	{_State274, QuestionToken}:                    _GotoState87Action,
	{_State274, ExclaimToken}:                     _GotoState82Action,
	{_State274, TildeTildeToken}:                  _GotoState88Action,
	{_State274, BitNegToken}:                      _GotoState79Action,
	{_State274, BitAndToken}:                      _GotoState78Action,
	{_State274, LexErrorToken}:                    _GotoState85Action,
	{_State274, UnsafeStatementType}:              _GotoState177Action,
	{_State274, InitializableTypeType}:            _GotoState95Action,
	{_State274, AtomTypeType}:                     _GotoState90Action,
	{_State274, ReturnableTypeType}:               _GotoState96Action,
	{_State274, ValueTypeType}:                    _GotoState178Action,
	{_State274, FieldDefType}:                     _GotoState262Action,
	{_State274, ImplicitStructDefType}:            _GotoState94Action,
	{_State274, ExplicitStructDefType}:            _GotoState48Action,
	{_State274, EnumValueDefType}:                 _GotoState372Action,
	{_State274, ImplicitEnumDefType}:              _GotoState93Action,
	{_State274, ExplicitEnumDefType}:              _GotoState91Action,
	{_State274, TraitDefType}:                     _GotoState97Action,
	{_State274, FuncTypeType}:                     _GotoState92Action,
	{_State275, DefaultToken}:                     _GotoState373Action,
	{_State276, IdentifierToken}:                  _GotoState170Action,
	{_State276, UnsafeToken}:                      _GotoState171Action,
	{_State276, StructToken}:                      _GotoState36Action,
	{_State276, EnumToken}:                        _GotoState81Action,
	{_State276, TraitToken}:                       _GotoState89Action,
	{_State276, FuncToken}:                        _GotoState83Action,
	{_State276, LparenToken}:                      _GotoState86Action,
	{_State276, LbracketToken}:                    _GotoState28Action,
	{_State276, DotToken}:                         _GotoState80Action,
	{_State276, QuestionToken}:                    _GotoState87Action,
	{_State276, ExclaimToken}:                     _GotoState82Action,
	{_State276, TildeTildeToken}:                  _GotoState88Action,
	{_State276, BitNegToken}:                      _GotoState79Action,
	{_State276, BitAndToken}:                      _GotoState78Action,
	{_State276, LexErrorToken}:                    _GotoState85Action,
	{_State276, UnsafeStatementType}:              _GotoState177Action,
	{_State276, InitializableTypeType}:            _GotoState95Action,
	{_State276, AtomTypeType}:                     _GotoState90Action,
	{_State276, ReturnableTypeType}:               _GotoState96Action,
	{_State276, ValueTypeType}:                    _GotoState178Action,
	{_State276, FieldDefType}:                     _GotoState262Action,
	{_State276, ImplicitStructDefType}:            _GotoState94Action,
	{_State276, ExplicitStructDefType}:            _GotoState48Action,
	{_State276, EnumValueDefType}:                 _GotoState374Action,
	{_State276, ImplicitEnumDefType}:              _GotoState93Action,
	{_State276, ExplicitEnumDefType}:              _GotoState91Action,
	{_State276, TraitDefType}:                     _GotoState97Action,
	{_State276, FuncTypeType}:                     _GotoState92Action,
	{_State278, IdentifierToken}:                  _GotoState170Action,
	{_State278, UnsafeToken}:                      _GotoState171Action,
	{_State278, StructToken}:                      _GotoState36Action,
	{_State278, EnumToken}:                        _GotoState81Action,
	{_State278, TraitToken}:                       _GotoState89Action,
	{_State278, FuncToken}:                        _GotoState83Action,
	{_State278, LparenToken}:                      _GotoState86Action,
	{_State278, LbracketToken}:                    _GotoState28Action,
	{_State278, DotToken}:                         _GotoState80Action,
	{_State278, QuestionToken}:                    _GotoState87Action,
	{_State278, ExclaimToken}:                     _GotoState82Action,
	{_State278, TildeTildeToken}:                  _GotoState88Action,
	{_State278, BitNegToken}:                      _GotoState79Action,
	{_State278, BitAndToken}:                      _GotoState78Action,
	{_State278, LexErrorToken}:                    _GotoState85Action,
	{_State278, UnsafeStatementType}:              _GotoState177Action,
	{_State278, InitializableTypeType}:            _GotoState95Action,
	{_State278, AtomTypeType}:                     _GotoState90Action,
	{_State278, ReturnableTypeType}:               _GotoState96Action,
	{_State278, ValueTypeType}:                    _GotoState178Action,
	{_State278, FieldDefType}:                     _GotoState375Action,
	{_State278, ImplicitStructDefType}:            _GotoState94Action,
	{_State278, ExplicitStructDefType}:            _GotoState48Action,
	{_State278, ImplicitEnumDefType}:              _GotoState93Action,
	{_State278, ExplicitEnumDefType}:              _GotoState91Action,
	{_State278, TraitDefType}:                     _GotoState97Action,
	{_State278, FuncTypeType}:                     _GotoState92Action,
	{_State280, IdentifierToken}:                  _GotoState376Action,
	{_State280, LparenToken}:                      _GotoState167Action,
	{_State283, RparenToken}:                      _GotoState377Action,
	{_State284, NewlinesToken}:                    _GotoState379Action,
	{_State284, CommaToken}:                       _GotoState378Action,
	{_State287, RbracketToken}:                    _GotoState380Action,
	{_State287, AddToken}:                         _GotoState182Action,
	{_State287, SubToken}:                         _GotoState187Action,
	{_State287, MulToken}:                         _GotoState185Action,
	{_State288, RbracketToken}:                    _GotoState381Action,
	{_State296, IdentifierToken}:                  _GotoState170Action,
	{_State296, UnsafeToken}:                      _GotoState171Action,
	{_State296, StructToken}:                      _GotoState36Action,
	{_State296, EnumToken}:                        _GotoState81Action,
	{_State296, TraitToken}:                       _GotoState89Action,
	{_State296, FuncToken}:                        _GotoState83Action,
	{_State296, LparenToken}:                      _GotoState86Action,
	{_State296, LbracketToken}:                    _GotoState28Action,
	{_State296, DotToken}:                         _GotoState80Action,
	{_State296, QuestionToken}:                    _GotoState87Action,
	{_State296, ExclaimToken}:                     _GotoState82Action,
	{_State296, TildeTildeToken}:                  _GotoState88Action,
	{_State296, BitNegToken}:                      _GotoState79Action,
	{_State296, BitAndToken}:                      _GotoState78Action,
	{_State296, LexErrorToken}:                    _GotoState85Action,
	{_State296, UnsafeStatementType}:              _GotoState177Action,
	{_State296, InitializableTypeType}:            _GotoState95Action,
	{_State296, AtomTypeType}:                     _GotoState90Action,
	{_State296, ReturnableTypeType}:               _GotoState96Action,
	{_State296, ValueTypeType}:                    _GotoState178Action,
	{_State296, FieldDefType}:                     _GotoState382Action,
	{_State296, ImplicitStructDefType}:            _GotoState94Action,
	{_State296, ExplicitStructDefType}:            _GotoState48Action,
	{_State296, ImplicitEnumDefType}:              _GotoState93Action,
	{_State296, ExplicitEnumDefType}:              _GotoState91Action,
	{_State296, TraitDefType}:                     _GotoState97Action,
	{_State296, FuncTypeType}:                     _GotoState92Action,
	{_State297, IdentifierToken}:                  _GotoState170Action,
	{_State297, UnsafeToken}:                      _GotoState171Action,
	{_State297, StructToken}:                      _GotoState36Action,
	{_State297, EnumToken}:                        _GotoState81Action,
	{_State297, TraitToken}:                       _GotoState89Action,
	{_State297, FuncToken}:                        _GotoState83Action,
	{_State297, LparenToken}:                      _GotoState86Action,
	{_State297, LbracketToken}:                    _GotoState28Action,
	{_State297, DotToken}:                         _GotoState80Action,
	{_State297, QuestionToken}:                    _GotoState87Action,
	{_State297, ExclaimToken}:                     _GotoState82Action,
	{_State297, TildeTildeToken}:                  _GotoState88Action,
	{_State297, BitNegToken}:                      _GotoState79Action,
	{_State297, BitAndToken}:                      _GotoState78Action,
	{_State297, LexErrorToken}:                    _GotoState85Action,
	{_State297, UnsafeStatementType}:              _GotoState177Action,
	{_State297, InitializableTypeType}:            _GotoState95Action,
	{_State297, AtomTypeType}:                     _GotoState90Action,
	{_State297, ReturnableTypeType}:               _GotoState96Action,
	{_State297, ValueTypeType}:                    _GotoState178Action,
	{_State297, FieldDefType}:                     _GotoState383Action,
	{_State297, ImplicitStructDefType}:            _GotoState94Action,
	{_State297, ExplicitStructDefType}:            _GotoState48Action,
	{_State297, ImplicitEnumDefType}:              _GotoState93Action,
	{_State297, ExplicitEnumDefType}:              _GotoState91Action,
	{_State297, TraitDefType}:                     _GotoState97Action,
	{_State297, FuncTypeType}:                     _GotoState92Action,
	{_State299, IdentifierToken}:                  _GotoState84Action,
	{_State299, StructToken}:                      _GotoState36Action,
	{_State299, EnumToken}:                        _GotoState81Action,
	{_State299, TraitToken}:                       _GotoState89Action,
	{_State299, FuncToken}:                        _GotoState83Action,
	{_State299, LparenToken}:                      _GotoState86Action,
	{_State299, LbracketToken}:                    _GotoState28Action,
	{_State299, DotToken}:                         _GotoState80Action,
	{_State299, QuestionToken}:                    _GotoState87Action,
	{_State299, ExclaimToken}:                     _GotoState82Action,
	{_State299, TildeTildeToken}:                  _GotoState88Action,
	{_State299, BitNegToken}:                      _GotoState79Action,
	{_State299, BitAndToken}:                      _GotoState78Action,
	{_State299, LexErrorToken}:                    _GotoState85Action,
	{_State299, InitializableTypeType}:            _GotoState95Action,
	{_State299, AtomTypeType}:                     _GotoState90Action,
	{_State299, ReturnableTypeType}:               _GotoState96Action,
	{_State299, ValueTypeType}:                    _GotoState384Action,
	{_State299, ImplicitStructDefType}:            _GotoState94Action,
	{_State299, ExplicitStructDefType}:            _GotoState48Action,
	{_State299, ImplicitEnumDefType}:              _GotoState93Action,
	{_State299, ExplicitEnumDefType}:              _GotoState91Action,
	{_State299, TraitDefType}:                     _GotoState97Action,
	{_State299, FuncTypeType}:                     _GotoState92Action,
	{_State302, CommaToken}:                       _GotoState189Action,
	{_State303, RparenToken}:                      _GotoState385Action,
	{_State305, IntegerLiteralToken}:              _GotoState26Action,
	{_State305, FloatLiteralToken}:                _GotoState22Action,
	{_State305, RuneLiteralToken}:                 _GotoState34Action,
	{_State305, StringLiteralToken}:               _GotoState35Action,
	{_State305, IdentifierToken}:                  _GotoState25Action,
	{_State305, TrueToken}:                        _GotoState38Action,
	{_State305, FalseToken}:                       _GotoState21Action,
	{_State305, StructToken}:                      _GotoState36Action,
	{_State305, FuncToken}:                        _GotoState23Action,
	{_State305, VarToken}:                         _GotoState39Action,
	{_State305, LetToken}:                         _GotoState29Action,
	{_State305, LabelDeclToken}:                   _GotoState27Action,
	{_State305, LparenToken}:                      _GotoState31Action,
	{_State305, LbracketToken}:                    _GotoState28Action,
	{_State305, NotToken}:                         _GotoState33Action,
	{_State305, SubToken}:                         _GotoState37Action,
	{_State305, MulToken}:                         _GotoState32Action,
	{_State305, BitNegToken}:                      _GotoState20Action,
	{_State305, BitAndToken}:                      _GotoState19Action,
	{_State305, GreaterToken}:                     _GotoState24Action,
	{_State305, LexErrorToken}:                    _GotoState30Action,
	{_State305, VarDeclPatternType}:               _GotoState59Action,
	{_State305, VarOrLetType}:                     _GotoState60Action,
	{_State305, OptionalLabelDeclType}:            _GotoState76Action,
	{_State305, SequenceExprType}:                 _GotoState386Action,
	{_State305, BlockExprType}:                    _GotoState45Action,
	{_State305, CallExprType}:                     _GotoState46Action,
	{_State305, AtomExprType}:                     _GotoState44Action,
	{_State305, LiteralType}:                      _GotoState51Action,
	{_State305, ImplicitStructExprType}:           _GotoState49Action,
	{_State305, AccessExprType}:                   _GotoState40Action,
	{_State305, PostfixUnaryExprType}:             _GotoState55Action,
	{_State305, PrefixUnaryOpType}:                _GotoState57Action,
	{_State305, PrefixUnaryExprType}:              _GotoState56Action,
	{_State305, MulExprType}:                      _GotoState52Action,
	{_State305, AddExprType}:                      _GotoState41Action,
	{_State305, CmpExprType}:                      _GotoState47Action,
	{_State305, AndExprType}:                      _GotoState42Action,
	{_State305, OrExprType}:                       _GotoState54Action,
	{_State305, InitializableTypeType}:            _GotoState50Action,
	{_State305, ExplicitStructDefType}:            _GotoState48Action,
	{_State305, AnonymousFuncExprType}:            _GotoState43Action,
	{_State306, IntegerLiteralToken}:              _GotoState26Action,
	{_State306, FloatLiteralToken}:                _GotoState22Action,
	{_State306, RuneLiteralToken}:                 _GotoState34Action,
	{_State306, StringLiteralToken}:               _GotoState35Action,
	{_State306, IdentifierToken}:                  _GotoState25Action,
	{_State306, TrueToken}:                        _GotoState38Action,
	{_State306, FalseToken}:                       _GotoState21Action,
	{_State306, StructToken}:                      _GotoState36Action,
	{_State306, FuncToken}:                        _GotoState23Action,
	{_State306, VarToken}:                         _GotoState39Action,
	{_State306, LetToken}:                         _GotoState29Action,
	{_State306, LabelDeclToken}:                   _GotoState27Action,
	{_State306, LparenToken}:                      _GotoState31Action,
	{_State306, LbracketToken}:                    _GotoState28Action,
	{_State306, NotToken}:                         _GotoState33Action,
	{_State306, SubToken}:                         _GotoState37Action,
	{_State306, MulToken}:                         _GotoState32Action,
	{_State306, BitNegToken}:                      _GotoState20Action,
	{_State306, BitAndToken}:                      _GotoState19Action,
	{_State306, GreaterToken}:                     _GotoState24Action,
	{_State306, LexErrorToken}:                    _GotoState30Action,
	{_State306, VarDeclPatternType}:               _GotoState59Action,
	{_State306, VarOrLetType}:                     _GotoState60Action,
	{_State306, OptionalLabelDeclType}:            _GotoState76Action,
	{_State306, SequenceExprType}:                 _GotoState387Action,
	{_State306, BlockExprType}:                    _GotoState45Action,
	{_State306, CallExprType}:                     _GotoState46Action,
	{_State306, AtomExprType}:                     _GotoState44Action,
	{_State306, LiteralType}:                      _GotoState51Action,
	{_State306, ImplicitStructExprType}:           _GotoState49Action,
	{_State306, AccessExprType}:                   _GotoState40Action,
	{_State306, PostfixUnaryExprType}:             _GotoState55Action,
	{_State306, PrefixUnaryOpType}:                _GotoState57Action,
	{_State306, PrefixUnaryExprType}:              _GotoState56Action,
	{_State306, MulExprType}:                      _GotoState52Action,
	{_State306, AddExprType}:                      _GotoState41Action,
	{_State306, CmpExprType}:                      _GotoState47Action,
	{_State306, AndExprType}:                      _GotoState42Action,
	{_State306, OrExprType}:                       _GotoState54Action,
	{_State306, InitializableTypeType}:            _GotoState50Action,
	{_State306, ExplicitStructDefType}:            _GotoState48Action,
	{_State306, AnonymousFuncExprType}:            _GotoState43Action,
	{_State307, IntegerLiteralToken}:              _GotoState26Action,
	{_State307, FloatLiteralToken}:                _GotoState22Action,
	{_State307, RuneLiteralToken}:                 _GotoState34Action,
	{_State307, StringLiteralToken}:               _GotoState35Action,
	{_State307, IdentifierToken}:                  _GotoState25Action,
	{_State307, TrueToken}:                        _GotoState38Action,
	{_State307, FalseToken}:                       _GotoState21Action,
	{_State307, StructToken}:                      _GotoState36Action,
	{_State307, FuncToken}:                        _GotoState23Action,
	{_State307, VarToken}:                         _GotoState39Action,
	{_State307, LetToken}:                         _GotoState29Action,
	{_State307, LabelDeclToken}:                   _GotoState27Action,
	{_State307, LparenToken}:                      _GotoState31Action,
	{_State307, LbracketToken}:                    _GotoState28Action,
	{_State307, NotToken}:                         _GotoState33Action,
	{_State307, SubToken}:                         _GotoState37Action,
	{_State307, MulToken}:                         _GotoState32Action,
	{_State307, BitNegToken}:                      _GotoState20Action,
	{_State307, BitAndToken}:                      _GotoState19Action,
	{_State307, GreaterToken}:                     _GotoState24Action,
	{_State307, LexErrorToken}:                    _GotoState30Action,
	{_State307, VarDeclPatternType}:               _GotoState59Action,
	{_State307, VarOrLetType}:                     _GotoState60Action,
	{_State307, OptionalLabelDeclType}:            _GotoState76Action,
	{_State307, SequenceExprType}:                 _GotoState388Action,
	{_State307, BlockExprType}:                    _GotoState45Action,
	{_State307, CallExprType}:                     _GotoState46Action,
	{_State307, AtomExprType}:                     _GotoState44Action,
	{_State307, LiteralType}:                      _GotoState51Action,
	{_State307, ImplicitStructExprType}:           _GotoState49Action,
	{_State307, AccessExprType}:                   _GotoState40Action,
	{_State307, PostfixUnaryExprType}:             _GotoState55Action,
	{_State307, PrefixUnaryOpType}:                _GotoState57Action,
	{_State307, PrefixUnaryExprType}:              _GotoState56Action,
	{_State307, MulExprType}:                      _GotoState52Action,
	{_State307, AddExprType}:                      _GotoState41Action,
	{_State307, CmpExprType}:                      _GotoState47Action,
	{_State307, AndExprType}:                      _GotoState42Action,
	{_State307, OrExprType}:                       _GotoState54Action,
	{_State307, InitializableTypeType}:            _GotoState50Action,
	{_State307, ExplicitStructDefType}:            _GotoState48Action,
	{_State307, AnonymousFuncExprType}:            _GotoState43Action,
	{_State308, IntegerLiteralToken}:              _GotoState26Action,
	{_State308, FloatLiteralToken}:                _GotoState22Action,
	{_State308, RuneLiteralToken}:                 _GotoState34Action,
	{_State308, StringLiteralToken}:               _GotoState35Action,
	{_State308, IdentifierToken}:                  _GotoState25Action,
	{_State308, TrueToken}:                        _GotoState38Action,
	{_State308, FalseToken}:                       _GotoState21Action,
	{_State308, StructToken}:                      _GotoState36Action,
	{_State308, FuncToken}:                        _GotoState23Action,
	{_State308, VarToken}:                         _GotoState39Action,
	{_State308, LetToken}:                         _GotoState29Action,
	{_State308, LabelDeclToken}:                   _GotoState27Action,
	{_State308, LparenToken}:                      _GotoState31Action,
	{_State308, LbracketToken}:                    _GotoState28Action,
	{_State308, NotToken}:                         _GotoState33Action,
	{_State308, SubToken}:                         _GotoState37Action,
	{_State308, MulToken}:                         _GotoState32Action,
	{_State308, BitNegToken}:                      _GotoState20Action,
	{_State308, BitAndToken}:                      _GotoState19Action,
	{_State308, GreaterToken}:                     _GotoState24Action,
	{_State308, LexErrorToken}:                    _GotoState30Action,
	{_State308, VarDeclPatternType}:               _GotoState59Action,
	{_State308, VarOrLetType}:                     _GotoState60Action,
	{_State308, OptionalLabelDeclType}:            _GotoState76Action,
	{_State308, OptionalSequenceExprType}:         _GotoState389Action,
	{_State308, SequenceExprType}:                 _GotoState390Action,
	{_State308, BlockExprType}:                    _GotoState45Action,
	{_State308, CallExprType}:                     _GotoState46Action,
	{_State308, AtomExprType}:                     _GotoState44Action,
	{_State308, LiteralType}:                      _GotoState51Action,
	{_State308, ImplicitStructExprType}:           _GotoState49Action,
	{_State308, AccessExprType}:                   _GotoState40Action,
	{_State308, PostfixUnaryExprType}:             _GotoState55Action,
	{_State308, PrefixUnaryOpType}:                _GotoState57Action,
	{_State308, PrefixUnaryExprType}:              _GotoState56Action,
	{_State308, MulExprType}:                      _GotoState52Action,
	{_State308, AddExprType}:                      _GotoState41Action,
	{_State308, CmpExprType}:                      _GotoState47Action,
	{_State308, AndExprType}:                      _GotoState42Action,
	{_State308, OrExprType}:                       _GotoState54Action,
	{_State308, InitializableTypeType}:            _GotoState50Action,
	{_State308, ExplicitStructDefType}:            _GotoState48Action,
	{_State308, AnonymousFuncExprType}:            _GotoState43Action,
	{_State309, LbraceToken}:                      _GotoState63Action,
	{_State309, BlockBodyType}:                    _GotoState391Action,
	{_State310, IdentifierToken}:                  _GotoState392Action,
	{_State311, DotToken}:                         _GotoState393Action,
	{_State314, CommaToken}:                       _GotoState395Action,
	{_State314, AssignToken}:                      _GotoState394Action,
	{_State316, ElseToken}:                        _GotoState396Action,
	{_State317, CaseToken}:                        _GotoState397Action,
	{_State317, CaseBranchesType}:                 _GotoState399Action,
	{_State317, CaseBranchType}:                   _GotoState398Action,
	{_State318, IdentifierToken}:                  _GotoState143Action,
	{_State318, LparenToken}:                      _GotoState144Action,
	{_State318, VarPatternType}:                   _GotoState400Action,
	{_State318, TuplePatternType}:                 _GotoState145Action,
	{_State319, IdentifierToken}:                  _GotoState217Action,
	{_State319, LparenToken}:                      _GotoState144Action,
	{_State319, DotDotDotToken}:                   _GotoState216Action,
	{_State319, VarPatternType}:                   _GotoState220Action,
	{_State319, TuplePatternType}:                 _GotoState145Action,
	{_State319, FieldVarPatternType}:              _GotoState401Action,
	{_State321, LbracketToken}:                    _GotoState109Action,
	{_State321, DotToken}:                         _GotoState108Action,
	{_State321, DollarLbracketToken}:              _GotoState107Action,
	{_State321, OptionalGenericBindingType}:       _GotoState111Action,
	{_State337, IntegerLiteralToken}:              _GotoState26Action,
	{_State337, FloatLiteralToken}:                _GotoState22Action,
	{_State337, RuneLiteralToken}:                 _GotoState34Action,
	{_State337, StringLiteralToken}:               _GotoState35Action,
	{_State337, IdentifierToken}:                  _GotoState25Action,
	{_State337, TrueToken}:                        _GotoState38Action,
	{_State337, FalseToken}:                       _GotoState21Action,
	{_State337, StructToken}:                      _GotoState36Action,
	{_State337, FuncToken}:                        _GotoState23Action,
	{_State337, VarToken}:                         _GotoState39Action,
	{_State337, LetToken}:                         _GotoState29Action,
	{_State337, LabelDeclToken}:                   _GotoState27Action,
	{_State337, LparenToken}:                      _GotoState31Action,
	{_State337, LbracketToken}:                    _GotoState28Action,
	{_State337, NotToken}:                         _GotoState33Action,
	{_State337, SubToken}:                         _GotoState37Action,
	{_State337, MulToken}:                         _GotoState32Action,
	{_State337, BitNegToken}:                      _GotoState20Action,
	{_State337, BitAndToken}:                      _GotoState19Action,
	{_State337, GreaterToken}:                     _GotoState24Action,
	{_State337, LexErrorToken}:                    _GotoState30Action,
	{_State337, VarDeclPatternType}:               _GotoState59Action,
	{_State337, VarOrLetType}:                     _GotoState60Action,
	{_State337, ExpressionType}:                   _GotoState402Action,
	{_State337, OptionalLabelDeclType}:            _GotoState53Action,
	{_State337, SequenceExprType}:                 _GotoState58Action,
	{_State337, BlockExprType}:                    _GotoState45Action,
	{_State337, CallExprType}:                     _GotoState46Action,
	{_State337, AtomExprType}:                     _GotoState44Action,
	{_State337, LiteralType}:                      _GotoState51Action,
	{_State337, ImplicitStructExprType}:           _GotoState49Action,
	{_State337, AccessExprType}:                   _GotoState40Action,
	{_State337, PostfixUnaryExprType}:             _GotoState55Action,
	{_State337, PrefixUnaryOpType}:                _GotoState57Action,
	{_State337, PrefixUnaryExprType}:              _GotoState56Action,
	{_State337, MulExprType}:                      _GotoState52Action,
	{_State337, AddExprType}:                      _GotoState41Action,
	{_State337, CmpExprType}:                      _GotoState47Action,
	{_State337, AndExprType}:                      _GotoState42Action,
	{_State337, OrExprType}:                       _GotoState54Action,
	{_State337, InitializableTypeType}:            _GotoState50Action,
	{_State337, ExplicitStructDefType}:            _GotoState48Action,
	{_State337, AnonymousFuncExprType}:            _GotoState43Action,
	{_State339, IntegerLiteralToken}:              _GotoState26Action,
	{_State339, FloatLiteralToken}:                _GotoState22Action,
	{_State339, RuneLiteralToken}:                 _GotoState34Action,
	{_State339, StringLiteralToken}:               _GotoState35Action,
	{_State339, IdentifierToken}:                  _GotoState25Action,
	{_State339, TrueToken}:                        _GotoState38Action,
	{_State339, FalseToken}:                       _GotoState21Action,
	{_State339, StructToken}:                      _GotoState36Action,
	{_State339, FuncToken}:                        _GotoState23Action,
	{_State339, VarToken}:                         _GotoState39Action,
	{_State339, LetToken}:                         _GotoState29Action,
	{_State339, LabelDeclToken}:                   _GotoState27Action,
	{_State339, LparenToken}:                      _GotoState31Action,
	{_State339, LbracketToken}:                    _GotoState28Action,
	{_State339, NotToken}:                         _GotoState33Action,
	{_State339, SubToken}:                         _GotoState37Action,
	{_State339, MulToken}:                         _GotoState32Action,
	{_State339, BitNegToken}:                      _GotoState20Action,
	{_State339, BitAndToken}:                      _GotoState19Action,
	{_State339, GreaterToken}:                     _GotoState24Action,
	{_State339, LexErrorToken}:                    _GotoState30Action,
	{_State339, VarDeclPatternType}:               _GotoState59Action,
	{_State339, VarOrLetType}:                     _GotoState60Action,
	{_State339, ExpressionType}:                   _GotoState403Action,
	{_State339, OptionalLabelDeclType}:            _GotoState53Action,
	{_State339, SequenceExprType}:                 _GotoState58Action,
	{_State339, BlockExprType}:                    _GotoState45Action,
	{_State339, CallExprType}:                     _GotoState46Action,
	{_State339, AtomExprType}:                     _GotoState44Action,
	{_State339, LiteralType}:                      _GotoState51Action,
	{_State339, ImplicitStructExprType}:           _GotoState49Action,
	{_State339, AccessExprType}:                   _GotoState40Action,
	{_State339, PostfixUnaryExprType}:             _GotoState55Action,
	{_State339, PrefixUnaryOpType}:                _GotoState57Action,
	{_State339, PrefixUnaryExprType}:              _GotoState56Action,
	{_State339, MulExprType}:                      _GotoState52Action,
	{_State339, AddExprType}:                      _GotoState41Action,
	{_State339, CmpExprType}:                      _GotoState47Action,
	{_State339, AndExprType}:                      _GotoState42Action,
	{_State339, OrExprType}:                       _GotoState54Action,
	{_State339, InitializableTypeType}:            _GotoState50Action,
	{_State339, ExplicitStructDefType}:            _GotoState48Action,
	{_State339, AnonymousFuncExprType}:            _GotoState43Action,
	{_State340, IntegerLiteralToken}:              _GotoState26Action,
	{_State340, FloatLiteralToken}:                _GotoState22Action,
	{_State340, RuneLiteralToken}:                 _GotoState34Action,
	{_State340, StringLiteralToken}:               _GotoState35Action,
	{_State340, IdentifierToken}:                  _GotoState25Action,
	{_State340, TrueToken}:                        _GotoState38Action,
	{_State340, FalseToken}:                       _GotoState21Action,
	{_State340, StructToken}:                      _GotoState36Action,
	{_State340, FuncToken}:                        _GotoState23Action,
	{_State340, VarToken}:                         _GotoState39Action,
	{_State340, LetToken}:                         _GotoState29Action,
	{_State340, LabelDeclToken}:                   _GotoState27Action,
	{_State340, LparenToken}:                      _GotoState31Action,
	{_State340, LbracketToken}:                    _GotoState28Action,
	{_State340, NotToken}:                         _GotoState33Action,
	{_State340, SubToken}:                         _GotoState37Action,
	{_State340, MulToken}:                         _GotoState32Action,
	{_State340, BitNegToken}:                      _GotoState20Action,
	{_State340, BitAndToken}:                      _GotoState19Action,
	{_State340, GreaterToken}:                     _GotoState24Action,
	{_State340, LexErrorToken}:                    _GotoState30Action,
	{_State340, VarDeclPatternType}:               _GotoState59Action,
	{_State340, VarOrLetType}:                     _GotoState60Action,
	{_State340, ExpressionType}:                   _GotoState404Action,
	{_State340, OptionalLabelDeclType}:            _GotoState53Action,
	{_State340, SequenceExprType}:                 _GotoState58Action,
	{_State340, BlockExprType}:                    _GotoState45Action,
	{_State340, CallExprType}:                     _GotoState46Action,
	{_State340, AtomExprType}:                     _GotoState44Action,
	{_State340, LiteralType}:                      _GotoState51Action,
	{_State340, ImplicitStructExprType}:           _GotoState49Action,
	{_State340, AccessExprType}:                   _GotoState40Action,
	{_State340, PostfixUnaryExprType}:             _GotoState55Action,
	{_State340, PrefixUnaryOpType}:                _GotoState57Action,
	{_State340, PrefixUnaryExprType}:              _GotoState56Action,
	{_State340, MulExprType}:                      _GotoState52Action,
	{_State340, AddExprType}:                      _GotoState41Action,
	{_State340, CmpExprType}:                      _GotoState47Action,
	{_State340, AndExprType}:                      _GotoState42Action,
	{_State340, OrExprType}:                       _GotoState54Action,
	{_State340, InitializableTypeType}:            _GotoState50Action,
	{_State340, ExplicitStructDefType}:            _GotoState48Action,
	{_State340, AnonymousFuncExprType}:            _GotoState43Action,
	{_State342, IntegerLiteralToken}:              _GotoState26Action,
	{_State342, FloatLiteralToken}:                _GotoState22Action,
	{_State342, RuneLiteralToken}:                 _GotoState34Action,
	{_State342, StringLiteralToken}:               _GotoState35Action,
	{_State342, IdentifierToken}:                  _GotoState25Action,
	{_State342, TrueToken}:                        _GotoState38Action,
	{_State342, FalseToken}:                       _GotoState21Action,
	{_State342, StructToken}:                      _GotoState36Action,
	{_State342, FuncToken}:                        _GotoState23Action,
	{_State342, VarToken}:                         _GotoState39Action,
	{_State342, LetToken}:                         _GotoState29Action,
	{_State342, LabelDeclToken}:                   _GotoState27Action,
	{_State342, LparenToken}:                      _GotoState31Action,
	{_State342, LbracketToken}:                    _GotoState28Action,
	{_State342, NotToken}:                         _GotoState33Action,
	{_State342, SubToken}:                         _GotoState37Action,
	{_State342, MulToken}:                         _GotoState32Action,
	{_State342, BitNegToken}:                      _GotoState20Action,
	{_State342, BitAndToken}:                      _GotoState19Action,
	{_State342, GreaterToken}:                     _GotoState24Action,
	{_State342, LexErrorToken}:                    _GotoState30Action,
	{_State342, VarDeclPatternType}:               _GotoState59Action,
	{_State342, VarOrLetType}:                     _GotoState60Action,
	{_State342, ExpressionType}:                   _GotoState231Action,
	{_State342, OptionalLabelDeclType}:            _GotoState53Action,
	{_State342, SequenceExprType}:                 _GotoState58Action,
	{_State342, BlockExprType}:                    _GotoState45Action,
	{_State342, ExpressionsType}:                  _GotoState405Action,
	{_State342, OptionalExpressionsType}:          _GotoState406Action,
	{_State342, CallExprType}:                     _GotoState46Action,
	{_State342, AtomExprType}:                     _GotoState44Action,
	{_State342, LiteralType}:                      _GotoState51Action,
	{_State342, ImplicitStructExprType}:           _GotoState49Action,
	{_State342, AccessExprType}:                   _GotoState40Action,
	{_State342, PostfixUnaryExprType}:             _GotoState55Action,
	{_State342, PrefixUnaryOpType}:                _GotoState57Action,
	{_State342, PrefixUnaryExprType}:              _GotoState56Action,
	{_State342, MulExprType}:                      _GotoState52Action,
	{_State342, AddExprType}:                      _GotoState41Action,
	{_State342, CmpExprType}:                      _GotoState47Action,
	{_State342, AndExprType}:                      _GotoState42Action,
	{_State342, OrExprType}:                       _GotoState54Action,
	{_State342, InitializableTypeType}:            _GotoState50Action,
	{_State342, ExplicitStructDefType}:            _GotoState48Action,
	{_State342, AnonymousFuncExprType}:            _GotoState43Action,
	{_State345, StringLiteralToken}:               _GotoState346Action,
	{_State345, ImportClauseType}:                 _GotoState407Action,
	{_State345, ImportClauseTerminalType}:         _GotoState408Action,
	{_State345, ImportClausesType}:                _GotoState409Action,
	{_State346, AsToken}:                          _GotoState410Action,
	{_State350, AddToken}:                         _GotoState182Action,
	{_State350, SubToken}:                         _GotoState187Action,
	{_State350, MulToken}:                         _GotoState185Action,
	{_State351, IdentifierToken}:                  _GotoState248Action,
	{_State351, GenericParameterDefType}:          _GotoState411Action,
	{_State353, IdentifierToken}:                  _GotoState84Action,
	{_State353, StructToken}:                      _GotoState36Action,
	{_State353, EnumToken}:                        _GotoState81Action,
	{_State353, TraitToken}:                       _GotoState89Action,
	{_State353, FuncToken}:                        _GotoState83Action,
	{_State353, LparenToken}:                      _GotoState86Action,
	{_State353, LbracketToken}:                    _GotoState28Action,
	{_State353, DotToken}:                         _GotoState80Action,
	{_State353, QuestionToken}:                    _GotoState87Action,
	{_State353, ExclaimToken}:                     _GotoState82Action,
	{_State353, TildeTildeToken}:                  _GotoState88Action,
	{_State353, BitNegToken}:                      _GotoState79Action,
	{_State353, BitAndToken}:                      _GotoState78Action,
	{_State353, LexErrorToken}:                    _GotoState85Action,
	{_State353, InitializableTypeType}:            _GotoState95Action,
	{_State353, AtomTypeType}:                     _GotoState90Action,
	{_State353, ReturnableTypeType}:               _GotoState96Action,
	{_State353, ValueTypeType}:                    _GotoState412Action,
	{_State353, ImplicitStructDefType}:            _GotoState94Action,
	{_State353, ExplicitStructDefType}:            _GotoState48Action,
	{_State353, ImplicitEnumDefType}:              _GotoState93Action,
	{_State353, ExplicitEnumDefType}:              _GotoState91Action,
	{_State353, TraitDefType}:                     _GotoState97Action,
	{_State353, FuncTypeType}:                     _GotoState92Action,
	{_State354, RparenToken}:                      _GotoState413Action,
	{_State355, AddToken}:                         _GotoState182Action,
	{_State355, SubToken}:                         _GotoState187Action,
	{_State355, MulToken}:                         _GotoState185Action,
	{_State356, DollarLbracketToken}:              _GotoState153Action,
	{_State356, OptionalGenericParametersType}:    _GotoState414Action,
	{_State357, LbraceToken}:                      _GotoState63Action,
	{_State357, BlockBodyType}:                    _GotoState415Action,
	{_State360, IdentifierToken}:                  _GotoState170Action,
	{_State360, UnsafeToken}:                      _GotoState171Action,
	{_State360, StructToken}:                      _GotoState36Action,
	{_State360, EnumToken}:                        _GotoState81Action,
	{_State360, TraitToken}:                       _GotoState89Action,
	{_State360, FuncToken}:                        _GotoState83Action,
	{_State360, LparenToken}:                      _GotoState86Action,
	{_State360, LbracketToken}:                    _GotoState28Action,
	{_State360, DotToken}:                         _GotoState80Action,
	{_State360, QuestionToken}:                    _GotoState87Action,
	{_State360, ExclaimToken}:                     _GotoState82Action,
	{_State360, TildeTildeToken}:                  _GotoState88Action,
	{_State360, BitNegToken}:                      _GotoState79Action,
	{_State360, BitAndToken}:                      _GotoState78Action,
	{_State360, LexErrorToken}:                    _GotoState85Action,
	{_State360, UnsafeStatementType}:              _GotoState177Action,
	{_State360, InitializableTypeType}:            _GotoState95Action,
	{_State360, AtomTypeType}:                     _GotoState90Action,
	{_State360, ReturnableTypeType}:               _GotoState96Action,
	{_State360, ValueTypeType}:                    _GotoState178Action,
	{_State360, FieldDefType}:                     _GotoState262Action,
	{_State360, ImplicitStructDefType}:            _GotoState94Action,
	{_State360, ExplicitStructDefType}:            _GotoState48Action,
	{_State360, EnumValueDefType}:                 _GotoState416Action,
	{_State360, ImplicitEnumDefType}:              _GotoState93Action,
	{_State360, ExplicitEnumDefType}:              _GotoState91Action,
	{_State360, TraitDefType}:                     _GotoState97Action,
	{_State360, FuncTypeType}:                     _GotoState92Action,
	{_State361, IdentifierToken}:                  _GotoState170Action,
	{_State361, UnsafeToken}:                      _GotoState171Action,
	{_State361, StructToken}:                      _GotoState36Action,
	{_State361, EnumToken}:                        _GotoState81Action,
	{_State361, TraitToken}:                       _GotoState89Action,
	{_State361, FuncToken}:                        _GotoState83Action,
	{_State361, LparenToken}:                      _GotoState86Action,
	{_State361, LbracketToken}:                    _GotoState28Action,
	{_State361, DotToken}:                         _GotoState80Action,
	{_State361, QuestionToken}:                    _GotoState87Action,
	{_State361, ExclaimToken}:                     _GotoState82Action,
	{_State361, TildeTildeToken}:                  _GotoState88Action,
	{_State361, BitNegToken}:                      _GotoState79Action,
	{_State361, BitAndToken}:                      _GotoState78Action,
	{_State361, LexErrorToken}:                    _GotoState85Action,
	{_State361, UnsafeStatementType}:              _GotoState177Action,
	{_State361, InitializableTypeType}:            _GotoState95Action,
	{_State361, AtomTypeType}:                     _GotoState90Action,
	{_State361, ReturnableTypeType}:               _GotoState96Action,
	{_State361, ValueTypeType}:                    _GotoState178Action,
	{_State361, FieldDefType}:                     _GotoState262Action,
	{_State361, ImplicitStructDefType}:            _GotoState94Action,
	{_State361, ExplicitStructDefType}:            _GotoState48Action,
	{_State361, EnumValueDefType}:                 _GotoState417Action,
	{_State361, ImplicitEnumDefType}:              _GotoState93Action,
	{_State361, ExplicitEnumDefType}:              _GotoState91Action,
	{_State361, TraitDefType}:                     _GotoState97Action,
	{_State361, FuncTypeType}:                     _GotoState92Action,
	{_State363, IdentifierToken}:                  _GotoState170Action,
	{_State363, UnsafeToken}:                      _GotoState171Action,
	{_State363, StructToken}:                      _GotoState36Action,
	{_State363, EnumToken}:                        _GotoState81Action,
	{_State363, TraitToken}:                       _GotoState89Action,
	{_State363, FuncToken}:                        _GotoState83Action,
	{_State363, LparenToken}:                      _GotoState86Action,
	{_State363, LbracketToken}:                    _GotoState28Action,
	{_State363, DotToken}:                         _GotoState80Action,
	{_State363, QuestionToken}:                    _GotoState87Action,
	{_State363, ExclaimToken}:                     _GotoState82Action,
	{_State363, TildeTildeToken}:                  _GotoState88Action,
	{_State363, BitNegToken}:                      _GotoState79Action,
	{_State363, BitAndToken}:                      _GotoState78Action,
	{_State363, LexErrorToken}:                    _GotoState85Action,
	{_State363, UnsafeStatementType}:              _GotoState177Action,
	{_State363, InitializableTypeType}:            _GotoState95Action,
	{_State363, AtomTypeType}:                     _GotoState90Action,
	{_State363, ReturnableTypeType}:               _GotoState96Action,
	{_State363, ValueTypeType}:                    _GotoState178Action,
	{_State363, FieldDefType}:                     _GotoState262Action,
	{_State363, ImplicitStructDefType}:            _GotoState94Action,
	{_State363, ExplicitStructDefType}:            _GotoState48Action,
	{_State363, EnumValueDefType}:                 _GotoState418Action,
	{_State363, ImplicitEnumDefType}:              _GotoState93Action,
	{_State363, ExplicitEnumDefType}:              _GotoState91Action,
	{_State363, TraitDefType}:                     _GotoState97Action,
	{_State363, FuncTypeType}:                     _GotoState92Action,
	{_State364, IdentifierToken}:                  _GotoState170Action,
	{_State364, UnsafeToken}:                      _GotoState171Action,
	{_State364, StructToken}:                      _GotoState36Action,
	{_State364, EnumToken}:                        _GotoState81Action,
	{_State364, TraitToken}:                       _GotoState89Action,
	{_State364, FuncToken}:                        _GotoState83Action,
	{_State364, LparenToken}:                      _GotoState86Action,
	{_State364, LbracketToken}:                    _GotoState28Action,
	{_State364, DotToken}:                         _GotoState80Action,
	{_State364, QuestionToken}:                    _GotoState87Action,
	{_State364, ExclaimToken}:                     _GotoState82Action,
	{_State364, TildeTildeToken}:                  _GotoState88Action,
	{_State364, BitNegToken}:                      _GotoState79Action,
	{_State364, BitAndToken}:                      _GotoState78Action,
	{_State364, LexErrorToken}:                    _GotoState85Action,
	{_State364, UnsafeStatementType}:              _GotoState177Action,
	{_State364, InitializableTypeType}:            _GotoState95Action,
	{_State364, AtomTypeType}:                     _GotoState90Action,
	{_State364, ReturnableTypeType}:               _GotoState96Action,
	{_State364, ValueTypeType}:                    _GotoState178Action,
	{_State364, FieldDefType}:                     _GotoState262Action,
	{_State364, ImplicitStructDefType}:            _GotoState94Action,
	{_State364, ExplicitStructDefType}:            _GotoState48Action,
	{_State364, EnumValueDefType}:                 _GotoState419Action,
	{_State364, ImplicitEnumDefType}:              _GotoState93Action,
	{_State364, ExplicitEnumDefType}:              _GotoState91Action,
	{_State364, TraitDefType}:                     _GotoState97Action,
	{_State364, FuncTypeType}:                     _GotoState92Action,
	{_State365, AddToken}:                         _GotoState182Action,
	{_State365, SubToken}:                         _GotoState187Action,
	{_State365, MulToken}:                         _GotoState185Action,
	{_State366, IdentifierToken}:                  _GotoState84Action,
	{_State366, StructToken}:                      _GotoState36Action,
	{_State366, EnumToken}:                        _GotoState81Action,
	{_State366, TraitToken}:                       _GotoState89Action,
	{_State366, FuncToken}:                        _GotoState83Action,
	{_State366, LparenToken}:                      _GotoState86Action,
	{_State366, LbracketToken}:                    _GotoState28Action,
	{_State366, DotToken}:                         _GotoState80Action,
	{_State366, QuestionToken}:                    _GotoState87Action,
	{_State366, ExclaimToken}:                     _GotoState82Action,
	{_State366, TildeTildeToken}:                  _GotoState88Action,
	{_State366, BitNegToken}:                      _GotoState79Action,
	{_State366, BitAndToken}:                      _GotoState78Action,
	{_State366, LexErrorToken}:                    _GotoState85Action,
	{_State366, InitializableTypeType}:            _GotoState95Action,
	{_State366, AtomTypeType}:                     _GotoState90Action,
	{_State366, ReturnableTypeType}:               _GotoState96Action,
	{_State366, ValueTypeType}:                    _GotoState420Action,
	{_State366, ImplicitStructDefType}:            _GotoState94Action,
	{_State366, ExplicitStructDefType}:            _GotoState48Action,
	{_State366, ImplicitEnumDefType}:              _GotoState93Action,
	{_State366, ExplicitEnumDefType}:              _GotoState91Action,
	{_State366, TraitDefType}:                     _GotoState97Action,
	{_State366, FuncTypeType}:                     _GotoState92Action,
	{_State367, AddToken}:                         _GotoState182Action,
	{_State367, SubToken}:                         _GotoState187Action,
	{_State367, MulToken}:                         _GotoState185Action,
	{_State368, IdentifierToken}:                  _GotoState84Action,
	{_State368, StructToken}:                      _GotoState36Action,
	{_State368, EnumToken}:                        _GotoState81Action,
	{_State368, TraitToken}:                       _GotoState89Action,
	{_State368, FuncToken}:                        _GotoState83Action,
	{_State368, LparenToken}:                      _GotoState86Action,
	{_State368, LbracketToken}:                    _GotoState28Action,
	{_State368, DotToken}:                         _GotoState80Action,
	{_State368, QuestionToken}:                    _GotoState87Action,
	{_State368, ExclaimToken}:                     _GotoState82Action,
	{_State368, TildeTildeToken}:                  _GotoState88Action,
	{_State368, BitNegToken}:                      _GotoState79Action,
	{_State368, BitAndToken}:                      _GotoState78Action,
	{_State368, LexErrorToken}:                    _GotoState85Action,
	{_State368, InitializableTypeType}:            _GotoState95Action,
	{_State368, AtomTypeType}:                     _GotoState90Action,
	{_State368, ReturnableTypeType}:               _GotoState358Action,
	{_State368, ImplicitStructDefType}:            _GotoState94Action,
	{_State368, ExplicitStructDefType}:            _GotoState48Action,
	{_State368, ImplicitEnumDefType}:              _GotoState93Action,
	{_State368, ExplicitEnumDefType}:              _GotoState91Action,
	{_State368, TraitDefType}:                     _GotoState97Action,
	{_State368, ReturnTypeType}:                   _GotoState421Action,
	{_State368, FuncTypeType}:                     _GotoState92Action,
	{_State369, IdentifierToken}:                  _GotoState265Action,
	{_State369, StructToken}:                      _GotoState36Action,
	{_State369, EnumToken}:                        _GotoState81Action,
	{_State369, TraitToken}:                       _GotoState89Action,
	{_State369, FuncToken}:                        _GotoState83Action,
	{_State369, LparenToken}:                      _GotoState86Action,
	{_State369, LbracketToken}:                    _GotoState28Action,
	{_State369, DotToken}:                         _GotoState80Action,
	{_State369, QuestionToken}:                    _GotoState87Action,
	{_State369, ExclaimToken}:                     _GotoState82Action,
	{_State369, DotDotDotToken}:                   _GotoState264Action,
	{_State369, TildeTildeToken}:                  _GotoState88Action,
	{_State369, BitNegToken}:                      _GotoState79Action,
	{_State369, BitAndToken}:                      _GotoState78Action,
	{_State369, LexErrorToken}:                    _GotoState85Action,
	{_State369, InitializableTypeType}:            _GotoState95Action,
	{_State369, AtomTypeType}:                     _GotoState90Action,
	{_State369, ReturnableTypeType}:               _GotoState96Action,
	{_State369, ValueTypeType}:                    _GotoState269Action,
	{_State369, ImplicitStructDefType}:            _GotoState94Action,
	{_State369, ExplicitStructDefType}:            _GotoState48Action,
	{_State369, ImplicitEnumDefType}:              _GotoState93Action,
	{_State369, ExplicitEnumDefType}:              _GotoState91Action,
	{_State369, TraitDefType}:                     _GotoState97Action,
	{_State369, ParameterDeclType}:                _GotoState422Action,
	{_State369, FuncTypeType}:                     _GotoState92Action,
	{_State371, GreaterToken}:                     _GotoState423Action,
	{_State376, LparenToken}:                      _GotoState424Action,
	{_State378, IdentifierToken}:                  _GotoState170Action,
	{_State378, UnsafeToken}:                      _GotoState171Action,
	{_State378, StructToken}:                      _GotoState36Action,
	{_State378, EnumToken}:                        _GotoState81Action,
	{_State378, TraitToken}:                       _GotoState89Action,
	{_State378, FuncToken}:                        _GotoState280Action,
	{_State378, LparenToken}:                      _GotoState86Action,
	{_State378, LbracketToken}:                    _GotoState28Action,
	{_State378, DotToken}:                         _GotoState80Action,
	{_State378, QuestionToken}:                    _GotoState87Action,
	{_State378, ExclaimToken}:                     _GotoState82Action,
	{_State378, TildeTildeToken}:                  _GotoState88Action,
	{_State378, BitNegToken}:                      _GotoState79Action,
	{_State378, BitAndToken}:                      _GotoState78Action,
	{_State378, LexErrorToken}:                    _GotoState85Action,
	{_State378, UnsafeStatementType}:              _GotoState177Action,
	{_State378, InitializableTypeType}:            _GotoState95Action,
	{_State378, AtomTypeType}:                     _GotoState90Action,
	{_State378, ReturnableTypeType}:               _GotoState96Action,
	{_State378, ValueTypeType}:                    _GotoState178Action,
	{_State378, FieldDefType}:                     _GotoState281Action,
	{_State378, ImplicitStructDefType}:            _GotoState94Action,
	{_State378, ExplicitStructDefType}:            _GotoState48Action,
	{_State378, ImplicitEnumDefType}:              _GotoState93Action,
	{_State378, ExplicitEnumDefType}:              _GotoState91Action,
	{_State378, TraitPropertyType}:                _GotoState425Action,
	{_State378, TraitDefType}:                     _GotoState97Action,
	{_State378, FuncTypeType}:                     _GotoState92Action,
	{_State378, MethodSignatureType}:              _GotoState282Action,
	{_State379, IdentifierToken}:                  _GotoState170Action,
	{_State379, UnsafeToken}:                      _GotoState171Action,
	{_State379, StructToken}:                      _GotoState36Action,
	{_State379, EnumToken}:                        _GotoState81Action,
	{_State379, TraitToken}:                       _GotoState89Action,
	{_State379, FuncToken}:                        _GotoState280Action,
	{_State379, LparenToken}:                      _GotoState86Action,
	{_State379, LbracketToken}:                    _GotoState28Action,
	{_State379, DotToken}:                         _GotoState80Action,
	{_State379, QuestionToken}:                    _GotoState87Action,
	{_State379, ExclaimToken}:                     _GotoState82Action,
	{_State379, TildeTildeToken}:                  _GotoState88Action,
	{_State379, BitNegToken}:                      _GotoState79Action,
	{_State379, BitAndToken}:                      _GotoState78Action,
	{_State379, LexErrorToken}:                    _GotoState85Action,
	{_State379, UnsafeStatementType}:              _GotoState177Action,
	{_State379, InitializableTypeType}:            _GotoState95Action,
	{_State379, AtomTypeType}:                     _GotoState90Action,
	{_State379, ReturnableTypeType}:               _GotoState96Action,
	{_State379, ValueTypeType}:                    _GotoState178Action,
	{_State379, FieldDefType}:                     _GotoState281Action,
	{_State379, ImplicitStructDefType}:            _GotoState94Action,
	{_State379, ExplicitStructDefType}:            _GotoState48Action,
	{_State379, ImplicitEnumDefType}:              _GotoState93Action,
	{_State379, ExplicitEnumDefType}:              _GotoState91Action,
	{_State379, TraitPropertyType}:                _GotoState426Action,
	{_State379, TraitDefType}:                     _GotoState97Action,
	{_State379, FuncTypeType}:                     _GotoState92Action,
	{_State379, MethodSignatureType}:              _GotoState282Action,
	{_State384, AddToken}:                         _GotoState182Action,
	{_State384, SubToken}:                         _GotoState187Action,
	{_State384, MulToken}:                         _GotoState185Action,
	{_State388, DoToken}:                          _GotoState427Action,
	{_State389, SemicolonToken}:                   _GotoState428Action,
	{_State392, LparenToken}:                      _GotoState31Action,
	{_State392, ImplicitStructExprType}:           _GotoState429Action,
	{_State393, IdentifierToken}:                  _GotoState430Action,
	{_State394, IntegerLiteralToken}:              _GotoState26Action,
	{_State394, FloatLiteralToken}:                _GotoState22Action,
	{_State394, RuneLiteralToken}:                 _GotoState34Action,
	{_State394, StringLiteralToken}:               _GotoState35Action,
	{_State394, IdentifierToken}:                  _GotoState25Action,
	{_State394, TrueToken}:                        _GotoState38Action,
	{_State394, FalseToken}:                       _GotoState21Action,
	{_State394, StructToken}:                      _GotoState36Action,
	{_State394, FuncToken}:                        _GotoState23Action,
	{_State394, VarToken}:                         _GotoState39Action,
	{_State394, LetToken}:                         _GotoState29Action,
	{_State394, LabelDeclToken}:                   _GotoState27Action,
	{_State394, LparenToken}:                      _GotoState31Action,
	{_State394, LbracketToken}:                    _GotoState28Action,
	{_State394, NotToken}:                         _GotoState33Action,
	{_State394, SubToken}:                         _GotoState37Action,
	{_State394, MulToken}:                         _GotoState32Action,
	{_State394, BitNegToken}:                      _GotoState20Action,
	{_State394, BitAndToken}:                      _GotoState19Action,
	{_State394, GreaterToken}:                     _GotoState24Action,
	{_State394, LexErrorToken}:                    _GotoState30Action,
	{_State394, VarDeclPatternType}:               _GotoState59Action,
	{_State394, VarOrLetType}:                     _GotoState60Action,
	{_State394, OptionalLabelDeclType}:            _GotoState76Action,
	{_State394, SequenceExprType}:                 _GotoState431Action,
	{_State394, BlockExprType}:                    _GotoState45Action,
	{_State394, CallExprType}:                     _GotoState46Action,
	{_State394, AtomExprType}:                     _GotoState44Action,
	{_State394, LiteralType}:                      _GotoState51Action,
	{_State394, ImplicitStructExprType}:           _GotoState49Action,
	{_State394, AccessExprType}:                   _GotoState40Action,
	{_State394, PostfixUnaryExprType}:             _GotoState55Action,
	{_State394, PrefixUnaryOpType}:                _GotoState57Action,
	{_State394, PrefixUnaryExprType}:              _GotoState56Action,
	{_State394, MulExprType}:                      _GotoState52Action,
	{_State394, AddExprType}:                      _GotoState41Action,
	{_State394, CmpExprType}:                      _GotoState47Action,
	{_State394, AndExprType}:                      _GotoState42Action,
	{_State394, OrExprType}:                       _GotoState54Action,
	{_State394, InitializableTypeType}:            _GotoState50Action,
	{_State394, ExplicitStructDefType}:            _GotoState48Action,
	{_State394, AnonymousFuncExprType}:            _GotoState43Action,
	{_State395, IntegerLiteralToken}:              _GotoState26Action,
	{_State395, FloatLiteralToken}:                _GotoState22Action,
	{_State395, RuneLiteralToken}:                 _GotoState34Action,
	{_State395, StringLiteralToken}:               _GotoState35Action,
	{_State395, IdentifierToken}:                  _GotoState25Action,
	{_State395, TrueToken}:                        _GotoState38Action,
	{_State395, FalseToken}:                       _GotoState21Action,
	{_State395, StructToken}:                      _GotoState36Action,
	{_State395, FuncToken}:                        _GotoState23Action,
	{_State395, VarToken}:                         _GotoState311Action,
	{_State395, LetToken}:                         _GotoState29Action,
	{_State395, LabelDeclToken}:                   _GotoState27Action,
	{_State395, LparenToken}:                      _GotoState31Action,
	{_State395, LbracketToken}:                    _GotoState28Action,
	{_State395, DotToken}:                         _GotoState310Action,
	{_State395, NotToken}:                         _GotoState33Action,
	{_State395, SubToken}:                         _GotoState37Action,
	{_State395, MulToken}:                         _GotoState32Action,
	{_State395, BitNegToken}:                      _GotoState20Action,
	{_State395, BitAndToken}:                      _GotoState19Action,
	{_State395, GreaterToken}:                     _GotoState24Action,
	{_State395, LexErrorToken}:                    _GotoState30Action,
	{_State395, VarDeclPatternType}:               _GotoState59Action,
	{_State395, VarOrLetType}:                     _GotoState60Action,
	{_State395, AssignPatternType}:                _GotoState312Action,
	{_State395, CasePatternType}:                  _GotoState432Action,
	{_State395, OptionalLabelDeclType}:            _GotoState76Action,
	{_State395, SequenceExprType}:                 _GotoState315Action,
	{_State395, BlockExprType}:                    _GotoState45Action,
	{_State395, CallExprType}:                     _GotoState46Action,
	{_State395, AtomExprType}:                     _GotoState44Action,
	{_State395, LiteralType}:                      _GotoState51Action,
	{_State395, ImplicitStructExprType}:           _GotoState49Action,
	{_State395, AccessExprType}:                   _GotoState40Action,
	{_State395, PostfixUnaryExprType}:             _GotoState55Action,
	{_State395, PrefixUnaryOpType}:                _GotoState57Action,
	{_State395, PrefixUnaryExprType}:              _GotoState56Action,
	{_State395, MulExprType}:                      _GotoState52Action,
	{_State395, AddExprType}:                      _GotoState41Action,
	{_State395, CmpExprType}:                      _GotoState47Action,
	{_State395, AndExprType}:                      _GotoState42Action,
	{_State395, OrExprType}:                       _GotoState54Action,
	{_State395, InitializableTypeType}:            _GotoState50Action,
	{_State395, ExplicitStructDefType}:            _GotoState48Action,
	{_State395, AnonymousFuncExprType}:            _GotoState43Action,
	{_State396, IfToken}:                          _GotoState135Action,
	{_State396, LbraceToken}:                      _GotoState63Action,
	{_State396, IfExprType}:                       _GotoState434Action,
	{_State396, BlockBodyType}:                    _GotoState433Action,
	{_State397, IntegerLiteralToken}:              _GotoState26Action,
	{_State397, FloatLiteralToken}:                _GotoState22Action,
	{_State397, RuneLiteralToken}:                 _GotoState34Action,
	{_State397, StringLiteralToken}:               _GotoState35Action,
	{_State397, IdentifierToken}:                  _GotoState25Action,
	{_State397, TrueToken}:                        _GotoState38Action,
	{_State397, FalseToken}:                       _GotoState21Action,
	{_State397, StructToken}:                      _GotoState36Action,
	{_State397, FuncToken}:                        _GotoState23Action,
	{_State397, VarToken}:                         _GotoState311Action,
	{_State397, LetToken}:                         _GotoState29Action,
	{_State397, LabelDeclToken}:                   _GotoState27Action,
	{_State397, LparenToken}:                      _GotoState31Action,
	{_State397, LbracketToken}:                    _GotoState28Action,
	{_State397, DotToken}:                         _GotoState310Action,
	{_State397, NotToken}:                         _GotoState33Action,
	{_State397, SubToken}:                         _GotoState37Action,
	{_State397, MulToken}:                         _GotoState32Action,
	{_State397, BitNegToken}:                      _GotoState20Action,
	{_State397, BitAndToken}:                      _GotoState19Action,
	{_State397, GreaterToken}:                     _GotoState24Action,
	{_State397, LexErrorToken}:                    _GotoState30Action,
	{_State397, VarDeclPatternType}:               _GotoState59Action,
	{_State397, VarOrLetType}:                     _GotoState60Action,
	{_State397, AssignPatternType}:                _GotoState312Action,
	{_State397, CasePatternType}:                  _GotoState313Action,
	{_State397, OptionalLabelDeclType}:            _GotoState76Action,
	{_State397, CasePatternsType}:                 _GotoState435Action,
	{_State397, SequenceExprType}:                 _GotoState315Action,
	{_State397, BlockExprType}:                    _GotoState45Action,
	{_State397, CallExprType}:                     _GotoState46Action,
	{_State397, AtomExprType}:                     _GotoState44Action,
	{_State397, LiteralType}:                      _GotoState51Action,
	{_State397, ImplicitStructExprType}:           _GotoState49Action,
	{_State397, AccessExprType}:                   _GotoState40Action,
	{_State397, PostfixUnaryExprType}:             _GotoState55Action,
	{_State397, PrefixUnaryOpType}:                _GotoState57Action,
	{_State397, PrefixUnaryExprType}:              _GotoState56Action,
	{_State397, MulExprType}:                      _GotoState52Action,
	{_State397, AddExprType}:                      _GotoState41Action,
	{_State397, CmpExprType}:                      _GotoState47Action,
	{_State397, AndExprType}:                      _GotoState42Action,
	{_State397, OrExprType}:                       _GotoState54Action,
	{_State397, InitializableTypeType}:            _GotoState50Action,
	{_State397, ExplicitStructDefType}:            _GotoState48Action,
	{_State397, AnonymousFuncExprType}:            _GotoState43Action,
	{_State399, CaseToken}:                        _GotoState397Action,
	{_State399, DefaultToken}:                     _GotoState436Action,
	{_State399, CaseBranchType}:                   _GotoState437Action,
	{_State399, OptionalDefaultBranchType}:        _GotoState439Action,
	{_State399, DefaultBranchType}:                _GotoState438Action,
	{_State405, CommaToken}:                       _GotoState340Action,
	{_State407, NewlinesToken}:                    _GotoState441Action,
	{_State407, CommaToken}:                       _GotoState440Action,
	{_State409, StringLiteralToken}:               _GotoState346Action,
	{_State409, RparenToken}:                      _GotoState442Action,
	{_State409, ImportClauseType}:                 _GotoState407Action,
	{_State409, ImportClauseTerminalType}:         _GotoState443Action,
	{_State410, IdentifierToken}:                  _GotoState444Action,
	{_State412, AddToken}:                         _GotoState182Action,
	{_State412, SubToken}:                         _GotoState187Action,
	{_State412, MulToken}:                         _GotoState185Action,
	{_State413, IdentifierToken}:                  _GotoState84Action,
	{_State413, StructToken}:                      _GotoState36Action,
	{_State413, EnumToken}:                        _GotoState81Action,
	{_State413, TraitToken}:                       _GotoState89Action,
	{_State413, FuncToken}:                        _GotoState83Action,
	{_State413, LparenToken}:                      _GotoState86Action,
	{_State413, LbracketToken}:                    _GotoState28Action,
	{_State413, DotToken}:                         _GotoState80Action,
	{_State413, QuestionToken}:                    _GotoState87Action,
	{_State413, ExclaimToken}:                     _GotoState82Action,
	{_State413, TildeTildeToken}:                  _GotoState88Action,
	{_State413, BitNegToken}:                      _GotoState79Action,
	{_State413, BitAndToken}:                      _GotoState78Action,
	{_State413, LexErrorToken}:                    _GotoState85Action,
	{_State413, InitializableTypeType}:            _GotoState95Action,
	{_State413, AtomTypeType}:                     _GotoState90Action,
	{_State413, ReturnableTypeType}:               _GotoState358Action,
	{_State413, ImplicitStructDefType}:            _GotoState94Action,
	{_State413, ExplicitStructDefType}:            _GotoState48Action,
	{_State413, ImplicitEnumDefType}:              _GotoState93Action,
	{_State413, ExplicitEnumDefType}:              _GotoState91Action,
	{_State413, TraitDefType}:                     _GotoState97Action,
	{_State413, ReturnTypeType}:                   _GotoState445Action,
	{_State413, FuncTypeType}:                     _GotoState92Action,
	{_State414, LparenToken}:                      _GotoState446Action,
	{_State420, AddToken}:                         _GotoState182Action,
	{_State420, SubToken}:                         _GotoState187Action,
	{_State420, MulToken}:                         _GotoState185Action,
	{_State423, StringLiteralToken}:               _GotoState447Action,
	{_State424, IdentifierToken}:                  _GotoState265Action,
	{_State424, StructToken}:                      _GotoState36Action,
	{_State424, EnumToken}:                        _GotoState81Action,
	{_State424, TraitToken}:                       _GotoState89Action,
	{_State424, FuncToken}:                        _GotoState83Action,
	{_State424, LparenToken}:                      _GotoState86Action,
	{_State424, LbracketToken}:                    _GotoState28Action,
	{_State424, DotToken}:                         _GotoState80Action,
	{_State424, QuestionToken}:                    _GotoState87Action,
	{_State424, ExclaimToken}:                     _GotoState82Action,
	{_State424, DotDotDotToken}:                   _GotoState264Action,
	{_State424, TildeTildeToken}:                  _GotoState88Action,
	{_State424, BitNegToken}:                      _GotoState79Action,
	{_State424, BitAndToken}:                      _GotoState78Action,
	{_State424, LexErrorToken}:                    _GotoState85Action,
	{_State424, InitializableTypeType}:            _GotoState95Action,
	{_State424, AtomTypeType}:                     _GotoState90Action,
	{_State424, ReturnableTypeType}:               _GotoState96Action,
	{_State424, ValueTypeType}:                    _GotoState269Action,
	{_State424, ImplicitStructDefType}:            _GotoState94Action,
	{_State424, ExplicitStructDefType}:            _GotoState48Action,
	{_State424, ImplicitEnumDefType}:              _GotoState93Action,
	{_State424, ExplicitEnumDefType}:              _GotoState91Action,
	{_State424, TraitDefType}:                     _GotoState97Action,
	{_State424, ParameterDeclType}:                _GotoState267Action,
	{_State424, ParameterDeclsType}:               _GotoState268Action,
	{_State424, OptionalParameterDeclsType}:       _GotoState448Action,
	{_State424, FuncTypeType}:                     _GotoState92Action,
	{_State427, LbraceToken}:                      _GotoState63Action,
	{_State427, BlockBodyType}:                    _GotoState449Action,
	{_State428, IntegerLiteralToken}:              _GotoState26Action,
	{_State428, FloatLiteralToken}:                _GotoState22Action,
	{_State428, RuneLiteralToken}:                 _GotoState34Action,
	{_State428, StringLiteralToken}:               _GotoState35Action,
	{_State428, IdentifierToken}:                  _GotoState25Action,
	{_State428, TrueToken}:                        _GotoState38Action,
	{_State428, FalseToken}:                       _GotoState21Action,
	{_State428, StructToken}:                      _GotoState36Action,
	{_State428, FuncToken}:                        _GotoState23Action,
	{_State428, VarToken}:                         _GotoState39Action,
	{_State428, LetToken}:                         _GotoState29Action,
	{_State428, LabelDeclToken}:                   _GotoState27Action,
	{_State428, LparenToken}:                      _GotoState31Action,
	{_State428, LbracketToken}:                    _GotoState28Action,
	{_State428, NotToken}:                         _GotoState33Action,
	{_State428, SubToken}:                         _GotoState37Action,
	{_State428, MulToken}:                         _GotoState32Action,
	{_State428, BitNegToken}:                      _GotoState20Action,
	{_State428, BitAndToken}:                      _GotoState19Action,
	{_State428, GreaterToken}:                     _GotoState24Action,
	{_State428, LexErrorToken}:                    _GotoState30Action,
	{_State428, VarDeclPatternType}:               _GotoState59Action,
	{_State428, VarOrLetType}:                     _GotoState60Action,
	{_State428, OptionalLabelDeclType}:            _GotoState76Action,
	{_State428, OptionalSequenceExprType}:         _GotoState450Action,
	{_State428, SequenceExprType}:                 _GotoState390Action,
	{_State428, BlockExprType}:                    _GotoState45Action,
	{_State428, CallExprType}:                     _GotoState46Action,
	{_State428, AtomExprType}:                     _GotoState44Action,
	{_State428, LiteralType}:                      _GotoState51Action,
	{_State428, ImplicitStructExprType}:           _GotoState49Action,
	{_State428, AccessExprType}:                   _GotoState40Action,
	{_State428, PostfixUnaryExprType}:             _GotoState55Action,
	{_State428, PrefixUnaryOpType}:                _GotoState57Action,
	{_State428, PrefixUnaryExprType}:              _GotoState56Action,
	{_State428, MulExprType}:                      _GotoState52Action,
	{_State428, AddExprType}:                      _GotoState41Action,
	{_State428, CmpExprType}:                      _GotoState47Action,
	{_State428, AndExprType}:                      _GotoState42Action,
	{_State428, OrExprType}:                       _GotoState54Action,
	{_State428, InitializableTypeType}:            _GotoState50Action,
	{_State428, ExplicitStructDefType}:            _GotoState48Action,
	{_State428, AnonymousFuncExprType}:            _GotoState43Action,
	{_State430, LparenToken}:                      _GotoState144Action,
	{_State430, TuplePatternType}:                 _GotoState451Action,
	{_State435, CommaToken}:                       _GotoState395Action,
	{_State435, ColonToken}:                       _GotoState452Action,
	{_State436, ColonToken}:                       _GotoState453Action,
	{_State439, RbraceToken}:                      _GotoState454Action,
	{_State445, LbraceToken}:                      _GotoState63Action,
	{_State445, BlockBodyType}:                    _GotoState455Action,
	{_State446, IdentifierToken}:                  _GotoState157Action,
	{_State446, ParameterDefType}:                 _GotoState160Action,
	{_State446, ParameterDefsType}:                _GotoState161Action,
	{_State446, OptionalParameterDefsType}:        _GotoState456Action,
	{_State448, RparenToken}:                      _GotoState457Action,
	{_State450, DoToken}:                          _GotoState458Action,
	{_State452, IntegerLiteralToken}:              _GotoState26Action,
	{_State452, FloatLiteralToken}:                _GotoState22Action,
	{_State452, RuneLiteralToken}:                 _GotoState34Action,
	{_State452, StringLiteralToken}:               _GotoState35Action,
	{_State452, IdentifierToken}:                  _GotoState25Action,
	{_State452, TrueToken}:                        _GotoState38Action,
	{_State452, FalseToken}:                       _GotoState21Action,
	{_State452, StructToken}:                      _GotoState36Action,
	{_State452, FuncToken}:                        _GotoState23Action,
	{_State452, VarToken}:                         _GotoState39Action,
	{_State452, LetToken}:                         _GotoState29Action,
	{_State452, LabelDeclToken}:                   _GotoState27Action,
	{_State452, LparenToken}:                      _GotoState31Action,
	{_State452, LbracketToken}:                    _GotoState28Action,
	{_State452, NotToken}:                         _GotoState33Action,
	{_State452, SubToken}:                         _GotoState37Action,
	{_State452, MulToken}:                         _GotoState32Action,
	{_State452, BitNegToken}:                      _GotoState20Action,
	{_State452, BitAndToken}:                      _GotoState19Action,
	{_State452, GreaterToken}:                     _GotoState24Action,
	{_State452, LexErrorToken}:                    _GotoState30Action,
	{_State452, VarDeclPatternType}:               _GotoState59Action,
	{_State452, VarOrLetType}:                     _GotoState60Action,
	{_State452, OptionalLabelDeclType}:            _GotoState76Action,
	{_State452, SequenceExprType}:                 _GotoState459Action,
	{_State452, BlockExprType}:                    _GotoState45Action,
	{_State452, CallExprType}:                     _GotoState46Action,
	{_State452, AtomExprType}:                     _GotoState44Action,
	{_State452, LiteralType}:                      _GotoState51Action,
	{_State452, ImplicitStructExprType}:           _GotoState49Action,
	{_State452, AccessExprType}:                   _GotoState40Action,
	{_State452, PostfixUnaryExprType}:             _GotoState55Action,
	{_State452, PrefixUnaryOpType}:                _GotoState57Action,
	{_State452, PrefixUnaryExprType}:              _GotoState56Action,
	{_State452, MulExprType}:                      _GotoState52Action,
	{_State452, AddExprType}:                      _GotoState41Action,
	{_State452, CmpExprType}:                      _GotoState47Action,
	{_State452, AndExprType}:                      _GotoState42Action,
	{_State452, OrExprType}:                       _GotoState54Action,
	{_State452, InitializableTypeType}:            _GotoState50Action,
	{_State452, ExplicitStructDefType}:            _GotoState48Action,
	{_State452, AnonymousFuncExprType}:            _GotoState43Action,
	{_State453, IntegerLiteralToken}:              _GotoState26Action,
	{_State453, FloatLiteralToken}:                _GotoState22Action,
	{_State453, RuneLiteralToken}:                 _GotoState34Action,
	{_State453, StringLiteralToken}:               _GotoState35Action,
	{_State453, IdentifierToken}:                  _GotoState25Action,
	{_State453, TrueToken}:                        _GotoState38Action,
	{_State453, FalseToken}:                       _GotoState21Action,
	{_State453, StructToken}:                      _GotoState36Action,
	{_State453, FuncToken}:                        _GotoState23Action,
	{_State453, VarToken}:                         _GotoState39Action,
	{_State453, LetToken}:                         _GotoState29Action,
	{_State453, LabelDeclToken}:                   _GotoState27Action,
	{_State453, LparenToken}:                      _GotoState31Action,
	{_State453, LbracketToken}:                    _GotoState28Action,
	{_State453, NotToken}:                         _GotoState33Action,
	{_State453, SubToken}:                         _GotoState37Action,
	{_State453, MulToken}:                         _GotoState32Action,
	{_State453, BitNegToken}:                      _GotoState20Action,
	{_State453, BitAndToken}:                      _GotoState19Action,
	{_State453, GreaterToken}:                     _GotoState24Action,
	{_State453, LexErrorToken}:                    _GotoState30Action,
	{_State453, VarDeclPatternType}:               _GotoState59Action,
	{_State453, VarOrLetType}:                     _GotoState60Action,
	{_State453, OptionalLabelDeclType}:            _GotoState76Action,
	{_State453, SequenceExprType}:                 _GotoState460Action,
	{_State453, BlockExprType}:                    _GotoState45Action,
	{_State453, CallExprType}:                     _GotoState46Action,
	{_State453, AtomExprType}:                     _GotoState44Action,
	{_State453, LiteralType}:                      _GotoState51Action,
	{_State453, ImplicitStructExprType}:           _GotoState49Action,
	{_State453, AccessExprType}:                   _GotoState40Action,
	{_State453, PostfixUnaryExprType}:             _GotoState55Action,
	{_State453, PrefixUnaryOpType}:                _GotoState57Action,
	{_State453, PrefixUnaryExprType}:              _GotoState56Action,
	{_State453, MulExprType}:                      _GotoState52Action,
	{_State453, AddExprType}:                      _GotoState41Action,
	{_State453, CmpExprType}:                      _GotoState47Action,
	{_State453, AndExprType}:                      _GotoState42Action,
	{_State453, OrExprType}:                       _GotoState54Action,
	{_State453, InitializableTypeType}:            _GotoState50Action,
	{_State453, ExplicitStructDefType}:            _GotoState48Action,
	{_State453, AnonymousFuncExprType}:            _GotoState43Action,
	{_State456, RparenToken}:                      _GotoState461Action,
	{_State457, IdentifierToken}:                  _GotoState84Action,
	{_State457, StructToken}:                      _GotoState36Action,
	{_State457, EnumToken}:                        _GotoState81Action,
	{_State457, TraitToken}:                       _GotoState89Action,
	{_State457, FuncToken}:                        _GotoState83Action,
	{_State457, LparenToken}:                      _GotoState86Action,
	{_State457, LbracketToken}:                    _GotoState28Action,
	{_State457, DotToken}:                         _GotoState80Action,
	{_State457, QuestionToken}:                    _GotoState87Action,
	{_State457, ExclaimToken}:                     _GotoState82Action,
	{_State457, TildeTildeToken}:                  _GotoState88Action,
	{_State457, BitNegToken}:                      _GotoState79Action,
	{_State457, BitAndToken}:                      _GotoState78Action,
	{_State457, LexErrorToken}:                    _GotoState85Action,
	{_State457, InitializableTypeType}:            _GotoState95Action,
	{_State457, AtomTypeType}:                     _GotoState90Action,
	{_State457, ReturnableTypeType}:               _GotoState358Action,
	{_State457, ImplicitStructDefType}:            _GotoState94Action,
	{_State457, ExplicitStructDefType}:            _GotoState48Action,
	{_State457, ImplicitEnumDefType}:              _GotoState93Action,
	{_State457, ExplicitEnumDefType}:              _GotoState91Action,
	{_State457, TraitDefType}:                     _GotoState97Action,
	{_State457, ReturnTypeType}:                   _GotoState462Action,
	{_State457, FuncTypeType}:                     _GotoState92Action,
	{_State458, LbraceToken}:                      _GotoState63Action,
	{_State458, BlockBodyType}:                    _GotoState463Action,
	{_State461, IdentifierToken}:                  _GotoState84Action,
	{_State461, StructToken}:                      _GotoState36Action,
	{_State461, EnumToken}:                        _GotoState81Action,
	{_State461, TraitToken}:                       _GotoState89Action,
	{_State461, FuncToken}:                        _GotoState83Action,
	{_State461, LparenToken}:                      _GotoState86Action,
	{_State461, LbracketToken}:                    _GotoState28Action,
	{_State461, DotToken}:                         _GotoState80Action,
	{_State461, QuestionToken}:                    _GotoState87Action,
	{_State461, ExclaimToken}:                     _GotoState82Action,
	{_State461, TildeTildeToken}:                  _GotoState88Action,
	{_State461, BitNegToken}:                      _GotoState79Action,
	{_State461, BitAndToken}:                      _GotoState78Action,
	{_State461, LexErrorToken}:                    _GotoState85Action,
	{_State461, InitializableTypeType}:            _GotoState95Action,
	{_State461, AtomTypeType}:                     _GotoState90Action,
	{_State461, ReturnableTypeType}:               _GotoState358Action,
	{_State461, ImplicitStructDefType}:            _GotoState94Action,
	{_State461, ExplicitStructDefType}:            _GotoState48Action,
	{_State461, ImplicitEnumDefType}:              _GotoState93Action,
	{_State461, ExplicitEnumDefType}:              _GotoState91Action,
	{_State461, TraitDefType}:                     _GotoState97Action,
	{_State461, ReturnTypeType}:                   _GotoState464Action,
	{_State461, FuncTypeType}:                     _GotoState92Action,
	{_State464, LbraceToken}:                      _GotoState63Action,
	{_State464, BlockBodyType}:                    _GotoState465Action,
	{_State1, _EndMarker}:                         _ReduceNilToOptionalDefinitionsAction,
	{_State1, _WildcardMarker}:                    _ReduceNilToOptionalNewlinesAction,
	{_State5, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State13, _EndMarker}:                        _ReduceNewlinesToOptionalDefinitionsAction,
	{_State13, _WildcardMarker}:                   _ReduceNewlinesToOptionalNewlinesAction,
	{_State14, _EndMarker}:                        _ReduceToSourceAction,
	{_State16, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State19, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryOpAction,
	{_State20, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryOpAction,
	{_State21, _WildcardMarker}:                   _ReduceFalseToLiteralAction,
	{_State22, _WildcardMarker}:                   _ReduceFloatLiteralToLiteralAction,
	{_State24, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State25, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State26, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralAction,
	{_State27, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State29, _WildcardMarker}:                   _ReduceLetToVarOrLetAction,
	{_State30, _WildcardMarker}:                   _ReduceLexErrorToAtomExprAction,
	{_State31, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State31, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State32, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State33, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State34, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State35, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State37, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State38, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
	{_State39, _WildcardMarker}:                   _ReduceVarToVarOrLetAction,
	{_State40, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State40, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State41, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State42, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State43, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State44, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State45, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State46, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State47, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State48, _WildcardMarker}:                   _ReduceExplicitStructDefToInitializableTypeAction,
	{_State49, _WildcardMarker}:                   _ReduceImplicitStructExprToAtomExprAction,
	{_State51, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State52, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State54, _WildcardMarker}:                   _ReduceOrExprToSequenceExprAction,
	{_State55, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State56, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State57, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State58, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State59, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State61, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State62, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State63, _WildcardMarker}:                   _ReduceEmptyListToStatementsAction,
	{_State64, _WildcardMarker}:                   _ReduceBlockBodyToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State66, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State67, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State68, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State69, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State70, _WildcardMarker}:                   _ReduceGlobalVarDefToDefinitionAction,
	{_State71, _WildcardMarker}:                   _ReduceEmptyListToPackageStatementsAction,
	{_State72, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State73, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State75, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State77, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State80, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State84, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State85, _WildcardMarker}:                   _ReduceLexErrorToAtomTypeAction,
	{_State86, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State90, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State91, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State92, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State93, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State94, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State95, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State96, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State97, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
	{_State99, _WildcardMarker}:                   _ReduceDotDotDotToArgumentAction,
	{_State100, _WildcardMarker}:                  _ReduceIdentifierToAtomExprAction,
	{_State101, _WildcardMarker}:                  _ReduceArgumentToArgumentsAction,
	{_State103, _WildcardMarker}:                  _ReduceColonExpressionsToArgumentAction,
	{_State104, _WildcardMarker}:                  _ReducePositionalToArgumentAction,
	{_State104, ColonToken}:                       _ReduceExpressionToOptionalExpressionAction,
	{_State106, RparenToken}:                      _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State107, RbracketToken}:                    _ReduceNilToOptionalGenericArgumentsAction,
	{_State109, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State109, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State110, _WildcardMarker}:                  _ReduceQuestionToPostfixUnaryExprAction,
	{_State112, _WildcardMarker}:                  _ReduceAddToAddOpAction,
	{_State113, _WildcardMarker}:                  _ReduceBitOrToAddOpAction,
	{_State114, _WildcardMarker}:                  _ReduceBitXorToAddOpAction,
	{_State115, _WildcardMarker}:                  _ReduceSubToAddOpAction,
	{_State116, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State117, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State118, _WildcardMarker}:                  _ReduceEqualToCmpOpAction,
	{_State119, _WildcardMarker}:                  _ReduceGreaterToCmpOpAction,
	{_State120, _WildcardMarker}:                  _ReduceGreaterOrEqualToCmpOpAction,
	{_State121, _WildcardMarker}:                  _ReduceLessToCmpOpAction,
	{_State122, _WildcardMarker}:                  _ReduceLessOrEqualToCmpOpAction,
	{_State123, _WildcardMarker}:                  _ReduceNotEqualToCmpOpAction,
	{_State124, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State125, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State125, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State126, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State127, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State128, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State129, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State130, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State131, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State132, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State134, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State135, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State136, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State137, _WildcardMarker}:                  _ReduceToBlockExprAction,
	{_State138, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State139, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State140, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State141, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State142, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State143, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State145, _WildcardMarker}:                  _ReduceTuplePatternToVarPatternAction,
	{_State146, _WildcardMarker}:                  _ReduceNilToOptionalValueTypeAction,
	{_State147, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State148, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State149, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State150, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State153, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State155, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State157, _WildcardMarker}:                  _ReduceInferredRefArgToParameterDefAction,
	{_State160, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State161, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State162, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State163, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State164, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State166, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State167, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State169, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State170, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State173, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State173, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State175, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State177, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State178, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State179, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State180, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State181, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State186, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
	{_State188, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State189, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State189, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State190, _WildcardMarker}:                  _ReduceToImplicitStructExprAction,
	{_State191, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State191, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State191, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State191, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State191, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State192, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State192, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State192, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State192, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State192, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State193, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State194, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State196, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State198, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State199, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State201, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State201, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State201, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State202, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State203, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State204, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State206, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State207, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State210, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State210, SemicolonToken}:                   _ReduceSequenceExprToForAssignmentAction,
	{_State211, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State213, LbraceToken}:                      _ReduceSequenceExprToConditionAction,
	{_State215, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State216, _WildcardMarker}:                  _ReduceDotDotDotToFieldVarPatternAction,
	{_State217, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State218, _WildcardMarker}:                  _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State220, _WildcardMarker}:                  _ReducePositionalToFieldVarPatternAction,
	{_State221, _EndMarker}:                       _ReduceToVarDeclPatternAction,
	{_State222, _WildcardMarker}:                  _ReduceValueTypeToOptionalValueTypeAction,
	{_State223, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State224, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State225, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State226, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State227, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State228, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State229, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State229, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State231, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State232, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State233, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State234, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State235, AssignToken}:                      _ReduceToAssignPatternAction,
	{_State235, _WildcardMarker}:                  _ReduceSequenceExprToExpressionAction,
	{_State236, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State238, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State239, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State240, _WildcardMarker}:                  _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State242, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State243, _WildcardMarker}:                  _ReduceImportStatementToPackageStatementBodyAction,
	{_State244, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State246, _WildcardMarker}:                  _ReduceUnsafeStatementToPackageStatementBodyAction,
	{_State247, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State248, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State249, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State250, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State252, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State253, _EndMarker}:                       _ReduceAliasToNamedFuncDefAction,
	{_State254, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State255, _WildcardMarker}:                  _ReduceInferredRefVarargToParameterDefAction,
	{_State256, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State258, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State262, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State265, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State267, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State268, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State269, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State270, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State271, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State272, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State277, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State279, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State281, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State282, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State284, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State285, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State286, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State289, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State290, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State291, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State292, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State293, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State294, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State295, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State298, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State300, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State301, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State302, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State304, _WildcardMarker}:                  _ReduceInitializeExprToAtomExprAction,
	{_State305, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State306, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State307, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State308, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State308, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State311, _WildcardMarker}:                  _ReduceVarToVarOrLetAction,
	{_State312, _WildcardMarker}:                  _ReduceAssignPatternToCasePatternAction,
	{_State313, _WildcardMarker}:                  _ReduceCasePatternToCasePatternsAction,
	{_State315, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State316, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State320, _WildcardMarker}:                  _ReduceToTuplePatternAction,
	{_State321, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State322, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State322, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State322, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State323, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State323, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State323, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State324, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State325, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State326, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State327, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State328, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State329, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State330, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State331, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State332, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State333, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State334, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State335, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State336, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State337, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State338, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State339, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State340, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State341, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State342, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State342, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State342, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State343, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State344, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State346, _WildcardMarker}:                  _ReduceStringLiteralToImportClauseAction,
	{_State347, _WildcardMarker}:                  _ReduceSingleToImportStatementAction,
	{_State348, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State349, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State350, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State352, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State355, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State356, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State358, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State359, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State362, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State365, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State367, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State368, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State370, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State372, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State373, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State374, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State375, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State377, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State380, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State381, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State382, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State383, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State384, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State385, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State386, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State387, SemicolonToken}:                   _ReduceAssignToForAssignmentAction,
	{_State390, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State391, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State394, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State395, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State397, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State398, _WildcardMarker}:                  _ReduceCaseBranchToCaseBranchesAction,
	{_State399, RbraceToken}:                      _ReduceNilToOptionalDefaultBranchAction,
	{_State400, _WildcardMarker}:                  _ReduceNamedToFieldVarPatternAction,
	{_State401, _WildcardMarker}:                  _ReduceAddToFieldVarPatternsAction,
	{_State402, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State403, _WildcardMarker}:                  _ReduceAssignStatementToStatementBodyAction,
	{_State404, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State405, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State406, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State408, _WildcardMarker}:                  _ReduceFirstToImportClausesAction,
	{_State411, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State412, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State413, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State415, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State416, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State417, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State417, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State418, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State419, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State419, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State420, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State421, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State422, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State424, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State425, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State426, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State428, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State428, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State429, _WildcardMarker}:                  _ReduceEnumMatchPatternToCasePatternAction,
	{_State431, LbraceToken}:                      _ReduceCaseToConditionAction,
	{_State432, _WildcardMarker}:                  _ReduceMultipleToCasePatternsAction,
	{_State433, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State434, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State437, _WildcardMarker}:                  _ReduceAddToCaseBranchesAction,
	{_State438, RbraceToken}:                      _ReduceDefaultBranchToOptionalDefaultBranchAction,
	{_State440, _WildcardMarker}:                  _ReduceExplicitToImportClauseTerminalAction,
	{_State441, _WildcardMarker}:                  _ReduceImplicitToImportClauseTerminalAction,
	{_State442, _WildcardMarker}:                  _ReduceMultipleToImportStatementAction,
	{_State443, _WildcardMarker}:                  _ReduceAddToImportClausesAction,
	{_State444, _WildcardMarker}:                  _ReduceAliasToImportClauseAction,
	{_State446, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State447, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State449, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State451, _WildcardMarker}:                  _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State452, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State453, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State454, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State455, _EndMarker}:                       _ReduceFuncDefToNamedFuncDefAction,
	{_State457, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State459, _WildcardMarker}:                  _ReduceToCaseBranchAction,
	{_State460, RbraceToken}:                      _ReduceToDefaultBranchAction,
	{_State461, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State462, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State463, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State465, _EndMarker}:                       _ReduceMethodDefToNamedFuncDefAction,
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
      NEWLINES -> State 13
      source -> State 7
      optional_definitions -> State 14
      optional_newlines -> State 15

  State 2:
    Kernel Items:
      #accept: ^.package_def
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 16
      package_def -> State 8

  State 3:
    Kernel Items:
      #accept: ^.type_def
    Reduce:
      (nil)
    Goto:
      TYPE -> State 17
      type_def -> State 9

  State 4:
    Kernel Items:
      #accept: ^.named_func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 18
      named_func_def -> State 10

  State 5:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 11
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 62
      COMMENT -> State 61
      lex_internal_tokens -> State 12

  State 7:
    Kernel Items:
      #accept: ^ source., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ package_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      #accept: ^ type_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      #accept: ^ named_func_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      #accept: ^ lex_internal_tokens., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      optional_definitions: NEWLINES., $
      optional_newlines: NEWLINES., *
    Reduce:
      * -> [optional_newlines]
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      source: optional_definitions., $
    Reduce:
      $ -> [source]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      optional_definitions: optional_newlines.definitions optional_newlines
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 16
      TYPE -> State 17
      FUNC -> State 18
      VAR -> State 39
      LET -> State 29
      LBRACE -> State 63
      definitions -> State 66
      definition -> State 65
      var_decl_pattern -> State 70
      var_or_let -> State 60
      block_body -> State 64
      type_def -> State 69
      named_func_def -> State 67
      package_def -> State 68

  State 16:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 71

  State 17:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 72

  State 18:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 73
      LPAREN -> State 74

  State 19:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 75

  State 24:
    Kernel Items:
      sequence_expr: GREATER.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 77
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 25:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      initializable_type: LBRACKET.value_type RBRACKET
      initializable_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 98
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 29:
    Kernel Items:
      var_or_let: LET., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 100
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT_DOT_DOT -> State 99
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 104
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      arguments -> State 102
      argument -> State 101
      colon_expressions -> State 103
      optional_expression -> State 105
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 32:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 106

  State 37:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      var_or_let: VAR., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 40:
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
      LBRACKET -> State 109
      DOT -> State 108
      QUESTION -> State 110
      DOLLAR_LBRACKET -> State 107
      optional_generic_binding -> State 111

  State 41:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 112
      SUB -> State 115
      BIT_XOR -> State 114
      BIT_OR -> State 113
      add_op -> State 116

  State 42:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 117

  State 43:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 118
      NOT_EQUAL -> State 123
      LESS -> State 121
      LESS_OR_EQUAL -> State 122
      GREATER -> State 119
      GREATER_OR_EQUAL -> State 120
      cmp_op -> State 124

  State 48:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      atom_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 125

  State 51:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 131
      DIV -> State 129
      MOD -> State 130
      BIT_AND -> State 126
      BIT_LSHIFT -> State 127
      BIT_RSHIFT -> State 128
      mul_op -> State 132

  State 53:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 135
      SWITCH -> State 136
      FOR -> State 134
      DO -> State 133
      LBRACE -> State 63
      if_expr -> State 138
      switch_expr -> State 140
      loop_expr -> State 139
      block_body -> State 137

  State 54:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 141

  State 55:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 142
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 58:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      LPAREN -> State 144
      var_pattern -> State 146
      tuple_pattern -> State 145

  State 61:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 147

  State 64:
    Kernel Items:
      definition: block_body., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 148
      optional_newlines -> State 149

  State 67:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 150

  State 71:
    Kernel Items:
      package_def: PACKAGE LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 151

  State 72:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 153
      ASSIGN -> State 152
      optional_generic_parameters -> State 154

  State 73:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 153
      ASSIGN -> State 155
      optional_generic_parameters -> State 156

  State 74:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      parameter_def -> State 158

  State 75:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 157
      parameter_def -> State 160
      parameter_defs -> State 161
      optional_parameter_defs -> State 159

  State 76:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 137

  State 77:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 162
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 79:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 163
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 80:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 107
      optional_generic_binding -> State 164

  State 81:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 165

  State 82:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 166
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 83:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 167

  State 84:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 168
      DOLLAR_LBRACKET -> State 107
      optional_generic_binding -> State 169

  State 85:
    Kernel Items:
      atom_type: LEX_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 173
      implicit_field_defs -> State 175
      optional_implicit_field_defs -> State 176
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 172
      implicit_enum_value_defs -> State 174
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 87:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 179
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 88:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 180
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 89:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 181

  State 90:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 98:
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
      RBRACKET -> State 186
      COMMA -> State 184
      COLON -> State 183
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 99:
    Kernel Items:
      argument: DOT_DOT_DOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 188

  State 101:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      arguments: arguments.COMMA argument
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 190
      COMMA -> State 189

  State 103:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 191

  State 104:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 192

  State 106:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 194
      implicit_struct_def -> State 94
      explicit_field_defs -> State 193
      optional_explicit_field_defs -> State 195
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 107:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      optional_generic_arguments -> State 197
      generic_arguments -> State 196
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 198
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 108:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 199

  State 109:
    Kernel Items:
      access_expr: access_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 100
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT_DOT_DOT -> State 99
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 104
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      argument -> State 200
      colon_expressions -> State 103
      optional_expression -> State 105
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 110:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 201

  State 112:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 202
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 117:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 203
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 118:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 204
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 125:
    Kernel Items:
      atom_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 100
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT_DOT_DOT -> State 99
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 104
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      arguments -> State 205
      argument -> State 101
      colon_expressions -> State 103
      optional_expression -> State 105
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 126:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 206
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 133:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 207

  State 134:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.assign_pattern IN sequence_expr DO block_body
      loop_expr: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      assign_pattern -> State 208
      optional_label_decl -> State 76
      for_assignment -> State 209
      sequence_expr -> State 210
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 135:
    Kernel Items:
      if_expr: IF.condition block_body
      if_expr: IF.condition block_body ELSE block_body
      if_expr: IF.condition block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      CASE -> State 211
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      condition -> State 212
      sequence_expr -> State 213
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 136:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 214
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 137:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 215
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 142:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

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
      IDENTIFIER -> State 217
      LPAREN -> State 144
      DOT_DOT_DOT -> State 216
      var_pattern -> State 220
      tuple_pattern -> State 145
      field_var_patterns -> State 219
      field_var_pattern -> State 218

  State 145:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      optional_value_type -> State 221
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 222
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 147:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
      statements: statements.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      RETURN -> State 228
      BREAK -> State 224
      CONTINUE -> State 225
      UNSAFE -> State 171
      STRUCT -> State 36
      FUNC -> State 23
      ASYNC -> State 223
      DEFER -> State 226
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      RBRACE -> State 227
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      assign_pattern -> State 230
      expression -> State 231
      optional_label_decl -> State 53
      sequence_expr -> State 235
      block_expr -> State 45
      statement -> State 236
      statement_body -> State 237
      unsafe_statement -> State 238
      jump_statement -> State 233
      jump_type -> State 234
      expressions -> State 232
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 229
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 148:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      PACKAGE -> State 16
      TYPE -> State 17
      FUNC -> State 18
      VAR -> State 39
      LET -> State 29
      LBRACE -> State 63
      definition -> State 239
      var_decl_pattern -> State 70
      var_or_let -> State 60
      block_body -> State 64
      type_def -> State 69
      named_func_def -> State 67
      package_def -> State 68

  State 149:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 150:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 240
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 151:
    Kernel Items:
      package_def: PACKAGE LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      IMPORT -> State 241
      UNSAFE -> State 171
      RPAREN -> State 242
      unsafe_statement -> State 246
      package_statement_body -> State 245
      package_statement -> State 244
      import_statement -> State 243

  State 152:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 247
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 153:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 248
      generic_parameter_def -> State 249
      generic_parameter_defs -> State 250
      optional_generic_parameter_defs -> State 251

  State 154:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 252
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 155:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 253
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 156:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 254

  State 157:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.DOT_DOT_DOT
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      DOT_DOT_DOT -> State 255
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 256
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 158:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 257

  State 159:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 258

  State 160:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 259

  State 162:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 165:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 260
      implicit_enum_value_defs -> State 263
      implicit_enum_def -> State 93
      explicit_enum_value_defs -> State 261
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 166:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 265
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      DOT_DOT_DOT -> State 264
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 269
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      parameter_decl -> State 267
      parameter_decls -> State 268
      optional_parameter_decls -> State 266
      func_type -> State 92

  State 168:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 270

  State 169:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 271
      QUESTION -> State 87
      EXCLAIM -> State 82
      DOLLAR_LBRACKET -> State 107
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      optional_generic_binding -> State 169
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 272
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 171:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 273

  State 172:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 274

  State 173:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 275

  State 174:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 277
      OR -> State 276

  State 175:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 278

  State 176:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 279

  State 177:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 179:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 280
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 281
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_property -> State 285
      trait_properties -> State 284
      optional_trait_properties -> State 283
      trait_def -> State 97
      func_type -> State 92
      method_signature -> State 282

  State 182:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 286
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 183:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 287
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 184:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 288

  State 185:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 289
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 186:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 290
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 188:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 291
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 189:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 100
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT_DOT_DOT -> State 99
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 104
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      argument -> State 292
      colon_expressions -> State 103
      optional_expression -> State 105
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 190:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      colon_expressions: colon_expressions COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 293
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      optional_expression -> State 294
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 192:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 293
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      optional_expression -> State 295
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 193:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 297
      COMMA -> State 296

  State 194:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 298

  State 196:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 299

  State 197:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 300

  State 198:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 199:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 301

  State 201:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 100
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT_DOT_DOT -> State 99
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 104
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      optional_arguments -> State 303
      arguments -> State 302
      argument -> State 101
      colon_expressions -> State 103
      optional_expression -> State 105
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 202:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 131
      DIV -> State 129
      MOD -> State 130
      BIT_AND -> State 126
      BIT_LSHIFT -> State 127
      BIT_RSHIFT -> State 128
      mul_op -> State 132

  State 203:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 118
      NOT_EQUAL -> State 123
      LESS -> State 121
      LESS_OR_EQUAL -> State 122
      GREATER -> State 119
      GREATER_OR_EQUAL -> State 120
      cmp_op -> State 124

  State 204:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 112
      SUB -> State 115
      BIT_XOR -> State 114
      BIT_OR -> State 113
      add_op -> State 116

  State 205:
    Kernel Items:
      arguments: arguments.COMMA argument
      atom_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 304
      COMMA -> State 189

  State 206:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 305

  State 208:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO block_body
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 307
      ASSIGN -> State 306

  State 209:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 308

  State 210:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO block_body
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 309

  State 211:
    Kernel Items:
      condition: CASE.case_patterns ASSIGN sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 311
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT -> State 310
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      assign_pattern -> State 312
      case_pattern -> State 313
      optional_label_decl -> State 76
      case_patterns -> State 314
      sequence_expr -> State 315
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 212:
    Kernel Items:
      if_expr: IF condition.block_body
      if_expr: IF condition.block_body ELSE block_body
      if_expr: IF condition.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 316

  State 213:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 317

  State 215:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 117

  State 216:
    Kernel Items:
      field_var_pattern: DOT_DOT_DOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 318

  State 218:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 320
      COMMA -> State 319

  State 220:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      optional_value_type: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 223:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 322
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 321
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 224:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      statement_body: DEFER.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      LEX_ERROR -> State 30
      optional_label_decl -> State 76
      block_expr -> State 45
      call_expr -> State 323
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 321
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 227:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 229:
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
      LBRACKET -> State 109
      DOT -> State 108
      QUESTION -> State 110
      DOLLAR_LBRACKET -> State 107
      ADD_ASSIGN -> State 324
      SUB_ASSIGN -> State 335
      MUL_ASSIGN -> State 334
      DIV_ASSIGN -> State 332
      MOD_ASSIGN -> State 333
      ADD_ONE_ASSIGN -> State 325
      SUB_ONE_ASSIGN -> State 336
      BIT_NEG_ASSIGN -> State 328
      BIT_AND_ASSIGN -> State 326
      BIT_OR_ASSIGN -> State 329
      BIT_XOR_ASSIGN -> State 331
      BIT_LSHIFT_ASSIGN -> State 327
      BIT_RSHIFT_ASSIGN -> State 330
      unary_op_assign -> State 338
      binary_op_assign -> State 337
      optional_generic_binding -> State 111

  State 230:
    Kernel Items:
      statement_body: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 339

  State 231:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 340

  State 233:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 341
      optional_jump_label -> State 342

  State 235:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 343
      SEMICOLON -> State 344

  State 238:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 346
      LPAREN -> State 345
      import_clause -> State 347

  State 242:
    Kernel Items:
      package_def: PACKAGE LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      package_statement_body: import_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 349
      COMMA -> State 348

  State 246:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 248:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 350
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 249:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 351

  State 251:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 352

  State 252:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 353
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 253:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 157
      parameter_def -> State 160
      parameter_defs -> State 161
      optional_parameter_defs -> State 354

  State 255:
    Kernel Items:
      parameter_def: IDENTIFIER DOT_DOT_DOT., *
      parameter_def: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 355
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 256:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 257:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 356

  State 258:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 358
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      return_type -> State 357
      func_type -> State 92

  State 259:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      parameter_def -> State 359

  State 260:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 360
      OR -> State 361

  State 261:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 362

  State 262:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 275

  State 263:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 363
      OR -> State 364

  State 264:
    Kernel Items:
      parameter_decl: DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 365
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 265:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 271
      QUESTION -> State 87
      EXCLAIM -> State 82
      DOLLAR_LBRACKET -> State 107
      DOT_DOT_DOT -> State 366
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      optional_generic_binding -> State 169
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 367
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 266:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 368

  State 267:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 369

  State 269:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 270:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 107
      optional_generic_binding -> State 370

  State 271:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 270
      DOLLAR_LBRACKET -> State 107
      optional_generic_binding -> State 164

  State 272:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 273:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 371

  State 274:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 372
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 275:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 373

  State 276:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 374
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 277:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 375
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 279:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 376
      LPAREN -> State 167

  State 281:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 377

  State 284:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 379
      COMMA -> State 378

  State 285:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 380
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 288:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 381

  State 289:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 382
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 297:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 383
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 298:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 384
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 300:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 189

  State 303:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 385

  State 304:
    Kernel Items:
      atom_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 386
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 306:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 387
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 307:
    Kernel Items:
      loop_expr: FOR assign_pattern IN.sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 388
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 308:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      optional_sequence_expr -> State 389
      sequence_expr -> State 390
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 309:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 391

  State 310:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 392

  State 311:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 393

  State 312:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      condition: CASE case_patterns.ASSIGN sequence_expr
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 395
      ASSIGN -> State 394

  State 315:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      if_expr: IF condition block_body., *
      if_expr: IF condition block_body.ELSE block_body
      if_expr: IF condition block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 396

  State 317:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 397
      case_branches -> State 399
      case_branch -> State 398

  State 318:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      LPAREN -> State 144
      var_pattern -> State 400
      tuple_pattern -> State 145

  State 319:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 217
      LPAREN -> State 144
      DOT_DOT_DOT -> State 216
      var_pattern -> State 220
      tuple_pattern -> State 145
      field_var_pattern -> State 401

  State 320:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 109
      DOT -> State 108
      DOLLAR_LBRACKET -> State 107
      optional_generic_binding -> State 111

  State 322:
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

  State 323:
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

  State 324:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 402
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 338:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      statement_body: assign_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 403
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 340:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 404
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 341:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      expression -> State 231
      optional_label_decl -> State 53
      sequence_expr -> State 58
      block_expr -> State 45
      expressions -> State 405
      optional_expressions -> State 406
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 343:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 346
      import_clause -> State 407
      import_clause_terminal -> State 408
      import_clauses -> State 409

  State 346:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 410

  State 347:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 348:
    Kernel Items:
      package_statement: package_statement_body COMMA., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 351:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 248
      generic_parameter_def -> State 411

  State 352:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 412
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 354:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 413

  State 355:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 356:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 153
      optional_generic_parameters -> State 414

  State 357:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 415

  State 358:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 360:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 416
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 361:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 417
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 362:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 418
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 364:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 262
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      enum_value_def -> State 419
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 365:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 366:
    Kernel Items:
      parameter_decl: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 420
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      func_type -> State 92

  State 367:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 368:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 358
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      return_type -> State 421
      func_type -> State 92

  State 369:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 265
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      DOT_DOT_DOT -> State 264
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 269
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      parameter_decl -> State 422
      func_type -> State 92

  State 370:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 423

  State 372:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 373:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 424

  State 377:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 280
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 281
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_property -> State 425
      trait_def -> State 97
      func_type -> State 92
      method_signature -> State 282

  State 379:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      UNSAFE -> State 171
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 280
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      unsafe_statement -> State 177
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 178
      field_def -> State 281
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_property -> State 426
      trait_def -> State 97
      func_type -> State 92
      method_signature -> State 282

  State 380:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 383:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 385:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 427

  State 389:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 428

  State 390:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 31
      implicit_struct_expr -> State 429

  State 393:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 430

  State 394:
    Kernel Items:
      condition: CASE case_patterns ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 431
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 395:
    Kernel Items:
      case_patterns: case_patterns COMMA.case_pattern
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 311
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT -> State 310
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      assign_pattern -> State 312
      case_pattern -> State 432
      optional_label_decl -> State 76
      sequence_expr -> State 315
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 396:
    Kernel Items:
      if_expr: IF condition block_body ELSE.block_body
      if_expr: IF condition block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 135
      LBRACE -> State 63
      if_expr -> State 434
      block_body -> State 433

  State 397:
    Kernel Items:
      case_branch: CASE.case_patterns COLON sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 311
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      DOT -> State 310
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      assign_pattern -> State 312
      case_pattern -> State 313
      optional_label_decl -> State 76
      case_patterns -> State 435
      sequence_expr -> State 315
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 398:
    Kernel Items:
      case_branches: case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches.optional_default_branch RBRACE
      case_branches: case_branches.case_branch
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      CASE -> State 397
      DEFAULT -> State 436
      case_branch -> State 437
      optional_default_branch -> State 439
      default_branch -> State 438

  State 400:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      statement_body: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 404:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 340

  State 406:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 407:
    Kernel Items:
      import_clause_terminal: import_clause.NEWLINES
      import_clause_terminal: import_clause.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 441
      COMMA -> State 440

  State 408:
    Kernel Items:
      import_clauses: import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 409:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
      import_clauses: import_clauses.import_clause_terminal
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 346
      RPAREN -> State 442
      import_clause -> State 407
      import_clause_terminal -> State 443

  State 410:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 444

  State 411:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 413:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 358
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      return_type -> State 445
      func_type -> State 92

  State 414:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 446

  State 415:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 182
      SUB -> State 187
      MUL -> State 185

  State 421:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 422:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 447

  State 424:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 265
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      DOT_DOT_DOT -> State 264
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 96
      value_type -> State 269
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      parameter_decl -> State 267
      parameter_decls -> State 268
      optional_parameter_decls -> State 448
      func_type -> State 92

  State 425:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 449

  State 428:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      optional_sequence_expr -> State 450
      sequence_expr -> State 390
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 429:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 144
      tuple_pattern -> State 451

  State 431:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 433:
    Kernel Items:
      if_expr: IF condition block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      if_expr: IF condition block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      case_branch: CASE case_patterns.COLON sequence_expr
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 395
      COLON -> State 452

  State 436:
    Kernel Items:
      default_branch: DEFAULT.COLON sequence_expr
    Reduce:
      (nil)
    Goto:
      COLON -> State 453

  State 437:
    Kernel Items:
      case_branches: case_branches case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      optional_default_branch: default_branch., RBRACE
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 454

  State 440:
    Kernel Items:
      import_clause_terminal: import_clause COMMA., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      import_clause_terminal: import_clause NEWLINES., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 443:
    Kernel Items:
      import_clauses: import_clauses import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 445:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 455

  State 446:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 157
      parameter_def -> State 160
      parameter_defs -> State 161
      optional_parameter_defs -> State 456

  State 447:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 457

  State 449:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 450:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 458

  State 451:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 452:
    Kernel Items:
      case_branch: CASE case_patterns COLON.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 459
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 453:
    Kernel Items:
      default_branch: DEFAULT COLON.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 34
      STRING_LITERAL -> State 35
      IDENTIFIER -> State 25
      TRUE -> State 38
      FALSE -> State 21
      STRUCT -> State 36
      FUNC -> State 23
      VAR -> State 39
      LET -> State 29
      LABEL_DECL -> State 27
      LPAREN -> State 31
      LBRACKET -> State 28
      NOT -> State 33
      SUB -> State 37
      MUL -> State 32
      BIT_NEG -> State 20
      BIT_AND -> State 19
      GREATER -> State 24
      LEX_ERROR -> State 30
      var_decl_pattern -> State 59
      var_or_let -> State 60
      optional_label_decl -> State 76
      sequence_expr -> State 460
      block_expr -> State 45
      call_expr -> State 46
      atom_expr -> State 44
      literal -> State 51
      implicit_struct_expr -> State 49
      access_expr -> State 40
      postfix_unary_expr -> State 55
      prefix_unary_op -> State 57
      prefix_unary_expr -> State 56
      mul_expr -> State 52
      add_expr -> State 41
      cmp_expr -> State 47
      and_expr -> State 42
      or_expr -> State 54
      initializable_type -> State 50
      explicit_struct_def -> State 48
      anonymous_func_expr -> State 43

  State 454:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 455:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 461

  State 457:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 358
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      return_type -> State 462
      func_type -> State 92

  State 458:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 463

  State 459:
    Kernel Items:
      case_branch: CASE case_patterns COLON sequence_expr., *
    Reduce:
      * -> [case_branch]
    Goto:
      (nil)

  State 460:
    Kernel Items:
      default_branch: DEFAULT COLON sequence_expr., RBRACE
    Reduce:
      RBRACE -> [default_branch]
    Goto:
      (nil)

  State 461:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 84
      STRUCT -> State 36
      ENUM -> State 81
      TRAIT -> State 89
      FUNC -> State 83
      LPAREN -> State 86
      LBRACKET -> State 28
      DOT -> State 80
      QUESTION -> State 87
      EXCLAIM -> State 82
      TILDE_TILDE -> State 88
      BIT_NEG -> State 79
      BIT_AND -> State 78
      LEX_ERROR -> State 85
      initializable_type -> State 95
      atom_type -> State 90
      returnable_type -> State 358
      implicit_struct_def -> State 94
      explicit_struct_def -> State 48
      implicit_enum_def -> State 93
      explicit_enum_def -> State 91
      trait_def -> State 97
      return_type -> State 464
      func_type -> State 92

  State 462:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 463:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 464:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 63
      block_body -> State 465

  State 465:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 465
Number of shift actions: 3196
Number of reduce actions: 374
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 4276
Number of unoptimized shift actions: 32209
Number of unoptimized reduce actions: 24780
*/
