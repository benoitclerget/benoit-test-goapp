package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Default http port
const (
	DefaultPort = "8080"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift!"
	}
	fmt.Fprintln(w, response)
	fmt.Println("Servicing request /health")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Working dir: " + path) // for example /home/user
	public := os.Getenv("OPENSHIFT_DATA_DIR")
	fmt.Println("OPENSHIFT_DATA_DIR: " + public) // openshift data path
	if public != "" {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(public))))
	} else {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path+"/html"))))
	}
	http.HandleFunc("/health", helloHandler)
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
