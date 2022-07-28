package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mrdibre/hexagonal-arch-go/application"
	"github.com/urfave/negroni"
	"net/http"
)

func MakeProductHandler(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface)  {
	router.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(service)))).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(req)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(res).Encode(product)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}