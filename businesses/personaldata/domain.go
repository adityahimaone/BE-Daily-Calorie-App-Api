package personaldata

import "time"

type Domain struct {
	ID int
	Calorie float64
	Weight int
	Height int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Repository interface{
	Insert(personaldata *Domain) (*Domain,error)
	Update(personaldata *Domain) (*Domain,error)
	Delete(personaldata *Domain) (*Domain,error)
	FindByID(id int) (*Domain,error)
}