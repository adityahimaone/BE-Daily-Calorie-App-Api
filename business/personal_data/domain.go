package personal_data

type Domain struct {
	ID       int
	Calories float64
	Weight   int
	Height   int
}
type Repository interface {
	Insert(personaldata *Domain) (*Domain, error)
	Update(id int, personaldata *Domain) (*Domain, error)
	Delete(id int) error
}
