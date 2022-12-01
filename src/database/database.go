package database

import (
	"lapsv/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Dbconnection() {

	DB, err = gorm.Open(mysql.Open(config.StringBanco), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panic("erro ao conectar no db")
	}
}
