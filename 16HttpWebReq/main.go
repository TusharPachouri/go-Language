package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T",res)
	defer res.Body.Close()

	data,err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(data))
}
