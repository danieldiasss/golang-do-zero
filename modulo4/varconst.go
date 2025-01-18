package main

import(
    "fmt"
)

func main () {
    //Variáveis
    // var+ nome + tipo
    var nome string
    nome= "Bento"
    fmt.Println(nome)

    nome="daniel"
    fmt.Println(nome)
    
    var idade int
    idade=4
    fmt.Println(idade)

    var a = "daniel"
    fmt.Println(a)

    var b,c int= 1, 2
    fmt.Println(b,c)

    var d = true
    fmt.Println(d)

    f := "apple"
    fmt.Println(f)

    //Constantes
    const idadeBento = 4
    fmt.Println(idadeBento)
}