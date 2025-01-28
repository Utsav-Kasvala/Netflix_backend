package router

import (
	controller "github.com/Utsav_Kasvala/Netflix_backend/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.AddMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkasWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
