package main

import "fmt"

func main()  {
	
	//Operadores Aritiméticos

	soma := 1 + 2
	subtracao := 1 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 4
	modulo := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, modulo)

	var num1 int16 = 25
	var num2 int16 = 35
	soma2 := num1 + num2

	fmt.Println(soma2)

	// Operadores de Atribuição

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais

	fmt.Println("Maior", 1 > 2)
	fmt.Println("Maior Igual", 1 >= 2)
	fmt.Println("Menor", 1 < 2)
	fmt.Println("Menor Igual", 1 <= 2)
	fmt.Println("Igual", 1 == 2)
	fmt.Println("Diferente", 1 != 2)

	// Operadores Lógicos

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Operadores Unitários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	// Go não possui operador ternário
}