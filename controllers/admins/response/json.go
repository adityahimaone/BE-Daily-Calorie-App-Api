package response

import (
	"Daily-Calorie-App-API/business/admins"
	"time"
)

type Admin struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminLogin struct {
	Token string `json:"token"`
}

func FromDomain(domain admins.Domain) Admin {
	return Admin{
		ID:        domain.ID,
		Username:  domain.Username,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
