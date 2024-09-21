@echo off
setlocal enabledelayedexpansion

set schema=orders

if "%1" == "up" (
  docker-compose -f docker-compose-4mock.yaml -p %schema% --env-file mockenv/.%schema%.env up -d
  exit /b
)

if "%1" == "down" (
  docker-compose -f docker-compose-4mock.yaml -p %schema% --env-file mockenv/.%schema%.env down
  exit /b
)
