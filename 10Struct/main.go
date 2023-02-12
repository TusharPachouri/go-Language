package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Struct in go lang : ")
	tushar := User{"Tushar", "tusharpachouri001@gmail.com", "Btech", 20}
	fmt.Println("Details about tushar: ", tushar)
	fmt.Printf("%v is full detail.\n", tushar)
	fmt.Printf("%v is %v years old.\n", tushar.Name, tushar.Age)
	fmt.Printf("Details are %+v\n", tushar)

	if tushar.Name == "Tushar" {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	if num := 3; num%2==0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

}

type User struct {
	Name   string
	Email  string
	Course string
	Age    int
}
