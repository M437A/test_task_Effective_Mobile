package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/viper"
)

var DB *sql.DB

func createDatabaseConnection() *sql.DB {
	host := viper.Get("db.host")
	port := viper.Get("db.port")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := viper.Get("db.database")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("Error opening database: %v", err))
	}

	log.Println("Successfully connected to database")

	runMigrations(db)
	log.Println("Migrations completed successfully")

	DB = db
	return db
}

func runMigrations(db *sql.DB) {
	migrationsDir := "migrations"

	if err := goose.Up(db, migrationsDir); err != nil {
		panic(fmt.Sprintf("Error running migrations: %v", err))
	}
}
