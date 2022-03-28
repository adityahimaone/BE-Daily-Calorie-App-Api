package request

import (
	"Daily-Calorie-App-API/business/personal_data"
	"Daily-Calorie-App-API/business/users"
	"Daily-Calorie-App-API/controllers/personal_data/request"
)

type User struct {
	Name           string               `json:"name"`
	Email          string               `json:"email"`
	Password       string               `json:"password"`
	AvatarUrl      string               `json:"avatar_url"`
	Gender         string               `json:"gender"`
	Age            int                  `json:"age"`
	ActivityType   float64              `json:"activity_type"`
	PersonalDataID int                  `json:"personal_data_id"`
	PersonalData   request.PersonalData `json:"personal_data"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(request *User) (*users.Domain, *personal_data.Domain) {
	return &users.Domain{
			Name:           request.Name,
			Email:          request.Email,
			Password:       request.Password,
			AvatarUrl:      request.AvatarUrl,
			Gender:         request.Gender,
			Age:            request.Age,
			ActivityType:   request.ActivityType,
			PersonalDataID: request.PersonalDataID,
		}, &personal_data.Domain{
			Calories: request.PersonalData.Calories,
			Weight:   request.PersonalData.Weight,
			Height:   request.PersonalData.Height,
		}
}
