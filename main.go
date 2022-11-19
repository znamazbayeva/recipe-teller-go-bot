package main

import (
	"flag"
	"log"
	"upgrade/cmd/bot"
	// "upgrade/internal/interfaces"
	// "upgrade/internal/models"
	"github.com/BurntSushi/toml"
	// "gopkg.in/telebot.v3"
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
	upgradeBot.Bot.Handle("/start",   upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/random", upgradeBot.ShowRandomRecipe)
	upgradeBot.Bot.Handle("/name", upgradeBot.ShowRecipeByName)
	upgradeBot.Bot.Handle("/ingredient", upgradeBot.ShowRecipeByIngredient)
	upgradeBot.Bot.Start()
}