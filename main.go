package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}
	r := NewRouter()
	http.Handle("/", r)
	fmt.Println(http.ListenAndServe(port, nil))
}
