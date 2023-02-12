package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println("Value of ptr(pointer) is: ", ptr)

	myNumber := 69
	var ptr1 = &myNumber                      // It is pointing or refering to the address of the var and we can do oprations on this ptr1
	fmt.Println("Value of pointer : ", ptr1)  // Address value of myNumber
	fmt.Println("Value of pointer : ", *ptr1) // Value of myNumber

	*ptr1 = *ptr1 + 2 //actually chaanging the value of the variable

	fmt.Println("Value of myNumber is now updated and became : ", myNumber)
}
