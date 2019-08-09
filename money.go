package money

type Money struct {
	amount   int
	currency string
}

func NewDollar(amount int) Money {
	return Money{amount, "USD"}
}

func NewFranc(amount int) Money {
	return Money{amount, "CHF"}
}

func Equals(a, b Money) bool {
	return a.amount == b.amount && a.currency == b.currency
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
