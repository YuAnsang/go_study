package main

import "fmt"

type Student struct {
	Name   string
	Class  string
	Number int
	Score  float64
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	// 각자 초기화
	var student Student
	student.Name = "안상"
	student.Class = "A"
	student.Number = 1
	student.Score = 100.0
	fmt.Println(student.Name)
	fmt.Println(student.Class)
	fmt.Println(student.Number)
	fmt.Println(student.Score)

	// 모든 필드 초기화
	var student2 = Student{"지오", "B", 2, 99.9}
	fmt.Println(student2.Name)
	fmt.Println(student2.Class)
	fmt.Println(student2.Number)
	fmt.Println(student2.Score)

	// 일부 필드 초기화
	var student3 = Student{Name: "현민", Class: "C"}
	fmt.Println(student3.Name)
	fmt.Println(student3.Class)
	fmt.Println(student3.Number)
	fmt.Println(student3.Score)

	// 필드 포함
	var vip = VIPUser{
		User{"안상", "kahlman", 36},
		1,
		100000,
	}
	fmt.Println(vip.Name)
	fmt.Println(vip.ID)
	fmt.Println(vip.Age)
	fmt.Println(vip.VIPLevel)
	fmt.Println(vip.Price)
}
