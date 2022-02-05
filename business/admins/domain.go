package admins

import "time"

type Domain struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(domain *Domain) (*Domain, error)
	Login(username, password string) (string, error)
}

type Repository interface {
	Insert(domain *Domain) (*Domain, error)
	GetAdminByUsername(username string) (*Domain, error)
}
