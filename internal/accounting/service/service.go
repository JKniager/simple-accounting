package service

import "time"

type Ledger interface {
	GetAccounts(fltr AccountFilter) []*Account
	CreateAccount(name string, accT AccountType, startBal float64) error
	GetTransactions() []*Transaction
	AddTransaction(date time.Time, amt float64, creditAcc, debitAcc *Account, comment string) error
}

type AccountSummary struct {
	Name    string
	AccType string
	Bal     float64
}

type TransactionFilterer interface {
	Filter(t *Transaction) bool
}

type SpendingSummary struct {
	AccName  string
	AccType  string
	StartBal float64
	EndBal   float64
	DeltaBal float64
}

func NewAccountingService(ldgr Ledger) *AccountingService {
	return &AccountingService{
		ldgr: ldgr,
	}
}

type AccountingService struct {
	ldgr Ledger
}

func (s *AccountingService) CreateAccount(name string, accT AccountType, startBal float64) error {
	return s.ldgr.CreateAccount(name, accT, startBal)
}

func (s *AccountingService) GetAccountSummaries(fltr AccountFilter) []AccountSummary {
	accs := s.ldgr.GetAccounts(fltr)
	summary := make([]AccountSummary, 0, len(accs))
	for _, a := range accs {
		summary = append(summary, AccountSummary{
			Name:    a.Name,
			AccType: a.AccType.Name(),
			Bal:     a.Balance(),
		})
	}

	return summary
}

func (s *AccountingService) GetSpendingSummaries() []SpendingSummary {
	return []SpendingSummary{}
}

func (s *AccountingService) AddTransaction(date time.Time, amt float64, creditAcc, debitAcc *Account, comment string) error {
	return s.ldgr.AddTransaction(date, amt, creditAcc, debitAcc, comment)
}
