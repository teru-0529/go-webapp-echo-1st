@echo off
setlocal enabledelayedexpansion

set schema=orders

if "%1" == "lint" (
  call :func_lint %schema%
  exit /b
)

if "%1" == "bundle" (
  call :func_bundle %schema%
  exit /b
)

if "%1" == "build-docs" (
  call :func_build_doc %schema%
  exit /b
)

call :func_lint %schema%
call :func_bundle %schema%
call :func_build_doc %schema%
exit /b

:func_lint
echo [lint %1 openapi]...
docker-compose run --rm redocly lint ./src/services/%1/root.yaml
exit /b

:func_bundle
echo [bundle %1 openapi]...
docker-compose run --rm redocly bundle ./src/services/%1/root.yaml -o ./openapi/%1/openapi.yaml
exit /b

:func_build_doc
echo [build-docs %1 openapi]...
docker-compose run --rm redocly build-docs ./src/services/%1/root.yaml -o ./redoc/%1.html
exit /b
