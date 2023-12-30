package main

import (
	"errors"
	"fmt"
)

func main() {

	//int 8, int16, int32, int64 (tamanho em bits)
	//uint 8, uint16, uint32, uint64 (aceita somente números positivos)
	//int -> usa arquitetura padrão do PC x32 ou x64
	var idade uint8 = 29
	fmt.Println(idade)

	//float32, float64
	var salario float32 = 3200.65
	altura := 1.75
	fmt.Println(salario)
	fmt.Println(altura)

	//strings
	nome := "Wendel Segadilha"
	fmt.Println(nome)

	//char
	letra := 'W' //tipo "rune" é o mesmo int
	fmt.Println(letra) //87 tabela ASCII	

	//valor zero (null) valor atribuido a variaveis não inicializadas
	var texto string
	var numero int
	fmt.Println(texto)
	fmt.Println(numero)

	//booleano (true, false)
	casado := true
	var solteiro bool
	fmt.Println(casado, solteiro)

	//tipo error
	var erro error = errors.New("Meu erro") //peculiaridade de GO (nil)
	fmt.Println(erro)
	println(erro)
}