package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	WithTx(tx *sql.Tx) *Queries
	Begin() (*sql.Tx, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) Begin() (*sql.Tx, error) {
	return store.db.Begin()
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.Begin()
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rbErr: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
