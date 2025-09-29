#!/bin/bash

# Tasker CLI Uninstall Script for macOS/Linux
# Usage: ./uninstall.sh [--global]

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
GRAY='\033[0;37m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${RED}Uninstalling Tasker CLI...${NC}"

# Determine install path
if [[ "$1" == "--global" ]]; then
    INSTALL_PATH="/usr/local/bin"
    echo -e "${YELLOW}Removing from global location: $INSTALL_PATH${NC}"
    SUDO_CMD="sudo"
else
    INSTALL_PATH="$HOME/.local/bin"
    echo -e "${YELLOW}Removing from local location: $INSTALL_PATH${NC}"
    SUDO_CMD=""
fi

BINARY_PATH="$INSTALL_PATH/tasker"
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