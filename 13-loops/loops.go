package main

import (
	"fmt"
	"time"
)

func main() {

	//uso 1
	i := 0 
	for i < 10 {
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second) // coloca a execução para dormir 1 segundo
	}
	fmt.Println(i)

	//uso 2
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J: ", j)
		time.Sleep(time.Second)
	}

	//uso 3 (range)
	nomes := [3]string {"Venes", "Wendel", "Anny"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//uso 4
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	//uso 5
	usuario := map[string]string {
		"nome": "Wendel",
		"sobrenome": "Segadilha",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//uso 6
	for {} //loop infinito

}