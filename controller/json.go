package controller

import (
	js "encoding/json"
	"log"
	ht "net/http"

	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

func writeResponse(wr ResWr, status int, message []byte) (e error) {
	wr.WriteHeader(status)
	if _, e = wr.Write(message); e != nil {
		ht.Error(wr, "Failed to write response", ht.StatusInternalServerError)
		log.Printf("Failed to write response: %v", e.Error())
		return
	}
	log.Printf("Wrote response (%s)", message)
	return
}

func encode[T any](wr ResWr, items []T) (e error) {
	name := ut.Name(items)
	if e = js.NewEncoder(wr).Encode(items); e != nil {
		ht.Error(wr, "Failed to encode "+name, ht.StatusInternalServerError)
		log.Printf("Failed to encode %s: %v", name, e.Error())
		return
	}
	log.Printf("Encoded %d %ss", len(items), name)
	return
}

func decode[T any](wr ResWr, req *Req, item *T) (e error) {
	name := ut.Name(item)
	if e = js.NewDecoder(req.Body).Decode(&item); e != nil {
		ht.Error(wr, "Failed to decode "+name, ht.StatusInternalServerError)
		log.Printf("Failed to decode %s: %v", name, e.Error())
		return
	}
	log.Printf("Decoded %s (%s)", name, item)
	return
}
