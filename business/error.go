package business

import "errors"

var (
	ErrDuplicateData  = errors.New("Duplicate data")
	ErrUserNotFound   = errors.New("User not found")
	ErrInsertData     = errors.New("Insert Data Failed")
	ErrUpdateData     = errors.New("Update Data Failed")
	ErrInvalidLogin   = errors.New("Username or Password is invalid")
	ErrGetData        = errors.New("Get Data Failed")
	ErrDeleteData     = errors.New("Delete Data Failed")
	ErrInternalServer = errors.New("Internal Server Error")
	ErrNotFound       = errors.New("Not Found")
	ErrUnAuthorized   = errors.New("UnAuthorized")
)
