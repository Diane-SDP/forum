package main

import (
	"database/sql"
	controllers "forum/controller"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        password TEXT,
		bio TEXT,
		image TEXT DEFAULT './Assets/img/user.png',
        email TEXT,
		uuid TEXT
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
		image TEXT,
		likes INTEGER NOT NULL,
		idUser INTEGER,
		idCategory INTEGER,
        FOREIGN KEY (idUser) REFERENCES users(id),
		FOREIGN KEY (idCategory) REFERENCES categories(id)
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
		image TEXT DEFAULT './Assets/img/categories.png',
        description TEXT
    )`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS likes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        idUser INTEGER,
        idPost INTEGER,
		uuidUser TEXT,
		FOREIGN KEY (idUser) REFERENCES user(id),
		FOREIGN KEY (idPost) REFERENCES post(id)
    )`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        idUser INTEGER,
        idPost INTEGER,
		commentaire TEXT
    )`)

	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("Assets"))
	http.Handle("/Assets/", http.StripPrefix("/Assets", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.HomeHandler(w, r, db)
	})

	http.HandleFunc("/postcategory", func(w http.ResponseWriter, r *http.Request) {
		controllers.CategoryHandler(w, r, db)
	})

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreatePostHandler(w, r, db)
	})
	http.HandleFunc("/like/", func(w http.ResponseWriter, r *http.Request) {
		controllers.LikedPostHandler(w, r, db)
	})
	http.HandleFunc("/loginsuccess", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoginSuccessHandler(w, r, db)
	})
	http.HandleFunc("/deco", func(w http.ResponseWriter, r *http.Request) {
		controllers.DisconnectHandler(w, r, db)
	})
	http.HandleFunc("/topic/", func(w http.ResponseWriter, r *http.Request) {
		controllers.TopicHandler(w, r, db)
	})
	http.HandleFunc("/edit/", func(w http.ResponseWriter, r *http.Request) {
		controllers.EditHandler(w, r, db)
	})
	http.HandleFunc("/createcategory", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateCategoryHandler(w, r, db)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		controllers.PostHandler(w, r, db)
	})
	http.HandleFunc("/connection", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConnectionHandler(w, r, db)
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		controllers.SignupHandler(w, r, db)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.LoginHandler(w, r, db)
	})
	http.HandleFunc("/profil/", func(w http.ResponseWriter, r *http.Request) {
		controllers.ProfilHandler(w, r, db)
	})
	http.HandleFunc("/changeprofil", func(w http.ResponseWriter, r *http.Request) {
		controllers.ChangeProfilHandler(w, r, db)
	})
	http.HandleFunc("/commentaire/", func(w http.ResponseWriter, r *http.Request) {
		controllers.CommentaireHandler(w, r, db)
	})
	http.HandleFunc("/createcommentaire/", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateCommentHandler(w, r, db)
	})
	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		controllers.SaveHandler(w, r, db)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
