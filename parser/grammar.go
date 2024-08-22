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
	BoolLiteralToken     = SymbolId(259)
	IntegerLiteralToken  = SymbolId(260)
	FloatLiteralToken    = SymbolId(261)
	RuneLiteralToken     = SymbolId(262)
	StringLiteralToken   = SymbolId(263)
	LbraceToken          = SymbolId(264)
	RbraceToken          = SymbolId(265)
	LparenToken          = SymbolId(266)
	RparenToken          = SymbolId(267)
	LbracketToken        = SymbolId(268)
	RbracketToken        = SymbolId(269)
	DotToken             = SymbolId(270)
	SemicolonToken       = SymbolId(271)
	IdentifierToken      = SymbolId(272)
	IfToken              = SymbolId(273)
	ElseToken            = SymbolId(274)
	MatchToken           = SymbolId(275)
	CaseToken            = SymbolId(276)
	DefaultToken         = SymbolId(277)
	ForToken             = SymbolId(278)
	ReturnToken          = SymbolId(279)
	BreakToken           = SymbolId(280)
	ContinueToken        = SymbolId(281)
	FuncToken            = SymbolId(282)
	LabelDeclToken       = SymbolId(283)
	JumpLabelToken       = SymbolId(284)
	AddAssignToken       = SymbolId(285)
	SubAssignToken       = SymbolId(286)
	MulAssignToken       = SymbolId(287)
	DivAssignToken       = SymbolId(288)
	ModAssignToken       = SymbolId(289)
	AddOneAssignToken    = SymbolId(290)
	SubOneAssignToken    = SymbolId(291)
	BitNegAssignToken    = SymbolId(292)
	BitAndAssignToken    = SymbolId(293)
	BitOrAssignToken     = SymbolId(294)
	BitXorAssignToken    = SymbolId(295)
	BitLshiftAssignToken = SymbolId(296)
	BitRshiftAssignToken = SymbolId(297)
	NotToken             = SymbolId(298)
	AndToken             = SymbolId(299)
	OrToken              = SymbolId(300)
	AddToken             = SymbolId(301)
	SubToken             = SymbolId(302)
	MulToken             = SymbolId(303)
	DivToken             = SymbolId(304)
	ModToken             = SymbolId(305)
	BitNegToken          = SymbolId(306)
	BitAndToken          = SymbolId(307)
	BitXorToken          = SymbolId(308)
	BitOrToken           = SymbolId(309)
	BitLshiftToken       = SymbolId(310)
	BitRshiftToken       = SymbolId(311)
	EqualToken           = SymbolId(312)
	NotEqualToken        = SymbolId(313)
	LessToken            = SymbolId(314)
	LessOrEqualToken     = SymbolId(315)
	GreaterToken         = SymbolId(316)
	GreaterOrEqualToken  = SymbolId(317)
	LexErrorToken        = SymbolId(318)
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
	// 65:2: literal -> BOOL_LITERAL: ...
	BoolLiteralToLiteral(BoolLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 66:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 67:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 68:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 69:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 72:23: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 74:24: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ *GenericSymbol, ArgumentList_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 81:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 82:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 83:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 84:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 85:2: atom_expr -> implicit_struct_expr: ...
	ImplicitStructExprToAtomExpr(ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 86:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 89:0: argument_list -> ...
	ToArgumentList() (*GenericSymbol, error)

	// 92:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 93:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 94:2: access_expr -> call: ...
	CallToAccessExpr(AccessExpr_ *GenericSymbol, Lparen_ *GenericSymbol, ArgumentList_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 96:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 99:2: unary_op -> NOT: ...
	NotToUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 100:2: unary_op -> BIT_NEG: ...
	BitNegToUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: unary_op -> SUB: ...
	SubToUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 104:2: unary_expr -> access_expr: ...
	AccessExprToUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 105:2: unary_expr -> op: ...
	OpToUnaryExpr(UnaryOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 108:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 109:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 110:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 111:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 113:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 116:2: mul_expr -> unary_expr: ...
	UnaryExprToMulExpr(UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 117:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 121:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 123:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 126:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 131:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 134:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 135:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 138:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 139:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 142:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 143:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 146:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 149:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 153:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 156:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 160:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 162:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 165:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 166:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 170:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 171:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 172:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 175:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 176:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 178:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 179:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: statement_body -> expression: ...
	ExpressionToStatementBody(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 183:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 191:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 192:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 196:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 200:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 202:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 203:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 208:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 210:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 213:14: match_expr -> ...
	ToMatchExpr(Match_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 223:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 224:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: expression -> match_expr: ...
	MatchExprToExpression(OptionalLabelDecl_ *GenericSymbol, MatchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 226:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 230:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 231:2: lex_internal_tokens -> COMMENT: ...
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

	item, err := _Parse(lexer, reducer, errHandler, _State1)
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

	item, err := _Parse(lexer, reducer, errHandler, _State2)
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
	case BoolLiteralToken:
		return "BOOL_LITERAL"
	case IntegerLiteralToken:
		return "INTEGER_LITERAL"
	case FloatLiteralToken:
		return "FLOAT_LITERAL"
	case RuneLiteralToken:
		return "RUNE_LITERAL"
	case StringLiteralToken:
		return "STRING_LITERAL"
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
	case SemicolonToken:
		return "SEMICOLON"
	case IdentifierToken:
		return "IDENTIFIER"
	case IfToken:
		return "IF"
	case ElseToken:
		return "ELSE"
	case MatchToken:
		return "MATCH"
	case CaseToken:
		return "CASE"
	case DefaultToken:
		return "DEFAULT"
	case ForToken:
		return "FOR"
	case ReturnToken:
		return "RETURN"
	case BreakToken:
		return "BREAK"
	case ContinueToken:
		return "CONTINUE"
	case FuncToken:
		return "FUNC"
	case LabelDeclToken:
		return "LABEL_DECL"
	case JumpLabelToken:
		return "JUMP_LABEL"
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
	case LiteralType:
		return "literal"
	case AnonymousFuncExprType:
		return "anonymous_func_expr"
	case ImplicitStructExprType:
		return "implicit_struct_expr"
	case AtomExprType:
		return "atom_expr"
	case ArgumentListType:
		return "argument_list"
	case AccessExprType:
		return "access_expr"
	case UnaryOpType:
		return "unary_op"
	case UnaryExprType:
		return "unary_expr"
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
	case SequenceExprType:
		return "sequence_expr"
	case OptionalJumpLabelType:
		return "optional_jump_label"
	case OptionalExpressionType:
		return "optional_expression"
	case JumpTypeType:
		return "jump_type"
	case OpOneAssignType:
		return "op_one_assign"
	case BinaryOpAssignType:
		return "binary_op_assign"
	case StatementBodyType:
		return "statement_body"
	case StatementType:
		return "statement"
	case StatementsType:
		return "statements"
	case OptionalLabelDeclType:
		return "optional_label_decl"
	case BlockBodyType:
		return "block_body"
	case BlockExprType:
		return "block_expr"
	case IfExprType:
		return "if_expr"
	case MatchExprType:
		return "match_expr"
	case LoopExprType:
		return "loop_expr"
	case ExpressionType:
		return "expression"
	case LexInternalTokensType:
		return "lex_internal_tokens"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	LiteralType            = SymbolId(319)
	AnonymousFuncExprType  = SymbolId(320)
	ImplicitStructExprType = SymbolId(321)
	AtomExprType           = SymbolId(322)
	ArgumentListType       = SymbolId(323)
	AccessExprType         = SymbolId(324)
	UnaryOpType            = SymbolId(325)
	UnaryExprType          = SymbolId(326)
	MulOpType              = SymbolId(327)
	MulExprType            = SymbolId(328)
	AddOpType              = SymbolId(329)
	AddExprType            = SymbolId(330)
	CmpOpType              = SymbolId(331)
	CmpExprType            = SymbolId(332)
	AndExprType            = SymbolId(333)
	OrExprType             = SymbolId(334)
	SequenceExprType       = SymbolId(335)
	OptionalJumpLabelType  = SymbolId(336)
	OptionalExpressionType = SymbolId(337)
	JumpTypeType           = SymbolId(338)
	OpOneAssignType        = SymbolId(339)
	BinaryOpAssignType     = SymbolId(340)
	StatementBodyType      = SymbolId(341)
	StatementType          = SymbolId(342)
	StatementsType         = SymbolId(343)
	OptionalLabelDeclType  = SymbolId(344)
	BlockBodyType          = SymbolId(345)
	BlockExprType          = SymbolId(346)
	IfExprType             = SymbolId(347)
	MatchExprType          = SymbolId(348)
	LoopExprType           = SymbolId(349)
	ExpressionType         = SymbolId(350)
	LexInternalTokensType  = SymbolId(351)
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
	_ReduceBoolLiteralToLiteral            = _ReduceType(1)
	_ReduceIntegerLiteralToLiteral         = _ReduceType(2)
	_ReduceFloatLiteralToLiteral           = _ReduceType(3)
	_ReduceRuneLiteralToLiteral            = _ReduceType(4)
	_ReduceStringLiteralToLiteral          = _ReduceType(5)
	_ReduceToAnonymousFuncExpr             = _ReduceType(6)
	_ReduceToImplicitStructExpr            = _ReduceType(7)
	_ReduceLiteralToAtomExpr               = _ReduceType(8)
	_ReduceIdentifierToAtomExpr            = _ReduceType(9)
	_ReduceBlockExprToAtomExpr             = _ReduceType(10)
	_ReduceAnonymousFuncExprToAtomExpr     = _ReduceType(11)
	_ReduceImplicitStructExprToAtomExpr    = _ReduceType(12)
	_ReduceLexErrorToAtomExpr              = _ReduceType(13)
	_ReduceToArgumentList                  = _ReduceType(14)
	_ReduceAtomExprToAccessExpr            = _ReduceType(15)
	_ReduceAccessToAccessExpr              = _ReduceType(16)
	_ReduceCallToAccessExpr                = _ReduceType(17)
	_ReduceIndexToAccessExpr               = _ReduceType(18)
	_ReduceNotToUnaryOp                    = _ReduceType(19)
	_ReduceBitNegToUnaryOp                 = _ReduceType(20)
	_ReduceSubToUnaryOp                    = _ReduceType(21)
	_ReduceAccessExprToUnaryExpr           = _ReduceType(22)
	_ReduceOpToUnaryExpr                   = _ReduceType(23)
	_ReduceMulToMulOp                      = _ReduceType(24)
	_ReduceDivToMulOp                      = _ReduceType(25)
	_ReduceModToMulOp                      = _ReduceType(26)
	_ReduceBitAndToMulOp                   = _ReduceType(27)
	_ReduceBitLshiftToMulOp                = _ReduceType(28)
	_ReduceBitRshiftToMulOp                = _ReduceType(29)
	_ReduceUnaryExprToMulExpr              = _ReduceType(30)
	_ReduceOpToMulExpr                     = _ReduceType(31)
	_ReduceAddToAddOp                      = _ReduceType(32)
	_ReduceSubToAddOp                      = _ReduceType(33)
	_ReduceBitOrToAddOp                    = _ReduceType(34)
	_ReduceBitXorToAddOp                   = _ReduceType(35)
	_ReduceMulExprToAddExpr                = _ReduceType(36)
	_ReduceOpToAddExpr                     = _ReduceType(37)
	_ReduceEqualToCmpOp                    = _ReduceType(38)
	_ReduceNotEqualToCmpOp                 = _ReduceType(39)
	_ReduceLessToCmpOp                     = _ReduceType(40)
	_ReduceLessOrEqualToCmpOp              = _ReduceType(41)
	_ReduceGreaterToCmpOp                  = _ReduceType(42)
	_ReduceGreaterOrEqualToCmpOp           = _ReduceType(43)
	_ReduceAddExprToCmpExpr                = _ReduceType(44)
	_ReduceOpToCmpExpr                     = _ReduceType(45)
	_ReduceCmpExprToAndExpr                = _ReduceType(46)
	_ReduceOpToAndExpr                     = _ReduceType(47)
	_ReduceAndExprToOrExpr                 = _ReduceType(48)
	_ReduceOpToOrExpr                      = _ReduceType(49)
	_ReduceToSequenceExpr                  = _ReduceType(50)
	_ReduceJumpLabelToOptionalJumpLabel    = _ReduceType(51)
	_ReduceUnlabelledToOptionalJumpLabel   = _ReduceType(52)
	_ReduceExpressionToOptionalExpression  = _ReduceType(53)
	_ReduceNilToOptionalExpression         = _ReduceType(54)
	_ReduceReturnToJumpType                = _ReduceType(55)
	_ReduceBreakToJumpType                 = _ReduceType(56)
	_ReduceContinueToJumpType              = _ReduceType(57)
	_ReduceAddOneAssignToOpOneAssign       = _ReduceType(58)
	_ReduceSubOneAssignToOpOneAssign       = _ReduceType(59)
	_ReduceAddAssignToBinaryOpAssign       = _ReduceType(60)
	_ReduceSubAssignToBinaryOpAssign       = _ReduceType(61)
	_ReduceMulAssignToBinaryOpAssign       = _ReduceType(62)
	_ReduceDivAssignToBinaryOpAssign       = _ReduceType(63)
	_ReduceModAssignToBinaryOpAssign       = _ReduceType(64)
	_ReduceBitNegAssignToBinaryOpAssign    = _ReduceType(65)
	_ReduceBitAndAssignToBinaryOpAssign    = _ReduceType(66)
	_ReduceBitOrAssignToBinaryOpAssign     = _ReduceType(67)
	_ReduceBitXorAssignToBinaryOpAssign    = _ReduceType(68)
	_ReduceBitLshiftAssignToBinaryOpAssign = _ReduceType(69)
	_ReduceBitRshiftAssignToBinaryOpAssign = _ReduceType(70)
	_ReduceExpressionToStatementBody       = _ReduceType(71)
	_ReduceJumpToStatementBody             = _ReduceType(72)
	_ReduceOpOneAssignToStatementBody      = _ReduceType(73)
	_ReduceBinaryOpAssignToStatementBody   = _ReduceType(74)
	_ReduceImplicitToStatement             = _ReduceType(75)
	_ReduceExplicitToStatement             = _ReduceType(76)
	_ReduceEmptyListToStatements           = _ReduceType(77)
	_ReduceAddToStatements                 = _ReduceType(78)
	_ReduceLabelDeclToOptionalLabelDecl    = _ReduceType(79)
	_ReduceUnlabelledToOptionalLabelDecl   = _ReduceType(80)
	_ReduceToBlockBody                     = _ReduceType(81)
	_ReduceToBlockExpr                     = _ReduceType(82)
	_ReduceNoElseToIfExpr                  = _ReduceType(83)
	_ReduceIfElseToIfExpr                  = _ReduceType(84)
	_ReduceMultiIfElseToIfExpr             = _ReduceType(85)
	_ReduceToMatchExpr                     = _ReduceType(86)
	_ReduceInfiniteToLoopExpr              = _ReduceType(87)
	_ReduceWhileToLoopExpr                 = _ReduceType(88)
	_ReduceSequenceExprToExpression        = _ReduceType(89)
	_ReduceIfExprToExpression              = _ReduceType(90)
	_ReduceMatchExprToExpression           = _ReduceType(91)
	_ReduceLoopExprToExpression            = _ReduceType(92)
	_ReduceSpacesToLexInternalTokens       = _ReduceType(93)
	_ReduceCommentToLexInternalTokens      = _ReduceType(94)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceBoolLiteralToLiteral:
		return "BoolLiteralToLiteral"
	case _ReduceIntegerLiteralToLiteral:
		return "IntegerLiteralToLiteral"
	case _ReduceFloatLiteralToLiteral:
		return "FloatLiteralToLiteral"
	case _ReduceRuneLiteralToLiteral:
		return "RuneLiteralToLiteral"
	case _ReduceStringLiteralToLiteral:
		return "StringLiteralToLiteral"
	case _ReduceToAnonymousFuncExpr:
		return "ToAnonymousFuncExpr"
	case _ReduceToImplicitStructExpr:
		return "ToImplicitStructExpr"
	case _ReduceLiteralToAtomExpr:
		return "LiteralToAtomExpr"
	case _ReduceIdentifierToAtomExpr:
		return "IdentifierToAtomExpr"
	case _ReduceBlockExprToAtomExpr:
		return "BlockExprToAtomExpr"
	case _ReduceAnonymousFuncExprToAtomExpr:
		return "AnonymousFuncExprToAtomExpr"
	case _ReduceImplicitStructExprToAtomExpr:
		return "ImplicitStructExprToAtomExpr"
	case _ReduceLexErrorToAtomExpr:
		return "LexErrorToAtomExpr"
	case _ReduceToArgumentList:
		return "ToArgumentList"
	case _ReduceAtomExprToAccessExpr:
		return "AtomExprToAccessExpr"
	case _ReduceAccessToAccessExpr:
		return "AccessToAccessExpr"
	case _ReduceCallToAccessExpr:
		return "CallToAccessExpr"
	case _ReduceIndexToAccessExpr:
		return "IndexToAccessExpr"
	case _ReduceNotToUnaryOp:
		return "NotToUnaryOp"
	case _ReduceBitNegToUnaryOp:
		return "BitNegToUnaryOp"
	case _ReduceSubToUnaryOp:
		return "SubToUnaryOp"
	case _ReduceAccessExprToUnaryExpr:
		return "AccessExprToUnaryExpr"
	case _ReduceOpToUnaryExpr:
		return "OpToUnaryExpr"
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
	case _ReduceUnaryExprToMulExpr:
		return "UnaryExprToMulExpr"
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
	case _ReduceToSequenceExpr:
		return "ToSequenceExpr"
	case _ReduceJumpLabelToOptionalJumpLabel:
		return "JumpLabelToOptionalJumpLabel"
	case _ReduceUnlabelledToOptionalJumpLabel:
		return "UnlabelledToOptionalJumpLabel"
	case _ReduceExpressionToOptionalExpression:
		return "ExpressionToOptionalExpression"
	case _ReduceNilToOptionalExpression:
		return "NilToOptionalExpression"
	case _ReduceReturnToJumpType:
		return "ReturnToJumpType"
	case _ReduceBreakToJumpType:
		return "BreakToJumpType"
	case _ReduceContinueToJumpType:
		return "ContinueToJumpType"
	case _ReduceAddOneAssignToOpOneAssign:
		return "AddOneAssignToOpOneAssign"
	case _ReduceSubOneAssignToOpOneAssign:
		return "SubOneAssignToOpOneAssign"
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
	case _ReduceExpressionToStatementBody:
		return "ExpressionToStatementBody"
	case _ReduceJumpToStatementBody:
		return "JumpToStatementBody"
	case _ReduceOpOneAssignToStatementBody:
		return "OpOneAssignToStatementBody"
	case _ReduceBinaryOpAssignToStatementBody:
		return "BinaryOpAssignToStatementBody"
	case _ReduceImplicitToStatement:
		return "ImplicitToStatement"
	case _ReduceExplicitToStatement:
		return "ExplicitToStatement"
	case _ReduceEmptyListToStatements:
		return "EmptyListToStatements"
	case _ReduceAddToStatements:
		return "AddToStatements"
	case _ReduceLabelDeclToOptionalLabelDecl:
		return "LabelDeclToOptionalLabelDecl"
	case _ReduceUnlabelledToOptionalLabelDecl:
		return "UnlabelledToOptionalLabelDecl"
	case _ReduceToBlockBody:
		return "ToBlockBody"
	case _ReduceToBlockExpr:
		return "ToBlockExpr"
	case _ReduceNoElseToIfExpr:
		return "NoElseToIfExpr"
	case _ReduceIfElseToIfExpr:
		return "IfElseToIfExpr"
	case _ReduceMultiIfElseToIfExpr:
		return "MultiIfElseToIfExpr"
	case _ReduceToMatchExpr:
		return "ToMatchExpr"
	case _ReduceInfiniteToLoopExpr:
		return "InfiniteToLoopExpr"
	case _ReduceWhileToLoopExpr:
		return "WhileToLoopExpr"
	case _ReduceSequenceExprToExpression:
		return "SequenceExprToExpression"
	case _ReduceIfExprToExpression:
		return "IfExprToExpression"
	case _ReduceMatchExprToExpression:
		return "MatchExprToExpression"
	case _ReduceLoopExprToExpression:
		return "LoopExprToExpression"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, BoolLiteralToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, SemicolonToken, IdentifierToken, IfToken, ElseToken, MatchToken, CaseToken, DefaultToken, ForToken, ReturnToken, BreakToken, ContinueToken, FuncToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceBoolLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.BoolLiteralToLiteral(args[0].Generic_)
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
	case _ReduceToAnonymousFuncExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AnonymousFuncExprType
		symbol.Generic_, err = reducer.ToAnonymousFuncExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceToImplicitStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructExprType
		symbol.Generic_, err = reducer.ToImplicitStructExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceToArgumentList:
		symbol.SymbolId_ = ArgumentListType
		symbol.Generic_, err = reducer.ToArgumentList()
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
	case _ReduceCallToAccessExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.CallToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceIndexToAccessExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.IndexToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceNotToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.NotToUnaryOp(args[0].Generic_)
	case _ReduceBitNegToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.BitNegToUnaryOp(args[0].Generic_)
	case _ReduceSubToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.SubToUnaryOp(args[0].Generic_)
	case _ReduceAccessExprToUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryExprType
		symbol.Generic_, err = reducer.AccessExprToUnaryExpr(args[0].Generic_)
	case _ReduceOpToUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = UnaryExprType
		symbol.Generic_, err = reducer.OpToUnaryExpr(args[0].Generic_, args[1].Generic_)
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
	case _ReduceUnaryExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		symbol.Generic_, err = reducer.UnaryExprToMulExpr(args[0].Generic_)
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
	case _ReduceToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		symbol.Generic_, err = reducer.ToSequenceExpr(args[0].Generic_)
	case _ReduceJumpLabelToOptionalJumpLabel:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.JumpLabelToOptionalJumpLabel(args[0].Generic_)
	case _ReduceUnlabelledToOptionalJumpLabel:
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.UnlabelledToOptionalJumpLabel()
	case _ReduceExpressionToOptionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.ExpressionToOptionalExpression(args[0].Generic_)
	case _ReduceNilToOptionalExpression:
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.NilToOptionalExpression()
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
	case _ReduceAddOneAssignToOpOneAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpOneAssignType
		symbol.Generic_, err = reducer.AddOneAssignToOpOneAssign(args[0].Generic_)
	case _ReduceSubOneAssignToOpOneAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpOneAssignType
		symbol.Generic_, err = reducer.SubOneAssignToOpOneAssign(args[0].Generic_)
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
	case _ReduceExpressionToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.ExpressionToStatementBody(args[0].Generic_)
	case _ReduceJumpToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.JumpToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceOpOneAssignToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.OpOneAssignToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceBinaryOpAssignToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.BinaryOpAssignToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceEmptyListToStatements:
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.EmptyListToStatements()
	case _ReduceAddToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.AddToStatements(args[0].Generic_, args[1].Generic_)
	case _ReduceLabelDeclToOptionalLabelDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.LabelDeclToOptionalLabelDecl(args[0].Generic_)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.UnlabelledToOptionalLabelDecl()
	case _ReduceToBlockBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BlockBodyType
		symbol.Generic_, err = reducer.ToBlockBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToBlockExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BlockExprType
		symbol.Generic_, err = reducer.ToBlockExpr(args[0].Generic_, args[1].Generic_)
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
	case _ReduceToMatchExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = MatchExprType
		symbol.Generic_, err = reducer.ToMatchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceInfiniteToLoopExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.InfiniteToLoopExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceWhileToLoopExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.WhileToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceSequenceExprToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.SequenceExprToExpression(args[0].Generic_)
	case _ReduceIfExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.IfExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceMatchExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.MatchExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceLoopExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LoopExprToExpression(args[0].Generic_, args[1].Generic_)
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
	_GotoState1Action                            = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                            = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                            = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                            = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                            = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                            = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                            = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                            = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                            = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                           = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                           = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                           = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                           = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                           = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                           = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                           = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                           = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                           = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                           = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                           = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                           = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                           = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                           = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                           = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                           = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                           = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                           = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                           = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                           = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                           = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                           = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                           = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                           = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                           = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                           = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                           = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                           = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                           = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                           = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                           = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                           = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                           = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                           = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                           = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                           = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                           = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                           = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                           = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                           = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                           = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                           = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                           = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                           = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                           = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                           = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                           = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                           = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                           = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                           = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                           = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                           = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                           = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                           = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                           = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                           = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                           = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                           = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                           = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                           = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                           = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                           = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                           = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                           = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                           = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                           = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                           = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                           = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                           = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                           = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                           = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                           = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                           = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                           = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                           = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                           = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                           = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                           = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                           = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                           = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                           = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                           = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                           = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                           = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                           = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                           = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                           = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                           = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                           = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                           = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                          = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                          = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                          = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                          = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                          = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                          = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                          = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                          = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                          = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                          = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                          = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                          = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                          = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                          = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                          = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                          = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                          = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                          = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                          = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                          = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                          = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                          = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                          = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                          = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                          = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                          = &_Action{_ShiftAction, _State125, 0}
	_ReduceBoolLiteralToLiteralAction            = &_Action{_ReduceAction, 0, _ReduceBoolLiteralToLiteral}
	_ReduceIntegerLiteralToLiteralAction         = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction           = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction            = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction          = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceToAnonymousFuncExprAction             = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceToImplicitStructExprAction            = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceLiteralToAtomExprAction               = &_Action{_ReduceAction, 0, _ReduceLiteralToAtomExpr}
	_ReduceIdentifierToAtomExprAction            = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction             = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction     = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction    = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceLexErrorToAtomExprAction              = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomExpr}
	_ReduceToArgumentListAction                  = &_Action{_ReduceAction, 0, _ReduceToArgumentList}
	_ReduceAtomExprToAccessExprAction            = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessExpr}
	_ReduceAccessToAccessExprAction              = &_Action{_ReduceAction, 0, _ReduceAccessToAccessExpr}
	_ReduceCallToAccessExprAction                = &_Action{_ReduceAction, 0, _ReduceCallToAccessExpr}
	_ReduceIndexToAccessExprAction               = &_Action{_ReduceAction, 0, _ReduceIndexToAccessExpr}
	_ReduceNotToUnaryOpAction                    = &_Action{_ReduceAction, 0, _ReduceNotToUnaryOp}
	_ReduceBitNegToUnaryOpAction                 = &_Action{_ReduceAction, 0, _ReduceBitNegToUnaryOp}
	_ReduceSubToUnaryOpAction                    = &_Action{_ReduceAction, 0, _ReduceSubToUnaryOp}
	_ReduceAccessExprToUnaryExprAction           = &_Action{_ReduceAction, 0, _ReduceAccessExprToUnaryExpr}
	_ReduceOpToUnaryExprAction                   = &_Action{_ReduceAction, 0, _ReduceOpToUnaryExpr}
	_ReduceMulToMulOpAction                      = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                      = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                      = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                   = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceUnaryExprToMulExprAction              = &_Action{_ReduceAction, 0, _ReduceUnaryExprToMulExpr}
	_ReduceOpToMulExprAction                     = &_Action{_ReduceAction, 0, _ReduceOpToMulExpr}
	_ReduceAddToAddOpAction                      = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                      = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                    = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                   = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceMulExprToAddExprAction                = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceOpToAddExprAction                     = &_Action{_ReduceAction, 0, _ReduceOpToAddExpr}
	_ReduceEqualToCmpOpAction                    = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                 = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                     = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction              = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                  = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction           = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceAddExprToCmpExprAction                = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceOpToCmpExprAction                     = &_Action{_ReduceAction, 0, _ReduceOpToCmpExpr}
	_ReduceCmpExprToAndExprAction                = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceOpToAndExprAction                     = &_Action{_ReduceAction, 0, _ReduceOpToAndExpr}
	_ReduceAndExprToOrExprAction                 = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceOpToOrExprAction                      = &_Action{_ReduceAction, 0, _ReduceOpToOrExpr}
	_ReduceToSequenceExprAction                  = &_Action{_ReduceAction, 0, _ReduceToSequenceExpr}
	_ReduceJumpLabelToOptionalJumpLabelAction    = &_Action{_ReduceAction, 0, _ReduceJumpLabelToOptionalJumpLabel}
	_ReduceUnlabelledToOptionalJumpLabelAction   = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalJumpLabel}
	_ReduceExpressionToOptionalExpressionAction  = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReduceReturnToJumpTypeAction                = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                 = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction              = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceAddOneAssignToOpOneAssignAction       = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToOpOneAssign}
	_ReduceSubOneAssignToOpOneAssignAction       = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToOpOneAssign}
	_ReduceAddAssignToBinaryOpAssignAction       = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction       = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction       = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction       = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction       = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction    = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction    = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction     = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction    = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceExpressionToStatementBodyAction       = &_Action{_ReduceAction, 0, _ReduceExpressionToStatementBody}
	_ReduceJumpToStatementBodyAction             = &_Action{_ReduceAction, 0, _ReduceJumpToStatementBody}
	_ReduceOpOneAssignToStatementBodyAction      = &_Action{_ReduceAction, 0, _ReduceOpOneAssignToStatementBody}
	_ReduceBinaryOpAssignToStatementBodyAction   = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignToStatementBody}
	_ReduceImplicitToStatementAction             = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction             = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceEmptyListToStatementsAction           = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                 = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceLabelDeclToOptionalLabelDeclAction    = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction   = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceToBlockBodyAction                     = &_Action{_ReduceAction, 0, _ReduceToBlockBody}
	_ReduceToBlockExprAction                     = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceNoElseToIfExprAction                  = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                  = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction             = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceToMatchExprAction                     = &_Action{_ReduceAction, 0, _ReduceToMatchExpr}
	_ReduceInfiniteToLoopExprAction              = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceWhileToLoopExprAction                 = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceSequenceExprToExpressionAction        = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceIfExprToExpressionAction              = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceMatchExprToExpressionAction           = &_Action{_ReduceAction, 0, _ReduceMatchExprToExpression}
	_ReduceLoopExprToExpressionAction            = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceSpacesToLexInternalTokensAction       = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction      = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State3, _EndMarker}:               &_Action{_AcceptAction, 0, 0},
	{_State4, _EndMarker}:               &_Action{_AcceptAction, 0, 0},
	{_State1, BoolLiteralToken}:         _GotoState6Action,
	{_State1, IntegerLiteralToken}:      _GotoState10Action,
	{_State1, FloatLiteralToken}:        _GotoState7Action,
	{_State1, RuneLiteralToken}:         _GotoState15Action,
	{_State1, StringLiteralToken}:       _GotoState16Action,
	{_State1, LparenToken}:              _GotoState13Action,
	{_State1, IdentifierToken}:          _GotoState9Action,
	{_State1, FuncToken}:                _GotoState8Action,
	{_State1, LabelDeclToken}:           _GotoState11Action,
	{_State1, NotToken}:                 _GotoState14Action,
	{_State1, SubToken}:                 _GotoState17Action,
	{_State1, BitNegToken}:              _GotoState5Action,
	{_State1, LexErrorToken}:            _GotoState12Action,
	{_State1, LiteralType}:              _GotoState26Action,
	{_State1, AnonymousFuncExprType}:    _GotoState21Action,
	{_State1, ImplicitStructExprType}:   _GotoState25Action,
	{_State1, AtomExprType}:             _GotoState22Action,
	{_State1, AccessExprType}:           _GotoState18Action,
	{_State1, UnaryOpType}:              _GotoState32Action,
	{_State1, UnaryExprType}:            _GotoState31Action,
	{_State1, MulExprType}:              _GotoState27Action,
	{_State1, AddExprType}:              _GotoState19Action,
	{_State1, CmpExprType}:              _GotoState24Action,
	{_State1, AndExprType}:              _GotoState20Action,
	{_State1, OrExprType}:               _GotoState29Action,
	{_State1, SequenceExprType}:         _GotoState30Action,
	{_State1, OptionalLabelDeclType}:    _GotoState28Action,
	{_State1, BlockExprType}:            _GotoState23Action,
	{_State1, ExpressionType}:           _GotoState3Action,
	{_State2, SpacesToken}:              _GotoState34Action,
	{_State2, CommentToken}:             _GotoState33Action,
	{_State2, LexInternalTokensType}:    _GotoState4Action,
	{_State8, LparenToken}:              _GotoState35Action,
	{_State13, ArgumentListType}:        _GotoState36Action,
	{_State18, LparenToken}:             _GotoState39Action,
	{_State18, LbracketToken}:           _GotoState38Action,
	{_State18, DotToken}:                _GotoState37Action,
	{_State19, AddToken}:                _GotoState40Action,
	{_State19, SubToken}:                _GotoState43Action,
	{_State19, BitXorToken}:             _GotoState42Action,
	{_State19, BitOrToken}:              _GotoState41Action,
	{_State19, AddOpType}:               _GotoState44Action,
	{_State20, AndToken}:                _GotoState45Action,
	{_State24, EqualToken}:              _GotoState46Action,
	{_State24, NotEqualToken}:           _GotoState51Action,
	{_State24, LessToken}:               _GotoState49Action,
	{_State24, LessOrEqualToken}:        _GotoState50Action,
	{_State24, GreaterToken}:            _GotoState47Action,
	{_State24, GreaterOrEqualToken}:     _GotoState48Action,
	{_State24, CmpOpType}:               _GotoState52Action,
	{_State27, MulToken}:                _GotoState58Action,
	{_State27, DivToken}:                _GotoState56Action,
	{_State27, ModToken}:                _GotoState57Action,
	{_State27, BitAndToken}:             _GotoState53Action,
	{_State27, BitLshiftToken}:          _GotoState54Action,
	{_State27, BitRshiftToken}:          _GotoState55Action,
	{_State27, MulOpType}:               _GotoState59Action,
	{_State28, LbraceToken}:             _GotoState62Action,
	{_State28, IfToken}:                 _GotoState61Action,
	{_State28, MatchToken}:              _GotoState63Action,
	{_State28, ForToken}:                _GotoState60Action,
	{_State28, BlockBodyType}:           _GotoState64Action,
	{_State28, IfExprType}:              _GotoState65Action,
	{_State28, MatchExprType}:           _GotoState67Action,
	{_State28, LoopExprType}:            _GotoState66Action,
	{_State29, OrToken}:                 _GotoState68Action,
	{_State32, BoolLiteralToken}:        _GotoState6Action,
	{_State32, IntegerLiteralToken}:     _GotoState10Action,
	{_State32, FloatLiteralToken}:       _GotoState7Action,
	{_State32, RuneLiteralToken}:        _GotoState15Action,
	{_State32, StringLiteralToken}:      _GotoState16Action,
	{_State32, LparenToken}:             _GotoState13Action,
	{_State32, IdentifierToken}:         _GotoState9Action,
	{_State32, FuncToken}:               _GotoState8Action,
	{_State32, LabelDeclToken}:          _GotoState11Action,
	{_State32, NotToken}:                _GotoState14Action,
	{_State32, SubToken}:                _GotoState17Action,
	{_State32, BitNegToken}:             _GotoState5Action,
	{_State32, LexErrorToken}:           _GotoState12Action,
	{_State32, LiteralType}:             _GotoState26Action,
	{_State32, AnonymousFuncExprType}:   _GotoState21Action,
	{_State32, ImplicitStructExprType}:  _GotoState25Action,
	{_State32, AtomExprType}:            _GotoState22Action,
	{_State32, AccessExprType}:          _GotoState18Action,
	{_State32, UnaryOpType}:             _GotoState32Action,
	{_State32, UnaryExprType}:           _GotoState70Action,
	{_State32, OptionalLabelDeclType}:   _GotoState69Action,
	{_State32, BlockExprType}:           _GotoState23Action,
	{_State35, RparenToken}:             _GotoState71Action,
	{_State36, RparenToken}:             _GotoState72Action,
	{_State37, IdentifierToken}:         _GotoState73Action,
	{_State38, BoolLiteralToken}:        _GotoState6Action,
	{_State38, IntegerLiteralToken}:     _GotoState10Action,
	{_State38, FloatLiteralToken}:       _GotoState7Action,
	{_State38, RuneLiteralToken}:        _GotoState15Action,
	{_State38, StringLiteralToken}:      _GotoState16Action,
	{_State38, LparenToken}:             _GotoState13Action,
	{_State38, IdentifierToken}:         _GotoState9Action,
	{_State38, FuncToken}:               _GotoState8Action,
	{_State38, LabelDeclToken}:          _GotoState11Action,
	{_State38, NotToken}:                _GotoState14Action,
	{_State38, SubToken}:                _GotoState17Action,
	{_State38, BitNegToken}:             _GotoState5Action,
	{_State38, LexErrorToken}:           _GotoState12Action,
	{_State38, LiteralType}:             _GotoState26Action,
	{_State38, AnonymousFuncExprType}:   _GotoState21Action,
	{_State38, ImplicitStructExprType}:  _GotoState25Action,
	{_State38, AtomExprType}:            _GotoState22Action,
	{_State38, AccessExprType}:          _GotoState18Action,
	{_State38, UnaryOpType}:             _GotoState32Action,
	{_State38, UnaryExprType}:           _GotoState31Action,
	{_State38, MulExprType}:             _GotoState27Action,
	{_State38, AddExprType}:             _GotoState19Action,
	{_State38, CmpExprType}:             _GotoState24Action,
	{_State38, AndExprType}:             _GotoState20Action,
	{_State38, OrExprType}:              _GotoState29Action,
	{_State38, SequenceExprType}:        _GotoState30Action,
	{_State38, OptionalLabelDeclType}:   _GotoState28Action,
	{_State38, BlockExprType}:           _GotoState23Action,
	{_State38, ExpressionType}:          _GotoState74Action,
	{_State39, ArgumentListType}:        _GotoState75Action,
	{_State44, BoolLiteralToken}:        _GotoState6Action,
	{_State44, IntegerLiteralToken}:     _GotoState10Action,
	{_State44, FloatLiteralToken}:       _GotoState7Action,
	{_State44, RuneLiteralToken}:        _GotoState15Action,
	{_State44, StringLiteralToken}:      _GotoState16Action,
	{_State44, LparenToken}:             _GotoState13Action,
	{_State44, IdentifierToken}:         _GotoState9Action,
	{_State44, FuncToken}:               _GotoState8Action,
	{_State44, LabelDeclToken}:          _GotoState11Action,
	{_State44, NotToken}:                _GotoState14Action,
	{_State44, SubToken}:                _GotoState17Action,
	{_State44, BitNegToken}:             _GotoState5Action,
	{_State44, LexErrorToken}:           _GotoState12Action,
	{_State44, LiteralType}:             _GotoState26Action,
	{_State44, AnonymousFuncExprType}:   _GotoState21Action,
	{_State44, ImplicitStructExprType}:  _GotoState25Action,
	{_State44, AtomExprType}:            _GotoState22Action,
	{_State44, AccessExprType}:          _GotoState18Action,
	{_State44, UnaryOpType}:             _GotoState32Action,
	{_State44, UnaryExprType}:           _GotoState31Action,
	{_State44, MulExprType}:             _GotoState76Action,
	{_State44, OptionalLabelDeclType}:   _GotoState69Action,
	{_State44, BlockExprType}:           _GotoState23Action,
	{_State45, BoolLiteralToken}:        _GotoState6Action,
	{_State45, IntegerLiteralToken}:     _GotoState10Action,
	{_State45, FloatLiteralToken}:       _GotoState7Action,
	{_State45, RuneLiteralToken}:        _GotoState15Action,
	{_State45, StringLiteralToken}:      _GotoState16Action,
	{_State45, LparenToken}:             _GotoState13Action,
	{_State45, IdentifierToken}:         _GotoState9Action,
	{_State45, FuncToken}:               _GotoState8Action,
	{_State45, LabelDeclToken}:          _GotoState11Action,
	{_State45, NotToken}:                _GotoState14Action,
	{_State45, SubToken}:                _GotoState17Action,
	{_State45, BitNegToken}:             _GotoState5Action,
	{_State45, LexErrorToken}:           _GotoState12Action,
	{_State45, LiteralType}:             _GotoState26Action,
	{_State45, AnonymousFuncExprType}:   _GotoState21Action,
	{_State45, ImplicitStructExprType}:  _GotoState25Action,
	{_State45, AtomExprType}:            _GotoState22Action,
	{_State45, AccessExprType}:          _GotoState18Action,
	{_State45, UnaryOpType}:             _GotoState32Action,
	{_State45, UnaryExprType}:           _GotoState31Action,
	{_State45, MulExprType}:             _GotoState27Action,
	{_State45, AddExprType}:             _GotoState19Action,
	{_State45, CmpExprType}:             _GotoState77Action,
	{_State45, OptionalLabelDeclType}:   _GotoState69Action,
	{_State45, BlockExprType}:           _GotoState23Action,
	{_State52, BoolLiteralToken}:        _GotoState6Action,
	{_State52, IntegerLiteralToken}:     _GotoState10Action,
	{_State52, FloatLiteralToken}:       _GotoState7Action,
	{_State52, RuneLiteralToken}:        _GotoState15Action,
	{_State52, StringLiteralToken}:      _GotoState16Action,
	{_State52, LparenToken}:             _GotoState13Action,
	{_State52, IdentifierToken}:         _GotoState9Action,
	{_State52, FuncToken}:               _GotoState8Action,
	{_State52, LabelDeclToken}:          _GotoState11Action,
	{_State52, NotToken}:                _GotoState14Action,
	{_State52, SubToken}:                _GotoState17Action,
	{_State52, BitNegToken}:             _GotoState5Action,
	{_State52, LexErrorToken}:           _GotoState12Action,
	{_State52, LiteralType}:             _GotoState26Action,
	{_State52, AnonymousFuncExprType}:   _GotoState21Action,
	{_State52, ImplicitStructExprType}:  _GotoState25Action,
	{_State52, AtomExprType}:            _GotoState22Action,
	{_State52, AccessExprType}:          _GotoState18Action,
	{_State52, UnaryOpType}:             _GotoState32Action,
	{_State52, UnaryExprType}:           _GotoState31Action,
	{_State52, MulExprType}:             _GotoState27Action,
	{_State52, AddExprType}:             _GotoState78Action,
	{_State52, OptionalLabelDeclType}:   _GotoState69Action,
	{_State52, BlockExprType}:           _GotoState23Action,
	{_State59, BoolLiteralToken}:        _GotoState6Action,
	{_State59, IntegerLiteralToken}:     _GotoState10Action,
	{_State59, FloatLiteralToken}:       _GotoState7Action,
	{_State59, RuneLiteralToken}:        _GotoState15Action,
	{_State59, StringLiteralToken}:      _GotoState16Action,
	{_State59, LparenToken}:             _GotoState13Action,
	{_State59, IdentifierToken}:         _GotoState9Action,
	{_State59, FuncToken}:               _GotoState8Action,
	{_State59, LabelDeclToken}:          _GotoState11Action,
	{_State59, NotToken}:                _GotoState14Action,
	{_State59, SubToken}:                _GotoState17Action,
	{_State59, BitNegToken}:             _GotoState5Action,
	{_State59, LexErrorToken}:           _GotoState12Action,
	{_State59, LiteralType}:             _GotoState26Action,
	{_State59, AnonymousFuncExprType}:   _GotoState21Action,
	{_State59, ImplicitStructExprType}:  _GotoState25Action,
	{_State59, AtomExprType}:            _GotoState22Action,
	{_State59, AccessExprType}:          _GotoState18Action,
	{_State59, UnaryOpType}:             _GotoState32Action,
	{_State59, UnaryExprType}:           _GotoState79Action,
	{_State59, OptionalLabelDeclType}:   _GotoState69Action,
	{_State59, BlockExprType}:           _GotoState23Action,
	{_State60, BoolLiteralToken}:        _GotoState6Action,
	{_State60, IntegerLiteralToken}:     _GotoState10Action,
	{_State60, FloatLiteralToken}:       _GotoState7Action,
	{_State60, RuneLiteralToken}:        _GotoState15Action,
	{_State60, StringLiteralToken}:      _GotoState16Action,
	{_State60, LparenToken}:             _GotoState13Action,
	{_State60, IdentifierToken}:         _GotoState9Action,
	{_State60, FuncToken}:               _GotoState8Action,
	{_State60, LabelDeclToken}:          _GotoState11Action,
	{_State60, NotToken}:                _GotoState14Action,
	{_State60, SubToken}:                _GotoState17Action,
	{_State60, BitNegToken}:             _GotoState5Action,
	{_State60, LexErrorToken}:           _GotoState12Action,
	{_State60, LiteralType}:             _GotoState26Action,
	{_State60, AnonymousFuncExprType}:   _GotoState21Action,
	{_State60, ImplicitStructExprType}:  _GotoState25Action,
	{_State60, AtomExprType}:            _GotoState22Action,
	{_State60, AccessExprType}:          _GotoState18Action,
	{_State60, UnaryOpType}:             _GotoState32Action,
	{_State60, UnaryExprType}:           _GotoState31Action,
	{_State60, MulExprType}:             _GotoState27Action,
	{_State60, AddExprType}:             _GotoState19Action,
	{_State60, CmpExprType}:             _GotoState24Action,
	{_State60, AndExprType}:             _GotoState20Action,
	{_State60, OrExprType}:              _GotoState29Action,
	{_State60, SequenceExprType}:        _GotoState81Action,
	{_State60, OptionalLabelDeclType}:   _GotoState69Action,
	{_State60, BlockExprType}:           _GotoState80Action,
	{_State61, BoolLiteralToken}:        _GotoState6Action,
	{_State61, IntegerLiteralToken}:     _GotoState10Action,
	{_State61, FloatLiteralToken}:       _GotoState7Action,
	{_State61, RuneLiteralToken}:        _GotoState15Action,
	{_State61, StringLiteralToken}:      _GotoState16Action,
	{_State61, LparenToken}:             _GotoState13Action,
	{_State61, IdentifierToken}:         _GotoState9Action,
	{_State61, FuncToken}:               _GotoState8Action,
	{_State61, LabelDeclToken}:          _GotoState11Action,
	{_State61, NotToken}:                _GotoState14Action,
	{_State61, SubToken}:                _GotoState17Action,
	{_State61, BitNegToken}:             _GotoState5Action,
	{_State61, LexErrorToken}:           _GotoState12Action,
	{_State61, LiteralType}:             _GotoState26Action,
	{_State61, AnonymousFuncExprType}:   _GotoState21Action,
	{_State61, ImplicitStructExprType}:  _GotoState25Action,
	{_State61, AtomExprType}:            _GotoState22Action,
	{_State61, AccessExprType}:          _GotoState18Action,
	{_State61, UnaryOpType}:             _GotoState32Action,
	{_State61, UnaryExprType}:           _GotoState31Action,
	{_State61, MulExprType}:             _GotoState27Action,
	{_State61, AddExprType}:             _GotoState19Action,
	{_State61, CmpExprType}:             _GotoState24Action,
	{_State61, AndExprType}:             _GotoState20Action,
	{_State61, OrExprType}:              _GotoState29Action,
	{_State61, SequenceExprType}:        _GotoState82Action,
	{_State61, OptionalLabelDeclType}:   _GotoState69Action,
	{_State61, BlockExprType}:           _GotoState23Action,
	{_State62, StatementsType}:          _GotoState83Action,
	{_State63, CaseToken}:               _GotoState84Action,
	{_State68, BoolLiteralToken}:        _GotoState6Action,
	{_State68, IntegerLiteralToken}:     _GotoState10Action,
	{_State68, FloatLiteralToken}:       _GotoState7Action,
	{_State68, RuneLiteralToken}:        _GotoState15Action,
	{_State68, StringLiteralToken}:      _GotoState16Action,
	{_State68, LparenToken}:             _GotoState13Action,
	{_State68, IdentifierToken}:         _GotoState9Action,
	{_State68, FuncToken}:               _GotoState8Action,
	{_State68, LabelDeclToken}:          _GotoState11Action,
	{_State68, NotToken}:                _GotoState14Action,
	{_State68, SubToken}:                _GotoState17Action,
	{_State68, BitNegToken}:             _GotoState5Action,
	{_State68, LexErrorToken}:           _GotoState12Action,
	{_State68, LiteralType}:             _GotoState26Action,
	{_State68, AnonymousFuncExprType}:   _GotoState21Action,
	{_State68, ImplicitStructExprType}:  _GotoState25Action,
	{_State68, AtomExprType}:            _GotoState22Action,
	{_State68, AccessExprType}:          _GotoState18Action,
	{_State68, UnaryOpType}:             _GotoState32Action,
	{_State68, UnaryExprType}:           _GotoState31Action,
	{_State68, MulExprType}:             _GotoState27Action,
	{_State68, AddExprType}:             _GotoState19Action,
	{_State68, CmpExprType}:             _GotoState24Action,
	{_State68, AndExprType}:             _GotoState85Action,
	{_State68, OptionalLabelDeclType}:   _GotoState69Action,
	{_State68, BlockExprType}:           _GotoState23Action,
	{_State69, LbraceToken}:             _GotoState62Action,
	{_State69, BlockBodyType}:           _GotoState64Action,
	{_State71, LbraceToken}:             _GotoState62Action,
	{_State71, BlockBodyType}:           _GotoState86Action,
	{_State74, RbracketToken}:           _GotoState87Action,
	{_State75, RparenToken}:             _GotoState88Action,
	{_State76, MulToken}:                _GotoState58Action,
	{_State76, DivToken}:                _GotoState56Action,
	{_State76, ModToken}:                _GotoState57Action,
	{_State76, BitAndToken}:             _GotoState53Action,
	{_State76, BitLshiftToken}:          _GotoState54Action,
	{_State76, BitRshiftToken}:          _GotoState55Action,
	{_State76, MulOpType}:               _GotoState59Action,
	{_State77, EqualToken}:              _GotoState46Action,
	{_State77, NotEqualToken}:           _GotoState51Action,
	{_State77, LessToken}:               _GotoState49Action,
	{_State77, LessOrEqualToken}:        _GotoState50Action,
	{_State77, GreaterToken}:            _GotoState47Action,
	{_State77, GreaterOrEqualToken}:     _GotoState48Action,
	{_State77, CmpOpType}:               _GotoState52Action,
	{_State78, AddToken}:                _GotoState40Action,
	{_State78, SubToken}:                _GotoState43Action,
	{_State78, BitXorToken}:             _GotoState42Action,
	{_State78, BitOrToken}:              _GotoState41Action,
	{_State78, AddOpType}:               _GotoState44Action,
	{_State81, LabelDeclToken}:          _GotoState11Action,
	{_State81, OptionalLabelDeclType}:   _GotoState69Action,
	{_State81, BlockExprType}:           _GotoState89Action,
	{_State82, LbraceToken}:             _GotoState62Action,
	{_State82, BlockBodyType}:           _GotoState90Action,
	{_State83, BoolLiteralToken}:        _GotoState6Action,
	{_State83, IntegerLiteralToken}:     _GotoState10Action,
	{_State83, FloatLiteralToken}:       _GotoState7Action,
	{_State83, RuneLiteralToken}:        _GotoState15Action,
	{_State83, StringLiteralToken}:      _GotoState16Action,
	{_State83, RbraceToken}:             _GotoState93Action,
	{_State83, LparenToken}:             _GotoState13Action,
	{_State83, IdentifierToken}:         _GotoState9Action,
	{_State83, ReturnToken}:             _GotoState94Action,
	{_State83, BreakToken}:              _GotoState91Action,
	{_State83, ContinueToken}:           _GotoState92Action,
	{_State83, FuncToken}:               _GotoState8Action,
	{_State83, LabelDeclToken}:          _GotoState11Action,
	{_State83, NotToken}:                _GotoState14Action,
	{_State83, SubToken}:                _GotoState17Action,
	{_State83, BitNegToken}:             _GotoState5Action,
	{_State83, LexErrorToken}:           _GotoState12Action,
	{_State83, LiteralType}:             _GotoState26Action,
	{_State83, AnonymousFuncExprType}:   _GotoState21Action,
	{_State83, ImplicitStructExprType}:  _GotoState25Action,
	{_State83, AtomExprType}:            _GotoState22Action,
	{_State83, AccessExprType}:          _GotoState95Action,
	{_State83, UnaryOpType}:             _GotoState32Action,
	{_State83, UnaryExprType}:           _GotoState31Action,
	{_State83, MulExprType}:             _GotoState27Action,
	{_State83, AddExprType}:             _GotoState19Action,
	{_State83, CmpExprType}:             _GotoState24Action,
	{_State83, AndExprType}:             _GotoState20Action,
	{_State83, OrExprType}:              _GotoState29Action,
	{_State83, SequenceExprType}:        _GotoState30Action,
	{_State83, JumpTypeType}:            _GotoState97Action,
	{_State83, StatementBodyType}:       _GotoState99Action,
	{_State83, StatementType}:           _GotoState98Action,
	{_State83, OptionalLabelDeclType}:   _GotoState28Action,
	{_State83, BlockExprType}:           _GotoState23Action,
	{_State83, ExpressionType}:          _GotoState96Action,
	{_State84, DefaultToken}:            _GotoState100Action,
	{_State85, AndToken}:                _GotoState45Action,
	{_State90, ElseToken}:               _GotoState101Action,
	{_State95, LparenToken}:             _GotoState39Action,
	{_State95, LbracketToken}:           _GotoState38Action,
	{_State95, DotToken}:                _GotoState37Action,
	{_State95, AddAssignToken}:          _GotoState102Action,
	{_State95, SubAssignToken}:          _GotoState113Action,
	{_State95, MulAssignToken}:          _GotoState112Action,
	{_State95, DivAssignToken}:          _GotoState110Action,
	{_State95, ModAssignToken}:          _GotoState111Action,
	{_State95, AddOneAssignToken}:       _GotoState103Action,
	{_State95, SubOneAssignToken}:       _GotoState114Action,
	{_State95, BitNegAssignToken}:       _GotoState106Action,
	{_State95, BitAndAssignToken}:       _GotoState104Action,
	{_State95, BitOrAssignToken}:        _GotoState107Action,
	{_State95, BitXorAssignToken}:       _GotoState109Action,
	{_State95, BitLshiftAssignToken}:    _GotoState105Action,
	{_State95, BitRshiftAssignToken}:    _GotoState108Action,
	{_State95, OpOneAssignType}:         _GotoState116Action,
	{_State95, BinaryOpAssignType}:      _GotoState115Action,
	{_State97, JumpLabelToken}:          _GotoState117Action,
	{_State97, OptionalJumpLabelType}:   _GotoState118Action,
	{_State99, NewlinesToken}:           _GotoState119Action,
	{_State99, SemicolonToken}:          _GotoState120Action,
	{_State101, LbraceToken}:            _GotoState62Action,
	{_State101, IfToken}:                _GotoState61Action,
	{_State101, BlockBodyType}:          _GotoState121Action,
	{_State101, IfExprType}:             _GotoState122Action,
	{_State115, BoolLiteralToken}:       _GotoState6Action,
	{_State115, IntegerLiteralToken}:    _GotoState10Action,
	{_State115, FloatLiteralToken}:      _GotoState7Action,
	{_State115, RuneLiteralToken}:       _GotoState15Action,
	{_State115, StringLiteralToken}:     _GotoState16Action,
	{_State115, LparenToken}:            _GotoState13Action,
	{_State115, IdentifierToken}:        _GotoState9Action,
	{_State115, FuncToken}:              _GotoState8Action,
	{_State115, LabelDeclToken}:         _GotoState11Action,
	{_State115, NotToken}:               _GotoState14Action,
	{_State115, SubToken}:               _GotoState17Action,
	{_State115, BitNegToken}:            _GotoState5Action,
	{_State115, LexErrorToken}:          _GotoState12Action,
	{_State115, LiteralType}:            _GotoState26Action,
	{_State115, AnonymousFuncExprType}:  _GotoState21Action,
	{_State115, ImplicitStructExprType}: _GotoState25Action,
	{_State115, AtomExprType}:           _GotoState22Action,
	{_State115, AccessExprType}:         _GotoState18Action,
	{_State115, UnaryOpType}:            _GotoState32Action,
	{_State115, UnaryExprType}:          _GotoState31Action,
	{_State115, MulExprType}:            _GotoState27Action,
	{_State115, AddExprType}:            _GotoState19Action,
	{_State115, CmpExprType}:            _GotoState24Action,
	{_State115, AndExprType}:            _GotoState20Action,
	{_State115, OrExprType}:             _GotoState29Action,
	{_State115, SequenceExprType}:       _GotoState30Action,
	{_State115, OptionalLabelDeclType}:  _GotoState28Action,
	{_State115, BlockExprType}:          _GotoState23Action,
	{_State115, ExpressionType}:         _GotoState123Action,
	{_State118, BoolLiteralToken}:       _GotoState6Action,
	{_State118, IntegerLiteralToken}:    _GotoState10Action,
	{_State118, FloatLiteralToken}:      _GotoState7Action,
	{_State118, RuneLiteralToken}:       _GotoState15Action,
	{_State118, StringLiteralToken}:     _GotoState16Action,
	{_State118, LparenToken}:            _GotoState13Action,
	{_State118, IdentifierToken}:        _GotoState9Action,
	{_State118, FuncToken}:              _GotoState8Action,
	{_State118, LabelDeclToken}:         _GotoState11Action,
	{_State118, NotToken}:               _GotoState14Action,
	{_State118, SubToken}:               _GotoState17Action,
	{_State118, BitNegToken}:            _GotoState5Action,
	{_State118, LexErrorToken}:          _GotoState12Action,
	{_State118, LiteralType}:            _GotoState26Action,
	{_State118, AnonymousFuncExprType}:  _GotoState21Action,
	{_State118, ImplicitStructExprType}: _GotoState25Action,
	{_State118, AtomExprType}:           _GotoState22Action,
	{_State118, AccessExprType}:         _GotoState18Action,
	{_State118, UnaryOpType}:            _GotoState32Action,
	{_State118, UnaryExprType}:          _GotoState31Action,
	{_State118, MulExprType}:            _GotoState27Action,
	{_State118, AddExprType}:            _GotoState19Action,
	{_State118, CmpExprType}:            _GotoState24Action,
	{_State118, AndExprType}:            _GotoState20Action,
	{_State118, OrExprType}:             _GotoState29Action,
	{_State118, SequenceExprType}:       _GotoState30Action,
	{_State118, OptionalExpressionType}: _GotoState125Action,
	{_State118, OptionalLabelDeclType}:  _GotoState28Action,
	{_State118, BlockExprType}:          _GotoState23Action,
	{_State118, ExpressionType}:         _GotoState124Action,
	{_State1, _WildcardMarker}:          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State5, _WildcardMarker}:          _ReduceBitNegToUnaryOpAction,
	{_State6, _WildcardMarker}:          _ReduceBoolLiteralToLiteralAction,
	{_State7, _WildcardMarker}:          _ReduceFloatLiteralToLiteralAction,
	{_State9, _WildcardMarker}:          _ReduceIdentifierToAtomExprAction,
	{_State10, _WildcardMarker}:         _ReduceIntegerLiteralToLiteralAction,
	{_State11, _WildcardMarker}:         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State12, _WildcardMarker}:         _ReduceLexErrorToAtomExprAction,
	{_State13, RparenToken}:             _ReduceToArgumentListAction,
	{_State14, _WildcardMarker}:         _ReduceNotToUnaryOpAction,
	{_State15, _WildcardMarker}:         _ReduceRuneLiteralToLiteralAction,
	{_State16, _WildcardMarker}:         _ReduceStringLiteralToLiteralAction,
	{_State17, _WildcardMarker}:         _ReduceSubToUnaryOpAction,
	{_State18, _WildcardMarker}:         _ReduceAccessExprToUnaryExprAction,
	{_State19, _WildcardMarker}:         _ReduceAddExprToCmpExprAction,
	{_State20, _WildcardMarker}:         _ReduceAndExprToOrExprAction,
	{_State21, _WildcardMarker}:         _ReduceAnonymousFuncExprToAtomExprAction,
	{_State22, _WildcardMarker}:         _ReduceAtomExprToAccessExprAction,
	{_State23, _WildcardMarker}:         _ReduceBlockExprToAtomExprAction,
	{_State24, _WildcardMarker}:         _ReduceCmpExprToAndExprAction,
	{_State25, _WildcardMarker}:         _ReduceImplicitStructExprToAtomExprAction,
	{_State26, _WildcardMarker}:         _ReduceLiteralToAtomExprAction,
	{_State27, _WildcardMarker}:         _ReduceMulExprToAddExprAction,
	{_State29, _EndMarker}:              _ReduceToSequenceExprAction,
	{_State30, _EndMarker}:              _ReduceSequenceExprToExpressionAction,
	{_State31, _WildcardMarker}:         _ReduceUnaryExprToMulExprAction,
	{_State32, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State33, _EndMarker}:              _ReduceCommentToLexInternalTokensAction,
	{_State34, _EndMarker}:              _ReduceSpacesToLexInternalTokensAction,
	{_State38, _WildcardMarker}:         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State39, RparenToken}:             _ReduceToArgumentListAction,
	{_State40, _WildcardMarker}:         _ReduceAddToAddOpAction,
	{_State41, _WildcardMarker}:         _ReduceBitOrToAddOpAction,
	{_State42, _WildcardMarker}:         _ReduceBitXorToAddOpAction,
	{_State43, _WildcardMarker}:         _ReduceSubToAddOpAction,
	{_State44, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State45, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State46, _WildcardMarker}:         _ReduceEqualToCmpOpAction,
	{_State47, _WildcardMarker}:         _ReduceGreaterToCmpOpAction,
	{_State48, _WildcardMarker}:         _ReduceGreaterOrEqualToCmpOpAction,
	{_State49, _WildcardMarker}:         _ReduceLessToCmpOpAction,
	{_State50, _WildcardMarker}:         _ReduceLessOrEqualToCmpOpAction,
	{_State51, _WildcardMarker}:         _ReduceNotEqualToCmpOpAction,
	{_State52, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State53, _WildcardMarker}:         _ReduceBitAndToMulOpAction,
	{_State54, _WildcardMarker}:         _ReduceBitLshiftToMulOpAction,
	{_State55, _WildcardMarker}:         _ReduceBitRshiftToMulOpAction,
	{_State56, _WildcardMarker}:         _ReduceDivToMulOpAction,
	{_State57, _WildcardMarker}:         _ReduceModToMulOpAction,
	{_State58, _WildcardMarker}:         _ReduceMulToMulOpAction,
	{_State59, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State60, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State61, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State62, _WildcardMarker}:         _ReduceEmptyListToStatementsAction,
	{_State64, _EndMarker}:              _ReduceToBlockExprAction,
	{_State65, _EndMarker}:              _ReduceIfExprToExpressionAction,
	{_State66, _EndMarker}:              _ReduceLoopExprToExpressionAction,
	{_State67, _EndMarker}:              _ReduceMatchExprToExpressionAction,
	{_State68, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State70, _WildcardMarker}:         _ReduceOpToUnaryExprAction,
	{_State72, _WildcardMarker}:         _ReduceToImplicitStructExprAction,
	{_State73, _WildcardMarker}:         _ReduceAccessToAccessExprAction,
	{_State76, _WildcardMarker}:         _ReduceOpToAddExprAction,
	{_State77, _WildcardMarker}:         _ReduceOpToAndExprAction,
	{_State78, _WildcardMarker}:         _ReduceOpToCmpExprAction,
	{_State79, _WildcardMarker}:         _ReduceOpToMulExprAction,
	{_State80, _WildcardMarker}:         _ReduceBlockExprToAtomExprAction,
	{_State80, _EndMarker}:              _ReduceInfiniteToLoopExprAction,
	{_State81, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State83, _WildcardMarker}:         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State85, _WildcardMarker}:         _ReduceOpToOrExprAction,
	{_State86, _WildcardMarker}:         _ReduceToAnonymousFuncExprAction,
	{_State87, _WildcardMarker}:         _ReduceIndexToAccessExprAction,
	{_State88, _WildcardMarker}:         _ReduceCallToAccessExprAction,
	{_State89, _EndMarker}:              _ReduceWhileToLoopExprAction,
	{_State90, _WildcardMarker}:         _ReduceNoElseToIfExprAction,
	{_State91, _WildcardMarker}:         _ReduceBreakToJumpTypeAction,
	{_State92, _WildcardMarker}:         _ReduceContinueToJumpTypeAction,
	{_State93, _EndMarker}:              _ReduceToBlockBodyAction,
	{_State94, _WildcardMarker}:         _ReduceReturnToJumpTypeAction,
	{_State95, _WildcardMarker}:         _ReduceAccessExprToUnaryExprAction,
	{_State96, _WildcardMarker}:         _ReduceExpressionToStatementBodyAction,
	{_State97, _WildcardMarker}:         _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State98, _WildcardMarker}:         _ReduceAddToStatementsAction,
	{_State100, _EndMarker}:             _ReduceToMatchExprAction,
	{_State102, _WildcardMarker}:        _ReduceAddAssignToBinaryOpAssignAction,
	{_State103, _WildcardMarker}:        _ReduceAddOneAssignToOpOneAssignAction,
	{_State104, _WildcardMarker}:        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State105, _WildcardMarker}:        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State106, _WildcardMarker}:        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State107, _WildcardMarker}:        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State108, _WildcardMarker}:        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State109, _WildcardMarker}:        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State110, _WildcardMarker}:        _ReduceDivAssignToBinaryOpAssignAction,
	{_State111, _WildcardMarker}:        _ReduceModAssignToBinaryOpAssignAction,
	{_State112, _WildcardMarker}:        _ReduceMulAssignToBinaryOpAssignAction,
	{_State113, _WildcardMarker}:        _ReduceSubAssignToBinaryOpAssignAction,
	{_State114, _WildcardMarker}:        _ReduceSubOneAssignToOpOneAssignAction,
	{_State115, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State116, _WildcardMarker}:        _ReduceOpOneAssignToStatementBodyAction,
	{_State117, _WildcardMarker}:        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State118, NewlinesToken}:          _ReduceNilToOptionalExpressionAction,
	{_State118, SemicolonToken}:         _ReduceNilToOptionalExpressionAction,
	{_State118, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State119, _WildcardMarker}:        _ReduceImplicitToStatementAction,
	{_State120, _WildcardMarker}:        _ReduceExplicitToStatementAction,
	{_State121, _EndMarker}:             _ReduceIfElseToIfExprAction,
	{_State122, _EndMarker}:             _ReduceMultiIfElseToIfExprAction,
	{_State123, _WildcardMarker}:        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State124, _WildcardMarker}:        _ReduceExpressionToOptionalExpressionAction,
	{_State125, _WildcardMarker}:        _ReduceJumpToStatementBodyAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 30
      optional_label_decl -> State 28
      block_expr -> State 23
      expression -> State 3

  State 2:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 34
      COMMENT -> State 33
      lex_internal_tokens -> State 4

  State 3:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 4:
    Kernel Items:
      #accept: ^ lex_internal_tokens., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      unary_op: BIT_NEG., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      literal: BOOL_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN RPAREN block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 35

  State 9:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      implicit_struct_expr: LPAREN.argument_list RPAREN
    Reduce:
      RPAREN -> [argument_list]
    Goto:
      argument_list -> State 36

  State 14:
    Kernel Items:
      unary_op: NOT., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      unary_op: SUB., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LPAREN argument_list RPAREN
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 39
      LBRACKET -> State 38
      DOT -> State 37

  State 19:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 40
      SUB -> State 43
      BIT_XOR -> State 42
      BIT_OR -> State 41
      add_op -> State 44

  State 20:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 45

  State 21:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 46
      NOT_EQUAL -> State 51
      LESS -> State 49
      LESS_OR_EQUAL -> State 50
      GREATER -> State 47
      GREATER_OR_EQUAL -> State 48
      cmp_op -> State 52

  State 25:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 58
      DIV -> State 56
      MOD -> State 57
      BIT_AND -> State 53
      BIT_LSHIFT -> State 54
      BIT_RSHIFT -> State 55
      mul_op -> State 59

  State 28:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.match_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      IF -> State 61
      MATCH -> State 63
      FOR -> State 60
      block_body -> State 64
      if_expr -> State 65
      match_expr -> State 67
      loop_expr -> State 66

  State 29:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 68

  State 30:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      mul_expr: unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      unary_expr: unary_op.unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 70
      optional_label_decl -> State 69
      block_expr -> State 23

  State 33:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.RPAREN block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 71

  State 36:
    Kernel Items:
      implicit_struct_expr: LPAREN argument_list.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 72

  State 37:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 73

  State 38:
    Kernel Items:
      access_expr: access_expr LBRACKET.expression RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 30
      optional_label_decl -> State 28
      block_expr -> State 23
      expression -> State 74

  State 39:
    Kernel Items:
      access_expr: access_expr LPAREN.argument_list RPAREN
    Reduce:
      RPAREN -> [argument_list]
    Goto:
      argument_list -> State 75

  State 40:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 76
      optional_label_decl -> State 69
      block_expr -> State 23

  State 45:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 77
      optional_label_decl -> State 69
      block_expr -> State 23

  State 46:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 78
      optional_label_decl -> State 69
      block_expr -> State 23

  State 53:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      mul_expr: mul_expr mul_op.unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 79
      optional_label_decl -> State 69
      block_expr -> State 23

  State 60:
    Kernel Items:
      loop_expr: FOR.block_expr
      loop_expr: FOR.sequence_expr block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 81
      optional_label_decl -> State 69
      block_expr -> State 80

  State 61:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 82
      optional_label_decl -> State 69
      block_expr -> State 23

  State 62:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 83

  State 63:
    Kernel Items:
      match_expr: MATCH.CASE DEFAULT
    Reduce:
      (nil)
    Goto:
      CASE -> State 84

  State 64:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      expression: optional_label_decl match_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 85
      optional_label_decl -> State 69
      block_expr -> State 23

  State 69:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 64

  State 70:
    Kernel Items:
      unary_expr: unary_op unary_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN RPAREN.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 86

  State 72:
    Kernel Items:
      implicit_struct_expr: LPAREN argument_list RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 87

  State 75:
    Kernel Items:
      access_expr: access_expr LPAREN argument_list.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 88

  State 76:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 58
      DIV -> State 56
      MOD -> State 57
      BIT_AND -> State 53
      BIT_LSHIFT -> State 54
      BIT_RSHIFT -> State 55
      mul_op -> State 59

  State 77:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 46
      NOT_EQUAL -> State 51
      LESS -> State 49
      LESS_OR_EQUAL -> State 50
      GREATER -> State 47
      GREATER_OR_EQUAL -> State 48
      cmp_op -> State 52

  State 78:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 40
      SUB -> State 43
      BIT_XOR -> State 42
      BIT_OR -> State 41
      add_op -> State 44

  State 79:
    Kernel Items:
      mul_expr: mul_expr mul_op unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 11
      optional_label_decl -> State 69
      block_expr -> State 89

  State 82:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      block_body -> State 90

  State 83:
    Kernel Items:
      statements: statements.statement
      block_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      RBRACE -> State 93
      LPAREN -> State 13
      IDENTIFIER -> State 9
      RETURN -> State 94
      BREAK -> State 91
      CONTINUE -> State 92
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 95
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 30
      jump_type -> State 97
      statement_body -> State 99
      statement -> State 98
      optional_label_decl -> State 28
      block_expr -> State 23
      expression -> State 96

  State 84:
    Kernel Items:
      match_expr: MATCH CASE.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 100

  State 85:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 45

  State 86:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN RPAREN block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      access_expr: access_expr LPAREN argument_list RPAREN., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 101

  State 91:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LPAREN argument_list RPAREN
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
      statement_body: access_expr.op_one_assign
      statement_body: access_expr.binary_op_assign expression
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 39
      LBRACKET -> State 38
      DOT -> State 37
      ADD_ASSIGN -> State 102
      SUB_ASSIGN -> State 113
      MUL_ASSIGN -> State 112
      DIV_ASSIGN -> State 110
      MOD_ASSIGN -> State 111
      ADD_ONE_ASSIGN -> State 103
      SUB_ONE_ASSIGN -> State 114
      BIT_NEG_ASSIGN -> State 106
      BIT_AND_ASSIGN -> State 104
      BIT_OR_ASSIGN -> State 107
      BIT_XOR_ASSIGN -> State 109
      BIT_LSHIFT_ASSIGN -> State 105
      BIT_RSHIFT_ASSIGN -> State 108
      op_one_assign -> State 116
      binary_op_assign -> State 115

  State 96:
    Kernel Items:
      statement_body: expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 117
      optional_jump_label -> State 118

  State 98:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 119
      SEMICOLON -> State 120

  State 100:
    Kernel Items:
      match_expr: MATCH CASE DEFAULT., $
    Reduce:
      $ -> [match_expr]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 62
      IF -> State 61
      block_body -> State 121
      if_expr -> State 122

  State 102:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 30
      optional_label_decl -> State 28
      block_expr -> State 23
      expression -> State 123

  State 116:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression]
      SEMICOLON -> [optional_expression]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 10
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 15
      STRING_LITERAL -> State 16
      LPAREN -> State 13
      IDENTIFIER -> State 9
      FUNC -> State 8
      LABEL_DECL -> State 11
      NOT -> State 14
      SUB -> State 17
      BIT_NEG -> State 5
      LEX_ERROR -> State 12
      literal -> State 26
      anonymous_func_expr -> State 21
      implicit_struct_expr -> State 25
      atom_expr -> State 22
      access_expr -> State 18
      unary_op -> State 32
      unary_expr -> State 31
      mul_expr -> State 27
      add_expr -> State 19
      cmp_expr -> State 24
      and_expr -> State 20
      or_expr -> State 29
      sequence_expr -> State 30
      optional_expression -> State 125
      optional_label_decl -> State 28
      block_expr -> State 23
      expression -> State 124

  State 119:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

Number of states: 125
Number of shift actions: 453
Number of reduce actions: 113
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
