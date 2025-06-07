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
	chatUC "core_chat/application/chat/usecase"
	notificationUC "core_chat/application/notification/usecase"
	personUC "core_chat/application/person/usecase"
	websocketApp "core_chat/application/websocket/service"
	"core_chat/db/authentication"
	"core_chat/db/chat"
	"core_chat/db/person"
	"core_chat/infra/mysql"
	"core_chat/infra/serviceimpl"
	"core_chat/web/rest"
	"core_chat/web/rest/routes"
	"core_chat/web/ws"
	"log"
	"net/http"
	"os"

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
	personValidatorService := serviceimpl.NewPersonValidatorService()
	personAntivirusService := serviceimpl.NewPersonAntivirusService()
	imageService := serviceimpl.NewImageService()
	emailAvailabilityUC := personUC.NewEmailAvailabilityUseCase(personPersonRepo, personValidatorService)
	identifierAvailabilityUC := personUC.NewIdentifierAvailabilityUseCase(personPersonRepo, personValidatorService)
	registerUC := personUC.NewRegisterPersonUseCase(personPersonRepo, hashService, personAntivirusService, personValidatorService, imageService)
	getProfileUC := personUC.NewGetProfileUseCase(personPersonRepo)
	defaultProfileImagePath := os.Getenv("DEFAULT_PROFILE_IMAGE")
	getProfileImageUC := personUC.NewGetProfileImageUseCase(personPersonRepo, imageService, defaultProfileImagePath)
	personHandler := routes.NewPersonHandler(emailAvailabilityUC, identifierAvailabilityUC, registerUC, getProfileUC, getProfileImageUC)

	chatRepo := chat.NewChatRepository(db)
	chatValidatorService := serviceimpl.NewChatValidatorService()
	chatAntivirusService := serviceimpl.NewChatAntivirusService()
	getChatImageUC := chatUC.NewGetChatImageUseCase(chatRepo)
	uploadChatImageUC := chatUC.NewUploadChatImageUseCase(chatAntivirusService, chatRepo, chatValidatorService)
	chatHandler := routes.NewChatImageHandler(uploadChatImageUC, getChatImageUC)

	var wsInitRouter websocketApp.WebSocketRouter
	serviceimpl.InitWebSocketManagerImpl(wsInitRouter)
	chatWsManager := serviceimpl.NewChatWebSocketManager(wsInitRouter)
	firebaseCredPath := os.Getenv("FIREBASE_CREDENTIAL_PATH")
	pushNotifierService, _ := serviceimpl.NewFirebasePushNotifier(firebaseCredPath)
	directMessageService := serviceimpl.NewDirectMessageServiceImpl()
	msgNotifierService := serviceimpl.NewMessageNotifierServiceImpl()
	sendMessageUC := chatUC.NewSendMessageUseCase(chatRepo, chatWsManager, pushNotifierService, directMessageService, msgNotifierService)
	sendImageUC := chatUC.NewSendImageUseCase(chatRepo, chatWsManager, pushNotifierService, directMessageService, msgNotifierService)
	joinChatRoomUC := chatUC.NewJoinChatRoomUseCase(chatRepo, chatWsManager)
	leaveChatRoomUC := chatUC.NewLeaveChatRoomUseCase(chatRepo, chatWsManager)

	notificationWsManager := serviceimpl.NewNotificationWebSocketManager(wsInitRouter)
	notifierService := serviceimpl.NewNotifierServiceImpl()
	serverResponseUC := notificationUC.NewServerResponseUseCase(notificationWsManager, notifierService)

	wsRouter := ws.NewDefaultRouter(sendMessageUC, sendImageUC, joinChatRoomUC, leaveChatRoomUC, serverResponseUC)
	wsManager := serviceimpl.NewWebSocketManager(wsInitRouter)
	wsManager.SetRouter(wsRouter)
	wsHandler := ws.NewWebSocketHandler(wsManager)

	r := mux.NewRouter()
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/auth/token/refresh", routes.Authenticate(authHandler.RefreshToken, "user", false)).Methods("POST")

	r.HandleFunc("/person/identifier", personHandler.CheckIdentifierAvailability).Methods("GET")
	r.HandleFunc("/person/email", personHandler.CheckEmailAvailability).Methods("GET")
	r.HandleFunc("/person", personHandler.Register).Methods("POST")
	r.HandleFunc("/person/profile", routes.Authenticate(personHandler.GetProfile, "user", false)).Methods("GET")
	r.HandleFunc("/person/profile/image", routes.Authenticate(personHandler.GetProfileImage, "user", false)).Methods("GET")

	r.HandleFunc("/chat/image", routes.Authenticate(chatHandler.UploadChatImage, "user", true)).Methods("POST")
	r.HandleFunc("/chat/image", routes.Authenticate(chatHandler.GetChatImage, "user", true)).Methods("GET")

	r.HandleFunc("/ws", routes.Authenticate(wsHandler.HandleWebSocketConn, "user", true)).Methods("GET")

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
