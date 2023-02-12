package main

import "fmt"

func main() {
	fmt.Println("Arrays:  \n\n ")
	var arr [4]string

	arr[0] = "tushar"
	arr[1] = "tushar1"
	arr[2] = "tushar2"

	fmt.Println("Array is : ", arr)
	fmt.Printf("type of an element in arr is %T\n", arr[0])
	fmt.Println("Length of array arr is: ", len(arr)) // output will be 4 'cause i declare size of arr "4"

	var heros = [5]string{"superman","Ironman","captain america","thor"};
	fmt.Println("Heros are: ", heros)

}
