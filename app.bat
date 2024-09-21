@ECHO OFF

if "%1" == "up" (

  echo db setup...
  cd db
  docker-compose up -d
  dx2-migration.exe up
  cd ..

  echo mock setup...
  cd api
  CALL   mock.bat up
  cd ..

  exit /b
)

if "%1" == "down" (

  echo db cleanup...
  cd db
  docker-compose down
  cd ..

  echo mock cleanup...
  cd api
  CALL mock.bat down
  cd ..

  exit /b
)

  echo parameter required (up/down)...
