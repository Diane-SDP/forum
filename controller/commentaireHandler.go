package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func CommentaireHandler(w http.ResponseWriter, r *http.Request) {
	type dataComment struct {
		IdPost      int
		CurrentUser models.User
		Lescoms     []models.Comment
		Post        models.Post
		Categories  []models.Category
		IsConnected bool
	}
	var data dataComment
	var id int
	cookie, err := r.Cookie("user")
	if err != nil {
		println("non connecté")
	} else {
		data.IsConnected = true
		id = models.GetIDFromUUID(cookie.Value)
	}
	parts := strings.Split(r.URL.Path, "/")
	idint, _ := strconv.Atoi(parts[len(parts)-1])
	var exists bool
	erreur := models.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = ?)", idint).Scan(&exists)
	if erreur != nil {
		panic(erreur)
	}
	if !exists { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	data.IdPost = idint
	data.Lescoms = models.GetComment(idint)
	data.Post = models.GetPost(idint)
	if data.IsConnected {
		data.CurrentUser = models.GetUser(id)
		data.Post.IsLiked = models.IsLikedBy(data.IdPost, models.GetIDFromUUID(cookie.Value))
		data.Post.IsMine = models.IsMine(idint, cookie.Value)
	}
	categories, _ := models.GetCategories()
	data.Categories = categories

	tmpl, err := template.ParseFiles("./view/commentaire.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
