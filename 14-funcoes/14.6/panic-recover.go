package main

import "fmt"

func recuperarExecução()  {
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso!")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecução() //evita a parada total da aplicação
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("MEIDA = 6") //mata a execução da aplicação (entra em pânico) caso não se use o recover 
}

func main() {
	fmt.Println(alunoAprovado(9, 8))
	fmt.Println(alunoAprovado(5, 4))
	fmt.Println(alunoAprovado(6, 6))
}