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

func Sha3512(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "SHA3 - 512", "")
		return
	}
	hash := sha3.Sum512([]byte(request))
	s := hex.EncodeToString(hash[:])
	RenderTemplate(w, "tool", "SHA3 - 512", s)
}

func RawSha3512(w http.ResponseWriter, request string) {
	hash := sha3.Sum512([]byte(request))
	fmt.Fprintf(w, "%x", hash)
}

func Sha3512Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawSha3512(w, body)
	} else {
		Sha3512(w, body)
	}
}

//end sha3-512

//sha3-256

func Sha3(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "SHA3 - 256", "")
		return
	}
	hash := sha3.Sum256([]byte(request))
	s := hex.EncodeToString(hash[:])
	RenderTemplate(w, "tool", "SHA3 - 256", s)
}

func RawSha3(w http.ResponseWriter, request string) {
	hash := sha3.Sum256([]byte(request))
	fmt.Fprintf(w, "%x", hash)
}

func Sha3Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawSha3(w, body)
	} else {
		Sha3(w, body)
	}
}

//end sha3-256

//md5

func Md5(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "MD5", "")
		return
	}
	hash := md5.Sum([]byte(request))
	s := hex.EncodeToString(hash[:])
	RenderTemplate(w, "tool", "MD5", s)
}

func RawMd5(w http.ResponseWriter, request string) {
	hash := md5.Sum([]byte(request))
	fmt.Fprintf(w, "%x", hash)
}

func Md5Handler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawMd5(w, body)
	} else {
		Md5(w, body)
	}
}

///end md5
