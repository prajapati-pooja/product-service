package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewProductsRoute(router *mux.Router) {
	router.Use(setContentTypeHeadersssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss)
	router.HandleFunc("/products", productsHandler)
}

func setContentTypeHeadersssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(responseWriter, request)
	})
}

func productsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	service := NewProductService(NewProductRepository())
	productType := request.URL.Query().Get("type")
	products := service.GetAll(productType)

	json.NewEncoder(responseWriter).Encode(products)
	//if err != nil {
	//	log.Printf("failed to write response %s", err.Error())
	//}
}