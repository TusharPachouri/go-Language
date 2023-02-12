package main

import "fmt"

func Adder(values ...int) int { // Here we are getting the value ...can be one or more than one but all should be int type
	sum := 0
	for _, val := range values {
		sum = sum + val
	}
	return sum
}

func minus(values ...int) (int,string) { // Here we are getting the value ...can be one or more than one but all should be int type
	sum := 0
	for _, val := range values {
		sum = sum - val
	}
	return sum , "string from minus function"
}

func main() {
	res := Adder(2, 3, 4, 1, 3, 4, 4)
	fmt.Println("Sum is : ", res)
    res2,_ := minus(2,3,4,6,7,3,4)  //for error handling
	ans,mes := minus(2232,423,23,4,32,2)
	fmt.Println("ans and message ",res2)
	fmt.Printf("ans : %v and message : %v",ans,mes)
}
