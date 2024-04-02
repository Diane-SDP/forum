package models

import (
	"database/sql"
	"log"
)

type User struct {
	Id       int
	Pseudo   string
	Bio      string
	Mail     string
	Image    string
	Password string
}

func GetUser(id int, db *sql.DB) User {
	rows, err := db.Query("SELECT username, password, bio, email, image FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Pseudo, &user.Password, &user.Bio, &user.Mail, &user.Image)
		if err != nil {
			log.Fatal(err)
		}
	}
	user.Id = id
	return user
}

func GetIDFromUUID(uuid string, db *sql.DB) int {
	rows, err := db.Query("SELECT id FROM users WHERE uuid = ?", uuid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return id
}
