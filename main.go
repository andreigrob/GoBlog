package main

import (
	ct "context"
	"log"
	ht "net/http"

	cr "github.com/andreigrob/web_quiz_andrei/controller"
	hr "github.com/andreigrob/web_quiz_andrei/handler"
)

const port = `:8080`

func start(handler ht.Handler) (e error) {
	log.Printf("Server started at %s", port)
	if e = ht.ListenAndServe(port, handler); e != nil {
		log.Fatalf("Error starting server on port %s: %v", port, e)
	}
	return
}

func main() {
	dbStr := getConnectionStr()
	conn, _ := connectToDb(dbStr)
	defer conn.Close(ct.Background())

	tmpl = templates()
	log.Printf("tmpl: %v", tmpl)

	formCr := cr.FormCrT{}
	formCr.Init(conn, tmpl)

	var handler hr.Handler
	handler.Init(formCr)

	_ = start(&handler)
}
