package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsettable(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		t.Run("error on never set value", func(t *testing.T) {
			s := NewUnsettableFloat64()
			_, err := s.Value()
			assert.ErrorIs(t, ErrUnsetValue, err)
		})

		t.Run("returns value after set", func(t *testing.T) {
			s := NewUnsettableFloat64()
			s.Set(64.5)
			v, err := s.Value()
			assert.NoError(t, err)
			assert.Equal(t, 64.5, v)
		})

		t.Run("error when getting an unset value", func(t *testing.T) {
			s := NewUnsettableFloat64()
			s.Set(64.5)
			v, err := s.Value()
			assert.NoError(t, err)
			assert.Equal(t, 64.5, v)
			s.Unset()
			_, err = s.Value()
			assert.ErrorIs(t, ErrUnsetValue, err)
		})
	})
}
