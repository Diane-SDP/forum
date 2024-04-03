package controllers

import (
	models "forum/model"
	"log"
	"net/http"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postcategory" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	titleInput := r.FormValue("titleInput")
	descriptionInput := r.FormValue("descriptionInput")
	file, fileheader, _ := r.FormFile("mediaInput")
	var path string
	if file != nil {
		path = models.Upload(file, fileheader)
	} else {
		path = "./Assets/img/categories.png"
	}

	_, err := models.DB.Exec("INSERT INTO categories (title, description, image) VALUES (?, ?, ?)", titleInput, descriptionInput, path)
	if err != nil {
		http.Error(w, "Erreur lors de la publication de la catégorie", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
