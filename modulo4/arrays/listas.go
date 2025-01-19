package main

import "fmt"

func main (){
	var array [2] string
	array[0]= "Hello"
	array[1]= "World"

	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[0],array[1])
	fmt.Println(array)

	numPrimos := [6]int{2,4,5,7,8,6}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3])
	fmt.Println(numPrimos[1])

	//var slice [] string
	slice := make([]string,5)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)

	slice[2]="Daniel"
	fmt.Println(len(slice))
	fmt.Println(slice[2])
	fmt.Println(slice[4])
	fmt.Println(slice[3])

	numPares := []int{2,4,6,8,10,12}
	fmt.Println(numPares)
	numPares = append(numPares, 14,20,24,22,26)
	fmt.Println(numPares)
	}