package main

import (
	"fmt"
	"reflect"
)

// estruturas (coleção de campos) se parecem com classes em linguagens OO
type Pessoa struct {
	Nome string
	Idade uint8
	Salario float32
	Endereco Endereco
}

type Endereco struct {
	Rua string
	Numero uint
}

func main()  {

	end := Endereco {"Rua A", 100}

	p1 := Pessoa {
		Nome: "Wendel Segadilha",
		//Idade: 29,
		Salario: 3200.98,
	}
	fmt.Println("Tipo: ", reflect.TypeOf(p1)) // verificando tipo de dado	
	fmt.Println(p1.Nome)

	p2 := Pessoa{"Venes Lopes", 30, 2500.65, end}
	fmt.Println(p2.Nome, p2.Endereco.Rua)


	var p3 Pessoa
	p3.Nome = "Anny"
	fmt.Println(p3.Nome)


}

