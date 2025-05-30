package controller

import (
	sc "strconv"

	ml "github.com/andreigrob/web_quiz_andrei/model"
)

type toItem[T ml.IEntity] = func(*Req) (_ T)

func toArticle(req *Req) (_ *ml.Article) {
	return &ml.Article{
		Name:    req.FormValue("name"),
		Email:   req.FormValue("email"),
		Message: req.FormValue("message"),
	}
}

func toAnswer(req *Req) (_ *ml.Answer) {
	return &ml.Answer{
		Name:     req.FormValue("name"),
		Email:    req.FormValue("email"),
		MessageA: req.FormValue("message0"),
		MessageB: req.FormValue("message1"),
		MessageC: req.FormValue("message2"),
	}
}

func toComment(req *Req) (_ *ml.Comment) {
	articleId, _ := sc.ParseInt(req.FormValue("articleId"), 10, 64)
	return &ml.Comment{
		Name:      req.FormValue("name"),
		Email:     req.FormValue("email"),
		Message:   req.FormValue("message"),
		ArticleId: articleId,
	}
}
