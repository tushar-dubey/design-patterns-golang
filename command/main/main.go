package main

func main() {
	b := &BankAccount{}
	cmd := NewBankAccountCommand(b, Deposit, 100)
	cmd.Call()
	cmd2 := NewBankAccountCommand(b, Withdraw, 25)
	cmd2.Call()
	cmd2.Undo()
}
