package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(slice)

	var array2 [5]string
	array2[0] = "Posição 1"
	fmt.Println(array2)

	array3 := [5]string {"a", "b", "c", "d", "e"}
	fmt.Println(array3)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice2 := array3[1:3]
	fmt.Println(slice2)

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 6)

	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}