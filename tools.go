package main

import (
	"html/template"
	"net/http"
	"log"
)

type Page struct {
	Title   string
	Request string
	Body    []byte
}

var templates_path = "./templates/"
var templates = template.Must(template.ParseGlob(templates_path + "*"))
const PORT = "1337"

func renderTemplate(w http.ResponseWriter, tmpl string, t string, r string) {
	err := templates.ExecuteTemplate(w, tmpl, &Page{Title: t, Request: r})
	if err != nil {
		log.Print("ExecuteTemplate: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	log.Print("Escuchando en el puerto "+PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
