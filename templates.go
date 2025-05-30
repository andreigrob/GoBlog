package main

import (
	tp "html/template"
	"log"
	"os"
	fp "path/filepath"
)

const (
	viewDir    = "view"
	nTemplates = 100
)

var (
	templatePaths []string
)

func loadTemplates() (err error) {
	templatePaths = make([]string, 0, nTemplates)
	var files []os.DirEntry
	if files, err = os.ReadDir(viewDir); err != nil {
		log.Fatalf("Error reading view directory: %v", err)
		return err
	}
	var i int
	for i = range files {
		if files[i].IsDir() {
			continue
		}
		templatePaths = append(templatePaths, fp.Join(viewDir, files[i].Name()))
	}
	return
}

func parseTemplates() (tmpl *tp.Template, err error) {
	if tmpl, err = tp.ParseFiles(templatePaths...); err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	return
}
