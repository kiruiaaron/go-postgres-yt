package router

import (
	"github.com/gorilla/mux"
	"github.com/kiruiaaron/go-postgres-yt/middleware"
)



func Router() *mux.Router{
	router := mux.NewRouter()


	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newstock",middleware.CreateStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{}",middleware.DeleteStock).Methods("DELETE","OPTIONS")

	return router

}