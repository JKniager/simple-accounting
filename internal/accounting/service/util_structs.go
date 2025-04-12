package service

import "errors"

var (
	ErrUnsetValue = errors.New("cannot access an unset value")
)

type Unsettable[T any] struct {
	isSet bool
	val   T
}

func (u *Unsettable[T]) IsSet() bool {
	return u.isSet
}

func (u *Unsettable[T]) Set(v T) {
	u.val = v
	u.isSet = true
}

func (u *Unsettable[T]) Unset() {
	u.isSet = false
}

func (u *Unsettable[T]) Value() (T, error) {
	if u.isSet {
		return u.val, nil
	}
	return u.val, ErrUnsetValue
}

type UnsettableFloat64 = Unsettable[float64]

func NewUnsettableFloat64() *UnsettableFloat64 {
	return &UnsettableFloat64{}
}
