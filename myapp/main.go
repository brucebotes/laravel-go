package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/brucebotes/celeritas"
	"github.com/docker/docker/api/server/middleware"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
