package main

import "fmt"

func main() {
	str1 := "Hello\t'World'\n"
	str2 := `"Hello\t'World'\n"` // 백쿼트는 특수문자가 동작하지 않음
	str3 := `
안녕하세요.
유안상입니다.
여러줄입니다.
	` // 여러줄로 출력 가능

	println(str1)
	println(str2)
	println(str3)

	// rune은 int32와 이름만 다를뿐 같은 타입
	var char rune = '한'
	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)

	// lenth
	str4 := "가나다라마"
	str5 := "ABCDE"
	fmt.Printf("len(str4) : %d\n", len(str4)) // 한글은 1글자당 3byte이기때문에 15가 출력
	fmt.Printf("len(str5) : %d\n", len(str5))

	// str -> rune[]
	str6 := "Hello 월드"
	runes := []rune(str6)
	fmt.Printf("len(str6) : %d\n", len(str6))
	fmt.Printf("len(runes) : %d\n", len(runes))
}
