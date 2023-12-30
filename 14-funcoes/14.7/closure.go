package main

import "fmt"

func closure() func() {
	texto := "Função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {

	texto := "Dentro da main"
	fmt.Println(texto)

	novaFucao := closure()
	novaFucao()

}