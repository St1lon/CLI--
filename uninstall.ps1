# Tasker Uninstaller
param(
    [switch]$KeepData,
    [switch]$Force,
    [switch]$Help
)

if ($Help) {
    Write-Host @"
Tasker Uninstaller

USAGE:
    .\uninstall.ps1 [options]

OPTIONS:
    -KeepData    Keep tasks.json data file
    -Force       Remove without confirmation  
    -Help        Show this help

EXAMPLES:
    .\uninstall.ps1                    # Normal uninstall
    .\uninstall.ps1 -KeepData          # Keep data
    .\uninstall.ps1 -Force             # Force remove
    .\uninstall.ps1 -KeepData -Force   # Quick remove, keep data

"@ -ForegroundColor Cyan
    exit 0
}

function Write-Log {
    param([string]$Message, [string]$Level = "INFO")
    $colors = @{
        "ERROR" = "Red"
        "SUCCESS" = "Green" 
        "WARNING" = "Yellow"
        "INFO" = "Cyan"
    }
    Write-Host "[$Level] $Message" -ForegroundColor $colors[$Level]
}

function Remove-SafeFile {
    param([string]$Path)
    if (Test-Path $Path) {
        try {
            Remove-Item $Path -Force -ErrorAction Stop
            Write-Log "Removed: $Path" "SUCCESS"
            return $true
        }
        catch {
            Write-Log "Error removing $Path" "ERROR"
            return $false
        }
    }
    return $true
}

function Remove-FromPath {
    param([string]$PathToRemove)
    try {
        $currentPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::User)
        if (-not $currentPath) { return $true }

        $pathParts = $currentPath -split ';' | Where-Object { $_ -ne $PathToRemove -and $_ -ne '' }
        $newPath = $pathParts -join ';'
        
        [Environment]::SetEnvironmentVariable("Path", $newPath, [EnvironmentVariableTarget]::User)
        Write-Log "Removed from PATH: $PathToRemove" "SUCCESS"
        return $true
    }
    catch {
        Write-Log "Error removing from PATH" "ERROR"
        return $false
    }
}

Clear-Host
Write-Host "TASKER UNINSTALLER" -ForegroundColor Red -BackgroundColor White
Write-Host "==================" -ForegroundColor Red

# Find installations
$locations = @(
    $PWD.Path,
    "$env:LOCALAPPDATA\Programs\Tasker",
    "$env:USERPROFILE\.local\bin"
)

$found = @()
foreach ($location in $locations) {
    $exePath = Join-Path $location "tasker.exe"
    if (Test-Path $exePath) {
        $found += @{ Path = $location; ExePath = $exePath }
        Write-Log "Found Tasker at: $location" "INFO"
    }
}

# Find data files
$dataFiles = @(
    "$PWD\tasks.json",
    "$env:USERPROFILE\tasks.json"
)

$foundData = @()
foreach ($dataPath in $dataFiles) {
    if (Test-Path $dataPath) {
        $foundData += $dataPath
        if ($KeepData) {
            Write-Log "Data (keeping): $dataPath" "INFO"
        } else {
            Write-Log "Data (removing): $dataPath" "WARNING"
        }
    }
}

if ($found.Count -eq 0 -and $foundData.Count -eq 0) {
    Write-Log "Tasker not found" "WARNING"
    exit 0
}

# Confirmation
if (-not $Force) {
    $confirm = Read-Host "`nRemove Tasker? (y/N)"
    if ($confirm -notmatch '^[Yy]') {
        Write-Log "Cancelled" "INFO"
        exit 0
    }
}

Write-Host "`nRemoving..." -ForegroundColor Yellow

$success = 0
$errors = 0

# Remove program files
foreach ($item in $found) {
    if (Remove-SafeFile $item.ExePath) { $success++ } else { $errors++ }
    
    if ($item.Path -ne $PWD.Path) {
        if (Remove-FromPath $item.Path) { $success++ } else { $errors++ }
    }
}

# Remove data files
if (-not $KeepData) {
    foreach ($dataPath in $foundData) {
        if (Remove-SafeFile $dataPath) { $success++ } else { $errors++ }
    }
}

# Results
Write-Host "`n" + "="*40 -ForegroundColor Gray
if ($errors -eq 0) {
    Write-Host "UNINSTALL COMPLETE!" -ForegroundColor Green
    Write-Host "Items removed: $success" -ForegroundColor Green
} else {
    Write-Host "COMPLETED WITH ERRORS" -ForegroundColor Yellow
    Write-Host "Success: $success | Errors: $errors" -ForegroundColor Yellow
}

if ($KeepData -and $foundData.Count -gt 0) {
    Write-Host "`nData files preserved:" -ForegroundColor Cyan
    $foundData | ForEach-Object { Write-Host "  $_" -ForegroundColor White }
}

Write-Host "`nRestart terminal to update PATH" -ForegroundColor Yellow
Write-Host "Thank you for using Tasker!" -ForegroundColor Green

if (-not $Force) {
    Read-Host "`nPress Enter to exit"
}