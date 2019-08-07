package money

import "testing"

func TestMultiplication(t *testing.T) {
	cases := []struct {
		times    int
		expected Dollar
	}{
		{times: 2, expected: NewDollar(10)},
		{times: 3, expected: NewDollar(15)},
	}

	five := NewDollar(5)
	for _, c := range cases {
		d := five.Times(c.times)
		if d != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, d)
		}
	}
}

func TestEquality(t *testing.T) {
	cases := []struct {
		a        Dollar
		b        Dollar
		expected bool
	}{
		{a: NewDollar(5), b: NewDollar(5), expected: true},
		{a: NewDollar(5), b: NewDollar(6), expected: false},
	}

	for _, c := range cases {
		if (c.a == c.b) != c.expected {
			t.Errorf("expected: %v, actual: %v", c.expected, (c.a == c.b))
		}
	}
}

func TestFrancMultiplication(t *testing.T) {
	cases := []struct {
		times    int
		expected Franc
	}{
		{times: 2, expected: NewFranc(10)},
		{times: 3, expected: NewFranc(15)},
	}

	five := NewFranc(5)
	for _, c := range cases {
		d := five.Times(c.times)
		if d != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, d)
		}
	}
}
