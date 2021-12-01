package app

import (
	"github.com/go-chi/chi"
)

// InitRouter creates a chi router instance.
func (h *Handler) InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Route("/location/{order_id}", func(r chi.Router) {
		r.Put("/", h.updateLocationHistory)
		r.Get("/", h.getLocationHistory)
		r.Delete("/", h.deleteLocationHistory)
	})

	return r
}
