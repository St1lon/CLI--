package main

import (
	//"cli-track/internal/domain"
	//"cli-track/internal/application/services"
	"cli-track/internal/infrastructure/storage"
	"fmt"
	//"fmt"
)

func main() {
	taskManager, _ := storage.LoadJson()
	err := taskManager.AddTask("vs sc sc scvsdcmvk", "to-do")
	if err != nil{
		fmt.Printf("Ошибка добавления %v\n", err)
		return
	}
	taskManager.PrintTasks()
	err = storage.SaveToJson(taskManager)
	if err != nil{
		fmt.Printf("Ошибка сохранения данных: %v\n", err)
		return
	}

}
