package home

import (
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("./home/home.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}
