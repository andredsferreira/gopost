package main

import (
	"fmt"
	"goweb01/data"
	"goweb01/db"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	db.ConnectDatabase()

	files := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("POST /register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("GET /requestinfo", handleRequestInfo)
	http.HandleFunc("GET /randomnumbers", handleRandomNumbers)
	http.HandleFunc("GET /date", handleDate)

	fmt.Println("server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func handleRequestInfo(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/request-info.html"))
	data := struct {
		URL           string
		Host          string
		Protocol      string
		Method        string
		ContentLength int64
	}{
		URL:           r.URL.String(),
		Host:          r.Host,
		Protocol:      r.Proto,
		Method:        r.Method,
		ContentLength: r.ContentLength,
	}
	tmpl.Execute(w, data)
}

func handleRandomNumbers(w http.ResponseWriter, r *http.Request) {
	nums := make([]int, 10)
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	for i := 0; i < 10; i++ {
		nums[i] = rng.Intn(50) + 1
	}
	tmpl := template.Must(template.ParseFiles("templates/random-numbers.html"))
	tmpl.Execute(w, nums)
}

func handleDate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("date.html").Funcs(template.FuncMap{
		"fdate": formatDate,
	}).ParseFiles("templates/date.html"))
	tmpl.Execute(w, time.Now())

}
