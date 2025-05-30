package handler

import (
	ht "net/http"

	cr "github.com/andreigrob/web_quiz_andrei/controller"
	ml "github.com/andreigrob/web_quiz_andrei/model"
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

type Routes = map[string]ht.HandlerFunc
type HandlerF = ht.HandlerFunc

type Handler struct {
	fc              cr.FormCr
	routesGet       Routes
	routesPost      Routes
	articleHtmlGet  ht.HandlerFunc
	commentHtmlGet  ht.HandlerFunc
	answerHtmlGet   ht.HandlerFunc
	articleHtmlPost ht.HandlerFunc
	commentHtmlPost ht.HandlerFunc
	answerHtmlPost  ht.HandlerFunc
}

func (h *Handler) Init(fc cr.FormCr) {
	h.fc = fc
	h.articleHtmlGet = cr.FormHandler[*ml.Article](&fc)
	h.commentHtmlGet = cr.FormHandler[*ml.Comment](&fc)
	h.answerHtmlGet = cr.FormHandler[*ml.Answer](&fc)
	h.articleHtmlPost = cr.SubmitHandler[*ml.Article](&fc)
	h.commentHtmlPost = cr.SubmitHandler[*ml.Comment](&fc)
	h.answerHtmlPost = cr.SubmitHandler[*ml.Comment](&fc)

	h.routesGet = Routes{
		"/article.html":  h.articleHtmlGet,
		"/":              h.commentHtmlGet,
		"/answer.html":   h.answerHtmlGet,
		"/articles.html": fc.ArticlesHtmlGet,
		"/comments.html": fc.CommentsHtmlGet,
		"/answers.html":  fc.AnswersHtmlGet,
		// API
		"/articles.json": fc.ArticlesJsonGet,
		"/comments.json": fc.CommentsJsonGet,
		"/answers.json":  fc.AnswersJsonGet,
	}

	h.routesPost = Routes{
		"/article.html": h.articleHtmlPost,
		"/":             h.commentHtmlPost,
		"/answer.html":  h.answerHtmlPost,
		// API
		"/article.json": fc.ArticleJsonPost,
		"/comment.json": fc.CommentJsonPost,
		"/answer.json":  fc.AnswerJsonPost,
	}
}

/*
func initRoutes(fc *cr.FormCr) {
	routes = []Route{
		// GET
		{"GET /article.html", cr.FormHandler[ml.Article](fc)},
		{"GET /articles.html", fc.ArticleList},
		{"GET /", cr.FormHandler[ml.Comment](fc)},
		{"GET /comments.html", fc.CommentList},
		{"GET /answer.html", cr.FormHandler[ml.Answer](fc)},
		{"GET /answers.html", fc.AnswerList},
		// API
		{"GET /article.json", fc.GetArticles},
		{"GET /comments.json", fc.GetComments},
		{"GET /answers.json", fc.GetAnswers},
		// POST
		{"POST /article.html", cr.SubmitHandler[ml.Article](fc)},
		{"POST /", cr.SubmitHandler[ml.Comment](fc)},
		{"POST /answer.html", cr.SubmitHandler[ml.Answer](fc)},
		// API
		{"POST /article.json", fc.PostArticle},
		{"POST /answer.json", fc.PostAnswer},
		{"POST /comment.json", fc.PostComment},
	}
}
*/

func (h *Handler) ServeHTTP(w ut.RW, req *ut.Req) {
	switch req.Method {
	case "GET":
		h.Get(w, req)
	case "POST":
		h.Post(w, req)
	}
}

func (h *Handler) handle(routes Routes, w ut.RW, req *ut.Req) {
	if handler, ok := routes[req.URL.Path]; ok {
		handler(w, req)
		return
	}
	w.WriteHeader(ht.StatusNotFound)
	_, _ = w.Write([]byte("404 Not Found"))
}

func (h *Handler) Get(w ut.RW, req *ut.Req) {
	h.handle(h.routesGet, w, req)
}

func (h *Handler) Post(w ut.RW, req *ut.Req) {
	h.handle(h.routesPost, w, req)
}
