package pointer_errors

type Wallet struct {
	balance int
}

//struct pointers
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
