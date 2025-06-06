package main

import (
	tp "html/template"
	"log"
	"os"
	fp "path/filepath"
)

const (
	viewDir    = `views`
	nTemplates = 100
)

var (
	tmpl *tp.Template
)

func loadTemplates() (templatePaths []string, e error) {
	files, e := os.ReadDir(viewDir)
	if e != nil {
		log.Fatalf("Error reading view directory: %v", e)
		return
	}
	templatePaths = make([]string, 0, nTemplates)
	var i int = len(files) - 1
	for ; i >= 0; i-- {
		if files[i].IsDir() {
			continue
		}
		templatePaths = append(templatePaths, fp.Join(viewDir, files[i].Name()))
	}
	return
}

func parseTemplates(templatePaths []string) (t *tp.Template, e error) {
	log.Printf("parseTemplates: templatePaths: %v", templatePaths)
	if t, e = tp.ParseFiles(templatePaths...); e != nil {
		return
	}
	log.Printf("parseTemplates: tmpl: %v", tmpl)
	return
}

func getTemplates() (t *tp.Template, e error) {
	templatePaths, e := loadTemplates()
	if e != nil {
		log.Fatalf("Error loading templates: %v", e)
		return
	}
	t, e = parseTemplates(templatePaths)
	if e != nil {
		log.Fatalf("Error parsing templates: %v", e)
		return
	}
	log.Printf("getTemplates: t: %v", t)
	return
}

func templates() (_ *tp.Template) {
	if tmpl == nil {
		tmpl, _ = getTemplates()
	}
	log.Printf("templates: tmpl: %v", tmpl)
	return tmpl
}
