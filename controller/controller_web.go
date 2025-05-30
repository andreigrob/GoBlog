package controller

import (
	"log"
	ht "net/http"

	ml "github.com/andreigrob/web_quiz_andrei/model"
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

// View Form
/*
// GET article.html
func (c *FormCr) ArticleForm(w RW, _ *Req) {
	renderForm[ml.Article](c, w)
}

// GET comment.html
func (c *FormCr) CommentForm(w RW, _ *Req) {
	renderForm[ml.Comment](c, w)
}

// GET answer.html
func (c *FormCr) AnswerForm(w RW, _ *Req) {
	renderForm[ml.Answer](c, w)
}
*/

// FormHandler returns a handler that renders a form template for an Entity.
func FormHandler[T ml.IEntity](c *FormCr) (_ ht.HandlerFunc) {
	return func(w RW, _ *Req) {
		renderForm[T](c, w)
	}
}

// renderForm renders a form template for an Entity.
func renderForm[T ml.IEntity](c *FormCr, w RW) {
	var val T
	_ = c.writeForm(w, val.Class().GetNameLower())
}

// writeForm renders a form template and writes it to the response.
func (c *FormCr) writeForm(w RW, name string) (err error) {
	if err = c.render(w, name+".html", nil); err != nil {
		return
	}
	return
}

var to = map[string]any{
	"article": toArticle,
	"comment": toComment,
	"answer":  toAnswer,
}

func SubmitHandler[T ml.IEntity](c *FormCr) (_ ht.HandlerFunc) {
	return func(w RW, req *Req) {
		var val T
		_ = submit(c, w, req, to[val.Class().GetNameLower()].(toItem[T]))
	}
}

// Submit Form
// POST article.html
func (c *FormCr) SubmitArticle(w RW, req *Req) {
	_ = submit(c, w, req, toArticle)
}

// POST comment.html
func (c *FormCr) SubmitComment(w RW, req *Req) {
	_ = submit(c, w, req, toComment)
}

// POST answer.html
func (c *FormCr) SubmitAnswer(w RW, req *Req) {
	_ = submit(c, w, req, toAnswer)
}

// submit reads values from the request, saves them to the database, and redirects to the list page.
func submit[T ml.IEntity](c *FormCr, w RW, req *Req, toT toItem[T]) (err error) {
	item := toT(req)
	name := ut.NameLower(item)
	newUrl := "/" + name + "s.html?success=true"
	log.Printf("newUrl: %s", newUrl)
	var id int64
	if id, err = save(c, w, item); err != nil {
		return
	}
	log.Printf("Saved %s with id %d, redirecting to (%s)", name, id, newUrl)
	ht.Redirect(w, req, newUrl, ht.StatusSeeOther)
	return
}

// List
func (c *FormCr) ArticlesHtmlGet(w RW, _ *Req) {
	const view = "articles.html"
	var articles []*ml.Article
	if err := find(c, w, &articles); err != nil {
		log.Printf("Failed to find articles: %v", err)
		return
	}
	if c.render(w, view, articles) != nil {
		return
	}
}

func (c *FormCr) CommentsHtmlGet(w RW, _ *Req) {
	var comments []*ml.Comment
	_ = writeList(c, w, comments)
}

func (c *FormCr) AnswersHtmlGet(w RW, _ *Req) {
	var answers []*ml.Answer
	_ = writeList(c, w, answers)
}

func writeList[T ml.IEntity](c *FormCr, w RW, items []T) (err error) {
	var val T
	Class := val.Class()
	name := Class.GetNameLower()
	view := name + "s.html"
	log.Printf("Writing view %s", view)
	if err = find(c, w, &items); err != nil {
		return
	}
	log.Printf("Found %d %s", len(items), name)
	if err = c.render(w, view, items); err != nil {
		return
	}
	log.Printf("Rendered %s", view)
	return
}
