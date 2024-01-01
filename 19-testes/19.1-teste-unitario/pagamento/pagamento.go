package pagamento

import "strings"

func TipoPagamento(tipoPagamento string) string {

	tipoRecebido := strings.Split(tipoPagamento, " ")[1]
	tipos := []string{"boleto", "cartão", "pix"}

	for _, tipo := range tipos {
		if tipoRecebido == tipo {
			return tipo
		}
	}

	return "tipo inválido"
}