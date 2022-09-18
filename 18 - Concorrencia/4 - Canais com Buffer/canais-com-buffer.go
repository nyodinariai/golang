package main

import "fmt"

func main(){
	// um canal com buffer é um canal com seu tamanho determinado
	canal := make(chan string, 2)
	canal <- "Ola mundo!"
	canal <- "Coisa boa"

	mensagem := <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}