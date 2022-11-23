package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./assets/*.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", nil)
	})

	http.Handle("/assets/", http.StripPrefix("/assets",
		http.FileServer(http.Dir("./assets"))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
