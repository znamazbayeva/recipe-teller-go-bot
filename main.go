package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sync"
	"upgrade/cmd/bot"
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
	fmt.Println(cfg, *configPath)
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		startBotHandler(cfg, db)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		startRMQ()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		startHttp()
		wg.Done()
	}()

	wg.Wait()
}

func startBotHandler(cfg *Config, db *gorm.DB) {
	upgradeBot := bot.UpgradeBot{
		Bot:   bot.InitBot(cfg.BotToken),
		Users: &repository.UserModel{Db: db},
	}

	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Handle("/random", upgradeBot.ShowRandomRecipe)
	upgradeBot.Bot.Handle("/name", upgradeBot.ShowRecipeByName)
	upgradeBot.Bot.Handle("/ingredient", upgradeBot.ShowRecipeByIngredient)

	upgradeBot.Bot.Start()

}

func startRMQ() {
	bot.GetLetterFromAdmin()
}

func startHttp() error {
	http.HandleFunc("/mails", getRoot)
	fmt.Print("starting there")
	return http.ListenAndServe("localhost:8080", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
