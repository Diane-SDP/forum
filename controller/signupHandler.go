package controllers

import (
	"crypto/md5"
	"encoding/hex"
	models "forum/model"
	"log"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	uuid := models.GenerateUUID()

	exist, msg := models.UserCheck(username, email)

	if exist {
		if msg == "Username déjà utilisé" {
			log.Println("Nom d'utilisateur déjà utilisé")
			http.Redirect(w, r, "/connection?error=badnameSignUp", http.StatusSeeOther)
			return
		}
		log.Println("Email déjà utilisé")
		http.Redirect(w, r, "/connection?error=badmailSignUp", http.StatusSeeOther)
		return
	}

	h := md5.New()
	h.Write([]byte(password))
	passwordHash := hex.EncodeToString(h.Sum(nil))

	// Insérer les données dans la base de données
	_, err := models.DB.Exec("INSERT INTO users (username, password, email, uuid, bio) VALUES (?, ?, ?, ?, ?)", username, passwordHash, email, uuid, "")
	if err != nil {
		http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		Value:  uuid,
		MaxAge: 3600,
	})
	// Réponse de succès
	http.Redirect(w, r, "/loginsuccess", http.StatusSeeOther)
}
