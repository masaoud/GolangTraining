package main

import "fmt"

func main() {
	//num in Decimal
	fmt.Println(42)

	//num in Binary
	fmt.Printf("%d - %b \n", 42, 42)

	//num in hexadecimal
	fmt.Printf("%x\n", 42)
	fmt.Printf("%#x\n", 42)
	fmt.Printf("%#X\n", 42)

	//loop
	for i:=0; i <= 100; i++{
		fmt.Printf("%d \t %b \t %x \n",i, i, i)
		

	}
}
