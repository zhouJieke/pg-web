package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Engin *gorm.DB

func InitMysql(config MysqlBase) {
	sprintf := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.UserName, config.Password, config.Host, config.Database)
	var err error
	Engin, err = gorm.Open("mysql", sprintf)
	if err != nil {
		panic("failed to connect database")
	}

}
