package main

import (
	"goweb01/db"
	"goweb01/handler"
	"log"
	"net/http"
)

func init() {
	db.ConnectDatabase()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", handler.LoginHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
