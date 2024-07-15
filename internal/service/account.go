package service

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	id string

	depositOperations  chan operation[error]
	withdrawOperations chan operation[error]
	getOperations      chan operation[float64]
}

func NewAccount(
	id string,
	depositOperations chan operation[error],
	withdrawOperations chan operation[error],
	getOperations chan operation[float64],
) BankAccount {
	return &Account{
		id: id,

		depositOperations:  depositOperations,
		withdrawOperations: withdrawOperations,
		getOperations:      getOperations,
	}
}

func (a *Account) Deposit(amount float64) error {
	op := operation[error]{
		id:     a.id,
		amount: amount,
		resp:   make(chan error, 1),
	}
	a.depositOperations <- op
	return <-op.resp
}

func (a *Account) Withdraw(amount float64) error {
	op := operation[error]{
		id:     a.id,
		amount: amount,
		resp:   make(chan error, 1),
	}
	a.withdrawOperations <- op
	return <-op.resp
}

func (a *Account) GetBalance() float64 {
	op := operation[float64]{
		id:   a.id,
		resp: make(chan float64, 1),
	}
	a.getOperations <- op
	return <-op.resp
}
