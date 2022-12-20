package db

import (
	"grpc_dev/server/db/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	sql_string = "root:241503332@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	Client *gorm.DB
	err    error
)

func init() {
	Client, err = gorm.Open(mysql.Open(sql_string), &gorm.Config{})
	if err != nil {
		panic("Can't connect db")
	}

	err = migration()
	if err != nil {
		panic("Migrate error")
	}
}

func migration() error {

	if errors := Client.AutoMigrate(&dto.Country{}); errors != nil {
		return errors
	}

	if errors := Client.Table("user_account").AutoMigrate(&dto.UserAccount{}); errors != nil {
		return errors
	}
	return nil
}
