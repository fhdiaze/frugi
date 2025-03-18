package types

import (
	"math"
)

type Money int64

func MoneyFromMajor(amount float64) Money {
	return Money(math.Round(amount * 100))
}
