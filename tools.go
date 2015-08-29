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

func RenderTemplate(w http.ResponseWriter, tmpl string, t string, r string) {
	err := templates.ExecuteTemplate(w, tmpl, &Page{Title: t, Request: r})
	if err != nil {
		log.Print("ExecuteTemplate: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index", "index", "")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/md5/", Md5Handler)
	http.HandleFunc("/sha3256/", Sha3Handler)
	http.HandleFunc("/sha3512/", Sha3512Handler)
	http.HandleFunc("/encodeBase64/", EncodeBase64Handler)
	http.HandleFunc("/decodeBase64/", DecodeBase64Handler)
	http.HandleFunc("/passgen/", PassGenHandler)
	http.HandleFunc("/encodeHex/", EncodeHexHandler)
	http.HandleFunc("/decodeHex/", DecodeHexHandler)
	http.HandleFunc("/ip/", IpHandler)
	http.HandleFunc("/checkport/", OpenportHandler)

	log.Print("Escuchando en el puerto "+PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
