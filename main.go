package main

import (
	"fmt"
	"goweb01/data"
	"goweb01/db"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db.ConnectDatabase()

	files := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("POST /register", handleRegister)
	http.HandleFunc("POST /login", handleLogin)
	http.HandleFunc("POST /logout", handleLogout)
	http.HandleFunc("GET /users", handleUsers)

	fmt.Println("server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("username")
	p := r.FormValue("password")
	if len(u) < 3 || len(p) < 3 {
		http.Error(w, "invalid username/password", http.StatusNotAcceptable)
		return
	}
	if eu, _ := data.GetUserByUsername(u); eu.Username == u {
		http.Error(w, "user already registered", http.StatusNotAcceptable)
		return
	}
	hp, err := hashPassword(p)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}
	err = data.AddUser(u, hp)
	if err != nil {
		http.Error(w, "could not register user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User registered successfully.")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

}

func handleLogout(w http.ResponseWriter, r *http.Request) {

}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	
}
