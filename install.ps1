param([switch]$Global)

Write-Host "Installing Tasker CLI..." -ForegroundColor Green

if ($Global) {
    $InstallPath = "C:\tools\bin"
    $PathScope = "Machine"
} else {
    $InstallPath = "$env:USERPROFILE\.local\bin"
    $PathScope = "User"
}

Write-Host "Install path: $InstallPath" -ForegroundColor Yellow

if (-not (Test-Path $InstallPath)) {
    New-Item -ItemType Directory -Path $InstallPath -Force | Out-Null
    Write-Host "Created directory: $InstallPath" -ForegroundColor Cyan
}

$SourcePath = ".\tasker.exe"
$DestinationPath = Join-Path $InstallPath "tasker.exe"

if (Test-Path $SourcePath) {
    Copy-Item $SourcePath $DestinationPath -Force
    Write-Host "Copied tasker.exe to $DestinationPath" -ForegroundColor Cyan
} else {
    Write-Error "tasker.exe not found!"
    exit 1
}

$CurrentPath = [Environment]::GetEnvironmentVariable("PATH", $PathScope)
if ($CurrentPath -notlike "*$InstallPath*") {
    $NewPath = $CurrentPath + ";" + $InstallPath
    [Environment]::SetEnvironmentVariable("PATH", $NewPath, $PathScope)
    Write-Host "Added to PATH" -ForegroundColor Green
} else {
    Write-Host "Already in PATH" -ForegroundColor Gray
}

Write-Host ""
Write-Host "Installation completed!" -ForegroundColor Green
Write-Host "Usage examples:" -ForegroundColor White
Write-Host '  tasker add "My task" "to-do"' -ForegroundColor Gray
Write-Host "  tasker list" -ForegroundColor Gray
Write-Host "  tasker list-todo" -ForegroundColor Gray
Write-Host "  tasker list-in-progress" -ForegroundColor Gray
Write-Host "  tasker list-done" -ForegroundColor Gray
Write-Host '  tasker update 1 "New description"' -ForegroundColor Gray
Write-Host "  tasker mark-in-progress 1" -ForegroundColor Gray
Write-Host "  tasker mark-done 1" -ForegroundColor Gray
Write-Host "  tasker mark-to-do 1" -ForegroundColor Gray
Write-Host ""
Write-Host "For autocompletion, run: . .\completion.ps1" -ForegroundColor Cyan
Write-Host "Please restart your terminal to use the command." -ForegroundColor Yellow