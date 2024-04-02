package controllers

import (
	"database/sql"
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func LikedPostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var exists bool
	var id int
	parts := strings.Split(r.URL.Path, "/")
	idint, _ := strconv.Atoi(parts[len(parts)-1])
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", idint).Scan(&exists)
	if err != nil {
		panic(err)
	}
	if !exists { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	var data models.DataProfil
	cookie, err := r.Cookie("user")
	if err != nil {
		data.IsMe = false
		data.IsConnected = false
	} else {
		id = models.GetIDFromUUID(cookie.Value, db)
		data.IsConnected = true
		data.CurrentUser = models.GetUser(id, db)
		if id == idint {
			data.IsMe = true
		}
	}
	info := models.GetUser(idint, db)
	data.User = info
	posts := models.GetPosts(db)
	var likedPosts []models.Post
	for _, elt := range posts {
		if models.IsLikedBy(elt.Id, idint, db) {
			if data.IsConnected {
				elt.IsLiked = models.IsLikedBy(elt.Id, id, db)
			}
			likedPosts = append(likedPosts, elt)
		}
	}
	data.Posts = likedPosts
	tmpl, err := template.ParseFiles("./view/profil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
