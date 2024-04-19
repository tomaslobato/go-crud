package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tomaslobato/go-crud/utils"
)

type Body struct {
	Content string
	User    string
}

func EditComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	//getting request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newComment Body
	err = json.Unmarshal(body, &newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := utils.GetCommentsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	found := false
	sameContent := false
	for i, comm := range comments {
		if fmt.Sprint(comm.ID) == id {
			if comm.Content == newComment.Content {
				sameContent = true
				break
			}
			comments[i].Content = newComment.Content
			comments[i].User = newComment.User
			found = true
			break
		}
	}
	if sameContent {
		http.Error(w, "Same content, input something different", 422)
		return
	}
	if !found {
		http.Error(w, "Comment not found", 404)
		return
	}

	updatedJson, err := json.Marshal(comments)
	err = os.WriteFile("mock/comments.json", updatedJson, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Comment edited \n Content: %s, User: %s \n", newComment.Content, newComment.User)
}
