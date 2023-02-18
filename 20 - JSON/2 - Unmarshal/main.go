package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	erro := json.Unmarshal([]byte(cachorroEmJson), &c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJson := `{"nome":"Toby","raca":"Poodle"}`

	c2 := make(map[string]string)

	erro2 := json.Unmarshal([]byte(cachorro2EmJson), &c2)

	if erro2 != nil {
		log.Fatal(erro2)
	}
	
	fmt.Println(c2)
}