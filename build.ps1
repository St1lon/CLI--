# build.ps1 - PowerShell скрипт сборки для Windows

Write-Host "Building Tasker CLI for Windows..." -ForegroundColor Green

# Сборка с оптимизацией
$output = "tasker.exe"
go build -ldflags="-w -s" -o $output cmd/cli-task/main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ Build successful: $output" -ForegroundColor Green
    Get-Item $output | Format-List Name, Length, LastWriteTime
} else {
    Write-Host "❌ Build failed" -ForegroundColor Red
    exit 1
}