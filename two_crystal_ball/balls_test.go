package twocrystalball

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrystalBallsSearch(t *testing.T) {
	t.Run("Should return index where crystal balls break", func(t *testing.T) {
		breaks := []bool{false, false, false, false, false, true, true, true, true, true}
		index := Search(breaks)
		assert.Equal(t, index, 5)
	})

	t.Run("Should return -1 when crystall balls don't break", func(t *testing.T) {
		breaks := []bool{false, false, false, false, false, false, false, false, false, false}
		index := Search(breaks)
		assert.Equal(t, index, -1)
	})
}
