package main

import (
	"api/src/banco"
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()
	banco.Conectar()

	
	r := router.Gerar()


	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}