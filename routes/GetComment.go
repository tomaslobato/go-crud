package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tomaslobato/go-api/models"
	"github.com/tomaslobato/go-api/utils"
)

func GetComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	comments, err := utils.GetCommentsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundComment *models.Comment //it's a pointer so it can be nil
	for _, comm := range comments {
		if fmt.Sprint(comm.ID) == id {
			foundComment = &comm
			break
		}
	}

	if foundComment == nil {
		fmt.Fprintf(w, "Comment not found")
		return
	}

	jsonData, err := json.Marshal(foundComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", string(jsonData))
}
