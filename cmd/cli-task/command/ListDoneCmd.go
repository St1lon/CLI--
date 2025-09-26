package command

import (
	"cli-track/internal/infrastructure/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var ListDoneCmd = &cobra.Command{
	Use:   "list-done",
	Short: "Показать сделанные задачи",
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}
		taskManager.PrintTasksFilter("done")
	},
}
