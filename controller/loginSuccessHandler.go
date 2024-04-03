package controllers

import (
	"html/template"
	"net/http"
)

func LoginSuccessHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/loginsuccess" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
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
