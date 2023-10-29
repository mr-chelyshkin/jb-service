package app

import (
	"github.com/go-chi/chi/v5"
)

func (a App) routerTools() {
	router := func() chi.Router {
		r := chi.NewRouter()
		r.Get("/liveness-change", handlerChangeLivenessProbe)
		r.Get("/readness-change", handlerChangeReadnessProbe)

		r.Get("/liveness", handlerLivenessProbe)
		r.Get("/readness", handlerReadnessProbe)

		r.Get("/oomkill", handlerCallOOM)
		r.Get("/throttling", handlerThrottlingCPU)
		return r
	}
	a.server.Router().Mount("/api/v1", router())
}

func (a App) routerViews() {
	router := func() chi.Router {
		r := chi.NewRouter()
		r.Get("/3xx", handlerRedirectReq)
		r.Get("/5xx", handlerServerErr)
		r.Get("/4xx", handlerBadReq)
		r.Get("/2xx", handlerOkReq)
		r.Get("/", handlerInfo)
		return r
	}
	a.server.Router().Mount("/", router())
}
