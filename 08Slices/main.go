package main

import (
	"fmt"  //fmt refers to Format package
	"sort" //sort funciton for sort
)

func main() {
	fmt.Println("Slices: ")
	var heros = []string{"iron man", "captain america", "superman", "shaktimaan"} //array have fix number of elements and we can make slices if we don't know how many elements we wanna fill in our array
	fmt.Printf("Type of the variable \"heros\" is : %T\n", heros)                 //[4]string --> it show the sliceing in goLang
	heros = append(heros, "hulk", "hawkeye")
	fmt.Println(heros)

	//slicing of arrays :
	heros = append(heros[1:])
	fmt.Println(heros)
	heros = append(heros[1:3])
	fmt.Println(heros)
	heros = append(heros[:3])
	fmt.Println(heros)

	//make function
	marks := make([]int, 4)
	marks[0] = 69
	marks[1] = 87
	marks[2] = 8
	marks[3] = 7
	// marks[4] = 69 ---> Show error 'cause we defined the range above i.e '4'

	//we can append and then it will re allocate the space again and then it will show no error
	marks = append(marks, 43, 4, 3, 2, 1)

	// then we can sort with the help of sort functioon.
	sort.Ints(marks)

	fmt.Println(marks)
	fmt.Println("Size of the slice is : ", len(marks))
	fmt.Println(sort.IntsAreSorted(marks)) // It'll return bool i.e true in this case since we sort the array in line no. 35

	//Delete element from the slice
	var serie1 = []string{"123", "456", "789", "101112", "131415"}
	index := 2
	fmt.Println("Length of slice before append: ", len(serie1))
	serie1 = append(serie1[:index], serie1[index+1:]...) //Here the value of index 2 will be deleted from the slice
	fmt.Println(serie1)
	fmt.Println("Length of slice after append : ", len(serie1))
	// arr := []int{1,2,3,4,6,7}
	// missing = 0
    // for i :=0 in range(len(arr)){
    //     if arr[i] != i+1:
    //         missing = i+1
    //         break
	// }
    // fmt.Println("missing: ",missing)

	
}
