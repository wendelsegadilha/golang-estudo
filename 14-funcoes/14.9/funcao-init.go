package main

import "fmt"

func main() {
	fmt.Println("Função main")
}

func init()  { // executa sempre antes da main (serve para inicializar configurações)
	fmt.Println("Função init")
}
