package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"regexp"
)

//pass gen

func passgenHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched && body == "" {
		return
	}

	if body == "" {
		renderTemplate(w, "tool", "Pass Generator", "")
		return
	}

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

	if matched {
		fmt.Fprint(w, string(pass))
		return
	}

	renderTemplate(w, "tool", "Pass Generator", string(pass))
}
