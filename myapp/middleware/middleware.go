package middleware

import (
	"myapp/data"

	"github.com/brucebotes/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
