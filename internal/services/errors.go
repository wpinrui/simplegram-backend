// internal/services/errors.go
package services

import "errors"

var ErrUsernameExists = errors.New("username already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
var ErrTokenInvalid = errors.New("token is invalid")
var ErrTokenExpired = errors.New("token is expired")
var ErrTokenNotFound = errors.New("token not found")
var ErrTokenMalformed = errors.New("token is malformed")
var ErrTokenSignatureInvalid = errors.New("token signature is invalid")
