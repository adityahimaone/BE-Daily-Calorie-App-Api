package users

import (
	"Daily-Calorie-App-API/business/personal_data"
	"time"
)

type Domain struct {
	ID             int
	Name           string
	Email          string
	Password       string
	AvatarUrl      string
	Gender         string
	PersonalDataID int
	Calories       float64
	Age            int
	ActivityType   int
	Weight         int
	Height         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Service interface {
	RegisterUser(userData *Domain, personalData *personal_data.Domain) (*Domain, error)
	Login(email string, password string) (string, error)
	GetUserByID(id int) (*Domain, error)
	EmailAvailable(email string) (bool, error)
	GetAllUser() (*[]Domain, error)
	EditUser(id int, user *Domain, personalData *personal_data.Domain) (*Domain, error)
	DeleteUser(id int) (string, error)
	CountCalories(userData *Domain, personalData *personal_data.Domain) (float64, error)
}

type Repository interface {
	Insert(user *Domain) (*Domain, error)
	Update(id int, user *Domain) (*Domain, error)
	GetUserByID(id int) (*Domain, error)
	GetUserByEmail(email string) (*Domain, error)
	GetAllUser() (*[]Domain, error)
	Delete(id int) (string, error)
}
