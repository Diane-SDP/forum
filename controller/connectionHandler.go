package controllers

import (
	"html/template"
	"net/http"
)

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/connection" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	tmpl, err := template.ParseFiles("./view/connection.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
