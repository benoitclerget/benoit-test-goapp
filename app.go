package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

// Default http port
const (
	DefaultPort = "3000"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DefaultPort
	}

	router := mux.NewRouter().StrictSlash(true)

	/*
		app.get('/org.ibm.gdps/images/:flavorsvg', function (req, res) {
			res.sendFile(path.join(__dirname, 'html/images/' + req.params.flavorsvg))
		  })
	*/
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("html"))))

	log.Printf("Starting app on port %+v\n", port)
	//http.ListenAndServe(":"+port, nil)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// https://golang.org/pkg/net/http/
//https://thenewstack.io/make-a-restful-json-api-go/
//http://www.gorillatoolkit.org/pkg/mux
