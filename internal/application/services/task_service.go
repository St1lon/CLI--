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
	task.SetID(tm.nextID)
	task.SetDescription(description)

	err := task.SetStatus(status)
	if err != nil {
		return fmt.Errorf("ошибка установки статуса: %w", err)
	}

	now := time.Now()
	task.SetCreatedAt(now)
	task.SetUpdatedAt(now)

	tm.Tasks[tm.nextID] = task
	tm.nextID++

	return nil
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
	if _, exists := tm.Tasks[id]; exists {
		delete(tm.Tasks, id)
	} else {
		return domain.ErrNotExistKey
	}
	return nil
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
