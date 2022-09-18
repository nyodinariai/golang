package enderecos

import (
	"testing"
)

type cenarioDeTeste struct{
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T){

	cenariosDeTeste := []cenarioDeTeste{
		{ "Rua ABC", "Rua" },
		{ "Avenida Paulista", "Avenida" },
		{ "Rodovia dos Imigrantes", "Rodovia" },
		{ "Praça das Rosas", "Tipo Inválido" },
		{ "Estrada ABC", "Estrada" },
		{ "RUA DOS BOBOS ABC", "Rua" },
		{ "AVENIDA NOVA", "Avenida" },
		{ "", "Tipo Inválido" },
		
	}

	for _, cenario := range cenariosDeTeste{
		tipoDeEnredecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnredecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnredecoRecebido, cenario.retornoEsperado)
		}
	}
}