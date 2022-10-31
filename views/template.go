package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseTemplateFile(filePath string) (Template, error) {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("parse template file %s: %w", filePath, err)
	}
	return Template{
		htmlTpl: t,
	}, nil
}

func (t Template) Execute(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(rw, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(rw, "Failed to execute template file", http.StatusInternalServerError)
		return
	}
}
