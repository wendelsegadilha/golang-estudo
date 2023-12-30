package main

import "fmt"

func soma(op string, numeros ...int) (string, int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return op, total
}

func main() {

	op, total := soma("+", 1, 2, 3, 5, 10)
	fmt.Println(op, total)

}