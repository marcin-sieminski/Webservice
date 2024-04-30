package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/marcin-sieminski/webservice/data"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *application) getCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		items := []data.Item{
			{
				ID:        1,
				CreatedAt: time.Now(),
				Name:      "Name1",
				Version:   1,
			},
			{
				ID:        2,
				CreatedAt: time.Now(),
				Name:      "Name2",
				Version:   1,
			},
		}

		if err := app.writeJSON(w, http.StatusOK, envelope{"items": items}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodPost {
		var input struct {
			Name string `json:"name"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%+v\n", input)
	}
}

func (app *application) getUpdateDeleteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.get(w, r)
	case http.MethodPut:
		app.update(w, r)
	case http.MethodDelete:
		app.delete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/items/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	item := data.Item{
		ID:        idInt,
		CreatedAt: time.Now(),
		Name:      "Name1",
		Version:   1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"item": item}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/items/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var input struct {
		Name *string `json:"name"`
	}

	item := data.Item{
		ID:        idInt,
		CreatedAt: time.Now(),
		Name:      "Name1",
		Version:   1,
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if input.Name != nil {
		item.Name = *input.Name
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"item": item}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/items/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delete a specific item with ID: %d", idInt)
}
