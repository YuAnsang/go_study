package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintHangul() {
	var hanguls = []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

//func main() {
//	go PrintHangul()
//	go PrintNumbers()
//
//	time.Sleep(3 * time.Second)
//}

var wg sync.WaitGroup

func sumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합계 : %d\n", a, b, sum)
	wg.Done()
}

//func main() {
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go sumAtoB(1, 1000000000)
//	}
//	wg.Wait()
//}
