package command

import (
	"cli-track/internal/domain"
	"cli-track/internal/infrastructure/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)


var MarkDoneCmd = &cobra.Command{
	Use:   "mark-done <id>",
	Short: "Отметить задачу как выполненную",
	Long: `Изменяет статус задачи на 'done' (выполнено)

Аргументы:
  id  ID задачи для изменения статуса

Пример:
  tasker mark-done 1  # Отмечает задачу №1 как выполненную`,
	Args: cobra.ExactArgs(1),
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
		err = taskManager.Mark_done(val)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = storage.SaveToJson(taskManager)
		if err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			return
		}
		fmt.Println("Статус задачи обновлен!")
	},
}