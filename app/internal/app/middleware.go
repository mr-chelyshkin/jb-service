package app

import (
	"net/http"
	"time"

	"github.com/mr-chelyshkin/jb-service/app"
	"github.com/mr-chelyshkin/jb-service/app/pkg/ws"

	"github.com/google/uuid"
)

func customLogger(l ws.Logger) func(w http.ResponseWriter, r *http.Request) error {
	f := func(w http.ResponseWriter, r *http.Request) error {
		t := time.Now()
		l.Printf(
			"served request",
			map[string]any{
				"method":     r.Method,
				"path":       r.URL.String(),
				"duration":   time.Since(t).String(),
				"replica_id": app.ReplicaID(),
				"request_id": uuid.New(),
			},
		)
		return nil
	}
	return f
}
