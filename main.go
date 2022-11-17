package main

import (
	"flag"
	"log"
	"upgrade/cmd/bot"

	"github.com/BurntSushi/toml"
	"gopkg.in/telebot.v3"
)

type Config struct {
	Env      string
	BotToken string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	upgradeBot := bot.UpgradeBot{
		Bot: bot.InitBot(cfg.BotToken),
	}

	upgradeBot.Bot.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("Привет, " + ctx.Sender().FirstName)
	})
	upgradeBot.Bot.Start()
}