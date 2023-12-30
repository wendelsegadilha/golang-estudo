package main

import (
	"fmt"
	"reflect"
)

func main() {

	// === ARRAYS  (tem tamanho fixo)
	//declaração 1
	var frutas [5]string
	frutas[0] = "manga"
	frutas[1] = "uva"
	frutas[2] = "banana"
	frutas[3] = "pêra"
	frutas[4] = "goiaba"
	fmt.Println((frutas))

	//declaração 2
	nomes := [3]string {"Wendel", "Venes", "Anny"}
	fmt.Println(nomes)

	//declaração 3
	notas := [...]int {1, 2, 3, 4, 5, 6}
	fmt.Println(notas)

	// === SLICES (não tem tamanho fixo)
	cidades := []string {"Santa Inês", "Zé Doca", "Bom Jardim"}
	cidades = append(cidades, "Santa Luzia") // add um novo item ao slice cidades
	fmt.Println(cidades)

	// verificando tipos
	fmt.Println("Array: ", reflect.TypeOf(nomes), "Slice: ", reflect.TypeOf(cidades))


}