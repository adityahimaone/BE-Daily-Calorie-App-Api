package response

import (
	"Daily-Calorie-App-API/businesses/personaldata"
	"Daily-Calorie-App-API/businesses/users"
	"Daily-Calorie-App-API/controllers/personaldata/response"
	"time"
)

type User struct {
	Name           string                `json:"name"`
	Email          string                `json:"email"`
	Password       string                `json:"password"`
	AvatarUrl      string                `json:"avatar_url"`
	Gender         string                `json:"gender"`
	PersonalDataID int                   `json:"personal_data_id"`
	PersonalData   response.PersonalData `json:"personalData"`
	CreatedAt      time.Time             `json:"createdAt"`
	UpdatedAt      time.Time             `json:"updatedAt"`
}

func FromDomain(domain users.Domain, domainData personaldata.Domain) User {
	return User{
		Name:           domain.Name,
		Email:          domain.Email,
		Password:       domain.Password,
		AvatarUrl:      domain.AvatarUrl,
		PersonalDataID: domain.PersonalDataID,
		PersonalData: response.PersonalData{
			Calorie: domainData.Calorie,
			Weight:  domainData.Weight,
			Height:  domainData.Height,
		},
		Gender:    domain.Gender,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type UserLogin struct {
	Token string `json:"token"`
}
