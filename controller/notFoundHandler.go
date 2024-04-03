package controllers

import (
	//"database/sql"
	//models "forum/model"

	models "forum/model"
	"html/template"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("./view/notFound.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data models.DataNotFound
	cookie, err := r.Cookie("user")
	if err != nil {
		data.IsConnected = false
	} else {
		data.IsConnected = true
		uuid := cookie.Value
		data.CurrentUser = models.GetUser(models.GetIDFromUUID(uuid))
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
