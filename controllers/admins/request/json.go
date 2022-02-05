package request

import "Daily-Calorie-App-API/business/admins"

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToDomain(request *Admin) *admins.Domain {
	return &admins.Domain{
		Username: request.Username,
		Password: request.Password,
	}
}
