package main

import (
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance float64
	OwnerName string
}

func (acc *Account) Deposit(amount float64) error {
	if(amount <= 0){
		return fmt.Errorf("Deposit amount should be a positive number")
	}
	acc.Balance+=amount
	fmt.Println("Paytm par", amount, "Rupay Prapt Hue. Paytm Krooooooooo...(AccountNumber:",acc.AccountNumber,"Current Balance:",acc.Balance,")")
	return nil
}
func(acc *Account) Withdraw(amount float64) error {
		if(amount <= 0){
		return fmt.Errorf("Withdraw amount should be a positive number")
	}
	if(amount > acc.Balance){
		return fmt.Errorf("Withdraw amount should be less than account Balance. Current Balance:",acc.Balance)
	}
	acc.Balance-=amount
	fmt.Println("Paytm se", amount, "Rupay nikal gye. Paytm Krooooooooo...(AccountNumber:",acc.AccountNumber,"Current Balance:",acc.Balance,")")
	return nil

}

func (acc *Account) GetBalance() float64 {return acc.Balance}

func (acc *Account) String() string {
	return fmt.Sprintln("Saving Account Number:",acc.AccountNumber,"A/c Holder Name:", acc.OwnerName)
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func(sa *SavingsAccount)  AddInterest() error {
	if(sa.Balance <= 0){
		return fmt.Errorf("Interest cannot be applied if the balance is 0")
	}
	interest:= sa.InterestRate * sa.Balance
	err:= sa.Deposit(interest)
	if(err!= nil){
		return fmt.Errorf("AddInterest: ",err) 
	}
	fmt.Println("Interest of",sa.InterestRate,"applied to Savings Account number",sa.AccountNumber,".Current balance is",sa.Balance)
	return nil
}

type OverDraftAccount struct {
	Account
	OverDraftLimit float64
}

func (oa *OverDraftAccount) Withdraw(amount float64) error{
	if(amount <= 0){
		return fmt.Errorf("Withdraw amount should be a positive number")
	}
	if(amount >( oa.Balance + oa.OverDraftLimit) ){
		return fmt.Errorf("Withdraw amount should be less than account Balance. Current Balance:",oa.Balance)
	}
	oa.Balance-=amount
	fmt.Println("Paytm se", amount, "Rupay nikal gye. Paytm Krooooooooo...(AccountNumber:",oa.AccountNumber,"Current Balance:",oa.Balance,")")
	return nil
}

func main() {

	// ----- Savings Account -----
	sa := SavingsAccount{
		Account: Account{
			AccountNumber: "SA123",
			Balance:       1000,
			OwnerName:     "Amit Sharma",
		},
		InterestRate: 0.05,
	}

	fmt.Println(sa.String())

	// Deposit tests
	fmt.Println(sa.Deposit(500))
	fmt.Println(sa.Deposit(-100)) // should fail

	// Withdraw tests
	fmt.Println(sa.Withdraw(300))
	fmt.Println(sa.Withdraw(5000)) // should fail
	fmt.Println(sa.Withdraw(-50))  // should fail

	// Interest tests
	fmt.Println(sa.AddInterest())

	// Force zero balance for edge case
	sa.Balance = 0
	fmt.Println(sa.AddInterest()) // should fail


	// ----- OverDraft Account -----
	oa := OverDraftAccount{
		Account: Account{
			AccountNumber: "OD456",
			Balance:       2000,
			OwnerName:     "Priya Deshmukh",
		},
		OverDraftLimit: 1000,
	}

	fmt.Println(oa.String())

	// Deposit tests
	fmt.Println(oa.Deposit(1000))
	fmt.Println(oa.Deposit(-200)) // should fail

	// Withdraw tests (within balance)
	fmt.Println(oa.Withdraw(500))

	// Withdraw using overdraft
	fmt.Println(oa.Withdraw(3000)) // allowed (2000 + 1000 limit)

	// Exceed overdraft
	fmt.Println(oa.Withdraw(100)) // should fail now
	fmt.Println(oa.Withdraw(1000))
	// Negative withdraw
	fmt.Println(oa.Withdraw(-10)) // should fail


	// ----- Base Account -----
	acc := Account{
		AccountNumber: "AC789",
		Balance:       500,
		OwnerName:     "Rahul Verma",
	}

	fmt.Println(acc.String())

	fmt.Println(acc.Deposit(200))
	fmt.Println(acc.Withdraw(100))
	fmt.Println(acc.Withdraw(1000)) // should fail

	fmt.Println("Final Balance:", acc.GetBalance())
}

