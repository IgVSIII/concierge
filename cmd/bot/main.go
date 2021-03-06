package main

import (
	"concierge-bot/src/config"
	"concierge-bot/src/db"
	"concierge-bot/src/tgapp"
	"log"
)

func main() {
	conf := config.GetConfig()
	log.Println(conf)

	connect := db.GetConnect(db.DBConfig{
		Host:     conf.DB.Host,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Name:     conf.DB.Name,
		Port:     conf.DB.Port,
	})

	connect.InitDB()
	connect.FillDB()

	tg := tgapp.TgApp{
		Token: conf.BotToken,
		DB:    connect,
	}

	tg.Run()

}
