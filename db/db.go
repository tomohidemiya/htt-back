package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var (
	db *gorm.DB
	err error
)

func Init()  {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass :=os.Getenv("POSTGRES_PASSWORD")
	dbName :=os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, pass)
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
}

func getDB() *gorm.DB {
	return db
}