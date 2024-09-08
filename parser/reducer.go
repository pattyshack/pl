package parser

type ReducerImpl struct {
	ParseErrorReducer

	AccessExprReducerImpl
	BinaryExprReducer
	CallExprReducerImpl
	ColonExprReducerImpl
	InitializeExprReducerImpl
	ImplicitStructExprReducerImpl
	IndexExprReducerImpl
	LiteralExprReducerImpl
	NamedExprReducerImpl
	UnaryExprReducer

	SliceTypeExprReducerImpl
	ArrayTypeExprReducerImpl
	MapTypeExprReducerImpl
	InferredTypeExprReducerImpl
	NamedTypeExprReducerImpl

	ArgumentListReducer
	ArgumentReducerImpl
	TypeArgumentListReducer
	StatementListReducer
}
