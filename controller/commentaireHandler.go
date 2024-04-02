package controllers

import (
	"database/sql"
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func CommentaireHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
		id = models.GetIDFromUUID(cookie.Value, db)
	}
	parts := strings.Split(r.URL.Path, "/")
	idint, _ := strconv.Atoi(parts[len(parts)-1])
	var exists bool
	erreur := db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = ?)", idint).Scan(&exists)
	if erreur != nil {
		panic(erreur)
	}
	if !exists { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	data.IdPost = idint
	data.Lescoms = models.GetComment(idint, db)
	data.Post = models.GetPost(idint, db)
	if data.IsConnected {
		data.CurrentUser = models.GetUser(id, db)
		data.Post.IsLiked = models.IsLikedBy(data.IdPost, models.GetIDFromUUID(cookie.Value, db), db)
		data.Post.IsMine = models.IsMine(idint, cookie.Value, db)
	}
	categories, _ := models.GetCategories(db)
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
