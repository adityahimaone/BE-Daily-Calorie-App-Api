package historiesdetail

import "time"

type Domain struct {
	ID          int
	FoodID      int
	HistoriesID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Insert(historiesDetail *Domain) (*Domain, error)
	Update(id int, historiesDetail *Domain) (*Domain, error)
	Delete(id int) error
}
type Service interface {
	Insert(historiesDetail *Domain) (*Domain, error)
	Update(id int, historiesDetail *Domain) (*Domain, error)
	Delete(id int) error
}
