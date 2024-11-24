package handler

import (
	"fmt"
	"goweb01/model"
	"goweb01/service"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	u, err := model.GetUserByUsername(r.FormValue("username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	p := r.FormValue("password")
	if !service.CheckPasswordHash(p, u.Password) {
		http.Error(w, "wrong password", http.StatusNotAcceptable)
	}
	t, err := service.GenerateJWT(u.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    t,
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "successful login")
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! You are authorized.")
}