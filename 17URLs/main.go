package main

import (
	"fmt"
	"net/url"
)

const url1 string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=28&ab_channel=HiteshChoudhary"

func err1(err error) {
	if err != nil {
		panic(err)
	}
} 

func main() {
	fmt.Print("hello\n")
	fmt.Printf(url1)
	res1, err := url.Parse(url1)
	err1(err)
	// fmt.Println(res)
	// fmt.Println("host : ", res.Host)
	// fmt.Println("scheme : ", res.Scheme)
	// fmt.Println("Path : ", res.Path)
	// fmt.Println("Port : ", res.Port())
	// fmt.Println("RAw query : ", res.RawQuery)

	qparams := res1.Query()

	fmt.Printf("Query is type of %T\n: ", qparams) // it is type of key value pairs...
	fmt.Println(qparams["ab_channel"])

	for _, val := range qparams {
		fmt.Println("Params is : ", val)
	}

	PartOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "youtube.com",
		Path:    "/watch",
		RawPath: "user=tushar pachouri",
	}
	Url_string := PartOfUrl.String()
	fmt.Println("URL is : ", Url_string)

}
