package command

import (
	"cli-track/internal/infrastructure/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Удалить задачу",
	Long:  "Удаляет задачу по её идентификатору (ID)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return
		}

		err = taskManager.DeleteTask(id)
		if err != nil {
			fmt.Printf("Ошибка удаления: %v\n", err)
			return
		}

		err = storage.SaveToJson(taskManager)
		if err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			return
		}

		fmt.Printf("Задача #%d успешно удалена!\n", id)
	},
}
