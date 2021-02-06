package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"../config"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	//c := config.GetConfig()
	dsn := "host=localhost user=gorm password=gorm DB.name=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
