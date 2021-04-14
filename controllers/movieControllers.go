package controllers

import (
	"encoding/json"
	"go_project/models"
	"go_project/static"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovies(resp http.ResponseWriter, req *http.Request) {
	
	resp.Header().Set("Content-Type","application/json")
	json.NewEncoder(resp).Encode(static.Movies)
}

func GetMovie(resp http.ResponseWriter, req *http.Request) {
	
	resp.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	
	for _,movie := range static.Movies{
		if movie.ID == params["id"] {
			json.NewEncoder(resp).Encode(movie)
			break
		}
	}
}

func CreateMovie(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type","application/json")
	var movie models.Movie
	json.NewDecoder(req.Body).Decode(&movie)
	static.Movies = append(static.Movies, movie)
}

func UpdateMovie(resp http.ResponseWriter, req *http.Request) {
	
	var movie models.Movie
	resp.Header().Set("Content-Type","application/json") 
	DeleteMovie(resp,req)

	json.NewDecoder(req.Body).Decode(&movie)
	static.Movies = append(static.Movies, movie)
	json.NewEncoder(resp).Encode(movie)


}

func DeleteMovie(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)
	for i,movie := range static.Movies{
		if movie.ID == params["id"] {
			static.Movies = append(static.Movies[:i],static.Movies[i+1:]... )
			break
		}
	}
	// json.NewEncoder(resp).Encode(static.Movies)
}

/*
    ptr = m			
	json.NewDecoder(req.Body).Decode(&movie)
	ptr.Isbn = movie.Isbn
	ptr.Title = movie.Title
	ptr.Director = movie.Director

	json.NewEncoder(resp).Encode(m)
	break
*/