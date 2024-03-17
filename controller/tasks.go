package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"link-walker-go/cache"
	"link-walker-go/types"
	"net/http"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	tasks := cache.GetAllTasks()
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		http.Error(writer, "Failed to marshal tasks", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(tasksJSON)
}

func PostTask(writer http.ResponseWriter, request *http.Request) {
	token, err := GetAuthorizationToken(request)
	if err != nil {
		http.Error(writer, "Authorization must have Bearer <token>", http.StatusUnauthorized)
		return
	}

	url, err := GetUrlFromBody(request)
	if err != nil {
		http.Error(writer, "url required in body", http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	task := types.NewTask(id, url, token)
	cache.AddTask(task)
	go StartTask(task)
}
