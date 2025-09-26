package command


import (
	"cli-track/internal/infrastructure/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать все задачи",
	Long: `Показывает все задачи из вашего списка

Задачи отображаются с подробной информацией:
  • ID задачи
  • Описание
  • Текущий статус
  • Время создания и последнего обновления

Для просмотра задач по конкретному статусу используйте:
  tasker list-todo      - только задачи к выполнению
  tasker list-progress  - только задачи в процессе
  tasker list-done      - только выполненные задачи`,
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}
		taskManager.PrintTasks()
	},
}