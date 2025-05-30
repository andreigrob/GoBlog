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

func (a *Article) Init(Name string, Email string, Message string) {
	a.Name = Name
	a.Email = Email
	a.Message = Message
}

func (a *Article) InitAny(args ...any) {
	a.Init(args[0].(string), args[1].(string), args[2].(string))
}

func (a *Article) Scan(rows Rows) (err error) {
	if err = rows.Scan(&a.Id, &a.Name, &a.Email, &a.Message); err != nil {
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

func (c *Comment) Init(Name string, Email string, Message string, ArticleId int64, CommentId int64) {
	c.Name = Name
	c.Email = Email
	c.Message = Message
	c.ArticleId = ArticleId
	c.CommentId = CommentId
}

func (c *Comment) InitAny(args ...any) {
	c.Init(args[0].(string), args[1].(string), args[2].(string), args[3].(int64), args[4].(int64))
}

func (c *Comment) Scan(rows Rows) (err error) {
	var ArticleId, CommentId *int64
	if err = rows.Scan(&c.Id, &c.Name, &c.Email, &c.Message, &ArticleId, &CommentId); err != nil {
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

func (a *Answer) Init(Name string, Email string, MessageA string, MessageB string, MessageC string) {
	a.Name = Name
	a.Email = Email
	a.MessageA = MessageA
	a.MessageB = MessageB
	a.MessageC = MessageC
}

func (a *Answer) InitAny(args ...any) {
	a.Init(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(string))
}

func (a *Answer) Scan(rows Rows) (err error) {
	if err = rows.Scan(&a.Id, &a.Name, &a.Email, &a.MessageA, &a.MessageB, &a.MessageC); err != nil {
		return
	}
	return
}
