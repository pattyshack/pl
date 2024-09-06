package parser

type ReducerImpl struct {
	ParseErrorReducer

	LiteralExprReducerImpl
	IdentifierExprReducerImpl
	AccessExprReducerImpl
	UnaryExprReducer
	BinaryExprReducer
	ImplicitStructExprReducerImpl

	ArgumentReducerImpl
	ArgumentListReducer
}
