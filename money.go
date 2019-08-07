package money

type Money interface {
	GetAmount() int
}

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount}
}

func (d Dollar) Times(t int) Dollar {
	return Dollar{d.amount * t}
}

func (d Dollar) GetAmount() int {
	return d.amount
}

// Franc.
type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount}
}

func (d Franc) Times(t int) Franc {
	return Franc{d.amount * t}
}

func (d Franc) GetAmount() int {
	return d.amount
}
