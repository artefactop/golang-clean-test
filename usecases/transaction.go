package usecases

import (
	"errors"
	"fmt"

	"github.com/artefactop/test_arch/domain"
)

type Logger interface {
	Log(message string) error
}

type Transaction struct {
	Id       string
	Type     string
	User     User
	Amount   int
	Currency string
}

type TransactionInteractor struct {
	UserRepository        UserRepository
	TransactionRepository domain.TransactionRepository
	Logger                Logger
}

func (interactor *TransactionInteractor) Charge(userId string, amount int, currency string) (Transaction, error) {
	// Interact with repository (interface)
	user := interactor.UserRepository.FindById(userId)
	if user == (User{}) {
		return Transaction{}, errors.New("User does not exists")
	}
	// Check business rules in Domain layer
	err := user.Customer.CheckChargeLimits(amount, currency)
	if err != nil {
		return Transaction{}, errors.New("Limits exceeded")
	}
	// Translate from usecase to domain
	domainTx := domain.Transaction{Amount: amount, Currency: currency}
	// Check business rules
	if domainTx.IsValidCharge() {
		return Transaction{}, errors.New("Not a valid charge")
	}

	// Interact with repository (interface)
	newDomainTx, _ := interactor.TransactionRepository.Store(domainTx)

	//Translate from domain to usecase
	tx := Transaction{
		Id:       newDomainTx.Id,
		Type:     newDomainTx.Type,
		User:     user,
		Amount:   newDomainTx.Amount,
		Currency: newDomainTx.Currency,
	} // there is a bug here because domain.Customer is not usecase.User but we should use an interface

	// Interact with logger (interface)
	// Log usecase event
	interactor.Logger.Log(fmt.Sprintf(
		"User '%s' made a tx of %d %s",
		tx.User.Id, tx.Amount, tx.Currency))
	return tx, nil
}
