package dbmngr

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"simple_accounting/internal/accounting/account"

	_ "github.com/mattn/go-sqlite3"
)

const (
	// Table initialization commands

	dbAccInitStr = `
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER NOT NULL PRIMARY KEY,
		type INTEGER NOT NULL,
		name TEXT NOT NULL,
		balance REAL NOT NULL
	);`

	dbTransInitStr = `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER NOT NULL PRIMARY KEY,
		credit_account INTEGER NOT NULL,
		debit_account INTEGER NOT NULL,
		amount REAL NOT NULL,
		FOREIGN KEY(credit_account) REFERENCES accounts(id),
		FOREIGN KEY(debit_account) REFERENCES accounts(id)
	);`

	// Select commands

	dbAccListSelectStr = `SELECT * FROM accounts;`
)

var (
	// Errors
	ErrTableNotExist      = errors.New("table does not exist")
	ErrUnknownAccountType = errors.New("unknown account type")
)

type DatabaseManager struct {
	ctx context.Context
	db  *sql.DB
}

func NewDatabaseManager(ctx context.Context, dbnm string) (*DatabaseManager, error) {
	if _, err := os.Stat(dbnm); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to open database from path %s", dbnm)
	}
	return connectToDatabase(ctx, dbnm)
}

func CreateDatabaseAndGetManager(ctx context.Context, dbPath string) (*DatabaseManager, error) {
	if _, err := os.Stat(dbPath); err == nil {
		return nil, errors.New("simple-account database already exists")
	}
	dbmgr, err := connectToDatabase(ctx, dbPath)
	if err != nil {
		return nil, err
	}
	if _, err := dbmgr.db.ExecContext(dbmgr.ctx, dbAccInitStr); err != nil {
		return nil, err
	}
	if _, err := dbmgr.db.ExecContext(dbmgr.ctx, dbTransInitStr); err != nil {
		return nil, err
	}

	return dbmgr, nil
}

func connectToDatabase(ctx context.Context, dbnm string) (*DatabaseManager, error) {
	db, err := sql.Open("sqlite3", dbnm)
	if err != nil {
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &DatabaseManager{
		ctx: ctx,
		db:  db,
	}, nil
}

func (dbmg *DatabaseManager) GetAccountList() ([]account.Account, error) {
	rows, err := dbmg.db.QueryContext(dbmg.ctx, dbAccListSelectStr)
	if err != nil {
		return nil, err
	}
	accs := []account.Account{}
	for rows.Next() {
		var (
			accId   int
			accType int
			accName string
			accBal  float64
		)

		if err := rows.Scan(&accId, &accType, &accName, &accBal); err != nil {
			return accs, err
		}
		a, err := createAccount(accId, account.AccountType(accType), accName, accBal)
		if err != nil {
			log.Printf("Error while creating account: %s", err.Error())
			continue
		}
		accs = append(accs, a)
	}

	if rows.Err() != nil {
		return accs, rows.Err()
	}

	return accs, nil
}

func createAccount(id account.AccountId, accType account.AccountType, name string, bal float64) (account.Account, error) {
	switch accType {
	case account.AccountTypeSavings:
		return account.NewSavingsAcct(id, name, bal), nil
	case account.AccountTypeExpense:
		return account.NewExpenseAcct(id, name, bal), nil
	}
	return nil, ErrUnknownAccountType
}

func (dbmg *DatabaseManager) AddAccount(aType account.AccountType, name string, balance float64) (account.AccountId, error) {

	return -1, nil
}
