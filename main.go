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
	http.Handle("GET /hello", middleware.AuthMiddleware(http.HandlerFunc(handler.HelloHandler)))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
