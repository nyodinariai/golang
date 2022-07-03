package main

import "fmt"


type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	lograouro string
	numero uint8
}

func main(){

	fmt.Println("Arquivo structs")

	endereco := endereco{"Rua dos bobos", 0}

	var u usuario
	u.nome = "Nathan"
	u.idade = 5
	u.endereco = endereco

	fmt.Println(u)
}