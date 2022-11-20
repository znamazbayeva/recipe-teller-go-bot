package main

import (
	"flag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"upgrade/cmd/bot"
	// "upgrade/internal/interfaces"
	// "upgrade/internal/models"
	"github.com/BurntSushi/toml"
	// "gopkg.in/telebot.v3"
	"upgrade/internal/repository"
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	upgradeBot := bot.UpgradeBot{
		Bot:   bot.InitBot(cfg.BotToken),
		Users: &repository.UserModel{Db: db},
	}
	upgradeBot.Bot.Handle("/start",   upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/random", upgradeBot.ShowRandomRecipe)
	upgradeBot.Bot.Handle("/name", upgradeBot.ShowRecipeByName)
	upgradeBot.Bot.Handle("/ingredient", upgradeBot.ShowRecipeByIngredient)

	upgradeBot.Bot.Start()
}
