package main

import (
	"github.com/gorilla/mux"
	"github.factset.com/bloom-filter/add"
	"github.factset.com/bloom-filter/hash"
	"github.factset.com/bloom-filter/member"
	"github.factset.com/bloom-filter/size"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/add", add.Handler)
	r.HandleFunc("/ismember", member.Handler)
	r.HandleFunc("/size", size.Handler)
	r.HandleFunc("/hashcount", hash.Handler)
	return r
}
