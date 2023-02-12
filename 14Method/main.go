package main

import "fmt"

type Movie struct {
	Id       int
	Name     string
	Director string
}

func (m Movie) GetName() (string){
	// fmt.Println("Name is : ", m.Name)
	return m.Name
}

func (m *Movie) newDirector() {      // we use * as a reference variable here
	m.Director = "Tushar Pachouri"
	fmt.Println("New dir is: ",m.Director)
}

func main() {
	user1 := Movie{01,"Harry porter","JK rolin"}
	user2 := Movie{02,"Harry porter 2","JK rolin"}
	// user3 := Movie{03,"Harry porter 3","JK rolin"}
	fmt.Println("User1 movie is: ",user1.GetName())
	user2.newDirector()
	fmt.Println(user2.Director)
    num := 0

	for num <= 10 {
		defer fmt.Print(num," ")
		fmt.Print(num," ")
		num ++
	}

	
}
