package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "hello world!", "layout", "public_navbar", "top")

	// t, _ := template.ParseFiles("app/views/templates/top.html")
	// t.Execute(w, "hello world!")
}
