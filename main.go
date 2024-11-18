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
	user := r.FormValue("username")
	password := r.FormValue("password")
	if len(user) < 3 || len(password) < 3 {
		http.Error(w, "invalid username/password", http.StatusNotAcceptable)
		return
	}
	if eu, _ := data.GetUserByUsername(user); eu.Username == user {
		http.Error(w, "user already registered", http.StatusNotAcceptable)
		return
	}
	hp, err := hashPassword(password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}
	err = data.AddUser(user, hp)
	if err != nil {
		http.Error(w, "could not register user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User registered successfully.")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
	user := r.FormValue("username")
	password := r.FormValue("password")
	eu, err := data.GetUserByUsername(user)
	if err != nil {
		http.Error(w, "invalid username", http.StatusUnauthorized)
	}
	if !checkPasswordHash(password, eu.HashedPassword) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
	}
	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, "error generating JWT token", http.StatusInternalServerError)
	}
	http.SetCookie(w, &http.Cookie{
		Name: "jwt_token",
		Value: token,
		HttpOnly: true,
	})
}

func handleLogout(w http.ResponseWriter, r *http.Request) {

}

func handleUsers(w http.ResponseWriter, r *http.Request) {

}
