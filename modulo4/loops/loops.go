package main

import "fmt"
//import "time"

func main(){

	// i++ -> soma 1
	// i-- -> subtratindo 1
	//sum := 0
	//for i := 0; i < 10; i++ {
		//fmt.Println("Sum", sum)
		//fmt.Println("Indice", i)
		//sum += i
		// é a mesma coisa que: sum=sum + i
		//sum -= i-> sum = sum - i

		//É como se ao ginal do loop fosse adicionado 1 ao índice i
		// i++
		// i = i + 1
	//}
	//fmt.Println(sum)

	//loop infinito 
	//for{
	//	fmt.Println("Golang do zero")
	//	time.Sleep(2 * time.Second)
	//}

	//for range
	frutas := []string{"laranja","maça","banana","uva","kiwi"}
	for i, fruta := range frutas{
		fmt.Println("fruta",fruta,"indice", i)
	}

}