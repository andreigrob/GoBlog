package controller

import (
	"log"
	ht "net/http"
)

func (c *Controller) render(w RW, view string, data any) (err error) {
	if err = c.tmpl.ExecuteTemplate(w, view, data); err != nil {
		ht.Error(w, "Failed to render form template", ht.StatusInternalServerError)
		log.Printf("Failed to render form template (%s): %v", view, err)
		return
	}
	log.Printf("Rendered form template (%s)", view)
	return
}
