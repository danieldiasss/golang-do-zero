package main

import "fmt"

func main (){
	var array [2] string
	array[0]= "Hello"
	array[1]= "World"

	fmt.Println(array[0],array[1])
	fmt.Println(array)

	numPrimos := [6]int{2,4,5,7,8,6}
	fmt.Println(numPrimos)
	
}