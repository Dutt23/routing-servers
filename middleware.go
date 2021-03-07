package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type logger struct {
	hanlder http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Start")
	l.hanlder.ServeHTTP(w, req)
	log.Println("Finish")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hi foo")
}
func greetUser(w http.ResponseWriter, req *http.Request) {
	user := mux.Vars(req)["user"]
	fmt.Fprintf(w, "hi %s\n", user)
}
func defineRouters(router *mux.Router) {
	router.HandleFunc("/foo", greet).Methods("GET")
	router.HandleFunc("/host", greet).Methods("GET").Host("www.foo.com")
	// Url will take only characters
	router.HandleFunc("/users/{user:[a-z,A-Z]+}", greetUser).Methods("GET")
}

func main() {
	// h := http.HandlerFunc(hello)
	r := mux.NewRouter()
	defineRouters(r)
	logger := logger{hanlder: r}
	http.ListenAndServe(":8000", &logger)
}
