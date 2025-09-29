# 🚀 Tasker CLI - Утилита управления задачами

Мощная и удобная кроссплатформенная CLI утилита для управления задачами, написанная на Go с использованием Cobra framework.

## ⚡ Быстрая установка

### 🪟 Windows (PowerShell)

```powershell
# Клонируйте репозиторий
git clone https://github.com/St1lon/CLI--
cd CLI--

# Соберите проект
./build.ps1

# Автоматическая установка
.\install.ps1

# Глобальная установка (требует прав администратора)
.\install.ps1 -Global
```

### 🐧 Linux / 🍎 macOS (Bash)

```bash
# Клонируйте репозиторий
git clone https://github.com/St1lon/CLI--
cd CLI--

# Сделайте скрипт исполняемым и установите
chmod +x install.sh
./install.sh

# Соберите проект
./build.sh

# Глобальная установка (требует sudo)
./install.sh --global
```

### 🛠️ Использование готовых скриптов сборки

#### Windows
```powershell
# Автоматическая сборка для Windows
.\build.ps1
```

#### Linux/macOS  
```bash
# Автоматическая сборка для Unix систем
chmod +x build.sh
./build.sh
```

## 📋 Использование

### Основные команды

#### Управление задачами
```bash
# Добавить новую задачу
tasker add "Написать проект" "to-do"

# Показать все задачи
tasker list

# Обновить описание задачи
tasker update 1 "Протестировать проект"

# Удалить задачу
tasker delete 1
```

#### Просмотр задач по статусам
```bash
# Показать задачи к выполнению
tasker list-todo

# Показать задачи в процессе
tasker list-in-progress

# Показать выполненные задачи
tasker list-done
```

#### Изменение статусов задач
```bash
# Отметить задачу как выполняемую
tasker mark-in-progress 1

# Отметить задачу как выполненную
tasker mark-done 1

# Вернуть задачу к выполнению
tasker mark-to-do 1
```

### Полный список команд

| Команда | Описание |
|---------|----------|
| `add` | Добавить новую задачу |
| `list` | Показать все задачи |
| `list-todo` | Показать задачи к выполнению |
| `list-in-progress` | Показать задачи в процессе |
| `list-done` | Показать выполненные задачи |
| `update` | Изменить описание задачи |
| `mark-to-do` | Отметить как требующую выполнения |
| `mark-in-progress` | Отметить как выполняемую |
| `mark-done` | Отметить как выполненную |
| `delete` | Удалить задачу |

### Доступные статусы
- `to-do` - задача к выполнению
- `in-progress` - задача в процессе
- `done` - задача выполнена

## 📁 Архитектура проекта

```
CLI--/
├── 📁 cmd/
│   └── cli-task/
│       ├── main.go              # 🚀 Главный файл - точка входа
│       └── command/             # 📁 Модульная организация команд
│           ├── AddCmd.go        # ➕ Команда добавления задач
│           ├── ListCmd.go       # 📋 Команда просмотра всех задач
│           ├── ListTodoCmd.go   # 📝 Фильтр задач к выполнению
│           ├── ListProgressCmd.go # 🔄 Фильтр задач в процессе
│           ├── ListDoneCmd.go   # ✅ Фильтр выполненных задач
│           ├── UpdateCmd.go     # ✏️ Команда обновления описания
│           ├── MarkTodoCmd.go   # 📋 Отметить как требующую выполнения
│           ├── MarkProgressCmd.go # 🔄 Отметить как выполняемую
│           ├── MarkDoneCmd.go   # ✅ Отметить как выполненную
│           └── DeleteCmd.go     # 🗑️ Команда удаления задач
├── 📁 internal/
│   ├── domain/                  # 🏗️ Доменная модель
│   │   ├── task.go             # 📝 Структура задачи с инкапсуляцией
│   │   └── errors.go           # ❌ Доменные ошибки
│   ├── application/
│   │   └── services/
│   │       ├── task_service.go # 💼 Бизнес-логика управления задачами
│   │       └── functions.go    # 🔧 Вспомогательные функции
│   └── infrastructure/
│       └── storage/
│           └── json_rep.go     # 💾 Слой персистентности JSON
├── 🔧 Скрипты установки и сборки:
│   ├── install.ps1             # 🪟 Windows установщик (PowerShell)
│   ├── install.sh              # 🐧 Linux/macOS установщик (Bash)
│   ├── uninstall.ps1           # 🗑️ Windows деинсталлятор
│   ├── uninstall.sh            # 🗑️ Unix деинсталлятор
│   ├── build.ps1               # 🔨 Windows сборщик
│   └── build.sh                # 🔨 Unix сборщик
├── completion.ps1              # 🎯 PowerShell автодополнение
├── go.mod                      # 📦 Go модуль с зависимостями
├── go.sum                      # 🔒 Хеши зависимостей
└── README.md                   # 📖 Документация проекта
```
## 💡 Примеры использования

### Базовый рабочий процесс

```bash
# 1. Добавить новую задачу
tasker add "Изучить Go" 

# 2. Добавить задачу с определенным статусом
tasker add "Написать тесты" --status in-progress

# 3. Посмотреть все задачи
tasker list

# 4. Посмотреть только задачи к выполнению
tasker list-todo

# 5. Обновить описание задачи
tasker update 1 "Изучить Go и создать CLI приложение"

# 6. Изменить статус на "в процессе"
tasker mark-in-progress 1

# 7. Отметить задачу как выполненную
tasker mark-done 1

# 8. Удалить задачу
tasker delete 2
```

### Продвинутые сценарии

```bash
# Посмотреть только задачи в процессе выполнения
tasker list-in-progress

# Посмотреть только выполненные задачи
tasker list-done

# Вернуть выполненную задачу к статусу "к выполнению"
tasker mark-to-do 1

# Получить справку по любой команде
tasker add --help
tasker list --help
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

## 🗂️ Местоположение данных

- **После установки**: `%USERPROFILE%\.local\bin\tasks.json`
- **При разработке**: `.\tasks.json` (в директории проекта)
- **Файл создается автоматически** при первом добавлении задачи

## 🗑️ Деинсталляция

### 🪟 Windows (PowerShell)

```powershell
# Полная деинсталляция
.\uninstall.ps1

# Деинсталляция с сохранением данных
.\uninstall.ps1 -KeepData

# Принудительная деинсталляция без подтверждений
.\uninstall.ps1 -Force

# Быстрое удаление с сохранением данных
.\uninstall.ps1 -KeepData -Force
```

### 🐧 Linux/macOS

```bash
# Запустить деинсталлятор
./uninstall.sh

# Глобальная деинсталляция
./uninstall.sh --global

# Или удалить вручную
rm -f /usr/local/bin/tasker
rm -f ~/.local/bin/tasker
rm -f ~/tasks.json  # файл данных (опционально)
```

### ❌ Что удаляется при деинсталляции:

✅ **Всегда удаляется:**
- Исполняемый файл `tasker.exe` / `tasker`
- Записи в переменной PATH
- Автодополнение PowerShell (Windows)
- Временные файлы сборки

⚙️ **Опционально удаляется:**
- `tasks.json` (файл с вашими задачами)
- Пустые папки установки

💡 **Используйте `-KeepData` для сохранения ваших задач**

## 🛠️ Разработка

### Требования

- **Go 1.19+** - для сборки приложения
- **Git** - для клонирования репозитория
- **PowerShell** (Windows) или **Bash** (Linux/macOS) - для запуска скриптов

### Локальная разработка

```bash
# Клонировать репозиторий
git clone https://github.com/St1lon/CLI--
cd CLI--

# Загрузить зависимости
go mod download

# Запустить без сборки
go run cmd/cli-task/main.go --help

# Запустить тесты (когда будут добавлены)
go test ./...

# Проверить код
go vet ./...
go fmt ./...
```

### Быстрые команды разработчика

```bash
# Сборка для текущей платформы
go build -ldflags="-w -s" -o tasker cmd/cli-task/main.go   # Linux/macOS
go build -ldflags="-w -s" -o tasker.exe cmd/cli-task/main.go  # Windows

# Тестовый запуск
./tasker add "Test task" to-do
./tasker list

# Форматирование кода
go fmt ./...
```
