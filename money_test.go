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
		b := c.a.Equals(c.b)
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

func TestSimpleAdditiom(t *testing.T) {
	cases := []struct {
		m1       Money
		m2       Money
		expected Money
	}{
		{m1: NewDollar(5), m2: NewDollar(5), expected: NewDollar(10)},
	}

	bank := NewBank()
	for _, c := range cases {
		sum := c.m1.Plus(c.m2)
		reduced := bank.Reduce(sum, "USD")
		if !reduced.Equals(c.expected) {
			t.Errorf("expected: %v, actual: %v", c.expected, sum)
		}
	}
}

func TestIdentityRate(t *testing.T) {
	bank := NewBank()
	rate := bank.rate("USD", "USD")
	if 1 != rate {
		t.Errorf("expected: %d, actual: %d", 1, rate)
	}
}

func TestReduceSum(t *testing.T) {
	sum := Sum{NewDollar(3), NewDollar(4)}
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	seven := NewDollar(7)
	if !seven.Equals(result) {
		t.Errorf("expected: %v, actual:%v", seven, result)
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	one := NewDollar(1)
	if !one.Equals(result) {
		t.Errorf("expected: %v, actual:%v", one, result)
	}
}

func TestReduceMoneyDifferencCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	oned := NewDollar(1)

	if !oned.Equals(result) {
		t.Errorf("expected: %v, actual:%v", oned, result)
	}
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrance := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrance), "USD")

	tenBucks := NewDollar(10)

	if !result.Equals(tenBucks) {
		t.Errorf("expected: %v, actual: %v", tenBucks, result)
	}
}
