package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

//Puerto donde escuchar
const PORT = "1337"

//Page contiene datos de la página
type Page struct {
	Title   string
	Request string
}

const templatesPath = "./templates/"

var templates = template.Must(template.ParseGlob(templatesPath + "[a-z]*"))

//RenderTemplate renderiza la template con el texto adecuado
func RenderTemplate(w http.ResponseWriter, tmpl string, title string, r interface{}) {
	var err error
	switch request := r.(type) {
	case string:
		err = templates.ExecuteTemplate(w, tmpl, &Page{title, request})
	case error:
		//TODO: Si es un error se imprima en otro formato
		err = templates.ExecuteTemplate(w, tmpl, &Page{title, request.Error()})
	default:
		log.Print("Strange error ocurred")
	}

	if err != nil {
		log.Print("ExecuteTemplate: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RootHandler muestra el index
func RootHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index", "index", "")
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
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

	log.Print("Escuchando en el puerto " + PORT)

	s := &http.Server{
		Addr:         ":" + PORT,
		ReadTimeout:  7 * time.Second,
		WriteTimeout: 7 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}