package router

import (
	"github.com/gorilla/mux"
	"github.com/messaging-go-service/internal/controller"
	"github.com/messaging-go-service/internal/repository"
	"github.com/messaging-go-service/middleware"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Init repositories
	userRepo := repository.NewUserRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)
	conversationRepo := repository.NewConversationRepository(db)

	// Init controllers
	authController := controller.NewAuthController(userRepo)
	userController := controller.NewUserController(userRepo)
	notificationController := controller.NewNotificationController(notificationRepo)
	conversationController := controller.NewConversationController(conversationRepo, userRepo)

	// Init routers
	authRouter := router.PathPrefix("/api/auth").Subrouter()
	authRouter.HandleFunc("/login", authController.Login).Methods("POST")
	authRouter.HandleFunc("/register", authController.Register).Methods("POST")
	authRouter.HandleFunc("/resetPassword/{id}", authController.ResetPassword).Methods("POST")

	userRouter := router.PathPrefix("/api/user").Subrouter()
	userRouter.Use(middleware.CheckAuth)
	userRouter.HandleFunc("", userController.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", userController.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/search", userController.SearchUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", userController.GetUserDetail).Methods("GET")

	notificationRouter := router.PathPrefix("/api/notification").Subrouter()
	notificationRouter.HandleFunc("/list/{user_id}", notificationController.GetNotificationsByUser).Methods("GET")

	conversationRouter := router.PathPrefix("/api/conversation").Subrouter()
	conversationRouter.Use(middleware.CheckAuth)
	conversationRouter.HandleFunc("", conversationController.AddConversation).Methods("POST")
	conversationRouter.HandleFunc("/{id}", conversationController.DeleteConversation).Methods("DELETE")
	conversationRouter.HandleFunc("/{id}", conversationController.GetConversationDetail).Methods("GET")
	conversationRouter.HandleFunc("/participant", conversationController.AddParticipant).Methods("POST")
	conversationRouter.HandleFunc("/message", conversationController.AddMessage).Methods("POST")
	conversationRouter.HandleFunc("/message/{conversation_id}", conversationController.RetrieveMessages).Methods("GET")
	conversationRouter.HandleFunc("/all/{user_id}", conversationController.GetConversationsByUserID).Methods("GET")

	return router
}
