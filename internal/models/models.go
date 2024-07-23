package models

type HelloResponse struct {
	Message string `json:"message"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}

type Users []User

type Response struct {
	Message string `json:"message"`
}

type DatabaseType struct {
	Users Users `json:"users"`
}
