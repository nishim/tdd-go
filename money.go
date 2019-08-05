package money

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(t int) Dollar {
	return Dollar{d.Amount * t}
}
