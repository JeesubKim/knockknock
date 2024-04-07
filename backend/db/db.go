package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"knockknock/config"
	"log"
)

func InitDB() *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBAddress, config.Envs.DBName), // data source name
		DefaultStringSize:         256,                                                                                                                                                          // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                                                                         // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                                                                         // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                                                                         // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                                                                        // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connected!")
	return db
}
