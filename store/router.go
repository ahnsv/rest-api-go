package store

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var ctrlr = &Controller{Repository: Repository{}}

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandleFunc
}

type Routes []Route