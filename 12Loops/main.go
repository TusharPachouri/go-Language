package main

import "fmt"

func main() {
	fmt.Println("Loops: ")
	student := []string{"tushar", "sarthak", "saif", "anmol", "nikhil"}
	for d := 0; d < len(student); d++ {
		fmt.Printf("Name of %v No. is %v\n", d+1, student[d]) //here we are traversing the slice and printing the values
	}
	for d := range student { //range is a function which treverse the "d" variable through our slice or array
	    
		fmt.Printf("Name of %v No. is %v\n", d+1, student[d])
	}
	for index, name := range student {      // Here index value initialize from 0 and go up-to range i.e "4" and their respective value of the index...
		fmt.Printf("Name at index %v is: %v\n", index+1, name)
	}
	for _, name := range student {      // Here we are not interested in indexing and only want the values of student array 
		fmt.Printf("Name is: %v\n", name)
	}
	num := 0
	for num < 11{  //just like while loop
		fmt.Println("number is: ",num)
		if num == 2{
			goto v //when value of num becomes 2 it will go to "v"
		}
		num++
	}
	v:fmt.Println("Hello jump ")
}
