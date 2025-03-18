package money

import (
	"testing"

	"github.com/fhdiaze/frugi/internal/pkg/types"
)

func TestMoneyFromMajor(t *testing.T) {
	// Arrange
	amount := 5.3399333
	wanted := types.MoneyFromMajor(amount)

	// Act
	got := types.MoneyFromMajor(amount)

	// Assert
	if got != 534 {
		t.Errorf("Output %q not equal to the expected %q", got, wanted)
	}
}
