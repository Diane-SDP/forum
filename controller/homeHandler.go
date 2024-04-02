package controllers

import (
	"database/sql"
	"encoding/json"
	models "forum/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	var data models.Data
	var uuid string
	var idUser int
	cookie, err := r.Cookie("user")
	if err != nil {
		data.IsConnected = false
	} else {
		data.IsConnected = true
		uuid = cookie.Value
		idUser = models.GetIDFromUUID(uuid, db)
		if err != nil {
			log.Fatal("prbl lors de la récup de la valeur du cookie")
		}
	}
	type DataJSON struct {
		Action string `json:"action"`
		Id     string `json:"id"`
	}
	if r.Method == "POST" {
		var dataJS DataJSON
		if err := json.NewDecoder(r.Body).Decode(&dataJS); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if dataJS.Action == "like" {
			idpost, _ := strconv.Atoi(dataJS.Id)
			models.AddLike(idpost, idUser, uuid, db)
		} else if dataJS.Action == "dislike" {
			idpost, _ := strconv.Atoi(dataJS.Id)
			models.RemoveLike(idpost, idUser, db)
		}
	}

	posts := models.GetPosts(db)
	categories, _ := models.GetCategories(db)
	tmpl, err := template.ParseFiles("./view/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Categories = categories
	if data.IsConnected {
		iduser := models.GetIDFromUUID(cookie.Value, db)
		for i := range posts {
			posts[i].IsLiked = models.IsLikedBy(posts[i].Id, iduser, db)
			posts[i].IsMine = models.IsMine(posts[i].Id, uuid, db)
		}
		data.CurrentUser = models.GetUser(iduser, db)
	}
	data.Posts = posts
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
