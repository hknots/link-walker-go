package cache

import (
	"link-walker-go/types"
	"sync"
)

var (
	tasks map[string]*types.Task = make(map[string]*types.Task)
	mu    sync.RWMutex
)

func addTask(task *types.Task) {
	mu.Lock()
	defer mu.Unlock()
	tasks[task.ID] = task
}

func GetTask(id string) (*types.Task, bool) {
	mu.RLock()
	defer mu.RUnlock()
	task, exists := tasks[id]
	return task, exists
}

func GetAllTasks() []*types.Task {
	mu.RLock()
	defer mu.RUnlock()
	allTasks := make([]*types.Task, 0, len(tasks))
	for _, task := range tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

