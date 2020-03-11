package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func landing(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", "Landing")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", "Dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
	tpl.ExecuteTemplate(w, "index.gohtml", "Me")
}

func main() {
	http.Handle("/", http.HandlerFunc(landing))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8000", nil)
}
