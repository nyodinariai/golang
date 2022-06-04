package main

import "fmt"

func somar(n1 int8, n2 int8) int8{
	return n1 + n2
}
func calculosMatematicos(n1, n2 int8)(int8,int8){
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main(){
	soma := somar(5,2)
	fmt.Println(soma)

	var f = func (txt string)  string{
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da Funcao")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(25, 30)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}