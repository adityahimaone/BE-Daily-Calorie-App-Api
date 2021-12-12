package response

import "Daily-Calorie-App-API/businesses/personaldata"

type PersonalData struct {
	Calorie float64 `json:"calorie"`
	Weight  int     `json:"weight"`
	Height  int     `json:"height"`
}

func FromDomain(domain personaldata.Domain) PersonalData {
	return PersonalData{
		Calorie: domain.Calorie,
		Weight:  domain.Weight,
		Height:  domain.Height,
	}
}
