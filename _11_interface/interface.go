package main

import "fmt"

type Animal interface {
	Eat()
	Walk(distance int) int
}

type Cat struct {
	Name string
	Age  int
}

func (c *Cat) Eat() {
	fmt.Printf("%s Eat!!!!\n", c.Name)
}

func (c *Cat) Walk(distance int) int {
	fmt.Printf("%s walk... %d\n", c.Name, distance)
	return distance
}

func main() {
	cat := Cat{"mya", 9}
	var animal Animal

	animal = &cat

	cat.Eat()
	cat.Walk(2000)

	animal.Eat()
	animal.Walk(1500)
}
