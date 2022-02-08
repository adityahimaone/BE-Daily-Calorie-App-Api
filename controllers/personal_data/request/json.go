package request

type PersonalData struct {
	Calories float64 `json:"calorie"`
	Weight   int     `json:"weight"`
	Height   int     `json:"height"`
}
