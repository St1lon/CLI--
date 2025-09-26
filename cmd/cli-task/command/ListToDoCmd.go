package command

import (
	"cli-track/internal/infrastructure/storage"
	"fmt"

	"github.com/spf13/cobra"
)



var ListToDoCmd = &cobra.Command{
	Use:   "list-to-do",
	Short: "Показать задачи которые нужно сделать",
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}
		taskManager.PrintTasksFilter("to-do")
	},
}