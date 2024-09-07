package parser

type ReducerImpl struct {
	ParseErrorReducer

	AccessExprReducerImpl
	BinaryExprReducer
	CallExprReducerImpl
	ColonExprReducerImpl
	IdentifierExprReducerImpl
	InitializeExprReducerImpl
	ImplicitStructExprReducerImpl
	IndexExprReducerImpl
	LiteralExprReducerImpl
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
