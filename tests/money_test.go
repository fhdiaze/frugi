package money

import (
	"testing"

	"github.com/fhdiaze/frugi/internal/pkg/types"
)

func TestMoneyFromMajor(t *testing.T) {
	// Arrange
	amount := 5.3399333
	var wanted int64 = 534

	// Act
	got := types.MoneyFromMajor(amount)

	// Assert
	if got.Minor() != wanted {
		t.Errorf("Output %q not equal to the expected %q", got.Minor(), wanted)
	}
}
