package main

import "fmt"

func main() {

	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	nomes := [3]string{"Gabriel", "Jesus", "Santos"} 

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "Gabriel" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Santos",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}