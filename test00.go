package main

import "fmt"

const (
	a = iota
	b
	c
)

func main()  {
	fmt.Println(a )
	fmt.Println(b)
	fmt.Println(c)
	
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n ", s)
	fmt.Println("Test")
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n ",bs)
	 for i :=0; i < len(s); i++{
		 fmt.Printf("%#U ", s[i])
	 }

	 fmt.Println("")

	 for i, v := range s{
	 	fmt.Printf("at position %d we have hex %#x \n", i , v)
	 }
	// Slice
	//creat an array
	//var x [4] int
	//assign a value
	//x[3]=44

	//Slice
	x := []int{11, 22, 33, 34, 45}
	fmt.Println(x)
	fmt.Println(x[1:3])
	x = append(x, 666, 777)
		fmt.Println(x)


	//append
	y := []int{123, 345, 5566}
	// falsch // x = append(x, y)
		x = append(x, y...)
	fmt.Println(x)


   // Delet

   x = append(x[:2],x[4:]...)
   	fmt.Println(x)


   //muli dimintional slice
   mp := [] string {"masaoud","khallouf"}
   fmt.Println(mp)
   xp := [] string{"Ahmad","Ali"}
   	fmt.Println(xp)

   pp := [][]string{mp,xp}
   	fmt.Println(pp)

   //map -->     m := map [Key] value {}

	m := map[string]int{
		"Masaoud": 32,
		"Ahmad":   4,
		"Talia":   6,
	}
	fmt.Println(m)
	fmt.Println(m["Masaoud"])
	fmt.Println(m["Ali"])

	if v, ok := m["Masaoud"]; ok{
		fmt.Println("this is the if print ", v)
		fmt.Println(ok)
	}

	m["Sarah"] = 25
	for k, v:= range m{
		fmt.Println(k,v)
	}


	//delet
	delete(m, "Sarah")
	fmt.Println(m)


	//Recursion
	/*

	func main() {

		n := factorial(4)
		fmt.Println(n)
	}

	func factorial(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}
	 */

/*
   func main() {
   	a := 42
   	fmt.Println(a)
   	fmt.Println(&a) // & gives you the address

   	fmt.Printf("%T\n", a)
   	fmt.Printf("%T\n", &a)

   	b := &a
   	fmt.Println(b)
   	fmt.Println(*b) // * gives you the value stored at an address when you have the address
   	fmt.Println(*&a)

   	*b = 43
   	fmt.Println(a)

   }
 */


}


