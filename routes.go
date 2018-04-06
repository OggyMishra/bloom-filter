package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r = mux.Router.StrictSlash(true)
	r.Handle("/add", Handler),
	r.Handle("/ismember", Handler),
	r.Handle("/size", Handler),
	r.Handle("/hashcount", Handler),

	return r
}
