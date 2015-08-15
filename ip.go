package main

import (
	"fmt"
	"net/http"
	"strings"
)

//show ip

func ipHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "ip", "ip", strings.Split(r.RemoteAddr,":")[0])
}

func raw_ipHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, strings.Split(r.RemoteAddr,":")[0] )
}
