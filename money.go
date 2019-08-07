package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount}
}

func (d Dollar) Times(t int) Dollar {
	return Dollar{d.amount * t}
}
