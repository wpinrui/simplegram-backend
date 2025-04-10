package controllers

type UserServiceInterface interface {
	CreateUser(username, password string) (string, error)
	Login(username, password string) (string, error)
}
