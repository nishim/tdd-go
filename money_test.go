package money

import "testing"

func TestMultiplication(t *testing.T) {
	cases := []struct {
		times    int
		expected Money
	}{
		{times: 2, expected: NewDollar(10)},
		{times: 3, expected: NewDollar(15)},
	}

	five := NewDollar(5)
	for _, c := range cases {
		d := five.Times(c.times)
		if d != c.expected {
			t.Errorf("expected: %v, actual: %v", c.expected, d)
		}
	}
}

func TestEquality(t *testing.T) {
	cases := []struct {
		a        Money
		b        Money
		expected bool
	}{
		{a: NewDollar(5), b: NewDollar(5), expected: true},
		{a: NewDollar(5), b: NewDollar(6), expected: false},
		{a: NewDollar(5), b: NewFranc(5), expected: false},
	}

	for _, c := range cases {
		b := Equals(c.a, c.b)
		if b != c.expected {
			t.Errorf("expected: %v, actual: %v, a:%v, b:%v", c.expected, b, c.a, c.b)
		}
	}
}

func TestCurrenvy(t *testing.T) {
	cases := []struct {
		currency string
		money    Money
	}{
		{currency: "USD", money: NewDollar(1)},
		{currency: "CHF", money: NewFranc(1)},
	}

	for _, c := range cases {
		if c.currency != c.money.Currency() {
			t.Errorf("expected: %s, actual: %s", c.currency, c.money.Currency())
		}
	}
}
