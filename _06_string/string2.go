package main

import (
	"reflect"
	"unsafe"
)

func main() {
	// string의 구조는 필드가 2개인 구조체 (reflect 패키지의 StringHeader)
	str1 := "안녕하세요. 한글입니다."
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	println(stringHeader1)
	println(stringHeader2) // 주소 확인

}
