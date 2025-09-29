#!/bin/bash

# Tasker CLI Installation Script for macOS/Linux
# Usage: ./install.sh [--global]

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
GRAY='\033[0;37m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}Installing Tasker CLI...${NC}"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go is not installed. Please install Go first.${NC}"
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

# Determine install path
if [[ "$1" == "--global" ]]; then
    INSTALL_PATH="/usr/local/bin"
    echo -e "${YELLOW}Installing globally to: $INSTALL_PATH${NC}"
    SUDO_CMD="sudo"
else
    INSTALL_PATH="$HOME/.local/bin"
    echo -e "${YELLOW}Installing locally to: $INSTALL_PATH${NC}"
    SUDO_CMD=""
fi

# Create install directory if it doesn't exist
if [[ ! -d "$INSTALL_PATH" ]]; then
    $SUDO_CMD mkdir -p "$INSTALL_PATH"
    echo -e "${CYAN}Created directory: $INSTALL_PATH${NC}"
fi

# Build the application
echo -e "${CYAN}Building tasker...${NC}"
go build -o tasker cmd/cli-task/main.go

if [[ ! -f "./tasker" ]]; then
    echo -e "${RED}Build failed! tasker binary not found.${NC}"
    exit 1
fi

# Install the binary
echo -e "${CYAN}Installing tasker to $INSTALL_PATH...${NC}"
$SUDO_CMD cp ./tasker "$INSTALL_PATH/"
$SUDO_CMD chmod +x "$INSTALL_PATH/tasker"

# Add to PATH if necessary
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

# Check if install path is in PATH
if [[ ":$PATH:" != *":$INSTALL_PATH:"* ]] && [[ "$1" != "--global" ]]; then
    echo -e "${YELLOW}Adding $INSTALL_PATH to PATH...${NC}"
    
    if [[ -n "$SHELL_CONFIG" ]]; then
        echo "" >> "$SHELL_CONFIG"
        echo "# Added by tasker installer" >> "$SHELL_CONFIG"
        echo "export PATH=\"$INSTALL_PATH:\$PATH\"" >> "$SHELL_CONFIG"
        echo -e "${CYAN}Added to $SHELL_CONFIG${NC}"
        echo -e "${YELLOW}Please restart your terminal or run: source $SHELL_CONFIG${NC}"
    else
        echo -e "${YELLOW}Please add $INSTALL_PATH to your PATH manually${NC}"
    fi
fi

# Clean up build artifact
rm -f ./tasker

echo ""
echo -e "${GREEN}✅ Installation completed!${NC}"
echo ""
echo -e "${GREEN}Usage examples:${NC}"
echo -e "${GRAY}  tasker add \"My task\"${NC}"
echo -e "${GRAY}  tasker list${NC}"
echo -e "${GRAY}  tasker list-todo${NC}"
echo -e "${GRAY}  tasker list-in-progress${NC}"
echo -e "${GRAY}  tasker list-done${NC}"
echo -e "${GRAY}  tasker update 1 \"New description\"${NC}"
echo -e "${GRAY}  tasker mark-in-progress 1${NC}"
echo -e "${GRAY}  tasker mark-done 1${NC}"
echo -e "${GRAY}  tasker delete 1${NC}"
echo ""
echo -e "${CYAN}For help: tasker --help${NC}"