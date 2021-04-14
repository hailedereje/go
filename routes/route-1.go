package routes

import (
	"fmt"
	"go_project/controllers"
	
	"log"
	"net/http"

	"github.com/gorilla/mux"
)





func Router(){

	r := mux.NewRouter()

	r.HandleFunc("/movies",controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies/create",controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",controllers.DeleteMovie).Methods("DELETE")


	fmt.Println("running at port 5000")
	log.Fatal(http.ListenAndServe(":5000",r))
}