package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in goLan")
	lang := make(map[int]string)

	lang[1] = "English"
	lang[2] = "Hindi"
	lang[3] = "Sanskrit"
	lang[4] = "Japanese"

	fmt.Println(" Map or (dictionary in python) is: ", lang)
	fmt.Printf("Type of maps in goLang is : %T", lang)
	fmt.Println("LAnguage of key 2", lang[2]) // key 2(int) --> value "hindi"(string)
	delete(lang, 3)                           //it will delete the value that is present at position where value of key is 3
	fmt.Println("After delete opration has been done out map become: ", lang)

	for key, value := range lang {
		fmt.Printf("for key %v the value is :%v\n", key, value) // we can print all the key : value in the lang map

	}
	fmt.Print("The values are : ")
	for _, value := range lang {
		fmt.Printf("%v ", value) // we can print all values in the lang map withot needing the keys
	}

}