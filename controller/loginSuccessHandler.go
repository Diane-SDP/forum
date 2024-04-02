package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
)

func LoginSuccessHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/loginsuccess" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                                  // Et on arrête notre code ici !
	}

	tmpl, err := template.ParseFiles("./view/loginsuccess.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
