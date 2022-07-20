package main

import "fmt"

// ... parametro variático
func soma(numeros ...int) int{
	total := 0

	for _, numero := range numeros{
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int){
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main(){
	totalDasoma :=	soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)

	fmt.Println(totalDasoma)

	escrever("Olá mundo", 10, 20, 30, 10)
}