package handlers

import (
    "net/http"
    "encoding/json"
)


// Handle JSON request to return value
func (h *Handler) HandleJSONValue(w http.ResponseWriter, r *http.Request) {
    var m Metrics

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    switch m.MType {
        case counterType:
            v, ok := h.Store.CounterData[m.ID]
            if !ok {
                http.Error(w, "not found", http.StatusNotFound)
                return
            }
            vPtr := int64(v)
            m.Delta = &vPtr
        case gaugeType:
            v, ok := h.Store.GaugeData[m.ID]
            if !ok {
                http.Error(w, "not found", http.StatusNotFound)
                return
            }
            vPtr := float64(v)
            m.Value = &vPtr
    }


    resp, err := json.Marshal(m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // respond to agent
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(resp)
}
