package parser

type ReducerImpl struct {
	ParseErrorReducer

	LiteralExprReducerImpl
	IdentifierExprReducerImpl
	AccessExprReducerImpl
	UnaryExprReducer
	BinaryExprReducer

	ArgumentReducerImpl
	ArgumentListReducer
}
