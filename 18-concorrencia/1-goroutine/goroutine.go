package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("chamada 1") // crio um goroutine para executar a chamada 1
	escrever("chamada 2")
}

func escrever(msg string) {
	for { // laço infinito
		fmt.Println(msg)
		time.Sleep(time.Second) // atraso de 1 segundo na execução
	}
}
