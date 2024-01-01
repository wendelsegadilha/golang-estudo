package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Wendel"), escrever("Venes"))

	for msg := range canal {
		fmt.Println(msg)
	}
}

// unifica toda a comunicação em apenas um canal
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string  {
	canalSaida := make(chan string)
	go func ()  {
		for {
			select {
			case msg := <-canalEntrada1:
				canalSaida <-msg
			case msg := <-canalEntrada2:
				canalSaida <-msg
			}
		}
	}()
	return canalSaida
}

func escrever(msg string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", msg)
			time.Sleep(time.Microsecond * 500) //0.5 segundos
		}
	}()
	return canal // canal somente de leitura
}