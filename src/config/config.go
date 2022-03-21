package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	BotToken   string   `json:"botToken" ,yaml:"botToken"`     // токен для телеграм бота
	AdminToken string   `json:"adminToken" ,yaml:"adminToken"` // токен для доступа к админке
	DB         DBConfig `json:"db" ,yaml:"db"`                 // настройки БД
}

type DBConfig struct {
	Host     string `json:"host" ,yaml:"host"`
	Name     string `json:"name" ,yaml:"name"`
	Port     int    `json:"port" ,yaml:"port"`
	User     string `json:"user" ,yaml:"user"`
	Password string `json:"password" ,yaml:"password"`
}

func (c Config) String() string {
	return fmt.Sprintf("{BotToken: %s AdminToken: %s, DB - Host: %s, Name: %s, Port: %d, User: %s, Password: %s}",
		c.BotToken,
		c.AdminToken,
		c.DB.Host,
		c.DB.Name,
		c.DB.Port,
		c.DB.User,
		c.DB.Password)
}

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config:GetConfig: Config file was not found")
		} else {
			log.Fatalln("Config:GetConfig: %w ", err)
		}
	}

	viper.AutomaticEnv()

	return Config{
		BotToken:   viper.GetString("botToken"),
		AdminToken: viper.GetString("adminToken"),
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Name:     viper.GetString("db.name"),
			Port:     viper.GetInt("db.port"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
		},
	}
}
