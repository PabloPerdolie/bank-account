package models

import (
	"errors"
	"log"
	"sync"
)

type account struct {
	ID         int
	balance    float64
	mu         sync.Mutex
	operations chan operation
}

type operation struct {
	amount float64
	opType string
	result chan result
}

type result struct {
	err     error
	balance float64
}

func NewAccount(id int) BankAccount {
	a := &account{
		ID:         id,
		balance:    0,
		operations: make(chan operation),
	}
	go a.processOperations()
	return a
}

func (a *account) processOperations() {
	for op := range a.operations {
		switch op.opType {
		case "deposit":
			op.result <- a.deposit(op.amount)
		case "withdraw":
			op.result <- a.withdraw(op.amount)
		case "getBalance":
			op.result <- result{balance: a.balance}
		}
	}
}

func (a *account) deposit(amount float64) result {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return result{err: errors.New("deposit amount must be positive")}
	}
	a.balance += amount
	logOperation("Deposit", a.ID, amount)
	return result{}
}

func (a *account) withdraw(amount float64) result {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount <= 0 {
		return result{err: errors.New("withdraw amount must be positive")}
	}
	if amount > a.balance {
		return result{err: errors.New("insufficient funds")}
	}
	a.balance -= amount
	logOperation("Withdraw", a.ID, amount)
	return result{}
}

func (a *account) Deposit(amount float64) error {
	resultChan := make(chan result)
	a.operations <- operation{amount: amount, opType: "deposit", result: resultChan}
	return (<-resultChan).err
}

func (a *account) Withdraw(amount float64) error {
	resultChan := make(chan result)
	a.operations <- operation{amount: amount, opType: "withdraw", result: resultChan}
	return (<-resultChan).err
}

func (a *account) GetBalance() float64 {
	resultChan := make(chan result)
	a.operations <- operation{opType: "getBalance", result: resultChan}
	return (<-resultChan).balance
}

func logOperation(opType string, accountID int, amount float64) {
	log.Printf("Operation: %s, Account ID: %d, Amount: %.2f\n", opType, accountID, amount)
}
