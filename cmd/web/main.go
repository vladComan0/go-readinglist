package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")
	app := &application{}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting the web application on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
