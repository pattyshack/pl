package parser

import (
	"fmt"
)

const (
	DecimalInteger            = "decimal integer literal"
	HexadecimalInteger        = "hexadecimal integer literal"
	ZeroOPrefixedOctalInteger = "0o-prefixed octal integer literal"
	ZeroPrefixedOctalInteger  = "0-prefixed octal integer literal"
	BinaryInteger             = "binary integer literal"

	SingleLineString    = "single line string literal"
	MultiLineString     = "mutli line string literal"
	RawSingleLineString = "raw single line string literal"
	RawMultiLineString  = "raw mutli line string literal"

	DecimalFloat     = "decimal float literal"
	HexadecimalFloat = "hexadecimal float literal"
)

type Node interface {
	Loc() Location
	End() Location

	TreeString(indent string, label string) string

	PrependToLeading(CommentGroups)
	AppendToLeading(CommentGroups)

	PrependToTrailing(CommentGroups)
	AppendToTrailing(CommentGroups)

	// Return the node's original leading comment groups, and reset the node's
	// leading comment groups.
	TakeLeading() CommentGroups

	// Return the node's original trailing comment groups, and reset the node's
	// trailing comment groups.
	TakeTrailing() CommentGroups
}

type ValueNode interface {
	Node
	Val() string
}

type Expression interface {
	Node
	Statement
	Definition
	IsExpression()
}

type IsExpr struct {
	IsStmt
	IsDef
}

func (IsExpr) IsExpression() {}

type TypeExpression interface {
	Node
	IsTypeExpression()
}

type IsTypeExpr struct{}

func (IsTypeExpr) IsTypeExpression() {}

type TypeProperty interface {
	Node
	IsTypeProperty()
}

type IsTypeProp struct{}

func (IsTypeProp) IsTypeProperty() {}

type Statement interface {
	Node
	IsStatement()
}

type IsStmt struct{}

func (IsStmt) IsStatement() {}

type Definition interface {
	Node
	IsDefinition()
}

type IsDef struct{}

func (IsDef) IsDefinition() {}

// A comment group is a single block comment, or a group of line comments
// separated by single newlines (ignoring spaces).
type CommentGroup []ValueNode

func (CommentGroup) Id() SymbolId {
	return commentGroupToken
}

func (cg CommentGroup) Loc() Location {
	return cg[0].Loc()
}

func (cg CommentGroup) End() Location {
	return cg[len(cg)-1].End()
}

// A symbol's trailing comment groups are comment groups that are on the same
// line as the symbol and immediately follows the symbol.
//
// A symbol's leading comment groups are any comment group immediately before
// the symbol that do not belong to previous symbols.
//
// For example,
//
//	OTHER_SYMBOL /* group 0 */
//	// group 1 line 1
//	// group 1 line 2
//
//	// group 2
//	/* group 3 */ SYMBOL /* group 4 */ // group 5 line 1
//	// group 5 line 2
//
// /* other group 6 */ OTHER_SYMBOL
//
// SYMBOL's leading comment groups are {1, 2, 3}.  SYMBOL's trailing comment
// groups are {4, 5}. Comment groups {0, 6} belong to other symbols.
type CommentGroups struct {
	Groups []CommentGroup
}

func (CommentGroups) Id() SymbolId { return CommentGroupsToken }

func (cgs CommentGroups) Loc() Location {
	return cgs.Groups[0].Loc()
}

func (cgs CommentGroups) End() Location {
	return cgs.Groups[len(cgs.Groups)-1].End()
}

func (cgs *CommentGroups) Prepend(other CommentGroups) {
	cgs.Groups = append(other.Groups, cgs.Groups...)
}

func (cgs *CommentGroups) Append(other CommentGroups) {
	cgs.Groups = append(cgs.Groups, other.Groups...)
}

type LeadingTrailingComments struct {
	LeadingComment  CommentGroups
	TrailingComment CommentGroups
}

func (s *LeadingTrailingComments) PrependToLeading(other CommentGroups) {
	s.LeadingComment.Prepend(other)
}

func (s *LeadingTrailingComments) AppendToLeading(other CommentGroups) {
	s.LeadingComment.Append(other)
}

func (s *LeadingTrailingComments) TakeLeading() CommentGroups {
	tmp := s.LeadingComment
	s.LeadingComment = CommentGroups{}
	return tmp
}

func (s *LeadingTrailingComments) PrependToTrailing(other CommentGroups) {
	s.TrailingComment.Prepend(other)
}

func (s *LeadingTrailingComments) AppendToTrailing(other CommentGroups) {
	s.TrailingComment.Append(other)
}

func (s *LeadingTrailingComments) TakeTrailing() CommentGroups {
	tmp := s.TrailingComment
	s.TrailingComment = CommentGroups{}
	return tmp
}

type StartEndPos struct {
	StartPos Location
	EndPos   Location
}

func NewStartEndPos(start Location, end Location) StartEndPos {
	return StartEndPos{
		StartPos: start,
		EndPos:   end,
	}
}

func (sep StartEndPos) Loc() Location {
	return sep.StartPos
}

func (sep StartEndPos) End() Location {
	return sep.EndPos
}

type TokenCount struct {
	SymbolId
	StartEndPos
	Count int
}

func (s TokenCount) Id() SymbolId {
	return s.SymbolId
}

func (s TokenCount) String() string {
	return fmt.Sprintf("[Symbol:%s Count=%d]", s.SymbolId, s.Count)
}

type TokenValue struct {
	SymbolId
	StartEndPos
	LeadingTrailingComments

	// The value string is optional if the value is fully determined by SymbolId.
	Value string
	// Used for classifying literal subtypes.
	SubType string
}

func (s TokenValue) Id() SymbolId {
	return s.SymbolId
}

func (s TokenValue) Val() string {
	return s.Value
}

func (s TokenValue) TreeString(indent string, label string) string {
	return fmt.Sprintf(
		"%s%s[TokenValue: SymbolId=%s Value=%s]",
		indent,
		label,
		s.SymbolId,
		s.Value)
}

type ParseErrorSymbol struct {
	IsExpr
	IsTypeExpr
	IsDef

	StartEndPos
	LeadingTrailingComments

	Error error
}

func (ParseErrorSymbol) Id() SymbolId {
	return ParseErrorToken
}

func (s ParseErrorSymbol) TreeString(indent string, label string) string {
	return fmt.Sprintf(
		"%s%s[ParseError: %s %s]",
		indent,
		label,
		s.Error,
		s.StartPos)
}

type ParseErrorReducerImpl struct {
	ParseErrors []*ParseErrorSymbol
}

var _ ParseErrorExprReducer = &ParseErrorReducerImpl{}
var _ ParseErrorTypeExprReducer = &ParseErrorReducerImpl{}

func (Reducer *ParseErrorReducerImpl) ToParseErrorExpr(
	pe ParseErrorSymbol,
) (Expression, error) {
	ptr := &pe
	Reducer.ParseErrors = append(Reducer.ParseErrors, ptr)
	return ptr, nil
}

func (Reducer *ParseErrorReducerImpl) ToParseErrorTypeExpr(
	pe ParseErrorSymbol,
) (TypeExpression, error) {
	ptr := &pe
	Reducer.ParseErrors = append(Reducer.ParseErrors, ptr)
	return ptr, nil
}

type NodeList[T Node] struct {
	StartEndPos

	// The leading / trailing comments from the list's start/end marker tokens.
	LeadingTrailingComments

	// MiddleComment is only used when the list is empty, but contains comments.
	MiddleComment CommentGroups

	Elements []T

	// Only used for TreeString() / String()
	ListType string
}

func (list NodeList[T]) elementsString(indent string) string {
	result := ""
	for idx, element := range list.Elements {
		result += "\n" + element.TreeString(
			indent,
			fmt.Sprintf("Element%d=", idx))
	}
	return result
}

func (list NodeList[T]) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[%s:", indent, label, list.ListType)
	if len(list.Elements) == 0 {
		return result + "]"
	}

	result += list.elementsString(indent + "  ")
	result += "\n" + indent + "]"
	return result
}

func NewNodeList[T Node](listType string) *NodeList[T] {
	return &NodeList[T]{
		ListType: listType,
	}
}

func (list *NodeList[T]) Add(element T) {
	if len(list.Elements) == 0 {
		list.StartPos = element.Loc()
	}
	list.EndPos = element.End()
	list.Elements = append(list.Elements, element)
}

func (list *NodeList[T]) ReduceAdd(separator ValueNode, element T) {
	prev := list.Elements[len(list.Elements)-1]
	prev.AppendToTrailing(separator.TakeLeading())
	prev.AppendToTrailing(separator.TakeTrailing())

	list.Add(element)
}

func (list *NodeList[T]) ReduceImproper(separator ValueNode) {
	list.EndPos = separator.End()

	lastElement := list.Elements[len(list.Elements)-1]
	lastElement.AppendToTrailing(separator.TakeLeading())
	lastElement.AppendToTrailing(separator.TakeTrailing())
}

func (list *NodeList[T]) ReduceMarkers(start ValueNode, end ValueNode) {
	list.StartPos = start.Loc()
	list.EndPos = end.End()

	list.LeadingComment = start.TakeLeading()
	if len(list.Elements) > 0 {
		list.Elements[0].PrependToLeading(start.TakeTrailing())
		list.Elements[len(list.Elements)-1].AppendToTrailing(end.TakeLeading())
	} else {
		list.MiddleComment = start.TakeTrailing()
		list.MiddleComment.Append(end.TakeLeading())
	}
	list.TrailingComment = end.TakeTrailing()
}
