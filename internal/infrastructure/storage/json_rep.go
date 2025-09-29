package storage

import (
	"cli-track/internal/application/services"
	"cli-track/internal/domain"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func SaveToJson(tm *services.TaskManager) error {
	var tasks []struct {
		ID          int       `json:"id"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	for _, task := range tm.Tasks {
		temp := struct {
			ID          int       `json:"id"`
			Description string    `json:"description"`
			Status      string    `json:"status"`
			CreatedAt   time.Time `json:"createdAt"`
			UpdatedAt   time.Time `json:"updatedAt"`
		}{
			ID:          task.GetID(),
			Description: task.GetDescription(),
			Status:      task.GetStatus(),
			CreatedAt:   task.GetCreatedAt(),
			UpdatedAt:   task.GetUpdatedAt(),
		}
		tasks = append(tasks, temp)
	}

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %w", err)
	}

	filename := "tasks.json"
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	return nil
}

func LoadJson() (*services.TaskManager, error) {
	filename := "tasks.json"
	tm := services.NewTaskManager()

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Файл tasks.json не найден, начинаем с пустого списка задач")
		return tm, nil
	}

	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %w", filename, err)
	}

	var tasks []struct {
		ID          int       `json:"id"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	tm.Tasks = make(map[int]*domain.Task)
	maxID := 0

	for _, taskData := range tasks {
		task := &domain.Task{}
		task.SetID(taskData.ID)
		task.SetDescription(taskData.Description)

		err := task.SetStatus(taskData.Status)
		if err != nil {
			return nil, fmt.Errorf("ошибка установки статуса для задачи %d: %w", taskData.ID, err)
		}

		task.SetCreatedAt(taskData.CreatedAt)
		task.SetUpdatedAt(taskData.UpdatedAt)

		tm.Tasks[taskData.ID] = task

		if taskData.ID > maxID {
			maxID = taskData.ID
		}
	}

	tm.SetNextID(maxID + 1)
	return tm, nil
}
