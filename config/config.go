package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App   *App      `mapstructure:"app"`
	Db    *Database `mapstructure:"mysql"`
	Redis *Redis    `mapstructure:"redis"`
	File  *File     `mapstructure:"file"`
}

var Cfg Config

func InitConfig() {
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取错误, %s", err)
		return
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("配置文件解析错误: %v", err)
	}
}
