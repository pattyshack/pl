package ast

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/errors"
)

type Visitor interface {
	// Call before visiting any children.
	Enter(Node)

	// Call after visiting all children.
	Exit(Node)
}

type Commentable interface {
	PrependToLeading(CommentGroups)
	AppendToLeading(CommentGroups)

	PrependToTrailing(CommentGroups)
	AppendToTrailing(CommentGroups)

	TakeComments() LeadingTrailingComments

	// Return the node's original leading comment groups, and reset the node's
	// leading comment groups.
	TakeLeading() CommentGroups

	// Return the node's original trailing comment groups, and reset the node's
	// trailing comment groups.
	TakeTrailing() CommentGroups
}

type TokenValue interface {
	Node

	Val() string
}

type Node interface {
	lexutil.Locatable
	Commentable

	Walk(Visitor)
}

// A Node may optionally implement Validator, which non-recursively validate
// the node's basic syntactic structure (without parent node information).
type Validator interface {
	Validate(*errors.Emitter)
}

type Expression interface {
	Node
	Statement
	IsExpression()
}

type ControlFlowExpr interface {
	Expression
	SetLabel(string)
}

type IsExpr struct {
	IsStmt
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

type IsStmt struct {
}

func (IsStmt) IsStatement() {}

type DirectiveExpression interface {
	Node
	IsDirectiveExpression()
}

type IsDirectiveExpr struct{}

func (IsDirectiveExpr) IsDirectiveExpression() {}

// A comment group is a single block comment, or a group of line comments
// separated by single newlines (ignoring spaces).
type CommentGroup struct {
	lexutil.StartEndPos
	Comments []string
}

func NewCommentGroup(value TokenValue) *CommentGroup {
	return &CommentGroup{
		StartEndPos: lexutil.NewStartEndPos(value.Loc(), value.End()),
		Comments:    []string{value.Val()},
	}
}

func (group *CommentGroup) Add(value TokenValue) {
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

func (s *LeadingTrailingComments) TakeComments() LeadingTrailingComments {
	tmp := *s
	*s = LeadingTrailingComments{}
	return tmp
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

// Control flow label
type Label string

func (label *Label) SetLabel(l string) {
	*label = Label(l)
}
