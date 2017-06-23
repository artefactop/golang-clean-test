package domain

import (
	"errors"
)

type CustomerRepository interface {
	FindById(id string) Customer
}

type Customer struct {
	Id      string
	Name    string
	Blocked bool
	Limits  bool
}

func (customer *Customer) CheckChargeLimits(amount int, currency string) error {
	if customer.Blocked {
		return errors.New("Cannot charge a blocked user")
	}
	if ok, err := hasExceededLimits(customer.Limits); ok {
		return err
	}
	return nil
}

func hasExceededLimits(limits interface{}) (bool, error) {
	// TODO get general limits
	// TODO check
	return false, errors.New("Limit by year")
}
