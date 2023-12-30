package main

import "fmt"

func main() {

	/*
		Em Go, os ponteiros são usados para transmitir o endereço de memória de uma variável. Eles são usados principalmente para permitir que funções modifiquem os valores originais das variáveis, passando o endereço de memória em vez de fazer cópias dos dados. Isso pode ser útil para otimizações de desempenho e manipulação eficiente de grandes conjuntos de dados.
	*/

	var x int
	var y *int // ponteiro para int
	fmt.Println(x, y)
	
	x = 10
	y = &x // atribui o endereço de memória onde se encontra o valor de x à variável y
	fmt.Println(x, y) // aqui exibe o endereço de memória
	fmt.Println(x, *y) // o "*" exibe o valor e não o endereço de memória (desreferenciação)
	
	x = 15 // a alteração será refletida em y
	fmt.Println(x, *y) // o "*" exibe o valor e não o endereço de memória (desreferenciação)


}