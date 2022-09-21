package banco

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o banco de dados
	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string{
viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

  	if err != nil {
    	log.Fatalf("Error while reading config file %s", err)
  	}

   value, ok := viper.Get(key).(string)
	if !ok {
    	log.Fatalf("Invalid type assertion")
  	}

  return value
}



// Abre a conexão com banco de dados
func Conectar() (*sql.DB, error){
	
	

	
	stringConexao := viperEnvVariable("DB_USER")+":"+viperEnvVariable("DB_PASSWORD")+"@/"+viperEnvVariable("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}