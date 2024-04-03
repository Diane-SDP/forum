package controllers

import (
	models "forum/model"
	"log"
	"net/http"
	"strconv"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("titleInput")
	content := r.FormValue("contentInput")
	id := r.FormValue("id")
	var path string
	idint, _ := strconv.Atoi(id)
	file, fileheader, err := r.FormFile("mediaInput")
	if err == nil {
		path = models.Upload(file, fileheader)
	} else {
		path = models.GetPost(idint).Image
	}
	_, err = models.DB.Exec("UPDATE posts SET title = ?, content = ?, image = ? WHERE id = ?", title, content, path, idint)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/commentaire/"+id, http.StatusSeeOther)

}
