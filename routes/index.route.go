package routes

import (
	"net/http"
)

type indexRouteHandler interface {
	Index(w http.ResponseWriter, r *http.Request)
	Product(w http.ResponseWriter, r *http.Request)
	ProcessProduct(w http.ResponseWriter, r *http.Request)
}

func AddHandlers(mux *http.ServeMux, handlers indexRouteHandler) {
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/product", handlers.Product)
	mux.HandleFunc("/process", handlers.ProcessProduct)
}
