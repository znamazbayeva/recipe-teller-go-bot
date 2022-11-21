package bot

import (
	"log"
	"time"
	"fmt"
	// "io/ioutil"
	// "net/http"
	"upgrade/internal/models"
	"gopkg.in/telebot.v3"
	"upgrade/internal/repository"
)

type UpgradeBot struct {
	Bot   *telebot.Bot
	Users *repository.UserModel
}

func (bot *UpgradeBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Такой пользователь уже существует %v", err)
	}

	if existUser == nil {
		err := bot.Users.Create(newUser)
		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Привет, " + ctx.Sender().FirstName + "\nDon't know what to cook today? Reciper Teller is here to help you. \n\n 1. Get random recipe with /random \n 2. Get recipe by name /name {enter here recipe name}\n 3. Get random recipe with one ingredient /ingredient {enter ingredient name}")
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}

func (bot *UpgradeBot) ShowRandomRecipe(ctx telebot.Context) error {
	var randomRecipe string
	newRecipe, err := repository.GetRandomRecipe()
	if err != nil {
		randomRecipe += fmt.Sprintf("%s", &err)
	} else {
		randomRecipe += fmt.Sprintf("%s", &newRecipe)
	}
  
	return ctx.Send(randomRecipe)
}


func (bot *UpgradeBot) ShowRecipeByName(ctx telebot.Context) error { 
	args := ctx.Args()
	if len(args) != 1 {
		return ctx.Send("Введите только одно имя")
	}
	var recipeByName string

	newRecipe, err := repository.GetRecipeByName(args[0])
	if err != nil {
		recipeByName += fmt.Sprintf("%s", "There is no receipts with this name. \n Try again: /name {recipename}")
	} else {
		recipeByName += fmt.Sprintf("%s", &newRecipe)
	}
  
	return ctx.Send(recipeByName)
}

func (bot *UpgradeBot) ShowRecipeByIngredient(ctx telebot.Context) error { 
	args := ctx.Args()
	if len(args) != 1 {
		return ctx.Send("Введите только один ингредиент")
	}
	var recipeByIngredient string
	
	newRecipe, err := repository.GetrecipeByIngredient(args[0])
	if err != nil {
		recipeByIngredient += fmt.Sprintf("%s", "There is no receipts with this ingredient. \n Try again: /ingredient {recipe ingredient}")
	} else {
		recipeByIngredient += fmt.Sprintf("%s", &newRecipe)
	}
  
	return ctx.Send(recipeByIngredient)
}
