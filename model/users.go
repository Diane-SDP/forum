package models

import (
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	Id       int
	Pseudo   string
	Bio      string
	Mail     string
	Image    string
	Password string
	IsGoogle bool
}

func GetUser(id int) User {
	rows, err := DB.Query("SELECT username, password, bio, email, image FROM users WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Pseudo, &user.Password, &user.Bio, &user.Mail, &user.Image)
		if err != nil {
			panic(err)
		}
	}
	user.Id = id
	return user
}

func GetIDFromUUID(uuid string) int {
	rows, err := DB.Query("SELECT id FROM users WHERE uuid = ?", uuid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}
	}
	return id
}

func CreateUser(username string, email string, uuid string) {
	_, err := DB.Exec("INSERT INTO users (username, email, uuid, googleAccount) VALUES (?, ?, ?, TRUE)", username, email, uuid)
	if err != nil {
		panic(err)
		return
	}
}

func Exist(email string) {

}

func GenerateUUID() string {
	const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, 12)
	for i := range result {
		result[i] = char[random.Intn(len(char))]
	}
	uuid := string(result)
	rows, err := DB.Query("SELECT uuid FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var usedUUID string
	for rows.Next() {
		err := rows.Scan(&usedUUID)
		if err != nil {
			panic(err)
		}
		if usedUUID == uuid {
			return GenerateUUID()
		}
	}
	return uuid
}

func UserCheck(username string, mail string) (bool, string) {
	rows, err := DB.Query("SELECT username, email, uuid FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var usedName string
	var usedMail string
	for rows.Next() {
		err := rows.Scan(&usedName, &usedMail)
		if err != nil {
			panic(err)
		}
		if usedName == username {
			return true, "Username déjà utilisé"
		} else if usedMail == mail {
			return true, "Mail déjà utilisé"
		}
	}
	return false, ""
}

func GoogleUserCheck(mail string) (bool, string) {
	rows, err := DB.Query("SELECT uuid FROM users WHERE email = ?", mail)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var uuid string
	for rows.Next() {
		err := rows.Scan(&uuid)
		if err != nil {
			panic(err)
		}
	}
	if uuid == "" {
		return false, uuid
	}
	return true, uuid
}

func CreateUsername() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(999999)
	username := "User" + strconv.Itoa(randomNumber)
	return username
}
