package main

import (
  "encoding/base64"
  "encoding/hex"
  "fmt"
  "net/http"
)

//hex decode

func raw_decodeHexHandler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash, err := hex.DecodeString(body)
  if err != nil {
    fmt.Println("error:", err)
    return
  }
  fmt.Fprint(w, string(hash))
}

func decodeHexHandler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "Hex decoder")
    return
  }
  hash, err := hex.DecodeString(body)
  if err != nil {
    fmt.Println("error:", err)
    return
  }
  templates.Execute(w, &Page{Title: "Hex decoder", R: string(hash) })
}

//end hex decode

//hex encode

func raw_encodeHexHandler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash := hex.EncodeToString([]byte(body))
  fmt.Fprint(w, hash)
}

func encodeHexHandler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "Hex encoder")
    return
  }
  hash := hex.EncodeToString([]byte(body))
  templates.Execute(w, &Page{Title: "Hex encoder", R: hash})
}

//end hex encode


//base64  decode

func raw_decodeBase64Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash, err := base64.StdEncoding.DecodeString(string(body))
  if err != nil {
    fmt.Println("error:", err)
    return
  }
  fmt.Fprint(w, string(hash))
}

func decodeBase64Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "Base64 decoder")
    return
  }
  hash, err := base64.StdEncoding.DecodeString(body)
  if err != nil {
    fmt.Println("error:", err)
    return
  }
  templates.Execute(w, &Page{Title: "Base64 decoder", R: string(hash) })
}

//end base64 decode

//base64 encode

func raw_encodeBase64Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  hash := base64.StdEncoding.EncodeToString([]byte(body))
  fmt.Fprint(w, hash)
}

func encodeBase64Handler(w http.ResponseWriter, r *http.Request) {
  body := r.FormValue("body")

  if body == "" {
    renderTemplate(w, "html", "Base64 encoder")
    return
  }
  hash := base64.StdEncoding.EncodeToString([]byte(body))
  templates.Execute(w, &Page{Title: "Base64 encoder", R: hash})
}

//end base64 encode
