package main

import (
	"net/http"

	"github.com/artefactop/test_arch/infrastructure"
	"github.com/artefactop/test_arch/interfaces/repositories"
	"github.com/artefactop/test_arch/interfaces/webhandlers"
	"github.com/artefactop/test_arch/usecases"
)

func main() {
	// TODO get configuration (infrastructure)
	// Interact with infrastructure
	dbHandler := infrastructure.NewPostgresHandler("configuration")

	// Dependency injection
	handlers := make(map[string]repositories.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbTransactionRepo"] = dbHandler

	transactionInteractor := new(usecases.TransactionInteractor)
	transactionInteractor.UserRepository = repositories.NewDbUserRepo(handlers)
	transactionInteractor.TransactionRepository = repositories.NewDbTransactionRepo(handlers)

	webserviceHandler := webhandlers.WebServiceHandler{}
	webserviceHandler.TransactionInteractor = transactionInteractor

	http.HandleFunc("/charges", func(res http.ResponseWriter, req *http.Request) {
		// Our entry point is the REST interface
		webserviceHandler.Charge(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
