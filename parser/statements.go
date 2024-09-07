package parser

//
// Statements
//

type StatementList = NodeList[Statement]

type StatementListReducer struct{}

var _ ProperStatementsReducer = &StatementListReducer{}
var _ StatementsReducer = &StatementListReducer{}

func (StatementListReducer) AddImplicitToProperStatements(
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

func (StatementListReducer) AddExplicitToProperStatements(
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

func (StatementListReducer) StatementToProperStatements(
	statement Statement,
) (
	*StatementList,
	error,
) {
	return newNodeList[Statement]("StatementList", statement), nil
}

func (StatementListReducer) ImproperImplicitToStatements(
	statements *StatementList,
	newlines TokenCount,
) (
	*StatementList,
	error,
) {
	return statements, nil
}

func (StatementListReducer) ImproperExplicitToStatements(
	statements *StatementList,
	semicolon TokenValue,
) (
	*StatementList,
	error,
) {
	statements.reduceImproper(semicolon)
	return statements, nil
}

func (StatementListReducer) NilToStatements() (*StatementList, error) {
	return nil, nil
}
