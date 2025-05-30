package controller

import (
	ft "fmt"
	ht "net/http"

	ml "github.com/andreigrob/web_quiz_andrei/model"
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

type RW = ut.RW   // ResponseWriter
type Req = ut.Req // Request

// Read
// GET articles.json
func (c *FormCr) ArticlesJsonGet(w RW, req *Req) {
	var articles []*ml.Article
	_ = get(c, w, req, articles)
}

// GET comments.json
func (c *FormCr) CommentsJsonGet(w RW, req *Req) {
	var comments []*ml.Comment
	_ = get(c, w, req, comments)
}

// GET answers.json
func (c *FormCr) AnswersJsonGet(w RW, req *Req) {
	var answers []*ml.Answer
	_ = get(c, w, req, answers)
}

// get gets a list of Entities from the database and writes them as JSON.
func get[T ml.IEntity](c *FormCr, w RW, _ *Req, items []T) (err error) {
	if err = find(c, w, &items); err != nil {
		return
	}
	if err = encode(w, items); err != nil {
		return
	}
	return
}

// Create
// POST article.json
func (c *FormCr) ArticleJsonPost(w RW, req *Req) {
	var article ml.Article
	_ = post(c, w, req, &article)
}

// POST answer.json
func (c *FormCr) AnswerJsonPost(w RW, req *Req) {
	var answer ml.Answer
	_ = post(c, w, req, &answer)
}

// POST comment.json
func (c *FormCr) CommentJsonPost(w RW, req *Req) {
	var comment ml.Comment
	_ = post(c, w, req, &comment)
}

// post reads a JSON object from the request, saves it to the database, and writes a REST response.
func post[T ml.IEntity](c *FormCr, w RW, req *Req, item T) (err error) {
	if err = decode(w, req, &item); err != nil {
		return
	}
	var id int64
	if id, err = save(c, w, item); err != nil {
		return
	}
	name := ut.Name(item)
	message := ft.Appendf(make([]byte, 0, 50), "%s saved successfully with id %d", name, id)
	if err = writeResponse(w, ht.StatusCreated, message); err != nil {
		return
	}
	return
}
