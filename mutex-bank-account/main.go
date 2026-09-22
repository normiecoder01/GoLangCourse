package main

import(
	"sync"
	"fmt"
)

type Account struct {
	Balance int
	Mutex sync.Mutex
}

func( acc *Account) Deposit(amount int) {
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()

	acc.Balance+= amount
	fmt.Println("Deposited Rs.",amount,". Current Balance is Rs.",acc.Balance,".")
}

func (acc *Account) Withdraw(amount int){
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()

	if(amount > acc.Balance){
		fmt.Println("Insufficient Balance: Cannot withdraw Rs.", amount,". Current balance is Rs.", acc.Balance,".")
		return 
	}
	acc.Balance-= amount
	fmt.Println("Withdrawn Rs.",amount,". Current Balance is Rs.",acc.Balance,".")

}

func (acc *Account) CheckBalance() {
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()
	fmt.Println("Current balance is Rs.",acc.Balance,".")
}

func main(){
	var wg sync.WaitGroup
	acc := Account{Balance: 1000}
	for i:= 0; i<10 ; i++{
		wg.Add(1)
		go func(amount int){			
			defer wg.Done()
			acc.Deposit(amount)
		}(i*100)

	}
		for i:= 0; i<10 ; i++{
		wg.Add(1)
		go func(amount int){			
			defer wg.Done()
			acc.Withdraw(amount)
		}(i*100)

	}
			wg.Wait()
}