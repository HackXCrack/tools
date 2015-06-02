package main

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "golang.org/x/crypto/sha3"
  "net/http"
)

//sha3-512

func raw_sha3512Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash := sha3.Sum512([]byte(body))
  fmt.Fprintf(w, "%x", hash)
}

func sha3512Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "SHA3 - 512")
    return
  }
  hash := sha3.Sum512([]byte(body))
  s := hex.EncodeToString(hash[:])
  templates.Execute(w, &Page{Title: "SHA3 - 512", R: s})
}

//end sha3-512

//sha3-256

func raw_sha3Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash := sha3.Sum256([]byte(body))
  fmt.Fprintf(w, "%x", hash)
}

func sha3Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "SHA3 - 256")
    return
  } 
  hash := sha3.Sum256([]byte(body))
  s := hex.EncodeToString(hash[:])
  templates.Execute(w, &Page{Title: "SHA3 - 256", R: s})
}

//end sha3-256

//md5

func md5Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "MD5")
    return
  }
  hash := md5.Sum([]byte(body))
  s := hex.EncodeToString(hash[:])
  templates.Execute(w, &Page{Title: "MD5", R: s}) 
}

func raw_md5Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash := md5.Sum([]byte(body))
  fmt.Fprintf(w, "%x", hash)
}

///end md5
