package main

import (
	"log"
	"net/http"
)

const port = ":3000"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Serving on", port)
	http.ListenAndServe(port, nil)
}
