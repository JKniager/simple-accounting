package dbmngr

import (
	"context"
	"os"
	"path/filepath"
	"simple_accounting/internal/accounting/account"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDbManager(t *testing.T) {
	t.Run("non-existent database", func(t *testing.T) {
		_, err := NewDatabaseManager(context.Background(), "fake.db")
		assert.ErrorContains(t, err, "failed to open database")
	})

	t.Run("opens database", func(t *testing.T) {
		db, err := NewDatabaseManager(context.Background(), filepath.Join("testdata", "empty.db"))
		require.NoError(t, err)
		assert.NotNil(t, db)
		assert.NotNil(t, db.db)
	})

	t.Run("database creation", func(t *testing.T) {
		tDir := t.TempDir()
		dbPath := filepath.Join(tDir, "test.db")
		_, err := CreateDatabaseAndGetManager(context.Background(), dbPath)
		require.NoError(t, err)
		_, err = os.Stat(dbPath)
		assert.NoError(t, err)
		_, err = CreateDatabaseAndGetManager(context.Background(), dbPath)
		assert.ErrorContains(t, err, "database already exists")

	})

	t.Run("account access", func(t *testing.T) {
		t.Run("empty account table", func(t *testing.T) {
			tDir := t.TempDir()
			dbPath := filepath.Join(tDir, "test.db")
			dm, err := CreateDatabaseAndGetManager(context.Background(), dbPath)
			require.NoError(t, err)
			require.NotNil(t, dm)
			accList, err := dm.GetAccountList()
			assert.NoError(t, err)
			assert.Empty(t, accList)
		})

		t.Run("add account", func(t *testing.T) {
			tDir := t.TempDir()
			dbPath := filepath.Join(tDir, "test.db")
			dm, err := CreateDatabaseAndGetManager(context.Background(), dbPath)
			require.NoError(t, err)
			require.NotNil(t, dm)

			id, err := dm.AddAccount(account.AccountTypeSavings, "MySavings", 64.0)
			assert.NoError(t, err)
			assert.NotEqual(t, -1, id)

			accList, err := dm.GetAccountList()
			assert.NoError(t, err)
			assert.Equal(t, 1, len(accList))
			acc := accList[0]
			assert.Equal(t, id, acc.Id())
			_, ok := acc.(*account.SavingsAcct)
			assert.True(t, ok)
			assert.Equal(t, "MySavings", acc.Name())
			assert.Equal(t, 64.0, acc.Balance())
		})
	})
}
