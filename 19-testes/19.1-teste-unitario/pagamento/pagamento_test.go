package pagamento

import "testing"

type cenario struct {
	tipoRecebido string
	tipoRetornado string
}

func TestTipoPagamento(t *testing.T) {

	cenarios := []cenario{
		{"pagamento boleto", "boleto"},
		{"pagamento cartão", "cartão"},
		{"pagamento pix", "pix"},
		{" ", "tipo inválido"},
	}

	for _, cenario:= range cenarios {
		retorno := TipoPagamento(cenario.tipoRecebido)
		if retorno != cenario.tipoRetornado {
			t.Errorf("O tipo recebido %s é diferente do tpo esperado %s", retorno, cenario.tipoRetornado)
		}
	}

}