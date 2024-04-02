package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/login" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}
	// Récupérer les données du formulaire
	username := r.FormValue("usernameLog")
	password := r.FormValue("passwordLog")

	h := md5.New()
	h.Write([]byte(password))
	passwordHash := hex.EncodeToString(h.Sum(nil))

	rows, err := db.Query("SELECT uuid,password FROM users WHERE username = ?", username)
	if err != nil {
		http.Error(w, "Erreur lors de la connexion", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()
	
	var passwordDB string
	var uuidUser string
	for rows.Next() {
		err := rows.Scan(&uuidUser, &passwordDB)
		if err != nil {
			http.Error(w, "Erreur lors de la connexion", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	if passwordDB == "" {
		log.Println("Nom d'utilisateur incorrect")
		http.Redirect(w, r, "/connection?error=badname", http.StatusSeeOther)
		return
	}
	if passwordHash != passwordDB {
		log.Println("Mot de passe incorrect")
		http.Redirect(w, r, "/connection?error=badpassword", http.StatusSeeOther)
		return
	}
		

	http.SetCookie(w, &http.Cookie{
		Name:  "user",
		Value: uuidUser,
	})
	http.Redirect(w, r, "/loginsuccess", http.StatusSeeOther)
}
