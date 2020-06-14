package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	fileExt    = ".html"
	filePrefix = "views/"
	layOutPath = "views/layout/"
)

// Views defines the shape of the views
type Views struct {
	t        *template.Template
	template string
}

func handleExts(fileNames []string) {
	for i, f := range fileNames {
		fileNames[i] = filePrefix + f + fileExt
	}
}

func layOutFiles() []string {
	files, err := filepath.Glob(layOutPath + "*")
	if err != nil {
		panic(err)
	}
	return files
}

// NewView returns the view struct
func NewView(layout string, fileNames ...string) *Views {
	handleExts(fileNames)
	fileNames = append(fileNames, layOutFiles()...)
	t, err := template.ParseFiles(fileNames...)
	if err != nil {
		panic(err)
	}
	return &Views{
		t:        t,
		template: layout,
	}
}

// Render is a function that renders a template
func (v *Views) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	data = SetData(r, data)
	v.t.ExecuteTemplate(w, v.template, data)
}
