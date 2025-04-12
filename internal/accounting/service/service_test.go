package service

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountingService(t *testing.T) {
	t.Run("CreateAccount", func(t *testing.T) {
		t.Run("returns error", func(t *testing.T) {
			testErr := errors.New("I are error")
			ml := &mockLedger{}
			ml.On("CreateAccount",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("AccountType"),
				mock.AnythingOfType("float64"),
			).Return(testErr)
			s := NewAccountingService(ml)
			err := s.CreateAccount("Test", Savings, 0.0)
			assert.ErrorIs(t, err, testErr)

			ml.AssertExpectations(t)
		})

		t.Run("passes arguments", func(t *testing.T) {
			ml := &mockLedger{}
			ml.On("CreateAccount",
				"Apples",
				Savings,
				2.0,
			).Return(nil).Once()
			ml.On("CreateAccount",
				"Oranges",
				Expense,
				0.0,
			).Return(nil).Once()
			ml.On("CreateAccount",
				"Pears",
				Savings,
				0.5,
			).Return(nil).Once()

			s := NewAccountingService(ml)
			assert.NoError(t, s.CreateAccount("Apples", Savings, 2.0))
			assert.NoError(t, s.CreateAccount("Oranges", Expense, 0.0))
			assert.NoError(t, s.CreateAccount("Pears", Savings, 0.5))

			ml.AssertExpectations(t)
		})
	})

	t.Run("AddTransaction", func(t *testing.T) {
		t.Run("returns error", func(t *testing.T) {
			testErr := errors.New("transaction terminated")
			testTime := time.Now()
			testCrd := &Account{}
			testDbt := &Account{}
			ml := &mockLedger{}
			ml.On("AddTransaction",
				testTime,
				mock.AnythingOfType("float64"),
				testCrd,
				testDbt,
				mock.AnythingOfType("string"),
			).Return(testErr)
			s := NewAccountingService(ml)
			err := s.AddTransaction(testTime, 23.0, testCrd, testDbt, "")
			assert.ErrorIs(t, err, testErr)

			ml.AssertExpectations(t)
		})
	})

	t.Run("GetAccountSummaries", func(t *testing.T) {
		t.Run("empty account list", func(t *testing.T) {
			ml := &mockLedger{}
			ml.On("GetAccounts", AccountFilter{}).Return([]*Account{})
			s := NewAccountingService(ml)
			res := s.GetAccountSummaries(AccountFilter{})
			assert.Empty(t, res)

			ml.AssertExpectations(t)
		})

		t.Run("list of accounts", func(t *testing.T) {
			mmb1 := &mockMoneyBucket{}
			mmb1.On("Balance").Return(25.34)
			ma1 := &Account{
				Name:    "Test 1",
				AccType: Savings,
				mb:      mmb1,
			}

			mmb2 := &mockMoneyBucket{}
			mmb2.On("Balance").Return(0.67)
			ma2 := &Account{
				Name:    "Test 2",
				AccType: Expense,
				mb:      mmb2,
			}

			ml := &mockLedger{}
			ml.On("GetAccounts", AccountFilter{}).Return([]*Account{
				ma1,
				ma2,
			})
			s := NewAccountingService(ml)
			res := s.GetAccountSummaries(AccountFilter{})
			assert.Equal(t, []AccountSummary{
				{
					Name:    "Test 1",
					AccType: "Savings",
					Bal:     25.34,
				},
				{
					Name:    "Test 2",
					AccType: "Expense",
					Bal:     0.67,
				},
			}, res)

			ml.AssertExpectations(t)
			mmb1.AssertExpectations(t)
			mmb2.AssertExpectations(t)
		})

		t.Run("filter", func(t *testing.T) {
			mmb := &mockMoneyBucket{}
			mmb.On("Balance").Return(25.34)
			ma := &Account{
				Name:    "Test 1",
				AccType: Savings,
				mb:      mmb,
			}

			testFilter := AccountFilter{
				AccType: SavingsFilter,
			}

			ml := &mockLedger{}
			ml.On("GetAccounts", testFilter).Return([]*Account{
				ma,
			})
			s := NewAccountingService(ml)
			res := s.GetAccountSummaries(AccountFilter{
				AccType: SavingsFilter,
			})
			assert.Equal(t, []AccountSummary{
				{
					Name:    "Test 1",
					AccType: "Savings",
					Bal:     25.34,
				},
			}, res)

			ml.AssertExpectations(t)
			mmb.AssertExpectations(t)
		})
	})
}

type mockLedger struct {
	mock.Mock
}

func (m *mockLedger) GetAccounts(fltr AccountFilter) []*Account {
	return m.Called(fltr).Get(0).([]*Account)
}

func (m *mockLedger) CreateAccount(name string, accT AccountType, startBal float64) error {
	return m.Called(name, accT, startBal).Error(0)
}

func (m *mockLedger) GetTransactions() []*Transaction {
	return m.Called().Get(0).([]*Transaction)
}

func (m *mockLedger) AddTransaction(date time.Time, amt float64, creditAcc, debitAcc *Account, comment string) error {
	return m.Called(date, amt, creditAcc, debitAcc, comment).Error(0)
}

type mockMoneyBucket struct {
	mock.Mock
}

func (m *mockMoneyBucket) Balance() float64 {
	return m.Called().Get(0).(float64)
}

func (m *mockMoneyBucket) Credit(amt float64) float64 {
	return m.Called(amt).Get(0).(float64)
}

func (m *mockMoneyBucket) Debit(amt float64) float64 {
	return m.Called(amt).Get(0).(float64)
}
