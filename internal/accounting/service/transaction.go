package service

import (
	"errors"
	"time"
)

var (
	ErrNilCreditAccount          = errors.New("nil credit account not allowed")
	ErrNilDebitAccount           = errors.New("nil debit account not allowed")
	ErrTransactionAlreadyApplied = errors.New("cannot apply an already applied transaction")
	ErrTransactionNotApplied     = errors.New("cannot undo an unapplied transaction")
	ErrZeroTransactionAmount     = errors.New("zero transaction amount not allowed")
)

type AccountTransactionInfo struct {
	Acc        *Account
	OldBalance *UnsettableFloat64
	NewBalance *UnsettableFloat64
}

func NewTransaction(date time.Time, amt float64, creditAcc, debitAcc *Account, comment string) (*Transaction, error) {
	if amt == 0.0 {
		return nil, ErrZeroTransactionAmount
	}

	if creditAcc == nil {
		return nil, ErrNilCreditAccount
	}

	if debitAcc == nil {
		return nil, ErrNilDebitAccount
	}

	return &Transaction{
		Date:   date,
		Amount: amt,
		CreditAccInfo: AccountTransactionInfo{
			Acc:        creditAcc,
			OldBalance: NewUnsettableFloat64(),
			NewBalance: NewUnsettableFloat64(),
		},
		DebitAccInfo: AccountTransactionInfo{
			Acc:        debitAcc,
			OldBalance: NewUnsettableFloat64(),
			NewBalance: NewUnsettableFloat64(),
		},
		Comment: comment,
	}, nil
}

type Transaction struct {
	Date          time.Time
	Amount        float64
	CreditAccInfo AccountTransactionInfo
	DebitAccInfo  AccountTransactionInfo
	Comment       string
	applied       bool
}

func (t *Transaction) Apply() error {
	if t.applied {
		return ErrTransactionAlreadyApplied
	}
	// Save old balances for record keeping.
	t.CreditAccInfo.OldBalance.Set(t.CreditAccInfo.Acc.Balance())
	t.DebitAccInfo.OldBalance.Set(t.DebitAccInfo.Acc.Balance())

	// Apply transaction to accounts and save new balances.
	t.CreditAccInfo.NewBalance.Set(t.CreditAccInfo.Acc.Credit(t.Amount))
	t.DebitAccInfo.NewBalance.Set(t.DebitAccInfo.Acc.Debit(t.Amount))
	t.applied = true
	return nil
}

func (t *Transaction) Undo() error {
	if !t.applied {
		return ErrTransactionNotApplied
	}
	// Reverse transaction.
	t.CreditAccInfo.Acc.Debit(t.Amount)
	t.DebitAccInfo.Acc.Credit(t.Amount)

	// Unset all recorded balances.
	t.CreditAccInfo.OldBalance.Unset()
	t.CreditAccInfo.NewBalance.Unset()
	t.DebitAccInfo.OldBalance.Unset()
	t.DebitAccInfo.NewBalance.Unset()
	t.applied = false
	return nil
}
