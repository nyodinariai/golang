package main

import "fmt"

func generica(interef interface{}) {
	fmt.Println(interef)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(123465))

	mapa := map[interface{}]interface{}{
		1: "String",
		float32(100): true,
	}

	fmt.Println(mapa)
}