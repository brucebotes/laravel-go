package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/brucebotes/celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}