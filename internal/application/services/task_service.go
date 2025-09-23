package services

import (
	"cli-track/internal/domain"
	"fmt"
	"time"
)

type TaskManager struct {
	tasks  map[int]*domain.Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]*domain.Task),
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

	tm.tasks[tm.nextID] = task
	tm.nextID++

	return nil
}

func (tm *TaskManager) PrintTasks(){
	if len(tm.tasks) == 0{
		fmt.Println("Менеджер задач пустой")
		return
	}
	for _, task := range tm.tasks {
	if task == nil {
		fmt.Println("Found nil task")
		continue
	}
		fmt.Println("Задача номер:", task.GetID())
		fmt.Println("Описание задачи:", task.GetDescription())
		fmt.Println("Статус задачи:", task.GetStatus())
		fmt.Println("Время создания задачи:", task.GetCreatedAt())
		fmt.Println("Время обновление задачи:", task.GetCreatedAt())
	}
}
 