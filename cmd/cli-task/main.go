package main

import (
	//"cli-track/internal/domain"
	//"cli-track/internal/application/services"
	//"cli-track/internal/application/services"
	"cli-track/internal/infrastructure/storage"
	"fmt"
	"strconv"

	//"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Ваш список задач и целей",
	Long:  `Утилита для менеджмента задач с функциями добавления удаления и изменения задач`,
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
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	// rootCmd.AddCommand(updateCmd)
	// rootCmd.AddCommand(mark-in-progressCmd)
	// rootCmd.AddCommand(mark-doneCmd)
	// rootCmd.AddCommand(listDoneCmd)
	// rootCmd.AddCommand(listToDoCmd)
	// rootCmd.AddCommand(listInprogressCmd)
	// rootCmd.AddCommand(deleteCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [description] [status]",
	Short: "Добавляет новую задачу",
	Long:  "Первый аргумент - описание задачи, второй аргумент - один из вариантов статуса задачи: 'to-do', 'in-progress', 'done'",
	Args:  cobra.ExactArgs(2),
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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать все задачи",
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}
		taskManager.PrintTasks()
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Изменить описание задачи",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskManager, err := storage.LoadJson()
		if err != nil {
			fmt.Printf("Ошибка загрузки: %v\n", err)
			return
		}
		val, err := strconv.Atoi(args[0]) 
		if err != nil {
    		fmt.Println("Введенный id не является числом")
			return
		}
		err = taskManager.UpdateTask(val, args[1])
		if err != nil{
			fmt.Println(err)
		}
		err = storage.SaveToJson(taskManager)
		if err != nil {
			fmt.Printf("Ошибка сохранения: %v\n", err)
			return
		}
		fmt.Println("Задача успешно обновлена!")

	},
}
