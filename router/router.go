package router

import (
	"github.com/gorilla/mux"
	"github.com/rapterx/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")

	return router
}
