package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler *Handler) http.Handler {
	router := mux.NewRouter()
	router.Path("/users").Methods(http.MethodPost).HandlerFunc(userHandler.Register)
	return router
}
