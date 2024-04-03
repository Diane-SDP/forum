package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		println("cookie inexistant")
		println(err.Error())
		return
	}
	uuid := cookie.Value
	idCurrent := models.GetIDFromUUID(uuid)
	parts := strings.Split(r.URL.Path, "/")
	idpost, _ := strconv.Atoi(parts[len(parts)-1])
	post := models.GetPost(idpost)
	user := models.GetUser(idCurrent)
	if post.User != user {
		http.Error(w, "Permission Denied", http.StatusForbidden)
		return
	}

	tmpl, err := template.ParseFiles("./view/edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
