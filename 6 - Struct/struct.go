package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	var u usuario
	u.nome = "Gabriel"
	u.idade = 20
	fmt.Println(u)

	var end endereco = endereco {"Rua das casas", 1}

	usuario2 := usuario{"Matheus", 25, end}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Augusto"}
	fmt.Println(usuario3)
}