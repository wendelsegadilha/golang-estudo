package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Wendel",
		"sobrenome": "Segadilha",
	}

	fmt.Println(usuario["nome"], usuario["sobrenome"])
	
	delete(usuario, "nome") // apagando chave
	
	fmt.Println(usuario["nome"], usuario["sobrenome"])
	
	usuario["salario"] = "3200.45" // adicionando chave

	fmt.Println(usuario["sobrenonome"], usuario["salario"])
	
}