package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	t.Run("MoneyBucket", func(t *testing.T) {
		t.Run("Asset", func(t *testing.T) {
			a := Asset{
				bal: 5.0,
			}
			assert.Equal(t, 5.0, a.Balance())
			assert.Equal(t, 3.5, a.Credit(1.5))
			assert.Equal(t, 4.5, a.Debit(1.0))
		})

		t.Run("Liability", func(t *testing.T) {
			l := Liability{
				bal: 5.0,
			}
			assert.Equal(t, 5.0, l.Balance())
			assert.Equal(t, 6.5, l.Credit(1.5))
			assert.Equal(t, 5.5, l.Debit(1.0))
		})
	})

	t.Run("NewAccount", func(t *testing.T) {
		t.Run("invalid account type", func(t *testing.T) {
			_, err := NewAccount("Test Account", AccountType(-1), 15.0)
			assert.ErrorIs(t, err, ErrInvalidAccountType)
		})

		t.Run("Assets", func(t *testing.T) {
			a, err := NewAccount("Test Account", Savings, 15.0)
			require.NoError(t, err)
			require.NotNil(t, a)
			assert.Equal(t, "Test Account", a.Name)
			assert.Equal(t, Savings, a.AccType)
			assert.IsType(t, &Asset{}, a.mb)

			assert.Equal(t, 15.0, a.Balance())
			assert.Equal(t, 13.5, a.Credit(1.5))
			assert.Equal(t, 14.0, a.Debit(0.5))
		})

		t.Run("Liability", func(t *testing.T) {
			a, err := NewAccount("Test Account", Expense, 15.0)
			require.NoError(t, err)
			require.NotNil(t, a)
			assert.Equal(t, "Test Account", a.Name)
			assert.Equal(t, Expense, a.AccType)
			assert.IsType(t, &Liability{}, a.mb)

			assert.Equal(t, 15.0, a.Balance())
			assert.Equal(t, 16.5, a.Credit(1.5))
			assert.Equal(t, 16.0, a.Debit(0.5))
		})
	})
}
