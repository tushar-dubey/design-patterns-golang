package main

import "fmt"

var overdraftLimit int = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited ", amount,
		"\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount > overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew ", amount,
			"\b, balance is now", b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	BankAccount *BankAccount
	Action      Action
	Amount      int
	succeeded   bool
}

func (b *BankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}

func (b *BankAccountCommand) Undo() {
	switch b.Action {
	case Withdraw:
		if b.succeeded == true {
			b.BankAccount.Deposit(b.Amount)
		}
	case Deposit:
		b.BankAccount.Withdraw(b.Amount)
	}
}

func NewBankAccountCommand(bankAccount *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{BankAccount: bankAccount, Action: action, Amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.Action {
	case Deposit:
		b.BankAccount.Deposit(b.Amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.BankAccount.Withdraw(b.Amount)
	}
}
