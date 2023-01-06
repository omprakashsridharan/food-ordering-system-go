package valueobject

import "github.com/shopspring/decimal"

type Money struct {
	Amount decimal.Decimal
}

func NewMoney(amount float64) *Money {
	return &Money{Amount: decimal.NewFromFloat(amount).Round(2)}
}

func NewMoneyFromDecimal(amount decimal.Decimal) *Money {
	return &Money{Amount: amount.Round(2)}
}

func (m *Money) GreaterThanZero() bool {
	return m.Amount.GreaterThan(decimal.Zero)
}

func (m *Money) GreaterThan(otherMoney *Money) bool {
	return m.Amount.GreaterThan(otherMoney.Amount)
}

func (m *Money) Add(otherMoney *Money) *Money {
	result := m.Amount.Add(otherMoney.Amount)
	return NewMoneyFromDecimal(result)
}

func (m *Money) Subtract(otherMoney *Money) *Money {
	result := m.Amount.Sub(otherMoney.Amount)
	return NewMoneyFromDecimal(result)
}

func (m *Money) Multiple(multiplier int64) *Money {
	result := m.Amount.Mul(decimal.NewFromInt(multiplier))
	return NewMoneyFromDecimal(result)
}
