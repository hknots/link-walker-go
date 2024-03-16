package controller

import (
	"encoding/json"
	"link-walker-go/cache"
	"net/http"
)

func GetTasks(writer http.ResponseWriter, response *http.Request) {
    tasks := cache.GetAllTasks()
    tasksJSON, err := json.Marshal(tasks)
    if err != nil {
        http.Error(writer, "Failed to marshal tasks", http.StatusInternalServerError)
        return
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.Write(tasksJSON)
}

func PostTask(writer http.ResponseWriter, response *http.Request) {

}


