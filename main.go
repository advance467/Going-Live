package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ListElements(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is a list of Elements</h1>"))
}

func SingleElement(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is a single new Element</h1>"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is a Tests/h1>"))
}

func LoadElementRoutes(mainRouter *mux.Router) {

	s := mainRouter.PathPrefix("/elements").Subrouter()
	s.HandleFunc("", ListElements).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", SingleElement).Methods("GET")

}

func LoadOtherRoutes(mainRouter *mux.Router) {

	s := mainRouter.PathPrefix("/other").Subrouter()
	s.HandleFunc("/test", TestHandler).Methods("GET")

}

func MainRoutes() *mux.Router {

	r := mux.NewRouter()

	LoadElementRoutes(r)
	LoadOtherRoutes(r)

	return r

}

func main() {

	http.ListenAndServe(":3000", MainRoutes())

}
