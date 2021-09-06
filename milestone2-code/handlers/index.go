package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Index(w http.ResponseWriter, r *http.Request) {
	htmlPath := "static/html"
	tpl := template.Must(template.ParseGlob(filepath.Join(htmlPath, "*.html")))
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}
}
