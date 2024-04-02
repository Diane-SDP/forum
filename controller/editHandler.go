package controllers

import (
	"database/sql"
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func EditHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	cookie, err := r.Cookie("user")
	if err != nil {
		println("cookie inexistant")
		println(err.Error())
		return
	}
	uuid := cookie.Value
	idCurrent := models.GetIDFromUUID(uuid, db)
	parts := strings.Split(r.URL.Path, "/")
	idpost, _ := strconv.Atoi(parts[len(parts)-1])
	post := models.GetPost(idpost, db)
	user := models.GetUser(idCurrent, db)
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
