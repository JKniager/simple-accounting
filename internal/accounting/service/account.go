package service

import "errors"

var (
	ErrInvalidAccountType = errors.New("invalid account type")
)

type AccountType int

const (
	Savings AccountType = iota
	Expense
)

var accTNames = map[AccountType]string{
	Savings: "Savings",
	Expense: "Expense",
}

func (a AccountType) Name() string {
	s, ok := accTNames[a]
	if !ok {
		return "UNKNOWN"
	}
	return s
}

type MoneyBucket interface {
	Balance() float64
	Credit(amt float64) float64
	Debit(amt float64) float64
}

type Asset struct {
	bal float64
}

func (a *Asset) Balance() float64 {
	return a.bal
}

func (a *Asset) Credit(amt float64) float64 {
	a.bal -= amt
	return a.Balance()
}

func (a *Asset) Debit(amt float64) float64 {
	a.bal += amt
	return a.Balance()
}

type Liability struct {
	bal float64
}

func (a *Liability) Balance() float64 {
	return a.bal
}

func (a *Liability) Credit(amt float64) float64 {
	a.bal += amt
	return a.Balance()
}

func (a *Liability) Debit(amt float64) float64 {
	a.bal -= amt
	return a.Balance()
}

func NewAccount(name string, accType AccountType, bal float64) (*Account, error) {
	var mnyBckt MoneyBucket
	switch accType {
	case Savings:
		mnyBckt = &Asset{bal: bal}
	case Expense:
		mnyBckt = &Liability{bal: bal}
	default:
		return nil, ErrInvalidAccountType
	}
	return &Account{
		Name:    name,
		AccType: accType,
		mb:      mnyBckt,
	}, nil
}

type Account struct {
	Name    string
	AccType AccountType
	mb      MoneyBucket
}

func (a *Account) Balance() float64 {
	return a.mb.Balance()
}

func (a *Account) Credit(amt float64) float64 {
	return a.mb.Credit(amt)
}

func (a *Account) Debit(amt float64) float64 {
	return a.mb.Debit(amt)
}

type AccountTypeFilter int

const (
	Any AccountTypeFilter = iota
	SavingsFilter
	ExpenseFilter
)

type AccountFilter struct {
	AccType AccountTypeFilter
}
