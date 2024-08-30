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
	DotdotdotToken       = SymbolId(306)
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

	// 102:2: field_var_pattern -> DOTDOTDOT: ...
	DotdotdotToFieldVarPattern(Dotdotdot_ *GenericSymbol) (*GenericSymbol, error)

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

	// 230:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 232:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 235:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 236:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 239:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 240:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 243:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 245:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 265:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 272:2: statement_body -> assign_statement: ...
	AssignStatementToStatementBody(AssignPattern_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 276:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 277:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 284:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 285:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 291:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 295:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 296:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 306:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 313:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 316:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 317:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 318:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 321:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 322:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 325:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 326:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 329:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 336:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 343:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 348:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 349:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 352:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 356:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 360:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 365:2: argument -> DOTDOTDOT: ...
	DotdotdotToArgument(Dotdotdot_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 369:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 383:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 386:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 387:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializableType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: atom_expr -> implicit_struct_expr: ...
	ImplicitStructExprToAtomExpr(ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 392:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 393:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 397:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 399:24: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 402:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 414:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 417:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 420:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 423:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 427:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 428:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 429:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 430:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 431:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 432:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 435:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 436:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 439:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 440:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 441:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 442:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 445:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 449:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 451:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 452:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 453:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 457:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 461:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 462:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 465:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 466:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 475:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 476:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 477:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Comma_ *GenericSymbol, IntegerLiteral_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 478:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Colon_ *GenericSymbol, ValueType_2 *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 481:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ *GenericSymbol) (*GenericSymbol, error)

	// 482:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 483:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_2 *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 486:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 487:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 490:2: atom_type -> LEX_ERROR: ...
	LexErrorToAtomType(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 496:2: returnable_type -> atom_type: ...
	AtomTypeToReturnableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 497:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 498:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 499:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 500:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 501:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 506:2: value_type -> returnable_type: ...
	ReturnableTypeToValueType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ *GenericSymbol, Mul_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 508:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ *GenericSymbol, Add_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 509:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ *GenericSymbol, Sub_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 512:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 513:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 514:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Assign_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 522:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 523:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 526:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 527:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 530:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 531:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 534:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 535:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 542:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 543:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 544:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 547:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 548:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 551:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 552:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 554:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 557:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 558:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 559:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 562:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 563:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 565:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 573:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 574:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 586:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 587:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 589:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 592:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 593:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 594:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 595:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 597:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 604:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 605:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 608:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 609:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 610:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 613:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 614:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 616:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 624:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 625:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 628:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 629:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 630:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 631:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 634:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 635:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 638:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 639:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 641:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 652:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 658:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 659:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol) (*GenericSymbol, error)

	// 660:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 661:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 664:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 665:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 668:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 669:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 672:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 673:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ *GenericSymbol, Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_2 *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_2 *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 674:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 678:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 690:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol) (*GenericSymbol, error)

	// 691:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 694:2: package_statement_body -> unsafe_statement: ...
	UnsafeStatementToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 695:2: package_statement_body -> import_statement: ...
	ImportStatementToPackageStatementBody(ImportStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 698:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 699:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 702:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 703:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 706:2: import_statement -> single: ...
	SingleToImportStatement(Import_ *GenericSymbol, ImportClause_ *GenericSymbol) (*GenericSymbol, error)

	// 707:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ *GenericSymbol, Lparen_ *GenericSymbol, ImportClauses_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 710:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 711:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ *GenericSymbol, As_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 714:2: import_clause_terminal -> implicit: ...
	ImplicitToImportClauseTerminal(ImportClause_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 715:2: import_clause_terminal -> explicit: ...
	ExplicitToImportClauseTerminal(ImportClause_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 718:2: import_clauses -> first: ...
	FirstToImportClauses(ImportClauseTerminal_ *GenericSymbol) (*GenericSymbol, error)

	// 719:2: import_clauses -> add: ...
	AddToImportClauses(ImportClauses_ *GenericSymbol, ImportClauseTerminal_ *GenericSymbol) (*GenericSymbol, error)

	// 723:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 724:2: lex_internal_tokens -> COMMENT: ...
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
	case DotdotdotToken:
		return "DOTDOTDOT"
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
	_ReduceDotdotdotToFieldVarPattern                         = _ReduceType(25)
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
	_ReduceToBlockExpr                                        = _ReduceType(63)
	_ReduceToBlockBody                                        = _ReduceType(64)
	_ReduceEmptyListToStatements                              = _ReduceType(65)
	_ReduceAddToStatements                                    = _ReduceType(66)
	_ReduceImplicitToStatement                                = _ReduceType(67)
	_ReduceExplicitToStatement                                = _ReduceType(68)
	_ReduceUnsafeStatementToStatementBody                     = _ReduceType(69)
	_ReduceExpressionOrImplicitStructToStatementBody          = _ReduceType(70)
	_ReduceAsyncToStatementBody                               = _ReduceType(71)
	_ReduceDeferToStatementBody                               = _ReduceType(72)
	_ReduceJumpStatementToStatementBody                       = _ReduceType(73)
	_ReduceAssignStatementToStatementBody                     = _ReduceType(74)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(75)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(76)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(77)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(78)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(79)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(80)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(81)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(82)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(83)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(84)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(85)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(86)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(87)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(88)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(89)
	_ReduceToUnsafeStatement                                  = _ReduceType(90)
	_ReduceToJumpStatement                                    = _ReduceType(91)
	_ReduceReturnToJumpType                                   = _ReduceType(92)
	_ReduceBreakToJumpType                                    = _ReduceType(93)
	_ReduceContinueToJumpType                                 = _ReduceType(94)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(95)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(96)
	_ReduceExpressionToExpressions                            = _ReduceType(97)
	_ReduceAddToExpressions                                   = _ReduceType(98)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(99)
	_ReduceNilToOptionalExpressions                           = _ReduceType(100)
	_ReduceToCallExpr                                         = _ReduceType(101)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(102)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(103)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(104)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(105)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(106)
	_ReduceAddToGenericArguments                              = _ReduceType(107)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(108)
	_ReduceNilToOptionalArguments                             = _ReduceType(109)
	_ReduceArgumentToArguments                                = _ReduceType(110)
	_ReduceAddToArguments                                     = _ReduceType(111)
	_ReducePositionalToArgument                               = _ReduceType(112)
	_ReduceNamedToArgument                                    = _ReduceType(113)
	_ReduceColonExpressionsToArgument                         = _ReduceType(114)
	_ReduceDotdotdotToArgument                                = _ReduceType(115)
	_ReducePairToColonExpressions                             = _ReduceType(116)
	_ReduceAddToColonExpressions                              = _ReduceType(117)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(118)
	_ReduceNilToOptionalExpression                            = _ReduceType(119)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(120)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(121)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(122)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(123)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(124)
	_ReduceImplicitStructExprToAtomExpr                       = _ReduceType(125)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(126)
	_ReduceTrueToLiteral                                      = _ReduceType(127)
	_ReduceFalseToLiteral                                     = _ReduceType(128)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(129)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(130)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(131)
	_ReduceStringLiteralToLiteral                             = _ReduceType(132)
	_ReduceToImplicitStructExpr                               = _ReduceType(133)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(134)
	_ReduceAccessToAccessExpr                                 = _ReduceType(135)
	_ReduceCallExprToAccessExpr                               = _ReduceType(136)
	_ReduceIndexToAccessExpr                                  = _ReduceType(137)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(138)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(139)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(140)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(141)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(142)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(143)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(144)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(145)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(146)
	_ReduceMulToMulOp                                         = _ReduceType(147)
	_ReduceDivToMulOp                                         = _ReduceType(148)
	_ReduceModToMulOp                                         = _ReduceType(149)
	_ReduceBitAndToMulOp                                      = _ReduceType(150)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(151)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(152)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(153)
	_ReduceOpToMulExpr                                        = _ReduceType(154)
	_ReduceAddToAddOp                                         = _ReduceType(155)
	_ReduceSubToAddOp                                         = _ReduceType(156)
	_ReduceBitOrToAddOp                                       = _ReduceType(157)
	_ReduceBitXorToAddOp                                      = _ReduceType(158)
	_ReduceMulExprToAddExpr                                   = _ReduceType(159)
	_ReduceOpToAddExpr                                        = _ReduceType(160)
	_ReduceEqualToCmpOp                                       = _ReduceType(161)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(162)
	_ReduceLessToCmpOp                                        = _ReduceType(163)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(164)
	_ReduceGreaterToCmpOp                                     = _ReduceType(165)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(166)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(167)
	_ReduceOpToCmpExpr                                        = _ReduceType(168)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(169)
	_ReduceOpToAndExpr                                        = _ReduceType(170)
	_ReduceAndExprToOrExpr                                    = _ReduceType(171)
	_ReduceOpToOrExpr                                         = _ReduceType(172)
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(173)
	_ReduceSliceToInitializableType                           = _ReduceType(174)
	_ReduceArrayToInitializableType                           = _ReduceType(175)
	_ReduceMapToInitializableType                             = _ReduceType(176)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(177)
	_ReduceNamedToAtomType                                    = _ReduceType(178)
	_ReduceExternNamedToAtomType                              = _ReduceType(179)
	_ReduceInferredToAtomType                                 = _ReduceType(180)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(181)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(182)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(183)
	_ReduceTraitDefToAtomType                                 = _ReduceType(184)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(185)
	_ReduceLexErrorToAtomType                                 = _ReduceType(186)
	_ReduceAtomTypeToReturnableType                           = _ReduceType(187)
	_ReduceOptionalToReturnableType                           = _ReduceType(188)
	_ReduceResultToReturnableType                             = _ReduceType(189)
	_ReduceReferenceToReturnableType                          = _ReduceType(190)
	_ReducePublicMethodsTraitToReturnableType                 = _ReduceType(191)
	_ReducePublicTraitToReturnableType                        = _ReduceType(192)
	_ReduceReturnableTypeToValueType                          = _ReduceType(193)
	_ReduceTraitIntersectToValueType                          = _ReduceType(194)
	_ReduceTraitUnionToValueType                              = _ReduceType(195)
	_ReduceTraitDifferenceToValueType                         = _ReduceType(196)
	_ReduceDefinitionToTypeDef                                = _ReduceType(197)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(198)
	_ReduceAliasToTypeDef                                     = _ReduceType(199)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(200)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(201)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(202)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(203)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(204)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(205)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(206)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(207)
	_ReduceExplicitToFieldDef                                 = _ReduceType(208)
	_ReduceImplicitToFieldDef                                 = _ReduceType(209)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(210)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(211)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(212)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(213)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(214)
	_ReduceToImplicitStructDef                                = _ReduceType(215)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(216)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(217)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(218)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(219)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(220)
	_ReduceToExplicitStructDef                                = _ReduceType(221)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(222)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(223)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(224)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(225)
	_ReduceToImplicitEnumDef                                  = _ReduceType(226)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(227)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(228)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(229)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(230)
	_ReduceToExplicitEnumDef                                  = _ReduceType(231)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(232)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(233)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(234)
	_ReduceImplicitToTraitProperties                          = _ReduceType(235)
	_ReduceExplicitToTraitProperties                          = _ReduceType(236)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(237)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(238)
	_ReduceToTraitDef                                         = _ReduceType(239)
	_ReduceReturnableTypeToReturnType                         = _ReduceType(240)
	_ReduceNilToReturnType                                    = _ReduceType(241)
	_ReduceArgToParameterDecl                                 = _ReduceType(242)
	_ReduceVarargToParameterDecl                              = _ReduceType(243)
	_ReduceUnamedToParameterDecl                              = _ReduceType(244)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(245)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(246)
	_ReduceAddToParameterDecls                                = _ReduceType(247)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(248)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(249)
	_ReduceToFuncType                                         = _ReduceType(250)
	_ReduceToMethodSignature                                  = _ReduceType(251)
	_ReduceInferredRefArgToParameterDef                       = _ReduceType(252)
	_ReduceInferredRefVarargToParameterDef                    = _ReduceType(253)
	_ReduceArgToParameterDef                                  = _ReduceType(254)
	_ReduceVarargToParameterDef                               = _ReduceType(255)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(256)
	_ReduceAddToParameterDefs                                 = _ReduceType(257)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(258)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(259)
	_ReduceFuncDefToNamedFuncDef                              = _ReduceType(260)
	_ReduceMethodDefToNamedFuncDef                            = _ReduceType(261)
	_ReduceAliasToNamedFuncDef                                = _ReduceType(262)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(263)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(264)
	_ReduceWithSpecToPackageDef                               = _ReduceType(265)
	_ReduceUnsafeStatementToPackageStatementBody              = _ReduceType(266)
	_ReduceImportStatementToPackageStatementBody              = _ReduceType(267)
	_ReduceImplicitToPackageStatement                         = _ReduceType(268)
	_ReduceExplicitToPackageStatement                         = _ReduceType(269)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(270)
	_ReduceAddToPackageStatements                             = _ReduceType(271)
	_ReduceSingleToImportStatement                            = _ReduceType(272)
	_ReduceMultipleToImportStatement                          = _ReduceType(273)
	_ReduceStringLiteralToImportClause                        = _ReduceType(274)
	_ReduceAliasToImportClause                                = _ReduceType(275)
	_ReduceImplicitToImportClauseTerminal                     = _ReduceType(276)
	_ReduceExplicitToImportClauseTerminal                     = _ReduceType(277)
	_ReduceFirstToImportClauses                               = _ReduceType(278)
	_ReduceAddToImportClauses                                 = _ReduceType(279)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(280)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(281)
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
	case _ReduceDotdotdotToFieldVarPattern:
		return "DotdotdotToFieldVarPattern"
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
	case _ReduceDotdotdotToArgument:
		return "DotdotdotToArgument"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceDotdotdotToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.DotdotdotToFieldVarPattern(args[0].Generic_)
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
	case _ReduceDotdotdotToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.DotdotdotToArgument(args[0].Generic_)
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
	_ReduceDotdotdotToFieldVarPatternAction                         = &_Action{_ReduceAction, 0, _ReduceDotdotdotToFieldVarPattern}
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
	_ReduceDotdotdotToArgumentAction                                = &_Action{_ReduceAction, 0, _ReduceDotdotdotToArgument}
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
	{_State5, IntegerLiteralToken}:                _GotoState25Action,
	{_State5, FloatLiteralToken}:                  _GotoState22Action,
	{_State5, RuneLiteralToken}:                   _GotoState33Action,
	{_State5, StringLiteralToken}:                 _GotoState34Action,
	{_State5, IdentifierToken}:                    _GotoState24Action,
	{_State5, TrueToken}:                          _GotoState37Action,
	{_State5, FalseToken}:                         _GotoState21Action,
	{_State5, StructToken}:                        _GotoState35Action,
	{_State5, FuncToken}:                          _GotoState23Action,
	{_State5, VarToken}:                           _GotoState38Action,
	{_State5, LetToken}:                           _GotoState28Action,
	{_State5, LabelDeclToken}:                     _GotoState26Action,
	{_State5, LparenToken}:                        _GotoState30Action,
	{_State5, LbracketToken}:                      _GotoState27Action,
	{_State5, NotToken}:                           _GotoState32Action,
	{_State5, SubToken}:                           _GotoState36Action,
	{_State5, MulToken}:                           _GotoState31Action,
	{_State5, BitNegToken}:                        _GotoState20Action,
	{_State5, BitAndToken}:                        _GotoState19Action,
	{_State5, LexErrorToken}:                      _GotoState29Action,
	{_State5, VarDeclPatternType}:                 _GotoState58Action,
	{_State5, VarOrLetType}:                       _GotoState59Action,
	{_State5, ExpressionType}:                     _GotoState11Action,
	{_State5, OptionalLabelDeclType}:              _GotoState52Action,
	{_State5, SequenceExprType}:                   _GotoState57Action,
	{_State5, BlockExprType}:                      _GotoState44Action,
	{_State5, CallExprType}:                       _GotoState45Action,
	{_State5, AtomExprType}:                       _GotoState43Action,
	{_State5, LiteralType}:                        _GotoState50Action,
	{_State5, ImplicitStructExprType}:             _GotoState48Action,
	{_State5, AccessExprType}:                     _GotoState39Action,
	{_State5, PostfixUnaryExprType}:               _GotoState54Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState56Action,
	{_State5, PrefixUnaryExprType}:                _GotoState55Action,
	{_State5, MulExprType}:                        _GotoState51Action,
	{_State5, AddExprType}:                        _GotoState40Action,
	{_State5, CmpExprType}:                        _GotoState46Action,
	{_State5, AndExprType}:                        _GotoState41Action,
	{_State5, OrExprType}:                         _GotoState53Action,
	{_State5, InitializableTypeType}:              _GotoState49Action,
	{_State5, ExplicitStructDefType}:              _GotoState47Action,
	{_State5, AnonymousFuncExprType}:              _GotoState42Action,
	{_State6, SpacesToken}:                        _GotoState61Action,
	{_State6, CommentToken}:                       _GotoState60Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State15, PackageToken}:                      _GotoState16Action,
	{_State15, TypeToken}:                         _GotoState17Action,
	{_State15, FuncToken}:                         _GotoState18Action,
	{_State15, VarToken}:                          _GotoState38Action,
	{_State15, LetToken}:                          _GotoState28Action,
	{_State15, LbraceToken}:                       _GotoState62Action,
	{_State15, DefinitionsType}:                   _GotoState65Action,
	{_State15, DefinitionType}:                    _GotoState64Action,
	{_State15, VarDeclPatternType}:                _GotoState69Action,
	{_State15, VarOrLetType}:                      _GotoState59Action,
	{_State15, BlockBodyType}:                     _GotoState63Action,
	{_State15, TypeDefType}:                       _GotoState68Action,
	{_State15, NamedFuncDefType}:                  _GotoState66Action,
	{_State15, PackageDefType}:                    _GotoState67Action,
	{_State16, LparenToken}:                       _GotoState70Action,
	{_State17, IdentifierToken}:                   _GotoState71Action,
	{_State18, IdentifierToken}:                   _GotoState72Action,
	{_State18, LparenToken}:                       _GotoState73Action,
	{_State23, LparenToken}:                       _GotoState74Action,
	{_State27, IdentifierToken}:                   _GotoState81Action,
	{_State27, StructToken}:                       _GotoState35Action,
	{_State27, EnumToken}:                         _GotoState78Action,
	{_State27, TraitToken}:                        _GotoState86Action,
	{_State27, FuncToken}:                         _GotoState80Action,
	{_State27, LparenToken}:                       _GotoState83Action,
	{_State27, LbracketToken}:                     _GotoState27Action,
	{_State27, DotToken}:                          _GotoState77Action,
	{_State27, QuestionToken}:                     _GotoState84Action,
	{_State27, ExclaimToken}:                      _GotoState79Action,
	{_State27, TildeTildeToken}:                   _GotoState85Action,
	{_State27, BitNegToken}:                       _GotoState76Action,
	{_State27, BitAndToken}:                       _GotoState75Action,
	{_State27, LexErrorToken}:                     _GotoState82Action,
	{_State27, InitializableTypeType}:             _GotoState92Action,
	{_State27, AtomTypeType}:                      _GotoState87Action,
	{_State27, ReturnableTypeType}:                _GotoState93Action,
	{_State27, ValueTypeType}:                     _GotoState95Action,
	{_State27, ImplicitStructDefType}:             _GotoState91Action,
	{_State27, ExplicitStructDefType}:             _GotoState47Action,
	{_State27, ImplicitEnumDefType}:               _GotoState90Action,
	{_State27, ExplicitEnumDefType}:               _GotoState88Action,
	{_State27, TraitDefType}:                      _GotoState94Action,
	{_State27, FuncTypeType}:                      _GotoState89Action,
	{_State30, IntegerLiteralToken}:               _GotoState25Action,
	{_State30, FloatLiteralToken}:                 _GotoState22Action,
	{_State30, RuneLiteralToken}:                  _GotoState33Action,
	{_State30, StringLiteralToken}:                _GotoState34Action,
	{_State30, IdentifierToken}:                   _GotoState97Action,
	{_State30, TrueToken}:                         _GotoState37Action,
	{_State30, FalseToken}:                        _GotoState21Action,
	{_State30, StructToken}:                       _GotoState35Action,
	{_State30, FuncToken}:                         _GotoState23Action,
	{_State30, VarToken}:                          _GotoState38Action,
	{_State30, LetToken}:                          _GotoState28Action,
	{_State30, LabelDeclToken}:                    _GotoState26Action,
	{_State30, LparenToken}:                       _GotoState30Action,
	{_State30, LbracketToken}:                     _GotoState27Action,
	{_State30, DotdotdotToken}:                    _GotoState96Action,
	{_State30, NotToken}:                          _GotoState32Action,
	{_State30, SubToken}:                          _GotoState36Action,
	{_State30, MulToken}:                          _GotoState31Action,
	{_State30, BitNegToken}:                       _GotoState20Action,
	{_State30, BitAndToken}:                       _GotoState19Action,
	{_State30, LexErrorToken}:                     _GotoState29Action,
	{_State30, VarDeclPatternType}:                _GotoState58Action,
	{_State30, VarOrLetType}:                      _GotoState59Action,
	{_State30, ExpressionType}:                    _GotoState101Action,
	{_State30, OptionalLabelDeclType}:             _GotoState52Action,
	{_State30, SequenceExprType}:                  _GotoState57Action,
	{_State30, BlockExprType}:                     _GotoState44Action,
	{_State30, CallExprType}:                      _GotoState45Action,
	{_State30, ArgumentsType}:                     _GotoState99Action,
	{_State30, ArgumentType}:                      _GotoState98Action,
	{_State30, ColonExpressionsType}:              _GotoState100Action,
	{_State30, OptionalExpressionType}:            _GotoState102Action,
	{_State30, AtomExprType}:                      _GotoState43Action,
	{_State30, LiteralType}:                       _GotoState50Action,
	{_State30, ImplicitStructExprType}:            _GotoState48Action,
	{_State30, AccessExprType}:                    _GotoState39Action,
	{_State30, PostfixUnaryExprType}:              _GotoState54Action,
	{_State30, PrefixUnaryOpType}:                 _GotoState56Action,
	{_State30, PrefixUnaryExprType}:               _GotoState55Action,
	{_State30, MulExprType}:                       _GotoState51Action,
	{_State30, AddExprType}:                       _GotoState40Action,
	{_State30, CmpExprType}:                       _GotoState46Action,
	{_State30, AndExprType}:                       _GotoState41Action,
	{_State30, OrExprType}:                        _GotoState53Action,
	{_State30, InitializableTypeType}:             _GotoState49Action,
	{_State30, ExplicitStructDefType}:             _GotoState47Action,
	{_State30, AnonymousFuncExprType}:             _GotoState42Action,
	{_State35, LparenToken}:                       _GotoState103Action,
	{_State39, LbracketToken}:                     _GotoState106Action,
	{_State39, DotToken}:                          _GotoState105Action,
	{_State39, QuestionToken}:                     _GotoState107Action,
	{_State39, DollarLbracketToken}:               _GotoState104Action,
	{_State39, OptionalGenericBindingType}:        _GotoState108Action,
	{_State40, AddToken}:                          _GotoState109Action,
	{_State40, SubToken}:                          _GotoState112Action,
	{_State40, BitXorToken}:                       _GotoState111Action,
	{_State40, BitOrToken}:                        _GotoState110Action,
	{_State40, AddOpType}:                         _GotoState113Action,
	{_State41, AndToken}:                          _GotoState114Action,
	{_State46, EqualToken}:                        _GotoState115Action,
	{_State46, NotEqualToken}:                     _GotoState120Action,
	{_State46, LessToken}:                         _GotoState118Action,
	{_State46, LessOrEqualToken}:                  _GotoState119Action,
	{_State46, GreaterToken}:                      _GotoState116Action,
	{_State46, GreaterOrEqualToken}:               _GotoState117Action,
	{_State46, CmpOpType}:                         _GotoState121Action,
	{_State49, LparenToken}:                       _GotoState122Action,
	{_State51, MulToken}:                          _GotoState128Action,
	{_State51, DivToken}:                          _GotoState126Action,
	{_State51, ModToken}:                          _GotoState127Action,
	{_State51, BitAndToken}:                       _GotoState123Action,
	{_State51, BitLshiftToken}:                    _GotoState124Action,
	{_State51, BitRshiftToken}:                    _GotoState125Action,
	{_State51, MulOpType}:                         _GotoState129Action,
	{_State52, IfToken}:                           _GotoState132Action,
	{_State52, SwitchToken}:                       _GotoState133Action,
	{_State52, ForToken}:                          _GotoState131Action,
	{_State52, DoToken}:                           _GotoState130Action,
	{_State52, LbraceToken}:                       _GotoState62Action,
	{_State52, IfExprType}:                        _GotoState135Action,
	{_State52, SwitchExprType}:                    _GotoState137Action,
	{_State52, LoopExprType}:                      _GotoState136Action,
	{_State52, BlockBodyType}:                     _GotoState134Action,
	{_State53, OrToken}:                           _GotoState138Action,
	{_State56, IntegerLiteralToken}:               _GotoState25Action,
	{_State56, FloatLiteralToken}:                 _GotoState22Action,
	{_State56, RuneLiteralToken}:                  _GotoState33Action,
	{_State56, StringLiteralToken}:                _GotoState34Action,
	{_State56, IdentifierToken}:                   _GotoState24Action,
	{_State56, TrueToken}:                         _GotoState37Action,
	{_State56, FalseToken}:                        _GotoState21Action,
	{_State56, StructToken}:                       _GotoState35Action,
	{_State56, FuncToken}:                         _GotoState23Action,
	{_State56, LabelDeclToken}:                    _GotoState26Action,
	{_State56, LparenToken}:                       _GotoState30Action,
	{_State56, LbracketToken}:                     _GotoState27Action,
	{_State56, NotToken}:                          _GotoState32Action,
	{_State56, SubToken}:                          _GotoState36Action,
	{_State56, MulToken}:                          _GotoState31Action,
	{_State56, BitNegToken}:                       _GotoState20Action,
	{_State56, BitAndToken}:                       _GotoState19Action,
	{_State56, LexErrorToken}:                     _GotoState29Action,
	{_State56, OptionalLabelDeclType}:             _GotoState139Action,
	{_State56, BlockExprType}:                     _GotoState44Action,
	{_State56, CallExprType}:                      _GotoState45Action,
	{_State56, AtomExprType}:                      _GotoState43Action,
	{_State56, LiteralType}:                       _GotoState50Action,
	{_State56, ImplicitStructExprType}:            _GotoState48Action,
	{_State56, AccessExprType}:                    _GotoState39Action,
	{_State56, PostfixUnaryExprType}:              _GotoState54Action,
	{_State56, PrefixUnaryOpType}:                 _GotoState56Action,
	{_State56, PrefixUnaryExprType}:               _GotoState140Action,
	{_State56, InitializableTypeType}:             _GotoState49Action,
	{_State56, ExplicitStructDefType}:             _GotoState47Action,
	{_State56, AnonymousFuncExprType}:             _GotoState42Action,
	{_State59, IdentifierToken}:                   _GotoState141Action,
	{_State59, LparenToken}:                       _GotoState142Action,
	{_State59, VarPatternType}:                    _GotoState144Action,
	{_State59, TuplePatternType}:                  _GotoState143Action,
	{_State62, StatementsType}:                    _GotoState145Action,
	{_State65, NewlinesToken}:                     _GotoState146Action,
	{_State65, OptionalNewlinesType}:              _GotoState147Action,
	{_State69, AssignToken}:                       _GotoState148Action,
	{_State70, PackageStatementsType}:             _GotoState149Action,
	{_State71, DollarLbracketToken}:               _GotoState151Action,
	{_State71, AssignToken}:                       _GotoState150Action,
	{_State71, OptionalGenericParametersType}:     _GotoState152Action,
	{_State72, DollarLbracketToken}:               _GotoState151Action,
	{_State72, AssignToken}:                       _GotoState153Action,
	{_State72, OptionalGenericParametersType}:     _GotoState154Action,
	{_State73, IdentifierToken}:                   _GotoState155Action,
	{_State73, ParameterDefType}:                  _GotoState156Action,
	{_State74, IdentifierToken}:                   _GotoState155Action,
	{_State74, ParameterDefType}:                  _GotoState158Action,
	{_State74, ParameterDefsType}:                 _GotoState159Action,
	{_State74, OptionalParameterDefsType}:         _GotoState157Action,
	{_State75, IdentifierToken}:                   _GotoState81Action,
	{_State75, StructToken}:                       _GotoState35Action,
	{_State75, EnumToken}:                         _GotoState78Action,
	{_State75, TraitToken}:                        _GotoState86Action,
	{_State75, FuncToken}:                         _GotoState80Action,
	{_State75, LparenToken}:                       _GotoState83Action,
	{_State75, LbracketToken}:                     _GotoState27Action,
	{_State75, DotToken}:                          _GotoState77Action,
	{_State75, QuestionToken}:                     _GotoState84Action,
	{_State75, ExclaimToken}:                      _GotoState79Action,
	{_State75, TildeTildeToken}:                   _GotoState85Action,
	{_State75, BitNegToken}:                       _GotoState76Action,
	{_State75, BitAndToken}:                       _GotoState75Action,
	{_State75, LexErrorToken}:                     _GotoState82Action,
	{_State75, InitializableTypeType}:             _GotoState92Action,
	{_State75, AtomTypeType}:                      _GotoState87Action,
	{_State75, ReturnableTypeType}:                _GotoState160Action,
	{_State75, ImplicitStructDefType}:             _GotoState91Action,
	{_State75, ExplicitStructDefType}:             _GotoState47Action,
	{_State75, ImplicitEnumDefType}:               _GotoState90Action,
	{_State75, ExplicitEnumDefType}:               _GotoState88Action,
	{_State75, TraitDefType}:                      _GotoState94Action,
	{_State75, FuncTypeType}:                      _GotoState89Action,
	{_State76, IdentifierToken}:                   _GotoState81Action,
	{_State76, StructToken}:                       _GotoState35Action,
	{_State76, EnumToken}:                         _GotoState78Action,
	{_State76, TraitToken}:                        _GotoState86Action,
	{_State76, FuncToken}:                         _GotoState80Action,
	{_State76, LparenToken}:                       _GotoState83Action,
	{_State76, LbracketToken}:                     _GotoState27Action,
	{_State76, DotToken}:                          _GotoState77Action,
	{_State76, QuestionToken}:                     _GotoState84Action,
	{_State76, ExclaimToken}:                      _GotoState79Action,
	{_State76, TildeTildeToken}:                   _GotoState85Action,
	{_State76, BitNegToken}:                       _GotoState76Action,
	{_State76, BitAndToken}:                       _GotoState75Action,
	{_State76, LexErrorToken}:                     _GotoState82Action,
	{_State76, InitializableTypeType}:             _GotoState92Action,
	{_State76, AtomTypeType}:                      _GotoState87Action,
	{_State76, ReturnableTypeType}:                _GotoState161Action,
	{_State76, ImplicitStructDefType}:             _GotoState91Action,
	{_State76, ExplicitStructDefType}:             _GotoState47Action,
	{_State76, ImplicitEnumDefType}:               _GotoState90Action,
	{_State76, ExplicitEnumDefType}:               _GotoState88Action,
	{_State76, TraitDefType}:                      _GotoState94Action,
	{_State76, FuncTypeType}:                      _GotoState89Action,
	{_State77, DollarLbracketToken}:               _GotoState104Action,
	{_State77, OptionalGenericBindingType}:        _GotoState162Action,
	{_State78, LparenToken}:                       _GotoState163Action,
	{_State79, IdentifierToken}:                   _GotoState81Action,
	{_State79, StructToken}:                       _GotoState35Action,
	{_State79, EnumToken}:                         _GotoState78Action,
	{_State79, TraitToken}:                        _GotoState86Action,
	{_State79, FuncToken}:                         _GotoState80Action,
	{_State79, LparenToken}:                       _GotoState83Action,
	{_State79, LbracketToken}:                     _GotoState27Action,
	{_State79, DotToken}:                          _GotoState77Action,
	{_State79, QuestionToken}:                     _GotoState84Action,
	{_State79, ExclaimToken}:                      _GotoState79Action,
	{_State79, TildeTildeToken}:                   _GotoState85Action,
	{_State79, BitNegToken}:                       _GotoState76Action,
	{_State79, BitAndToken}:                       _GotoState75Action,
	{_State79, LexErrorToken}:                     _GotoState82Action,
	{_State79, InitializableTypeType}:             _GotoState92Action,
	{_State79, AtomTypeType}:                      _GotoState87Action,
	{_State79, ReturnableTypeType}:                _GotoState164Action,
	{_State79, ImplicitStructDefType}:             _GotoState91Action,
	{_State79, ExplicitStructDefType}:             _GotoState47Action,
	{_State79, ImplicitEnumDefType}:               _GotoState90Action,
	{_State79, ExplicitEnumDefType}:               _GotoState88Action,
	{_State79, TraitDefType}:                      _GotoState94Action,
	{_State79, FuncTypeType}:                      _GotoState89Action,
	{_State80, LparenToken}:                       _GotoState165Action,
	{_State81, DotToken}:                          _GotoState166Action,
	{_State81, DollarLbracketToken}:               _GotoState104Action,
	{_State81, OptionalGenericBindingType}:        _GotoState167Action,
	{_State83, IdentifierToken}:                   _GotoState168Action,
	{_State83, UnsafeToken}:                       _GotoState169Action,
	{_State83, StructToken}:                       _GotoState35Action,
	{_State83, EnumToken}:                         _GotoState78Action,
	{_State83, TraitToken}:                        _GotoState86Action,
	{_State83, FuncToken}:                         _GotoState80Action,
	{_State83, LparenToken}:                       _GotoState83Action,
	{_State83, LbracketToken}:                     _GotoState27Action,
	{_State83, DotToken}:                          _GotoState77Action,
	{_State83, QuestionToken}:                     _GotoState84Action,
	{_State83, ExclaimToken}:                      _GotoState79Action,
	{_State83, TildeTildeToken}:                   _GotoState85Action,
	{_State83, BitNegToken}:                       _GotoState76Action,
	{_State83, BitAndToken}:                       _GotoState75Action,
	{_State83, LexErrorToken}:                     _GotoState82Action,
	{_State83, UnsafeStatementType}:               _GotoState175Action,
	{_State83, InitializableTypeType}:             _GotoState92Action,
	{_State83, AtomTypeType}:                      _GotoState87Action,
	{_State83, ReturnableTypeType}:                _GotoState93Action,
	{_State83, ValueTypeType}:                     _GotoState176Action,
	{_State83, FieldDefType}:                      _GotoState171Action,
	{_State83, ImplicitFieldDefsType}:             _GotoState173Action,
	{_State83, OptionalImplicitFieldDefsType}:     _GotoState174Action,
	{_State83, ImplicitStructDefType}:             _GotoState91Action,
	{_State83, ExplicitStructDefType}:             _GotoState47Action,
	{_State83, EnumValueDefType}:                  _GotoState170Action,
	{_State83, ImplicitEnumValueDefsType}:         _GotoState172Action,
	{_State83, ImplicitEnumDefType}:               _GotoState90Action,
	{_State83, ExplicitEnumDefType}:               _GotoState88Action,
	{_State83, TraitDefType}:                      _GotoState94Action,
	{_State83, FuncTypeType}:                      _GotoState89Action,
	{_State84, IdentifierToken}:                   _GotoState81Action,
	{_State84, StructToken}:                       _GotoState35Action,
	{_State84, EnumToken}:                         _GotoState78Action,
	{_State84, TraitToken}:                        _GotoState86Action,
	{_State84, FuncToken}:                         _GotoState80Action,
	{_State84, LparenToken}:                       _GotoState83Action,
	{_State84, LbracketToken}:                     _GotoState27Action,
	{_State84, DotToken}:                          _GotoState77Action,
	{_State84, QuestionToken}:                     _GotoState84Action,
	{_State84, ExclaimToken}:                      _GotoState79Action,
	{_State84, TildeTildeToken}:                   _GotoState85Action,
	{_State84, BitNegToken}:                       _GotoState76Action,
	{_State84, BitAndToken}:                       _GotoState75Action,
	{_State84, LexErrorToken}:                     _GotoState82Action,
	{_State84, InitializableTypeType}:             _GotoState92Action,
	{_State84, AtomTypeType}:                      _GotoState87Action,
	{_State84, ReturnableTypeType}:                _GotoState177Action,
	{_State84, ImplicitStructDefType}:             _GotoState91Action,
	{_State84, ExplicitStructDefType}:             _GotoState47Action,
	{_State84, ImplicitEnumDefType}:               _GotoState90Action,
	{_State84, ExplicitEnumDefType}:               _GotoState88Action,
	{_State84, TraitDefType}:                      _GotoState94Action,
	{_State84, FuncTypeType}:                      _GotoState89Action,
	{_State85, IdentifierToken}:                   _GotoState81Action,
	{_State85, StructToken}:                       _GotoState35Action,
	{_State85, EnumToken}:                         _GotoState78Action,
	{_State85, TraitToken}:                        _GotoState86Action,
	{_State85, FuncToken}:                         _GotoState80Action,
	{_State85, LparenToken}:                       _GotoState83Action,
	{_State85, LbracketToken}:                     _GotoState27Action,
	{_State85, DotToken}:                          _GotoState77Action,
	{_State85, QuestionToken}:                     _GotoState84Action,
	{_State85, ExclaimToken}:                      _GotoState79Action,
	{_State85, TildeTildeToken}:                   _GotoState85Action,
	{_State85, BitNegToken}:                       _GotoState76Action,
	{_State85, BitAndToken}:                       _GotoState75Action,
	{_State85, LexErrorToken}:                     _GotoState82Action,
	{_State85, InitializableTypeType}:             _GotoState92Action,
	{_State85, AtomTypeType}:                      _GotoState87Action,
	{_State85, ReturnableTypeType}:                _GotoState178Action,
	{_State85, ImplicitStructDefType}:             _GotoState91Action,
	{_State85, ExplicitStructDefType}:             _GotoState47Action,
	{_State85, ImplicitEnumDefType}:               _GotoState90Action,
	{_State85, ExplicitEnumDefType}:               _GotoState88Action,
	{_State85, TraitDefType}:                      _GotoState94Action,
	{_State85, FuncTypeType}:                      _GotoState89Action,
	{_State86, LparenToken}:                       _GotoState179Action,
	{_State95, RbracketToken}:                     _GotoState184Action,
	{_State95, CommaToken}:                        _GotoState182Action,
	{_State95, ColonToken}:                        _GotoState181Action,
	{_State95, AddToken}:                          _GotoState180Action,
	{_State95, SubToken}:                          _GotoState185Action,
	{_State95, MulToken}:                          _GotoState183Action,
	{_State97, AssignToken}:                       _GotoState186Action,
	{_State99, RparenToken}:                       _GotoState188Action,
	{_State99, CommaToken}:                        _GotoState187Action,
	{_State100, ColonToken}:                       _GotoState189Action,
	{_State102, ColonToken}:                       _GotoState190Action,
	{_State103, IdentifierToken}:                  _GotoState168Action,
	{_State103, UnsafeToken}:                      _GotoState169Action,
	{_State103, StructToken}:                      _GotoState35Action,
	{_State103, EnumToken}:                        _GotoState78Action,
	{_State103, TraitToken}:                       _GotoState86Action,
	{_State103, FuncToken}:                        _GotoState80Action,
	{_State103, LparenToken}:                      _GotoState83Action,
	{_State103, LbracketToken}:                    _GotoState27Action,
	{_State103, DotToken}:                         _GotoState77Action,
	{_State103, QuestionToken}:                    _GotoState84Action,
	{_State103, ExclaimToken}:                     _GotoState79Action,
	{_State103, TildeTildeToken}:                  _GotoState85Action,
	{_State103, BitNegToken}:                      _GotoState76Action,
	{_State103, BitAndToken}:                      _GotoState75Action,
	{_State103, LexErrorToken}:                    _GotoState82Action,
	{_State103, UnsafeStatementType}:              _GotoState175Action,
	{_State103, InitializableTypeType}:            _GotoState92Action,
	{_State103, AtomTypeType}:                     _GotoState87Action,
	{_State103, ReturnableTypeType}:               _GotoState93Action,
	{_State103, ValueTypeType}:                    _GotoState176Action,
	{_State103, FieldDefType}:                     _GotoState192Action,
	{_State103, ImplicitStructDefType}:            _GotoState91Action,
	{_State103, ExplicitFieldDefsType}:            _GotoState191Action,
	{_State103, OptionalExplicitFieldDefsType}:    _GotoState193Action,
	{_State103, ExplicitStructDefType}:            _GotoState47Action,
	{_State103, ImplicitEnumDefType}:              _GotoState90Action,
	{_State103, ExplicitEnumDefType}:              _GotoState88Action,
	{_State103, TraitDefType}:                     _GotoState94Action,
	{_State103, FuncTypeType}:                     _GotoState89Action,
	{_State104, IdentifierToken}:                  _GotoState81Action,
	{_State104, StructToken}:                      _GotoState35Action,
	{_State104, EnumToken}:                        _GotoState78Action,
	{_State104, TraitToken}:                       _GotoState86Action,
	{_State104, FuncToken}:                        _GotoState80Action,
	{_State104, LparenToken}:                      _GotoState83Action,
	{_State104, LbracketToken}:                    _GotoState27Action,
	{_State104, DotToken}:                         _GotoState77Action,
	{_State104, QuestionToken}:                    _GotoState84Action,
	{_State104, ExclaimToken}:                     _GotoState79Action,
	{_State104, TildeTildeToken}:                  _GotoState85Action,
	{_State104, BitNegToken}:                      _GotoState76Action,
	{_State104, BitAndToken}:                      _GotoState75Action,
	{_State104, LexErrorToken}:                    _GotoState82Action,
	{_State104, OptionalGenericArgumentsType}:     _GotoState195Action,
	{_State104, GenericArgumentsType}:             _GotoState194Action,
	{_State104, InitializableTypeType}:            _GotoState92Action,
	{_State104, AtomTypeType}:                     _GotoState87Action,
	{_State104, ReturnableTypeType}:               _GotoState93Action,
	{_State104, ValueTypeType}:                    _GotoState196Action,
	{_State104, ImplicitStructDefType}:            _GotoState91Action,
	{_State104, ExplicitStructDefType}:            _GotoState47Action,
	{_State104, ImplicitEnumDefType}:              _GotoState90Action,
	{_State104, ExplicitEnumDefType}:              _GotoState88Action,
	{_State104, TraitDefType}:                     _GotoState94Action,
	{_State104, FuncTypeType}:                     _GotoState89Action,
	{_State105, IdentifierToken}:                  _GotoState197Action,
	{_State106, IntegerLiteralToken}:              _GotoState25Action,
	{_State106, FloatLiteralToken}:                _GotoState22Action,
	{_State106, RuneLiteralToken}:                 _GotoState33Action,
	{_State106, StringLiteralToken}:               _GotoState34Action,
	{_State106, IdentifierToken}:                  _GotoState97Action,
	{_State106, TrueToken}:                        _GotoState37Action,
	{_State106, FalseToken}:                       _GotoState21Action,
	{_State106, StructToken}:                      _GotoState35Action,
	{_State106, FuncToken}:                        _GotoState23Action,
	{_State106, VarToken}:                         _GotoState38Action,
	{_State106, LetToken}:                         _GotoState28Action,
	{_State106, LabelDeclToken}:                   _GotoState26Action,
	{_State106, LparenToken}:                      _GotoState30Action,
	{_State106, LbracketToken}:                    _GotoState27Action,
	{_State106, DotdotdotToken}:                   _GotoState96Action,
	{_State106, NotToken}:                         _GotoState32Action,
	{_State106, SubToken}:                         _GotoState36Action,
	{_State106, MulToken}:                         _GotoState31Action,
	{_State106, BitNegToken}:                      _GotoState20Action,
	{_State106, BitAndToken}:                      _GotoState19Action,
	{_State106, LexErrorToken}:                    _GotoState29Action,
	{_State106, VarDeclPatternType}:               _GotoState58Action,
	{_State106, VarOrLetType}:                     _GotoState59Action,
	{_State106, ExpressionType}:                   _GotoState101Action,
	{_State106, OptionalLabelDeclType}:            _GotoState52Action,
	{_State106, SequenceExprType}:                 _GotoState57Action,
	{_State106, BlockExprType}:                    _GotoState44Action,
	{_State106, CallExprType}:                     _GotoState45Action,
	{_State106, ArgumentType}:                     _GotoState198Action,
	{_State106, ColonExpressionsType}:             _GotoState100Action,
	{_State106, OptionalExpressionType}:           _GotoState102Action,
	{_State106, AtomExprType}:                     _GotoState43Action,
	{_State106, LiteralType}:                      _GotoState50Action,
	{_State106, ImplicitStructExprType}:           _GotoState48Action,
	{_State106, AccessExprType}:                   _GotoState39Action,
	{_State106, PostfixUnaryExprType}:             _GotoState54Action,
	{_State106, PrefixUnaryOpType}:                _GotoState56Action,
	{_State106, PrefixUnaryExprType}:              _GotoState55Action,
	{_State106, MulExprType}:                      _GotoState51Action,
	{_State106, AddExprType}:                      _GotoState40Action,
	{_State106, CmpExprType}:                      _GotoState46Action,
	{_State106, AndExprType}:                      _GotoState41Action,
	{_State106, OrExprType}:                       _GotoState53Action,
	{_State106, InitializableTypeType}:            _GotoState49Action,
	{_State106, ExplicitStructDefType}:            _GotoState47Action,
	{_State106, AnonymousFuncExprType}:            _GotoState42Action,
	{_State108, LparenToken}:                      _GotoState199Action,
	{_State113, IntegerLiteralToken}:              _GotoState25Action,
	{_State113, FloatLiteralToken}:                _GotoState22Action,
	{_State113, RuneLiteralToken}:                 _GotoState33Action,
	{_State113, StringLiteralToken}:               _GotoState34Action,
	{_State113, IdentifierToken}:                  _GotoState24Action,
	{_State113, TrueToken}:                        _GotoState37Action,
	{_State113, FalseToken}:                       _GotoState21Action,
	{_State113, StructToken}:                      _GotoState35Action,
	{_State113, FuncToken}:                        _GotoState23Action,
	{_State113, LabelDeclToken}:                   _GotoState26Action,
	{_State113, LparenToken}:                      _GotoState30Action,
	{_State113, LbracketToken}:                    _GotoState27Action,
	{_State113, NotToken}:                         _GotoState32Action,
	{_State113, SubToken}:                         _GotoState36Action,
	{_State113, MulToken}:                         _GotoState31Action,
	{_State113, BitNegToken}:                      _GotoState20Action,
	{_State113, BitAndToken}:                      _GotoState19Action,
	{_State113, LexErrorToken}:                    _GotoState29Action,
	{_State113, OptionalLabelDeclType}:            _GotoState139Action,
	{_State113, BlockExprType}:                    _GotoState44Action,
	{_State113, CallExprType}:                     _GotoState45Action,
	{_State113, AtomExprType}:                     _GotoState43Action,
	{_State113, LiteralType}:                      _GotoState50Action,
	{_State113, ImplicitStructExprType}:           _GotoState48Action,
	{_State113, AccessExprType}:                   _GotoState39Action,
	{_State113, PostfixUnaryExprType}:             _GotoState54Action,
	{_State113, PrefixUnaryOpType}:                _GotoState56Action,
	{_State113, PrefixUnaryExprType}:              _GotoState55Action,
	{_State113, MulExprType}:                      _GotoState200Action,
	{_State113, InitializableTypeType}:            _GotoState49Action,
	{_State113, ExplicitStructDefType}:            _GotoState47Action,
	{_State113, AnonymousFuncExprType}:            _GotoState42Action,
	{_State114, IntegerLiteralToken}:              _GotoState25Action,
	{_State114, FloatLiteralToken}:                _GotoState22Action,
	{_State114, RuneLiteralToken}:                 _GotoState33Action,
	{_State114, StringLiteralToken}:               _GotoState34Action,
	{_State114, IdentifierToken}:                  _GotoState24Action,
	{_State114, TrueToken}:                        _GotoState37Action,
	{_State114, FalseToken}:                       _GotoState21Action,
	{_State114, StructToken}:                      _GotoState35Action,
	{_State114, FuncToken}:                        _GotoState23Action,
	{_State114, LabelDeclToken}:                   _GotoState26Action,
	{_State114, LparenToken}:                      _GotoState30Action,
	{_State114, LbracketToken}:                    _GotoState27Action,
	{_State114, NotToken}:                         _GotoState32Action,
	{_State114, SubToken}:                         _GotoState36Action,
	{_State114, MulToken}:                         _GotoState31Action,
	{_State114, BitNegToken}:                      _GotoState20Action,
	{_State114, BitAndToken}:                      _GotoState19Action,
	{_State114, LexErrorToken}:                    _GotoState29Action,
	{_State114, OptionalLabelDeclType}:            _GotoState139Action,
	{_State114, BlockExprType}:                    _GotoState44Action,
	{_State114, CallExprType}:                     _GotoState45Action,
	{_State114, AtomExprType}:                     _GotoState43Action,
	{_State114, LiteralType}:                      _GotoState50Action,
	{_State114, ImplicitStructExprType}:           _GotoState48Action,
	{_State114, AccessExprType}:                   _GotoState39Action,
	{_State114, PostfixUnaryExprType}:             _GotoState54Action,
	{_State114, PrefixUnaryOpType}:                _GotoState56Action,
	{_State114, PrefixUnaryExprType}:              _GotoState55Action,
	{_State114, MulExprType}:                      _GotoState51Action,
	{_State114, AddExprType}:                      _GotoState40Action,
	{_State114, CmpExprType}:                      _GotoState201Action,
	{_State114, InitializableTypeType}:            _GotoState49Action,
	{_State114, ExplicitStructDefType}:            _GotoState47Action,
	{_State114, AnonymousFuncExprType}:            _GotoState42Action,
	{_State121, IntegerLiteralToken}:              _GotoState25Action,
	{_State121, FloatLiteralToken}:                _GotoState22Action,
	{_State121, RuneLiteralToken}:                 _GotoState33Action,
	{_State121, StringLiteralToken}:               _GotoState34Action,
	{_State121, IdentifierToken}:                  _GotoState24Action,
	{_State121, TrueToken}:                        _GotoState37Action,
	{_State121, FalseToken}:                       _GotoState21Action,
	{_State121, StructToken}:                      _GotoState35Action,
	{_State121, FuncToken}:                        _GotoState23Action,
	{_State121, LabelDeclToken}:                   _GotoState26Action,
	{_State121, LparenToken}:                      _GotoState30Action,
	{_State121, LbracketToken}:                    _GotoState27Action,
	{_State121, NotToken}:                         _GotoState32Action,
	{_State121, SubToken}:                         _GotoState36Action,
	{_State121, MulToken}:                         _GotoState31Action,
	{_State121, BitNegToken}:                      _GotoState20Action,
	{_State121, BitAndToken}:                      _GotoState19Action,
	{_State121, LexErrorToken}:                    _GotoState29Action,
	{_State121, OptionalLabelDeclType}:            _GotoState139Action,
	{_State121, BlockExprType}:                    _GotoState44Action,
	{_State121, CallExprType}:                     _GotoState45Action,
	{_State121, AtomExprType}:                     _GotoState43Action,
	{_State121, LiteralType}:                      _GotoState50Action,
	{_State121, ImplicitStructExprType}:           _GotoState48Action,
	{_State121, AccessExprType}:                   _GotoState39Action,
	{_State121, PostfixUnaryExprType}:             _GotoState54Action,
	{_State121, PrefixUnaryOpType}:                _GotoState56Action,
	{_State121, PrefixUnaryExprType}:              _GotoState55Action,
	{_State121, MulExprType}:                      _GotoState51Action,
	{_State121, AddExprType}:                      _GotoState202Action,
	{_State121, InitializableTypeType}:            _GotoState49Action,
	{_State121, ExplicitStructDefType}:            _GotoState47Action,
	{_State121, AnonymousFuncExprType}:            _GotoState42Action,
	{_State122, IntegerLiteralToken}:              _GotoState25Action,
	{_State122, FloatLiteralToken}:                _GotoState22Action,
	{_State122, RuneLiteralToken}:                 _GotoState33Action,
	{_State122, StringLiteralToken}:               _GotoState34Action,
	{_State122, IdentifierToken}:                  _GotoState97Action,
	{_State122, TrueToken}:                        _GotoState37Action,
	{_State122, FalseToken}:                       _GotoState21Action,
	{_State122, StructToken}:                      _GotoState35Action,
	{_State122, FuncToken}:                        _GotoState23Action,
	{_State122, VarToken}:                         _GotoState38Action,
	{_State122, LetToken}:                         _GotoState28Action,
	{_State122, LabelDeclToken}:                   _GotoState26Action,
	{_State122, LparenToken}:                      _GotoState30Action,
	{_State122, LbracketToken}:                    _GotoState27Action,
	{_State122, DotdotdotToken}:                   _GotoState96Action,
	{_State122, NotToken}:                         _GotoState32Action,
	{_State122, SubToken}:                         _GotoState36Action,
	{_State122, MulToken}:                         _GotoState31Action,
	{_State122, BitNegToken}:                      _GotoState20Action,
	{_State122, BitAndToken}:                      _GotoState19Action,
	{_State122, LexErrorToken}:                    _GotoState29Action,
	{_State122, VarDeclPatternType}:               _GotoState58Action,
	{_State122, VarOrLetType}:                     _GotoState59Action,
	{_State122, ExpressionType}:                   _GotoState101Action,
	{_State122, OptionalLabelDeclType}:            _GotoState52Action,
	{_State122, SequenceExprType}:                 _GotoState57Action,
	{_State122, BlockExprType}:                    _GotoState44Action,
	{_State122, CallExprType}:                     _GotoState45Action,
	{_State122, ArgumentsType}:                    _GotoState203Action,
	{_State122, ArgumentType}:                     _GotoState98Action,
	{_State122, ColonExpressionsType}:             _GotoState100Action,
	{_State122, OptionalExpressionType}:           _GotoState102Action,
	{_State122, AtomExprType}:                     _GotoState43Action,
	{_State122, LiteralType}:                      _GotoState50Action,
	{_State122, ImplicitStructExprType}:           _GotoState48Action,
	{_State122, AccessExprType}:                   _GotoState39Action,
	{_State122, PostfixUnaryExprType}:             _GotoState54Action,
	{_State122, PrefixUnaryOpType}:                _GotoState56Action,
	{_State122, PrefixUnaryExprType}:              _GotoState55Action,
	{_State122, MulExprType}:                      _GotoState51Action,
	{_State122, AddExprType}:                      _GotoState40Action,
	{_State122, CmpExprType}:                      _GotoState46Action,
	{_State122, AndExprType}:                      _GotoState41Action,
	{_State122, OrExprType}:                       _GotoState53Action,
	{_State122, InitializableTypeType}:            _GotoState49Action,
	{_State122, ExplicitStructDefType}:            _GotoState47Action,
	{_State122, AnonymousFuncExprType}:            _GotoState42Action,
	{_State129, IntegerLiteralToken}:              _GotoState25Action,
	{_State129, FloatLiteralToken}:                _GotoState22Action,
	{_State129, RuneLiteralToken}:                 _GotoState33Action,
	{_State129, StringLiteralToken}:               _GotoState34Action,
	{_State129, IdentifierToken}:                  _GotoState24Action,
	{_State129, TrueToken}:                        _GotoState37Action,
	{_State129, FalseToken}:                       _GotoState21Action,
	{_State129, StructToken}:                      _GotoState35Action,
	{_State129, FuncToken}:                        _GotoState23Action,
	{_State129, LabelDeclToken}:                   _GotoState26Action,
	{_State129, LparenToken}:                      _GotoState30Action,
	{_State129, LbracketToken}:                    _GotoState27Action,
	{_State129, NotToken}:                         _GotoState32Action,
	{_State129, SubToken}:                         _GotoState36Action,
	{_State129, MulToken}:                         _GotoState31Action,
	{_State129, BitNegToken}:                      _GotoState20Action,
	{_State129, BitAndToken}:                      _GotoState19Action,
	{_State129, LexErrorToken}:                    _GotoState29Action,
	{_State129, OptionalLabelDeclType}:            _GotoState139Action,
	{_State129, BlockExprType}:                    _GotoState44Action,
	{_State129, CallExprType}:                     _GotoState45Action,
	{_State129, AtomExprType}:                     _GotoState43Action,
	{_State129, LiteralType}:                      _GotoState50Action,
	{_State129, ImplicitStructExprType}:           _GotoState48Action,
	{_State129, AccessExprType}:                   _GotoState39Action,
	{_State129, PostfixUnaryExprType}:             _GotoState54Action,
	{_State129, PrefixUnaryOpType}:                _GotoState56Action,
	{_State129, PrefixUnaryExprType}:              _GotoState204Action,
	{_State129, InitializableTypeType}:            _GotoState49Action,
	{_State129, ExplicitStructDefType}:            _GotoState47Action,
	{_State129, AnonymousFuncExprType}:            _GotoState42Action,
	{_State130, LbraceToken}:                      _GotoState62Action,
	{_State130, BlockBodyType}:                    _GotoState205Action,
	{_State131, IntegerLiteralToken}:              _GotoState25Action,
	{_State131, FloatLiteralToken}:                _GotoState22Action,
	{_State131, RuneLiteralToken}:                 _GotoState33Action,
	{_State131, StringLiteralToken}:               _GotoState34Action,
	{_State131, IdentifierToken}:                  _GotoState24Action,
	{_State131, TrueToken}:                        _GotoState37Action,
	{_State131, FalseToken}:                       _GotoState21Action,
	{_State131, StructToken}:                      _GotoState35Action,
	{_State131, FuncToken}:                        _GotoState23Action,
	{_State131, VarToken}:                         _GotoState38Action,
	{_State131, LetToken}:                         _GotoState28Action,
	{_State131, LabelDeclToken}:                   _GotoState26Action,
	{_State131, LparenToken}:                      _GotoState30Action,
	{_State131, LbracketToken}:                    _GotoState27Action,
	{_State131, NotToken}:                         _GotoState32Action,
	{_State131, SubToken}:                         _GotoState36Action,
	{_State131, MulToken}:                         _GotoState31Action,
	{_State131, BitNegToken}:                      _GotoState20Action,
	{_State131, BitAndToken}:                      _GotoState19Action,
	{_State131, LexErrorToken}:                    _GotoState29Action,
	{_State131, VarDeclPatternType}:               _GotoState58Action,
	{_State131, VarOrLetType}:                     _GotoState59Action,
	{_State131, AssignPatternType}:                _GotoState206Action,
	{_State131, OptionalLabelDeclType}:            _GotoState139Action,
	{_State131, ForAssignmentType}:                _GotoState207Action,
	{_State131, SequenceExprType}:                 _GotoState208Action,
	{_State131, BlockExprType}:                    _GotoState44Action,
	{_State131, CallExprType}:                     _GotoState45Action,
	{_State131, AtomExprType}:                     _GotoState43Action,
	{_State131, LiteralType}:                      _GotoState50Action,
	{_State131, ImplicitStructExprType}:           _GotoState48Action,
	{_State131, AccessExprType}:                   _GotoState39Action,
	{_State131, PostfixUnaryExprType}:             _GotoState54Action,
	{_State131, PrefixUnaryOpType}:                _GotoState56Action,
	{_State131, PrefixUnaryExprType}:              _GotoState55Action,
	{_State131, MulExprType}:                      _GotoState51Action,
	{_State131, AddExprType}:                      _GotoState40Action,
	{_State131, CmpExprType}:                      _GotoState46Action,
	{_State131, AndExprType}:                      _GotoState41Action,
	{_State131, OrExprType}:                       _GotoState53Action,
	{_State131, InitializableTypeType}:            _GotoState49Action,
	{_State131, ExplicitStructDefType}:            _GotoState47Action,
	{_State131, AnonymousFuncExprType}:            _GotoState42Action,
	{_State132, IntegerLiteralToken}:              _GotoState25Action,
	{_State132, FloatLiteralToken}:                _GotoState22Action,
	{_State132, RuneLiteralToken}:                 _GotoState33Action,
	{_State132, StringLiteralToken}:               _GotoState34Action,
	{_State132, IdentifierToken}:                  _GotoState24Action,
	{_State132, TrueToken}:                        _GotoState37Action,
	{_State132, FalseToken}:                       _GotoState21Action,
	{_State132, CaseToken}:                        _GotoState209Action,
	{_State132, StructToken}:                      _GotoState35Action,
	{_State132, FuncToken}:                        _GotoState23Action,
	{_State132, VarToken}:                         _GotoState38Action,
	{_State132, LetToken}:                         _GotoState28Action,
	{_State132, LabelDeclToken}:                   _GotoState26Action,
	{_State132, LparenToken}:                      _GotoState30Action,
	{_State132, LbracketToken}:                    _GotoState27Action,
	{_State132, NotToken}:                         _GotoState32Action,
	{_State132, SubToken}:                         _GotoState36Action,
	{_State132, MulToken}:                         _GotoState31Action,
	{_State132, BitNegToken}:                      _GotoState20Action,
	{_State132, BitAndToken}:                      _GotoState19Action,
	{_State132, LexErrorToken}:                    _GotoState29Action,
	{_State132, VarDeclPatternType}:               _GotoState58Action,
	{_State132, VarOrLetType}:                     _GotoState59Action,
	{_State132, OptionalLabelDeclType}:            _GotoState139Action,
	{_State132, ConditionType}:                    _GotoState210Action,
	{_State132, SequenceExprType}:                 _GotoState211Action,
	{_State132, BlockExprType}:                    _GotoState44Action,
	{_State132, CallExprType}:                     _GotoState45Action,
	{_State132, AtomExprType}:                     _GotoState43Action,
	{_State132, LiteralType}:                      _GotoState50Action,
	{_State132, ImplicitStructExprType}:           _GotoState48Action,
	{_State132, AccessExprType}:                   _GotoState39Action,
	{_State132, PostfixUnaryExprType}:             _GotoState54Action,
	{_State132, PrefixUnaryOpType}:                _GotoState56Action,
	{_State132, PrefixUnaryExprType}:              _GotoState55Action,
	{_State132, MulExprType}:                      _GotoState51Action,
	{_State132, AddExprType}:                      _GotoState40Action,
	{_State132, CmpExprType}:                      _GotoState46Action,
	{_State132, AndExprType}:                      _GotoState41Action,
	{_State132, OrExprType}:                       _GotoState53Action,
	{_State132, InitializableTypeType}:            _GotoState49Action,
	{_State132, ExplicitStructDefType}:            _GotoState47Action,
	{_State132, AnonymousFuncExprType}:            _GotoState42Action,
	{_State133, IntegerLiteralToken}:              _GotoState25Action,
	{_State133, FloatLiteralToken}:                _GotoState22Action,
	{_State133, RuneLiteralToken}:                 _GotoState33Action,
	{_State133, StringLiteralToken}:               _GotoState34Action,
	{_State133, IdentifierToken}:                  _GotoState24Action,
	{_State133, TrueToken}:                        _GotoState37Action,
	{_State133, FalseToken}:                       _GotoState21Action,
	{_State133, StructToken}:                      _GotoState35Action,
	{_State133, FuncToken}:                        _GotoState23Action,
	{_State133, VarToken}:                         _GotoState38Action,
	{_State133, LetToken}:                         _GotoState28Action,
	{_State133, LabelDeclToken}:                   _GotoState26Action,
	{_State133, LparenToken}:                      _GotoState30Action,
	{_State133, LbracketToken}:                    _GotoState27Action,
	{_State133, NotToken}:                         _GotoState32Action,
	{_State133, SubToken}:                         _GotoState36Action,
	{_State133, MulToken}:                         _GotoState31Action,
	{_State133, BitNegToken}:                      _GotoState20Action,
	{_State133, BitAndToken}:                      _GotoState19Action,
	{_State133, LexErrorToken}:                    _GotoState29Action,
	{_State133, VarDeclPatternType}:               _GotoState58Action,
	{_State133, VarOrLetType}:                     _GotoState59Action,
	{_State133, OptionalLabelDeclType}:            _GotoState139Action,
	{_State133, SequenceExprType}:                 _GotoState212Action,
	{_State133, BlockExprType}:                    _GotoState44Action,
	{_State133, CallExprType}:                     _GotoState45Action,
	{_State133, AtomExprType}:                     _GotoState43Action,
	{_State133, LiteralType}:                      _GotoState50Action,
	{_State133, ImplicitStructExprType}:           _GotoState48Action,
	{_State133, AccessExprType}:                   _GotoState39Action,
	{_State133, PostfixUnaryExprType}:             _GotoState54Action,
	{_State133, PrefixUnaryOpType}:                _GotoState56Action,
	{_State133, PrefixUnaryExprType}:              _GotoState55Action,
	{_State133, MulExprType}:                      _GotoState51Action,
	{_State133, AddExprType}:                      _GotoState40Action,
	{_State133, CmpExprType}:                      _GotoState46Action,
	{_State133, AndExprType}:                      _GotoState41Action,
	{_State133, OrExprType}:                       _GotoState53Action,
	{_State133, InitializableTypeType}:            _GotoState49Action,
	{_State133, ExplicitStructDefType}:            _GotoState47Action,
	{_State133, AnonymousFuncExprType}:            _GotoState42Action,
	{_State138, IntegerLiteralToken}:              _GotoState25Action,
	{_State138, FloatLiteralToken}:                _GotoState22Action,
	{_State138, RuneLiteralToken}:                 _GotoState33Action,
	{_State138, StringLiteralToken}:               _GotoState34Action,
	{_State138, IdentifierToken}:                  _GotoState24Action,
	{_State138, TrueToken}:                        _GotoState37Action,
	{_State138, FalseToken}:                       _GotoState21Action,
	{_State138, StructToken}:                      _GotoState35Action,
	{_State138, FuncToken}:                        _GotoState23Action,
	{_State138, LabelDeclToken}:                   _GotoState26Action,
	{_State138, LparenToken}:                      _GotoState30Action,
	{_State138, LbracketToken}:                    _GotoState27Action,
	{_State138, NotToken}:                         _GotoState32Action,
	{_State138, SubToken}:                         _GotoState36Action,
	{_State138, MulToken}:                         _GotoState31Action,
	{_State138, BitNegToken}:                      _GotoState20Action,
	{_State138, BitAndToken}:                      _GotoState19Action,
	{_State138, LexErrorToken}:                    _GotoState29Action,
	{_State138, OptionalLabelDeclType}:            _GotoState139Action,
	{_State138, BlockExprType}:                    _GotoState44Action,
	{_State138, CallExprType}:                     _GotoState45Action,
	{_State138, AtomExprType}:                     _GotoState43Action,
	{_State138, LiteralType}:                      _GotoState50Action,
	{_State138, ImplicitStructExprType}:           _GotoState48Action,
	{_State138, AccessExprType}:                   _GotoState39Action,
	{_State138, PostfixUnaryExprType}:             _GotoState54Action,
	{_State138, PrefixUnaryOpType}:                _GotoState56Action,
	{_State138, PrefixUnaryExprType}:              _GotoState55Action,
	{_State138, MulExprType}:                      _GotoState51Action,
	{_State138, AddExprType}:                      _GotoState40Action,
	{_State138, CmpExprType}:                      _GotoState46Action,
	{_State138, AndExprType}:                      _GotoState213Action,
	{_State138, InitializableTypeType}:            _GotoState49Action,
	{_State138, ExplicitStructDefType}:            _GotoState47Action,
	{_State138, AnonymousFuncExprType}:            _GotoState42Action,
	{_State139, LbraceToken}:                      _GotoState62Action,
	{_State139, BlockBodyType}:                    _GotoState134Action,
	{_State142, IdentifierToken}:                  _GotoState215Action,
	{_State142, LparenToken}:                      _GotoState142Action,
	{_State142, DotdotdotToken}:                   _GotoState214Action,
	{_State142, VarPatternType}:                   _GotoState218Action,
	{_State142, TuplePatternType}:                 _GotoState143Action,
	{_State142, FieldVarPatternsType}:             _GotoState217Action,
	{_State142, FieldVarPatternType}:              _GotoState216Action,
	{_State144, IdentifierToken}:                  _GotoState81Action,
	{_State144, StructToken}:                      _GotoState35Action,
	{_State144, EnumToken}:                        _GotoState78Action,
	{_State144, TraitToken}:                       _GotoState86Action,
	{_State144, FuncToken}:                        _GotoState80Action,
	{_State144, LparenToken}:                      _GotoState83Action,
	{_State144, LbracketToken}:                    _GotoState27Action,
	{_State144, DotToken}:                         _GotoState77Action,
	{_State144, QuestionToken}:                    _GotoState84Action,
	{_State144, ExclaimToken}:                     _GotoState79Action,
	{_State144, TildeTildeToken}:                  _GotoState85Action,
	{_State144, BitNegToken}:                      _GotoState76Action,
	{_State144, BitAndToken}:                      _GotoState75Action,
	{_State144, LexErrorToken}:                    _GotoState82Action,
	{_State144, OptionalValueTypeType}:            _GotoState219Action,
	{_State144, InitializableTypeType}:            _GotoState92Action,
	{_State144, AtomTypeType}:                     _GotoState87Action,
	{_State144, ReturnableTypeType}:               _GotoState93Action,
	{_State144, ValueTypeType}:                    _GotoState220Action,
	{_State144, ImplicitStructDefType}:            _GotoState91Action,
	{_State144, ExplicitStructDefType}:            _GotoState47Action,
	{_State144, ImplicitEnumDefType}:              _GotoState90Action,
	{_State144, ExplicitEnumDefType}:              _GotoState88Action,
	{_State144, TraitDefType}:                     _GotoState94Action,
	{_State144, FuncTypeType}:                     _GotoState89Action,
	{_State145, IntegerLiteralToken}:              _GotoState25Action,
	{_State145, FloatLiteralToken}:                _GotoState22Action,
	{_State145, RuneLiteralToken}:                 _GotoState33Action,
	{_State145, StringLiteralToken}:               _GotoState34Action,
	{_State145, IdentifierToken}:                  _GotoState24Action,
	{_State145, TrueToken}:                        _GotoState37Action,
	{_State145, FalseToken}:                       _GotoState21Action,
	{_State145, ReturnToken}:                      _GotoState226Action,
	{_State145, BreakToken}:                       _GotoState222Action,
	{_State145, ContinueToken}:                    _GotoState223Action,
	{_State145, UnsafeToken}:                      _GotoState169Action,
	{_State145, StructToken}:                      _GotoState35Action,
	{_State145, FuncToken}:                        _GotoState23Action,
	{_State145, AsyncToken}:                       _GotoState221Action,
	{_State145, DeferToken}:                       _GotoState224Action,
	{_State145, VarToken}:                         _GotoState38Action,
	{_State145, LetToken}:                         _GotoState28Action,
	{_State145, LabelDeclToken}:                   _GotoState26Action,
	{_State145, RbraceToken}:                      _GotoState225Action,
	{_State145, LparenToken}:                      _GotoState30Action,
	{_State145, LbracketToken}:                    _GotoState27Action,
	{_State145, NotToken}:                         _GotoState32Action,
	{_State145, SubToken}:                         _GotoState36Action,
	{_State145, MulToken}:                         _GotoState31Action,
	{_State145, BitNegToken}:                      _GotoState20Action,
	{_State145, BitAndToken}:                      _GotoState19Action,
	{_State145, LexErrorToken}:                    _GotoState29Action,
	{_State145, VarDeclPatternType}:               _GotoState58Action,
	{_State145, VarOrLetType}:                     _GotoState59Action,
	{_State145, AssignPatternType}:                _GotoState228Action,
	{_State145, ExpressionType}:                   _GotoState229Action,
	{_State145, OptionalLabelDeclType}:            _GotoState52Action,
	{_State145, SequenceExprType}:                 _GotoState233Action,
	{_State145, BlockExprType}:                    _GotoState44Action,
	{_State145, StatementType}:                    _GotoState234Action,
	{_State145, StatementBodyType}:                _GotoState235Action,
	{_State145, UnsafeStatementType}:              _GotoState236Action,
	{_State145, JumpStatementType}:                _GotoState231Action,
	{_State145, JumpTypeType}:                     _GotoState232Action,
	{_State145, ExpressionsType}:                  _GotoState230Action,
	{_State145, CallExprType}:                     _GotoState45Action,
	{_State145, AtomExprType}:                     _GotoState43Action,
	{_State145, LiteralType}:                      _GotoState50Action,
	{_State145, ImplicitStructExprType}:           _GotoState48Action,
	{_State145, AccessExprType}:                   _GotoState227Action,
	{_State145, PostfixUnaryExprType}:             _GotoState54Action,
	{_State145, PrefixUnaryOpType}:                _GotoState56Action,
	{_State145, PrefixUnaryExprType}:              _GotoState55Action,
	{_State145, MulExprType}:                      _GotoState51Action,
	{_State145, AddExprType}:                      _GotoState40Action,
	{_State145, CmpExprType}:                      _GotoState46Action,
	{_State145, AndExprType}:                      _GotoState41Action,
	{_State145, OrExprType}:                       _GotoState53Action,
	{_State145, InitializableTypeType}:            _GotoState49Action,
	{_State145, ExplicitStructDefType}:            _GotoState47Action,
	{_State145, AnonymousFuncExprType}:            _GotoState42Action,
	{_State146, PackageToken}:                     _GotoState16Action,
	{_State146, TypeToken}:                        _GotoState17Action,
	{_State146, FuncToken}:                        _GotoState18Action,
	{_State146, VarToken}:                         _GotoState38Action,
	{_State146, LetToken}:                         _GotoState28Action,
	{_State146, LbraceToken}:                      _GotoState62Action,
	{_State146, DefinitionType}:                   _GotoState237Action,
	{_State146, VarDeclPatternType}:               _GotoState69Action,
	{_State146, VarOrLetType}:                     _GotoState59Action,
	{_State146, BlockBodyType}:                    _GotoState63Action,
	{_State146, TypeDefType}:                      _GotoState68Action,
	{_State146, NamedFuncDefType}:                 _GotoState66Action,
	{_State146, PackageDefType}:                   _GotoState67Action,
	{_State148, IntegerLiteralToken}:              _GotoState25Action,
	{_State148, FloatLiteralToken}:                _GotoState22Action,
	{_State148, RuneLiteralToken}:                 _GotoState33Action,
	{_State148, StringLiteralToken}:               _GotoState34Action,
	{_State148, IdentifierToken}:                  _GotoState24Action,
	{_State148, TrueToken}:                        _GotoState37Action,
	{_State148, FalseToken}:                       _GotoState21Action,
	{_State148, StructToken}:                      _GotoState35Action,
	{_State148, FuncToken}:                        _GotoState23Action,
	{_State148, VarToken}:                         _GotoState38Action,
	{_State148, LetToken}:                         _GotoState28Action,
	{_State148, LabelDeclToken}:                   _GotoState26Action,
	{_State148, LparenToken}:                      _GotoState30Action,
	{_State148, LbracketToken}:                    _GotoState27Action,
	{_State148, NotToken}:                         _GotoState32Action,
	{_State148, SubToken}:                         _GotoState36Action,
	{_State148, MulToken}:                         _GotoState31Action,
	{_State148, BitNegToken}:                      _GotoState20Action,
	{_State148, BitAndToken}:                      _GotoState19Action,
	{_State148, LexErrorToken}:                    _GotoState29Action,
	{_State148, VarDeclPatternType}:               _GotoState58Action,
	{_State148, VarOrLetType}:                     _GotoState59Action,
	{_State148, ExpressionType}:                   _GotoState238Action,
	{_State148, OptionalLabelDeclType}:            _GotoState52Action,
	{_State148, SequenceExprType}:                 _GotoState57Action,
	{_State148, BlockExprType}:                    _GotoState44Action,
	{_State148, CallExprType}:                     _GotoState45Action,
	{_State148, AtomExprType}:                     _GotoState43Action,
	{_State148, LiteralType}:                      _GotoState50Action,
	{_State148, ImplicitStructExprType}:           _GotoState48Action,
	{_State148, AccessExprType}:                   _GotoState39Action,
	{_State148, PostfixUnaryExprType}:             _GotoState54Action,
	{_State148, PrefixUnaryOpType}:                _GotoState56Action,
	{_State148, PrefixUnaryExprType}:              _GotoState55Action,
	{_State148, MulExprType}:                      _GotoState51Action,
	{_State148, AddExprType}:                      _GotoState40Action,
	{_State148, CmpExprType}:                      _GotoState46Action,
	{_State148, AndExprType}:                      _GotoState41Action,
	{_State148, OrExprType}:                       _GotoState53Action,
	{_State148, InitializableTypeType}:            _GotoState49Action,
	{_State148, ExplicitStructDefType}:            _GotoState47Action,
	{_State148, AnonymousFuncExprType}:            _GotoState42Action,
	{_State149, ImportToken}:                      _GotoState239Action,
	{_State149, UnsafeToken}:                      _GotoState169Action,
	{_State149, RparenToken}:                      _GotoState240Action,
	{_State149, UnsafeStatementType}:              _GotoState244Action,
	{_State149, PackageStatementBodyType}:         _GotoState243Action,
	{_State149, PackageStatementType}:             _GotoState242Action,
	{_State149, ImportStatementType}:              _GotoState241Action,
	{_State150, IdentifierToken}:                  _GotoState81Action,
	{_State150, StructToken}:                      _GotoState35Action,
	{_State150, EnumToken}:                        _GotoState78Action,
	{_State150, TraitToken}:                       _GotoState86Action,
	{_State150, FuncToken}:                        _GotoState80Action,
	{_State150, LparenToken}:                      _GotoState83Action,
	{_State150, LbracketToken}:                    _GotoState27Action,
	{_State150, DotToken}:                         _GotoState77Action,
	{_State150, QuestionToken}:                    _GotoState84Action,
	{_State150, ExclaimToken}:                     _GotoState79Action,
	{_State150, TildeTildeToken}:                  _GotoState85Action,
	{_State150, BitNegToken}:                      _GotoState76Action,
	{_State150, BitAndToken}:                      _GotoState75Action,
	{_State150, LexErrorToken}:                    _GotoState82Action,
	{_State150, InitializableTypeType}:            _GotoState92Action,
	{_State150, AtomTypeType}:                     _GotoState87Action,
	{_State150, ReturnableTypeType}:               _GotoState93Action,
	{_State150, ValueTypeType}:                    _GotoState245Action,
	{_State150, ImplicitStructDefType}:            _GotoState91Action,
	{_State150, ExplicitStructDefType}:            _GotoState47Action,
	{_State150, ImplicitEnumDefType}:              _GotoState90Action,
	{_State150, ExplicitEnumDefType}:              _GotoState88Action,
	{_State150, TraitDefType}:                     _GotoState94Action,
	{_State150, FuncTypeType}:                     _GotoState89Action,
	{_State151, IdentifierToken}:                  _GotoState246Action,
	{_State151, GenericParameterDefType}:          _GotoState247Action,
	{_State151, GenericParameterDefsType}:         _GotoState248Action,
	{_State151, OptionalGenericParameterDefsType}: _GotoState249Action,
	{_State152, IdentifierToken}:                  _GotoState81Action,
	{_State152, StructToken}:                      _GotoState35Action,
	{_State152, EnumToken}:                        _GotoState78Action,
	{_State152, TraitToken}:                       _GotoState86Action,
	{_State152, FuncToken}:                        _GotoState80Action,
	{_State152, LparenToken}:                      _GotoState83Action,
	{_State152, LbracketToken}:                    _GotoState27Action,
	{_State152, DotToken}:                         _GotoState77Action,
	{_State152, QuestionToken}:                    _GotoState84Action,
	{_State152, ExclaimToken}:                     _GotoState79Action,
	{_State152, TildeTildeToken}:                  _GotoState85Action,
	{_State152, BitNegToken}:                      _GotoState76Action,
	{_State152, BitAndToken}:                      _GotoState75Action,
	{_State152, LexErrorToken}:                    _GotoState82Action,
	{_State152, InitializableTypeType}:            _GotoState92Action,
	{_State152, AtomTypeType}:                     _GotoState87Action,
	{_State152, ReturnableTypeType}:               _GotoState93Action,
	{_State152, ValueTypeType}:                    _GotoState250Action,
	{_State152, ImplicitStructDefType}:            _GotoState91Action,
	{_State152, ExplicitStructDefType}:            _GotoState47Action,
	{_State152, ImplicitEnumDefType}:              _GotoState90Action,
	{_State152, ExplicitEnumDefType}:              _GotoState88Action,
	{_State152, TraitDefType}:                     _GotoState94Action,
	{_State152, FuncTypeType}:                     _GotoState89Action,
	{_State153, IntegerLiteralToken}:              _GotoState25Action,
	{_State153, FloatLiteralToken}:                _GotoState22Action,
	{_State153, RuneLiteralToken}:                 _GotoState33Action,
	{_State153, StringLiteralToken}:               _GotoState34Action,
	{_State153, IdentifierToken}:                  _GotoState24Action,
	{_State153, TrueToken}:                        _GotoState37Action,
	{_State153, FalseToken}:                       _GotoState21Action,
	{_State153, StructToken}:                      _GotoState35Action,
	{_State153, FuncToken}:                        _GotoState23Action,
	{_State153, VarToken}:                         _GotoState38Action,
	{_State153, LetToken}:                         _GotoState28Action,
	{_State153, LabelDeclToken}:                   _GotoState26Action,
	{_State153, LparenToken}:                      _GotoState30Action,
	{_State153, LbracketToken}:                    _GotoState27Action,
	{_State153, NotToken}:                         _GotoState32Action,
	{_State153, SubToken}:                         _GotoState36Action,
	{_State153, MulToken}:                         _GotoState31Action,
	{_State153, BitNegToken}:                      _GotoState20Action,
	{_State153, BitAndToken}:                      _GotoState19Action,
	{_State153, LexErrorToken}:                    _GotoState29Action,
	{_State153, VarDeclPatternType}:               _GotoState58Action,
	{_State153, VarOrLetType}:                     _GotoState59Action,
	{_State153, ExpressionType}:                   _GotoState251Action,
	{_State153, OptionalLabelDeclType}:            _GotoState52Action,
	{_State153, SequenceExprType}:                 _GotoState57Action,
	{_State153, BlockExprType}:                    _GotoState44Action,
	{_State153, CallExprType}:                     _GotoState45Action,
	{_State153, AtomExprType}:                     _GotoState43Action,
	{_State153, LiteralType}:                      _GotoState50Action,
	{_State153, ImplicitStructExprType}:           _GotoState48Action,
	{_State153, AccessExprType}:                   _GotoState39Action,
	{_State153, PostfixUnaryExprType}:             _GotoState54Action,
	{_State153, PrefixUnaryOpType}:                _GotoState56Action,
	{_State153, PrefixUnaryExprType}:              _GotoState55Action,
	{_State153, MulExprType}:                      _GotoState51Action,
	{_State153, AddExprType}:                      _GotoState40Action,
	{_State153, CmpExprType}:                      _GotoState46Action,
	{_State153, AndExprType}:                      _GotoState41Action,
	{_State153, OrExprType}:                       _GotoState53Action,
	{_State153, InitializableTypeType}:            _GotoState49Action,
	{_State153, ExplicitStructDefType}:            _GotoState47Action,
	{_State153, AnonymousFuncExprType}:            _GotoState42Action,
	{_State154, LparenToken}:                      _GotoState252Action,
	{_State155, IdentifierToken}:                  _GotoState81Action,
	{_State155, StructToken}:                      _GotoState35Action,
	{_State155, EnumToken}:                        _GotoState78Action,
	{_State155, TraitToken}:                       _GotoState86Action,
	{_State155, FuncToken}:                        _GotoState80Action,
	{_State155, LparenToken}:                      _GotoState83Action,
	{_State155, LbracketToken}:                    _GotoState27Action,
	{_State155, DotToken}:                         _GotoState77Action,
	{_State155, QuestionToken}:                    _GotoState84Action,
	{_State155, ExclaimToken}:                     _GotoState79Action,
	{_State155, DotdotdotToken}:                   _GotoState253Action,
	{_State155, TildeTildeToken}:                  _GotoState85Action,
	{_State155, BitNegToken}:                      _GotoState76Action,
	{_State155, BitAndToken}:                      _GotoState75Action,
	{_State155, LexErrorToken}:                    _GotoState82Action,
	{_State155, InitializableTypeType}:            _GotoState92Action,
	{_State155, AtomTypeType}:                     _GotoState87Action,
	{_State155, ReturnableTypeType}:               _GotoState93Action,
	{_State155, ValueTypeType}:                    _GotoState254Action,
	{_State155, ImplicitStructDefType}:            _GotoState91Action,
	{_State155, ExplicitStructDefType}:            _GotoState47Action,
	{_State155, ImplicitEnumDefType}:              _GotoState90Action,
	{_State155, ExplicitEnumDefType}:              _GotoState88Action,
	{_State155, TraitDefType}:                     _GotoState94Action,
	{_State155, FuncTypeType}:                     _GotoState89Action,
	{_State156, RparenToken}:                      _GotoState255Action,
	{_State157, RparenToken}:                      _GotoState256Action,
	{_State159, CommaToken}:                       _GotoState257Action,
	{_State163, IdentifierToken}:                  _GotoState168Action,
	{_State163, UnsafeToken}:                      _GotoState169Action,
	{_State163, StructToken}:                      _GotoState35Action,
	{_State163, EnumToken}:                        _GotoState78Action,
	{_State163, TraitToken}:                       _GotoState86Action,
	{_State163, FuncToken}:                        _GotoState80Action,
	{_State163, LparenToken}:                      _GotoState83Action,
	{_State163, LbracketToken}:                    _GotoState27Action,
	{_State163, DotToken}:                         _GotoState77Action,
	{_State163, QuestionToken}:                    _GotoState84Action,
	{_State163, ExclaimToken}:                     _GotoState79Action,
	{_State163, TildeTildeToken}:                  _GotoState85Action,
	{_State163, BitNegToken}:                      _GotoState76Action,
	{_State163, BitAndToken}:                      _GotoState75Action,
	{_State163, LexErrorToken}:                    _GotoState82Action,
	{_State163, UnsafeStatementType}:              _GotoState175Action,
	{_State163, InitializableTypeType}:            _GotoState92Action,
	{_State163, AtomTypeType}:                     _GotoState87Action,
	{_State163, ReturnableTypeType}:               _GotoState93Action,
	{_State163, ValueTypeType}:                    _GotoState176Action,
	{_State163, FieldDefType}:                     _GotoState260Action,
	{_State163, ImplicitStructDefType}:            _GotoState91Action,
	{_State163, ExplicitStructDefType}:            _GotoState47Action,
	{_State163, EnumValueDefType}:                 _GotoState258Action,
	{_State163, ImplicitEnumValueDefsType}:        _GotoState261Action,
	{_State163, ImplicitEnumDefType}:              _GotoState90Action,
	{_State163, ExplicitEnumValueDefsType}:        _GotoState259Action,
	{_State163, ExplicitEnumDefType}:              _GotoState88Action,
	{_State163, TraitDefType}:                     _GotoState94Action,
	{_State163, FuncTypeType}:                     _GotoState89Action,
	{_State165, IdentifierToken}:                  _GotoState263Action,
	{_State165, StructToken}:                      _GotoState35Action,
	{_State165, EnumToken}:                        _GotoState78Action,
	{_State165, TraitToken}:                       _GotoState86Action,
	{_State165, FuncToken}:                        _GotoState80Action,
	{_State165, LparenToken}:                      _GotoState83Action,
	{_State165, LbracketToken}:                    _GotoState27Action,
	{_State165, DotToken}:                         _GotoState77Action,
	{_State165, QuestionToken}:                    _GotoState84Action,
	{_State165, ExclaimToken}:                     _GotoState79Action,
	{_State165, DotdotdotToken}:                   _GotoState262Action,
	{_State165, TildeTildeToken}:                  _GotoState85Action,
	{_State165, BitNegToken}:                      _GotoState76Action,
	{_State165, BitAndToken}:                      _GotoState75Action,
	{_State165, LexErrorToken}:                    _GotoState82Action,
	{_State165, InitializableTypeType}:            _GotoState92Action,
	{_State165, AtomTypeType}:                     _GotoState87Action,
	{_State165, ReturnableTypeType}:               _GotoState93Action,
	{_State165, ValueTypeType}:                    _GotoState267Action,
	{_State165, ImplicitStructDefType}:            _GotoState91Action,
	{_State165, ExplicitStructDefType}:            _GotoState47Action,
	{_State165, ImplicitEnumDefType}:              _GotoState90Action,
	{_State165, ExplicitEnumDefType}:              _GotoState88Action,
	{_State165, TraitDefType}:                     _GotoState94Action,
	{_State165, ParameterDeclType}:                _GotoState265Action,
	{_State165, ParameterDeclsType}:               _GotoState266Action,
	{_State165, OptionalParameterDeclsType}:       _GotoState264Action,
	{_State165, FuncTypeType}:                     _GotoState89Action,
	{_State166, IdentifierToken}:                  _GotoState268Action,
	{_State168, IdentifierToken}:                  _GotoState81Action,
	{_State168, StructToken}:                      _GotoState35Action,
	{_State168, EnumToken}:                        _GotoState78Action,
	{_State168, TraitToken}:                       _GotoState86Action,
	{_State168, FuncToken}:                        _GotoState80Action,
	{_State168, LparenToken}:                      _GotoState83Action,
	{_State168, LbracketToken}:                    _GotoState27Action,
	{_State168, DotToken}:                         _GotoState269Action,
	{_State168, QuestionToken}:                    _GotoState84Action,
	{_State168, ExclaimToken}:                     _GotoState79Action,
	{_State168, DollarLbracketToken}:              _GotoState104Action,
	{_State168, TildeTildeToken}:                  _GotoState85Action,
	{_State168, BitNegToken}:                      _GotoState76Action,
	{_State168, BitAndToken}:                      _GotoState75Action,
	{_State168, LexErrorToken}:                    _GotoState82Action,
	{_State168, OptionalGenericBindingType}:       _GotoState167Action,
	{_State168, InitializableTypeType}:            _GotoState92Action,
	{_State168, AtomTypeType}:                     _GotoState87Action,
	{_State168, ReturnableTypeType}:               _GotoState93Action,
	{_State168, ValueTypeType}:                    _GotoState270Action,
	{_State168, ImplicitStructDefType}:            _GotoState91Action,
	{_State168, ExplicitStructDefType}:            _GotoState47Action,
	{_State168, ImplicitEnumDefType}:              _GotoState90Action,
	{_State168, ExplicitEnumDefType}:              _GotoState88Action,
	{_State168, TraitDefType}:                     _GotoState94Action,
	{_State168, FuncTypeType}:                     _GotoState89Action,
	{_State169, LessToken}:                        _GotoState271Action,
	{_State170, OrToken}:                          _GotoState272Action,
	{_State171, AssignToken}:                      _GotoState273Action,
	{_State172, RparenToken}:                      _GotoState275Action,
	{_State172, OrToken}:                          _GotoState274Action,
	{_State173, CommaToken}:                       _GotoState276Action,
	{_State174, RparenToken}:                      _GotoState277Action,
	{_State176, AddToken}:                         _GotoState180Action,
	{_State176, SubToken}:                         _GotoState185Action,
	{_State176, MulToken}:                         _GotoState183Action,
	{_State179, IdentifierToken}:                  _GotoState168Action,
	{_State179, UnsafeToken}:                      _GotoState169Action,
	{_State179, StructToken}:                      _GotoState35Action,
	{_State179, EnumToken}:                        _GotoState78Action,
	{_State179, TraitToken}:                       _GotoState86Action,
	{_State179, FuncToken}:                        _GotoState278Action,
	{_State179, LparenToken}:                      _GotoState83Action,
	{_State179, LbracketToken}:                    _GotoState27Action,
	{_State179, DotToken}:                         _GotoState77Action,
	{_State179, QuestionToken}:                    _GotoState84Action,
	{_State179, ExclaimToken}:                     _GotoState79Action,
	{_State179, TildeTildeToken}:                  _GotoState85Action,
	{_State179, BitNegToken}:                      _GotoState76Action,
	{_State179, BitAndToken}:                      _GotoState75Action,
	{_State179, LexErrorToken}:                    _GotoState82Action,
	{_State179, UnsafeStatementType}:              _GotoState175Action,
	{_State179, InitializableTypeType}:            _GotoState92Action,
	{_State179, AtomTypeType}:                     _GotoState87Action,
	{_State179, ReturnableTypeType}:               _GotoState93Action,
	{_State179, ValueTypeType}:                    _GotoState176Action,
	{_State179, FieldDefType}:                     _GotoState279Action,
	{_State179, ImplicitStructDefType}:            _GotoState91Action,
	{_State179, ExplicitStructDefType}:            _GotoState47Action,
	{_State179, ImplicitEnumDefType}:              _GotoState90Action,
	{_State179, ExplicitEnumDefType}:              _GotoState88Action,
	{_State179, TraitPropertyType}:                _GotoState283Action,
	{_State179, TraitPropertiesType}:              _GotoState282Action,
	{_State179, OptionalTraitPropertiesType}:      _GotoState281Action,
	{_State179, TraitDefType}:                     _GotoState94Action,
	{_State179, FuncTypeType}:                     _GotoState89Action,
	{_State179, MethodSignatureType}:              _GotoState280Action,
	{_State180, IdentifierToken}:                  _GotoState81Action,
	{_State180, StructToken}:                      _GotoState35Action,
	{_State180, EnumToken}:                        _GotoState78Action,
	{_State180, TraitToken}:                       _GotoState86Action,
	{_State180, FuncToken}:                        _GotoState80Action,
	{_State180, LparenToken}:                      _GotoState83Action,
	{_State180, LbracketToken}:                    _GotoState27Action,
	{_State180, DotToken}:                         _GotoState77Action,
	{_State180, QuestionToken}:                    _GotoState84Action,
	{_State180, ExclaimToken}:                     _GotoState79Action,
	{_State180, TildeTildeToken}:                  _GotoState85Action,
	{_State180, BitNegToken}:                      _GotoState76Action,
	{_State180, BitAndToken}:                      _GotoState75Action,
	{_State180, LexErrorToken}:                    _GotoState82Action,
	{_State180, InitializableTypeType}:            _GotoState92Action,
	{_State180, AtomTypeType}:                     _GotoState87Action,
	{_State180, ReturnableTypeType}:               _GotoState284Action,
	{_State180, ImplicitStructDefType}:            _GotoState91Action,
	{_State180, ExplicitStructDefType}:            _GotoState47Action,
	{_State180, ImplicitEnumDefType}:              _GotoState90Action,
	{_State180, ExplicitEnumDefType}:              _GotoState88Action,
	{_State180, TraitDefType}:                     _GotoState94Action,
	{_State180, FuncTypeType}:                     _GotoState89Action,
	{_State181, IdentifierToken}:                  _GotoState81Action,
	{_State181, StructToken}:                      _GotoState35Action,
	{_State181, EnumToken}:                        _GotoState78Action,
	{_State181, TraitToken}:                       _GotoState86Action,
	{_State181, FuncToken}:                        _GotoState80Action,
	{_State181, LparenToken}:                      _GotoState83Action,
	{_State181, LbracketToken}:                    _GotoState27Action,
	{_State181, DotToken}:                         _GotoState77Action,
	{_State181, QuestionToken}:                    _GotoState84Action,
	{_State181, ExclaimToken}:                     _GotoState79Action,
	{_State181, TildeTildeToken}:                  _GotoState85Action,
	{_State181, BitNegToken}:                      _GotoState76Action,
	{_State181, BitAndToken}:                      _GotoState75Action,
	{_State181, LexErrorToken}:                    _GotoState82Action,
	{_State181, InitializableTypeType}:            _GotoState92Action,
	{_State181, AtomTypeType}:                     _GotoState87Action,
	{_State181, ReturnableTypeType}:               _GotoState93Action,
	{_State181, ValueTypeType}:                    _GotoState285Action,
	{_State181, ImplicitStructDefType}:            _GotoState91Action,
	{_State181, ExplicitStructDefType}:            _GotoState47Action,
	{_State181, ImplicitEnumDefType}:              _GotoState90Action,
	{_State181, ExplicitEnumDefType}:              _GotoState88Action,
	{_State181, TraitDefType}:                     _GotoState94Action,
	{_State181, FuncTypeType}:                     _GotoState89Action,
	{_State182, IntegerLiteralToken}:              _GotoState286Action,
	{_State183, IdentifierToken}:                  _GotoState81Action,
	{_State183, StructToken}:                      _GotoState35Action,
	{_State183, EnumToken}:                        _GotoState78Action,
	{_State183, TraitToken}:                       _GotoState86Action,
	{_State183, FuncToken}:                        _GotoState80Action,
	{_State183, LparenToken}:                      _GotoState83Action,
	{_State183, LbracketToken}:                    _GotoState27Action,
	{_State183, DotToken}:                         _GotoState77Action,
	{_State183, QuestionToken}:                    _GotoState84Action,
	{_State183, ExclaimToken}:                     _GotoState79Action,
	{_State183, TildeTildeToken}:                  _GotoState85Action,
	{_State183, BitNegToken}:                      _GotoState76Action,
	{_State183, BitAndToken}:                      _GotoState75Action,
	{_State183, LexErrorToken}:                    _GotoState82Action,
	{_State183, InitializableTypeType}:            _GotoState92Action,
	{_State183, AtomTypeType}:                     _GotoState87Action,
	{_State183, ReturnableTypeType}:               _GotoState287Action,
	{_State183, ImplicitStructDefType}:            _GotoState91Action,
	{_State183, ExplicitStructDefType}:            _GotoState47Action,
	{_State183, ImplicitEnumDefType}:              _GotoState90Action,
	{_State183, ExplicitEnumDefType}:              _GotoState88Action,
	{_State183, TraitDefType}:                     _GotoState94Action,
	{_State183, FuncTypeType}:                     _GotoState89Action,
	{_State185, IdentifierToken}:                  _GotoState81Action,
	{_State185, StructToken}:                      _GotoState35Action,
	{_State185, EnumToken}:                        _GotoState78Action,
	{_State185, TraitToken}:                       _GotoState86Action,
	{_State185, FuncToken}:                        _GotoState80Action,
	{_State185, LparenToken}:                      _GotoState83Action,
	{_State185, LbracketToken}:                    _GotoState27Action,
	{_State185, DotToken}:                         _GotoState77Action,
	{_State185, QuestionToken}:                    _GotoState84Action,
	{_State185, ExclaimToken}:                     _GotoState79Action,
	{_State185, TildeTildeToken}:                  _GotoState85Action,
	{_State185, BitNegToken}:                      _GotoState76Action,
	{_State185, BitAndToken}:                      _GotoState75Action,
	{_State185, LexErrorToken}:                    _GotoState82Action,
	{_State185, InitializableTypeType}:            _GotoState92Action,
	{_State185, AtomTypeType}:                     _GotoState87Action,
	{_State185, ReturnableTypeType}:               _GotoState288Action,
	{_State185, ImplicitStructDefType}:            _GotoState91Action,
	{_State185, ExplicitStructDefType}:            _GotoState47Action,
	{_State185, ImplicitEnumDefType}:              _GotoState90Action,
	{_State185, ExplicitEnumDefType}:              _GotoState88Action,
	{_State185, TraitDefType}:                     _GotoState94Action,
	{_State185, FuncTypeType}:                     _GotoState89Action,
	{_State186, IntegerLiteralToken}:              _GotoState25Action,
	{_State186, FloatLiteralToken}:                _GotoState22Action,
	{_State186, RuneLiteralToken}:                 _GotoState33Action,
	{_State186, StringLiteralToken}:               _GotoState34Action,
	{_State186, IdentifierToken}:                  _GotoState24Action,
	{_State186, TrueToken}:                        _GotoState37Action,
	{_State186, FalseToken}:                       _GotoState21Action,
	{_State186, StructToken}:                      _GotoState35Action,
	{_State186, FuncToken}:                        _GotoState23Action,
	{_State186, VarToken}:                         _GotoState38Action,
	{_State186, LetToken}:                         _GotoState28Action,
	{_State186, LabelDeclToken}:                   _GotoState26Action,
	{_State186, LparenToken}:                      _GotoState30Action,
	{_State186, LbracketToken}:                    _GotoState27Action,
	{_State186, NotToken}:                         _GotoState32Action,
	{_State186, SubToken}:                         _GotoState36Action,
	{_State186, MulToken}:                         _GotoState31Action,
	{_State186, BitNegToken}:                      _GotoState20Action,
	{_State186, BitAndToken}:                      _GotoState19Action,
	{_State186, LexErrorToken}:                    _GotoState29Action,
	{_State186, VarDeclPatternType}:               _GotoState58Action,
	{_State186, VarOrLetType}:                     _GotoState59Action,
	{_State186, ExpressionType}:                   _GotoState289Action,
	{_State186, OptionalLabelDeclType}:            _GotoState52Action,
	{_State186, SequenceExprType}:                 _GotoState57Action,
	{_State186, BlockExprType}:                    _GotoState44Action,
	{_State186, CallExprType}:                     _GotoState45Action,
	{_State186, AtomExprType}:                     _GotoState43Action,
	{_State186, LiteralType}:                      _GotoState50Action,
	{_State186, ImplicitStructExprType}:           _GotoState48Action,
	{_State186, AccessExprType}:                   _GotoState39Action,
	{_State186, PostfixUnaryExprType}:             _GotoState54Action,
	{_State186, PrefixUnaryOpType}:                _GotoState56Action,
	{_State186, PrefixUnaryExprType}:              _GotoState55Action,
	{_State186, MulExprType}:                      _GotoState51Action,
	{_State186, AddExprType}:                      _GotoState40Action,
	{_State186, CmpExprType}:                      _GotoState46Action,
	{_State186, AndExprType}:                      _GotoState41Action,
	{_State186, OrExprType}:                       _GotoState53Action,
	{_State186, InitializableTypeType}:            _GotoState49Action,
	{_State186, ExplicitStructDefType}:            _GotoState47Action,
	{_State186, AnonymousFuncExprType}:            _GotoState42Action,
	{_State187, IntegerLiteralToken}:              _GotoState25Action,
	{_State187, FloatLiteralToken}:                _GotoState22Action,
	{_State187, RuneLiteralToken}:                 _GotoState33Action,
	{_State187, StringLiteralToken}:               _GotoState34Action,
	{_State187, IdentifierToken}:                  _GotoState97Action,
	{_State187, TrueToken}:                        _GotoState37Action,
	{_State187, FalseToken}:                       _GotoState21Action,
	{_State187, StructToken}:                      _GotoState35Action,
	{_State187, FuncToken}:                        _GotoState23Action,
	{_State187, VarToken}:                         _GotoState38Action,
	{_State187, LetToken}:                         _GotoState28Action,
	{_State187, LabelDeclToken}:                   _GotoState26Action,
	{_State187, LparenToken}:                      _GotoState30Action,
	{_State187, LbracketToken}:                    _GotoState27Action,
	{_State187, DotdotdotToken}:                   _GotoState96Action,
	{_State187, NotToken}:                         _GotoState32Action,
	{_State187, SubToken}:                         _GotoState36Action,
	{_State187, MulToken}:                         _GotoState31Action,
	{_State187, BitNegToken}:                      _GotoState20Action,
	{_State187, BitAndToken}:                      _GotoState19Action,
	{_State187, LexErrorToken}:                    _GotoState29Action,
	{_State187, VarDeclPatternType}:               _GotoState58Action,
	{_State187, VarOrLetType}:                     _GotoState59Action,
	{_State187, ExpressionType}:                   _GotoState101Action,
	{_State187, OptionalLabelDeclType}:            _GotoState52Action,
	{_State187, SequenceExprType}:                 _GotoState57Action,
	{_State187, BlockExprType}:                    _GotoState44Action,
	{_State187, CallExprType}:                     _GotoState45Action,
	{_State187, ArgumentType}:                     _GotoState290Action,
	{_State187, ColonExpressionsType}:             _GotoState100Action,
	{_State187, OptionalExpressionType}:           _GotoState102Action,
	{_State187, AtomExprType}:                     _GotoState43Action,
	{_State187, LiteralType}:                      _GotoState50Action,
	{_State187, ImplicitStructExprType}:           _GotoState48Action,
	{_State187, AccessExprType}:                   _GotoState39Action,
	{_State187, PostfixUnaryExprType}:             _GotoState54Action,
	{_State187, PrefixUnaryOpType}:                _GotoState56Action,
	{_State187, PrefixUnaryExprType}:              _GotoState55Action,
	{_State187, MulExprType}:                      _GotoState51Action,
	{_State187, AddExprType}:                      _GotoState40Action,
	{_State187, CmpExprType}:                      _GotoState46Action,
	{_State187, AndExprType}:                      _GotoState41Action,
	{_State187, OrExprType}:                       _GotoState53Action,
	{_State187, InitializableTypeType}:            _GotoState49Action,
	{_State187, ExplicitStructDefType}:            _GotoState47Action,
	{_State187, AnonymousFuncExprType}:            _GotoState42Action,
	{_State189, IntegerLiteralToken}:              _GotoState25Action,
	{_State189, FloatLiteralToken}:                _GotoState22Action,
	{_State189, RuneLiteralToken}:                 _GotoState33Action,
	{_State189, StringLiteralToken}:               _GotoState34Action,
	{_State189, IdentifierToken}:                  _GotoState24Action,
	{_State189, TrueToken}:                        _GotoState37Action,
	{_State189, FalseToken}:                       _GotoState21Action,
	{_State189, StructToken}:                      _GotoState35Action,
	{_State189, FuncToken}:                        _GotoState23Action,
	{_State189, VarToken}:                         _GotoState38Action,
	{_State189, LetToken}:                         _GotoState28Action,
	{_State189, LabelDeclToken}:                   _GotoState26Action,
	{_State189, LparenToken}:                      _GotoState30Action,
	{_State189, LbracketToken}:                    _GotoState27Action,
	{_State189, NotToken}:                         _GotoState32Action,
	{_State189, SubToken}:                         _GotoState36Action,
	{_State189, MulToken}:                         _GotoState31Action,
	{_State189, BitNegToken}:                      _GotoState20Action,
	{_State189, BitAndToken}:                      _GotoState19Action,
	{_State189, LexErrorToken}:                    _GotoState29Action,
	{_State189, VarDeclPatternType}:               _GotoState58Action,
	{_State189, VarOrLetType}:                     _GotoState59Action,
	{_State189, ExpressionType}:                   _GotoState291Action,
	{_State189, OptionalLabelDeclType}:            _GotoState52Action,
	{_State189, SequenceExprType}:                 _GotoState57Action,
	{_State189, BlockExprType}:                    _GotoState44Action,
	{_State189, CallExprType}:                     _GotoState45Action,
	{_State189, OptionalExpressionType}:           _GotoState292Action,
	{_State189, AtomExprType}:                     _GotoState43Action,
	{_State189, LiteralType}:                      _GotoState50Action,
	{_State189, ImplicitStructExprType}:           _GotoState48Action,
	{_State189, AccessExprType}:                   _GotoState39Action,
	{_State189, PostfixUnaryExprType}:             _GotoState54Action,
	{_State189, PrefixUnaryOpType}:                _GotoState56Action,
	{_State189, PrefixUnaryExprType}:              _GotoState55Action,
	{_State189, MulExprType}:                      _GotoState51Action,
	{_State189, AddExprType}:                      _GotoState40Action,
	{_State189, CmpExprType}:                      _GotoState46Action,
	{_State189, AndExprType}:                      _GotoState41Action,
	{_State189, OrExprType}:                       _GotoState53Action,
	{_State189, InitializableTypeType}:            _GotoState49Action,
	{_State189, ExplicitStructDefType}:            _GotoState47Action,
	{_State189, AnonymousFuncExprType}:            _GotoState42Action,
	{_State190, IntegerLiteralToken}:              _GotoState25Action,
	{_State190, FloatLiteralToken}:                _GotoState22Action,
	{_State190, RuneLiteralToken}:                 _GotoState33Action,
	{_State190, StringLiteralToken}:               _GotoState34Action,
	{_State190, IdentifierToken}:                  _GotoState24Action,
	{_State190, TrueToken}:                        _GotoState37Action,
	{_State190, FalseToken}:                       _GotoState21Action,
	{_State190, StructToken}:                      _GotoState35Action,
	{_State190, FuncToken}:                        _GotoState23Action,
	{_State190, VarToken}:                         _GotoState38Action,
	{_State190, LetToken}:                         _GotoState28Action,
	{_State190, LabelDeclToken}:                   _GotoState26Action,
	{_State190, LparenToken}:                      _GotoState30Action,
	{_State190, LbracketToken}:                    _GotoState27Action,
	{_State190, NotToken}:                         _GotoState32Action,
	{_State190, SubToken}:                         _GotoState36Action,
	{_State190, MulToken}:                         _GotoState31Action,
	{_State190, BitNegToken}:                      _GotoState20Action,
	{_State190, BitAndToken}:                      _GotoState19Action,
	{_State190, LexErrorToken}:                    _GotoState29Action,
	{_State190, VarDeclPatternType}:               _GotoState58Action,
	{_State190, VarOrLetType}:                     _GotoState59Action,
	{_State190, ExpressionType}:                   _GotoState291Action,
	{_State190, OptionalLabelDeclType}:            _GotoState52Action,
	{_State190, SequenceExprType}:                 _GotoState57Action,
	{_State190, BlockExprType}:                    _GotoState44Action,
	{_State190, CallExprType}:                     _GotoState45Action,
	{_State190, OptionalExpressionType}:           _GotoState293Action,
	{_State190, AtomExprType}:                     _GotoState43Action,
	{_State190, LiteralType}:                      _GotoState50Action,
	{_State190, ImplicitStructExprType}:           _GotoState48Action,
	{_State190, AccessExprType}:                   _GotoState39Action,
	{_State190, PostfixUnaryExprType}:             _GotoState54Action,
	{_State190, PrefixUnaryOpType}:                _GotoState56Action,
	{_State190, PrefixUnaryExprType}:              _GotoState55Action,
	{_State190, MulExprType}:                      _GotoState51Action,
	{_State190, AddExprType}:                      _GotoState40Action,
	{_State190, CmpExprType}:                      _GotoState46Action,
	{_State190, AndExprType}:                      _GotoState41Action,
	{_State190, OrExprType}:                       _GotoState53Action,
	{_State190, InitializableTypeType}:            _GotoState49Action,
	{_State190, ExplicitStructDefType}:            _GotoState47Action,
	{_State190, AnonymousFuncExprType}:            _GotoState42Action,
	{_State191, NewlinesToken}:                    _GotoState295Action,
	{_State191, CommaToken}:                       _GotoState294Action,
	{_State193, RparenToken}:                      _GotoState296Action,
	{_State194, CommaToken}:                       _GotoState297Action,
	{_State195, RbracketToken}:                    _GotoState298Action,
	{_State196, AddToken}:                         _GotoState180Action,
	{_State196, SubToken}:                         _GotoState185Action,
	{_State196, MulToken}:                         _GotoState183Action,
	{_State198, RbracketToken}:                    _GotoState299Action,
	{_State199, IntegerLiteralToken}:              _GotoState25Action,
	{_State199, FloatLiteralToken}:                _GotoState22Action,
	{_State199, RuneLiteralToken}:                 _GotoState33Action,
	{_State199, StringLiteralToken}:               _GotoState34Action,
	{_State199, IdentifierToken}:                  _GotoState97Action,
	{_State199, TrueToken}:                        _GotoState37Action,
	{_State199, FalseToken}:                       _GotoState21Action,
	{_State199, StructToken}:                      _GotoState35Action,
	{_State199, FuncToken}:                        _GotoState23Action,
	{_State199, VarToken}:                         _GotoState38Action,
	{_State199, LetToken}:                         _GotoState28Action,
	{_State199, LabelDeclToken}:                   _GotoState26Action,
	{_State199, LparenToken}:                      _GotoState30Action,
	{_State199, LbracketToken}:                    _GotoState27Action,
	{_State199, DotdotdotToken}:                   _GotoState96Action,
	{_State199, NotToken}:                         _GotoState32Action,
	{_State199, SubToken}:                         _GotoState36Action,
	{_State199, MulToken}:                         _GotoState31Action,
	{_State199, BitNegToken}:                      _GotoState20Action,
	{_State199, BitAndToken}:                      _GotoState19Action,
	{_State199, LexErrorToken}:                    _GotoState29Action,
	{_State199, VarDeclPatternType}:               _GotoState58Action,
	{_State199, VarOrLetType}:                     _GotoState59Action,
	{_State199, ExpressionType}:                   _GotoState101Action,
	{_State199, OptionalLabelDeclType}:            _GotoState52Action,
	{_State199, SequenceExprType}:                 _GotoState57Action,
	{_State199, BlockExprType}:                    _GotoState44Action,
	{_State199, CallExprType}:                     _GotoState45Action,
	{_State199, OptionalArgumentsType}:            _GotoState301Action,
	{_State199, ArgumentsType}:                    _GotoState300Action,
	{_State199, ArgumentType}:                     _GotoState98Action,
	{_State199, ColonExpressionsType}:             _GotoState100Action,
	{_State199, OptionalExpressionType}:           _GotoState102Action,
	{_State199, AtomExprType}:                     _GotoState43Action,
	{_State199, LiteralType}:                      _GotoState50Action,
	{_State199, ImplicitStructExprType}:           _GotoState48Action,
	{_State199, AccessExprType}:                   _GotoState39Action,
	{_State199, PostfixUnaryExprType}:             _GotoState54Action,
	{_State199, PrefixUnaryOpType}:                _GotoState56Action,
	{_State199, PrefixUnaryExprType}:              _GotoState55Action,
	{_State199, MulExprType}:                      _GotoState51Action,
	{_State199, AddExprType}:                      _GotoState40Action,
	{_State199, CmpExprType}:                      _GotoState46Action,
	{_State199, AndExprType}:                      _GotoState41Action,
	{_State199, OrExprType}:                       _GotoState53Action,
	{_State199, InitializableTypeType}:            _GotoState49Action,
	{_State199, ExplicitStructDefType}:            _GotoState47Action,
	{_State199, AnonymousFuncExprType}:            _GotoState42Action,
	{_State200, MulToken}:                         _GotoState128Action,
	{_State200, DivToken}:                         _GotoState126Action,
	{_State200, ModToken}:                         _GotoState127Action,
	{_State200, BitAndToken}:                      _GotoState123Action,
	{_State200, BitLshiftToken}:                   _GotoState124Action,
	{_State200, BitRshiftToken}:                   _GotoState125Action,
	{_State200, MulOpType}:                        _GotoState129Action,
	{_State201, EqualToken}:                       _GotoState115Action,
	{_State201, NotEqualToken}:                    _GotoState120Action,
	{_State201, LessToken}:                        _GotoState118Action,
	{_State201, LessOrEqualToken}:                 _GotoState119Action,
	{_State201, GreaterToken}:                     _GotoState116Action,
	{_State201, GreaterOrEqualToken}:              _GotoState117Action,
	{_State201, CmpOpType}:                        _GotoState121Action,
	{_State202, AddToken}:                         _GotoState109Action,
	{_State202, SubToken}:                         _GotoState112Action,
	{_State202, BitXorToken}:                      _GotoState111Action,
	{_State202, BitOrToken}:                       _GotoState110Action,
	{_State202, AddOpType}:                        _GotoState113Action,
	{_State203, RparenToken}:                      _GotoState302Action,
	{_State203, CommaToken}:                       _GotoState187Action,
	{_State205, ForToken}:                         _GotoState303Action,
	{_State206, InToken}:                          _GotoState305Action,
	{_State206, AssignToken}:                      _GotoState304Action,
	{_State207, SemicolonToken}:                   _GotoState306Action,
	{_State208, DoToken}:                          _GotoState307Action,
	{_State209, IntegerLiteralToken}:              _GotoState25Action,
	{_State209, FloatLiteralToken}:                _GotoState22Action,
	{_State209, RuneLiteralToken}:                 _GotoState33Action,
	{_State209, StringLiteralToken}:               _GotoState34Action,
	{_State209, IdentifierToken}:                  _GotoState24Action,
	{_State209, TrueToken}:                        _GotoState37Action,
	{_State209, FalseToken}:                       _GotoState21Action,
	{_State209, StructToken}:                      _GotoState35Action,
	{_State209, FuncToken}:                        _GotoState23Action,
	{_State209, VarToken}:                         _GotoState309Action,
	{_State209, LetToken}:                         _GotoState28Action,
	{_State209, LabelDeclToken}:                   _GotoState26Action,
	{_State209, LparenToken}:                      _GotoState30Action,
	{_State209, LbracketToken}:                    _GotoState27Action,
	{_State209, DotToken}:                         _GotoState308Action,
	{_State209, NotToken}:                         _GotoState32Action,
	{_State209, SubToken}:                         _GotoState36Action,
	{_State209, MulToken}:                         _GotoState31Action,
	{_State209, BitNegToken}:                      _GotoState20Action,
	{_State209, BitAndToken}:                      _GotoState19Action,
	{_State209, LexErrorToken}:                    _GotoState29Action,
	{_State209, VarDeclPatternType}:               _GotoState58Action,
	{_State209, VarOrLetType}:                     _GotoState59Action,
	{_State209, AssignPatternType}:                _GotoState310Action,
	{_State209, CasePatternType}:                  _GotoState311Action,
	{_State209, OptionalLabelDeclType}:            _GotoState139Action,
	{_State209, CasePatternsType}:                 _GotoState312Action,
	{_State209, SequenceExprType}:                 _GotoState313Action,
	{_State209, BlockExprType}:                    _GotoState44Action,
	{_State209, CallExprType}:                     _GotoState45Action,
	{_State209, AtomExprType}:                     _GotoState43Action,
	{_State209, LiteralType}:                      _GotoState50Action,
	{_State209, ImplicitStructExprType}:           _GotoState48Action,
	{_State209, AccessExprType}:                   _GotoState39Action,
	{_State209, PostfixUnaryExprType}:             _GotoState54Action,
	{_State209, PrefixUnaryOpType}:                _GotoState56Action,
	{_State209, PrefixUnaryExprType}:              _GotoState55Action,
	{_State209, MulExprType}:                      _GotoState51Action,
	{_State209, AddExprType}:                      _GotoState40Action,
	{_State209, CmpExprType}:                      _GotoState46Action,
	{_State209, AndExprType}:                      _GotoState41Action,
	{_State209, OrExprType}:                       _GotoState53Action,
	{_State209, InitializableTypeType}:            _GotoState49Action,
	{_State209, ExplicitStructDefType}:            _GotoState47Action,
	{_State209, AnonymousFuncExprType}:            _GotoState42Action,
	{_State210, LbraceToken}:                      _GotoState62Action,
	{_State210, BlockBodyType}:                    _GotoState314Action,
	{_State212, LbraceToken}:                      _GotoState315Action,
	{_State213, AndToken}:                         _GotoState114Action,
	{_State215, AssignToken}:                      _GotoState316Action,
	{_State217, RparenToken}:                      _GotoState318Action,
	{_State217, CommaToken}:                       _GotoState317Action,
	{_State220, AddToken}:                         _GotoState180Action,
	{_State220, SubToken}:                         _GotoState185Action,
	{_State220, MulToken}:                         _GotoState183Action,
	{_State221, IntegerLiteralToken}:              _GotoState25Action,
	{_State221, FloatLiteralToken}:                _GotoState22Action,
	{_State221, RuneLiteralToken}:                 _GotoState33Action,
	{_State221, StringLiteralToken}:               _GotoState34Action,
	{_State221, IdentifierToken}:                  _GotoState24Action,
	{_State221, TrueToken}:                        _GotoState37Action,
	{_State221, FalseToken}:                       _GotoState21Action,
	{_State221, StructToken}:                      _GotoState35Action,
	{_State221, FuncToken}:                        _GotoState23Action,
	{_State221, LabelDeclToken}:                   _GotoState26Action,
	{_State221, LparenToken}:                      _GotoState30Action,
	{_State221, LbracketToken}:                    _GotoState27Action,
	{_State221, LexErrorToken}:                    _GotoState29Action,
	{_State221, OptionalLabelDeclType}:            _GotoState139Action,
	{_State221, BlockExprType}:                    _GotoState44Action,
	{_State221, CallExprType}:                     _GotoState320Action,
	{_State221, AtomExprType}:                     _GotoState43Action,
	{_State221, LiteralType}:                      _GotoState50Action,
	{_State221, ImplicitStructExprType}:           _GotoState48Action,
	{_State221, AccessExprType}:                   _GotoState319Action,
	{_State221, InitializableTypeType}:            _GotoState49Action,
	{_State221, ExplicitStructDefType}:            _GotoState47Action,
	{_State221, AnonymousFuncExprType}:            _GotoState42Action,
	{_State224, IntegerLiteralToken}:              _GotoState25Action,
	{_State224, FloatLiteralToken}:                _GotoState22Action,
	{_State224, RuneLiteralToken}:                 _GotoState33Action,
	{_State224, StringLiteralToken}:               _GotoState34Action,
	{_State224, IdentifierToken}:                  _GotoState24Action,
	{_State224, TrueToken}:                        _GotoState37Action,
	{_State224, FalseToken}:                       _GotoState21Action,
	{_State224, StructToken}:                      _GotoState35Action,
	{_State224, FuncToken}:                        _GotoState23Action,
	{_State224, LabelDeclToken}:                   _GotoState26Action,
	{_State224, LparenToken}:                      _GotoState30Action,
	{_State224, LbracketToken}:                    _GotoState27Action,
	{_State224, LexErrorToken}:                    _GotoState29Action,
	{_State224, OptionalLabelDeclType}:            _GotoState139Action,
	{_State224, BlockExprType}:                    _GotoState44Action,
	{_State224, CallExprType}:                     _GotoState321Action,
	{_State224, AtomExprType}:                     _GotoState43Action,
	{_State224, LiteralType}:                      _GotoState50Action,
	{_State224, ImplicitStructExprType}:           _GotoState48Action,
	{_State224, AccessExprType}:                   _GotoState319Action,
	{_State224, InitializableTypeType}:            _GotoState49Action,
	{_State224, ExplicitStructDefType}:            _GotoState47Action,
	{_State224, AnonymousFuncExprType}:            _GotoState42Action,
	{_State227, LbracketToken}:                    _GotoState106Action,
	{_State227, DotToken}:                         _GotoState105Action,
	{_State227, QuestionToken}:                    _GotoState107Action,
	{_State227, DollarLbracketToken}:              _GotoState104Action,
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
	{_State227, OptionalGenericBindingType}:       _GotoState108Action,
	{_State228, AssignToken}:                      _GotoState337Action,
	{_State230, CommaToken}:                       _GotoState338Action,
	{_State232, JumpLabelToken}:                   _GotoState339Action,
	{_State232, OptionalJumpLabelType}:            _GotoState340Action,
	{_State235, NewlinesToken}:                    _GotoState341Action,
	{_State235, SemicolonToken}:                   _GotoState342Action,
	{_State239, StringLiteralToken}:               _GotoState344Action,
	{_State239, LparenToken}:                      _GotoState343Action,
	{_State239, ImportClauseType}:                 _GotoState345Action,
	{_State243, NewlinesToken}:                    _GotoState347Action,
	{_State243, CommaToken}:                       _GotoState346Action,
	{_State245, AddToken}:                         _GotoState180Action,
	{_State245, SubToken}:                         _GotoState185Action,
	{_State245, MulToken}:                         _GotoState183Action,
	{_State246, IdentifierToken}:                  _GotoState81Action,
	{_State246, StructToken}:                      _GotoState35Action,
	{_State246, EnumToken}:                        _GotoState78Action,
	{_State246, TraitToken}:                       _GotoState86Action,
	{_State246, FuncToken}:                        _GotoState80Action,
	{_State246, LparenToken}:                      _GotoState83Action,
	{_State246, LbracketToken}:                    _GotoState27Action,
	{_State246, DotToken}:                         _GotoState77Action,
	{_State246, QuestionToken}:                    _GotoState84Action,
	{_State246, ExclaimToken}:                     _GotoState79Action,
	{_State246, TildeTildeToken}:                  _GotoState85Action,
	{_State246, BitNegToken}:                      _GotoState76Action,
	{_State246, BitAndToken}:                      _GotoState75Action,
	{_State246, LexErrorToken}:                    _GotoState82Action,
	{_State246, InitializableTypeType}:            _GotoState92Action,
	{_State246, AtomTypeType}:                     _GotoState87Action,
	{_State246, ReturnableTypeType}:               _GotoState93Action,
	{_State246, ValueTypeType}:                    _GotoState348Action,
	{_State246, ImplicitStructDefType}:            _GotoState91Action,
	{_State246, ExplicitStructDefType}:            _GotoState47Action,
	{_State246, ImplicitEnumDefType}:              _GotoState90Action,
	{_State246, ExplicitEnumDefType}:              _GotoState88Action,
	{_State246, TraitDefType}:                     _GotoState94Action,
	{_State246, FuncTypeType}:                     _GotoState89Action,
	{_State248, CommaToken}:                       _GotoState349Action,
	{_State249, RbracketToken}:                    _GotoState350Action,
	{_State250, ImplementsToken}:                  _GotoState351Action,
	{_State250, AddToken}:                         _GotoState180Action,
	{_State250, SubToken}:                         _GotoState185Action,
	{_State250, MulToken}:                         _GotoState183Action,
	{_State252, IdentifierToken}:                  _GotoState155Action,
	{_State252, ParameterDefType}:                 _GotoState158Action,
	{_State252, ParameterDefsType}:                _GotoState159Action,
	{_State252, OptionalParameterDefsType}:        _GotoState352Action,
	{_State253, IdentifierToken}:                  _GotoState81Action,
	{_State253, StructToken}:                      _GotoState35Action,
	{_State253, EnumToken}:                        _GotoState78Action,
	{_State253, TraitToken}:                       _GotoState86Action,
	{_State253, FuncToken}:                        _GotoState80Action,
	{_State253, LparenToken}:                      _GotoState83Action,
	{_State253, LbracketToken}:                    _GotoState27Action,
	{_State253, DotToken}:                         _GotoState77Action,
	{_State253, QuestionToken}:                    _GotoState84Action,
	{_State253, ExclaimToken}:                     _GotoState79Action,
	{_State253, TildeTildeToken}:                  _GotoState85Action,
	{_State253, BitNegToken}:                      _GotoState76Action,
	{_State253, BitAndToken}:                      _GotoState75Action,
	{_State253, LexErrorToken}:                    _GotoState82Action,
	{_State253, InitializableTypeType}:            _GotoState92Action,
	{_State253, AtomTypeType}:                     _GotoState87Action,
	{_State253, ReturnableTypeType}:               _GotoState93Action,
	{_State253, ValueTypeType}:                    _GotoState353Action,
	{_State253, ImplicitStructDefType}:            _GotoState91Action,
	{_State253, ExplicitStructDefType}:            _GotoState47Action,
	{_State253, ImplicitEnumDefType}:              _GotoState90Action,
	{_State253, ExplicitEnumDefType}:              _GotoState88Action,
	{_State253, TraitDefType}:                     _GotoState94Action,
	{_State253, FuncTypeType}:                     _GotoState89Action,
	{_State254, AddToken}:                         _GotoState180Action,
	{_State254, SubToken}:                         _GotoState185Action,
	{_State254, MulToken}:                         _GotoState183Action,
	{_State255, IdentifierToken}:                  _GotoState354Action,
	{_State256, IdentifierToken}:                  _GotoState81Action,
	{_State256, StructToken}:                      _GotoState35Action,
	{_State256, EnumToken}:                        _GotoState78Action,
	{_State256, TraitToken}:                       _GotoState86Action,
	{_State256, FuncToken}:                        _GotoState80Action,
	{_State256, LparenToken}:                      _GotoState83Action,
	{_State256, LbracketToken}:                    _GotoState27Action,
	{_State256, DotToken}:                         _GotoState77Action,
	{_State256, QuestionToken}:                    _GotoState84Action,
	{_State256, ExclaimToken}:                     _GotoState79Action,
	{_State256, TildeTildeToken}:                  _GotoState85Action,
	{_State256, BitNegToken}:                      _GotoState76Action,
	{_State256, BitAndToken}:                      _GotoState75Action,
	{_State256, LexErrorToken}:                    _GotoState82Action,
	{_State256, InitializableTypeType}:            _GotoState92Action,
	{_State256, AtomTypeType}:                     _GotoState87Action,
	{_State256, ReturnableTypeType}:               _GotoState356Action,
	{_State256, ImplicitStructDefType}:            _GotoState91Action,
	{_State256, ExplicitStructDefType}:            _GotoState47Action,
	{_State256, ImplicitEnumDefType}:              _GotoState90Action,
	{_State256, ExplicitEnumDefType}:              _GotoState88Action,
	{_State256, TraitDefType}:                     _GotoState94Action,
	{_State256, ReturnTypeType}:                   _GotoState355Action,
	{_State256, FuncTypeType}:                     _GotoState89Action,
	{_State257, IdentifierToken}:                  _GotoState155Action,
	{_State257, ParameterDefType}:                 _GotoState357Action,
	{_State258, NewlinesToken}:                    _GotoState358Action,
	{_State258, OrToken}:                          _GotoState359Action,
	{_State259, RparenToken}:                      _GotoState360Action,
	{_State260, AssignToken}:                      _GotoState273Action,
	{_State261, NewlinesToken}:                    _GotoState361Action,
	{_State261, OrToken}:                          _GotoState362Action,
	{_State262, IdentifierToken}:                  _GotoState81Action,
	{_State262, StructToken}:                      _GotoState35Action,
	{_State262, EnumToken}:                        _GotoState78Action,
	{_State262, TraitToken}:                       _GotoState86Action,
	{_State262, FuncToken}:                        _GotoState80Action,
	{_State262, LparenToken}:                      _GotoState83Action,
	{_State262, LbracketToken}:                    _GotoState27Action,
	{_State262, DotToken}:                         _GotoState77Action,
	{_State262, QuestionToken}:                    _GotoState84Action,
	{_State262, ExclaimToken}:                     _GotoState79Action,
	{_State262, TildeTildeToken}:                  _GotoState85Action,
	{_State262, BitNegToken}:                      _GotoState76Action,
	{_State262, BitAndToken}:                      _GotoState75Action,
	{_State262, LexErrorToken}:                    _GotoState82Action,
	{_State262, InitializableTypeType}:            _GotoState92Action,
	{_State262, AtomTypeType}:                     _GotoState87Action,
	{_State262, ReturnableTypeType}:               _GotoState93Action,
	{_State262, ValueTypeType}:                    _GotoState363Action,
	{_State262, ImplicitStructDefType}:            _GotoState91Action,
	{_State262, ExplicitStructDefType}:            _GotoState47Action,
	{_State262, ImplicitEnumDefType}:              _GotoState90Action,
	{_State262, ExplicitEnumDefType}:              _GotoState88Action,
	{_State262, TraitDefType}:                     _GotoState94Action,
	{_State262, FuncTypeType}:                     _GotoState89Action,
	{_State263, IdentifierToken}:                  _GotoState81Action,
	{_State263, StructToken}:                      _GotoState35Action,
	{_State263, EnumToken}:                        _GotoState78Action,
	{_State263, TraitToken}:                       _GotoState86Action,
	{_State263, FuncToken}:                        _GotoState80Action,
	{_State263, LparenToken}:                      _GotoState83Action,
	{_State263, LbracketToken}:                    _GotoState27Action,
	{_State263, DotToken}:                         _GotoState269Action,
	{_State263, QuestionToken}:                    _GotoState84Action,
	{_State263, ExclaimToken}:                     _GotoState79Action,
	{_State263, DollarLbracketToken}:              _GotoState104Action,
	{_State263, DotdotdotToken}:                   _GotoState364Action,
	{_State263, TildeTildeToken}:                  _GotoState85Action,
	{_State263, BitNegToken}:                      _GotoState76Action,
	{_State263, BitAndToken}:                      _GotoState75Action,
	{_State263, LexErrorToken}:                    _GotoState82Action,
	{_State263, OptionalGenericBindingType}:       _GotoState167Action,
	{_State263, InitializableTypeType}:            _GotoState92Action,
	{_State263, AtomTypeType}:                     _GotoState87Action,
	{_State263, ReturnableTypeType}:               _GotoState93Action,
	{_State263, ValueTypeType}:                    _GotoState365Action,
	{_State263, ImplicitStructDefType}:            _GotoState91Action,
	{_State263, ExplicitStructDefType}:            _GotoState47Action,
	{_State263, ImplicitEnumDefType}:              _GotoState90Action,
	{_State263, ExplicitEnumDefType}:              _GotoState88Action,
	{_State263, TraitDefType}:                     _GotoState94Action,
	{_State263, FuncTypeType}:                     _GotoState89Action,
	{_State264, RparenToken}:                      _GotoState366Action,
	{_State266, CommaToken}:                       _GotoState367Action,
	{_State267, AddToken}:                         _GotoState180Action,
	{_State267, SubToken}:                         _GotoState185Action,
	{_State267, MulToken}:                         _GotoState183Action,
	{_State268, DollarLbracketToken}:              _GotoState104Action,
	{_State268, OptionalGenericBindingType}:       _GotoState368Action,
	{_State269, IdentifierToken}:                  _GotoState268Action,
	{_State269, DollarLbracketToken}:              _GotoState104Action,
	{_State269, OptionalGenericBindingType}:       _GotoState162Action,
	{_State270, AddToken}:                         _GotoState180Action,
	{_State270, SubToken}:                         _GotoState185Action,
	{_State270, MulToken}:                         _GotoState183Action,
	{_State271, IdentifierToken}:                  _GotoState369Action,
	{_State272, IdentifierToken}:                  _GotoState168Action,
	{_State272, UnsafeToken}:                      _GotoState169Action,
	{_State272, StructToken}:                      _GotoState35Action,
	{_State272, EnumToken}:                        _GotoState78Action,
	{_State272, TraitToken}:                       _GotoState86Action,
	{_State272, FuncToken}:                        _GotoState80Action,
	{_State272, LparenToken}:                      _GotoState83Action,
	{_State272, LbracketToken}:                    _GotoState27Action,
	{_State272, DotToken}:                         _GotoState77Action,
	{_State272, QuestionToken}:                    _GotoState84Action,
	{_State272, ExclaimToken}:                     _GotoState79Action,
	{_State272, TildeTildeToken}:                  _GotoState85Action,
	{_State272, BitNegToken}:                      _GotoState76Action,
	{_State272, BitAndToken}:                      _GotoState75Action,
	{_State272, LexErrorToken}:                    _GotoState82Action,
	{_State272, UnsafeStatementType}:              _GotoState175Action,
	{_State272, InitializableTypeType}:            _GotoState92Action,
	{_State272, AtomTypeType}:                     _GotoState87Action,
	{_State272, ReturnableTypeType}:               _GotoState93Action,
	{_State272, ValueTypeType}:                    _GotoState176Action,
	{_State272, FieldDefType}:                     _GotoState260Action,
	{_State272, ImplicitStructDefType}:            _GotoState91Action,
	{_State272, ExplicitStructDefType}:            _GotoState47Action,
	{_State272, EnumValueDefType}:                 _GotoState370Action,
	{_State272, ImplicitEnumDefType}:              _GotoState90Action,
	{_State272, ExplicitEnumDefType}:              _GotoState88Action,
	{_State272, TraitDefType}:                     _GotoState94Action,
	{_State272, FuncTypeType}:                     _GotoState89Action,
	{_State273, DefaultToken}:                     _GotoState371Action,
	{_State274, IdentifierToken}:                  _GotoState168Action,
	{_State274, UnsafeToken}:                      _GotoState169Action,
	{_State274, StructToken}:                      _GotoState35Action,
	{_State274, EnumToken}:                        _GotoState78Action,
	{_State274, TraitToken}:                       _GotoState86Action,
	{_State274, FuncToken}:                        _GotoState80Action,
	{_State274, LparenToken}:                      _GotoState83Action,
	{_State274, LbracketToken}:                    _GotoState27Action,
	{_State274, DotToken}:                         _GotoState77Action,
	{_State274, QuestionToken}:                    _GotoState84Action,
	{_State274, ExclaimToken}:                     _GotoState79Action,
	{_State274, TildeTildeToken}:                  _GotoState85Action,
	{_State274, BitNegToken}:                      _GotoState76Action,
	{_State274, BitAndToken}:                      _GotoState75Action,
	{_State274, LexErrorToken}:                    _GotoState82Action,
	{_State274, UnsafeStatementType}:              _GotoState175Action,
	{_State274, InitializableTypeType}:            _GotoState92Action,
	{_State274, AtomTypeType}:                     _GotoState87Action,
	{_State274, ReturnableTypeType}:               _GotoState93Action,
	{_State274, ValueTypeType}:                    _GotoState176Action,
	{_State274, FieldDefType}:                     _GotoState260Action,
	{_State274, ImplicitStructDefType}:            _GotoState91Action,
	{_State274, ExplicitStructDefType}:            _GotoState47Action,
	{_State274, EnumValueDefType}:                 _GotoState372Action,
	{_State274, ImplicitEnumDefType}:              _GotoState90Action,
	{_State274, ExplicitEnumDefType}:              _GotoState88Action,
	{_State274, TraitDefType}:                     _GotoState94Action,
	{_State274, FuncTypeType}:                     _GotoState89Action,
	{_State276, IdentifierToken}:                  _GotoState168Action,
	{_State276, UnsafeToken}:                      _GotoState169Action,
	{_State276, StructToken}:                      _GotoState35Action,
	{_State276, EnumToken}:                        _GotoState78Action,
	{_State276, TraitToken}:                       _GotoState86Action,
	{_State276, FuncToken}:                        _GotoState80Action,
	{_State276, LparenToken}:                      _GotoState83Action,
	{_State276, LbracketToken}:                    _GotoState27Action,
	{_State276, DotToken}:                         _GotoState77Action,
	{_State276, QuestionToken}:                    _GotoState84Action,
	{_State276, ExclaimToken}:                     _GotoState79Action,
	{_State276, TildeTildeToken}:                  _GotoState85Action,
	{_State276, BitNegToken}:                      _GotoState76Action,
	{_State276, BitAndToken}:                      _GotoState75Action,
	{_State276, LexErrorToken}:                    _GotoState82Action,
	{_State276, UnsafeStatementType}:              _GotoState175Action,
	{_State276, InitializableTypeType}:            _GotoState92Action,
	{_State276, AtomTypeType}:                     _GotoState87Action,
	{_State276, ReturnableTypeType}:               _GotoState93Action,
	{_State276, ValueTypeType}:                    _GotoState176Action,
	{_State276, FieldDefType}:                     _GotoState373Action,
	{_State276, ImplicitStructDefType}:            _GotoState91Action,
	{_State276, ExplicitStructDefType}:            _GotoState47Action,
	{_State276, ImplicitEnumDefType}:              _GotoState90Action,
	{_State276, ExplicitEnumDefType}:              _GotoState88Action,
	{_State276, TraitDefType}:                     _GotoState94Action,
	{_State276, FuncTypeType}:                     _GotoState89Action,
	{_State278, IdentifierToken}:                  _GotoState374Action,
	{_State278, LparenToken}:                      _GotoState165Action,
	{_State281, RparenToken}:                      _GotoState375Action,
	{_State282, NewlinesToken}:                    _GotoState377Action,
	{_State282, CommaToken}:                       _GotoState376Action,
	{_State285, RbracketToken}:                    _GotoState378Action,
	{_State285, AddToken}:                         _GotoState180Action,
	{_State285, SubToken}:                         _GotoState185Action,
	{_State285, MulToken}:                         _GotoState183Action,
	{_State286, RbracketToken}:                    _GotoState379Action,
	{_State294, IdentifierToken}:                  _GotoState168Action,
	{_State294, UnsafeToken}:                      _GotoState169Action,
	{_State294, StructToken}:                      _GotoState35Action,
	{_State294, EnumToken}:                        _GotoState78Action,
	{_State294, TraitToken}:                       _GotoState86Action,
	{_State294, FuncToken}:                        _GotoState80Action,
	{_State294, LparenToken}:                      _GotoState83Action,
	{_State294, LbracketToken}:                    _GotoState27Action,
	{_State294, DotToken}:                         _GotoState77Action,
	{_State294, QuestionToken}:                    _GotoState84Action,
	{_State294, ExclaimToken}:                     _GotoState79Action,
	{_State294, TildeTildeToken}:                  _GotoState85Action,
	{_State294, BitNegToken}:                      _GotoState76Action,
	{_State294, BitAndToken}:                      _GotoState75Action,
	{_State294, LexErrorToken}:                    _GotoState82Action,
	{_State294, UnsafeStatementType}:              _GotoState175Action,
	{_State294, InitializableTypeType}:            _GotoState92Action,
	{_State294, AtomTypeType}:                     _GotoState87Action,
	{_State294, ReturnableTypeType}:               _GotoState93Action,
	{_State294, ValueTypeType}:                    _GotoState176Action,
	{_State294, FieldDefType}:                     _GotoState380Action,
	{_State294, ImplicitStructDefType}:            _GotoState91Action,
	{_State294, ExplicitStructDefType}:            _GotoState47Action,
	{_State294, ImplicitEnumDefType}:              _GotoState90Action,
	{_State294, ExplicitEnumDefType}:              _GotoState88Action,
	{_State294, TraitDefType}:                     _GotoState94Action,
	{_State294, FuncTypeType}:                     _GotoState89Action,
	{_State295, IdentifierToken}:                  _GotoState168Action,
	{_State295, UnsafeToken}:                      _GotoState169Action,
	{_State295, StructToken}:                      _GotoState35Action,
	{_State295, EnumToken}:                        _GotoState78Action,
	{_State295, TraitToken}:                       _GotoState86Action,
	{_State295, FuncToken}:                        _GotoState80Action,
	{_State295, LparenToken}:                      _GotoState83Action,
	{_State295, LbracketToken}:                    _GotoState27Action,
	{_State295, DotToken}:                         _GotoState77Action,
	{_State295, QuestionToken}:                    _GotoState84Action,
	{_State295, ExclaimToken}:                     _GotoState79Action,
	{_State295, TildeTildeToken}:                  _GotoState85Action,
	{_State295, BitNegToken}:                      _GotoState76Action,
	{_State295, BitAndToken}:                      _GotoState75Action,
	{_State295, LexErrorToken}:                    _GotoState82Action,
	{_State295, UnsafeStatementType}:              _GotoState175Action,
	{_State295, InitializableTypeType}:            _GotoState92Action,
	{_State295, AtomTypeType}:                     _GotoState87Action,
	{_State295, ReturnableTypeType}:               _GotoState93Action,
	{_State295, ValueTypeType}:                    _GotoState176Action,
	{_State295, FieldDefType}:                     _GotoState381Action,
	{_State295, ImplicitStructDefType}:            _GotoState91Action,
	{_State295, ExplicitStructDefType}:            _GotoState47Action,
	{_State295, ImplicitEnumDefType}:              _GotoState90Action,
	{_State295, ExplicitEnumDefType}:              _GotoState88Action,
	{_State295, TraitDefType}:                     _GotoState94Action,
	{_State295, FuncTypeType}:                     _GotoState89Action,
	{_State297, IdentifierToken}:                  _GotoState81Action,
	{_State297, StructToken}:                      _GotoState35Action,
	{_State297, EnumToken}:                        _GotoState78Action,
	{_State297, TraitToken}:                       _GotoState86Action,
	{_State297, FuncToken}:                        _GotoState80Action,
	{_State297, LparenToken}:                      _GotoState83Action,
	{_State297, LbracketToken}:                    _GotoState27Action,
	{_State297, DotToken}:                         _GotoState77Action,
	{_State297, QuestionToken}:                    _GotoState84Action,
	{_State297, ExclaimToken}:                     _GotoState79Action,
	{_State297, TildeTildeToken}:                  _GotoState85Action,
	{_State297, BitNegToken}:                      _GotoState76Action,
	{_State297, BitAndToken}:                      _GotoState75Action,
	{_State297, LexErrorToken}:                    _GotoState82Action,
	{_State297, InitializableTypeType}:            _GotoState92Action,
	{_State297, AtomTypeType}:                     _GotoState87Action,
	{_State297, ReturnableTypeType}:               _GotoState93Action,
	{_State297, ValueTypeType}:                    _GotoState382Action,
	{_State297, ImplicitStructDefType}:            _GotoState91Action,
	{_State297, ExplicitStructDefType}:            _GotoState47Action,
	{_State297, ImplicitEnumDefType}:              _GotoState90Action,
	{_State297, ExplicitEnumDefType}:              _GotoState88Action,
	{_State297, TraitDefType}:                     _GotoState94Action,
	{_State297, FuncTypeType}:                     _GotoState89Action,
	{_State300, CommaToken}:                       _GotoState187Action,
	{_State301, RparenToken}:                      _GotoState383Action,
	{_State303, IntegerLiteralToken}:              _GotoState25Action,
	{_State303, FloatLiteralToken}:                _GotoState22Action,
	{_State303, RuneLiteralToken}:                 _GotoState33Action,
	{_State303, StringLiteralToken}:               _GotoState34Action,
	{_State303, IdentifierToken}:                  _GotoState24Action,
	{_State303, TrueToken}:                        _GotoState37Action,
	{_State303, FalseToken}:                       _GotoState21Action,
	{_State303, StructToken}:                      _GotoState35Action,
	{_State303, FuncToken}:                        _GotoState23Action,
	{_State303, VarToken}:                         _GotoState38Action,
	{_State303, LetToken}:                         _GotoState28Action,
	{_State303, LabelDeclToken}:                   _GotoState26Action,
	{_State303, LparenToken}:                      _GotoState30Action,
	{_State303, LbracketToken}:                    _GotoState27Action,
	{_State303, NotToken}:                         _GotoState32Action,
	{_State303, SubToken}:                         _GotoState36Action,
	{_State303, MulToken}:                         _GotoState31Action,
	{_State303, BitNegToken}:                      _GotoState20Action,
	{_State303, BitAndToken}:                      _GotoState19Action,
	{_State303, LexErrorToken}:                    _GotoState29Action,
	{_State303, VarDeclPatternType}:               _GotoState58Action,
	{_State303, VarOrLetType}:                     _GotoState59Action,
	{_State303, OptionalLabelDeclType}:            _GotoState139Action,
	{_State303, SequenceExprType}:                 _GotoState384Action,
	{_State303, BlockExprType}:                    _GotoState44Action,
	{_State303, CallExprType}:                     _GotoState45Action,
	{_State303, AtomExprType}:                     _GotoState43Action,
	{_State303, LiteralType}:                      _GotoState50Action,
	{_State303, ImplicitStructExprType}:           _GotoState48Action,
	{_State303, AccessExprType}:                   _GotoState39Action,
	{_State303, PostfixUnaryExprType}:             _GotoState54Action,
	{_State303, PrefixUnaryOpType}:                _GotoState56Action,
	{_State303, PrefixUnaryExprType}:              _GotoState55Action,
	{_State303, MulExprType}:                      _GotoState51Action,
	{_State303, AddExprType}:                      _GotoState40Action,
	{_State303, CmpExprType}:                      _GotoState46Action,
	{_State303, AndExprType}:                      _GotoState41Action,
	{_State303, OrExprType}:                       _GotoState53Action,
	{_State303, InitializableTypeType}:            _GotoState49Action,
	{_State303, ExplicitStructDefType}:            _GotoState47Action,
	{_State303, AnonymousFuncExprType}:            _GotoState42Action,
	{_State304, IntegerLiteralToken}:              _GotoState25Action,
	{_State304, FloatLiteralToken}:                _GotoState22Action,
	{_State304, RuneLiteralToken}:                 _GotoState33Action,
	{_State304, StringLiteralToken}:               _GotoState34Action,
	{_State304, IdentifierToken}:                  _GotoState24Action,
	{_State304, TrueToken}:                        _GotoState37Action,
	{_State304, FalseToken}:                       _GotoState21Action,
	{_State304, StructToken}:                      _GotoState35Action,
	{_State304, FuncToken}:                        _GotoState23Action,
	{_State304, VarToken}:                         _GotoState38Action,
	{_State304, LetToken}:                         _GotoState28Action,
	{_State304, LabelDeclToken}:                   _GotoState26Action,
	{_State304, LparenToken}:                      _GotoState30Action,
	{_State304, LbracketToken}:                    _GotoState27Action,
	{_State304, NotToken}:                         _GotoState32Action,
	{_State304, SubToken}:                         _GotoState36Action,
	{_State304, MulToken}:                         _GotoState31Action,
	{_State304, BitNegToken}:                      _GotoState20Action,
	{_State304, BitAndToken}:                      _GotoState19Action,
	{_State304, LexErrorToken}:                    _GotoState29Action,
	{_State304, VarDeclPatternType}:               _GotoState58Action,
	{_State304, VarOrLetType}:                     _GotoState59Action,
	{_State304, OptionalLabelDeclType}:            _GotoState139Action,
	{_State304, SequenceExprType}:                 _GotoState385Action,
	{_State304, BlockExprType}:                    _GotoState44Action,
	{_State304, CallExprType}:                     _GotoState45Action,
	{_State304, AtomExprType}:                     _GotoState43Action,
	{_State304, LiteralType}:                      _GotoState50Action,
	{_State304, ImplicitStructExprType}:           _GotoState48Action,
	{_State304, AccessExprType}:                   _GotoState39Action,
	{_State304, PostfixUnaryExprType}:             _GotoState54Action,
	{_State304, PrefixUnaryOpType}:                _GotoState56Action,
	{_State304, PrefixUnaryExprType}:              _GotoState55Action,
	{_State304, MulExprType}:                      _GotoState51Action,
	{_State304, AddExprType}:                      _GotoState40Action,
	{_State304, CmpExprType}:                      _GotoState46Action,
	{_State304, AndExprType}:                      _GotoState41Action,
	{_State304, OrExprType}:                       _GotoState53Action,
	{_State304, InitializableTypeType}:            _GotoState49Action,
	{_State304, ExplicitStructDefType}:            _GotoState47Action,
	{_State304, AnonymousFuncExprType}:            _GotoState42Action,
	{_State305, IntegerLiteralToken}:              _GotoState25Action,
	{_State305, FloatLiteralToken}:                _GotoState22Action,
	{_State305, RuneLiteralToken}:                 _GotoState33Action,
	{_State305, StringLiteralToken}:               _GotoState34Action,
	{_State305, IdentifierToken}:                  _GotoState24Action,
	{_State305, TrueToken}:                        _GotoState37Action,
	{_State305, FalseToken}:                       _GotoState21Action,
	{_State305, StructToken}:                      _GotoState35Action,
	{_State305, FuncToken}:                        _GotoState23Action,
	{_State305, VarToken}:                         _GotoState38Action,
	{_State305, LetToken}:                         _GotoState28Action,
	{_State305, LabelDeclToken}:                   _GotoState26Action,
	{_State305, LparenToken}:                      _GotoState30Action,
	{_State305, LbracketToken}:                    _GotoState27Action,
	{_State305, NotToken}:                         _GotoState32Action,
	{_State305, SubToken}:                         _GotoState36Action,
	{_State305, MulToken}:                         _GotoState31Action,
	{_State305, BitNegToken}:                      _GotoState20Action,
	{_State305, BitAndToken}:                      _GotoState19Action,
	{_State305, LexErrorToken}:                    _GotoState29Action,
	{_State305, VarDeclPatternType}:               _GotoState58Action,
	{_State305, VarOrLetType}:                     _GotoState59Action,
	{_State305, OptionalLabelDeclType}:            _GotoState139Action,
	{_State305, SequenceExprType}:                 _GotoState386Action,
	{_State305, BlockExprType}:                    _GotoState44Action,
	{_State305, CallExprType}:                     _GotoState45Action,
	{_State305, AtomExprType}:                     _GotoState43Action,
	{_State305, LiteralType}:                      _GotoState50Action,
	{_State305, ImplicitStructExprType}:           _GotoState48Action,
	{_State305, AccessExprType}:                   _GotoState39Action,
	{_State305, PostfixUnaryExprType}:             _GotoState54Action,
	{_State305, PrefixUnaryOpType}:                _GotoState56Action,
	{_State305, PrefixUnaryExprType}:              _GotoState55Action,
	{_State305, MulExprType}:                      _GotoState51Action,
	{_State305, AddExprType}:                      _GotoState40Action,
	{_State305, CmpExprType}:                      _GotoState46Action,
	{_State305, AndExprType}:                      _GotoState41Action,
	{_State305, OrExprType}:                       _GotoState53Action,
	{_State305, InitializableTypeType}:            _GotoState49Action,
	{_State305, ExplicitStructDefType}:            _GotoState47Action,
	{_State305, AnonymousFuncExprType}:            _GotoState42Action,
	{_State306, IntegerLiteralToken}:              _GotoState25Action,
	{_State306, FloatLiteralToken}:                _GotoState22Action,
	{_State306, RuneLiteralToken}:                 _GotoState33Action,
	{_State306, StringLiteralToken}:               _GotoState34Action,
	{_State306, IdentifierToken}:                  _GotoState24Action,
	{_State306, TrueToken}:                        _GotoState37Action,
	{_State306, FalseToken}:                       _GotoState21Action,
	{_State306, StructToken}:                      _GotoState35Action,
	{_State306, FuncToken}:                        _GotoState23Action,
	{_State306, VarToken}:                         _GotoState38Action,
	{_State306, LetToken}:                         _GotoState28Action,
	{_State306, LabelDeclToken}:                   _GotoState26Action,
	{_State306, LparenToken}:                      _GotoState30Action,
	{_State306, LbracketToken}:                    _GotoState27Action,
	{_State306, NotToken}:                         _GotoState32Action,
	{_State306, SubToken}:                         _GotoState36Action,
	{_State306, MulToken}:                         _GotoState31Action,
	{_State306, BitNegToken}:                      _GotoState20Action,
	{_State306, BitAndToken}:                      _GotoState19Action,
	{_State306, LexErrorToken}:                    _GotoState29Action,
	{_State306, VarDeclPatternType}:               _GotoState58Action,
	{_State306, VarOrLetType}:                     _GotoState59Action,
	{_State306, OptionalLabelDeclType}:            _GotoState139Action,
	{_State306, OptionalSequenceExprType}:         _GotoState387Action,
	{_State306, SequenceExprType}:                 _GotoState388Action,
	{_State306, BlockExprType}:                    _GotoState44Action,
	{_State306, CallExprType}:                     _GotoState45Action,
	{_State306, AtomExprType}:                     _GotoState43Action,
	{_State306, LiteralType}:                      _GotoState50Action,
	{_State306, ImplicitStructExprType}:           _GotoState48Action,
	{_State306, AccessExprType}:                   _GotoState39Action,
	{_State306, PostfixUnaryExprType}:             _GotoState54Action,
	{_State306, PrefixUnaryOpType}:                _GotoState56Action,
	{_State306, PrefixUnaryExprType}:              _GotoState55Action,
	{_State306, MulExprType}:                      _GotoState51Action,
	{_State306, AddExprType}:                      _GotoState40Action,
	{_State306, CmpExprType}:                      _GotoState46Action,
	{_State306, AndExprType}:                      _GotoState41Action,
	{_State306, OrExprType}:                       _GotoState53Action,
	{_State306, InitializableTypeType}:            _GotoState49Action,
	{_State306, ExplicitStructDefType}:            _GotoState47Action,
	{_State306, AnonymousFuncExprType}:            _GotoState42Action,
	{_State307, LbraceToken}:                      _GotoState62Action,
	{_State307, BlockBodyType}:                    _GotoState389Action,
	{_State308, IdentifierToken}:                  _GotoState390Action,
	{_State309, DotToken}:                         _GotoState391Action,
	{_State312, CommaToken}:                       _GotoState393Action,
	{_State312, AssignToken}:                      _GotoState392Action,
	{_State314, ElseToken}:                        _GotoState394Action,
	{_State315, CaseToken}:                        _GotoState395Action,
	{_State315, CaseBranchesType}:                 _GotoState397Action,
	{_State315, CaseBranchType}:                   _GotoState396Action,
	{_State316, IdentifierToken}:                  _GotoState141Action,
	{_State316, LparenToken}:                      _GotoState142Action,
	{_State316, VarPatternType}:                   _GotoState398Action,
	{_State316, TuplePatternType}:                 _GotoState143Action,
	{_State317, IdentifierToken}:                  _GotoState215Action,
	{_State317, LparenToken}:                      _GotoState142Action,
	{_State317, DotdotdotToken}:                   _GotoState214Action,
	{_State317, VarPatternType}:                   _GotoState218Action,
	{_State317, TuplePatternType}:                 _GotoState143Action,
	{_State317, FieldVarPatternType}:              _GotoState399Action,
	{_State319, LbracketToken}:                    _GotoState106Action,
	{_State319, DotToken}:                         _GotoState105Action,
	{_State319, DollarLbracketToken}:              _GotoState104Action,
	{_State319, OptionalGenericBindingType}:       _GotoState108Action,
	{_State335, IntegerLiteralToken}:              _GotoState25Action,
	{_State335, FloatLiteralToken}:                _GotoState22Action,
	{_State335, RuneLiteralToken}:                 _GotoState33Action,
	{_State335, StringLiteralToken}:               _GotoState34Action,
	{_State335, IdentifierToken}:                  _GotoState24Action,
	{_State335, TrueToken}:                        _GotoState37Action,
	{_State335, FalseToken}:                       _GotoState21Action,
	{_State335, StructToken}:                      _GotoState35Action,
	{_State335, FuncToken}:                        _GotoState23Action,
	{_State335, VarToken}:                         _GotoState38Action,
	{_State335, LetToken}:                         _GotoState28Action,
	{_State335, LabelDeclToken}:                   _GotoState26Action,
	{_State335, LparenToken}:                      _GotoState30Action,
	{_State335, LbracketToken}:                    _GotoState27Action,
	{_State335, NotToken}:                         _GotoState32Action,
	{_State335, SubToken}:                         _GotoState36Action,
	{_State335, MulToken}:                         _GotoState31Action,
	{_State335, BitNegToken}:                      _GotoState20Action,
	{_State335, BitAndToken}:                      _GotoState19Action,
	{_State335, LexErrorToken}:                    _GotoState29Action,
	{_State335, VarDeclPatternType}:               _GotoState58Action,
	{_State335, VarOrLetType}:                     _GotoState59Action,
	{_State335, ExpressionType}:                   _GotoState400Action,
	{_State335, OptionalLabelDeclType}:            _GotoState52Action,
	{_State335, SequenceExprType}:                 _GotoState57Action,
	{_State335, BlockExprType}:                    _GotoState44Action,
	{_State335, CallExprType}:                     _GotoState45Action,
	{_State335, AtomExprType}:                     _GotoState43Action,
	{_State335, LiteralType}:                      _GotoState50Action,
	{_State335, ImplicitStructExprType}:           _GotoState48Action,
	{_State335, AccessExprType}:                   _GotoState39Action,
	{_State335, PostfixUnaryExprType}:             _GotoState54Action,
	{_State335, PrefixUnaryOpType}:                _GotoState56Action,
	{_State335, PrefixUnaryExprType}:              _GotoState55Action,
	{_State335, MulExprType}:                      _GotoState51Action,
	{_State335, AddExprType}:                      _GotoState40Action,
	{_State335, CmpExprType}:                      _GotoState46Action,
	{_State335, AndExprType}:                      _GotoState41Action,
	{_State335, OrExprType}:                       _GotoState53Action,
	{_State335, InitializableTypeType}:            _GotoState49Action,
	{_State335, ExplicitStructDefType}:            _GotoState47Action,
	{_State335, AnonymousFuncExprType}:            _GotoState42Action,
	{_State337, IntegerLiteralToken}:              _GotoState25Action,
	{_State337, FloatLiteralToken}:                _GotoState22Action,
	{_State337, RuneLiteralToken}:                 _GotoState33Action,
	{_State337, StringLiteralToken}:               _GotoState34Action,
	{_State337, IdentifierToken}:                  _GotoState24Action,
	{_State337, TrueToken}:                        _GotoState37Action,
	{_State337, FalseToken}:                       _GotoState21Action,
	{_State337, StructToken}:                      _GotoState35Action,
	{_State337, FuncToken}:                        _GotoState23Action,
	{_State337, VarToken}:                         _GotoState38Action,
	{_State337, LetToken}:                         _GotoState28Action,
	{_State337, LabelDeclToken}:                   _GotoState26Action,
	{_State337, LparenToken}:                      _GotoState30Action,
	{_State337, LbracketToken}:                    _GotoState27Action,
	{_State337, NotToken}:                         _GotoState32Action,
	{_State337, SubToken}:                         _GotoState36Action,
	{_State337, MulToken}:                         _GotoState31Action,
	{_State337, BitNegToken}:                      _GotoState20Action,
	{_State337, BitAndToken}:                      _GotoState19Action,
	{_State337, LexErrorToken}:                    _GotoState29Action,
	{_State337, VarDeclPatternType}:               _GotoState58Action,
	{_State337, VarOrLetType}:                     _GotoState59Action,
	{_State337, ExpressionType}:                   _GotoState401Action,
	{_State337, OptionalLabelDeclType}:            _GotoState52Action,
	{_State337, SequenceExprType}:                 _GotoState57Action,
	{_State337, BlockExprType}:                    _GotoState44Action,
	{_State337, CallExprType}:                     _GotoState45Action,
	{_State337, AtomExprType}:                     _GotoState43Action,
	{_State337, LiteralType}:                      _GotoState50Action,
	{_State337, ImplicitStructExprType}:           _GotoState48Action,
	{_State337, AccessExprType}:                   _GotoState39Action,
	{_State337, PostfixUnaryExprType}:             _GotoState54Action,
	{_State337, PrefixUnaryOpType}:                _GotoState56Action,
	{_State337, PrefixUnaryExprType}:              _GotoState55Action,
	{_State337, MulExprType}:                      _GotoState51Action,
	{_State337, AddExprType}:                      _GotoState40Action,
	{_State337, CmpExprType}:                      _GotoState46Action,
	{_State337, AndExprType}:                      _GotoState41Action,
	{_State337, OrExprType}:                       _GotoState53Action,
	{_State337, InitializableTypeType}:            _GotoState49Action,
	{_State337, ExplicitStructDefType}:            _GotoState47Action,
	{_State337, AnonymousFuncExprType}:            _GotoState42Action,
	{_State338, IntegerLiteralToken}:              _GotoState25Action,
	{_State338, FloatLiteralToken}:                _GotoState22Action,
	{_State338, RuneLiteralToken}:                 _GotoState33Action,
	{_State338, StringLiteralToken}:               _GotoState34Action,
	{_State338, IdentifierToken}:                  _GotoState24Action,
	{_State338, TrueToken}:                        _GotoState37Action,
	{_State338, FalseToken}:                       _GotoState21Action,
	{_State338, StructToken}:                      _GotoState35Action,
	{_State338, FuncToken}:                        _GotoState23Action,
	{_State338, VarToken}:                         _GotoState38Action,
	{_State338, LetToken}:                         _GotoState28Action,
	{_State338, LabelDeclToken}:                   _GotoState26Action,
	{_State338, LparenToken}:                      _GotoState30Action,
	{_State338, LbracketToken}:                    _GotoState27Action,
	{_State338, NotToken}:                         _GotoState32Action,
	{_State338, SubToken}:                         _GotoState36Action,
	{_State338, MulToken}:                         _GotoState31Action,
	{_State338, BitNegToken}:                      _GotoState20Action,
	{_State338, BitAndToken}:                      _GotoState19Action,
	{_State338, LexErrorToken}:                    _GotoState29Action,
	{_State338, VarDeclPatternType}:               _GotoState58Action,
	{_State338, VarOrLetType}:                     _GotoState59Action,
	{_State338, ExpressionType}:                   _GotoState402Action,
	{_State338, OptionalLabelDeclType}:            _GotoState52Action,
	{_State338, SequenceExprType}:                 _GotoState57Action,
	{_State338, BlockExprType}:                    _GotoState44Action,
	{_State338, CallExprType}:                     _GotoState45Action,
	{_State338, AtomExprType}:                     _GotoState43Action,
	{_State338, LiteralType}:                      _GotoState50Action,
	{_State338, ImplicitStructExprType}:           _GotoState48Action,
	{_State338, AccessExprType}:                   _GotoState39Action,
	{_State338, PostfixUnaryExprType}:             _GotoState54Action,
	{_State338, PrefixUnaryOpType}:                _GotoState56Action,
	{_State338, PrefixUnaryExprType}:              _GotoState55Action,
	{_State338, MulExprType}:                      _GotoState51Action,
	{_State338, AddExprType}:                      _GotoState40Action,
	{_State338, CmpExprType}:                      _GotoState46Action,
	{_State338, AndExprType}:                      _GotoState41Action,
	{_State338, OrExprType}:                       _GotoState53Action,
	{_State338, InitializableTypeType}:            _GotoState49Action,
	{_State338, ExplicitStructDefType}:            _GotoState47Action,
	{_State338, AnonymousFuncExprType}:            _GotoState42Action,
	{_State340, IntegerLiteralToken}:              _GotoState25Action,
	{_State340, FloatLiteralToken}:                _GotoState22Action,
	{_State340, RuneLiteralToken}:                 _GotoState33Action,
	{_State340, StringLiteralToken}:               _GotoState34Action,
	{_State340, IdentifierToken}:                  _GotoState24Action,
	{_State340, TrueToken}:                        _GotoState37Action,
	{_State340, FalseToken}:                       _GotoState21Action,
	{_State340, StructToken}:                      _GotoState35Action,
	{_State340, FuncToken}:                        _GotoState23Action,
	{_State340, VarToken}:                         _GotoState38Action,
	{_State340, LetToken}:                         _GotoState28Action,
	{_State340, LabelDeclToken}:                   _GotoState26Action,
	{_State340, LparenToken}:                      _GotoState30Action,
	{_State340, LbracketToken}:                    _GotoState27Action,
	{_State340, NotToken}:                         _GotoState32Action,
	{_State340, SubToken}:                         _GotoState36Action,
	{_State340, MulToken}:                         _GotoState31Action,
	{_State340, BitNegToken}:                      _GotoState20Action,
	{_State340, BitAndToken}:                      _GotoState19Action,
	{_State340, LexErrorToken}:                    _GotoState29Action,
	{_State340, VarDeclPatternType}:               _GotoState58Action,
	{_State340, VarOrLetType}:                     _GotoState59Action,
	{_State340, ExpressionType}:                   _GotoState229Action,
	{_State340, OptionalLabelDeclType}:            _GotoState52Action,
	{_State340, SequenceExprType}:                 _GotoState57Action,
	{_State340, BlockExprType}:                    _GotoState44Action,
	{_State340, ExpressionsType}:                  _GotoState403Action,
	{_State340, OptionalExpressionsType}:          _GotoState404Action,
	{_State340, CallExprType}:                     _GotoState45Action,
	{_State340, AtomExprType}:                     _GotoState43Action,
	{_State340, LiteralType}:                      _GotoState50Action,
	{_State340, ImplicitStructExprType}:           _GotoState48Action,
	{_State340, AccessExprType}:                   _GotoState39Action,
	{_State340, PostfixUnaryExprType}:             _GotoState54Action,
	{_State340, PrefixUnaryOpType}:                _GotoState56Action,
	{_State340, PrefixUnaryExprType}:              _GotoState55Action,
	{_State340, MulExprType}:                      _GotoState51Action,
	{_State340, AddExprType}:                      _GotoState40Action,
	{_State340, CmpExprType}:                      _GotoState46Action,
	{_State340, AndExprType}:                      _GotoState41Action,
	{_State340, OrExprType}:                       _GotoState53Action,
	{_State340, InitializableTypeType}:            _GotoState49Action,
	{_State340, ExplicitStructDefType}:            _GotoState47Action,
	{_State340, AnonymousFuncExprType}:            _GotoState42Action,
	{_State343, StringLiteralToken}:               _GotoState344Action,
	{_State343, ImportClauseType}:                 _GotoState405Action,
	{_State343, ImportClauseTerminalType}:         _GotoState406Action,
	{_State343, ImportClausesType}:                _GotoState407Action,
	{_State344, AsToken}:                          _GotoState408Action,
	{_State348, AddToken}:                         _GotoState180Action,
	{_State348, SubToken}:                         _GotoState185Action,
	{_State348, MulToken}:                         _GotoState183Action,
	{_State349, IdentifierToken}:                  _GotoState246Action,
	{_State349, GenericParameterDefType}:          _GotoState409Action,
	{_State351, IdentifierToken}:                  _GotoState81Action,
	{_State351, StructToken}:                      _GotoState35Action,
	{_State351, EnumToken}:                        _GotoState78Action,
	{_State351, TraitToken}:                       _GotoState86Action,
	{_State351, FuncToken}:                        _GotoState80Action,
	{_State351, LparenToken}:                      _GotoState83Action,
	{_State351, LbracketToken}:                    _GotoState27Action,
	{_State351, DotToken}:                         _GotoState77Action,
	{_State351, QuestionToken}:                    _GotoState84Action,
	{_State351, ExclaimToken}:                     _GotoState79Action,
	{_State351, TildeTildeToken}:                  _GotoState85Action,
	{_State351, BitNegToken}:                      _GotoState76Action,
	{_State351, BitAndToken}:                      _GotoState75Action,
	{_State351, LexErrorToken}:                    _GotoState82Action,
	{_State351, InitializableTypeType}:            _GotoState92Action,
	{_State351, AtomTypeType}:                     _GotoState87Action,
	{_State351, ReturnableTypeType}:               _GotoState93Action,
	{_State351, ValueTypeType}:                    _GotoState410Action,
	{_State351, ImplicitStructDefType}:            _GotoState91Action,
	{_State351, ExplicitStructDefType}:            _GotoState47Action,
	{_State351, ImplicitEnumDefType}:              _GotoState90Action,
	{_State351, ExplicitEnumDefType}:              _GotoState88Action,
	{_State351, TraitDefType}:                     _GotoState94Action,
	{_State351, FuncTypeType}:                     _GotoState89Action,
	{_State352, RparenToken}:                      _GotoState411Action,
	{_State353, AddToken}:                         _GotoState180Action,
	{_State353, SubToken}:                         _GotoState185Action,
	{_State353, MulToken}:                         _GotoState183Action,
	{_State354, DollarLbracketToken}:              _GotoState151Action,
	{_State354, OptionalGenericParametersType}:    _GotoState412Action,
	{_State355, LbraceToken}:                      _GotoState62Action,
	{_State355, BlockBodyType}:                    _GotoState413Action,
	{_State358, IdentifierToken}:                  _GotoState168Action,
	{_State358, UnsafeToken}:                      _GotoState169Action,
	{_State358, StructToken}:                      _GotoState35Action,
	{_State358, EnumToken}:                        _GotoState78Action,
	{_State358, TraitToken}:                       _GotoState86Action,
	{_State358, FuncToken}:                        _GotoState80Action,
	{_State358, LparenToken}:                      _GotoState83Action,
	{_State358, LbracketToken}:                    _GotoState27Action,
	{_State358, DotToken}:                         _GotoState77Action,
	{_State358, QuestionToken}:                    _GotoState84Action,
	{_State358, ExclaimToken}:                     _GotoState79Action,
	{_State358, TildeTildeToken}:                  _GotoState85Action,
	{_State358, BitNegToken}:                      _GotoState76Action,
	{_State358, BitAndToken}:                      _GotoState75Action,
	{_State358, LexErrorToken}:                    _GotoState82Action,
	{_State358, UnsafeStatementType}:              _GotoState175Action,
	{_State358, InitializableTypeType}:            _GotoState92Action,
	{_State358, AtomTypeType}:                     _GotoState87Action,
	{_State358, ReturnableTypeType}:               _GotoState93Action,
	{_State358, ValueTypeType}:                    _GotoState176Action,
	{_State358, FieldDefType}:                     _GotoState260Action,
	{_State358, ImplicitStructDefType}:            _GotoState91Action,
	{_State358, ExplicitStructDefType}:            _GotoState47Action,
	{_State358, EnumValueDefType}:                 _GotoState414Action,
	{_State358, ImplicitEnumDefType}:              _GotoState90Action,
	{_State358, ExplicitEnumDefType}:              _GotoState88Action,
	{_State358, TraitDefType}:                     _GotoState94Action,
	{_State358, FuncTypeType}:                     _GotoState89Action,
	{_State359, IdentifierToken}:                  _GotoState168Action,
	{_State359, UnsafeToken}:                      _GotoState169Action,
	{_State359, StructToken}:                      _GotoState35Action,
	{_State359, EnumToken}:                        _GotoState78Action,
	{_State359, TraitToken}:                       _GotoState86Action,
	{_State359, FuncToken}:                        _GotoState80Action,
	{_State359, LparenToken}:                      _GotoState83Action,
	{_State359, LbracketToken}:                    _GotoState27Action,
	{_State359, DotToken}:                         _GotoState77Action,
	{_State359, QuestionToken}:                    _GotoState84Action,
	{_State359, ExclaimToken}:                     _GotoState79Action,
	{_State359, TildeTildeToken}:                  _GotoState85Action,
	{_State359, BitNegToken}:                      _GotoState76Action,
	{_State359, BitAndToken}:                      _GotoState75Action,
	{_State359, LexErrorToken}:                    _GotoState82Action,
	{_State359, UnsafeStatementType}:              _GotoState175Action,
	{_State359, InitializableTypeType}:            _GotoState92Action,
	{_State359, AtomTypeType}:                     _GotoState87Action,
	{_State359, ReturnableTypeType}:               _GotoState93Action,
	{_State359, ValueTypeType}:                    _GotoState176Action,
	{_State359, FieldDefType}:                     _GotoState260Action,
	{_State359, ImplicitStructDefType}:            _GotoState91Action,
	{_State359, ExplicitStructDefType}:            _GotoState47Action,
	{_State359, EnumValueDefType}:                 _GotoState415Action,
	{_State359, ImplicitEnumDefType}:              _GotoState90Action,
	{_State359, ExplicitEnumDefType}:              _GotoState88Action,
	{_State359, TraitDefType}:                     _GotoState94Action,
	{_State359, FuncTypeType}:                     _GotoState89Action,
	{_State361, IdentifierToken}:                  _GotoState168Action,
	{_State361, UnsafeToken}:                      _GotoState169Action,
	{_State361, StructToken}:                      _GotoState35Action,
	{_State361, EnumToken}:                        _GotoState78Action,
	{_State361, TraitToken}:                       _GotoState86Action,
	{_State361, FuncToken}:                        _GotoState80Action,
	{_State361, LparenToken}:                      _GotoState83Action,
	{_State361, LbracketToken}:                    _GotoState27Action,
	{_State361, DotToken}:                         _GotoState77Action,
	{_State361, QuestionToken}:                    _GotoState84Action,
	{_State361, ExclaimToken}:                     _GotoState79Action,
	{_State361, TildeTildeToken}:                  _GotoState85Action,
	{_State361, BitNegToken}:                      _GotoState76Action,
	{_State361, BitAndToken}:                      _GotoState75Action,
	{_State361, LexErrorToken}:                    _GotoState82Action,
	{_State361, UnsafeStatementType}:              _GotoState175Action,
	{_State361, InitializableTypeType}:            _GotoState92Action,
	{_State361, AtomTypeType}:                     _GotoState87Action,
	{_State361, ReturnableTypeType}:               _GotoState93Action,
	{_State361, ValueTypeType}:                    _GotoState176Action,
	{_State361, FieldDefType}:                     _GotoState260Action,
	{_State361, ImplicitStructDefType}:            _GotoState91Action,
	{_State361, ExplicitStructDefType}:            _GotoState47Action,
	{_State361, EnumValueDefType}:                 _GotoState416Action,
	{_State361, ImplicitEnumDefType}:              _GotoState90Action,
	{_State361, ExplicitEnumDefType}:              _GotoState88Action,
	{_State361, TraitDefType}:                     _GotoState94Action,
	{_State361, FuncTypeType}:                     _GotoState89Action,
	{_State362, IdentifierToken}:                  _GotoState168Action,
	{_State362, UnsafeToken}:                      _GotoState169Action,
	{_State362, StructToken}:                      _GotoState35Action,
	{_State362, EnumToken}:                        _GotoState78Action,
	{_State362, TraitToken}:                       _GotoState86Action,
	{_State362, FuncToken}:                        _GotoState80Action,
	{_State362, LparenToken}:                      _GotoState83Action,
	{_State362, LbracketToken}:                    _GotoState27Action,
	{_State362, DotToken}:                         _GotoState77Action,
	{_State362, QuestionToken}:                    _GotoState84Action,
	{_State362, ExclaimToken}:                     _GotoState79Action,
	{_State362, TildeTildeToken}:                  _GotoState85Action,
	{_State362, BitNegToken}:                      _GotoState76Action,
	{_State362, BitAndToken}:                      _GotoState75Action,
	{_State362, LexErrorToken}:                    _GotoState82Action,
	{_State362, UnsafeStatementType}:              _GotoState175Action,
	{_State362, InitializableTypeType}:            _GotoState92Action,
	{_State362, AtomTypeType}:                     _GotoState87Action,
	{_State362, ReturnableTypeType}:               _GotoState93Action,
	{_State362, ValueTypeType}:                    _GotoState176Action,
	{_State362, FieldDefType}:                     _GotoState260Action,
	{_State362, ImplicitStructDefType}:            _GotoState91Action,
	{_State362, ExplicitStructDefType}:            _GotoState47Action,
	{_State362, EnumValueDefType}:                 _GotoState417Action,
	{_State362, ImplicitEnumDefType}:              _GotoState90Action,
	{_State362, ExplicitEnumDefType}:              _GotoState88Action,
	{_State362, TraitDefType}:                     _GotoState94Action,
	{_State362, FuncTypeType}:                     _GotoState89Action,
	{_State363, AddToken}:                         _GotoState180Action,
	{_State363, SubToken}:                         _GotoState185Action,
	{_State363, MulToken}:                         _GotoState183Action,
	{_State364, IdentifierToken}:                  _GotoState81Action,
	{_State364, StructToken}:                      _GotoState35Action,
	{_State364, EnumToken}:                        _GotoState78Action,
	{_State364, TraitToken}:                       _GotoState86Action,
	{_State364, FuncToken}:                        _GotoState80Action,
	{_State364, LparenToken}:                      _GotoState83Action,
	{_State364, LbracketToken}:                    _GotoState27Action,
	{_State364, DotToken}:                         _GotoState77Action,
	{_State364, QuestionToken}:                    _GotoState84Action,
	{_State364, ExclaimToken}:                     _GotoState79Action,
	{_State364, TildeTildeToken}:                  _GotoState85Action,
	{_State364, BitNegToken}:                      _GotoState76Action,
	{_State364, BitAndToken}:                      _GotoState75Action,
	{_State364, LexErrorToken}:                    _GotoState82Action,
	{_State364, InitializableTypeType}:            _GotoState92Action,
	{_State364, AtomTypeType}:                     _GotoState87Action,
	{_State364, ReturnableTypeType}:               _GotoState93Action,
	{_State364, ValueTypeType}:                    _GotoState418Action,
	{_State364, ImplicitStructDefType}:            _GotoState91Action,
	{_State364, ExplicitStructDefType}:            _GotoState47Action,
	{_State364, ImplicitEnumDefType}:              _GotoState90Action,
	{_State364, ExplicitEnumDefType}:              _GotoState88Action,
	{_State364, TraitDefType}:                     _GotoState94Action,
	{_State364, FuncTypeType}:                     _GotoState89Action,
	{_State365, AddToken}:                         _GotoState180Action,
	{_State365, SubToken}:                         _GotoState185Action,
	{_State365, MulToken}:                         _GotoState183Action,
	{_State366, IdentifierToken}:                  _GotoState81Action,
	{_State366, StructToken}:                      _GotoState35Action,
	{_State366, EnumToken}:                        _GotoState78Action,
	{_State366, TraitToken}:                       _GotoState86Action,
	{_State366, FuncToken}:                        _GotoState80Action,
	{_State366, LparenToken}:                      _GotoState83Action,
	{_State366, LbracketToken}:                    _GotoState27Action,
	{_State366, DotToken}:                         _GotoState77Action,
	{_State366, QuestionToken}:                    _GotoState84Action,
	{_State366, ExclaimToken}:                     _GotoState79Action,
	{_State366, TildeTildeToken}:                  _GotoState85Action,
	{_State366, BitNegToken}:                      _GotoState76Action,
	{_State366, BitAndToken}:                      _GotoState75Action,
	{_State366, LexErrorToken}:                    _GotoState82Action,
	{_State366, InitializableTypeType}:            _GotoState92Action,
	{_State366, AtomTypeType}:                     _GotoState87Action,
	{_State366, ReturnableTypeType}:               _GotoState356Action,
	{_State366, ImplicitStructDefType}:            _GotoState91Action,
	{_State366, ExplicitStructDefType}:            _GotoState47Action,
	{_State366, ImplicitEnumDefType}:              _GotoState90Action,
	{_State366, ExplicitEnumDefType}:              _GotoState88Action,
	{_State366, TraitDefType}:                     _GotoState94Action,
	{_State366, ReturnTypeType}:                   _GotoState419Action,
	{_State366, FuncTypeType}:                     _GotoState89Action,
	{_State367, IdentifierToken}:                  _GotoState263Action,
	{_State367, StructToken}:                      _GotoState35Action,
	{_State367, EnumToken}:                        _GotoState78Action,
	{_State367, TraitToken}:                       _GotoState86Action,
	{_State367, FuncToken}:                        _GotoState80Action,
	{_State367, LparenToken}:                      _GotoState83Action,
	{_State367, LbracketToken}:                    _GotoState27Action,
	{_State367, DotToken}:                         _GotoState77Action,
	{_State367, QuestionToken}:                    _GotoState84Action,
	{_State367, ExclaimToken}:                     _GotoState79Action,
	{_State367, DotdotdotToken}:                   _GotoState262Action,
	{_State367, TildeTildeToken}:                  _GotoState85Action,
	{_State367, BitNegToken}:                      _GotoState76Action,
	{_State367, BitAndToken}:                      _GotoState75Action,
	{_State367, LexErrorToken}:                    _GotoState82Action,
	{_State367, InitializableTypeType}:            _GotoState92Action,
	{_State367, AtomTypeType}:                     _GotoState87Action,
	{_State367, ReturnableTypeType}:               _GotoState93Action,
	{_State367, ValueTypeType}:                    _GotoState267Action,
	{_State367, ImplicitStructDefType}:            _GotoState91Action,
	{_State367, ExplicitStructDefType}:            _GotoState47Action,
	{_State367, ImplicitEnumDefType}:              _GotoState90Action,
	{_State367, ExplicitEnumDefType}:              _GotoState88Action,
	{_State367, TraitDefType}:                     _GotoState94Action,
	{_State367, ParameterDeclType}:                _GotoState420Action,
	{_State367, FuncTypeType}:                     _GotoState89Action,
	{_State369, GreaterToken}:                     _GotoState421Action,
	{_State374, LparenToken}:                      _GotoState422Action,
	{_State376, IdentifierToken}:                  _GotoState168Action,
	{_State376, UnsafeToken}:                      _GotoState169Action,
	{_State376, StructToken}:                      _GotoState35Action,
	{_State376, EnumToken}:                        _GotoState78Action,
	{_State376, TraitToken}:                       _GotoState86Action,
	{_State376, FuncToken}:                        _GotoState278Action,
	{_State376, LparenToken}:                      _GotoState83Action,
	{_State376, LbracketToken}:                    _GotoState27Action,
	{_State376, DotToken}:                         _GotoState77Action,
	{_State376, QuestionToken}:                    _GotoState84Action,
	{_State376, ExclaimToken}:                     _GotoState79Action,
	{_State376, TildeTildeToken}:                  _GotoState85Action,
	{_State376, BitNegToken}:                      _GotoState76Action,
	{_State376, BitAndToken}:                      _GotoState75Action,
	{_State376, LexErrorToken}:                    _GotoState82Action,
	{_State376, UnsafeStatementType}:              _GotoState175Action,
	{_State376, InitializableTypeType}:            _GotoState92Action,
	{_State376, AtomTypeType}:                     _GotoState87Action,
	{_State376, ReturnableTypeType}:               _GotoState93Action,
	{_State376, ValueTypeType}:                    _GotoState176Action,
	{_State376, FieldDefType}:                     _GotoState279Action,
	{_State376, ImplicitStructDefType}:            _GotoState91Action,
	{_State376, ExplicitStructDefType}:            _GotoState47Action,
	{_State376, ImplicitEnumDefType}:              _GotoState90Action,
	{_State376, ExplicitEnumDefType}:              _GotoState88Action,
	{_State376, TraitPropertyType}:                _GotoState423Action,
	{_State376, TraitDefType}:                     _GotoState94Action,
	{_State376, FuncTypeType}:                     _GotoState89Action,
	{_State376, MethodSignatureType}:              _GotoState280Action,
	{_State377, IdentifierToken}:                  _GotoState168Action,
	{_State377, UnsafeToken}:                      _GotoState169Action,
	{_State377, StructToken}:                      _GotoState35Action,
	{_State377, EnumToken}:                        _GotoState78Action,
	{_State377, TraitToken}:                       _GotoState86Action,
	{_State377, FuncToken}:                        _GotoState278Action,
	{_State377, LparenToken}:                      _GotoState83Action,
	{_State377, LbracketToken}:                    _GotoState27Action,
	{_State377, DotToken}:                         _GotoState77Action,
	{_State377, QuestionToken}:                    _GotoState84Action,
	{_State377, ExclaimToken}:                     _GotoState79Action,
	{_State377, TildeTildeToken}:                  _GotoState85Action,
	{_State377, BitNegToken}:                      _GotoState76Action,
	{_State377, BitAndToken}:                      _GotoState75Action,
	{_State377, LexErrorToken}:                    _GotoState82Action,
	{_State377, UnsafeStatementType}:              _GotoState175Action,
	{_State377, InitializableTypeType}:            _GotoState92Action,
	{_State377, AtomTypeType}:                     _GotoState87Action,
	{_State377, ReturnableTypeType}:               _GotoState93Action,
	{_State377, ValueTypeType}:                    _GotoState176Action,
	{_State377, FieldDefType}:                     _GotoState279Action,
	{_State377, ImplicitStructDefType}:            _GotoState91Action,
	{_State377, ExplicitStructDefType}:            _GotoState47Action,
	{_State377, ImplicitEnumDefType}:              _GotoState90Action,
	{_State377, ExplicitEnumDefType}:              _GotoState88Action,
	{_State377, TraitPropertyType}:                _GotoState424Action,
	{_State377, TraitDefType}:                     _GotoState94Action,
	{_State377, FuncTypeType}:                     _GotoState89Action,
	{_State377, MethodSignatureType}:              _GotoState280Action,
	{_State382, AddToken}:                         _GotoState180Action,
	{_State382, SubToken}:                         _GotoState185Action,
	{_State382, MulToken}:                         _GotoState183Action,
	{_State386, DoToken}:                          _GotoState425Action,
	{_State387, SemicolonToken}:                   _GotoState426Action,
	{_State390, LparenToken}:                      _GotoState30Action,
	{_State390, ImplicitStructExprType}:           _GotoState427Action,
	{_State391, IdentifierToken}:                  _GotoState428Action,
	{_State392, IntegerLiteralToken}:              _GotoState25Action,
	{_State392, FloatLiteralToken}:                _GotoState22Action,
	{_State392, RuneLiteralToken}:                 _GotoState33Action,
	{_State392, StringLiteralToken}:               _GotoState34Action,
	{_State392, IdentifierToken}:                  _GotoState24Action,
	{_State392, TrueToken}:                        _GotoState37Action,
	{_State392, FalseToken}:                       _GotoState21Action,
	{_State392, StructToken}:                      _GotoState35Action,
	{_State392, FuncToken}:                        _GotoState23Action,
	{_State392, VarToken}:                         _GotoState38Action,
	{_State392, LetToken}:                         _GotoState28Action,
	{_State392, LabelDeclToken}:                   _GotoState26Action,
	{_State392, LparenToken}:                      _GotoState30Action,
	{_State392, LbracketToken}:                    _GotoState27Action,
	{_State392, NotToken}:                         _GotoState32Action,
	{_State392, SubToken}:                         _GotoState36Action,
	{_State392, MulToken}:                         _GotoState31Action,
	{_State392, BitNegToken}:                      _GotoState20Action,
	{_State392, BitAndToken}:                      _GotoState19Action,
	{_State392, LexErrorToken}:                    _GotoState29Action,
	{_State392, VarDeclPatternType}:               _GotoState58Action,
	{_State392, VarOrLetType}:                     _GotoState59Action,
	{_State392, OptionalLabelDeclType}:            _GotoState139Action,
	{_State392, SequenceExprType}:                 _GotoState429Action,
	{_State392, BlockExprType}:                    _GotoState44Action,
	{_State392, CallExprType}:                     _GotoState45Action,
	{_State392, AtomExprType}:                     _GotoState43Action,
	{_State392, LiteralType}:                      _GotoState50Action,
	{_State392, ImplicitStructExprType}:           _GotoState48Action,
	{_State392, AccessExprType}:                   _GotoState39Action,
	{_State392, PostfixUnaryExprType}:             _GotoState54Action,
	{_State392, PrefixUnaryOpType}:                _GotoState56Action,
	{_State392, PrefixUnaryExprType}:              _GotoState55Action,
	{_State392, MulExprType}:                      _GotoState51Action,
	{_State392, AddExprType}:                      _GotoState40Action,
	{_State392, CmpExprType}:                      _GotoState46Action,
	{_State392, AndExprType}:                      _GotoState41Action,
	{_State392, OrExprType}:                       _GotoState53Action,
	{_State392, InitializableTypeType}:            _GotoState49Action,
	{_State392, ExplicitStructDefType}:            _GotoState47Action,
	{_State392, AnonymousFuncExprType}:            _GotoState42Action,
	{_State393, IntegerLiteralToken}:              _GotoState25Action,
	{_State393, FloatLiteralToken}:                _GotoState22Action,
	{_State393, RuneLiteralToken}:                 _GotoState33Action,
	{_State393, StringLiteralToken}:               _GotoState34Action,
	{_State393, IdentifierToken}:                  _GotoState24Action,
	{_State393, TrueToken}:                        _GotoState37Action,
	{_State393, FalseToken}:                       _GotoState21Action,
	{_State393, StructToken}:                      _GotoState35Action,
	{_State393, FuncToken}:                        _GotoState23Action,
	{_State393, VarToken}:                         _GotoState309Action,
	{_State393, LetToken}:                         _GotoState28Action,
	{_State393, LabelDeclToken}:                   _GotoState26Action,
	{_State393, LparenToken}:                      _GotoState30Action,
	{_State393, LbracketToken}:                    _GotoState27Action,
	{_State393, DotToken}:                         _GotoState308Action,
	{_State393, NotToken}:                         _GotoState32Action,
	{_State393, SubToken}:                         _GotoState36Action,
	{_State393, MulToken}:                         _GotoState31Action,
	{_State393, BitNegToken}:                      _GotoState20Action,
	{_State393, BitAndToken}:                      _GotoState19Action,
	{_State393, LexErrorToken}:                    _GotoState29Action,
	{_State393, VarDeclPatternType}:               _GotoState58Action,
	{_State393, VarOrLetType}:                     _GotoState59Action,
	{_State393, AssignPatternType}:                _GotoState310Action,
	{_State393, CasePatternType}:                  _GotoState430Action,
	{_State393, OptionalLabelDeclType}:            _GotoState139Action,
	{_State393, SequenceExprType}:                 _GotoState313Action,
	{_State393, BlockExprType}:                    _GotoState44Action,
	{_State393, CallExprType}:                     _GotoState45Action,
	{_State393, AtomExprType}:                     _GotoState43Action,
	{_State393, LiteralType}:                      _GotoState50Action,
	{_State393, ImplicitStructExprType}:           _GotoState48Action,
	{_State393, AccessExprType}:                   _GotoState39Action,
	{_State393, PostfixUnaryExprType}:             _GotoState54Action,
	{_State393, PrefixUnaryOpType}:                _GotoState56Action,
	{_State393, PrefixUnaryExprType}:              _GotoState55Action,
	{_State393, MulExprType}:                      _GotoState51Action,
	{_State393, AddExprType}:                      _GotoState40Action,
	{_State393, CmpExprType}:                      _GotoState46Action,
	{_State393, AndExprType}:                      _GotoState41Action,
	{_State393, OrExprType}:                       _GotoState53Action,
	{_State393, InitializableTypeType}:            _GotoState49Action,
	{_State393, ExplicitStructDefType}:            _GotoState47Action,
	{_State393, AnonymousFuncExprType}:            _GotoState42Action,
	{_State394, IfToken}:                          _GotoState132Action,
	{_State394, LbraceToken}:                      _GotoState62Action,
	{_State394, IfExprType}:                       _GotoState432Action,
	{_State394, BlockBodyType}:                    _GotoState431Action,
	{_State395, IntegerLiteralToken}:              _GotoState25Action,
	{_State395, FloatLiteralToken}:                _GotoState22Action,
	{_State395, RuneLiteralToken}:                 _GotoState33Action,
	{_State395, StringLiteralToken}:               _GotoState34Action,
	{_State395, IdentifierToken}:                  _GotoState24Action,
	{_State395, TrueToken}:                        _GotoState37Action,
	{_State395, FalseToken}:                       _GotoState21Action,
	{_State395, StructToken}:                      _GotoState35Action,
	{_State395, FuncToken}:                        _GotoState23Action,
	{_State395, VarToken}:                         _GotoState309Action,
	{_State395, LetToken}:                         _GotoState28Action,
	{_State395, LabelDeclToken}:                   _GotoState26Action,
	{_State395, LparenToken}:                      _GotoState30Action,
	{_State395, LbracketToken}:                    _GotoState27Action,
	{_State395, DotToken}:                         _GotoState308Action,
	{_State395, NotToken}:                         _GotoState32Action,
	{_State395, SubToken}:                         _GotoState36Action,
	{_State395, MulToken}:                         _GotoState31Action,
	{_State395, BitNegToken}:                      _GotoState20Action,
	{_State395, BitAndToken}:                      _GotoState19Action,
	{_State395, LexErrorToken}:                    _GotoState29Action,
	{_State395, VarDeclPatternType}:               _GotoState58Action,
	{_State395, VarOrLetType}:                     _GotoState59Action,
	{_State395, AssignPatternType}:                _GotoState310Action,
	{_State395, CasePatternType}:                  _GotoState311Action,
	{_State395, OptionalLabelDeclType}:            _GotoState139Action,
	{_State395, CasePatternsType}:                 _GotoState433Action,
	{_State395, SequenceExprType}:                 _GotoState313Action,
	{_State395, BlockExprType}:                    _GotoState44Action,
	{_State395, CallExprType}:                     _GotoState45Action,
	{_State395, AtomExprType}:                     _GotoState43Action,
	{_State395, LiteralType}:                      _GotoState50Action,
	{_State395, ImplicitStructExprType}:           _GotoState48Action,
	{_State395, AccessExprType}:                   _GotoState39Action,
	{_State395, PostfixUnaryExprType}:             _GotoState54Action,
	{_State395, PrefixUnaryOpType}:                _GotoState56Action,
	{_State395, PrefixUnaryExprType}:              _GotoState55Action,
	{_State395, MulExprType}:                      _GotoState51Action,
	{_State395, AddExprType}:                      _GotoState40Action,
	{_State395, CmpExprType}:                      _GotoState46Action,
	{_State395, AndExprType}:                      _GotoState41Action,
	{_State395, OrExprType}:                       _GotoState53Action,
	{_State395, InitializableTypeType}:            _GotoState49Action,
	{_State395, ExplicitStructDefType}:            _GotoState47Action,
	{_State395, AnonymousFuncExprType}:            _GotoState42Action,
	{_State397, CaseToken}:                        _GotoState395Action,
	{_State397, DefaultToken}:                     _GotoState434Action,
	{_State397, CaseBranchType}:                   _GotoState435Action,
	{_State397, OptionalDefaultBranchType}:        _GotoState437Action,
	{_State397, DefaultBranchType}:                _GotoState436Action,
	{_State403, CommaToken}:                       _GotoState338Action,
	{_State405, NewlinesToken}:                    _GotoState439Action,
	{_State405, CommaToken}:                       _GotoState438Action,
	{_State407, StringLiteralToken}:               _GotoState344Action,
	{_State407, RparenToken}:                      _GotoState440Action,
	{_State407, ImportClauseType}:                 _GotoState405Action,
	{_State407, ImportClauseTerminalType}:         _GotoState441Action,
	{_State408, IdentifierToken}:                  _GotoState442Action,
	{_State410, AddToken}:                         _GotoState180Action,
	{_State410, SubToken}:                         _GotoState185Action,
	{_State410, MulToken}:                         _GotoState183Action,
	{_State411, IdentifierToken}:                  _GotoState81Action,
	{_State411, StructToken}:                      _GotoState35Action,
	{_State411, EnumToken}:                        _GotoState78Action,
	{_State411, TraitToken}:                       _GotoState86Action,
	{_State411, FuncToken}:                        _GotoState80Action,
	{_State411, LparenToken}:                      _GotoState83Action,
	{_State411, LbracketToken}:                    _GotoState27Action,
	{_State411, DotToken}:                         _GotoState77Action,
	{_State411, QuestionToken}:                    _GotoState84Action,
	{_State411, ExclaimToken}:                     _GotoState79Action,
	{_State411, TildeTildeToken}:                  _GotoState85Action,
	{_State411, BitNegToken}:                      _GotoState76Action,
	{_State411, BitAndToken}:                      _GotoState75Action,
	{_State411, LexErrorToken}:                    _GotoState82Action,
	{_State411, InitializableTypeType}:            _GotoState92Action,
	{_State411, AtomTypeType}:                     _GotoState87Action,
	{_State411, ReturnableTypeType}:               _GotoState356Action,
	{_State411, ImplicitStructDefType}:            _GotoState91Action,
	{_State411, ExplicitStructDefType}:            _GotoState47Action,
	{_State411, ImplicitEnumDefType}:              _GotoState90Action,
	{_State411, ExplicitEnumDefType}:              _GotoState88Action,
	{_State411, TraitDefType}:                     _GotoState94Action,
	{_State411, ReturnTypeType}:                   _GotoState443Action,
	{_State411, FuncTypeType}:                     _GotoState89Action,
	{_State412, LparenToken}:                      _GotoState444Action,
	{_State418, AddToken}:                         _GotoState180Action,
	{_State418, SubToken}:                         _GotoState185Action,
	{_State418, MulToken}:                         _GotoState183Action,
	{_State421, StringLiteralToken}:               _GotoState445Action,
	{_State422, IdentifierToken}:                  _GotoState263Action,
	{_State422, StructToken}:                      _GotoState35Action,
	{_State422, EnumToken}:                        _GotoState78Action,
	{_State422, TraitToken}:                       _GotoState86Action,
	{_State422, FuncToken}:                        _GotoState80Action,
	{_State422, LparenToken}:                      _GotoState83Action,
	{_State422, LbracketToken}:                    _GotoState27Action,
	{_State422, DotToken}:                         _GotoState77Action,
	{_State422, QuestionToken}:                    _GotoState84Action,
	{_State422, ExclaimToken}:                     _GotoState79Action,
	{_State422, DotdotdotToken}:                   _GotoState262Action,
	{_State422, TildeTildeToken}:                  _GotoState85Action,
	{_State422, BitNegToken}:                      _GotoState76Action,
	{_State422, BitAndToken}:                      _GotoState75Action,
	{_State422, LexErrorToken}:                    _GotoState82Action,
	{_State422, InitializableTypeType}:            _GotoState92Action,
	{_State422, AtomTypeType}:                     _GotoState87Action,
	{_State422, ReturnableTypeType}:               _GotoState93Action,
	{_State422, ValueTypeType}:                    _GotoState267Action,
	{_State422, ImplicitStructDefType}:            _GotoState91Action,
	{_State422, ExplicitStructDefType}:            _GotoState47Action,
	{_State422, ImplicitEnumDefType}:              _GotoState90Action,
	{_State422, ExplicitEnumDefType}:              _GotoState88Action,
	{_State422, TraitDefType}:                     _GotoState94Action,
	{_State422, ParameterDeclType}:                _GotoState265Action,
	{_State422, ParameterDeclsType}:               _GotoState266Action,
	{_State422, OptionalParameterDeclsType}:       _GotoState446Action,
	{_State422, FuncTypeType}:                     _GotoState89Action,
	{_State425, LbraceToken}:                      _GotoState62Action,
	{_State425, BlockBodyType}:                    _GotoState447Action,
	{_State426, IntegerLiteralToken}:              _GotoState25Action,
	{_State426, FloatLiteralToken}:                _GotoState22Action,
	{_State426, RuneLiteralToken}:                 _GotoState33Action,
	{_State426, StringLiteralToken}:               _GotoState34Action,
	{_State426, IdentifierToken}:                  _GotoState24Action,
	{_State426, TrueToken}:                        _GotoState37Action,
	{_State426, FalseToken}:                       _GotoState21Action,
	{_State426, StructToken}:                      _GotoState35Action,
	{_State426, FuncToken}:                        _GotoState23Action,
	{_State426, VarToken}:                         _GotoState38Action,
	{_State426, LetToken}:                         _GotoState28Action,
	{_State426, LabelDeclToken}:                   _GotoState26Action,
	{_State426, LparenToken}:                      _GotoState30Action,
	{_State426, LbracketToken}:                    _GotoState27Action,
	{_State426, NotToken}:                         _GotoState32Action,
	{_State426, SubToken}:                         _GotoState36Action,
	{_State426, MulToken}:                         _GotoState31Action,
	{_State426, BitNegToken}:                      _GotoState20Action,
	{_State426, BitAndToken}:                      _GotoState19Action,
	{_State426, LexErrorToken}:                    _GotoState29Action,
	{_State426, VarDeclPatternType}:               _GotoState58Action,
	{_State426, VarOrLetType}:                     _GotoState59Action,
	{_State426, OptionalLabelDeclType}:            _GotoState139Action,
	{_State426, OptionalSequenceExprType}:         _GotoState448Action,
	{_State426, SequenceExprType}:                 _GotoState388Action,
	{_State426, BlockExprType}:                    _GotoState44Action,
	{_State426, CallExprType}:                     _GotoState45Action,
	{_State426, AtomExprType}:                     _GotoState43Action,
	{_State426, LiteralType}:                      _GotoState50Action,
	{_State426, ImplicitStructExprType}:           _GotoState48Action,
	{_State426, AccessExprType}:                   _GotoState39Action,
	{_State426, PostfixUnaryExprType}:             _GotoState54Action,
	{_State426, PrefixUnaryOpType}:                _GotoState56Action,
	{_State426, PrefixUnaryExprType}:              _GotoState55Action,
	{_State426, MulExprType}:                      _GotoState51Action,
	{_State426, AddExprType}:                      _GotoState40Action,
	{_State426, CmpExprType}:                      _GotoState46Action,
	{_State426, AndExprType}:                      _GotoState41Action,
	{_State426, OrExprType}:                       _GotoState53Action,
	{_State426, InitializableTypeType}:            _GotoState49Action,
	{_State426, ExplicitStructDefType}:            _GotoState47Action,
	{_State426, AnonymousFuncExprType}:            _GotoState42Action,
	{_State428, LparenToken}:                      _GotoState142Action,
	{_State428, TuplePatternType}:                 _GotoState449Action,
	{_State433, CommaToken}:                       _GotoState393Action,
	{_State433, ColonToken}:                       _GotoState450Action,
	{_State434, ColonToken}:                       _GotoState451Action,
	{_State437, RbraceToken}:                      _GotoState452Action,
	{_State443, LbraceToken}:                      _GotoState62Action,
	{_State443, BlockBodyType}:                    _GotoState453Action,
	{_State444, IdentifierToken}:                  _GotoState155Action,
	{_State444, ParameterDefType}:                 _GotoState158Action,
	{_State444, ParameterDefsType}:                _GotoState159Action,
	{_State444, OptionalParameterDefsType}:        _GotoState454Action,
	{_State446, RparenToken}:                      _GotoState455Action,
	{_State448, DoToken}:                          _GotoState456Action,
	{_State450, IntegerLiteralToken}:              _GotoState25Action,
	{_State450, FloatLiteralToken}:                _GotoState22Action,
	{_State450, RuneLiteralToken}:                 _GotoState33Action,
	{_State450, StringLiteralToken}:               _GotoState34Action,
	{_State450, IdentifierToken}:                  _GotoState24Action,
	{_State450, TrueToken}:                        _GotoState37Action,
	{_State450, FalseToken}:                       _GotoState21Action,
	{_State450, StructToken}:                      _GotoState35Action,
	{_State450, FuncToken}:                        _GotoState23Action,
	{_State450, VarToken}:                         _GotoState38Action,
	{_State450, LetToken}:                         _GotoState28Action,
	{_State450, LabelDeclToken}:                   _GotoState26Action,
	{_State450, LparenToken}:                      _GotoState30Action,
	{_State450, LbracketToken}:                    _GotoState27Action,
	{_State450, NotToken}:                         _GotoState32Action,
	{_State450, SubToken}:                         _GotoState36Action,
	{_State450, MulToken}:                         _GotoState31Action,
	{_State450, BitNegToken}:                      _GotoState20Action,
	{_State450, BitAndToken}:                      _GotoState19Action,
	{_State450, LexErrorToken}:                    _GotoState29Action,
	{_State450, VarDeclPatternType}:               _GotoState58Action,
	{_State450, VarOrLetType}:                     _GotoState59Action,
	{_State450, OptionalLabelDeclType}:            _GotoState139Action,
	{_State450, SequenceExprType}:                 _GotoState457Action,
	{_State450, BlockExprType}:                    _GotoState44Action,
	{_State450, CallExprType}:                     _GotoState45Action,
	{_State450, AtomExprType}:                     _GotoState43Action,
	{_State450, LiteralType}:                      _GotoState50Action,
	{_State450, ImplicitStructExprType}:           _GotoState48Action,
	{_State450, AccessExprType}:                   _GotoState39Action,
	{_State450, PostfixUnaryExprType}:             _GotoState54Action,
	{_State450, PrefixUnaryOpType}:                _GotoState56Action,
	{_State450, PrefixUnaryExprType}:              _GotoState55Action,
	{_State450, MulExprType}:                      _GotoState51Action,
	{_State450, AddExprType}:                      _GotoState40Action,
	{_State450, CmpExprType}:                      _GotoState46Action,
	{_State450, AndExprType}:                      _GotoState41Action,
	{_State450, OrExprType}:                       _GotoState53Action,
	{_State450, InitializableTypeType}:            _GotoState49Action,
	{_State450, ExplicitStructDefType}:            _GotoState47Action,
	{_State450, AnonymousFuncExprType}:            _GotoState42Action,
	{_State451, IntegerLiteralToken}:              _GotoState25Action,
	{_State451, FloatLiteralToken}:                _GotoState22Action,
	{_State451, RuneLiteralToken}:                 _GotoState33Action,
	{_State451, StringLiteralToken}:               _GotoState34Action,
	{_State451, IdentifierToken}:                  _GotoState24Action,
	{_State451, TrueToken}:                        _GotoState37Action,
	{_State451, FalseToken}:                       _GotoState21Action,
	{_State451, StructToken}:                      _GotoState35Action,
	{_State451, FuncToken}:                        _GotoState23Action,
	{_State451, VarToken}:                         _GotoState38Action,
	{_State451, LetToken}:                         _GotoState28Action,
	{_State451, LabelDeclToken}:                   _GotoState26Action,
	{_State451, LparenToken}:                      _GotoState30Action,
	{_State451, LbracketToken}:                    _GotoState27Action,
	{_State451, NotToken}:                         _GotoState32Action,
	{_State451, SubToken}:                         _GotoState36Action,
	{_State451, MulToken}:                         _GotoState31Action,
	{_State451, BitNegToken}:                      _GotoState20Action,
	{_State451, BitAndToken}:                      _GotoState19Action,
	{_State451, LexErrorToken}:                    _GotoState29Action,
	{_State451, VarDeclPatternType}:               _GotoState58Action,
	{_State451, VarOrLetType}:                     _GotoState59Action,
	{_State451, OptionalLabelDeclType}:            _GotoState139Action,
	{_State451, SequenceExprType}:                 _GotoState458Action,
	{_State451, BlockExprType}:                    _GotoState44Action,
	{_State451, CallExprType}:                     _GotoState45Action,
	{_State451, AtomExprType}:                     _GotoState43Action,
	{_State451, LiteralType}:                      _GotoState50Action,
	{_State451, ImplicitStructExprType}:           _GotoState48Action,
	{_State451, AccessExprType}:                   _GotoState39Action,
	{_State451, PostfixUnaryExprType}:             _GotoState54Action,
	{_State451, PrefixUnaryOpType}:                _GotoState56Action,
	{_State451, PrefixUnaryExprType}:              _GotoState55Action,
	{_State451, MulExprType}:                      _GotoState51Action,
	{_State451, AddExprType}:                      _GotoState40Action,
	{_State451, CmpExprType}:                      _GotoState46Action,
	{_State451, AndExprType}:                      _GotoState41Action,
	{_State451, OrExprType}:                       _GotoState53Action,
	{_State451, InitializableTypeType}:            _GotoState49Action,
	{_State451, ExplicitStructDefType}:            _GotoState47Action,
	{_State451, AnonymousFuncExprType}:            _GotoState42Action,
	{_State454, RparenToken}:                      _GotoState459Action,
	{_State455, IdentifierToken}:                  _GotoState81Action,
	{_State455, StructToken}:                      _GotoState35Action,
	{_State455, EnumToken}:                        _GotoState78Action,
	{_State455, TraitToken}:                       _GotoState86Action,
	{_State455, FuncToken}:                        _GotoState80Action,
	{_State455, LparenToken}:                      _GotoState83Action,
	{_State455, LbracketToken}:                    _GotoState27Action,
	{_State455, DotToken}:                         _GotoState77Action,
	{_State455, QuestionToken}:                    _GotoState84Action,
	{_State455, ExclaimToken}:                     _GotoState79Action,
	{_State455, TildeTildeToken}:                  _GotoState85Action,
	{_State455, BitNegToken}:                      _GotoState76Action,
	{_State455, BitAndToken}:                      _GotoState75Action,
	{_State455, LexErrorToken}:                    _GotoState82Action,
	{_State455, InitializableTypeType}:            _GotoState92Action,
	{_State455, AtomTypeType}:                     _GotoState87Action,
	{_State455, ReturnableTypeType}:               _GotoState356Action,
	{_State455, ImplicitStructDefType}:            _GotoState91Action,
	{_State455, ExplicitStructDefType}:            _GotoState47Action,
	{_State455, ImplicitEnumDefType}:              _GotoState90Action,
	{_State455, ExplicitEnumDefType}:              _GotoState88Action,
	{_State455, TraitDefType}:                     _GotoState94Action,
	{_State455, ReturnTypeType}:                   _GotoState460Action,
	{_State455, FuncTypeType}:                     _GotoState89Action,
	{_State456, LbraceToken}:                      _GotoState62Action,
	{_State456, BlockBodyType}:                    _GotoState461Action,
	{_State459, IdentifierToken}:                  _GotoState81Action,
	{_State459, StructToken}:                      _GotoState35Action,
	{_State459, EnumToken}:                        _GotoState78Action,
	{_State459, TraitToken}:                       _GotoState86Action,
	{_State459, FuncToken}:                        _GotoState80Action,
	{_State459, LparenToken}:                      _GotoState83Action,
	{_State459, LbracketToken}:                    _GotoState27Action,
	{_State459, DotToken}:                         _GotoState77Action,
	{_State459, QuestionToken}:                    _GotoState84Action,
	{_State459, ExclaimToken}:                     _GotoState79Action,
	{_State459, TildeTildeToken}:                  _GotoState85Action,
	{_State459, BitNegToken}:                      _GotoState76Action,
	{_State459, BitAndToken}:                      _GotoState75Action,
	{_State459, LexErrorToken}:                    _GotoState82Action,
	{_State459, InitializableTypeType}:            _GotoState92Action,
	{_State459, AtomTypeType}:                     _GotoState87Action,
	{_State459, ReturnableTypeType}:               _GotoState356Action,
	{_State459, ImplicitStructDefType}:            _GotoState91Action,
	{_State459, ExplicitStructDefType}:            _GotoState47Action,
	{_State459, ImplicitEnumDefType}:              _GotoState90Action,
	{_State459, ExplicitEnumDefType}:              _GotoState88Action,
	{_State459, TraitDefType}:                     _GotoState94Action,
	{_State459, ReturnTypeType}:                   _GotoState462Action,
	{_State459, FuncTypeType}:                     _GotoState89Action,
	{_State462, LbraceToken}:                      _GotoState62Action,
	{_State462, BlockBodyType}:                    _GotoState463Action,
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
	{_State24, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State25, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralAction,
	{_State26, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State28, _WildcardMarker}:                   _ReduceLetToVarOrLetAction,
	{_State29, _WildcardMarker}:                   _ReduceLexErrorToAtomExprAction,
	{_State30, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State30, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State31, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State32, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State33, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State34, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State36, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State37, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
	{_State38, _WildcardMarker}:                   _ReduceVarToVarOrLetAction,
	{_State39, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State39, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State40, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State41, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State42, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State43, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State44, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State45, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State46, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State47, _WildcardMarker}:                   _ReduceExplicitStructDefToInitializableTypeAction,
	{_State48, _WildcardMarker}:                   _ReduceImplicitStructExprToAtomExprAction,
	{_State50, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State51, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State53, _WildcardMarker}:                   _ReduceOrExprToSequenceExprAction,
	{_State54, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State55, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State56, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State57, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State58, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State60, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State61, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State62, _WildcardMarker}:                   _ReduceEmptyListToStatementsAction,
	{_State63, _WildcardMarker}:                   _ReduceBlockBodyToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State65, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State66, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State67, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State68, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State69, _WildcardMarker}:                   _ReduceGlobalVarDefToDefinitionAction,
	{_State70, _WildcardMarker}:                   _ReduceEmptyListToPackageStatementsAction,
	{_State71, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State72, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State74, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State77, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State81, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State82, _WildcardMarker}:                   _ReduceLexErrorToAtomTypeAction,
	{_State83, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State87, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State88, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State89, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State90, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State91, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State92, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State93, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State94, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
	{_State96, _WildcardMarker}:                   _ReduceDotdotdotToArgumentAction,
	{_State97, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State98, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State100, _WildcardMarker}:                  _ReduceColonExpressionsToArgumentAction,
	{_State101, _WildcardMarker}:                  _ReducePositionalToArgumentAction,
	{_State101, ColonToken}:                       _ReduceExpressionToOptionalExpressionAction,
	{_State103, RparenToken}:                      _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State104, RbracketToken}:                    _ReduceNilToOptionalGenericArgumentsAction,
	{_State106, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State106, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State107, _WildcardMarker}:                  _ReduceQuestionToPostfixUnaryExprAction,
	{_State109, _WildcardMarker}:                  _ReduceAddToAddOpAction,
	{_State110, _WildcardMarker}:                  _ReduceBitOrToAddOpAction,
	{_State111, _WildcardMarker}:                  _ReduceBitXorToAddOpAction,
	{_State112, _WildcardMarker}:                  _ReduceSubToAddOpAction,
	{_State113, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State114, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State115, _WildcardMarker}:                  _ReduceEqualToCmpOpAction,
	{_State116, _WildcardMarker}:                  _ReduceGreaterToCmpOpAction,
	{_State117, _WildcardMarker}:                  _ReduceGreaterOrEqualToCmpOpAction,
	{_State118, _WildcardMarker}:                  _ReduceLessToCmpOpAction,
	{_State119, _WildcardMarker}:                  _ReduceLessOrEqualToCmpOpAction,
	{_State120, _WildcardMarker}:                  _ReduceNotEqualToCmpOpAction,
	{_State121, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State122, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State122, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State123, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State124, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State125, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State126, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State127, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State128, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State129, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State131, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State132, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State133, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State134, _WildcardMarker}:                  _ReduceToBlockExprAction,
	{_State135, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State136, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State137, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State138, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State140, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State141, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State143, _WildcardMarker}:                  _ReduceTuplePatternToVarPatternAction,
	{_State144, _WildcardMarker}:                  _ReduceNilToOptionalValueTypeAction,
	{_State145, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State146, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State147, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State148, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State151, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State153, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State155, _WildcardMarker}:                  _ReduceInferredRefArgToParameterDefAction,
	{_State158, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State159, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State160, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State161, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State162, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State164, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State165, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State167, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State168, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State171, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State171, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State173, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State175, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State176, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State177, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State178, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State179, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State184, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
	{_State186, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State187, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State187, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State188, _WildcardMarker}:                  _ReduceToImplicitStructExprAction,
	{_State189, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State189, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State189, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State189, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State189, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State190, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State190, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State190, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State190, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State190, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State191, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State192, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State194, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State196, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State197, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State199, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State199, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State199, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State200, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State201, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State202, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State204, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State205, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State208, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State208, SemicolonToken}:                   _ReduceSequenceExprToForAssignmentAction,
	{_State209, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State211, LbraceToken}:                      _ReduceSequenceExprToConditionAction,
	{_State213, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State214, _WildcardMarker}:                  _ReduceDotdotdotToFieldVarPatternAction,
	{_State215, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State216, _WildcardMarker}:                  _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State218, _WildcardMarker}:                  _ReducePositionalToFieldVarPatternAction,
	{_State219, _EndMarker}:                       _ReduceToVarDeclPatternAction,
	{_State220, _WildcardMarker}:                  _ReduceValueTypeToOptionalValueTypeAction,
	{_State221, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State222, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State223, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State224, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State225, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State226, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State227, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State227, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State229, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State230, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State231, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State232, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State233, AssignToken}:                      _ReduceToAssignPatternAction,
	{_State233, _WildcardMarker}:                  _ReduceSequenceExprToExpressionAction,
	{_State234, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State236, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State237, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State238, _WildcardMarker}:                  _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State240, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State241, _WildcardMarker}:                  _ReduceImportStatementToPackageStatementBodyAction,
	{_State242, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State244, _WildcardMarker}:                  _ReduceUnsafeStatementToPackageStatementBodyAction,
	{_State245, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State246, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State247, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State248, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State250, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State251, _EndMarker}:                       _ReduceAliasToNamedFuncDefAction,
	{_State252, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State253, _WildcardMarker}:                  _ReduceInferredRefVarargToParameterDefAction,
	{_State254, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State256, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State260, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State263, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State265, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State266, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State267, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State268, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State269, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State270, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State275, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State277, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State279, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State280, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State282, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State283, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State284, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State287, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State288, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State289, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State290, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State291, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State292, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State293, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State296, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State298, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State299, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State300, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State302, _WildcardMarker}:                  _ReduceInitializeExprToAtomExprAction,
	{_State303, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State304, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State305, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State306, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State306, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State309, _WildcardMarker}:                  _ReduceVarToVarOrLetAction,
	{_State310, _WildcardMarker}:                  _ReduceAssignPatternToCasePatternAction,
	{_State311, _WildcardMarker}:                  _ReduceCasePatternToCasePatternsAction,
	{_State313, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State314, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State318, _WildcardMarker}:                  _ReduceToTuplePatternAction,
	{_State319, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State320, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State320, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State320, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State321, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State321, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State321, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
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
	{_State336, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State337, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State338, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State339, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State340, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State340, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State340, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State341, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State342, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State344, _WildcardMarker}:                  _ReduceStringLiteralToImportClauseAction,
	{_State345, _WildcardMarker}:                  _ReduceSingleToImportStatementAction,
	{_State346, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State347, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State348, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State350, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State353, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State354, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State356, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State357, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State360, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State363, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State365, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State366, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State368, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State370, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State371, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State372, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State373, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State375, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State378, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State379, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State380, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State381, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State382, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State383, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State384, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State385, SemicolonToken}:                   _ReduceAssignToForAssignmentAction,
	{_State388, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State389, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State392, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State393, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State395, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State396, _WildcardMarker}:                  _ReduceCaseBranchToCaseBranchesAction,
	{_State397, RbraceToken}:                      _ReduceNilToOptionalDefaultBranchAction,
	{_State398, _WildcardMarker}:                  _ReduceNamedToFieldVarPatternAction,
	{_State399, _WildcardMarker}:                  _ReduceAddToFieldVarPatternsAction,
	{_State400, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State401, _WildcardMarker}:                  _ReduceAssignStatementToStatementBodyAction,
	{_State402, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State403, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State404, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State406, _WildcardMarker}:                  _ReduceFirstToImportClausesAction,
	{_State409, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State410, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State411, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State413, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State414, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State415, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State415, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State416, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State417, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State417, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State418, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State419, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State420, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State422, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State423, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State424, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State426, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State426, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State427, _WildcardMarker}:                  _ReduceEnumMatchPatternToCasePatternAction,
	{_State429, LbraceToken}:                      _ReduceCaseToConditionAction,
	{_State430, _WildcardMarker}:                  _ReduceMultipleToCasePatternsAction,
	{_State431, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State432, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State435, _WildcardMarker}:                  _ReduceAddToCaseBranchesAction,
	{_State436, RbraceToken}:                      _ReduceDefaultBranchToOptionalDefaultBranchAction,
	{_State438, _WildcardMarker}:                  _ReduceExplicitToImportClauseTerminalAction,
	{_State439, _WildcardMarker}:                  _ReduceImplicitToImportClauseTerminalAction,
	{_State440, _WildcardMarker}:                  _ReduceMultipleToImportStatementAction,
	{_State441, _WildcardMarker}:                  _ReduceAddToImportClausesAction,
	{_State442, _WildcardMarker}:                  _ReduceAliasToImportClauseAction,
	{_State444, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State445, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State447, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State449, _WildcardMarker}:                  _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State450, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State451, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State452, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State453, _EndMarker}:                       _ReduceFuncDefToNamedFuncDefAction,
	{_State455, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State457, _WildcardMarker}:                  _ReduceToCaseBranchAction,
	{_State458, RbraceToken}:                      _ReduceToDefaultBranchAction,
	{_State459, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State460, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State461, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State463, _EndMarker}:                       _ReduceMethodDefToNamedFuncDefAction,
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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 11
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 61
      COMMENT -> State 60
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
      VAR -> State 38
      LET -> State 28
      LBRACE -> State 62
      definitions -> State 65
      definition -> State 64
      var_decl_pattern -> State 69
      var_or_let -> State 59
      block_body -> State 63
      type_def -> State 68
      named_func_def -> State 66
      package_def -> State 67

  State 16:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 70

  State 17:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 71

  State 18:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 72
      LPAREN -> State 73

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
      LPAREN -> State 74

  State 24:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      initializable_type: LBRACKET.value_type RBRACKET
      initializable_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 95
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 28:
    Kernel Items:
      var_or_let: LET., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 97
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOTDOTDOT -> State 96
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 101
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      arguments -> State 99
      argument -> State 98
      colon_expressions -> State 100
      optional_expression -> State 102
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 31:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 103

  State 36:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      var_or_let: VAR., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 39:
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
      LBRACKET -> State 106
      DOT -> State 105
      QUESTION -> State 107
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 108

  State 40:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 109
      SUB -> State 112
      BIT_XOR -> State 111
      BIT_OR -> State 110
      add_op -> State 113

  State 41:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 114

  State 42:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 115
      NOT_EQUAL -> State 120
      LESS -> State 118
      LESS_OR_EQUAL -> State 119
      GREATER -> State 116
      GREATER_OR_EQUAL -> State 117
      cmp_op -> State 121

  State 47:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      atom_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 122

  State 50:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 128
      DIV -> State 126
      MOD -> State 127
      BIT_AND -> State 123
      BIT_LSHIFT -> State 124
      BIT_RSHIFT -> State 125
      mul_op -> State 129

  State 52:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 132
      SWITCH -> State 133
      FOR -> State 131
      DO -> State 130
      LBRACE -> State 62
      if_expr -> State 135
      switch_expr -> State 137
      loop_expr -> State 136
      block_body -> State 134

  State 53:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 138

  State 54:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 140
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 57:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      LPAREN -> State 142
      var_pattern -> State 144
      tuple_pattern -> State 143

  State 60:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 145

  State 63:
    Kernel Items:
      definition: block_body., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 146
      optional_newlines -> State 147

  State 66:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 148

  State 70:
    Kernel Items:
      package_def: PACKAGE LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 149

  State 71:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 151
      ASSIGN -> State 150
      optional_generic_parameters -> State 152

  State 72:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 151
      ASSIGN -> State 153
      optional_generic_parameters -> State 154

  State 73:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      parameter_def -> State 156

  State 74:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 155
      parameter_def -> State 158
      parameter_defs -> State 159
      optional_parameter_defs -> State 157

  State 75:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 160
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 76:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 161
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 77:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 162

  State 78:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 163

  State 79:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 164
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 80:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 165

  State 81:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 166
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 167

  State 82:
    Kernel Items:
      atom_type: LEX_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 171
      implicit_field_defs -> State 173
      optional_implicit_field_defs -> State 174
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 170
      implicit_enum_value_defs -> State 172
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 84:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 177
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 85:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 178
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 86:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 179

  State 87:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 95:
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
      RBRACKET -> State 184
      COMMA -> State 182
      COLON -> State 181
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 96:
    Kernel Items:
      argument: DOTDOTDOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 186

  State 98:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      arguments: arguments.COMMA argument
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 188
      COMMA -> State 187

  State 100:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 189

  State 101:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 190

  State 103:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 192
      implicit_struct_def -> State 91
      explicit_field_defs -> State 191
      optional_explicit_field_defs -> State 193
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 104:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      optional_generic_arguments -> State 195
      generic_arguments -> State 194
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 196
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 105:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 197

  State 106:
    Kernel Items:
      access_expr: access_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 97
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOTDOTDOT -> State 96
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 101
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      argument -> State 198
      colon_expressions -> State 100
      optional_expression -> State 102
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 107:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 199

  State 109:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 200
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 114:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 201
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 115:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 202
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 122:
    Kernel Items:
      atom_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 97
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOTDOTDOT -> State 96
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 101
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      arguments -> State 203
      argument -> State 98
      colon_expressions -> State 100
      optional_expression -> State 102
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 123:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 204
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 130:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 205

  State 131:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.assign_pattern IN sequence_expr DO block_body
      loop_expr: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      assign_pattern -> State 206
      optional_label_decl -> State 139
      for_assignment -> State 207
      sequence_expr -> State 208
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 132:
    Kernel Items:
      if_expr: IF.condition block_body
      if_expr: IF.condition block_body ELSE block_body
      if_expr: IF.condition block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      CASE -> State 209
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      condition -> State 210
      sequence_expr -> State 211
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 133:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 212
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 134:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 213
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 139:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 134

  State 140:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

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
      IDENTIFIER -> State 215
      LPAREN -> State 142
      DOTDOTDOT -> State 214
      var_pattern -> State 218
      tuple_pattern -> State 143
      field_var_patterns -> State 217
      field_var_pattern -> State 216

  State 143:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 144:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      optional_value_type -> State 219
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 220
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 145:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
      statements: statements.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      RETURN -> State 226
      BREAK -> State 222
      CONTINUE -> State 223
      UNSAFE -> State 169
      STRUCT -> State 35
      FUNC -> State 23
      ASYNC -> State 221
      DEFER -> State 224
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      RBRACE -> State 225
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      assign_pattern -> State 228
      expression -> State 229
      optional_label_decl -> State 52
      sequence_expr -> State 233
      block_expr -> State 44
      statement -> State 234
      statement_body -> State 235
      unsafe_statement -> State 236
      jump_statement -> State 231
      jump_type -> State 232
      expressions -> State 230
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 227
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 146:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      PACKAGE -> State 16
      TYPE -> State 17
      FUNC -> State 18
      VAR -> State 38
      LET -> State 28
      LBRACE -> State 62
      definition -> State 237
      var_decl_pattern -> State 69
      var_or_let -> State 59
      block_body -> State 63
      type_def -> State 68
      named_func_def -> State 66
      package_def -> State 67

  State 147:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 148:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 238
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 149:
    Kernel Items:
      package_def: PACKAGE LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      IMPORT -> State 239
      UNSAFE -> State 169
      RPAREN -> State 240
      unsafe_statement -> State 244
      package_statement_body -> State 243
      package_statement -> State 242
      import_statement -> State 241

  State 150:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 245
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 151:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 246
      generic_parameter_def -> State 247
      generic_parameter_defs -> State 248
      optional_generic_parameter_defs -> State 249

  State 152:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 250
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 153:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 251
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 154:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 252

  State 155:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.DOTDOTDOT
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      DOTDOTDOT -> State 253
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 254
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 156:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 255

  State 157:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 256

  State 158:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 257

  State 160:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 258
      implicit_enum_value_defs -> State 261
      implicit_enum_def -> State 90
      explicit_enum_value_defs -> State 259
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 164:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 165:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 263
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      DOTDOTDOT -> State 262
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 267
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      parameter_decl -> State 265
      parameter_decls -> State 266
      optional_parameter_decls -> State 264
      func_type -> State 89

  State 166:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 268

  State 167:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 168:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 269
      QUESTION -> State 84
      EXCLAIM -> State 79
      DOLLAR_LBRACKET -> State 104
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      optional_generic_binding -> State 167
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 270
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 169:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 271

  State 170:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 272

  State 171:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 273

  State 172:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 275
      OR -> State 274

  State 173:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 276

  State 174:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 277

  State 175:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 177:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 278
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 279
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_property -> State 283
      trait_properties -> State 282
      optional_trait_properties -> State 281
      trait_def -> State 94
      func_type -> State 89
      method_signature -> State 280

  State 180:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 284
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 181:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 285
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 182:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 286

  State 183:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 287
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 184:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 288
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 186:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 289
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 187:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 97
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOTDOTDOT -> State 96
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 101
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      argument -> State 290
      colon_expressions -> State 100
      optional_expression -> State 102
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 188:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      colon_expressions: colon_expressions COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 291
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      optional_expression -> State 292
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 190:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 291
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      optional_expression -> State 293
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 191:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 295
      COMMA -> State 294

  State 192:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 296

  State 194:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 297

  State 195:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 298

  State 196:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 197:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 299

  State 199:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 97
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOTDOTDOT -> State 96
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 101
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      optional_arguments -> State 301
      arguments -> State 300
      argument -> State 98
      colon_expressions -> State 100
      optional_expression -> State 102
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 200:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 128
      DIV -> State 126
      MOD -> State 127
      BIT_AND -> State 123
      BIT_LSHIFT -> State 124
      BIT_RSHIFT -> State 125
      mul_op -> State 129

  State 201:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 115
      NOT_EQUAL -> State 120
      LESS -> State 118
      LESS_OR_EQUAL -> State 119
      GREATER -> State 116
      GREATER_OR_EQUAL -> State 117
      cmp_op -> State 121

  State 202:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 109
      SUB -> State 112
      BIT_XOR -> State 111
      BIT_OR -> State 110
      add_op -> State 113

  State 203:
    Kernel Items:
      arguments: arguments.COMMA argument
      atom_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 302
      COMMA -> State 187

  State 204:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 303

  State 206:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO block_body
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 305
      ASSIGN -> State 304

  State 207:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 306

  State 208:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO block_body
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 307

  State 209:
    Kernel Items:
      condition: CASE.case_patterns ASSIGN sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 309
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOT -> State 308
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      assign_pattern -> State 310
      case_pattern -> State 311
      optional_label_decl -> State 139
      case_patterns -> State 312
      sequence_expr -> State 313
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 210:
    Kernel Items:
      if_expr: IF condition.block_body
      if_expr: IF condition.block_body ELSE block_body
      if_expr: IF condition.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 314

  State 211:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 315

  State 213:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 114

  State 214:
    Kernel Items:
      field_var_pattern: DOTDOTDOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 316

  State 216:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 318
      COMMA -> State 317

  State 218:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      optional_value_type: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 221:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 320
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 319
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 222:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      statement_body: DEFER.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      LEX_ERROR -> State 29
      optional_label_decl -> State 139
      block_expr -> State 44
      call_expr -> State 321
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 319
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 225:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
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
      LBRACKET -> State 106
      DOT -> State 105
      QUESTION -> State 107
      DOLLAR_LBRACKET -> State 104
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
      optional_generic_binding -> State 108

  State 228:
    Kernel Items:
      statement_body: assign_pattern.ASSIGN expression
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
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 338

  State 231:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
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
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 341
      SEMICOLON -> State 342

  State 236:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 344
      LPAREN -> State 343
      import_clause -> State 345

  State 240:
    Kernel Items:
      package_def: PACKAGE LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      package_statement_body: import_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 347
      COMMA -> State 346

  State 244:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 246:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 348
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 247:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 248:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 349

  State 249:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 350

  State 250:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 351
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 251:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 155
      parameter_def -> State 158
      parameter_defs -> State 159
      optional_parameter_defs -> State 352

  State 253:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT., *
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 353
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 254:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 255:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 354

  State 256:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 356
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      return_type -> State 355
      func_type -> State 89

  State 257:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      parameter_def -> State 357

  State 258:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 358
      OR -> State 359

  State 259:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 360

  State 260:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 273

  State 261:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 361
      OR -> State 362

  State 262:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 363
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 263:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 269
      QUESTION -> State 84
      EXCLAIM -> State 79
      DOLLAR_LBRACKET -> State 104
      DOTDOTDOT -> State 364
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      optional_generic_binding -> State 167
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 365
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 264:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 366

  State 265:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 367

  State 267:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 268:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 368

  State 269:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 268
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 162

  State 270:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 271:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 369

  State 272:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 370
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 273:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 371

  State 274:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 372
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 275:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 373
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 277:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 374
      LPAREN -> State 165

  State 279:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 375

  State 282:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 377
      COMMA -> State 376

  State 283:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 378
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 286:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 379

  State 287:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 380
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 295:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 381
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 296:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 382
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 298:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 187

  State 301:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 383

  State 302:
    Kernel Items:
      atom_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 384
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 304:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 385
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 305:
    Kernel Items:
      loop_expr: FOR assign_pattern IN.sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 386
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 306:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      optional_sequence_expr -> State 387
      sequence_expr -> State 388
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 307:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 389

  State 308:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 390

  State 309:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 391

  State 310:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      condition: CASE case_patterns.ASSIGN sequence_expr
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 393
      ASSIGN -> State 392

  State 313:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      if_expr: IF condition block_body., *
      if_expr: IF condition block_body.ELSE block_body
      if_expr: IF condition block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 394

  State 315:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 395
      case_branches -> State 397
      case_branch -> State 396

  State 316:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      LPAREN -> State 142
      var_pattern -> State 398
      tuple_pattern -> State 143

  State 317:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 215
      LPAREN -> State 142
      DOTDOTDOT -> State 214
      var_pattern -> State 218
      tuple_pattern -> State 143
      field_var_pattern -> State 399

  State 318:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 106
      DOT -> State 105
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 108

  State 320:
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

  State 321:
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
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 400
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 336:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      statement_body: assign_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 401
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 338:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 402
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 339:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      expression -> State 229
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 44
      expressions -> State 403
      optional_expressions -> State 404
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

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
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 344
      import_clause -> State 405
      import_clause_terminal -> State 406
      import_clauses -> State 407

  State 344:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 408

  State 345:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      package_statement: package_statement_body COMMA., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 348:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 349:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 246
      generic_parameter_def -> State 409

  State 350:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 410
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 352:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 411

  State 353:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 354:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 151
      optional_generic_parameters -> State 412

  State 355:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 413

  State 356:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 414
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 359:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 415
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 360:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 361:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 416
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 362:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 260
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      enum_value_def -> State 417
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 363:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 364:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 418
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      func_type -> State 89

  State 365:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 366:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 356
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      return_type -> State 419
      func_type -> State 89

  State 367:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 263
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      DOTDOTDOT -> State 262
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 267
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      parameter_decl -> State 420
      func_type -> State 89

  State 368:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 421

  State 370:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 373:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 422

  State 375:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 278
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 279
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_property -> State 423
      trait_def -> State 94
      func_type -> State 89
      method_signature -> State 280

  State 377:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      UNSAFE -> State 169
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 278
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      unsafe_statement -> State 175
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 176
      field_def -> State 279
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_property -> State 424
      trait_def -> State 94
      func_type -> State 89
      method_signature -> State 280

  State 378:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 383:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 385:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 425

  State 387:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 426

  State 388:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 389:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 390:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 30
      implicit_struct_expr -> State 427

  State 391:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 428

  State 392:
    Kernel Items:
      condition: CASE case_patterns ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 429
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 393:
    Kernel Items:
      case_patterns: case_patterns COMMA.case_pattern
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 309
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOT -> State 308
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      assign_pattern -> State 310
      case_pattern -> State 430
      optional_label_decl -> State 139
      sequence_expr -> State 313
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 394:
    Kernel Items:
      if_expr: IF condition block_body ELSE.block_body
      if_expr: IF condition block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 132
      LBRACE -> State 62
      if_expr -> State 432
      block_body -> State 431

  State 395:
    Kernel Items:
      case_branch: CASE.case_patterns COLON sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 309
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      DOT -> State 308
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      assign_pattern -> State 310
      case_pattern -> State 311
      optional_label_decl -> State 139
      case_patterns -> State 433
      sequence_expr -> State 313
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 396:
    Kernel Items:
      case_branches: case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 397:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches.optional_default_branch RBRACE
      case_branches: case_branches.case_branch
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      CASE -> State 395
      DEFAULT -> State 434
      case_branch -> State 435
      optional_default_branch -> State 437
      default_branch -> State 436

  State 398:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      statement_body: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 338

  State 404:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      import_clause_terminal: import_clause.NEWLINES
      import_clause_terminal: import_clause.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 439
      COMMA -> State 438

  State 406:
    Kernel Items:
      import_clauses: import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 407:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
      import_clauses: import_clauses.import_clause_terminal
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 344
      RPAREN -> State 440
      import_clause -> State 405
      import_clause_terminal -> State 441

  State 408:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 442

  State 409:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 410:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 411:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 356
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      return_type -> State 443
      func_type -> State 89

  State 412:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 444

  State 413:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 180
      SUB -> State 185
      MUL -> State 183

  State 419:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 445

  State 422:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 263
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      DOTDOTDOT -> State 262
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 93
      value_type -> State 267
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      parameter_decl -> State 265
      parameter_decls -> State 266
      optional_parameter_decls -> State 446
      func_type -> State 89

  State 423:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 447

  State 426:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      optional_sequence_expr -> State 448
      sequence_expr -> State 388
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 427:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 142
      tuple_pattern -> State 449

  State 429:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      if_expr: IF condition block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      if_expr: IF condition block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 433:
    Kernel Items:
      case_branch: CASE case_patterns.COLON sequence_expr
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 393
      COLON -> State 450

  State 434:
    Kernel Items:
      default_branch: DEFAULT.COLON sequence_expr
    Reduce:
      (nil)
    Goto:
      COLON -> State 451

  State 435:
    Kernel Items:
      case_branches: case_branches case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 436:
    Kernel Items:
      optional_default_branch: default_branch., RBRACE
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      (nil)

  State 437:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 452

  State 438:
    Kernel Items:
      import_clause_terminal: import_clause COMMA., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      import_clause_terminal: import_clause NEWLINES., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      import_clauses: import_clauses import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 443:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 453

  State 444:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 155
      parameter_def -> State 158
      parameter_defs -> State 159
      optional_parameter_defs -> State 454

  State 445:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 446:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 455

  State 447:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 456

  State 449:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 450:
    Kernel Items:
      case_branch: CASE case_patterns COLON.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 457
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 451:
    Kernel Items:
      default_branch: DEFAULT COLON.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 24
      TRUE -> State 37
      FALSE -> State 21
      STRUCT -> State 35
      FUNC -> State 23
      VAR -> State 38
      LET -> State 28
      LABEL_DECL -> State 26
      LPAREN -> State 30
      LBRACKET -> State 27
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      var_decl_pattern -> State 58
      var_or_let -> State 59
      optional_label_decl -> State 139
      sequence_expr -> State 458
      block_expr -> State 44
      call_expr -> State 45
      atom_expr -> State 43
      literal -> State 50
      implicit_struct_expr -> State 48
      access_expr -> State 39
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 40
      cmp_expr -> State 46
      and_expr -> State 41
      or_expr -> State 53
      initializable_type -> State 49
      explicit_struct_def -> State 47
      anonymous_func_expr -> State 42

  State 452:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 453:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 454:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 459

  State 455:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 356
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      return_type -> State 460
      func_type -> State 89

  State 456:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 461

  State 457:
    Kernel Items:
      case_branch: CASE case_patterns COLON sequence_expr., *
    Reduce:
      * -> [case_branch]
    Goto:
      (nil)

  State 458:
    Kernel Items:
      default_branch: DEFAULT COLON sequence_expr., RBRACE
    Reduce:
      RBRACE -> [default_branch]
    Goto:
      (nil)

  State 459:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 81
      STRUCT -> State 35
      ENUM -> State 78
      TRAIT -> State 86
      FUNC -> State 80
      LPAREN -> State 83
      LBRACKET -> State 27
      DOT -> State 77
      QUESTION -> State 84
      EXCLAIM -> State 79
      TILDE_TILDE -> State 85
      BIT_NEG -> State 76
      BIT_AND -> State 75
      LEX_ERROR -> State 82
      initializable_type -> State 92
      atom_type -> State 87
      returnable_type -> State 356
      implicit_struct_def -> State 91
      explicit_struct_def -> State 47
      implicit_enum_def -> State 90
      explicit_enum_def -> State 88
      trait_def -> State 94
      return_type -> State 462
      func_type -> State 89

  State 460:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 461:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 462:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 463

  State 463:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 463
Number of shift actions: 3124
Number of reduce actions: 372
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 4242
Number of unoptimized shift actions: 31295
Number of unoptimized reduce actions: 24666
*/
