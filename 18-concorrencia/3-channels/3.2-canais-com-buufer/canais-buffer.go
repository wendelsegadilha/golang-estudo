package main

import "fmt"

func main() {
	// defido buffer de tamanho 2
	canal := make(chan string, 2)

	canal <- "Wendel"
	canal <- "Venes"

	// so bloqueia se exceder o tamanho do buffer
	fmt.Println(<- canal)
	fmt.Println(<- canal)
}