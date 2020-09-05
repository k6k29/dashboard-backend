package postgresql

import (
	"dashboard/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

var ConnectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s TimeZone=%s", config.PGHost, config.PGPort, config.PGUser, config.PGPassword, config.PGDatabase, "disable", "Asia/Shanghai")

func GetInstance() *gorm.DB {
	if DB == nil {
		var err error
		DB, err = gorm.Open(postgres.Open(ConnectionString), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
		if sqlDb, err := DB.DB(); err == nil {
			sqlDb.SetMaxIdleConns(100)
			sqlDb.SetMaxIdleConns(20)
		} else {
			panic(err.Error())
		}
	}
	return DB
}

func init() {
	if DB, err := gorm.Open(postgres.Open(ConnectionString), &gorm.Config{}); err != nil {
		log.Printf("postgres connect error %v", err)
	} else {
		if DB.Error != nil {
			log.Printf("database error %v", DB.Error)
		}
	}
}
