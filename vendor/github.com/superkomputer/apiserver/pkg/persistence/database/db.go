package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	Psql_Server   = "PSQL_SERVER"
	Psql_Port     = "PSQL_PORT"
	Psql_User     = "PSQL_USER"
	Psql_Password = "PSQL_PASSWORD"
	Psql_Database = "PSQL_DB"
)

var db *sql.DB

type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// InitDB initialize DB
func InitDB() {
	var err error
	config := GetConfiguration()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

// GetConfiguration initializes configuration of psql
func GetConfiguration() Configuration {
	server := "127.0.0.1"
	if value, ok := os.LookupEnv(Psql_Server); ok {
		server = value
	}
	port := "5432"
	if value, ok := os.LookupEnv(Psql_Port); ok {
		port = value
	}
	user := "admin"
	if value, ok := os.LookupEnv(Psql_User); ok {
		user = value
	}
	password := "admin"
	if value, ok := os.LookupEnv(Psql_Password); ok {
		password = value
	}
	database := "mydb"
	if value, ok := os.LookupEnv(Psql_Database); ok {
		database = value
	}
	return Configuration{
		Server:   server,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
}
