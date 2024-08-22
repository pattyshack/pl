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
	DotToken             = SymbolId(266)
	SemicolonToken       = SymbolId(267)
	IdentifierToken      = SymbolId(268)
	IfToken              = SymbolId(269)
	ElseToken            = SymbolId(270)
	MatchToken           = SymbolId(271)
	CaseToken            = SymbolId(272)
	ForToken             = SymbolId(273)
	ReturnToken          = SymbolId(274)
	BreakToken           = SymbolId(275)
	ContinueToken        = SymbolId(276)
	LabelDeclToken       = SymbolId(277)
	JumpLabelToken       = SymbolId(278)
	AddAssignToken       = SymbolId(279)
	SubAssignToken       = SymbolId(280)
	MulAssignToken       = SymbolId(281)
	DivAssignToken       = SymbolId(282)
	ModAssignToken       = SymbolId(283)
	AddOneAssignToken    = SymbolId(284)
	SubOneAssignToken    = SymbolId(285)
	BitNegAssignToken    = SymbolId(286)
	BitAndAssignToken    = SymbolId(287)
	BitOrAssignToken     = SymbolId(288)
	BitXorAssignToken    = SymbolId(289)
	BitLshiftAssignToken = SymbolId(290)
	BitRshiftAssignToken = SymbolId(291)
	LexErrorToken        = SymbolId(292)
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
	// 42:2: literal -> bool: ...
	BoolToLiteral(BoolLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 43:2: literal -> integer: ...
	IntegerToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 44:2: literal -> float: ...
	FloatToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 45:2: literal -> rune: ...
	RuneToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 46:2: literal -> string: ...
	StringToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 50:2: expression -> literal: ...
	LiteralToExpression(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 51:2: expression -> variable: ...
	VariableToExpression(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 52:2: expression -> accessor: ...
	AccessorToExpression(Expression_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 53:2: expression -> control_flow: ...
	ControlFlowToExpression(ControlFlow_ *GenericSymbol) (*GenericSymbol, error)

	// 54:2: expression -> lex_error: ...
	LexErrorToExpression(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 57:2: optional_label_decl -> labelled: ...
	LabelledToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 58:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 61:2: optional_jump_label -> labelled: ...
	LabelledToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 62:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 65:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 66:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 69:2: jump_type -> return: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 70:2: jump_type -> break: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 71:2: jump_type -> continue: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 74:2: op_one_assign -> add_one: ...
	AddOneToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 75:2: op_one_assign -> sub_one: ...
	SubOneToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 78:2: binary_op_assign -> add: ...
	AddToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 79:2: binary_op_assign -> sub: ...
	SubToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 80:2: binary_op_assign -> mul: ...
	MulToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 81:2: binary_op_assign -> div: ...
	DivToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 82:2: binary_op_assign -> mod: ...
	ModToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 83:2: binary_op_assign -> bit_neg: ...
	BitNegToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 84:2: binary_op_assign -> bit_and: ...
	BitAndToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 85:2: binary_op_assign -> bit_or: ...
	BitOrToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 86:2: binary_op_assign -> bit_xor: ...
	BitXorToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 87:2: binary_op_assign -> bit_lshift: ...
	BitLshiftToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 88:2: binary_op_assign -> bit_rshift: ...
	BitRshiftToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 91:2: statement_body -> expression: ...
	ExpressionToStatementBody(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 92:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 93:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(Expression_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 94:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(Expression_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_2 *GenericSymbol) (*GenericSymbol, error)

	// 97:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 98:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 102:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 104:17: sequence_body -> ...
	ToSequenceBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 105:12: sequence -> ...
	ToSequence(OptionalLabelDecl_ *GenericSymbol, SequenceBody_ *GenericSymbol) (*GenericSymbol, error)

	// 108:2: if_branch -> no_else: ...
	NoElseToIfBranch(If_ *GenericSymbol, Expression_ *GenericSymbol, Sequence_ *GenericSymbol) (*GenericSymbol, error)

	// 109:2: if_branch -> if_else: ...
	IfElseToIfBranch(If_ *GenericSymbol, Expression_ *GenericSymbol, Sequence_ *GenericSymbol, Else_ *GenericSymbol, Sequence_2 *GenericSymbol) (*GenericSymbol, error)

	// 110:2: if_branch -> multi_if_else: ...
	MultiIfElseToIfBranch(If_ *GenericSymbol, Expression_ *GenericSymbol, Sequence_ *GenericSymbol, Else_ *GenericSymbol, IfBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 113:16: match_branch -> ...
	ToMatchBranch(Match_ *GenericSymbol, Case_ *GenericSymbol) (*GenericSymbol, error)

	// 116:8: loop -> ...
	ToLoop(For_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: control_flow -> sequence: ...
	SequenceToControlFlow(Sequence_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: control_flow -> if_branch: ...
	IfBranchToControlFlow(OptionalLabelDecl_ *GenericSymbol, IfBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 121:2: control_flow -> match_branch: ...
	MatchBranchToControlFlow(OptionalLabelDecl_ *GenericSymbol, MatchBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: control_flow -> loop: ...
	LoopToControlFlow(OptionalLabelDecl_ *GenericSymbol, Loop_ *GenericSymbol) (*GenericSymbol, error)

	// 126:2: lex_internal_tokens -> spaces: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: lex_internal_tokens -> comment: ...
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
	case LexErrorToken:
		return "LEX_ERROR"
	case LiteralType:
		return "literal"
	case ExpressionType:
		return "expression"
	case OptionalLabelDeclType:
		return "optional_label_decl"
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
	case SequenceBodyType:
		return "sequence_body"
	case SequenceType:
		return "sequence"
	case IfBranchType:
		return "if_branch"
	case MatchBranchType:
		return "match_branch"
	case LoopType:
		return "loop"
	case ControlFlowType:
		return "control_flow"
	case LexInternalTokensType:
		return "lex_internal_tokens"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	LiteralType            = SymbolId(293)
	ExpressionType         = SymbolId(294)
	OptionalLabelDeclType  = SymbolId(295)
	OptionalJumpLabelType  = SymbolId(296)
	OptionalExpressionType = SymbolId(297)
	JumpTypeType           = SymbolId(298)
	OpOneAssignType        = SymbolId(299)
	BinaryOpAssignType     = SymbolId(300)
	StatementBodyType      = SymbolId(301)
	StatementType          = SymbolId(302)
	StatementsType         = SymbolId(303)
	SequenceBodyType       = SymbolId(304)
	SequenceType           = SymbolId(305)
	IfBranchType           = SymbolId(306)
	MatchBranchType        = SymbolId(307)
	LoopType               = SymbolId(308)
	ControlFlowType        = SymbolId(309)
	LexInternalTokensType  = SymbolId(310)
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
	_ReduceBoolToLiteral                  = _ReduceType(1)
	_ReduceIntegerToLiteral               = _ReduceType(2)
	_ReduceFloatToLiteral                 = _ReduceType(3)
	_ReduceRuneToLiteral                  = _ReduceType(4)
	_ReduceStringToLiteral                = _ReduceType(5)
	_ReduceLiteralToExpression            = _ReduceType(6)
	_ReduceVariableToExpression           = _ReduceType(7)
	_ReduceAccessorToExpression           = _ReduceType(8)
	_ReduceControlFlowToExpression        = _ReduceType(9)
	_ReduceLexErrorToExpression           = _ReduceType(10)
	_ReduceLabelledToOptionalLabelDecl    = _ReduceType(11)
	_ReduceUnlabelledToOptionalLabelDecl  = _ReduceType(12)
	_ReduceLabelledToOptionalJumpLabel    = _ReduceType(13)
	_ReduceUnlabelledToOptionalJumpLabel  = _ReduceType(14)
	_ReduceExpressionToOptionalExpression = _ReduceType(15)
	_ReduceNilToOptionalExpression        = _ReduceType(16)
	_ReduceReturnToJumpType               = _ReduceType(17)
	_ReduceBreakToJumpType                = _ReduceType(18)
	_ReduceContinueToJumpType             = _ReduceType(19)
	_ReduceAddOneToOpOneAssign            = _ReduceType(20)
	_ReduceSubOneToOpOneAssign            = _ReduceType(21)
	_ReduceAddToBinaryOpAssign            = _ReduceType(22)
	_ReduceSubToBinaryOpAssign            = _ReduceType(23)
	_ReduceMulToBinaryOpAssign            = _ReduceType(24)
	_ReduceDivToBinaryOpAssign            = _ReduceType(25)
	_ReduceModToBinaryOpAssign            = _ReduceType(26)
	_ReduceBitNegToBinaryOpAssign         = _ReduceType(27)
	_ReduceBitAndToBinaryOpAssign         = _ReduceType(28)
	_ReduceBitOrToBinaryOpAssign          = _ReduceType(29)
	_ReduceBitXorToBinaryOpAssign         = _ReduceType(30)
	_ReduceBitLshiftToBinaryOpAssign      = _ReduceType(31)
	_ReduceBitRshiftToBinaryOpAssign      = _ReduceType(32)
	_ReduceExpressionToStatementBody      = _ReduceType(33)
	_ReduceJumpToStatementBody            = _ReduceType(34)
	_ReduceOpOneAssignToStatementBody     = _ReduceType(35)
	_ReduceBinaryOpAssignToStatementBody  = _ReduceType(36)
	_ReduceImplicitToStatement            = _ReduceType(37)
	_ReduceExplicitToStatement            = _ReduceType(38)
	_ReduceEmptyListToStatements          = _ReduceType(39)
	_ReduceAddToStatements                = _ReduceType(40)
	_ReduceToSequenceBody                 = _ReduceType(41)
	_ReduceToSequence                     = _ReduceType(42)
	_ReduceNoElseToIfBranch               = _ReduceType(43)
	_ReduceIfElseToIfBranch               = _ReduceType(44)
	_ReduceMultiIfElseToIfBranch          = _ReduceType(45)
	_ReduceToMatchBranch                  = _ReduceType(46)
	_ReduceToLoop                         = _ReduceType(47)
	_ReduceSequenceToControlFlow          = _ReduceType(48)
	_ReduceIfBranchToControlFlow          = _ReduceType(49)
	_ReduceMatchBranchToControlFlow       = _ReduceType(50)
	_ReduceLoopToControlFlow              = _ReduceType(51)
	_ReduceSpacesToLexInternalTokens      = _ReduceType(52)
	_ReduceCommentToLexInternalTokens     = _ReduceType(53)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceBoolToLiteral:
		return "BoolToLiteral"
	case _ReduceIntegerToLiteral:
		return "IntegerToLiteral"
	case _ReduceFloatToLiteral:
		return "FloatToLiteral"
	case _ReduceRuneToLiteral:
		return "RuneToLiteral"
	case _ReduceStringToLiteral:
		return "StringToLiteral"
	case _ReduceLiteralToExpression:
		return "LiteralToExpression"
	case _ReduceVariableToExpression:
		return "VariableToExpression"
	case _ReduceAccessorToExpression:
		return "AccessorToExpression"
	case _ReduceControlFlowToExpression:
		return "ControlFlowToExpression"
	case _ReduceLexErrorToExpression:
		return "LexErrorToExpression"
	case _ReduceLabelledToOptionalLabelDecl:
		return "LabelledToOptionalLabelDecl"
	case _ReduceUnlabelledToOptionalLabelDecl:
		return "UnlabelledToOptionalLabelDecl"
	case _ReduceLabelledToOptionalJumpLabel:
		return "LabelledToOptionalJumpLabel"
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
	case _ReduceAddOneToOpOneAssign:
		return "AddOneToOpOneAssign"
	case _ReduceSubOneToOpOneAssign:
		return "SubOneToOpOneAssign"
	case _ReduceAddToBinaryOpAssign:
		return "AddToBinaryOpAssign"
	case _ReduceSubToBinaryOpAssign:
		return "SubToBinaryOpAssign"
	case _ReduceMulToBinaryOpAssign:
		return "MulToBinaryOpAssign"
	case _ReduceDivToBinaryOpAssign:
		return "DivToBinaryOpAssign"
	case _ReduceModToBinaryOpAssign:
		return "ModToBinaryOpAssign"
	case _ReduceBitNegToBinaryOpAssign:
		return "BitNegToBinaryOpAssign"
	case _ReduceBitAndToBinaryOpAssign:
		return "BitAndToBinaryOpAssign"
	case _ReduceBitOrToBinaryOpAssign:
		return "BitOrToBinaryOpAssign"
	case _ReduceBitXorToBinaryOpAssign:
		return "BitXorToBinaryOpAssign"
	case _ReduceBitLshiftToBinaryOpAssign:
		return "BitLshiftToBinaryOpAssign"
	case _ReduceBitRshiftToBinaryOpAssign:
		return "BitRshiftToBinaryOpAssign"
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
	case _ReduceToSequenceBody:
		return "ToSequenceBody"
	case _ReduceToSequence:
		return "ToSequence"
	case _ReduceNoElseToIfBranch:
		return "NoElseToIfBranch"
	case _ReduceIfElseToIfBranch:
		return "IfElseToIfBranch"
	case _ReduceMultiIfElseToIfBranch:
		return "MultiIfElseToIfBranch"
	case _ReduceToMatchBranch:
		return "ToMatchBranch"
	case _ReduceToLoop:
		return "ToLoop"
	case _ReduceSequenceToControlFlow:
		return "SequenceToControlFlow"
	case _ReduceIfBranchToControlFlow:
		return "IfBranchToControlFlow"
	case _ReduceMatchBranchToControlFlow:
		return "MatchBranchToControlFlow"
	case _ReduceLoopToControlFlow:
		return "LoopToControlFlow"
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
	_State1  = _StateId(1)
	_State2  = _StateId(2)
	_State3  = _StateId(3)
	_State4  = _StateId(4)
	_State5  = _StateId(5)
	_State6  = _StateId(6)
	_State7  = _StateId(7)
	_State8  = _StateId(8)
	_State9  = _StateId(9)
	_State10 = _StateId(10)
	_State11 = _StateId(11)
	_State12 = _StateId(12)
	_State13 = _StateId(13)
	_State14 = _StateId(14)
	_State15 = _StateId(15)
	_State16 = _StateId(16)
	_State17 = _StateId(17)
	_State18 = _StateId(18)
	_State19 = _StateId(19)
	_State20 = _StateId(20)
	_State21 = _StateId(21)
	_State22 = _StateId(22)
	_State23 = _StateId(23)
	_State24 = _StateId(24)
	_State25 = _StateId(25)
	_State26 = _StateId(26)
	_State27 = _StateId(27)
	_State28 = _StateId(28)
	_State29 = _StateId(29)
	_State30 = _StateId(30)
	_State31 = _StateId(31)
	_State32 = _StateId(32)
	_State33 = _StateId(33)
	_State34 = _StateId(34)
	_State35 = _StateId(35)
	_State36 = _StateId(36)
	_State37 = _StateId(37)
	_State38 = _StateId(38)
	_State39 = _StateId(39)
	_State40 = _StateId(40)
	_State41 = _StateId(41)
	_State42 = _StateId(42)
	_State43 = _StateId(43)
	_State44 = _StateId(44)
	_State45 = _StateId(45)
	_State46 = _StateId(46)
	_State47 = _StateId(47)
	_State48 = _StateId(48)
	_State49 = _StateId(49)
	_State50 = _StateId(50)
	_State51 = _StateId(51)
	_State52 = _StateId(52)
	_State53 = _StateId(53)
	_State54 = _StateId(54)
	_State55 = _StateId(55)
	_State56 = _StateId(56)
	_State57 = _StateId(57)
	_State58 = _StateId(58)
	_State59 = _StateId(59)
	_State60 = _StateId(60)
	_State61 = _StateId(61)
	_State62 = _StateId(62)
	_State63 = _StateId(63)
	_State64 = _StateId(64)
	_State65 = _StateId(65)
	_State66 = _StateId(66)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, BoolLiteralToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, LbraceToken, RbraceToken, DotToken, SemicolonToken, IdentifierToken, IfToken, ElseToken, MatchToken, CaseToken, ForToken, ReturnToken, BreakToken, ContinueToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, LexErrorToken:
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
	case _ReduceBoolToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.BoolToLiteral(args[0].Generic_)
	case _ReduceIntegerToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.IntegerToLiteral(args[0].Generic_)
	case _ReduceFloatToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.FloatToLiteral(args[0].Generic_)
	case _ReduceRuneToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.RuneToLiteral(args[0].Generic_)
	case _ReduceStringToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.StringToLiteral(args[0].Generic_)
	case _ReduceLiteralToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LiteralToExpression(args[0].Generic_)
	case _ReduceVariableToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.VariableToExpression(args[0].Generic_)
	case _ReduceAccessorToExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.AccessorToExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceControlFlowToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.ControlFlowToExpression(args[0].Generic_)
	case _ReduceLexErrorToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LexErrorToExpression(args[0].Generic_)
	case _ReduceLabelledToOptionalLabelDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.LabelledToOptionalLabelDecl(args[0].Generic_)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.UnlabelledToOptionalLabelDecl()
	case _ReduceLabelledToOptionalJumpLabel:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.LabelledToOptionalJumpLabel(args[0].Generic_)
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
	case _ReduceAddOneToOpOneAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpOneAssignType
		symbol.Generic_, err = reducer.AddOneToOpOneAssign(args[0].Generic_)
	case _ReduceSubOneToOpOneAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpOneAssignType
		symbol.Generic_, err = reducer.SubOneToOpOneAssign(args[0].Generic_)
	case _ReduceAddToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.AddToBinaryOpAssign(args[0].Generic_)
	case _ReduceSubToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.SubToBinaryOpAssign(args[0].Generic_)
	case _ReduceMulToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.MulToBinaryOpAssign(args[0].Generic_)
	case _ReduceDivToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.DivToBinaryOpAssign(args[0].Generic_)
	case _ReduceModToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.ModToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitNegToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitNegToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitAndToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitAndToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitOrToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitOrToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitXorToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitXorToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitLshiftToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitLshiftToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitRshiftToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitRshiftToBinaryOpAssign(args[0].Generic_)
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
	case _ReduceToSequenceBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SequenceBodyType
		symbol.Generic_, err = reducer.ToSequenceBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToSequence:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SequenceType
		symbol.Generic_, err = reducer.ToSequence(args[0].Generic_, args[1].Generic_)
	case _ReduceNoElseToIfBranch:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfBranchType
		symbol.Generic_, err = reducer.NoElseToIfBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceIfElseToIfBranch:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfBranchType
		symbol.Generic_, err = reducer.IfElseToIfBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceMultiIfElseToIfBranch:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfBranchType
		symbol.Generic_, err = reducer.MultiIfElseToIfBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceToMatchBranch:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = MatchBranchType
		symbol.Generic_, err = reducer.ToMatchBranch(args[0].Generic_, args[1].Generic_)
	case _ReduceToLoop:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LoopType
		symbol.Generic_, err = reducer.ToLoop(args[0].Generic_)
	case _ReduceSequenceToControlFlow:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ControlFlowType
		symbol.Generic_, err = reducer.SequenceToControlFlow(args[0].Generic_)
	case _ReduceIfBranchToControlFlow:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ControlFlowType
		symbol.Generic_, err = reducer.IfBranchToControlFlow(args[0].Generic_, args[1].Generic_)
	case _ReduceMatchBranchToControlFlow:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ControlFlowType
		symbol.Generic_, err = reducer.MatchBranchToControlFlow(args[0].Generic_, args[1].Generic_)
	case _ReduceLoopToControlFlow:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ControlFlowType
		symbol.Generic_, err = reducer.LoopToControlFlow(args[0].Generic_, args[1].Generic_)
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
	_GotoState1Action                           = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                           = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                           = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                           = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                           = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                           = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                           = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                           = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                           = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                          = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                          = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                          = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                          = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                          = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                          = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                          = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                          = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                          = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                          = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                          = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                          = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                          = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                          = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                          = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                          = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                          = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                          = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                          = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                          = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                          = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                          = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                          = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                          = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                          = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                          = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                          = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                          = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                          = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                          = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                          = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                          = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                          = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                          = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                          = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                          = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                          = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                          = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                          = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                          = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                          = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                          = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                          = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                          = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                          = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                          = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                          = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                          = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                          = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                          = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                          = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                          = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                          = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                          = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                          = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                          = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                          = &_Action{_ShiftAction, _State66, 0}
	_ReduceBoolToLiteralAction                  = &_Action{_ReduceAction, 0, _ReduceBoolToLiteral}
	_ReduceIntegerToLiteralAction               = &_Action{_ReduceAction, 0, _ReduceIntegerToLiteral}
	_ReduceFloatToLiteralAction                 = &_Action{_ReduceAction, 0, _ReduceFloatToLiteral}
	_ReduceRuneToLiteralAction                  = &_Action{_ReduceAction, 0, _ReduceRuneToLiteral}
	_ReduceStringToLiteralAction                = &_Action{_ReduceAction, 0, _ReduceStringToLiteral}
	_ReduceLiteralToExpressionAction            = &_Action{_ReduceAction, 0, _ReduceLiteralToExpression}
	_ReduceVariableToExpressionAction           = &_Action{_ReduceAction, 0, _ReduceVariableToExpression}
	_ReduceAccessorToExpressionAction           = &_Action{_ReduceAction, 0, _ReduceAccessorToExpression}
	_ReduceControlFlowToExpressionAction        = &_Action{_ReduceAction, 0, _ReduceControlFlowToExpression}
	_ReduceLexErrorToExpressionAction           = &_Action{_ReduceAction, 0, _ReduceLexErrorToExpression}
	_ReduceLabelledToOptionalLabelDeclAction    = &_Action{_ReduceAction, 0, _ReduceLabelledToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction  = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceLabelledToOptionalJumpLabelAction    = &_Action{_ReduceAction, 0, _ReduceLabelledToOptionalJumpLabel}
	_ReduceUnlabelledToOptionalJumpLabelAction  = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalJumpLabel}
	_ReduceExpressionToOptionalExpressionAction = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReduceReturnToJumpTypeAction               = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction             = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceAddOneToOpOneAssignAction            = &_Action{_ReduceAction, 0, _ReduceAddOneToOpOneAssign}
	_ReduceSubOneToOpOneAssignAction            = &_Action{_ReduceAction, 0, _ReduceSubOneToOpOneAssign}
	_ReduceAddToBinaryOpAssignAction            = &_Action{_ReduceAction, 0, _ReduceAddToBinaryOpAssign}
	_ReduceSubToBinaryOpAssignAction            = &_Action{_ReduceAction, 0, _ReduceSubToBinaryOpAssign}
	_ReduceMulToBinaryOpAssignAction            = &_Action{_ReduceAction, 0, _ReduceMulToBinaryOpAssign}
	_ReduceDivToBinaryOpAssignAction            = &_Action{_ReduceAction, 0, _ReduceDivToBinaryOpAssign}
	_ReduceModToBinaryOpAssignAction            = &_Action{_ReduceAction, 0, _ReduceModToBinaryOpAssign}
	_ReduceBitNegToBinaryOpAssignAction         = &_Action{_ReduceAction, 0, _ReduceBitNegToBinaryOpAssign}
	_ReduceBitAndToBinaryOpAssignAction         = &_Action{_ReduceAction, 0, _ReduceBitAndToBinaryOpAssign}
	_ReduceBitOrToBinaryOpAssignAction          = &_Action{_ReduceAction, 0, _ReduceBitOrToBinaryOpAssign}
	_ReduceBitXorToBinaryOpAssignAction         = &_Action{_ReduceAction, 0, _ReduceBitXorToBinaryOpAssign}
	_ReduceBitLshiftToBinaryOpAssignAction      = &_Action{_ReduceAction, 0, _ReduceBitLshiftToBinaryOpAssign}
	_ReduceBitRshiftToBinaryOpAssignAction      = &_Action{_ReduceAction, 0, _ReduceBitRshiftToBinaryOpAssign}
	_ReduceExpressionToStatementBodyAction      = &_Action{_ReduceAction, 0, _ReduceExpressionToStatementBody}
	_ReduceJumpToStatementBodyAction            = &_Action{_ReduceAction, 0, _ReduceJumpToStatementBody}
	_ReduceOpOneAssignToStatementBodyAction     = &_Action{_ReduceAction, 0, _ReduceOpOneAssignToStatementBody}
	_ReduceBinaryOpAssignToStatementBodyAction  = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignToStatementBody}
	_ReduceImplicitToStatementAction            = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction            = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceEmptyListToStatementsAction          = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceToSequenceBodyAction                 = &_Action{_ReduceAction, 0, _ReduceToSequenceBody}
	_ReduceToSequenceAction                     = &_Action{_ReduceAction, 0, _ReduceToSequence}
	_ReduceNoElseToIfBranchAction               = &_Action{_ReduceAction, 0, _ReduceNoElseToIfBranch}
	_ReduceIfElseToIfBranchAction               = &_Action{_ReduceAction, 0, _ReduceIfElseToIfBranch}
	_ReduceMultiIfElseToIfBranchAction          = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfBranch}
	_ReduceToMatchBranchAction                  = &_Action{_ReduceAction, 0, _ReduceToMatchBranch}
	_ReduceToLoopAction                         = &_Action{_ReduceAction, 0, _ReduceToLoop}
	_ReduceSequenceToControlFlowAction          = &_Action{_ReduceAction, 0, _ReduceSequenceToControlFlow}
	_ReduceIfBranchToControlFlowAction          = &_Action{_ReduceAction, 0, _ReduceIfBranchToControlFlow}
	_ReduceMatchBranchToControlFlowAction       = &_Action{_ReduceAction, 0, _ReduceMatchBranchToControlFlow}
	_ReduceLoopToControlFlowAction              = &_Action{_ReduceAction, 0, _ReduceLoopToControlFlow}
	_ReduceSpacesToLexInternalTokensAction      = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction     = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State3, _EndMarker}:              &_Action{_AcceptAction, 0, 0},
	{_State4, _EndMarker}:              &_Action{_AcceptAction, 0, 0},
	{_State1, BoolLiteralToken}:        _GotoState5Action,
	{_State1, IntegerLiteralToken}:     _GotoState8Action,
	{_State1, FloatLiteralToken}:       _GotoState6Action,
	{_State1, RuneLiteralToken}:        _GotoState11Action,
	{_State1, StringLiteralToken}:      _GotoState12Action,
	{_State1, IdentifierToken}:         _GotoState7Action,
	{_State1, LabelDeclToken}:          _GotoState9Action,
	{_State1, LexErrorToken}:           _GotoState10Action,
	{_State1, LiteralType}:             _GotoState14Action,
	{_State1, ExpressionType}:          _GotoState3Action,
	{_State1, OptionalLabelDeclType}:   _GotoState15Action,
	{_State1, SequenceType}:            _GotoState16Action,
	{_State1, ControlFlowType}:         _GotoState13Action,
	{_State2, SpacesToken}:             _GotoState18Action,
	{_State2, CommentToken}:            _GotoState17Action,
	{_State2, LexInternalTokensType}:   _GotoState4Action,
	{_State3, DotToken}:                _GotoState19Action,
	{_State15, LbraceToken}:            _GotoState22Action,
	{_State15, IfToken}:                _GotoState21Action,
	{_State15, MatchToken}:             _GotoState23Action,
	{_State15, ForToken}:               _GotoState20Action,
	{_State15, SequenceBodyType}:       _GotoState27Action,
	{_State15, IfBranchType}:           _GotoState24Action,
	{_State15, MatchBranchType}:        _GotoState26Action,
	{_State15, LoopType}:               _GotoState25Action,
	{_State19, IdentifierToken}:        _GotoState28Action,
	{_State21, BoolLiteralToken}:       _GotoState5Action,
	{_State21, IntegerLiteralToken}:    _GotoState8Action,
	{_State21, FloatLiteralToken}:      _GotoState6Action,
	{_State21, RuneLiteralToken}:       _GotoState11Action,
	{_State21, StringLiteralToken}:     _GotoState12Action,
	{_State21, IdentifierToken}:        _GotoState7Action,
	{_State21, LabelDeclToken}:         _GotoState9Action,
	{_State21, LexErrorToken}:          _GotoState10Action,
	{_State21, LiteralType}:            _GotoState14Action,
	{_State21, ExpressionType}:         _GotoState29Action,
	{_State21, OptionalLabelDeclType}:  _GotoState15Action,
	{_State21, SequenceType}:           _GotoState16Action,
	{_State21, ControlFlowType}:        _GotoState13Action,
	{_State22, StatementsType}:         _GotoState30Action,
	{_State23, CaseToken}:              _GotoState31Action,
	{_State29, DotToken}:               _GotoState19Action,
	{_State29, LabelDeclToken}:         _GotoState9Action,
	{_State29, OptionalLabelDeclType}:  _GotoState32Action,
	{_State29, SequenceType}:           _GotoState33Action,
	{_State30, BoolLiteralToken}:       _GotoState5Action,
	{_State30, IntegerLiteralToken}:    _GotoState8Action,
	{_State30, FloatLiteralToken}:      _GotoState6Action,
	{_State30, RuneLiteralToken}:       _GotoState11Action,
	{_State30, StringLiteralToken}:     _GotoState12Action,
	{_State30, RbraceToken}:            _GotoState36Action,
	{_State30, IdentifierToken}:        _GotoState7Action,
	{_State30, ReturnToken}:            _GotoState37Action,
	{_State30, BreakToken}:             _GotoState34Action,
	{_State30, ContinueToken}:          _GotoState35Action,
	{_State30, LabelDeclToken}:         _GotoState9Action,
	{_State30, LexErrorToken}:          _GotoState10Action,
	{_State30, LiteralType}:            _GotoState14Action,
	{_State30, ExpressionType}:         _GotoState38Action,
	{_State30, OptionalLabelDeclType}:  _GotoState15Action,
	{_State30, JumpTypeType}:           _GotoState39Action,
	{_State30, StatementBodyType}:      _GotoState41Action,
	{_State30, StatementType}:          _GotoState40Action,
	{_State30, SequenceType}:           _GotoState16Action,
	{_State30, ControlFlowType}:        _GotoState13Action,
	{_State32, LbraceToken}:            _GotoState22Action,
	{_State32, SequenceBodyType}:       _GotoState27Action,
	{_State33, ElseToken}:              _GotoState42Action,
	{_State38, DotToken}:               _GotoState19Action,
	{_State38, AddAssignToken}:         _GotoState43Action,
	{_State38, SubAssignToken}:         _GotoState54Action,
	{_State38, MulAssignToken}:         _GotoState53Action,
	{_State38, DivAssignToken}:         _GotoState51Action,
	{_State38, ModAssignToken}:         _GotoState52Action,
	{_State38, AddOneAssignToken}:      _GotoState44Action,
	{_State38, SubOneAssignToken}:      _GotoState55Action,
	{_State38, BitNegAssignToken}:      _GotoState47Action,
	{_State38, BitAndAssignToken}:      _GotoState45Action,
	{_State38, BitOrAssignToken}:       _GotoState48Action,
	{_State38, BitXorAssignToken}:      _GotoState50Action,
	{_State38, BitLshiftAssignToken}:   _GotoState46Action,
	{_State38, BitRshiftAssignToken}:   _GotoState49Action,
	{_State38, OpOneAssignType}:        _GotoState57Action,
	{_State38, BinaryOpAssignType}:     _GotoState56Action,
	{_State39, JumpLabelToken}:         _GotoState58Action,
	{_State39, OptionalJumpLabelType}:  _GotoState59Action,
	{_State41, NewlinesToken}:          _GotoState60Action,
	{_State41, SemicolonToken}:         _GotoState61Action,
	{_State42, IfToken}:                _GotoState21Action,
	{_State42, LabelDeclToken}:         _GotoState9Action,
	{_State42, OptionalLabelDeclType}:  _GotoState32Action,
	{_State42, SequenceType}:           _GotoState63Action,
	{_State42, IfBranchType}:           _GotoState62Action,
	{_State56, BoolLiteralToken}:       _GotoState5Action,
	{_State56, IntegerLiteralToken}:    _GotoState8Action,
	{_State56, FloatLiteralToken}:      _GotoState6Action,
	{_State56, RuneLiteralToken}:       _GotoState11Action,
	{_State56, StringLiteralToken}:     _GotoState12Action,
	{_State56, IdentifierToken}:        _GotoState7Action,
	{_State56, LabelDeclToken}:         _GotoState9Action,
	{_State56, LexErrorToken}:          _GotoState10Action,
	{_State56, LiteralType}:            _GotoState14Action,
	{_State56, ExpressionType}:         _GotoState64Action,
	{_State56, OptionalLabelDeclType}:  _GotoState15Action,
	{_State56, SequenceType}:           _GotoState16Action,
	{_State56, ControlFlowType}:        _GotoState13Action,
	{_State59, BoolLiteralToken}:       _GotoState5Action,
	{_State59, IntegerLiteralToken}:    _GotoState8Action,
	{_State59, FloatLiteralToken}:      _GotoState6Action,
	{_State59, RuneLiteralToken}:       _GotoState11Action,
	{_State59, StringLiteralToken}:     _GotoState12Action,
	{_State59, IdentifierToken}:        _GotoState7Action,
	{_State59, LabelDeclToken}:         _GotoState9Action,
	{_State59, LexErrorToken}:          _GotoState10Action,
	{_State59, LiteralType}:            _GotoState14Action,
	{_State59, ExpressionType}:         _GotoState65Action,
	{_State59, OptionalLabelDeclType}:  _GotoState15Action,
	{_State59, OptionalExpressionType}: _GotoState66Action,
	{_State59, SequenceType}:           _GotoState16Action,
	{_State59, ControlFlowType}:        _GotoState13Action,
	{_State64, DotToken}:               _GotoState19Action,
	{_State65, DotToken}:               _GotoState19Action,
	{_State1, _WildcardMarker}:         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State5, _WildcardMarker}:         _ReduceBoolToLiteralAction,
	{_State6, _WildcardMarker}:         _ReduceFloatToLiteralAction,
	{_State7, _WildcardMarker}:         _ReduceVariableToExpressionAction,
	{_State8, _WildcardMarker}:         _ReduceIntegerToLiteralAction,
	{_State9, _WildcardMarker}:         _ReduceLabelledToOptionalLabelDeclAction,
	{_State10, _WildcardMarker}:        _ReduceLexErrorToExpressionAction,
	{_State11, _WildcardMarker}:        _ReduceRuneToLiteralAction,
	{_State12, _WildcardMarker}:        _ReduceStringToLiteralAction,
	{_State13, _WildcardMarker}:        _ReduceControlFlowToExpressionAction,
	{_State14, _WildcardMarker}:        _ReduceLiteralToExpressionAction,
	{_State16, _WildcardMarker}:        _ReduceSequenceToControlFlowAction,
	{_State17, _EndMarker}:             _ReduceCommentToLexInternalTokensAction,
	{_State18, _EndMarker}:             _ReduceSpacesToLexInternalTokensAction,
	{_State20, _WildcardMarker}:        _ReduceToLoopAction,
	{_State21, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State22, _WildcardMarker}:        _ReduceEmptyListToStatementsAction,
	{_State24, _WildcardMarker}:        _ReduceIfBranchToControlFlowAction,
	{_State25, _WildcardMarker}:        _ReduceLoopToControlFlowAction,
	{_State26, _WildcardMarker}:        _ReduceMatchBranchToControlFlowAction,
	{_State27, _WildcardMarker}:        _ReduceToSequenceAction,
	{_State28, _WildcardMarker}:        _ReduceAccessorToExpressionAction,
	{_State29, LbraceToken}:            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State30, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State31, _WildcardMarker}:        _ReduceToMatchBranchAction,
	{_State33, _WildcardMarker}:        _ReduceNoElseToIfBranchAction,
	{_State34, _WildcardMarker}:        _ReduceBreakToJumpTypeAction,
	{_State35, _WildcardMarker}:        _ReduceContinueToJumpTypeAction,
	{_State36, _WildcardMarker}:        _ReduceToSequenceBodyAction,
	{_State37, _WildcardMarker}:        _ReduceReturnToJumpTypeAction,
	{_State38, _WildcardMarker}:        _ReduceExpressionToStatementBodyAction,
	{_State39, _WildcardMarker}:        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State40, _WildcardMarker}:        _ReduceAddToStatementsAction,
	{_State42, LbraceToken}:            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State43, _WildcardMarker}:        _ReduceAddToBinaryOpAssignAction,
	{_State44, _WildcardMarker}:        _ReduceAddOneToOpOneAssignAction,
	{_State45, _WildcardMarker}:        _ReduceBitAndToBinaryOpAssignAction,
	{_State46, _WildcardMarker}:        _ReduceBitLshiftToBinaryOpAssignAction,
	{_State47, _WildcardMarker}:        _ReduceBitNegToBinaryOpAssignAction,
	{_State48, _WildcardMarker}:        _ReduceBitOrToBinaryOpAssignAction,
	{_State49, _WildcardMarker}:        _ReduceBitRshiftToBinaryOpAssignAction,
	{_State50, _WildcardMarker}:        _ReduceBitXorToBinaryOpAssignAction,
	{_State51, _WildcardMarker}:        _ReduceDivToBinaryOpAssignAction,
	{_State52, _WildcardMarker}:        _ReduceModToBinaryOpAssignAction,
	{_State53, _WildcardMarker}:        _ReduceMulToBinaryOpAssignAction,
	{_State54, _WildcardMarker}:        _ReduceSubToBinaryOpAssignAction,
	{_State55, _WildcardMarker}:        _ReduceSubOneToOpOneAssignAction,
	{_State56, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State57, _WildcardMarker}:        _ReduceOpOneAssignToStatementBodyAction,
	{_State58, _WildcardMarker}:        _ReduceLabelledToOptionalJumpLabelAction,
	{_State59, _WildcardMarker}:        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State59, NewlinesToken}:          _ReduceNilToOptionalExpressionAction,
	{_State59, SemicolonToken}:         _ReduceNilToOptionalExpressionAction,
	{_State60, _WildcardMarker}:        _ReduceImplicitToStatementAction,
	{_State61, _WildcardMarker}:        _ReduceExplicitToStatementAction,
	{_State62, _WildcardMarker}:        _ReduceMultiIfElseToIfBranchAction,
	{_State63, _WildcardMarker}:        _ReduceIfElseToIfBranchAction,
	{_State64, _WildcardMarker}:        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State65, _WildcardMarker}:        _ReduceExpressionToOptionalExpressionAction,
	{_State66, _WildcardMarker}:        _ReduceJumpToStatementBodyAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 5
      INTEGER_LITERAL -> State 8
      FLOAT_LITERAL -> State 6
      RUNE_LITERAL -> State 11
      STRING_LITERAL -> State 12
      IDENTIFIER -> State 7
      LABEL_DECL -> State 9
      LEX_ERROR -> State 10
      literal -> State 14
      expression -> State 3
      optional_label_decl -> State 15
      sequence -> State 16
      control_flow -> State 13

  State 2:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 18
      COMMENT -> State 17
      lex_internal_tokens -> State 4

  State 3:
    Kernel Items:
      #accept: ^ expression., $
      expression: expression.DOT IDENTIFIER
    Reduce:
      $ -> [#accept]
    Goto:
      DOT -> State 19

  State 4:
    Kernel Items:
      #accept: ^ lex_internal_tokens., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      literal: BOOL_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      expression: IDENTIFIER., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      expression: LEX_ERROR., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      expression: control_flow., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      expression: literal., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      sequence: optional_label_decl.sequence_body
      control_flow: optional_label_decl.if_branch
      control_flow: optional_label_decl.match_branch
      control_flow: optional_label_decl.loop
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 22
      IF -> State 21
      MATCH -> State 23
      FOR -> State 20
      sequence_body -> State 27
      if_branch -> State 24
      match_branch -> State 26
      loop -> State 25

  State 16:
    Kernel Items:
      control_flow: sequence., *
    Reduce:
      * -> [control_flow]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      expression: expression DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 28

  State 20:
    Kernel Items:
      loop: FOR., *
    Reduce:
      * -> [loop]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      if_branch: IF.expression sequence
      if_branch: IF.expression sequence ELSE sequence
      if_branch: IF.expression sequence ELSE if_branch
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 5
      INTEGER_LITERAL -> State 8
      FLOAT_LITERAL -> State 6
      RUNE_LITERAL -> State 11
      STRING_LITERAL -> State 12
      IDENTIFIER -> State 7
      LABEL_DECL -> State 9
      LEX_ERROR -> State 10
      literal -> State 14
      expression -> State 29
      optional_label_decl -> State 15
      sequence -> State 16
      control_flow -> State 13

  State 22:
    Kernel Items:
      sequence_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 30

  State 23:
    Kernel Items:
      match_branch: MATCH.CASE
    Reduce:
      (nil)
    Goto:
      CASE -> State 31

  State 24:
    Kernel Items:
      control_flow: optional_label_decl if_branch., *
    Reduce:
      * -> [control_flow]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      control_flow: optional_label_decl loop., *
    Reduce:
      * -> [control_flow]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      control_flow: optional_label_decl match_branch., *
    Reduce:
      * -> [control_flow]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      sequence: optional_label_decl sequence_body., *
    Reduce:
      * -> [sequence]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      expression: expression DOT IDENTIFIER., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      expression: expression.DOT IDENTIFIER
      if_branch: IF expression.sequence
      if_branch: IF expression.sequence ELSE sequence
      if_branch: IF expression.sequence ELSE if_branch
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      DOT -> State 19
      LABEL_DECL -> State 9
      optional_label_decl -> State 32
      sequence -> State 33

  State 30:
    Kernel Items:
      statements: statements.statement
      sequence_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 5
      INTEGER_LITERAL -> State 8
      FLOAT_LITERAL -> State 6
      RUNE_LITERAL -> State 11
      STRING_LITERAL -> State 12
      RBRACE -> State 36
      IDENTIFIER -> State 7
      RETURN -> State 37
      BREAK -> State 34
      CONTINUE -> State 35
      LABEL_DECL -> State 9
      LEX_ERROR -> State 10
      literal -> State 14
      expression -> State 38
      optional_label_decl -> State 15
      jump_type -> State 39
      statement_body -> State 41
      statement -> State 40
      sequence -> State 16
      control_flow -> State 13

  State 31:
    Kernel Items:
      match_branch: MATCH CASE., *
    Reduce:
      * -> [match_branch]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      sequence: optional_label_decl.sequence_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 22
      sequence_body -> State 27

  State 33:
    Kernel Items:
      if_branch: IF expression sequence., *
      if_branch: IF expression sequence.ELSE sequence
      if_branch: IF expression sequence.ELSE if_branch
    Reduce:
      * -> [if_branch]
    Goto:
      ELSE -> State 42

  State 34:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      sequence_body: LBRACE statements RBRACE., *
    Reduce:
      * -> [sequence_body]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      expression: expression.DOT IDENTIFIER
      statement_body: expression., *
      statement_body: expression.op_one_assign
      statement_body: expression.binary_op_assign expression
    Reduce:
      * -> [statement_body]
    Goto:
      DOT -> State 19
      ADD_ASSIGN -> State 43
      SUB_ASSIGN -> State 54
      MUL_ASSIGN -> State 53
      DIV_ASSIGN -> State 51
      MOD_ASSIGN -> State 52
      ADD_ONE_ASSIGN -> State 44
      SUB_ONE_ASSIGN -> State 55
      BIT_NEG_ASSIGN -> State 47
      BIT_AND_ASSIGN -> State 45
      BIT_OR_ASSIGN -> State 48
      BIT_XOR_ASSIGN -> State 50
      BIT_LSHIFT_ASSIGN -> State 46
      BIT_RSHIFT_ASSIGN -> State 49
      op_one_assign -> State 57
      binary_op_assign -> State 56

  State 39:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 58
      optional_jump_label -> State 59

  State 40:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 60
      SEMICOLON -> State 61

  State 42:
    Kernel Items:
      if_branch: IF expression sequence ELSE.sequence
      if_branch: IF expression sequence ELSE.if_branch
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      IF -> State 21
      LABEL_DECL -> State 9
      optional_label_decl -> State 32
      sequence -> State 63
      if_branch -> State 62

  State 43:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      statement_body: expression binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 5
      INTEGER_LITERAL -> State 8
      FLOAT_LITERAL -> State 6
      RUNE_LITERAL -> State 11
      STRING_LITERAL -> State 12
      IDENTIFIER -> State 7
      LABEL_DECL -> State 9
      LEX_ERROR -> State 10
      literal -> State 14
      expression -> State 64
      optional_label_decl -> State 15
      sequence -> State 16
      control_flow -> State 13

  State 57:
    Kernel Items:
      statement_body: expression op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression]
      SEMICOLON -> [optional_expression]
    Goto:
      BOOL_LITERAL -> State 5
      INTEGER_LITERAL -> State 8
      FLOAT_LITERAL -> State 6
      RUNE_LITERAL -> State 11
      STRING_LITERAL -> State 12
      IDENTIFIER -> State 7
      LABEL_DECL -> State 9
      LEX_ERROR -> State 10
      literal -> State 14
      expression -> State 65
      optional_label_decl -> State 15
      optional_expression -> State 66
      sequence -> State 16
      control_flow -> State 13

  State 60:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      if_branch: IF expression sequence ELSE if_branch., *
    Reduce:
      * -> [if_branch]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      if_branch: IF expression sequence ELSE sequence., *
    Reduce:
      * -> [if_branch]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      expression: expression.DOT IDENTIFIER
      statement_body: expression binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      DOT -> State 19

  State 65:
    Kernel Items:
      expression: expression.DOT IDENTIFIER
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      DOT -> State 19

  State 66:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

Number of states: 66
Number of shift actions: 122
Number of reduce actions: 62
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
