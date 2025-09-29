#!/bin/bash

# ================================================================
# 🗑️ Быстрый деинсталлятор для Unix/Linux/macOS
# ================================================================

set -e

echo "🗑️ Деинсталляция CLI Task Manager (Tasker)"
echo "========================================"

# Функция для логирования
log() {
    echo "$(date '+%H:%M:%S') $1"
}

# Поиск и удаление исполняемого файла
INSTALL_LOCATIONS=(
    "/usr/local/bin/tasker"
    "$HOME/.local/bin/tasker"
    "$HOME/bin/tasker"
    "./tasker"
)

FOUND=false

for location in "${INSTALL_LOCATIONS[@]}"; do
    if [ -f "$location" ]; then
        log "✅ Найден tasker в: $location"
        rm -f "$location" && log "✅ Удален: $location" || log "❌ Ошибка удаления: $location"
        FOUND=true
    fi
done

# Поиск и удаление файлов данных
DATA_LOCATIONS=(
    "./tasks.json"
    "$HOME/tasks.json"
    "$HOME/.config/tasker/tasks.json"
    "$HOME/.local/share/tasker/tasks.json"
)

for location in "${DATA_LOCATIONS[@]}"; do
    if [ -f "$location" ]; then
        read -p "🗂️ Найден файл данных: $location. Удалить? (y/N): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            rm -f "$location" && log "✅ Удален файл данных: $location"
            # Удаляем пустые папки
            rmdir "$(dirname "$location")" 2>/dev/null || true
        fi
    fi
done

if [ "$FOUND" = true ]; then
    log "✅ Деинсталляция завершена успешно"
    log "🔄 Перезагрузите терминал для обновления PATH"
else
    log "⚠️ Tasker не найден в стандартных локациях"
fi

echo "🙏 Спасибо за использование CLI Task Manager!"