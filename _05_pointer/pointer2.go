package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	ChangeData(&data)

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	// 포인터 변수가 아무리 많아도 인스턴스는 1개
	var p1 *Data = &Data{}
	var p2 *Data = p1
	var p3 *Data = p1
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
