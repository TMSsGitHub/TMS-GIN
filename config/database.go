package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Driver    string
	Host      string
	Port      int
	User      string
	Password  string
	Name      string
	Charset   string
	ParseTime bool
	Loc       string
}

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		Cfg.Db.User,
		Cfg.Db.Password,
		Cfg.Db.Host,
		Cfg.Db.Port,
		Cfg.Db.Name,
		Cfg.Db.Charset,
		Cfg.Db.ParseTime,
		Cfg.Db.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
