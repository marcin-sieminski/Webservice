package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	items, err := app.itemslist.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "<html><head><title>Items List</title></head><body><h1>Items List</h1><ul>")
	for _, item := range *items {
		fmt.Fprintf(w, "<li>%s</li>", item.Name)
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func (app *application) itemView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	item, err := app.itemslist.Get(int64(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", item.Name)
}

func (app *application) itemCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.itemCreateForm(w, r)
	case http.MethodPost:
		app.itemCreateProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) itemCreateForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Create Item</title></head>"+
		"<body><h1>Create Item</h1><form action=\"/item/create\" method=\"post\">"+
		"<label for=\name\">Name</label><input type=\"text\" name=\"name\" id=\"name\">"+
		"<button type=\"submit\">Create</button></form></body></html>")
}

func (app *application) itemCreateProcess(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	if name == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	data, err := json.Marshal(item)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req, _ := http.NewRequest("POST", app.itemslist.Endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("unexpected status: %s", resp.Status)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
