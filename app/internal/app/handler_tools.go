package app

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"github.com/mr-chelyshkin/jb-service/app"
)

type probeResponse struct {
	Host    string `json:"host"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

// @Tags        Service State
// @Summary		Return service LivenessProbe.
// @Description	This is LivenessProbe for K8S.
// @Produce		json
// @Success		200
// @Router		/api/v1/liveness [get]
func handlerLivenessProbe(w http.ResponseWriter, _ *http.Request) {
	if !cfg.livenessIsOk {
		http.Error(w, fmt.Sprintf("LivenessProbe [%s] internal error", app.ReplicaID()), http.StatusInternalServerError)
		return
	}

	response := probeResponse{
		Status:  "OK",
		Message: "LivenessProbe",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Service State
// @Summary		Change LivenessProbe: Success/Failure.
// @Description	Change LivenessProbe service flag for check K8S reaction (expect pod restart).
// @Produce		json
// @Success		200
// @Router		/api/v1/liveness-change [get]
func handlerChangeLivenessProbe(w http.ResponseWriter, _ *http.Request) {
	switch cfg.livenessIsOk {
	case true:
		cfg.livenessIsOk = false
	case false:
		cfg.livenessIsOk = true
	}

	response := probeResponse{
		Status:  "OK",
		Message: fmt.Sprintf("LivenessProbe flag changed, now it's: %t", cfg.livenessIsOk),
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Service State
// @Summary		Return service ReadnessProbe.
// @Description	This is ReadnessProbe for K8S.
// @Produce		json
// @Success		200
// @Router		/api/v1/readness [get]
func handlerReadnessProbe(w http.ResponseWriter, _ *http.Request) {
	if !cfg.readnessIsOk {
		http.Error(w, fmt.Sprintf("ReadnessProbe [%s] internal error", app.ReplicaID()), http.StatusInternalServerError)
		return
	}

	response := probeResponse{
		Status:  "OK",
		Message: "ReadnessProbe",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Service State
// @Summary		Change ReadnessProbe: Success/Failure.
// @Description	Change ReadnessProbe service flag for check K8S reaction (expect traffic interrupt, check by req logs: "replica_id").
// @Produce		json
// @Success		200
// @Router		/api/v1/readness-change [get]
func handlerChangeReadnessProbe(w http.ResponseWriter, _ *http.Request) {
	switch cfg.readnessIsOk {
	case true:
		cfg.readnessIsOk = false
	case false:
		cfg.readnessIsOk = true
	}

	response := probeResponse{
		Status:  "OK",
		Message: fmt.Sprintf("ReadnessProbe flag changed, now it's: %t", cfg.readnessIsOk),
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Service State
// @Summary		Increase memory consumption.
// @Description	Gradual increase in memory consumption for OOM.
// @Produce		json
// @Success		200
// @Router		/api/v1/oomkill [get]
func handlerCallOOM(w http.ResponseWriter, _ *http.Request) {
	debug.SetGCPercent(-1)

	go func() {
		var oom [][]byte
		for {
			hog := make([]byte, 1024*1024*100)
			oom = append(oom, hog)
		}
	}()

	response := probeResponse{
		Status:  "OK",
		Message: "The process of gradual memory recycling has started, wait for OOM",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

// @Tags        Service State
// @Summary		Increase cpu consumption.
// @Description	Increase cpu consumption for throttling with duration 1 minute.
// @Produce		json
// @Success		200
// @Router		/api/v1/throttling [get]
func handlerThrottlingCPU(w http.ResponseWriter, _ *http.Request) {
	numCPU := runtime.NumCPU()
	work := func(wg *sync.WaitGroup) {
		defer wg.Done()

		end := time.Now().Add(time.Second * 180)
		for time.Now().Before(end) {
			_ = math.Sqrt(float64(time.Now().UnixNano()))
		}
	}
	go func() {
		var wg sync.WaitGroup
		wg.Add(numCPU)

		for i := 0; i < numCPU; i++ {
			go work(&wg)
		}
	}()

	response := probeResponse{
		Status:  "OK",
		Message: "The process of cpu recycling has started, duration: 3 minutes",
		Host:    app.ReplicaID(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
