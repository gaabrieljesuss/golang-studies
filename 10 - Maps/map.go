package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "Gabriel",
			"ultimo": "Jesus",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus Arapiraca",
		},
	}
	fmt.Println(usuario2["nome"])
	delete(usuario2, "nome")

	fmt.Println(usuario2)
}