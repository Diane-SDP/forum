package models

type Data struct {
	Posts      []Post
	Categories []Category
	Users []User
	CurrentUser User
	IsConnected bool
}

type DataProfil struct {
	User       User
	Posts []Post
	IsMe bool
	IsConnected bool
	CurrentUser User
}

type DataNotFound struct {
	IsConnected bool
	CurrentUser User
}