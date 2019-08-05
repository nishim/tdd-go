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
