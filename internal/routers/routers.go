package routers

import (
    // "context"

    "github.com/go-chi/chi/v5"
    "github.com/shipherman/go-metrics/internal/handlers"
    "github.com/shipherman/go-metrics/internal/middleware/logger"
    "github.com/shipherman/go-metrics/internal/middleware/gzip"
    "github.com/shipherman/go-metrics/internal/middleware/crypt"
    "github.com/shipherman/go-metrics/internal/options"

)


func InitRouter(cfg options.Options, h handlers.Handler) (chi.Router, error) {
    // Context for router
    // ctx := context.WithValue(context.Background(), key, cfg.Key)

    // Routers
    router := chi.NewRouter()
    router.Use(logger.LogHandler)
    router.Use(gzip.GzipHandle)
    router.Use(crypt.CheckReqSign)
    router.Get("/", h.HandleMain)
    router.Get("/ping", h.HandlePing)
    router.Post("/updates/", h.HandleBatchUpdate)
    router.Post("/update/{type}/{metric}/{value}", h.HandleUpdate)
    router.Get("/value/gauge/{metric}", h.HandleValue)
    router.Get("/value/counter/{metric}", h.HandleValue)
    router.Post("/value/",h.HandleJSONValue)
    router.Post("/update/",h.HandleJSONUpdate)

    return router, nil
}
