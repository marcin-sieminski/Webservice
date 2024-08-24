package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/marcin-sieminski/Webservice/model"
)

type application struct {
	itemslist *model.ItemslistModel
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/item", "Endpoint for the items list web service")

	app := &application{
		itemslist: &model.ItemslistModel{Endpoint: *endpoint},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
