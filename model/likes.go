package models

import (
	"database/sql"
	"log"
)

func IsLikedBy(id int, iduser int, db *sql.DB) bool {
	rows, err := db.Query("SELECT idUser FROM likes WHERE idPost = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var idLike int
	for rows.Next() {
		err := rows.Scan(&idLike)
		if err != nil {
			log.Fatal(err)
		}
		if iduser == idLike {
			return true
		}
	}
	return false
}

func AddLike(idpost int, iduser int, uuid string, db *sql.DB) {
	_, err := db.Exec("INSERT INTO likes (idPost, idUser, uuidUser) VALUES (?, ?, ?)", idpost, iduser, uuid)
	if err != nil {
		log.Println(err)
	}
	nblikes := GetLikes(idpost, db) + 1
	_, err = db.Exec("UPDATE posts SET likes = ? WHERE id = ?", nblikes, idpost)
	if err != nil {
		log.Println(err)
	}
}

func RemoveLike(idpost int, iduser int, db *sql.DB) {
	_, err := db.Exec("DELETE FROM likes WHERE idpost = ? AND iduser = ?", idpost, iduser)
	if err != nil {
		log.Println(err)
	}
	nblikes := GetLikes(idpost, db) - 1
	_, err = db.Exec("UPDATE posts SET likes = ? WHERE id = ?", nblikes, idpost)
	if err != nil {
		log.Println(err)
	}
}

