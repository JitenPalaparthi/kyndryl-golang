package main

import (
	"log"
	"net/http"
	"os"
)

var (
	PORT string
)

func main() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			//fmt.Fprintln(w, "Hello Kyndryl!")
			w.Write([]byte("Hello Kyndryl!"))
			w.WriteHeader(http.StatusOK) // default is 200
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No method implemented on this action"))
		}
	})
	http.HandleFunc("/ping", ping)

	log.Println("Application is up and running on Port:", PORT)
	http.ListenAndServe(":"+PORT, nil) // if you dont give any thing it is 0.0.0.0

}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Pong"))
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// localhost:8085
