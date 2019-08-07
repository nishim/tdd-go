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

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount}
}

func (d Franc) Times(t int) Franc {
	return Franc{d.amount * t}
}
