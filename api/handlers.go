package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) getCreateItemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of the items")
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added a new item")
		return
	}
}

func (app *application) getUpdateDeleteItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getItem(w, r)
	case http.MethodPut:
		app.updateItem(w, r)
	case http.MethodDelete:
		app.deleteItem(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/items/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Display the details of a specific item with ID: %d", idInt)
}

func (app *application) updateItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/items/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Update the details of a specific item with ID: %d", idInt)
}

func (app *application) deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/items/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete a specific item with ID: %d", idInt)
}
