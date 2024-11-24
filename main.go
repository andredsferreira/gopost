package main

import (
	"goweb01/db"
	"goweb01/handler"
	"goweb01/middleware"
	"log"
	"net/http"
)

func init() {
	db.ConnectDatabase()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", handler.LoginHandler)
	mux.HandleFunc("POST /logout", handler.LogoutHandler)
	mux.HandleFunc("POST /register", handler.RegisterHandler)
	mux.HandleFunc("GET /hello", middleware.AuthMiddleware(handler.HelloHandler))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
