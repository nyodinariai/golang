package main

import (
	"errors"
	"fmt"
)

func main()  {
	var inteiro1 int8 = 25
	var inteiro2 int16 = 250
	var inteiro3 int32 = 25000000
	var inteiro4 int64 = 2500000000000000000

	fmt.Println(inteiro1)
	fmt.Println(inteiro2)
	fmt.Println(inteiro3)
	fmt.Println(inteiro4)

	var inteiroPositivo uint = 25

	fmt.Println(inteiroPositivo)

	//alias
	// int32 = rune

	var inteiro5 rune = 123456


	// BYTE = UINT8
	var inteiro6 byte = 123

	fmt.Println(inteiro5)
	fmt.Println(inteiro6)

	var numeroreal1 float32 = 123.650
	var numeroreal2 float64 = 123.6504545641132132123

	fmt.Println(numeroreal1)
	fmt.Println(numeroreal2)

	numeroreal3 := 123456.456789123

	fmt.Println(numeroreal3)


	var str string = "Texto Simples"
	fmt.Println(str)

	str1 := "Texto simples marmota"

	fmt.Println(str1)

	char := 'A'
	fmt.Println(char)

	//Texto

	texto := 5
	fmt.Println(texto)

	var boolean bool
	fmt.Println(boolean)

	var erro error = errors.New("Erro do nada")
	fmt.Println(erro)
}