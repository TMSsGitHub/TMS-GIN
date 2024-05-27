package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Database struct {
	Host      string `yml:"host"`
	Port      int    `yml:"port"`
	User      string `yml:"user"`
	Password  string `yml:"password"`
	Table     string `yml:"table"`
	Charset   string `yml:"charset"`
	ParseTime bool   `yml:"parse_time"`
	Loc       string `yml:"loc"`

	Prefix        string `yml:"prefix"`
	SingularTable bool   `yml:"singular_table"`
}

var DB *gorm.DB

func InitDB() {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		})
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		Cfg.Db.User,
		Cfg.Db.Password,
		Cfg.Db.Host,
		Cfg.Db.Port,
		Cfg.Db.Table,
		Cfg.Db.Charset,
		Cfg.Db.ParseTime,
		Cfg.Db.Loc,
	)
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: dbLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Cfg.Db.Prefix,
			SingularTable: Cfg.Db.SingularTable,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
