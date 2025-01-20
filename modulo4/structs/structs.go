package main

import "fmt"

//type <nome da nossa estrutura> struct {<campos>}
type Pessoa struct {
	Nome string
	Idade int	
}

type Profissao struct {
	Pessoa 
	Tipo string	
}

func main (){
	fmt.Println(Pessoa{"Daniel",18})
	fmt.Println(Pessoa{Nome:"Mirelly",Idade: 17})
	fmt.Println(Pessoa{Nome:"Daniel"})
	p1 := Pessoa{Nome:"João", Idade: 20}
	
	//go run structs.gofmt.Println(p1.Nome)
	//fmt.Println(p1.Idade)

	p1.Idade = 21
	//fmt.Println(p1.Idade)

	p2 := Pessoa{Nome:"Patrick", Idade: 2}
	p3 := Pessoa{Nome:"Daniel",Idade: 18}
	p4 := Pessoa{Nome:"Mirelly", Idade: 17}

	pessoas:= []Pessoa{}
	pessoas= append(pessoas,p1,p2,p3,p4 )
	//fmt.Println(pessoas)

	//structs +map
	
	//alunos:= map[string][]Pessoa{}
	//alunos["programação"] = pessoas
	//fmt.Println(alunos)

	var alunos = map[string][]Pessoa{
	"programação": {{Nome: "Daniel",Idade: 18}, {Nome: "Mirelly",Idade: 17}},

	"engenharia": {{Nome: "Daniel2", Idade: 18}, {Nome: "Mirelly2", Idade: 17}},
}
	fmt.Println(alunos)

	//struct herdando campos de outra struct
	prof:= Profissao{p2,"dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}