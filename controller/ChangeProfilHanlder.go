package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	models "forum/model"
	"html/template"
	"log"
	"net/http"
)

func ChangeProfilHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/changeprofil" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	cookie, err := r.Cookie("user")
	if err != nil {
		log.Fatal(err)
	}
	id := models.GetIDFromUUID(cookie.Value, db)
	newname := r.FormValue("changename")

	var exists bool
	var info models.User
	info = models.GetUser(id, db)

	erreur := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", newname).Scan(&exists)
	if erreur != nil {
		panic(erreur)
	}
	if !exists {
		if newname != "" {
			_, err = db.Exec("UPDATE users Set username = ? Where id = ?", newname, id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	newpsd := r.FormValue("changepsd")
	vraipswd := r.FormValue("psd")

	prbl := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE password = ?)", newpsd).Scan(&exists)
	if prbl != nil {
		panic(prbl)
	}
	h := md5.New()
	h.Write([]byte(vraipswd))
	vraipswd = hex.EncodeToString(h.Sum(nil))
	if vraipswd == info.Password {
		if !exists {
			if newpsd != "" {
				h := md5.New()
				h.Write([]byte(newpsd))
				newpsd := hex.EncodeToString(h.Sum(nil))
				_, err = db.Exec("UPDATE users set password = ? where id = ?", newpsd, id)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	bio := r.FormValue("biographie")
	if bio != "" {
		_, err = db.Exec("UPDATE users SET bio = ? WHERE id = ?", bio, id)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, fileheader, _ := r.FormFile("img")
	var path string
	if file != nil {
		path = models.Upload(file, fileheader)
		_, err = db.Exec("UPDATE users SET image = ? WHERE id = ?", path, id)
		if err != nil {
			print(err.Error())
		}
	}

	store := r.FormValue("photo")
	if store != "" {
		_, err = db.Exec("UPDATE users SET image = ? WHERE id = ?", store, id)
		if err != nil {
			print(err.Error())
		}
	}

	
	info = models.GetUser(id, db)
	tmpl, err := template.ParseFiles("./view/changeprofil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
