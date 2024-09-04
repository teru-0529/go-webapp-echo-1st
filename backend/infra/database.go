/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package infra

// TITLE:DB設定

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// FUNCTION: DB setting
func InitDB(isDebug bool) func() {

	dns := getDnsEnv()

	// PROCESS:database open
	con, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	// PROCESS:connection pool settings
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

	// PROCESS:global connection setting
	boil.SetDB(con)
	if isDebug {
		boil.DebugMode = true
	}

	// PROCESS:connection test
	if err = con.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("*** db connection prepared *** [%s]\n", dns)
	return func() { con.Close() }
}

// FUNCTION:
func getDnsEnv() string {
	// dns from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST_NAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	return dns
}
