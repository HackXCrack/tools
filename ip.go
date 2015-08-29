package main

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
)

//show ip

func Ip(w http.ResponseWriter, ip_s string) {
	renderTemplate(w, "ip", "ip", ip_s)
}

func Raw_ip(w http.ResponseWriter, ip_s string) {
	fmt.Fprint(w, ip_s)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)
	ip_s := strings.Split(r.RemoteAddr, ":")[0]

	if matched { //Show raw format
		Raw_ip(w, ip_s)
	} else {
		Ip(w, ip_s)
	}
}

//end show ip

//check port

func openport(w http.ResponseWriter, request string) {
	if strings.Split(request, ":")[1] == "" {
		renderTemplate(w, "tool", "Check open port", "")
		return
	}
	_, err := net.Dial("tcp", request)
	if err != nil {
		renderTemplate(w, "tool", "Check open port", "Cerrado")
	} else {
		renderTemplate(w, "tool", "Check open port", "Abierto")
	}
}

func raw_openport(w http.ResponseWriter, request string) {
	if strings.Split(request, ":")[1] == "" {
		return
	}
	_, err := net.Dial("tcp", request)
	if err != nil {
		fmt.Fprint(w, "Cerrado")
	} else {
		fmt.Fprint(w, "Abierto")
	}
}

func openportHandler(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	body := ip + ":" + r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		raw_openport(w, body)
	} else {
		openport(w, body)
	}
}

//end check port
