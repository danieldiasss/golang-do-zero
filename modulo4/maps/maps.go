package main

import "fmt"

func main() {
	idade  := map[string]int{}
	idade["daniel"] = 18
	idade["mirelly"] = 17
	fmt.Println(idade)
	fmt.Println(idade["daniel"])
	fmt.Println(idade["mirelly"])

	anonasc := map[string]int{
	"daniel": 2006,
	"mirelly": 2008,
}

	fmt.Println(anonasc)
	fmt.Println(anonasc["daniel"])
	fmt.Println(anonasc["mirelly"])
	anonasc["golangdozero"] = 2024
	fmt.Println(anonasc)
}