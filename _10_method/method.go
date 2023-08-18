package main

import "fmt"

type account struct {
	balance int
}

func (a *account) withdraw(amount int) {
	a.balance -= amount
}

func main() {
	a := account{100}
	a.withdraw(30)
	fmt.Printf("%d \n", a.balance)
}
