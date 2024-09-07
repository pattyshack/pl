package parser

type ReducerImpl struct {
	ParseErrorReducer

	AccessExprReducerImpl
	BinaryExprReducer
	CallExprReducerImpl
	ColonExprReducerImpl
	IdentifierExprReducerImpl
	ImplicitStructExprReducerImpl
	IndexExprReducerImpl
	LiteralExprReducerImpl
	UnaryExprReducer

	ArgumentListReducer
	ArgumentReducerImpl
	TypeArgumentListReducer
}
