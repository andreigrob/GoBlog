package model

import (
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

// Entity Objects
var (
	ArticleObject EntityObject
	CommentObject EntityObject
	AnswerObject  EntityObject
)

// Initialize Entity Objects
func init() {
	ArticleObject.Init("article", "Articles", "Name, Email, Message")
	CommentObject.Init("comment", "Comments", "Name, Email, Message, ArticleId, CommentId")
	AnswerObject.Init("answer", "Answers", "Name, Email, MessageA, MessageB, MessageC")
}

// Entity is a base struct for all entities.
type Entity struct {
	Id int64
}

// Id
func (e *Entity) GetId() (_ int64) {
	return e.Id
}

func (e *Entity) SetId(id int64) {
	e.Id = id
}

// An Article is an Entity with a Name, Email, and Message.
type Article struct {
	Entity
	Name    string
	Email   string
	Message string
}

// Article Entity Object
func (*Article) Class() (_ *EntityObject) {
	return &ArticleObject
}

// Fields Array
func (a *Article) GetFields() (_ []any) {
	return []any{a.Name, a.Email, a.Message}
}

// A Comment is associated either with an Article or another Comment.
type Comment struct {
	Entity
	Name      string
	Email     string
	Message   string
	ArticleId int64
	CommentId int64
}

// Comment Entity Object
func (*Comment) Class() (_ *EntityObject) {
	return &CommentObject
}

// Fields Array
func (c *Comment) GetFields() (_ []any) {
	return []any{c.Name, c.Email, c.Message, ut.Nil(c.ArticleId), ut.Nil(c.CommentId)}
}

// Questionaire Answer
type Answer struct {
	Entity
	Name     string
	Email    string
	MessageA string
	MessageB string
	MessageC string
}

// Answer Entity Object
func (*Answer) Class() (_ *EntityObject) {
	return &AnswerObject
}

// Fields Array
func (a *Answer) GetFields() (_ []any) {
	return []any{a.Name, a.Email, a.MessageA, a.MessageB, a.MessageC}
}

// Interface Checks
var (
	_ IEntity = &Article{}
	_ IEntity = &Comment{}
	_ IEntity = &Answer{}
)
