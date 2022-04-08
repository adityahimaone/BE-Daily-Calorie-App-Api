package response

import (
	"Daily-Calorie-App-API/business/users"
	"time"
)

type User struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	AvatarUrl      string    `json:"avatar_url"`
	Gender         string    `json:"gender"`
	PersonalDataID int       `json:"personal_data_id"`
	Calories       float64   `json:"calories"`
	Weight         int       `json:"weight"`
	Height         int       `json:"height"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:             uint64(domain.ID),
		Name:           domain.Name,
		Email:          domain.Email,
		Password:       domain.Password,
		AvatarUrl:      domain.AvatarUrl,
		PersonalDataID: domain.PersonalDataID,
		Calories:       domain.Calories,
		Weight:         domain.Weight,
		Height:         domain.Height,
		Gender:         domain.Gender,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func FromDomainArray(domain []users.Domain) []User {
	var result []User
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}

type UserLogin struct {
	Token string `json:"token"`
}

type CountCalories struct {
	Calories float64 `json:"calories"`
}
