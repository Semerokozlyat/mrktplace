package views

import (
	"bytes"
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
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
	tpl := template.New(fileNames[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField is not implemented")
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, fileNames...)
	if err != nil {
		return Template{}, fmt.Errorf("parse template file %s on FS: %w", fileNames, err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

//func ParseTemplateFile(filePaths ...string) (Template, error) {
//	t, err := template.ParseFiles(filePaths...)
//	if err != nil {
//		return Template{}, fmt.Errorf("parse template file %s: %w", filePaths, err)
//	}
//	return Template{
//		htmlTpl: t,
//	}, nil
//}

func (t Template) Execute(rw http.ResponseWriter, r *http.Request, data interface{}) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	htmlTpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("failed to clone template: %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	htmlTpl = htmlTpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)
	// TODO: remember this could consume much memory on large pages
	var buf bytes.Buffer
	err = htmlTpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(rw, "Failed to execute template file", http.StatusInternalServerError)
		return
	}
	io.Copy(rw, &buf)
}
