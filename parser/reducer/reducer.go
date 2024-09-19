package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

type Reducer struct {
	ParseErrors []error

	BoolLiteralExprs    []*BoolLiteralExpr
	IntLiteralExprs     []*IntLiteralExpr
	FloatLiteralExprs   []*FloatLiteralExpr
	RuneLiteralExprs    []*RuneLiteralExpr
	StringLiteralExprs  []*StringLiteralExpr
	NamedExprs          []*NamedExpr
	AccessExprs         []*AccessExpr
	UnaryExprs          []*UnaryExpr
	BinaryExprs         []*BinaryExpr
	ImplicitStructExprs []*ImplicitStructExpr
	CallExprs           []*CallExpr
	IndexExprs          []*IndexExpr
	AsExprs             []*AsExpr
	InitializeExprs     []*InitializeExpr
	IfExprs             []*IfExpr
	SwitchExprs         []*SwitchExpr
	SelectExprs         []*SelectExpr
	LoopExprs           []*LoopExpr

	SliceTypeExprs    []*SliceTypeExpr
	ArrayTypeExprs    []*ArrayTypeExpr
	MapTypeExprs      []*MapTypeExpr
	InferredTypeExprs []*InferredTypeExpr
	NamedTypeExprs    []*NamedTypeExpr
	UnaryTypeExprs    []*UnaryTypeExpr
	BinaryTypeExprs   []*BinaryTypeExpr
	StructTypeExprs   []*PropertiesTypeExpr
	EnumTypeExprs     []*PropertiesTypeExpr
	TraitTypeExprs    []*PropertiesTypeExpr

	TypeDefs []*TypeDef

	ImportClauses []*ImportClause

	FuncTypeExprs []*FuncSignature

	FuncDefinitions    []*FuncDefinition
	MethodDefinitions  []*FuncDefinition
	AnonymousFuncExprs []*FuncDefinition
}

var _ lr.Reducer = &Reducer{}

func NewReducer() *Reducer {
	return &Reducer{}
}
