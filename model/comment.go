package models

type Comment struct {
	Content string
	User    User
	Id      int
}

func GetComment(idpost int) []Comment {
	rows, err := DB.Query("SELECT id, idUser, commentaire FROM comment WHERE idPost = ?", idpost)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var commentaire []Comment
	var iduser int
	for rows.Next() {
		var com Comment
		err := rows.Scan(&com.Id, &iduser, &com.Content)
		if err != nil {
			panic(err)
		}
		user := GetUser(iduser)
		com.User = user
		commentaire = append(commentaire, com)
	}
	return commentaire
}
