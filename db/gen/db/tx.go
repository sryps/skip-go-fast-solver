package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// InTx executes a callback fn of database operations within a transaction.
func (q *Queries) InTx(ctx context.Context, fn func(ctx context.Context, q Querier) error, opts *sql.TxOptions) error {
	if _, ok := q.db.(*sql.Tx); ok {
		// if we are already in a transaction, reuse it, and leave handling
		// commit/rollback to the parent tx.
		if err := fn(ctx, q); err != nil {
			return fmt.Errorf("executing nested transaction: %w", err)
		}
		return nil
	}

	txer, ok := q.db.(*sql.DB)
	if !ok {
		return fmt.Errorf("db is not an *sql.DB")
	}

	tx, err := txer.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("beginning tx: %w", err)
	}

	defer func() {
		rerr := tx.Rollback()
		if rerr == nil || errors.Is(rerr, sql.ErrTxDone) {
			return
		}
		err = fmt.Errorf("defer (%s): %w", rerr.Error(), err)
	}()

	if err = fn(ctx, q.WithTx(tx)); err != nil {
		return fmt.Errorf("executing transaction: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
