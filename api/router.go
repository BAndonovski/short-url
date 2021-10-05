package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router(root *mux.Router) *mux.Router {
	root.HandleFunc("/encode", encode).Methods(http.MethodPost)
	root.HandleFunc("/decode", decode).Methods(http.MethodGet)
	return root
}
