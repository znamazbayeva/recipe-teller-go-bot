package repository

import (
	"log"
	"io/ioutil"
	"net/http"
	"github.com/tidwall/gjson"
    "upgrade/internal/models"
)

func GetHttpData(url string, query string) (string, error) {
	var finalUrl string
	finalUrl = url + query
	response, err := http.Get(finalUrl)
	if err != nil {
		log.Printf("Ошибка получения рецепта %v", err)
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
	   return "", err
	}

	recipe := string(body)
	return recipe, nil 
}

func GetRecipeModel(recipe string) (models.Recipe, error) {
	id := (gjson.Get(recipe, "meals.#.idMeal")).Array()[0]
	title := (gjson.Get(recipe, "meals.#.strMeal")).Array()[0]
	category := (gjson.Get(recipe, "meals.#.strCategory")).Array()[0]
	cuisine := (gjson.Get(recipe, "meals.#.strArea")).Array()[0]
	instructions := (gjson.Get(recipe, "meals.#.strInstructions")).Array()[0]
	source := (gjson.Get(recipe, "meals.#.strSource")).Array()[0]
	
	newRecipe := models.Recipe{
		ID:       		id.Int(),
		Title: 			title.String(),
		Category:  		category.String(),
		Cuisine:   		cuisine.String(),
		Instructions:   instructions.String(),
		Source:			source.String(),
	}
	return newRecipe, nil
}

func GetRandomRecipe() (models.Recipe, error) {
	recipe, err := GetHttpData("https://www.themealdb.com/api/json/v1/1/random.php", "")

	if err != nil {
		return models.Recipe{}, err
	}

	newRecipe, err := GetRecipeModel(recipe)

	if err != nil {
		return models.Recipe{}, err
	}

	return newRecipe, nil
}

func GetRecipeByName(name string) (models.Recipe, error) { 
	recipe, err := GetHttpData("https://www.themealdb.com/api/json/v1/1/search.php?s=", name)
	if err != nil {
		return models.Recipe{}, err
	}
	newRecipe, err := GetRecipeModel(recipe)

	if err != nil {
		return models.Recipe{}, err
	}

	return newRecipe, nil
}

func GetrecipeByIngredient(ingredient string) (models.Recipe, error) { 
	recipe, err := GetHttpData("https://www.themealdb.com/api/json/v1/1/search.php?i=", ingredient)
	if err != nil {
		return models.Recipe{}, err
	}
	newRecipe, err := GetRecipeModel(recipe)

	if err != nil {
		return models.Recipe{}, err
	}

	return newRecipe, nil
}



