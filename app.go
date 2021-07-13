package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Default http port
const (
	DefaultPort = "8080"
)

func systemInfosHandler(w http.ResponseWriter, r *http.Request) {
	response := ""
	public := os.Getenv("APP_ROOT")
	response += "APP_ROOT: " + public
	home := os.Getenv("HOME")
	response += "<br />HOME: " + home
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		response += "<br />Working dir: " + path
	}
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Go version: %s\n", strings.TrimSuffix(string(out), "\n"))
	response += "<br />Go version: " + strings.TrimSuffix(string(out), "\n")
	out, err = exec.Command("uname", "-m").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Computer arch: %s\n", strings.TrimSuffix(string(out), "\n"))
	response += "<br />Computer architecture: " + strings.TrimSuffix(string(out), "\n")

	name := ""
	version := ""
	releaseContent, err := exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`NAME=(.*)`)
	//fmt.Printf("%q\n", re.FindStringSubmatch(string(releaseContent)))
	matches := re.FindStringSubmatch(string(releaseContent))
	if len(matches) > 0 {
		//fmt.Printf("Os Name : %s", matches[1])
		name = matches[1]
	}
	re = regexp.MustCompile(`VERSION=(.*)`)
	//fmt.Printf("%q\n", re.FindStringSubmatch(string(releaseContent)))
	matches = re.FindStringSubmatch(string(releaseContent))
	if len(matches) > 0 {
		//fmt.Printf("Os Version : %s\n", matches[1])
		version = matches[1]
	}
	fmt.Printf("Os Name / version: %s / %s\n", name, version)
	response += "<br />OS Name: " + strings.TrimSuffix(string(name), "\n")
	response += "<br />OS Version: " + strings.TrimSuffix(string(version), "\n")
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, response)
	fmt.Println("Servicing request /api/system_infos")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/api/system_infos", systemInfosHandler)

	htmlPath := os.Getenv("HTML_PATH")
	if len(htmlPath) == 0 {
		htmlPath = "/tmp/src/public"
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(htmlPath))))
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = DefaultPort
	}
	log.Printf("Starting app on port %+v\n", port)

	go listenAndServe(port)

	/*
		port = os.Getenv("SECOND_PORT")
		if len(port) == 0 {
			port = "8888"
		}
		go listenAndServe(port)
	*/

	select {}

}
