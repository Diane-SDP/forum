package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	categories, _ := models.GetCategories()
	parts := strings.Split(r.URL.Path, "/")
	idstr := parts[len(parts)-1]

	var exists bool

	erreur := models.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)", idstr).Scan(&exists)
	if erreur != nil {
		panic(erreur)
	}
	if !exists { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	type DataTopic struct {
		Posts       []models.Post
		Current     models.Category
		Categories  []models.Category
		IsConnected bool
		CurrentUser models.User
	}

	var data DataTopic
	id, _ := strconv.Atoi(idstr)
	data.Current = models.GetCategory(id)
	data.Posts = models.GetTopicPosts(id)
	data.Categories = categories

	cookie, err := r.Cookie("user")
	if err != nil {
		data.IsConnected = false
	} else {
		data.IsConnected = true
	}

	if data.IsConnected {
		iduser := models.GetIDFromUUID(cookie.Value)
		for i := range data.Posts {
			data.Posts[i].IsLiked = models.IsLikedBy(data.Posts[i].Id, iduser)
			data.Posts[i].IsMine = models.IsMine(data.Posts[i].Id, cookie.Value)
		}

		data.CurrentUser = models.GetUser(iduser)
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
