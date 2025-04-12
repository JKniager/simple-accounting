package ledger

import (
	"fmt"
	"log"
	"simple_accounting/internal/accounting/account"
	"time"
)

type Transaction struct {
	date time.Time
	crdt account.AccountId
	dbt  account.AccountId
	amt  float64
}

func (t *Transaction) Date() time.Time {
	return t.date
}

func (t *Transaction) CreditAccount() account.AccountId {
	return t.crdt
}

func (t *Transaction) DebitAccount() account.AccountId {
	return t.dbt
}

func (t *Transaction) Amount() float64 {
	return t.amt
}

type Ledger struct {
	accts map[account.AccountId]account.Account
	trns  []Transaction
}

func NewLedger() *Ledger {
	return &Ledger{
		accts: map[account.AccountId]account.Account{},
		trns:  []Transaction{},
	}
}

func (l *Ledger) AddAccount(acc account.Account) error {
	if acc == nil {
		log.Panic("tried to add nil account!")
	}

	if _, ok := l.accts[acc.Id()]; ok {
		return fmt.Errorf("tried to add account with a duplicate id '%s'", acc.Id())
	}

	l.accts[acc.Id()] = acc
	return nil
}

func (l *Ledger) AddTransaction(date time.Time, caId account.AccountId, daId account.AccountId, amt float64) {

	if caId == daId {
		log.Panic("credit and debit accout ids can't match!")
	}

	if amt == 0 {
		log.Panic("transaction amount can't be 0!")
	}

	ca, ok := l.accts[caId]
	if !ok {
		log.Panicf("ledger does not contain account with id %v!", caId)
	}

	da, ok := l.accts[daId]
	if !ok {
		log.Panicf("ledger does not contain account with id %v!", daId)
	}

	ca.Credit(amt)
	da.Debit(amt)
	l.trns = append(l.trns, Transaction{
		date: date,
		crdt: caId,
		dbt:  daId,
		amt:  amt,
	})
}

func (l *Ledger) GetAccountBalance(id account.AccountId) (float64, error) {
	if acc, ok := l.accts[id]; ok {
		return acc.Balance(), nil
	}
	return 0, fmt.Errorf("could not find account with id %v", id)
}
