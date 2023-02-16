package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}