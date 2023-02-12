package main

import "fmt"

func main() {
	var name int
	fmt.Scan(&name)

	
	switch name {
	case 1:
		fmt.Println("Saiful")
	case 2:
		fmt.Println("Sarthak")  // fallthrough will also run the next case wheater it is true or not... 
		fallthrough
	case 3:
		fmt.Println("Tushar")
		fallthrough
	case 4:
		fmt.Println("nikhil")
	case 5:
		fmt.Println("ankit")
	default:
		fmt.Println("Anmol")
	}

}
