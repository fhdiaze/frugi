package money

import (
	"testing"

	"github.com/fhdiaze/frugi/internal/money"
)

func TestFromMajor(t *testing.T) {
	// Arrange
	amount := 5.3399333
	wanted := money.Money(amount)

	// Act
	got := money.FromMajor(amount)

	// Assert
	if got != 534 {
		t.Errorf("Output %q not equal to the expected %q", got, wanted)
	}
}
