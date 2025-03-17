package person

import (
	"github.com/Pxe2k/kaspi-task/pkg"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	*mux.Router
	uc UseCase
}

func New(uc UseCase) *Handler {
	h := Handler{
		mux.NewRouter(),
		uc,
	}

	api := h.PathPrefix("/api").Subrouter()
	api.Use(pkg.SetMiddlewareJSON)

	persons := api.PathPrefix("/persons").Subrouter()

	persons.HandleFunc("/store", h.store).Methods(http.MethodPost)
	persons.HandleFunc("/find/{name}", h.find).Methods(http.MethodGet)
	persons.HandleFunc("/get/{document_id}", h.get).Methods(http.MethodGet)
	persons.HandleFunc("/check/{document_id}", h.check).Methods(http.MethodGet)

	return &h
}
