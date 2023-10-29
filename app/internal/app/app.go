package app

import (
	"context"

	"github.com/mr-chelyshkin/jb-service/app/pkg/log"
	"github.com/mr-chelyshkin/jb-service/app/pkg/ws"

	"github.com/go-chi/chi/v5"
)

type webServer interface {
	Run(ctx context.Context, l ws.Logger)
	Router() *chi.Mux
}

// App object.
type App struct {
	logger ws.Logger
	server webServer
}

// New create Service object.
func New() (App, error) {
	l, err := log.New([]byte("INFO"))
	if err != nil {
		return App{}, err
	}
	s, err := ws.New(
		ws.WithMiddlewareReqID(),
		ws.WithMiddlewareRealIP(),
		ws.WithMiddlewareTimeout(3),

		ws.WithShutdownTimeout(5),
		ws.WithMiddlewareCustomAfter(customLogger(l)),
	)

	return App{
		server: s,
		logger: l,
	}, nil
}

// Run service.
func (a App) Run(ctx context.Context) {
	a.routerTools()
	a.routerViews()

	a.server.Run(ctx, a.logger)
}
