package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	myhttp "webserver/http"
)

func main() {
	port := flag.Int("port", 8080, "port of the server")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", &myhttp.TextHandler{Response: "root"})
	mux.Handle("/home", &myhttp.TextHandler{Response: "home"})
	mux.Handle("/welcome", &myhttp.TextHandler{Response: "welcome"})

	log.Printf("listening on port: %d", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), mux)
}
