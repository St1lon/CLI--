package main

import (
	//"cli-track/internal/domain"
	//"cli-track/internal/application/services"
	//"cli-track/internal/application/services"
	"cli-track/internal/domain"
	"cli-track/internal/infrastructure/storage"
	"cli-track/cmd/cli-task/command"
	"fmt"
	"strconv"

	//"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "CLI утилита для управления задачами",
	Long: `Tasker CLI - Простая и мощная утилита для управления задачами

Возможности:
  • Добавление задач с различными статусами
  • Просмотр задач с фильтрацией по статусам
  • Обновление описаний и статусов задач
  • Автоматическое сохранение в JSON файл

Примеры использования:
  tasker add "Изучить Go" "to-do"          # Добавить новую задачу
  tasker list                              # Показать все задачи
  tasker list-todo                         # Показать только задачи к выполнению
  tasker update 1 "Изучить Go углубленно"  # Обновить описание задачи`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	Execute()
}

func init() {
	rootCmd.AddCommand(command.AddCmd)
	rootCmd.AddCommand(command.ListCmd)
	rootCmd.AddCommand(command.UpdateCmd)
	rootCmd.AddCommand(command.MarkInProgressCmd)
	rootCmd.AddCommand(command.MarkDoneCmd)
	rootCmd.AddCommand(command.ListDoneCmd)
	rootCmd.AddCommand(command.ListToDoCmd)
	rootCmd.AddCommand(command.ListInprogressCmd)
	rootCmd.AddCommand(command.DeleteCmd)
}
