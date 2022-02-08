package edamamAPI

type FoodSource struct {
	Hints []struct {
		Food struct {
			FoodID    string `json:"foodId"`
			Label     string `json:"label"`
			Nutrients struct {
				EnercKcal float64 `json:"ENERC_KCAL"`
				Procnt    float64 `json:"PROCNT"`
				Fat       float64 `json:"FAT"`
				Chocdf    float64 `json:"CHOCDF"`
				Fibtg     float64 `json:"FIBTG"`
			} `json:"nutrients"`
			Category             string  `json:"category"`
			CategoryLabel        string  `json:"categoryLabel"`
			Image                string  `json:"image"`
			ServingsPerContainer float64 `json:"servingsPerContainer"`
		} `json:"food,omitempty"`
	} `json:"hints"`
}

type RecipeSource struct {
	Q     string `json:"q"`
	From  int    `json:"from"`
	To    int    `json:"to"`
	More  bool   `json:"more"`
	Count int    `json:"count"`
	Hits  []struct {
		Recipe struct {
			URI             string   `json:"uri"`
			Label           string   `json:"label"`
			Image           string   `json:"image"`
			Source          string   `json:"source"`
			URL             string   `json:"url"`
			IngredientLines []string `json:"ingredientLines"`
			Ingredients     []struct {
				Text         string      `json:"text"`
				Quantity     float64     `json:"quantity"`
				Measure      interface{} `json:"measure"`
				Food         string      `json:"food"`
				Weight       float64     `json:"weight"`
				FoodCategory string      `json:"foodCategory"`
				FoodID       string      `json:"foodId"`
				Image        string      `json:"image"`
			} `json:"ingredients"`
			Calories float64 `json:"calories"`
		} `json:"recipe"`
	} `json:"hits"`
}

type Foods struct {
	Title       string
	ImgURL      string
	Calories    float64
	Fat         float64
	Protein     float64
	Carbs       float64
	ServingSize float64
}
