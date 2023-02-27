package main

import (
	"html/template"
	"net/http"
)

// // Page représente une page HTML
// type Page struct {
// 	Name string
// 	Body string
// }

// // Les pages HTML
// var pages = map[string]*Page{
// 	"page1": &Page{
// 		Name: "Page 1",
// 		Body: "<h1>Page 1</h1><p>Redirection vers <a href=\"/page2\">page 2</a></p>",
// 	},
// 	"page2": &Page{
// 		Name: "Page 2",
// 		Body: "<h1>Page 2</h1><p>Redirection vers <a href=\"/page3\">page 3</a></p>",
// 	},
// 	"page3": &Page{
// 		Name: "Page 3",
// 		Body: "<h1>Page 3</h1><p>Redirection vers <a href=\"/\">page d'accueil</a></p>",
// 	},
// }

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	tmpl1 := template.Must(template.ParseFiles("page1.html"))
	tmpl2 := template.Must(template.ParseFiles("page2.html"))
	tmpl3 := template.Must(template.ParseFiles("page3.html"))

	http.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/page1" {
			tmpl1.Execute(w, nil)
		} else {
			tmpl := template.Must(template.ParseFiles("404.html"))
			tmpl.Execute(w, nil)
		}
	})
	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/page2" {
			tmpl2.Execute(w, nil)
		} else {
			tmpl := template.Must(template.ParseFiles("404.html"))
			tmpl.Execute(w, nil)
		}
	})
	http.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/page3" {
			tmpl3.Execute(w, nil)
		} else {
			tmpl := template.Must(template.ParseFiles("404.html"))
			tmpl.Execute(w, nil)
		}
	})

	http.ListenAndServe(":8080", nil)
}
