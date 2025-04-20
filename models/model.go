package controllers

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"name"`
	Password string
	UserType int `json:"type"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
