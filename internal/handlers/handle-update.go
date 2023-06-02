package handlers

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "strconv"
    "github.com/shipherman/go-metrics/internal/storage"

)


// Handle URI request to update metric value
func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
    // get context params
    metricType := chi.URLParam(r, "type")
    metric := chi.URLParam(r, "metric")
    value := chi.URLParam(r, "value")

    // find out metric type
    switch metricType {
        case counterType:
            v, err := strconv.Atoi(value)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            h.Store.UpdateCounter(metric, storage.Counter(v))
        case gaugeType:
            v, err := strconv.ParseFloat(value, 64)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
            }
            h.Store.UpdateGauge(metric, storage.Gauge(v))
        default:
            http.Error(w, "Incorrect metric type", http.StatusBadRequest)
    }
}
