package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// string은 immutable
	var str = "Hello World"
	str = "How are you?" // 가능
	//str[2] = 'a' // 불가능 (Error)
	println(str)

	// 문자열 합산
	var str1 = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	addr1 := stringheader.Data

	str1 += " World"
	addr2 := stringheader.Data

	str1 += " Welcome!"
	addr3 := stringheader.Data
	fmt.Println(str1)
	fmt.Printf("addr1: %x\n", addr1)
	fmt.Printf("addr2: %x\n", addr2)
	fmt.Printf("addr3: %x\n", addr3)
}
