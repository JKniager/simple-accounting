package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	t.Run("NewTransaction", func(t *testing.T) {
		t.Run("nil credit", func(t *testing.T) {
			_, err := NewTransaction(time.Now(), 10.0, nil, &Account{}, "")
			assert.ErrorIs(t, err, ErrNilCreditAccount)
		})

		t.Run("nil debit", func(t *testing.T) {
			_, err := NewTransaction(time.Now(), 10.0, &Account{}, nil, "")
			assert.ErrorIs(t, err, ErrNilDebitAccount)
		})

		t.Run("zero amount", func(t *testing.T) {
			_, err := NewTransaction(time.Now(), 0.0, &Account{}, &Account{}, "")
			assert.ErrorIs(t, err, ErrZeroTransactionAmount)
		})

		t.Run("happy path", func(t *testing.T) {
			ca := &Account{
				mb: &Asset{
					bal: 35.0,
				},
			}
			da := &Account{
				mb: &Liability{
					bal: 25.0,
				},
			}
			testDate := time.Now()
			tr, err := NewTransaction(testDate, 10.0, ca, da, "I'm a Comment")
			assert.NoError(t, err)
			require.NotNil(t, tr)
			assert.Equal(t, testDate, tr.Date)
			assert.Equal(t, 10.0, tr.Amount)
			assert.Equal(t, ca, tr.CreditAccInfo.Acc)
			assert.False(t, tr.CreditAccInfo.OldBalance.IsSet())
			assert.False(t, tr.CreditAccInfo.NewBalance.IsSet())
			assert.Equal(t, da, tr.DebitAccInfo.Acc)
			assert.False(t, tr.DebitAccInfo.OldBalance.IsSet())
			assert.False(t, tr.DebitAccInfo.NewBalance.IsSet())
		})
	})

	t.Run("Apply;Undo", func(t *testing.T) {
		t.Run("already applied", func(t *testing.T) {
			tr := &Transaction{
				applied: true,
			}
			assert.ErrorIs(t, tr.Apply(), ErrTransactionAlreadyApplied)
		})

		t.Run("already undone", func(t *testing.T) {
			tr := &Transaction{}
			assert.ErrorIs(t, tr.Undo(), ErrTransactionNotApplied)
		})

		t.Run("full test", func(t *testing.T) {
			ca := &Account{
				mb: &Asset{
					bal: 35.0,
				},
			}
			da := &Account{
				mb: &Liability{
					bal: 25.0,
				},
			}
			testDate := time.Now()
			tr, err := NewTransaction(testDate, 10.0, ca, da, "I'm a Comment")
			require.NoError(t, err)
			require.NotNil(t, tr)

			assert.False(t, tr.CreditAccInfo.NewBalance.IsSet())
			assert.False(t, tr.CreditAccInfo.OldBalance.IsSet())
			assert.False(t, tr.DebitAccInfo.NewBalance.IsSet())
			assert.False(t, tr.DebitAccInfo.OldBalance.IsSet())

			assert.NoError(t, tr.Apply())
			assert.Equal(t, 25.0, ca.Balance())
			assert.Equal(t, 15.0, da.Balance())
			v, err := tr.CreditAccInfo.OldBalance.Value()
			assert.NoError(t, err)
			assert.Equal(t, 35.0, v)
			v, err = tr.CreditAccInfo.NewBalance.Value()
			assert.NoError(t, err)
			assert.Equal(t, 25.0, v)
			v, err = tr.DebitAccInfo.OldBalance.Value()
			assert.NoError(t, err)
			assert.Equal(t, 25.0, v)
			v, err = tr.DebitAccInfo.NewBalance.Value()
			assert.NoError(t, err)
			assert.Equal(t, 15.0, v)
			assert.ErrorIs(t, tr.Apply(), ErrTransactionAlreadyApplied)

			assert.NoError(t, tr.Undo())
			assert.Equal(t, 35.0, ca.Balance())
			assert.Equal(t, 25.0, da.Balance())
			assert.False(t, tr.CreditAccInfo.NewBalance.IsSet())
			assert.False(t, tr.CreditAccInfo.OldBalance.IsSet())
			assert.False(t, tr.DebitAccInfo.NewBalance.IsSet())
			assert.False(t, tr.DebitAccInfo.OldBalance.IsSet())
			assert.ErrorIs(t, tr.Undo(), ErrTransactionNotApplied)
		})
	})
}
