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
	NamedExprReducerImpl
	StatementsExprReducerImpl
	UnaryExprReducerImpl

	ArrayTypeExprReducerImpl
	BinaryTypeExprReducerImpl
	InferredTypeExprReducerImpl
	MapTypeExprReducerImpl
	NamedTypeExprReducerImpl
	PropertiesTypeExprReducer
	SliceTypeExprReducerImpl
	UnaryTypeExprReducerImpl

	ImportClauseReducerImpl
	ImportStatementReducerImpl

	ArgumentListReducerImpl
	ArgumentReducerImpl
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
	StatementReducer
	OptionalSimpleStatementReducer
	UnsafeStatementReducer
	UnsafeStatementPropertyReducer
	JumpStatementReducer
	FallthroughStatementReducer
	AssignStatementReducer
	UnaryOpAssignStatementReducer
	BinaryOpAssignStatementReducer
	CasePatternsReducer
	VarDeclPatternReducer
	VarPatternReducer
	TuplePatternReducer
	FieldVarPatternsReducer
	FieldVarPatternReducer
	OptionalTypeExprReducer
	AssignPatternReducer
	CasePatternReducer
	SequenceExprReducer
	ConditionReducer
	SwitchExprReducer
	SwitchExprBodyReducer
	LoopExprReducer
	LoopExprBodyReducer
	OptionalSequenceExprReducer
	ForAssignmentReducer
	TypeDefReducer
	PackageDefReducer
}

func NewReducer() *ReducerImpl {
	return &ReducerImpl{}
}
