package models

type Recipe struct {
    ID              int64       `json:"idMeal"`
    Title           string      `json:"strMeal"`
    Category        string      `json:"strCategory"`
    Cuisine         string      `json:"strArea"`
    Instructions    string      `json:"strInstructions"`
	Source			string      `json:"strSource"`
}