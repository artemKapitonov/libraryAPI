package models

import (
	"mime/multipart"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	File   multipart.File
	Path   string `json:"path"`
}

type UserBook struct {
	ID     int
	UserID int
	BookID int
}
