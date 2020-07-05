package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	//"runtime"
)

func main() {

	// any other request will be handled by this function
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running GoMatha Request Handler")


		// read the body
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		
		if strings.Compare(string(b), "How I can save the World from Corona") == 0 {
			log.Println("Prayer Received")
			fmt.Fprintf(rw, "Just sit at home , My blessings are with you !!!!!!!!!!")
			log.Println("Blessings Provided")
		 }

	})

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting GoMatha Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}