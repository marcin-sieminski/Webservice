package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/marcin-sieminski/webservice/models"
)

type application struct {
	itemslist *models.ItemslistModel
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/items", "Endpoint for the items list web service")

	app := &application{
		itemslist: &models.ItemslistModel{Endpoint: *endpoint},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
