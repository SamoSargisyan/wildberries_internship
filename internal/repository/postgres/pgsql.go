package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection(dbcfg string) (*sqlx.DB, error) {
	
	db, err := sqlx.Open("postgres", dbcfg)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
