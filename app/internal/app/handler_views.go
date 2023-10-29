package app

import (
	"net/http"
)

// @Tags        Response from service
// @Summary		Get OK response
// @Description	Get an OK response [status: 200] for testing purposes
// @Produce		text/plain
// @Success		200
// @Router		/2xx [get]
func handlerOkReq(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello JetBrains!"))
}

// @Tags        Response from service
// @Summary		Get OK response
// @Description	Make redirect [status: 301] to "/2xx" for testing purposes
// @Produce		text/plain
// @Success		200
// @Router	    /3xx [get]
func handlerRedirectReq(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/2xx", http.StatusMovedPermanently)
}

// @Tags        Response from service
// @Summary		Get BadRequest response
// @Description	Get an error response [status: 400] for testing purposes
// @Produce		text/plain
// @Success		400
// @Router		/4xx [get]
func handlerBadReq(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "This is a bad request.", http.StatusBadRequest)
}

// @Tags        Response from service
// @Summary		Get InternalServerError response
// @Description	Get an error response [status: 500] for testing purposes
// @Produce		text/plain
// @Success		500
// @Router		/5xx [get]
func handlerServerErr(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "This is an internal error.", http.StatusInternalServerError)
}

func handlerInfo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/sys/swagger/index.html", http.StatusMovedPermanently)
}
