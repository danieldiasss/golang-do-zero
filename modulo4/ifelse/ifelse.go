package main

import "fmt"

//IF/ ELSE
//SE/ SENÃO

//if (condição) {<ação>}
func main(){
	numero:=8
	if numero== 1 {
		fmt.Println("Valor é igual a 1")
	} else if numero == 2 {
			fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor é diferente de 1 e 2")
	}

	//Operações
	fmt.Println(1+1)
	fmt.Println(2-1)
	fmt.Println(2*2)
	fmt.Println(10/2)
	//resto da divisão
	fmt.Println(10%2)

	if numero%2 == 0 {
		fmt.Printf("%d é par", numero)
	} else {
	   fmt.Printf("%d é ímpar", numero)
	}
}


	