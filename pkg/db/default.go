package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mncPaymentAPI/pkg/reader"
	"os"
)

func Default() (*gorm.DB, error) {
	var (
		username string
		password string
		host     string
		port     string
		dbName   string
	)

	reader.GetEnv(".env")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, dbName, port)

	dbConn, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			CreateBatchSize: 500,
		},
	)
	return dbConn, err
}
