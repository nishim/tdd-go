package money

import "testing"

func TestMultiplication(t *testing.T) {
	cases := []struct {
		dollar   int
		times    int
		expected int
	}{
		{dollar: 5, times: 2, expected: 10},
	}

	for _, c := range cases {
		d := Dollar{c.dollar}
		d.Times{c.times}
		if d.Amount != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, d.Amount)
		}
	}
}
