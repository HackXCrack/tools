package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Request string
	Body    []byte
}

var templates_path = "./templates/"
var templates = template.Must(template.ParseGlob(templates_path + "*"))

func renderTemplate(w http.ResponseWriter, tmpl string, t string, r string) {
	err := templates.ExecuteTemplate(w, tmpl, &Page{Title: t, Request: r})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
		return
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "index", "")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/md5/", makeHandler(md5Handler))
	http.HandleFunc("/sha3256/", makeHandler(sha3Handler))
	http.HandleFunc("/sha3512/", makeHandler(sha3512Handler))
	http.HandleFunc("/encodeBase64/", makeHandler(encodeBase64Handler))
	http.HandleFunc("/decodeBase64/", makeHandler(decodeBase64Handler))
	http.HandleFunc("/passgen/", makeHandler(passgenHandler))
	http.HandleFunc("/encodeHex/", makeHandler(encodeHexHandler))
	http.HandleFunc("/decodeHex/", makeHandler(decodeHexHandler))
	http.HandleFunc("/ip/", makeHandler(ipHandler))

	http.ListenAndServe(":1337", nil)
}
