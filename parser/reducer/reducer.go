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

func (reducer *Reducer) MergeFrom(other *Reducer) {
	reducer.ParseErrors = append(reducer.ParseErrors, other.ParseErrors...)
	reducer.BoolLiteralExprs = append(
		reducer.BoolLiteralExprs,
		other.BoolLiteralExprs...)
	reducer.IntLiteralExprs = append(
		reducer.IntLiteralExprs,
		other.IntLiteralExprs...)
	reducer.FloatLiteralExprs = append(
		reducer.FloatLiteralExprs,
		other.FloatLiteralExprs...)
	reducer.RuneLiteralExprs = append(
		reducer.RuneLiteralExprs,
		other.RuneLiteralExprs...)
	reducer.StringLiteralExprs = append(
		reducer.StringLiteralExprs,
		other.StringLiteralExprs...)
	reducer.NamedExprs = append(reducer.NamedExprs, other.NamedExprs...)
	reducer.AccessExprs = append(reducer.AccessExprs, other.AccessExprs...)
	reducer.UnaryExprs = append(reducer.UnaryExprs, other.UnaryExprs...)
	reducer.BinaryExprs = append(reducer.BinaryExprs, other.BinaryExprs...)
	reducer.ImplicitStructExprs = append(
		reducer.ImplicitStructExprs,
		other.ImplicitStructExprs...)
	reducer.CallExprs = append(reducer.CallExprs, other.CallExprs...)
	reducer.IndexExprs = append(reducer.IndexExprs, other.IndexExprs...)
	reducer.AsExprs = append(reducer.AsExprs, other.AsExprs...)
	reducer.InitializeExprs = append(
		reducer.InitializeExprs,
		other.InitializeExprs...)
	reducer.IfExprs = append(reducer.IfExprs, other.IfExprs...)
	reducer.SwitchExprs = append(reducer.SwitchExprs, other.SwitchExprs...)
	reducer.SelectExprs = append(reducer.SelectExprs, other.SelectExprs...)
	reducer.LoopExprs = append(reducer.LoopExprs, other.LoopExprs...)
	reducer.SliceTypeExprs = append(
		reducer.SliceTypeExprs,
		other.SliceTypeExprs...)
	reducer.ArrayTypeExprs = append(
		reducer.ArrayTypeExprs,
		other.ArrayTypeExprs...)
	reducer.MapTypeExprs = append(reducer.MapTypeExprs, other.MapTypeExprs...)
	reducer.InferredTypeExprs = append(
		reducer.InferredTypeExprs,
		other.InferredTypeExprs...)
	reducer.NamedTypeExprs = append(
		reducer.NamedTypeExprs,
		other.NamedTypeExprs...)
	reducer.UnaryTypeExprs = append(
		reducer.UnaryTypeExprs,
		other.UnaryTypeExprs...)
	reducer.BinaryTypeExprs = append(
		reducer.BinaryTypeExprs,
		other.BinaryTypeExprs...)
	reducer.StructTypeExprs = append(
		reducer.StructTypeExprs,
		other.StructTypeExprs...)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, other.EnumTypeExprs...)
	reducer.TraitTypeExprs = append(
		reducer.TraitTypeExprs,
		other.TraitTypeExprs...)
	reducer.TypeDefs = append(reducer.TypeDefs, other.TypeDefs...)
	reducer.ImportClauses = append(reducer.ImportClauses, other.ImportClauses...)
	reducer.FuncTypeExprs = append(reducer.FuncTypeExprs, other.FuncTypeExprs...)
	reducer.FuncDefinitions = append(
		reducer.FuncDefinitions,
		other.FuncDefinitions...)
	reducer.MethodDefinitions = append(
		reducer.MethodDefinitions,
		other.MethodDefinitions...)
	reducer.AnonymousFuncExprs = append(
		reducer.AnonymousFuncExprs,
		other.AnonymousFuncExprs...)
}
