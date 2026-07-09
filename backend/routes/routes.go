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
	http.HandleFunc("/session",middleware.MethodeMiddleware("GET", authHandlers.CheckSessionHandler))
	http.HandleFunc("/logout", middleware.MethodeMiddleware("GET", middleware.AuthMiddleware(authHandlers.LogoutHandler)))
	http.HandleFunc("/posts", middleware.MethodeMiddleware("GET", middleware.AuthMiddleware(posthandlers.GetPostsHandler)))
	http.HandleFunc("/post/{id}", middleware.MethodeMiddleware("GET", middleware.AuthMiddleware(posthandlers.GetPostHandler)))
	http.HandleFunc("/creatPost", middleware.MethodeMiddleware("POST", middleware.AuthMiddleware(posthandlers.CreatPostHandler)))
}
