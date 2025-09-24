package services

import (
	"cli-track/internal/domain"
	"errors"
	"fmt"
	"sort"
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

func (tm *TaskManager) UpdateTask(id int, description string) error{
	if id < 0{
		return domain.ErrWrongID
	}
	task := tm.Tasks[id]
	task.SetDescription(description)

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

	// Собираем все ID и сортируем их
	var ids []int
	for id := range tm.Tasks {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	// Выводим задачи в отсортированном порядке
	for _, id := range ids {
		task := tm.Tasks[id]
		fmt.Println("Задача номер:", task.GetID())
		fmt.Println("Описание задачи:", task.GetDescription())
		fmt.Println("Статус задачи:", task.GetStatus())
		fmt.Println("Время создания задачи:", task.GetCreatedAt())
		fmt.Println("Время обновление задачи:", task.GetUpdatedAt()) // Исправлено: GetUpdatedAt вместо GetCreatedAt
		fmt.Println()
	}
}

// GetNextID возвращает следующий ID для новой задачи
func (tm *TaskManager) GetNextID() int {
	return tm.nextID
}

// SetNextID устанавливает следующий ID (для загрузки из JSON)
func (tm *TaskManager) SetNextID(id int) {
	tm.nextID = id
}
