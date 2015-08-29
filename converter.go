package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"regexp"
)

//hex decode

func DecodeHex(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "Hex decoder", "")
		return
	}

	hash, err := hex.DecodeString(request)
	if err != nil {
		//fmt.Println("error:", err)
		return
	}
	RenderTemplate(w, "tool", "Hex decoder", string(hash))
}

func RawDecodeHex(w http.ResponseWriter, request string) {
	hash, err := hex.DecodeString(request)
	if err != nil {
		//fmt.Println("error:", err)
		return
	}
	fmt.Fprint(w, string(hash))
}

func DecodeHexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")

	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawDecodeHex(w, body)
	} else {
		DecodeHex(w, body)
	}
}

//end hex decode

//hex encode

func EncodeHex(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "Hex encoder", "")
		return
	}
	hash := hex.EncodeToString([]byte(request))
	RenderTemplate(w, "tool", "Hex encoder", hash)
}

func RawEncodeHex(w http.ResponseWriter, request string) {
	hash := hex.EncodeToString([]byte(request))
	fmt.Fprint(w, hash)
}

func EncodeHexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawEncodeHex(w, body)
	} else {
		EncodeHex(w, body)
	}
}

//end hex encode

//base64  decode

func DecodeBase64(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "Base64 decoder", "")
		return
	}

	hash, err := base64.StdEncoding.DecodeString(request)
	if err != nil {
		//se debería mostrar el error al user???
		return
	}
	RenderTemplate(w, "tool", "Base64 decoder", string(hash))
}

func RawDecodeBase64(w http.ResponseWriter, request string) {
	hash, err := base64.StdEncoding.DecodeString(request)
	if err != nil {
		//se debería mostrar el error al user???
		return
	}
	fmt.Fprint(w, string(hash))
}

func DecodeBase64Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawDecodeBase64(w, body)
	} else {
		DecodeBase64(w, body)
	}
}

//end base64 decode

//base64 encode

func EncodeBase64(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "Base64 encoder", "")
		return
	}

	hash := base64.StdEncoding.EncodeToString([]byte(request))
	RenderTemplate(w, "tool", "Base64 encoder", hash)

}

func RawEncodeBase64(w http.ResponseWriter, request string) {
	hash := base64.StdEncoding.EncodeToString([]byte(request))
	fmt.Fprint(w, hash)
}

func EncodeBase64Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawEncodeBase64(w, body)
	} else {
		EncodeBase64(w, body)
	}
}

//end base64 encode
