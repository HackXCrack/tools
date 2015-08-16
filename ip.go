package main

import (
	"fmt"
	"net"
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

//end show ip

//check port

func openportHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("p")
	if body == "" {
		renderTemplate(w, "tool", "Check open port", "")
		return
	}

	ip := strings.Split(r.RemoteAddr, ":")[0]
	_, err := net.Dial("tcp", ip+":"+body)
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)
	//Show raw format
	if matched {
    if err != nil {
      fmt.Fprint(w, "Close") 
    } else {
		  fmt.Fprint(w, "Open")
    }
		return
	}
	
  if err != nil {
		renderTemplate(w, "tool", "Check open port", "Close")
		return
	}

	renderTemplate(w, "tool", "Check open port", "Open")
}
