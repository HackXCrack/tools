package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"regexp"
)

//pass gen

func doPass(seed string) string {
	rand_bytes := make([]byte, len(seed))
	rand.Read(rand_bytes)
	rand_bytes2 := make([]byte, len(seed))
	rand.Read(rand_bytes2)
	reduced_ascii := make([]byte, 94)
	for i := 0; i < 94; i++ {
		var tmp byte
		tmp = byte('!' + i)
		reduced_ascii[i] = tmp
	}

	pass := make([]byte, len(seed))
	for i := 0; i < len(seed); i++ {
		tmp := rand_bytes[i] % 94
		if tmp == 0 {
			tmp = 94
		}
		pass[i] = reduced_ascii[(seed[i]+rand_bytes2[i])%tmp]
	}
	return string(pass)
}

func passGen(w http.ResponseWriter, request string) {
	if request == "" {
		renderTemplate(w, "tool", "Pass Generator", "")
		return
	}
	renderTemplate(w, "tool", "Pass Generator", doPass(request))
}

func raw_passGen(w http.ResponseWriter, request string) {
	if request == "" {
		return
	}
	fmt.Fprint(w, doPass(request))
}

func passgenHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_passGen(w, body)
	} else {
		passGen(w, body)
	}
}
