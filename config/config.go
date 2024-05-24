package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App App
	Db  Database
}

var Cfg Config

func init() {
	viper.SetConfigName("application.ini")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取错误, %s", err)
		return
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("配置文件解析错误: %v", err)
	}

	InitDB()
}
