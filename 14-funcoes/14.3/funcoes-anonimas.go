package main

import "fmt"

func main() {

	retorno := func(msg string) string {
		return fmt.Sprintf("Recebido: %s", msg)
	}("Wendel Segadilha")

	fmt.Println(retorno)

}