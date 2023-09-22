package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	HostStr     = "7ac003ab7fac.c.methodot.com:34485"
	UserStr     = "root"
	PasswordStr = "123456"
	Dbname      = "test"
)

var (
	DbConn = new(gorm.DB)
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserStr, PasswordStr, HostStr, Dbname)
	DbConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
