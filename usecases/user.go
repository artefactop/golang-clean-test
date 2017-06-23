package usecases

import "github.com/artefactop/test_arch/domain"

type UserRepository interface {
	FindById(id string) User
}

type User struct {
	Id       string
	IsAdmin  bool
	Customer domain.Customer
}
