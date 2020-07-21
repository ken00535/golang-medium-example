#!bat
@echo off

if not exist bin (
    mkdir bin
)

if "%1"=="rebuild" (
    del /S /Q bin
)

for /f "usebackq" %%i in (`dir /b /on /a:d .\cmd`) do (
    echo %%i
    go build -o ./bin/%%i.exe -tags windows ./cmd/%%i
)
