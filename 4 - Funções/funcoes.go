package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f1 = func (txt string) {
		fmt.Println(txt)
	}

	f1("Texto da função 1")
}