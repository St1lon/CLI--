# Tasker CLI - Утилита управления задачами

Простая и эффективная CLI утилита для управления задачами, написанная на Go.

## 🚀 Быстрая установка

### Автоматическая установка (рекомендуется)

```powershell
# Клонируйте репозиторий
git clone https://github.com/St1lon/CLI--

# Перейдите в директорию
cd CLI--

# Запустите скрипт установки
.\install.ps1
```

### Ручная установка

```powershell
# Соберите приложение
go build -ldflags="-w -s" -o tasker.exe cmd/cli-task/main.go

# Скопируйте в PATH
copy tasker.exe C:\Windows\System32\
```

## 📋 Использование

### Основные команды

```bash
# Добавить новую задачу
tasker add "Изучить Go" "to-do"

# Показать все задачи
tasker list

# Показать задачи по статусам
tasker list-todo         # Задачи к выполнению
tasker list-in-progress  # Задачи в процессе
tasker list-done         # Выполненные задачи

# Изменить описание задачи
tasker update 1 "Новое описание"

# Изменить статус задачи
tasker mark-in-progress 1
tasker mark-done 1

# Справка
tasker --help
tasker add --help
```

### Доступные статусы

- `to-do` - к выполнению
- `in-progress` - в процессе  
- `done` - выполнено

## 🔧 Возможности

- ✅ **Добавление задач** с описанием и статусом
- ✅ **Просмотр всех задач** с сортировкой по ID
- ✅ **Фильтрация по статусам** - отдельные команды для разных статусов
- ✅ **Обновление описания задач** по ID
- ✅ **Изменение статуса задач** - mark-in-progress, mark-done
- ✅ **Персистентное хранение** в JSON файле
- ✅ **Валидация статусов** с информативными ошибками
- ✅ **Автоматическое создание файла** при первом запуске
- ✅ **Кроссплатформенность** Windows/Linux/macOS
- ✅ **Автодополнение команд** для PowerShell

## 📁 Структура проекта

```
CLI--/
├── cmd/cli-task/main.go        # Точка входа и CLI команды
├── internal/
│   ├── domain/
│   │   ├── task.go             # Модель задачи
│   │   └── errors.go           # Ошибки домена
│   ├── application/services/
│   │   └── task_service.go     # Бизнес-логика
│   └── infrastructure/storage/
│       └── json_rep.go         # JSON хранилище
├── tasker.exe                  # Исполняемый файл
├── tasks.json                  # Файл данных
├── install.ps1                 # Скрипт установки
└── completion.ps1              # Автодополнение
```

## 🎯 Примеры использования

```bash
# Добавление различных задач
tasker add "Купить продукты" "to-do"
tasker add "Написать отчет" "in-progress"
tasker add "Закончить проект" "done"

# Просмотр задач
tasker list                    # Все задачи
tasker list-todo              # Только задачи к выполнению
tasker list-in-progress       # Только задачи в процессе  
tasker list-done              # Только выполненные задачи

# Управление задачами
tasker update 1 "Новое описание задачи"  # Изменить описание
tasker mark-in-progress 1                # Перевести в процесс
tasker mark-done 2                       # Пометить как выполненную
```

## ⚙️ Настройка автодополнения

Для включения автодополнения в PowerShell:

```powershell
# Добавьте в ваш PowerShell профиль
. .\completion.ps1
```

## 🗂️ Местоположение данных

- **После установки**: `%USERPROFILE%\.local\bin\tasks.json`
- **При разработке**: `.\tasks.json` (в директории проекта)
- **Файл создается автоматически** при первом добавлении задачи

## 🛠️ Разработка

### Требования

- Go 1.19+
- PowerShell (для скриптов установки)

### Команды разработчика

```bash
# Запуск в режиме разработки
go run cmd/cli-task/main.go add "Test task" "to-do"
go run cmd/cli-task/main.go list

# Сборка оптимизированного исполняемого файла
go build -ldflags="-w -s" -o tasker.exe cmd/cli-task/main.go

# Установка в систему
.\install.ps1

# Тесты
go test ./...

# Проверка доступных команд
.\tasker.exe --help
```

## 📄 Лицензия

MIT License

## 🤝 Вклад в проект

1. Fork репозитория
2. Создайте feature branch (`git checkout -b feature/amazing-feature`)
3. Commit изменения (`git commit -m 'Add some amazing feature'`)
4. Push в branch (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

---

**Tasker CLI** - делайте больше, управляйте задачами проще! 🎯