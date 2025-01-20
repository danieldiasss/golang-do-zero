package main

import "fmt"

func main (){

	posição := 2
	switch posição {
	case 1:
		fmt.Println("Primeiro Lugar")
	case 2:
		fmt.Println("Segundo Lugar")
	case 3:
		fmt.Println("Terceiro Lugar")
	}
	
	familia := "Mirelly"
	switch familia{
case "Mirelly":
	fmt.Println("É minha namorada")
case "Franciane":
	fmt.Println("É minha mãe")
case "Isaías":
	fmt.Println("É meu pai")
}
}