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

	FloatingCommentReducerImpl
	PackageDefReducerImpl

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

	TypeDefReducer
}

func NewReducer() *ReducerImpl {
	return &ReducerImpl{}
}
