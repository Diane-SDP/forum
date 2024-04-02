package controllers

import (
	"database/sql"
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func TopicHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	categories, _ := models.GetCategories(db)
	parts := strings.Split(r.URL.Path, "/")
	idstr := parts[len(parts)-1]

	var exists bool

	erreur := db.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)", idstr).Scan(&exists)
	if erreur != nil {
		panic(erreur)
	}
	if !exists { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	type DataTopic struct {
		Posts       []models.Post
		Current models.Category
		Categories  []models.Category
		IsConnected bool
		CurrentUser models.User
	}
 
	var data DataTopic
	id, _ := strconv.Atoi(idstr)
	data.Current = models.GetCategory(id, db)
	data.Posts = models.GetTopicPosts(id, db)
	data.Categories = categories

	cookie, err := r.Cookie("user")
	if err != nil {
		data.IsConnected = false
	} else {
		data.IsConnected = true
	}
	
	if data.IsConnected {
		iduser := models.GetIDFromUUID(cookie.Value, db)
		for i := range data.Posts {
			data.Posts[i].IsLiked = models.IsLikedBy(data.Posts[i].Id, iduser, db)
			data.Posts[i].IsMine = models.IsMine(data.Posts[i].Id, cookie.Value, db)
		}
	
		data.CurrentUser = models.GetUser(iduser, db)
	}

	tmpl, err := template.ParseFiles("./view/topic.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
