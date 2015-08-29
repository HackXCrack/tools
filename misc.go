package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"regexp"
)

//PassGen

//DoPass genera y devuelve la password pseudo-aleatoria
func DoPass(seed string) string {
	randBytes := make([]byte, len(seed))
	rand.Read(randBytes)
	randBytes2 := make([]byte, len(seed))
	rand.Read(randBytes2)
	reducedASCII := make([]byte, 94)
	for i := 0; i < 94; i++ {
		var tmp byte
		tmp = byte('!' + i)
		reducedASCII[i] = tmp
	}

	pass := make([]byte, len(seed))
	for i := 0; i < len(seed); i++ {
		tmp := randBytes[i] % 94
		if tmp == 0 {
			tmp = 94
		}
		pass[i] = reducedASCII[(seed[i]+randBytes2[i])%tmp]
	}
	return string(pass)
}

//PassGen genera la password en formato normal
func PassGen(w http.ResponseWriter, request string) {
	if request == "" {
		RenderTemplate(w, "tool", "Pass Generator", "")
		return
	}
	RenderTemplate(w, "tool", "Pass Generator", DoPass(request))
}

//RawPassGen genera la password en formato raw
func RawPassGen(w http.ResponseWriter, request string) {
	if request == "" {
		return
	}
	fmt.Fprint(w, DoPass(request))
}

//PassGenHandler decide el formato en que se mostrara la password
func PassGenHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawPassGen(w, body)
	} else {
		PassGen(w, body)
	}
}

//end PassGen
