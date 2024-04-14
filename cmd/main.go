package main

import (
	"fmt"
	"net/http"

	"github.com/tomaslobato/go-api/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", routes.GetCommentsList)
	mux.HandleFunc("POST /comment", routes.PostComment)
	mux.HandleFunc("GET /comment/{id}", routes.GetComment)

	fmt.Println("Server running on http://localhost:8000")

	err := http.ListenAndServe("localhost:8000", mux)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
