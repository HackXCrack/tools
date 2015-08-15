package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//show ip

func ipHandler(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	//Show raw format
	if matched {
		fmt.Fprint(w, strings.Split(r.RemoteAddr, ":")[0])
		return
	}

	renderTemplate(w, "ip", "ip", strings.Split(r.RemoteAddr, ":")[0])
}
