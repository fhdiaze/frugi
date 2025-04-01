package types

import (
	"github.com/shopspring/decimal"
)

type Money decimal.Decimal

func MoneyFromMajor(amount float64) Money {
	return Money(decimal.NewFromFloat(amount))
}

func (m Money) Div(d float64) Money {
	return Money(decimal.Decimal(m).Div(decimal.NewFromFloat(d)))
}

func (m Money) ToFloat64() float64 {
	t, _ := decimal.Decimal(m).Float64()
	return t
}

func (m Money) Minor() int64 {
	return decimal.Decimal(m).Mul(decimal.NewFromInt(100)).IntPart()
}
