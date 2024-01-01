package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup // declarando o grupo de espera
	waitGroup.Add(2) // informando quantas goroutines estarao rodando

	go func()  {
		escrever("chamada 1")
		waitGroup.Done() // remove a goroutine da lista de execução
	}()

	go func()  {
		escrever("chamada 2")
		waitGroup.Done() // remove a goroutine da lista de execução
	}()

	waitGroup.Wait() // aguarda até que as goroutines terminem suas execuções

}

func escrever(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
