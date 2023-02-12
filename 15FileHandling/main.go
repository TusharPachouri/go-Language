package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func read(file string) {
	data, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	checkError(err)
	fmt.Println("Data inside file is: ", string(data))
}

func main() {

	text := "Tushar Pachouri is great"
	file, err := os.Create("./file.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	len, err := io.WriteString(file, text)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of the string in file file: ", len)

	read("./file.txt")

}
