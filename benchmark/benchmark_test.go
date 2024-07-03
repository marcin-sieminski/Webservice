package benchmark

import (
	"flag"
	"testing"

	"github.com/marcin-sieminski/webservice/model"
)

var (
	endpoint = flag.String("endpoint", "http://localhost:4000/v1/items", "Endpoint for the items list web service")
)

type application struct {
	itemslist *model.ItemslistModel
}

func BenchmarkApi(b *testing.B) {
	b.StopTimer()
	app := &application{
		itemslist: &model.ItemslistModel{Endpoint: *endpoint},
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, err := app.itemslist.GetAll()
		b.StopTimer()
		if err != nil {
			b.Fatal("api error")
		}
	}
}
