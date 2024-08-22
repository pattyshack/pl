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
	ForToken             = SymbolId(277)
	ReturnToken          = SymbolId(278)
	BreakToken           = SymbolId(279)
	ContinueToken        = SymbolId(280)
	LabelDeclToken       = SymbolId(281)
	JumpLabelToken       = SymbolId(282)
	AddAssignToken       = SymbolId(283)
	SubAssignToken       = SymbolId(284)
	MulAssignToken       = SymbolId(285)
	DivAssignToken       = SymbolId(286)
	ModAssignToken       = SymbolId(287)
	AddOneAssignToken    = SymbolId(288)
	SubOneAssignToken    = SymbolId(289)
	BitNegAssignToken    = SymbolId(290)
	BitAndAssignToken    = SymbolId(291)
	BitOrAssignToken     = SymbolId(292)
	BitXorAssignToken    = SymbolId(293)
	BitLshiftAssignToken = SymbolId(294)
	BitRshiftAssignToken = SymbolId(295)
	NotToken             = SymbolId(296)
	AndToken             = SymbolId(297)
	OrToken              = SymbolId(298)
	AddToken             = SymbolId(299)
	SubToken             = SymbolId(300)
	MulToken             = SymbolId(301)
	DivToken             = SymbolId(302)
	ModToken             = SymbolId(303)
	BitNegToken          = SymbolId(304)
	BitAndToken          = SymbolId(305)
	BitXorToken          = SymbolId(306)
	BitOrToken           = SymbolId(307)
	BitLshiftToken       = SymbolId(308)
	BitRshiftToken       = SymbolId(309)
	EqualToken           = SymbolId(310)
	NotEqualToken        = SymbolId(311)
	LessToken            = SymbolId(312)
	LessOrEqualToken     = SymbolId(313)
	GreaterToken         = SymbolId(314)
	GreaterOrEqualToken  = SymbolId(315)
	LexErrorToken        = SymbolId(316)
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
	// 61:2: literal -> BOOL_LITERAL: ...
	BoolLiteralToLiteral(BoolLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 62:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 63:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 64:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 65:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 69:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 70:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 71:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 72:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 75:0: argument_list -> ...
	ToArgumentList() (*GenericSymbol, error)

	// 78:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 79:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 80:2: access_expr -> call: ...
	CallToAccessExpr(AccessExpr_ *GenericSymbol, Lparen_ *GenericSymbol, ArgumentList_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 82:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 85:2: unary_op -> NOT: ...
	NotToUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 86:2: unary_op -> BIT_NEG: ...
	BitNegToUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 87:2: unary_op -> SUB: ...
	SubToUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 90:2: unary_expr -> access_expr: ...
	AccessExprToUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 91:2: unary_expr -> op: ...
	OpToUnaryExpr(UnaryOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 94:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 95:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 96:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 97:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 98:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 99:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 102:2: mul_expr -> unary_expr: ...
	UnaryExprToMulExpr(UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 103:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 106:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 107:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 108:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 109:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 113:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 116:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 117:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 121:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 124:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 125:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 135:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 138:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 139:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 142:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 143:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 146:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 148:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 156:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 158:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 159:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 162:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 163:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 164:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 165:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: statement_body -> expression: ...
	ExpressionToStatementBody(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 178:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 182:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 188:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 189:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 196:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 199:14: match_expr -> ...
	ToMatchExpr(Match_ *GenericSymbol, Case_ *GenericSymbol) (*GenericSymbol, error)

	// 202:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: expression -> match_expr: ...
	MatchExprToExpression(OptionalLabelDecl_ *GenericSymbol, MatchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 212:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: lex_internal_tokens -> COMMENT: ...
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
	case ForToken:
		return "FOR"
	case ReturnToken:
		return "RETURN"
	case BreakToken:
		return "BREAK"
	case ContinueToken:
		return "CONTINUE"
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

	LiteralType            = SymbolId(317)
	AtomExprType           = SymbolId(318)
	ArgumentListType       = SymbolId(319)
	AccessExprType         = SymbolId(320)
	UnaryOpType            = SymbolId(321)
	UnaryExprType          = SymbolId(322)
	MulOpType              = SymbolId(323)
	MulExprType            = SymbolId(324)
	AddOpType              = SymbolId(325)
	AddExprType            = SymbolId(326)
	CmpOpType              = SymbolId(327)
	CmpExprType            = SymbolId(328)
	AndExprType            = SymbolId(329)
	OrExprType             = SymbolId(330)
	SequenceExprType       = SymbolId(331)
	OptionalJumpLabelType  = SymbolId(332)
	OptionalExpressionType = SymbolId(333)
	JumpTypeType           = SymbolId(334)
	OpOneAssignType        = SymbolId(335)
	BinaryOpAssignType     = SymbolId(336)
	StatementBodyType      = SymbolId(337)
	StatementType          = SymbolId(338)
	StatementsType         = SymbolId(339)
	OptionalLabelDeclType  = SymbolId(340)
	BlockBodyType          = SymbolId(341)
	BlockExprType          = SymbolId(342)
	IfExprType             = SymbolId(343)
	MatchExprType          = SymbolId(344)
	LoopExprType           = SymbolId(345)
	ExpressionType         = SymbolId(346)
	LexInternalTokensType  = SymbolId(347)
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
	_ReduceLiteralToAtomExpr               = _ReduceType(6)
	_ReduceIdentifierToAtomExpr            = _ReduceType(7)
	_ReduceBlockExprToAtomExpr             = _ReduceType(8)
	_ReduceLexErrorToAtomExpr              = _ReduceType(9)
	_ReduceToArgumentList                  = _ReduceType(10)
	_ReduceAtomExprToAccessExpr            = _ReduceType(11)
	_ReduceAccessToAccessExpr              = _ReduceType(12)
	_ReduceCallToAccessExpr                = _ReduceType(13)
	_ReduceIndexToAccessExpr               = _ReduceType(14)
	_ReduceNotToUnaryOp                    = _ReduceType(15)
	_ReduceBitNegToUnaryOp                 = _ReduceType(16)
	_ReduceSubToUnaryOp                    = _ReduceType(17)
	_ReduceAccessExprToUnaryExpr           = _ReduceType(18)
	_ReduceOpToUnaryExpr                   = _ReduceType(19)
	_ReduceMulToMulOp                      = _ReduceType(20)
	_ReduceDivToMulOp                      = _ReduceType(21)
	_ReduceModToMulOp                      = _ReduceType(22)
	_ReduceBitAndToMulOp                   = _ReduceType(23)
	_ReduceBitLshiftToMulOp                = _ReduceType(24)
	_ReduceBitRshiftToMulOp                = _ReduceType(25)
	_ReduceUnaryExprToMulExpr              = _ReduceType(26)
	_ReduceOpToMulExpr                     = _ReduceType(27)
	_ReduceAddToAddOp                      = _ReduceType(28)
	_ReduceSubToAddOp                      = _ReduceType(29)
	_ReduceBitOrToAddOp                    = _ReduceType(30)
	_ReduceBitXorToAddOp                   = _ReduceType(31)
	_ReduceMulExprToAddExpr                = _ReduceType(32)
	_ReduceOpToAddExpr                     = _ReduceType(33)
	_ReduceEqualToCmpOp                    = _ReduceType(34)
	_ReduceNotEqualToCmpOp                 = _ReduceType(35)
	_ReduceLessToCmpOp                     = _ReduceType(36)
	_ReduceLessOrEqualToCmpOp              = _ReduceType(37)
	_ReduceGreaterToCmpOp                  = _ReduceType(38)
	_ReduceGreaterOrEqualToCmpOp           = _ReduceType(39)
	_ReduceAddExprToCmpExpr                = _ReduceType(40)
	_ReduceOpToCmpExpr                     = _ReduceType(41)
	_ReduceCmpExprToAndExpr                = _ReduceType(42)
	_ReduceOpToAndExpr                     = _ReduceType(43)
	_ReduceAndExprToOrExpr                 = _ReduceType(44)
	_ReduceOpToOrExpr                      = _ReduceType(45)
	_ReduceToSequenceExpr                  = _ReduceType(46)
	_ReduceJumpLabelToOptionalJumpLabel    = _ReduceType(47)
	_ReduceUnlabelledToOptionalJumpLabel   = _ReduceType(48)
	_ReduceExpressionToOptionalExpression  = _ReduceType(49)
	_ReduceNilToOptionalExpression         = _ReduceType(50)
	_ReduceReturnToJumpType                = _ReduceType(51)
	_ReduceBreakToJumpType                 = _ReduceType(52)
	_ReduceContinueToJumpType              = _ReduceType(53)
	_ReduceAddOneAssignToOpOneAssign       = _ReduceType(54)
	_ReduceSubOneAssignToOpOneAssign       = _ReduceType(55)
	_ReduceAddAssignToBinaryOpAssign       = _ReduceType(56)
	_ReduceSubAssignToBinaryOpAssign       = _ReduceType(57)
	_ReduceMulAssignToBinaryOpAssign       = _ReduceType(58)
	_ReduceDivAssignToBinaryOpAssign       = _ReduceType(59)
	_ReduceModAssignToBinaryOpAssign       = _ReduceType(60)
	_ReduceBitNegAssignToBinaryOpAssign    = _ReduceType(61)
	_ReduceBitAndAssignToBinaryOpAssign    = _ReduceType(62)
	_ReduceBitOrAssignToBinaryOpAssign     = _ReduceType(63)
	_ReduceBitXorAssignToBinaryOpAssign    = _ReduceType(64)
	_ReduceBitLshiftAssignToBinaryOpAssign = _ReduceType(65)
	_ReduceBitRshiftAssignToBinaryOpAssign = _ReduceType(66)
	_ReduceExpressionToStatementBody       = _ReduceType(67)
	_ReduceJumpToStatementBody             = _ReduceType(68)
	_ReduceOpOneAssignToStatementBody      = _ReduceType(69)
	_ReduceBinaryOpAssignToStatementBody   = _ReduceType(70)
	_ReduceImplicitToStatement             = _ReduceType(71)
	_ReduceExplicitToStatement             = _ReduceType(72)
	_ReduceEmptyListToStatements           = _ReduceType(73)
	_ReduceAddToStatements                 = _ReduceType(74)
	_ReduceLabelDeclToOptionalLabelDecl    = _ReduceType(75)
	_ReduceUnlabelledToOptionalLabelDecl   = _ReduceType(76)
	_ReduceToBlockBody                     = _ReduceType(77)
	_ReduceToBlockExpr                     = _ReduceType(78)
	_ReduceNoElseToIfExpr                  = _ReduceType(79)
	_ReduceIfElseToIfExpr                  = _ReduceType(80)
	_ReduceMultiIfElseToIfExpr             = _ReduceType(81)
	_ReduceToMatchExpr                     = _ReduceType(82)
	_ReduceInfiniteToLoopExpr              = _ReduceType(83)
	_ReduceWhileToLoopExpr                 = _ReduceType(84)
	_ReduceSequenceExprToExpression        = _ReduceType(85)
	_ReduceIfExprToExpression              = _ReduceType(86)
	_ReduceMatchExprToExpression           = _ReduceType(87)
	_ReduceLoopExprToExpression            = _ReduceType(88)
	_ReduceSpacesToLexInternalTokens       = _ReduceType(89)
	_ReduceCommentToLexInternalTokens      = _ReduceType(90)
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
	case _ReduceLiteralToAtomExpr:
		return "LiteralToAtomExpr"
	case _ReduceIdentifierToAtomExpr:
		return "IdentifierToAtomExpr"
	case _ReduceBlockExprToAtomExpr:
		return "BlockExprToAtomExpr"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, BoolLiteralToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, SemicolonToken, IdentifierToken, IfToken, ElseToken, MatchToken, CaseToken, ForToken, ReturnToken, BreakToken, ContinueToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = MatchExprType
		symbol.Generic_, err = reducer.ToMatchExpr(args[0].Generic_, args[1].Generic_)
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
	_ReduceBoolLiteralToLiteralAction            = &_Action{_ReduceAction, 0, _ReduceBoolLiteralToLiteral}
	_ReduceIntegerLiteralToLiteralAction         = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction           = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction            = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction          = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceLiteralToAtomExprAction               = &_Action{_ReduceAction, 0, _ReduceLiteralToAtomExpr}
	_ReduceIdentifierToAtomExprAction            = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction             = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
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
	{_State1, IntegerLiteralToken}:      _GotoState9Action,
	{_State1, FloatLiteralToken}:        _GotoState7Action,
	{_State1, RuneLiteralToken}:         _GotoState13Action,
	{_State1, StringLiteralToken}:       _GotoState14Action,
	{_State1, IdentifierToken}:          _GotoState8Action,
	{_State1, LabelDeclToken}:           _GotoState10Action,
	{_State1, NotToken}:                 _GotoState12Action,
	{_State1, SubToken}:                 _GotoState15Action,
	{_State1, BitNegToken}:              _GotoState5Action,
	{_State1, LexErrorToken}:            _GotoState11Action,
	{_State1, LiteralType}:              _GotoState22Action,
	{_State1, AtomExprType}:             _GotoState19Action,
	{_State1, AccessExprType}:           _GotoState16Action,
	{_State1, UnaryOpType}:              _GotoState28Action,
	{_State1, UnaryExprType}:            _GotoState27Action,
	{_State1, MulExprType}:              _GotoState23Action,
	{_State1, AddExprType}:              _GotoState17Action,
	{_State1, CmpExprType}:              _GotoState21Action,
	{_State1, AndExprType}:              _GotoState18Action,
	{_State1, OrExprType}:               _GotoState25Action,
	{_State1, SequenceExprType}:         _GotoState26Action,
	{_State1, OptionalLabelDeclType}:    _GotoState24Action,
	{_State1, BlockExprType}:            _GotoState20Action,
	{_State1, ExpressionType}:           _GotoState3Action,
	{_State2, SpacesToken}:              _GotoState30Action,
	{_State2, CommentToken}:             _GotoState29Action,
	{_State2, LexInternalTokensType}:    _GotoState4Action,
	{_State16, LparenToken}:             _GotoState33Action,
	{_State16, LbracketToken}:           _GotoState32Action,
	{_State16, DotToken}:                _GotoState31Action,
	{_State17, AddToken}:                _GotoState34Action,
	{_State17, SubToken}:                _GotoState37Action,
	{_State17, BitXorToken}:             _GotoState36Action,
	{_State17, BitOrToken}:              _GotoState35Action,
	{_State17, AddOpType}:               _GotoState38Action,
	{_State18, AndToken}:                _GotoState39Action,
	{_State21, EqualToken}:              _GotoState40Action,
	{_State21, NotEqualToken}:           _GotoState45Action,
	{_State21, LessToken}:               _GotoState43Action,
	{_State21, LessOrEqualToken}:        _GotoState44Action,
	{_State21, GreaterToken}:            _GotoState41Action,
	{_State21, GreaterOrEqualToken}:     _GotoState42Action,
	{_State21, CmpOpType}:               _GotoState46Action,
	{_State23, MulToken}:                _GotoState52Action,
	{_State23, DivToken}:                _GotoState50Action,
	{_State23, ModToken}:                _GotoState51Action,
	{_State23, BitAndToken}:             _GotoState47Action,
	{_State23, BitLshiftToken}:          _GotoState48Action,
	{_State23, BitRshiftToken}:          _GotoState49Action,
	{_State23, MulOpType}:               _GotoState53Action,
	{_State24, LbraceToken}:             _GotoState56Action,
	{_State24, IfToken}:                 _GotoState55Action,
	{_State24, MatchToken}:              _GotoState57Action,
	{_State24, ForToken}:                _GotoState54Action,
	{_State24, BlockBodyType}:           _GotoState58Action,
	{_State24, IfExprType}:              _GotoState59Action,
	{_State24, MatchExprType}:           _GotoState61Action,
	{_State24, LoopExprType}:            _GotoState60Action,
	{_State25, OrToken}:                 _GotoState62Action,
	{_State28, BoolLiteralToken}:        _GotoState6Action,
	{_State28, IntegerLiteralToken}:     _GotoState9Action,
	{_State28, FloatLiteralToken}:       _GotoState7Action,
	{_State28, RuneLiteralToken}:        _GotoState13Action,
	{_State28, StringLiteralToken}:      _GotoState14Action,
	{_State28, IdentifierToken}:         _GotoState8Action,
	{_State28, LabelDeclToken}:          _GotoState10Action,
	{_State28, NotToken}:                _GotoState12Action,
	{_State28, SubToken}:                _GotoState15Action,
	{_State28, BitNegToken}:             _GotoState5Action,
	{_State28, LexErrorToken}:           _GotoState11Action,
	{_State28, LiteralType}:             _GotoState22Action,
	{_State28, AtomExprType}:            _GotoState19Action,
	{_State28, AccessExprType}:          _GotoState16Action,
	{_State28, UnaryOpType}:             _GotoState28Action,
	{_State28, UnaryExprType}:           _GotoState64Action,
	{_State28, OptionalLabelDeclType}:   _GotoState63Action,
	{_State28, BlockExprType}:           _GotoState20Action,
	{_State31, IdentifierToken}:         _GotoState65Action,
	{_State32, BoolLiteralToken}:        _GotoState6Action,
	{_State32, IntegerLiteralToken}:     _GotoState9Action,
	{_State32, FloatLiteralToken}:       _GotoState7Action,
	{_State32, RuneLiteralToken}:        _GotoState13Action,
	{_State32, StringLiteralToken}:      _GotoState14Action,
	{_State32, IdentifierToken}:         _GotoState8Action,
	{_State32, LabelDeclToken}:          _GotoState10Action,
	{_State32, NotToken}:                _GotoState12Action,
	{_State32, SubToken}:                _GotoState15Action,
	{_State32, BitNegToken}:             _GotoState5Action,
	{_State32, LexErrorToken}:           _GotoState11Action,
	{_State32, LiteralType}:             _GotoState22Action,
	{_State32, AtomExprType}:            _GotoState19Action,
	{_State32, AccessExprType}:          _GotoState16Action,
	{_State32, UnaryOpType}:             _GotoState28Action,
	{_State32, UnaryExprType}:           _GotoState27Action,
	{_State32, MulExprType}:             _GotoState23Action,
	{_State32, AddExprType}:             _GotoState17Action,
	{_State32, CmpExprType}:             _GotoState21Action,
	{_State32, AndExprType}:             _GotoState18Action,
	{_State32, OrExprType}:              _GotoState25Action,
	{_State32, SequenceExprType}:        _GotoState26Action,
	{_State32, OptionalLabelDeclType}:   _GotoState24Action,
	{_State32, BlockExprType}:           _GotoState20Action,
	{_State32, ExpressionType}:          _GotoState66Action,
	{_State33, ArgumentListType}:        _GotoState67Action,
	{_State38, BoolLiteralToken}:        _GotoState6Action,
	{_State38, IntegerLiteralToken}:     _GotoState9Action,
	{_State38, FloatLiteralToken}:       _GotoState7Action,
	{_State38, RuneLiteralToken}:        _GotoState13Action,
	{_State38, StringLiteralToken}:      _GotoState14Action,
	{_State38, IdentifierToken}:         _GotoState8Action,
	{_State38, LabelDeclToken}:          _GotoState10Action,
	{_State38, NotToken}:                _GotoState12Action,
	{_State38, SubToken}:                _GotoState15Action,
	{_State38, BitNegToken}:             _GotoState5Action,
	{_State38, LexErrorToken}:           _GotoState11Action,
	{_State38, LiteralType}:             _GotoState22Action,
	{_State38, AtomExprType}:            _GotoState19Action,
	{_State38, AccessExprType}:          _GotoState16Action,
	{_State38, UnaryOpType}:             _GotoState28Action,
	{_State38, UnaryExprType}:           _GotoState27Action,
	{_State38, MulExprType}:             _GotoState68Action,
	{_State38, OptionalLabelDeclType}:   _GotoState63Action,
	{_State38, BlockExprType}:           _GotoState20Action,
	{_State39, BoolLiteralToken}:        _GotoState6Action,
	{_State39, IntegerLiteralToken}:     _GotoState9Action,
	{_State39, FloatLiteralToken}:       _GotoState7Action,
	{_State39, RuneLiteralToken}:        _GotoState13Action,
	{_State39, StringLiteralToken}:      _GotoState14Action,
	{_State39, IdentifierToken}:         _GotoState8Action,
	{_State39, LabelDeclToken}:          _GotoState10Action,
	{_State39, NotToken}:                _GotoState12Action,
	{_State39, SubToken}:                _GotoState15Action,
	{_State39, BitNegToken}:             _GotoState5Action,
	{_State39, LexErrorToken}:           _GotoState11Action,
	{_State39, LiteralType}:             _GotoState22Action,
	{_State39, AtomExprType}:            _GotoState19Action,
	{_State39, AccessExprType}:          _GotoState16Action,
	{_State39, UnaryOpType}:             _GotoState28Action,
	{_State39, UnaryExprType}:           _GotoState27Action,
	{_State39, MulExprType}:             _GotoState23Action,
	{_State39, AddExprType}:             _GotoState17Action,
	{_State39, CmpExprType}:             _GotoState69Action,
	{_State39, OptionalLabelDeclType}:   _GotoState63Action,
	{_State39, BlockExprType}:           _GotoState20Action,
	{_State46, BoolLiteralToken}:        _GotoState6Action,
	{_State46, IntegerLiteralToken}:     _GotoState9Action,
	{_State46, FloatLiteralToken}:       _GotoState7Action,
	{_State46, RuneLiteralToken}:        _GotoState13Action,
	{_State46, StringLiteralToken}:      _GotoState14Action,
	{_State46, IdentifierToken}:         _GotoState8Action,
	{_State46, LabelDeclToken}:          _GotoState10Action,
	{_State46, NotToken}:                _GotoState12Action,
	{_State46, SubToken}:                _GotoState15Action,
	{_State46, BitNegToken}:             _GotoState5Action,
	{_State46, LexErrorToken}:           _GotoState11Action,
	{_State46, LiteralType}:             _GotoState22Action,
	{_State46, AtomExprType}:            _GotoState19Action,
	{_State46, AccessExprType}:          _GotoState16Action,
	{_State46, UnaryOpType}:             _GotoState28Action,
	{_State46, UnaryExprType}:           _GotoState27Action,
	{_State46, MulExprType}:             _GotoState23Action,
	{_State46, AddExprType}:             _GotoState70Action,
	{_State46, OptionalLabelDeclType}:   _GotoState63Action,
	{_State46, BlockExprType}:           _GotoState20Action,
	{_State53, BoolLiteralToken}:        _GotoState6Action,
	{_State53, IntegerLiteralToken}:     _GotoState9Action,
	{_State53, FloatLiteralToken}:       _GotoState7Action,
	{_State53, RuneLiteralToken}:        _GotoState13Action,
	{_State53, StringLiteralToken}:      _GotoState14Action,
	{_State53, IdentifierToken}:         _GotoState8Action,
	{_State53, LabelDeclToken}:          _GotoState10Action,
	{_State53, NotToken}:                _GotoState12Action,
	{_State53, SubToken}:                _GotoState15Action,
	{_State53, BitNegToken}:             _GotoState5Action,
	{_State53, LexErrorToken}:           _GotoState11Action,
	{_State53, LiteralType}:             _GotoState22Action,
	{_State53, AtomExprType}:            _GotoState19Action,
	{_State53, AccessExprType}:          _GotoState16Action,
	{_State53, UnaryOpType}:             _GotoState28Action,
	{_State53, UnaryExprType}:           _GotoState71Action,
	{_State53, OptionalLabelDeclType}:   _GotoState63Action,
	{_State53, BlockExprType}:           _GotoState20Action,
	{_State54, BoolLiteralToken}:        _GotoState6Action,
	{_State54, IntegerLiteralToken}:     _GotoState9Action,
	{_State54, FloatLiteralToken}:       _GotoState7Action,
	{_State54, RuneLiteralToken}:        _GotoState13Action,
	{_State54, StringLiteralToken}:      _GotoState14Action,
	{_State54, IdentifierToken}:         _GotoState8Action,
	{_State54, LabelDeclToken}:          _GotoState10Action,
	{_State54, NotToken}:                _GotoState12Action,
	{_State54, SubToken}:                _GotoState15Action,
	{_State54, BitNegToken}:             _GotoState5Action,
	{_State54, LexErrorToken}:           _GotoState11Action,
	{_State54, LiteralType}:             _GotoState22Action,
	{_State54, AtomExprType}:            _GotoState19Action,
	{_State54, AccessExprType}:          _GotoState16Action,
	{_State54, UnaryOpType}:             _GotoState28Action,
	{_State54, UnaryExprType}:           _GotoState27Action,
	{_State54, MulExprType}:             _GotoState23Action,
	{_State54, AddExprType}:             _GotoState17Action,
	{_State54, CmpExprType}:             _GotoState21Action,
	{_State54, AndExprType}:             _GotoState18Action,
	{_State54, OrExprType}:              _GotoState25Action,
	{_State54, SequenceExprType}:        _GotoState73Action,
	{_State54, OptionalLabelDeclType}:   _GotoState63Action,
	{_State54, BlockExprType}:           _GotoState72Action,
	{_State55, BoolLiteralToken}:        _GotoState6Action,
	{_State55, IntegerLiteralToken}:     _GotoState9Action,
	{_State55, FloatLiteralToken}:       _GotoState7Action,
	{_State55, RuneLiteralToken}:        _GotoState13Action,
	{_State55, StringLiteralToken}:      _GotoState14Action,
	{_State55, IdentifierToken}:         _GotoState8Action,
	{_State55, LabelDeclToken}:          _GotoState10Action,
	{_State55, NotToken}:                _GotoState12Action,
	{_State55, SubToken}:                _GotoState15Action,
	{_State55, BitNegToken}:             _GotoState5Action,
	{_State55, LexErrorToken}:           _GotoState11Action,
	{_State55, LiteralType}:             _GotoState22Action,
	{_State55, AtomExprType}:            _GotoState19Action,
	{_State55, AccessExprType}:          _GotoState16Action,
	{_State55, UnaryOpType}:             _GotoState28Action,
	{_State55, UnaryExprType}:           _GotoState27Action,
	{_State55, MulExprType}:             _GotoState23Action,
	{_State55, AddExprType}:             _GotoState17Action,
	{_State55, CmpExprType}:             _GotoState21Action,
	{_State55, AndExprType}:             _GotoState18Action,
	{_State55, OrExprType}:              _GotoState25Action,
	{_State55, SequenceExprType}:        _GotoState74Action,
	{_State55, OptionalLabelDeclType}:   _GotoState63Action,
	{_State55, BlockExprType}:           _GotoState20Action,
	{_State56, StatementsType}:          _GotoState75Action,
	{_State57, CaseToken}:               _GotoState76Action,
	{_State62, BoolLiteralToken}:        _GotoState6Action,
	{_State62, IntegerLiteralToken}:     _GotoState9Action,
	{_State62, FloatLiteralToken}:       _GotoState7Action,
	{_State62, RuneLiteralToken}:        _GotoState13Action,
	{_State62, StringLiteralToken}:      _GotoState14Action,
	{_State62, IdentifierToken}:         _GotoState8Action,
	{_State62, LabelDeclToken}:          _GotoState10Action,
	{_State62, NotToken}:                _GotoState12Action,
	{_State62, SubToken}:                _GotoState15Action,
	{_State62, BitNegToken}:             _GotoState5Action,
	{_State62, LexErrorToken}:           _GotoState11Action,
	{_State62, LiteralType}:             _GotoState22Action,
	{_State62, AtomExprType}:            _GotoState19Action,
	{_State62, AccessExprType}:          _GotoState16Action,
	{_State62, UnaryOpType}:             _GotoState28Action,
	{_State62, UnaryExprType}:           _GotoState27Action,
	{_State62, MulExprType}:             _GotoState23Action,
	{_State62, AddExprType}:             _GotoState17Action,
	{_State62, CmpExprType}:             _GotoState21Action,
	{_State62, AndExprType}:             _GotoState77Action,
	{_State62, OptionalLabelDeclType}:   _GotoState63Action,
	{_State62, BlockExprType}:           _GotoState20Action,
	{_State63, LbraceToken}:             _GotoState56Action,
	{_State63, BlockBodyType}:           _GotoState58Action,
	{_State66, RbracketToken}:           _GotoState78Action,
	{_State67, RparenToken}:             _GotoState79Action,
	{_State68, MulToken}:                _GotoState52Action,
	{_State68, DivToken}:                _GotoState50Action,
	{_State68, ModToken}:                _GotoState51Action,
	{_State68, BitAndToken}:             _GotoState47Action,
	{_State68, BitLshiftToken}:          _GotoState48Action,
	{_State68, BitRshiftToken}:          _GotoState49Action,
	{_State68, MulOpType}:               _GotoState53Action,
	{_State69, EqualToken}:              _GotoState40Action,
	{_State69, NotEqualToken}:           _GotoState45Action,
	{_State69, LessToken}:               _GotoState43Action,
	{_State69, LessOrEqualToken}:        _GotoState44Action,
	{_State69, GreaterToken}:            _GotoState41Action,
	{_State69, GreaterOrEqualToken}:     _GotoState42Action,
	{_State69, CmpOpType}:               _GotoState46Action,
	{_State70, AddToken}:                _GotoState34Action,
	{_State70, SubToken}:                _GotoState37Action,
	{_State70, BitXorToken}:             _GotoState36Action,
	{_State70, BitOrToken}:              _GotoState35Action,
	{_State70, AddOpType}:               _GotoState38Action,
	{_State73, LabelDeclToken}:          _GotoState10Action,
	{_State73, OptionalLabelDeclType}:   _GotoState63Action,
	{_State73, BlockExprType}:           _GotoState80Action,
	{_State74, LbraceToken}:             _GotoState56Action,
	{_State74, BlockBodyType}:           _GotoState81Action,
	{_State75, BoolLiteralToken}:        _GotoState6Action,
	{_State75, IntegerLiteralToken}:     _GotoState9Action,
	{_State75, FloatLiteralToken}:       _GotoState7Action,
	{_State75, RuneLiteralToken}:        _GotoState13Action,
	{_State75, StringLiteralToken}:      _GotoState14Action,
	{_State75, RbraceToken}:             _GotoState84Action,
	{_State75, IdentifierToken}:         _GotoState8Action,
	{_State75, ReturnToken}:             _GotoState85Action,
	{_State75, BreakToken}:              _GotoState82Action,
	{_State75, ContinueToken}:           _GotoState83Action,
	{_State75, LabelDeclToken}:          _GotoState10Action,
	{_State75, NotToken}:                _GotoState12Action,
	{_State75, SubToken}:                _GotoState15Action,
	{_State75, BitNegToken}:             _GotoState5Action,
	{_State75, LexErrorToken}:           _GotoState11Action,
	{_State75, LiteralType}:             _GotoState22Action,
	{_State75, AtomExprType}:            _GotoState19Action,
	{_State75, AccessExprType}:          _GotoState86Action,
	{_State75, UnaryOpType}:             _GotoState28Action,
	{_State75, UnaryExprType}:           _GotoState27Action,
	{_State75, MulExprType}:             _GotoState23Action,
	{_State75, AddExprType}:             _GotoState17Action,
	{_State75, CmpExprType}:             _GotoState21Action,
	{_State75, AndExprType}:             _GotoState18Action,
	{_State75, OrExprType}:              _GotoState25Action,
	{_State75, SequenceExprType}:        _GotoState26Action,
	{_State75, JumpTypeType}:            _GotoState88Action,
	{_State75, StatementBodyType}:       _GotoState90Action,
	{_State75, StatementType}:           _GotoState89Action,
	{_State75, OptionalLabelDeclType}:   _GotoState24Action,
	{_State75, BlockExprType}:           _GotoState20Action,
	{_State75, ExpressionType}:          _GotoState87Action,
	{_State77, AndToken}:                _GotoState39Action,
	{_State81, ElseToken}:               _GotoState91Action,
	{_State86, LparenToken}:             _GotoState33Action,
	{_State86, LbracketToken}:           _GotoState32Action,
	{_State86, DotToken}:                _GotoState31Action,
	{_State86, AddAssignToken}:          _GotoState92Action,
	{_State86, SubAssignToken}:          _GotoState103Action,
	{_State86, MulAssignToken}:          _GotoState102Action,
	{_State86, DivAssignToken}:          _GotoState100Action,
	{_State86, ModAssignToken}:          _GotoState101Action,
	{_State86, AddOneAssignToken}:       _GotoState93Action,
	{_State86, SubOneAssignToken}:       _GotoState104Action,
	{_State86, BitNegAssignToken}:       _GotoState96Action,
	{_State86, BitAndAssignToken}:       _GotoState94Action,
	{_State86, BitOrAssignToken}:        _GotoState97Action,
	{_State86, BitXorAssignToken}:       _GotoState99Action,
	{_State86, BitLshiftAssignToken}:    _GotoState95Action,
	{_State86, BitRshiftAssignToken}:    _GotoState98Action,
	{_State86, OpOneAssignType}:         _GotoState106Action,
	{_State86, BinaryOpAssignType}:      _GotoState105Action,
	{_State88, JumpLabelToken}:          _GotoState107Action,
	{_State88, OptionalJumpLabelType}:   _GotoState108Action,
	{_State90, NewlinesToken}:           _GotoState109Action,
	{_State90, SemicolonToken}:          _GotoState110Action,
	{_State91, LbraceToken}:             _GotoState56Action,
	{_State91, IfToken}:                 _GotoState55Action,
	{_State91, BlockBodyType}:           _GotoState111Action,
	{_State91, IfExprType}:              _GotoState112Action,
	{_State105, BoolLiteralToken}:       _GotoState6Action,
	{_State105, IntegerLiteralToken}:    _GotoState9Action,
	{_State105, FloatLiteralToken}:      _GotoState7Action,
	{_State105, RuneLiteralToken}:       _GotoState13Action,
	{_State105, StringLiteralToken}:     _GotoState14Action,
	{_State105, IdentifierToken}:        _GotoState8Action,
	{_State105, LabelDeclToken}:         _GotoState10Action,
	{_State105, NotToken}:               _GotoState12Action,
	{_State105, SubToken}:               _GotoState15Action,
	{_State105, BitNegToken}:            _GotoState5Action,
	{_State105, LexErrorToken}:          _GotoState11Action,
	{_State105, LiteralType}:            _GotoState22Action,
	{_State105, AtomExprType}:           _GotoState19Action,
	{_State105, AccessExprType}:         _GotoState16Action,
	{_State105, UnaryOpType}:            _GotoState28Action,
	{_State105, UnaryExprType}:          _GotoState27Action,
	{_State105, MulExprType}:            _GotoState23Action,
	{_State105, AddExprType}:            _GotoState17Action,
	{_State105, CmpExprType}:            _GotoState21Action,
	{_State105, AndExprType}:            _GotoState18Action,
	{_State105, OrExprType}:             _GotoState25Action,
	{_State105, SequenceExprType}:       _GotoState26Action,
	{_State105, OptionalLabelDeclType}:  _GotoState24Action,
	{_State105, BlockExprType}:          _GotoState20Action,
	{_State105, ExpressionType}:         _GotoState113Action,
	{_State108, BoolLiteralToken}:       _GotoState6Action,
	{_State108, IntegerLiteralToken}:    _GotoState9Action,
	{_State108, FloatLiteralToken}:      _GotoState7Action,
	{_State108, RuneLiteralToken}:       _GotoState13Action,
	{_State108, StringLiteralToken}:     _GotoState14Action,
	{_State108, IdentifierToken}:        _GotoState8Action,
	{_State108, LabelDeclToken}:         _GotoState10Action,
	{_State108, NotToken}:               _GotoState12Action,
	{_State108, SubToken}:               _GotoState15Action,
	{_State108, BitNegToken}:            _GotoState5Action,
	{_State108, LexErrorToken}:          _GotoState11Action,
	{_State108, LiteralType}:            _GotoState22Action,
	{_State108, AtomExprType}:           _GotoState19Action,
	{_State108, AccessExprType}:         _GotoState16Action,
	{_State108, UnaryOpType}:            _GotoState28Action,
	{_State108, UnaryExprType}:          _GotoState27Action,
	{_State108, MulExprType}:            _GotoState23Action,
	{_State108, AddExprType}:            _GotoState17Action,
	{_State108, CmpExprType}:            _GotoState21Action,
	{_State108, AndExprType}:            _GotoState18Action,
	{_State108, OrExprType}:             _GotoState25Action,
	{_State108, SequenceExprType}:       _GotoState26Action,
	{_State108, OptionalExpressionType}: _GotoState115Action,
	{_State108, OptionalLabelDeclType}:  _GotoState24Action,
	{_State108, BlockExprType}:          _GotoState20Action,
	{_State108, ExpressionType}:         _GotoState114Action,
	{_State1, _WildcardMarker}:          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State5, _WildcardMarker}:          _ReduceBitNegToUnaryOpAction,
	{_State6, _WildcardMarker}:          _ReduceBoolLiteralToLiteralAction,
	{_State7, _WildcardMarker}:          _ReduceFloatLiteralToLiteralAction,
	{_State8, _WildcardMarker}:          _ReduceIdentifierToAtomExprAction,
	{_State9, _WildcardMarker}:          _ReduceIntegerLiteralToLiteralAction,
	{_State10, _WildcardMarker}:         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State11, _WildcardMarker}:         _ReduceLexErrorToAtomExprAction,
	{_State12, _WildcardMarker}:         _ReduceNotToUnaryOpAction,
	{_State13, _WildcardMarker}:         _ReduceRuneLiteralToLiteralAction,
	{_State14, _WildcardMarker}:         _ReduceStringLiteralToLiteralAction,
	{_State15, _WildcardMarker}:         _ReduceSubToUnaryOpAction,
	{_State16, _WildcardMarker}:         _ReduceAccessExprToUnaryExprAction,
	{_State17, _WildcardMarker}:         _ReduceAddExprToCmpExprAction,
	{_State18, _WildcardMarker}:         _ReduceAndExprToOrExprAction,
	{_State19, _WildcardMarker}:         _ReduceAtomExprToAccessExprAction,
	{_State20, _WildcardMarker}:         _ReduceBlockExprToAtomExprAction,
	{_State21, _WildcardMarker}:         _ReduceCmpExprToAndExprAction,
	{_State22, _WildcardMarker}:         _ReduceLiteralToAtomExprAction,
	{_State23, _WildcardMarker}:         _ReduceMulExprToAddExprAction,
	{_State25, _EndMarker}:              _ReduceToSequenceExprAction,
	{_State26, _EndMarker}:              _ReduceSequenceExprToExpressionAction,
	{_State27, _WildcardMarker}:         _ReduceUnaryExprToMulExprAction,
	{_State28, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State29, _EndMarker}:              _ReduceCommentToLexInternalTokensAction,
	{_State30, _EndMarker}:              _ReduceSpacesToLexInternalTokensAction,
	{_State32, _WildcardMarker}:         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State33, RparenToken}:             _ReduceToArgumentListAction,
	{_State34, _WildcardMarker}:         _ReduceAddToAddOpAction,
	{_State35, _WildcardMarker}:         _ReduceBitOrToAddOpAction,
	{_State36, _WildcardMarker}:         _ReduceBitXorToAddOpAction,
	{_State37, _WildcardMarker}:         _ReduceSubToAddOpAction,
	{_State38, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State39, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State40, _WildcardMarker}:         _ReduceEqualToCmpOpAction,
	{_State41, _WildcardMarker}:         _ReduceGreaterToCmpOpAction,
	{_State42, _WildcardMarker}:         _ReduceGreaterOrEqualToCmpOpAction,
	{_State43, _WildcardMarker}:         _ReduceLessToCmpOpAction,
	{_State44, _WildcardMarker}:         _ReduceLessOrEqualToCmpOpAction,
	{_State45, _WildcardMarker}:         _ReduceNotEqualToCmpOpAction,
	{_State46, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State47, _WildcardMarker}:         _ReduceBitAndToMulOpAction,
	{_State48, _WildcardMarker}:         _ReduceBitLshiftToMulOpAction,
	{_State49, _WildcardMarker}:         _ReduceBitRshiftToMulOpAction,
	{_State50, _WildcardMarker}:         _ReduceDivToMulOpAction,
	{_State51, _WildcardMarker}:         _ReduceModToMulOpAction,
	{_State52, _WildcardMarker}:         _ReduceMulToMulOpAction,
	{_State53, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State54, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State55, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State56, _WildcardMarker}:         _ReduceEmptyListToStatementsAction,
	{_State58, _EndMarker}:              _ReduceToBlockExprAction,
	{_State59, _EndMarker}:              _ReduceIfExprToExpressionAction,
	{_State60, _EndMarker}:              _ReduceLoopExprToExpressionAction,
	{_State61, _EndMarker}:              _ReduceMatchExprToExpressionAction,
	{_State62, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State64, _WildcardMarker}:         _ReduceOpToUnaryExprAction,
	{_State65, _WildcardMarker}:         _ReduceAccessToAccessExprAction,
	{_State68, _WildcardMarker}:         _ReduceOpToAddExprAction,
	{_State69, _WildcardMarker}:         _ReduceOpToAndExprAction,
	{_State70, _WildcardMarker}:         _ReduceOpToCmpExprAction,
	{_State71, _WildcardMarker}:         _ReduceOpToMulExprAction,
	{_State72, _WildcardMarker}:         _ReduceBlockExprToAtomExprAction,
	{_State72, _EndMarker}:              _ReduceInfiniteToLoopExprAction,
	{_State73, LbraceToken}:             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State75, _WildcardMarker}:         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State76, _EndMarker}:              _ReduceToMatchExprAction,
	{_State77, _WildcardMarker}:         _ReduceOpToOrExprAction,
	{_State78, _WildcardMarker}:         _ReduceIndexToAccessExprAction,
	{_State79, _WildcardMarker}:         _ReduceCallToAccessExprAction,
	{_State80, _EndMarker}:              _ReduceWhileToLoopExprAction,
	{_State81, _WildcardMarker}:         _ReduceNoElseToIfExprAction,
	{_State82, _WildcardMarker}:         _ReduceBreakToJumpTypeAction,
	{_State83, _WildcardMarker}:         _ReduceContinueToJumpTypeAction,
	{_State84, _EndMarker}:              _ReduceToBlockBodyAction,
	{_State85, _WildcardMarker}:         _ReduceReturnToJumpTypeAction,
	{_State86, _WildcardMarker}:         _ReduceAccessExprToUnaryExprAction,
	{_State87, _WildcardMarker}:         _ReduceExpressionToStatementBodyAction,
	{_State88, _WildcardMarker}:         _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State89, _WildcardMarker}:         _ReduceAddToStatementsAction,
	{_State92, _WildcardMarker}:         _ReduceAddAssignToBinaryOpAssignAction,
	{_State93, _WildcardMarker}:         _ReduceAddOneAssignToOpOneAssignAction,
	{_State94, _WildcardMarker}:         _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State95, _WildcardMarker}:         _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State96, _WildcardMarker}:         _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State97, _WildcardMarker}:         _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State98, _WildcardMarker}:         _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State99, _WildcardMarker}:         _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State100, _WildcardMarker}:        _ReduceDivAssignToBinaryOpAssignAction,
	{_State101, _WildcardMarker}:        _ReduceModAssignToBinaryOpAssignAction,
	{_State102, _WildcardMarker}:        _ReduceMulAssignToBinaryOpAssignAction,
	{_State103, _WildcardMarker}:        _ReduceSubAssignToBinaryOpAssignAction,
	{_State104, _WildcardMarker}:        _ReduceSubOneAssignToOpOneAssignAction,
	{_State105, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State106, _WildcardMarker}:        _ReduceOpOneAssignToStatementBodyAction,
	{_State107, _WildcardMarker}:        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State108, NewlinesToken}:          _ReduceNilToOptionalExpressionAction,
	{_State108, SemicolonToken}:         _ReduceNilToOptionalExpressionAction,
	{_State108, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State109, _WildcardMarker}:        _ReduceImplicitToStatementAction,
	{_State110, _WildcardMarker}:        _ReduceExplicitToStatementAction,
	{_State111, _EndMarker}:             _ReduceIfElseToIfExprAction,
	{_State112, _EndMarker}:             _ReduceMultiIfElseToIfExprAction,
	{_State113, _WildcardMarker}:        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State114, _WildcardMarker}:        _ReduceExpressionToOptionalExpressionAction,
	{_State115, _WildcardMarker}:        _ReduceJumpToStatementBodyAction,
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
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 26
      optional_label_decl -> State 24
      block_expr -> State 20
      expression -> State 3

  State 2:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 30
      COMMENT -> State 29
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
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      unary_op: NOT., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      unary_op: SUB., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LPAREN argument_list RPAREN
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 33
      LBRACKET -> State 32
      DOT -> State 31

  State 17:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 34
      SUB -> State 37
      BIT_XOR -> State 36
      BIT_OR -> State 35
      add_op -> State 38

  State 18:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 39

  State 19:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 40
      NOT_EQUAL -> State 45
      LESS -> State 43
      LESS_OR_EQUAL -> State 44
      GREATER -> State 41
      GREATER_OR_EQUAL -> State 42
      cmp_op -> State 46

  State 22:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 52
      DIV -> State 50
      MOD -> State 51
      BIT_AND -> State 47
      BIT_LSHIFT -> State 48
      BIT_RSHIFT -> State 49
      mul_op -> State 53

  State 24:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.match_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      IF -> State 55
      MATCH -> State 57
      FOR -> State 54
      block_body -> State 58
      if_expr -> State 59
      match_expr -> State 61
      loop_expr -> State 60

  State 25:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 62

  State 26:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      mul_expr: unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      unary_expr: unary_op.unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 64
      optional_label_decl -> State 63
      block_expr -> State 20

  State 29:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 65

  State 32:
    Kernel Items:
      access_expr: access_expr LBRACKET.expression RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 26
      optional_label_decl -> State 24
      block_expr -> State 20
      expression -> State 66

  State 33:
    Kernel Items:
      access_expr: access_expr LPAREN.argument_list RPAREN
    Reduce:
      RPAREN -> [argument_list]
    Goto:
      argument_list -> State 67

  State 34:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 68
      optional_label_decl -> State 63
      block_expr -> State 20

  State 39:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 69
      optional_label_decl -> State 63
      block_expr -> State 20

  State 40:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 70
      optional_label_decl -> State 63
      block_expr -> State 20

  State 47:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      mul_expr: mul_expr mul_op.unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 71
      optional_label_decl -> State 63
      block_expr -> State 20

  State 54:
    Kernel Items:
      loop_expr: FOR.block_expr
      loop_expr: FOR.sequence_expr block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 73
      optional_label_decl -> State 63
      block_expr -> State 72

  State 55:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 74
      optional_label_decl -> State 63
      block_expr -> State 20

  State 56:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 75

  State 57:
    Kernel Items:
      match_expr: MATCH.CASE
    Reduce:
      (nil)
    Goto:
      CASE -> State 76

  State 58:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      expression: optional_label_decl match_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 77
      optional_label_decl -> State 63
      block_expr -> State 20

  State 63:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      block_body -> State 58

  State 64:
    Kernel Items:
      unary_expr: unary_op unary_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 78

  State 67:
    Kernel Items:
      access_expr: access_expr LPAREN argument_list.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 79

  State 68:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 52
      DIV -> State 50
      MOD -> State 51
      BIT_AND -> State 47
      BIT_LSHIFT -> State 48
      BIT_RSHIFT -> State 49
      mul_op -> State 53

  State 69:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 40
      NOT_EQUAL -> State 45
      LESS -> State 43
      LESS_OR_EQUAL -> State 44
      GREATER -> State 41
      GREATER_OR_EQUAL -> State 42
      cmp_op -> State 46

  State 70:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 34
      SUB -> State 37
      BIT_XOR -> State 36
      BIT_OR -> State 35
      add_op -> State 38

  State 71:
    Kernel Items:
      mul_expr: mul_expr mul_op unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 10
      optional_label_decl -> State 63
      block_expr -> State 80

  State 74:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      block_body -> State 81

  State 75:
    Kernel Items:
      statements: statements.statement
      block_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      RBRACE -> State 84
      IDENTIFIER -> State 8
      RETURN -> State 85
      BREAK -> State 82
      CONTINUE -> State 83
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 86
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 26
      jump_type -> State 88
      statement_body -> State 90
      statement -> State 89
      optional_label_decl -> State 24
      block_expr -> State 20
      expression -> State 87

  State 76:
    Kernel Items:
      match_expr: MATCH CASE., $
    Reduce:
      $ -> [match_expr]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 39

  State 78:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      access_expr: access_expr LPAREN argument_list RPAREN., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 91

  State 82:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 86:
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
      LPAREN -> State 33
      LBRACKET -> State 32
      DOT -> State 31
      ADD_ASSIGN -> State 92
      SUB_ASSIGN -> State 103
      MUL_ASSIGN -> State 102
      DIV_ASSIGN -> State 100
      MOD_ASSIGN -> State 101
      ADD_ONE_ASSIGN -> State 93
      SUB_ONE_ASSIGN -> State 104
      BIT_NEG_ASSIGN -> State 96
      BIT_AND_ASSIGN -> State 94
      BIT_OR_ASSIGN -> State 97
      BIT_XOR_ASSIGN -> State 99
      BIT_LSHIFT_ASSIGN -> State 95
      BIT_RSHIFT_ASSIGN -> State 98
      op_one_assign -> State 106
      binary_op_assign -> State 105

  State 87:
    Kernel Items:
      statement_body: expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 107
      optional_jump_label -> State 108

  State 89:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 109
      SEMICOLON -> State 110

  State 91:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      IF -> State 55
      block_body -> State 111
      if_expr -> State 112

  State 92:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 26
      optional_label_decl -> State 24
      block_expr -> State 20
      expression -> State 113

  State 106:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression]
      SEMICOLON -> [optional_expression]
    Goto:
      BOOL_LITERAL -> State 6
      INTEGER_LITERAL -> State 9
      FLOAT_LITERAL -> State 7
      RUNE_LITERAL -> State 13
      STRING_LITERAL -> State 14
      IDENTIFIER -> State 8
      LABEL_DECL -> State 10
      NOT -> State 12
      SUB -> State 15
      BIT_NEG -> State 5
      LEX_ERROR -> State 11
      literal -> State 22
      atom_expr -> State 19
      access_expr -> State 16
      unary_op -> State 28
      unary_expr -> State 27
      mul_expr -> State 23
      add_expr -> State 17
      cmp_expr -> State 21
      and_expr -> State 18
      or_expr -> State 25
      sequence_expr -> State 26
      optional_expression -> State 115
      optional_label_decl -> State 24
      block_expr -> State 20
      expression -> State 114

  State 109:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

Number of states: 115
Number of shift actions: 394
Number of reduce actions: 108
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
