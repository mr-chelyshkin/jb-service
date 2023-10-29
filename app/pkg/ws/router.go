package ws

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func sysRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	return r
}
