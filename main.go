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

	middlewareStack := middleware.CreateStack(
		middleware.LoggerMiddleware,
		middleware.LoggerMiddleware,
	)

	mux.HandleFunc("POST /login", handler.LoginHandler)
	mux.HandleFunc("POST /logout", handler.LogoutHandler)
	mux.HandleFunc("POST /register", handler.RegisterHandler)

	mux.HandleFunc("GET /hello", handler.HelloHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: middlewareStack(mux),
	}
	log.Fatal(s.ListenAndServe())
}
