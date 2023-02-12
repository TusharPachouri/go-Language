package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello")
	fmt.Println("Enter a number btw 1-5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your number is : ", input)
	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println("Add 1 to your number : ", numInput+1)
	}


	//or	
	var str string
	fmt.Println("Enter your name: ")
	fmt.Scan(&str)   //Tushar Pachori === Tushar????
	fmt.Println(str)
}
