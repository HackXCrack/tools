package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"net/http"
	"regexp"
)

//sha3-512

func sha3512(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "SHA3 - 512", "")
		return
	}
	hash := sha3.Sum512([]byte(request))
	s := hex.EncodeToString(hash[:])
	renderTemplate(w, "tool", "SHA3 - 512", s)
}

func raw_sha3512(w http.ResponseWriter, request string) {
	hash := sha3.Sum512([]byte(request))
	fmt.Fprintf(w, "%x", hash)
}

func sha3512Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_sha3512(w, body)
	} else {
		sha3512(w, body)
	}
}

//end sha3-512

//sha3-256

func sha3_f(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "SHA3 - 256", "")
		return
	}
	hash := sha3.Sum256([]byte(request))
	s := hex.EncodeToString(hash[:])
	renderTemplate(w, "tool", "SHA3 - 256", s)
}

func raw_sha3(w http.ResponseWriter, request string) {
	hash := sha3.Sum256([]byte(request))
	fmt.Fprintf(w, "%x", hash)
}

func sha3Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_sha3(w, body)
	} else {
		sha3_f(w, body)
	}
}

//end sha3-256

//md5

func md5_f(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "MD5", "")
		return
	}
	hash := md5.Sum([]byte(request))
	s := hex.EncodeToString(hash[:])
	renderTemplate(w, "tool", "MD5", s)
}

func raw_md5(w http.ResponseWriter, request string) {
	hash := md5.Sum([]byte(request))
	fmt.Fprintf(w, "%x", hash)
}

func md5Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_md5(w, body)
	} else {
		md5_f(w, body)
	}
}

///end md5
