package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/shipherman/go-metrics/internal/storage"
)

// Handler accept slice of metrics structure in json format
func (h *Handler) HandleBatchUpdate(w http.ResponseWriter, r *http.Request) {
	var m []Metrics
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, v := range m {
		switch v.MType {
		case counterType:
			if v.Delta == nil {
				http.Error(w, "metric value should not be empty", http.StatusBadRequest)
				return
			}
			h.Store.UpdateCounter(v.ID, storage.Counter(*v.Delta))
		case gaugeType:
			if v.Value == nil {
				http.Error(w, "metric value should not be empty", http.StatusBadRequest)
				return
			}
			h.Store.UpdateGauge(v.ID, storage.Gauge(*v.Value))
		default:
			http.Error(w, "Incorrect metric type", http.StatusBadRequest)
		}
	}
	w.WriteHeader(http.StatusOK)

}
