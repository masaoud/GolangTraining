package main

import "fmt"
//import "strconv"

type person struct {
	first string
	last  string
	age int
}
type secretAgent struct {
	person
	address string

}

func main() {


	// struct  Structs allow us to compose together values of different types.
	p1 := person{
		first: "masaoud",
		last:  "khallouf",
		age : 32,
	}
	p2 := person{
		first: "tim",
		last:  "lol",
	}
	//embedded struct take one struct and embed it in another struct.

	sa1 := secretAgent{
		person: person{
			first: "Alli",
			last: "Alli",
			age: 33,
		},
		address: "Memhardstr.8",

	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.address)

	fmt.Println(p1.first, p1.age)
	fmt.Println(p2.first)

	//An anonymous struct is a struct which is not associated with a specific identifier.

	p3:= struct {
		first string
		last  string
		age   int
	}{
		first: "Ahmad",
		last:  "khallouf",
		age:   4,
	}

	fmt.Println(p3)


	/*myArr := make([]int, 0)
		for i:=0; i<10; i++ {
			myArr = append(myArr, i)
		}
		myArr[4] = 22

		a,_ := strconv.Atoi("4")

		fmt.Println(a )
	}

	func ChangeName(in string)  {
		fmt.Println(in)
	}

	type person struct {
		name string
		age int

		nr int
		city string


// anynom func
	func() {
		fmt.Println("Anynom func")
	}()

	func(x int) {
		fmt.Println("Anynom func with param",x)
	}(32)


	//returning function

	import (
		"fmt"
	)

	func main() {
		fmt.Println(bar()())

	}

	func bar() func() int {
		return func() int {
			return 451
		}
	}
	*/

}
