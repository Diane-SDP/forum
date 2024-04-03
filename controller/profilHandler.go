package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	var exists bool
	parts := strings.Split(r.URL.Path, "/")
	idint, _ := strconv.Atoi(parts[len(parts)-1])
	err := models.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", idint).Scan(&exists)
	if err != nil {
		panic(err)
	}
	if !exists { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	var data models.DataProfil
	var id int
	cookie, err := r.Cookie("user")
	if err != nil {
		data.IsMe = false
		data.IsConnected = false
	} else {
		id = models.GetIDFromUUID(cookie.Value)
		data.IsConnected = true
		data.CurrentUser = models.GetUser(id)
		if id == idint {
			data.IsMe = true
		}
	}
	data.Posts = models.GetMyPosts(idint)
	data.User = models.GetUser(idint)
	for i := range data.Posts {
		if data.IsConnected {
			data.Posts[i].IsLiked = models.IsLikedBy(data.Posts[i].Id, id)
		}
		if data.IsMe {
			data.Posts[i].IsMine = true
		}
	}

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
