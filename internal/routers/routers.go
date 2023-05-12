package routers

import (
    "github.com/go-chi/chi/v5"
    "github.com/shipherman/go-metrics/internal/handlers"

)


func InitRouter() chi.Router {
    h := handlers.NewHandler()

    // Routers
    router := chi.NewRouter()
    router.Get("/", h.HandleMain)
    router.Post("/update/{type}/{metric}/{value}", h.HandleUpdate)
    router.Get("/value/gauge/{metric}", h.HandleValue)
    router.Get("/value/counter/{metric}", h.HandleValue)

    return router
}
