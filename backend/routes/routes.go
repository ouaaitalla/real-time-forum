package routes

import (
	"net/http"

	"real-time-forum/backend/handlers/authHandlers"
	posthandlers "real-time-forum/backend/handlers/post-Handlers"
	"real-time-forum/backend/middleware"
)

func Routes() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/register", middleware.MethodeMiddleware("POST", authHandlers.RegisterHandler))
	http.HandleFunc("/login", middleware.MethodeMiddleware("POST", authHandlers.LoginHandler))
	// http.HandleFunc("/logout", middleware.MethodeMiddleware("POST", authHandlers.LogoutHandler))
	http.HandleFunc("/posts", middleware.MethodeMiddleware("GET", middleware.AuthMiddleware(posthandlers.GetPostsHandler)))
	http.HandleFunc("/creatPost", middleware.MethodeMiddleware("POST", middleware.AuthMiddleware(posthandlers.CreatPostHandler)))
}
