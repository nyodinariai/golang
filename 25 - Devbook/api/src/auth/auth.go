package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Gera o token de autenticação
func Token(usuarioID uint) (string, error){
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.Secret_key)) //secret
}

//Verifica se o token é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil{
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido")
}

func extrairToken(r *http.Request) string{
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token)(interface{},error){
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.Secret_key, nil
}

func ExtrairUsuarioId(r *http.Request) (uint, error){
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil{
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f",permissions["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return uint(usuarioID), nil
	}

	return 0, errors.New("Token invalido")
}