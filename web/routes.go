package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("../view/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/item/view/", app.itemView)
	mux.HandleFunc("/item/delete/", app.itemDelete)
	mux.HandleFunc("/item/create", app.itemCreate)

	return mux
}
