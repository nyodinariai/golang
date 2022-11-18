package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Conexao com o banco de dados
	StringConexaoBanco = ""
	
	//Porta do Servidor
	Porta = 0
)

// Inicializa as variáveis de ambiente
func Carregar(){
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%/)/%s?charset=utf8&parseTime=True&Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)	
}