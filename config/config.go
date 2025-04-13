package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config 用 Viper 读取配置文件
type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn         string
		MaxIdleConn int
		MaxOpenConn int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")   //操作配置文件的文件名
	viper.SetConfigType("yml")      //文件类型
	viper.AddConfigPath("./config") //文件路径
	//读取文件内容
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v ", err)
	}
	//通过结构体字段去直接访问
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshalling config %v", err)
	}
	initDB()
}
