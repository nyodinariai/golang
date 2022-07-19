package main

import (
	"fmt"
)

func main(){
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Printf("Incrementando %d \n", i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Printf("Incrementando %d \n", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3] string {"João", "Davi", "Lucas"}

	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}
}