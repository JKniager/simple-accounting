package ledger

import (
	"simple_accounting/internal/accounting/account"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LedgerTestSuite struct {
	suite.Suite
}

func TestLedgerTestSuite(t *testing.T) {
	suite.Run(t, new(LedgerTestSuite))
}

func (s *LedgerTestSuite) TestAddAccount() {
	ldg := NewLedger()

	// Panics on nil
	s.Assert().Panics(func() {
		ldg.AddAccount(nil)
	})

	// Does not panic if given real account.
	acc := account.NewSavingsAcct("test", 10.0)
	err := ldg.AddAccount(acc)
	s.Require().NoError(err)
	s.Assert().Equal(1, len(ldg.accts))

	// Errors when given an account with a duplicate ID.
	acc = account.NewSavingsAcct("test", 1.0)
	err = ldg.AddAccount(acc)
	s.Assert().Error(err)
}

func (s *LedgerTestSuite) TestAddTransaction() {
	ldg := NewLedger()

	// Panics when credit id and debit id match
	s.Assert().Panics(func() {
		ldg.AddTransaction(time.Now(), "acc0", "acc0", 5)
	})

	// Panics when transaction amount is 0
	s.Assert().Panics(func() {
		ldg.AddTransaction(time.Now(), "acc0", "acc1", 0)
	})

	// Panics when transaction account ids aren't in account map.
	s.Assert().Panics(func() {
		ldg.AddTransaction(time.Now(), "acc0", "acc1", 5)
	})

	ldg.AddAccount(account.NewSavingsAcct("acc0", 10.0))

	s.Assert().Panics(func() {
		ldg.AddTransaction(time.Now(), "acc0", "acc1", 5)
	})

	ldg.AddAccount(account.NewExpenseAcct("acc1", 20.0))

	ldg.AddTransaction(time.Now(), "acc0", "acc1", 5)
	s.Assert().Equal(5.0, ldg.accts["acc0"].Balance())
	s.Assert().Equal(25.0, ldg.accts["acc1"].Balance())
	s.Assert().Equal(1, len(ldg.trns))
}

func (s *LedgerTestSuite) TestGetAccountBalance() {
	ldg := NewLedger()
	// Should return error when account ID is not in ledger
	amt, err := ldg.GetAccountBalance("acc0")
	s.Assert().Equal(0.0, amt)
	s.Assert().ErrorContains(err, "could not find account")

	ldg.AddAccount(account.NewSavingsAcct("acc0", 10.0))
	amt, err = ldg.GetAccountBalance("acc0")
	s.Assert().Equal(10.0, amt)
	s.Assert().NoError(err)
}
