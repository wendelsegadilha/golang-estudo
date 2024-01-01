package main

import "fmt"

func main() { 
	// cria o canal mensagens
	mensagens := make(chan string)

	// cria uma goroutine que enviar a mensagem para o canal
	go func() { mensagens <- "ping" }()

	// aguarda a mensagem chegar no canal
	msg := <- mensagens

	//imprime a mensagem recebida
	fmt.Println(msg)
}