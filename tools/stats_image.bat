@echo off
setlocal enabledelayedexpansion

if "%~1"=="" (
  set /p testName=Enter the name of the test: 
) else (
  set "testName=%~1"
  for %%i in (%~1) do set testName=%%~ni
)

hakutest statistics "%testName%" image
