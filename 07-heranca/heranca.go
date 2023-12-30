package main

import "fmt"

type pessoa struct {
	nome  string
	email string
}

type aluno struct {
	pessoa
	curso string
	ano   int16
}

func main() {
	p1 := pessoa{"Wendel", "wendel@gmail.com"}
	a1 := aluno{p1, "Engenharia da Computação", 2020}
	fmt.Println(a1.email)
}
