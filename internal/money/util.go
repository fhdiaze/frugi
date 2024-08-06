package money

import (
	"math"
)

type Money int64

func FromMajor(amount float64) Money {
	return Money(math.Round(amount * 100))
}
