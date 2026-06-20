package store

import (
	"fmt"
	"sync"
)

type Task struct {
	ID          string
	Title       string
	Description *string
	Done        bool
}

var (
	mu sync.RWMutex

	Tasks = []*Task{
		{ID: "t_001", Title: "Первая задача", Description: strPtr("Учебный пример"), Done: false},
		{ID: "t_002", Title: "Вторая задача", Description: strPtr("GraphQL API"), Done: true},
	}
)

func strPtr(s string) *string {
	return &s
}

func All() []*Task {
	mu.RLock()
	defer mu.RUnlock()

	result := make([]*Task, len(Tasks))
	copy(result, Tasks)
	return result
}

func ByID(id string) (*Task, bool) {
	mu.RLock()
	defer mu.RUnlock()

	for _, t := range Tasks {
		if t.ID == id {
			return t, true
		}
	}
	return nil, false
}

func Create(title string, description *string) *Task {
	mu.Lock()
	defer mu.Unlock()

	task := &Task{
		ID:          fmt.Sprintf("t_%03d", len(Tasks)+1),
		Title:       title,
		Description: description,
		Done:        false,
	}
	Tasks = append(Tasks, task)
	return task
}

func Update(id string, title *string, description *string, done *bool) (*Task, bool) {
	mu.Lock()
	defer mu.Unlock()

	for _, t := range Tasks {
		if t.ID == id {
			if title != nil {
				t.Title = *title
			}
			if description != nil {
				t.Description = description
			}
			if done != nil {
				t.Done = *done
			}
			return t, true
		}
	}
	return nil, false
}

func Delete(id string) bool {
	mu.Lock()
	defer mu.Unlock()

	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true
		}
	}
	return false
}
