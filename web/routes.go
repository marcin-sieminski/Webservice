package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/item/view/", app.itemView)
	mux.HandleFunc("/item/create", app.itemCreate)

	return mux
}
