/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package infra

// TITLE:DB設定

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// FUNCTION: DB setting
func InitDB(config *Config) func() {

	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Hostname,
		config.Postgres.Port,
		config.Postgres.Db,
	)

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

	boil.DebugMode = config.DebugMode

	// PROCESS:connection test
	if err = con.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Printf("db connection prepared [%s]\n", dns)
	return func() { con.Close() }
}
