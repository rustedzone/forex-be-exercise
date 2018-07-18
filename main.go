package main

import (
	"flag"
	"fmt"
	"forex-be-exercise/route"
	"net/http"
)

func main() {
	r := route.Router()
	port := flag.String("port", "80", "App Port")
	flag.Parse()

	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", *port),
	}
	server.ListenAndServe()
}
