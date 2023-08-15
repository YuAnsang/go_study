package main

import "fmt"

func main() {
	// 포인터 기본
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가르키는 메모리의 값 : %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)

	// 포인터 비교
	var test = 10
	var test2 = 10

	var p1 *int = &test
	var p2 *int = &test
	var p3 *int = &test2
	var p4 *int

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
	fmt.Printf("p4 == nil : %v\n", p4 == nil)
}
