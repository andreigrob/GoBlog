package controller

import (
	ft "fmt"
	ht "net/http"

	ml "github.com/andreigrob/web_quiz_andrei/model"
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

type ResWr = ut.ResWr // ResponseWriter
type Req = ut.Req     // Request

// Read
// GET articles.json
func (c *FormCrT) ArticlesJsonGet(wr ResWr, req *Req) {
	var articles []*ml.Article
	_ = get(c, wr, req, articles)
}

// GET comments.json
func (c *FormCrT) CommentsJsonGet(wr ResWr, req *Req) {
	var comments []*ml.Comment
	_ = get(c, wr, req, comments)
}

// GET answers.json
func (c *FormCrT) AnswersJsonGet(wr ResWr, req *Req) {
	var answers []*ml.Answer
	_ = get(c, wr, req, answers)
}

// get gets a list of Entities from the database and writes them as JSON.
func get[T any, P ml.EntityPI[T]](c *FormCrT, wr ResWr, _ *Req, items []P) (err error) {
	if err = find(c, wr, &items); err != nil {
		return
	}
	if err = encode(wr, items); err != nil {
		return
	}
	return
}

// Create
// POST article.json
func (c *FormCrT) ArticleJsonPost(wr ResWr, req *Req) {
	var article ml.Article
	_ = post(c, wr, req, &article)
}

// POST answer.json
func (c *FormCrT) AnswerJsonPost(wr ResWr, req *Req) {
	var answer ml.Answer
	_ = post(c, wr, req, &answer)
}

// POST comment.json
func (c *FormCrT) CommentJsonPost(wr ResWr, req *Req) {
	var comment ml.Comment
	_ = post(c, wr, req, &comment)
}

// post reads a JSON object from the request, saves it to the database, and writes a REST response.
func post[T ml.EntityI](c *FormCrT, wr ResWr, req *Req, item T) (err error) {
	if err = decode(wr, req, &item); err != nil {
		return
	}
	var id int64
	if id, err = save(c, wr, item); err != nil {
		return
	}
	name := item.Class().GetNameLower()
	message := ft.Appendf(make([]byte, 0, 50), "%s saved successfully with id %d", name, id)
	if err = writeResponse(wr, ht.StatusCreated, message); err != nil {
		return
	}
	return
}
