package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main(){

	tipoEndereco := enderecos.TipoDeEndereco("Rodovia Das rodas")

	fmt.Println(tipoEndereco)
}