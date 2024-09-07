package parser

type ReducerImpl struct {
	ParseErrorReducer

	AccessExprReducerImpl
	BinaryExprReducer
	CallExprReducerImpl
	ColonExprReducerImpl
	IdentifierExprReducerImpl
	ImplicitStructExprReducerImpl
	LiteralExprReducerImpl
	UnaryExprReducer

	ArgumentListReducer
	ArgumentReducerImpl
	TypeArgumentListReducer
}
