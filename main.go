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

	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", middleware.LoggerMiddleware(handler.HomeHandler))

	mux.HandleFunc("POST /login", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.LoginHandler)))
	mux.HandleFunc("POST /logout", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.LogoutHandler)))
	mux.HandleFunc("POST /register", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.RegisterHandler)))

	mux.HandleFunc("GET /hello", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.HelloHandler)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
