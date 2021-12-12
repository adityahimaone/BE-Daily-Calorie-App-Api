package businesses

import "errors"

var (
	ErrDuplicateData = errors.New("Duplicate data")
	ErrUserNotFound  = errors.New("User not found")
	ErrEmailNotFound = errors.New("Email not found")
	ErrInsertData    = errors.New("Insert Data Failed")
	ErrUpdateData    = errors.New("Update Data Failed")
	ErrInvalidLogin  = errors.New("Invalid Login")
)
