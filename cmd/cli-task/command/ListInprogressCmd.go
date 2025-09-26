package command

import (
	"cli-track/internal/infrastructure/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var ListInprogressCmd = &cobra.Command{
	Use:   "list-in-progress",
	Short: "Показать задачи которые в прогрессе",
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}
		taskManager.PrintTasksFilter("in-progress")
	},
}
