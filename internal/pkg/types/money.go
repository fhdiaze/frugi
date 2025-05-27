package types

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type Money decimal.Decimal

func MoneyFromMajor(amount float64) Money {
	return Money(decimal.NewFromFloat(amount))
}

func (m Money) Div(d float64) Money {
	return Money(decimal.Decimal(m).Div(decimal.NewFromFloat(d)))
}

func (m Money) Mul(c float64) Money {
	return Money(decimal.Decimal(m).Mul(decimal.NewFromFloat(c)))
}

func (m Money) Add(c Money) Money {
	return Money(decimal.Decimal(m).Add(decimal.Decimal(c)))
}

func (m Money) ToFloat64() float64 {
	t, _ := decimal.Decimal(m).Float64()
	return t
}

func (m Money) Minor() int64 {
	return decimal.Decimal(m).Mul(decimal.NewFromInt(100)).IntPart()
}

func (m *Money) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return fmt.Errorf("invalid money amount: %w", err)
	}

	*m = MoneyFromMajor(v)

	return nil
}
