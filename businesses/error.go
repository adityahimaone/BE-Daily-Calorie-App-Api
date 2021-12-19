package businesses

import "errors"

var (
	ErrDuplicateData = errors.New("Duplicate data")
	ErrUserNotFound  = errors.New("User not found")
	ErrEmailNotFound = errors.New("Email not found")
	ErrInsertData    = errors.New("Insert Data Failed")
	ErrUpdateData    = errors.New("Update Data Failed")
	ErrInvalidLogin  = errors.New("Invalid Login")
	ErrInvalidToken  = errors.New("Invalid Token")
	ErrInvalidData   = errors.New("Invalid Data")
	ErrInvalidUser   = errors.New("Invalid User")
	ErrInvalidEmail  = errors.New("Invalid Email")
	ErrGetData       = errors.New("Get Data Failed")
)
