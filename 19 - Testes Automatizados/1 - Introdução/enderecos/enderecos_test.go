package enderecos_test

import (
	"testing"
	. "introducao-testes/enderecos"
	"github.com/stretchr/testify/assert"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		enderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		assert.Equal(t, cenario.retornoEsperado, enderecoRecebido)
	}
}