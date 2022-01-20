package controllers

import (
	"html/template"
	"net/http"
	"path"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Welcome In My Website...",
		"desc":  "Enjoy with this page...",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
