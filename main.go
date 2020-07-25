package main

import (
	"net/http"
	"product-service/internal"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	internal.NewProductsRoute(router)
	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
