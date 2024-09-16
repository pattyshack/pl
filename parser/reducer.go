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

	DefinitionReducer
	BranchStatementReducer
	TypeDefReducer
	PackageDefReducer
}

func NewReducer() *ReducerImpl {
	return &ReducerImpl{}
}
