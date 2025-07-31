package constants

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("username already exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCreatingUser       = errors.New("failed to create user")
	ErrFetchingUser       = errors.New("error fetching user")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidPass       = errors.New("invalid  password")
)


