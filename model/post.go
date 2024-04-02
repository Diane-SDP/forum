package models

import (
	"database/sql"
	"log"
	"strconv"
)

type Post struct {
	Id            int
	User          User
	Title         string
	Content       string
	Category      Category
	Image         string
	Likes         int
	IsLiked       bool
	Comment       string
	IsMine        bool
}

func GetPosts(db *sql.DB) []Post {
	rows, err := db.Query("SELECT id, title, content, idUser, idCategory, image, likes FROM posts")
	if err != nil {
		log.Fatal(err)
		print("GETPOSTS")
		return nil
	}
	defer rows.Close()
	var posts []Post
	var idUser int
	var idCategory string
	var pathImg string
	var likes int
	var idpost int
	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &post.Title, &post.Content, &idUser, &idCategory, &pathImg, &likes)
		if err != nil {
			log.Fatal(err)
		}
		idCategoryINT, _ := strconv.Atoi(idCategory)
		post.Category = GetCategory(idCategoryINT, db)
		post.User = GetUser(idUser, db)
		if err != nil {
			log.Fatal(err)
			print("GETPOSTS")
			return nil
		}
		post.Image = pathImg
		post.Likes = likes
		post.Id = idpost
		if err != nil {
			log.Fatal(err)
		}
		post.Comment = CountComment(post.Id, db)
		posts = append(posts, post)
	}
	return posts
}

func GetPost(id int, db *sql.DB) Post {
	rows, err := db.Query("SELECT id, title, content, idUser, idCategory, image, likes FROM posts WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var post Post
	var idpost int
	var idUser int
	var idCategory int
	for rows.Next() {
		err := rows.Scan(&idpost, &post.Title, &post.Content, &idUser, &idCategory, &post.Image, &post.Likes)
		if err != nil {
			log.Fatal(err)
		}
		post.Category = GetCategory(idCategory, db)
		post.User = GetUser(idUser, db)
		post.Id = idpost
		post.Comment = CountComment(post.Id, db)
		if err != nil {
			log.Fatal(err)
		}
	}
	return post
}

func GetMyPosts(id int, db *sql.DB) []Post {
	rows, err := db.Query("SELECT id, title, content, idUser, idCategory, image, likes FROM posts WHERE idUser = ?", id)
	if err != nil {
		log.Fatal(err)
		print("GETPOSTS")
		return nil
	}
	defer rows.Close()
	var posts []Post
	var idUser int
	var idCategory string
	var pathImg string
	var likes int
	var idpost int
	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &post.Title, &post.Content, &idUser, &idCategory, &pathImg, &likes)
		if err != nil {
			log.Fatal(err)
			print("GETPOSTS")
			return nil
		}
		idCategoryINT, _ := strconv.Atoi(idCategory)
		post.Category = GetCategory(idCategoryINT, db)
		post.User = GetUser(idUser, db)
		post.Image = pathImg
		post.Likes = likes
		post.Id = idpost
		if err != nil {
			log.Fatal(err)
			print("GETPOSTS")
			return nil
		}
		post.Comment = CountComment(post.Id, db)
		posts = append(posts, post)
	}
	return posts
}

func GetTopicPosts(id int, db *sql.DB) []Post {
	rows, err := db.Query("SELECT id, title, content, idUser, idcategory, image, likes FROM posts WHERE idcategory = ?", id)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var posts []Post
	var idUser int
	var idCategory string
	var pathImg string
	var likes int
	var idpost int

	for rows.Next() {
		var post Post
		err := rows.Scan(&idpost, &post.Title, &post.Content, &idUser, &idCategory, &pathImg, &likes)
		if err != nil {
			return nil
		}
		idCategoryINT, _ := strconv.Atoi(idCategory)
		post.Category = GetCategory(idCategoryINT, db)
		post.User = GetUser(idUser, db)
		post.Image = pathImg
		post.Likes = likes
		post.Id = idpost
		if err != nil {
			return nil
		}
		post.Comment = CountComment(post.Id, db)
		posts = append(posts, post)
	}
	return posts
}

func GetLikes(id int, db *sql.DB) int {
	rows, err := db.Query("SELECT likes FROM posts WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var likes int
	for rows.Next() {
		err := rows.Scan(&likes)
		if err != nil {
			log.Fatal(err)
		}
	}
	return likes
}

func IsMine(idpost int, uuiduser string, db *sql.DB) bool {
	iduser := GetIDFromUUID(uuiduser, db)
	rows, err := db.Query("SELECT idUser FROM posts WHERE id = ?", idpost)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var idOwner int
	for rows.Next() {
		err := rows.Scan(&idOwner)
		if err != nil {
			log.Fatal(err)
		}
		if idOwner == iduser {
			return true
		}
	}
	return false
}
