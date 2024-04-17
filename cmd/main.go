package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomaslobato/go-crud/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", routes.GetCommentsList)
	mux.HandleFunc("POST /comment", routes.PostComment)
	mux.HandleFunc("GET /comment/{id}", routes.GetComment)
	mux.HandleFunc("DELETE /comment/{id}", routes.DeleteComment)
	mux.HandleFunc("PUT /comment/{id}", routes.EditComment)

	fmt.Println("Server running on http://localhost:8000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe("0.0.0.0:"+port, mux)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
