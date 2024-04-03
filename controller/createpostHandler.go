package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	categories, _ := models.GetCategories()
	type createPostData struct {
		Categories  []models.Category
		CurrentUser models.User
	}
	var data createPostData
	data.Categories = categories
	cookie, err := r.Cookie("user")
	if err != nil {
		panic(err)
	}
	uuid := cookie.Value
	if err != nil {
		panic(err)
	}
	id := models.GetIDFromUUID(uuid)
	tmpl, err := template.ParseFiles("./view/create.html")

	data.CurrentUser = models.GetUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
