package routes

import (
	"net/http"
	"real-time-forum/backend/handlers/authHandlers"
	"real-time-forum/backend/middleware"
)

func Routes() {
	http.HandleFunc("/register", middleware.MethodeMiddleware("POST", authHandlers.RegisterHandler))
	// http.HandleFunc("/login", middleware.MethodeMiddleware("GET", middleware.()))
}
