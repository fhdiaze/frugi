package types

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type Money decimal.Decimal

func MoneyFromMajor(amount float64) *Money {
	m := Money(decimal.NewFromFloat(amount))
	return &m
}

func (m *Money) Div(d float64) *Money {
	n := Money(decimal.Decimal(*m).Div(decimal.NewFromFloat(d)))
	return &n
}

func (m *Money) Mul(c float64) *Money {
	n := Money(decimal.Decimal(*m).Mul(decimal.NewFromFloat(c)))
	return &n
}

func (m *Money) Add(c *Money) *Money {
	n := Money(decimal.Decimal(*m).Add(decimal.Decimal(*c)))
	return &n
}

func (m *Money) ToFloat64() float64 {
	t, _ := decimal.Decimal(*m).Float64()
	return t
}

func (m *Money) Major() float64 {
	major, _ := decimal.Decimal(*m).Round(2).Float64()

	return major
}

func (m *Money) Minor() int64 {
	return decimal.Decimal(*m).Mul(decimal.NewFromInt(100)).IntPart()
}

func (m *Money) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return fmt.Errorf("invalid money amount: %w", err)
	}

	*m = *MoneyFromMajor(v)

	return nil
}

func (m *Money) MajorPart() int64 {
	return decimal.Decimal(*m).IntPart()
}

func (m *Money) MinorPart() int8 {
	minor := decimal.Decimal(*m).Round(2).Mul(decimal.NewFromInt(100)).IntPart() % 100

	return int8(minor)
}
