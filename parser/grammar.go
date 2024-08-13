// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type SymbolId int

const (
	SpacesToken         = SymbolId(256)
	NewlinesToken       = SymbolId(257)
	CommentToken        = SymbolId(258)
	BoolLiteralToken    = SymbolId(259)
	IntegerLiteralToken = SymbolId(260)
	FloatLiteralToken   = SymbolId(261)
	RuneLiteralToken    = SymbolId(262)
	StringLiteralToken  = SymbolId(263)
	LexErrorToken       = SymbolId(264)
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
	// 15:2: expression -> literal: ...
	LiteralToExpression(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 16:2: expression -> lex_error: ...
	LexErrorToExpression(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 19:2: literal -> bool: ...
	BoolToLiteral(BoolLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 20:2: literal -> integer: ...
	IntegerToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 21:2: literal -> float: ...
	FloatToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 22:2: literal -> rune: ...
	RuneToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 23:2: literal -> string: ...
	StringToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 27:2: lex_internal_tokens -> spaces: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 28:2: lex_internal_tokens -> newlines: ...
	NewlinesToLexInternalTokens(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 29:2: lex_internal_tokens -> comment: ...
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
	case LexErrorToken:
		return "LEX_ERROR"
	case ExpressionType:
		return "expression"
	case LiteralType:
		return "literal"
	case LexInternalTokensType:
		return "lex_internal_tokens"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	ExpressionType        = SymbolId(265)
	LiteralType           = SymbolId(266)
	LexInternalTokensType = SymbolId(267)
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
	_ReduceLiteralToExpression         = _ReduceType(1)
	_ReduceLexErrorToExpression        = _ReduceType(2)
	_ReduceBoolToLiteral               = _ReduceType(3)
	_ReduceIntegerToLiteral            = _ReduceType(4)
	_ReduceFloatToLiteral              = _ReduceType(5)
	_ReduceRuneToLiteral               = _ReduceType(6)
	_ReduceStringToLiteral             = _ReduceType(7)
	_ReduceSpacesToLexInternalTokens   = _ReduceType(8)
	_ReduceNewlinesToLexInternalTokens = _ReduceType(9)
	_ReduceCommentToLexInternalTokens  = _ReduceType(10)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceLiteralToExpression:
		return "LiteralToExpression"
	case _ReduceLexErrorToExpression:
		return "LexErrorToExpression"
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
	case _ReduceSpacesToLexInternalTokens:
		return "SpacesToLexInternalTokens"
	case _ReduceNewlinesToLexInternalTokens:
		return "NewlinesToLexInternalTokens"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, BoolLiteralToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, LexErrorToken:
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
	case _ReduceLiteralToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LiteralToExpression(args[0].Generic_)
	case _ReduceLexErrorToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LexErrorToExpression(args[0].Generic_)
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
	case _ReduceSpacesToLexInternalTokens:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LexInternalTokensType
		symbol.Generic_, err = reducer.SpacesToLexInternalTokens(args[0].Generic_)
	case _ReduceNewlinesToLexInternalTokens:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LexInternalTokensType
		symbol.Generic_, err = reducer.NewlinesToLexInternalTokens(args[0].Generic_)
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
	_GotoState1Action                        = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                        = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                        = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                        = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                        = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                        = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                        = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                        = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                        = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                       = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                       = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                       = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                       = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                       = &_Action{_ShiftAction, _State14, 0}
	_ReduceLiteralToExpressionAction         = &_Action{_ReduceAction, 0, _ReduceLiteralToExpression}
	_ReduceLexErrorToExpressionAction        = &_Action{_ReduceAction, 0, _ReduceLexErrorToExpression}
	_ReduceBoolToLiteralAction               = &_Action{_ReduceAction, 0, _ReduceBoolToLiteral}
	_ReduceIntegerToLiteralAction            = &_Action{_ReduceAction, 0, _ReduceIntegerToLiteral}
	_ReduceFloatToLiteralAction              = &_Action{_ReduceAction, 0, _ReduceFloatToLiteral}
	_ReduceRuneToLiteralAction               = &_Action{_ReduceAction, 0, _ReduceRuneToLiteral}
	_ReduceStringToLiteralAction             = &_Action{_ReduceAction, 0, _ReduceStringToLiteral}
	_ReduceSpacesToLexInternalTokensAction   = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceNewlinesToLexInternalTokensAction = &_Action{_ReduceAction, 0, _ReduceNewlinesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction  = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State3, _EndMarker}:            &_Action{_AcceptAction, 0, 0},
	{_State4, _EndMarker}:            &_Action{_AcceptAction, 0, 0},
	{_State1, BoolLiteralToken}:      _GotoState5Action,
	{_State1, IntegerLiteralToken}:   _GotoState7Action,
	{_State1, FloatLiteralToken}:     _GotoState6Action,
	{_State1, RuneLiteralToken}:      _GotoState9Action,
	{_State1, StringLiteralToken}:    _GotoState10Action,
	{_State1, LexErrorToken}:         _GotoState8Action,
	{_State1, ExpressionType}:        _GotoState3Action,
	{_State1, LiteralType}:           _GotoState11Action,
	{_State2, SpacesToken}:           _GotoState14Action,
	{_State2, NewlinesToken}:         _GotoState13Action,
	{_State2, CommentToken}:          _GotoState12Action,
	{_State2, LexInternalTokensType}: _GotoState4Action,
	{_State5, _EndMarker}:            _ReduceBoolToLiteralAction,
	{_State6, _EndMarker}:            _ReduceFloatToLiteralAction,
	{_State7, _EndMarker}:            _ReduceIntegerToLiteralAction,
	{_State8, _EndMarker}:            _ReduceLexErrorToExpressionAction,
	{_State9, _EndMarker}:            _ReduceRuneToLiteralAction,
	{_State10, _EndMarker}:           _ReduceStringToLiteralAction,
	{_State11, _EndMarker}:           _ReduceLiteralToExpressionAction,
	{_State12, _EndMarker}:           _ReduceCommentToLexInternalTokensAction,
	{_State13, _EndMarker}:           _ReduceNewlinesToLexInternalTokensAction,
	{_State14, _EndMarker}:           _ReduceSpacesToLexInternalTokensAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      (nil)
    Goto:
      BOOL_LITERAL -> State 5
      INTEGER_LITERAL -> State 7
      FLOAT_LITERAL -> State 6
      RUNE_LITERAL -> State 9
      STRING_LITERAL -> State 10
      LEX_ERROR -> State 8
      expression -> State 3
      literal -> State 11

  State 2:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 14
      NEWLINES -> State 13
      COMMENT -> State 12
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
      literal: BOOL_LITERAL., $
    Reduce:
      $ -> [literal]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      literal: FLOAT_LITERAL., $
    Reduce:
      $ -> [literal]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      literal: INTEGER_LITERAL., $
    Reduce:
      $ -> [literal]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      expression: LEX_ERROR., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      literal: RUNE_LITERAL., $
    Reduce:
      $ -> [literal]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      literal: STRING_LITERAL., $
    Reduce:
      $ -> [literal]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      expression: literal., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      lex_internal_tokens: NEWLINES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

Number of states: 14
Number of shift actions: 12
Number of reduce actions: 12
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
