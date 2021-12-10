package users

type serviceUsers struct {
	repository Repository
}

func NewService(repositoryUser Repository) Service {
	return &serviceUsers{
		repository: repositoryUser,
	}
}

func (s serviceUsers) RegisterUser(user *Domain) (*Domain, error) {
	panic("implement me")
}

func (s serviceUsers) Login(email string, password string) (string, error) {
	panic("implement me")
}

func (s serviceUsers) FindByID(id int) (*Domain, error) {
	panic("implement me")
}

func (s serviceUsers) FindByEmail(email string) (*Domain, error) {
	panic("implement me")
}

func (s serviceUsers) GetAllUsers() ([]*Domain, error) {
	panic("implement me")
}

func (s serviceUsers) EditUser(id int, user *Domain) (*Domain, error) {
	panic("implement me")
}

func (s serviceUsers) DeleteUser(id int, user *Domain) (*Domain, error) {
	panic("implement me")
}


