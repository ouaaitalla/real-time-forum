package main

import (
	"log"
	"net/http"
	"real-time-forum/backend/db"
	posthandlers "real-time-forum/backend/handlers/post-Handlers"
)

func main() {
	if err := db.Init("database/forum.db"); err != nil {
		log.Fatalf("database init failed: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/posts", posthandlers.PostsHandler)
	mux.Handle("/", http.FileServer(http.Dir("frontend")))
	log.Println("real-time forum running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}
