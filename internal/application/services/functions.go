package services

import (
	"cli-track/internal/domain"
	"fmt"
	"sort"
)

func printAllAreas(task *domain.Task) {
	fmt.Println("Задача номер:", task.GetID())
	fmt.Println("Описание задачи:", task.GetDescription())
	fmt.Println("Статус задачи:", task.GetStatus())
	fmt.Println("Время создания задачи:", task.GetCreatedAt())
	fmt.Println("Время обновление задачи:", task.GetUpdatedAt())
	fmt.Println()
}

func sortingId(tm *TaskManager) []int {
	var ids []int
	for id := range tm.Tasks {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	return ids
}
