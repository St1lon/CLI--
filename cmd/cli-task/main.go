package main

import (
	//"cli-track/internal/domain"
	//"cli-track/internal/application/services"
	//"cli-track/internal/application/services"
	"cli-track/internal/domain"
	"cli-track/internal/infrastructure/storage"
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
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(mark_in_progressCmd)
	rootCmd.AddCommand(mark_doneCmd)
	rootCmd.AddCommand(listDoneCmd)
	rootCmd.AddCommand(listToDoCmd)
	rootCmd.AddCommand(listInprogressCmd)
	rootCmd.AddCommand(deleteCmd)
}

var addCmd = &cobra.Command{
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

var listCmd = &cobra.Command{
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

var listDoneCmd = &cobra.Command{
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
var listToDoCmd = &cobra.Command{
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

var listInprogressCmd = &cobra.Command{
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

var updateCmd = &cobra.Command{
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

var mark_in_progressCmd = &cobra.Command{
	Use:   "mark-in-progress <id>",
	Short: "Отметить задачу как выполняемую",
	Long: `Изменяет статус задачи на 'in-progress' (в процессе выполнения)

Аргументы:
  id  ID задачи для изменения статуса

Пример:
  tasker mark-in-progress 1  # Отмечает задачу №1 как выполняемую`,
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
		err = taskManager.Mark_in_progress(val)
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

var mark_doneCmd = &cobra.Command{
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

var deleteCmd = &cobra.Command{
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
