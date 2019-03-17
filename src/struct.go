package main

import "fmt"

type person struct {
	name string
	age int
	data map[string]int
}

func newPerson () *person {
	p := person{}
	p.data = map[string]int{}
	return &p
}

func main() {

	p1 := person{}

	p1.name = "joonhak"
	p1.age = 18
	p1.data = make(map[string]int)

	// 선언과 동시에 초기화
	p2 := person{
		name : "joonhak",
		age : 18,
		data : map[string]int {"joonhak" : 18},
	}
	p3 := person{"joonhak", 18, make(map[string]int)}

	p4 := new(person)   // return pointer
	p4.name = "jonnhak" // `pointer`도 . 으로 접근 가능 (이 경우 포인터가 자동으로 dereference 된다.)

	// `person`의 data 필드가 초기화되지 않았기 때문에 에러 발생
	// p4.data["joonhak"] = 18

	p5 := newPerson()
	p5.data["joonhak"] = 18

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)

}