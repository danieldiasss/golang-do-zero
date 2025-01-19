package main

import "fmt"

func main () {
	fmt.Println(soma(42,13))

	//Apliquei uma variável e dei nome a ela,variável men:=
	men:= menos(20,5)
	fmt.Println(men)

	fmt.Println(divisao(42,3))

	nome1, nome2:= printName("DANIEL",)
	fmt.Println(nome1)
	fmt.Println(nome2)

	nome,sobrenome:= printNomeCompleto("Daniel","Ribeiro")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}


func soma (x int, y int) int {
	return x+y
	
}

func menos (x int, y int) int {
	return x-y

}

func divisao (x int, y int) int{
	return x/y
}

func printName (nome string) (string,string){
	return nome, nome
}

//Função começando com letra minúscula:
//Função é privada
//Função ela só pode ser utilizada no próprio pacote
func printNomeCompleto (nome, sobrenome string) (string, string){
	return nome,sobrenome
}