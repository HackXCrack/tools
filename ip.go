package main

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
)

//show ip

//Ip muestra la ip en formato html
func Ip(w http.ResponseWriter, ip_s string) {
	RenderTemplate(w, "ip", "IP", ip_s)
}

//RawIp muestra la ip en formato raw
func RawIp(w http.ResponseWriter, ip_s string) {
	fmt.Fprint(w, ip_s)
}

//IpHandler decide el formato
func IpHandler(w http.ResponseWriter, r *http.Request) {
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)
	ip_s := strings.Split(r.RemoteAddr, ":")[0]

	if matched { //Show raw format
		RawIp(w, ip_s)
	} else {
		Ip(w, ip_s)
	}
}

//end show ip

//check port

//Openport muestra si el puerto esta abierto o cerrado en formato html
func Openport(w http.ResponseWriter, request string) {
	if strings.Split(request, ":")[1] == "" {
		RenderTemplate(w, "tool", "Check open port", "")
		return
	}
	_, err := net.Dial("tcp", request)
	if err != nil {
		RenderTemplate(w, "tool", "Check open port", "Cerrado")
	} else {
		RenderTemplate(w, "tool", "Check open port", "Abierto")
	}
}

//RawOpenport muestra si el puerto esta abierto o cerrado en formato raw
func RawOpenport(w http.ResponseWriter, request string) {
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

//OpenportHandler decide el formato
func OpenportHandler(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	body := ip + ":" + r.FormValue("p")
	matched, _ := regexp.MatchString("^*/raw", r.URL.Path)

	if matched { //Show raw format
		RawOpenport(w, body)
	} else {
		Openport(w, body)
	}
}

//end check port
