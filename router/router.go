package router

import (
	"shopping-list/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/product/{id}", middleware.GetOneProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", middleware.GetAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product", middleware.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/product/{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/product/{id}", middleware.DeleteProduct).Methods("DELETE", "OPTIONS")

	return router
}
