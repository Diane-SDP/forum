package models

import (
	"database/sql"
	"log"
)

type Comment struct {
	Content string
	User    User
	Id      int
}

func GetComment(idpost int, db *sql.DB) []Comment {
	rows, err := db.Query("SELECT id, idUser, commentaire FROM comment WHERE idPost = ?", idpost)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var commentaire []Comment
	var iduser int
	for rows.Next() {
		var com Comment
		err := rows.Scan(&com.Id, &iduser, &com.Content)
		if err != nil {
			log.Fatal(err)
		}
		user := GetUser(iduser, db)
		com.User = user
		commentaire = append(commentaire, com)
	}
	return commentaire
}
