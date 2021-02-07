package db

import (
	"app/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
	err error
)

func Init()  {
	host := config.Getenv("POSTGRES_HOST", "")
	port := config.Getenv("POSTGRES_PORT", "")
	user := config.Getenv("POSTGRES_USER", "")
	pwd := config.Getenv("POSTGRES_PASSWORD", "")
	dbName := config.Getenv("POSTGRES_DB", "")
	sslMode := config.Getenv("POSTGRES_SSL_ENABLED", "require")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbName, pwd, sslMode)
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
}

func getDB() *gorm.DB {
	return db
}