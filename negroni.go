package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8000", n)
}
