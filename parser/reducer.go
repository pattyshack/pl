package parser

type ReducerImpl struct {
	ParseErrorReducerImpl

	FuncReducerImpl

	AccessExprReducerImpl
	BinaryExprReducerImpl
	CallExprReducerImpl
	ColonExprReducerImpl
	IfExprReducerImpl
	InitializeExprReducerImpl
	ImplicitStructExprReducerImpl
	IndexExprReducerImpl
	LiteralExprReducerImpl
	LoopExprReducerImpl
	NamedExprReducerImpl
	SelectExprReducerImpl
	StatementsExprReducerImpl
	SwitchExprReducerImpl
	UnaryExprReducerImpl

	ArrayTypeExprReducerImpl
	BinaryTypeExprReducerImpl
	InferredTypeExprReducerImpl
	MapTypeExprReducerImpl
	NamedTypeExprReducerImpl
	PropertiesTypeExprReducer
	SliceTypeExprReducerImpl
	UnaryTypeExprReducerImpl
	VarPatternReducerImpl

	BranchStatementReducerImpl
	ImportClauseReducerImpl
	ImportStatementReducerImpl
	JumpStatementReducerImpl
	UnsafeStatementReducerImpl

	ArgumentListReducerImpl
	ArgumentReducerImpl
	CasePatternsReducerImpl
	CaseAssignPatternReducerImpl
	CaseEnumPatternReducerImpl
	DefinitionListReducerImpl
	FieldDefReducerImpl
	GenericArgumentListReducerImpl
	GenericParameterReducerImpl
	GenericParameterListReducerImpl
	ParameterReducerImpl
	ParameterListReducerImpl
	TypePropertyListReducerImpl

	// TODO

	FloatingCommentReducer
	TypeDefReducer
	PackageDefReducer
}

func NewReducer() *ReducerImpl {
	return &ReducerImpl{}
}
