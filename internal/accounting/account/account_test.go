package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSavingsAcct(t *testing.T) {
	t.Run("TestCredit", func(t *testing.T) {
		sact := NewSavingsAcct("test", 10.0)
		err := sact.Credit(5.0)
		assert.NoError(t, err)
		assert.Equal(t, 5.0, sact.Balance())
	})

	t.Run("TestDedit", func(t *testing.T) {
		sact := NewSavingsAcct("test", 10.0)
		err := sact.Debit(5.0)
		assert.NoError(t, err)
		assert.Equal(t, 15.0, sact.Balance())
	})
}

func TestExpenseAcct(t *testing.T) {
	t.Run("TestCredit", func(t *testing.T) {
		eact := NewExpenseAcct("test", 10.0)
		err := eact.Credit(5.0)
		assert.NoError(t, err)
		assert.Equal(t, 5.0, eact.Balance())
	})

	t.Run("TestDedit", func(t *testing.T) {
		eact := NewExpenseAcct("test", 10.0)
		err := eact.Debit(5.0)
		assert.NoError(t, err)
		assert.Equal(t, 15.0, eact.Balance())
	})
}
