package services

import (
	"cli-track/internal/domain"
	"fmt"
	"time"
)

type TaskManager struct {
	Tasks  map[int]*domain.Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks:  make(map[int]*domain.Task),
		nextID: 1,
	}
}

func (tm *TaskManager) AddTask(description, status string) error {

	task := &domain.Task{}

	// Используем следующий доступный ID
	newID := tm.getNextAvailableID()
	task.SetID(newID)
	task.SetDescription(description)

	err := task.SetStatus(status)
	if err != nil {
		return fmt.Errorf("ошибка установки статуса: %w", err)
	}

	now := time.Now()
	task.SetCreatedAt(now)
	task.SetUpdatedAt(now)

	tm.Tasks[newID] = task
	tm.nextID = newID + 1

	return nil
}

// getNextAvailableID возвращает следующий доступный ID
func (tm *TaskManager) getNextAvailableID() int {
	if len(tm.Tasks) == 0 {
		return 1
	}

	// Проверяем, есть ли "дыры" в нумерации
	for i := 1; i <= len(tm.Tasks)+1; i++ {
		if _, exists := tm.Tasks[i]; !exists {
			return i
		}
	}

	// Если дыр нет, возвращаем следующий ID
	return len(tm.Tasks) + 1
}

func (tm *TaskManager) UpdateTask(id int, description string) error {
	if id < 0 {
		return domain.ErrWrongID
	}
	if _, exists := tm.Tasks[id]; exists {
	} else {
		return domain.ErrNotExistKey
	}
	task := tm.Tasks[id]
	task.SetDescription(description)

	now := time.Now()
	task.SetUpdatedAt(now)

	tm.Tasks[id] = task

	return nil
}

func (tm *TaskManager) DeleteTask(id int) error {
	if id < 0 {
		return domain.ErrWrongID
	}
	if _, exists := tm.Tasks[id]; !exists {
		return domain.ErrNotExistKey
	}

	delete(tm.Tasks, id)

	tm.reindexTasks()

	return nil
}

func (tm *TaskManager) reindexTasks() {
	if len(tm.Tasks) == 0 {
		tm.nextID = 1
		return
	}

	ids := sortingId(tm)

	newTasks := make(map[int]*domain.Task)

	for newID, oldID := range ids {
		task := tm.Tasks[oldID]
		task.SetID(newID + 1)
		newTasks[newID+1] = task
	}

	tm.Tasks = newTasks

	tm.nextID = len(newTasks) + 1
}

func (tm *TaskManager) Mark_in_progress(id int) error {
	if id < 0 {
		return domain.ErrWrongID
	}
	if _, exists := tm.Tasks[id]; exists {
	} else {
		return domain.ErrNotExistKey
	}
	task := tm.Tasks[id]
	task.SetStatus("in-progress")
	now := time.Now()
	task.SetUpdatedAt(now)

	tm.Tasks[id] = task

	return nil
}

func (tm *TaskManager) Mark_done(id int) error {
	if id < 0 {
		return domain.ErrWrongID
	}
	if _, exists := tm.Tasks[id]; exists {
	} else {
		return domain.ErrNotExistKey
	}
	task := tm.Tasks[id]
	task.SetStatus("done")
	now := time.Now()
	task.SetUpdatedAt(now)

	tm.Tasks[id] = task

	return nil
}

func (tm *TaskManager) MarkToDo(id int) error {
	if id < 0 {
		return domain.ErrWrongID
	}
	if _, exists := tm.Tasks[id]; exists {
	} else {
		return domain.ErrNotExistKey
	}
	task := tm.Tasks[id]
	task.SetStatus("to-do")
	now := time.Now()
	task.SetUpdatedAt(now)

	tm.Tasks[id] = task

	return nil
}

func (tm *TaskManager) PrintTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("Менеджер задач пустой")
		return
	}

	ids := sortingId(tm)

	for _, id := range ids {
		task := tm.Tasks[id]
		printAllAreas(task)
	}
}

func (tm *TaskManager) PrintTasksFilter(status string) {
	if len(tm.Tasks) == 0 {
		fmt.Println("Менеджер задач пустой")
		return
	}

	ids := sortingId(tm)

	for _, id := range ids {
		if tm.Tasks[id].GetStatus() == status {
			task := tm.Tasks[id]
			printAllAreas(task)
		}
	}
}

func (tm *TaskManager) GetNextID() int {
	return tm.nextID
}

func (tm *TaskManager) SetNextID(id int) {
	tm.nextID = id
}
