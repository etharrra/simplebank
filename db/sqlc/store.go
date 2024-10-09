package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to create db quries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new  Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

/**
 * Executes a transaction with the provided function using the given context. 
 * If an error occurs during the transaction or the function execution, it rolls back the transaction and returns the error. 
 * Otherwise, it commits the transaction.
 * 
 * @param ctx The context for the transaction
 * @param fn The function to be executed within the transaction
 * @return An error if any occurred during the transaction or function execution
*/
func (store *Store) execTr(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

/**
 * TransferTx performs a transaction to transfer funds between two accounts.
 * It creates a transfer record, debits the amount from the sender's account, credits the amount to the receiver's account,
 * and updates the account balances accordingly. If the sender's account has a lower ID than the receiver's account,
 * the transfer is processed in that order to prevent deadlocks.
 *
 * @param ctx The context for the transaction
 * @param arg The TransferTxParams containing the necessary parameters for the transfer
 * @return TransferTxResult The result of the transfer transaction
 * @return error An error if the transaction encounters any issues
 */
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTr(ctx, func(q *Queries) error {
		var err error

		// use arg in CrateTransferParams cuz fields in arg and CrateTransferParams are identical
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}
		return nil
	})

	return result, err
}

/**
 * addMoney adds money to two different accounts.
 *
 * Parameters:
 * - ctx: the context for the operation
 * - q: the set of queries to execute
 * - accountID1: the ID of the first account to add money to
 * - amount1: the amount to add to the first account
 * - accountID2: the ID of the second account to add money to
 * - amount2: the amount to add to the second account
 *
 * Returns:
 * - account1: the updated first account after adding money
 * - account2: the updated second account after adding money
 * - err: an error if any occurred during the operation
 */
func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	if err != nil {
		return
	}
	retur
}
