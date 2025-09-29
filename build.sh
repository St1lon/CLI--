#!/bin/bash

# build.sh - Универсальный скрипт сборки

# Определяем платформу
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Определяем имя файла
if [[ "$OS" == "windows" ]] || [[ "$OS" == *"mingw"* ]] || [[ "$OS" == *"cygwin"* ]]; then
    OUTPUT="tasker.exe"
else
    OUTPUT="tasker"
fi

echo "Building for platform: $OS"
echo "Output file: $OUTPUT"

# Сборка с оптимизацией
go build -ldflags="-w -s" -o "$OUTPUT" cmd/cli-task/main.go

if [[ $? -eq 0 ]]; then
    echo "✅ Build successful: $OUTPUT"
    ls -la "$OUTPUT"
else
    echo "❌ Build failed"
    exit 1
fi