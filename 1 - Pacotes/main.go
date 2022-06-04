package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escrevendo seu primeiro código")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("nt@gmail.com")
    if err != nil {
        fmt.Println(err)
    }
}