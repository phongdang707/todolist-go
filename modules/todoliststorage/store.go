package todoliststorage

import "database/sql"

type sqlStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *sqlStore {
	return &sqlStore{db: db}
}
