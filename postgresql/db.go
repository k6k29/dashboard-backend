package postgresql

import (
	"dashboard/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

var ConnectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s", config.PGHost, config.PGPort, config.PGUser, config.PGPassword, config.PGDatabase, "disable")

func GetInstance() *gorm.DB {
	if DB == nil {
		DB, _ = gorm.Open("postgres", ConnectionString)
		DB.DB().SetMaxOpenConns(100)
		DB.DB().SetMaxIdleConns(20)
	}
	return DB
}

func init() {
	if DB, err := gorm.Open("postgres", ConnectionString); err != nil {
		log.Printf("postgres connect error %v", err)
	} else {
		if DB.Error != nil {
			log.Printf("database error %v", DB.Error)
		}
	}
}
