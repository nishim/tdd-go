package money

type pair struct {
	from, to string
}

type Bank struct {
	rates map[pair]int
}

func NewBank() *Bank {
	bank := Bank{rates: make(map[pair]int)}
	return &bank
}

func (b *Bank) Reduce(source Expression, currency string) Money {
	return source.Reduce(b, currency)
}

func (b *Bank) AddRate(from, to string, rate int) {
	pair := pair{from, to}
	b.rates[pair] = rate
}

func (b *Bank) rate(from, to string) int {
	if from == to {
		return 1
	}
	pair := pair{from, to}
	return b.rates[pair]
}

type Expression interface {
	Reduce(bank *Bank, to string) Money
}

type Sum struct {
	augend Money
	addend Money
}

func (s Sum) Reduce(bank *Bank, to string) Money {
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

func (m Money) Reduce(bank *Bank, to string) Money {
	rate := bank.rate(m.currency, to)
	return New(m.amount/rate, to)
}
