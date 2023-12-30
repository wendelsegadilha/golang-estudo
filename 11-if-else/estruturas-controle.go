package main

import "fmt"

func main() {

	idade := 10

	// IF ELSE comum
	if idade >= 18 {
		fmt.Println("Adulto")
	} else if idade > 12 && idade < 18 {
		fmt.Println("Adolecente")
	} else {
		fmt.Println("Criança")
	}

	// IF init (Inicializa uma variável direto na condição)
	if salario := 3200; salario > 2500 { // cria e inicializa a variavel salario (limitada ao escopo de IF ELSE)
		fmt.Println("Emprego aceito")
	} else {
		fmt.Println("Emprego rejeitado")	
	}

}