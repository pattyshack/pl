package ast

import (
	"fmt"

	"github.com/pattyshack/gt/lexutil"
)

type Node interface {
	Loc() lexutil.Location
	End() lexutil.Location

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

type ValuedNode interface {
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
type CommentGroup struct {
	StartEndPos
	Comments []string
}

func NewCommentGroup(value ValuedNode) *CommentGroup {
	return &CommentGroup{
		StartEndPos: NewStartEndPos(value.Loc(), value.End()),
		Comments:    []string{value.Val()},
	}
}

func (group *CommentGroup) Add(value ValuedNode) {
	group.EndPos = value.End()
	group.Comments = append(group.Comments, value.Val())
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

func (cgs CommentGroups) Loc() lexutil.Location {
	return cgs.Groups[0].Loc()
}

func (cgs CommentGroups) End() lexutil.Location {
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
	StartPos lexutil.Location
	EndPos   lexutil.Location
}

func NewStartEndPos(start lexutil.Location, end lexutil.Location) StartEndPos {
	return StartEndPos{
		StartPos: start,
		EndPos:   end,
	}
}

func (sep StartEndPos) Loc() lexutil.Location {
	return sep.StartPos
}

func (sep StartEndPos) End() lexutil.Location {
	return sep.EndPos
}

type ParseErrorNode struct {
	IsExpr
	IsTypeExpr

	StartEndPos
	LeadingTrailingComments

	Errors []error
}

func NewParseErrorNode(startEnd StartEndPos, err error) *ParseErrorNode {
	return &ParseErrorNode{
		StartEndPos: startEnd,
		Errors:      []error{err},
	}
}

func (s ParseErrorNode) TreeString(indent string, label string) string {
	more := ""
	if len(s.Errors) > 1 {
		fmt.Sprintf("(and %d more errors) ", len(s.Errors)-1)
	}
	return fmt.Sprintf(
		"%s%s[ParseErrorNode: %s %s%s]",
		indent,
		label,
		s.Errors[0],
		more,
		s.StartPos)
}
