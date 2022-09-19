package banco

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o banco de dados
)

// Abre a conexão com banco de dados
func Conectar() (*sql.DB, error){
	stringConexao := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}