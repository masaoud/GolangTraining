package main

import "fmt"

func main() {
	/*fmt.Println("Hello world")
	foo()
}
func foo()  {
	for i := 0 ; i < 100 ; i++ {
		if i %2 == 0{
			fmt.Println(i)
		} else if i %2 != 0 {

			fmt.Println("here is odd num",i)
		}

	}*/
	n := "Masaoud"
	switch n {
	case "M","masaoud","Mm":
		fmt.Println("any think")
	case "Masaoud":
		fmt.Println("true")
	default:
		fmt.Println("this is default")



	}



}


