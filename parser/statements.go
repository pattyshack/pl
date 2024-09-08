package parser

//
// Statements
//

type StatementList = NodeList[Statement]

type StatementListReducerImpl struct{}

var _ ProperStatementsReducer = &StatementListReducerImpl{}
var _ StatementsReducer = &StatementListReducerImpl{}

func (StatementListReducerImpl) AddImplicitToProperStatements(
	statements *StatementList,
	newlines TokenCount,
	statement Statement,
) (
	*StatementList,
	error,
) {
	statements.reduceAdd(TokenValue{}, statement)
	return statements, nil
}

func (StatementListReducerImpl) AddExplicitToProperStatements(
	statements *StatementList,
	semicolon TokenValue,
	statement Statement,
) (
	*StatementList,
	error,
) {
	statements.reduceAdd(semicolon, statement)
	return statements, nil
}

func (StatementListReducerImpl) StatementToProperStatements(
	statement Statement,
) (
	*StatementList,
	error,
) {
	return newNodeList[Statement]("StatementList", statement), nil
}

func (StatementListReducerImpl) ImproperImplicitToStatements(
	statements *StatementList,
	newlines TokenCount,
) (
	*StatementList,
	error,
) {
	return statements, nil
}

func (StatementListReducerImpl) ImproperExplicitToStatements(
	statements *StatementList,
	semicolon TokenValue,
) (
	*StatementList,
	error,
) {
	statements.reduceImproper(semicolon)
	return statements, nil
}

func (StatementListReducerImpl) NilToStatements() (*StatementList, error) {
	return nil, nil
}
