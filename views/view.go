package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir 		string = "views/layouts/"
	TemplateExt 	string = ".gohtml"
)

type View struct {
	Template 	*template.Template
	Layout 		string
}

func layoutFiles() []string {
	files, error := filepath.Glob(LayoutDir + "*" + TemplateExt)

	if error != nil { panic(error) }

	return files
}

func NewView(layout string, files ...string) *View {
	files = append(files, 
		"views/layouts/default.gohtml",
		"views/layouts/header.gohtml",
		"views/layouts/footer.gohtml",
	)

	t, error := template.ParseFiles(files...)

	if error != nil {
		panic(error)
	}

	return &View {
		Template: t,
		Layout: layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
