# PowerShell автодополнение для Tasker CLI
# Добавьте эту функцию в ваш PowerShell профиль

Register-ArgumentCompleter -Native -CommandName 'tasker' -ScriptBlock {
    param($commandName, $wordToComplete, $cursorPosition)
    
    $commands = @(
        'add',
        'list',
        'list-todo',
        'list-in-progress',
        'list-done',
        'update',
        'mark-in-progress',
        'mark-done',
        'help'
    )
    
    $statuses = @(
        'to-do',
        'in-progress', 
        'done'
    )
    
    # Получаем текущие аргументы
    $commandElements = $wordToComplete -split ' '
    $currentCommand = if ($commandElements.Count -gt 1) { $commandElements[1] } else { '' }
    
    switch ($currentCommand) {
        'add' {
            # Для команды add предлагаем статусы для второго аргумента
            if ($commandElements.Count -eq 4) {
                $statuses | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
                    [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
                }
            }
        }
        default {
            # Предлагаем основные команды
            $commands | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
                [System.Management.Automation.CompletionResult]::new($_, $_, 'Command', $_)
            }
        }
    }
}

Write-Host "✅ Автодополнение для tasker активировано!" -ForegroundColor Green
Write-Host "Доступные команды:" -ForegroundColor Cyan
Write-Host "  tasker add <описание> <статус>" -ForegroundColor Gray
Write-Host "  tasker list" -ForegroundColor Gray
Write-Host "  tasker list-todo" -ForegroundColor Gray
Write-Host "  tasker list-in-progress" -ForegroundColor Gray
Write-Host "  tasker list-done" -ForegroundColor Gray
Write-Host "  tasker update <id> <новое_описание>" -ForegroundColor Gray
Write-Host "  tasker mark-in-progress <id>" -ForegroundColor Gray
Write-Host "  tasker mark-done <id>" -ForegroundColor Gray