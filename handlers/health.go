// Package handlers : Health check handler
package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := map[string]string{"status": "up"}

	err := enc.Encode(resp)
	if err != nil {
		panic("unable to encode response")
	}
}
