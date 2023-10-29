package app

import (
	"os"
)

var ReplicaID = func() string {
	h, _ := os.Hostname()
	return h
}
