package controllers

import (
	m "core_chat/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("test")
var tokenName = "token"

type Claims struct {
	UserName string `json:"username"`
	UserType int    `json:"type"`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, username string, userType int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	//create claims with user data
	claims := &Claims{
		UserName: username,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	//encrypt claim to jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//jwtKey := os.Getenv("JWT_TOKEN")
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	//set token to cookies
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(r, accessType)
		if !isValidToken {
			sendModifiedResponse(w, 400, "Unauthorized")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, username, userType := validateTokenFromCookies(r)
	fmt.Print(username, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.UserName, accessClaims.UserType
		}
	}
	return false, "", -1
}

func ProtectedContent(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	sendModifiedResponse(w, 200, "Login OK!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	username := r.FormValue("username")
	password := r.FormValue("password")

	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)

	var user m.User
	if err := row.Scan(&user.UserName, &user.Password, &user.UserType); err != nil {
		sendModifiedResponse(w, 400, "Error")
	} else {
		if CheckPasswordHash(user.Password, password) {
			generateToken(w, user.UserName, user.UserType)
			sendModifiedResponse(w, 200, "Success")
			return
		}
		sendModifiedResponse(w, 400, "Invalid username or password")
	}
}

func CheckPasswordHash(storedHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)

	sendModifiedResponse(w, 200, "Logout Success!")
}

func sendModifiedResponse(w http.ResponseWriter, stat int, msg string) {
	var response m.Response
	response.Status = stat
	response.Message = msg
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
