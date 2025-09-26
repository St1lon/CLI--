package command

import (
	"cli-track/internal/domain"
	"cli-track/internal/infrastructure/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update <id> <новое_описание>",
	Short: "Изменить описание задачи",
	Long: `Обновляет описание существующей задачи

Аргументы:
  id               ID задачи для обновления
  новое_описание   Новое описание задачи

Пример:
  tasker update 1 "Изучить Go более углубленно"
  tasker update 2 "Купить молоко и хлеб"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Println(err)
			return
		}
		val, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(domain.ErrWrongID)
			return
		}
		err = taskManager.UpdateTask(val, args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = storage.SaveToJson(taskManager)
		if err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			return
		}
		fmt.Println("Задача успешно обновлена!")

	},
}
