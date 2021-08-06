package db

import (
	"kanbanmusume_ura/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

const (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "postgres"
	PORT     = "5432"
	DBNAME   = "kanbanmusume"
)

func Init() {
	db, err = gorm.Open("postgres", "host="+HOST+" port="+PORT+" user="+USER+" password="+PASSWORD+" dbname="+DBNAME+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigrate()
}

func Connect() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigrate() {
	db.AutoMigrate(&models.User{})
}
