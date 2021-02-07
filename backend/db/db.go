package db

import (
	"app/backend/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
	err error
)

func Init()  {
	host := config.GetEnv("POSTGRES_HOST", "")
	port := config.GetEnv("POSTGRES_PORT", "")
	user := config.GetEnv("POSTGRES_USER", "")
	pwd := config.GetEnv("POSTGRES_PASSWORD", "")
	dbName := config.GetEnv("POSTGRES_DB", "")
	sslMode := config.GetEnv("POSTGRES_SSL_ENABLED", "require")
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

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}