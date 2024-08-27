@echo off
setlocal enabledelayedexpansion

if "%1" == "up" (
  docker-compose -f docker-compose-4mock.yaml -p orders --env-file mockenv/.orders.env up -d
  exit /b
)

if "%1" == "down" (
  docker-compose -f docker-compose-4mock.yaml -p orders --env-file mockenv/.orders.env down
  exit /b
)
