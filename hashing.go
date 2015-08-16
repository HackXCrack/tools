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

func sha3512Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash := sha3.Sum512([]byte(body))
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched {
		fmt.Fprintf(w, "%x", hash)
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "SHA3 - 512", "")
		return
	}
	s := hex.EncodeToString(hash[:])
	renderTemplate(w, "tool", "SHA3 - 512", s)
}

//end sha3-512

//sha3-256

func sha3Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash := sha3.Sum256([]byte(body))
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched {
		fmt.Fprintf(w, "%x", hash)
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "SHA3 - 256", "")
		return
	}
	s := hex.EncodeToString(hash[:])
	renderTemplate(w, "tool", "SHA3 - 256", s)
}

//end sha3-256

//md5

func md5Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	hash := md5.Sum([]byte(body))
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched {
		fmt.Fprintf(w, "%x", hash)
		return
	}

	//Normal format
	if body == "" {
		renderTemplate(w, "tool", "MD5", "")
		return
	}

	s := hex.EncodeToString(hash[:])
	renderTemplate(w, "tool", "MD5", s)
}

///end md5
