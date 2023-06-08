package models

import "os"

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
	File   os.File
}

type UserBook struct {
	ID     int
	UserID int
	BookID int
}
