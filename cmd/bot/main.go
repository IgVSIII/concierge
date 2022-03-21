package main

import (
	"concierge-bot/src/config"
	"concierge-bot/src/db"
	"concierge-bot/src/tgapp"
	"fmt"
)

func main() {
	conf := config.GetConfig()

	fmt.Println(conf)

	connect := db.GetConnect(db.DBConfig{
		Host:     conf.DB.Host,
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Name:     conf.DB.Name,
		Port:     conf.DB.Port,
	})

	connect.InitDB()

	tg := tgapp.TgApp{
		Token: conf.BotToken,
		DB:    connect,
	}

	tg.Run()

}
