package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var database *sql.DB

func init() {
	godotenv.Load()
	connection, error := sql.Open("mysql", os.Getenv("DATABASE_LOGIN")+":"+os.Getenv("DATABASE_PASSWORD")+"@/"+os.Getenv("DATABASE_NAME"))
	if error != nil {
		panic(error.Error())
	}

	database = connection
}

func GetDatabase() *sql.DB {
	return database
}
