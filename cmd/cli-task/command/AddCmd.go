package command

import (
	"cli-track/internal/infrastructure/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add <описание> <статус>",
	Short: "Добавить новую задачу",
	Long: `Добавляет новую задачу в список

Аргументы:
  описание  Описание задачи (обязательно)
  статус    Статус задачи (обязательно)

Доступные статусы:
  • to-do        - задача к выполнению
  • in-progress  - задача в процессе выполнения  
  • done         - выполненная задача

Примеры:
  tasker add "Купить молоко" "to-do"
  tasker add "Написать отчет" "in-progress"
  tasker add "Изучить Go" "done"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}

		err = taskManager.AddTask(args[0], args[1])
		if err != nil {
			fmt.Printf("Ошибка добавления задачи: %v\n", err)
			return
		}

		err = storage.SaveToJson(taskManager)
		if err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			return
		}

		fmt.Println("Задача успешно добавлена!")
	},
}
