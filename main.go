// package main

// import (
// 	"core_chat/controllers"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// )

//	func main() {
//		router := mux.NewRouter()
//		router.HandleFunc("/test", controllers.Authenticate(controllers.ProtectedContent, 1)).Methods("GET")
//		router.HandleFunc("/login", controllers.Login).Methods("POST")
//		router.HandleFunc("/logout", controllers.Logout).Methods("POST")
//		http.Handle("/", router)
//		fmt.Println(controllers.HashPassword("test"))
//		fmt.Println("Connected to port 8888")
//		log.Println("Connected to port 8888")
//		log.Fatal(http.ListenAndServe(":8888", router))
//	}
package main

import (
	"core_chat/application/authentication/usecase"
	"core_chat/db/authentication"
	"core_chat/infra/mysql"
	"core_chat/infra/serviceimpl"
	"core_chat/web/rest"
	"core_chat/web/rest/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func configure() http.Handler {
	db := mysql.Connect()
	userRepo := authentication.NewUserRepository(db)
	tokenService := serviceimpl.NewJWTTokenService()
	hashService := serviceimpl.NewBcryptHashService()
	httpService := serviceimpl.NewHTTPService()
	loginUC := usecase.NewLoginUseCase(userRepo, tokenService, hashService, httpService)
	logoutUC := usecase.NewLogoutUseCase(httpService)
	handler := routes.NewAuthHandler(loginUC, logoutUC)

	r := mux.NewRouter()
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("POST")
	r.HandleFunc("/protected", routes.Authenticate(func(w http.ResponseWriter, r *http.Request) {
		rest.SendResponse(w, 200, "Login OK!")
	}, 1)).Methods("GET")

	return r
}

func main() {
	handler := configure()
	log.Println("Server running at :8888")
	log.Fatal(http.ListenAndServe(":8888", handler))
}
