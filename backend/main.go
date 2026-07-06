package main

import (
	"fmt"
	"log"
	"net/http"

	"real-time-forum/backend/routes"
	"real-time-forum/database"
	// "net/http"
)

func main() {
	err := database.InitDb()
	if err != nil {
		// error
		fmt.Println(err)
		return
	}
	routes.Routes()
	fmt.Println("http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ERROR STARTING SERVER", err)
	}
}
