package main

import (
	ct "context"
	"log"
	ht "net/http"

	cr "github.com/andreigrob/web_quiz_andrei/controller"
	hr "github.com/andreigrob/web_quiz_andrei/handler"
)

const port = ":8080"

func start(handler ht.Handler) (err error) {
	log.Printf("Server started at %s", port)
	if err = ht.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}
	return
}

func main() {
	dbStr := getConnectionStr()
	conn, _ := connectToDb(dbStr)
	defer conn.Close(ct.Background())

	_ = loadTemplates()
	tmpl, _ := parseTemplates()

	formCr := cr.FormCr{}
	formCr.Init(conn, tmpl)

	var handler hr.Handler
	handler.Init(formCr)

	_ = start(&handler)
}
