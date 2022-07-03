package main

import "fmt"

type pessoa struct{
	nome string
	idade uint8
	altura uint8
	sobrenome string
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main(){
	fmt.Println("Herança")

	p1 := pessoa{"Marcos", 20, 150, "Silva"}
	estudante1 := estudante{p1, "Administração de Empresas", "Universidade Estadual do Ceará"}

	fmt.Println(estudante1)
}