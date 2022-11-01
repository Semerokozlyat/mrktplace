package views

import (
	"fmt"
	"html/template"
	"io/fs"
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

func ParseTemplateFS(fs fs.FS, fileNames ...string) (Template, error) {
	t, err := template.ParseFS(fs, fileNames...)
	if err != nil {
		return Template{}, fmt.Errorf("parse template file %s on FS: %w", fileNames, err)
	}
	return Template{
		htmlTpl: t,
	}, nil
}

func ParseTemplateFile(filePaths ...string) (Template, error) {
	t, err := template.ParseFiles(filePaths...)
	if err != nil {
		return Template{}, fmt.Errorf("parse template file %s: %w", filePaths, err)
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
