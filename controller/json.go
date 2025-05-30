package controller

import (
	js "encoding/json"
	"log"
	ht "net/http"

	ml "github.com/andreigrob/web_quiz_andrei/model"
	ut "github.com/andreigrob/web_quiz_andrei/utils"
)

func writeResponse(w RW, status int, message []byte) (err error) {
	w.WriteHeader(status)
	if _, err = w.Write(message); err != nil {
		ht.Error(w, "Failed to write response", ht.StatusInternalServerError)
		log.Printf("Failed to write response: %v", err)
		return
	}
	log.Printf("Wrote response (%s)", message)
	return
}

func encode[T ml.IEntity](w RW, items []T) (err error) {
	name := ut.Name(items)
	if err = js.NewEncoder(w).Encode(items); err != nil {
		ht.Error(w, "Failed to encode "+name, ht.StatusInternalServerError)
		log.Printf("Failed to encode %s: %v", name, err)
		return
	}
	log.Printf("Encoded %d %ss", len(items), name)
	return
}

func decode[T any](w RW, req *Req, item *T) (err error) {
	name := ut.Name(item)
	if err = js.NewDecoder(req.Body).Decode(&item); err != nil {
		ht.Error(w, "Failed to decode "+name, ht.StatusInternalServerError)
		log.Printf("Failed to decode %s: %v", name, err)
		return
	}
	log.Printf("Decoded %s (%s)", name, item)
	return
}
