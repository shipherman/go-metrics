package routers

import (
    "github.com/go-chi/chi/v5"
    "github.com/shipherman/go-metrics/internal/handlers"
    "github.com/shipherman/go-metrics/internal/logger"

)


func InitRouter() chi.Router {
    h := handlers.NewHandler()

    // Routers
    router := chi.NewRouter()
    router.Get("/", logger.LogHandler(h.HandleMain()))
    router.Post("/update/{type}/{metric}/{value}", logger.LogHandler(h.HandleUpdate()))
    router.Get("/value/gauge/{metric}", logger.LogHandler(h.HandleValue()))
    router.Get("/value/counter/{metric}", logger.LogHandler(h.HandleValue()))

    return router
}
