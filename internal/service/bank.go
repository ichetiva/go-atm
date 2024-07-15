package service

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

type operation[T any] struct {
	id     string
	amount float64

	resp chan T
}

// Bank contains main logic for operations
type Bank struct {
	log      *slog.Logger
	accounts map[string]float64
	mu       sync.Mutex

	createOperations   chan operation[string]
	depositOperations  chan operation[error]
	withdrawOperations chan operation[error]
	getOperations      chan operation[float64]
}

func NewBank(log *slog.Logger) *Bank {
	return &Bank{
		log:      log,
		accounts: make(map[string]float64),
		mu:       sync.Mutex{},

		createOperations:   make(chan operation[string]),
		depositOperations:  make(chan operation[error]),
		withdrawOperations: make(chan operation[error]),
		getOperations:      make(chan operation[float64]),
	}
}

func (b *Bank) RunWorkers(count int) {
	for i := 0; i < count; i++ {
		go func() {
			for {
				select {
				case operation := <-b.createOperations:
					b.handleCreateOperation(operation)
				case operation := <-b.depositOperations:
					b.handleUpdateBalanceOperation(operation, 1)
				case operation := <-b.withdrawOperations:
					b.handleUpdateBalanceOperation(operation, -1)
				case operation := <-b.getOperations:
					b.handleGetBalanceOperation(operation)
				}
			}
		}()
	}
}

func (b *Bank) CreateAccount() string {
	op := operation[string]{
		resp: make(chan string, 1),
	}
	b.createOperations <- op
	return <-op.resp
}

func (b *Bank) GetAccount(id string) BankAccount {
	return NewAccount(id, b.depositOperations, b.withdrawOperations, b.getOperations)
}

func (b *Bank) Close() {
	close(b.createOperations)
	close(b.depositOperations)
	close(b.withdrawOperations)
	close(b.getOperations)
}

func (b *Bank) handleCreateOperation(op operation[string]) {
	defer close(op.resp)

	b.log.Debug("Create new account")
	id := uuid.New().String()
	b.createAccount(id)
	b.log.Debug("Create new account successful", "accountID", id)
	op.resp <- id
}

func (b *Bank) handleGetBalanceOperation(op operation[float64]) {
	defer close(op.resp)

	b.log.Debug("Get account balance", "accountID", op.id)
	balance, _ := b.getAccountBalance(op.id)
	b.log.Debug("Get account balance successful", "accountID", op.id, "balance", balance)
	op.resp <- balance
}

func (b *Bank) handleUpdateBalanceOperation(op operation[error], amountMultiplier float64) {
	defer close(op.resp)

	b.log.Debug("Update account balance", "accountID", op.id, "amount", op.amount)
	if _, ok := b.getAccountBalance(op.id); !ok {
		b.log.Debug("Update account balance failed", "reason", "account not found", "accountID", op.id, "amount", op.amount)
		op.resp <- fmt.Errorf("balance '%s' not found", op.id)
		return
	}
	b.updateAccountBalance(op.id, amountMultiplier*op.amount)
	b.log.Debug("Update account balance successful", "accountID", op.id, "amount", op.amount)
	op.resp <- nil
}

func (b *Bank) createAccount(id string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.accounts[id] = 0
}

func (b *Bank) getAccountBalance(id string) (float64, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	val, ok := b.accounts[id]
	return val, ok
}

func (b *Bank) updateAccountBalance(id string, amount float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.accounts[id] += amount
}
