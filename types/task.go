package types

import "github.com/google/uuid"


type Task struct {
    ID       string
	Status   string
    URL      string
    Requests int64 
	token    string
}

func NewTask(url string, token string) *Task {
	return &Task{
		ID:       uuid.NewString(),
		Status:   "FETCHING_RESOURCES",
		URL:      url,
		Requests: 0,
		token:    token,
	}
}

func (task *Task) GetToken() string {
	return task.token
}
