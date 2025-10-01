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

# Check if binary exists
if [[ -f "$BINARY_PATH" ]]; then
    echo -e "${CYAN}Removing tasker binary...${NC}"
    $SUDO_CMD rm -f "$BINARY_PATH"
    echo -e "${GREEN}✅ Binary removed: $BINARY_PATH${NC}"
else
    echo -e "${GRAY}Binary not found at: $BINARY_PATH${NC}"
fi

# Remove from PATH (only for local installations)
if [[ "$1" != "--global" ]]; then
    SHELL_CONFIG=""
    case $SHELL in
        */bash)
            SHELL_CONFIG="$HOME/.bashrc"
            if [[ -f "$HOME/.bash_profile" ]]; then
                SHELL_CONFIG="$HOME/.bash_profile"
            fi
            ;;
        */zsh)
            SHELL_CONFIG="$HOME/.zshrc"
            ;;
        */fish)
            SHELL_CONFIG="$HOME/.config/fish/config.fish"
            ;;
    esac

    if [[ -n "$SHELL_CONFIG" ]] && [[ -f "$SHELL_CONFIG" ]]; then
        if grep -q "# Added by tasker installer" "$SHELL_CONFIG"; then
            echo -e "${CYAN}Removing PATH entry from $SHELL_CONFIG...${NC}"
            # Remove the tasker installer lines
            sed -i.bak '/# Added by tasker installer/,/export PATH.*tasker/d' "$SHELL_CONFIG"
            echo -e "${GREEN}✅ PATH entry removed${NC}"
            echo -e "${YELLOW}Please restart your terminal or run: source $SHELL_CONFIG${NC}"
        fi
    fi
fi

# Optional: Remove tasks data (ask user)
TASKS_FILE="$HOME/tasks.json"
if [[ -f "$TASKS_FILE" ]]; then
    echo ""
    read -p "Do you want to remove task data file ($TASKS_FILE)? [y/N]: " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -f "$TASKS_FILE"
        echo -e "${GREEN}✅ Task data removed: $TASKS_FILE${NC}"
    else
        echo -e "${GRAY}Task data preserved: $TASKS_FILE${NC}"
    fi
fi

echo ""
echo -e "${GREEN}✅ Uninstallation completed!${NC}"
echo -e "${GRAY}Tasker CLI has been removed from your system.${NC}"