package repositories

import (
	"fmt"

	"github.com/artefactop/test_arch/domain"
)

type DbTransactionRepo DbRepo

func NewDbTransactionRepo(dbHandlers map[string]DbHandler) *DbTransactionRepo {
	dbTxRepo := new(DbTransactionRepo)
	dbTxRepo.dbHandlers = dbHandlers
	dbTxRepo.dbHandler = dbHandlers["DbTransactionRepo"]
	return dbTxRepo
}

func (repo *DbTransactionRepo) Store(tx domain.Transaction) (domain.Transaction, error) {
	tx.Id = "snowflake"
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO transactions (id, type, amount, currency, customer_id)
                                        VALUES ('%v', '%v', '%d', '%v', '%v')`,
		tx.Id, tx.Type, tx.Amount, tx.Currency, tx.Customer.Id))
	return tx, nil
}
