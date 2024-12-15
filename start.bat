@echo off
echo Выполняется запуск контейнеров...
command1
timeout /t 4 >nul

echo Выполняется развертывание модели...
command2
timeout /t 4 >nul

echo Приложение запущено!
pause
