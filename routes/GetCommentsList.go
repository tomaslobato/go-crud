package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tomaslobato/go-crud/utils"
)

func GetCommentsList(w http.ResponseWriter, r *http.Request) {
	comments, err := utils.GetCommentsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(comments) //Comment type to json data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", string(jsonData))
}
