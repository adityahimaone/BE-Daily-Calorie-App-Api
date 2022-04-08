package admins

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business"
	"Daily-Calorie-App-API/helpers"
)

type serviceAdmin struct {
	adminRepository Repository
	jwtAuth         *auth.ConfigJWT
}

func NewServiceAdmin(adminRepository Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceAdmin{
		adminRepository: adminRepository,
		jwtAuth:         jwtAuth,
	}
}

func (service serviceAdmin) Register(admin *Domain) (*Domain, error) {
	passwordHash, err := helpers.PasswordHash(admin.Password)
	if err != nil {
		return &Domain{}, business.ErrInternalServer
	}
	admin.Password = passwordHash
	result, err := service.adminRepository.Insert(admin)
	if result == (&Domain{}) {
		return &Domain{}, business.ErrDuplicateData
	}
	if err != nil {
		return &Domain{}, business.ErrDuplicateData
	}
	return result, nil
}

func (service serviceAdmin) Login(username, password string) (string, error) {
	admin, err := service.adminRepository.GetAdminByUsername(username)
	if err != nil {
		return "User Not Found", business.ErrInternalServer
	}
	if admin.ID == 0 {
		return "User Not Found", business.ErrUserNotFound
	}
	if !helpers.ValidateHash(password, admin.Password) {
		return "Error Validate Hash", business.ErrInvalidLogin
	}
	token := service.jwtAuth.GenerateToken(admin.ID, "admin")
	return token, err
}
