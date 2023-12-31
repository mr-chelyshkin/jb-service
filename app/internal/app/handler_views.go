package app

import (
	"encoding/json"
	"net/http"

	"github.com/mr-chelyshkin/jb-service/app"
)

type viewResponse struct {
	Host    string `json:"host"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
}

// @Tags        Response from service
// @Summary		Get OK response.
// @Description	Get an OK response [status: 200] for testing purposes.
// @Produce		json
// @Success		200
// @Router		/2xx [get]
func handlerOkReq(w http.ResponseWriter, _ *http.Request) {
	response := viewResponse{
		Status:  http.StatusOK,
		Message: "StatusOK",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Response from service
// @Summary		Get OK response.
// @Description	Make redirect [status: 301] to "/2xx" for testing purposes.
// @Produce		json
// @Success		200
// @Router	    /3xx [get]
func handlerRedirectReq(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/2xx", http.StatusMovedPermanently)
}

// @Tags        Response from service
// @Summary		Get BadRequest response.
// @Description	Get an error response [status: 400] for testing purposes.
// @Produce		json
// @Success		400
// @Router		/4xx [get]
func handlerBadReq(w http.ResponseWriter, _ *http.Request) {
	response := viewResponse{
		Status:  http.StatusBadRequest,
		Message: "StatusBadRequest",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Response from service
// @Summary		Get InternalServerError response
// @Description	Get an error response [status: 500] for testing purposes.
// @Produce		json
// @Success		500
// @Router		/5xx [get]
func handlerServerErr(w http.ResponseWriter, _ *http.Request) {
	response := viewResponse{
		Status:  http.StatusInternalServerError,
		Message: "StatusInternalServerError",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(response)
}

func handlerInfo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/sys/swagger/index.html", http.StatusMovedPermanently)
}
