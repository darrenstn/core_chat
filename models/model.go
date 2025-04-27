package controllers

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"name"`
	Password string
	UserType int `json:"type"`
}

type Claims struct {
	UserName string `json:"username"`
	UserType int    `json:"type"`
	jwt.StandardClaims
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
