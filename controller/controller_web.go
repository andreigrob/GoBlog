package controller

import (
	"log"
	ht "net/http"

	ml "github.com/andreigrob/web_quiz_andrei/model"
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
func FormHandler[T ml.EntityI](c *FormCrT) (_ ht.HandlerFunc) {
	return func(wr ResWr, _ *Req) {
		renderForm[T](c, wr)
	}
}

// renderForm renders a form template for an Entity.
func renderForm[T ml.EntityI](c *FormCrT, wr ResWr) {
	var val T
	_ = c.writeForm(wr, val.Class().GetNameLower())
}

// writeForm renders a form template and writes it to the response.
func (c *FormCrT) writeForm(wr ResWr, name string) (err error) {
	if err = c.render(wr, name+".html", nil); err != nil {
		return
	}
	return
}

var to = map[string]any{
	`article`: toArticle,
	`comment`: toComment,
	`answer`:  toAnswer,
}

func SubmitHandler[T ml.EntityI](c *FormCrT) (_ ht.HandlerFunc) {
	return func(wr ResWr, req *Req) {
		var val T
		_ = submit(c, wr, req, to[val.Class().GetNameLower()].(toItem[T]))
	}
}

// Submit Form
// POST article.html
func (c *FormCrT) SubmitArticle(wr ResWr, req *Req) {
	_ = submit(c, wr, req, toArticle)
}

// POST comment.html
func (c *FormCrT) SubmitComment(wr ResWr, req *Req) {
	_ = submit(c, wr, req, toComment)
}

// POST answer.html
func (c *FormCrT) SubmitAnswer(wr ResWr, req *Req) {
	_ = submit(c, wr, req, toAnswer)
}

// submit reads values from the request, saves them to the database, and redirects to the list page.
func submit[T ml.EntityI](c *FormCrT, wr ResWr, req *Req, toT toItem[T]) (err error) {
	item := toT(req)
	name := item.Class().GetNameLower()
	newUrl := "/" + name + "s.html?success=true"
	log.Printf("newUrl: %s", newUrl)
	var id int64
	if id, err = save(c, wr, item); err != nil {
		return
	}
	log.Printf("Saved %s with id %d, redirecting to (%s)", name, id, newUrl)
	ht.Redirect(wr, req, newUrl, ht.StatusSeeOther)
	return
}

// List
func (c *FormCrT) ArticlesHtmlGet(wr ResWr, _ *Req) {
	const view = `articles.html`
	var articles []*ml.Article
	if err := find(c, wr, &articles); err != nil {
		log.Printf("Failed to find articles: %v", err)
		return
	}
	var Len = len(articles)
	articleDTOs := make([]*ml.ArticleDTO, Len)
	var i int
	for ; i < Len; i++ {
		articleDTOs[i] = &ml.ArticleDTO{
			Article:  *articles[i],
			Comments: make([]ml.Comment, 0),
		}
	}
	if c.render(wr, view, articleDTOs) != nil {
		return
	}
}

func (c *FormCrT) CommentsHtmlGet(wr ResWr, _ *Req) {
	var comments []*ml.Comment
	_ = writeList(c, wr, comments)
}

func (c *FormCrT) AnswersHtmlGet(wr ResWr, _ *Req) {
	var answers []*ml.Answer
	_ = writeList(c, wr, answers)
}

func writeList[T any, P ml.EntityPI[T]](c *FormCrT, wr ResWr, items []P) (err error) {
	var val P
	Class := val.Class()
	name := Class.GetNameLower()
	view := name + "s.html"
	log.Printf("Writing view %s", view)
	if err = find(c, wr, &items); err != nil {
		return
	}
	log.Printf("Found %d %s", len(items), name)
	if err = c.render(wr, view, items); err != nil {
		return
	}
	log.Printf("Rendered %s", view)
	return
}
