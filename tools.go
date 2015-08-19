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
	http.HandleFunc("/md5/", md5Handler)
	http.HandleFunc("/sha3256/", sha3Handler)
	http.HandleFunc("/sha3512/", sha3512Handler)
	http.HandleFunc("/encodeBase64/", encodeBase64Handler)
	http.HandleFunc("/decodeBase64/", decodeBase64Handler)
	http.HandleFunc("/passgen/", passgenHandler)
	http.HandleFunc("/encodeHex/", encodeHexHandler)
	http.HandleFunc("/decodeHex/", decodeHexHandler)
	http.HandleFunc("/ip/", ipHandler)
	http.HandleFunc("/checkport/", openportHandler)

	http.ListenAndServe(":1337", nil)
}
