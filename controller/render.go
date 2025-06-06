package controller

import (
	"log"
	ht "net/http"
)

func (c *ControllerT) render(wr ResWr, view string, data any) (e error) {
	log.Printf("render")
	log.Printf("Rendering form template (%s)", view)
	log.Printf("data: %v", data)
	log.Printf("c.tmpl: %v", c.tmpl)
	if e = c.tmpl.ExecuteTemplate(wr, view, data); e != nil {
		ht.Error(wr, "Failed to render form template", ht.StatusInternalServerError)
		log.Printf("Failed to render form template (%s): %v", view, e.Error())
		return
	}
	log.Printf("Rendered form template (%s)", view)
	return
}
