package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tomaslobato/go-crud/models"
)

func GetCommentsHandler() ([]models.Comment, error) {
	data, err := ioutil.ReadFile("mock/comments.json")
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	err = json.Unmarshal(data, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
