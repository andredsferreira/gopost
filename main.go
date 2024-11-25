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

	mux.HandleFunc("/", middleware.LoggerMiddleware(handler.IndexHandler))

	mux.HandleFunc("GET /home", middleware.LoggerMiddleware(handler.HomeHandler))

	mux.HandleFunc("POST /login", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.LoginHandler)))
	mux.HandleFunc("POST /logout", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.LogoutHandler)))
	mux.HandleFunc("POST /register", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.RegisterHandler)))

	mux.HandleFunc("GET /post", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.GetUserPostsHandler)))
	mux.HandleFunc("POST /post", middleware.LoggerMiddleware(
		middleware.AuthMiddleware(handler.CreatePostHandler)))
	mux.HandleFunc("GET /post/all", middleware.LoggerMiddleware(handler.GetAllPostsHandler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
