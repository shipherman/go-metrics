package routers

import (
    "github.com/go-chi/chi/v5"
    "github.com/shipherman/go-metrics/internal/handlers"
    "github.com/shipherman/go-metrics/internal/logger"
    "github.com/shipherman/go-metrics/internal/gzip"
    "github.com/shipherman/go-metrics/internal/options"

)


func InitRouter(cfg options.Options) (chi.Router, handlers.Handler, error) {
    h, err := handlers.NewHandler(cfg)
    if err != nil {
        return nil, h, err
    }

    // Routers
    router := chi.NewRouter()
    router.Use(logger.LogHandler)
    router.Use(gzip.GzipHandle)
    router.Get("/", h.HandleMain)
    router.Post("/update/{type}/{metric}/{value}", h.HandleUpdate)
    router.Get("/value/gauge/{metric}", h.HandleValue)
    router.Get("/value/counter/{metric}", h.HandleValue)
    router.Post("/value/",h.HandleJSONValue)
    router.Post("/update/",h.HandleJSONUpdate)

    return router, h, nil
}
