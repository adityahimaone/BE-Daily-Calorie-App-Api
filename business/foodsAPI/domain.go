package foodsAPI

type Domain struct {
	ID                 int
	Title              string
	ImgURL             string
	Calories           float64
	Fat                float64
	Protein            float64
	Carbs              float64
	ServingSize        float64
	MealTime           string
	DietaryPreferences string
	PlanType           string
	RangeCalories      string
	RecipeLabel        string
	RecipeImageURL     string
	RecipeURL          string
	RecipeCalories     float64
	RecipeIngredients  []string
	Lunch              []DomainRecipe
	Dinner             []DomainRecipe
	Breakfast          []DomainRecipe
}

type DomainRecipe struct {
	RecipeLabel       string   `json:"recipeLabel"`
	RecipeImageURL    string   `json:"recipeImageURL"`
	RecipeURL         string   `json:"recipeURL"`
	RecipeCalories    float64  `json:"recipeCalories"`
	RecipeIngredients []string `json:"recipeIngredients"`
}

type Repository interface {
	GetFoodByName(name string) (*[]Domain, error)
	GetMealPlan(food *Domain) (*[]DomainRecipe, error)
}

type Service interface {
	GetFoodByName(name string) (*[]Domain, error)
	GetMealPlan(food *Domain) (*Domain, error)
}
