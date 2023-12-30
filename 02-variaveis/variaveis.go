package main

import "fmt"

func main()  {
	var nome string = "Wendel"
	fmt.Println(nome)

	idade := 29
	fmt.Println(idade)

	casado, salario := true, 2500.10
	fmt.Println(casado, salario)

	//inverter variavel (troca)
	x := 10
	y := 20
	fmt.Println(x, y)
	x , y = y, x
	fmt.Println(x, y)

}