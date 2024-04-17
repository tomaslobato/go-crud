package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tomaslobato/go-crud/utils"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	//getting comments
	comments, err := utils.GetCommentsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var deletedComm string
	for i, comm := range comments {
		if fmt.Sprint(comm.ID) == id {
			deletedComm = fmt.Sprintf(`ID: %v, Content: "%s", User: %s`, comm.ID, comm.Content, comm.User)
			comments = append(comments[:i], comments[i+1:]...)
			break
		}
	}
	if deletedComm == "" {
		http.Error(w, "Comment not found", 404)
		return
	}

	updatedJson, err := json.Marshal(comments)
	err = os.WriteFile("mock/comments.json", updatedJson, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `Comment deleted! 
	"%s"`,
		deletedComm)
}
