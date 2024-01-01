package main

import (
	"fmt"
	"time"
)

func main() {

	// cria o canal
	canal := make(chan string)

	// cria a goroutine passando a msg e o canal
	go escrever("wendel", canal)

	// recebe o dado do canal e imprime
	// finaliza a execuão assim que o canal for fechado ()executa enquanto o canal estiver aberto)
	for msg := range canal {
		fmt.Println(msg)
	}

	/* podemos capturar o retorno dessa forma também
	for {
		msg, aberto := <- canal // execução bloqueante
		if !aberto {
			break
		}
		fmt.Println(msg)
	}
	*/

}

func escrever(msg string, canal chan string) {
	for i := 0; i < 5; i++ {
		// envia o dado para o canal
		canal <- msg
		time.Sleep(time.Second)
	}
	// fecha o canal
	close(canal)
}