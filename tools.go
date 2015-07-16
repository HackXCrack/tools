package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Page struct {
	Title string
	R string
	Body  []byte
}

//show ip

func ipHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "ip.html", &Page{R: strings.Split(r.RemoteAddr,":")[0] })
}

func raw_ipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strings.Split(r.RemoteAddr,":")[0] )
}

//end show ip

//pass gen

func passgenHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")

	if body == "" {
		renderTemplate(w, "tool", "PASS GENERATOR")
		return
	}
	rand_bytes := make([]byte, len(body))
	rand_bytes2 := make([]byte, len(body))
	rand.Read(rand_bytes)
	rand.Read(rand_bytes2)

	reduced_ascii := make([]byte, 94)
	for i := 0; i < 94; i++ {
		var tmp byte
		tmp = byte('!' + i)
		reduced_ascii[i] = tmp
	}
		pass := make([]byte, len(body))
	for i := 0; i < len(body); i++ {
		tmp := rand_bytes[i] % 94
		if tmp == 0 {
			tmp = 94
		}
		pass[i] = reduced_ascii[(body[i]+rand_bytes2[i])%tmp]
	}
	templates.Execute(w, &Page{Title: "PASS GENERATOR", R: string(pass) })
}

func raw_passgenHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "raw", "raw")
	body := r.FormValue("body")

	if body == "" {return}

	rand_bytes := make([]byte, len(body))
	rand.Read(rand_bytes)
	rand_bytes2 := make([]byte, len(body))
	rand.Read(rand_bytes2)

	reduced_ascii := make([]byte, 94)
	for i := 0; i < 94; i++ {
		var tmp byte
		tmp = byte('!' + i)
		reduced_ascii[i] = tmp
	}

	pass := make([]byte, len(body))
	for i := 0; i < len(body); i++ {
		tmp := rand_bytes[i] % 94
		if tmp == 0 {
			tmp = 94
		}
		pass[i] = reduced_ascii[(body[i]+rand_bytes2[i])%tmp]

	}
	fmt.Fprint(w, string(pass))
}

//end passgen
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
