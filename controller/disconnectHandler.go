package controllers

import (
	"net/http"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/deco" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
	}

	c := &http.Cookie{
		Name:   "user",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
