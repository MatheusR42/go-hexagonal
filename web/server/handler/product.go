package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusr42/go-hexagonal/adpters/dto"
	"github.com/matheusr42/go-hexagonal/application"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDTO dto.Product
		if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		if err = json.NewEncoder(w).Encode(product); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(product); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
