package models

import (
	"fmt"
	"os"
	"todo/interfaces"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDataBase struct {}

var Database interfaces.DatabaseInterface = new(PostgresDataBase)

func (database PostgresDataBase) Connect() (*gorm.DB, error) {

	var psqlConnectionString string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	
	return gorm.Open(postgres.Open(psqlConnectionString), &gorm.Config{})
}