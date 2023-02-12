package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number here: ")

	//comma Ok , ok error
	input, _ := reader.ReadString('\n')
	fmt.Println("You string is : ", input)
	fmt.Printf("Type is: %T\n", input)

	fmt.Println("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	fmt.Println("Second number is: ", input2)

}
