package main

import "fmt"

func f1() {
	fmt.Println("função 1")
}

func f2() {
	fmt.Println("função 2")
}

func main() {
	defer f1() // adia (executa a função por último) a execução da função 1 (o defer executa antes do return)
	f2()
}