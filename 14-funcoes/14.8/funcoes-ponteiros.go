package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1 // "*" acessa o valor do endereço de memória
}

func main() {

	//passando copia de valor
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//passando endereço de memória (referencia)
	novoNumero := 40
	fmt.Println(novoNumero)
	inverteSinalComPonteiro(&novoNumero) // "&" acessa o endereço de memoria
	fmt.Println(novoNumero)

}