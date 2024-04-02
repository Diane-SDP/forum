package controllers

import (
	"database/sql"
	models "forum/model"
	"html/template"
	"log"
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/create" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	categories, _ := models.GetCategories(db)
	type createPostData struct {
		Categories []models.Category
		CurrentUser models.User
	}
	var data createPostData
	data.Categories = categories
	cookie, err := r.Cookie("user")
	if err != nil {
		log.Fatal(err)
	}
	uuid := cookie.Value
	if err != nil {
		log.Fatal(err)
	}
	id := models.GetIDFromUUID(uuid, db)
	tmpl, err := template.ParseFiles("./view/create.html")

	data.CurrentUser = models.GetUser(id, db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
