package database

import (
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
	stringDeConexao := "leo:atento@22@tcp(a853687:3306)/laps?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(stringDeConexao), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panic("erro ao conectar no db")
	}
}
