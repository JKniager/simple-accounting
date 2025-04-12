package service

import (
	"errors"
	"time"
)

var (
	// Errors
	ErrStartTimeAfterEnd  = errors.New("cannot set the start time after the end time")
	ErrStartTimeEqualsEnd = errors.New("cannot set the start and end times equal to each other")
)

type PeriodFilter struct {
	Start time.Time
	End   time.Time
}

func NewCustomPeriodFilter(start, end time.Time) (*PeriodFilter, error) {
	if start.After(end) {
		return nil, ErrStartTimeAfterEnd
	}

	if start == end {
		return nil, ErrStartTimeEqualsEnd
	}

	return &PeriodFilter{
		Start: start,
		End:   end,
	}, nil
}

func NewWeekAfterPeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t,
		End:   t.AddDate(0, 0, 7),
	}
}

func NewBiWeekBeforePeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t.AddDate(0, 0, -14),
		End:   t,
	}
}

func NewBiWeekAfterPeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t,
		End:   t.AddDate(0, 0, 14),
	}
}

func NewWeekBeforePeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t.AddDate(0, 0, -7),
		End:   t,
	}
}

func NewMonthAfterPeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t,
		End:   t.AddDate(0, 1, 0),
	}
}

func NewMonthBeforePeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t.AddDate(0, -1, 0),
		End:   t,
	}
}

func NewFirstQuarterPeriodFilter(year int) *PeriodFilter {
	return &PeriodFilter{
		Start: time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewSecondQuarterPeriodFilter(year int) *PeriodFilter {
	return &PeriodFilter{
		Start: time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewThirdQuarterPeriodFilter(year int) *PeriodFilter {
	return &PeriodFilter{
		Start: time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewFourthQuarterPeriodFilter(year int) *PeriodFilter {
	return &PeriodFilter{
		Start: time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewYearAfterPeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t,
		End:   t.AddDate(1, 0, 0),
	}
}

func NewYearBeforePeriodFilter(t time.Time) *PeriodFilter {
	return &PeriodFilter{
		Start: t.AddDate(-1, 0, 0),
		End:   t,
	}
}
