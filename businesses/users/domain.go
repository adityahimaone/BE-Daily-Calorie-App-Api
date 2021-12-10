package users

import (
	"Daily-Calorie-App-API/businesses/personaldata"
	"time"
)

type Domain struct {
	ID int
	Name string
	Email string
	Password string
	AvatarUrl string
	Gender string
	PersonalDataID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface{
	RegisterUser(userData *Domain, personalData *personaldata.Domain) (*Domain, error)
	Login(email string, password string) (string, error)
	FindByID(id int) (*Domain, error)
	EmailAvailable(email string) (bool, error)
	GetAllUsers() ([]*Domain, error)
	EditUser(id int,user *Domain) (*Domain, error)
	DeleteUser(id int, user *Domain) (*Domain,error)
}

type Repository interface{
	Insert(user *Domain) (*Domain,error)
	Update(id int, user *Domain) (*Domain,error)
	FindByID(id int) (*Domain,error)
	FindByEmail(email string) (*Domain,error)
	GetAllUsers() ([]*Domain,error)
	Delete(id int, user *Domain) (*Domain,error)
}


