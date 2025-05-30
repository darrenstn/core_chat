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
	authUC "core_chat/application/authentication/usecase"
	personUC "core_chat/application/person/usecase"
	"core_chat/db/authentication"
	"core_chat/db/person"
	"core_chat/infra/mysql"
	"core_chat/infra/serviceimpl"
	"core_chat/web/rest"
	"core_chat/web/rest/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func configure() http.Handler {
	db := mysql.Connect()

	authPersonRepo := authentication.NewPersonRepository(db)
	tokenService := serviceimpl.NewJWTTokenService()
	hashService := serviceimpl.NewBcryptHashService()
	loginUC := authUC.NewLoginUseCase(authPersonRepo, tokenService, hashService)
	refreshUC := authUC.NewRefreshTokenUseCase(authPersonRepo, tokenService)
	authHandler := routes.NewAuthHandler(loginUC, refreshUC)

	personPersonRepo := person.NewPersonRepository(db)
	validatorService := serviceimpl.NewValidatorService()
	antivirusService := serviceimpl.NewAntivirusService()
	imageService := serviceimpl.NewImageService()
	emailAvailabilityUC := personUC.NewEmailAvailabilityUseCase(personPersonRepo, validatorService)
	identifierAvailabilityUC := personUC.NewIdentifierAvailabilityUseCase(personPersonRepo, validatorService)
	registerUC := personUC.NewRegisterPersonUseCase(personPersonRepo, hashService, antivirusService, validatorService, imageService)
	personHandler := routes.NewPersonHandler(emailAvailabilityUC, identifierAvailabilityUC, registerUC)

	r := mux.NewRouter()
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/auth/token/refresh", routes.Authenticate(authHandler.RefreshToken, "user", false)).Methods("POST")

	r.HandleFunc("/person/identifier", personHandler.CheckIdentifierAvailability).Methods("GET")
	r.HandleFunc("/person/email", personHandler.CheckEmailAvailability).Methods("GET")
	r.HandleFunc("/person", personHandler.Register).Methods("POST")

	r.HandleFunc("/protected/email", routes.Authenticate(func(w http.ResponseWriter, r *http.Request) {
		rest.SendResponse(w, 200, "Login and Email OK!")
	}, "user", true)).Methods("GET")
	r.HandleFunc("/protected", routes.Authenticate(func(w http.ResponseWriter, r *http.Request) {
		rest.SendResponse(w, 200, "Login OK!")
	}, "user", false)).Methods("GET")

	return r
}

func main() {
	loadEnv()
	handler := configure()
	log.Println("Server running at :8888")
	log.Fatal(http.ListenAndServe(":8888", handler))
}
