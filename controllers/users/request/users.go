package request

import (
	"Daily-Calorie-App-API/businesses/personaldata"
	"Daily-Calorie-App-API/businesses/users"
	"Daily-Calorie-App-API/controllers/personaldata/request"
)

type User struct {
	Name           string               `json:"name"`
	Email          string               `json:"email"`
	Password       string               `json:"password"`
	AvatarUrl      string               `json:"avatar_url"`
	Gender         string               `json:"gender"`
	PersonalDataID int                  `json:"personal_data_id"`
	PersonalData   request.PersonalData `json:"personal_data"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(request *User) (*users.Domain, *personaldata.Domain) {
	return &users.Domain{
			Name:           request.Name,
			Email:          request.Email,
			Password:       request.Password,
			AvatarUrl:      request.AvatarUrl,
			Gender:         request.Gender,
			PersonalDataID: request.PersonalDataID,
		}, &personaldata.Domain{
			Calorie: request.PersonalData.Calorie,
			Weight:  request.PersonalData.Weight,
			Height:  request.PersonalData.Height,
		}
}
