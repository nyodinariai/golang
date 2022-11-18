package main

import (
	"api/src/banco"
	"api/src/config"
	"api/src/handlers"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()
	DB := banco.Conectar()
	h := handlers.New(DB)
	
	r := router.Gerar()


	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}