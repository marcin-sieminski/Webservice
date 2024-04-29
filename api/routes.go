package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.NotFound)
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/items", app.getCreateItemsHandler)
	mux.HandleFunc("/v1/items/", app.getUpdateDeleteItemsHandler)
	return mux
}
