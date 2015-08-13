package main

import (
	"fmt"
	"net/http"
	"strings"
)

//show ip

func ipHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "ip", &Page{R: strings.Split(r.RemoteAddr,":")    [0] })
}

func raw_ipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strings.Split(r.RemoteAddr,":")[0] )
}
