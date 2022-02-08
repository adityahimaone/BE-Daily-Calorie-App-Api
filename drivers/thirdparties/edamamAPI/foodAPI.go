package edamamAPI

import (
	"Daily-Calorie-App-API/business/foodsAPI"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

type foodAPI struct {
	httpClient http.Client
}

func NewFoodAPI() foodsAPI.Repository {
	return &foodAPI{
		httpClient: http.Client{},
	}
}

func (api foodAPI) GetFoodByName(name string) (*[]foodsAPI.Domain, error) {
	appID := viper.GetString("edamamFoodAPI.appId")
	appKey := viper.GetString(`edamamFoodAPI.appKey`)
	splitQuery := strings.Split(name, " ")
	joinQuery := strings.Join(splitQuery, "%20")
	url := fmt.Sprintf("https://api.edamam.com/api/food-database/v2/parser?app_id=%s&app_key=%s&ingr=%s", appID, appKey, joinQuery)
	response, err := http.Get(url)
	if err != nil {
		return &[]foodsAPI.Domain{}, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &[]foodsAPI.Domain{}, err
	}
	defer response.Body.Close()

	food := FoodSource{}
	err = json.Unmarshal(responseData, &food)
	if err != nil {
		return &[]foodsAPI.Domain{}, err
	}

	var FoodsTemp []struct {
		Title       string
		ImgURL      string
		Calories    float64
		Fat         float64
		Protein     float64
		Carbs       float64
		ServingSize float64
	}

	for _, foodItem := range food.Hints {
		FoodsTemp = append(FoodsTemp, struct {
			Title       string
			ImgURL      string
			Calories    float64
			Fat         float64
			Protein     float64
			Carbs       float64
			ServingSize float64
		}{
			Title:       foodItem.Food.Label,
			ImgURL:      foodItem.Food.Image,
			Calories:    foodItem.Food.Nutrients.EnercKcal,
			Fat:         foodItem.Food.Nutrients.Fat,
			Protein:     foodItem.Food.Nutrients.Procnt,
			Carbs:       foodItem.Food.Nutrients.Chocdf,
			ServingSize: foodItem.Food.ServingsPerContainer,
		})
	}
	result := make([]foodsAPI.Domain, len(FoodsTemp))
	for i, foodItem := range FoodsTemp {
		result[i] = foodsAPI.Domain{
			Title:       foodItem.Title,
			ImgURL:      foodItem.ImgURL,
			Calories:    foodItem.Calories,
			Fat:         foodItem.Fat,
			Protein:     foodItem.Protein,
			Carbs:       foodItem.Carbs,
			ServingSize: foodItem.ServingSize,
		}
	}
	return &result, nil
}

func (api foodAPI) GetMealPlan(food *foodsAPI.Domain) (*[]foodsAPI.DomainRecipe, error) {
	appID := viper.GetString("edamamRecipeAPI.appId")
	appKey := viper.GetString(`edamamRecipeAPI.appKey`)
	url := fmt.Sprintf("https://api.edamam.com/search?q=%s&app_id=%s&app_key=%s&to=%s&diet=%s&calories=%s", food.MealTime, appID, appKey, food.PlanType, food.DietaryPreferences, food.RangeCalories)
	response, err := http.Get(url)
	if err != nil {
		return &[]foodsAPI.DomainRecipe{}, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	//log.Println(responseData)
	if err != nil {
		return &[]foodsAPI.DomainRecipe{}, err
	}
	defer response.Body.Close()

	recipe := RecipeSource{}
	err = json.Unmarshal(responseData, &recipe)
	if err != nil {
		return &[]foodsAPI.DomainRecipe{}, err
	}

	var RecipesTemp []struct {
		Label       string
		ImgURL      string
		Url         string
		Calories    float64
		Ingredients []string
	}

	for _, recipeItem := range recipe.Hits {
		RecipesTemp = append(RecipesTemp, struct {
			Label       string
			ImgURL      string
			Url         string
			Calories    float64
			Ingredients []string
		}{
			Label:       recipeItem.Recipe.Label,
			ImgURL:      recipeItem.Recipe.Image,
			Url:         recipeItem.Recipe.URL,
			Calories:    recipeItem.Recipe.Calories,
			Ingredients: recipeItem.Recipe.IngredientLines,
		})
	}

	result := make([]foodsAPI.DomainRecipe, len(RecipesTemp))
	for i, recipeItem := range RecipesTemp {
		result[i] = foodsAPI.DomainRecipe{
			RecipeLabel:       recipeItem.Label,
			RecipeImageURL:    recipeItem.ImgURL,
			RecipeURL:         recipeItem.Url,
			RecipeCalories:    recipeItem.Calories,
			RecipeIngredients: recipeItem.Ingredients,
		}
	}
	return &result, nil
}
