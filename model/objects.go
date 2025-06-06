package model

import (
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

type Rows = ut.Rows

func (*Article) New() (_ *Article) {
	return &Article{}
}

func (a *Article) NewAny() (_ any) {
	return a.New()
}

func (a *Article) Init(Name, Email, Message string) {
	a.Name = Name
	a.Email = Email
	a.Message = Message
}

func (a *Article) InitAny(args ...any) {
	a.Init(args[0].(string), args[1].(string), args[2].(string))
}

func (a *Article) Scan(rows Rows) (e error) {
	if e = rows.Scan(&a.Id, &a.Name, &a.Email, &a.Message); e != nil {
		return
	}
	return
}

func (*Comment) New() (_ *Comment) {
	return &Comment{}
}

func (c *Comment) NewAny() (_ any) {
	return c.New()
}

func (c *Comment) Init(Name, Email, Message string, ArticleId, CommentId int64) {
	c.Name = Name
	c.Email = Email
	c.Message = Message
	c.ArticleId = ArticleId
	c.CommentId = CommentId
}

func (c *Comment) InitAny(args ...any) {
	c.Init(args[0].(string), args[1].(string), args[2].(string), args[3].(int64), args[4].(int64))
}

func (c *Comment) Scan(rows Rows) (e error) {
	var ArticleId, CommentId *int64
	if e = rows.Scan(&c.Id, &c.Name, &c.Email, &c.Message, &ArticleId, &CommentId); e != nil {
		return
	}
	//comment.ArticleId = ut.Get[int64](ArticleId)
	//comment.CommentId = ut.Get[int64](CommentId)
	c.ArticleId = ut.Star(ArticleId)
	c.CommentId = ut.Star(CommentId)
	return
}

func (*Answer) New() (_ *Answer) {
	return &Answer{}
}

func (a *Answer) NewAny() (_ any) {
	return a.New()
}

func (a *Answer) Init(Name, Email, MessageA, MessageB, MessageC string) {
	a.Name = Name
	a.Email = Email
	a.MessageA = MessageA
	a.MessageB = MessageB
	a.MessageC = MessageC
}

func (a *Answer) InitAny(args ...any) {
	a.Init(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
}

func (a *Answer) Scan(rows Rows) (e error) {
	if e = rows.Scan(&a.Id, &a.Name, &a.Email, &a.MessageA, &a.MessageB, &a.MessageC); e != nil {
		return
	}
	return
}
