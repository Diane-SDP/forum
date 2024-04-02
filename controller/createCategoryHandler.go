package controllers

import (
	"database/sql"
	models "forum/model"
	"html/template"
	"log"
	"net/http"
)

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/createcategory" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	tmpl, err := template.ParseFiles("./view/createCategory.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data models.Data
	cookie, err := r.Cookie("user")
	if err != nil {
		log.Fatal(err)
	}
	uuid := cookie.Value
	data.CurrentUser = models.GetUser(models.GetIDFromUUID(uuid, db), db)

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
