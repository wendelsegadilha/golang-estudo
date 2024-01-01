package main

import (
	"fmt"
	"testes/pagamento"
)

func main()  {
	
	retorno := pagamento.TipoPagamento("pagamento pix")
	fmt.Println(retorno)
	
}