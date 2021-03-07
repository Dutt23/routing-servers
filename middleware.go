package main

import (
	"fmt"
	"log"
	"net/http"
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

func main() {
	h := http.HandlerFunc(hello)
	logger := logger{hanlder: h}
	http.ListenAndServe(":8000", &logger)
}
