package webhandlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/artefactop/test_arch/usecases"
)

type TransactionInteractor interface {
	Charge(customerId string, amount int, currency string) (usecases.Transaction, error)
}

type WebServiceHandler struct {
	TransactionInteractor TransactionInteractor
}

func (handler WebServiceHandler) Charge(res http.ResponseWriter, req *http.Request) {
	// TODO get customerId
	// TODO get amount
	// TODO get currency
	charge, _ := handler.TransactionInteractor.Charge("customer_id", 100, "EUR")
	io.WriteString(res, fmt.Sprintf("charge object: %v\n", charge)) // Here we should translate to json if needed
}
