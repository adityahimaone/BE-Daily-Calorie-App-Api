package personaldata

type Domain struct {
	ID      int
	Calorie float64
	Weight  int
	Height  int
}
type Repository interface {
	Insert(personaldata *Domain) (*Domain, error)
	Update(id int, personaldata *Domain) (*Domain, error)
	Delete(personaldata *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}
