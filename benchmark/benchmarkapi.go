package benchmark

import (
	"flag"

	"github.com/marcin-sieminski/webservice/model"
)

var (
	endpoint = flag.String("endpoint", "http://localhost:4000/v1/items", "Endpoint for the items list web service")
)

type application struct {
	itemslist *model.ItemslistModel
}

func benchmarkApi() bool {
	app := &application{
		itemslist: &model.ItemslistModel{Endpoint: *endpoint},
	}
	_, err := app.itemslist.GetAll()

	return err != nil
}
