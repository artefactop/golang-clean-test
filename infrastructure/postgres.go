package infrastructure

import (
	"database/sql"

	"github.com/artefactop/test_arch/interfaces/repositories"
)

type PostgresHandler struct {
	Conn string
}

func (handler *PostgresHandler) Execute(statement string) {

}

func (handler *PostgresHandler) Query(statement string) repositories.Row {
	return PostgresRow{}
}

type PostgresRow struct {
	Rows *sql.Rows
}

func (r PostgresRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r PostgresRow) Next() bool {
	return r.Rows.Next()
}

func NewPostgresHandler(config string) *PostgresHandler {
	return new(PostgresHandler)
}
