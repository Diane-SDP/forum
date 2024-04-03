package controllers

import (
	models "forum/model"
	"log"
	"net/http"
	"strconv"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	cookie, err := r.Cookie("user")
	if err != nil {
		http.Error(w, "Utilisateur non connecté", http.StatusUnauthorized)
		return
	}
	idUser := models.GetIDFromUUID(cookie.Value)
	titleInput := r.FormValue("titleInput")
	contentInput := r.FormValue("contentInput")
	categoryInput, _ := strconv.Atoi(r.FormValue("posts-category"))

	file, fileheader, _ := r.FormFile("mediaInput")
	var path string
	if file != nil {
		path = models.Upload(file, fileheader)
	}

	_, err = models.DB.Exec("INSERT INTO posts (title, content, idUser, idCategory, image, likes) VALUES (?, ?, ?, ?, ?, ?)", titleInput, contentInput, idUser, categoryInput, path, 0)
	if err != nil {
		http.Error(w, "Erreur lors de la publication du post", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
