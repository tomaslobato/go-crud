package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tomaslobato/go-crud/models"
	"github.com/tomaslobato/go-crud/utils"
)

func PostComment(w http.ResponseWriter, r *http.Request) {
	//getting request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//getting comments
	comments, err := utils.GetCommentsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//creating the new comment and appending it
	var newComm models.Comment
	err = json.Unmarshal(body, &newComm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the request body contains the required fields
	if newComm.Content == "" || newComm.User == "" {
		http.Error(w, "Missing required fields 'Content' or 'User'", 422)
		return
	}

	newComm.ID = len(comments) + 1
	comments = append(comments, newComm)

	updatedJson, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//writing the updated comments in the mocked file
	err = os.WriteFile("mock/comments.json", updatedJson, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Comment Posted!\n %s", string(updatedJson))
}
