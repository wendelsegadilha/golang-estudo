package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

// método adicionado à estrutura de usuário
func (u usuario) salvar() {
	fmt.Printf("Usuário %s salvo com sucesso!\n", u.nome)
}

// quando não queremos alterar algum valor dentro da struct passamos como cópia
func (u usuario) maiorDeIdade() bool  {
	return u.idade >= 18
}

// quando queremos alterar algum valor dentro da struct passamos como ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++ //altera o valor na memoria (reflexão global)
}

func main() {

	u1 := usuario{"Wendel", 29}
	u1.fazerAniversario()
	fmt.Println(u1.maiorDeIdade())
	fmt.Println(u1.idade)
	u1.salvar()

	u2 := usuario{"Venes", 30}
	fmt.Println(u2.maiorDeIdade())
	u2.salvar()

}