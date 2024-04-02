package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func SignupHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.URL.Path != "/signup" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound, db) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	uuid := generateUUID(db)

	exist, msg := UserCheck(username, email, db)

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
	_, err := db.Exec("INSERT INTO users (username, password, email, uuid, bio) VALUES (?, ?, ?, ?, ?)", username, passwordHash, email, uuid, "")
	if err != nil {
		http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "user",
		Value: uuid,
	})
	// Réponse de succès
	http.Redirect(w, r, "/loginsuccess", http.StatusSeeOther)
}

func generateUUID(db *sql.DB) string {
	const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, 12)
	for i := range result {
		result[i] = char[random.Intn(len(char))]
	}
	uuid := string(result)
	rows, err := db.Query("SELECT uuid FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var usedUUID string
	for rows.Next() {
		err := rows.Scan(&usedUUID)
		if err != nil {
			log.Fatal(err)
		}
		if usedUUID == uuid {
			return generateUUID(db)
		}
	}
	return uuid
}

func UserCheck(username string, mail string, db *sql.DB) (bool, string) {
	rows, err := db.Query("SELECT username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var usedName string
	var usedMail string
	for rows.Next() {
		err := rows.Scan(&usedName, &usedMail)
		if err != nil {
			log.Fatal(err)
		}
		if usedName == username {
			return true, "Username déjà utilisé"
		} else if usedMail == mail {
			return true, "Mail déjà utilisé"
		}
	}
	return false, ""
}
