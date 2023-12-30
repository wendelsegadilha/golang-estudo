package main

import "fmt"

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	// função anônima
	var f = func (text string)  {
		fmt.Println(text)
	}

	f("Wendel Segadilha")

	sum, sub, _ := calculos(10, 20) //ignora o resultado da multiplicação
	fmt.Println(sum, sub)

}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// retorno múltiplos
func calculos(n1, n2 int) (int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	return soma, subtracao, multiplicacao
}