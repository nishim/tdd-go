package money

import "reflect"

type Bank struct{}

func (b *Bank) Reduce(source Expression, currency string) Money {
	if reflect.TypeOf(source).Name() == "Money" {
		return source.(Money)
	}
	sum := source.(Sum)
	return sum.Reduce(currency)
}

type Expression interface{}

type Sum struct {
	augend Money
	addend Money
}

func (s *Sum) Reduce(to string) Money {
	amount := s.augend.GetAmount() + s.addend.GetAmount()
	return New(amount, to)
}

type Money struct {
	amount   int
	currency string
}

func New(amount int, currency string) Money {
	return Money{amount, currency}
}

func NewDollar(amount int) Money {
	return Money{amount, "USD"}
}

func NewFranc(amount int) Money {
	return Money{amount, "CHF"}
}

func (m Money) Equals(a Money) bool {
	return m.amount == a.amount && m.currency == a.currency
}

func (m Money) Times(t int) Money {
	return Money{m.amount * t, m.currency}
}

func (m Money) GetAmount() int {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Plus(a Money) Expression {
	return Sum{m, a}
}
