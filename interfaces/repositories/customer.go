package repositories

import (
	"fmt"

	"github.com/artefactop/test_arch/domain"
)

type DbCustomerRepo DbRepo

func NewDbCustomerRepo(dbHandlers map[string]DbHandler) *DbCustomerRepo {
	dbCustomerRepo := new(DbCustomerRepo)
	dbCustomerRepo.dbHandlers = dbHandlers
	dbCustomerRepo.dbHandler = dbHandlers["DbCustomerRepo"]
	return dbCustomerRepo
}

func (repo *DbCustomerRepo) Store(customer domain.Customer) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO customers (id, name)
                                        VALUES ('%s', '%v')`,
		customer.Id, customer.Name))
}

func (repo *DbCustomerRepo) FindById(id string) domain.Customer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name FROM customers
                                             WHERE id = '%s' LIMIT 1`, id))
	var name string
	row.Next()
	row.Scan(&name)
	return domain.Customer{Id: id, Name: name}
}
