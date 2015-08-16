package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"regexp"
)

//hex decode

func decodeHexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash, err := hex.DecodeString(body)
	if err != nil {
		//fmt.Println("error:", err)
		return
	}

	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched {
		fmt.Fprint(w, string(hash))
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "Hex decoder", "")
		return
	}

	renderTemplate(w, "tool", "Hex decoder", string(hash))
}

//end hex decode

//hex encode

func encodeHexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash := hex.EncodeToString([]byte(body))
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched {
		fmt.Fprint(w, hash)
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "Hex encoder", "")
		return
	}
	renderTemplate(w, "tool", "Hex encoder", hash)
}

//end hex encode

//base64  decode

func decodeBase64Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		//se debería mostrar el error al user???
		return
	}

	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched {
		fmt.Fprint(w, string(hash))
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "Base64 decoder", "")
		return
	}

	renderTemplate(w, "tool", "Base64 decoder", string(hash))
}

//end base64 decode

//base64 encode

func encodeBase64Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash := base64.StdEncoding.EncodeToString([]byte(body))
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched {
		fmt.Fprint(w, hash)
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "Base64 encoder", "")
		return
	}
	renderTemplate(w, "tool", "Base64 encoder", hash)
}

//end base64 encode
