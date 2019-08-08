package money

import "reflect"

type Money interface {
	GetAmount() int
	Times(int) Money
}

func Equals(a, b Money) bool {
	return a.GetAmount() == b.GetAmount() &&
		reflect.TypeOf(a) == reflect.TypeOf(b)
}

type dollar struct {
	amount int
}

func NewDollar(amount int) Money {
	return dollar{amount}
}

func (d dollar) Times(t int) Money {
	return dollar{d.amount * t}
}

func (d dollar) GetAmount() int {
	return d.amount
}

// Franc.
type franc struct {
	amount int
}

func NewFranc(amount int) Money {
	return franc{amount}
}

func (d franc) Times(t int) Money {
	return franc{d.amount * t}
}

func (d franc) GetAmount() int {
	return d.amount
}
