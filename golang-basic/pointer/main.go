package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ptr *int

	ptr = &a

	fmt.Println(a)
	fmt.Println("Address of a:", ptr)
	fmt.Println("Value at the address stored in ptr:", *ptr)

	*ptr = 20
	fmt.Println("New value of a:", a)
}
