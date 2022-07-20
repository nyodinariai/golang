package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ola mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Função com parametro")

}