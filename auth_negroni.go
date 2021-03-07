package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type auth struct {
	Username string
	Password string
}

func (auth *auth) ServeHTTP(writer http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")

	if username != auth.Username || password != auth.Password {
		http.Error(writer, "Unauthorized", 401)
		return
	}
	ctx := context.WithValue(req.Context(), "username", username)
	req = req.WithContext(ctx)
	next(writer, req)
}

func hello(writer http.ResponseWriter, req *http.Request) {
	username := req.Context().Value("username").(string)
	fmt.Fprintf(writer, "Hello %s", username)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")
	n := negroni.Classic()
	n.Use(&auth{
		Username: "admin",
		Password: "password",
	})
	// Always put the use part in the end
	n.UseHandler(router)
	http.ListenAndServe(":8000", n)

}
