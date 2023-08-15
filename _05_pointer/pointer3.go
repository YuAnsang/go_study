package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	// Go 언어에서는 메모리 공간이 함수 외부로 공개되는지 여부를 자동으로 검사해 스택 메모리 또는 힙 메모리에 할당할지를 결정함.
	userPointer := NewUser("안상", 36)
	fmt.Println(userPointer)
}
