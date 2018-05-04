package main

import (
	"github.com/gorilla/mux"
	"github.com/bloom-filter/add"
	"github.com/bloom-filter/hash"
	"github.com/bloom-filter/member"
	"github.com/bloom-filter/size"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/add", add.Handler)
	r.HandleFunc("/ismember", member.Handler)
	r.HandleFunc("/size", size.Handler)
	r.HandleFunc("/hashcount", hash.Handler)
	return r
}
