package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPeriodFilter(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		t.Run("Custom", func(t *testing.T) {
			t.Run("error when start after end time", func(t *testing.T) {
				_, err := NewCustomPeriodFilter(
					time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
					time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC),
				)
				assert.ErrorIs(t, err, ErrStartTimeAfterEnd)
			})

			t.Run("error when start and end time are equal", func(t *testing.T) {
				_, err := NewCustomPeriodFilter(
					time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
					time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				)
				assert.ErrorIs(t, err, ErrStartTimeEqualsEnd)
			})

			t.Run("happy path", func(t *testing.T) {
				fltr, err := NewCustomPeriodFilter(
					time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC),
					time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				)
				assert.NoError(t, err)
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})
		})

		t.Run("Week", func(t *testing.T) {
			t.Run("WeekAfter", func(t *testing.T) {
				fltr := NewWeekAfterPeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, 1, 8, 1, 1, 1, 1, time.UTC), fltr.End)
			})

			t.Run("WeekBefore", func(t *testing.T) {
				fltr := NewWeekBeforePeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1998, time.December, 25, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})
		})

		t.Run("BiWeek", func(t *testing.T) {
			t.Run("BiWeekAfter", func(t *testing.T) {
				fltr := NewBiWeekAfterPeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, 1, 15, 1, 1, 1, 1, time.UTC), fltr.End)
			})

			t.Run("BiWeekBefore", func(t *testing.T) {
				fltr := NewBiWeekBeforePeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1998, time.December, 18, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})
		})

		t.Run("Month", func(t *testing.T) {
			t.Run("MonthAfter", func(t *testing.T) {
				fltr := NewMonthAfterPeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, time.January, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, time.February, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})

			t.Run("MonthBefore", func(t *testing.T) {
				fltr := NewMonthBeforePeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1998, time.December, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, time.January, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})
		})

		t.Run("Quarter", func(t *testing.T) {
			t.Run("FirstQuarter", func(t *testing.T) {
				fltr := NewFirstQuarterPeriodFilter(1999)
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, time.April, 1, 0, 0, 0, 0, time.UTC), fltr.End)
			})

			t.Run("SecondQuarter", func(t *testing.T) {
				fltr := NewSecondQuarterPeriodFilter(1999)
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, time.April, 1, 0, 0, 0, 0, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, time.July, 1, 0, 0, 0, 0, time.UTC), fltr.End)
			})

			t.Run("ThirdQuarter", func(t *testing.T) {
				fltr := NewThirdQuarterPeriodFilter(1999)
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, time.July, 1, 0, 0, 0, 0, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, time.October, 1, 0, 0, 0, 0, time.UTC), fltr.End)
			})

			t.Run("FourthQuarter", func(t *testing.T) {
				fltr := NewFourthQuarterPeriodFilter(1999)
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, time.October, 1, 0, 0, 0, 0, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC), fltr.End)
			})
		})

		t.Run("Year", func(t *testing.T) {
			t.Run("YearAfter", func(t *testing.T) {
				fltr := NewYearAfterPeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1999, time.January, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})

			t.Run("YearBefore", func(t *testing.T) {
				fltr := NewYearBeforePeriodFilter(time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC))
				require.NotNil(t, fltr)
				assert.Equal(t, time.Date(1998, time.January, 1, 1, 1, 1, 1, time.UTC), fltr.Start)
				assert.Equal(t, time.Date(1999, time.January, 1, 1, 1, 1, 1, time.UTC), fltr.End)
			})
		})
	})

	type testAccounts struct {
		credit *Account
		debit  *Account
	}

	var setupTestAccounts = func() testAccounts {
		credit, err := NewAccount("Savings", Savings, 1234)
		if err != nil {
			panic(fmt.Errorf("got an error: %s", err.Error()))
		}
		debit, err := NewAccount("MoneySink", Expense, 9999.99)
		if err != nil {
			panic(fmt.Errorf("got an error: %s", err.Error()))
		}

		return testAccounts{
			credit: credit,
			debit:  debit,
		}
	}

	t.Run("Filter", func(t *testing.T) {
		testAccs := setupTestAccounts()

		f, err := NewCustomPeriodFilter(
			time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC),
			time.Date(1999, time.January, 1, 0, 1, 0, 0, time.UTC),
		)
		require.NoError(t, err)

		testTran, err := NewTransaction(
			time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC),
			12.24,
			testAccs.debit,
			testAccs.credit,
			"test",
		)
		require.NoError(t, err)

		assert.True(t, f.Filter(testTran))

		testTran, err = NewTransaction(
			time.Date(1999, time.January, 1, 0, 0, 5, 0, time.UTC),
			12.24,
			testAccs.debit,
			testAccs.credit,
			"test",
		)
		require.NoError(t, err)

		assert.True(t, f.Filter(testTran))

		testTran, err = NewTransaction(
			time.Date(1999, time.January, 1, 0, 1, 0, 0, time.UTC),
			1.23,
			testAccs.debit,
			testAccs.credit,
			"test",
		)
		require.NoError(t, err)

		assert.False(t, f.Filter(testTran))

		testTran, err = NewTransaction(
			time.Date(1999, time.January, 1, 0, 2, 0, 0, time.UTC),
			1.23,
			testAccs.debit,
			testAccs.credit,
			"test",
		)
		require.NoError(t, err)

		assert.False(t, f.Filter(testTran))
	})
}
