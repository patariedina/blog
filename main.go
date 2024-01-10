package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	var fileName = "home.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template")
		return
	}
}
func es_lambda(w http.ResponseWriter, r *http.Request) {
	var fileName = "es_lambda.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template")
		return
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	var fileName = "about.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template")
		return
	}
}
func blog(w http.ResponseWriter, r *http.Request) {
	var fileName = "blog.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template")
		return
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/es_lambda":
		es_lambda(w, r)
	case "/home":
		home(w, r)
	case "/about":
		about(w, r)
	case "/blog":
		blog(w, r)
	default:
		fmt.Fprint(w, "hello world")
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe("", nil)
}
