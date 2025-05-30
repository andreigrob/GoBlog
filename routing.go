package main

/*import (
	ht "net/http"

	cr "github.com/andreigrob/web_quiz_andrei/controller"
	ml "github.com/andreigrob/web_quiz_andrei/model"
)

type Route struct {
	Path    string
	Handler ht.HandlerFunc
}

var (
	routes []Route
)

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
// Write this using the default router
func addRoutes(router *ht.ServeMux) {
	var i int
	for ; i < len(routes); i++ {
		router.HandleFunc(routes[i].Path, routes[i].Handler)
	}
}
*/