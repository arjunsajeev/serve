package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port")
	flag.Parse()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir('.'))
	mux.Handle("/", fileServer)

	err := http.ListenAndServe(":"+strconv.Itoa(*port), mux)
	log.Fatal(err)
}
