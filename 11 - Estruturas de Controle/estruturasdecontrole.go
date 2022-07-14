package main

import "fmt"

func main(){
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15{
		fmt.Println("Maior que 15")
	}else {
		fmt.Println("Menor ou igual a 15")
	}

	if outronumero := numero; outronumero > 0{
		fmt.Println("Número é maior que 0")
	}
}