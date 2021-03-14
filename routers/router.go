package routers

import (
	"net/http"

	"github.com/0B1t322/Note-API/note"
	"github.com/gorilla/mux"
)

func New(ns note.NoteService) *mux.Router {
	r := mux.NewRouter()


	r.Handle(
		"/notes",
		http.HandlerFunc(ns.Add),
	).Methods("POST")

	r.Handle(
		"/notes",
		http.HandlerFunc(ns.GetAll),
	).Methods("GET")

	r.Handle(
		"/notes/{id:[0-9]+}",
		http.HandlerFunc(ns.Get),
	).Methods("GET")

	r.Handle(
		"/notes/{id:[0-9]+}",
		http.HandlerFunc(ns.Update),
	).Methods("PUT")

	
	r.Handle(
		"/notes/{id:[0-9]+}",
		http.HandlerFunc(ns.Delete),
	).Methods("DELETE")


	return r
}