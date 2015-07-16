package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	R string
	Body  []byte
}

var templates_path = "./templates/"
var templates = template.Must(template.ParseFiles(templates_path + "tool.html", templates_path + "index.html", templates_path + "ip.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, title string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", &Page{Title: title})
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

func main() {
	Hello := func(w http.ResponseWriter, r *http.Request) { templates.ExecuteTemplate(w, "index.html", nil) }
	http.HandleFunc("/", Hello)
	http.HandleFunc("/md5", makeHandler(md5Handler))
	http.HandleFunc("/raw_md5", makeHandler(raw_md5Handler))
	http.HandleFunc("/sha3256", makeHandler(sha3Handler))
	http.HandleFunc("/raw_sha3256", makeHandler(raw_sha3Handler))
	http.HandleFunc("/sha3512", makeHandler(sha3512Handler))
	http.HandleFunc("/raw_sha3512", makeHandler(raw_sha3512Handler))
	http.HandleFunc("/encodeBase64", makeHandler(encodeBase64Handler))
	http.HandleFunc("/raw_encodeBase64", makeHandler(raw_encodeBase64Handler))
	http.HandleFunc("/decodeBase64", makeHandler(decodeBase64Handler))
	http.HandleFunc("/raw_decodeBase64", makeHandler(raw_decodeBase64Handler))
	http.HandleFunc("/passgen", makeHandler(passgenHandler))
	http.HandleFunc("/raw_passgen", makeHandler(raw_passgenHandler))
	http.HandleFunc("/encodeHex", makeHandler(encodeHexHandler))
	http.HandleFunc("/raw_encodeHex", makeHandler(raw_encodeHexHandler))
	http.HandleFunc("/decodeHex", makeHandler(decodeHexHandler))
	http.HandleFunc("/raw_decodeHex", makeHandler(raw_decodeHexHandler))
	http.HandleFunc("/ip", makeHandler(ipHandler))
	http.HandleFunc("/raw_ip", makeHandler(raw_ipHandler))

	http.ListenAndServe(":1337", nil)
}
