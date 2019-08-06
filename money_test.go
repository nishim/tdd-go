package money

import "testing"

func TestMultiplication(t *testing.T) {
	cases := []struct {
		times    int
		expected int
	}{
		{times: 2, expected: 10},
		{times: 3, expected: 15},
	}

	five := Dollar{5}
	for _, c := range cases {
		d := five.Times(c.times)
		if d.Amount != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, d.Amount)
		}
	}
}

func TestEquality(t *testing.T) {
	cases := []struct {
		a        Dollar
		b        Dollar
		expected bool
	}{
		{a: Dollar{5}, b: Dollar{5}, expected: true},
		{a: Dollar{5}, b: Dollar{6}, expected: false},
	}

	for _, c := range cases {
		if (c.a == c.b) != c.expected {
			t.Errorf("expected: %v, actual: %v", c.expected, (c.a == c.b))
		}
	}
}
