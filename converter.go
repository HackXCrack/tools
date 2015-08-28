package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"regexp"
)

//hex decode

func decodeHex(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "Hex decoder", "")
		return
	}

	hash, err := hex.DecodeString(request)
	if err != nil {
		//fmt.Println("error:", err)
		return
	}
	renderTemplate(w, "tool", "Hex decoder", string(hash))
}

func raw_decodeHex(w http.ResponseWriter, request string) {
	hash, err := hex.DecodeString(request)
	if err != nil {
		//fmt.Println("error:", err)
		return
	}
	fmt.Fprint(w, string(hash))
}

func decodeHexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")

	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_decodeHex(w, body)
	} else {
		decodeHex(w, body)
	}
}

//end hex decode

//hex encode

func encodeHex(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "Hex encoder", "")
		return
	}
	hash := hex.EncodeToString([]byte(request))
	renderTemplate(w, "tool", "Hex encoder", hash)
}

func raw_encodeHex(w http.ResponseWriter, request string) {
	hash := hex.EncodeToString([]byte(request))
	fmt.Fprint(w, hash)
}

func encodeHexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_encodeHex(w, body)
	} else {
		encodeHex(w, body)
	}
}

//end hex encode

//base64  decode

func decodeBase64(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "Base64 decoder", "")
		return
	}

	hash, err := base64.StdEncoding.DecodeString(request)
	if err != nil {
		//se debería mostrar el error al user???
		return
	}
	renderTemplate(w, "tool", "Base64 decoder", string(hash))
}

func raw_decodeBase64(w http.ResponseWriter, request string) {
	hash, err := base64.StdEncoding.DecodeString(request)
	if err != nil {
		//se debería mostrar el error al user???
		return
	}
	fmt.Fprint(w, string(hash))
}

func decodeBase64Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_decodeBase64(w, body)
	} else {
		decodeBase64(w, body)
	}
}

//end base64 decode

//base64 encode

func encodeBase64(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "Base64 encoder", "")
		return
	}

	hash := base64.StdEncoding.EncodeToString([]byte(request))
	renderTemplate(w, "tool", "Base64 encoder", hash)

}

func raw_encodeBase64(w http.ResponseWriter, request string) {
	hash := base64.StdEncoding.EncodeToString([]byte(request))
	fmt.Fprint(w, hash)
}

func encodeBase64Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_encodeBase64(w, body)
	} else {
		encodeBase64(w, body)
	}
}

//end base64 encode
