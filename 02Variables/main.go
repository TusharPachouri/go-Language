package main

import "fmt"

const Name string = "tushar pachouri"

func main() {
	var str string = "Tushar Pachouri"
	fmt.Println(str)
	fmt.Printf("Variable is type of: %T \n", str)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("Variable is type of: %T \n", isBool)

	var smallval uint8 = 255  //there are uint8,32,64 and int8,32,64
	fmt.Println(smallval)
	fmt.Printf("Variable is type of: %T \n", smallval)

	var smallFloat float32 = 69.69696996996969696969969696
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type of: %T \n", smallFloat)

	var largeFloat float64 = 69.69696996996969696969969696
	fmt.Println(largeFloat)
	fmt.Printf("Variable is type of: %T \n", largeFloat)

// Default values: 

	var exampleVariable int
	fmt.Println(exampleVariable) //it will print the value of 0 that is 0 is the pre defined value in int

    var examplestr string
	fmt.Println(examplestr)

//implicit type
    var name ="Tushar Pachouri"
	fmt.Println(name)
	//** name = 8 cant do that 'cause lexar already decide that name is a str type cant change the value now** 
// NO var style
    age := 69
	fmt.Println(age)

// Constant
    fmt.Println(Name)
}