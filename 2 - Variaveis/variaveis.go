package main

import "fmt"

func main(){

	var variavel1 = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Muito Louco"
		variavel4 string = "Coisas Boas"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	
}