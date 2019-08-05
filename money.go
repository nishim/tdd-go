package money

type Dollar struct {
	Amount int
}

func (d *Dollar) Times(t int) {
	d.Amount = d.Amount * t
}
