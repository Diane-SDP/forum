package models

import (
	"log"
)

func IsLikedBy(id int, iduser int) bool {
	rows, err := DB.Query("SELECT idUser FROM likes WHERE idPost = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var idLike int
	for rows.Next() {
		err := rows.Scan(&idLike)
		if err != nil {
			panic(err)
		}
		if iduser == idLike {
			return true
		}
	}
	return false
}

func AddLike(idpost int, iduser int, uuid string) {
	_, err := DB.Exec("INSERT INTO likes (idPost, idUser, uuidUser) VALUES (?, ?, ?)", idpost, iduser, uuid)
	if err != nil {
		log.Println(err)
	}
	nblikes := GetLikes(idpost) + 1
	_, err = DB.Exec("UPDATE posts SET likes = ? WHERE id = ?", nblikes, idpost)
	if err != nil {
		log.Println(err)
	}
}

func RemoveLike(idpost int, iduser int) {
	_, err := DB.Exec("DELETE FROM likes WHERE idpost = ? AND iduser = ?", idpost, iduser)
	if err != nil {
		log.Println(err)
	}
	nblikes := GetLikes(idpost) - 1
	_, err = DB.Exec("UPDATE posts SET likes = ? WHERE id = ?", nblikes, idpost)
	if err != nil {
		log.Println(err)
	}
}
