package main

import "fmt"

func main() {
	b := &BankAccount{}
	cmd := NewBankAccountCommand(b, Deposit, 100)
	cmd.Call()
	cmd2 := NewBankAccountCommand(b, Withdraw, 25)
	cmd2.Call()
	cmd2.Undo()

	// Composite example with command pattern
	from := &BankAccount{100}
	to := &BankAccount{0}
	c := NewMoneyTransferCommand(from, to, 25)
	c.Call()
	fmt.Println(from, to)
	c.Undo()
	fmt.Println(from, to)
}
