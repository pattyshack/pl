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

	SliceTypeExprReducerImpl
	ArrayTypeExprReducerImpl
	MapTypeExprReducerImpl

	ArgumentListReducer
	ArgumentReducerImpl
	TypeArgumentListReducer
	StatementListReducer
}
