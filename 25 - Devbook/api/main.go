package main

import (
	"api/src/config"
	"api/src/migration"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Carregar()
	migration.Migrate()
	r := router.Gerar()


	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}